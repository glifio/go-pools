package util

import (
	"errors"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
)

func MethodStrToBytes(methodStr constants.Method) ([4]byte, error) {
	abi, err := abigen.AgentMetaData.GetAbi()
	if err != nil {
		return [4]byte{}, err
	}

	methodID := ToByte4Selector(abi.Methods[string(methodStr)].ID)

	if methodID == [4]byte{} {
		return [4]byte{}, errors.New("method not found")
	}

	return methodID, nil
}

func ToByte4Selector(bytes []byte) [4]byte {
	var byte4 [4]byte
	copy(byte4[:], bytes)
	return byte4
}
