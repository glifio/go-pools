package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) RampWithdraw(
	ctx context.Context,
	assets *big.Int,
	receiver common.Address,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	rampTransactor, err := abigen.NewSimpleRampTransactor(a.queries.SimpleRamp(), client)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.Address)
	if err != nil {
		return nil, err
	}

	args := []interface{}{assets, receiver, senderAccount.Address, common.Big0}

	return util.WriteTx(ctx, senderWallet, senderAccount, senderPassphrase, a.queries.ChainID(), common.Big0, nonce, args, rampTransactor.WithdrawF, "Withdraw from Ramp")
}

func (a *fevmActions) RampRedeem(
	ctx context.Context,
	shares *big.Int,
	receiver common.Address,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	rampTransactor, err := abigen.NewSimpleRampTransactor(a.queries.SimpleRamp(), client)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.Address)
	if err != nil {
		return nil, err
	}

	args := []interface{}{shares, receiver, senderAccount.Address, common.Big0}

	return util.WriteTx(ctx, senderWallet, senderAccount, senderPassphrase, a.queries.ChainID(), common.Big0, nonce, args, rampTransactor.RedeemF, "Redeem from Ramp")
}
