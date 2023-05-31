package econ

import (
	"context"
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/glifio/go-pools/abigen"
)

// @dev rates have 2 WADs worth of precision (1e36) to maintain per epoch rate precision
func InterestOwed(ctx context.Context, account abigen.Account, rate *big.Int, chainHeadHeight abi.ChainEpoch) *big.Int {
	currentEpoch := new(big.Int).SetUint64(uint64(chainHeadHeight))

	interestOwed := big.NewInt(0)
	if account.EpochsPaid.Cmp(currentEpoch) == -1 {
		// Compute the number of epochs that are owed to get current
		epochsToPay := new(big.Int).Sub(currentEpoch, account.EpochsPaid)

		// Multiply the rate by the principal to get the per epoch interest rate, now has an extra WAD of precision after multiplication
		interestPerEpoch := new(big.Int).Mul(account.Principal, rate)

		// div out the extra wad from the multiplication above
		interestPerEpoch.Div(interestPerEpoch, big.NewInt(1e18))

		// Compute the total interest owed by multiplying how many epochs to pay, by the per epoch interest payment
		interestOwed = new(big.Int).Mul(interestPerEpoch, epochsToPay)

		// div out the extra wad in interestOwed from the WAD embedded in `rate` for precision
		interestOwed.Div(interestOwed, big.NewInt(1e18))
	}

	return interestOwed
}
