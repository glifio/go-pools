package econ

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/deploy"
	"github.com/glifio/go-pools/sdk"
	"github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/vc"
)

var TIPSET_HEIGHT = big.NewInt(4198539)

func setupTest() (types.PoolsSDK, error) {
	extern := deploy.Extern
	extern.LotusDialAddr = os.Getenv("LOTUS_DIAL_ADDR")
	extern.LotusToken = os.Getenv("LOTUS_TOKEN")

	return sdk.New(context.Background(), big.NewInt(constants.MainnetChainID), extern)
}

type agentDataTest struct {
	collateralValue *big.Int
	liveSectors     *big.Int
	principal       *big.Int
}

var agentDataTests = []struct {
	name  string
	agent string
	want  agentDataTest
}{
	{"empty agent", "0xDBa96B0FDbc87C7eEb641Ee37EAFC55B355079E4", agentDataTest{
		collateralValue: bigFromStr("0"),
		liveSectors:     bigFromStr("0"),
		principal:       bigFromStr("0"),
	}},
	{"single miner agent", "0xf0F1ceCCF78D411EeD9Ca190ca7F157140cCB2d3", agentDataTest{
		collateralValue: bigFromStr("120419585663471646317"),
		liveSectors:     bigFromStr("323"),
		principal:       bigFromStr("2000000000000000000"),
	}},
	{"agent with fee debt", "0xFF65F5f3D309fEA7aA2d4cB2727E918FAb0aE7F7", agentDataTest{
		collateralValue: bigFromStr("0"),
		liveSectors:     bigFromStr("0"),
		principal:       bigFromStr("0"),
	}},
	{"multiminer agent", "0x64Ea1aD49A78B6A6878c4975566b8953b1Ef1e79", agentDataTest{
		collateralValue: bigFromStr("1882771200768867984814855"),
		liveSectors:     bigFromStr("2805488"),
		principal:       bigFromStr("1081255000003082191780823"),
	}},
}

func TestComputeAgentData(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(TIPSET_HEIGHT.Int64()), ltypes.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range agentDataTests {
		t.Run(tt.name, func(t *testing.T) {
			principal, err := psdk.Query().AgentPrincipal(context.Background(), common.HexToAddress(tt.agent), TIPSET_HEIGHT)
			if err != nil {
				t.Fatal(err)
			}

			ad, err := ComputeAgentData(context.Background(), common.HexToAddress(tt.agent), big.NewInt(0), principal, address.Undef, psdk, ts)
			if err != nil {
				t.Fatal(err)
			}

			testAgentData(t, ad, tt.want.liveSectors, tt.want.collateralValue, tt.want.principal)
		})
	}
}

var borrowAmount = big.NewInt(1e18)

func TestComputeBorrowAgentData(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(TIPSET_HEIGHT.Int64()), ltypes.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range agentDataTests {
		t.Run(tt.name, func(t *testing.T) {
			ad, err := ComputeBorrowAgentData(context.Background(), common.HexToAddress(tt.agent), borrowAmount, psdk, ts)
			if err != nil {
				t.Fatal(err)
			}

			expectedLiveSectors := tt.want.liveSectors
			expectedCollateralValue := new(big.Int).Add(tt.want.collateralValue, borrowAmount)
			expectedPrincipal := new(big.Int).Add(tt.want.principal, borrowAmount)

			testAgentData(t, ad, expectedLiveSectors, expectedCollateralValue, expectedPrincipal)
		})
	}
}

func TestComputeRemoveMinerAgentData(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(TIPSET_HEIGHT.Int64()), ltypes.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	// multiminer agent
	agentAddr := common.HexToAddress(agentDataTests[3].agent)
	minerToRm, _ := address.NewFromString("f03055018")

	termRes, err := EstimateTerminationFeeMiner(context.Background(), lapi, minerToRm, ts)
	if err != nil {
		t.Fatal(err)
	}

	ad, err := ComputeRmMinerAgentData(context.Background(), agentAddr, minerToRm, psdk, ts)
	if err != nil {
		t.Fatal(err)
	}

	expectedLiveSectors := new(big.Int).Sub(agentDataTests[3].want.liveSectors, big.NewInt(int64(termRes.LiveSectors)))
	expectedCollateralValue := new(big.Int).Sub(agentDataTests[3].want.collateralValue, termRes.ToBaseFi().LiquidationValue())
	expectedPrincipal := agentDataTests[3].want.principal

	testAgentData(t, ad, expectedLiveSectors, expectedCollateralValue, expectedPrincipal)
}

