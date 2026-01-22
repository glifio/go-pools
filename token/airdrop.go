package token

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/util"
)

var (
	// Duplicate owner case: one owner with two agents
	duplicateOwner = common.HexToAddress("0x3e39a95489a06aeb044c4438790f74cf27a8ca82")
	agent139       = common.HexToAddress("0x207201C13A0640c99152f9aF05C16f95f27d659F")
	agent149       = common.HexToAddress("0x992309c1116Bb9DFB52d6793Be258E3eA9F0D76a")
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

		// Special case: handle duplicate owner with multiple agents
		// Check if this agent is one of the known duplicates
		// this is a hacky workaround to handle the single duplicate owner case, otherwise we could do it programmatically, but it's computationally heavy
		if address.Hex() == agent139.Hex() && claimer.Hex() == duplicateOwner.Hex() {
			// Agent 139 gets 5,903.43 FIL
			amount := new(big.Int)
			amount.SetString("5903426997075036796116", 10)
			return util.ToFIL(amount), claimer, nil
		} else if address.Hex() == agent149.Hex() && claimer.Hex() == duplicateOwner.Hex() {
			// Agent 149 gets 487,247.74 FIL
			amount := new(big.Int)
			amount.SetString("487247735941371334694747", 10)
			return util.ToFIL(amount), claimer, nil
		}
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
