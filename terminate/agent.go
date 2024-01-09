package terminate

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-state-types/abi"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/util"
)

func PreviewAgentTermination(ctx context.Context, sdk types.PoolsSDK, agentAddr common.Address, tipset *ltypes.TipSet) (*AgentCollateralStats, error) {
	lapi, closer, err := sdk.Extern().ConnectLotusClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	// if no tipset is found, we use 3 epochs behind chainhead (so we dont get epoch syncronization errors)
	if tipset == nil {
		ch, err := lapi.ChainHead(ctx)
		if err != nil {
			return nil, err
		}

		tipset, err = lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(ch.Height()-3), ltypes.EmptyTSK)
		if err != nil {
			return nil, err
		}
	}

	bigHeight := big.NewInt(int64(tipset.Height()))

	miners, err := sdk.Query().AgentMiners(ctx, agentAddr, bigHeight)
	if err != nil {
		return nil, err
	}

	minerCount := int64(len(miners))

	tasks := make([]util.TaskFunc, minerCount)
	for i := int64(0); i < minerCount; i++ {
		minerAddr := miners[i]
		tasks[i] = func() (interface{}, error) {
			return PreviewTerminateSectorsQuick(context.Background(), *lapi, minerAddr, tipset)
		}
	}

	results, err := util.Multiread(tasks)
	if err != nil {
		return nil, err
	}

	terminationPenalty := big.NewInt(0)
	collateral := big.NewInt(0)
	minersTerminationStats := make([]*TerminateSectorsSummary, minerCount)
	for i, terminationStats := range results {
		terminationStat := terminationStats.(*PreviewTerminateSectorsReturn)
		// add the miners termination penalty to the aggregate
		terminationPenalty = new(big.Int).Add(terminationPenalty, terminationStat.SectorStats.TerminationPenalty)
		// add the miners balance to the aggregate
		collateral = new(big.Int).Add(collateral, terminationStat.Actor.Balance.Int)
		// append to the array of individual miners so the frontend can show better data
		minersTerminationStats[i] = &TerminateSectorsSummary{
			MinerAddr:          terminationStat.Actor.Address.String(),
			MinerBal:           terminationStat.Actor.Balance.String(),
			TerminationPenalty: terminationStat.SectorStats.TerminationPenalty.String(),
			SectorsTerminated:  terminationStat.SectorsTerminated,
			SectorsCount:       terminationStat.SectorsCount,
			MinExpiration:      terminationStat.SectorStats.MinExpiration,
			MaxExpiration:      terminationStat.SectorStats.MaxExpiration,
			MinAge:             terminationStat.SectorStats.MinAge,
			MaxAge:             terminationStat.SectorStats.MaxAge,
		}
	}

	agentLiquidFIL, err := sdk.Query().AgentLiquidAssets(ctx, agentAddr, bigHeight)
	if err != nil {
		return nil, err
	}

	liquidationValue := new(big.Int).Add(collateral, agentLiquidFIL)
	liquidationValue.Sub(liquidationValue, terminationPenalty)

	return &AgentCollateralStats{
		LiquidationValue:       liquidationValue.String(),
		AggTerminationPenalty:  terminationPenalty.String(),
		AgentLiquidCollateral:  agentLiquidFIL.String(),
		MinersTerminationStats: minersTerminationStats,
		Epoch:                  tipset.Height(),
	}, nil
}
