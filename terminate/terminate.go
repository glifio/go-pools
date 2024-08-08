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

var PERC_SECTORS_TO_SAMPLE = big.NewRat(1, 100)
var MAX_SAMPLED_SECTORS = big.NewInt(10000)
var MIN_SAMPLED_SECTORS = big.NewInt(1000)

func EstimateTerminationPenalty(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*big.Int, float64, error) {
	sectors, err := AllSectors(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, 0, err
	}

	// Sample 1% of sectors, with a max of 10000 and a minimum of 1000
	sample := bitfield.NewFromSet(SampleSectors(sectors, PERC_SECTORS_TO_SAMPLE, MAX_SAMPLED_SECTORS, MIN_SAMPLED_SECTORS))

	penalty, avgPenaltyPerPledge, err := TerminateSectors(ctx, api, minerAddr, &sample, ts)
	if err != nil {
		return nil, 0, err
	}

	return penalty, avgPenaltyPerPledge, nil
}

func AllSectors(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) ([]uint64, error) {
	allSectors, err := api.StateMinerAllocated(context.Background(), minerAddr, ts.Key())
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

func TerminateSectors(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, sectors *bitfield.BitField, ts *types.TipSet) (*big.Int, float64, error) {
	_, mstate, err := mstat.LoadMinerActor(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, 0, err
	}

	minerInfo, err := mstate.Info()
	if err != nil {
		return nil, 0, err
	}

	smoothedPow, err := util.TotalPowerSmoothed(ctx, api, ts)
	if err != nil {
		return nil, 0, err
	}

	smoothedRew, err := util.ThisEpochRewardsSmoothed(ctx, api, ts)
	if err != nil {
		return nil, 0, err
	}

	sectorInfos, err := api.StateMinerSectors(context.Background(), minerAddr, sectors, ts.Key())
	if err != nil {
		return nil, 0, err
	}

	penaltyPerPledges := make([]float64, 0)

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
			return nil, 0, xerrors.Errorf("negative termination fee: %v", termFee)
		}

		if s.InitialPledge.Int.Cmp(big.NewInt(0)) == 0 {
			// if the initial pledge is 0, something is up?
			return nil, 0, xerrors.Errorf("0 initial pledge for sector")
		}

		// compute a float64 representation of the termination fee/initial pledge
		termFeeRat := big.NewRat(0, 1).SetFrac(termFee, s.InitialPledge.Int)
		termFeeFloat, _ := termFeeRat.Float64()
		penaltyPerPledges = append(penaltyPerPledges, termFeeFloat)
	}

	// loop through the penalties and average them
	sum := float64(0)
	for _, penalty := range penaltyPerPledges {
		sum += penalty
	}

	avgPenaltyPerPledge := sum / float64(len(penaltyPerPledges))

	// interpolate the penalty to the entire sector set based on initial pledge
	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, 0, err
	}

	initialPledge := new(big.Int).Add(lf.InitialPledgeRequirement.Int, lf.PreCommitDeposits.Int)

	// the total penalty is the average penalty per pledge times the total initial pledge
	totalPenalty := avgPenaltyPerPledge * float64(initialPledge.Int64())

	return big.NewInt(int64(totalPenalty)), avgPenaltyPerPledge, nil
}
