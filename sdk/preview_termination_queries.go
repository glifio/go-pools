package sdk

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"strings"

	corebig "math/big"

	"github.com/filecoin-project/go-state-types/big"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	poolstypes "github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/util"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)

// PreviewTerminateSector gets the burnt funds for when a single sector is terminated
func (q *fevmQueries) PreviewTerminateSector(
	ctx context.Context,
	minerAddr address.Address,
	tipset string,
	vmHeight uint64,
	sectorNumber uint64,
	gasLimit uint64,
) (actor *types.ActorV5, totalBurn *corebig.Int, epoch abi.ChainEpoch, err error) {
	lClient, closer, err := q.extern.ConnectLotusClient()
	if err != nil {
		return nil, nil, 0, err
	}
	defer closer()
	api := *lClient

	h, ts, err := parseTipSetAndHeight(ctx, api, tipset, vmHeight)
	if err != nil {
		return nil, nil, 0, err
	}

	// Lookup actor balance
	actor, err = api.StateGetActor(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, nil, 0, err
	}

	// Lookup current owner / worker
	minerInfo, err := api.StateMinerInfo(ctx, minerAddr, ts.Key())
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

	totalBurn, err = terminateSectors(ctx, *lClient, h, ts, minerAddr, minerInfo,
		params, int64(gasLimit))
	if err != nil {
		return nil, nil, 0, err
	}

	return actor, totalBurn, h, nil
}

type terminationTask struct {
	termination         miner.TerminationDeclaration
	deadlineHeight      abi.ChainEpoch
	deadlineTs          *types.TipSet
	sectorsCount        int64
	sampledSectorsCount int64
}

// PreviewTerminateSectors gets the burnt funds for when all the sectors are terminated
func (q *fevmQueries) PreviewTerminateSectors(
	ctx context.Context,
	minerAddr address.Address,
	tipset string,
	vmHeight uint64,
	batchSize uint64,
	gasLimit uint64,
	useSampling bool,
	errorCh chan error,
	progressCh chan *poolstypes.PreviewTerminateSectorsProgress,
	resultCh chan *poolstypes.PreviewTerminateSectorsReturn,
) {
	lClient, closer, err := q.extern.ConnectLotusClient()
	if err != nil {
		errorCh <- err
		return
	}
	defer closer()
	api := *lClient

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

	workerActor, err := api.StateGetActor(ctx, minerInfo.Worker, ts.Key())
	if err != nil {
		errorCh <- err
		return
	}

	var params miner.TerminateSectorsParams

	totalBurn := new(corebig.Int)

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

	if batchSize == 0 && gasLimit == 0 {
		batchSize = 2500
		gasLimit = 90000000000

		workerBal, _ := util.ToFIL(workerActorPrev.Balance.Int).Float64()
		if workerBal < 1.0 {
			ratio := workerBal / 1.0
			batchSize = uint64(float64(batchSize) * ratio)
			gasLimit = uint64(float64(gasLimit) * ratio)
		}

		if batchSize < 10 {
			errorCh <- xerrors.Errorf("Not enough worker funds! Batch size too small!")
			return
		}
	}

	progressCh <- &poolstypes.PreviewTerminateSectorsProgress{
		Epoch:                  h,
		MinerInfo:              minerInfo,
		WorkerActor:            workerActor,
		PrevHeightForImmutable: prevHeight,
		WorkerActorPrev:        workerActorPrev,
		BatchSize:              batchSize,
		GasLimit:               gasLimit,
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
	var dlIdx uint64

	deadlinePartitions := make([]deadlinePartition, 0)

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

	terminations := make(chan terminationTask)
	burns := make(chan *corebig.Int)

	go runTerminationsInBatches(terminations, burns)

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
			progressCh <- &poolstypes.PreviewTerminateSectorsProgress{
				DeadlinePartitionCount: len(deadlinePartitions),
				DeadlinePartitionIndex: deadlinePartitionIdx,
				Deadline:               dlIdx,
				DeadlineImmutable:      dlImmutable,
				Partition:              partIdx,
				SectorsCount:           sc,
			}
			if sc > 0 {
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

						progressCh <- &poolstypes.PreviewTerminateSectorsProgress{
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
						terminations <- terminationTask{
							termination:         termination,
							deadlineHeight:      deadlineHeight,
							deadlineTs:          deadlineTs,
							sectorsCount:        int64(sc),
							sampledSectorsCount: int64(sc),
						}

						params = miner.TerminateSectorsParams{
							Terminations: []miner.TerminationDeclaration{termination},
						}
						burn, err := terminateSectors(ctx, *lClient, deadlineHeight, deadlineTs,
							minerAddr, minerInfo, params, int64(gasLimit))
						if err != nil {
							errorCh <- err
							return
						}
						sectorsTerminated += sliceCount
						sectorsCount += sliceCount
						totalBurn = totalBurn.Add(totalBurn, burn)
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
					// fmt.Printf("Jim sampledSectors len: %v\n", len(sampledSectors))
					// fmt.Printf("Jim sampledSectors: %v\n", sampledSectors)
					sampledSectorsBitfield := bitfield.NewFromSet(sampledSectors)
					termination := miner.TerminationDeclaration{
						Deadline:  uint64(dlIdx),
						Partition: uint64(partIdx),
						Sectors:   sampledSectorsBitfield,
					}
					terminations <- terminationTask{
						termination:         termination,
						deadlineHeight:      deadlineHeight,
						deadlineTs:          deadlineTs,
						sectorsCount:        int64(sc),
						sampledSectorsCount: int64(len(sampledSectors)),
					}
					params = miner.TerminateSectorsParams{
						Terminations: []miner.TerminationDeclaration{termination},
					}
					burn, err := terminateSectors(ctx, *lClient, deadlineHeight, deadlineTs,
						minerAddr, minerInfo, params, int64(gasLimit))
					if err != nil {
						errorCh <- err
						return
					}
					scaledBurn := new(corebig.Int).Div(
						new(corebig.Int).Mul(burn, corebig.NewInt(int64(sc))),
						corebig.NewInt(int64(len(sampledSectors))),
					)
					sectorsTerminated += uint64(len(sampledSectors))
					sectorsCount += uint64(len(sectors))
					totalBurn = totalBurn.Add(totalBurn, scaledBurn)
				}
			}
		}
		close(terminations)
	}()

	for burn := range burns {
		fmt.Printf("Jim burn2: %v\n", burn)
	}

	resultCh <- &poolstypes.PreviewTerminateSectorsReturn{
		Actor:             actor,
		TotalBurn:         totalBurn,
		SectorsTerminated: sectorsTerminated,
		SectorsCount:      sectorsCount,
		Epoch:             h,
	}
}

