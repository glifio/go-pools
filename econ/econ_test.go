package econ

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	name        string
	agentAddr   common.Address
	agentID     *big.Int
	blockNumber *big.Int
}{
	{
		name:        "agent with interest",
		agentAddr:   common.HexToAddress("0x64Ea1aD49A78B6A6878c4975566b8953b1Ef1e79"),
		agentID:     big.NewInt(59),
		blockNumber: big.NewInt(4234390),
	},
	{
		name:        "empty agent",
		agentAddr:   common.HexToAddress("0x3EDA9b357Ea222F6D51a77e136a91cd8b9F461BD"),
		agentID:     big.NewInt(1),
		blockNumber: big.NewInt(4234390),
	},
}

func TestAgentInterestOwed(t *testing.T) {
	ctx := context.Background()
	psdk, err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	ethClient, err := psdk.Extern().ConnectEthClient()
	assert.NoError(t, err)
	defer ethClient.Close()

	caller, err := abigen.NewInfinityPoolV2Caller(psdk.Query().InfinityPool(), ethClient)
	assert.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interestOwed, err := caller.GetAgentInterestOwed(&bind.CallOpts{
				BlockNumber: tc.blockNumber,
			}, tc.agentID)
			assert.NoError(t, err)

			owedFromMSDK, err := psdk.Query().AgentInterestOwed(ctx, tc.agentAddr, tc.blockNumber)
			require.NoError(t, err)

			if !util.AssertApproxEqAbs(owedFromMSDK, interestOwed, big.NewInt(1)) {
				t.Fatalf("Interest owed from AgentInterestOwed should match abigen InfinityPoolV2 for %s", tc.name)
			}
		})
	}
}
