package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/vc"
)

var ADOClient struct {
	SignCredential func(ctx context.Context, jws string) (abigen.SignedCredential, error)
	AgentData      func(ctx context.Context, agentAddr common.Address) (*vc.AgentData, error)
	PreviewAction  func(ctx context.Context, agentAddr common.Address, target address.Address, value *big.Int, method constants.Method) (*vc.AgentData, error)
}

func NewADOClient(ctx context.Context, rpcurl string, namespace string) (jsonrpc.ClientCloser, error) {
	return jsonrpc.NewClient(ctx, rpcurl, namespace, &ADOClient, nil)
}
