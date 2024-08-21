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

func EstimateTerminationFeeAgent(
	ctx context.Context,
	agentAddr common.Address,
	withoutMiner address.Address,
	psdk poolstypes.PoolsSDK,
	tsk *types.TipSet,
) (*AgentFi, error) {
	agentFi := &AgentFi{}

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

	agentFi.AvailableBalance = agentAvail

	miners, err := psdk.Query().AgentMiners(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	// remove any miner that is supplised as the withoutMiner
	if withoutMiner != address.Undef {
		for i, miner := range miners {
			if miner == withoutMiner {
				miners = append(miners[:i], miners[i+1:]...)
				break
			}
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

	for _, res := range results {
		termResult := res.(*TerminateSectorResult)
		agentFi.AvailableBalance = big.NewInt(0).Add(agentFi.AvailableBalance, termResult.AvailableBalance)
		agentFi.LockedRewards = big.NewInt(0).Add(agentFi.LockedRewards, termResult.VestingFunds)
		agentFi.InitialPledge = big.NewInt(0).Add(agentFi.InitialPledge, termResult.InitialPledge)
		agentFi.FeeDebt = big.NewInt(0).Add(agentFi.FeeDebt, termResult.FeeDebt)
		agentFi.TerminationFee = big.NewInt(0).Add(agentFi.TerminationFee, termResult.EstimatedTerminationFee)
	}

	return agentFi, nil
}

func EstimateTerminationFeeMiner(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*TerminateSectorResult, error) {
	sectors, err := AllSectors(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	allSectors, err := api.StateMinerSectorCount(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}

	if allSectors.Live != uint64(len(sectors)) {
		return nil, xerrors.Errorf("sector count mismatch: %d != %d", allSectors.Live, len(sectors))
	}

	sampled := SampleSectors(sectors, MAX_SAMPLED_SECTORS)

	sample := bitfield.NewFromSet(sampled)

	res, err := TerminateSectors(ctx, api, minerAddr, &sample, ts)
	if err != nil {
		return nil, err
	}

	res.FaultySectors = allSectors.Faulty
	res.LiveSectors = allSectors.Live

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
	result := TerminateSectorResult{}

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

	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, err
	}

	result.TotalBalance = actor.Balance.Int
	result.AvailableBalance = big.NewInt(0).Sub(result.TotalBalance, big.NewInt(0).Add(result.InitialPledge, result.VestingFunds))
	result.VestingFunds = lf.VestingFunds.Int
	result.InitialPledge = lf.InitialPledgeRequirement.Int
	// the available balance is the total balance minus the initial pledge and vesting funds
	feeDebt, err := mstate.FeeDebt()
	if err != nil {
		return nil, err
	}
	result.FeeDebt = feeDebt.Int

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
			return nil, xerrors.Errorf("negative termination fee: %v", termFee)
		}

		if s.InitialPledge.Int.Cmp(big.NewInt(0)) == 0 {
			// if the initial pledge is 0, something is up?
			return nil, xerrors.Errorf("0 initial pledge for sector")
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
