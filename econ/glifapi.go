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
)

func GetAgentFiFromAPI(agentAddr common.Address, eventsURL string) (*AgentFi, error) {
	url := fmt.Sprintf("%s/agent/%s/margin", eventsURL, agentAddr)
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

	var response AgentMarginJSON
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	balance, ok := new(big.Int).SetString(response.Balance, 10)
	if !ok {
		return nil, fmt.Errorf("error converting balance to big.Int")
	}

	spendableBalance, ok := new(big.Int).SetString(response.SpendableBalance, 10)
	if !ok {
		return nil, fmt.Errorf("error converting spendable balance to big.Int")
	}

	agentAvail, ok := new(big.Int).SetString(response.AvailableBalance, 10)
	if !ok {
		return nil, fmt.Errorf("error converting available balance to big.Int")
	}

	lockedRewards, ok := new(big.Int).SetString(response.LockedRewards, 10)
	if !ok {
		return nil, fmt.Errorf("error converting locked rewards to big.Int")
	}

	initialPledge, ok := new(big.Int).SetString(response.InitialPledge, 10)
	if !ok {
		return nil, fmt.Errorf("error converting initial pledge to big.Int")
	}

	feeDebt, ok := new(big.Int).SetString(response.FeeDebt, 10)
	if !ok {
		return nil, fmt.Errorf("error converting fee debt to big.Int")
	}

	terminationFee, ok := new(big.Int).SetString(response.TerminationFee, 10)
	if !ok {
		return nil, fmt.Errorf("error converting termination fee to big.Int")
	}

	liveSectors, ok := new(big.Int).SetString(response.LiveSectors, 10)
	if !ok {
		return nil, fmt.Errorf("error converting live sectors to big.Int")
	}

	faultySectors, ok := new(big.Int).SetString(response.FaultySectors, 10)
	if !ok {
		return nil, fmt.Errorf("error converting faulty sectors to big.Int")
	}

	principal, ok := new(big.Int).SetString(response.Principal, 10)
	if !ok {
		return nil, fmt.Errorf("error converting principal to big.Int")
	}

	interest, ok := new(big.Int).SetString(response.Interest, 10)
	if !ok {
		return nil, fmt.Errorf("error converting interest to big.Int")
	}

	afi := &AgentFi{
		BaseFi: BaseFi{
			Balance:          balance,
			AvailableBalance: agentAvail,
			LockedRewards:    lockedRewards,
			InitialPledge:    initialPledge,
			FeeDebt:          feeDebt,
			TerminationFee:   terminationFee,
			LiveSectors:      liveSectors,
			FaultySectors:    faultySectors,
		},
		Liability: Liability{
			Principal: principal,
			Interest:  interest,
		},
		SpendableBalance: spendableBalance,
	}
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

func GetBaseFisFromAPI(agentAddr common.Address, eventsURL string) (miners []address.Address, baseFis []*BaseFi, err error) {
	baseFis = make([]*BaseFi, 0)
	miners = make([]address.Address, 0)

	url := fmt.Sprintf("%s/agent/%s/collateral-value", eventsURL, agentAddr)
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
