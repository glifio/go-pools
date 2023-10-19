package econ

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/mstat"
	poolstypes "github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/vc"
)

// TODO https://github.com/glif-confidential/ado/issues/9
func ComputeAgentData(
	ctx context.Context,
	sdk poolstypes.PoolsSDK,
	agentLiquidAssets *big.Int,
	principal *big.Int,
	minerIDs []address.Address,
	agentAddr common.Address,
	tsk *types.TipSet,
) (*vc.AgentData, error) {
	lapi, closer, err := sdk.Extern().ConnectLotusClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	// here we just work our way through the AgentData, computing each key
	// TODO: could probably parellize this elegently to speed things up
	data := &vc.AgentData{}

	aggMinerStats, err := mstat.ComputeMinersStats(ctx, minerIDs, tsk, lapi)
	if err != nil {
		return nil, err
	}

	data.QaPower = aggMinerStats.QualityAdjPower

	data.GreenScore = aggMinerStats.GreenScore

	data.AgentValue = big.NewInt(0).Add(agentLiquidAssets, aggMinerStats.Balance)

	data.LiveSectors = aggMinerStats.LiveSectors
	data.FaultySectors = aggMinerStats.FaultySectors

	/* ~~~~~ CollateralValue ~~~~~ */

	// CV = AgentValue * terminationPenalty
	data.CollateralValue = CollateralValue(data.AgentValue)

	/* ~~~~~ ExpectedDailyFaultPenalties ~~~~~ */

	// COULD REMOVE
	data.ExpectedDailyFaultPenalties = aggMinerStats.PenaltyFaultPerDay

	/* ~~~~~ ExpectedDailyRewards ~~~~~ */

	data.ExpectedDailyRewards = aggMinerStats.ExpectedDailyReward

	/* ~~~~~ Principal ~~~~~ */
	data.Principal = principal

	// TODO: What to do on a Pay action type? -> deduce principal _after_ the payment is made

	/* ~~~~~ GCRED ~~~~~ */

	// exposure at default is current principal, including newPrincipal
	var ead = data.Principal
	// loan to value computation uses agent's total value, less principal as its denominator
	// valueDenominator := lockedFunds.Balance.Add(lockedFunds.Balance, agentLiquidAssets)
	// subtract liabilities (upcoming payment or withdrawal) from the denominator
	// valueDenominator = valueDenominator.Sub(valueDenominator, deductibles)

	var ltv = LoanToValueRatio(ead, data.AgentValue)
	// loan to collateral value
	var ltcv = LoanToCollateralRatio(ead, data.CollateralValue)

	var equity = big.NewInt(0).Sub(data.AgentValue, data.Principal)

	var dte = DebtToEquityRatio(ead, equity)

	var faultRatio = big.NewFloat(0)
	var vestingToPledgeRatio = big.NewFloat(0)

	if data.LiveSectors.Int64() > 0 {
		faultRatio = faultRatio.Quo(new(big.Float).SetInt(data.FaultySectors), new(big.Float).SetInt(data.LiveSectors))
		vestingToPledgeRatio = vestingToPledgeRatio.Quo(new(big.Float).SetInt(aggMinerStats.VestingFunds), new(big.Float).SetInt(aggMinerStats.PledgedFunds))
	}

	data.Gcred = CreditScoreSimple(ead, ltv, ltcv, data.ExpectedDailyRewards, dte, faultRatio, vestingToPledgeRatio)

	return data, nil
}
