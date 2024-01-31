package util

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
)

func SetupSuite(t *testing.T) (*api.FullNodeStruct, jsonrpc.ClientCloser) {
	lotusDialAddr := os.Getenv("LOTUS_DIAL_ADDR")
	lotusToken := os.Getenv("LOTUS_TOKEN")

	if lotusDialAddr == "" {
		t.Fatal("LOTUS_DIAL_ADDR env var must be set")
	}

	var lcli api.FullNodeStruct = api.FullNodeStruct{}
	head := http.Header{}

	if lotusToken != "" {
		head.Add("Authorization", "Bearer "+lotusToken)
	}

	closer, err := jsonrpc.NewMergeClient(
		context.Background(),
		lotusDialAddr,
		"Filecoin",
		api.GetInternalStructs(&lcli),
		head,
	)
	if err != nil {
		t.Fatal(err)
	}

	networkName, err := lcli.StateNetworkName(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if err := build.UseNetworkBundle(string(networkName)); err != nil {
		t.Fatal(err)
	}

	return &lcli, closer
}

func TeardownSuite(close jsonrpc.ClientCloser) {
	defer close()
}
