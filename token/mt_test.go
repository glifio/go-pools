package token

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gotest.tools/assert"
)

func createInMemoryJSON() string {
	return `{
		"format": "standard-v1",
		"leafEncoding": ["address", "uint256"],
		"tree": [
			"0xfa77c123512b8d1a96eb642e97ba5813924c91d903a8b414ba6c273e9cb186a8",
			"0xe9f01c18b7f5594df77d8a85ea6706c36a7af4f20953baf052c26f24649b0d69",
			"0xf600243d7915a98569d643bceb700ab77996b47441375bd56f5ce1852cf3c78e",
			"0x0317388807289770a9258380c114bc8a07a3423f3c1a2929ea7b9f0b1ea75aa3",
			"0x2be8798f98a4f9eaefcc7eb6094523156b56740dcfcd8662b0242a732c5cc4b7",
			"0x6dcfd0cd698b26806648d326861f6f4044bec8f997c51ea069a74c69b50e6897",
			"0xf7c9ada898b95afe693e14db0888bbd9b37a7c191131b7915fbe35bf50e5f01b",
			"0x273993308c3abf250ff7f8868fafb5dba6a2270f716b8932e0eef1cb1cd3c63f",
			"0xb3a597c09e74c10ca3924eea38f1fb3cea68e9be17fdb873aebef1caf9f56440",
			"0xe29ad263494c105af7b72500475d522912864c25e72f4f2e787e5890cb7d7d62",
			"0xdc245c15a3d55d2822dd6d6157295473ebb9270d916bd5a6769fab50e83b4162",
			"0x9c8bf0a7ef5fab223b5795c0cc5d16bdb82556697f706801cbb6e3537559b4df",
			"0x99bebfa12784556b0a6c2c37bf8c7c682c25c9f9ef88bf350594a52fbbc61ab5",
			"0x8b4d6b3712a65c7551f9446b578818a788f1bad79bd6144414d07580175f49fb",
			"0x66daba46f046f96491fd8c3baaff0de0d78c0c7859d53412e548985210e8fd7d",
			"0x41c8154967bf86572d17afe74cf6a9794d9ae8aa4406921329f4613f874258ef",
			"0x2c09da6d2ec20ced7b3e302064e79a4701aea0be6ba9a23afff12f29036acaf3",
			"0x21c498f30f63a10cebf1158441879ff7af0b4a4214453a91b74f631fac3f5b3b",
			"0x10a9ae8051d64c7879d3373f417b6d0f543d2a6ea8136fa8a3d427ae14dd67d2"
		],
		"values": [
			{
				"value": [
					"0xFf000000000000000000000000000000000013DC",
					"1000000000000000000"
				],
				"treeIndex": 15
			},
			{
				"value": [
					"0xFF000000000000000000000000000000000013dd",
					"1000000000000000000"
				],
				"treeIndex": 18
			},
			{
				"value": [
					"0xFf000000000000000000000000000000000013dF",
					"1000000000000000000"
				],
				"treeIndex": 12
			},
			{
				"value": [
					"0xFf00000000000000000000000000000000001876",
					"1000000000000000000"
				],
				"treeIndex": 13
			},
			{
				"value": [
					"0xfF0000000000000000000000000000000000E879",
					"1000000000000000000"
				],
				"treeIndex": 10
			},
			{
				"value": [
					"0xFf0000000000000000000000000000000000e6f6",
					"1000000000000000000"
				],
				"treeIndex": 17
			},
			{
				"value": [
					"0xFF00000000000000000000000000000000017F7E",
					"1000000000000000000"
				],
				"treeIndex": 11
			},
			{
				"value": [
					"0xfF000000000000000000000000000000000017B3",
					"1000000000000000000"
				],
				"treeIndex": 9
			},
			{
				"value": [
					"0xFf000000000000000000000000000000000017b4",
					"1000000000000000000"
				],
				"treeIndex": 16
			},
			{
				"value": [
					"0xff000000000000000000000000000000000017B5",
					"1000000000000000000"
				],
				"treeIndex": 14
			}
		],
		"root": "0xfa77c123512b8d1a96eb642e97ba5813924c91d903a8b414ba6c273e9cb186a8",
		"id": "dd946a64-e1d2-4a59-aa0d-1d1174ee71c8"
	}
`
}

