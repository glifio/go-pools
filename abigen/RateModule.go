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

// Account is an auto generated low-level Go binding around an user-defined struct.
type Account_duplicate_RateModule struct {
	StartEpoch *big.Int
	Principal  *big.Int
	EpochsPaid *big.Int
	Defaulted  bool
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_RateModule struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// RateModuleMetaData contains all meta data concerning the RateModule contract.
var RateModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256[61]\",\"name\":\"_rateLookup\",\"type\":\"uint256[61]\"},{\"internalType\":\"uint256[10]\",\"name\":\"_levels\",\"type\":\"uint256[10]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agentTotalValue\",\"type\":\"uint256\"}],\"name\":\"computeDTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedDailyRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountPrincipal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipal\",\"type\":\"uint256\"}],\"name\":\"computeDTI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"name\":\"computeLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credParser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"defaulted\",\"type\":\"bool\"}],\"internalType\":\"structAccount\",\"name\":\"account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDTI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGCRED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rateLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"agentIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"}],\"name\":\"setAgentLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseRate\",\"type\":\"uint256\"}],\"name\":\"setBaseRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"_levels\",\"type\":\"uint256[10]\"}],\"name\":\"setLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDTE\",\"type\":\"uint256\"}],\"name\":\"setMaxDTE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDTI\",\"type\":\"uint256\"}],\"name\":\"setMaxDTI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLTV\",\"type\":\"uint256\"}],\"name\":\"setMaxLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minGCRED\",\"type\":\"uint256\"}],\"name\":\"setMinGCRED\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[61]\",\"name\":\"_rateLookup\",\"type\":\"uint256[61]\"}],\"name\":\"setRateLookup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCredParser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RateModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use RateModuleMetaData.ABI instead.
var RateModuleABI = RateModuleMetaData.ABI

// RateModule is an auto generated Go binding around an Ethereum contract.
type RateModule struct {
	RateModuleCaller     // Read-only binding to the contract
	RateModuleTransactor // Write-only binding to the contract
	RateModuleFilterer   // Log filterer for contract events
}

// RateModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RateModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RateModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RateModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RateModuleSession struct {
	Contract     *RateModule       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RateModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RateModuleCallerSession struct {
	Contract *RateModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RateModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RateModuleTransactorSession struct {
	Contract     *RateModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RateModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RateModuleRaw struct {
	Contract *RateModule // Generic contract binding to access the raw methods on
}

// RateModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RateModuleCallerRaw struct {
	Contract *RateModuleCaller // Generic read-only contract binding to access the raw methods on
}

// RateModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RateModuleTransactorRaw struct {
	Contract *RateModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRateModule creates a new instance of RateModule, bound to a specific deployed contract.
func NewRateModule(address common.Address, backend bind.ContractBackend) (*RateModule, error) {
	contract, err := bindRateModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RateModule{RateModuleCaller: RateModuleCaller{contract: contract}, RateModuleTransactor: RateModuleTransactor{contract: contract}, RateModuleFilterer: RateModuleFilterer{contract: contract}}, nil
}

// NewRateModuleCaller creates a new read-only instance of RateModule, bound to a specific deployed contract.
func NewRateModuleCaller(address common.Address, caller bind.ContractCaller) (*RateModuleCaller, error) {
	contract, err := bindRateModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RateModuleCaller{contract: contract}, nil
}

// NewRateModuleTransactor creates a new write-only instance of RateModule, bound to a specific deployed contract.
func NewRateModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*RateModuleTransactor, error) {
	contract, err := bindRateModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RateModuleTransactor{contract: contract}, nil
}

// NewRateModuleFilterer creates a new log filterer instance of RateModule, bound to a specific deployed contract.
func NewRateModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*RateModuleFilterer, error) {
	contract, err := bindRateModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RateModuleFilterer{contract: contract}, nil
}

