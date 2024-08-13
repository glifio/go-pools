package terminate

import (
	"context"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/mstat"
)

// PreviewTerminateSectorsQuick will preview the cost of terminating all
// the sectors for a miner
//
// It samples a subset of sectors and uses the "off-chain" calculation method.
//
// It has been "tuned" to provide reasonably accurate results (typically <1% error)
// while executing in a fair quick constant amount of time.

// Use the more complex [PreviewTerminateSectors] function for more control over the number of
// sectors to sample, or to use the "on-chain" calculation method, and also
// to stream results.
func PreviewTerminateSectorsQuick(
	ctx context.Context,
	api *lotusapi.FullNodeStruct,
	minerAddr address.Address,
	ts *types.TipSet,
) (*PreviewTerminateSectorsReturn, error) {

	actor, mstate, err := mstat.LoadMinerActor(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	minerInfo, err := mstate.Info()
	if err != nil {
		return nil, err
	}

	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, err
	}

	// Fetch all sectors
	sectors, err := AllSectors(context.Background(), api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	// Sample 1% of the sectors
	var PERC_SECTORS_TO_SAMPLE = big.NewRat(1, 100)
	var MAX_SAMPLED_SECTORS = big.NewInt(1000)
	var MIN_SAMPLED_SECTORS = big.NewInt(300)
	sampled := SampleSectors(sectors, PERC_SECTORS_TO_SAMPLE, MAX_SAMPLED_SECTORS, MIN_SAMPLED_SECTORS)

	// create a bitfield from the sampled sectors
	sample := bitfield.NewFromSet(sampled)

	// Get estimate for terminating the sampled sectors
	res, err := TerminateSectors(context.Background(), api, minerAddr, &sample, ts)
	if err != nil {
		return nil, err
	}

	stats := SectorStats{
		TerminationPenalty: res.EstimatedTotalTerminationFee,
		InitialPledge:      res.InitialPledge,
		SectorCount:        res.SectorCount,
	}

	return &PreviewTerminateSectorsReturn{
		Actor:                      actor,
		MinerInfo:                  minerInfo,
		VestingBalance:             lf.VestingFunds.Int,
		InitialPledge:              new(big.Int).Add(lf.InitialPledgeRequirement.Int, lf.PreCommitDeposits.Int),
		SectorStats:                &stats,
		SectorsTerminated:          uint64(res.SectorCount),
		SectorsCount:               uint64(res.SectorCount),
		SectorsInSkippedPartitions: 0,
		Epoch:                      ts.Height(),
		PartitionsCount:            64,
		SampledPartitionsCount:     64,
		Tipset:                     ts,
	}, nil
}
