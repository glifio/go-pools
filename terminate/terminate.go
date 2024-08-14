package terminate

import (
	"context"
	"math"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"

	"github.com/glifio/go-pools/mstat"
	"github.com/glifio/go-pools/util"
	"golang.org/x/xerrors"
)

var MAX_SAMPLED_SECTORS = 1000

func EstimateTerminationPenalty(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*TerminateSectorResult, error) {
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
	// create a new terminate sector result
	result := TerminateSectorResult{
		InitialPledgeFromSample:    big.NewInt(0),
		TerminationFeeFromSample:   big.NewInt(0),
		AvgInitialPledgePerSector:  big.NewInt(0),
		AvgTerminationFeePerPledge: big.NewInt(0),
		EstimatedInitialPledge:     big.NewInt(0),
		EstimatedTerminationFee:    big.NewInt(0),
		InitialPledge:              big.NewInt(0),
	}

	_, mstate, err := mstat.LoadMinerActor(ctx, api, minerAddr, ts)
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

	// set the total number of live sectors
	result.LiveSectors, err = mstate.NumLiveSectors()
	if err != nil {
		return nil, err
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

	// interpolate the penalty to the entire sector set based on initial pledge
	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, err
	}

	result.InitialPledge = lf.InitialPledgeRequirement.Int

	// the total penalty is the average penalty per pledge times the total initial pledge
	result.EstimatedTerminationFee = util.MulWad(result.AvgTerminationFeePerPledge, lf.InitialPledgeRequirement.Int)

	// the total initial pledge is the average initial pledge per sector times the total number of sectors
	result.EstimatedInitialPledge = new(big.Int).Mul(result.AvgInitialPledgePerSector, big.NewInt(int64(result.LiveSectors)))

	return &result, nil
}
