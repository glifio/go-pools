package terminate

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sort"
	"strconv"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
)

// this test compares N random miner termination penalties computed in the most precise (time intensive) way against the quick, less precise sampling method used in the ADO
var N = 3

// the wad based percentage that is acceptible for imprecision
var DIFF = big.NewInt(3e16)

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
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// r.Shuffle(len(miners), func(i, j int) { miners[i], miners[j] = miners[j], miners[i] })
	sort.Slice(miners, func(i, j int) bool {
		miner1, _ := strconv.Atoi(miners[i].String()[2:])
		miner2, _ := strconv.Atoi(miners[j].String()[2:])
		return miner1 < miner2
	})

	// // make a slice the length of N
	chosenMiners := make([]address.Address, 0)

	for i, miner := range miners {
		sectorCount, err := lapi.StateMinerSectorCount(context.Background(), miner, ts.Key())
		if err != nil {
			t.Fatal(err)
		}
		// if the miner has no active sectors, move to the next miner
		// fmt.Println(miner, sectorCount.Active)
		if i%100 == 0 {
			fmt.Printf("%v\n", miner)
		}
		if sectorCount.Active > 0 {
			fmt.Printf("Miner: %v - %d sectors @%v\n", miner, sectorCount.Active, ts.Height())
			chosenMiners = append(chosenMiners, miner)
			/*
				if len(chosenMiners) >= N {
					break
				}
			*/
		}
	}

	for i, miner := range chosenMiners {
		fmt.Printf("(%d/%d) %v:\n", i+1, N, miner)
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
				fmt.Println("PRECISE: ", result.SectorStats.TerminationPenalty)
				fmt.Println("IMPRECISE: ", imprecise.SectorStats.TerminationPenalty)
				if !util.AssertApproxEqRel(result.SectorStats.TerminationPenalty, imprecise.SectorStats.TerminationPenalty, DIFF) {
					t.Fatalf("TERMINATION PENALTIES DO NOT MATCH: precise: %v, imprecise: %v", result.SectorStats.TerminationPenalty, imprecise.SectorStats.TerminationPenalty)
				}
				break loop

			case err := <-errorCh:
				log.Fatal(err)
			}
		}

	}

	// TODO: test the 10 random miners instead of just 1
}
