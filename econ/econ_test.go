package econ

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/big"
// 	"os"
// 	"testing"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/filecoin-project/go-address"
// 	biggy "github.com/filecoin-project/go-state-types/big"
// 	"github.com/filecoin-project/lotus/chain/types"
// 	"github.com/glif-confidential/ado/singleton"
// 	"github.com/glifio/go-pools/vc"
// 	"github.com/spf13/viper"
// )

// var epoch = int64(567326)

// func setupSuite(tb testing.TB, config string) (func(tb testing.TB), *types.TipSet) {
// 	viper.SetConfigFile(config)
// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}

// 	// if err := chain.ConnectLotus(chain.ChainOptions{
// 		// DialAddr: viper.GetString("lotus_addr"),
// 	// 	Token:    viper.GetString("lotus_token"),
// 	// }); err != nil {
// 	// 	tb.Errorf("Expected no error, received %v", err)
// 	// }

// 	if err := singleton.ConnectLotus(singleton.ChainOptions{
// 		DialAddr: viper.GetString("lotus_addr"),
// 		Token:    viper.GetString("lotus_token"),
// 	}); err != nil {
// 		tb.Errorf("Expected no error, received %v", err)
// 	}

// 	log.Printf("epocs: %v", epoch)
// 	// defer chain.LotusClient.Close()
// 	ts, err := singleton.Lotus().Api.ChainHead(context.Background())
// 	// ts, err := singleton.Lotus().Api.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(epoch), types.EmptyTSK)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// rpcurl := "http://127.0.0.1:8545"
// 	// privatekey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
// 	// router := "0x5FbDB2315678afecb367f032d93F642f64180aa3"

// 	// connectParams, err := eth_utils.NewConnectParams(
// 	// 	rpcurl,
// 	// 	privatekey,
// 	// 	router,
// 	// )

// 	// owner := "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
// 	// operator := "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
// 	// contract_utils.AgentCreate(&connectParams, owner, operator)

// 	// Return a function to teardown the test
// 	return func(tb testing.TB) {
// 		// chain.LotusClient.Close()
// 	}, ts
// }

// func setupTest(tb testing.TB) func(tb testing.TB) {
// 	return func(tb testing.TB) {
// 	}
// }

// func MustAddressFromString(addr string) address.Address {
// 	res, err := address.NewFromString(addr)
// 	if err != nil {
// 		log.Panic("big int failed")
// 	}
// 	return res
// }

// func MustNewBigInt(val string) *big.Int {
// 	res, success := big.NewInt(0).SetString(val, 10)
// 	if !success {
// 		log.Panic("big int failed")
// 	}
// 	return res
// }

// // func TestAggAssets(t *testing.T) {
// // 	teardownSuite, ts := setupSuite(t)
// // 	defer teardownSuite(t)

// // 	var tests = []struct {
// // 		addr address.Address
// // 		want *big.Int
// // 	}{
// // 		{MustAddressFromString("f01491738"), MustNewBigInt("81426425080002391402032")},
// // 		{MustAddressFromString("f0440919"), MustNewBigInt("329088353412830271096879")},
// // 		{MustAddressFromString("f078770"), MustNewBigInt("150230014031415685918284")},
// // 		{MustAddressFromString("f0427688"), MustNewBigInt("660018902332302912698762")},
// // 		{MustAddressFromString("f0501283"), MustNewBigInt("9517348124497036602373")},
// // 	}

// // 	for _, tt := range tests {
// // 		testname := fmt.Sprintf("%s,%s", tt.addr, tt.want)
// // 		teardownTest := setupTest(t)
// // 		defer teardownTest(t)
// // 		t.Run(testname, func(t *testing.T) {
// // 			answer, err := ScoreAggAssets(context.Background(), []address.Address{tt.addr}, ts)
// // 			if err != nil {
// // 				t.Errorf("ScoreAggAssets returned: %s", err)
// // 			}
// // 			if answer.Cmp(tt.want) != 0 {
// // 				t.Errorf("Expected %v, received %v", tt.want, answer)
// // 			}
// // 		})
// // 	}
// // }

