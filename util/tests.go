package util

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"reflect"
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

func AssertApproxEqAbs(a, b, DUST *big.Int) bool {
	diff := new(big.Int).Sub(a, b)
	diff.Abs(diff)
	return diff.Cmp(DUST) <= 0
}

// DIFF is a WAD math based percentage, such that 1e18 is 100%
func AssertApproxEqRel(a, b, DIFF *big.Int) bool {
	// compute the diff
	diff := new(big.Int).Sub(a, b)
	diff.Abs(diff)

	// Calculate the difference in terms of percentage: (|a - b| / a) * 1e18
	// To avoid losing precision, first multiply diff by 1e18, then divide by a
	percentageDiff := new(big.Int).Mul(diff, big.NewInt(1e18))
	percentageDiff.Div(percentageDiff, a)

	// Check if the calculated percentage difference is within the specified range
	return percentageDiff.Cmp(DIFF) <= 0
}

func PrintStructKeysAndValues(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
