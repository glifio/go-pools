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

// IERC4626MetaData contains all meta data concerning the IERC4626 contract.
var IERC4626MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC4626ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC4626MetaData.ABI instead.
var IERC4626ABI = IERC4626MetaData.ABI

// IERC4626 is an auto generated Go binding around an Ethereum contract.
type IERC4626 struct {
	IERC4626Caller     // Read-only binding to the contract
	IERC4626Transactor // Write-only binding to the contract
	IERC4626Filterer   // Log filterer for contract events
}

// IERC4626Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC4626Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC4626Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC4626Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC4626Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC4626Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC4626Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC4626Session struct {
	Contract     *IERC4626         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC4626CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC4626CallerSession struct {
	Contract *IERC4626Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC4626TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC4626TransactorSession struct {
	Contract     *IERC4626Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC4626Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC4626Raw struct {
	Contract *IERC4626 // Generic contract binding to access the raw methods on
}

// IERC4626CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC4626CallerRaw struct {
	Contract *IERC4626Caller // Generic read-only contract binding to access the raw methods on
}

// IERC4626TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC4626TransactorRaw struct {
	Contract *IERC4626Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC4626 creates a new instance of IERC4626, bound to a specific deployed contract.
func NewIERC4626(address common.Address, backend bind.ContractBackend) (*IERC4626, error) {
	contract, err := bindIERC4626(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC4626{IERC4626Caller: IERC4626Caller{contract: contract}, IERC4626Transactor: IERC4626Transactor{contract: contract}, IERC4626Filterer: IERC4626Filterer{contract: contract}}, nil
}

// NewIERC4626Caller creates a new read-only instance of IERC4626, bound to a specific deployed contract.
func NewIERC4626Caller(address common.Address, caller bind.ContractCaller) (*IERC4626Caller, error) {
	contract, err := bindIERC4626(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC4626Caller{contract: contract}, nil
}

// NewIERC4626Transactor creates a new write-only instance of IERC4626, bound to a specific deployed contract.
func NewIERC4626Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC4626Transactor, error) {
	contract, err := bindIERC4626(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC4626Transactor{contract: contract}, nil
}

// NewIERC4626Filterer creates a new log filterer instance of IERC4626, bound to a specific deployed contract.
func NewIERC4626Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC4626Filterer, error) {
	contract, err := bindIERC4626(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC4626Filterer{contract: contract}, nil
}

// bindIERC4626 binds a generic wrapper to an already deployed contract.
func bindIERC4626(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC4626MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC4626 *IERC4626Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC4626.Contract.IERC4626Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC4626 *IERC4626Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC4626.Contract.IERC4626Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC4626 *IERC4626Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC4626.Contract.IERC4626Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC4626 *IERC4626CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC4626.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC4626 *IERC4626TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC4626.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC4626 *IERC4626TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC4626.Contract.contract.Transact(opts, method, params...)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_IERC4626 *IERC4626Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_IERC4626 *IERC4626Session) Asset() (common.Address, error) {
	return _IERC4626.Contract.Asset(&_IERC4626.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_IERC4626 *IERC4626CallerSession) Asset() (common.Address, error) {
	return _IERC4626.Contract.Asset(&_IERC4626.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_IERC4626 *IERC4626Caller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_IERC4626 *IERC4626Session) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.ConvertToAssets(&_IERC4626.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.ConvertToAssets(&_IERC4626.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626Caller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626Session) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.ConvertToShares(&_IERC4626.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.ConvertToShares(&_IERC4626.CallOpts, assets)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_IERC4626 *IERC4626Caller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_IERC4626 *IERC4626Session) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxDeposit(&_IERC4626.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxDeposit(&_IERC4626.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_IERC4626 *IERC4626Caller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_IERC4626 *IERC4626Session) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxMint(&_IERC4626.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxMint(&_IERC4626.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_IERC4626 *IERC4626Caller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_IERC4626 *IERC4626Session) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxRedeem(&_IERC4626.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxRedeem(&_IERC4626.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_IERC4626 *IERC4626Caller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_IERC4626 *IERC4626Session) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxWithdraw(&_IERC4626.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _IERC4626.Contract.MaxWithdraw(&_IERC4626.CallOpts, owner)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626Caller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626Session) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.PreviewDeposit(&_IERC4626.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.PreviewDeposit(&_IERC4626.CallOpts, assets)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_IERC4626 *IERC4626Caller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_IERC4626 *IERC4626Session) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.PreviewRedeem(&_IERC4626.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.PreviewRedeem(&_IERC4626.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626Caller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626Session) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.PreviewWithdraw(&_IERC4626.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _IERC4626.Contract.PreviewWithdraw(&_IERC4626.CallOpts, assets)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_IERC4626 *IERC4626Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC4626.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_IERC4626 *IERC4626Session) TotalAssets() (*big.Int, error) {
	return _IERC4626.Contract.TotalAssets(&_IERC4626.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_IERC4626 *IERC4626CallerSession) TotalAssets() (*big.Int, error) {
	return _IERC4626.Contract.TotalAssets(&_IERC4626.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_IERC4626 *IERC4626Transactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _IERC4626.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_IERC4626 *IERC4626Session) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Deposit(&_IERC4626.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_IERC4626 *IERC4626TransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Deposit(&_IERC4626.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_IERC4626 *IERC4626Transactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _IERC4626.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_IERC4626 *IERC4626Session) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Mint(&_IERC4626.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_IERC4626 *IERC4626TransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Mint(&_IERC4626.TransactOpts, shares, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_IERC4626 *IERC4626Transactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _IERC4626.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_IERC4626 *IERC4626Session) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Redeem(&_IERC4626.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_IERC4626 *IERC4626TransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Redeem(&_IERC4626.TransactOpts, shares, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_IERC4626 *IERC4626Transactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _IERC4626.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_IERC4626 *IERC4626Session) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Withdraw(&_IERC4626.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_IERC4626 *IERC4626TransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _IERC4626.Contract.Withdraw(&_IERC4626.TransactOpts, assets, receiver, owner)
}
