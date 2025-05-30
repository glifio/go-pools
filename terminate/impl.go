package terminate

import (
	"bytes"
	"context"
	"errors"
	"math"
	corebig "math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v14/miner"
	minertypes "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/v6/actors/util/smoothing"
	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"
	"github.com/glifio/go-pools/util"
	cbor "github.com/ipfs/go-ipld-cbor"
	"golang.org/x/xerrors"
)

type terminationTask struct {
	termination         miner.TerminationDeclaration
	deadlineHeight      abi.ChainEpoch
	deadlineTs          *types.TipSet
	sectorsCount        int64
	sampledSectorsCount int64
}

const maxPartitionsPerTx = 3

func runTerminationsInBatches(
	ctx context.Context,
	api *lotusapi.FullNodeStruct,
	minerAddr address.Address,
	minerInfo minertypes.MinerInfo,
	gasLimit int64,
	offchain bool,
	tasks chan terminationTask,
	statsCh chan *SectorStats,
	errorCh chan error,
) {
	defer func() {
		close(statsCh)
	}()
	if offchain {
		errorCh <- xerrors.Errorf("offchain termination is not supported")
		return
	}
	pending := make([]terminationTask, 0, maxPartitionsPerTx)
	flush := func() {
		err := runPendingTerminations(ctx, api, minerAddr, minerInfo, gasLimit, pending, offchain, statsCh)
		if err != nil {
			errorCh <- err
			return
		}
	}
	var pendingDeadlineHeight abi.ChainEpoch
	for task := range tasks {
		if pendingDeadlineHeight != 0 && task.deadlineHeight != pendingDeadlineHeight {
			flush()
			pending = pending[:0]
			pendingDeadlineHeight = 0
		}
		pending = append(pending, task)
		pendingDeadlineHeight = task.deadlineHeight
		if len(pending) == maxPartitionsPerTx {
			flush()
			pending = pending[:0]
			pendingDeadlineHeight = 0
		}
	}
	flush()
}

func runPendingTerminations(
	ctx context.Context,
	api *lotusapi.FullNodeStruct,
	minerAddr address.Address,
	minerInfo minertypes.MinerInfo,
	gasLimit int64,
	tasks []terminationTask,
	offchain bool,
	statsCh chan *SectorStats,
) error {
	if offchain {
		return xerrors.Errorf("offchain termination is not supported")
	}
	if len(tasks) > 0 {
		// fmt.Printf("Terminating: %+v\n", len(tasks))
		height := tasks[0].deadlineHeight
		ts := tasks[0].deadlineTs
		var sectorsCount int64
		var sampledSectorsCount int64
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
		stats, err := terminateSectors(ctx, api, height, ts,
			minerAddr, minerInfo, params, gasLimit, offchain)
		if err != nil {
			return err
		}
		statsCh <- stats.ScaleUp(sectorsCount, sampledSectorsCount)
	}
	return nil
}

func terminateSectors(
	ctx context.Context,
	api *lotusapi.FullNodeStruct,
	height abi.ChainEpoch,
	ts *types.TipSet,
	minerAddr address.Address,
	minerInfo minertypes.MinerInfo,
	params miner.TerminateSectorsParams,
	gasLimit int64,
	offchain bool,
) (*SectorStats, error) {
	stats := NewSectorStats()
	if offchain {
		smoothedPow, err := util.TotalPowerSmoothed(ctx, api, ts)
		if err != nil {
			return nil, err
		}

		powFilterEstimate := smoothing.NewEstimate(smoothedPow.PositionEstimate, smoothedPow.VelocityEstimate)

		smoothedRew, err := util.ThisEpochRewardsSmoothed(ctx, api, ts)
		if err != nil {
			return nil, err
		}

		rewFilterEstimate := smoothing.NewEstimate(smoothedRew.PositionEstimate, smoothedRew.VelocityEstimate)

		allTermSectors := make([]bitfield.BitField, 0)
		for _, term := range params.Terminations {
			allTermSectors = append(allTermSectors, term.Sectors)
		}
		allSectorsBitfield, err := bitfield.MultiMerge(allTermSectors...)
		if err != nil {
			return nil, err
		}

		sectors, err := api.StateMinerSectors(context.Background(), minerAddr, &allSectorsBitfield, ts.Key())
		if err != nil {
			return nil, err
		}

		for _, s := range sectors {
			sectorPower := miner8.QAPowerForSector(minerInfo.SectorSize, util.ConvertSectorType(s))

			// the termination penalty calculation
			var expectedDayReward, expectedStoragePledge, replacedDayReward abi.TokenAmount
			if s.ExpectedDayReward != nil {
				expectedDayReward = *s.ExpectedDayReward
			}
			if s.ExpectedStoragePledge != nil {
				expectedStoragePledge = *s.ExpectedStoragePledge
			}
			if s.ReplacedDayReward != nil {
				replacedDayReward = *s.ReplacedDayReward
			}

			termFee := miner8.PledgePenaltyForTermination(
				expectedDayReward,
				height-s.PowerBaseEpoch,
				expectedStoragePledge,
				powFilterEstimate,
				sectorPower,
				rewFilterEstimate,
				replacedDayReward,
				s.PowerBaseEpoch-s.Activation,
			)

			// this should never happen, but termination fee should never be negative
			if termFee.Cmp(big.Zero().Int) < 0 {
				return nil, xerrors.Errorf("negative termination fee: %v", termFee)
			}

			// the daily sector fee calculation
			sectorFee := miner8.PledgePenaltyForContinuedFault(rewFilterEstimate, powFilterEstimate, sectorPower)

			// incur the sector fee of Min(41 days, remaining days b4 sector expiry)
			epochsUntilTerm := s.Expiration - height
			daysUntilTerm := float64(epochsUntilTerm / builtin.EpochsInDay)
			sectorFeeDaysIncurred := int64(math.Min(41, daysUntilTerm))
			sectorFeesUntilTerm := sectorFee.Mul(sectorFee.Int, corebig.NewInt(sectorFeeDaysIncurred))
			// add the term fee and sector fees to the total fee

			stats = stats.AddSector(s, height, termFee.Int, sectorFeesUntilTerm)
		}

		return stats, nil
	} else {
		enc := new(bytes.Buffer)
		err := params.MarshalCBOR(enc)
		if err != nil {
			return nil, err
		}

		workerActor, err := api.StateGetActor(ctx, minerInfo.Worker, ts.Key())
		if err != nil {
			return nil, err
		}

		_, minerState, err := util.LoadMinerActor(ctx, api, minerAddr, ts)
		if err != nil {
			return nil, err
		}
		state := minerState.GetState().(*miner.State)
		empty, err := state.EarlyTerminations.IsEmpty()
		if err != nil {
			return nil, err
		}
		if !empty {
			return nil, errors.New("Early terminations exist in state, aborting")
		}

		nonce := workerActor.Nonce

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

		var msgs []*types.Message

		msgs = append(msgs, msg)

		trace, err := api.StateCall(context.Background(), msg, ts.Key())
		if err != nil {
			return nil, err
		}

		burnAddr, _ := address.NewFromString("f099")

		if trace.MsgRct.ExitCode != 0 {
			return nil, xerrors.Errorf("Exit Code: %v", trace.MsgRct.ExitCode)
		}
		for _, subMsg := range trace.ExecutionTrace.Subcalls {
			if subMsg.Msg.To == burnAddr {
				stats.TerminationPenalty = new(corebig.Int).Add(stats.TerminationPenalty,
					subMsg.Msg.Value.Int)
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

		return stats, nil
	}
}
