package sdk

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMaxBorrowDTICap(t *testing.T) {
	rate, _ := new(big.Int).SetString("171232876712328767123287671233", 10)
	maxDTI := big.NewInt(25e16)

	var tests = []struct {
		name      string
		edr       *big.Int
		principal *big.Int
		want      *big.Int
	}{
		{"No principal", big.NewInt(100), big.NewInt(0), big.NewInt(50694)},
		{"Existing principal", big.NewInt(100), big.NewInt(10000), big.NewInt(40694)},
		{"No rewards", big.NewInt(0), big.NewInt(0), big.NewInt(0)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			answer := computeMaxDTICap(rate, tt.edr, tt.principal, maxDTI)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestMaxBorrowDTECap(t *testing.T) {
	// MAX_DTE = 2e18
	var tests = []struct {
		name       string
		agentValue *big.Int
		principal  *big.Int
		want       *big.Int
	}{
		{"No principal", big.NewInt(100), big.NewInt(0), big.NewInt(200)},
		{"Existing principal", big.NewInt(100), big.NewInt(50), big.NewInt(150)},
		{"No collateral", big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{"Capped", big.NewInt(100), big.NewInt(200), big.NewInt(0)},
		{"Over", big.NewInt(100), big.NewInt(300), big.NewInt(0)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			answer := computeMaxDTECap(tt.agentValue, tt.principal)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}

func TestMaxBorrowLTVCap(t *testing.T) {
	// MAX_LTV := 80%
	var tests = []struct {
		name             string
		liquidationValue *big.Int
		principal        *big.Int
		recoveryRate     *big.Int
		want             *big.Int
	}{
		{"No principal", big.NewInt(100), big.NewInt(0), big.NewInt(9e17), big.NewInt(286)},
		{"Existing principal", big.NewInt(100), big.NewInt(50), big.NewInt(9e17), big.NewInt(108)},
		{"Existing principal2", big.NewInt(945e3), big.NewInt(750e3), big.NewInt(9e17), big.NewInt(21429)},
		{"No collateral", big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{"Capped", big.NewInt(100), big.NewInt(286), big.NewInt(9e17), big.NewInt(0)},
		{"WAD Math", big.NewInt(1e18), big.NewInt(0), big.NewInt(9e17), big.NewInt(2857142857142857143)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			answer := computeMaxLTVCap(tt.liquidationValue, tt.principal, tt.recoveryRate)
			if answer.Cmp(tt.want) != 0 {
				t.Errorf("Expected %v, received %v", tt.want, answer)
			}
		})
	}
}
