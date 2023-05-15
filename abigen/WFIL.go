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

// WFILMetaData contains all meta data concerning the WFIL contract.
var WFILMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoveryTimelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052620000134262dd7c0062000253565b6080523480156200002357600080fd5b50604051620012c0380380620012c083398101604081905262000046916200027b565b806040518060400160405280600b81526020016a15dc985c1c19590811925360aa1b8152506040518060400160405280600481526020016315d1925360e21b815250601282600090816200009b919062000352565b506001620000aa838262000352565b506002805460ff191660ff9290921691909117905550506001600160a01b038116620000e95760405163e6c4247b60e01b815260040160405180910390fd5b62000108816001600160a01b03166200015a60201b620009d81760201c565b600680546001600160a01b0319166001600160a01b039290921691821790556040516000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350506200041e565b600080806200016984620001a6565b91509150816200017b57509192915050565b6000806200018983620001d9565b91509150816200019d575093949350505050565b95945050505050565b600080600160401b600160a01b03831660ff60981b8103620001d357600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a81146200022d5760009250600091505b508115806200023d57503d601614155b156200024e57506000928392509050565b915091565b808201808211156200027557634e487b7160e01b600052601160045260246000fd5b92915050565b6000602082840312156200028e57600080fd5b81516001600160a01b0381168114620002a657600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620002d857607f821691505b602082108103620002f957634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200034d57600081815260208120601f850160051c81016020861015620003285750805b601f850160051c820191505b81811015620003495782815560010162000334565b5050505b505050565b81516001600160401b038111156200036e576200036e620002ad565b62000386816200037f8454620002c3565b84620002ff565b602080601f831160018114620003be5760008415620003a55750858301515b600019600386901b1c1916600185901b17855562000349565b600085815260208120601f198616915b82811015620003ef57888601518255948401946001909101908401620003ce565b50858210156200040e5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b608051610e7f620004416000396000818161024801526106d20152610e7f6000f3fe6080604052600436106101025760003560e01c8063670a43c41161009557806395d89b411161006457806395d89b41146102d7578063a9059cbb146102ec578063d0e30db01461030c578063dd62ed3e14610314578063f2fde38b1461033457600080fd5b8063670a43c41461023657806370a082311461026a57806379ba50971461028a5780638da5cb5b1461029f57600080fd5b80632b968958116100d15780632b968958146101b55780632e1a7d4d146101ca578063313ce567146101ea578063436f3b601461021657600080fd5b806306fdde0314610116578063095ea7b31461014157806318160ddd1461017157806323b872dd1461019557600080fd5b366101115761010f610354565b005b600080fd5b34801561012257600080fd5b5061012b610395565b6040516101389190610c75565b60405180910390f35b34801561014d57600080fd5b5061016161015c366004610cdf565b610423565b6040519015158152602001610138565b34801561017d57600080fd5b5061018760035481565b604051908152602001610138565b3480156101a157600080fd5b506101616101b0366004610d09565b61049f565b3480156101c157600080fd5b5061010f6105c8565b3480156101d657600080fd5b5061010f6101e5366004610d45565b61065a565b3480156101f657600080fd5b506002546102049060ff1681565b60405160ff9091168152602001610138565b34801561022257600080fd5b5061010f610231366004610cdf565b6106a6565b34801561024257600080fd5b506101877f000000000000000000000000000000000000000000000000000000000000000081565b34801561027657600080fd5b50610187610285366004610d5e565b610796565b34801561029657600080fd5b5061010f6107ce565b3480156102ab57600080fd5b506006546102bf906001600160a01b031681565b6040516001600160a01b039091168152602001610138565b3480156102e357600080fd5b5061012b610853565b3480156102f857600080fd5b50610161610307366004610cdf565b610860565b61010f610354565b34801561032057600080fd5b5061018761032f366004610d80565b6108f5565b34801561034057600080fd5b5061010f61034f366004610d5e565b610962565b61035e3334610a1f565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546103a290610db3565b80601f01602080910402602001604051908101604052809291908181526020018280546103ce90610db3565b801561041b5780601f106103f05761010080835404028352916020019161041b565b820191906000526020600020905b8154815290600101906020018083116103fe57829003601f168201915b505050505081565b6000610437836001600160a01b03166109d8565b3360008181526005602090815260408083206001600160a01b038616808552908352928190208790555186815293965090927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a35060015b92915050565b60006104b3846001600160a01b03166109d8565b93506104c7836001600160a01b03166109d8565b6001600160a01b03851660009081526005602090815260408083203384529091529020549093506000198114610526576105018382610e03565b6001600160a01b03861660009081526005602090815260408083203384529091529020555b6001600160a01b0385166000908152600460205260408120805485929061054e908490610e03565b90915550506001600160a01b0384166000908152600460205260408120805485929061057b908490610e16565b92505081905550836001600160a01b0316856001600160a01b0316600080516020610e2a833981519152856040516105b591815260200190565b60405180910390a3506001949350505050565b6006546001600160a01b031633146105f2576040516282b42960e81b815260040160405180910390fd5b6007546001600160a01b03161561061b576040516282b42960e81b815260040160405180910390fd5b600680546001600160a01b031916905560405160009033907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3565b6106643382610aab565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a26106a33382610b34565b50565b6006546001600160a01b031633146106d0576040516282b42960e81b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000042101561071157604051637d857b6760e01b815260040160405180910390fd5b6000600354476107219190610e03565b90508082111561073057600080fd5b610742836001600160a01b03166109d8565b925061074e8383610a1f565b826001600160a01b03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c8360405161078991815260200190565b60405180910390a2505050565b6000600460006107ae846001600160a01b03166109d8565b6001600160a01b0316815260208101919091526040016000205492915050565b6007546001600160a01b031633146107f8576040516282b42960e81b815260040160405180910390fd5b600680546001600160a01b0319808216339081179093556007805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b600180546103a290610db3565b6000610874836001600160a01b03166109d8565b33600090815260046020526040812080549295508492909190610898908490610e03565b90915550506001600160a01b038316600090815260046020526040812080548492906108c5908490610e16565b90915550506040518281526001600160a01b038416903390600080516020610e2a8339815191529060200161048d565b60006005600061090d856001600160a01b03166109d8565b6001600160a01b03166001600160a01b031681526020019081526020016000206000610941846001600160a01b03166109d8565b6001600160a01b031681526020810191909152604001600020549392505050565b6006546001600160a01b0316331461098c576040516282b42960e81b815260040160405180910390fd5b600780546001600160a01b0319166001600160a01b03831690811790915560405133907f3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a90600090a350565b60008060006109e684610bce565b91509150816109f757509192915050565b600080610a0383610c0b565b9150915081610a16575093949350505050565b95945050505050565b610a31826001600160a01b03166109d8565b91508060036000828254610a459190610e16565b90915550506001600160a01b03821660009081526004602052604081208054839290610a72908490610e16565b90915550506040518181526001600160a01b03831690600090600080516020610e2a833981519152906020015b60405180910390a35050565b610abd826001600160a01b03166109d8565b6001600160a01b038116600090815260046020526040812080549294508392909190610aea908490610e03565b925050819055508060036000828254610b039190610e03565b90915550506040518181526000906001600160a01b03841690600080516020610e2a83398151915290602001610a9f565b80471015610b555760405163356680b760e01b815260040160405180910390fd5b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610ba2576040519150601f19603f3d011682016040523d82523d6000602084013e610ba7565b606091505b5050905080610bc957604051633204506f60e01b815260040160405180910390fd5b505050565b60008073ffffffffffffffffffffffff0000000000000000831660ff60981b8103610c05576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a602060006002607f60991b015afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114610c515760009250600091505b50811580610c6057503d601614155b15610c7057506000928392509050565b915091565b600060208083528351808285015260005b81811015610ca257858101830151858201604001528201610c86565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610cda57600080fd5b919050565b60008060408385031215610cf257600080fd5b610cfb83610cc3565b946020939093013593505050565b600080600060608486031215610d1e57600080fd5b610d2784610cc3565b9250610d3560208501610cc3565b9150604084013590509250925092565b600060208284031215610d5757600080fd5b5035919050565b600060208284031215610d7057600080fd5b610d7982610cc3565b9392505050565b60008060408385031215610d9357600080fd5b610d9c83610cc3565b9150610daa60208401610cc3565b90509250929050565b600181811c90821680610dc757607f821691505b602082108103610de757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561049957610499610ded565b8082018082111561049957610499610ded56feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212204ebefe9b936138732eb202276acbb2fbf47fad86f8721a2d74d28d675015b14e64736f6c63430008110033",
}

