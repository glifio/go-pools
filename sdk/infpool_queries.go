package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/util"
)

func (q *fevmQueries) InfPoolGetAccount(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (abigen.Account, error) {
	return q.AgentAccount(ctx, agentAddr, constants.INFINITY_POOL_ID, blockNumber)
}

func (q *fevmQueries) InfPoolGetAgentLvl(ctx context.Context, agentID *big.Int) (*big.Int, float64, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, 0, err
	}
	defer client.Close()

	police, err := abigen.NewAgentPoliceV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, 0, err
	}

	lvl, err := police.AccountLevel(nil, agentID)
	if err != nil {
		return nil, 0, err
	}

	cap, err := police.Levels(nil, lvl)
	if err != nil {
		return nil, 0, err
	}

	capInFIL, _ := util.ToFIL(cap).Float64()

	return lvl, capInFIL, nil
}

func (q *fevmQueries) InfPoolGetRate(ctx context.Context, cred abigen.VerifiableCredential) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	infCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, ethClient)
	if err != nil {
		return nil, err
	}

	rate, err := infCaller.GetRate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return rate, nil
}

func (q *fevmQueries) InfPoolTotalAssets(ctx context.Context, blockNumber *big.Int) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolBorrowableLiquidity(ctx context.Context, blockNumber *big.Int) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalBorrowableAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolTotalBorrowed(ctx context.Context, blockNumber *big.Int) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalBorrowed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolExitReserve(ctx context.Context, blockNumber *big.Int) (*big.Int, *big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, nil, err
	}

	minLiquidity, err := poolCaller.GetAbsMinLiquidity(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, nil, err
	}

	liquidAssets, err := poolCaller.GetLiquidAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, nil, err
	}

	if liquidAssets.Cmp(minLiquidity) == -1 {
		return liquidAssets, minLiquidity, nil
	}

	return minLiquidity, minLiquidity, nil
}

func findMinCap(values []*big.Int) *big.Int {
	min := new(big.Int).Set(values[0]) // Copy the first value

	for _, value := range values {
		// If value is smaller than min, replace min
		if value.Cmp(min) == -1 {
			min = value
		}
	}

	return min
}

func (q *fevmQueries) InfPoolFeesAccrued(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	return poolCaller.TreasuryFeesOwed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
}

func (q *fevmQueries) InfPoolApy(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	perEpochRate, err := poolCaller.GetRate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	baseApr := new(big.Int).Mul(perEpochRate, big.NewInt(constants.EpochsInYear))
	// take out the treasury fee from the APR
	tFeeRate, err := q.TreasuryFeeRate(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	// apr * (1 - tFeeRate)
	baseApr.Mul(baseApr, new(big.Int).Sub(big.NewInt(1e18), tFeeRate))
	// div out the WAD precision
	baseApr.Div(baseApr, big.NewInt(1e18))

	// use the utilization rate to get the current APY
	borrowed, err := q.InfPoolTotalBorrowed(ctx, nil)
	if err != nil {
		return nil, err
	}
	assets, err := q.InfPoolTotalAssets(ctx, nil)
	if err != nil {
		return nil, err
	}

	borrowedAtto := util.ToAtto(borrowed)
	assetsAtto := util.ToAtto(assets)

	borrowedAtto.Mul(borrowedAtto, big.NewInt(1e18))
	utilizationRate := new(big.Int).Div(borrowedAtto, assetsAtto)

	curApy := new(big.Int).Mul(baseApr, utilizationRate)
	// div out the WAD precision twice to get back to 1e18 = 100% rate
	curApy.Div(curApy, big.NewInt(1e18))
	curApy.Div(curApy, big.NewInt(1e18))

	return curApy, nil
}
