package util

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/glifio/go-pools/abigen"
)

type humanReadableErrorCode struct {
	errCode [4]byte
	errMsg  string
}

var humanReadableErrorCodes = []humanReadableErrorCode{
	{
		errCode: [4]byte{212, 187, 102, 113},
		errMsg:  "Caller type not supported - is your Agent the owner of the miner you're trying to add?",
	},
}

func humanReadableRevert(errMsg error) error {
	errStr := errMsg.Error()

	// Check if "revert reason: " exists in the errMsg
	if !strings.Contains(errStr, "revert reason: ") {
		return fmt.Errorf("error message does not contain a revert reason %s", errStr)
	}

	// Split the errMsg by "revert reason: "
	parts := strings.Split(errStr, "revert reason: ")

	// Check if we have at least two parts
	if len(parts) < 2 {
		return fmt.Errorf("error message is not in the expected format %s", errStr)
	}

	// The error selector should be the next 10 characters after "revert reason: "
	errorSelectorStr := parts[1][:10]

	// Check if the error selector starts with "0x"
	if !strings.HasPrefix(errorSelectorStr, "0x") {
		return fmt.Errorf("error selector does not start with '0x' %s", errStr)
	}

	// Remove the "0x" prefix
	errorSelectorStr = errorSelectorStr[2:]

	// Decode the hex string to a byte slice
	errorSelector, err := hex.DecodeString(errorSelectorStr)
	if err != nil {
		return fmt.Errorf("failed to decode error selector: %v", err)
	}

	matched, err := matchSelector([4]byte(errorSelector))

	if matched {
		return err
	}

	return fmt.Errorf("Error calling contract (could not decode): %s", errStr)
}

func matchSelector(thrownSelector [4]byte) (bool, error) {
	abis, err := gatherABIS()
	// not sure if we should throw the actual error here...
	if err != nil {
		return false, fmt.Errorf("Transaction reverted with error: %s", thrownSelector)
	}

	var humanErr error

	// Loop through all the ABIs
	for _, abi := range abis {
		for name, error := range abi.Errors {
			// Compute the selector for the error
			selector := crypto.Keccak256Hash([]byte(error.Sig))

			if [4]byte(selector[:4]) == thrownSelector {
				humanErr = fmt.Errorf("Transaction reverted with error: %s", name)
			}
		}
	}

	for _, humanReadableErrorCode := range humanReadableErrorCodes {
		if humanReadableErrorCode.errCode == thrownSelector {
			humanErr = fmt.Errorf("Transaction reverted with error: %s", humanReadableErrorCode.errMsg)
		}
	}

	if humanErr == nil {
		return false, fmt.Errorf("Transaction reverted with error: %s", thrownSelector)
	}

	return true, humanErr
}

func gatherABIS() ([]*abi.ABI, error) {
	abis := []*abi.ABI{}

	agentAbi, err := abigen.AgentMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, agentAbi)

	agentFactoryAbi, err := abigen.AgentFactoryMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, agentFactoryAbi)

	agentPoliceAbi, err := abigen.AgentPoliceMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, agentPoliceAbi)

	credParserAbi, err := abigen.CredParserMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, credParserAbi)

	infPoolAbi, err := abigen.InfinityPoolMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, infPoolAbi)

	minerRegistryAbi, err := abigen.MinerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, minerRegistryAbi)

	poolRegistryAbi, err := abigen.PoolRegistryMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, poolRegistryAbi)

	poolTokenAbi, err := abigen.PoolTokenMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, poolTokenAbi)

	rateModuleAbi, err := abigen.RateModuleMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, rateModuleAbi)

	routerAbi, err := abigen.RouterMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, routerAbi)

	wrappedFILAbi, err := abigen.WFILMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, wrappedFILAbi)

	return abis, nil
}
