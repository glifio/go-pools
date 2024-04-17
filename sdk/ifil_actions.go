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

func (a *fevmActions) IFILTransfer(ctx context.Context, auth *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ifil, err := abigen.NewPoolTokenTransactor(a.queries.IFIL(), client)
	if err != nil {
		return nil, err
	}

	tx, err := ifil.Transfer(auth, receiver, amount)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) IFILApprove(ctx context.Context, auth *bind.TransactOpts, spender common.Address, allowance *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ifil, err := abigen.NewPoolTokenTransactor(a.queries.IFIL(), client)
	if err != nil {
		return nil, err
	}

	tx, err := ifil.Approve(auth, spender, allowance)

	return util.TxPostProcess(tx, err)
}
