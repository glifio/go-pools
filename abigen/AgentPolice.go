// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SignedCredential is an auto generated low-level Go binding around an user-defined struct.
type SignedCredential_duplicate struct {
	Vc     VerifiableCredential
	V      uint8
	R      [32]byte
	S      [32]byte
	Blssig []byte
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_duplicate3 struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// AgentPoliceMetaData contains all meta data concerning the AgentPolice contract.
var AgentPoliceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_defaultWindow\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"contractIPoolRegistry\",\"name\":\"_poolRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentStateRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCredential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"Defaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agentID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"faultEpoch\",\"type\":\"uint256\"}],\"name\":\"FaultySectors\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"OnAdministration\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_VERIFIABLE_CREDENTIAL_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrationWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"agentApproved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"agentLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmAdministration\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmEquity\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"credentialUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"deriveStructHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distributeLiquidatedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"isValidCredential\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgent[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"name\":\"markAsFaulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxConsecutiveFaultEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEpochsOwedTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoolsPerAgent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidator\",\"type\":\"uint64\"}],\"name\":\"prepareMinerForLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministrationDueToFaultySectorDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"registerCredentialUseBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sectorFaultyTolerancePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_administrationWindow\",\"type\":\"uint256\"}],\"name\":\"setAdministrationWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setAgentDefaultDueToFaultySectorDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setAgentDefaulted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultWindow\",\"type\":\"uint256\"}],\"name\":\"setDefaultWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxConsecutiveFaultEpochs\",\"type\":\"uint256\"}],\"name\":\"setMaxConsecutiveFaultEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDTE\",\"type\":\"uint256\"}],\"name\":\"setMaxDTE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxEpochsOwedTolerance\",\"type\":\"uint256\"}],\"name\":\"setMaxEpochsOwedTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLTV\",\"type\":\"uint256\"}],\"name\":\"setMaxLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPoolsPerAgent\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolsPerAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sectorFaultyTolerancePercent\",\"type\":\"uint256\"}],\"name\":\"setSectorFaultyTolerancePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"transferOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"validateCred\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610180604052670de0b6b3a7640000600b819055600c5562000025601e610e10620004ed565b6200003290601862000510565b6200003f90600362000510565b600d5562000051601e610e10620004ed565b6200005e90601862000510565b6200006b90600162000510565b600e5566038d7ea4c68000600f556010805460ff191690553480156200009057600080fd5b5060405162004bde38038062004bde833981016040819052620000b3916200061c565b8484818a8a878282620000d6600083620002af60201b62001af71790919060201c565b61012052620000f3816001620002af602090811b62001af717901c565b61014052815160208084019190912060e052815190820120610100524660a0526200018160e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c0526001600160a01b0390811661016052620001b5925083169050620002ff602090811b62001b2817901c565b90506001600160a01b038116620001df57604051635435b28960e11b815260040160405180910390fd5b620001ea816200034b565b506200020a816001600160a01b0316620002ff60201b62001b281760201c565b90506001600160a01b0381166200023457604051635435b28960e11b815260040160405180910390fd5b6200023f816200039e565b5050600886905562000255601e610e10620004ed565b6200026290601862000510565b6200026f90600762000510565b600955600a8055600680546001600160a01b039384166001600160a01b0319918216179091556007805492909316911617905550620008b2945050505050565b6000602083511015620002cf57620002c783620003f1565b9050620002f9565b82620002e6836200043d60201b62001b6f1760201c565b90620002f390826200078c565b5060ff90505b92915050565b600080806200030e8462000440565b91509150816200032057509192915050565b6000806200032e8362000473565b915091508162000342575093949350505050565b95945050505050565b600380546001600160a01b03191690556200037b6001600160a01b038216620002ff602090811b62001b2817901c565b600280546001600160a01b0319166001600160a01b039290921691909117905550565b600580546001600160a01b0319169055620003ce6001600160a01b038216620002ff602090811b62001b2817901c565b600480546001600160a01b0319166001600160a01b039290921691909117905550565b600080829050601f8151111562000428578260405163305a27a960e01b81526004016200041f919062000858565b60405180910390fd5b805162000435826200088d565b179392505050565b90565b600080600160401b600160a01b03831660ff60981b81036200046d57600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620004c75760009250600091505b50811580620004d757503d601614155b15620004e857506000928392509050565b915091565b6000826200050b57634e487b7160e01b600052601260045260246000fd5b500490565b8082028115828204841417620002f957634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60005b83811015620005695781810151838201526020016200054f565b50506000910152565b600082601f8301126200058457600080fd5b81516001600160401b0380821115620005a157620005a162000536565b604051601f8301601f19908116603f01168101908282118183101715620005cc57620005cc62000536565b81604052838152866020858801011115620005e657600080fd5b620005f98460208301602089016200054c565b9695505050505050565b6001600160a01b03811681146200061957600080fd5b50565b600080600080600080600080610100898b0312156200063a57600080fd5b88516001600160401b03808211156200065257600080fd5b620006608c838d0162000572565b995060208b01519150808211156200067757600080fd5b50620006868b828c0162000572565b975050604089015195506060890151620006a08162000603565b60808a0151909550620006b38162000603565b60a08a0151909450620006c68162000603565b60c08a0151909350620006d98162000603565b60e08a0151909250620006ec8162000603565b809150509295985092959890939650565b600181811c908216806200071257607f821691505b6020821081036200073357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200078757600081815260208120601f850160051c81016020861015620007625750805b601f850160051c820191505b8181101562000783578281556001016200076e565b5050505b505050565b81516001600160401b03811115620007a857620007a862000536565b620007c081620007b98454620006fd565b8462000739565b602080601f831160018114620007f85760008415620007df5750858301515b600019600386901b1c1916600185901b17855562000783565b600085815260208120601f198616915b82811015620008295788860151825594840194600190910190840162000808565b5085821015620008485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020815260008251806020840152620008798160408501602087016200054c565b601f01601f19169190910160400192915050565b80516020808301519190811015620007335760001960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516101605161428e6200095060003960008181610a8201528181610af001528181610ddf01528181610e3a01528181610f7b0152818161177401528181611d1c01528181611fc10152612656015260006113ef015260006113c50152600061301701526000612fef01526000612f4a01526000612f7401526000612f9e015261428e6000f3fe608060405234801561001057600080fd5b50600436106103205760003560e01c806369cdb53c116101a7578063ac7e534e116100ee578063cb70f56311610097578063f384bd0511610071578063f384bd051461067f578063fc3e517414610688578063fcad44481461069b57600080fd5b8063cb70f56314610639578063e30c39781461064c578063f2fde38b1461066c57600080fd5b8063bc4c294b116100c8578063bc4c294b14610600578063c6dc030e14610613578063cada34b11461062657600080fd5b8063ac7e534e146105ba578063acec9564146105da578063af10015e146105ed57600080fd5b80638da5cb5b11610150578063a2f3c2101161012a578063a2f3c2101461058b578063a86beff81461059e578063ac76eb40146105b157600080fd5b80638da5cb5b1461054f5780639b83b8cd1461056f5780639fc2f56b1461058257600080fd5b80638456cb59116101815780638456cb591461051957806384b0196e146105215780638b21505a1461053c57600080fd5b806369cdb53c146104f557806378ee44fd1461050857806379ba50971461051157600080fd5b8063395bbfe91161026b578063570ca735116102145780635e7981de116101ee5780635e7981de146104bc57806364af5690146104cf578063680af724146104e257600080fd5b8063570ca735146104575780635c975abb1461049c5780635cf6862f146104a957600080fd5b806351692149116102455780635169214914610428578063561463801461043b5780635654657a1461044457600080fd5b8063395bbfe9146103f957806340444c7c146104025780634107e6441461041557600080fd5b806327ddb79d116102cd5780632c6de7c9116102a75780632c6de7c9146103d5578063338891eb146103e857806337279488146103f057600080fd5b806327ddb79d146103a757806329605e77146103ba5780632b85012b146103cd57600080fd5b80631387b4ed116102fe5780631387b4ed146103555780631f2218de14610371578063260e63c61461039457600080fd5b8063046f7da21461032557806308a0c3751461032f578063117573c514610342575b600080fd5b61032d6106a3565b005b61032d61033d3660046133f7565b6106d5565b61032d6103503660046133f7565b6106e2565b61035e600e5481565b6040519081526020015b60405180910390f35b61038461037f366004613410565b6106ef565b6040519015158152602001610368565b61032d6103a2366004613410565b610718565b61032d6103b536600461344c565b610724565b61032d6103c83660046134f3565b610829565b61035e610878565b61032d6103e3366004613510565b6108ba565b61032d610a7d565b61035e600f5481565b61035e600d5481565b61035e61041036600461378e565b610b5b565b61032d6104233660046133f7565b610cae565b6103846104363660046133f7565b610d0c565b61035e600a5481565b61032d610452366004613410565b610e2a565b6004546104779073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610368565b6010546103849060ff1681565b61032d6104b7366004613410565b610f74565b61032d6104ca3660046137c3565b611057565b61032d6104dd3660046133f7565b61130c565b61032d6104f03660046133f7565b611319565b61032d61050336600461388e565b505050565b61035e60085481565b61032d611326565b61032d611382565b6105296113b7565b6040516103689796959493929190613953565b61032d61054a3660046134f3565b61145b565b6002546104779073ffffffffffffffffffffffffffffffffffffffff1681565b61032d61057d366004613a12565b61159b565b61035e600b5481565b61035e61059936600461378e565b611639565b61032d6105ac3660046134f3565b611652565b61035e60095481565b6005546104779073ffffffffffffffffffffffffffffffffffffffff1681565b61032d6105e83660046133f7565b611748565b61032d6105fb3660046133f7565b611755565b61032d61060e3660046133f7565b611762565b61032d610621366004613a57565b61176f565b61032d610634366004613a8c565b61186e565b61032d610647366004613510565b6118c0565b6003546104779073ffffffffffffffffffffffffffffffffffffffff1681565b61032d61067a3660046134f3565b611a08565b61035e600c5481565b610477610696366004613a57565b611a77565b61032d611a9d565b6106ab611b72565b601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6106dd611b72565b600c55565b6106ea611b72565b600b55565b60008060118161070161059986613aea565b815260200190815260200160002054119050919050565b61072181611bc3565b50565b61072c611e73565b6000805b828110156108235783838281811061074a5761074a613af6565b905060200201602081019061075f91906134f3565b91508173ffffffffffffffffffffffffffffffffffffffff1663a71804496040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156107a957600080fd5b505af11580156107bd573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff167fae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd64360405161080991815260200190565b60405180910390a28061081b81613b54565b915050610730565b50505050565b610831611b72565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6040518060c00160405280609681526020016141c3609691396040516020016108a19190613b8c565b6040516020818303038152906040528051906020012081565b6108c2611b72565b8160095461093e8273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610914573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109389190613ba8565b82611eea565b610974576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff16639e4d74cc6109af8573ffffffffffffffffffffffffffffffffffffffff16611b28565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401600060405180830381600087803b158015610a1557600080fd5b505af1158015610a29573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff871681527fcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c29250602001905060405180910390a150505050565b610aa67f0000000000000000000000000000000000000000000000000000000000000000612023565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055610b147f0000000000000000000000000000000000000000000000000000000000000000612107565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006040518060c00160405280609681526020016141c360969139604051602001610b869190613b8c565b60405160208183030381529060405280519060200120826000015183602001518460400151856060015186608001518760a001518860c001518960e00151604051602001610bd49190613b8c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083019a909a5273ffffffffffffffffffffffffffffffffffffffff909816978101979097526060870195909552608086019390935260a085019190915260c08401527fffffffff000000000000000000000000000000000000000000000000000000001660e083015267ffffffffffffffff1661010082015261012081019190915261014001604051602081830303815290604052805190602001209050919050565b610cb6611b72565b610cc3601e610e10613bc1565b610cce906018613bfc565b811115610d07576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e55565b6006546040517f9a014e2400000000000000000000000000000000000000000000000000000000815260048101839052600091829173ffffffffffffffffffffffffffffffffffffffff90911690639a014e2490602401600060405180830381865afa158015610d80573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610dc69190810190613c13565b90508051600003610dda5750600092915050565b610e1f7f00000000000000000000000000000000000000000000000000000000000000008483600081518110610e1257610e12613af6565b60200260200101516121ae565b606001519392505050565b610e3381611bc3565b6000610e5e7f0000000000000000000000000000000000000000000000000000000000000000612279565b90506000610e7582610e6f85613aea565b90612320565b905080600003610e8457505050565b6000610e9983610e9386613aea565b906123c0565b9050818111610ed4576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b54610eeb610ee48484613cb9565b8490612418565b1115610f23576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c54610f3c610ee485610f3688613aea565b9061242d565b1115610823576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610f9f7f0000000000000000000000000000000000000000000000000000000000000000612279565b9050600f54610fd4610fbb8385610fb590613aea565b90612485565b610fce84610fc887613aea565b906124dd565b90612418565b111561100c576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61101c8260200135600954611eea565b15611053576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b61105f611b72565b60008273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d09190613ba8565b90506110db81610d0c565b15611112576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af115801561118f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b39190613cdc565b5060006111c08284612535565b9050600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8573ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561124b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126f9190613cf7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af11580156112e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113059190613cdc565b5050505050565b611314611b72565b600a55565b611321611b72565b600f55565b60035473ffffffffffffffffffffffffffffffffffffffff163314611377576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611380336128c9565b565b61138a611b72565b601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b6000606080828080836113ea7f000000000000000000000000000000000000000000000000000000000000000083612958565b6114157f00000000000000000000000000000000000000000000000000000000000000006001612958565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b611463611b72565b806008546114b58273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610914573d6000803e3d6000fd5b6114eb576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663615d92116040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561153357600080fd5b505af1158015611547573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff861681527f0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1925060200190505b60405180910390a1505050565b6115a3611b72565b6040517ff794a92300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff80841660048301528216602482015273ffffffffffffffffffffffffffffffffffffffff84169063f794a92390604401600060405180830381600087803b15801561161c57600080fd5b505af1158015611630573d6000803e3d6000fd5b50505050505050565b600061164c61164783610b5b565b6129fc565b92915050565b61165a611b72565b8061166481612a44565b61169a576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663615d92116040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156116e257600080fd5b505af11580156116f6573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff851681527f0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b19250602001905060405180910390a15050565b611750611b72565b600855565b61175d611b72565b600d55565b61176a611b72565b600955565b6117997f000000000000000000000000000000000000000000000000000000000000000033612ae0565b8060000151602001513373ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118119190613ba8565b14611848576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b436011600061185a8460000151611639565b815260208101919091526040016000205550565b61187c838361050384613d14565b61188961037f8280613d20565b15610503576040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118c8611b72565b816118d281612a44565b611908576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16639e4d74cc6119438473ffffffffffffffffffffffffffffffffffffffff16611b28565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401600060405180830381600087803b1580156119a957600080fd5b505af11580156119bd573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff861681527fcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c29250602001905061158e565b611a10611b72565b611a2f8173ffffffffffffffffffffffffffffffffffffffff16611b28565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b600061164c611a898360000151611639565b836020015184604001518560600151612bb1565b60055473ffffffffffffffffffffffffffffffffffffffff163314611aee576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61138033612bd9565b6000602083511015611b1357611b0c83612c68565b905061164c565b81611b1e8482613df5565b5060ff9392505050565b6000806000611b3684612cc8565b9150915081611b4757509192915050565b600080611b5383612d15565b9150915081611b66575093949350505050565b95945050505050565b90565b60025473ffffffffffffffffffffffffffffffffffffffff163314611380576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006546040517f9a014e240000000000000000000000000000000000000000000000000000000081526020830135600482015260009173ffffffffffffffffffffffffffffffffffffffff1690639a014e2490602401600060405180830381865afa158015611c36573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611c7c9190810190613c13565b9050600a5481511115611cbb576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8151811015610503576000828281518110611cdb57611cdb613af6565b602002602001015190506000611d13600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683612d99565b90506000611d467f00000000000000000000000000000000000000000000000000000000000000008760200135856121ae565b6040517f9741fbfa00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff831690639741fbfa90611d9d9084908a90600401613fc3565b602060405180830381865afa158015611dba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dde9190613cdc565b611e14576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b43600e548260400151611e2791906140c1565b1015611e5f576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505080611e6c90613b54565b9050611cbe565b60045473ffffffffffffffffffffffffffffffffffffffff163314801590611eb3575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15611380576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006546040517f9a014e2400000000000000000000000000000000000000000000000000000000815260048101849052600091829173ffffffffffffffffffffffffffffffffffffffff90911690639a014e2490602401600060405180830381865afa158015611f5e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611fa49190810190613c13565b905060005b815181101561201857611fbc8443613cb9565b611ff37f000000000000000000000000000000000000000000000000000000000000000087858581518110610e1257610e12613af6565b6040015110156120085760019250505061164c565b61201181613b54565b9050611fa9565b506000949350505050565b604080518082018252601481527f524f555445525f504f4f4c5f5245474953545259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f390c2e8200000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024015b602060405180830381865afa1580156120e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164c9190613cf7565b604080518082018252601181527f524f555445525f5746494c5f544f4b454e000000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527faee0a13500000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016120c6565b6121db60405180608001604052806000815260200160008152602001600081526020016000151581525090565b6040517f6361f6de000000000000000000000000000000000000000000000000000000008152600481018490526024810183905273ffffffffffffffffffffffffffffffffffffffff851690636361f6de90604401608060405180830381865afa15801561224d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061227191906140d4565b949350505050565b604080518082018252601281527f524f555445525f435245445f5041525345520000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527fd26df3b700000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016120c6565b60e08201516040517f17bbd29a00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff8416916317bbd29a916123789160040161413e565b602060405180830381865afa158015612395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123b99190613ba8565b9392505050565b60e08201516040517f6bb0d0cc00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff841691636bb0d0cc916123789160040161413e565b60006123b983670de0b6b3a764000084612ed2565b60e08201516040517f1b2b5fad00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff841691631b2b5fad916123789160040161413e565b60e08201516040517f402ed8cb00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff84169163402ed8cb916123789160040161413e565b60e08201516040517f07124c0600000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff8416916307124c06916123789160040161413e565b6006546040517f9a014e240000000000000000000000000000000000000000000000000000000081526004810184905260009182918291829182918291829173ffffffffffffffffffffffffffffffffffffffff1690639a014e2490602401600060405180830381865afa1580156125b1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526125f79190810190613c13565b80519091506000808267ffffffffffffffff81111561261857612618613549565b604051908082528060200260200182016040528015612641578160200160208202803683370190505b509050600097505b828810156126c8576126887f00000000000000000000000000000000000000000000000000000000000000008d868b81518110610e1257610e12613af6565b602001519450848189815181106126a1576126a1613af6565b60209081029190910101526126b6858a6140c1565b98506126c188613b54565b9750612649565b600097505b828810156128ba578388815181106126e7576126e7613af6565b60200260200101519650888b828a8151811061270557612705613af6565b60200260200101516127179190613bfc565b6127219190613bc1565b60065490965060009061274a9073ffffffffffffffffffffffffffffffffffffffff1689612d99565b6007546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8084166004830152602482018b905292935091169063095ea7b3906044016020604051808303816000875af11580156127c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127e99190613cdc565b506040517fc564f772000000000000000000000000000000000000000000000000000000008152600481018e90526024810188905273ffffffffffffffffffffffffffffffffffffffff82169063c564f772906044016020604051808303816000875af115801561285e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128829190613ba8565b925082871161289257600061289c565b61289c8388613cb9565b6128a6908c6140c1565b9a5050876128b390613b54565b97506126cd565b50505050505050505092915050565b600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016905561291073ffffffffffffffffffffffffffffffffffffffff8216611b28565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b606060ff831461296b57611b0c83612ef1565b81805461297790613d54565b80601f01602080910402602001604051908101604052809291908181526020018280546129a390613d54565b80156129f05780601f106129c5576101008083540402835291602001916129f0565b820191906000526020600020905b8154815290600101906020018083116129d357829003601f168201915b5050505050905061164c565b600061164c612a09612f30565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000808273ffffffffffffffffffffffffffffffffffffffff16638903f7f06040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ab69190613ba8565b905080600003612ac95750600092915050565b600d54612ad690826140c1565b4310159392505050565b612ae98261306d565b6040517f1ffbb06400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529190911690631ffbb06490602401602060405180830381865afa158015612b57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b7b9190613cdc565b611053576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806000612bc287878787613114565b91509150612bcf81613203565b5095945050505050565b600580547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055612c2073ffffffffffffffffffffffffffffffffffffffff8216611b28565b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b600080829050601f81511115612cb557826040517f305a27a9000000000000000000000000000000000000000000000000000000008152600401612cac919061413e565b60405180910390fd5b8051612cc082614151565b179392505050565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103612d0f576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a8114612d755760009250600091505b50811580612d8457503d601614155b15612d9457506000928392509050565b915091565b60008273ffffffffffffffffffffffffffffffffffffffff1663efde4e646040518163ffffffff1660e01b8152600401602060405180830381865afa158015612de6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e0a9190613ba8565b821115612e43576040517f8f5d631800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f41d1de970000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8416906341d1de9790602401602060405180830381865afa158015612eae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123b99190613cf7565b828202811515841585830485141716612eea57600080fd5b0492915050565b60606000612efe836133b6565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b60003073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015612f9657507f000000000000000000000000000000000000000000000000000000000000000046145b15612fc057507f000000000000000000000000000000000000000000000000000000000000000090565b613068604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b905090565b604080518082018252601481527f524f555445525f4147454e545f464143544f5259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f29f616da00000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016120c6565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561314b57506000905060036131fa565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561319f573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166131f3576000600192509250506131fa565b9150600090505b94509492505050565b600081600481111561321757613217614193565b0361321f5750565b600181600481111561323357613233614193565b0361329a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401612cac565b60028160048111156132ae576132ae614193565b03613315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401612cac565b600381600481111561332957613329614193565b03610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401612cac565b600060ff8216601f81111561164c576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561340957600080fd5b5035919050565b60006020828403121561342257600080fd5b813567ffffffffffffffff81111561343957600080fd5b820161010081850312156123b957600080fd5b6000806020838503121561345f57600080fd5b823567ffffffffffffffff8082111561347757600080fd5b818501915085601f83011261348b57600080fd5b81358181111561349a57600080fd5b8660208260051b85010111156134af57600080fd5b60209290920196919550909350505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461072157600080fd5b80356134ee816134c1565b919050565b60006020828403121561350557600080fd5b81356123b9816134c1565b6000806040838503121561352357600080fd5b823561352e816134c1565b9150602083013561353e816134c1565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff8111828210171561359c5761359c613549565b60405290565b60405160a0810167ffffffffffffffff8111828210171561359c5761359c613549565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561360c5761360c613549565b604052919050565b80357fffffffff00000000000000000000000000000000000000000000000000000000811681146134ee57600080fd5b803567ffffffffffffffff811681146134ee57600080fd5b600082601f83011261366d57600080fd5b813567ffffffffffffffff81111561368757613687613549565b6136b860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016135c5565b8181528460208386010111156136cd57600080fd5b816020850160208301376000918101602001919091529392505050565b600061010082840312156136fd57600080fd5b613705613578565b9050613710826134e3565b81526020820135602082015260408201356040820152606082013560608201526080820135608082015261374660a08301613614565b60a082015261375760c08301613644565b60c082015260e082013567ffffffffffffffff81111561377657600080fd5b6137828482850161365c565b60e08301525092915050565b6000602082840312156137a057600080fd5b813567ffffffffffffffff8111156137b757600080fd5b612271848285016136ea565b600080604083850312156137d657600080fd5b82356137e1816134c1565b946020939093013593505050565b600060a0828403121561380157600080fd5b6138096135a2565b9050813567ffffffffffffffff8082111561382357600080fd5b61382f858386016136ea565b83526020840135915060ff8216821461384757600080fd5b8160208401526040840135604084015260608401356060840152608084013591508082111561387557600080fd5b506138828482850161365c565b60808301525092915050565b6000806000606084860312156138a357600080fd5b833592506138b360208501613614565b9150604084013567ffffffffffffffff8111156138cf57600080fd5b6138db868287016137ef565b9150509250925092565b60005b838110156139005781810151838201526020016138e8565b50506000910152565b600081518084526139218160208601602086016138e5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b7fff00000000000000000000000000000000000000000000000000000000000000881681526000602060e08184015261398f60e084018a613909565b83810360408501526139a1818a613909565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825283870192509083019060005b81811015613a00578351835292840192918401916001016139e4565b50909c9b505050505050505050505050565b600080600060608486031215613a2757600080fd5b8335613a32816134c1565b9250613a4060208501613644565b9150613a4e60408501613644565b90509250925092565b600060208284031215613a6957600080fd5b813567ffffffffffffffff811115613a8057600080fd5b612271848285016137ef565b600080600060608486031215613aa157600080fd5b83359250613ab160208501613614565b9150604084013567ffffffffffffffff811115613acd57600080fd5b840160a08187031215613adf57600080fd5b809150509250925092565b600061164c36836136ea565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613b8557613b85613b25565b5060010190565b60008251613b9e8184602087016138e5565b9190910192915050565b600060208284031215613bba57600080fd5b5051919050565b600082613bf7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b808202811582820484141761164c5761164c613b25565b60006020808385031215613c2657600080fd5b825167ffffffffffffffff80821115613c3e57600080fd5b818501915085601f830112613c5257600080fd5b815181811115613c6457613c64613549565b8060051b9150613c758483016135c5565b8181529183018401918481019088841115613c8f57600080fd5b938501935b83851015613cad57845182529385019390850190613c94565b98975050505050505050565b8181038181111561164c5761164c613b25565b805180151581146134ee57600080fd5b600060208284031215613cee57600080fd5b6123b982613ccc565b600060208284031215613d0957600080fd5b81516123b9816134c1565b600061164c36836137ef565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01833603018112613b9e57600080fd5b600181811c90821680613d6857607f821691505b602082108103613da1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561050357600081815260208120601f850160051c81016020861015613dce5750805b601f850160051c820191505b81811015613ded57828155600101613dda565b505050505050565b815167ffffffffffffffff811115613e0f57613e0f613549565b613e2381613e1d8454613d54565b84613da7565b602080601f831160018114613e765760008415613e405750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613ded565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015613ec357888601518255948401946001909101908401613ea4565b5085821015613eff57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613f4457600080fd5b830160208101925035905067ffffffffffffffff811115613f6457600080fd5b803603821315613f7357600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b82518152602083015160208201526040830151604082015260608301511515606082015260a0608082015260008235613ffb816134c1565b73ffffffffffffffffffffffffffffffffffffffff811660a084015250602083013560c0830152604083013560e0830152610100606084013581840152608084013561012084015261404f60a08501613614565b7fffffffff000000000000000000000000000000000000000000000000000000001661014084015261408360c08501613644565b67ffffffffffffffff166101608401526140a060e0850185613f0f565b826101808601526140b66101a086018284613f7a565b979650505050505050565b8082018082111561164c5761164c613b25565b6000608082840312156140e657600080fd5b6040516080810181811067ffffffffffffffff8211171561410957614109613549565b806040525082518152602083015160208201526040830151604082015261413260608401613ccc565b60608201529392505050565b6020815260006123b96020830184613909565b80516020808301519190811015613da1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfe56657269666961626c6543726564656e7469616c2861646472657373206973737565722c75696e74323536207375626a6563742c75696e743235362065706f63684973737565642c75696e743235362065706f636856616c6964556e74696c2c75696e743235362076616c75652c62797465733420616374696f6e2c75696e743634207461726765742c627974657320636c61696d29a2646970667358221220c2014f77231e5cadab6e06536fd15a22e6f8928396665ba277b1090a0048250264736f6c63430008110033",
}

// AgentPoliceABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentPoliceMetaData.ABI instead.
var AgentPoliceABI = AgentPoliceMetaData.ABI

// AgentPoliceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgentPoliceMetaData.Bin instead.
var AgentPoliceBin = AgentPoliceMetaData.Bin

// DeployAgentPolice deploys a new Ethereum contract, binding an instance of AgentPolice to it.
func DeployAgentPolice(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _version string, _defaultWindow *big.Int, _owner common.Address, _operator common.Address, _router common.Address, _poolRegistry common.Address, _wFIL common.Address) (common.Address, *types.Transaction, *AgentPolice, error) {
	parsed, err := AgentPoliceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgentPoliceBin), backend, _name, _version, _defaultWindow, _owner, _operator, _router, _poolRegistry, _wFIL)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentPolice{AgentPoliceCaller: AgentPoliceCaller{contract: contract}, AgentPoliceTransactor: AgentPoliceTransactor{contract: contract}, AgentPoliceFilterer: AgentPoliceFilterer{contract: contract}}, nil
}

// AgentPolice is an auto generated Go binding around an Ethereum contract.
type AgentPolice struct {
	AgentPoliceCaller     // Read-only binding to the contract
	AgentPoliceTransactor // Write-only binding to the contract
	AgentPoliceFilterer   // Log filterer for contract events
}

// AgentPoliceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentPoliceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentPoliceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentPoliceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentPoliceSession struct {
	Contract     *AgentPolice      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentPoliceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentPoliceCallerSession struct {
	Contract *AgentPoliceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AgentPoliceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentPoliceTransactorSession struct {
	Contract     *AgentPoliceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AgentPoliceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentPoliceRaw struct {
	Contract *AgentPolice // Generic contract binding to access the raw methods on
}

// AgentPoliceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentPoliceCallerRaw struct {
	Contract *AgentPoliceCaller // Generic read-only contract binding to access the raw methods on
}

// AgentPoliceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentPoliceTransactorRaw struct {
	Contract *AgentPoliceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentPolice creates a new instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPolice(address common.Address, backend bind.ContractBackend) (*AgentPolice, error) {
	contract, err := bindAgentPolice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentPolice{AgentPoliceCaller: AgentPoliceCaller{contract: contract}, AgentPoliceTransactor: AgentPoliceTransactor{contract: contract}, AgentPoliceFilterer: AgentPoliceFilterer{contract: contract}}, nil
}

// NewAgentPoliceCaller creates a new read-only instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPoliceCaller(address common.Address, caller bind.ContractCaller) (*AgentPoliceCaller, error) {
	contract, err := bindAgentPolice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceCaller{contract: contract}, nil
}

