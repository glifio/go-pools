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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_defaultWindow\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"contractIPoolRegistry\",\"name\":\"_poolRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentStateRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCredential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"Defaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agentID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"faultEpoch\",\"type\":\"uint256\"}],\"name\":\"FaultySectors\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"OnAdministration\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_VERIFIABLE_CREDENTIAL_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrationWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"agentApproved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"agentLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"accountId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"authMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmAdministration\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmEquity\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"credentialUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"deriveStructHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distributeLiquidatedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"isValidCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgent[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"name\":\"markAsFaulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxConsecutiveFaultEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEpochsOwedTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoolsPerAgent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidator\",\"type\":\"uint64\"}],\"name\":\"prepareMinerForLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministrationDueToFaultySectorDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"registerCredentialUseBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sectorFaultyTolerancePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_administrationWindow\",\"type\":\"uint256\"}],\"name\":\"setAdministrationWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setAgentDefaultDueToFaultySectorDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setAgentDefaulted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultWindow\",\"type\":\"uint256\"}],\"name\":\"setDefaultWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxConsecutiveFaultEpochs\",\"type\":\"uint256\"}],\"name\":\"setMaxConsecutiveFaultEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDTE\",\"type\":\"uint256\"}],\"name\":\"setMaxDTE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxEpochsOwedTolerance\",\"type\":\"uint256\"}],\"name\":\"setMaxEpochsOwedTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLTV\",\"type\":\"uint256\"}],\"name\":\"setMaxLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPoolsPerAgent\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolsPerAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sectorFaultyTolerancePercent\",\"type\":\"uint256\"}],\"name\":\"setSectorFaultyTolerancePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"transferOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blssig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"validateCred\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610180604052670de0b6b3a7640000600b819055600c5562000025601e610e10620004ed565b6200003290601862000510565b6200003f90600362000510565b600d5562000051601e610e10620004ed565b6200005e90601862000510565b6200006b90600162000510565b600e5566038d7ea4c68000600f556010805460ff191690553480156200009057600080fd5b50604051620058ce380380620058ce833981016040819052620000b3916200061c565b8484818a8a878282620000d6600083620002af60201b62001ea41790919060201c565b61012052620000f3816001620002af602090811b62001ea417901c565b61014052815160208084019190912060e052815190820120610100524660a0526200018160e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c0526001600160a01b0390811661016052620001b5925083169050620002ff602090811b62001ed517901c565b90506001600160a01b038116620001df57604051635435b28960e11b815260040160405180910390fd5b620001ea816200034b565b506200020a816001600160a01b0316620002ff60201b62001ed51760201c565b90506001600160a01b0381166200023457604051635435b28960e11b815260040160405180910390fd5b6200023f816200039e565b5050600886905562000255601e610e10620004ed565b6200026290601862000510565b6200026f90600762000510565b600955600a8055600680546001600160a01b039384166001600160a01b0319918216179091556007805492909316911617905550620008b2945050505050565b6000602083511015620002cf57620002c783620003f1565b9050620002f9565b82620002e6836200043d60201b62001f1c1760201c565b90620002f390826200078c565b5060ff90505b92915050565b600080806200030e8462000440565b91509150816200032057509192915050565b6000806200032e8362000473565b915091508162000342575093949350505050565b95945050505050565b600380546001600160a01b03191690556200037b6001600160a01b038216620002ff602090811b62001ed517901c565b600280546001600160a01b0319166001600160a01b039290921691909117905550565b600580546001600160a01b0319169055620003ce6001600160a01b038216620002ff602090811b62001ed517901c565b600480546001600160a01b0319166001600160a01b039290921691909117905550565b600080829050601f8151111562000428578260405163305a27a960e01b81526004016200041f919062000858565b60405180910390fd5b805162000435826200088d565b179392505050565b90565b600080600160401b600160a01b03831660ff60981b81036200046d57600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620004c75760009250600091505b50811580620004d757503d601614155b15620004e857506000928392509050565b915091565b6000826200050b57634e487b7160e01b600052601260045260246000fd5b500490565b8082028115828204841417620002f957634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60005b83811015620005695781810151838201526020016200054f565b50506000910152565b600082601f8301126200058457600080fd5b81516001600160401b0380821115620005a157620005a162000536565b604051601f8301601f19908116603f01168101908282118183101715620005cc57620005cc62000536565b81604052838152866020858801011115620005e657600080fd5b620005f98460208301602089016200054c565b9695505050505050565b6001600160a01b03811681146200061957600080fd5b50565b600080600080600080600080610100898b0312156200063a57600080fd5b88516001600160401b03808211156200065257600080fd5b620006608c838d0162000572565b995060208b01519150808211156200067757600080fd5b50620006868b828c0162000572565b975050604089015195506060890151620006a08162000603565b60808a0151909550620006b38162000603565b60a08a0151909450620006c68162000603565b60c08a0151909350620006d98162000603565b60e08a0151909250620006ec8162000603565b809150509295985092959890939650565b600181811c908216806200071257607f821691505b6020821081036200073357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200078757600081815260208120601f850160051c81016020861015620007625750805b601f850160051c820191505b8181101562000783578281556001016200076e565b5050505b505050565b81516001600160401b03811115620007a857620007a862000536565b620007c081620007b98454620006fd565b8462000739565b602080601f831160018114620007f85760008415620007df5750858301515b600019600386901b1c1916600185901b17855562000783565b600085815260208120601f198616915b82811015620008295788860151825594840194600190910190840162000808565b5085821015620008485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020815260008251806020840152620008798160408501602087016200054c565b601f01601f19169190910160400192915050565b80516020808301519190811015620007335760001960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051614f776200095760003960008181610c5501528181610cc301528181610fb20152818161100d0152818161114e01528181611b1c015281816120c90152818161236e01528181612b4b0152612e5a015260006117970152600061176d015260006138d5015260006138ad01526000613808015260006138320152600061385c0152614f776000f3fe608060405234801561001057600080fd5b506004361061032b5760003560e01c8063680af724116101b2578063ac76eb40116100f9578063cada34b1116100a2578063f2fde38b1161007c578063f2fde38b1461068a578063f384bd051461069d578063fc3e5174146106a6578063fcad4448146106b957600080fd5b8063cada34b114610644578063cb70f56314610657578063e30c39781461066a57600080fd5b8063af10015e116100d3578063af10015e1461060b578063bc4c294b1461061e578063c6dc030e1461063157600080fd5b8063ac76eb40146105cf578063ac7e534e146105d8578063acec9564146105f857600080fd5b80638b21505a1161015b5780639fc2f56b116101355780639fc2f56b146105a0578063a2f3c210146105a9578063a86beff8146105bc57600080fd5b80638b21505a1461055a5780638da5cb5b1461056d5780639b83b8cd1461058d57600080fd5b806379ba50971161018c57806379ba50971461052f5780638456cb591461053757806384b0196e1461053f57600080fd5b8063680af7241461050057806369cdb53c1461051357806378ee44fd1461052657600080fd5b806337279488116102765780635654657a1161021f5780635cf6862f116101f95780635cf6862f146104c75780635e7981de146104da57806364af5690146104ed57600080fd5b80635654657a14610462578063570ca735146104755780635c975abb146104ba57600080fd5b80634107e644116102505780634107e644146104335780635169214914610446578063561463801461045957600080fd5b8063372794881461040e578063395bbfe91461041757806340444c7c1461042057600080fd5b806327ddb79d116102d85780632c6de7c9116102b25780632c6de7c9146103e05780632e446475146103f3578063338891eb1461040657600080fd5b806327ddb79d146103b257806329605e77146103c55780632b85012b146103d857600080fd5b80631387b4ed116103095780631387b4ed146103605780631f2218de1461037c578063260e63c61461039f57600080fd5b8063046f7da21461033057806308a0c3751461033a578063117573c51461034d575b600080fd5b6103386106c1565b005b610338610348366004613ede565b6106f3565b61033861035b366004613ede565b610700565b610369600e5481565b6040519081526020015b60405180910390f35b61038f61038a366004613ef7565b61070d565b6040519015158152602001610373565b6103386103ad366004613ef7565b610736565b6103386103c0366004613f33565b610742565b6103386103d3366004613fda565b610847565b610369610896565b6103386103ee366004613ff7565b6108d8565b6103386104013660046141a1565b610a9b565b610338610c50565b610369600f5481565b610369600d5481565b61036961042e3660046142e9565b610d2e565b610338610441366004613ede565b610e81565b61038f610454366004613ede565b610edf565b610369600a5481565b610338610470366004613ef7565b610ffd565b6004546104959073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610373565b60105461038f9060ff1681565b6103386104d5366004613ef7565b611147565b6103386104e836600461431e565b61122a565b6103386104fb366004613ede565b6114df565b61033861050e366004613ede565b6114ec565b6103386105213660046143e9565b6114f9565b61036960085481565b6103386116ce565b61033861172a565b61054761175f565b60405161037397969594939291906144a4565b610338610568366004613fda565b611803565b6002546104959073ffffffffffffffffffffffffffffffffffffffff1681565b61033861059b366004614563565b611943565b610369600b5481565b6103696105b73660046142e9565b6119e1565b6103386105ca366004613fda565b6119fa565b61036960095481565b6005546104959073ffffffffffffffffffffffffffffffffffffffff1681565b610338610606366004613ede565b611af0565b610338610619366004613ede565b611afd565b61033861062c366004613ede565b611b0a565b61033861063f3660046145a8565b611b17565b6103386106523660046145dd565b611c16565b610338610665366004613ff7565b611c6d565b6003546104959073ffffffffffffffffffffffffffffffffffffffff1681565b610338610698366004613fda565b611db5565b610369600c5481565b6104956106b43660046145a8565b611e24565b610338611e4a565b6106c9611f1f565b601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6106fb611f1f565b600c55565b610708611f1f565b600b55565b60008060118161071f6105b78661463b565b815260200190815260200160002054119050919050565b61073f81611f70565b50565b61074a612220565b6000805b828110156108415783838281811061076857610768614647565b905060200201602081019061077d9190613fda565b91508173ffffffffffffffffffffffffffffffffffffffff1663a71804496040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156107c757600080fd5b505af11580156107db573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff167fae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd64360405161082791815260200190565b60405180910390a280610839816146a5565b91505061074e565b50505050565b61084f611f1f565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6040518060c0016040528060968152602001614eac609691396040516020016108bf91906146dd565b6040516020818303038152906040528051906020012081565b6108e0611f1f565b8160095461095c8273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610932573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095691906146f9565b82612297565b610992576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff16639e4d74cc6109cd8573ffffffffffffffffffffffffffffffffffffffff16611ed5565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401600060405180830381600087803b158015610a3357600080fd5b505af1158015610a47573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff871681527fcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c29250602001905060405180910390a150505050565b6000610aa760026123d0565b610ab19082614712565b9050610abc8361241e565b610ac69082614712565b9050610ad18261241e565b610adb9082614712565b90506000610ae882612436565b9050610af5816002612457565b610aff8185612463565b610b098184612463565b6000610b148261247c565b9050639d8b067860006051816060610b2f8b8684868a6124f7565b9092509050818015610b425750805160a0145b8015610ba7575080607f81518110610b5c57610b5c614647565b6020910101517fff00000000000000000000000000000000000000000000000000000000000000167f0100000000000000000000000000000000000000000000000000000000000000145b8015610c0c575080608081518110610bc157610bc1614647565b6020910101517fff00000000000000000000000000000000000000000000000000000000000000167ff500000000000000000000000000000000000000000000000000000000000000145b15610c1e575050505050505050505050565b6040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c797f0000000000000000000000000000000000000000000000000000000000000000612518565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055610ce77f00000000000000000000000000000000000000000000000000000000000000006125fc565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006040518060c0016040528060968152602001614eac60969139604051602001610d5991906146dd565b60405160208183030381529060405280519060200120826000015183602001518460400151856060015186608001518760a001518860c001518960e00151604051602001610da791906146dd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083019a909a5273ffffffffffffffffffffffffffffffffffffffff909816978101979097526060870195909552608086019390935260a085019190915260c08401527fffffffff000000000000000000000000000000000000000000000000000000001660e083015267ffffffffffffffff1661010082015261012081019190915261014001604051602081830303815290604052805190602001209050919050565b610e89611f1f565b610e96601e610e10614754565b610ea1906018614768565b811115610eda576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e55565b6006546040517f9a014e2400000000000000000000000000000000000000000000000000000000815260048101839052600091829173ffffffffffffffffffffffffffffffffffffffff90911690639a014e2490602401600060405180830381865afa158015610f53573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610f99919081019061477f565b90508051600003610fad5750600092915050565b610ff27f00000000000000000000000000000000000000000000000000000000000000008483600081518110610fe557610fe5614647565b60200260200101516126a3565b606001519392505050565b61100681611f70565b60006110317f000000000000000000000000000000000000000000000000000000000000000061276e565b90506000611048826110428561463b565b90612815565b90508060000361105757505050565b600061106c836110668661463b565b906128b5565b90508181116110a7576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b546110be6110b78484614825565b849061290d565b11156110f6576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c5461110f6110b7856111098861463b565b90612922565b1115610841576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006111727f000000000000000000000000000000000000000000000000000000000000000061276e565b9050600f546111a761118e83856111889061463b565b9061297a565b6111a18461119b8761463b565b906129d2565b9061290d565b11156111df576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111ef8260200135600954612297565b15611226576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b611232611f1f565b60008273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561127f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a391906146f9565b90506112ae81610edf565b156112e5576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af1158015611362573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113869190614848565b5060006113938284612a2a565b9050600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8573ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561141e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114429190614863565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af11580156114b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114d89190614848565b5050505050565b6114e7611f1f565b600a55565b6114f4611f1f565b600f55565b608081015151156115ec57600061151382600001516119e1565b60405160200161152591815260200190565b6040516020818303038152906040529050611547610402836080015183610a9b565b815160200151841415806115855750815160a001517fffffffff00000000000000000000000000000000000000000000000000000000848116911614155b806115a7575081516040015143108015906115a557508151606001514311155b155b806115b55750815160200151155b15610841576040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006115f782611e24565b82515190915073ffffffffffffffffffffffffffffffffffffffff808316911614158061162a575061162881612dbe565b155b8061163a57508151602001518414155b806115855750815160a001517fffffffff0000000000000000000000000000000000000000000000000000000084811691161415806115a7575081516040015143108015906115a557508151606001514311806115b55750815160200151610841576040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035473ffffffffffffffffffffffffffffffffffffffff16331461171f576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61172833612ee4565b565b611732611f1f565b601080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b6000606080828080836117927f000000000000000000000000000000000000000000000000000000000000000083612f73565b6117bd7f00000000000000000000000000000000000000000000000000000000000000006001612f73565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b61180b611f1f565b8060085461185d8273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610932573d6000803e3d6000fd5b611893576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663615d92116040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156118db57600080fd5b505af11580156118ef573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff861681527f0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1925060200190505b60405180910390a1505050565b61194b611f1f565b6040517ff794a92300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff80841660048301528216602482015273ffffffffffffffffffffffffffffffffffffffff84169063f794a92390604401600060405180830381600087803b1580156119c457600080fd5b505af11580156119d8573d6000803e3d6000fd5b50505050505050565b60006119f46119ef83610d2e565b613017565b92915050565b611a02611f1f565b80611a0c8161305f565b611a42576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663615d92116040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611a8a57600080fd5b505af1158015611a9e573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff851681527f0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b19250602001905060405180910390a15050565b611af8611f1f565b600855565b611b05611f1f565b600d55565b611b12611f1f565b600955565b611b417f0000000000000000000000000000000000000000000000000000000000000000336130fb565b8060000151602001513373ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bb991906146f9565b14611bf0576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4360116000611c0284600001516119e1565b815260208101919091526040016000205550565b611c24838361052184614880565b611c3161038a828061488c565b15611c68576040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b611c75611f1f565b81611c7f8161305f565b611cb5576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16639e4d74cc611cf08473ffffffffffffffffffffffffffffffffffffffff16611ed5565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401600060405180830381600087803b158015611d5657600080fd5b505af1158015611d6a573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff861681527fcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c292506020019050611936565b611dbd611f1f565b611ddc8173ffffffffffffffffffffffffffffffffffffffff16611ed5565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60006119f4611e3683600001516119e1565b8360200151846040015185606001516131cc565b60055473ffffffffffffffffffffffffffffffffffffffff163314611e9b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611728336131f4565b6000602083511015611ec057611eb983613283565b90506119f4565b81611ecb8482614961565b5060ff9392505050565b6000806000611ee3846132da565b9150915081611ef457509192915050565b600080611f0083613327565b9150915081611f13575093949350505050565b95945050505050565b90565b60025473ffffffffffffffffffffffffffffffffffffffff163314611728576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006546040517f9a014e240000000000000000000000000000000000000000000000000000000081526020830135600482015260009173ffffffffffffffffffffffffffffffffffffffff1690639a014e2490602401600060405180830381865afa158015611fe3573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612029919081019061477f565b9050600a5481511115612068576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8151811015611c6857600082828151811061208857612088614647565b6020026020010151905060006120c0600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836133ab565b905060006120f37f00000000000000000000000000000000000000000000000000000000000000008760200135856126a3565b6040517f9741fbfa00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff831690639741fbfa9061214a9084908a90600401614b2f565b602060405180830381865afa158015612167573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061218b9190614848565b6121c1576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b43600e5482604001516121d49190614712565b101561220c576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505080612219906146a5565b905061206b565b60045473ffffffffffffffffffffffffffffffffffffffff163314801590612260575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15611728576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006546040517f9a014e2400000000000000000000000000000000000000000000000000000000815260048101849052600091829173ffffffffffffffffffffffffffffffffffffffff90911690639a014e2490602401600060405180830381865afa15801561230b573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612351919081019061477f565b905060005b81518110156123c5576123698443614825565b6123a07f000000000000000000000000000000000000000000000000000000000000000087858581518110610fe557610fe5614647565b6040015110156123b5576001925050506119f4565b6123be816146a5565b9050612356565b506000949350505050565b6000601782116123e257506001919050565b60ff82116123f257506002919050565b61ffff821161240357506003919050565b63ffffffff821161241657506005919050565b506009919050565b6000815161242c83516123d0565b6119f49190614712565b61243e613ea9565b805161244a90836134e4565b5060006020820152919050565b6112268260048361355b565b6124708260028351613562565b8151611c689082613683565b606081602001516000146124f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f496e76616c69642043424f52000000000000000000000000000000000000000060448201526064015b60405180910390fd5b50515190565b6000606061250a600088888888886136a4565b915091509550959350505050565b604080518082018252601481527f524f555445525f504f4f4c5f5245474953545259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f390c2e8200000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024015b602060405180830381865afa1580156125d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119f49190614863565b604080518082018252601181527f524f555445525f5746494c5f544f4b454e000000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527faee0a13500000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016125bb565b6126d060405180608001604052806000815260200160008152602001600081526020016000151581525090565b6040517f6361f6de000000000000000000000000000000000000000000000000000000008152600481018490526024810183905273ffffffffffffffffffffffffffffffffffffffff851690636361f6de90604401608060405180830381865afa158015612742573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127669190614c2d565b949350505050565b604080518082018252601281527f524f555445525f435245445f5041525345520000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527fd26df3b700000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016125bb565b60e08201516040517f17bbd29a00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff8416916317bbd29a9161286d91600401614c97565b602060405180830381865afa15801561288a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ae91906146f9565b9392505050565b60e08201516040517f6bb0d0cc00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff841691636bb0d0cc9161286d91600401614c97565b60006128ae83670de0b6b3a764000084613790565b60e08201516040517f1b2b5fad00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff841691631b2b5fad9161286d91600401614c97565b60e08201516040517f402ed8cb00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff84169163402ed8cb9161286d91600401614c97565b60e08201516040517f07124c0600000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff8416916307124c069161286d91600401614c97565b6006546040517f9a014e240000000000000000000000000000000000000000000000000000000081526004810184905260009182918291829182918291829173ffffffffffffffffffffffffffffffffffffffff1690639a014e2490602401600060405180830381865afa158015612aa6573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612aec919081019061477f565b80519091506000808267ffffffffffffffff811115612b0d57612b0d614048565b604051908082528060200260200182016040528015612b36578160200160208202803683370190505b509050600097505b82881015612bbd57612b7d7f00000000000000000000000000000000000000000000000000000000000000008d868b81518110610fe557610fe5614647565b60200151945084818981518110612b9657612b96614647565b6020908102919091010152612bab858a614712565b9850612bb6886146a5565b9750612b3e565b600097505b82881015612daf57838881518110612bdc57612bdc614647565b60200260200101519650888b828a81518110612bfa57612bfa614647565b6020026020010151612c0c9190614768565b612c169190614754565b600654909650600090612c3f9073ffffffffffffffffffffffffffffffffffffffff16896133ab565b6007546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8084166004830152602482018b905292935091169063095ea7b3906044016020604051808303816000875af1158015612cba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cde9190614848565b506040517fc564f772000000000000000000000000000000000000000000000000000000008152600481018e90526024810188905273ffffffffffffffffffffffffffffffffffffffff82169063c564f772906044016020604051808303816000875af1158015612d53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d7791906146f9565b9250828711612d87576000612d91565b612d918388614825565b612d9b908c614712565b9a505087612da8906146a5565b9750612bc2565b50505050505050505092915050565b604080518082018252601081527f524f555445525f56435f49535355455200000000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f1603a47a00000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff838116917f000000000000000000000000000000000000000000000000000000000000000090911690631a6e649e90602401602060405180830381865afa158015612ea3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ec79190614863565b73ffffffffffffffffffffffffffffffffffffffff161492915050565b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055612f2b73ffffffffffffffffffffffffffffffffffffffff8216611ed5565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b606060ff8314612f8657611eb9836137af565b818054612f92906148c0565b80601f0160208091040260200160405190810160405280929190818152602001828054612fbe906148c0565b801561300b5780601f10612fe05761010080835404028352916020019161300b565b820191906000526020600020905b815481529060010190602001808311612fee57829003601f168201915b505050505090506119f4565b60006119f46130246137ee565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000808273ffffffffffffffffffffffffffffffffffffffff16638903f7f06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130d191906146f9565b9050806000036130e45750600092915050565b600d546130f19082614712565b4310159392505050565b6131048261392b565b6040517f1ffbb06400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529190911690631ffbb06490602401602060405180830381865afa158015613172573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131969190614848565b611226576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060006131dd878787876139d2565b915091506131ea81613ac1565b5095945050505050565b600580547fffffffffffffffffffffffff000000000000000000000000000000000000000016905561323b73ffffffffffffffffffffffffffffffffffffffff8216611ed5565b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b600080829050601f815111156132c757826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016124e89190614c97565b80516132d282614caa565b179392505050565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103613321576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a81146133875760009250600091505b5081158061339657503d601614155b156133a657506000928392509050565b915091565b60008273ffffffffffffffffffffffffffffffffffffffff1663efde4e646040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061341c91906146f9565b821115613455576040517f8f5d631800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f41d1de970000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8416906341d1de9790602401602060405180830381865afa1580156134c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ae9190614863565b604080518082019091526060815260006020820152613504602083614cec565b1561352c57613514602083614cec565b61351f906020614825565b6135299083614712565b91505b60208084018390526040518085526000815290818401018181101561355057600080fd5b604052509192915050565b611c688383835b60178167ffffffffffffffff16116135895782516108419060e0600585901b168317613c74565b60ff8167ffffffffffffffff16116135cb5782516135b2906018611fe0600586901b1617613c74565b5082516108419067ffffffffffffffff83166001613cdd565b61ffff8167ffffffffffffffff161161360e5782516135f5906019611fe0600586901b1617613c74565b5082516108419067ffffffffffffffff83166002613cdd565b63ffffffff8167ffffffffffffffff161161365357825161363a90601a611fe0600586901b1617613c74565b5082516108419067ffffffffffffffff83166004613cdd565b825161366a90601b611fe0600586901b1617613c74565b5082516108419067ffffffffffffffff83166008613cdd565b6040805180820190915260608152600060208201526128ae83838451613d62565b600060606000886136b65760006136b9565b60015b90508815806136c6575085155b6136cf57600080fd5b600087878388888d6040516020016136ec96959493929190614d00565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052915073fe000000000000000000000000000000000000059061373e9083906146dd565b600060405180830381855af49150503d8060008114613779576040519150601f19603f3d011682016040523d82523d6000602084013e61377e565b606091505b50935093505050965096945050505050565b8282028115158415858304851417166137a857600080fd5b0492915050565b606060006137bc83613e51565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b60003073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561385457507f000000000000000000000000000000000000000000000000000000000000000046145b1561387e57507f000000000000000000000000000000000000000000000000000000000000000090565b613926604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b905090565b604080518082018252601481527f524f555445525f4147454e545f464143544f5259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f29f616da00000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016125bb565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613a095750600090506003613ab8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613a5d573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613ab157600060019250925050613ab8565b9150600090505b94509492505050565b6000816004811115613ad557613ad5614d50565b03613add5750565b6001816004811115613af157613af1614d50565b03613b58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016124e8565b6002816004811115613b6c57613b6c614d50565b03613bd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016124e8565b6003816004811115613be757613be7614d50565b0361073f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016124e8565b6040805180820190915260608152600060208201528251516000613c99826001614712565b905084602001518210613cba57613cba85613cb5836002614768565b613e92565b8451602083820101858153508051821115613cd3578181525b5093949350505050565b6040805180820190915260608152600060208201528351516000613d018285614712565b90508560200151811115613d1e57613d1e86613cb5836002614768565b60006001613d2e86610100614e9f565b613d389190614825565b90508651828101878319825116178152508051831115613d56578281525b50959695505050505050565b6040805180820190915260608152600060208201528251821115613d8557600080fd5b8351516000613d948483614712565b90508560200151811115613db157613db186613cb5836002614768565b855180518382016020019160009180851115613dcb578482525b505050602086015b60208610613e0b5780518252613dea602083614712565b9150613df7602082614712565b9050613e04602087614825565b9550613dd3565b5181517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60208890036101000a0190811690199190911617905250849150509392505050565b600060ff8216601f8111156119f4576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151613e9e83836134e4565b506108418382613683565b6040518060400160405280613ed1604051806040016040528060608152602001600081525090565b8152602001600081525090565b600060208284031215613ef057600080fd5b5035919050565b600060208284031215613f0957600080fd5b813567ffffffffffffffff811115613f2057600080fd5b820161010081850312156128ae57600080fd5b60008060208385031215613f4657600080fd5b823567ffffffffffffffff80821115613f5e57600080fd5b818501915085601f830112613f7257600080fd5b813581811115613f8157600080fd5b8660208260051b8501011115613f9657600080fd5b60209290920196919550909350505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461073f57600080fd5b8035613fd581613fa8565b919050565b600060208284031215613fec57600080fd5b81356128ae81613fa8565b6000806040838503121561400a57600080fd5b823561401581613fa8565b9150602083013561402581613fa8565b809150509250929050565b803567ffffffffffffffff81168114613fd557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff8111828210171561409b5761409b614048565b60405290565b60405160a0810167ffffffffffffffff8111828210171561409b5761409b614048565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561410b5761410b614048565b604052919050565b600082601f83011261412457600080fd5b813567ffffffffffffffff81111561413e5761413e614048565b61416f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016140c4565b81815284602083860101111561418457600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000606084860312156141b657600080fd5b6141bf84614030565b9250602084013567ffffffffffffffff808211156141dc57600080fd5b6141e887838801614113565b935060408601359150808211156141fe57600080fd5b5061420b86828701614113565b9150509250925092565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114613fd557600080fd5b6000610100828403121561425857600080fd5b614260614077565b905061426b82613fca565b8152602082013560208201526040820135604082015260608201356060820152608082013560808201526142a160a08301614215565b60a08201526142b260c08301614030565b60c082015260e082013567ffffffffffffffff8111156142d157600080fd5b6142dd84828501614113565b60e08301525092915050565b6000602082840312156142fb57600080fd5b813567ffffffffffffffff81111561431257600080fd5b61276684828501614245565b6000806040838503121561433157600080fd5b823561433c81613fa8565b946020939093013593505050565b600060a0828403121561435c57600080fd5b6143646140a1565b9050813567ffffffffffffffff8082111561437e57600080fd5b61438a85838601614245565b83526020840135915060ff821682146143a257600080fd5b816020840152604084013560408401526060840135606084015260808401359150808211156143d057600080fd5b506143dd84828501614113565b60808301525092915050565b6000806000606084860312156143fe57600080fd5b8335925061440e60208501614215565b9150604084013567ffffffffffffffff81111561442a57600080fd5b61420b8682870161434a565b60005b83811015614451578181015183820152602001614439565b50506000910152565b60008151808452614472816020860160208601614436565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b7fff00000000000000000000000000000000000000000000000000000000000000881681526000602060e0818401526144e060e084018a61445a565b83810360408501526144f2818a61445a565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825283870192509083019060005b8181101561455157835183529284019291840191600101614535565b50909c9b505050505050505050505050565b60008060006060848603121561457857600080fd5b833561458381613fa8565b925061459160208501614030565b915061459f60408501614030565b90509250925092565b6000602082840312156145ba57600080fd5b813567ffffffffffffffff8111156145d157600080fd5b6127668482850161434a565b6000806000606084860312156145f257600080fd5b8335925061460260208501614215565b9150604084013567ffffffffffffffff81111561461e57600080fd5b840160a0818703121561463057600080fd5b809150509250925092565b60006119f43683614245565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036146d6576146d6614676565b5060010190565b600082516146ef818460208701614436565b9190910192915050565b60006020828403121561470b57600080fd5b5051919050565b808201808211156119f4576119f4614676565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261476357614763614725565b500490565b80820281158282048414176119f4576119f4614676565b6000602080838503121561479257600080fd5b825167ffffffffffffffff808211156147aa57600080fd5b818501915085601f8301126147be57600080fd5b8151818111156147d0576147d0614048565b8060051b91506147e18483016140c4565b81815291830184019184810190888411156147fb57600080fd5b938501935b8385101561481957845182529385019390850190614800565b98975050505050505050565b818103818111156119f4576119f4614676565b80518015158114613fd557600080fd5b60006020828403121561485a57600080fd5b6128ae82614838565b60006020828403121561487557600080fd5b81516128ae81613fa8565b60006119f4368361434a565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018336030181126146ef57600080fd5b600181811c908216806148d457607f821691505b60208210810361490d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115611c6857600081815260208120601f850160051c8101602086101561493a5750805b601f850160051c820191505b8181101561495957828155600101614946565b505050505050565b815167ffffffffffffffff81111561497b5761497b614048565b61498f8161498984546148c0565b84614913565b602080601f8311600181146149e257600084156149ac5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555614959565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015614a2f57888601518255948401946001909101908401614a10565b5085821015614a6b57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614ab057600080fd5b830160208101925035905067ffffffffffffffff811115614ad057600080fd5b803603821315614adf57600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b82518152602083015160208201526040830151604082015260608301511515606082015260a0608082015260008235614b6781613fa8565b73ffffffffffffffffffffffffffffffffffffffff811660a084015250602083013560c0830152604083013560e08301526101006060840135818401526080840135610120840152614bbb60a08501614215565b7fffffffff0000000000000000000000000000000000000000000000000000000016610140840152614bef60c08501614030565b67ffffffffffffffff16610160840152614c0c60e0850185614a7b565b82610180860152614c226101a086018284614ae6565b979650505050505050565b600060808284031215614c3f57600080fd5b6040516080810181811067ffffffffffffffff82111715614c6257614c62614048565b8060405250825181526020830151602082015260408301516040820152614c8b60608401614838565b60608201529392505050565b6020815260006128ae602083018461445a565b8051602080830151919081101561490d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b600082614cfb57614cfb614725565b500690565b600067ffffffffffffffff80891683528760208401528087166040840152808616606084015260c06080840152614d3a60c084018661445a565b915080841660a084015250979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600181815b80851115614dd857817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115614dbe57614dbe614676565b80851615614dcb57918102915b93841c9390800290614d84565b509250929050565b600082614def575060016119f4565b81614dfc575060006119f4565b8160018114614e125760028114614e1c57614e38565b60019150506119f4565b60ff841115614e2d57614e2d614676565b50506001821b6119f4565b5060208310610133831016604e8410600b8410161715614e5b575081810a6119f4565b614e658383614d7f565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115614e9757614e97614676565b029392505050565b60006128ae8383614de056fe56657269666961626c6543726564656e7469616c2861646472657373206973737565722c75696e74323536207375626a6563742c75696e743235362065706f63684973737565642c75696e743235362065706f636856616c6964556e74696c2c75696e743235362076616c75652c62797465733420616374696f6e2c75696e743634207461726765742c627974657320636c61696d29a2646970667358221220d3cb77b67f99ab6495ecbf7ce72b90374debdc58740992d480f35f91442cc1ca64736f6c63430008110033",
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

