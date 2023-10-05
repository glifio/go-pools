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

// FilForwarderMetaData contains all meta data concerning the FilForwarder contract.
var FilForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"errorCode\",\"type\":\"int256\"}],\"name\":\"ActorError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ActorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailToCallActor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addr\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"InvalidCodec\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NotEnoughBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"destination\",\"type\":\"bytes\"}],\"name\":\"forward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106f4806100206000396000f3fe60806040526004361061001e5760003560e01c8063d948d46814610023575b600080fd5b61003661003136600461047c565b610038565b005b600061007983838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061008a92505050565b905061008581346100d3565b505050565b60408051602081019091526060815260408051602081019091528281526100b08161011b565b6100cd5760405163e6c4247b60e01b815260040160405180910390fd5b92915050565b81516040805160008082526020820190925290916100f79183908190866000610231565b9050805160001461008557604051630e74990760e01b815260040160405180910390fd5b6000816000015160008151811061013457610134610504565b01602001516001600160f81b03191660000361015457505151600a101590565b8151805160009061016757610167610504565b6020910101516001600160f81b031916600160f81b14806101ad57508151805160009061019657610196610504565b6020910101516001600160f81b031916600160f91b145b156101bb5750515160151490565b815180516000906101ce576101ce610504565b01602001516001600160f81b031916600360f81b036101f05750515160311490565b8151805160009061020357610203610504565b01602001516001600160f81b031916600160fa1b03610226575051516040101590565b505151610100101590565b6060600287511015610261578660405163370d875f60e01b8152600401610258919061056a565b60405180910390fd5b6102726003607f60991b0184610339565b6000806003607f60991b0188868661028b57600061028e565b60015b8a8a8e6040516020016102a696959493929190610584565b60408051601f19818403018152908290526102c0916105d1565b600060405180830381855af49150503d80600081146102fb576040519150601f19603f3d011682016040523d82523d6000602084013e610300565b606091505b50915091508161032357604051638a7db5bf60e01b815260040160405180910390fd5b61032c8161038d565b9998505050505050505050565b478181101561036557604051634787a10360e11b81526004810182905260248101839052604401610258565b823f1515806103875760405163064d954b60e41b815260040160405180910390fd5b50505050565b60606000806000848060200190518101906103a891906105ed565b9194509250905067ffffffffffffffff82166103e3578051156103de57604051630e74990760e01b815260040160405180910390fd5b610452565b67ffffffffffffffff821660511480610406575067ffffffffffffffff82166071145b1561042d5780516000036103de57604051630e74990760e01b815260040160405180910390fd5b60405163f1f6bced60e01b815267ffffffffffffffff83166004820152602401610258565b82156104745760405163d4bb667160e01b815260048101849052602401610258565b949350505050565b6000806020838503121561048f57600080fd5b823567ffffffffffffffff808211156104a757600080fd5b818501915085601f8301126104bb57600080fd5b8135818111156104ca57600080fd5b8660208285010111156104dc57600080fd5b60209290920196919550909350505050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60005b8381101561053557818101518382015260200161051d565b50506000910152565b6000815180845261055681602086016020860161051a565b601f01601f19169290920160200192915050565b60208152600061057d602083018461053e565b9392505050565b600067ffffffffffffffff8089168352876020840152808716604084015280861660608401525060c060808301526105bf60c083018561053e565b82810360a084015261032c818561053e565b600082516105e381846020870161051a565b9190910192915050565b60008060006060848603121561060257600080fd5b83519250602084015167ffffffffffffffff808216821461062257600080fd5b60408601519193508082111561063757600080fd5b818601915086601f83011261064b57600080fd5b81518181111561065d5761065d6104ee565b604051601f8201601f19908116603f01168101908382118183101715610685576106856104ee565b8160405282815289602084870101111561069e57600080fd5b6106af83602083016020880161051a565b8095505050505050925092509256fea2646970667358221220bfb952ccb0bbf828c6dbef4fc93a941e2d5e2e135403570ec1e4c8d3c4fc298764736f6c63430008110033",
}

// FilForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use FilForwarderMetaData.ABI instead.
var FilForwarderABI = FilForwarderMetaData.ABI

// FilForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FilForwarderMetaData.Bin instead.
var FilForwarderBin = FilForwarderMetaData.Bin

// DeployFilForwarder deploys a new Ethereum contract, binding an instance of FilForwarder to it.
func DeployFilForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FilForwarder, error) {
	parsed, err := FilForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FilForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FilForwarder{FilForwarderCaller: FilForwarderCaller{contract: contract}, FilForwarderTransactor: FilForwarderTransactor{contract: contract}, FilForwarderFilterer: FilForwarderFilterer{contract: contract}}, nil
}

// FilForwarder is an auto generated Go binding around an Ethereum contract.
type FilForwarder struct {
	FilForwarderCaller     // Read-only binding to the contract
	FilForwarderTransactor // Write-only binding to the contract
	FilForwarderFilterer   // Log filterer for contract events
}

// FilForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type FilForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FilForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FilForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FilForwarderSession struct {
	Contract     *FilForwarder     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FilForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FilForwarderCallerSession struct {
	Contract *FilForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FilForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FilForwarderTransactorSession struct {
	Contract     *FilForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FilForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type FilForwarderRaw struct {
	Contract *FilForwarder // Generic contract binding to access the raw methods on
}

// FilForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FilForwarderCallerRaw struct {
	Contract *FilForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// FilForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FilForwarderTransactorRaw struct {
	Contract *FilForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFilForwarder creates a new instance of FilForwarder, bound to a specific deployed contract.
func NewFilForwarder(address common.Address, backend bind.ContractBackend) (*FilForwarder, error) {
	contract, err := bindFilForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FilForwarder{FilForwarderCaller: FilForwarderCaller{contract: contract}, FilForwarderTransactor: FilForwarderTransactor{contract: contract}, FilForwarderFilterer: FilForwarderFilterer{contract: contract}}, nil
}

// NewFilForwarderCaller creates a new read-only instance of FilForwarder, bound to a specific deployed contract.
func NewFilForwarderCaller(address common.Address, caller bind.ContractCaller) (*FilForwarderCaller, error) {
	contract, err := bindFilForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FilForwarderCaller{contract: contract}, nil
}

// NewFilForwarderTransactor creates a new write-only instance of FilForwarder, bound to a specific deployed contract.
func NewFilForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*FilForwarderTransactor, error) {
	contract, err := bindFilForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FilForwarderTransactor{contract: contract}, nil
}

// NewFilForwarderFilterer creates a new log filterer instance of FilForwarder, bound to a specific deployed contract.
func NewFilForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*FilForwarderFilterer, error) {
	contract, err := bindFilForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FilForwarderFilterer{contract: contract}, nil
}

// bindFilForwarder binds a generic wrapper to an already deployed contract.
func bindFilForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FilForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilForwarder *FilForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilForwarder.Contract.FilForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilForwarder *FilForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilForwarder.Contract.FilForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilForwarder *FilForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilForwarder.Contract.FilForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilForwarder *FilForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilForwarder *FilForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilForwarder *FilForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilForwarder.Contract.contract.Transact(opts, method, params...)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes destination) payable returns()
func (_FilForwarder *FilForwarderTransactor) Forward(opts *bind.TransactOpts, destination []byte) (*types.Transaction, error) {
	return _FilForwarder.contract.Transact(opts, "forward", destination)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes destination) payable returns()
func (_FilForwarder *FilForwarderSession) Forward(destination []byte) (*types.Transaction, error) {
	return _FilForwarder.Contract.Forward(&_FilForwarder.TransactOpts, destination)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes destination) payable returns()
func (_FilForwarder *FilForwarderTransactorSession) Forward(destination []byte) (*types.Transaction, error) {
	return _FilForwarder.Contract.Forward(&_FilForwarder.TransactOpts, destination)
}
