package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
)

func (q *fevmQueries) LPPlusFutureValidityDuration(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lpPlus, err := abigen.NewLPPlusCaller(q.lpPlus, client)
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	return lpPlus.FutureValidityDuration(opts)
}

func (q *fevmQueries) LPPlusWindowId(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lpPlus, err := abigen.NewLPPlusCaller(q.lpPlus, client)
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	return lpPlus.WindowId(opts)
}

func (q *fevmQueries) LPPlusWindowIdToStakingSnapshot(ctx context.Context, windowId *big.Int, blockNumber *big.Int) (abigen.StakingSnapshot, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return abigen.StakingSnapshot{}, err
	}
	defer client.Close()

	lpPlus, err := abigen.NewLPPlusCaller(q.lpPlus, client)
	if err != nil {
		return abigen.StakingSnapshot{}, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	return lpPlus.WindowIdToStakingSnapshot(opts, windowId)
}

func (q *fevmQueries) LPPlusWindowIdToMerkleRoot(ctx context.Context, windowId *big.Int, blockNumber *big.Int) ([32]byte, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return [32]byte{}, err
	}
	defer client.Close()

	lpPlus, err := abigen.NewLPPlusCaller(q.lpPlus, client)
	if err != nil {
		return [32]byte{}, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	return lpPlus.WindowIdToMerkleRoot(opts, windowId)
}

func (q *fevmQueries) LPPlusTokenOfOwnerByIndex(ctx context.Context, owner common.Address, index *big.Int, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lpPlus, err := abigen.NewLPPlusCaller(q.lpPlus, client)
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	return lpPlus.TokenOfOwnerByIndex(opts, owner, index)
}
