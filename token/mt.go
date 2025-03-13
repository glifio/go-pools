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

func (mt *MerkleTree) GetIdxForAddr(address common.Address) (int, error) {
	// loop through the leafs and get the associated leaf of the address
	entries := mt.Entries()
	for _, entry := range entries {
		addr := entry.Value[0].(common.Address)
		if addr.Hex() == address.Hex() {
			return entry.ValueIndex, nil
		}
	}

	return -1, nil
}

func (mt *MerkleTree) GetLeafValueForAddr(address common.Address) (*big.Int, error) {
	entries := mt.Entries()
	for _, entry := range entries {
		addr := entry.Value[0].(common.Address)
		if addr.Hex() == address.Hex() {
			return entry.Value[1].(*big.Int), nil
		}
	}

	return nil, nil
}

func (mt *MerkleTree) GetProofForAddr(address common.Address) ([][32]byte, error) {
	idx, err := mt.GetIdxForAddr(address)
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
