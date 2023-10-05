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

// CredParserMetaData contains all meta data concerning the CredParser contract.
var CredParserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getAgentValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getExpectedDailyRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getFaultySectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getGCRED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getGreenScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getLiveSectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getPrincipal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_agentData\",\"type\":\"bytes\"}],\"name\":\"getQAPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610469806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c806350543540116100765780636bb0d0cc1161005b5780636bb0d0cc1461012c578063b661e4821461013f578063d9bc1b731461015257600080fd5b8063505435401461010657806368fd63d31461011957600080fd5b806307124c06146100a857806317bbd29a146100cd5780631b2b5fad146100e0578063402ed8cb146100f3575b600080fd5b6100bb6100b63660046102dd565b610165565b60405190815260200160405180910390f35b6100bb6100db3660046102dd565b610185565b6100bb6100ee3660046102dd565b6101a5565b6100bb6101013660046102dd565b6101c5565b6100bb6101143660046102dd565b6101e6565b6100bb6101273660046102dd565b610206565b6100bb61013a3660046102dd565b610226565b6100bb61014d3660046102dd565b610243565b6100bb6101603660046102dd565b610264565b60008180602001905181019061017b91906103ac565b60e0015192915050565b60008180602001905181019061019b91906103ac565b60c0015192915050565b6000818060200190518101906101bb91906103ac565b6020015192915050565b6000818060200190518101906101db91906103ac565b610100015192915050565b6000818060200190518101906101fc91906103ac565b60a0015192915050565b60008180602001905181019061021c91906103ac565b6080015192915050565b60008180602001905181019061023c91906103ac565b5192915050565b60008180602001905181019061025991906103ac565b610120015192915050565b60008180602001905181019061027a91906103ac565b6060015192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff811182821017156102d7576102d7610284565b60405290565b6000602082840312156102ef57600080fd5b813567ffffffffffffffff8082111561030757600080fd5b818401915084601f83011261031b57600080fd5b81358181111561032d5761032d610284565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561037357610373610284565b8160405282815287602084870101111561038c57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600061014082840312156103bf57600080fd5b6103c76102b3565b825181526020830151602082015260408301516040820152606083015160608201526080830151608082015260a083015160a082015260c083015160c082015260e083015160e0820152610100808401518183015250610120808401518183015250809150509291505056fea264697066735822122080629255c8db8d5082c57af69d654fbdc134bf005e73e0e75eea3ed1853bb30264736f6c63430008110033",
}

// CredParserABI is the input ABI used to generate the binding from.
// Deprecated: Use CredParserMetaData.ABI instead.
var CredParserABI = CredParserMetaData.ABI

// CredParserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CredParserMetaData.Bin instead.
var CredParserBin = CredParserMetaData.Bin

// DeployCredParser deploys a new Ethereum contract, binding an instance of CredParser to it.
func DeployCredParser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CredParser, error) {
	parsed, err := CredParserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CredParserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CredParser{CredParserCaller: CredParserCaller{contract: contract}, CredParserTransactor: CredParserTransactor{contract: contract}, CredParserFilterer: CredParserFilterer{contract: contract}}, nil
}

// CredParser is an auto generated Go binding around an Ethereum contract.
type CredParser struct {
	CredParserCaller     // Read-only binding to the contract
	CredParserTransactor // Write-only binding to the contract
	CredParserFilterer   // Log filterer for contract events
}

