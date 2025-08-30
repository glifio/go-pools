package util

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
)

func MethodStrToBytes(methodStr constants.Method) ([4]byte, error) {
	abiAgent, _ := abigen.AgentMetaData.GetAbi()
	abiPlus, _ := abigen.PlusMetaData.GetAbi()
	abis := []*abi.ABI{
		abiAgent,
		abiPlus,
	}
	for _, abi := range abis {
		methodID := ToByte4Selector(abi.Methods[string(methodStr)].ID)

		if methodID != [4]byte{} {
			return methodID, nil
		}
	}

	return [4]byte{}, errors.New("method not found")
}

func ToByte4Selector(bytes []byte) [4]byte {
	var byte4 [4]byte
	copy(byte4[:], bytes)
	return byte4
}
