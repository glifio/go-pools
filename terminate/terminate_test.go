package terminate

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/util"
)

// this test compares N random miner termination penalties computed in the most precise (time intensive) way against the quick, less precise sampling method used in the ADO
var N = 3

// the percentage that is acceptible for imprecision
var DIFF = big.NewFloat(3.00) // 3%

type SectorInfo struct {
	Deadline  int    `json:"Deadline"`
	Partition int    `json:"Partition"`
	Sectors   uint64 `json:"Sectors"`
}

var tests = []struct {
	name              string
	miner             string
	height            abi.ChainEpoch
	percent_to_sample int64
	want              string
}{
	{"miner that was off", "f01824405", 4157809, 1, "15276221548081039917042"},
	{"miner that was off", "f08403", 4157809, 1, "669276103568731990330"},
	{"miner that was off", "f02366381", 4157809, 1, "0"},
	{"miner that was off", "f01847751", 4157809, 1, "10305264645060108083102"},
	{"miner that was off", "f01315096", 4157809, 1, "25865743620631274061184"},
	{"miner that was off", "f02177086", 4158864, 1, "206662397221857395692"},
	{"miner that was off", "f01889668", 4157809, 1, "25205954710368106840743"},
}

func TestTerminationOffChainFullMiner(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	for _, tt := range tests {
		testname := fmt.Sprintf("%s:miner %s", tt.name, tt.miner)
		t.Run(testname, func(t *testing.T) {
			ts, err := lapi.ChainGetTipSetByHeight(context.Background(), tt.height, types.EmptyTSK)
			if err != nil {
				t.Fatal(err)
			}

			minerAddr, err := address.NewFromString(tt.miner)
			if err != nil {
				t.Fatal(err)
			}

			log.Println("Fetching sectors ")
			sectors, err := AllSectors(context.Background(), lapi, minerAddr, ts)
			if err != nil {
				t.Fatal(err)
			}

			var PERC_SECTORS_TO_SAMPLE = big.NewRat(tt.percent_to_sample, 100)
			var MAX_SAMPLED_SECTORS = big.NewInt(1000)
			var MIN_SAMPLED_SECTORS = big.NewInt(300)
			sampled := SampleSectors(sectors, PERC_SECTORS_TO_SAMPLE, MAX_SAMPLED_SECTORS, MIN_SAMPLED_SECTORS)

			log.Println("Sampled sectors length: ", len(sampled))

			// Sample sectors
			sample := bitfield.NewFromSet(sampled)
			full := bitfield.NewFromSet(sectors)

			// if want is 0 fetch all sectors
			expectedTermFee := big.NewInt(0)
			if tt.want == "-1" {
				resFull, err := TerminateSectors(context.Background(), lapi, minerAddr, &full, ts)
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("Termination fee: %v", resFull.TerminationFeeFromSample)
				expectedTermFee = resFull.EstimatedTotalTerminationFee
			} else {
				expectedTermFee.SetString(tt.want, 10)
			}

			res, err := TerminateSectors(context.Background(), lapi, minerAddr, &sample, ts)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("Estimated total pledge: %v", util.ToFIL(res.EstimatedTotalInitialPledge))

			termFeeDiff := util.Diff(expectedTermFee, res.EstimatedTotalTerminationFee)
			initialPledgeDiff := util.Diff(res.InitialPledge, res.EstimatedTotalInitialPledge)
			t.Logf("The difference in term fee: %s%%", util.Diff(expectedTermFee, res.EstimatedTotalTerminationFee).Text('f', 2))
			t.Logf("The difference in initial pledge: %s%%", util.Diff(res.InitialPledge, res.EstimatedTotalInitialPledge).Text('f', 2))

			// if the difference in term fee is greater than 3% or the difference in initial pledge is greater than 3%, fail the test
			if termFeeDiff.Cmp(DIFF) == 1 {
				t.Fatalf("Expected term fee: %v, actual fee: %v", util.ToFIL(expectedTermFee), util.ToFIL(res.EstimatedTotalTerminationFee))
			}
			if initialPledgeDiff.Cmp(DIFF) == 1 {
				t.Fatalf("Expected initial pledge: %v, actual: %v", util.ToFIL(res.InitialPledge), util.ToFIL(res.EstimatedTotalInitialPledge))
			}
		})
	}
}