// CredParserCaller is an auto generated read-only Go binding around an Ethereum contract.
type CredParserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredParserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CredParserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredParserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CredParserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredParserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CredParserSession struct {
	Contract     *CredParser       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CredParserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CredParserCallerSession struct {
	Contract *CredParserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CredParserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CredParserTransactorSession struct {
	Contract     *CredParserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CredParserRaw is an auto generated low-level Go binding around an Ethereum contract.
type CredParserRaw struct {
	Contract *CredParser // Generic contract binding to access the raw methods on
}

// CredParserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CredParserCallerRaw struct {
	Contract *CredParserCaller // Generic read-only contract binding to access the raw methods on
}

// CredParserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CredParserTransactorRaw struct {
	Contract *CredParserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredParser creates a new instance of CredParser, bound to a specific deployed contract.
func NewCredParser(address common.Address, backend bind.ContractBackend) (*CredParser, error) {
	contract, err := bindCredParser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CredParser{CredParserCaller: CredParserCaller{contract: contract}, CredParserTransactor: CredParserTransactor{contract: contract}, CredParserFilterer: CredParserFilterer{contract: contract}}, nil
}

// NewCredParserCaller creates a new read-only instance of CredParser, bound to a specific deployed contract.
func NewCredParserCaller(address common.Address, caller bind.ContractCaller) (*CredParserCaller, error) {
	contract, err := bindCredParser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CredParserCaller{contract: contract}, nil
}

// NewCredParserTransactor creates a new write-only instance of CredParser, bound to a specific deployed contract.
func NewCredParserTransactor(address common.Address, transactor bind.ContractTransactor) (*CredParserTransactor, error) {
	contract, err := bindCredParser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CredParserTransactor{contract: contract}, nil
}

// NewCredParserFilterer creates a new log filterer instance of CredParser, bound to a specific deployed contract.
func NewCredParserFilterer(address common.Address, filterer bind.ContractFilterer) (*CredParserFilterer, error) {
	contract, err := bindCredParser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CredParserFilterer{contract: contract}, nil
}

// bindCredParser binds a generic wrapper to an already deployed contract.
func bindCredParser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CredParserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CredParser *CredParserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CredParser.Contract.CredParserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CredParser *CredParserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CredParser.Contract.CredParserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CredParser *CredParserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CredParser.Contract.CredParserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CredParser *CredParserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CredParser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CredParser *CredParserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CredParser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CredParser *CredParserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CredParser.Contract.contract.Transact(opts, method, params...)
}

// GetAgentValue is a free data retrieval call binding the contract method 0x6bb0d0cc.
//
// Solidity: function getAgentValue(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetAgentValue(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getAgentValue", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentValue is a free data retrieval call binding the contract method 0x6bb0d0cc.
//
// Solidity: function getAgentValue(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetAgentValue(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetAgentValue(&_CredParser.CallOpts, _agentData)
}

// GetAgentValue is a free data retrieval call binding the contract method 0x6bb0d0cc.
//
// Solidity: function getAgentValue(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetAgentValue(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetAgentValue(&_CredParser.CallOpts, _agentData)
}

// GetCollateralValue is a free data retrieval call binding the contract method 0x1b2b5fad.
//
// Solidity: function getCollateralValue(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetCollateralValue(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getCollateralValue", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralValue is a free data retrieval call binding the contract method 0x1b2b5fad.
//
// Solidity: function getCollateralValue(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetCollateralValue(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetCollateralValue(&_CredParser.CallOpts, _agentData)
}

// GetCollateralValue is a free data retrieval call binding the contract method 0x1b2b5fad.
//
// Solidity: function getCollateralValue(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetCollateralValue(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetCollateralValue(&_CredParser.CallOpts, _agentData)
}

// GetExpectedDailyRewards is a free data retrieval call binding the contract method 0xd9bc1b73.
//
// Solidity: function getExpectedDailyRewards(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetExpectedDailyRewards(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getExpectedDailyRewards", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedDailyRewards is a free data retrieval call binding the contract method 0xd9bc1b73.
//
// Solidity: function getExpectedDailyRewards(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetExpectedDailyRewards(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetExpectedDailyRewards(&_CredParser.CallOpts, _agentData)
}

// GetExpectedDailyRewards is a free data retrieval call binding the contract method 0xd9bc1b73.
//
// Solidity: function getExpectedDailyRewards(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetExpectedDailyRewards(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetExpectedDailyRewards(&_CredParser.CallOpts, _agentData)
}

