package terminate

import (
	"math/big"

	"github.com/glifio/go-pools/constants"
)

func (ats PreviewAgentTerminationSummary) LiquidationValue() *big.Int {
	// if the termination penalty is greater than the initial pledge,
	// OR if the initial pledge is 0,
	// we report liquidation value as 0
	if ats.InitialPledge.Cmp(ats.TerminationPenalty) < 0 || ats.InitialPledge.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	totalAvail := new(big.Int).Add(ats.MinersAvailableBal, ats.AgentAvailableBal)
	// first we multiply the available balance by the recovery rate
	// recovery rate = (initial pledge - termination penalty) / initial pledge
	// discounted avail = (available * initial pledge - available * termination penalty) / initial pledge
	availTimesPledge := new(big.Int).Mul(totalAvail, ats.InitialPledge)
	availTimesTermPenalty := new(big.Int).Mul(totalAvail, ats.TerminationPenalty)

	discountedAvail := new(big.Int).Sub(availTimesPledge, availTimesTermPenalty)
	discountedAvail.Div(discountedAvail, ats.InitialPledge)

	// we add the discounted available balance to the vesting balance and initial pledge, subtract the termination penalty to get the liquidation value
	liquidationValue := new(big.Int).Add(discountedAvail, ats.VestingBalance)
	liquidationValue.Add(liquidationValue, ats.InitialPledge)
	liquidationValue.Sub(liquidationValue, ats.TerminationPenalty)

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
