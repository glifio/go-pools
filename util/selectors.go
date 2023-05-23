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

	var methodID [4]byte
	copy(methodID[:], abi.Methods[string(methodStr)].ID[:4])

	if methodID == [4]byte{} {
		return [4]byte{}, errors.New("method not found")
	}

	return methodID, nil
}
