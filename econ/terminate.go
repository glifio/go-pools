package econ

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"

	poolstypes "github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/util"
	"golang.org/x/xerrors"
)

var MAX_SAMPLED_SECTORS = 1000
var INITIAL_PLEDGE_INTERPOLATION_REL_DIFF = big.NewFloat(3.00) // 3%

func EstimateTerminationFeeAgent(
	ctx context.Context,
	agentAddr common.Address,
	withoutMiner address.Address,
	psdk poolstypes.PoolsSDK,
	tsk *types.TipSet,
) (*AgentFi, error) {
	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	height := big.NewInt(int64(tsk.Height()) - 1)
	agentAvail, err := psdk.Query().AgentLiquidAssets(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	miners, err := psdk.Query().AgentMiners(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	// remove any miner that is supplied as the withoutMiner
	if withoutMiner != address.Undef {
		minerCount := len(miners)
		for i, miner := range miners {
			if miner == withoutMiner {
				miners = append(miners[:i], miners[i+1:]...)
				break
			}
		}
		// check that we actually removed a miner
		if minerCount == len(miners) {
			return nil, xerrors.Errorf("miner %s not found in agent's miners", withoutMiner)
		}
	}

	minerCount := int64(len(miners))

	tasks := make([]util.TaskFunc, minerCount)
	for i := int64(0); i < minerCount; i++ {
		minerAddr := miners[i]
		tasks[i] = func() (interface{}, error) {
			return EstimateTerminationFeeMiner(ctx, lapi, minerAddr, tsk)
		}
	}

	results, err := util.Multiread(tasks)
	if err != nil {
		return nil, err
	}

	baseFis := make([]*BaseFi, minerCount)

	for i, res := range results {
		baseFis[i] = res.(*TerminateSectorResult).ToBaseFi()
	}

	principal, err := psdk.Query().AgentPrincipal(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	interest, err := psdk.Query().AgentInterestOwed(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	l := Liability{
		Principal: principal,
		Interest:  interest,
	}

	return NewAgentFi(agentAvail, l, baseFis), nil
}

func EstimateTerminationFeeMiner(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*TerminateSectorResult, error) {
	sectors, err := AllSectors(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	sampled := SampleSectors(sectors, MAX_SAMPLED_SECTORS)

	sample := bitfield.NewFromSet(sampled)

	res, err := TerminateSectors(ctx, api, minerAddr, &sample, ts)
	if err != nil {
		return nil, err
	}

	if err := failOnBrokenAssumptions(minerAddr, res); err != nil {
		return nil, err
	}

	return res, nil
}

func AllSectors(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) ([]uint64, error) {
	allLiveSectors := make([]bitfield.BitField, 0)

	for dlIdx := uint64(0); dlIdx < 48; dlIdx++ {
		// 48 calls to the internet
		partitions, err := api.StateMinerPartitions(ctx, minerAddr, uint64(dlIdx), ts.Key())
		// if this error occurs, the miner hasn't been around long enough to process the deadlineTs, so we just skip it
		if err != nil {
			continue
		}

		for _, partition := range partitions {
			allLiveSectors = append(allLiveSectors, partition.LiveSectors)
		}
	}

	// one bitfield with all the sectors
	allSectors, err := bitfield.MultiMerge(allLiveSectors...)
	if err != nil {
		return nil, err
	}
	// sectors => all sectors as an array of uint64 sector IDs
	return allSectors.All(math.MaxUint64)
}

func SampleSectors(sectors []uint64, max int) []uint64 {
	sectorCount := len(sectors)

	// If there are no sectors, return an empty slice
	if sectorCount == 0 {
		return []uint64{}
	}

	// Initialize the number of samples to the maximum number of samples
	samples := max

	// If there are fewer sectors than the maximum number of samples, sample all sectors
	if sectorCount < max {
		samples = sectorCount
	}

	// Compute the step size for evenly spaced sampling
	step := float64(sectorCount-1) / float64(samples-1)

	// Sample the sectors
	sampledSectors := make([]uint64, samples)
	for i := 0; i < samples; i++ {
		index := int(float64(i) * step)
		sampledSectors[i] = sectors[index]
	}

	return sampledSectors
}

func TerminateSectors(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, sectors *bitfield.BitField, ts *types.TipSet) (*TerminateSectorResult, error) {
	result := TerminateSectorResult{
		TotalBalance:     big.NewInt(0),
		AvailableBalance: big.NewInt(0),
		VestingFunds:     big.NewInt(0),
		InitialPledge:    big.NewInt(0),
		FeeDebt:          big.NewInt(0),

		EstimatedInitialPledge:    big.NewInt(0),
		InitialPledgeFromSample:   big.NewInt(0),
		AvgInitialPledgePerSector: big.NewInt(0),

		EstimatedTerminationFee:    big.NewInt(0),
		TerminationFeeFromSample:   big.NewInt(0),
		AvgTerminationFeePerPledge: big.NewInt(0),

		SampledSectors: 0,
		LiveSectors:    0,
		FaultySectors:  0,
	}

	actor, mstate, err := util.LoadMinerActor(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	minerInfo, err := mstate.Info()
	if err != nil {
		return nil, err
	}

	smoothedPow, err := util.TotalPowerSmoothed(ctx, api, ts)
	if err != nil {
		return nil, err
	}

	smoothedRew, err := util.ThisEpochRewardsSmoothed(ctx, api, ts)
	if err != nil {
		return nil, err
	}

	sectorInfos, err := api.StateMinerSectors(context.Background(), minerAddr, sectors, ts.Key())
	if err != nil {
		return nil, err
	}

	sectorCount, err := api.StateMinerSectorCount(context.Background(), minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}
	result.LiveSectors = sectorCount.Live
	result.FaultySectors = sectorCount.Faulty

	feeDebt, err := mstate.FeeDebt()
	if err != nil {
		return nil, err
	}
	result.FeeDebt = feeDebt.Int

	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, err
	}
	result.TotalBalance = actor.Balance.Int
	result.VestingFunds = lf.VestingFunds.Int
	result.InitialPledge = lf.InitialPledgeRequirement.Int

	avail, err := mstate.AvailableBalance(actor.Balance)
	if err != nil {
		return nil, err
	}

	// sanity check that the available is a positive number
	if feeDebt.Int.Sign() > 0 && avail.Int.Sign() > 0 {
		// this should never happen
		return nil, xerrors.Errorf("miner: %s: fee debt and available balance are both positive", minerAddr)
	} else if feeDebt.Int.Sign() > 0 {
		result.AvailableBalance = big.NewInt(0)
	} else {
		// otherwise no fee debt, go on normally
		result.AvailableBalance = avail.Int
	}

	// if no sectors found return empty result
	if len(sectorInfos) == 0 {
		return &result, nil
	}

	for _, s := range sectorInfos {
		sectorPower := miner8.QAPowerForSector(minerInfo.SectorSize, util.ConvertSectorType(s))

		// the termination penalty calculation
		termFee := miner8.PledgePenaltyForTermination(
			s.ExpectedDayReward,
			ts.Height()-s.PowerBaseEpoch,
			s.ExpectedStoragePledge,
			smoothedPow,
			sectorPower,
			smoothedRew,
			s.ReplacedDayReward,
			s.PowerBaseEpoch-s.Activation,
		).Int

		// this should never happen, but termination fee should never be negative
		if termFee.Cmp(big.NewInt(0)) < 0 {
			return nil, xerrors.Errorf("miner %s: negative termination fee: %v", minerAddr, termFee)
		}

		if s.InitialPledge.Int.Cmp(big.NewInt(0)) == 0 {
			// if the initial pledge is 0, something is up?
			return nil, xerrors.Errorf("miner %s: 0 initial pledge for sector", minerAddr)
		}

		result.TerminationFeeFromSample = big.NewInt(0).Add(result.TerminationFeeFromSample, termFee)
		result.InitialPledgeFromSample = big.NewInt(0).Add(result.InitialPledgeFromSample, s.InitialPledge.Int)
		result.SampledSectors++
	}

	// calculate the average penalty per pledge as result.TerminationFee / result.InitialPledge and keep the reult in big.int
	result.AvgTerminationFeePerPledge = util.DivWad(result.TerminationFeeFromSample, result.InitialPledgeFromSample)

	// calculate the average initial pledge per sector as result.InitialPledge / result.SectorCount and keep the result in big.int
	result.AvgInitialPledgePerSector = big.NewInt(0).Div(result.InitialPledgeFromSample, big.NewInt(int64(result.SampledSectors)))

	// the total penalty is the average penalty per pledge times the total initial pledge
	result.EstimatedTerminationFee = util.MulWad(result.AvgTerminationFeePerPledge, lf.InitialPledgeRequirement.Int)

	// the total initial pledge is the average initial pledge per sector times the total number of sectors
	result.EstimatedInitialPledge = new(big.Int).Mul(result.AvgInitialPledgePerSector, big.NewInt(int64(result.LiveSectors)))

	return &result, nil
}

// Assumptions:
// 1. there cannot be negative fee debt
// 2. there cannot be positive fee debt with a positive balance
// 3. there cannot be positive fee debt with a positive available balance
// 4. there cannot be a positive fee debt with a positive initial pledge
// 5. there cannot be a positive fee debt with a positive vesting funds
// 6. there cannot be a negative termination fee
// 7. there cannot be a negative initial pledge
// 8. with no fee debt, available balance should always be positive (or 0)
// 9. with no fee debt, initial pledge should always be positive (or 0)
// 10. with no fee debt, balance should always be positive (or 0)
// 11. If estimated initial pledge is off from actual initial pledge, the sampling result is too imprecise
func failOnBrokenAssumptions(minerAddr address.Address, result *TerminateSectorResult) error {
	// 6. there cannot be a negative termination fee
	if result.EstimatedTerminationFee.Cmp(big.NewInt(0)) < 0 {
		return xerrors.Errorf("miner %s: negative termination fee: %v", minerAddr, result.EstimatedTerminationFee)
	}
	// 7. there cannot be a negative initial pledge
	if result.InitialPledge.Cmp(big.NewInt(0)) < 0 {
		return xerrors.Errorf("miner %s: negative initial pledge: %v", minerAddr, result.InitialPledge)
	}

	// fee debt exists
	if result.FeeDebt.Cmp(big.NewInt(0)) > 0 {
		// 2.there cannot be positive fee debt with a positive balance
		if result.TotalBalance.Cmp(big.NewInt(0)) > 0 {
			return xerrors.Errorf("miner %s: positive fee debt with positive balance: %v", minerAddr, result.FeeDebt)
		}
		// 3. there cannot be positive fee debt with a positive available balance
		if result.AvailableBalance.Cmp(big.NewInt(0)) > 0 {
			return xerrors.Errorf("miner %s: positive fee debt with positive available balance: %v", minerAddr, result.FeeDebt)
		}
		// 4. there cannot be a positive fee debt with a positive initial pledge
		if result.InitialPledge.Cmp(big.NewInt(0)) > 0 {
			return xerrors.Errorf("miner %s: positive fee debt with positive initial pledge: %v", minerAddr, result.FeeDebt)
		}
		// 5. there cannot be a positive fee debt with a positive vesting funds
		if result.VestingFunds.Cmp(big.NewInt(0)) > 0 {
			return xerrors.Errorf("miner %s: positive fee debt with positive vesting funds: %v", minerAddr, result.FeeDebt)
		}
	} else if result.FeeDebt.Cmp(big.NewInt(0)) == 0 {
		// no fee debt
		// 8. with no fee debt, available balance should always be positive (or 0)
		if result.AvailableBalance.Cmp(big.NewInt(0)) < 0 {
			return xerrors.Errorf("miner %s: no fee debt, but negative available balance: %v", minerAddr, result.AvailableBalance)
		}
		// 9. with no fee debt, initial pledge should always be positive (or 0)
		if result.InitialPledge.Cmp(big.NewInt(0)) < 0 {
			return xerrors.Errorf("miner %s: no fee debt, but negative initial pledge: %v", minerAddr, result.InitialPledge)
		}
		// 10. with no fee debt, balance should always be positive (or 0)
		if result.TotalBalance.Cmp(big.NewInt(0)) < 0 {
			return xerrors.Errorf("miner %s: no fee debt, but negative balance: %v", minerAddr, result.TotalBalance)
		}
	} else {
		// 1. fee debt is negative - error
		return xerrors.Errorf("miner %s: negative fee debt: %v", minerAddr, result.FeeDebt)
	}

	// 11. If estimated initial pledge is off from actual initial pledge, the sampling result is too imprecise
	initialPledgeDiff := util.Diff(result.InitialPledge, result.EstimatedInitialPledge)

	// we're out of the acceptance range
	if initialPledgeDiff.Cmp(INITIAL_PLEDGE_INTERPOLATION_REL_DIFF) > 0 {
		return xerrors.Errorf("miner %s: initial pledge diff too high: %v, %v, %v", minerAddr, result.InitialPledge, result.EstimatedInitialPledge, initialPledgeDiff)
	}

	return nil
}