// // func TestAggLockedAssets(t *testing.T) {
// // 	teardownSuite, ts := setupSuite(t)
// // 	defer teardownSuite(t)

// // 	var tests = []struct {
// // 		addr address.Address
// // 		want *LockedFunds
// // 	}{
// // 		{MustAddressFromString("f01491738"), &LockedFunds{
// // 			Balance:        MustNewBigInt("71120045605367595310007"),
// // 			InitialPledge:  MustNewBigInt("61056637385738071048442"),
// // 			VestingRewards: MustNewBigInt("10063408219629524261565")}},
// // 		{MustAddressFromString("f0440919"), &LockedFunds{
// // 			Balance:        MustNewBigInt("317314117491004651346884"),
// // 			InitialPledge:  MustNewBigInt("264453745959588296229054"),
// // 			VestingRewards: MustNewBigInt("52860371531416355117830")}},
// // 		{MustAddressFromString("f078770"), &LockedFunds{
// // 			Balance:        MustNewBigInt("150051442053291647154042"),
// // 			InitialPledge:  MustNewBigInt("127749760293281889092692"),
// // 			VestingRewards: MustNewBigInt("22301681760009758061350")}},
// // 		{MustAddressFromString("f0427688"), &LockedFunds{
// // 			Balance:        MustNewBigInt("656438524298914131498979"),
// // 			InitialPledge:  MustNewBigInt("575424691545850276702113"),
// // 			VestingRewards: MustNewBigInt("81013832753063854796866")}},
// // 	}

// // 	for _, tt := range tests {
// // 		testname := tt.addr.String()
// // 		t.Run(testname, func(t *testing.T) {
// // 			answer, err := ScoreAggLockedFunds(context.Background(), []address.Address{tt.addr}, ts)
// // 			if err != nil {
// // 				t.Errorf("ScoreAggLockedFunds returned: %s", err)
// // 			}
// // 			if answer.Balance.Cmp(tt.want.Balance) != 0 {
// // 				t.Errorf("Expected %v, received %v", tt.want.Balance, answer.Balance)
// // 			}
// // 			if answer.InitialPledge.Cmp(tt.want.InitialPledge) != 0 {
// // 				t.Errorf("Expected %v, received %v", tt.want.InitialPledge, answer.InitialPledge)
// // 			}
// // 			if answer.VestingRewards.Cmp(tt.want.VestingRewards) != 0 {
// // 				t.Errorf("Expected %v, received %v", tt.want.VestingRewards, answer.VestingRewards)
// // 			}
// // 		})
// // 	}
// // }

// // func TestAggPenaltyForTermination(t *testing.T) {
// // 	teardownSuite, ts := setupSuite(t)
// // 	defer teardownSuite(t)

// // 	var tests = []struct {
// // 		addr address.Address
// // 		want *big.Int
// // 	}{
// // 		{MustAddressFromString("f0501283"), MustNewBigInt("2922495706093776305685")},
// // 		{MustAddressFromString("f01938357"), MustNewBigInt("12413675509966745230972")},
// // 		{MustAddressFromString("f01491738"), MustNewBigInt("22384228806452773215101")},
// // 		// {MustAddressFromString("f0440919"), MustNewBigInt("329088353412830271096879")},  // slow
// // 		{MustAddressFromString("f078770"), MustNewBigInt("37163083654174384557328")},
// // 		// {MustAddressFromString("f0427688"), MustNewBigInt("660018902332302912698762")},  // slow
// // 	}

// // 	for _, tt := range tests {
// // 		testname := tt.addr.String()
// // 		t.Run(testname, func(t *testing.T) {
// // 			answer, err := ScoreAggPenaltyForTermination(context.Background(), []address.Address{tt.addr}, ts)
// // 			if err != nil {
// // 				t.Errorf("ScoreAggCollateralValue returned: %s", err)
// // 			}
// // 			if answer.PenaltyTermination.Cmp(tt.want) != 0 {
// // 				t.Errorf("Expected %v, received %v", tt.want, answer.PenaltyTermination)
// // 			}
// // 		})
// // 	}
// // }

// func TestAgent(t *testing.T) {
// 	teardownSuite, ts := setupSuite(t, "../testnet.env")
// 	defer teardownSuite(t)

