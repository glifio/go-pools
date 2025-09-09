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
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/constants"
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
	{"miner", "f08403", 4171951, 1000, "685608907628163727896"},
	{"miner", "f01344987", 4161576, 1000, "19784479924073946376310"},
	{"miner", "f01824405", 4157809, 1000, "15276221548081039917042"},
	{"miner", "f08403", 4157809, 1000, "669276103568731990330"},
	{"miner", "f02366381", 4157809, 1000, "0"},
	{"miner", "f01847751", 4157809, 1000, "10305264645060108083102"},
	{"miner", "f01315096", 4157809, 1000, "25865743620631274061184"},
	{"miner", "f02177086", 4158864, 1000, "206662397221857395692"},
	{"miner", "f01889668", 4157809, 1000, "25205954710368106840743"},
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

type baseFiTest struct {
	*BaseFi

	maxBorrowAndSeal     *big.Int
	maxBorrowAndWithdraw *big.Int
	liquidationValue     *big.Int
	height               *big.Int
}

var baseFiTests = []struct {
	name  string
	miner string
	want  baseFiTest
}{
	{"no sectors miner", "f01882569", baseFiTest{
		BaseFi: &BaseFi{
			Balance:          bigFromStr("927372141826678333"),
			AvailableBalance: bigFromStr("927372141826678333"),
			LockedRewards:    big.NewInt(0),
			InitialPledge:    big.NewInt(0),
			FeeDebt:          big.NewInt(0),
			TerminationFee:   big.NewInt(0),
			LiveSectors:      big.NewInt(0),
			FaultySectors:    big.NewInt(0),
		},
		maxBorrowAndSeal:     bigFromStr("2782116425480034999"),
		maxBorrowAndWithdraw: bigFromStr("695529106370008750"),
		liquidationValue:     bigFromStr("927372141826678333"),
		height:               BASE_BLOCK_HEIGHT,
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
		height:               BASE_BLOCK_HEIGHT,
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
		height:               BASE_BLOCK_HEIGHT,
	}},
	{"normal miner", "f01344987", baseFiTest{
		BaseFi: &BaseFi{
			Balance:          bigFromStr("39773757056590498353766"),
			AvailableBalance: bigFromStr("450324027859823088796"),
			LockedRewards:    bigFromStr("2485958269968226649253"),
			InitialPledge:    bigFromStr("36837474758762448615717"),
			FeeDebt:          big.NewInt(0),
			TerminationFee:   bigFromStr("3131185354494808132335"),
			LiveSectors:      big.NewInt(19363),
			FaultySectors:    big.NewInt(0),
		},
		maxBorrowAndSeal:     bigFromStr("109927715106287070664293"),
		maxBorrowAndWithdraw: bigFromStr("27481928776571767666074"),
		liquidationValue:     bigFromStr("36642571702095690221431"),
		height:               BASE_BLOCK_HEIGHT,
	}},
	// THIS TEST FAILS BECAUSE WE NEED A BETTER, MODERN EXAMPLE
	// {"miner with fee debt and positive balance", "f01931245", baseFiTest{
	// 	BaseFi: &BaseFi{
	// 		Balance:          bigFromStr("65041872611791874551"),
	// 		AvailableBalance: bigFromStr("0"),
	// 		LockedRewards:    bigFromStr("137666159110180561"),
	// 		InitialPledge:    bigFromStr("64904206452681693990"),
	// 		FeeDebt:          bigFromStr("957933257513729319"),
	// 		TerminationFee:   bigFromStr("12832724091624933434"),
	// 		LiveSectors:      big.NewInt(323),
	// 		FaultySectors:    big.NewInt(323),
	// 	},
	// 	maxBorrowAndSeal:     bigFromStr("156627445560500823351"),
	// 	maxBorrowAndWithdraw: bigFromStr("39156861390125205838"),
	// 	liquidationValue:     bigFromStr("52209148520166941117"),
	// 	height:               big.NewInt(4299545),
	// }},
}

