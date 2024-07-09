package terminate

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
)

func TestTermPenaltyOnPledge(t *testing.T) {
	ctx := context.Background()

	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	// get ChainHead and lookback 5 epochs
	/*
		head, err := lapi.ChainHead(context.Background())
		if err != nil {
			t.Fatal(err)
		}
	*/

	height := abi.ChainEpoch(4072417)
	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), height, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	filToPledge := big.NewInt(2 * 1e18)
	var sectorSize uint64 = 34359738368
	var ratioVerified float64 = 1.0
	cost, penalty, err := TermPenaltyOnPledge(ctx, lapi, ts, filToPledge, sectorSize, ratioVerified)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Jim cost: %v\n", cost)
	fmt.Printf("Jim penalty: %v\n", penalty)
	t.Fatal("Jim2")
}
