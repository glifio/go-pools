package terminate

import (
	"math/big"

	"github.com/glifio/go-pools/constants"
)

func (term PreviewAgentTerminationSummary) LiquidationValue() *big.Int {
	// if the termination penalty is greater than the initial pledge,
	// OR if the initial pledge is 0,
	// we report liquidation value as 0
	if term.InitialPledge.Cmp(term.TerminationPenalty) < 0 || term.InitialPledge.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	totalAvail := new(big.Int).Add(term.MinersAvailableBal, term.AgentAvailableBal)
	// first we multiply the available balance by the recovery rate
	// recovery rate = (initial pledge - termination penalty) / initial pledge
	// discounted avail = (available * initial pledge - available * termination penalty) / initial pledge
	availTimesPledge := new(big.Int).Mul(totalAvail, term.InitialPledge)
	availTimesTermPenalty := new(big.Int).Mul(totalAvail, term.TerminationPenalty)

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
