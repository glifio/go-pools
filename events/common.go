package events

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/deploy"
)

func getStartEpochIfNil(startEpoch *big.Int, chainID *big.Int) uint64 {
	// start epoch supplied, use it
	if startEpoch != nil && startEpoch.Cmp(big.NewInt(0)) == 1 {
		return startEpoch.Uint64()
	}

	// start epoch not supplied, use protocol deploy epoch
	switch chainID.Int64() {
	case constants.MainnetChainID:
		return deploy.ProtocolDeployEpoch.Uint64()
	case constants.CalibnetChainID:
		return deploy.TProtocolDeployEpoch.Uint64()
	default:
		return 0
	}
}

func getEndEpochIfNil(endEpoch *big.Int) *uint64 {
	// start epoch supplied, use it
	if endEpoch != nil && endEpoch.Cmp(big.NewInt(0)) == 1 {
		end := endEpoch.Uint64()
		return &end
	}

	return nil
}

func getFilterOpts(ctx context.Context, startEpoch *big.Int, endEpoch *big.Int, chainID *big.Int) *bind.FilterOpts {
	start := getStartEpochIfNil(startEpoch, chainID)
	end := getEndEpochIfNil(endEpoch)

	return &bind.FilterOpts{Context: ctx, Start: start, End: end}
}

func getWatchOpts(ctx context.Context, startEpoch *big.Int, chainID *big.Int) *bind.WatchOpts {
	start := getStartEpochIfNil(startEpoch, chainID)
	return &bind.WatchOpts{Start: &start, Context: ctx}
}