func TestComputeRemoveSingleMiner(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(TIPSET_HEIGHT.Int64()), ltypes.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	// multiminer agent
	agentAddr := common.HexToAddress(agentDataTests[1].agent)
	minerToRm, _ := address.NewFromString("f01931245")

	ad, err := ComputeRmMinerAgentData(context.Background(), agentAddr, minerToRm, psdk, ts)
	if err != nil {
		t.Fatal(err)
	}

	agentLiquidAssets, err := psdk.Query().AgentLiquidAssets(context.Background(), agentAddr, TIPSET_HEIGHT)
	if err != nil {
		t.Fatal(err)
	}

	expectedLiveSectors := big.NewInt(0)
	expectedCollateralValue := agentLiquidAssets
	expectedPrincipal := agentDataTests[1].want.principal

	testAgentData(t, ad, expectedLiveSectors, expectedCollateralValue, expectedPrincipal)

}

var withdrawAmount = big.NewInt(1e18)

func TestComputeWithdrawAgentData(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(TIPSET_HEIGHT.Int64()), ltypes.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	// multiminer agent
	agentAddr := common.HexToAddress(agentDataTests[3].agent)

	ad, err := ComputeWithdrawAgentData(context.Background(), agentAddr, withdrawAmount, psdk, ts)
	if err != nil {
		t.Fatal(err)
	}

	expectedLiveSectors := agentDataTests[3].want.liveSectors
	expectedCollateralValue := new(big.Int).Sub(agentDataTests[3].want.collateralValue, withdrawAmount)
	expectedPrincipal := agentDataTests[3].want.principal

	testAgentData(t, ad, expectedLiveSectors, expectedCollateralValue, expectedPrincipal)
}

func TestComputeWithdrawTooMuch(t *testing.T) {
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(TIPSET_HEIGHT.Int64()), ltypes.EmptyTSK)
	if err != nil {
		t.Fatal(err)
	}

	// multiminer agent
	agentAddr := common.HexToAddress(agentDataTests[3].agent)

	// withdraw a lot
	withdrawAmount := new(big.Int).Mul(big.NewInt(1e18), withdrawAmount)

	_, err = ComputeWithdrawAgentData(context.Background(), agentAddr, withdrawAmount, psdk, ts)
	if err == nil {
		t.Fatal("Expected error - withdraw amount is greater than agent's available balance")
	}
}

func testAgentData(t *testing.T, ad *vc.AgentData, expectedLiveSectors *big.Int, expectedCollateralValue *big.Int, expectedPrincipal *big.Int) {
	if ad.Principal.Cmp(expectedPrincipal) != 0 {
		t.Fatalf("principal comparison failed. have: %v, want: %v", ad.Principal, expectedPrincipal)
	}
	if ad.LiveSectors.Cmp(expectedLiveSectors) != 0 {
		t.Fatalf("live sectors comparison failed. have: %v, want: %v", ad.LiveSectors, expectedLiveSectors)
	}
	if ad.CollateralValue.Cmp(expectedCollateralValue) != 0 {
		t.Fatalf("collateral value comparison failed. have: %v, want: %v", ad.CollateralValue, expectedCollateralValue)
	}
	if ad.AgentValue.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("expected 0 AgentValue")
	}
	if ad.ExpectedDailyFaultPenalties.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("expected 0 ExpectedDailyFaultPenalties")
	}
	if ad.ExpectedDailyRewards.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("expected 0 ExpectedDailyRewards")
	}
	if ad.Gcred.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("expected 0 Gcred")
	}
	if ad.GreenScore.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("expected 0 GreenScore")
	}
	if ad.FaultySectors.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("expected 0 FaultySectors")
	}
	if ad.QaPower.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("expected QaPower")
	}
}
