package fevm

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/glifio/go-pools/rpc"
)

func (c *fevmExtern) ConnectEthClient() (*ethclient.Client, error) {
	return ethclient.Dial(c.dialAddr)
}

func (c *fevmExtern) ConnectLotusClient() (*lotusapi.FullNodeStruct, jsonrpc.ClientCloser, error) {
	head := http.Header{}

	if c.token != "" {
		head.Add("Authorization", "Bearer "+c.token)
	}

	lapi := &lotusapi.FullNodeStruct{}

	closer, err := jsonrpc.NewMergeClient(
		context.Background(),
		c.dialAddr,
		"Filecoin",
		lotusapi.GetInternalStructs(lapi),
		head,
	)

	if err != nil {
		return nil, nil, err
	}

	return lapi, closer, nil
}

func (c *fevmExtern) ConnectAdoClient(ctx context.Context) (jsonrpc.ClientCloser, error) {
	return rpc.NewADOClient(ctx, c.adoAddr, c.adoNamespace)
}
