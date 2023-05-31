package econ

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/glifio/go-pools/abigen"
)

func TestInterestOwed(t *testing.T) {
	var DUST = big.NewInt(10000)

	// this test makes sure, with known values, that the interest owed is calculated correctly
	// an account with GCRED of 80 should be paying roughly 22% per year

	// pretend the account owes a years worth of epochs
	currentEpoch := abi.ChainEpoch(builtin.EpochsInYear)
	// borrow 1, expect to owe .22
	principal := big.NewInt(1e18)
	expectedOwed := big.NewInt(22e16)

	testAccount := abigen.Account{
		Principal:  principal,
		EpochsPaid: common.Big0,
		StartEpoch: common.Big0,
		Defaulted:  false,
	}

	// this is the known per epoch rate for a GCRED of 80, denominated in attofil
	rate, _ := new(big.Int).SetString("209284627092846270890410958904", 10)

	owed := InterestOwed(context.Background(), testAccount, rate, currentEpoch)

	if !assertApproxEqAbs(owed, expectedOwed, DUST) {
		t.Errorf("expected %v, got %v", expectedOwed, owed)
	}
}

func assertApproxEqAbs(a, b, DUST *big.Int) bool {
	diff := new(big.Int).Sub(a, b)
	diff.Abs(diff)
	return diff.Cmp(DUST) <= 0
}
