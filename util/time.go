package util

import (
	"math/big"
	"time"

	"github.com/filecoin-project/lotus/build"
)

var (
	// TODO change to mainnet
	GENESIS_TS = uint64(1598306430)
)

func EpochHeightToTimestamp(height *big.Int) time.Time {
	ts := (build.BlockDelaySecs * (uint64(height.Uint64()))) + GENESIS_TS
	return time.Unix(int64(ts), 0)
}