// AuthMsg is a paid mutator transaction binding the contract method 0x2e446475.
//
// Solidity: function authMsg(uint64 accountId, bytes signature, bytes message) returns()
func (_AgentPolice *AgentPoliceTransactor) AuthMsg(opts *bind.TransactOpts, accountId uint64, signature []byte, message []byte) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "authMsg", accountId, signature, message)
}

// AuthMsg is a paid mutator transaction binding the contract method 0x2e446475.
//
// Solidity: function authMsg(uint64 accountId, bytes signature, bytes message) returns()
func (_AgentPolice *AgentPoliceSession) AuthMsg(accountId uint64, signature []byte, message []byte) (*types.Transaction, error) {
	return _AgentPolice.Contract.AuthMsg(&_AgentPolice.TransactOpts, accountId, signature, message)
}

// AuthMsg is a paid mutator transaction binding the contract method 0x2e446475.
//
// Solidity: function authMsg(uint64 accountId, bytes signature, bytes message) returns()
func (_AgentPolice *AgentPoliceTransactorSession) AuthMsg(accountId uint64, signature []byte, message []byte) (*types.Transaction, error) {
	return _AgentPolice.Contract.AuthMsg(&_AgentPolice.TransactOpts, accountId, signature, message)
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

// IsValidCredential is a paid mutator transaction binding the contract method 0xcada34b1.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceTransactor) IsValidCredential(opts *bind.TransactOpts, agent *big.Int, action [4]byte, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "isValidCredential", agent, action, sc)
}

// IsValidCredential is a paid mutator transaction binding the contract method 0xcada34b1.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceSession) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.IsValidCredential(&_AgentPolice.TransactOpts, agent, action, sc)
}

// IsValidCredential is a paid mutator transaction binding the contract method 0xcada34b1.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceTransactorSession) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.IsValidCredential(&_AgentPolice.TransactOpts, agent, action, sc)
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

// ValidateCred is a paid mutator transaction binding the contract method 0x69cdb53c.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceTransactor) ValidateCred(opts *bind.TransactOpts, agent *big.Int, selector [4]byte, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "validateCred", agent, selector, sc)
}

// ValidateCred is a paid mutator transaction binding the contract method 0x69cdb53c.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceSession) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.ValidateCred(&_AgentPolice.TransactOpts, agent, selector, sc)
}

// ValidateCred is a paid mutator transaction binding the contract method 0x69cdb53c.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32,bytes) sc) returns()
func (_AgentPolice *AgentPoliceTransactorSession) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.ValidateCred(&_AgentPolice.TransactOpts, agent, selector, sc)
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
