package sdk

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
	return connectLotusClient(c.dialAddr, c.token)
}

func connectLotusClient(lotusDialAddr string, lotusToken string) (*lotusapi.FullNodeStruct, jsonrpc.ClientCloser, error) {
	head := http.Header{}

	if lotusToken != "" {
		head.Add("Authorization", "Bearer "+lotusToken)
	}

	lapi := &lotusapi.FullNodeStruct{}

	closer, err := jsonrpc.NewMergeClient(
		context.Background(),
		lotusDialAddr,
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
