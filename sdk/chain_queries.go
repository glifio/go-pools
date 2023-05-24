package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	filtypes "github.com/filecoin-project/lotus/chain/types"
)

func (q *fevmQueries) ChainHead(ctx context.Context) (*filtypes.TipSet, error) {
	lClient, closer, err := q.extern.ConnectLotusClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	return lClient.ChainHead(ctx)
}

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

func (q *fevmQueries) ChainGetNonce(ctx context.Context, addr common.Address) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	nonce, err := ethClient.NonceAt(ctx, addr, nil)
	if err != nil {
		return nil, err
	}

	return big.NewInt(int64(nonce)), nil
}
