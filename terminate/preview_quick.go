package terminate

import (
	"context"
	"math"
	"math/big"

	"github.com/filecoin-project/go-state-types/builtin/v13/miner"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/mstat"
	"github.com/glifio/go-pools/util"
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
	var batchSize uint64 = 40
	var desiredMinSamples uint64 = 600
	var gasLimit uint64 = 270000000000
	var maxPartitions uint64 = 21

	h := ts.Height()

	// Lookup actor balance
	actor, mstate, err := mstat.LoadMinerActor(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	actor.DelegatedAddress = &minerAddr

	// Lookup current owner / worker
	minerInfo, err := mstate.Info()
	if err != nil {
		return nil, err
	}

	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, err
	}

	provingDeadline, err := api.StateMinerProvingDeadline(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}

	prevHeight := h - 120

	tsPrev, err := api.ChainGetTipSetByHeight(ctx, prevHeight, types.EmptyTSK)
	if err != nil {
		return nil, err
	}

	autoBatchSize := false
	if batchSize == 0 && gasLimit == 0 {
		autoBatchSize = true
		batchSize = 2500
		gasLimit = 90000000000 * 3 // 3 deadlines per batch

		workerActorPrev, err := api.StateGetActor(ctx, minerInfo.Worker, tsPrev.Key())
		if err != nil {
			return nil, err
		}

		workerBal, _ := util.ToFIL(workerActorPrev.Balance.Int).Float64()
		if workerBal < 3.0 {
			ratio := workerBal / 3.0
			batchSize = uint64(float64(batchSize) * ratio)
			gasLimit = uint64(float64(gasLimit) * ratio)
		}
	}

	if autoBatchSize && batchSize < 10 {
		return nil, err
	}

	type deadlinePartition struct {
		dlIdx          uint64
		dlImmutable    bool
		deadlineTs     *types.TipSet
		deadlineHeight abi.ChainEpoch
		partIdx        int
		partition      lotusapi.Partition
	}

	var sectorsTerminated uint64
	var sectorsCount uint64
	var sectorsInSkippedPartitions uint64
	var dlIdx uint64

	deadlinePartitions := make([]deadlinePartition, 0)

	allLiveSectors := make([]bitfield.BitField, 0)

	for dlIdx = 0; dlIdx < 48; dlIdx++ {
		dlImmutable := false
		deadlineTs := ts
		deadlineHeight := h
		if dlIdx == provingDeadline.Index || dlIdx == (provingDeadline.Index+1)%48 {
			dlImmutable = true
			deadlineTs = tsPrev
			deadlineHeight = prevHeight
		}
		partitions, err := api.StateMinerPartitions(ctx, minerAddr, uint64(dlIdx), deadlineTs.Key())
		// if this error occurs, the miner hasn't been around long enough to process the deadlineTs, so we just skip it
		if err != nil {
			continue
		}
		for partIdx, partition := range partitions {
			allLiveSectors = append(allLiveSectors, partition.LiveSectors)
			sc, err := partition.LiveSectors.Count()
			if err != nil {
				return nil, err
			}
			if sc > 0 {
				deadlinePartitions = append(deadlinePartitions, deadlinePartition{
					dlIdx:          dlIdx,
					dlImmutable:    dlImmutable,
					deadlineTs:     deadlineTs,
					deadlineHeight: deadlineHeight,
					partIdx:        partIdx,
					partition:      partition,
				})
			}
		}
	}

	sampledDeadlinePartitions := make([]bool, len(deadlinePartitions))
	maxSectorsToSample := max(maxPartitions, 3)
	allSectors, err := bitfield.MultiMerge(allLiveSectors...)
	if err != nil {
		return nil, err
	}
	sectors, err := allSectors.All(math.MaxUint64)
	if err != nil {
		return nil, err
	}
	sampledSectors := make([]uint64, 0)
	step := max(float64(len(sectors))/float64(maxSectorsToSample-1), 1.0)
	lastIndex := -1
	for sampleIndex := 0.0; int(sampleIndex) < len(sectors); sampleIndex += step {
		i := int(sampleIndex)
		if i > lastIndex {
			sampledSectors = append(sampledSectors, sectors[i])
			lastIndex = i
		}
	}
	if lastIndex < len(sectors)-1 {
		sampledSectors = append(sampledSectors, sectors[len(sectors)-1])
	}
	for _, sectorNumber := range sampledSectors {
		sectorPartition, err := api.StateSectorPartition(ctx, minerAddr, abi.SectorNumber(sectorNumber), ts.Key())
		if err != nil {
			return nil, err
		}
		for i, deadlinePartition := range deadlinePartitions {
			if deadlinePartition.dlIdx == sectorPartition.Deadline &&
				deadlinePartition.partIdx == int(sectorPartition.Partition) {
				sampledDeadlinePartitions[i] = true
			}
		}
	}

	sampledDeadlinePartitionCount := uint64(0)
	for i := range deadlinePartitions {
		if sampledDeadlinePartitions[i] &&
			(maxPartitions == 0 || sampledDeadlinePartitionCount < maxPartitions) {
			sampledDeadlinePartitionCount++
		}
	}

	terminationsCh := make(chan terminationTask)
	statsCh := make(chan *SectorStats)
	errorCh := make(chan error)

	offchain := true
	go runTerminationsInBatches(ctx, api, minerAddr, minerInfo,
		int64(gasLimit), offchain, terminationsCh, statsCh, errorCh)

	go func() {
		for deadlinePartitionIdx, deadlinePartition := range deadlinePartitions {
			dlIdx := deadlinePartition.dlIdx
			deadlineTs := deadlinePartition.deadlineTs
			deadlineHeight := deadlinePartition.deadlineHeight
			partIdx := deadlinePartition.partIdx
			partition := deadlinePartition.partition

			sc, err := partition.LiveSectors.Count()
			if err != nil {
				errorCh <- err
				return
			}
			if sc > 0 {
				if sampledDeadlinePartitions[deadlinePartitionIdx] {
					sectors, err := partition.LiveSectors.All(math.MaxUint64)
					if err != nil {
						errorCh <- err
						return
					}
					sampledSectors := make([]uint64, 0)
					partitionBatchSize := max(batchSize, desiredMinSamples/uint64(len(deadlinePartitions)))
					step := max(float64(len(sectors))/float64(partitionBatchSize-1), 1.0)
					lastIndex := -1
					for sampleIndex := 0.0; int(sampleIndex) < len(sectors); sampleIndex += step {
						i := int(sampleIndex)
						if i > lastIndex {
							sampledSectors = append(sampledSectors, sectors[i])
							lastIndex = i
						}
					}
					if lastIndex < len(sectors)-1 {
						sampledSectors = append(sampledSectors, sectors[len(sectors)-1])
					}
					sampledSectorsBitfield := bitfield.NewFromSet(sampledSectors)
					termination := miner.TerminationDeclaration{
						Deadline:  uint64(dlIdx),
						Partition: uint64(partIdx),
						Sectors:   sampledSectorsBitfield,
					}
					terminationsCh <- terminationTask{
						termination:         termination,
						deadlineHeight:      deadlineHeight,
						deadlineTs:          deadlineTs,
						sectorsCount:        int64(sc),
						sampledSectorsCount: int64(len(sampledSectors)),
					}
					sectorsTerminated += uint64(len(sampledSectors))
					sectorsCount += uint64(len(sectors))
				} else {
					// DeadlinePartition not sampled
					sectorsInSkippedPartitions += sc
					sectorsCount += sc
				}
			}
		}
		close(terminationsCh)
	}()

	allStats := NewSectorStats()

loop:
	for {
		select {
		case stats := <-statsCh:
			if stats == nil {
				break loop
			}
			allStats = allStats.Accumulate(stats)
		case err := <-errorCh:
			return nil, err
		}
	}

	if sectorsCount > 0 {
		allStats = allStats.ScaleUp(int64(sectorsCount),
			int64(sectorsCount-sectorsInSkippedPartitions))
	}

	return &PreviewTerminateSectorsReturn{
		Actor:                      actor,
		MinerInfo:                  minerInfo,
		VestingBalance:             lf.VestingFunds.Int,
		InitialPledge:              new(big.Int).Add(lf.InitialPledgeRequirement.Int, lf.PreCommitDeposits.Int),
		SectorStats:                allStats,
		SectorsTerminated:          sectorsTerminated,
		SectorsCount:               sectorsCount,
		SectorsInSkippedPartitions: sectorsInSkippedPartitions,
		Epoch:                      h,
		PartitionsCount:            uint64(len(deadlinePartitions)),
		SampledPartitionsCount:     sampledDeadlinePartitionCount,
		Tipset:                     ts,
	}, nil
}
