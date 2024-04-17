package sdk

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/jimpick/go-ethereum/accounts/abi/bind"
	"github.com/jimpick/go-ethereum/common"
)

func (q *fevmQueries) ListPools(ctx context.Context) ([]common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolRegCaller, err := abigen.NewPoolRegistryCaller(q.poolRegistry, client)
	if err != nil {
		return nil, err
	}

	numPools, err := poolRegCaller.AllPoolsLength(nil)
	if err != nil {
		return nil, err
	}

	var pools []common.Address

	for i := big.NewInt(0); i.Cmp(numPools) < 0; i.Add(i, big.NewInt(1)) {
		poolAddr, err := poolRegCaller.AllPools(nil, i)
		if err != nil {
			return nil, err
		}
		pools = append(pools, poolAddr)
	}

	return pools, nil
}

func (q *fevmQueries) TreasuryFeeRate(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolRegCaller, err := abigen.NewPoolRegistryCaller(q.poolRegistry, client)
	if err != nil {
		return nil, err
	}

	return poolRegCaller.TreasuryFeeRate(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
}
