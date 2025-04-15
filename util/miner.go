package util

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	cbor "github.com/ipfs/go-ipld-cbor"
)

func LoadMinerActor(ctx context.Context, api *api.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*types.Actor, miner.State, error) {
	mact, err := api.StateGetActor(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, nil, err
	}

	tbsMiner := NewTieredBlockstore(api)

	state, err := miner.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsMiner)), mact)

	return mact, state, err
}

func NewTieredBlockstore(lapi *api.FullNodeStruct) blockstore.Blockstore {
	return blockstore.NewTieredBstore(
		blockstore.NewAPIBlockstore(lapi),
		blockstore.NewMemory(),
	)
}