func TestMinerFi(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	for _, tt := range baseFiTests {
		testname := fmt.Sprintf("%s:%s", tt.name, tt.miner)
		t.Run(testname, func(t *testing.T) {
			ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(tt.want.height.Int64()), types.EmptyTSK)
			if err != nil {
				t.Fatal(err)
			}
			minerAddr, err := address.NewFromString(tt.miner)
			if err != nil {
				t.Fatal(err)
			}

			ret, err := EstimateTerminationFeeMiner(context.Background(), lapi, minerAddr, ts)
			if err != nil {
				t.Fatal(err)
			}

			minerFi := MinerFi{BaseFi: *ret.ToBaseFi()}
			lv := minerFi.LiquidationValue()

			assertBaseFiEqual(t, &tt.want, &minerFi)

			// if the miner can borrow, test the limits
			if lv.Sign() == 1 {
				// test the max borrow and seal by simulating a borrow
				maxBorrowAndSeal := minerFi.MaxBorrowAndSeal()
				// new collateral value after borrowing
				newCollateralValue := big.NewInt(0).Add(lv, maxBorrowAndSeal)
				// test under DTL
				dtl := util.DivWad(maxBorrowAndSeal, newCollateralValue)
				if dtl.Cmp(constants.MAX_BORROW_DTL) != 0 {
					t.Fatalf("DTL is not equal to max borrow dtl: %v", dtl)
				}

				// test max borrow and withdraw by simulating a borrow and withdraw
				maxBorrowAndWithdraw := minerFi.MaxBorrowAndWithdraw()
				// new collateral value after borrowing is the old liquidation value
				newCollateralValue = lv
				// test under DTL
				dtl = util.DivWad(maxBorrowAndWithdraw, newCollateralValue)
				if dtl.Cmp(constants.MAX_BORROW_DTL) == 1 {
					t.Fatalf("DTL is greater than max borrow dtl: %v", dtl)
				}
			}
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
	dtl                  float64
}

var agentTests = []struct {
	name  string
	agent string
	block *big.Int
	want  agentFiTest
}{
	// agent 59
	{"agent many miners", "0x64Ea1aD49A78B6A6878c4975566b8953b1Ef1e79", BASE_BLOCK_HEIGHT, agentFiTest{
		AgentFi: &AgentFi{
			BaseFi: BaseFi{
				Balance:          bigFromStr("2819592172520852577940854"),
				AvailableBalance: bigFromStr("12737055656823005978228"),
				LockedRewards:    bigFromStr("133496862716254206867810"),
				InitialPledge:    bigFromStr("2673358209388044369719498"),
				FeeDebt:          big.NewInt(0),
				TerminationFee:   bigFromStr("227235447797983771426153"),
				LiveSectors:      big.NewInt(1241723),
				FaultySectors:    big.NewInt(0),
			},
			Liability: Liability{
				Principal: bigFromStr("1550270393738119289518458"),
				Interest:  bigFromStr("0"),
			},
		},
		maxBorrowAndSeal:     bigFromStr("3123377454405613741932615"),
		maxBorrowAndWithdraw: bigFromStr("393036636954487379015530"),
		liquidationValue:     bigFromStr("2592356724722868806514701"),
		borrowLimit:          bigFromStr("1572146547817949516062119"),
		withdrawLimit:        bigFromStr("12737055656823005978228"),
		marginCall:           bigFromStr("1824977537161957912788818"),
		leverageRatio:        2.489955276842645206,
		dtl:                  0.5983863608714945,
	}},
	// agent 27
	{"agent no miners", "0xDBa96B0FDbc87C7eEb641Ee37EAFC55B355079E4", BASE_BLOCK_HEIGHT, agentFiTest{
		AgentFi:              EmptyAgentFi(),
		balance:              big.NewInt(0),
		maxBorrowAndSeal:     big.NewInt(0),
		maxBorrowAndWithdraw: big.NewInt(0),
		liquidationValue:     big.NewInt(0),
		borrowLimit:          big.NewInt(0),
		withdrawLimit:        big.NewInt(0),
		marginCall:           big.NewInt(0),
		leverageRatio:        1,
		dtl:                  0,
	}},
	// agent 212
	{"agent single miner", "0x7a46ac436cF1606eABd8BD5F6B9A169258c01452", BASE_BLOCK_HEIGHT, agentFiTest{
		AgentFi: &AgentFi{
			BaseFi: BaseFi{
				Balance:          bigFromStr("102616089602808694802964"),
				AvailableBalance: bigFromStr("80595336348363929738"),
				LockedRewards:    bigFromStr("2796443644134574248635"),
				InitialPledge:    bigFromStr("99738765355288049546651"),
				FeeDebt:          big.NewInt(0),
				TerminationFee:   bigFromStr("8477795055199484211465"),
				LiveSectors:      big.NewInt(80456),
				FaultySectors:    big.NewInt(0),
			},
			Liability: Liability{
				Principal: bigFromStr("73569777167058713695624"),
				Interest:  bigFromStr("0"),
			},
		},
		maxBorrowAndSeal:     bigFromStr("61513785714780677114127"),
		maxBorrowAndWithdraw: bigFromStr("-3029978398642076943165"),
		liquidationValue:     bigFromStr("94138294547609210591499"),
		borrowLimit:          bigFromStr("-12119913594568307772663"),
		withdrawLimit:        bigFromStr("-4039971198189435924221"),
		marginCall:           bigFromStr("86627881540410570455047"),
		leverageRatio:        4.591082801378754,
		dtl:                  0.7821864594339948,
	}},
	// // agent 91
	// {"agent miner with fee debt", "0xFF65F5f3D309fEA7aA2d4cB2727E918FAb0aE7F7", BASE_BLOCK_HEIGHT, agentFiTest{
	// 	AgentFi: &AgentFi{
	// 		BaseFi: BaseFi{
	// 			Balance:          bigFromStr("0"),
	// 			AvailableBalance: bigFromStr("145838619789"),
	// 			LockedRewards:    bigFromStr("0"),
	// 			InitialPledge:    bigFromStr("0"),
	// 			FeeDebt:          bigFromStr("43697948992036495293"),
	// 			TerminationFee:   bigFromStr("0"),
	// 			LiveSectors:      big.NewInt(0),
	// 			FaultySectors:    big.NewInt(0),
	// 		},
	// 		Liability: Liability{
	// 			Principal: bigFromStr("0"),
	// 			Interest:  bigFromStr("0"),
	// 		},
	// 	},
	// 	maxBorrowAndSeal:     bigFromStr("0"),
	// 	maxBorrowAndWithdraw: bigFromStr("0"),
	// 	liquidationValue:     bigFromStr("0"),
	// 	borrowLimit:          bigFromStr("0"),
	// 	withdrawLimit:        bigFromStr("0"),
	// 	marginCall:           bigFromStr("0"),
	// 	leverageRatio:        0,
	// 	dtl:                  0,
	// }},
	{"agent with balance no miners, with principal", "0xf0F1ceCCF78D411EeD9Ca190ca7F157140cCB2d3", BASE_BLOCK_HEIGHT, agentFiTest{
		AgentFi: &AgentFi{
			BaseFi: BaseFi{
				Balance:          bigFromStr("944117375094337319"),
				AvailableBalance: bigFromStr("944117375094337319"),
				LockedRewards:    bigFromStr("0"),
				InitialPledge:    bigFromStr("0"),
				FeeDebt:          bigFromStr("0"),
				TerminationFee:   bigFromStr("0"),
				LiveSectors:      big.NewInt(0),
				FaultySectors:    big.NewInt(0),
			},
			Liability: Liability{
				Principal: bigFromStr("659045655399482142"),
				Interest:  bigFromStr("0"),
			},
		},
		maxBorrowAndSeal:     bigFromStr("793544454054283425"),
		maxBorrowAndWithdraw: bigFromStr("28485474244510146"),
		liquidationValue:     bigFromStr("944117375094337319"),
		borrowLimit:          bigFromStr("113941896978040581"),
		withdrawLimit:        bigFromStr("37980632326013527"),
		marginCall:           bigFromStr("799532420089697463"),
		leverageRatio:        3.569241913055146,
		dtl:                  0.7198284609562833,
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

	for _, tt := range agentTests {
		testname := fmt.Sprintf("%s:%s", tt.name, tt.agent)
		t.Run(testname, func(t *testing.T) {
			ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(tt.block.Uint64()), types.EmptyTSK)
			if err != nil {
				t.Fatal(err)
			}

			agent := common.HexToAddress(tt.agent)

			agentFi, err := EstimateTerminationFeeAgent(context.Background(), agent, address.Undef, psdk, ts)
			if err != nil {
				t.Fatal(err)
			}

			assertAgentFiEqual(t, &tt.want, agentFi)
		})
	}
}

func TestComputePerc(t *testing.T) {
	lv := bigFromStr("871702482492442116952350")
	p := bigFromStr("343733044178305628615912")
	i := bigFromStr("28202982363374106229")

	perc := ComputePerc(new(big.Int).Add(p, i), lv)
	exp, _ := new(big.Float).SetString("0.394356164017979036")
	if perc.String() != exp.String() {
		t.Fatalf("Expected perc: %v, actual: %v", exp, perc)
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
		t.Fatalf("Expected balance: %v, actual: %v", expected.Balance, actual.Balance)
	}
	if expected.FaultySectors.Cmp(actual.FaultySectors) != 0 {
		t.Fatalf("Expected faulty sectors: %v, actual: %v", expected.FaultySectors, actual.FaultySectors)
	}
	if expected.LiveSectors.Cmp(actual.LiveSectors) != 0 {
		t.Fatalf("Expected live sectors: %v, actual: %v", expected.LiveSectors, actual.LiveSectors)
	}

	if expected.maxBorrowAndSeal.Cmp(actual.MaxBorrowAndSeal(constants.MAX_BORROW_DTL)) != 0 {
		t.Fatalf("Expected max borrow and seal: %v, actual: %v", expected.maxBorrowAndSeal, actual.MaxBorrowAndSeal(constants.MAX_BORROW_DTL))
	}
	if expected.maxBorrowAndWithdraw.Cmp(actual.MaxBorrowAndWithdraw(constants.MAX_BORROW_DTL)) != 0 {
		t.Fatalf("Expected max borrow and withdraw: %v, actual: %v", expected.maxBorrowAndWithdraw, actual.MaxBorrowAndWithdraw(constants.MAX_BORROW_DTL))
	}
	if expected.liquidationValue.Cmp(actual.LiquidationValue()) != 0 {
		t.Fatalf("Expected liquidation value: %v, actual: %v", expected.liquidationValue, actual.LiquidationValue())
	}
	if expected.borrowLimit.Cmp(actual.BorrowLimit(constants.MAX_BORROW_DTL)) != 0 {
		t.Fatalf("Expected borrow limit: %v, actual: %v", expected.borrowLimit, actual.BorrowLimit(constants.MAX_BORROW_DTL))
	}
	if expected.withdrawLimit.Cmp(actual.WithdrawLimit(constants.MAX_BORROW_DTL)) != 0 {
		t.Fatalf("Expected withdraw limit: %v, actual: %v", expected.withdrawLimit, actual.WithdrawLimit(constants.MAX_BORROW_DTL))
	}
	if expected.marginCall.Cmp(actual.MarginCall()) != 0 {
		t.Fatalf("Expected margin call: %v, actual: %v", expected.marginCall, actual.MarginCall())
	}
	lr, _ := actual.LeverageRatio().Float64()
	if expected.leverageRatio != lr {
		t.Fatalf("Expected leverage ratio: %v, actual: %v", expected.leverageRatio, lr)
	}

	dtl, _ := actual.DTL().Float64()
	if expected.dtl != dtl {
		t.Fatalf("Expected DTL: %v, actual: %v", expected.dtl, dtl)
	}
}

func assertBaseFiEqual(t *testing.T, expected *baseFiTest, actual *MinerFi) {
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
