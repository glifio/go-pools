package terminate

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/exp/constraints"
	"golang.org/x/xerrors"
)

// From Lotus

func parseTipSetRef(ctx context.Context, api lotusapi.FullNodeStruct, tss string) (*types.TipSet, error) {
	if tss[0] == '@' {
		if tss == "@head" {
			return api.ChainHead(ctx)
		}

		var h uint64
		if _, err := fmt.Sscanf(tss, "@%d", &h); err != nil {
			return nil, xerrors.Errorf("parsing height tipset ref: %w", err)
		}

		ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		return api.ChainGetTipSetByHeight(ctxTimeout, abi.ChainEpoch(h), types.EmptyTSK)
	}

	cids, err := parseTipSetString(tss)
	if err != nil {
		return nil, err
	}

	if len(cids) == 0 {
		return nil, nil
	}

	k := types.NewTipSetKey(cids...)
	ts, err := api.ChainGetTipSet(ctx, k)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func parseTipSetString(ts string) ([]cid.Cid, error) {
	strs := strings.Split(ts, ",")

	var cids []cid.Cid
	for _, s := range strs {
		c, err := cid.Parse(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		cids = append(cids, c)
	}

	return cids, nil
}

func parseTipSetAndHeight(ctx context.Context, api lotusapi.FullNodeStruct, tipset string, vmHeight uint64) (h abi.ChainEpoch, ts *types.TipSet, err error) {
	h = abi.ChainEpoch(vmHeight)
	if tss := tipset; tss != "" {
		ts, err = parseTipSetRef(ctx, api, tss)
	} else if h > 0 {
		ts, err = api.ChainGetTipSetByHeight(ctx, h, types.EmptyTSK)
	} else {
		ts, err = api.ChainHead(ctx)
	}
	if err != nil {
		return 0, nil, err
	}

	if h == 0 {
		h = ts.Height()
	}

	return h, ts, nil
}

// https://stackoverflow.com/questions/27516387/what-is-the-correct-way-to-find-the-min-between-two-integers-in-go
func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
