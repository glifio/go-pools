package terminate

import (
	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
)

// PreviewTerminateSectorsReturn contains the aggregated results from a query.
type PreviewTerminateSectorsReturn struct {
	Actor                      *types.ActorV5
	MinerInfo                  lotusapi.MinerInfo
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
	MinerInfo              lotusapi.MinerInfo
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

type TerminateSectorsSummary struct {
	MinerAddr          string         `json:"minerAddr"`
	MinerBal           string         `json:"minerBal"`
	TerminationPenalty string         `json:"terminationPenalty"`
	SectorsTerminated  uint64         `json:"sectorsTerminated"`
	SectorsCount       uint64         `json:"sectorCount"`
	MinExpiration      abi.ChainEpoch `json:"minExpiration"`
	MaxExpiration      abi.ChainEpoch `json:"maxExpiration"`
	MaxAge             abi.ChainEpoch `json:"maxAge"`
	MinAge             abi.ChainEpoch `json:"minAge"`
}
