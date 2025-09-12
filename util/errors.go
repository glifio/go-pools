package util

import (
	"encoding/hex"
	"fmt"
	"regexp"
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
	// override the PayUp error message
	{
		errCode: [4]byte{50, 19, 202, 175},
		errMsg:  "PayUp - Agent must pay off its fees owed to the pool before continuing. Run `glif agent pay to-current` first and try again",
	},
}

func HumanReadableRevert(errMsg error) error {
	errStr := errMsg.Error()

	// Check if "revert reason: " exists in the errMsg
	if !strings.Contains(errStr, "revert reason: ") {
		r, _ := regexp.Compile(`vm error=\[0x([0-9A-Fa-f]+)\]`)
		matches := r.FindStringSubmatch(errStr)
		if len(matches) == 2 {
			data, _ := hex.DecodeString(matches[1])
			if len(data) >= 4 {
				identifier := data[:4]
				abis, _ := gatherABIS()
				for _, abi := range abis {
					abiError, _ := abi.ErrorByID([4]byte(identifier))
					if abiError != nil {
						errorData, err := abiError.Unpack(data)
						if err == nil {
							return fmt.Errorf("vm error: %v %+v", abiError, errorData)
						} else {
							return fmt.Errorf("vm error: %v", abiError)
						}
					}
				}
			}
		}
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

	matched, err := matchSelector(ToByte4Selector(errorSelector))

	if matched {
		return err
	}

	return fmt.Errorf("error calling contract (could not decode): %s", errStr)
}

func matchSelector(thrownSelector [4]byte) (bool, error) {
	abis, err := gatherABIS()
	// not sure if we should throw the actual error here...
	if err != nil {
		return false, fmt.Errorf("transaction reverted with error: %s", thrownSelector)
	}

	var humanErr error

	// Loop through all the ABIs
	for _, abi := range abis {
		for name, error := range abi.Errors {
			// Compute the selector for the error
			selector := ToByte4Selector(crypto.Keccak256Hash([]byte(error.Sig)).Bytes())

			if selector == thrownSelector {
				humanErr = fmt.Errorf("transaction reverted with error: %s", name)
			}
		}
	}

	for _, humanReadableErrorCode := range humanReadableErrorCodes {
		if humanReadableErrorCode.errCode == thrownSelector {
			humanErr = fmt.Errorf("transaction reverted with error: %s", humanReadableErrorCode.errMsg)
		}
	}

	if humanErr == nil {
		return false, fmt.Errorf("transaction reverted with error: %s", thrownSelector)
	}

	return true, humanErr
}

func gatherABIS() ([]*abi.ABI, error) {
	abis := []*abi.ABI{}

	agentAbi, err := abigen.AgentMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, agentAbi)

	agentFactoryAbi, err := abigen.AgentFactoryMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, agentFactoryAbi)

	credParserAbi, err := abigen.CredParserMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, credParserAbi)

	minerRegistryAbi, err := abigen.MinerRegistryMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, minerRegistryAbi)

	poolTokenAbi, err := abigen.PoolTokenMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, poolTokenAbi)

	routerAbi, err := abigen.RouterMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, routerAbi)

	wrappedFILAbi, err := abigen.WFILMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, wrappedFILAbi)

	poolV2Abi, err := abigen.InfinityPoolV2MetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, poolV2Abi)

	policeV2Abi, err := abigen.AgentPoliceV2MetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, policeV2Abi)

	tokenAbi, err := abigen.TokenMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, tokenAbi)

	plusAbi, err := abigen.SPPlusMetaData.GetAbi()
	if err != nil {
		return []*abi.ABI{}, nil
	}
	abis = append(abis, plusAbi)

	return abis, nil
}