// bindRateModule binds a generic wrapper to an already deployed contract.
func bindRateModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RateModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RateModule *RateModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RateModule.Contract.RateModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RateModule *RateModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.Contract.RateModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RateModule *RateModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RateModule.Contract.RateModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RateModule *RateModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RateModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RateModule *RateModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RateModule *RateModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RateModule.Contract.contract.Transact(opts, method, params...)
}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCaller) AccountLevel(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "accountLevel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_RateModule *RateModuleSession) AccountLevel(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.AccountLevel(&_RateModule.CallOpts, arg0)
}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCallerSession) AccountLevel(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.AccountLevel(&_RateModule.CallOpts, arg0)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_RateModule *RateModuleCaller) BaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "baseRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_RateModule *RateModuleSession) BaseRate() (*big.Int, error) {
	return _RateModule.Contract.BaseRate(&_RateModule.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_RateModule *RateModuleCallerSession) BaseRate() (*big.Int, error) {
	return _RateModule.Contract.BaseRate(&_RateModule.CallOpts)
}

// ComputeDTE is a free data retrieval call binding the contract method 0xe4649162.
//
// Solidity: function computeDTE(uint256 principal, uint256 agentTotalValue) pure returns(uint256)
func (_RateModule *RateModuleCaller) ComputeDTE(opts *bind.CallOpts, principal *big.Int, agentTotalValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "computeDTE", principal, agentTotalValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeDTE is a free data retrieval call binding the contract method 0xe4649162.
//
// Solidity: function computeDTE(uint256 principal, uint256 agentTotalValue) pure returns(uint256)
func (_RateModule *RateModuleSession) ComputeDTE(principal *big.Int, agentTotalValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTE(&_RateModule.CallOpts, principal, agentTotalValue)
}

// ComputeDTE is a free data retrieval call binding the contract method 0xe4649162.
//
// Solidity: function computeDTE(uint256 principal, uint256 agentTotalValue) pure returns(uint256)
func (_RateModule *RateModuleCallerSession) ComputeDTE(principal *big.Int, agentTotalValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTE(&_RateModule.CallOpts, principal, agentTotalValue)
}

// ComputeDTI is a free data retrieval call binding the contract method 0xedf2460b.
//
// Solidity: function computeDTI(uint256 expectedDailyRewards, uint256 rate, uint256 accountPrincipal, uint256 totalPrincipal) pure returns(uint256)
func (_RateModule *RateModuleCaller) ComputeDTI(opts *bind.CallOpts, expectedDailyRewards *big.Int, rate *big.Int, accountPrincipal *big.Int, totalPrincipal *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "computeDTI", expectedDailyRewards, rate, accountPrincipal, totalPrincipal)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeDTI is a free data retrieval call binding the contract method 0xedf2460b.
//
// Solidity: function computeDTI(uint256 expectedDailyRewards, uint256 rate, uint256 accountPrincipal, uint256 totalPrincipal) pure returns(uint256)
func (_RateModule *RateModuleSession) ComputeDTI(expectedDailyRewards *big.Int, rate *big.Int, accountPrincipal *big.Int, totalPrincipal *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTI(&_RateModule.CallOpts, expectedDailyRewards, rate, accountPrincipal, totalPrincipal)
}

// ComputeDTI is a free data retrieval call binding the contract method 0xedf2460b.
//
// Solidity: function computeDTI(uint256 expectedDailyRewards, uint256 rate, uint256 accountPrincipal, uint256 totalPrincipal) pure returns(uint256)
func (_RateModule *RateModuleCallerSession) ComputeDTI(expectedDailyRewards *big.Int, rate *big.Int, accountPrincipal *big.Int, totalPrincipal *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTI(&_RateModule.CallOpts, expectedDailyRewards, rate, accountPrincipal, totalPrincipal)
}

// ComputeLTV is a free data retrieval call binding the contract method 0x3705afb6.
//
// Solidity: function computeLTV(uint256 principal, uint256 collateralValue) pure returns(uint256)
func (_RateModule *RateModuleCaller) ComputeLTV(opts *bind.CallOpts, principal *big.Int, collateralValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "computeLTV", principal, collateralValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeLTV is a free data retrieval call binding the contract method 0x3705afb6.
//
// Solidity: function computeLTV(uint256 principal, uint256 collateralValue) pure returns(uint256)
func (_RateModule *RateModuleSession) ComputeLTV(principal *big.Int, collateralValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeLTV(&_RateModule.CallOpts, principal, collateralValue)
}

// ComputeLTV is a free data retrieval call binding the contract method 0x3705afb6.
//
// Solidity: function computeLTV(uint256 principal, uint256 collateralValue) pure returns(uint256)
func (_RateModule *RateModuleCallerSession) ComputeLTV(principal *big.Int, collateralValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeLTV(&_RateModule.CallOpts, principal, collateralValue)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_RateModule *RateModuleCaller) CredParser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "credParser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_RateModule *RateModuleSession) CredParser() (common.Address, error) {
	return _RateModule.Contract.CredParser(&_RateModule.CallOpts)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_RateModule *RateModuleCallerSession) CredParser() (common.Address, error) {
	return _RateModule.Contract.CredParser(&_RateModule.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_RateModule *RateModuleCaller) GetRate(opts *bind.CallOpts, vc VerifiableCredential) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "getRate", vc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_RateModule *RateModuleSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _RateModule.Contract.GetRate(&_RateModule.CallOpts, vc)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_RateModule *RateModuleCallerSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _RateModule.Contract.GetRate(&_RateModule.CallOpts, vc)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_RateModule *RateModuleCaller) IsApproved(opts *bind.CallOpts, account Account, vc VerifiableCredential) (bool, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "isApproved", account, vc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_RateModule *RateModuleSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _RateModule.Contract.IsApproved(&_RateModule.CallOpts, account, vc)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_RateModule *RateModuleCallerSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _RateModule.Contract.IsApproved(&_RateModule.CallOpts, account, vc)
}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCaller) Levels(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "levels", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_RateModule *RateModuleSession) Levels(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.Levels(&_RateModule.CallOpts, arg0)
}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCallerSession) Levels(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.Levels(&_RateModule.CallOpts, arg0)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_RateModule *RateModuleCaller) MaxDTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "maxDTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_RateModule *RateModuleSession) MaxDTE() (*big.Int, error) {
	return _RateModule.Contract.MaxDTE(&_RateModule.CallOpts)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MaxDTE() (*big.Int, error) {
	return _RateModule.Contract.MaxDTE(&_RateModule.CallOpts)
}

