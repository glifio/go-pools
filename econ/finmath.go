package econ

import (
	"math/big"

	"github.com/glifio/go-pools/mstat"
)

// EL = EAD * PD% * LGD%
func ExpectedLoss(ead *big.Int, pd *big.Float, lgd *big.Float) *big.Int {
	fel := big.NewFloat(0)
	fel.Mul(new(big.Float).SetInt(ead), pd)
	fel.Mul(fel, lgd)
	el, _ := fel.Int(nil)
	// can't loose more than EAD
	if el.Cmp(ead) > 0 {
		return ead
	}
	return el
}

// CV = AV-TP
// AV Agent Value
// TP Termination Penalty
// LGD = 1 - (CV / EAD)
// LGD (in dollars) = Exposure at Risk (EAD) * (1 - Recovery Rate)
// LGD (as percentage) = 1 - (Agent Collateral Value / Outstanding Debt)
// LGD is loss given default and refers to the amount of money a bank loses
// when a borrower defaults on a loan. (Percentage)
func LossGivenDefault(cv *big.Int, ead *big.Int) *big.Float {
	prec := uint(200)
	lgd := big.NewFloat(0).SetPrec(prec)
	// no exposure -> no loss (divide by zero protection)
	if ead.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}
	// cv >= ead -> undercollateralized no loss
	if cv.Cmp(ead) >= 0 {
		return big.NewFloat(0)
	}
	// cv / ead
	lgd = lgd.Quo(new(big.Float).SetPrec(prec).SetInt(cv), new(big.Float).SetPrec(prec).SetInt(ead))

	// 1 - (cd/ead)
	return lgd.Sub(big.NewFloat(1).SetPrec(prec), lgd).SetPrec(5).SetMode(big.ToNearestEven)
}

// LTV = EAD / MV (Miner Value)
func LoanToValueRatio(ead *big.Int, mv *big.Int) *big.Float {
	if mv.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(1)
	}
	ltv := big.NewFloat(0)
	//ead.Mul(ead, new(big.Int).Exp(ead, big.NewInt(18), nil))
	ltv.Quo(new(big.Float).SetInt(ead), new(big.Float).SetInt(mv))
	return ltv
}

// LTC = EAD / CV (Collateral Value)
func LoanToCollateralRatio(ead *big.Int, cv *big.Int) *big.Float {
	if ead.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(1)
	}
	ltc := big.NewFloat(0)
	//ead.Mul(ead, new(big.Int).Exp(ead, big.NewInt(18), nil))
	ltc.Quo(new(big.Float).SetInt(ead), new(big.Float).SetInt(cv))
	return ltc
}

// PD = hard coded to 20%
// maybe a function of edr vs payments
func ProbabilityOfDefault() *big.Float {
	pd := big.NewFloat(0.2)
	// pd := new(big.Int).Mul(
	// 	big.NewInt(int64(percent)),
	// 	new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
	// )
	return pd
}

// EAD = current loans principal + new principal
func ExposureAtDefault(current *big.Int, principal *big.Int) *big.Int {
	ead := big.NewInt(0)
	return ead.Add(current, principal)
}

/**
 * @notice Collateral Value (CV)
 * CV = Agent Value * Termination Penalty Ratio
 * The collateral value of an agent is the agent's balance multiplied by the
 * termination penalty ratio. The termination penalty ratio is the estimated percentage
 * of the agent's balance that will be lost if the agent is terminated.
 */
func CollateralValue(av *big.Int) *big.Int {
	// convert the percentage to a big.Rat
	r := new(big.Rat).SetFloat64(mstat.TerminationPenaltyRatio)

	// take product of x * percent
	result := new(big.Int).Mul(av, r.Num())
	return result.Div(result, r.Denom())
}

/*
 * Debt to Income (DTI)
 * DTI = Estimated Daily Payments (EDP) / Estimated Daily Reward (EDR)
 * Debt to income ratio is debt payments divided by income
 */
func DebtToIncome(edp *big.Int, edr *big.Int) *big.Float {
	dti := big.NewFloat(0)
	// no income -> dti = 1 (div by zero protection)
	if edr.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(1)
	}
	// dti = edp / edr
	dti.Quo(new(big.Float).SetInt(edp), new(big.Float).SetInt(edr))
	return dti
}

/*
 * Debt to Equity (DTE)
 * DTE = Pricipal / Equity
 * Debt to equity ratio is debt divided by equity
 */
func DebtToEquityRatio(principal *big.Int, equity *big.Int) *big.Float {
	dte := big.NewFloat(0)
	// no equity -> dte = 1 (div by zero protection)
	if equity.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(1)
	}
	// dti = edp / edr
	dte.Quo(new(big.Float).SetInt(principal), new(big.Float).SetInt(equity))
	return dte
}