// WFILABI is the input ABI used to generate the binding from.
// Deprecated: Use WFILMetaData.ABI instead.
var WFILABI = WFILMetaData.ABI

// WFILBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WFILMetaData.Bin instead.
var WFILBin = WFILMetaData.Bin

// DeployWFIL deploys a new Ethereum contract, binding an instance of WFIL to it.
func DeployWFIL(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *WFIL, error) {
	parsed, err := WFILMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WFILBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WFIL{WFILCaller: WFILCaller{contract: contract}, WFILTransactor: WFILTransactor{contract: contract}, WFILFilterer: WFILFilterer{contract: contract}}, nil
}

// WFIL is an auto generated Go binding around an Ethereum contract.
type WFIL struct {
	WFILCaller     // Read-only binding to the contract
	WFILTransactor // Write-only binding to the contract
	WFILFilterer   // Log filterer for contract events
}

// WFILCaller is an auto generated read-only Go binding around an Ethereum contract.
type WFILCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WFILTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WFILTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WFILFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WFILFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WFILSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WFILSession struct {
	Contract     *WFIL             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WFILCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WFILCallerSession struct {
	Contract *WFILCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WFILTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WFILTransactorSession struct {
	Contract     *WFILTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WFILRaw is an auto generated low-level Go binding around an Ethereum contract.
type WFILRaw struct {
	Contract *WFIL // Generic contract binding to access the raw methods on
}

// WFILCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WFILCallerRaw struct {
	Contract *WFILCaller // Generic read-only contract binding to access the raw methods on
}

// WFILTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WFILTransactorRaw struct {
	Contract *WFILTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWFIL creates a new instance of WFIL, bound to a specific deployed contract.
func NewWFIL(address common.Address, backend bind.ContractBackend) (*WFIL, error) {
	contract, err := bindWFIL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WFIL{WFILCaller: WFILCaller{contract: contract}, WFILTransactor: WFILTransactor{contract: contract}, WFILFilterer: WFILFilterer{contract: contract}}, nil
}

// NewWFILCaller creates a new read-only instance of WFIL, bound to a specific deployed contract.
func NewWFILCaller(address common.Address, caller bind.ContractCaller) (*WFILCaller, error) {
	contract, err := bindWFIL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WFILCaller{contract: contract}, nil
}

// NewWFILTransactor creates a new write-only instance of WFIL, bound to a specific deployed contract.
func NewWFILTransactor(address common.Address, transactor bind.ContractTransactor) (*WFILTransactor, error) {
	contract, err := bindWFIL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WFILTransactor{contract: contract}, nil
}

// NewWFILFilterer creates a new log filterer instance of WFIL, bound to a specific deployed contract.
func NewWFILFilterer(address common.Address, filterer bind.ContractFilterer) (*WFILFilterer, error) {
	contract, err := bindWFIL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WFILFilterer{contract: contract}, nil
}

// bindWFIL binds a generic wrapper to an already deployed contract.
func bindWFIL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WFILMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WFIL *WFILRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WFIL.Contract.WFILCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WFIL *WFILRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WFIL.Contract.WFILTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WFIL *WFILRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WFIL.Contract.WFILTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WFIL *WFILCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WFIL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WFIL *WFILTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WFIL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WFIL *WFILTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WFIL.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WFIL *WFILCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WFIL *WFILSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WFIL.Contract.Allowance(&_WFIL.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WFIL *WFILCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WFIL.Contract.Allowance(&_WFIL.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_WFIL *WFILCaller) BalanceOf(opts *bind.CallOpts, _a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "balanceOf", _a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_WFIL *WFILSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _WFIL.Contract.BalanceOf(&_WFIL.CallOpts, _a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_WFIL *WFILCallerSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _WFIL.Contract.BalanceOf(&_WFIL.CallOpts, _a)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WFIL *WFILCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WFIL *WFILSession) Decimals() (uint8, error) {
	return _WFIL.Contract.Decimals(&_WFIL.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WFIL *WFILCallerSession) Decimals() (uint8, error) {
	return _WFIL.Contract.Decimals(&_WFIL.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WFIL *WFILCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WFIL *WFILSession) Name() (string, error) {
	return _WFIL.Contract.Name(&_WFIL.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WFIL *WFILCallerSession) Name() (string, error) {
	return _WFIL.Contract.Name(&_WFIL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WFIL *WFILCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WFIL *WFILSession) Owner() (common.Address, error) {
	return _WFIL.Contract.Owner(&_WFIL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WFIL *WFILCallerSession) Owner() (common.Address, error) {
	return _WFIL.Contract.Owner(&_WFIL.CallOpts)
}

// RecoveryTimelock is a free data retrieval call binding the contract method 0x670a43c4.
//
// Solidity: function recoveryTimelock() view returns(uint256)
func (_WFIL *WFILCaller) RecoveryTimelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "recoveryTimelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecoveryTimelock is a free data retrieval call binding the contract method 0x670a43c4.
//
// Solidity: function recoveryTimelock() view returns(uint256)
func (_WFIL *WFILSession) RecoveryTimelock() (*big.Int, error) {
	return _WFIL.Contract.RecoveryTimelock(&_WFIL.CallOpts)
}

// RecoveryTimelock is a free data retrieval call binding the contract method 0x670a43c4.
//
// Solidity: function recoveryTimelock() view returns(uint256)
func (_WFIL *WFILCallerSession) RecoveryTimelock() (*big.Int, error) {
	return _WFIL.Contract.RecoveryTimelock(&_WFIL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WFIL *WFILCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WFIL *WFILSession) Symbol() (string, error) {
	return _WFIL.Contract.Symbol(&_WFIL.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WFIL *WFILCallerSession) Symbol() (string, error) {
	return _WFIL.Contract.Symbol(&_WFIL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WFIL *WFILCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WFIL.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WFIL *WFILSession) TotalSupply() (*big.Int, error) {
	return _WFIL.Contract.TotalSupply(&_WFIL.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WFIL *WFILCallerSession) TotalSupply() (*big.Int, error) {
	return _WFIL.Contract.TotalSupply(&_WFIL.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WFIL *WFILTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WFIL *WFILSession) AcceptOwnership() (*types.Transaction, error) {
	return _WFIL.Contract.AcceptOwnership(&_WFIL.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WFIL *WFILTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WFIL.Contract.AcceptOwnership(&_WFIL.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_WFIL *WFILTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_WFIL *WFILSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.Approve(&_WFIL.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_WFIL *WFILTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.Approve(&_WFIL.TransactOpts, _spender, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WFIL *WFILTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WFIL *WFILSession) Deposit() (*types.Transaction, error) {
	return _WFIL.Contract.Deposit(&_WFIL.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WFIL *WFILTransactorSession) Deposit() (*types.Transaction, error) {
	return _WFIL.Contract.Deposit(&_WFIL.TransactOpts)
}

// RecoverDeposit is a paid mutator transaction binding the contract method 0x436f3b60.
//
// Solidity: function recoverDeposit(address _depositor, uint256 _amount) returns()
func (_WFIL *WFILTransactor) RecoverDeposit(opts *bind.TransactOpts, _depositor common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "recoverDeposit", _depositor, _amount)
}

// RecoverDeposit is a paid mutator transaction binding the contract method 0x436f3b60.
//
// Solidity: function recoverDeposit(address _depositor, uint256 _amount) returns()
func (_WFIL *WFILSession) RecoverDeposit(_depositor common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.RecoverDeposit(&_WFIL.TransactOpts, _depositor, _amount)
}

// RecoverDeposit is a paid mutator transaction binding the contract method 0x436f3b60.
//
// Solidity: function recoverDeposit(address _depositor, uint256 _amount) returns()
func (_WFIL *WFILTransactorSession) RecoverDeposit(_depositor common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.RecoverDeposit(&_WFIL.TransactOpts, _depositor, _amount)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_WFIL *WFILTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_WFIL *WFILSession) RevokeOwnership() (*types.Transaction, error) {
	return _WFIL.Contract.RevokeOwnership(&_WFIL.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_WFIL *WFILTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _WFIL.Contract.RevokeOwnership(&_WFIL.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_WFIL *WFILTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_WFIL *WFILSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.Transfer(&_WFIL.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_WFIL *WFILTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.Transfer(&_WFIL.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_WFIL *WFILTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "transferFrom", _owner, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_WFIL *WFILSession) TransferFrom(_owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.TransferFrom(&_WFIL.TransactOpts, _owner, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_WFIL *WFILTransactorSession) TransferFrom(_owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.TransferFrom(&_WFIL.TransactOpts, _owner, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_WFIL *WFILTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_WFIL *WFILSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _WFIL.Contract.TransferOwnership(&_WFIL.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_WFIL *WFILTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _WFIL.Contract.TransferOwnership(&_WFIL.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_WFIL *WFILTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _WFIL.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_WFIL *WFILSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.Withdraw(&_WFIL.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_WFIL *WFILTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _WFIL.Contract.Withdraw(&_WFIL.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WFIL *WFILTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WFIL.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WFIL *WFILSession) Receive() (*types.Transaction, error) {
	return _WFIL.Contract.Receive(&_WFIL.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WFIL *WFILTransactorSession) Receive() (*types.Transaction, error) {
	return _WFIL.Contract.Receive(&_WFIL.TransactOpts)
}

// WFILApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WFIL contract.
type WFILApprovalIterator struct {
	Event *WFILApproval // Event containing the contract specifics and raw log

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
func (it *WFILApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WFILApproval)
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
		it.Event = new(WFILApproval)
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
func (it *WFILApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WFILApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WFILApproval represents a Approval event raised by the WFIL contract.
type WFILApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_WFIL *WFILFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WFILApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WFIL.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WFILApprovalIterator{contract: _WFIL.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_WFIL *WFILFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WFILApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WFIL.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WFILApproval)
				if err := _WFIL.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_WFIL *WFILFilterer) ParseApproval(log types.Log) (*WFILApproval, error) {
	event := new(WFILApproval)
	if err := _WFIL.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WFILDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WFIL contract.
type WFILDepositIterator struct {
	Event *WFILDeposit // Event containing the contract specifics and raw log

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
func (it *WFILDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WFILDeposit)
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
		it.Event = new(WFILDeposit)
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
func (it *WFILDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WFILDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WFILDeposit represents a Deposit event raised by the WFIL contract.
type WFILDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_WFIL *WFILFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address) (*WFILDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WFIL.contract.FilterLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return &WFILDepositIterator{contract: _WFIL.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_WFIL *WFILFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WFILDeposit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _WFIL.contract.WatchLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WFILDeposit)
				if err := _WFIL.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_WFIL *WFILFilterer) ParseDeposit(log types.Log) (*WFILDeposit, error) {
	event := new(WFILDeposit)
	if err := _WFIL.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WFILOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the WFIL contract.
type WFILOwnershipPendingIterator struct {
	Event *WFILOwnershipPending // Event containing the contract specifics and raw log

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
func (it *WFILOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WFILOwnershipPending)
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
		it.Event = new(WFILOwnershipPending)
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
func (it *WFILOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WFILOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WFILOwnershipPending represents a OwnershipPending event raised by the WFIL contract.
type WFILOwnershipPending struct {
	CurrentOwner common.Address
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_WFIL *WFILFilterer) FilterOwnershipPending(opts *bind.FilterOpts, currentOwner []common.Address, pendingOwner []common.Address) (*WFILOwnershipPendingIterator, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _WFIL.contract.FilterLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WFILOwnershipPendingIterator{contract: _WFIL.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_WFIL *WFILFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *WFILOwnershipPending, currentOwner []common.Address, pendingOwner []common.Address) (event.Subscription, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _WFIL.contract.WatchLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WFILOwnershipPending)
				if err := _WFIL.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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

// ParseOwnershipPending is a log parse operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_WFIL *WFILFilterer) ParseOwnershipPending(log types.Log) (*WFILOwnershipPending, error) {
	event := new(WFILOwnershipPending)
	if err := _WFIL.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WFILOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WFIL contract.
type WFILOwnershipTransferredIterator struct {
	Event *WFILOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WFILOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WFILOwnershipTransferred)
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
		it.Event = new(WFILOwnershipTransferred)
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
func (it *WFILOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WFILOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WFILOwnershipTransferred represents a OwnershipTransferred event raised by the WFIL contract.
type WFILOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_WFIL *WFILFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*WFILOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WFIL.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WFILOwnershipTransferredIterator{contract: _WFIL.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_WFIL *WFILFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WFILOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WFIL.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WFILOwnershipTransferred)
				if err := _WFIL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_WFIL *WFILFilterer) ParseOwnershipTransferred(log types.Log) (*WFILOwnershipTransferred, error) {
	event := new(WFILOwnershipTransferred)
	if err := _WFIL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WFILTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WFIL contract.
type WFILTransferIterator struct {
	Event *WFILTransfer // Event containing the contract specifics and raw log

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
func (it *WFILTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WFILTransfer)
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
		it.Event = new(WFILTransfer)
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
func (it *WFILTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WFILTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WFILTransfer represents a Transfer event raised by the WFIL contract.
type WFILTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_WFIL *WFILFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WFILTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WFIL.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WFILTransferIterator{contract: _WFIL.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_WFIL *WFILFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WFILTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WFIL.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WFILTransfer)
				if err := _WFIL.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_WFIL *WFILFilterer) ParseTransfer(log types.Log) (*WFILTransfer, error) {
	event := new(WFILTransfer)
	if err := _WFIL.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WFILWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WFIL contract.
type WFILWithdrawalIterator struct {
	Event *WFILWithdrawal // Event containing the contract specifics and raw log

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
func (it *WFILWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WFILWithdrawal)
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
		it.Event = new(WFILWithdrawal)
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
func (it *WFILWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WFILWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WFILWithdrawal represents a Withdrawal event raised by the WFIL contract.
type WFILWithdrawal struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed to, uint256 amount)
func (_WFIL *WFILFilterer) FilterWithdrawal(opts *bind.FilterOpts, to []common.Address) (*WFILWithdrawalIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WFIL.contract.FilterLogs(opts, "Withdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return &WFILWithdrawalIterator{contract: _WFIL.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed to, uint256 amount)
func (_WFIL *WFILFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WFILWithdrawal, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WFIL.contract.WatchLogs(opts, "Withdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WFILWithdrawal)
				if err := _WFIL.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed to, uint256 amount)
func (_WFIL *WFILFilterer) ParseWithdrawal(log types.Log) (*WFILWithdrawal, error) {
	event := new(WFILWithdrawal)
	if err := _WFIL.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
