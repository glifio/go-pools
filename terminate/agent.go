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

var LookbackEpochs abi.ChainEpoch = 3

// PreviewAgentTermination preview terminating all the
// sectors on all the miners for an agent (using sampling and "off-chain"
// calculation) and will return the liquidation value of the agent.
func PreviewAgentTermination(ctx context.Context, sdk types.PoolsSDK, agentAddr common.Address, tipset *ltypes.TipSet) (PreviewAgentTerminationSummary, error) {
	lapi, closer, err := sdk.Extern().ConnectLotusClient()
	if err != nil {
		return PreviewAgentTerminationSummary{}, err
	}
	defer closer()

	// if no tipset is found, we use 3 epochs behind chainhead (so we dont get epoch syncronization errors)
	if tipset == nil {
		ch, err := lapi.ChainHead(ctx)
		if err != nil {
			return PreviewAgentTerminationSummary{}, err
		}

		tipset, err = lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(ch.Height()-LookbackEpochs), ltypes.EmptyTSK)
		if err != nil {
			return PreviewAgentTerminationSummary{}, err
		}
	}

	bigHeight := big.NewInt(int64(tipset.Height()))

	miners, err := sdk.Query().AgentMiners(ctx, agentAddr, bigHeight)
	if err != nil {
		return PreviewAgentTerminationSummary{}, err
	}

	minerCount := int64(len(miners))

	tasks := make([]util.TaskFunc, minerCount)
	for i := int64(0); i < minerCount; i++ {
		minerAddr := miners[i]
		tasks[i] = func() (interface{}, error) {
			return PreviewTerminateSectorsQuick(context.Background(), lapi, minerAddr, tipset)
		}
	}

	results, err := util.Multiread(tasks)
	if err != nil {
		return PreviewAgentTerminationSummary{}, err
	}

	terminationPenaltyAgg := big.NewInt(0)
	initialPledgeAgg := big.NewInt(0)
	vestingBalanceAgg := big.NewInt(0)
	availableBalanceAgg := big.NewInt(0)

	for _, terminationStats := range results {
		terminationStat := terminationStats.(*PreviewTerminateSectorsReturn)
		// add the miners termination penalty to the aggregate
		terminationPenaltyAgg = new(big.Int).Add(terminationPenaltyAgg, terminationStat.SectorStats.TerminationPenalty)
		// add the miners bals to their aggregate counterpart
		initialPledgeAgg = new(big.Int).Add(initialPledgeAgg, terminationStat.InitialPledge)
		vestingBalanceAgg = new(big.Int).Add(vestingBalanceAgg, terminationStat.VestingBalance)
		availableBalanceAgg = new(big.Int).Add(availableBalanceAgg, terminationStat.Actor.Balance.Int)
	}

	agentLiquidFIL, err := sdk.Query().AgentLiquidAssets(ctx, agentAddr, bigHeight)
	if err != nil {
		return PreviewAgentTerminationSummary{}, err
	}

	availableBalanceAgg.Add(availableBalanceAgg, agentLiquidFIL)

	return PreviewAgentTerminationSummary{
		TerminationPenalty: terminationPenaltyAgg,
		InitialPledge:      initialPledgeAgg,
		VestingBalance:     vestingBalanceAgg,
		AvailableBalance:   availableBalanceAgg,
	}, nil
}
