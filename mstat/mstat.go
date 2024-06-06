package mstat

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/actors/builtin/reward"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
	cbor "github.com/ipfs/go-ipld-cbor"
)

var TerminationPenaltyRatio float64 = 0.5

type SectorStatus struct {
	Live   *big.Int
	Active *big.Int
	Faults *big.Int
}

// TODO: add fault info
type MinerStats struct {
	MinerInfo                *api.MinerInfo
	Balance                  *big.Int
	ExpectedDailyReward      *big.Int
	ExpectedDailyBlockReward *big.Int
	PenaltyFaultPerDay       *big.Int
	PledgedFunds             *big.Int
	VestingFunds             *big.Int
	GreenScore               *big.Int
	QualityAdjPower          *big.Int
	LiveSectors              *big.Int
	FaultySectors            *big.Int
	HasMinPower              bool

	// DEPRECATED
	PenaltyTermination *big.Int
}

// create a new miner stats struct
func NewMinerStats() *MinerStats {
	return &MinerStats{
		MinerInfo:                &api.MinerInfo{},
		Balance:                  big.NewInt(0),
		PenaltyTermination:       big.NewInt(0),
		ExpectedDailyReward:      big.NewInt(0),
		ExpectedDailyBlockReward: big.NewInt(0),
		PenaltyFaultPerDay:       big.NewInt(0),
		PledgedFunds:             big.NewInt(0),
		VestingFunds:             big.NewInt(0),
		GreenScore:               big.NewInt(0),
		QualityAdjPower:          big.NewInt(0),
		LiveSectors:              big.NewInt(0),
		FaultySectors:            big.NewInt(0),
		HasMinPower:              false,
	}
}

func ComputeMinersStats(ctx context.Context, minerAddrs []address.Address, ts *types.TipSet, api *api.FullNodeStruct) (*MinerStats, error) {
	aggMinerStats := NewMinerStats()

	minerCount := int64(len(minerAddrs))

	tasks := make([]util.TaskFunc, minerCount)
	for i := int64(0); i < minerCount; i++ {
		minerAddr := minerAddrs[i]
		tasks[i] = func() (interface{}, error) {
			return ComputeMinerStats(ctx, minerAddr, ts, api)
		}
	}

	results, err := util.Multiread(tasks)
	if err != nil {
		return nil, err
	}

	for _, minerStats := range results {
		minerStats := minerStats.(*MinerStats)

		aggMinerStats.Balance = big.NewInt(0).Add(aggMinerStats.Balance, minerStats.Balance)
		aggMinerStats.PenaltyTermination = big.NewInt(0).Add(aggMinerStats.PenaltyTermination, minerStats.PenaltyTermination)
		aggMinerStats.PenaltyFaultPerDay = big.NewInt(0).Add(aggMinerStats.PenaltyFaultPerDay, minerStats.PenaltyFaultPerDay)
		aggMinerStats.ExpectedDailyReward = big.NewInt(0).Add(aggMinerStats.ExpectedDailyReward, minerStats.ExpectedDailyReward)
		aggMinerStats.ExpectedDailyBlockReward = big.NewInt(0).Add(aggMinerStats.ExpectedDailyBlockReward, minerStats.ExpectedDailyBlockReward)
		aggMinerStats.PledgedFunds = big.NewInt(0).Add(aggMinerStats.PledgedFunds, minerStats.PledgedFunds)
		aggMinerStats.VestingFunds = big.NewInt(0).Add(aggMinerStats.VestingFunds, minerStats.VestingFunds)

		// add sector stats
		aggMinerStats.LiveSectors = big.NewInt(0).Add(aggMinerStats.LiveSectors, minerStats.LiveSectors)
		aggMinerStats.FaultySectors = big.NewInt(0).Add(aggMinerStats.FaultySectors, minerStats.FaultySectors)

		aggMinerStats.GreenScore = big.NewInt(0).Add(aggMinerStats.GreenScore, minerStats.GreenScore)

		// only add power if miner has min power
		if minerStats.HasMinPower {
			aggMinerStats.QualityAdjPower = big.NewInt(0).Add(aggMinerStats.QualityAdjPower, minerStats.QualityAdjPower)
			aggMinerStats.HasMinPower = true
		}
	}

	return aggMinerStats, nil
}