// MaxDTI is a free data retrieval call binding the contract method 0xe08fdb55.
//
// Solidity: function maxDTI() view returns(uint256)
func (_RateModule *RateModuleCaller) MaxDTI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "maxDTI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDTI is a free data retrieval call binding the contract method 0xe08fdb55.
//
// Solidity: function maxDTI() view returns(uint256)
func (_RateModule *RateModuleSession) MaxDTI() (*big.Int, error) {
	return _RateModule.Contract.MaxDTI(&_RateModule.CallOpts)
}

// MaxDTI is a free data retrieval call binding the contract method 0xe08fdb55.
//
// Solidity: function maxDTI() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MaxDTI() (*big.Int, error) {
	return _RateModule.Contract.MaxDTI(&_RateModule.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_RateModule *RateModuleCaller) MaxLTV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "maxLTV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_RateModule *RateModuleSession) MaxLTV() (*big.Int, error) {
	return _RateModule.Contract.MaxLTV(&_RateModule.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MaxLTV() (*big.Int, error) {
	return _RateModule.Contract.MaxLTV(&_RateModule.CallOpts)
}

// MinGCRED is a free data retrieval call binding the contract method 0xdb29c170.
//
// Solidity: function minGCRED() view returns(uint256)
func (_RateModule *RateModuleCaller) MinGCRED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "minGCRED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinGCRED is a free data retrieval call binding the contract method 0xdb29c170.
//
// Solidity: function minGCRED() view returns(uint256)
func (_RateModule *RateModuleSession) MinGCRED() (*big.Int, error) {
	return _RateModule.Contract.MinGCRED(&_RateModule.CallOpts)
}

// MinGCRED is a free data retrieval call binding the contract method 0xdb29c170.
//
// Solidity: function minGCRED() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MinGCRED() (*big.Int, error) {
	return _RateModule.Contract.MinGCRED(&_RateModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RateModule *RateModuleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RateModule *RateModuleSession) Owner() (common.Address, error) {
	return _RateModule.Contract.Owner(&_RateModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RateModule *RateModuleCallerSession) Owner() (common.Address, error) {
	return _RateModule.Contract.Owner(&_RateModule.CallOpts)
}

// PenaltyRate is a free data retrieval call binding the contract method 0xd6b7494f.
//
// Solidity: function penaltyRate() view returns(uint256 rate)
func (_RateModule *RateModuleCaller) PenaltyRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "penaltyRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyRate is a free data retrieval call binding the contract method 0xd6b7494f.
//
// Solidity: function penaltyRate() view returns(uint256 rate)
func (_RateModule *RateModuleSession) PenaltyRate() (*big.Int, error) {
	return _RateModule.Contract.PenaltyRate(&_RateModule.CallOpts)
}

// PenaltyRate is a free data retrieval call binding the contract method 0xd6b7494f.
//
// Solidity: function penaltyRate() view returns(uint256 rate)
func (_RateModule *RateModuleCallerSession) PenaltyRate() (*big.Int, error) {
	return _RateModule.Contract.PenaltyRate(&_RateModule.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RateModule *RateModuleCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RateModule *RateModuleSession) PendingOwner() (common.Address, error) {
	return _RateModule.Contract.PendingOwner(&_RateModule.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RateModule *RateModuleCallerSession) PendingOwner() (common.Address, error) {
	return _RateModule.Contract.PendingOwner(&_RateModule.CallOpts)
}

// RateLookup is a free data retrieval call binding the contract method 0xa73292f0.
//
// Solidity: function rateLookup(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCaller) RateLookup(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "rateLookup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateLookup is a free data retrieval call binding the contract method 0xa73292f0.
//
// Solidity: function rateLookup(uint256 ) view returns(uint256)
func (_RateModule *RateModuleSession) RateLookup(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.RateLookup(&_RateModule.CallOpts, arg0)
}

// RateLookup is a free data retrieval call binding the contract method 0xa73292f0.
//
// Solidity: function rateLookup(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCallerSession) RateLookup(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.RateLookup(&_RateModule.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RateModule *RateModuleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RateModule *RateModuleSession) AcceptOwnership() (*types.Transaction, error) {
	return _RateModule.Contract.AcceptOwnership(&_RateModule.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RateModule *RateModuleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RateModule.Contract.AcceptOwnership(&_RateModule.TransactOpts)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_RateModule *RateModuleTransactor) SetAgentLevels(opts *bind.TransactOpts, agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setAgentLevels", agentIDs, level)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_RateModule *RateModuleSession) SetAgentLevels(agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetAgentLevels(&_RateModule.TransactOpts, agentIDs, level)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_RateModule *RateModuleTransactorSession) SetAgentLevels(agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetAgentLevels(&_RateModule.TransactOpts, agentIDs, level)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 _baseRate) returns()
func (_RateModule *RateModuleTransactor) SetBaseRate(opts *bind.TransactOpts, _baseRate *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setBaseRate", _baseRate)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 _baseRate) returns()
func (_RateModule *RateModuleSession) SetBaseRate(_baseRate *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetBaseRate(&_RateModule.TransactOpts, _baseRate)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 _baseRate) returns()
func (_RateModule *RateModuleTransactorSession) SetBaseRate(_baseRate *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetBaseRate(&_RateModule.TransactOpts, _baseRate)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_RateModule *RateModuleTransactor) SetLevels(opts *bind.TransactOpts, _levels [10]*big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setLevels", _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_RateModule *RateModuleSession) SetLevels(_levels [10]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetLevels(&_RateModule.TransactOpts, _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_RateModule *RateModuleTransactorSession) SetLevels(_levels [10]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetLevels(&_RateModule.TransactOpts, _levels)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_RateModule *RateModuleTransactor) SetMaxDTE(opts *bind.TransactOpts, _maxDTE *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMaxDTE", _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_RateModule *RateModuleSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTE(&_RateModule.TransactOpts, _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_RateModule *RateModuleTransactorSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTE(&_RateModule.TransactOpts, _maxDTE)
}

// SetMaxDTI is a paid mutator transaction binding the contract method 0x9915bcad.
//
// Solidity: function setMaxDTI(uint256 _maxDTI) returns()
func (_RateModule *RateModuleTransactor) SetMaxDTI(opts *bind.TransactOpts, _maxDTI *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMaxDTI", _maxDTI)
}

// SetMaxDTI is a paid mutator transaction binding the contract method 0x9915bcad.
//
// Solidity: function setMaxDTI(uint256 _maxDTI) returns()
func (_RateModule *RateModuleSession) SetMaxDTI(_maxDTI *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTI(&_RateModule.TransactOpts, _maxDTI)
}

// SetMaxDTI is a paid mutator transaction binding the contract method 0x9915bcad.
//
// Solidity: function setMaxDTI(uint256 _maxDTI) returns()
func (_RateModule *RateModuleTransactorSession) SetMaxDTI(_maxDTI *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTI(&_RateModule.TransactOpts, _maxDTI)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_RateModule *RateModuleTransactor) SetMaxLTV(opts *bind.TransactOpts, _maxLTV *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMaxLTV", _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_RateModule *RateModuleSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxLTV(&_RateModule.TransactOpts, _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_RateModule *RateModuleTransactorSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxLTV(&_RateModule.TransactOpts, _maxLTV)
}

// SetMinGCRED is a paid mutator transaction binding the contract method 0x509edfb8.
//
// Solidity: function setMinGCRED(uint256 _minGCRED) returns()
func (_RateModule *RateModuleTransactor) SetMinGCRED(opts *bind.TransactOpts, _minGCRED *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMinGCRED", _minGCRED)
}

// SetMinGCRED is a paid mutator transaction binding the contract method 0x509edfb8.
//
// Solidity: function setMinGCRED(uint256 _minGCRED) returns()
func (_RateModule *RateModuleSession) SetMinGCRED(_minGCRED *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMinGCRED(&_RateModule.TransactOpts, _minGCRED)
}

// SetMinGCRED is a paid mutator transaction binding the contract method 0x509edfb8.
//
// Solidity: function setMinGCRED(uint256 _minGCRED) returns()
func (_RateModule *RateModuleTransactorSession) SetMinGCRED(_minGCRED *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMinGCRED(&_RateModule.TransactOpts, _minGCRED)
}

// SetRateLookup is a paid mutator transaction binding the contract method 0x492bff0e.
//
// Solidity: function setRateLookup(uint256[61] _rateLookup) returns()
func (_RateModule *RateModuleTransactor) SetRateLookup(opts *bind.TransactOpts, _rateLookup [61]*big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setRateLookup", _rateLookup)
}

// SetRateLookup is a paid mutator transaction binding the contract method 0x492bff0e.
//
// Solidity: function setRateLookup(uint256[61] _rateLookup) returns()
func (_RateModule *RateModuleSession) SetRateLookup(_rateLookup [61]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetRateLookup(&_RateModule.TransactOpts, _rateLookup)
}

// SetRateLookup is a paid mutator transaction binding the contract method 0x492bff0e.
//
// Solidity: function setRateLookup(uint256[61] _rateLookup) returns()
func (_RateModule *RateModuleTransactorSession) SetRateLookup(_rateLookup [61]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetRateLookup(&_RateModule.TransactOpts, _rateLookup)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RateModule *RateModuleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RateModule *RateModuleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RateModule.Contract.TransferOwnership(&_RateModule.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RateModule *RateModuleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RateModule.Contract.TransferOwnership(&_RateModule.TransactOpts, newOwner)
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_RateModule *RateModuleTransactor) UpdateCredParser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "updateCredParser")
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_RateModule *RateModuleSession) UpdateCredParser() (*types.Transaction, error) {
	return _RateModule.Contract.UpdateCredParser(&_RateModule.TransactOpts)
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_RateModule *RateModuleTransactorSession) UpdateCredParser() (*types.Transaction, error) {
	return _RateModule.Contract.UpdateCredParser(&_RateModule.TransactOpts)
}
