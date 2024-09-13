package terminate

import (
	"context"
	"math/big"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/util"
)

func TestTermPenaltyOnPledge(t *testing.T) {
	ctx := context.Background()

	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	height := abi.ChainEpoch(4072417)
	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), height, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	filToPledge := new(big.Int).Mul(big.NewInt(1000), constants.WAD)
	var sectorSize uint64 = 34359738368
	var ratioVerified float64 = 1.0
	cost, penalty, err := TermPenaltyOnPledge(ctx, lapi, ts, filToPledge, sectorSize, ratioVerified)
	if err != nil {
		t.Fatal(err)
	}
	expectedCost, _ := new(big.Int).SetString("999985869063292958098", 10)
	if cost.Cmp(expectedCost) != 0 {
		t.Errorf("expected cost %v, got %v", expectedCost, cost)
	}
	expectedPenalty, _ := new(big.Int).SetString("41721371161825193117", 10)
	if penalty.Cmp(expectedPenalty) != 0 {
		t.Errorf("expected penalty %v, got %v", expectedPenalty, penalty)
	}
}