const maxPartitions = 3

func runTerminationsInBatches(tasks chan terminationTask, burns chan *corebig.Int) {
	pending := make([]terminationTask, 0, maxPartitions)
	var pendingDeadlineHeight abi.ChainEpoch
	for task := range tasks {
		if pendingDeadlineHeight != 0 && task.deadlineHeight != pendingDeadlineHeight {
			runPendingTerminations(pending, burns)
			pending = pending[:0]
			pendingDeadlineHeight = 0
		}
		pending = append(pending, task)
		pendingDeadlineHeight = task.deadlineHeight
		if len(pending) == maxPartitions {
			runPendingTerminations(pending, burns)
			pending = pending[:0]
			pendingDeadlineHeight = 0
		}
	}
	runPendingTerminations(pending, burns)
	close(burns)
}

func runPendingTerminations(tasks []terminationTask, burns chan *corebig.Int) {
	if len(tasks) > 0 {
		fmt.Printf("Terminating: %+v\n", tasks)
		height := tasks[0].deadlineHeight
		var sectorsCount int64
		var sampledSectorsCount int64
		// ts := tasks[0].deadlineTs
		terminations := make([]miner.TerminationDeclaration, 0)
		for _, task := range tasks {
			if task.deadlineHeight != height {
				panic("Bad height") // Should never happen
			}
			terminations = append(terminations, task.termination)
			sectorsCount += task.sectorsCount
			sampledSectorsCount += task.sampledSectorsCount
		}
		params := miner.TerminateSectorsParams{
			Terminations: terminations,
		}
		fmt.Printf("Jim sectorsCount: %v\n", sectorsCount)
		fmt.Printf("Jim sampledSectorsCount: %v\n", sampledSectorsCount)
		fmt.Printf("Jim params: %+v\n", params)
		burns <- corebig.NewInt(0)

	}
}