// 	var tests = []struct {
// 		addr address.Address
// 		want *big.Int
// 	}{
// 		{MustAddressFromString("f0501283"), MustNewBigInt("0")},
// 		{MustAddressFromString("f01938357"), MustNewBigInt("0")},
// 		{MustAddressFromString("f01491738"), MustNewBigInt("0")},
// 		{MustAddressFromString("f0440919"), MustNewBigInt("0")},
// 		{MustAddressFromString("f078770"), MustNewBigInt("0")},
// 		{MustAddressFromString("f0427688"), MustNewBigInt("0")},
// 	}

// 	for _, tt := range tests {
// 		testname := tt.addr.String()
// 		t.Run(testname, func(t *testing.T) {
// 			log.Printf("Skipping test for now %s", ts.Height().String())
// 			// answer, err := ScoreAgent(context.Background(), []address.Address{tt.addr}, ts)
// 			// if err != nil {
// 			// 	t.Errorf("ScoreAgent returned: %s", err)
// 			// }
// 			// if answer.Cmp(tt.want) != 0 {
// 			// 	t.Errorf("Expected %v, received %v", tt.want, answer)
// 			// }
// 		})
// 	}
// }

// func TestComputeAgentDataTestnet(t *testing.T) {
// 	teardownSuite, ts := setupSuite(t, "../testnet.env")
// 	defer teardownSuite(t)