// NewAgentPoliceTransactor creates a new write-only instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPoliceTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentPoliceTransactor, error) {
	contract, err := bindAgentPolice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceTransactor{contract: contract}, nil
}

// NewAgentPoliceFilterer creates a new log filterer instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPoliceFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentPoliceFilterer, error) {
	contract, err := bindAgentPolice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceFilterer{contract: contract}, nil
}

// bindAgentPolice binds a generic wrapper to an already deployed contract.
func bindAgentPolice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentPoliceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentPolice *AgentPoliceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentPolice.Contract.AgentPoliceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentPolice *AgentPoliceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.Contract.AgentPoliceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentPolice *AgentPoliceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentPolice.Contract.AgentPoliceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentPolice *AgentPoliceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentPolice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentPolice *AgentPoliceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentPolice *AgentPoliceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentPolice.Contract.contract.Transact(opts, method, params...)
}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPolice *AgentPoliceCaller) VERIFIABLECREDENTIALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "_VERIFIABLE_CREDENTIAL_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPolice *AgentPoliceSession) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _AgentPolice.Contract.VERIFIABLECREDENTIALTYPEHASH(&_AgentPolice.CallOpts)
}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPolice *AgentPoliceCallerSession) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _AgentPolice.Contract.VERIFIABLECREDENTIALTYPEHASH(&_AgentPolice.CallOpts)
}

