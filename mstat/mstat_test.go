package mstat

import (
	"context"
	"math/big"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
)

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
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

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

	_, err = ComputeEDRLazy1(context.Background(), minerTest, ts, lapi)
	if err != nil {
		t.Fatal(err)
	}

}

func TestComputeMinerStats(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

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

	bal, err := lapi.StateGetActor(context.Background(), minerTest, ts.Key())
	if err != nil {
		t.Fatal(err)
	}

	minerStats, err := ComputeMinerStats(context.Background(), minerTest, ts, lapi)
	if err != nil {
		t.Fatal(err)
	}

	if minerStats.Balance.Cmp(bal.Balance.Int) != 0 {
		t.Errorf("incorrect balance, expected %v got %v\n", bal.Balance.Int, minerStats.Balance)
	}
}

func TestComputeMinersStats(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	minerTest := toAddressType("f0501283")
	minerTest2 := toAddressType("f01660837")

	// just lookback 5 tipsets from HEAD to see if things look right (don't wanna look too close to head)
	chainHead, err := lapi.ChainHead(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), chainHead.Height(), types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	bal, err := lapi.StateGetActor(context.Background(), minerTest, ts.Key())
	if err != nil {
		t.Fatal(err)
	}

	bal2, err := lapi.StateGetActor(context.Background(), minerTest2, ts.Key())
	if err != nil {
		t.Fatal(err)
	}

	minerStats, err := ComputeMinersStats(
		context.Background(),
		[]address.Address{minerTest, minerTest2},
		ts,
		lapi,
	)
	if err != nil {
		t.Fatal(err)
	}

	summedBal := new(big.Int).Add(bal.Balance.Int, bal2.Balance.Int)

	if minerStats.Balance.Cmp(summedBal) != 0 {
		t.Errorf("incorrect balance, expected %v got %v\n", bal.Balance.Int, minerStats.Balance)
	}
}

func TestComputeMinersStatsWrongActor(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	// just lookback 5 tipsets from HEAD to see if things look right (don't wanna look too close to head)
	chainHead, err := lapi.ChainHead(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), chainHead.Height(), types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	minerTest := toAddressType("f0501283")
	// this is a random multisig
	minerTest2 := toAddressType("f02202404")
	// this is a fake address
	minerTest3 := toAddressType("f000000000")

	var testCases = []struct {
		minerAddrs []address.Address
		name       string
	}{
		{name: "wrong-actor", minerAddrs: []address.Address{minerTest, minerTest2}},
		{name: "bad-addr", minerAddrs: []address.Address{minerTest, minerTest3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			minerStats, err := ComputeMinersStats(
				context.Background(),
				[]address.Address{minerTest, minerTest2},
				ts,
				lapi,
			)
			if err == nil {
				t.Fatal(err)
			}

			if minerStats != nil {
				t.Fatal("miner stats should be nil on err")
			}
		})
	}
}

// func testMinerStats(t *testing.T) {
// 	lapi, closer := util.SetupSuite(t)
// 	defer util.TeardownSuite(closer)

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
