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
	var tests = []struct {
		name       string
		agentValue *big.Int
		principal  *big.Int
		want       *big.Int
	}{
		{"No principal", big.NewInt(100), big.NewInt(0), big.NewInt(100)},
		{"Existing principal", big.NewInt(100), big.NewInt(50), big.NewInt(50)},
		{"No collateral", big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{"Capped", big.NewInt(100), big.NewInt(100), big.NewInt(0)},
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

// given current LTV ratios, LTV and DTE should produce the same results
func TestMaxBorrowLTVCap(t *testing.T) {
	var tests = []struct {
		name       string
		agentValue *big.Int
		principal  *big.Int
		want       *big.Int
	}{
		{"No principal", big.NewInt(100), big.NewInt(0), big.NewInt(100)},
		{"Existing principal", big.NewInt(100), big.NewInt(50), big.NewInt(50)},
		{"No collateral", big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{"Capped", big.NewInt(100), big.NewInt(100), big.NewInt(0)},
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