func terminateSectors(
	ctx context.Context,
	api lotusapi.FullNodeStruct,
	height abi.ChainEpoch,
	ts *types.TipSet,
	minerAddr address.Address,
	minerInfo lotusapi.MinerInfo,
	params miner.TerminateSectorsParams,
	gasLimit int64,
) (*corebig.Int, error) {
	enc := new(bytes.Buffer)
	err := params.MarshalCBOR(enc)
	if err != nil {
		return nil, err
	}

	workerActor, err := api.StateGetActor(ctx, minerInfo.Worker, ts.Key())
	if err != nil {
		return nil, err
	}

	nonce := workerActor.Nonce

	workerAddress, err := api.StateAccountKey(ctx, minerInfo.Worker, ts.Key())
	if err != nil {
		return nil, err
	}

	tipsetMsgs, err := api.ChainGetMessagesInTipset(ctx, ts.Key())
	if err != nil {
		return nil, err
	}
	for _, tMsg := range tipsetMsgs {
		if (tMsg.Message.From == minerInfo.Worker ||
			tMsg.Message.From == workerAddress) &&
			tMsg.Message.Nonce >= nonce {
			nonce = tMsg.Message.Nonce + 1
		}
	}

	msg := &types.Message{
		Version:    0,
		To:         minerAddr,
		From:       minerInfo.Worker,
		Nonce:      nonce,
		Value:      big.NewInt(0),
		GasLimit:   gasLimit,
		GasFeeCap:  big.NewInt(10135200),
		GasPremium: big.NewInt(0),
		Method:     9, // Terminate sectors
		Params:     enc.Bytes(),
	}
	cid := msg.Cid()

	var msgs []*types.Message

	msgs = append(msgs, msg)

	result, err := api.StateCompute(context.Background(), height, msgs, ts.Key())
	if err != nil {
		return nil, err
	}

	burn := big.NewInt(0)
	burnAddr, _ := address.NewFromString("f099")

	for _, trace := range result.Trace {
		if trace.MsgCid == cid {
			if trace.MsgRct.ExitCode != 0 {
				return nil, xerrors.Errorf("Exit Code: %v", trace.MsgRct.ExitCode)
			}
			for _, subMsg := range trace.ExecutionTrace.Subcalls {
				if subMsg.Msg.To == burnAddr {
					burn = subMsg.Msg.Value
				}
			}
			if trace.Error != "" {
				return nil, xerrors.Errorf("Error: %v\n", trace.Error)
			}

			var ret []bool

			err = cbor.DecodeReader(bytes.NewReader(trace.MsgRct.Return), &ret)
			if err != nil {
				return nil, err
			}
			if !ret[0] {
				return nil, xerrors.Errorf("Error: Not all sectors terminated!")
			}
		}
	}
	return burn.Int, nil
}

func ParseTipSetRef(ctx context.Context, api lotusapi.FullNodeStruct, tss string) (*types.TipSet, error) {
	if tss[0] == '@' {
		if tss == "@head" {
			return api.ChainHead(ctx)
		}

		var h uint64
		if _, err := fmt.Sscanf(tss, "@%d", &h); err != nil {
			return nil, xerrors.Errorf("parsing height tipset ref: %w", err)
		}

		return api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(h), types.EmptyTSK)
	}

	cids, err := ParseTipSetString(tss)
	if err != nil {
		return nil, err
	}

	if len(cids) == 0 {
		return nil, nil
	}

	k := types.NewTipSetKey(cids...)
	ts, err := api.ChainGetTipSet(ctx, k)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func ParseTipSetString(ts string) ([]cid.Cid, error) {
	strs := strings.Split(ts, ",")

	var cids []cid.Cid
	for _, s := range strs {
		c, err := cid.Parse(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		cids = append(cids, c)
	}

	return cids, nil
}

func parseTipSetAndHeight(ctx context.Context, api lotusapi.FullNodeStruct, tipset string, vmHeight uint64) (h abi.ChainEpoch, ts *types.TipSet, err error) {
	h = abi.ChainEpoch(vmHeight)
	if tss := tipset; tss != "" {
		ts, err = ParseTipSetRef(ctx, api, tss)
	} else if h > 0 {
		ts, err = api.ChainGetTipSetByHeight(ctx, h, types.EmptyTSK)
	} else {
		ts, err = api.ChainHead(ctx)
	}
	if err != nil {
		return 0, nil, err
	}

	if h == 0 {
		h = ts.Height()
	}

	return h, ts, nil
}
