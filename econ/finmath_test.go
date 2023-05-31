package econ

import (
	"fmt"
	"math/big"
	"testing"
)

func TestEL(t *testing.T) {
	var tests = []struct {
		name string
		ead  *big.Int
		pd   *big.Float
		lgd  *big.Float
		want *big.Int
	}{
		{"normal", big.NewInt(20000), big.NewFloat(0.2), big.NewFloat(0.4), big.NewInt(1600)},
		{"normal", big.NewInt(20000), big.NewFloat(0.2), big.NewFloat(1.0), big.NewInt(4000)},
		{"normal", big.NewInt(20000), big.NewFloat(0.1), big.NewFloat(0.4), big.NewInt(800)},
		{"no loan", big.NewInt(0), big.NewFloat(0.1), big.NewFloat(0.4), big.NewInt(0)},
		{"normal", big.NewInt(0), big.NewFloat(0), big.NewFloat(0), big.NewInt(0)},
		{"normal", big.NewInt(100), big.NewFloat(2000), big.NewFloat(10000), big.NewInt(100)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %s %s = %s", tt.name, tt.ead, tt.pd.String(), tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := ExpectedLoss(tt.ead, tt.pd, tt.lgd)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestLGD(t *testing.T) {
	var tests = []struct {
		name string
		cv   *big.Int
		ead  *big.Int
		want *big.Float
	}{
		{"10% loss", big.NewInt(800), big.NewInt(1000), big.NewFloat(0.2).SetPrec(5).SetMode(big.ToNearestEven)},
		{"2.25% loss", big.NewInt(125), big.NewInt(1000), big.NewFloat(0.875).SetPrec(5).SetMode(big.ToNearestEven)},
		{"total loss (div by zero)", big.NewInt(0), big.NewInt(999), big.NewFloat(1)},
		{"minimal loss", big.NewInt(1000000), big.NewInt(1000001), big.NewFloat(0.000001).SetPrec(6).SetMode(big.ToNearestEven)},
		{"cv greater than loan", big.NewInt(1001), big.NewInt(1000), big.NewFloat(0)},
		{"100% loss", big.NewInt(1), big.NewInt(100000000), big.NewFloat(1)},
		{"no loan no loss", big.NewInt(12234), big.NewInt(0), big.NewFloat(0)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %s / %s = %s", tt.name, tt.cv, tt.ead, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := LossGivenDefault(tt.cv, tt.ead)
			if answer.SetPrec(60).SetMode(big.ToNearestEven).Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestLTV(t *testing.T) {
	var tests = []struct {
		name string
		mv   *big.Int
		ead  *big.Int
		want *big.Float
	}{
		{"10%", big.NewInt(1000), big.NewInt(100), big.NewFloat(0.1)},
		{"2.25%", big.NewInt(4000), big.NewInt(100), big.NewFloat(0.025)},
		{"no value (div by zero)", big.NewInt(0), big.NewInt(100), big.NewFloat(1)},
		{"minimal", big.NewInt(1000000), big.NewInt(1), big.NewFloat(0.000001)},
		{"loan greater than value", big.NewInt(100), big.NewInt(1000), big.NewFloat(10)},
		{"100%", big.NewInt(100), big.NewInt(100), big.NewFloat(1)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %s %s = %s", tt.name, tt.ead, tt.mv, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := LoanToValueRatio(tt.ead, tt.mv)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestPD(t *testing.T) {
	var tests = []struct {
		name string
		want *big.Float
	}{
		{"20%", big.NewFloat(0.2)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s = %s", tt.name, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := ProbabilityOfDefault()
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestEAD(t *testing.T) {
	var tests = []struct {
		name      string
		ead       *big.Int
		principal *big.Int
		want      *big.Int
	}{
		{"", big.NewInt(1000), big.NewInt(100), big.NewInt(1100)},
		{"", big.NewInt(4000), big.NewInt(100), big.NewInt(4100)},
		{"", big.NewInt(0), big.NewInt(100), big.NewInt(100)},
		{"", big.NewInt(1000000), big.NewInt(1), big.NewInt(1000001)},
		{"", big.NewInt(100), big.NewInt(1000), big.NewInt(1100)},
		{"", big.NewInt(0), big.NewInt(0), big.NewInt(0)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %s %s = %s", tt.name, tt.ead, tt.principal, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := ExposureAtDefault(tt.ead, tt.principal)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestDTI(t *testing.T) {
	var tests = []struct {
		name string
		edp  *big.Int
		edr  *big.Int
		want *big.Float
	}{
		{"40%", big.NewInt(40), big.NewInt(100), big.NewFloat(0.4)},
		{"2.25%", big.NewInt(100), big.NewInt(4000), big.NewFloat(0.025)},
		{"no value (div by zero)", big.NewInt(0), big.NewInt(100), big.NewFloat(0)},
		{"minimal", big.NewInt(1), big.NewInt(1000000), big.NewFloat(0.000001)},
		{"payment greater than income", big.NewInt(1200), big.NewInt(1000), big.NewFloat(1.2)},
		{"100%", big.NewInt(100), big.NewInt(100), big.NewFloat(1)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s %s %s = %s", tt.name, tt.edp, tt.edr, tt.want.String())
		t.Run(testname, func(t *testing.T) {
			answer := DebtToIncome(tt.edp, tt.edr)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}
