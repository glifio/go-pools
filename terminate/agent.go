package terminate

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-state-types/abi"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/constants"
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
		AgentLiquidBalance: agentLiquidFIL,
	}, nil
}

func (term PreviewAgentTerminationSummary) LiquidationValue() *big.Int {
	// if the termination penalty is greater than the initial pledge,
	// OR if the initial pledge is 0,
	// we report liquidation value as 0
	if term.InitialPledge.Cmp(term.TerminationPenalty) < 0 || term.InitialPledge.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	// first we multiply the available balance by the recovery rate
	// recovery rate = (initial pledge - termination penalty) / initial pledge
	// discounted avail = (available * initial pledge - available * termination penalty) / initial pledge
	availTimesPledge := new(big.Int).Mul(term.AvailableBalance, term.InitialPledge)
	availTimesTermPenalty := new(big.Int).Mul(term.AvailableBalance, term.TerminationPenalty)

	discountedAvail := new(big.Int).Sub(availTimesPledge, availTimesTermPenalty)
	discountedAvail.Div(discountedAvail, term.InitialPledge)

	// we add the discounted available balance to the vesting balance and initial pledge, subtract the termination penalty to get the liquidation value
	liquidationValue := new(big.Int).Add(discountedAvail, term.VestingBalance)
	liquidationValue.Add(liquidationValue, term.InitialPledge)
	liquidationValue.Sub(liquidationValue, term.TerminationPenalty)

	return liquidationValue
}

// 1e18 = 100%
func (term PreviewAgentTerminationSummary) RecoveryRate() *big.Int {
	if term.InitialPledge.Cmp(term.TerminationPenalty) == -1 || term.InitialPledge.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	// recovery rate = (initial pledge - termination penalty) / initial pledge
	recoveryRate := new(big.Int).Sub(term.InitialPledge, term.TerminationPenalty)
	recoveryRate.Mul(recoveryRate, constants.WAD)
	recoveryRate.Quo(recoveryRate, term.InitialPledge)

	return recoveryRate
}