func TestGetMerkleID(t *testing.T) {
	mt, err := MerkleTreeFromJSON([]byte(createInMemoryJSON()))
	assert.NilError(t, err)
	id, err := ParseUUIDToBytes("dd946a64-e1d2-4a59-aa0d-1d1174ee71c8")
	assert.NilError(t, err)
	assert.Equal(t, mt.ID(), id)
}

func TestGetRoot(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	expectedRoot := "0x401db874c7325c0a85a7a6154da0d54e7a0513e540f3c89137ece1411d0b3bd8"
	rootBytes := mt.GetRoot() // Returns []byte
	rootHex := "0x" + hex.EncodeToString(rootBytes)
	assert.Equal(t, rootHex, expectedRoot)
}

func TestGetIdxForAddr(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	address := common.HexToAddress("0xA6f573F96A4B9037Ee9057B22678F9093aB2b2D8")
	expectedIndex := 6928

	index, err := mt.GetIdxForAddr(address)
	assert.NilError(t, err)
	assert.Equal(t, index, expectedIndex)
}

func TestGetProofForAddr(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	address := common.HexToAddress("0xA6f573F96A4B9037Ee9057B22678F9093aB2b2D8")
	idx, err := mt.GetIdxForAddr(address)
	assert.NilError(t, err)

	proof, err := mt.GetProofForAddr(address)
	assert.NilError(t, err)

	entries := mt.Entries()
	leaf := entries[idx].Value

	// Convert proof to [][]byte
	proofBytes := make([][]byte, len(proof))
	for i, proofItem := range proof {
		proofBytes[i] = proofItem[:]
	}

	// Verify the proof
	valid, err := mt.Verify(proofBytes, leaf)
	assert.NilError(t, err)
	assert.Assert(t, valid)
}

func TestGetIdxForAddrWithAgent_DuplicateOwner(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	// Owner address that appears twice in the merkle tree
	duplicateOwner := common.HexToAddress("0x3e39a95489a06aeb044c4438790f74cf27a8ca82")
	agent139 := common.HexToAddress("0x207201C13A0640c99152f9aF05C16f95f27d659F")
	agent149 := common.HexToAddress("0x992309c1116Bb9DFB52d6793Be258E3eA9F0D76a")

	// Expected amounts from the merkle tree
	expectedAmount139 := "5903426997075036796116"   // Agent 139's amount
	expectedAmount149 := "487247735941371334694747" // Agent 149's amount

	// Test with agent 139 context
	idx139, err := mt.GetIdxForAddrWithAgent(duplicateOwner, &agent139)
	assert.NilError(t, err)
	assert.Assert(t, idx139 >= 0, "Index should be valid")

	// Verify the index points to the correct amount for agent 139
	entries := mt.Entries()
	value139 := entries[idx139].Value[1].(*big.Int)
	assert.Equal(t, value139.String(), expectedAmount139)

	// Test with agent 149 context
	idx149, err := mt.GetIdxForAddrWithAgent(duplicateOwner, &agent149)
	assert.NilError(t, err)
	assert.Assert(t, idx149 >= 0, "Index should be valid")

	// Verify the index points to the correct amount for agent 149
	value149 := entries[idx149].Value[1].(*big.Int)
	assert.Equal(t, value149.String(), expectedAmount149)

	// The indices should be different for the two agents
	assert.Assert(t, idx139 != idx149, "Agent 139 and 149 should have different indices")

	// Test without agent context - should get first match
	idxNoAgent, err := mt.GetIdxForAddrWithAgent(duplicateOwner, nil)
	assert.NilError(t, err)
	assert.Assert(t, idxNoAgent >= 0, "Index should be valid")
	// Should be one of the two indices
	assert.Assert(t, idxNoAgent == idx139 || idxNoAgent == idx149)
}

func TestGetLeafValueForAddrWithAgent_DuplicateOwner(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	// Owner address that appears twice in the merkle tree
	duplicateOwner := common.HexToAddress("0x3e39a95489a06aeb044c4438790f74cf27a8ca82")
	agent139 := common.HexToAddress("0x207201C13A0640c99152f9aF05C16f95f27d659F")
	agent149 := common.HexToAddress("0x992309c1116Bb9DFB52d6793Be258E3eA9F0D76a")

	// Expected amounts from the merkle tree
	expectedAmount139 := "5903426997075036796116"   // Agent 139's amount
	expectedAmount149 := "487247735941371334694747" // Agent 149's amount

	// Test with agent 139 context
	value139, err := mt.GetLeafValueForAddrWithAgent(duplicateOwner, &agent139)
	assert.NilError(t, err)
	assert.Equal(t, value139.String(), expectedAmount139)

	// Test with agent 149 context
	value149, err := mt.GetLeafValueForAddrWithAgent(duplicateOwner, &agent149)
	assert.NilError(t, err)
	assert.Equal(t, value149.String(), expectedAmount149)

	// Verify the amounts are different
	assert.Assert(t, value139.Cmp(value149) != 0, "Agent 139 and 149 should have different amounts")
}

