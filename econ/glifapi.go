package econ

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
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

		return nil, fmt.Errorf("error fetching margin. Status code: %d", resp.StatusCode)
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

	url := fmt.Sprintf("%s/agent/%s/miners", eventsURL, agentAddr)
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

	var response []MinerDetailsJSON
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, nil, err
	}

	// convert the miners stats to the econ package's type
	for _, miner := range response {
		bal, ok := new(big.Int).SetString(miner.Balance, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner balance to big.Int")
		}

		avail, ok := new(big.Int).SetString(miner.AvailableBalance, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner available balance to big.Int")
		}

		ip, ok := new(big.Int).SetString(miner.InitialPledge, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner initial pledge to big.Int")
		}

		locked, ok := new(big.Int).SetString(miner.LockedRewards, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner locked rewards to big.Int")
		}

		feeDebt, ok := new(big.Int).SetString(miner.FeeDebt, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner fee debt to big.Int")
		}

		termFee, ok := new(big.Int).SetString(miner.TerminationFee, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner termination fee to big.Int")
		}

		live, ok := new(big.Int).SetString(miner.LiveSectors, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner live sectors to big.Int")
		}

		faulty, ok := new(big.Int).SetString(miner.FaultySectors, 10)
		if !ok {
			return nil, nil, fmt.Errorf("error converting miner faulty sectors to big.Int")
		}

		miners = append(miners, miner.MinerAddr)
		baseFis = append(baseFis, NewBaseFi(
			bal,
			avail,
			locked,
			ip,
			feeDebt,
			termFee,
			live,
			faulty,
		))
	}

	return miners, baseFis, nil
}

func GetPoolMetricsFromAPI(eventsURL string) (*PoolMetrics, error) {
	url := fmt.Sprintf("%s/metrics", eventsURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// if the server can't find the agent we're asking for, it will return a 404
		// we treat it as 0 baseFis
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("error fetching collateral stats. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response PoolsMetricsJSON
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	poolTotalAssets := big.NewInt(0)
	poolTotalAssets.SetString(response.PoolTotalAssets, 10)
	poolTotalBorrowed := big.NewInt(0)
	poolTotalBorrowed.SetString(response.PoolTotalBorrowed, 10)
	poolTotalBorrowableAssets := big.NewInt(0)
	poolTotalBorrowableAssets.SetString(response.PoolTotalBorrowableAssets, 10)
	poolExitReserve := big.NewInt(0)
	poolExitReserve.SetString(response.PoolExitReserve, 10)
	totalMinerCollaterals := big.NewInt(0)
	totalMinerCollaterals.SetString(response.TotalMinerCollaterals, 10)
	totalValueLocked := big.NewInt(0)
	totalValueLocked.SetString(response.TotalValueLocked, 10)
	totalMinersSectors := big.NewInt(0)
	totalMinersSectors.SetString(response.TotalMinersSectors, 10)
	totalMinerQAP := big.NewInt(0)
	totalMinerQAP.SetString(response.TotalMinerQAP, 10)
	totalMinerRBP := big.NewInt(0)
	totalMinerRBP.SetString(response.TotalMinerRBP, 10)
	totalMinerEDR := big.NewInt(0)
	totalMinerEDR.SetString(response.TotalMinerEDR, 10)

	result := PoolMetrics{
		Height:                    response.Height,
		Timestamp:                 response.Timestamp,
		PoolTotalAssets:           poolTotalAssets,
		PoolTotalBorrowed:         poolTotalBorrowed,
		PoolTotalBorrowableAssets: poolTotalBorrowableAssets,
		PoolExitReserve:           poolExitReserve,
		TotalAgentCount:           response.TotalAgentCount,
		TotalMinerCollaterals:     totalMinerCollaterals,
		TotalMinersCount:          response.TotalMinersCount,
		TotalValueLocked:          totalValueLocked,
		TotalMinersSectors:        totalMinersSectors,
		TotalMinerQAP:             totalMinerQAP,
		TotalMinerRBP:             totalMinerRBP,
		TotalMinerEDR:             totalMinerEDR,
	}
	return &result, nil
}
