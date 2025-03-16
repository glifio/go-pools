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
	// first we check if this address is in the merkle tree
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
	// check to see if this is an agent address
	ownerAddress, exists := agentOwnerMap[address]
	if exists {
		claimer = ownerAddress
	}

	entries := mt.Entries()

	for _, entry := range entries {
		addr := entry.Value[0].(common.Address)
		if addr.Hex() == claimer.Hex() {
			amount := entry.Value[1].(*big.Int)
			return util.ToFIL(amount), claimer, nil
		}
	}

	return big.NewFloat(0), claimer, nil
}