// 	var tests = []struct {
// 		agent       common.Address
// 		assets      *big.Int
// 		liabilities *big.Int
// 		miners      []address.Address
// 		want        *vc.AgentData
// 	}{
// 		{common.HexToAddress("0xc5D6b79d96dc8EBf7caBE9a17af5Cf8DB80Cc8f2"),
// 			MustNewBigInt("0"),
// 			MustNewBigInt("0"),
// 			[]address.Address{MustAddressFromString("t06024")},
// 			&vc.AgentData{
// 				AgentValue:                  MustNewBigInt("152162778917550758427699"),
// 				Principal:                   MustNewBigInt("0"),
// 				ExpectedDailyRewards:        MustNewBigInt("6892136367183830803456"),
// 				CollateralValue:             MustNewBigInt("76081389458775379213850"),
// 				QaPower:                     MustNewBigInt("155683974545408"),
// 				ExpectedDailyFaultPenalties: MustNewBigInt("20676409101551492410368"),
// 				Gcred:                       MustNewBigInt("90"),
// 				FaultySectors:               MustNewBigInt("0"),
// 				LiveSectors:                 MustNewBigInt("4531"),
// 				GreenScore:                  MustNewBigInt("0"),
// 			},
// 		},
// 		{common.HexToAddress("0xc5D6b79d96dc8EBf7caBE9a17af5Cf8DB80Cc8f2"),
// 			MustNewBigInt("0"),
// 			MustNewBigInt("0"),
// 			[]address.Address{MustAddressFromString("t016539")},
// 			&vc.AgentData{
// 				AgentValue:                  MustNewBigInt("283638225449250911915"),
// 				Principal:                   MustNewBigInt("0"),
// 				ExpectedDailyRewards:        MustNewBigInt("20595935426884886736"),
// 				CollateralValue:             MustNewBigInt("141819112724625455958"),
// 				QaPower:                     MustNewBigInt("377957122048"),
// 				ExpectedDailyFaultPenalties: MustNewBigInt("61787806280654660208"),
// 				Gcred:                       MustNewBigInt("90"),
// 				FaultySectors:               MustNewBigInt("0"),
// 				LiveSectors:                 MustNewBigInt("11"),
// 				GreenScore:                  MustNewBigInt("0"),
// 			},
// 		},
// 		{common.HexToAddress("0xc5D6b79d96dc8EBf7caBE9a17af5Cf8DB80Cc8f2"),
// 			MustNewBigInt("0"),
// 			MustNewBigInt("0"),
// 			[]address.Address{MustAddressFromString("t01002")},
// 			&vc.AgentData{
// 				AgentValue:                  MustNewBigInt("2781608154511004151681615"),
// 				Principal:                   MustNewBigInt("0"),
// 				ExpectedDailyRewards:        MustNewBigInt("10200005652377628198912"),
// 				CollateralValue:             MustNewBigInt("1390804077255502075840808"),
// 				QaPower:                     MustNewBigInt("263882790666240"),
// 				ExpectedDailyFaultPenalties: MustNewBigInt("30600016957132884596736"),
// 				Gcred:                       MustNewBigInt("90"),
// 				FaultySectors:               MustNewBigInt("0"),
// 				LiveSectors:                 MustNewBigInt("768"),
// 				GreenScore:                  MustNewBigInt("0"),
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("%s", tt.agent)
// 		teardownTest := setupTest(t)
// 		defer teardownTest(t)
// 		t.Run(testname, func(t *testing.T) {
// 			answer, err := ComputeAgentData(context.Background(), tt.assets, tt.liabilities, tt.miners, tt.agent, ts)
// 			if err != nil {
// 				t.Errorf("ComputeAgentData returned: %s", err)
// 			}
// 			if answer.AgentValue.Cmp(tt.want.AgentValue) != 0 {
// 				t.Errorf("Agent value - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.AgentValue)).String(), types.FIL(biggy.NewFromGo(answer.AgentValue)).String())
// 			}
// 			if answer.CollateralValue.Cmp(tt.want.CollateralValue) != 0 {
// 				t.Errorf("Collateral value - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.CollateralValue)).String(), types.FIL(biggy.NewFromGo(answer.CollateralValue)).String())
// 			}
// 			if answer.ExpectedDailyRewards.Cmp(tt.want.ExpectedDailyRewards) != 0 {
// 				t.Errorf("DailyRewards - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.ExpectedDailyRewards)).String(), types.FIL(biggy.NewFromGo(answer.ExpectedDailyRewards)).String())
// 			}
// 			if answer.Principal.Cmp(tt.want.Principal) != 0 {
// 				t.Errorf("Principal - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.Principal)).String(), types.FIL(biggy.NewFromGo(answer.Principal)).String())
// 			}
// 			if answer.QaPower.Cmp(tt.want.QaPower) != 0 {
// 				t.Errorf("QaPower - Expected %v, received %v", types.DeciStr(biggy.NewFromGo(tt.want.QaPower)), types.DeciStr(biggy.NewFromGo(answer.QaPower)))
// 				t.Errorf("QaPower - Expected %v, received %v", tt.want.QaPower, answer.QaPower)
// 			}
// 			if answer.ExpectedDailyFaultPenalties.Cmp(tt.want.ExpectedDailyFaultPenalties) != 0 {
// 				t.Errorf("ExpectedDailyFaultPenalties - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.ExpectedDailyFaultPenalties)).String(), types.FIL(biggy.NewFromGo(answer.ExpectedDailyFaultPenalties)).String())
// 			}
// 			if answer.Gcred.Cmp(tt.want.Gcred) != 0 {
// 				t.Errorf("Gcred - Expected %v, received %v", tt.want.Gcred, answer.Gcred)
// 			}
// 			if answer.FaultySectors.Cmp(tt.want.FaultySectors) != 0 {
// 				t.Errorf("FaultySectors - Expected %v, received %v", tt.want.FaultySectors, answer.FaultySectors)
// 			}
// 			if answer.LiveSectors.Cmp(tt.want.LiveSectors) != 0 {
// 				t.Errorf("LiveSectors - Expected %v, received %v", tt.want.LiveSectors, answer.LiveSectors)
// 			}
// 			if answer.GreenScore.Cmp(tt.want.GreenScore) != 0 {
// 				t.Errorf("GreenScore - Expected %v, received %v", tt.want.GreenScore, answer.GreenScore)
// 			}
// 		})
// 	}
// }

// func TestComputeAgentDataMainnet(t *testing.T) {
// 	teardownSuite, ts := setupSuite(t, "../mainnet.env")
// 	defer teardownSuite(t)

