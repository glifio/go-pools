// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/jimpick/go-ethereum"
	"github.com/jimpick/go-ethereum/accounts/abi"
	"github.com/jimpick/go-ethereum/accounts/abi/bind"
	"github.com/jimpick/go-ethereum/common"
	"github.com/jimpick/go-ethereum/core/types"
	"github.com/jimpick/go-ethereum/event"
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
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"errorCode\",\"type\":\"int256\"}],\"name\":\"ActorError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailToCallActor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addr\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"InvalidCodec\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NotEnoughBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"destination\",\"type\":\"bytes\"}],\"name\":\"forward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// FilForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use FilForwarderMetaData.ABI instead.
var FilForwarderABI = FilForwarderMetaData.ABI

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