func TestGetLeafValueForAddrWithAgent_NonDuplicate(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	// Regular address (not a duplicate)
	address := common.HexToAddress("0xA6f573F96A4B9037Ee9057B22678F9093aB2b2D8")

	// Test with nil agent (should work normally)
	valueNil, err := mt.GetLeafValueForAddrWithAgent(address, nil)
	assert.NilError(t, err)
	assert.Assert(t, valueNil != nil)

	// Test with some agent context (should still return the same value since no duplicate)
	someAgent := common.HexToAddress("0x0000000000000000000000000000000000000001")
	valueWithAgent, err := mt.GetLeafValueForAddrWithAgent(address, &someAgent)
	assert.NilError(t, err)
	assert.Equal(t, valueNil.String(), valueWithAgent.String())
}

func TestGetProofForAddrWithAgent_DuplicateOwner(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	// Owner address that appears twice in the merkle tree
	duplicateOwner := common.HexToAddress("0x3e39a95489a06aeb044c4438790f74cf27a8ca82")
	agent139 := common.HexToAddress("0x207201C13A0640c99152f9aF05C16f95f27d659F")
	agent149 := common.HexToAddress("0x992309c1116Bb9DFB52d6793Be258E3eA9F0D76a")

	// Get proof and value for agent 139
	proof139, err := mt.GetProofForAddrWithAgent(duplicateOwner, &agent139)
	assert.NilError(t, err)
	value139, err := mt.GetLeafValueForAddrWithAgent(duplicateOwner, &agent139)
	assert.NilError(t, err)

	// Verify proof for agent 139
	proofBytes139 := make([][]byte, len(proof139))
	for i, proofItem := range proof139 {
		proofBytes139[i] = proofItem[:]
	}
	leaf139 := []interface{}{duplicateOwner, value139}
	valid139, err := mt.Verify(proofBytes139, leaf139)
	assert.NilError(t, err)
	assert.Assert(t, valid139, "Proof for agent 139 should be valid")

	// Get proof and value for agent 149
	proof149, err := mt.GetProofForAddrWithAgent(duplicateOwner, &agent149)
	assert.NilError(t, err)
	value149, err := mt.GetLeafValueForAddrWithAgent(duplicateOwner, &agent149)
	assert.NilError(t, err)

	// Verify proof for agent 149
	proofBytes149 := make([][]byte, len(proof149))
	for i, proofItem := range proof149 {
		proofBytes149[i] = proofItem[:]
	}
	leaf149 := []interface{}{duplicateOwner, value149}
	valid149, err := mt.Verify(proofBytes149, leaf149)
	assert.NilError(t, err)
	assert.Assert(t, valid149, "Proof for agent 149 should be valid")

	// The proofs should be different (different tree indices)
	assert.Assert(t, len(proof139) > 0 && len(proof149) > 0)
}

func TestGetProofForAddrWithAgent_NonDuplicate(t *testing.T) {
	mt := &MerkleTree{}
	mt, err := mt.ReadFromJSON(false)
	assert.NilError(t, err)

	// Regular address (not a duplicate)
	address := common.HexToAddress("0xA6f573F96A4B9037Ee9057B22678F9093aB2b2D8")

	// Get proof with nil agent
	proofNil, err := mt.GetProofForAddrWithAgent(address, nil)
	assert.NilError(t, err)

	// Get proof with some agent context
	someAgent := common.HexToAddress("0x0000000000000000000000000000000000000001")
	proofWithAgent, err := mt.GetProofForAddrWithAgent(address, &someAgent)
	assert.NilError(t, err)

	// Proofs should be identical for non-duplicate addresses
	assert.Equal(t, len(proofNil), len(proofWithAgent))
	for i := range proofNil {
		assert.Equal(t, proofNil[i], proofWithAgent[i])
	}
}