// AdministrationWindow is a free data retrieval call binding the contract method 0xac76eb40.
//
// Solidity: function administrationWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) AdministrationWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "administrationWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdministrationWindow is a free data retrieval call binding the contract method 0xac76eb40.
//
// Solidity: function administrationWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) AdministrationWindow() (*big.Int, error) {
	return _AgentPolice.Contract.AdministrationWindow(&_AgentPolice.CallOpts)
}

// AdministrationWindow is a free data retrieval call binding the contract method 0xac76eb40.
//
// Solidity: function administrationWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) AdministrationWindow() (*big.Int, error) {
	return _AgentPolice.Contract.AdministrationWindow(&_AgentPolice.CallOpts)
}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCaller) AgentApproved(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "agentApproved", vc)

	if err != nil {
		return err
	}

	return err

}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceSession) AgentApproved(vc VerifiableCredential) error {
	return _AgentPolice.Contract.AgentApproved(&_AgentPolice.CallOpts, vc)
}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) AgentApproved(vc VerifiableCredential) error {
	return _AgentPolice.Contract.AgentApproved(&_AgentPolice.CallOpts, vc)
}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPolice *AgentPoliceCaller) AgentLiquidated(opts *bind.CallOpts, agentID *big.Int) (bool, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "agentLiquidated", agentID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPolice *AgentPoliceSession) AgentLiquidated(agentID *big.Int) (bool, error) {
	return _AgentPolice.Contract.AgentLiquidated(&_AgentPolice.CallOpts, agentID)
}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPolice *AgentPoliceCallerSession) AgentLiquidated(agentID *big.Int) (bool, error) {
	return _AgentPolice.Contract.AgentLiquidated(&_AgentPolice.CallOpts, agentID)
}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCaller) ConfirmRmAdministration(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "confirmRmAdministration", vc)

	if err != nil {
		return err
	}

	return err

}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceSession) ConfirmRmAdministration(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmAdministration(&_AgentPolice.CallOpts, vc)
}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) ConfirmRmAdministration(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmAdministration(&_AgentPolice.CallOpts, vc)
}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCaller) ConfirmRmEquity(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "confirmRmEquity", vc)

	if err != nil {
		return err
	}

	return err

}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceSession) ConfirmRmEquity(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmEquity(&_AgentPolice.CallOpts, vc)
}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) ConfirmRmEquity(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmEquity(&_AgentPolice.CallOpts, vc)
}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_AgentPolice *AgentPoliceCaller) CredentialUsed(opts *bind.CallOpts, vc VerifiableCredential) (bool, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "credentialUsed", vc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_AgentPolice *AgentPoliceSession) CredentialUsed(vc VerifiableCredential) (bool, error) {
	return _AgentPolice.Contract.CredentialUsed(&_AgentPolice.CallOpts, vc)
}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_AgentPolice *AgentPoliceCallerSession) CredentialUsed(vc VerifiableCredential) (bool, error) {
	return _AgentPolice.Contract.CredentialUsed(&_AgentPolice.CallOpts, vc)
}

// DefaultWindow is a free data retrieval call binding the contract method 0x78ee44fd.
//
// Solidity: function defaultWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) DefaultWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "defaultWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultWindow is a free data retrieval call binding the contract method 0x78ee44fd.
//
// Solidity: function defaultWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) DefaultWindow() (*big.Int, error) {
	return _AgentPolice.Contract.DefaultWindow(&_AgentPolice.CallOpts)
}

// DefaultWindow is a free data retrieval call binding the contract method 0x78ee44fd.
//
// Solidity: function defaultWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) DefaultWindow() (*big.Int, error) {
	return _AgentPolice.Contract.DefaultWindow(&_AgentPolice.CallOpts)
}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPolice *AgentPoliceCaller) DeriveStructHash(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "deriveStructHash", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPolice *AgentPoliceSession) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.DeriveStructHash(&_AgentPolice.CallOpts, vc)
}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPolice *AgentPoliceCallerSession) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.DeriveStructHash(&_AgentPolice.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPolice *AgentPoliceCaller) Digest(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "digest", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPolice *AgentPoliceSession) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.Digest(&_AgentPolice.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPolice *AgentPoliceCallerSession) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.Digest(&_AgentPolice.CallOpts, vc)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentPolice *AgentPoliceCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentPolice *AgentPoliceSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentPolice.Contract.Eip712Domain(&_AgentPolice.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentPolice *AgentPoliceCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentPolice.Contract.Eip712Domain(&_AgentPolice.CallOpts)
}

// IsValidCredential is a free data retrieval call binding the contract method 0xcada34b1.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns()
func (_AgentPolice *AgentPoliceCaller) IsValidCredential(opts *bind.CallOpts, agent *big.Int, action [4]byte, sc SignedCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "isValidCredential", agent, action, sc)

	if err != nil {
		return err
	}

	return err

}

// IsValidCredential is a free data retrieval call binding the contract method 0xcada34b1.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns()
func (_AgentPolice *AgentPoliceSession) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.IsValidCredential(&_AgentPolice.CallOpts, agent, action, sc)
}

// IsValidCredential is a free data retrieval call binding the contract method 0xcada34b1.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.IsValidCredential(&_AgentPolice.CallOpts, agent, action, sc)
}

// MaxConsecutiveFaultEpochs is a free data retrieval call binding the contract method 0x395bbfe9.
//
// Solidity: function maxConsecutiveFaultEpochs() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxConsecutiveFaultEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxConsecutiveFaultEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxConsecutiveFaultEpochs is a free data retrieval call binding the contract method 0x395bbfe9.
//
// Solidity: function maxConsecutiveFaultEpochs() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxConsecutiveFaultEpochs() (*big.Int, error) {
	return _AgentPolice.Contract.MaxConsecutiveFaultEpochs(&_AgentPolice.CallOpts)
}

