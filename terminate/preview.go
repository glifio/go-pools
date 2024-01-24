package terminate

import (
	"context"
	"math"
	"math/big"

	"github.com/filecoin-project/go-state-types/builtin/v12/miner"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/mstat"
)

// PreviewTerminateSector will preview the cost of terminating a single sector
// on a miner.
//
//   - tipset takes a string selector, similar to Lotus, eg. "@1234", or "cid,cid,cid"
//   - offchain controls whether or not to use the "on-chain" or "off-chain"
//     calculation method.
func PreviewTerminateSector(
	ctx context.Context,
	api *lotusapi.FullNodeStruct,
	minerAddr address.Address,
	tipset string,
	vmHeight uint64,
	sectorNumber uint64,
	gasLimit uint64,
	offchain bool,
) (actor *types.ActorV5, sectorStats *SectorStats, epoch abi.ChainEpoch, err error) {
	h, ts, err := parseTipSetAndHeight(ctx, api, tipset, vmHeight)
	if err != nil {
		return nil, nil, 0, err
	}

	// Lookup actor balance
	mactor, mstate, err := mstat.LoadMinerActor(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, nil, 0, err
	}
	actor = mactor

	// Lookup current owner / worker
	minerInfo, err := mstate.Info()
	if err != nil {
		return nil, nil, 0, err
	}

	var params miner.TerminateSectorsParams

	if gasLimit == 0 {
		gasLimit = 6000000000
	}

	sectorPartition, err := api.StateSectorPartition(ctx, minerAddr, abi.SectorNumber(sectorNumber), ts.Key())
	if err != nil {
		return nil, nil, 0, err
	}

	termination := miner.TerminationDeclaration{
		Deadline:  sectorPartition.Deadline,
		Partition: sectorPartition.Partition,
		Sectors:   bitfield.NewFromSet([]uint64{sectorNumber}),
	}
	params = miner.TerminateSectorsParams{
		Terminations: []miner.TerminationDeclaration{termination},
	}

	stats, err := terminateSectors(ctx, api, h, ts, minerAddr, minerInfo,
		params, int64(gasLimit), offchain)
	if err != nil {
		return nil, nil, 0, err
	}

	return actor, stats, h, nil
}

