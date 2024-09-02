package terminate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/exp/constraints"
	"golang.org/x/xerrors"
)

// From Lotus

func parseTipSetRef(ctx context.Context, api *lotusapi.FullNodeStruct, tss string) (*types.TipSet, error) {
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

func parseTipSetAndHeight(ctx context.Context, api *lotusapi.FullNodeStruct, tipset string, vmHeight uint64) (h abi.ChainEpoch, ts *types.TipSet, err error) {
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

func FetchAgentCollateralStats(ctx context.Context, agentID *big.Int, eventsURL string) (*AgentCollateralStats, error) {
	url := fmt.Sprintf("%s/agent/%s/collateral-value", eventsURL, agentID)
	// Making an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// if the server can't find the agent we're asking for, it will return a 404
		// we treat it as 0 collateral
		if resp.StatusCode == http.StatusNotFound {
			return &AgentCollateralStats{
				AvailableBalance:       big.NewInt(0),
				TerminationPenalty:     big.NewInt(0),
				MinersTerminationStats: []*MinerCollateralStat{},
				Epoch:                  0,
			}, nil
		}

		return nil, fmt.Errorf("error fetching collateral stats. Status code: %d", resp.StatusCode)
	}

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type minerCollateralStat struct {
		Address            address.Address `json:"address"`
		Total              string          `json:"total"`
		Available          string          `json:"available"`
		Pledged            string          `json:"pledged"`
		Vesting            string          `json:"vesting"`
		TerminationPenalty string          `json:"terminationPenalty"`
	}

	type agentCollateralStats struct {
		AvailableBalance       string                 `json:"agentAvailableBalance"`
		TerminationPenalty     string                 `json:"terminationPenalty"`
		MinersTerminationStats []*minerCollateralStat `json:"minersCollateralStats"`
		Epoch                  abi.ChainEpoch
	}

	var response agentCollateralStats
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	minersStats := make([]*MinerCollateralStat, len(response.MinersTerminationStats))
	for i, miner := range response.MinersTerminationStats {
		// Convert string values to big.Int
		total, ok := new(big.Int).SetString(miner.Total, 10)
		if !ok {
			return nil, fmt.Errorf("Error converting total to big.Int")
		}

		available, ok := new(big.Int).SetString(miner.Available, 10)
		if !ok {
			return nil, fmt.Errorf("Error converting available to big.Int")
		}

		pledged, ok := new(big.Int).SetString(miner.Pledged, 10)
		if !ok {
			return nil, fmt.Errorf("Error converting pledged to big.Int")
		}

		vesting, ok := new(big.Int).SetString(miner.Vesting, 10)
		if !ok {
			return nil, fmt.Errorf("Error converting vesting to big.Int")
		}

		terminationPenalty, ok := new(big.Int).SetString(miner.TerminationPenalty, 10)
		if !ok {
			return nil, fmt.Errorf("Error converting termination penalty to big.Int")
		}

		minersStats[i] = &MinerCollateralStat{
			Address:            miner.Address,
			Total:              total,
			Available:          available,
			Pledged:            pledged,
			Vesting:            vesting,
			TerminationPenalty: terminationPenalty,
		}
	}

	// Convert string values to big.Int
	availableBalance, ok := new(big.Int).SetString(response.AvailableBalance, 10)
	if !ok {
		return nil, fmt.Errorf("Error converting available balance to big.Int")
	}

	terminationPenalty, ok := new(big.Int).SetString(response.TerminationPenalty, 10)
	if !ok {
		return nil, fmt.Errorf("Error converting termination penalty to big.Int")
	}

	return &AgentCollateralStats{
		AvailableBalance:       availableBalance,
		TerminationPenalty:     terminationPenalty,
		MinersTerminationStats: minersStats,
		Epoch:                  response.Epoch,
	}, nil
}
