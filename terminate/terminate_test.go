package terminate

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
)

// this test compares N random miner termination penalties computed in the most precise (time intensive) way against the quick, less precise sampling method used in the ADO
var N = 10

func TestTerminationPrecision(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	// get ChainHead and lookback 5 epochs
	head, err := lapi.ChainHead(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	// TODO: this doesn't align with `@head` passed to the PreviewTerminateSectors call
	lookback := head.Height() - 1
	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), lookback, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	miners, err := lapi.StateListMiners(context.Background(), ts.Key())
	if err != nil {
		t.Fatal(err)
	}

	// make a slice the length of N
	chosenMiners := make([]address.Address, N)

	for i := 0; i < N; i++ {
		// choose a random miner
		randomIndex := rand.Intn(len(miners))
		chosenMiners[i] = miners[randomIndex]
	}
	// eventually run this test on all 10 miners in parallel
	miner := chosenMiners[0]
	fmt.Println("RANDOM MINER CHOSEN: ", miner)

	imprecise, err := PreviewTerminateSectorsQuick(context.Background(), lapi, miner, ts)
	if err != nil {
		t.Fatal(err)
	}

	errorCh := make(chan error)
	resultCh := make(chan *PreviewTerminateSectorsReturn)

	go PreviewTerminateSectors(context.Background(), lapi, miner, "@head", 0, 0, 0, false, false, false, 0, errorCh, nil, resultCh)

loop:
	for {
		select {
		case result := <-resultCh:
			if result.SectorStats.TerminationPenalty.Cmp(imprecise.SectorStats.TerminationPenalty) != 0 {
				t.Fatalf("TERMINATION PENALTIES DO NOT MATCH: precise: %v, imprecise: %v", result.SectorStats.TerminationPenalty, imprecise.SectorStats.TerminationPenalty)
			}
			break loop

		case err := <-errorCh:
			log.Fatal(err)
		}

	}

	// TODO: test the 10 random miners instead of just 1
}
