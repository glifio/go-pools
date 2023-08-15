package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
	"github.com/glifio/go-wallet-utils/accounts"
)

func (a *fevmActions) IFILTransfer(
	ctx context.Context,
	receiver common.Address,
	amount *big.Int,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
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

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{receiver, amount}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		senderWallet,
		senderAccount,
		senderPassphrase,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewPoolTokenTransactor,
		a.queries.IFIL(),
		"Transfer",
		"Transfer iFIL",
	)
}

func (a *fevmActions) IFILApprove(
	ctx context.Context,
	spender common.Address,
	allowance *big.Int,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
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

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{spender, allowance}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		senderWallet,
		senderAccount,
		senderPassphrase,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewPoolTokenTransactor,
		a.queries.IFIL(),
		"Approve",
		"Approve iFIL",
	)
}
