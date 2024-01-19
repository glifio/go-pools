package sdk

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/deploy"
	"github.com/glifio/go-pools/util"
)

func TestAgentCollateralStats(t *testing.T) {
	sdk, err := New(context.Background(), big.NewInt(constants.MainnetChainID), deploy.Extern)
	if err != nil {
		t.Fatal(err)
	}

	agentCollateralStats, err := sdk.Query().AgentCollateralStats(context.Background(), common.HexToAddress("0xB286f63a1B18be6772A1cF828710E594ADcAEe38"), nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(util.ToFIL(agentCollateralStats.LiquidationValue()))

	t.Fatal()
}