// 	var tests = []struct {
// 		agent       common.Address
// 		assets      *big.Int
// 		liabilities *big.Int
// 		miners      []address.Address
// 		want        *vc.AgentData
// 	}{
// 		{common.HexToAddress("0xc5D6b79d96dc8EBf7caBE9a17af5Cf8DB80Cc8f2"),
// 			MustNewBigInt("0"),
// 			MustNewBigInt("0"),
// 			[]address.Address{MustAddressFromString("f0501283")},
// 			&vc.AgentData{
// 				AgentValue:                  MustNewBigInt("152162778917550758427699"),
// 				Principal:                   MustNewBigInt("0"),
// 				ExpectedDailyRewards:        MustNewBigInt("6892136367183830803456"),
// 				CollateralValue:             MustNewBigInt("76081389458775379213850"),
// 				QaPower:                     MustNewBigInt("155683974545408"),
// 				ExpectedDailyFaultPenalties: MustNewBigInt("20676409101551492410368"),
// 				Gcred:                       MustNewBigInt("90"),
// 				FaultySectors:               MustNewBigInt("0"),
// 				LiveSectors:                 MustNewBigInt("4531"),
// 				GreenScore:                  MustNewBigInt("0"),
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("%s", tt.agent)
// 		teardownTest := setupTest(t)
// 		defer teardownTest(t)
// 		t.Run(testname, func(t *testing.T) {
// 			answer, err := ComputeAgentData(context.Background(), tt.assets, tt.liabilities, tt.miners, tt.agent, ts)
// 			if err != nil {
// 				t.Errorf("ComputeAgentData returned: %s", err)
// 			}
// 			if answer.AgentValue.Cmp(tt.want.AgentValue) != 0 {
// 				t.Errorf("Agent value - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.AgentValue)).String(), types.FIL(biggy.NewFromGo(answer.AgentValue)).String())
// 			}
// 			if answer.CollateralValue.Cmp(tt.want.CollateralValue) != 0 {
// 				t.Errorf("Collateral value - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.CollateralValue)).String(), types.FIL(biggy.NewFromGo(answer.CollateralValue)).String())
// 			}
// 			if answer.ExpectedDailyRewards.Cmp(tt.want.ExpectedDailyRewards) != 0 {
// 				t.Errorf("DailyRewards - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.ExpectedDailyRewards)).String(), types.FIL(biggy.NewFromGo(answer.ExpectedDailyRewards)).String())
// 			}
// 			if answer.Principal.Cmp(tt.want.Principal) != 0 {
// 				t.Errorf("Principal - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.Principal)).String(), types.FIL(biggy.NewFromGo(answer.Principal)).String())
// 			}
// 			if answer.QaPower.Cmp(tt.want.QaPower) != 0 {
// 				t.Errorf("QaPower - Expected %v, received %v", types.DeciStr(biggy.NewFromGo(tt.want.QaPower)), types.DeciStr(biggy.NewFromGo(answer.QaPower)))
// 				t.Errorf("QaPower - Expected %v, received %v", tt.want.QaPower, answer.QaPower)
// 			}
// 			if answer.ExpectedDailyFaultPenalties.Cmp(tt.want.ExpectedDailyFaultPenalties) != 0 {
// 				t.Errorf("ExpectedDailyFaultPenalties - Expected %v, received %v", types.FIL(biggy.NewFromGo(tt.want.ExpectedDailyFaultPenalties)).String(), types.FIL(biggy.NewFromGo(answer.ExpectedDailyFaultPenalties)).String())
// 			}
// 			if answer.Gcred.Cmp(tt.want.Gcred) != 0 {
// 				t.Errorf("Gcred - Expected %v, received %v", tt.want.Gcred, answer.Gcred)
// 			}
// 			if answer.FaultySectors.Cmp(tt.want.FaultySectors) != 0 {
// 				t.Errorf("FaultySectors - Expected %v, received %v", tt.want.FaultySectors, answer.FaultySectors)
// 			}
// 			if answer.LiveSectors.Cmp(tt.want.LiveSectors) != 0 {
// 				t.Errorf("LiveSectors - Expected %v, received %v", tt.want.LiveSectors, answer.LiveSectors)
// 			}
// 			if answer.GreenScore.Cmp(tt.want.GreenScore) != 0 {
// 				t.Errorf("GreenScore - Expected %v, received %v", tt.want.GreenScore, answer.GreenScore)
// 			}
// 		})
// 	}
// }
