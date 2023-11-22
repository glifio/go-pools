package sdk

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	corebig "math/big"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/schollz/progressbar/v3"
)

// PreviewTerminateSectors gets the burnt funds for when the sectors are terminated
func (q *fevmQueries) PreviewTerminateSectors(
	ctx context.Context,
	rpcUrl string,
	minerID string,
	tipset string,
	vmHeight uint64,
	sectorNumber uint64,
	allSectors bool,
	batchSize uint64,
	gasLimit uint64,
	quiet bool,
) {
	lClient, closer, err := q.extern.ConnectLotusClient()
	if err != nil {
		// return nil, err
		return
	}
	defer closer()
	api := *lClient

	// return lClient.ChainHead(ctx)

	minerAddr, err := address.NewFromString(minerID)
	if err != nil {
		log.Fatal(err)
	}

	h := abi.ChainEpoch(vmHeight)
	var ts *types.TipSet
	if tss := tipset; tss != "" {
		ts, err = ParseTipSetRef(ctx, api, tss)
	} else if h > 0 {
		ts, err = api.ChainGetTipSetByHeight(ctx, h, types.EmptyTSK)
	} else {
		ts, err = api.ChainHead(ctx)
	}
	if err != nil {
		log.Fatal(err)
	}

	if h == 0 {
		h = ts.Height()
	}

	// Lookup actor balance
	actor, err := api.StateGetActor(ctx, minerAddr, ts.Key())
	if err != nil {
		log.Fatal(err)
	}

	// Lookup current owner / worker
	minerInfo, err := api.StateMinerInfo(ctx, minerAddr, ts.Key())
	if err != nil {
		log.Fatal(err)
	}

	workerActor, err := api.StateGetActor(ctx, minerInfo.Worker, ts.Key())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Miner: %v\n", minerAddr)
	fmt.Printf("Epoch: %d\n", h)
	fmt.Printf("Worker: %v (Balance: %v)\n", minerInfo.Worker, util.ToFIL(workerActor.Balance.Int))

	var params miner.TerminateSectorsParams

	totalBurn := new(corebig.Int)

	if allSectors {
		provingDeadline, err := api.StateMinerProvingDeadline(ctx, minerAddr, ts.Key())
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Proving deadline: %+v\n", provingDeadline)
		// fmt.Printf("Proving deadline Index: %+v\n", provingDeadline.Index)

		// prevHeight := provingDeadline.PeriodStart - 120
		prevHeight := h - 120

		tsPrev, err := api.ChainGetTipSetByHeight(ctx, prevHeight, types.EmptyTSK)
		if err != nil {
			log.Fatal(err)
		}

		workerActorPrev, err := api.StateGetActor(ctx, minerInfo.Worker, tsPrev.Key())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Epoch used for immutable deadlines: %d (Worker balance: %v)\n", prevHeight, util.ToFIL(workerActorPrev.Balance.Int))

		if batchSize == 0 && gasLimit == 0 {
			batchSize = 2500
			// gasLimit = 60000000000
			gasLimit = 90000000000

			workerBal, _ := util.ToFIL(workerActorPrev.Balance.Int).Float64()
			if workerBal < 1.0 {
				ratio := workerBal / 1.0
				batchSize = uint64(float64(batchSize) * ratio)
				gasLimit = uint64(float64(gasLimit) * ratio)
			}

			fmt.Printf("Batch Size: %d\n", batchSize)
			fmt.Printf("Gas Limit: %d\n", gasLimit)
			if batchSize < 10 {
				log.Fatal("Not enough worker funds! Batch size too small!")
			}
		}

		var bar *progressbar.ProgressBar
		if quiet {
			bar = progressbar.NewOptions(48,
				progressbar.OptionSetDescription("Deadlines"),
				progressbar.OptionSetWriter(os.Stderr),
				progressbar.OptionSetWidth(10),
				progressbar.OptionThrottle(65*time.Millisecond),
				progressbar.OptionShowCount(),
				progressbar.OptionShowIts(),
				/*
					OptionOnCompletion(func() {
						fmt.Fprint(os.Stderr, "\n")
					}),
				*/
				progressbar.OptionSpinnerType(14),
				progressbar.OptionFullWidth(),
				progressbar.OptionSetRenderBlankState(true),
				progressbar.OptionClearOnFinish())
		}
		defer func() {
			if bar != nil {
				bar.Close()
			}
		}()

		var dlIdx uint64
		for dlIdx = 0; dlIdx < 48; dlIdx++ {
			if quiet {
				bar.Add(1)
			}
			immutable := ""
			deadlineTs := ts
			deadlineHeight := h
			if dlIdx == provingDeadline.Index || dlIdx == (provingDeadline.Index+1)%48 {
				// Immutable deadline
				immutable = " (Immutable)"
				deadlineTs = tsPrev
				deadlineHeight = prevHeight
			}
			partitions, err := api.StateMinerPartitions(ctx, minerAddr, uint64(dlIdx), deadlineTs.Key())
			if err != nil {
				log.Fatal(err)
			}
			for partIdx, partition := range partitions {
				sc, err := partition.LiveSectors.Count()
				if err != nil {
					log.Fatal(err)
				}
				if sc > 0 {
					if !quiet {
						fmt.Printf("Deadline %d%s Partition %d Sectors %d\n", dlIdx, immutable, partIdx, sc)
					}
					var i uint64
					for i = 0; i < sc; i += batchSize {
						lastIndex := min(sc, i+batchSize)
						slicedSectors, err := partition.LiveSectors.Slice(i, lastIndex-i)
						if err != nil {
							log.Fatal(err)
						}
						sliceCount, err := slicedSectors.Count()
						if err != nil {
							log.Fatal(err)
						}

						if !quiet {
							fmt.Printf("  Slice: %d to %d (->%d): %d\n", i, lastIndex-1, sc-1, sliceCount)
						}

						termination := miner.TerminationDeclaration{
							Deadline:  uint64(dlIdx),
							Partition: uint64(partIdx),
							Sectors:   slicedSectors,
						}
						params = miner.TerminateSectorsParams{
							Terminations: []miner.TerminationDeclaration{termination},
						}
						burn, err := terminateSectors(ctx, *lClient, deadlineHeight, deadlineTs,
							minerAddr, minerInfo, params, int64(gasLimit))
						if err != nil {
							if quiet {
								bar.Close()
							}
							log.Fatal(err)
						}
						totalBurn = totalBurn.Add(totalBurn, burn)
					}
				}
			}
		}
		if quiet {
			bar.Close()
		}
	} else {
		if gasLimit == 0 {
			gasLimit = 6000000000
		}

		sectorPartition, err := api.StateSectorPartition(ctx, minerAddr, abi.SectorNumber(sectorNumber), ts.Key())
		if err != nil {
			log.Fatal(err)
		}

		termination := miner.TerminationDeclaration{
			Deadline:  sectorPartition.Deadline,
			Partition: sectorPartition.Partition,
			Sectors:   bitfield.NewFromSet([]uint64{sectorNumber}),
		}
		params = miner.TerminateSectorsParams{
			Terminations: []miner.TerminationDeclaration{termination},
		}
		fmt.Printf("Sector: %v\n", sectorNumber)
		fmt.Printf("Deadline: %v\n", sectorPartition.Deadline)
		fmt.Printf("Partition: %v\n", sectorPartition.Partition)
		burn, err := terminateSectors(ctx, *lClient, h, ts, minerAddr, minerInfo,
			params, int64(gasLimit))
		if err != nil {
			log.Fatal(err)
		}
		totalBurn = totalBurn.Add(totalBurn, burn)
	}
	fmt.Printf("\nMiner actor balance (attoFIL): %s\n", actor.Balance)
	fmt.Printf("Total burn (attoFIL): %s\n", totalBurn)
	fmt.Printf("Miner actor balance (FIL): %0.03f\n", util.ToFIL(actor.Balance.Int))
	fmt.Printf("Total burn (FIL): %0.03f\n", util.ToFIL(totalBurn))
	difference := new(corebig.Int).Sub(actor.Balance.Int, totalBurn)
	fmt.Printf("Approximate liquidation value (FIL): %0.03f\n", util.ToFIL(difference))
	diff, _ := util.ToFIL(difference).Float64()
	balance, _ := util.ToFIL(actor.Balance.Int).Float64()
	pct := diff / balance * 100
	fmt.Printf("Approximate recovery percentage: %0.03f%%\n", pct)
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

	// fmt.Printf("Miner Owner: %v\n", minerInfo.Owner)
	// fmt.Printf("Miner Worker: %v\n", minerInfo.Worker)
	// fmt.Printf("Worker Nonce: %v\n", workerActor.Nonce)
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
			// fmt.Printf("Jim found nonce %d\n", tMsg.Message.Nonce)
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

	/*
		out, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	*/

	burn := big.NewInt(0)
	burnAddr, _ := address.NewFromString("f099")

	for _, trace := range result.Trace {
		if trace.MsgCid == cid {
			// fmt.Printf("CID: %v\n", cid)

			if trace.MsgRct.ExitCode != 0 {
				fmt.Printf("Exit Code: %v\n", trace.MsgRct.ExitCode)
			}
			for _, subMsg := range trace.ExecutionTrace.Subcalls {
				if subMsg.Msg.To == burnAddr {
					burn = subMsg.Msg.Value
					// fmt.Printf("Burn: %v\n", burn)
				}
			}
			if trace.Error != "" {
				fmt.Printf("Error: %v\n", trace.Error)
			}

			var ret []bool

			err = cbor.DecodeReader(bytes.NewReader(trace.MsgRct.Return), &ret)
			if err != nil {
				return nil, err
			}
			//fmt.Printf("Return: %+v\n", ret[0])
			if !ret[0] {
				log.Fatal("Error: Not all sectors terminated!")
			}

			/*
				out, err := json.Marshal(trace.MsgRct)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(out))
			*/
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
			return nil, fmt.Errorf("parsing height tipset ref: %w", err)
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
