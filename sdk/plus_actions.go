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

	tx, err := plus.Mint(auth, auth.From)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) PlusMintAndActivate(ctx context.Context, auth *bind.TransactOpts, beneficiary common.Address, tier uint8) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewPlusTransactor(a.queries.Plus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.MintAndActivate(auth, auth.From, beneficiary, tier)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) PlusActivate(ctx context.Context, auth *bind.TransactOpts, beneficiary common.Address, tokenID *big.Int, tier uint8) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewPlusTransactor(a.queries.Plus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.Activate(auth, beneficiary, tokenID, tier)

	return util.TxPostProcess(tx, err)
}
