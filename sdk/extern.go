package sdk

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/glifio/go-pools/rpc"
)

func (c *fevmExtern) ConnectEthClient() (*ethclient.Client, error) {
	return connectEthClient(c.dialAddr, c.token)
}

func connectEthClient(dialAddr string, token string) (*ethclient.Client, error) {
	if token == "" {
		return ethclient.Dial(dialAddr)
	}

	tokenHeader := ethrpc.WithHeader("Authorization", "Bearer "+token)
	httpClient := ethrpc.WithHTTPClient(&http.Client{
		Timeout: 0,
	})

	client, err := ethrpc.DialOptions(context.Background(), dialAddr, httpClient, tokenHeader)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return ethclient.NewClient(client), nil
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

func (c *fevmExtern) GetEventsAPI() string {
	return c.eventsURL
}
