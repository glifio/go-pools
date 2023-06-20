package util

import (
	"math/big"
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/glifio/go-pools/constants"
)

var (
	GENESIS_TS   = uint64(1598306430)
	T_GENESIS_TS = uint64(1667326410)
)

func EpochHeightToTimestamp(height *big.Int, chainID *big.Int) time.Time {
	var genTs uint64
	switch chainID.Uint64() {
	case constants.MainnetChainID:
		genTs = GENESIS_TS
	case constants.CalibnetChainID:
		genTs = T_GENESIS_TS
	default:
		genTs = uint64(0)
	}

	ts := (build.BlockDelaySecs * (uint64(height.Uint64()))) + genTs
	return time.Unix(int64(ts), 0)
}
