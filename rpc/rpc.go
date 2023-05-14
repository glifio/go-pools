package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/types"
)

var ADOClient struct {
	Borrow      func(context.Context, common.Address, *big.Int) (abigen.SignedCredential, error)
	Pay         func(context.Context, common.Address, *big.Int) (abigen.SignedCredential, error)
	Withdraw    func(context.Context, common.Address, *big.Int) (abigen.SignedCredential, error)
	PushFunds   func(context.Context, common.Address, *big.Int, address.Address) (abigen.SignedCredential, error)
	PullFunds   func(context.Context, common.Address, *big.Int, address.Address) (abigen.SignedCredential, error)
	AddMiner    func(context.Context, common.Address, address.Address) (abigen.SignedCredential, error)
	RemoveMiner func(context.Context, common.Address, address.Address) (abigen.SignedCredential, error)
	AmountOwed  func(context.Context, common.Address, *big.Int) (types.AgentOwed, error)
}

func NewADOClient(ctx context.Context, rpcurl string, namespace string) (jsonrpc.ClientCloser, error) {
	return jsonrpc.NewClient(ctx, rpcurl, namespace, &ADOClient, nil)
}
