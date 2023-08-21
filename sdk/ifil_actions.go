package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) IFILTransfer(ctx context.Context, auth *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iFILTransactor, err := abigen.NewPoolTokenTransactor(a.queries.IFIL(), client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{receiver, amount}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, iFILTransactor.Transfer, "Transfer iFIL")
}

func (a *fevmActions) IFILApprove(ctx context.Context, auth *bind.TransactOpts, spender common.Address, allowance *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iFILTransactor, err := abigen.NewPoolTokenTransactor(a.queries.IFIL(), client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{spender, allowance}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, iFILTransactor.Approve, "Approve iFIL")
}
