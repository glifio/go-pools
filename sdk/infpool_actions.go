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

func (a *fevmActions) InfPoolDepositFIL(ctx context.Context, auth *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolTransactor, err := abigen.NewInfinityPoolTransactor(a.queries.InfinityPool(), client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{receiver}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), amount, nil, args, poolTransactor.Deposit0, "Deposit FIL")
}
