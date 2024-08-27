package econ

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/glifio/go-pools/constants"
)

func GetAgentFiFromAPI(agentAddr common.Address) (*AgentFi, error) {
	url := fmt.Sprintf("%s/agent/%s/collateral-value", constants.EventsURL, agentAddr)
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
			return EmptyAgentFi(), nil
		}

		return nil, fmt.Errorf("error fetching collateral stats. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response agentCollateralStats
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	minersStats := make([]*BaseFi, len(response.MinersTerminationStats))

	agentAvail, ok := new(big.Int).SetString(response.AvailableBalance, 10)
	if !ok {
		return nil, fmt.Errorf("error converting agent available balance to big.Int")
	}

	// convert the miners stats to the econ package's type
	for i, miner := range response.MinersTerminationStats {
		m, err := convertToBigInt(miner)
		if err != nil {
			return nil, err
		}
		minersStats[i] = NewBaseFi(
			m.Total,
			m.Available,
			m.Vesting,
			m.Pledged,
			// TODO: https://github.com/glifio/go-db/issues/3
			big.NewInt(0),
			m.TerminationPenalty,
			// TODO: merge liveSectors stuff in pools-events
			big.NewInt(0),
			big.NewInt(0),
		)
	}

	interest, principal, err := GetAgentDebtFromAPI(agentAddr)
	if err != nil {
		return nil, err
	}

	afi := NewAgentFi(
		agentAvail,
		Liability{
			Interest:  interest,
			Principal: principal,
		},
		minersStats,
	)
	return afi, nil
}

type AgentInfo struct {
	TxHash        string         `json:"txHash"`
	Height        uint64         `json:"height"`
	Id            uint64         `json:"id"`
	Address       string         `json:"address"`
	AddressNative common.Address `json:"addressNative"`
	// BALANCE DEPRECATED
	Balance          string `json:"balance"`
	Miners           uint64 `json:"miners"`
	AvailableBalance string `json:"availableBalance"`
	PrincipalBalance string `json:"principalBalance"`
}

func GetBaseFisFromAPI(agentAddr common.Address) (miners []address.Address, baseFis []*BaseFi, err error) {
	baseFis = make([]*BaseFi, 0)
	miners = make([]address.Address, 0)

	url := fmt.Sprintf("%s/agent/%s/collateral-value", constants.EventsURL, agentAddr)
	// Making an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// if the server can't find the agent we're asking for, it will return a 404
		// we treat it as 0 baseFis
		if resp.StatusCode == http.StatusNotFound {
			return miners, baseFis, nil
		}

		return nil, nil, fmt.Errorf("error fetching collateral stats. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var response agentCollateralStats
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, nil, err
	}

	// convert the miners stats to the econ package's type
	for _, miner := range response.MinersTerminationStats {
		m, err := convertToBigInt(miner)
		if err != nil {
			return nil, nil, err
		}
		miners = append(miners, m.Address)
		baseFis = append(baseFis, NewBaseFi(
			m.Total,
			m.Available,
			m.Vesting,
			m.Pledged,
			// TODO: https://github.com/glifio/go-db/issues/3
			big.NewInt(0),
			m.TerminationPenalty,
			// TODO: merge liveSectors stuff in pools-events
			big.NewInt(0),
			big.NewInt(0),
		))
	}

	return miners, baseFis, nil
}

func GetAgentDebtFromAPI(agentAddr common.Address) (interest *big.Int, principal *big.Int, err error) {
	url := fmt.Sprintf("%s/agent/%s", constants.EventsURL, agentAddr)
	// Making an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// if the server can't find the agent we're asking for, it will return a 404
		// we treat it as 0 principal
		if resp.StatusCode == http.StatusNotFound {
			return big.NewInt(0), big.NewInt(0), nil
		}

		return big.NewInt(0), big.NewInt(0), fmt.Errorf("error fetching principal. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	var response AgentInfo
	if err := json.Unmarshal(body, &response); err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	principal, ok := new(big.Int).SetString(response.PrincipalBalance, 10)
	if !ok {
		return big.NewInt(0), big.NewInt(0), fmt.Errorf("error converting principal to big.Int")
	}

	//TODO: https://github.com/glifio/pools-events/issues/146 - interest is not returned by the API
	return big.NewInt(0), principal, nil
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

type minerCollateralStatBig struct {
	Address            address.Address `json:"address"`
	Total              *big.Int        `json:"total"`
	Available          *big.Int        `json:"available"`
	Pledged            *big.Int        `json:"pledged"`
	Vesting            *big.Int        `json:"vesting"`
	TerminationPenalty *big.Int        `json:"terminationPenalty"`
}

func convertToBigInt(m *minerCollateralStat) (*minerCollateralStatBig, error) {
	// convert all string types to big.Int
	total, ok := new(big.Int).SetString(m.Total, 10)
	if !ok {
		return nil, fmt.Errorf("error converting miner total collateral to big.Int")
	}
	available, ok := new(big.Int).SetString(m.Available, 10)
	if !ok {
		return nil, fmt.Errorf("error converting miner available collateral to big.Int")
	}
	pledged, ok := new(big.Int).SetString(m.Pledged, 10)
	if !ok {
		return nil, fmt.Errorf("error converting miner pledged collateral to big.Int")
	}
	vesting, ok := new(big.Int).SetString(m.Vesting, 10)
	if !ok {
		return nil, fmt.Errorf("error converting miner vesting collateral to big.Int")
	}
	tp, ok := new(big.Int).SetString(m.TerminationPenalty, 10)
	if !ok {
		return nil, fmt.Errorf("error converting miner termination penalty to big.Int")
	}

	return &minerCollateralStatBig{
		Address:            m.Address,
		Total:              total,
		Available:          available,
		Pledged:            pledged,
		Vesting:            vesting,
		TerminationPenalty: tp,
	}, nil
}
