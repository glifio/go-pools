package terminate

import (
	"context"
	"log"
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

var PERC_SECTORS_TO_SAMPLE = big.NewRat(2, 100)
var MAX_SAMPLED_SECTORS = big.NewInt(1000)
var MIN_SAMPLED_SECTORS = big.NewInt(1)

func EstimateTerminationPenalty(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*big.Int, *big.Int, error) {
	log.Println("Fetching all sectors ")
	sectors, err := AllSectors(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, nil, err
	}
	log.Println("Fetched all sectors ", len(sectors))

	sampled := SampleSectors(sectors, PERC_SECTORS_TO_SAMPLE, MAX_SAMPLED_SECTORS, MIN_SAMPLED_SECTORS)

	log.Println("Sampled sectors length: ", len(sampled))

	// Sample 1% of sectors, with a max of 10000 and a minimum of 1000
	sample := bitfield.NewFromSet(sampled)

	res, err := TerminateSectors(ctx, api, minerAddr, &sample, ts)
	if err != nil {
		return nil, nil, err
	}

	log.Printf("Termination penalty: %0.02f\n", util.ToFIL(res.TerminationFeeFromSample))

	return res.EstimatedTotalTerminationFee, res.AvgTerminationFeePerPledge, nil
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

func SampleSectors(sectors []uint64, perc *big.Rat, max *big.Int, min *big.Int) []uint64 {
	sectorCount := len(sectors)
	if sectorCount == 0 {
		return []uint64{}
	}

	// Compute the percentage of sectors
	percFloat, _ := perc.Float64()
	desiredCount := int(percFloat * float64(sectorCount))

	// Apply the min and max boundaries
	if big.NewInt(int64(desiredCount)).Cmp(min) < 0 {
		desiredCount = int(min.Int64())
	}
	if big.NewInt(int64(desiredCount)).Cmp(max) > 0 {
		desiredCount = int(max.Int64())
	}

	// Compute the step size for evenly spaced sampling
	step := float64(sectorCount-1) / float64(desiredCount-1)

	// Sample the sectors
	sampledSectors := make([]uint64, desiredCount)
	for i := 0; i < desiredCount; i++ {
		index := int(float64(i) * step)
		sampledSectors[i] = sectors[index]
	}

	return sampledSectors
}

func TerminateSectors(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, sectors *bitfield.BitField, ts *types.TipSet) (*TerminateSectorResult, error) {
	// create a new terminate sector result
	result := TerminateSectorResult{
		InitialPledgeFromSample:      big.NewInt(0),
		TerminationFeeFromSample:     big.NewInt(0),
		AvgInitialPledgePerSector:    big.NewInt(0),
		AvgTerminationFeePerPledge:   big.NewInt(0),
		EstimatedTotalInitialPledge:  big.NewInt(0),
		EstimatedTotalTerminationFee: big.NewInt(0),
		InitialPledge:                big.NewInt(0),
		SectorCount:                  0,
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
		result.SectorCount++
	}

	// calculate the average penalty per pledge as result.TerminationFee / result.InitialPledge and keep the reult in big.int
	result.AvgTerminationFeePerPledge = util.DivWad(result.TerminationFeeFromSample, result.InitialPledgeFromSample)
	// log average penalty per pledge
	// log.Printf("AVERAGE PENALTY PER PLEDGE: %f\n", util.ToFIL(result.AvgTerminationFeePerPledge))

	// calculate the average initial pledge per sector as result.InitialPledge / result.SectorCount and keep the result in big.int
	result.AvgInitialPledgePerSector = big.NewInt(0).Div(result.InitialPledgeFromSample, big.NewInt(int64(result.SectorCount)))
	// log initial pledge per sector
	// log.Printf("AVERAGE INITIAL PLEDGE PER SECTOR: %d\n", result.AvgInitialPledgePerSector)

	// interpolate the penalty to the entire sector set based on initial pledge
	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, err
	}

	result.InitialPledge = lf.InitialPledgeRequirement.Int

	// the total penalty is the average penalty per pledge times the total initial pledge
	result.EstimatedTotalTerminationFee = util.MulWad(result.AvgTerminationFeePerPledge, lf.InitialPledgeRequirement.Int)
	// log.Printf("ESTIMATED TOTAL TERMINATION FEE: %f\n", util.ToFIL(result.EstimatedTotalTerminationFee))

	// the total initial pledge is the average initial pledge per sector times the total number of sectors
	mstate.GetAllocatedSectors()
	liveSectors, _ := mstate.NumLiveSectors()
	result.EstimatedTotalInitialPledge = new(big.Int).Mul(result.AvgInitialPledgePerSector, big.NewInt(int64(liveSectors)))
	// log.Printf("ESTIMATED TOTAL INITIAL PLEDGE: %f\n", util.ToFIL(result.EstimatedTotalInitialPledge))

	// log.Printf("AVERAGE PERCENTAGE PENALTY: %0.02f\n", avgPenaltyPerPledge)

	return &result, nil
}
