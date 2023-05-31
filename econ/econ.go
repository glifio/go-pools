package econ

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	biggy "github.com/filecoin-project/go-state-types/big"
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

	aggMinerStats := mstat.NewMinerStats()

	// loop through each miner and add address to MinerData
	for _, miner := range minerIDs {
		minerStats, err := mstat.ComputeMinerStats(ctx, miner, tsk, lapi)
		if err != nil {
			log.Printf("error getting miner stats: %v", err)
			return data, err
		}

		aggMinerStats.Balance = big.NewInt(0).Add(aggMinerStats.Balance, minerStats.Balance)
		log.Printf("Amount: %v", types.FIL(biggy.NewFromGo(minerStats.Balance)).String())
		aggMinerStats.PenaltyTermination = big.NewInt(0).Add(aggMinerStats.PenaltyTermination, minerStats.PenaltyTermination)
		aggMinerStats.PenaltyFaultPerDay = big.NewInt(0).Add(aggMinerStats.PenaltyFaultPerDay, minerStats.PenaltyFaultPerDay)
		aggMinerStats.ExpectedDailyReward = big.NewInt(0).Add(aggMinerStats.ExpectedDailyReward, minerStats.ExpectedDailyReward)
		aggMinerStats.PledgedFunds = big.NewInt(0).Add(aggMinerStats.PledgedFunds, minerStats.PledgedFunds)
		aggMinerStats.VestingFunds = big.NewInt(0).Add(aggMinerStats.VestingFunds, minerStats.VestingFunds)

		// add sector stats
		aggMinerStats.LiveSectors = big.NewInt(0).Add(aggMinerStats.LiveSectors, minerStats.LiveSectors)
		aggMinerStats.FaultySectors = big.NewInt(0).Add(aggMinerStats.FaultySectors, minerStats.FaultySectors)

		aggMinerStats.GreenScore = big.NewInt(0).Add(aggMinerStats.GreenScore, minerStats.GreenScore)

		// only add power if miner has min power
		if minerStats.HasMinPower {
			aggMinerStats.QualityAdjPower = big.NewInt(0).Add(aggMinerStats.QualityAdjPower, minerStats.QualityAdjPower)
			aggMinerStats.HasMinPower = true
		}
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

	var faultRatio = big.NewFloat(0).Quo(new(big.Float).SetInt(data.FaultySectors), new(big.Float).SetInt(data.LiveSectors))

	var vestingToPledgeRatio = big.NewFloat(0).Quo(new(big.Float).SetInt(aggMinerStats.VestingFunds), new(big.Float).SetInt(aggMinerStats.PledgedFunds))

	data.Gcred = CreditScore(ead, ltv, ltcv, data.ExpectedDailyRewards, dte, faultRatio, vestingToPledgeRatio)

	return data, nil
}
