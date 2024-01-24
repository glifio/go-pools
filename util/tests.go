package util

import (
	"context"
	"net/http"
	"testing"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
)

var DIAL_ADDR = ""
var TOKEN = ""

func SetupSuite(t *testing.T) (*api.FullNodeStruct, jsonrpc.ClientCloser) {
	if DIAL_ADDR == "" {
		t.Fatal("DIAL_ADDR must be set")
	}

	var lcli api.FullNodeStruct = api.FullNodeStruct{}
	head := http.Header{}

	if TOKEN != "" {
		head.Add("Authorization", "Bearer "+TOKEN)
	}

	closer, err := jsonrpc.NewMergeClient(
		context.Background(),
		DIAL_ADDR,
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
