package econ

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	pooltypes "github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/util"
	"github.com/glifio/go-pools/vc"
)

func GetAgentEcon(ctx context.Context, agentAddr common.Address, psdk pooltypes.PoolsSDK, ts *types.TipSet) (*big.Int, *big.Int, error) {
	query := psdk.Query()
	blockNumber := big.NewInt(int64(ts.Height()))

	agentAvailableBal, err := query.AgentLiquidAssets(ctx, agentAddr, blockNumber)
	if err != nil {
		return nil, nil, err
	}

	principal, err := getAgentPrincipal(ctx, agentAddr, psdk, ts)
	if err != nil {
		return nil, nil, err
	}

	return agentAvailableBal, principal, nil
}

func getAgentPrincipal(ctx context.Context, agentAddr common.Address, psdk pooltypes.PoolsSDK, ts *types.TipSet) (*big.Int, error) {
	blockNumber := big.NewInt(int64(ts.Height()))
	principal, err := psdk.Query().AgentPrincipal(ctx, agentAddr, blockNumber)
	if err != nil {
		return nil, err
	}

	return principal, nil
}

func ComputeBorrowAgentData(ctx context.Context, agentAddr common.Address, amount *big.Int, psdk pooltypes.PoolsSDK, ts *types.TipSet) (*vc.AgentData, error) {
	principal, err := getAgentPrincipal(ctx, agentAddr, psdk, ts)
	if err != nil {
		return nil, err
	}

	newPrincipal := big.NewInt(0).Add(principal, amount)

	data, err := ComputeAgentData(ctx, agentAddr, amount, newPrincipal, address.Undef, psdk, ts)
	if err != nil {
		return nil, err
	}

	return data, err
}

func ComputePayAgentData(ctx context.Context, agentAddr common.Address, value *big.Int, psdk pooltypes.PoolsSDK, ts *types.TipSet) (*vc.AgentData, error) {
	agentAvailableBal, principal, err := GetAgentEcon(ctx, agentAddr, psdk, ts)
	if err != nil {
		return nil, err
	}

	if condition := value.Cmp(agentAvailableBal); condition == 1 {
		return nil, errors.New("pay amount is greater than agent's available balance")
	}

	removingBal := big.NewInt(0).Sub(big.NewInt(0), value)

	return ComputeAgentData(ctx, agentAddr, removingBal, principal, address.Undef, psdk, ts)
}

func ComputeWithdrawAgentData(ctx context.Context, agentAddr common.Address, withdrawAmount *big.Int, psdk pooltypes.PoolsSDK, ts *types.TipSet) (*vc.AgentData, error) {
	agentAvailableBal, principal, err := GetAgentEcon(ctx, agentAddr, psdk, ts)
	if err != nil {
		return nil, err
	}

	// check if amount is greater than agentAvailableBal
	if condition := withdrawAmount.Cmp(agentAvailableBal); condition == 1 {
		return nil, errors.New("withdraw amount is greater than agent's available balance")
	}

	// here we want to pass a negative value to ComputeAgentData to indicate that the agent's balance is being withdrawn
	removingVal := big.NewInt(0).Sub(big.NewInt(0), withdrawAmount)

	data, err := ComputeAgentData(ctx, agentAddr, removingVal, principal, address.Undef, psdk, ts)
	if err != nil {
		return nil, err
	}

	return data, err
}

func ComputePushFundsAgentData(ctx context.Context, agentAddr common.Address, minerAddr address.Address, psdk pooltypes.PoolsSDK, ts *types.TipSet) (*vc.AgentData, error) {
	principal, err := getAgentPrincipal(ctx, agentAddr, psdk, ts)
	if err != nil {
		return nil, err
	}

	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	// make sure we aren't pushing a miner with fee debt
	_, mstate, err := util.LoadMinerActor(ctx, lapi, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	feeDebt, err := mstate.FeeDebt()
	if err != nil {
		return nil, err
	}

	if feeDebt.Sign() == 1 {
		return nil, errors.New("cannot push funds to miner with fee debt")
	}

	return ComputeAgentData(ctx, agentAddr, big.NewInt(0), principal, address.Undef, psdk, ts)
}

func ComputeRmMinerAgentData(ctx context.Context, agentAddr common.Address, miner address.Address, psdk pooltypes.PoolsSDK, ts *types.TipSet) (*vc.AgentData, error) {
	principal, err := getAgentPrincipal(ctx, agentAddr, psdk, ts)
	if err != nil {
		return nil, err
	}

	return ComputeAgentData(ctx, agentAddr, big.NewInt(0), principal, miner, psdk, ts)
}
