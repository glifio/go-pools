package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) PlusMint(ctx context.Context, auth *bind.TransactOpts) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewPlusTransactor(a.queries.Plus(), client)
	if err != nil {
		return nil, err
	}

	personalCashBackPercent := big.NewInt(0)
	tx, err := plus.Mint(auth, auth.From, personalCashBackPercent)

	return util.TxPostProcess(tx, err)
}
