package terminate

import (
	"context"
	"math"

	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
)

// PreviewTerminateSectorsQuick gets the burnt funds for when all the sectors are terminated using sampling and offchain calculations
func PreviewTerminateSectorsQuick(
	ctx context.Context,
	api lotusapi.FullNodeStruct,
	minerAddr address.Address,
	tipset string,
	vmHeight uint64,
	errorCh chan error,
	resultCh chan *PreviewTerminateSectorsReturn,
) {
	var batchSize uint64 = 40
	var gasLimit uint64 = 270000000000
	var maxPartitions uint64 = 21

	h, ts, err := parseTipSetAndHeight(ctx, api, tipset, vmHeight)
	if err != nil {
		errorCh <- err
		return
	}

	// Lookup actor balance
	actor, err := api.StateGetActor(ctx, minerAddr, ts.Key())
	if err != nil {
		errorCh <- err
		return
	}

	// Lookup current owner / worker
	minerInfo, err := api.StateMinerInfo(ctx, minerAddr, ts.Key())
	if err != nil {
		errorCh <- err
		return
	}

	provingDeadline, err := api.StateMinerProvingDeadline(ctx, minerAddr, ts.Key())
	if err != nil {
		errorCh <- err
		return
	}

	prevHeight := h - 120

	tsPrev, err := api.ChainGetTipSetByHeight(ctx, prevHeight, types.EmptyTSK)
	if err != nil {
		errorCh <- err
		return
	}

	workerActorPrev, err := api.StateGetActor(ctx, minerInfo.Worker, tsPrev.Key())
	if err != nil {
		errorCh <- err
		return
	}

	autoBatchSize := false
	if batchSize == 0 && gasLimit == 0 {
		autoBatchSize = true
		batchSize = 2500
		gasLimit = 90000000000 * 3 // 3 deadlines per batch

		workerBal, _ := util.ToFIL(workerActorPrev.Balance.Int).Float64()
		if workerBal < 3.0 {
			ratio := workerBal / 3.0
			batchSize = uint64(float64(batchSize) * ratio)
			gasLimit = uint64(float64(gasLimit) * ratio)
		}
	}

	if autoBatchSize && batchSize < 10 {
		errorCh <- xerrors.Errorf("Not enough worker funds! Batch size too small!")
		return
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
		if err != nil {
			errorCh <- err
			return
		}
		for partIdx, partition := range partitions {
			allLiveSectors = append(allLiveSectors, partition.LiveSectors)
			sc, err := partition.LiveSectors.Count()
			if err != nil {
				errorCh <- err
				return
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
		errorCh <- err
		return
	}
	sectors, err := allSectors.All(math.MaxUint64)
	if err != nil {
		errorCh <- err
		return
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
			errorCh <- err
			return
		}
		for i, deadlinePartition := range deadlinePartitions {
			if deadlinePartition.dlIdx == sectorPartition.Deadline &&
				deadlinePartition.partIdx == int(sectorPartition.Partition) {
				sampledDeadlinePartitions[i] = true
			}
		}
	}

	sampledDeadlinePartitionCount := uint64(0)
	for i, _ := range deadlinePartitions {
		if sampledDeadlinePartitions[i] &&
			(maxPartitions == 0 || sampledDeadlinePartitionCount < maxPartitions) {
			sampledDeadlinePartitionCount++
		}
	}

	terminationsCh := make(chan terminationTask)
	statsCh := make(chan *SectorStats)

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
					step := max(float64(len(sectors))/float64(batchSize-1), 1.0)
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
	for stats := range statsCh {
		allStats = allStats.Accumulate(stats)
	}
	if sectorsCount > 0 {
		allStats = allStats.ScaleUp(int64(sectorsCount),
			int64(sectorsCount-sectorsInSkippedPartitions))
	}

	resultCh <- &PreviewTerminateSectorsReturn{
		Actor:                      actor,
		MinerInfo:                  minerInfo,
		SectorStats:                allStats,
		SectorsTerminated:          sectorsTerminated,
		SectorsCount:               sectorsCount,
		SectorsInSkippedPartitions: sectorsInSkippedPartitions,
		Epoch:                      h,
		PartitionsCount:            uint64(len(deadlinePartitions)),
		SampledPartitionsCount:     sampledDeadlinePartitionCount,
		Tipset:                     ts,
	}
}
