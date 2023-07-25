package mstat

import (
	"context"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/util/smoothing"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/actors/builtin/power"
	"github.com/filecoin-project/lotus/chain/actors/builtin/reward"
	"github.com/filecoin-project/lotus/chain/types"
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
	MinerInfo           *api.MinerInfo
	Balance             *big.Int
	PenaltyTermination  *big.Int
	ExpectedDailyReward *big.Int
	PenaltyFaultPerDay  *big.Int
	PledgedFunds        *big.Int
	VestingFunds        *big.Int
	GreenScore          *big.Int
	QualityAdjPower     *big.Int
	LiveSectors         *big.Int
	FaultySectors       *big.Int
	HasMinPower         bool
}

// create a new miner stats struct
func NewMinerStats() *MinerStats {
	return &MinerStats{
		MinerInfo:           &api.MinerInfo{},
		Balance:             big.NewInt(0),
		PenaltyTermination:  big.NewInt(0),
		ExpectedDailyReward: big.NewInt(0),
		PenaltyFaultPerDay:  big.NewInt(0),
		PledgedFunds:        big.NewInt(0),
		VestingFunds:        big.NewInt(0),
		GreenScore:          big.NewInt(0),
		QualityAdjPower:     big.NewInt(0),
		LiveSectors:         big.NewInt(0),
		FaultySectors:       big.NewInt(0),
		HasMinPower:         false,
	}
}

func ComputeMinerStats(ctx context.Context, minerAddr address.Address, ts *types.TipSet, api api.FullNode) (*MinerStats, error) {
	minerStats := &MinerStats{}

	mi, err := api.StateMinerInfo(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}
	minerStats.MinerInfo = &mi

	actor, minerState, err := loadMinerActor(ctx, api, minerAddr, ts)
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
	sectorStatus, err := computeSectorStatus(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}
	minerStats.LiveSectors = new(big.Int).SetUint64(sectorStatus.Live)
	minerStats.FaultySectors = new(big.Int).SetUint64(sectorStatus.Faulty)

	// penalty is 50% of miner's assets
	minerStats.Balance = actor.Balance.Int
	minerStats.PenaltyTermination = computePercentOf(actor.Balance.Int)
	// lazy compute of PenaltyFaultPerDay is 3x the EDR
	edr, err := ComputeEDRLazy1(ctx, minerAddr, ts, api)
	if err != nil {
		return nil, err
	}

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

func ComputeFilPerTiB(ctx context.Context, ts *types.TipSet, api api.FullNode) (*float64, error) {
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

func ComputeEDRLazy1(ctx context.Context, minerAddr address.Address, ts *types.TipSet, api api.FullNode) (*big.Int, error) {
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
		log.Printf("block reward: %v", reward)

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

			weekly, _ := new(big.Rat).Mul(
				winRatio,
				new(big.Rat).SetInt64(7*builtin.EpochsInDay),
			).Float64()

			avgDuration, _ := new(big.Rat).Mul(
				new(big.Rat).SetInt64(builtin.EpochDurationSeconds),
				new(big.Rat).Inv(winRatio),
			).Float64()

			log.Print("Projected average block win rate: ")
			log.Printf(
				"%.02f/week (every %s)",
				weekly,
				(time.Second * time.Duration(avgDuration)).Truncate(time.Second).String(),
			)
			log.Printf(
				"%.02f/day (every %s)",
				daily,
				(time.Second * time.Duration(avgDuration)).Truncate(time.Second).String(),
			)

			// Geometric distribution of P(Y < k) calculated as described in https://en.wikipedia.org/wiki/Geometric_distribution#Probability_Outcomes_Examples
			// https://www.wolframalpha.com/input/?i=t+%3E+0%3B+p+%3E+0%3B+p+%3C+1%3B+c+%3E+0%3B+c+%3C1%3B+1-%281-p%29%5E%28t%29%3Dc%3B+solve+t
			// t == how many dice-rolls (epochs) before win
			// p == winRate == ( minerPower / netPower )
			// c == target probability of win ( 99.9% in this case )
			log.Print("Projected block win with ")
			log.Printf(
				"99.9%% probability every %s",
				(time.Second * time.Duration(
					builtin.EpochDurationSeconds*math.Log(1-0.999)/
						math.Log(1-winRatioFloat),
				)).Truncate(time.Second).String(),
			)
			log.Println("(projections DO NOT account for future network and miner growth)")

			dailyFloat := new(big.Float).SetFloat64(daily)
			log.Printf("daily win ratio: %v", dailyFloat)

			rewardFloat := new(big.Float).SetInt(reward)

			totalRewardsBigFloat := new(big.Float).Mul(dailyFloat, rewardFloat)
			log.Printf("total reward: %v", totalRewardsBigFloat)

			edr, _ := totalRewardsBigFloat.Int(nil)

			return edr, nil
		}
	}

	return big.NewInt(0), nil
}

