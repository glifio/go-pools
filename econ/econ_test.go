package econ

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/mock"
	"github.com/glifio/go-pools/mstat"
	"github.com/glifio/go-pools/terminate"
	"github.com/glifio/go-pools/util"

	mockery "github.com/stretchr/testify/mock"
)

func TestComputeAgentDataSimple(t *testing.T) {
	// mock structs with on-chain data to compute the agent data from
	termPenalty := big.NewInt(100)
	minerPledged := big.NewInt(1000)
	minerVesting := big.NewInt(100)
	minerAvail := big.NewInt(100)
	minerBal := new(big.Int).Add(minerPledged, minerVesting)
	minerBal.Add(minerBal, minerAvail)
	agentAvail := big.NewInt(100)
	principal := big.NewInt(100)

	edr := big.NewInt(1)
	qap := big.NewInt(10000)

	mockATS := terminate.PreviewAgentTerminationSummary{
		TerminationPenalty: termPenalty,
		InitialPledge:      minerPledged,
		VestingBalance:     minerVesting,
		MinersAvailableBal: minerAvail,
		AgentAvailableBal:  agentAvail,
	}

	mockMS := &mstat.MinerStats{
		MinerInfo:           &api.MinerInfo{},
		Balance:             minerBal,
		PenaltyTermination:  termPenalty,
		ExpectedDailyReward: edr,
		PenaltyFaultPerDay:  big.NewInt(0),
		PledgedFunds:        minerPledged,
		VestingFunds:        minerVesting,
		GreenScore:          big.NewInt(0),
		QualityAdjPower:     qap,
		LiveSectors:         big.NewInt(100),
		FaultySectors:       big.NewInt(0),
		HasMinPower:         true,
	}

	mockFEVMQueries := mock.NewFEVMQueries(t)
	mockFEVMQueries.EXPECT().AgentPreviewTerminationPrecise(
		mockery.Anything,
		mockery.Anything,
		mockery.Anything,
	).Return(mockATS, nil)

	sdk := mock.NewPoolsSDK(t)
	sdk.On("Query").Return(mockFEVMQueries)

	agentData, err := ComputeAgentData(
		context.Background(),
		sdk,
		agentAvail,
		principal,
		mockMS,
		common.HexToAddress(""),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	expAgentVal := new(big.Int).Add(agentAvail, minerBal)
	if agentData.AgentValue.Cmp(expAgentVal) != 0 {
		t.Errorf("Expected %v, received %v", expAgentVal, agentData.AgentValue)
	}
	if agentData.CollateralValue.Cmp(mockATS.LiquidationValue()) != 0 {
		t.Errorf("Expected %v, received %v", mockATS.LiquidationValue(), agentData.CollateralValue)
	}
	if agentData.ExpectedDailyRewards.Cmp(edr) != 0 {
		t.Errorf("Expected %v, received %v", edr, agentData.ExpectedDailyRewards)
	}
	if agentData.QaPower.Cmp(qap) != 0 {
		t.Errorf("Expected %v, received %v", qap, agentData.QaPower)
	}
	if agentData.Principal.Cmp(principal) != 0 {
		t.Errorf("Expected %v, received %v", principal, agentData.Principal)
	}
}

func TestFaultySectorsEqualLiveWhenOverLTV(t *testing.T) {
	// create a liquidation value that is less than the principal
	// 50% recovery
	termPenalty := big.NewInt(100)
	minerPledged := big.NewInt(200)
	minerVesting := big.NewInt(0)
	minerAvail := big.NewInt(0)
	minerBal := new(big.Int).Add(minerPledged, minerVesting)
	minerBal.Add(minerBal, minerAvail)
	agentAvail := big.NewInt(0)
	principal := big.NewInt(1000)

	edr := big.NewInt(1)
	qap := big.NewInt(10000)

	mockATS := terminate.PreviewAgentTerminationSummary{
		TerminationPenalty: termPenalty,
		InitialPledge:      minerPledged,
		VestingBalance:     minerVesting,
		MinersAvailableBal: minerAvail,
		AgentAvailableBal:  agentAvail,
	}

	liveSectors := big.NewInt(100)

	mockMS := &mstat.MinerStats{
		MinerInfo:           &api.MinerInfo{},
		Balance:             minerBal,
		PenaltyTermination:  termPenalty,
		ExpectedDailyReward: edr,
		PenaltyFaultPerDay:  big.NewInt(0),
		PledgedFunds:        minerPledged,
		VestingFunds:        minerVesting,
		GreenScore:          big.NewInt(0),
		QualityAdjPower:     qap,
		LiveSectors:         liveSectors,
		FaultySectors:       big.NewInt(0),
		HasMinPower:         true,
	}

	mockFEVMQueries := mock.NewFEVMQueries(t)
	mockFEVMQueries.EXPECT().AgentPreviewTerminationPrecise(
		mockery.Anything,
		mockery.Anything,
		mockery.Anything,
	).Return(mockATS, nil)

	sdk := mock.NewPoolsSDK(t)
	sdk.On("Query").Return(mockFEVMQueries)

	agentData, err := ComputeAgentData(
		context.Background(),
		sdk,
		agentAvail,
		principal,
		mockMS,
		common.HexToAddress(""),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if mockATS.LiquidationValue().Cmp(principal) != -1 {
		t.Errorf("Expected liquidation value to be less than principal")
	}

	if agentData.FaultySectors.Cmp(liveSectors) != 0 {
		t.Errorf("Expected faulty sectors to be equal to live sectors")
	}
}

// this test uses the actual SDK to compute the agent data and compare the balances
func TestMinerStatsBalanceInvariant(t *testing.T) {
	lapi, closer := util.SetupSuite(t)
	defer util.TeardownSuite(closer)

	chainHead, err := lapi.ChainHead(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// lookback 3 blocks to avoid reorgs
	ts, err := lapi.ChainGetTipSetByHeight(context.Background(), chainHead.Height()-3, types.EmptyTSK)

	agentID := big.NewInt(2)

	acs, err := terminate.FetchAgentCollateralStats(context.Background(), agentID)
	if err != nil {
		t.Fatal(err)
	}

	ats := acs.Summarize()

	bal := new(big.Int).Add(ats.InitialPledge, ats.MinersAvailableBal)
	bal.Add(bal, ats.VestingBalance)

	// using Jim's miner
	minerAddr, err := address.NewFromString("f01931245")
	if err != nil {
		t.Fatal(err)
	}

	ms, err := mstat.ComputeMinersStats(
		context.Background(),
		[]address.Address{minerAddr},
		ts,
		lapi,
	)
	if err != nil {
		t.Fatal(err)
	}

	if ms.Balance.Cmp(bal) != 0 {
		t.Errorf("Expected %v, received %v", bal, ms.Balance)
	}
}
