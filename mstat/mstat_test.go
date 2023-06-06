package mstat

import (
	"context"
	"math/big"
	"net/http"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

var DIAL_ADDR = ""
var TOKEN = ""

func setupSuite(t *testing.T) (api.FullNodeStruct, jsonrpc.ClientCloser) {
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

	return lcli, closer
}

func teardownSuite(close jsonrpc.ClientCloser) {
	defer close()
}

func TestComputePercentage(t *testing.T) {
	// Since we know the percentage is 50%, the results of these tests should always equal original num / 2

	var testCases = []struct {
		originalNum *big.Int
		expRes      *big.Int
	}{
		{big.NewInt(100), big.NewInt(50)},
		{big.NewInt(200), big.NewInt(100)},
		{big.NewInt(1), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(2), big.NewInt(1)},
		{big.NewInt(3), big.NewInt(1)},
		{big.NewInt(4), big.NewInt(2)},
		{big.NewInt(5), big.NewInt(2)},
		{big.NewInt(100000000000000), big.NewInt(50000000000000)},
	}

	for _, tc := range testCases {
		t.Run(tc.originalNum.String(), func(t *testing.T) {
			answer := computePercentOf(tc.originalNum)
			if answer.Cmp(tc.expRes) != 0 {
				t.Errorf("expected %s, got %s", tc.expRes, answer)
			}
		})
	}
}

func TestCmpWithOnChainSectorInfo(t *testing.T) {
	lapi, closer := setupSuite(t)
	defer teardownSuite(closer)

	minerTest := toAddressType("f0501283")

	// just lookback 5 tipsets from HEAD to see if things look right (don't wanna look too close to head)
	chainHead, err := lapi.ChainHead(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), chainHead.Height(), types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	edrPrecise, err := ComputeEDRPrecise(context.Background(), minerTest, ts, lapi)
	if err != nil {
		t.Fatal(err)
	}

	edrLazy1, err := ComputeEDRLazy1(context.Background(), minerTest, ts, lapi)
	if err != nil {
		t.Fatal(err)
	}

	edrLazy2, err := ComputeEDRLazy2(context.Background(), minerTest, ts, lapi)
	if err != nil {
		t.Fatal(err)
	}

	if edrPrecise.Cmp(edrLazy1) != 0 {
		t.Errorf("Precise vs lazy1 fail - expected %s, got %s", edrPrecise, edrLazy1.String())
	}

	if edrPrecise.Cmp(edrLazy2) != 0 {
		t.Errorf("Precise vs lazy2 fail - expected %s, got %s", edrPrecise, edrLazy2.String())
	}

	if edrLazy1.Cmp(edrLazy2) != 0 {
		t.Errorf("lazy1 vs laz2 fail - expected %s, got %s", edrLazy1, edrLazy2.String())
	}

}

// func testMinerStats(t *testing.T) {
// 	lapi, closer := setupSuite(t)
// 	defer teardownSuite(closer)

// 	var testCases = []struct {
// 		addr            address.Address
// 		height          abi.ChainEpoch
// 		expDailyRewards *big.Int
// 		expPenaltyTerm  *big.Int
// 	}{
// 		{toAddressType("f0501283"), 2845750, bigInt("30720941433723887280"), bigInt("2744389106752621606362")},
// 		{toAddressType("f01938357"), 2845753, bigInt("223177646823924333252"), bigInt("18079621691476321790913")},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.addr.String(), func(t *testing.T) {
// 			ts, err := lapi.ChainGetTipSetByHeight(context.Background(), tc.height, types.EmptyTSK)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			minerStats, err := Compute(context.Background(), tc.addr, ts, lapi)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			if minerStats.ExpectedDailyReward.Cmp(tc.expDailyRewards) != 0 {
// 				t.Errorf("expected %s, got %s", tc.expDailyRewards, minerStats.ExpectedDailyReward)
// 			}
// 		})
// 	}
// }

func toAddressType(addr string) address.Address {
	address, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}

	return address
}

func bigInt(val string) *big.Int {
	bigVal, ok := new(big.Int).SetString(val, 10)
	if !ok {
		panic("could not parse big int")
	}

	return bigVal
}