func TestTerminationPrecisionFromOffChain(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), 4158864, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	minerAddr, err := address.NewFromString("f01502887")
	if err != nil {
		t.Fatal(err)
	}
	sampled := []uint64{12607}

	sample := bitfield.NewFromSet(sampled)

	res, err := TerminateSectors(context.Background(), lapi, minerAddr, &sample, ts)
	if err != nil {
		t.Fatal(err)
	}

	// test if the penalty is within 3% of the expected value
	expectedPenalty := big.NewInt(9397221857395692)
	if !util.AssertApproxEqAbs(res.TerminationFeeFromSample, expectedPenalty, big.NewInt(3e1)) {
		t.Fatalf("Expected penalty: %v, actual penalty: %v", expectedPenalty, res.TerminationFeeFromSample)
	} else {
		t.Logf("Expected penalty: %v, actual penalty: %v", expectedPenalty, res.TerminationFeeFromSample)
	}

}

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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(miners), func(i, j int) { miners[i], miners[j] = miners[j], miners[i] })

	// // make a slice the length of N
	chosenMiners := make([]address.Address, 0)

	for _, miner := range miners {
		sectorCount, err := lapi.StateMinerSectorCount(context.Background(), miner, ts.Key())
		if err != nil {
			t.Fatal(err)
		}
		// if the miner has no active sectors, move to the next miner
		// fmt.Println(miner, sectorCount.Active)
		if sectorCount.Active > 0 {
			chosenMiners = append(chosenMiners, miner)
			if len(chosenMiners) >= N {
				break
			}
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
				if !util.AssertApproxEqRel(result.SectorStats.TerminationPenalty, imprecise.SectorStats.TerminationPenalty, big.NewInt(3e16)) {
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

func TestEstimateTerminationPenalty2(t *testing.T) {
	data := []SectorInfo{
		{Deadline: 7, Partition: 0, Sectors: 18722},
		{Deadline: 7, Partition: 0, Sectors: 18723},
		{Deadline: 7, Partition: 0, Sectors: 18724},
		{Deadline: 7, Partition: 0, Sectors: 18725},
		{Deadline: 7, Partition: 0, Sectors: 18726},
		{Deadline: 7, Partition: 0, Sectors: 18727},
		{Deadline: 7, Partition: 0, Sectors: 18728},
		{Deadline: 7, Partition: 0, Sectors: 18729},
		{Deadline: 7, Partition: 0, Sectors: 18730},
		{Deadline: 7, Partition: 0, Sectors: 18731},
		{Deadline: 7, Partition: 0, Sectors: 18732},
		{Deadline: 7, Partition: 0, Sectors: 18733},
		{Deadline: 7, Partition: 0, Sectors: 18734},
		{Deadline: 7, Partition: 0, Sectors: 18735},
		{Deadline: 7, Partition: 0, Sectors: 18736},
		{Deadline: 7, Partition: 0, Sectors: 18737},
		{Deadline: 7, Partition: 0, Sectors: 18738},
		{Deadline: 7, Partition: 0, Sectors: 18740},
		{Deadline: 7, Partition: 0, Sectors: 18741},
		{Deadline: 7, Partition: 0, Sectors: 18742},
		{Deadline: 7, Partition: 0, Sectors: 18743},
		{Deadline: 7, Partition: 0, Sectors: 18744},
		{Deadline: 7, Partition: 0, Sectors: 18745},
		{Deadline: 7, Partition: 0, Sectors: 18746},
		{Deadline: 7, Partition: 0, Sectors: 18747},
		{Deadline: 7, Partition: 0, Sectors: 18748},
		{Deadline: 7, Partition: 0, Sectors: 18749},
		{Deadline: 7, Partition: 0, Sectors: 18750},
		{Deadline: 7, Partition: 0, Sectors: 18751},
		{Deadline: 7, Partition: 0, Sectors: 18752},
		{Deadline: 7, Partition: 0, Sectors: 18753},
		{Deadline: 7, Partition: 0, Sectors: 18754},
		{Deadline: 7, Partition: 0, Sectors: 18755},
		{Deadline: 7, Partition: 0, Sectors: 18756},
		{Deadline: 7, Partition: 0, Sectors: 18757},
		{Deadline: 7, Partition: 0, Sectors: 18758},
		{Deadline: 7, Partition: 0, Sectors: 18759},
		{Deadline: 7, Partition: 0, Sectors: 18760},
		{Deadline: 7, Partition: 0, Sectors: 18761},
		{Deadline: 7, Partition: 0, Sectors: 18762},
		{Deadline: 7, Partition: 0, Sectors: 18763},
		{Deadline: 7, Partition: 0, Sectors: 18764},
		{Deadline: 7, Partition: 0, Sectors: 18765},
		{Deadline: 7, Partition: 0, Sectors: 18766},
		{Deadline: 7, Partition: 0, Sectors: 18767},
		{Deadline: 7, Partition: 0, Sectors: 18768},
		{Deadline: 7, Partition: 0, Sectors: 18769},
		{Deadline: 7, Partition: 0, Sectors: 18770},
		{Deadline: 7, Partition: 0, Sectors: 18771},
		{Deadline: 7, Partition: 0, Sectors: 18772},
		{Deadline: 7, Partition: 0, Sectors: 18773},
		{Deadline: 7, Partition: 0, Sectors: 18774},
		{Deadline: 7, Partition: 0, Sectors: 18776},
		{Deadline: 7, Partition: 0, Sectors: 18777},
		{Deadline: 7, Partition: 0, Sectors: 18778},
		{Deadline: 7, Partition: 0, Sectors: 18779},
		{Deadline: 7, Partition: 0, Sectors: 18780},
		{Deadline: 7, Partition: 0, Sectors: 18781},
		{Deadline: 7, Partition: 0, Sectors: 18782},
		{Deadline: 7, Partition: 0, Sectors: 18783},
		{Deadline: 7, Partition: 0, Sectors: 18784},
		{Deadline: 7, Partition: 0, Sectors: 18785},
		{Deadline: 7, Partition: 0, Sectors: 18786},
		{Deadline: 7, Partition: 0, Sectors: 18787},
		{Deadline: 7, Partition: 0, Sectors: 18788},
		{Deadline: 7, Partition: 0, Sectors: 18789},
		{Deadline: 7, Partition: 0, Sectors: 18790},
		{Deadline: 7, Partition: 0, Sectors: 18791},
		{Deadline: 7, Partition: 0, Sectors: 18792},
		{Deadline: 7, Partition: 0, Sectors: 18793},
	}

	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), 4161576, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	minerAddr, err := address.NewFromString("f01894158")
	if err != nil {
		t.Fatal(err)
	}

	// extract the sectors from the data into sampled
	sampled := make([]uint64, len(data))
	for i, sector := range data {
		sampled[i] = sector.Sectors
	}

	sample := bitfield.NewFromSet(sampled)

	res, err := TerminateSectors(context.Background(), lapi, minerAddr, &sample, ts)
	if err != nil {
		t.Fatal(err)
	}

	// test if the penalty is within 3% of the expected value
	// create big.int from string
	expectedPenalty := big.NewInt(0)
	expectedPenalty.SetString("23690174266209542000", 10)

	if !util.AssertApproxEqAbs(res.TerminationFeeFromSample, expectedPenalty, big.NewInt(3e1)) {
		t.Fatalf("Expected penalty: %v, actual penalty: %v", expectedPenalty, res.TerminationFeeFromSample)
	} else {
		t.Logf("Expected penalty: %v, actual penalty: %v", expectedPenalty, res.TerminationFeeFromSample)
	}
}
