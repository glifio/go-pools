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

func (q *fevmQueries) IFILBalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolTokenCaller, err := abigen.NewPoolTokenCaller(q.iFIL, client)
	if err != nil {
		return nil, err
	}

	bal, err := poolTokenCaller.BalanceOf(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}
	return util.ToFIL(bal), nil
}

func (q *fevmQueries) IFILSupply(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iFILCaller, err := abigen.NewPoolTokenCaller(q.iFIL, client)
	if err != nil {
		return nil, err
	}

	return iFILCaller.TotalSupply(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
}

func (q *fevmQueries) IFILPrice(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	infPoolCaller, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	return infPoolCaller.ConvertToAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, constants.WAD)
}

func (q *fevmQueries) IFILMinter(ctx context.Context) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	poolTokenCaller, err := abigen.NewPoolTokenCaller(q.iFIL, client)
	if err != nil {
		return common.Address{}, err
	}

	return poolTokenCaller.Minter(nil)
}