// MaxConsecutiveFaultEpochs is a free data retrieval call binding the contract method 0x395bbfe9.
//
// Solidity: function maxConsecutiveFaultEpochs() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxConsecutiveFaultEpochs() (*big.Int, error) {
	return _AgentPolice.Contract.MaxConsecutiveFaultEpochs(&_AgentPolice.CallOpts)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxDTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxDTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxDTE() (*big.Int, error) {
	return _AgentPolice.Contract.MaxDTE(&_AgentPolice.CallOpts)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxDTE() (*big.Int, error) {
	return _AgentPolice.Contract.MaxDTE(&_AgentPolice.CallOpts)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxEpochsOwedTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxEpochsOwedTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _AgentPolice.Contract.MaxEpochsOwedTolerance(&_AgentPolice.CallOpts)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _AgentPolice.Contract.MaxEpochsOwedTolerance(&_AgentPolice.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxLTV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxLTV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxLTV() (*big.Int, error) {
	return _AgentPolice.Contract.MaxLTV(&_AgentPolice.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxLTV() (*big.Int, error) {
	return _AgentPolice.Contract.MaxLTV(&_AgentPolice.CallOpts)
}

// MaxPoolsPerAgent is a free data retrieval call binding the contract method 0x56146380.
//
// Solidity: function maxPoolsPerAgent() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxPoolsPerAgent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxPoolsPerAgent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPoolsPerAgent is a free data retrieval call binding the contract method 0x56146380.
//
// Solidity: function maxPoolsPerAgent() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxPoolsPerAgent() (*big.Int, error) {
	return _AgentPolice.Contract.MaxPoolsPerAgent(&_AgentPolice.CallOpts)
}

// MaxPoolsPerAgent is a free data retrieval call binding the contract method 0x56146380.
//
// Solidity: function maxPoolsPerAgent() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxPoolsPerAgent() (*big.Int, error) {
	return _AgentPolice.Contract.MaxPoolsPerAgent(&_AgentPolice.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPolice *AgentPoliceCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPolice *AgentPoliceSession) Operator() (common.Address, error) {
	return _AgentPolice.Contract.Operator(&_AgentPolice.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) Operator() (common.Address, error) {
	return _AgentPolice.Contract.Operator(&_AgentPolice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPolice *AgentPoliceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPolice *AgentPoliceSession) Owner() (common.Address, error) {
	return _AgentPolice.Contract.Owner(&_AgentPolice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) Owner() (common.Address, error) {
	return _AgentPolice.Contract.Owner(&_AgentPolice.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPolice *AgentPoliceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPolice *AgentPoliceSession) Paused() (bool, error) {
	return _AgentPolice.Contract.Paused(&_AgentPolice.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPolice *AgentPoliceCallerSession) Paused() (bool, error) {
	return _AgentPolice.Contract.Paused(&_AgentPolice.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPolice *AgentPoliceCaller) PendingOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "pendingOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPolice *AgentPoliceSession) PendingOperator() (common.Address, error) {
	return _AgentPolice.Contract.PendingOperator(&_AgentPolice.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) PendingOperator() (common.Address, error) {
	return _AgentPolice.Contract.PendingOperator(&_AgentPolice.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPolice *AgentPoliceCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPolice *AgentPoliceSession) PendingOwner() (common.Address, error) {
	return _AgentPolice.Contract.PendingOwner(&_AgentPolice.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) PendingOwner() (common.Address, error) {
	return _AgentPolice.Contract.PendingOwner(&_AgentPolice.CallOpts)
}

// Recover is a free data retrieval call binding the contract method 0xfc3e5174.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns(address)
func (_AgentPolice *AgentPoliceCaller) Recover(opts *bind.CallOpts, sc SignedCredential) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "recover", sc)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0xfc3e5174.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns(address)
func (_AgentPolice *AgentPoliceSession) Recover(sc SignedCredential) (common.Address, error) {
	return _AgentPolice.Contract.Recover(&_AgentPolice.CallOpts, sc)
}

// Recover is a free data retrieval call binding the contract method 0xfc3e5174.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) Recover(sc SignedCredential) (common.Address, error) {
	return _AgentPolice.Contract.Recover(&_AgentPolice.CallOpts, sc)
}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) SectorFaultyTolerancePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "sectorFaultyTolerancePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _AgentPolice.Contract.SectorFaultyTolerancePercent(&_AgentPolice.CallOpts)
}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _AgentPolice.Contract.SectorFaultyTolerancePercent(&_AgentPolice.CallOpts)
}

// ValidateCred is a free data retrieval call binding the contract method 0x69cdb53c.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns()
func (_AgentPolice *AgentPoliceCaller) ValidateCred(opts *bind.CallOpts, agent *big.Int, selector [4]byte, sc SignedCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "validateCred", agent, selector, sc)

	if err != nil {
		return err
	}

	return err

}

// ValidateCred is a free data retrieval call binding the contract method 0x69cdb53c.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns()
func (_AgentPolice *AgentPoliceSession) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.ValidateCred(&_AgentPolice.CallOpts, agent, selector, sc)
}

// ValidateCred is a free data retrieval call binding the contract method 0x69cdb53c.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.ValidateCred(&_AgentPolice.CallOpts, agent, selector, sc)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPolice *AgentPoliceTransactor) AcceptOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "acceptOperator")
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPolice *AgentPoliceSession) AcceptOperator() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOperator(&_AgentPolice.TransactOpts)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPolice *AgentPoliceTransactorSession) AcceptOperator() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOperator(&_AgentPolice.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPolice *AgentPoliceTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPolice *AgentPoliceSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOwnership(&_AgentPolice.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPolice *AgentPoliceTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOwnership(&_AgentPolice.TransactOpts)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPolice *AgentPoliceTransactor) DistributeLiquidatedFunds(opts *bind.TransactOpts, agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "distributeLiquidatedFunds", agent, amount)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPolice *AgentPoliceSession) DistributeLiquidatedFunds(agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.DistributeLiquidatedFunds(&_AgentPolice.TransactOpts, agent, amount)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPolice *AgentPoliceTransactorSession) DistributeLiquidatedFunds(agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.DistributeLiquidatedFunds(&_AgentPolice.TransactOpts, agent, amount)
}

// MarkAsFaulty is a paid mutator transaction binding the contract method 0x27ddb79d.
//
// Solidity: function markAsFaulty(address[] agents) returns()
func (_AgentPolice *AgentPoliceTransactor) MarkAsFaulty(opts *bind.TransactOpts, agents []common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "markAsFaulty", agents)
}

// MarkAsFaulty is a paid mutator transaction binding the contract method 0x27ddb79d.
//
// Solidity: function markAsFaulty(address[] agents) returns()
func (_AgentPolice *AgentPoliceSession) MarkAsFaulty(agents []common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.MarkAsFaulty(&_AgentPolice.TransactOpts, agents)
}

// MarkAsFaulty is a paid mutator transaction binding the contract method 0x27ddb79d.
//
// Solidity: function markAsFaulty(address[] agents) returns()
func (_AgentPolice *AgentPoliceTransactorSession) MarkAsFaulty(agents []common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.MarkAsFaulty(&_AgentPolice.TransactOpts, agents)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPolice *AgentPoliceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPolice *AgentPoliceSession) Pause() (*types.Transaction, error) {
	return _AgentPolice.Contract.Pause(&_AgentPolice.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPolice *AgentPoliceTransactorSession) Pause() (*types.Transaction, error) {
	return _AgentPolice.Contract.Pause(&_AgentPolice.TransactOpts)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPolice *AgentPoliceTransactor) PrepareMinerForLiquidation(opts *bind.TransactOpts, agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "prepareMinerForLiquidation", agent, miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPolice *AgentPoliceSession) PrepareMinerForLiquidation(agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPolice.Contract.PrepareMinerForLiquidation(&_AgentPolice.TransactOpts, agent, miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPolice *AgentPoliceTransactorSession) PrepareMinerForLiquidation(agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPolice.Contract.PrepareMinerForLiquidation(&_AgentPolice.TransactOpts, agent, miner, liquidator)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x2c6de7c9.
//
// Solidity: function putAgentOnAdministration(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactor) PutAgentOnAdministration(opts *bind.TransactOpts, agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "putAgentOnAdministration", agent, administration)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x2c6de7c9.
//
// Solidity: function putAgentOnAdministration(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceSession) PutAgentOnAdministration(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministration(&_AgentPolice.TransactOpts, agent, administration)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x2c6de7c9.
//
// Solidity: function putAgentOnAdministration(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactorSession) PutAgentOnAdministration(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministration(&_AgentPolice.TransactOpts, agent, administration)
}

// PutAgentOnAdministrationDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xcb70f563.
//
// Solidity: function putAgentOnAdministrationDueToFaultySectorDays(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactor) PutAgentOnAdministrationDueToFaultySectorDays(opts *bind.TransactOpts, agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "putAgentOnAdministrationDueToFaultySectorDays", agent, administration)
}

// PutAgentOnAdministrationDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xcb70f563.
//
// Solidity: function putAgentOnAdministrationDueToFaultySectorDays(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceSession) PutAgentOnAdministrationDueToFaultySectorDays(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministrationDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent, administration)
}

// PutAgentOnAdministrationDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xcb70f563.
//
// Solidity: function putAgentOnAdministrationDueToFaultySectorDays(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactorSession) PutAgentOnAdministrationDueToFaultySectorDays(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministrationDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent, administration)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_AgentPolice *AgentPoliceTransactor) RefreshRoutes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "refreshRoutes")
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_AgentPolice *AgentPoliceSession) RefreshRoutes() (*types.Transaction, error) {
	return _AgentPolice.Contract.RefreshRoutes(&_AgentPolice.TransactOpts)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_AgentPolice *AgentPoliceTransactorSession) RefreshRoutes() (*types.Transaction, error) {
	return _AgentPolice.Contract.RefreshRoutes(&_AgentPolice.TransactOpts)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0xc6dc030e.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceTransactor) RegisterCredentialUseBlock(opts *bind.TransactOpts, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "registerCredentialUseBlock", sc)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0xc6dc030e.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceSession) RegisterCredentialUseBlock(sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.RegisterCredentialUseBlock(&_AgentPolice.TransactOpts, sc)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0xc6dc030e.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceTransactorSession) RegisterCredentialUseBlock(sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.RegisterCredentialUseBlock(&_AgentPolice.TransactOpts, sc)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AgentPolice *AgentPoliceTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AgentPolice *AgentPoliceSession) Resume() (*types.Transaction, error) {
	return _AgentPolice.Contract.Resume(&_AgentPolice.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AgentPolice *AgentPoliceTransactorSession) Resume() (*types.Transaction, error) {
	return _AgentPolice.Contract.Resume(&_AgentPolice.TransactOpts)
}

// SetAdministrationWindow is a paid mutator transaction binding the contract method 0xbc4c294b.
//
// Solidity: function setAdministrationWindow(uint256 _administrationWindow) returns()
func (_AgentPolice *AgentPoliceTransactor) SetAdministrationWindow(opts *bind.TransactOpts, _administrationWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setAdministrationWindow", _administrationWindow)
}

// SetAdministrationWindow is a paid mutator transaction binding the contract method 0xbc4c294b.
//
// Solidity: function setAdministrationWindow(uint256 _administrationWindow) returns()
func (_AgentPolice *AgentPoliceSession) SetAdministrationWindow(_administrationWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAdministrationWindow(&_AgentPolice.TransactOpts, _administrationWindow)
}

// SetAdministrationWindow is a paid mutator transaction binding the contract method 0xbc4c294b.
//
// Solidity: function setAdministrationWindow(uint256 _administrationWindow) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetAdministrationWindow(_administrationWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAdministrationWindow(&_AgentPolice.TransactOpts, _administrationWindow)
}

// SetAgentDefaultDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xa86beff8.
//
// Solidity: function setAgentDefaultDueToFaultySectorDays(address agent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetAgentDefaultDueToFaultySectorDays(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setAgentDefaultDueToFaultySectorDays", agent)
}

// SetAgentDefaultDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xa86beff8.
//
// Solidity: function setAgentDefaultDueToFaultySectorDays(address agent) returns()
func (_AgentPolice *AgentPoliceSession) SetAgentDefaultDueToFaultySectorDays(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaultDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent)
}

// SetAgentDefaultDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xa86beff8.
//
// Solidity: function setAgentDefaultDueToFaultySectorDays(address agent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetAgentDefaultDueToFaultySectorDays(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaultDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent)
}

// SetAgentDefaulted is a paid mutator transaction binding the contract method 0x8b21505a.
//
// Solidity: function setAgentDefaulted(address agent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetAgentDefaulted(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setAgentDefaulted", agent)
}

// SetAgentDefaulted is a paid mutator transaction binding the contract method 0x8b21505a.
//
// Solidity: function setAgentDefaulted(address agent) returns()
func (_AgentPolice *AgentPoliceSession) SetAgentDefaulted(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaulted(&_AgentPolice.TransactOpts, agent)
}

// SetAgentDefaulted is a paid mutator transaction binding the contract method 0x8b21505a.
//
// Solidity: function setAgentDefaulted(address agent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetAgentDefaulted(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaulted(&_AgentPolice.TransactOpts, agent)
}

// SetDefaultWindow is a paid mutator transaction binding the contract method 0xacec9564.
//
// Solidity: function setDefaultWindow(uint256 _defaultWindow) returns()
func (_AgentPolice *AgentPoliceTransactor) SetDefaultWindow(opts *bind.TransactOpts, _defaultWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setDefaultWindow", _defaultWindow)
}

// SetDefaultWindow is a paid mutator transaction binding the contract method 0xacec9564.
//
// Solidity: function setDefaultWindow(uint256 _defaultWindow) returns()
func (_AgentPolice *AgentPoliceSession) SetDefaultWindow(_defaultWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetDefaultWindow(&_AgentPolice.TransactOpts, _defaultWindow)
}

// SetDefaultWindow is a paid mutator transaction binding the contract method 0xacec9564.
//
// Solidity: function setDefaultWindow(uint256 _defaultWindow) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetDefaultWindow(_defaultWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetDefaultWindow(&_AgentPolice.TransactOpts, _defaultWindow)
}

// SetMaxConsecutiveFaultEpochs is a paid mutator transaction binding the contract method 0xaf10015e.
//
// Solidity: function setMaxConsecutiveFaultEpochs(uint256 _maxConsecutiveFaultEpochs) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxConsecutiveFaultEpochs(opts *bind.TransactOpts, _maxConsecutiveFaultEpochs *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxConsecutiveFaultEpochs", _maxConsecutiveFaultEpochs)
}

// SetMaxConsecutiveFaultEpochs is a paid mutator transaction binding the contract method 0xaf10015e.
//
// Solidity: function setMaxConsecutiveFaultEpochs(uint256 _maxConsecutiveFaultEpochs) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxConsecutiveFaultEpochs(_maxConsecutiveFaultEpochs *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxConsecutiveFaultEpochs(&_AgentPolice.TransactOpts, _maxConsecutiveFaultEpochs)
}

// SetMaxConsecutiveFaultEpochs is a paid mutator transaction binding the contract method 0xaf10015e.
//
// Solidity: function setMaxConsecutiveFaultEpochs(uint256 _maxConsecutiveFaultEpochs) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxConsecutiveFaultEpochs(_maxConsecutiveFaultEpochs *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxConsecutiveFaultEpochs(&_AgentPolice.TransactOpts, _maxConsecutiveFaultEpochs)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxDTE(opts *bind.TransactOpts, _maxDTE *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxDTE", _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxDTE(&_AgentPolice.TransactOpts, _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxDTE(&_AgentPolice.TransactOpts, _maxDTE)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxEpochsOwedTolerance(opts *bind.TransactOpts, _maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxEpochsOwedTolerance", _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxEpochsOwedTolerance(&_AgentPolice.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxEpochsOwedTolerance(&_AgentPolice.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxLTV(opts *bind.TransactOpts, _maxLTV *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxLTV", _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxLTV(&_AgentPolice.TransactOpts, _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxLTV(&_AgentPolice.TransactOpts, _maxLTV)
}

