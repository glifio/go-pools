package token

import (
	"embed"
	"encoding/json"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/common"
)

//go:embed data/*
var files embed.FS

type MerkleTree struct {
	id string
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

func (mt *MerkleTree) ID() string {
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

	mt := &MerkleTree{
		id:           data.Id,
		StandardTree: *tree,
	}

	return mt, nil
}

func (mt *MerkleTree) ReadFromJSON() (*MerkleTree, error) {
	merkleTree, err := files.ReadFile("data/merkle-tree.json")
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

func (mt *MerkleTree) GetProofForAddr(address common.Address) ([][]byte, error) {
	idx, err := mt.GetIdxForAddr(address)
	if err != nil {
		return nil, err
	}

	proof, err := mt.GetProofWithIndex(idx)
	if err != nil {
		return nil, err
	}

	return proof, nil
}
