package sdk

import (
	"context"

	filtypes "github.com/filecoin-project/lotus/chain/types"
)

func (q *fevmQueries) PreviewTerminateSectors(ctx context.Context) (*filtypes.TipSet, error) {
	lClient, closer, err := q.extern.ConnectLotusClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	return lClient.ChainHead(ctx)
}