// SetMaxPoolsPerAgent is a paid mutator transaction binding the contract method 0x64af5690.
//
// Solidity: function setMaxPoolsPerAgent(uint256 _maxPoolsPerAgent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxPoolsPerAgent(opts *bind.TransactOpts, _maxPoolsPerAgent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxPoolsPerAgent", _maxPoolsPerAgent)
}

// SetMaxPoolsPerAgent is a paid mutator transaction binding the contract method 0x64af5690.
//
// Solidity: function setMaxPoolsPerAgent(uint256 _maxPoolsPerAgent) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxPoolsPerAgent(_maxPoolsPerAgent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxPoolsPerAgent(&_AgentPolice.TransactOpts, _maxPoolsPerAgent)
}

// SetMaxPoolsPerAgent is a paid mutator transaction binding the contract method 0x64af5690.
//
// Solidity: function setMaxPoolsPerAgent(uint256 _maxPoolsPerAgent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxPoolsPerAgent(_maxPoolsPerAgent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxPoolsPerAgent(&_AgentPolice.TransactOpts, _maxPoolsPerAgent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetSectorFaultyTolerancePercent(opts *bind.TransactOpts, _sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setSectorFaultyTolerancePercent", _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPolice *AgentPoliceSession) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetSectorFaultyTolerancePercent(&_AgentPolice.TransactOpts, _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetSectorFaultyTolerancePercent(&_AgentPolice.TransactOpts, _sectorFaultyTolerancePercent)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPolice *AgentPoliceTransactor) TransferOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "transferOperator", newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPolice *AgentPoliceSession) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOperator(&_AgentPolice.TransactOpts, newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPolice *AgentPoliceTransactorSession) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOperator(&_AgentPolice.TransactOpts, newOperator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPolice *AgentPoliceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPolice *AgentPoliceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOwnership(&_AgentPolice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPolice *AgentPoliceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOwnership(&_AgentPolice.TransactOpts, newOwner)
}

// AgentPoliceDefaultedIterator is returned from FilterDefaulted and is used to iterate over the raw logs and unpacked data for Defaulted events raised by the AgentPolice contract.
type AgentPoliceDefaultedIterator struct {
	Event *AgentPoliceDefaulted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceDefaultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceDefaulted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceDefaulted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceDefaultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceDefaultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceDefaulted represents a Defaulted event raised by the AgentPolice contract.
type AgentPoliceDefaulted struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaulted is a free log retrieval operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPolice *AgentPoliceFilterer) FilterDefaulted(opts *bind.FilterOpts) (*AgentPoliceDefaultedIterator, error) {

	logs, sub, err := _AgentPolice.contract.FilterLogs(opts, "Defaulted")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceDefaultedIterator{contract: _AgentPolice.contract, event: "Defaulted", logs: logs, sub: sub}, nil
}

// WatchDefaulted is a free log subscription operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPolice *AgentPoliceFilterer) WatchDefaulted(opts *bind.WatchOpts, sink chan<- *AgentPoliceDefaulted) (event.Subscription, error) {

	logs, sub, err := _AgentPolice.contract.WatchLogs(opts, "Defaulted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceDefaulted)
				if err := _AgentPolice.contract.UnpackLog(event, "Defaulted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaulted is a log parse operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPolice *AgentPoliceFilterer) ParseDefaulted(log types.Log) (*AgentPoliceDefaulted, error) {
	event := new(AgentPoliceDefaulted)
	if err := _AgentPolice.contract.UnpackLog(event, "Defaulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AgentPolice contract.
type AgentPoliceEIP712DomainChangedIterator struct {
	Event *AgentPoliceEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceEIP712DomainChanged represents a EIP712DomainChanged event raised by the AgentPolice contract.
type AgentPoliceEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentPolice *AgentPoliceFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AgentPoliceEIP712DomainChangedIterator, error) {

	logs, sub, err := _AgentPolice.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceEIP712DomainChangedIterator{contract: _AgentPolice.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentPolice *AgentPoliceFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AgentPoliceEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AgentPolice.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceEIP712DomainChanged)
				if err := _AgentPolice.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentPolice *AgentPoliceFilterer) ParseEIP712DomainChanged(log types.Log) (*AgentPoliceEIP712DomainChanged, error) {
	event := new(AgentPoliceEIP712DomainChanged)
	if err := _AgentPolice.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceFaultySectorsIterator is returned from FilterFaultySectors and is used to iterate over the raw logs and unpacked data for FaultySectors events raised by the AgentPolice contract.
type AgentPoliceFaultySectorsIterator struct {
	Event *AgentPoliceFaultySectors // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceFaultySectorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceFaultySectors)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceFaultySectors)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceFaultySectorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceFaultySectorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceFaultySectors represents a FaultySectors event raised by the AgentPolice contract.
type AgentPoliceFaultySectors struct {
	AgentID    common.Address
	FaultEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFaultySectors is a free log retrieval operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPolice *AgentPoliceFilterer) FilterFaultySectors(opts *bind.FilterOpts, agentID []common.Address) (*AgentPoliceFaultySectorsIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPolice.contract.FilterLogs(opts, "FaultySectors", agentIDRule)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceFaultySectorsIterator{contract: _AgentPolice.contract, event: "FaultySectors", logs: logs, sub: sub}, nil
}

// WatchFaultySectors is a free log subscription operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPolice *AgentPoliceFilterer) WatchFaultySectors(opts *bind.WatchOpts, sink chan<- *AgentPoliceFaultySectors, agentID []common.Address) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPolice.contract.WatchLogs(opts, "FaultySectors", agentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceFaultySectors)
				if err := _AgentPolice.contract.UnpackLog(event, "FaultySectors", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFaultySectors is a log parse operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPolice *AgentPoliceFilterer) ParseFaultySectors(log types.Log) (*AgentPoliceFaultySectors, error) {
	event := new(AgentPoliceFaultySectors)
	if err := _AgentPolice.contract.UnpackLog(event, "FaultySectors", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceOnAdministrationIterator is returned from FilterOnAdministration and is used to iterate over the raw logs and unpacked data for OnAdministration events raised by the AgentPolice contract.
type AgentPoliceOnAdministrationIterator struct {
	Event *AgentPoliceOnAdministration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceOnAdministrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceOnAdministration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceOnAdministration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceOnAdministrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceOnAdministrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceOnAdministration represents a OnAdministration event raised by the AgentPolice contract.
type AgentPoliceOnAdministration struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnAdministration is a free log retrieval operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPolice *AgentPoliceFilterer) FilterOnAdministration(opts *bind.FilterOpts) (*AgentPoliceOnAdministrationIterator, error) {

	logs, sub, err := _AgentPolice.contract.FilterLogs(opts, "OnAdministration")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceOnAdministrationIterator{contract: _AgentPolice.contract, event: "OnAdministration", logs: logs, sub: sub}, nil
}

// WatchOnAdministration is a free log subscription operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPolice *AgentPoliceFilterer) WatchOnAdministration(opts *bind.WatchOpts, sink chan<- *AgentPoliceOnAdministration) (event.Subscription, error) {

	logs, sub, err := _AgentPolice.contract.WatchLogs(opts, "OnAdministration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceOnAdministration)
				if err := _AgentPolice.contract.UnpackLog(event, "OnAdministration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnAdministration is a log parse operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPolice *AgentPoliceFilterer) ParseOnAdministration(log types.Log) (*AgentPoliceOnAdministration, error) {
	event := new(AgentPoliceOnAdministration)
	if err := _AgentPolice.contract.UnpackLog(event, "OnAdministration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
