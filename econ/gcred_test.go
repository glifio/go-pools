package econ

import (
	"fmt"
	"math/big"
	"testing"
)

func TestDTICreditScore(t *testing.T) {
	var tests = []struct {
		name string
		dti  *big.Float
		want *big.Int
	}{
		{"poor", big.NewFloat(0.75), big.NewInt(4)},
		{"fair", big.NewFloat(0.45), big.NewInt(6)},
		{"good", big.NewFloat(0.25), big.NewInt(8)},
		{"excellent", big.NewFloat(0.19), big.NewInt(10)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s:dti(%s)=>score(%s)", tt.name, tt.dti.String(), tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := DebtToIncomeScore(tt.dti)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

// CV >= EAD: Score of 10
// CV is 75% to 99% of the loan amount: Score of 8
// CV is 50% to 74% of the loan amount: Score of 6
// CV is less than 50% of the loan amount: Score of 4
func TestCollateralCreditScore(t *testing.T) {

	var tests = []struct {
		name string
		ead  *big.Int
		av   *big.Int
		cv   *big.Int
		want *big.Int
	}{
		{"poor", big.NewInt(1100), big.NewInt(1000), big.NewInt(600), big.NewInt(4)},
		{"fair", big.NewInt(800), big.NewInt(1000), big.NewInt(600), big.NewInt(6)},
		{"good", big.NewInt(600), big.NewInt(1000), big.NewInt(600), big.NewInt(8)},
		{"excellent", big.NewInt(400), big.NewInt(1000), big.NewInt(600), big.NewInt(10)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s:ead(%s)av(%s)cv(%d)=>score(%s)", tt.name, tt.ead, tt.av, tt.cv, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			ltv := LoanToValueRatio(tt.ead, tt.av)
			ltc := LoanToCollateralRatio(tt.ead, tt.cv)
			answer := LoanToCollateralScore(tt.ead, ltv, ltc)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestInterestRate(t *testing.T) {

	var tests = []struct {
		name  string
		gcred *big.Int
		want  *big.Float
	}{
		{"poor", big.NewInt(40), big.NewFloat(0.3689404)},
		{"fair", big.NewInt(60), big.NewFloat(0.273318)},
		{"good", big.NewInt(80), big.NewFloat(0.202479)},
		{"excellent", big.NewInt(100), big.NewFloat(0.15)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s:gcred(%s)=>rate(%s)", tt.name, tt.gcred, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := InterestRate(tt.gcred)
			if answer.SetPrec(20).Cmp(tt.want.SetPrec(20)) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestCreditScore(t *testing.T) {

	var tests = []struct {
		name    string
		ead     *big.Int
		av      *big.Int
		edr     *big.Int
		dte     *big.Float
		faults  *big.Float
		vesting *big.Float
		want    *big.Int
	}{
		{"poor", big.NewInt(110000), big.NewInt(100000), big.NewInt(100), big.NewFloat(1.01), big.NewFloat(0.1), big.NewFloat(0.001), big.NewInt(40)},
		{"fair", big.NewInt(60000), big.NewInt(100000), big.NewInt(240), big.NewFloat(0.91), big.NewFloat(0.001), big.NewFloat(0.01), big.NewInt(60)},
		{"good", big.NewInt(50000), big.NewInt(100000), big.NewInt(300), big.NewFloat(0.89), big.NewFloat(0.0001), big.NewFloat(0.05), big.NewInt(80)},
		{"excellent", big.NewInt(40000), big.NewInt(100000), big.NewInt(500), big.NewFloat(0.79), big.NewFloat(0), big.NewFloat(0.10), big.NewInt(100)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s:ead(%s)av(%s)edr(%s)=>score(%s)", tt.name, tt.ead, tt.av, tt.edr, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			ltv := LoanToValueRatio(tt.ead, tt.av)
			ltc := LoanToCollateralRatio(tt.ead, CollateralValue(tt.av))
			answer := CreditScore(tt.ead, ltv, ltc, tt.edr, tt.dte, tt.faults, tt.vesting)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}
