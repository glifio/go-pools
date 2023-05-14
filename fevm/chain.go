package fevm

import (
	"context"
	"math/big"
)

func (q *fevmQueries) ChainHeight(ctx context.Context) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	bn, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	return big.NewInt(int64(bn)), nil
}

func (q *fevmQueries) ChainID() *big.Int {
	return q.chainID
}
