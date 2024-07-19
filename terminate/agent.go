package terminate

import (
	"math/big"

	"github.com/glifio/go-pools/constants"
)

func (ats PreviewAgentTerminationSummary) LiquidationValue() *big.Int {
	// total available balance is the sum of the miners available balance and the agent available balance
	totalAvailableBalance := new(big.Int).Add(ats.MinersAvailableBal, ats.AgentAvailableBal)

	// we add the total available balance to the vesting balance and initial pledge, subtract the termination penalty to get the liquidation value
	totalAssets := new(big.Int).Add(totalAvailableBalance, ats.VestingBalance)
	totalAssets.Add(totalAssets, ats.InitialPledge)

	// if the total assets is less than or equal the termination penalty, we return 0
	if totalAssets.Cmp(ats.TerminationPenalty) <= 0 {
		return big.NewInt(0)
	}

	// liquidation value = total assets - termination penalty
	liquidationValue := new(big.Int).Sub(totalAssets, ats.TerminationPenalty)

	return liquidationValue
}

// 1e18 = 100%
func (ats PreviewAgentTerminationSummary) RecoveryRate() *big.Int {
	if ats.InitialPledge.Cmp(ats.TerminationPenalty) == -1 || ats.InitialPledge.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	// recovery rate = (initial pledge - termination penalty) / initial pledge
	recoveryRate := new(big.Int).Sub(ats.InitialPledge, ats.TerminationPenalty)
	recoveryRate.Mul(recoveryRate, constants.WAD)
	recoveryRate.Quo(recoveryRate, ats.InitialPledge)

	return recoveryRate
}

func (ats PreviewAgentTerminationSummary) LTV(principal *big.Int) *big.Int {
	liquidationValue := ats.LiquidationValue()
	ltv := big.NewInt(0)
	if liquidationValue.Cmp(big.NewInt(0)) > 0 && principal.Cmp(big.NewInt(0)) > 0 {
		// add wad precision in prior to the div to keep atto denomination before we convert to percentage
		ltv.Mul(principal, constants.WAD)
		ltv.Div(ltv, liquidationValue)
	}

	return ltv
}

func (cs *AgentCollateralStats) Summarize() PreviewAgentTerminationSummary {
	availBal := big.NewInt(0)
	initialPledge := big.NewInt(0)
	vestingBal := big.NewInt(0)
	for _, miner := range cs.MinersTerminationStats {
		availBal.Add(availBal, miner.Available)
		initialPledge.Add(initialPledge, miner.Pledged)
		vestingBal.Add(vestingBal, miner.Vesting)
	}

	return PreviewAgentTerminationSummary{
		TerminationPenalty: cs.TerminationPenalty,
		InitialPledge:      initialPledge,
		VestingBalance:     vestingBal,
		MinersAvailableBal: availBal,
		AgentAvailableBal:  cs.AvailableBalance,
	}
}