// PreviewTerminateSectors will preview the cost of terminating all the sectors
// for a miner.
//
// This function is intended to be run as a goroutine. Channels are used to
// stream back progress and results/errors, so clients can report progress as
// the query progresses.
//
//   - tipset takes a string selector, similar to Lotus, eg. "@1234", or "cid,cid,cid"
//   - batchSize is the number of sectors to query per message. If set to zero,
//     and gasLimit is also zero, a default batch size will be calculated based
//     on the miner's worker balance.
//   - offchain controls whether or not to use the "on-chain" or "off-chain"
//     calculation method.
//   - optimize will make additional queries to try to select deadlines to
//     sample that are evenly distributed over historical time.
//   - maxPartitions controls how many partitions to sample. If set to zero,
//     all partitions will be sampled.
//
// Use the simpler, more opinionated [PreviewTerminateSectorsQuick] function
// for quick results using sampling, "off-chain" calcuation, and a simpler
// API (no channels for streaming progress)
func PreviewTerminateSectors(
	ctx context.Context,
	api *lotusapi.FullNodeStruct,
	minerAddr address.Address,
	tipset string,
	vmHeight uint64,
	batchSize uint64,
	gasLimit uint64,
	useSampling bool, // if set, not all sectors will be sampled
	optimize bool,
	offchain bool,
	maxPartitions uint64,
	errorCh chan error,
	progressCh chan *PreviewTerminateSectorsProgress,
	resultCh chan *PreviewTerminateSectorsReturn,
) {
	h, ts, err := parseTipSetAndHeight(ctx, api, tipset, vmHeight)
	if err != nil {
		errorCh <- err
		return
	}

	// Lookup actor balance
	actor, mstate, err := mstat.LoadMinerActor(ctx, api, minerAddr, ts)

	// Lookup current owner / worker
	minerInfo, err := mstate.Info()
	if err != nil {
		errorCh <- err
		return
	}

	lf, err := mstate.LockedFunds()
	if err != nil {
		errorCh <- err
		return
	}

	workerActor, err := api.StateGetActor(ctx, minerInfo.Worker, ts.Key())
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
	}

	progressCh <- &PreviewTerminateSectorsProgress{
		Epoch:                  h,
		MinerInfo:              minerInfo,
		WorkerActor:            workerActor,
		PrevHeightForImmutable: prevHeight,
		WorkerActorPrev:        workerActorPrev,
		BatchSize:              batchSize,
		GasLimit:               gasLimit,
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
			if optimize {
				allLiveSectors = append(allLiveSectors, partition.LiveSectors)
			}
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
	if optimize {
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
	} else {
		var step float64
		if maxPartitions > 0 {
			step = max(float64(len(deadlinePartitions))/float64(maxPartitions), 1.0)
		} else {
			step = 1
		}
		for sampleIndex := 0.0; int(sampleIndex) < len(deadlinePartitions); sampleIndex += step {
			i := int(sampleIndex)
			sampledDeadlinePartitions[i] = true
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

	go runTerminationsInBatches(ctx, api, minerAddr, minerInfo,
		int64(gasLimit), offchain, terminationsCh, statsCh, errorCh)

	go func() {
		for deadlinePartitionIdx, deadlinePartition := range deadlinePartitions {
			dlIdx := deadlinePartition.dlIdx
			dlImmutable := deadlinePartition.dlImmutable
			deadlineTs := deadlinePartition.deadlineTs
			deadlineHeight := deadlinePartition.deadlineHeight
			partIdx := deadlinePartition.partIdx
			partition := deadlinePartition.partition

			sc, err := partition.LiveSectors.Count()
			if err != nil {
				errorCh <- err
				return
			}
			progressCh <- &PreviewTerminateSectorsProgress{
				DeadlinePartitionCount: len(deadlinePartitions),
				DeadlinePartitionIndex: deadlinePartitionIdx,
				Deadline:               dlIdx,
				DeadlineImmutable:      dlImmutable,
				Partition:              partIdx,
				SectorsCount:           sc,
			}
			if sc > 0 {

				if sampledDeadlinePartitions[deadlinePartitionIdx] {
					if !useSampling {
						var i uint64
						for i = 0; i < sc; i += batchSize {
							lastIndex := min(sc, i+batchSize)
							slicedSectors, err := partition.LiveSectors.Slice(i, lastIndex-i)
							if err != nil {
								errorCh <- err
								return
							}
							sliceCount, err := slicedSectors.Count()
							if err != nil {
								errorCh <- err
								return
							}

							progressCh <- &PreviewTerminateSectorsProgress{
								DeadlinePartitionCount: len(deadlinePartitions),
								DeadlinePartitionIndex: deadlinePartitionIdx,
								Deadline:               dlIdx,
								DeadlineImmutable:      dlImmutable,
								Partition:              partIdx,
								SectorsCount:           sc,
								SliceStart:             i,
								SliceEnd:               lastIndex,
								SliceCount:             sliceCount,
							}

							termination := miner.TerminationDeclaration{
								Deadline:  uint64(dlIdx),
								Partition: uint64(partIdx),
								Sectors:   slicedSectors,
							}
							terminationsCh <- terminationTask{
								termination:         termination,
								deadlineHeight:      deadlineHeight,
								deadlineTs:          deadlineTs,
								sectorsCount:        int64(sc),
								sampledSectorsCount: int64(sc),
							}
							sectorsTerminated += sliceCount
							sectorsCount += sliceCount
						}
					} else {
						// useSampling
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
					}
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
	}
}
