package token

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/util"
)

func ReadAgentOwnerMap(testDrop bool) (map[common.Address]common.Address, error) {
	network := "mainnet"
	if testDrop {
		network = "calibnet"
	}

	agentOwnerMapPath := fmt.Sprintf("data/%s/agent-to-owner.json", network)
	agentOwnerMap, err := files.ReadFile(agentOwnerMapPath)
	if err != nil {
		return nil, err
	}

	var data map[common.Address]common.Address
	err = json.Unmarshal(agentOwnerMap, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CheckAirdropEligibility(address common.Address, testDrop bool) (eligibleAmount *big.Float, claimer common.Address, err error) {
	mt := &MerkleTree{}
	mt, err = mt.ReadFromJSON(testDrop)
	if err != nil {
		return nil, address, err
	}

	agentOwnerMap, err := ReadAgentOwnerMap(testDrop)
	if err != nil {
		return nil, address, err
	}

	claimer = address
	var agentAddr *common.Address

	// Check if this is an agent address
	ownerAddress, exists := agentOwnerMap[address]
	if exists {
		claimer = ownerAddress
		agentAddr = &address // Pass agent address for duplicate resolution
	}

	// Use the improved merkle tree method that handles duplicates
	amount, err := mt.GetLeafValueForAddrWithAgent(claimer, agentAddr)
	if err != nil {
		return nil, claimer, err
	}

	if amount == nil {
		return big.NewFloat(0), claimer, nil
	}

	return util.ToFIL(amount), claimer, nil
}
