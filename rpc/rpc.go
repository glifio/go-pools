package rpc

import (
	"context"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/glifio/go-pools/abigen"
)

var ADOClient struct {
	SignCredential func(ctx context.Context, jws string) (abigen.SignedCredential, error)
}

func NewADOClient(ctx context.Context, rpcurl string, namespace string) (jsonrpc.ClientCloser, error) {
	return jsonrpc.NewClient(ctx, rpcurl, namespace, &ADOClient, nil)
}
