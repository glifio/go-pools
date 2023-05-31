package econ

import (
	"log"
	"math/big"

	"github.com/ALTree/bigfloat"
)

// ead * rate / 365
func EstimatedDailyPayment(ead *big.Int, rate *big.Float) *big.Int {
	edp := big.NewFloat(0)
	days := big.NewFloat(365)

	edp = edp.Mul(big.NewFloat(0).SetInt(ead), big.NewFloat(0).Quo(rate, days))

	res, _ := edp.Int(nil)

	return res
}

/*
* Interest Rate
* Interest Rate = 15 * e^(0.015 * (100 - CreditScore))
 */
func InterestRate(gcred *big.Int) *big.Float {
	baserate := big.NewFloat(0.15)
	bias := big.NewFloat(0.015)
	maxScore := big.NewFloat(100)

	e := bigfloat.Exp(big.NewFloat(0).Mul(bias, big.NewFloat(0).Sub(maxScore, big.NewFloat(0).SetInt(gcred))))
	rate := big.NewFloat(0).Mul(baserate, e)
	return rate
}

/*
* Credit Score
* Credit Score = (DTE * 30%) - (LTC * 30%) - (FaultRatio * 20%) - (VestingRatio * 20%)
* Score = 90-100 is excellent
* Score = 80-90 is good
* Score = 60-80 is fair
* Score = 40-60 is poor
 */
func CreditScore(ead *big.Int, ltv *big.Float, ltc *big.Float, edr *big.Int, dte *big.Float, faultRatio *big.Float, vestingToPledgeRatio *big.Float) *big.Int {
	score := big.NewInt(0)

	// Calculate the weighted score for Debt to Equity Ratio
	// dte score count for 50% of the credit score
	dteScore := DebtToEquityScore(dte)
	log.Printf("DTE Score: %v", dteScore)
	weightedDTE := big.NewInt(0).Mul(dteScore, big.NewInt(3))

	// Calculate the weighted score for Loan to Collateral Ratio
	// ltc score count for 30% of the credit score
	ltcScore := LoanToCollateralScore(ead, ltv, ltc)
	log.Printf("LTC Score: %v", ltcScore)
	weightedLTC := big.NewInt(0).Mul(ltcScore, big.NewInt(3))

	// Calculate the weighted score for Fault ratio
	// fault score count for 20% of the credit score
	faultScore := FaultScore(faultRatio)
	log.Printf("Fault Score: %v", faultScore)
	weightedFault := big.NewInt(0).Mul(faultScore, big.NewInt(2))

	// Calculate the weighted score for Vesting to Pledge Ratio
	// vesting score count for 10% of the credit score
	vestingScore := VestingScore(vestingToPledgeRatio)
	log.Printf("Vesting Score: %v", vestingScore)
	weightedVesting := big.NewInt(0).Mul(vestingScore, big.NewInt(2))

	// Add all the weighted scores together
	score = big.NewInt(0).Add(score, weightedLTC)
	score = big.NewInt(0).Add(score, weightedDTE)
	score = big.NewInt(0).Add(score, weightedFault)
	score = big.NewInt(0).Add(score, weightedVesting)
	return score
}

/*
* Vesting to Pledge Credit Score
* Vesting to Pledge Ratio > 10%: Score of 10
* Vesting to Pledge Ratio between 5% and 10%: Score of 8
* Vesting to Pledge Ratio between 1% and 5%: Score of 6
* Vesting to Pledge Ratio < 1%: Score of 4
 */
func VestingScore(vestingToPledgeRatio *big.Float) *big.Int {
	switch {
	case vestingToPledgeRatio.Cmp(big.NewFloat(0.10)) >= 0:
		return big.NewInt(10)
	case vestingToPledgeRatio.Cmp(big.NewFloat(0.05)) >= 0:
		return big.NewInt(8)
	case vestingToPledgeRatio.Cmp(big.NewFloat(0.01)) >= 0:
		return big.NewInt(6)
	case vestingToPledgeRatio.Cmp(big.NewFloat(0.01)) < 0:
		return big.NewInt(4)
	default:
		return big.NewInt(0)
	}
}

/*
* Fault Ratio Credit Score
 */
func FaultScore(faultRatio *big.Float) *big.Int {
	switch {
	// less than 10 faults per million (0.001%)
	case faultRatio.Cmp(big.NewFloat(0.00001)) <= 0:
		return big.NewInt(10)
	// less than 100 faults per million (0.01%)
	case faultRatio.Cmp(big.NewFloat(0.0001)) <= 0:
		return big.NewInt(8)
	// less than 1000 faults per million (0.1%)
	case faultRatio.Cmp(big.NewFloat(0.001)) <= 0:
		return big.NewInt(6)
	// more than 1000 faults per million (0.1%)
	case faultRatio.Cmp(big.NewFloat(0.001)) > 0:
		return big.NewInt(4)
	default:
		return big.NewInt(0)
	}
}

/*
* Debt to Income Credit Score
* DTI < 20%: Score of 10
* DTI between 20% and 35%: Score of 8
* DTI between 35% and 50%: Score of 6
* DTI > 50%: Score of 4
 */
func DebtToIncomeScore(dti *big.Float) *big.Int {
	switch {
	case dti.Cmp(big.NewFloat(0)) >= 0 && dti.Cmp(big.NewFloat(0.20)) <= 0:
		return big.NewInt(10)
	case dti.Cmp(big.NewFloat(0.20)) == 1 && dti.Cmp(big.NewFloat(0.35)) <= 0:
		return big.NewInt(8)
	case dti.Cmp(big.NewFloat(0.35)) == 1 && dti.Cmp(big.NewFloat(0.50)) <= 0:
		return big.NewInt(6)
	case dti.Cmp(big.NewFloat(0.50)) == 1:
		return big.NewInt(4)
	default:
		return big.NewInt(0)
	}
}

/*
* Loan to Collateral Credit Score
* LTC < 80%: Score of 10
* LTC between 80% and 100%: Score of 8
* LTV between 80% and 100%: Score of 6
* LTV > 100%: Score of 4
 */
func LoanToCollateralScore(ead *big.Int, ltv *big.Float, ltc *big.Float) *big.Int {
	switch {
	case ltc.Cmp(big.NewFloat(0.8)) <= 0:
		return big.NewInt(10)
	case ltc.Cmp(big.NewFloat(1)) <= 0:
		return big.NewInt(8)
	case ltc.Cmp(big.NewFloat(1)) > 0 && ltv.Cmp(big.NewFloat(0.80)) <= 0:
		return big.NewInt(6)
	case ltc.Cmp(big.NewFloat(1)) > 0 && ltv.Cmp(big.NewFloat(0.80)) == 1:
		return big.NewInt(4)
	default:
		return big.NewInt(0)
	}
}

/*
* Debt to Equity Credit Score
* LTE < 80%: Score of 10
* LTE between 80% and 90%: Score of 8
* LTE between 90% and 100%: Score of 6
* LTE > 100%: Score of 4
 */
func DebtToEquityScore(dte *big.Float) *big.Int {
	switch {
	case dte.Cmp(big.NewFloat(0.80)) <= 0:
		return big.NewInt(10)
	case dte.Cmp(big.NewFloat(0.80)) > 0 && dte.Cmp(big.NewFloat(0.90)) <= 0:
		return big.NewInt(8)
	case dte.Cmp(big.NewFloat(0.90)) > 0 && dte.Cmp(big.NewFloat(1.00)) <= 0:
		return big.NewInt(6)
	case dte.Cmp(big.NewFloat(1.00)) > 0:
		return big.NewInt(4)
	default:
		return big.NewInt(0)
	}
}
