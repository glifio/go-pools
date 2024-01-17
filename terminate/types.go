package terminate

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	minertypes "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/lotus/chain/types"
)

// PreviewTerminateSectorsReturn contains the aggregated results from a query.
type PreviewTerminateSectorsReturn struct {
	Actor                      *types.ActorV5
	MinerInfo                  minertypes.MinerInfo
	VestingBalance             *big.Int
	InitialPledge              *big.Int
	SectorStats                *SectorStats
	SectorsTerminated          uint64
	SectorsCount               uint64
	SectorsInSkippedPartitions uint64
	Epoch                      abi.ChainEpoch
	PartitionsCount            uint64
	SampledPartitionsCount     uint64
	Tipset                     *types.TipSet
}

// PreviewTerminateSectorsProgress is returned multiple times during a streaming
// query to report progress.
type PreviewTerminateSectorsProgress struct {
	Epoch                  abi.ChainEpoch
	MinerInfo              minertypes.MinerInfo
	WorkerActor            *types.ActorV5
	PrevHeightForImmutable abi.ChainEpoch
	WorkerActorPrev        *types.ActorV5
	BatchSize              uint64
	GasLimit               uint64
	DeadlinePartitionCount int
	DeadlinePartitionIndex int
	Deadline               uint64
	DeadlineImmutable      bool
	Partition              int
	SectorsCount           uint64
	SliceStart             uint64
	SliceEnd               uint64
	SliceCount             uint64
}

type PreviewAgentTerminationSummary struct {
	TerminationPenalty *big.Int
	InitialPledge      *big.Int
	VestingBalance     *big.Int
	AvailableBalance   *big.Int
	AgentLiquidBalance *big.Int
}
