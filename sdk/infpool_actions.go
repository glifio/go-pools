package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
	"github.com/glifio/go-wallet-utils/accounts"
)

func (a *fevmActions) InfPoolDepositFIL(
	ctx context.Context,
	receiver common.Address,
	amount *big.Int,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
	proposer address.Address,
	approver address.Address,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	args := []interface{}{receiver}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		senderWallet,
		senderAccount,
		senderPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		amount,
		args,
		abigen.NewInfinityPoolTransactor,
		a.queries.InfinityPool(),
		"Deposit0",
		"Deposit FIL",
	)
}