func ComputeMinerStats(ctx context.Context, minerAddr address.Address, ts *types.TipSet, api *api.FullNodeStruct) (*MinerStats, error) {
	minerStats := &MinerStats{}

	mi, err := api.StateMinerInfo(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}
	minerStats.MinerInfo = &mi

	actor, minerState, err := LoadMinerActor(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	lf, err := minerState.LockedFunds()
	if err != nil {
		return nil, err
	}

	minerStats.PledgedFunds = new(big.Int).Add(lf.InitialPledgeRequirement.Int, lf.PreCommitDeposits.Int)
	minerStats.VestingFunds = lf.VestingFunds.Int

	// get miner power
	power, err := api.StateMinerPower(ctx, minerAddr, ts.Key())
	if err != nil {
		log.Printf("error getting miner power: %v", err)
		return nil, err
	}
	minerStats.QualityAdjPower = power.MinerPower.QualityAdjPower.Int
	minerStats.HasMinPower = power.HasMinPower

	// get sector status
	sectorStatus, err := api.StateMinerSectorCount(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}
	minerStats.LiveSectors = new(big.Int).SetUint64(sectorStatus.Live)
	minerStats.FaultySectors = new(big.Int).SetUint64(sectorStatus.Faulty)

	// penalty is 50% of miner's assets
	minerStats.Balance = actor.Balance.Int
	// TODO: Deprecate
	minerStats.PenaltyTermination = computePercentOf(actor.Balance.Int)
	// lazy compute of PenaltyFaultPerDay is 3x the EDR
	expectedDailyBlockRewards, err := ComputeBlockRewardsLazy(ctx, minerAddr, ts, api)
	if err != nil {
		return nil, err
	}
	// set the expected daily block reward
	minerStats.ExpectedDailyBlockReward = expectedDailyBlockRewards

	edr := expectedDailyBlockRewards

	// add optimistic earnings to the edr based on non pledged FIL
	availBal := big.NewInt(0).Sub(actor.Balance.Int, minerStats.PledgedFunds)
	availBal.Sub(availBal, minerStats.VestingFunds)

	// if we have pledged assets, we can optimistically add edr from available balance, assuming it will be pledged
	edr.Add(edr, ComputeOptimisticEarnings(availBal, minerStats.PledgedFunds, expectedDailyBlockRewards))

	// calculate estimated daily vesting rewards as 1/180 of vesting funds
	edvr := big.NewInt(0)
	if minerStats.VestingFunds != nil && minerStats.VestingFunds.Cmp(big.NewInt(0)) == 1 {
		edvr = big.NewInt(0).Div(minerStats.VestingFunds, big.NewInt(180))
	}

	// expected daily reward is the sum of edr and edvr
	minerStats.ExpectedDailyReward = big.NewInt(0).Add(edr, edvr)
	minerStats.PenaltyFaultPerDay = new(big.Int).Mul(edr, big.NewInt(3))

	minerStats.GreenScore = big.NewInt(0)

	return minerStats, nil
}

func ComputeFilPerTiB(ctx context.Context, ts *types.TipSet, api *api.FullNodeStruct) (*float64, error) {
	// get the block rewards for this epoch
	reward, err := getBlockReward(ctx, api, ts)
	if err != nil {
		return nil, err
	}

	// get the smoothed power for this epoch
	power, err := api.StateMinerPower(ctx, address.Undef, ts.Key())
	if err != nil {
		return nil, err
	}

	// calculate the ratio of reward to power
	ratio, _ := new(big.Rat).SetFrac(
		big.NewInt(0).Mul(reward, big.NewInt(0).SetUint64(build.BlocksPerEpoch)),
		big.NewInt(0).Div(power.TotalPower.QualityAdjPower.Int, big.NewInt(1099511627776)),
	).Float64()

	// multiply by 2880 to get the daily reward
	ratio = ratio * 2880

	return &ratio, nil
}

// This function takes an amount of FIL, an amount of pledged FIL, and an amount of block reward earnings, and computes a theoretical earnings for the available FIL if it were pledged
func ComputeOptimisticEarnings(available *big.Int, pledged *big.Int, expectedBlockRewards *big.Int) *big.Int {
	// if we dont have any pledge funds, we can't compute optimistic earnings
	if pledged.Cmp(big.NewInt(0)) == 1 {
		// compute ratio of earnings to initial pledge
		rewardsPledgeRatio := new(big.Rat).SetFrac(expectedBlockRewards, pledged)
		// multiply by available balance to get the expected daily reward
		optimisticEdr := new(big.Int).Mul(available, rewardsPledgeRatio.Num())
		optimisticEdr.Div(optimisticEdr, rewardsPledgeRatio.Denom())
		// add the optimistic edr to the edr
		return optimisticEdr
	}
	return big.NewInt(0)
}

// Deprecated: ComputeEDRLazy1 is deprecated.
func ComputeEDRLazy1(ctx context.Context, minerAddr address.Address, ts *types.TipSet, api *api.FullNodeStruct) (*big.Int, error) {
	return ComputeBlockRewardsLazy(ctx, minerAddr, ts, api)
}

func ComputeBlockRewardsLazy(ctx context.Context, minerAddr address.Address, ts *types.TipSet, api *api.FullNodeStruct) (*big.Int, error) {
	pow, err := api.StateMinerPower(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}

	if pow.HasMinPower {
		// get the block rewards for this epoch
		reward, err := getBlockReward(ctx, api, ts)
		if err != nil {
			return nil, err
		}

		winRatio := new(big.Rat).SetFrac(
			types.BigMul(pow.MinerPower.QualityAdjPower, types.NewInt(build.BlocksPerEpoch)).Int,
			pow.TotalPower.QualityAdjPower.Int,
		)

		if winRatioFloat, _ := winRatio.Float64(); winRatioFloat > 0 {

			// if the corresponding poisson distribution isn't infinitely small then
			// throw it into the mix as well, accounting for multi-wins
			winRationWithPoissonFloat := -math.Expm1(-winRatioFloat)
			winRationWithPoisson := new(big.Rat).SetFloat64(winRationWithPoissonFloat)
			if winRationWithPoisson != nil {
				winRatio = winRationWithPoisson
			}

			daily, _ := new(big.Rat).Mul(
				winRatio,
				new(big.Rat).SetInt64(builtin.EpochsInDay),
			).Float64()

			dailyFloat := new(big.Float).SetFloat64(daily)

			rewardFloat := new(big.Float).SetInt(reward)

			totalRewardsBigFloat := new(big.Float).Mul(dailyFloat, rewardFloat)

			edr, _ := totalRewardsBigFloat.Int(nil)

			return edr, nil
		}
	}

	return big.NewInt(0), nil
}

func getBlockReward(ctx context.Context, api *api.FullNodeStruct, ts *types.TipSet) (*big.Int, error) {
	ract, err := api.StateGetActor(ctx, reward.Address, ts.Key())
	if err != nil {
		return big.NewInt(0), err
	}

	tbsRew := newTieredBlockstore(api)

	rst, err := reward.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsRew)), ract)
	if err != nil {
		return big.NewInt(0), err
	}

	epochReward, err := rst.ThisEpochReward()
	if err != nil {
		return big.NewInt(0), err
	}

	// Divide by expected leaders per epoch to get block reward
	blockReward := types.BigDiv(epochReward, types.NewInt(uint64(builtin.ExpectedLeadersPerEpoch)))

	return blockReward.Int, nil
}

func LoadMinerActor(ctx context.Context, api *api.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*types.ActorV5, miner.State, error) {
	mact, err := api.StateGetActor(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, nil, err
	}

	tbsMiner := newTieredBlockstore(api)

	state, err := miner.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsMiner)), mact)

	return mact, state, err
}

func newTieredBlockstore(lapi *api.FullNodeStruct) blockstore.Blockstore {
	return blockstore.NewTieredBstore(
		blockstore.NewAPIBlockstore(lapi),
		blockstore.NewMemory(),
	)
}

func computePercentOf(x *big.Int) *big.Int {
	// convert the percentage to a big.Rat
	r := new(big.Rat).SetFloat64(TerminationPenaltyRatio)

	// take product of x * percent
	result := new(big.Int).Mul(x, r.Num())
	return result.Div(result, r.Denom())
}
