package econ

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/terminate"
	"github.com/glifio/go-pools/util"
)

// this test compares N random miner termination penalties computed in the most precise (time intensive) way against the quick, less precise sampling method used in the ADO
var N = 3

type SectorInfo struct {
	Deadline  int    `json:"Deadline"`
	Partition int    `json:"Partition"`
	Sectors   uint64 `json:"Sectors"`
}

var tests = []struct {
	name    string
	miner   string
	height  abi.ChainEpoch
	samples int
	want    string
}{
	{"empty miner", "f01882569", 4182470, 1000, "0"},
	{"miner", "f01344987", 4161576, 1000, "19784479924073946376310"},
	{"miner", "f01824405", 4157809, 1000, "15276221548081039917042"},
	{"miner", "f08403", 4157809, 1000, "669276103568731990330"},
	{"miner", "f02366381", 4157809, 1000, "0"},
	{"miner", "f01847751", 4157809, 1000, "10305264645060108083102"},
	{"miner", "f01315096", 4157809, 1000, "25865743620631274061184"},
	{"miner", "f02177086", 4158864, 1000, "206662397221857395692"},
	{"miner", "f01889668", 4157809, 1000, "25205954710368106840743"},
}

func TestTerminationOffChainFullMiner(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.name, tt.miner)
		t.Run(testname, func(t *testing.T) {
			ts, err := lapi.ChainGetTipSetByHeight(context.Background(), tt.height, types.EmptyTSK)
			if err != nil {
				t.Fatal(err)
			}

			minerAddr, err := address.NewFromString(tt.miner)
			if err != nil {
				t.Fatal(err)
			}

			log.Println("get all sectors ")
			sectors, err := AllSectors(context.Background(), lapi, minerAddr, ts)
			if err != nil {
				t.Fatal(err)
			}

			log.Println("sample sectors ")

			sampled := SampleSectors(sectors, tt.samples)

			log.Println("Sampled sectors length: ", len(sampled))

			// Sample sectors
			sample := bitfield.NewFromSet(sampled)
			full := bitfield.NewFromSet(sectors)

			// if want is -1 fetch all sectors
			expectedTermFee := big.NewInt(0)
			if tt.want == "-1" {
				resFull, err := TerminateSectors(context.Background(), lapi, minerAddr, &full, ts)
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("Termination fee: %v", resFull.TerminationFeeFromSample)
				expectedTermFee = resFull.EstimatedTerminationFee
			} else {
				expectedTermFee.SetString(tt.want, 10)
			}

			res, err := TerminateSectors(context.Background(), lapi, minerAddr, &sample, ts)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("Estimated total pledge: %v", util.ToFIL(res.EstimatedInitialPledge))

			termFeeDiff := util.Diff(expectedTermFee, res.EstimatedTerminationFee)
			initialPledgeDiff := util.Diff(res.InitialPledge, res.EstimatedInitialPledge)
			t.Logf("The difference in term fee: %s%%", util.Diff(expectedTermFee, res.EstimatedTerminationFee).Text('f', 2))
			t.Logf("The difference in initial pledge: %s%%", util.Diff(res.InitialPledge, res.EstimatedInitialPledge).Text('f', 2))

			// if the difference in term fee is greater than 3% or the difference in initial pledge is greater than 3%, fail the test
			if termFeeDiff.Cmp(INITIAL_PLEDGE_INTERPOLATION_REL_DIFF) == 1 {
				t.Fatalf("Expected term fee: %v, actual fee: %v", util.ToFIL(expectedTermFee), util.ToFIL(res.EstimatedTerminationFee))
			}
			if initialPledgeDiff.Cmp(INITIAL_PLEDGE_INTERPOLATION_REL_DIFF) == 1 {
				t.Fatalf("Expected initial pledge: %v, actual: %v", util.ToFIL(res.InitialPledge), util.ToFIL(res.EstimatedInitialPledge))
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

	expectedTermFee := big.NewInt(9397221857395692)
	termFeeDiff := util.Diff(expectedTermFee, res.TerminationFeeFromSample)
	if termFeeDiff.Cmp(INITIAL_PLEDGE_INTERPOLATION_REL_DIFF) == 1 {
		t.Fatalf("Expected term fee: %v, actual fee: %v", util.ToFIL(expectedTermFee), util.ToFIL(res.EstimatedTerminationFee))
	}
}

func XTestTerminationPrecision(t *testing.T) {
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
		imprecise, err := EstimateTerminationFeeMiner(context.Background(), lapi, miner, ts)
		if err != nil {
			t.Fatal(err)
		}

		errorCh := make(chan error)
		resultCh := make(chan *terminate.PreviewTerminateSectorsReturn)

		go terminate.PreviewTerminateSectors(context.Background(), lapi, miner, "@head", 0, 0, 0, false, false, false, 0, errorCh, nil, resultCh)

	loop:
		for {
			select {
			case result := <-resultCh:
				fmt.Println("PRECISE: ", result.SectorStats.TerminationPenalty)
				fmt.Println("IMPRECISE: ", imprecise.EstimatedTerminationFee)
				if !util.AssertApproxEqRel(result.SectorStats.TerminationPenalty, imprecise.EstimatedTerminationFee, big.NewInt(3e16)) {
					t.Fatalf("TERMINATION PENALTIES DO NOT MATCH: precise: %v, imprecise: %v", result.SectorStats.TerminationPenalty, imprecise.EstimatedTerminationFee)
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

var BLOCK_HEIGHT = abi.ChainEpoch(4198539)

type baseFiTest struct {
	*BaseFi

	maxBorrowAndSeal     *big.Int
	maxBorrowAndWithdraw *big.Int
	liquidationValue     *big.Int
}

var baseFiTests = []struct {
	name  string
	miner string
	want  baseFiTest
}{
	{"no sectors miner", "f01882569", baseFiTest{
		BaseFi: &BaseFi{
			Balance:          bigFromStr("16791927372141826678333"),
			AvailableBalance: bigFromStr("16753493784149168047801"),
			LockedRewards:    bigFromStr("38433587992658630532"),
			InitialPledge:    big.NewInt(0),
			FeeDebt:          big.NewInt(0),
			TerminationFee:   big.NewInt(0),
			LiveSectors:      big.NewInt(0),
			FaultySectors:    big.NewInt(0),
		},
		maxBorrowAndSeal:     bigFromStr("50375782116425480034999"),
		maxBorrowAndWithdraw: bigFromStr("12593945529106370008750"),
		liquidationValue:     bigFromStr("16791927372141826678333"),
	}},
	{"miner with fee debt", "f01836766", baseFiTest{
		BaseFi: &BaseFi{
			Balance:          big.NewInt(0),
			AvailableBalance: big.NewInt(0),
			LockedRewards:    big.NewInt(0),
			InitialPledge:    big.NewInt(0),
			FeeDebt:          bigFromStr("43697948992036495293"),
			TerminationFee:   big.NewInt(0),
			LiveSectors:      big.NewInt(0),
			FaultySectors:    big.NewInt(0),
		},
		maxBorrowAndSeal:     big.NewInt(0),
		maxBorrowAndWithdraw: big.NewInt(0),
		liquidationValue:     big.NewInt(0),
	}},
	{"inactive miner (all 0s)", "f01156", baseFiTest{
		BaseFi: &BaseFi{
			Balance:          big.NewInt(0),
			AvailableBalance: big.NewInt(0),
			LockedRewards:    big.NewInt(0),
			InitialPledge:    big.NewInt(0),
			FeeDebt:          big.NewInt(0),
			TerminationFee:   big.NewInt(0),
			LiveSectors:      big.NewInt(0),
			FaultySectors:    big.NewInt(0),
		},
		maxBorrowAndSeal:     big.NewInt(0),
		maxBorrowAndWithdraw: big.NewInt(0),
		liquidationValue:     big.NewInt(0),
	}},
	{"normal miner", "f01344987", baseFiTest{
		BaseFi: &BaseFi{
			Balance:          bigFromStr("183615910324483211964542"),
			AvailableBalance: bigFromStr("217841811153934568647"),
			LockedRewards:    bigFromStr("8591178717907362918730"),
			InitialPledge:    bigFromStr("174806889795421914477165"),
			FeeDebt:          big.NewInt(0),
			TerminationFee:   bigFromStr("19439988058419303079504"),
			LiveSectors:      big.NewInt(89417),
			FaultySectors:    big.NewInt(0),
		},
		maxBorrowAndSeal:     bigFromStr("492527766798191726655114"),
		maxBorrowAndWithdraw: bigFromStr("123131941699547931663779"),
		liquidationValue:     bigFromStr("164175922266063908885038"),
	}},
}

func TestBaseFi(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), BLOCK_HEIGHT, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range baseFiTests {
		testname := fmt.Sprintf("%s:%s", tt.name, tt.miner)
		t.Run(testname, func(t *testing.T) {
			minerAddr, err := address.NewFromString(tt.miner)
			if err != nil {
				t.Fatal(err)
			}

			ret, err := EstimateTerminationFeeMiner(context.Background(), lapi, minerAddr, ts)
			if err != nil {
				t.Fatal(err)
			}

			assertBaseFiEqual(t, &tt.want, ret.ToBaseFi())
		})
	}
}

type agentFiTest struct {
	*AgentFi

	balance              *big.Int
	maxBorrowAndSeal     *big.Int
	maxBorrowAndWithdraw *big.Int
	liquidationValue     *big.Int
	borrowLimit          *big.Int
	withdrawLimit        *big.Int
	marginCall           *big.Int
	leverageRatio        float64
}

var agentTests = []struct {
	name  string
	agent string
	want  agentFiTest
}{
	// agent 59
	{"agent many miners", "0x64Ea1aD49A78B6A6878c4975566b8953b1Ef1e79", agentFiTest{
		AgentFi: &AgentFi{
			BaseFi: BaseFi{
				Balance:          bigFromStr("2090733730144435384596917"),
				AvailableBalance: bigFromStr("8911457546764252091903"),
				LockedRewards:    bigFromStr("90756351468487272842162"),
				InitialPledge:    bigFromStr("1991050146969153602729858"),
				FeeDebt:          big.NewInt(0),
				TerminationFee:   bigFromStr("207962529375567399782062"),
				LiveSectors:      big.NewInt(2805488),
				FaultySectors:    big.NewInt(0),
			},
			Principal: bigFromStr("1081255000003082191780823"),
		},
		maxBorrowAndSeal:     bigFromStr("5648313602306603954444565"),
		maxBorrowAndWithdraw: bigFromStr("1412078400576650988611142"),
		liquidationValue:     bigFromStr("1882771200768867984814855"),
		borrowLimit:          bigFromStr("4567058602303521762663742"),
		withdrawLimit:        bigFromStr("441097867431425062440425"),
		marginCall:           bigFromStr("1272064705885979049153909"),
		leverageRatio:        0.5742891114764925,
	}},
	// agent 27
	{"agent no miners", "0xDBa96B0FDbc87C7eEb641Ee37EAFC55B355079E4", agentFiTest{
		AgentFi:              EmptyAgentFi(),
		balance:              big.NewInt(0),
		maxBorrowAndSeal:     big.NewInt(0),
		maxBorrowAndWithdraw: big.NewInt(0),
		liquidationValue:     big.NewInt(0),
		borrowLimit:          big.NewInt(0),
		withdrawLimit:        big.NewInt(0),
		marginCall:           big.NewInt(0),
		leverageRatio:        0,
	}},
	// agent 2
	{"agent single miner", "0xf0F1ceCCF78D411EeD9Ca190ca7F157140cCB2d3", agentFiTest{
		AgentFi: &AgentFi{
			BaseFi: BaseFi{
				Balance:          bigFromStr("133252309755096579751"),
				AvailableBalance: bigFromStr("57665444302030819344"),
				LockedRewards:    bigFromStr("10682659000384066417"),
				InitialPledge:    bigFromStr("64904206452681693990"),
				FeeDebt:          big.NewInt(0),
				TerminationFee:   bigFromStr("12832724091624933434"),
				LiveSectors:      big.NewInt(323),
				FaultySectors:    big.NewInt(0),
			},
			Principal: bigFromStr("2000000000000000000"),
		},
		maxBorrowAndSeal:     bigFromStr("361258756990414938951"),
		maxBorrowAndWithdraw: bigFromStr("90314689247603734738"),
		liquidationValue:     bigFromStr("120419585663471646317"),
		borrowLimit:          bigFromStr("359258756990414938951"),
		withdrawLimit:        bigFromStr("117752918996804979651"),
		marginCall:           bigFromStr("2352941176470588235"),
		leverageRatio:        0.0166085939341235,
	}},
	// agent 91
	{"agent miner with fee debt", "0xFF65F5f3D309fEA7aA2d4cB2727E918FAb0aE7F7", agentFiTest{
		AgentFi: &AgentFi{
			BaseFi: BaseFi{
				Balance:          bigFromStr("0"),
				AvailableBalance: bigFromStr("145838619789"),
				LockedRewards:    bigFromStr("0"),
				InitialPledge:    bigFromStr("0"),
				FeeDebt:          bigFromStr("43697948992036495293"),
				TerminationFee:   bigFromStr("0"),
				LiveSectors:      big.NewInt(0),
				FaultySectors:    big.NewInt(0),
			},
			Principal: bigFromStr("0"),
		},
		maxBorrowAndSeal:     bigFromStr("0"),
		maxBorrowAndWithdraw: bigFromStr("0"),
		liquidationValue:     bigFromStr("0"),
		borrowLimit:          bigFromStr("0"),
		withdrawLimit:        bigFromStr("0"),
		marginCall:           bigFromStr("0"),
		leverageRatio:        0,
	}},
}

func TestEstimateTermationFeeAgent(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), BLOCK_HEIGHT, types.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range agentTests {
		testname := fmt.Sprintf("%s:%s", tt.name, tt.agent)
		t.Run(testname, func(t *testing.T) {

			agent := common.HexToAddress(tt.agent)

			agentFi, err := EstimateTerminationFeeAgent(context.Background(), agent, address.Undef, psdk, ts)
			if err != nil {
				t.Fatal(err)
			}

			assertAgentFiEqual(t, &tt.want, agentFi)
		})
	}
}

func assertAgentFiEqual(t *testing.T, expected *agentFiTest, actual *AgentFi) {
	if expected.AvailableBalance.Cmp(actual.AvailableBalance) != 0 {
		t.Fatalf("Expected available balance: %v, actual: %v", expected.AvailableBalance, actual.AvailableBalance)
	}
	if expected.LockedRewards.Cmp(actual.LockedRewards) != 0 {
		t.Fatalf("Expected locked rewards: %v, actual: %v", expected.LockedRewards, actual.LockedRewards)
	}
	if expected.InitialPledge.Cmp(actual.InitialPledge) != 0 {
		t.Fatalf("Expected initial pledge: %v, actual: %v", expected.InitialPledge, actual.InitialPledge)
	}
	if expected.FeeDebt.Cmp(actual.FeeDebt) != 0 {
		t.Fatalf("Expected fee debt: %v, actual: %v", expected.FeeDebt, actual.FeeDebt)
	}
	if expected.TerminationFee.Cmp(actual.TerminationFee) != 0 {
		t.Fatalf("Expected termination fee: %v, actual: %v", expected.TerminationFee, actual.TerminationFee)
	}
	if expected.Principal.Cmp(actual.Principal) != 0 {
		t.Fatalf("Expected principal: %v, actual: %v", expected.Principal, actual.Principal)
	}
	if expected.Balance.Cmp(actual.Balance) != 0 {
		t.Fatalf("Expected balance: %v, actual: %v", expected.balance, actual.Balance)
	}
	if expected.FaultySectors.Cmp(actual.FaultySectors) != 0 {
		t.Fatalf("Expected faulty sectors: %v, actual: %v", expected.FaultySectors, actual.FaultySectors)
	}
	if expected.LiveSectors.Cmp(actual.LiveSectors) != 0 {
		t.Fatalf("Expected live sectors: %v, actual: %v", expected.LiveSectors, actual.LiveSectors)
	}

	if expected.maxBorrowAndSeal.Cmp(actual.MaxBorrowAndSeal()) != 0 {
		t.Fatalf("Expected max borrow and seal: %v, actual: %v", expected.maxBorrowAndSeal, actual.MaxBorrowAndSeal())
	}
	if expected.maxBorrowAndWithdraw.Cmp(actual.MaxBorrowAndWithdraw()) != 0 {
		t.Fatalf("Expected max borrow and withdraw: %v, actual: %v", expected.maxBorrowAndWithdraw, actual.MaxBorrowAndWithdraw())
	}
	if expected.liquidationValue.Cmp(actual.LiquidationValue()) != 0 {
		t.Fatalf("Expected liquidation value: %v, actual: %v", expected.liquidationValue, actual.LiquidationValue())
	}
	if expected.borrowLimit.Cmp(actual.BorrowLimit()) != 0 {
		t.Fatalf("Expected borrow limit: %v, actual: %v", expected.borrowLimit, actual.BorrowLimit())
	}
	if expected.withdrawLimit.Cmp(actual.WithdrawLimit()) != 0 {
		t.Fatalf("Expected withdraw limit: %v, actual: %v", expected.withdrawLimit, actual.WithdrawLimit())
	}
	if expected.marginCall.Cmp(actual.MarginCall()) != 0 {
		t.Fatalf("Expected margin call: %v, actual: %v", expected.marginCall, actual.MarginCall())
	}
	if expected.leverageRatio != actual.LeverageRatio() {
		t.Fatalf("Expected leverage ratio: %v, actual: %v", expected.leverageRatio, actual.LeverageRatio())
	}
}

func assertBaseFiEqual(t *testing.T, expected *baseFiTest, actual *BaseFi) {
	if expected.AvailableBalance.Cmp(actual.AvailableBalance) != 0 {
		t.Fatalf("Expected available balance: %v, actual: %v", expected.AvailableBalance, actual.AvailableBalance)
	}
	if expected.LockedRewards.Cmp(actual.LockedRewards) != 0 {
		t.Fatalf("Expected locked rewards: %v, actual: %v", expected.LockedRewards, actual.LockedRewards)
	}
	if expected.InitialPledge.Cmp(actual.InitialPledge) != 0 {
		t.Fatalf("Expected initial pledge: %v, actual: %v", expected.InitialPledge, actual.InitialPledge)
	}
	if expected.FeeDebt.Cmp(actual.FeeDebt) != 0 {
		t.Fatalf("Expected fee debt: %v, actual: %v", expected.FeeDebt, actual.FeeDebt)
	}
	if expected.TerminationFee.Cmp(actual.TerminationFee) != 0 {
		t.Fatalf("Expected termination fee: %v, actual: %v", expected.TerminationFee, actual.TerminationFee)
	}
	if expected.Balance.Cmp(actual.Balance) != 0 {
		t.Fatalf("Expected balance: %v, actual: %v", expected.Balance, actual.Balance)
	}
	if expected.FaultySectors.Cmp(actual.FaultySectors) != 0 {
		t.Fatalf("Expected faulty sectors: %v, actual: %v", expected.FaultySectors, actual.FaultySectors)
	}
	if expected.LiveSectors.Cmp(actual.LiveSectors) != 0 {
		t.Fatalf("Expected live sectors: %v, actual: %v", expected.LiveSectors, actual.LiveSectors)
	}

	if expected.maxBorrowAndSeal.Cmp(actual.MaxBorrowAndSeal()) != 0 {
		t.Fatalf("Expected max borrow and seal: %v, actual: %v", expected.maxBorrowAndSeal, actual.MaxBorrowAndSeal())
	}
	if expected.maxBorrowAndWithdraw.Cmp(actual.MaxBorrowAndWithdraw()) != 0 {
		t.Fatalf("Expected max borrow and withdraw: %v, actual: %v", expected.maxBorrowAndWithdraw, actual.MaxBorrowAndWithdraw())
	}
	if expected.liquidationValue.Cmp(actual.LiquidationValue()) != 0 {
		t.Fatalf("Expected liquidation value: %v, actual: %v", expected.liquidationValue, actual.LiquidationValue())
	}
}

func bigFromStr(s string) *big.Int {
	b := new(big.Int)
	b.SetString(s, 10)
	return b
}