// GetFaultySectors is a free data retrieval call binding the contract method 0x07124c06.
//
// Solidity: function getFaultySectors(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetFaultySectors(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getFaultySectors", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFaultySectors is a free data retrieval call binding the contract method 0x07124c06.
//
// Solidity: function getFaultySectors(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetFaultySectors(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetFaultySectors(&_CredParser.CallOpts, _agentData)
}

// GetFaultySectors is a free data retrieval call binding the contract method 0x07124c06.
//
// Solidity: function getFaultySectors(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetFaultySectors(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetFaultySectors(&_CredParser.CallOpts, _agentData)
}

// GetGCRED is a free data retrieval call binding the contract method 0x68fd63d3.
//
// Solidity: function getGCRED(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetGCRED(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getGCRED", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGCRED is a free data retrieval call binding the contract method 0x68fd63d3.
//
// Solidity: function getGCRED(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetGCRED(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetGCRED(&_CredParser.CallOpts, _agentData)
}

// GetGCRED is a free data retrieval call binding the contract method 0x68fd63d3.
//
// Solidity: function getGCRED(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetGCRED(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetGCRED(&_CredParser.CallOpts, _agentData)
}

// GetGreenScore is a free data retrieval call binding the contract method 0xb661e482.
//
// Solidity: function getGreenScore(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetGreenScore(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getGreenScore", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGreenScore is a free data retrieval call binding the contract method 0xb661e482.
//
// Solidity: function getGreenScore(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetGreenScore(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetGreenScore(&_CredParser.CallOpts, _agentData)
}

// GetGreenScore is a free data retrieval call binding the contract method 0xb661e482.
//
// Solidity: function getGreenScore(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetGreenScore(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetGreenScore(&_CredParser.CallOpts, _agentData)
}

// GetLiveSectors is a free data retrieval call binding the contract method 0x402ed8cb.
//
// Solidity: function getLiveSectors(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetLiveSectors(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getLiveSectors", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiveSectors is a free data retrieval call binding the contract method 0x402ed8cb.
//
// Solidity: function getLiveSectors(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetLiveSectors(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetLiveSectors(&_CredParser.CallOpts, _agentData)
}

// GetLiveSectors is a free data retrieval call binding the contract method 0x402ed8cb.
//
// Solidity: function getLiveSectors(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetLiveSectors(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetLiveSectors(&_CredParser.CallOpts, _agentData)
}

// GetPrincipal is a free data retrieval call binding the contract method 0x17bbd29a.
//
// Solidity: function getPrincipal(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetPrincipal(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getPrincipal", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrincipal is a free data retrieval call binding the contract method 0x17bbd29a.
//
// Solidity: function getPrincipal(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetPrincipal(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetPrincipal(&_CredParser.CallOpts, _agentData)
}

// GetPrincipal is a free data retrieval call binding the contract method 0x17bbd29a.
//
// Solidity: function getPrincipal(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetPrincipal(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetPrincipal(&_CredParser.CallOpts, _agentData)
}

// GetQAPower is a free data retrieval call binding the contract method 0x50543540.
//
// Solidity: function getQAPower(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCaller) GetQAPower(opts *bind.CallOpts, _agentData []byte) (*big.Int, error) {
	var out []interface{}
	err := _CredParser.contract.Call(opts, &out, "getQAPower", _agentData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQAPower is a free data retrieval call binding the contract method 0x50543540.
//
// Solidity: function getQAPower(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserSession) GetQAPower(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetQAPower(&_CredParser.CallOpts, _agentData)
}

// GetQAPower is a free data retrieval call binding the contract method 0x50543540.
//
// Solidity: function getQAPower(bytes _agentData) pure returns(uint256)
func (_CredParser *CredParserCallerSession) GetQAPower(_agentData []byte) (*big.Int, error) {
	return _CredParser.Contract.GetQAPower(&_CredParser.CallOpts, _agentData)
}
