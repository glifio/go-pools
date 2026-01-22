package token

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/big"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

//go:embed data/*
var files embed.FS

var (
	// Duplicate owner case: one owner with two agents
	// Map agent addresses to their claim amounts for resolving duplicates
	agentClaimAmounts = map[common.Address]*big.Int{
		common.HexToAddress("0x207201C13A0640c99152f9aF05C16f95f27d659F"): mustParseBigInt("5903426997075036796116"),   // Agent 139
		common.HexToAddress("0x992309c1116Bb9DFB52d6793Be258E3eA9F0D76a"): mustParseBigInt("487247735941371334694747"), // Agent 149
	}
)

func mustParseBigInt(s string) *big.Int {
	i := new(big.Int)
	i.SetString(s, 10)
	return i
}

type MerkleTree struct {
	id [16]byte
	smt.StandardTree
}

type MerkleTreeValue struct {
	Value     []string `json:"value"`
	TreeIndex int      `json:"treeIndex"`
}

type OZMerkleJSON struct {
	Format       string            `json:"format"`
	LeafEncoding []string          `json:"leafEncoding"`
	Tree         []string          `json:"tree"`
	Values       []MerkleTreeValue `json:"values"`
	Root         string            `json:"root"`
	Id           string            `json:"id"`
}

func (mt *MerkleTree) ID() [16]byte {
	return mt.id
}

func MerkleTreeFromJSON(jsonBytes []byte) (*MerkleTree, error) {
	var data OZMerkleJSON
	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}

	values := [][]interface{}{}
	for _, v := range data.Values {
		values = append(values, []interface{}{smt.SolAddress(v.Value[0]), smt.SolNumber(v.Value[1])})
	}

	leafEncodings := []string{
		smt.SOL_ADDRESS,
		smt.SOL_UINT256,
	}

	tree, err := smt.Of(values, leafEncodings)
	if err != nil {
		return nil, err
	}

	idBytes, err := ParseUUIDToBytes(data.Id)
	if err != nil {
		return nil, err
	}

	mt := &MerkleTree{
		id:           idBytes,
		StandardTree: *tree,
	}

	return mt, nil
}

func (mt *MerkleTree) ReadFromJSON(testDrop bool) (*MerkleTree, error) {
	network := "mainnet"
	if testDrop {
		network = "calibnet"
	}

	merkleTree, err := files.ReadFile(fmt.Sprintf("data/%s/merkle-tree.json", network))
	if err != nil {
		return nil, err
	}

	return MerkleTreeFromJSON(merkleTree)
}

// getIdxAndValueForAddr finds the correct merkle tree entry for an address.
// When there are duplicate owner addresses in the tree and agentAddr is provided,
// it uses the agent address to determine which specific entry to return.
// Returns (index, value, found) where index is -1 if not found.
func (mt *MerkleTree) getIdxAndValueForAddr(ownerAddr common.Address, agentAddr *common.Address) (int, *big.Int, bool) {
	entries := mt.Entries()

	type match struct {
		idx   int
		value *big.Int
	}
	var matches []match

	// Find all matching entries
	for _, entry := range entries {
		addr := entry.Value[0].(common.Address)
		if addr.Hex() == ownerAddr.Hex() {
			matches = append(matches, match{
				idx:   entry.ValueIndex,
				value: entry.Value[1].(*big.Int),
			})
		}
	}

	if len(matches) == 0 {
		return -1, nil, false
	}

	if len(matches) == 1 {
		return matches[0].idx, matches[0].value, true
	}

	// Multiple matches - use agent address to determine correct entry
	if agentAddr != nil {
		expectedAmount, exists := agentClaimAmounts[*agentAddr]
		if exists {
			for _, m := range matches {
				if m.value.Cmp(expectedAmount) == 0 {
					return m.idx, m.value, true
				}
			}
		}
	}

	// Fallback: return first match (maintains backward compatibility)
	return matches[0].idx, matches[0].value, true
}

func (mt *MerkleTree) GetIdxForAddr(address common.Address) (int, error) {
	return mt.GetIdxForAddrWithAgent(address, nil)
}

func (mt *MerkleTree) GetIdxForAddrWithAgent(ownerAddr common.Address, agentAddr *common.Address) (int, error) {
	idx, _, found := mt.getIdxAndValueForAddr(ownerAddr, agentAddr)
	if !found {
		return -1, nil
	}
	return idx, nil
}

func (mt *MerkleTree) GetLeafValueForAddr(address common.Address) (*big.Int, error) {
	return mt.GetLeafValueForAddrWithAgent(address, nil)
}

func (mt *MerkleTree) GetLeafValueForAddrWithAgent(ownerAddr common.Address, agentAddr *common.Address) (*big.Int, error) {
	_, value, found := mt.getIdxAndValueForAddr(ownerAddr, agentAddr)
	if !found {
		return nil, nil
	}
	return value, nil
}

func (mt *MerkleTree) GetProofForAddr(address common.Address) ([][32]byte, error) {
	return mt.GetProofForAddrWithAgent(address, nil)
}

func (mt *MerkleTree) GetProofForAddrWithAgent(ownerAddr common.Address, agentAddr *common.Address) ([][32]byte, error) {
	idx, err := mt.GetIdxForAddrWithAgent(ownerAddr, agentAddr)
	if err != nil {
		return nil, err
	}

	if idx == -1 {
		return nil, fmt.Errorf("address not found in merkle tree")
	}

	proof, err := mt.GetProofWithIndex(idx)
	if err != nil {
		return nil, err
	}

	var proofArray [][32]byte
	for _, p := range proof {
		var pArray [32]byte
		copy(pArray[:], p)
		proofArray = append(proofArray, pArray)
	}

	return proofArray, nil
}

func ParseUUIDToBytes(uuidStr string) ([16]byte, error) {
	var byteArray [16]byte
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return byteArray, err
	}
	copy(byteArray[:], parsedUUID[:])
	return byteArray, nil
}