func getSmoothedRew(ctx context.Context, api api.FullNode, ts *types.TipSet) (smoothing.FilterEstimate, error) {
	ract, err := api.StateGetActor(ctx, reward.Address, ts.Key())
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	tbsRew := newTieredBlockstore(api)

	rst, err := reward.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsRew)), ract)
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	smoothedRew, err := rst.ThisEpochRewardSmoothed()
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	return smoothing.NewEstimate(
		smoothedRew.PositionEstimate,
		smoothedRew.VelocityEstimate,
	), nil
}

func getBlockReward(ctx context.Context, api api.FullNode, ts *types.TipSet) (*big.Int, error) {
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

func getSmoothedPow(ctx context.Context, api api.FullNode, ts *types.TipSet) (smoothing.FilterEstimate, error) {
	pact, err := api.StateGetActor(ctx, power.Address, ts.Key())
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	tbsPow := newTieredBlockstore(api)

	pst, err := power.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsPow)), pact)
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	smoothedPow, err := pst.TotalPowerSmoothed()
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	return smoothing.NewEstimate(
		smoothedPow.PositionEstimate,
		smoothedPow.VelocityEstimate,
	), nil
}

func loadMinerActor(ctx context.Context, api api.FullNode, minerAddr address.Address, ts *types.TipSet) (*types.ActorV5, miner.State, error) {
	mact, err := api.StateGetActor(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, nil, err
	}

	tbsMiner := newTieredBlockstore(api)

	state, err := miner.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsMiner)), mact)

	return mact, state, err
}

func newTieredBlockstore(api api.FullNode) blockstore.Blockstore {
	return blockstore.NewTieredBstore(
		blockstore.NewAPIBlockstore(api),
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

func computeSectorStatus(ctx context.Context, lotus api.FullNode, miner address.Address, tsk *types.TipSet) (*api.MinerSectors, error) {

	sectorStatus := api.MinerSectors{}

	const LOOK_BACK_LIMIT = 3

	// loop back in steps of 1000 tipsets 10 times and avarage the sector count
	for i := 0; i < LOOK_BACK_LIMIT; i++ {
		ts, err := lotus.ChainGetTipSetByHeight(ctx, tsk.Height()-abi.ChainEpoch((i*200)), tsk.Key())
		if err != nil {
			return nil, err
		}
		sectorCount, err := lotus.StateMinerSectorCount(ctx, miner, ts.Key())
		if err != nil {
			return nil, err
		}
		log.Printf("Sector count(%d): %d", ts.Height(), sectorCount)

		// add result to the total and average it
		sectorStatus.Live += sectorCount.Live
		sectorStatus.Active += sectorCount.Active
		sectorStatus.Faulty += sectorCount.Faulty
	}

	sectorStatus.Live = sectorStatus.Live / LOOK_BACK_LIMIT
	sectorStatus.Active = sectorStatus.Active / LOOK_BACK_LIMIT
	sectorStatus.Faulty = sectorStatus.Faulty / LOOK_BACK_LIMIT

	log.Printf("Sector average(%d): %d", tsk.Height(), sectorStatus)

	return &sectorStatus, nil
}
