package sdk

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	token "github.com/glifio/go-pools/jws"
	"github.com/glifio/go-pools/rpc"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) SPPlusMint(ctx context.Context, auth *bind.TransactOpts) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.Mint(auth, auth.From)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusMintAndActivate(ctx context.Context, auth *bind.TransactOpts, beneficiary common.Address, tier uint8) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.MintAndActivate(auth, beneficiary, tier)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusActivate(ctx context.Context, auth *bind.TransactOpts, beneficiary common.Address, tokenID *big.Int, tier uint8) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.Activate(auth, beneficiary, tokenID, tier)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusMintActivateAndFund(ctx context.Context, auth *bind.TransactOpts, cashBackPercent *big.Int, beneficiary common.Address, tier uint8, amount *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.MintActivateAndFund(auth, cashBackPercent, beneficiary, tier, amount)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusSetPersonalCashBackPercent(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, cashBackPercent *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.SetPersonalCashBackPercent(auth, tokenID, cashBackPercent)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusFundGLFVault(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, amount *big.Int, cashBackPercent *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	if cashBackPercent == nil {
		caller, err := abigen.NewSPPlusCaller(a.queries.SPPlus(), client)
		if err != nil {
			return nil, err
		}

		opts := &bind.CallOpts{Context: ctx}

		cashBackPercent, err = caller.MaxCashBackPercent(opts)
		if err != nil {
			return nil, err
		}
	}

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.FundGlfVault0(auth, tokenID, amount, cashBackPercent)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusClaimCashBack(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, receiver common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.ClaimCashBack(auth, tokenID, receiver)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusUpgrade(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, tier uint8) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.Upgrade(auth, tokenID, tier)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusDowngrade(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, tier uint8, agentAddr common.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, address.Undef, big.NewInt(0), constants.MethodSPPlusDowngrade, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.Downgrade(auth, tokenID, tier, sc)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusWithdrawExtraLockedFunds(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.WithdrawExtraLockedFunds(auth, tokenID)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) SPPlusWithdrawGlfVault(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewSPPlusTransactor(a.queries.SPPlus(), client)
	if err != nil {
		return nil, err
	}

	tx, err := plus.WithdrawGlfVault(auth, tokenID, amount, receiver)

	return util.TxPostProcess(tx, err)
}
