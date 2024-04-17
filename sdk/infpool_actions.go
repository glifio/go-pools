package sdk

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
	"github.com/jimpick/go-ethereum/accounts/abi/bind"
	"github.com/jimpick/go-ethereum/common"
	"github.com/jimpick/go-ethereum/core/types"
)

func (a *fevmActions) InfPoolDepositFIL(ctx context.Context, auth *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	infpool, err := abigen.NewInfinityPoolTransactor(a.queries.InfinityPool(), client)
	if err != nil {
		return nil, err
	}

	auth.Value = amount
	tx, err := infpool.Deposit0(auth, receiver)

	return util.TxPostProcess(tx, err)
}
