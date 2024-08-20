package terminate

import (
	"context"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/constants"
)

func Margin(balance *big.Int, termPenalty *big.Int) (*big.Int, *big.Int, *big.Int) {
	margin := big.NewInt(0).Sub(balance, termPenalty)

	// calculate max borrow and seal as (margin * 1e18 / (1e18-borrowDTL)) - margin
	maxBorrowAndSeal := big.NewInt(0).Sub(
		big.NewInt(0).Div(
			big.NewInt(0).Mul(margin, big.NewInt(1e18)),
			big.NewInt(0).Sub(constants.WAD, constants.MAX_BORROW_DTL),
		),
		margin,
	)

	// calculate max borrow and withdraw as margin - ((margin * 25e16) / 1e18)
	maxBorrowAndWithdraw := big.NewInt(0).Sub(margin, big.NewInt(0).Div(big.NewInt(0).Mul(margin, big.NewInt(25e16)), big.NewInt(1e18)))

	return margin, maxBorrowAndSeal, maxBorrowAndWithdraw
}

func GetMinerMargin(ctx context.Context, api *api.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*MinerMargin, error) {
	termRes, err := EstimateTerminationPenalty(context.Background(), api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	margin, maxBorrowAndSeal, maxBorrowAndWithdraw := Margin(termRes.TotalBalance, termRes.EstimatedTerminationFee)

	baseMargin := BaseMargin{
		AvailableBalance:     termRes.AvailableBalance,
		LockedRewards:        termRes.VestingFunds,
		InitialPledge:        termRes.InitialPledge,
		Margin:               margin,
		TerminationPenalty:   termRes.EstimatedTerminationFee,
		MaxBorrowAndSeal:     maxBorrowAndSeal,
		MaxBorrowAndWithdraw: maxBorrowAndWithdraw,
	}

	return &MinerMargin{
		BaseMargin:   baseMargin,
		MinerAddr:    minerAddr,
		MinerBalance: termRes.TotalBalance,
	}, nil
}
