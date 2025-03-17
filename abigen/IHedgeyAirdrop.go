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

// IHedgeyAirdropCampaign is an auto generated low-level Go binding around an user-defined struct.
type IHedgeyAirdropCampaign struct {
	Manager     common.Address
	Token       common.Address
	Amount      *big.Int
	Start       *big.Int
	End         *big.Int
	TokenLockup uint8
	Root        [32]byte
	Delegating  bool
}

// IHedgeyAirdropClaimLockup is an auto generated low-level Go binding around an user-defined struct.
type IHedgeyAirdropClaimLockup struct {
	TokenLocker common.Address
	Start       *big.Int
	Cliff       *big.Int
	Period      *big.Int
	Periods     *big.Int
}

// IHedgeyAirdropSignatureParams is an auto generated low-level Go binding around an user-defined struct.
type IHedgeyAirdropSignatureParams struct {
	Nonce  *big.Int
	Expiry *big.Int
	V      uint8
	R      [32]byte
	S      [32]byte
}

// IHedgeyAirdropMetaData contains all meta data concerning the IHedgeyAirdrop contract.
var IHedgeyAirdropMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"CampaignCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"enumIHedgeyAirdrop.TokenLockup\",\"name\":\"tokenLockup\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"delegating\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIHedgeyAirdrop.Campaign\",\"name\":\"campaign\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalClaimers\",\"type\":\"uint256\"}],\"name\":\"CampaignStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenLocker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periods\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIHedgeyAirdrop.ClaimLockup\",\"name\":\"claimLockup\",\"type\":\"tuple\"}],\"name\":\"ClaimLockupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRemaining\",\"type\":\"uint256\"}],\"name\":\"LockedTokensClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRemaining\",\"type\":\"uint256\"}],\"name\":\"UnlockedTokensClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"campaigns\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"enumIHedgeyAirdrop.TokenLockup\",\"name\":\"tokenLockup\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"delegating\",\"type\":\"bool\"}],\"internalType\":\"structIHedgeyAirdrop.Campaign\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"campaignIds\",\"type\":\"bytes16[]\"}],\"name\":\"cancelCampaigns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"campaignId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"campaignId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIHedgeyAirdrop.SignatureParams\",\"name\":\"delegationSignature\",\"type\":\"tuple\"}],\"name\":\"claimAndDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"campaignId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIHedgeyAirdrop.SignatureParams\",\"name\":\"claimSignature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIHedgeyAirdrop.SignatureParams\",\"name\":\"delegationSignature\",\"type\":\"tuple\"}],\"name\":\"claimAndDelegateWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"claimLockups\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenLocker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periods\",\"type\":\"uint256\"}],\"internalType\":\"structIHedgeyAirdrop.ClaimLockup\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"campaignIds\",\"type\":\"bytes16[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"proofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimAmounts\",\"type\":\"uint256[]\"}],\"name\":\"claimMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"campaignIds\",\"type\":\"bytes16[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"proofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"claimAmounts\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIHedgeyAirdrop.SignatureParams\",\"name\":\"claimSignature\",\"type\":\"tuple\"}],\"name\":\"claimMultipleWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"campaignId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIHedgeyAirdrop.SignatureParams\",\"name\":\"claimSignature\",\"type\":\"tuple\"}],\"name\":\"claimWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"enumIHedgeyAirdrop.TokenLockup\",\"name\":\"tokenLockup\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"delegating\",\"type\":\"bool\"}],\"internalType\":\"structIHedgeyAirdrop.Campaign\",\"name\":\"campaign\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenLocker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periods\",\"type\":\"uint256\"}],\"internalType\":\"structIHedgeyAirdrop.ClaimLockup\",\"name\":\"claimLockup\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"vestingAdmin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimers\",\"type\":\"uint256\"}],\"name\":\"createLockedCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"enumIHedgeyAirdrop.TokenLockup\",\"name\":\"tokenLockup\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"delegating\",\"type\":\"bool\"}],\"internalType\":\"structIHedgeyAirdrop.Campaign\",\"name\":\"campaign\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimers\",\"type\":\"uint256\"}],\"name\":\"createUnlockedCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"usedIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// IHedgeyAirdropABI is the input ABI used to generate the binding from.
// Deprecated: Use IHedgeyAirdropMetaData.ABI instead.
var IHedgeyAirdropABI = IHedgeyAirdropMetaData.ABI

// IHedgeyAirdrop is an auto generated Go binding around an Ethereum contract.
type IHedgeyAirdrop struct {
	IHedgeyAirdropCaller     // Read-only binding to the contract
	IHedgeyAirdropTransactor // Write-only binding to the contract
	IHedgeyAirdropFilterer   // Log filterer for contract events
}

// IHedgeyAirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHedgeyAirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyAirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHedgeyAirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyAirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHedgeyAirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyAirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHedgeyAirdropSession struct {
	Contract     *IHedgeyAirdrop   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHedgeyAirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHedgeyAirdropCallerSession struct {
	Contract *IHedgeyAirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IHedgeyAirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHedgeyAirdropTransactorSession struct {
	Contract     *IHedgeyAirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IHedgeyAirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHedgeyAirdropRaw struct {
	Contract *IHedgeyAirdrop // Generic contract binding to access the raw methods on
}

// IHedgeyAirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHedgeyAirdropCallerRaw struct {
	Contract *IHedgeyAirdropCaller // Generic read-only contract binding to access the raw methods on
}

// IHedgeyAirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHedgeyAirdropTransactorRaw struct {
	Contract *IHedgeyAirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHedgeyAirdrop creates a new instance of IHedgeyAirdrop, bound to a specific deployed contract.
func NewIHedgeyAirdrop(address common.Address, backend bind.ContractBackend) (*IHedgeyAirdrop, error) {
	contract, err := bindIHedgeyAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdrop{IHedgeyAirdropCaller: IHedgeyAirdropCaller{contract: contract}, IHedgeyAirdropTransactor: IHedgeyAirdropTransactor{contract: contract}, IHedgeyAirdropFilterer: IHedgeyAirdropFilterer{contract: contract}}, nil
}

// NewIHedgeyAirdropCaller creates a new read-only instance of IHedgeyAirdrop, bound to a specific deployed contract.
func NewIHedgeyAirdropCaller(address common.Address, caller bind.ContractCaller) (*IHedgeyAirdropCaller, error) {
	contract, err := bindIHedgeyAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropCaller{contract: contract}, nil
}

// NewIHedgeyAirdropTransactor creates a new write-only instance of IHedgeyAirdrop, bound to a specific deployed contract.
func NewIHedgeyAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*IHedgeyAirdropTransactor, error) {
	contract, err := bindIHedgeyAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropTransactor{contract: contract}, nil
}

// NewIHedgeyAirdropFilterer creates a new log filterer instance of IHedgeyAirdrop, bound to a specific deployed contract.
func NewIHedgeyAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*IHedgeyAirdropFilterer, error) {
	contract, err := bindIHedgeyAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropFilterer{contract: contract}, nil
}

// bindIHedgeyAirdrop binds a generic wrapper to an already deployed contract.
func bindIHedgeyAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHedgeyAirdropMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHedgeyAirdrop *IHedgeyAirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHedgeyAirdrop.Contract.IHedgeyAirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHedgeyAirdrop *IHedgeyAirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.IHedgeyAirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHedgeyAirdrop *IHedgeyAirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.IHedgeyAirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHedgeyAirdrop *IHedgeyAirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHedgeyAirdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.contract.Transact(opts, method, params...)
}

// Campaigns is a free data retrieval call binding the contract method 0x50b4aebe.
//
// Solidity: function campaigns(bytes16 id) view returns((address,address,uint256,uint256,uint256,uint8,bytes32,bool))
func (_IHedgeyAirdrop *IHedgeyAirdropCaller) Campaigns(opts *bind.CallOpts, id [16]byte) (IHedgeyAirdropCampaign, error) {
	var out []interface{}
	err := _IHedgeyAirdrop.contract.Call(opts, &out, "campaigns", id)

	if err != nil {
		return *new(IHedgeyAirdropCampaign), err
	}

	out0 := *abi.ConvertType(out[0], new(IHedgeyAirdropCampaign)).(*IHedgeyAirdropCampaign)

	return out0, err

}

// Campaigns is a free data retrieval call binding the contract method 0x50b4aebe.
//
// Solidity: function campaigns(bytes16 id) view returns((address,address,uint256,uint256,uint256,uint8,bytes32,bool))
func (_IHedgeyAirdrop *IHedgeyAirdropSession) Campaigns(id [16]byte) (IHedgeyAirdropCampaign, error) {
	return _IHedgeyAirdrop.Contract.Campaigns(&_IHedgeyAirdrop.CallOpts, id)
}

// Campaigns is a free data retrieval call binding the contract method 0x50b4aebe.
//
// Solidity: function campaigns(bytes16 id) view returns((address,address,uint256,uint256,uint256,uint8,bytes32,bool))
func (_IHedgeyAirdrop *IHedgeyAirdropCallerSession) Campaigns(id [16]byte) (IHedgeyAirdropCampaign, error) {
	return _IHedgeyAirdrop.Contract.Campaigns(&_IHedgeyAirdrop.CallOpts, id)
}

// ClaimLockups is a free data retrieval call binding the contract method 0xd6eecb00.
//
// Solidity: function claimLockups(bytes16 id) view returns((address,uint256,uint256,uint256,uint256))
func (_IHedgeyAirdrop *IHedgeyAirdropCaller) ClaimLockups(opts *bind.CallOpts, id [16]byte) (IHedgeyAirdropClaimLockup, error) {
	var out []interface{}
	err := _IHedgeyAirdrop.contract.Call(opts, &out, "claimLockups", id)

	if err != nil {
		return *new(IHedgeyAirdropClaimLockup), err
	}

	out0 := *abi.ConvertType(out[0], new(IHedgeyAirdropClaimLockup)).(*IHedgeyAirdropClaimLockup)

	return out0, err

}

// ClaimLockups is a free data retrieval call binding the contract method 0xd6eecb00.
//
// Solidity: function claimLockups(bytes16 id) view returns((address,uint256,uint256,uint256,uint256))
func (_IHedgeyAirdrop *IHedgeyAirdropSession) ClaimLockups(id [16]byte) (IHedgeyAirdropClaimLockup, error) {
	return _IHedgeyAirdrop.Contract.ClaimLockups(&_IHedgeyAirdrop.CallOpts, id)
}

// ClaimLockups is a free data retrieval call binding the contract method 0xd6eecb00.
//
// Solidity: function claimLockups(bytes16 id) view returns((address,uint256,uint256,uint256,uint256))
func (_IHedgeyAirdrop *IHedgeyAirdropCallerSession) ClaimLockups(id [16]byte) (IHedgeyAirdropClaimLockup, error) {
	return _IHedgeyAirdrop.Contract.ClaimLockups(&_IHedgeyAirdrop.CallOpts, id)
}

// UsedIds is a free data retrieval call binding the contract method 0xbcecbd9d.
//
// Solidity: function usedIds(bytes16 id) view returns(bool)
func (_IHedgeyAirdrop *IHedgeyAirdropCaller) UsedIds(opts *bind.CallOpts, id [16]byte) (bool, error) {
	var out []interface{}
	err := _IHedgeyAirdrop.contract.Call(opts, &out, "usedIds", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedIds is a free data retrieval call binding the contract method 0xbcecbd9d.
//
// Solidity: function usedIds(bytes16 id) view returns(bool)
func (_IHedgeyAirdrop *IHedgeyAirdropSession) UsedIds(id [16]byte) (bool, error) {
	return _IHedgeyAirdrop.Contract.UsedIds(&_IHedgeyAirdrop.CallOpts, id)
}

// UsedIds is a free data retrieval call binding the contract method 0xbcecbd9d.
//
// Solidity: function usedIds(bytes16 id) view returns(bool)
func (_IHedgeyAirdrop *IHedgeyAirdropCallerSession) UsedIds(id [16]byte) (bool, error) {
	return _IHedgeyAirdrop.Contract.UsedIds(&_IHedgeyAirdrop.CallOpts, id)
}

// Verify is a free data retrieval call binding the contract method 0xcff144f7.
//
// Solidity: function verify(bytes32 root, bytes32[] proof, address claimer, uint256 amount) pure returns(bool)
func (_IHedgeyAirdrop *IHedgeyAirdropCaller) Verify(opts *bind.CallOpts, root [32]byte, proof [][32]byte, claimer common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _IHedgeyAirdrop.contract.Call(opts, &out, "verify", root, proof, claimer, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xcff144f7.
//
// Solidity: function verify(bytes32 root, bytes32[] proof, address claimer, uint256 amount) pure returns(bool)
func (_IHedgeyAirdrop *IHedgeyAirdropSession) Verify(root [32]byte, proof [][32]byte, claimer common.Address, amount *big.Int) (bool, error) {
	return _IHedgeyAirdrop.Contract.Verify(&_IHedgeyAirdrop.CallOpts, root, proof, claimer, amount)
}

// Verify is a free data retrieval call binding the contract method 0xcff144f7.
//
// Solidity: function verify(bytes32 root, bytes32[] proof, address claimer, uint256 amount) pure returns(bool)
func (_IHedgeyAirdrop *IHedgeyAirdropCallerSession) Verify(root [32]byte, proof [][32]byte, claimer common.Address, amount *big.Int) (bool, error) {
	return _IHedgeyAirdrop.Contract.Verify(&_IHedgeyAirdrop.CallOpts, root, proof, claimer, amount)
}

// CancelCampaigns is a paid mutator transaction binding the contract method 0x857c2c31.
//
// Solidity: function cancelCampaigns(bytes16[] campaignIds) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) CancelCampaigns(opts *bind.TransactOpts, campaignIds [][16]byte) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "cancelCampaigns", campaignIds)
}

// CancelCampaigns is a paid mutator transaction binding the contract method 0x857c2c31.
//
// Solidity: function cancelCampaigns(bytes16[] campaignIds) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) CancelCampaigns(campaignIds [][16]byte) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.CancelCampaigns(&_IHedgeyAirdrop.TransactOpts, campaignIds)
}

// CancelCampaigns is a paid mutator transaction binding the contract method 0x857c2c31.
//
// Solidity: function cancelCampaigns(bytes16[] campaignIds) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) CancelCampaigns(campaignIds [][16]byte) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.CancelCampaigns(&_IHedgeyAirdrop.TransactOpts, campaignIds)
}

// Claim is a paid mutator transaction binding the contract method 0x891f4e42.
//
// Solidity: function claim(bytes16 campaignId, bytes32[] proof, uint256 claimAmount) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) Claim(opts *bind.TransactOpts, campaignId [16]byte, proof [][32]byte, claimAmount *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "claim", campaignId, proof, claimAmount)
}

// Claim is a paid mutator transaction binding the contract method 0x891f4e42.
//
// Solidity: function claim(bytes16 campaignId, bytes32[] proof, uint256 claimAmount) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) Claim(campaignId [16]byte, proof [][32]byte, claimAmount *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.Claim(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimAmount)
}

// Claim is a paid mutator transaction binding the contract method 0x891f4e42.
//
// Solidity: function claim(bytes16 campaignId, bytes32[] proof, uint256 claimAmount) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) Claim(campaignId [16]byte, proof [][32]byte, claimAmount *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.Claim(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimAmount)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x4a3067b9.
//
// Solidity: function claimAndDelegate(bytes16 campaignId, bytes32[] proof, uint256 claimAmount, address delegatee, (uint256,uint256,uint8,bytes32,bytes32) delegationSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) ClaimAndDelegate(opts *bind.TransactOpts, campaignId [16]byte, proof [][32]byte, claimAmount *big.Int, delegatee common.Address, delegationSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "claimAndDelegate", campaignId, proof, claimAmount, delegatee, delegationSignature)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x4a3067b9.
//
// Solidity: function claimAndDelegate(bytes16 campaignId, bytes32[] proof, uint256 claimAmount, address delegatee, (uint256,uint256,uint8,bytes32,bytes32) delegationSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) ClaimAndDelegate(campaignId [16]byte, proof [][32]byte, claimAmount *big.Int, delegatee common.Address, delegationSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimAndDelegate(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimAmount, delegatee, delegationSignature)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x4a3067b9.
//
// Solidity: function claimAndDelegate(bytes16 campaignId, bytes32[] proof, uint256 claimAmount, address delegatee, (uint256,uint256,uint8,bytes32,bytes32) delegationSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) ClaimAndDelegate(campaignId [16]byte, proof [][32]byte, claimAmount *big.Int, delegatee common.Address, delegationSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimAndDelegate(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimAmount, delegatee, delegationSignature)
}

// ClaimAndDelegateWithSig is a paid mutator transaction binding the contract method 0xf0e1a347.
//
// Solidity: function claimAndDelegateWithSig(bytes16 campaignId, bytes32[] proof, address claimer, uint256 claimAmount, (uint256,uint256,uint8,bytes32,bytes32) claimSignature, address delegatee, (uint256,uint256,uint8,bytes32,bytes32) delegationSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) ClaimAndDelegateWithSig(opts *bind.TransactOpts, campaignId [16]byte, proof [][32]byte, claimer common.Address, claimAmount *big.Int, claimSignature IHedgeyAirdropSignatureParams, delegatee common.Address, delegationSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "claimAndDelegateWithSig", campaignId, proof, claimer, claimAmount, claimSignature, delegatee, delegationSignature)
}

// ClaimAndDelegateWithSig is a paid mutator transaction binding the contract method 0xf0e1a347.
//
// Solidity: function claimAndDelegateWithSig(bytes16 campaignId, bytes32[] proof, address claimer, uint256 claimAmount, (uint256,uint256,uint8,bytes32,bytes32) claimSignature, address delegatee, (uint256,uint256,uint8,bytes32,bytes32) delegationSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) ClaimAndDelegateWithSig(campaignId [16]byte, proof [][32]byte, claimer common.Address, claimAmount *big.Int, claimSignature IHedgeyAirdropSignatureParams, delegatee common.Address, delegationSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimAndDelegateWithSig(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimer, claimAmount, claimSignature, delegatee, delegationSignature)
}

// ClaimAndDelegateWithSig is a paid mutator transaction binding the contract method 0xf0e1a347.
//
// Solidity: function claimAndDelegateWithSig(bytes16 campaignId, bytes32[] proof, address claimer, uint256 claimAmount, (uint256,uint256,uint8,bytes32,bytes32) claimSignature, address delegatee, (uint256,uint256,uint8,bytes32,bytes32) delegationSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) ClaimAndDelegateWithSig(campaignId [16]byte, proof [][32]byte, claimer common.Address, claimAmount *big.Int, claimSignature IHedgeyAirdropSignatureParams, delegatee common.Address, delegationSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimAndDelegateWithSig(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimer, claimAmount, claimSignature, delegatee, delegationSignature)
}

// ClaimMultiple is a paid mutator transaction binding the contract method 0xd4b52e0b.
//
// Solidity: function claimMultiple(bytes16[] campaignIds, bytes32[][] proofs, uint256[] claimAmounts) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) ClaimMultiple(opts *bind.TransactOpts, campaignIds [][16]byte, proofs [][][32]byte, claimAmounts []*big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "claimMultiple", campaignIds, proofs, claimAmounts)
}

// ClaimMultiple is a paid mutator transaction binding the contract method 0xd4b52e0b.
//
// Solidity: function claimMultiple(bytes16[] campaignIds, bytes32[][] proofs, uint256[] claimAmounts) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) ClaimMultiple(campaignIds [][16]byte, proofs [][][32]byte, claimAmounts []*big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimMultiple(&_IHedgeyAirdrop.TransactOpts, campaignIds, proofs, claimAmounts)
}

// ClaimMultiple is a paid mutator transaction binding the contract method 0xd4b52e0b.
//
// Solidity: function claimMultiple(bytes16[] campaignIds, bytes32[][] proofs, uint256[] claimAmounts) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) ClaimMultiple(campaignIds [][16]byte, proofs [][][32]byte, claimAmounts []*big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimMultiple(&_IHedgeyAirdrop.TransactOpts, campaignIds, proofs, claimAmounts)
}

// ClaimMultipleWithSig is a paid mutator transaction binding the contract method 0x01a36e74.
//
// Solidity: function claimMultipleWithSig(bytes16[] campaignIds, bytes32[][] proofs, address claimer, uint256[] claimAmounts, (uint256,uint256,uint8,bytes32,bytes32) claimSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) ClaimMultipleWithSig(opts *bind.TransactOpts, campaignIds [][16]byte, proofs [][][32]byte, claimer common.Address, claimAmounts []*big.Int, claimSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "claimMultipleWithSig", campaignIds, proofs, claimer, claimAmounts, claimSignature)
}

// ClaimMultipleWithSig is a paid mutator transaction binding the contract method 0x01a36e74.
//
// Solidity: function claimMultipleWithSig(bytes16[] campaignIds, bytes32[][] proofs, address claimer, uint256[] claimAmounts, (uint256,uint256,uint8,bytes32,bytes32) claimSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) ClaimMultipleWithSig(campaignIds [][16]byte, proofs [][][32]byte, claimer common.Address, claimAmounts []*big.Int, claimSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimMultipleWithSig(&_IHedgeyAirdrop.TransactOpts, campaignIds, proofs, claimer, claimAmounts, claimSignature)
}

// ClaimMultipleWithSig is a paid mutator transaction binding the contract method 0x01a36e74.
//
// Solidity: function claimMultipleWithSig(bytes16[] campaignIds, bytes32[][] proofs, address claimer, uint256[] claimAmounts, (uint256,uint256,uint8,bytes32,bytes32) claimSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) ClaimMultipleWithSig(campaignIds [][16]byte, proofs [][][32]byte, claimer common.Address, claimAmounts []*big.Int, claimSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimMultipleWithSig(&_IHedgeyAirdrop.TransactOpts, campaignIds, proofs, claimer, claimAmounts, claimSignature)
}

// ClaimWithSig is a paid mutator transaction binding the contract method 0xd69564fb.
//
// Solidity: function claimWithSig(bytes16 campaignId, bytes32[] proof, address claimer, uint256 claimAmount, (uint256,uint256,uint8,bytes32,bytes32) claimSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) ClaimWithSig(opts *bind.TransactOpts, campaignId [16]byte, proof [][32]byte, claimer common.Address, claimAmount *big.Int, claimSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "claimWithSig", campaignId, proof, claimer, claimAmount, claimSignature)
}

// ClaimWithSig is a paid mutator transaction binding the contract method 0xd69564fb.
//
// Solidity: function claimWithSig(bytes16 campaignId, bytes32[] proof, address claimer, uint256 claimAmount, (uint256,uint256,uint8,bytes32,bytes32) claimSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) ClaimWithSig(campaignId [16]byte, proof [][32]byte, claimer common.Address, claimAmount *big.Int, claimSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimWithSig(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimer, claimAmount, claimSignature)
}

// ClaimWithSig is a paid mutator transaction binding the contract method 0xd69564fb.
//
// Solidity: function claimWithSig(bytes16 campaignId, bytes32[] proof, address claimer, uint256 claimAmount, (uint256,uint256,uint8,bytes32,bytes32) claimSignature) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) ClaimWithSig(campaignId [16]byte, proof [][32]byte, claimer common.Address, claimAmount *big.Int, claimSignature IHedgeyAirdropSignatureParams) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.ClaimWithSig(&_IHedgeyAirdrop.TransactOpts, campaignId, proof, claimer, claimAmount, claimSignature)
}

// CreateLockedCampaign is a paid mutator transaction binding the contract method 0x98a87e7f.
//
// Solidity: function createLockedCampaign(bytes16 id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, (address,uint256,uint256,uint256,uint256) claimLockup, address vestingAdmin, uint256 totalClaimers) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) CreateLockedCampaign(opts *bind.TransactOpts, id [16]byte, campaign IHedgeyAirdropCampaign, claimLockup IHedgeyAirdropClaimLockup, vestingAdmin common.Address, totalClaimers *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "createLockedCampaign", id, campaign, claimLockup, vestingAdmin, totalClaimers)
}

// CreateLockedCampaign is a paid mutator transaction binding the contract method 0x98a87e7f.
//
// Solidity: function createLockedCampaign(bytes16 id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, (address,uint256,uint256,uint256,uint256) claimLockup, address vestingAdmin, uint256 totalClaimers) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) CreateLockedCampaign(id [16]byte, campaign IHedgeyAirdropCampaign, claimLockup IHedgeyAirdropClaimLockup, vestingAdmin common.Address, totalClaimers *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.CreateLockedCampaign(&_IHedgeyAirdrop.TransactOpts, id, campaign, claimLockup, vestingAdmin, totalClaimers)
}

// CreateLockedCampaign is a paid mutator transaction binding the contract method 0x98a87e7f.
//
// Solidity: function createLockedCampaign(bytes16 id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, (address,uint256,uint256,uint256,uint256) claimLockup, address vestingAdmin, uint256 totalClaimers) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) CreateLockedCampaign(id [16]byte, campaign IHedgeyAirdropCampaign, claimLockup IHedgeyAirdropClaimLockup, vestingAdmin common.Address, totalClaimers *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.CreateLockedCampaign(&_IHedgeyAirdrop.TransactOpts, id, campaign, claimLockup, vestingAdmin, totalClaimers)
}

// CreateUnlockedCampaign is a paid mutator transaction binding the contract method 0x215506d5.
//
// Solidity: function createUnlockedCampaign(bytes16 id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, uint256 totalClaimers) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactor) CreateUnlockedCampaign(opts *bind.TransactOpts, id [16]byte, campaign IHedgeyAirdropCampaign, totalClaimers *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.contract.Transact(opts, "createUnlockedCampaign", id, campaign, totalClaimers)
}

// CreateUnlockedCampaign is a paid mutator transaction binding the contract method 0x215506d5.
//
// Solidity: function createUnlockedCampaign(bytes16 id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, uint256 totalClaimers) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropSession) CreateUnlockedCampaign(id [16]byte, campaign IHedgeyAirdropCampaign, totalClaimers *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.CreateUnlockedCampaign(&_IHedgeyAirdrop.TransactOpts, id, campaign, totalClaimers)
}

// CreateUnlockedCampaign is a paid mutator transaction binding the contract method 0x215506d5.
//
// Solidity: function createUnlockedCampaign(bytes16 id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, uint256 totalClaimers) returns()
func (_IHedgeyAirdrop *IHedgeyAirdropTransactorSession) CreateUnlockedCampaign(id [16]byte, campaign IHedgeyAirdropCampaign, totalClaimers *big.Int) (*types.Transaction, error) {
	return _IHedgeyAirdrop.Contract.CreateUnlockedCampaign(&_IHedgeyAirdrop.TransactOpts, id, campaign, totalClaimers)
}

// IHedgeyAirdropCampaignCancelledIterator is returned from FilterCampaignCancelled and is used to iterate over the raw logs and unpacked data for CampaignCancelled events raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropCampaignCancelledIterator struct {
	Event *IHedgeyAirdropCampaignCancelled // Event containing the contract specifics and raw log

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
func (it *IHedgeyAirdropCampaignCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyAirdropCampaignCancelled)
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
		it.Event = new(IHedgeyAirdropCampaignCancelled)
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
func (it *IHedgeyAirdropCampaignCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyAirdropCampaignCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyAirdropCampaignCancelled represents a CampaignCancelled event raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropCampaignCancelled struct {
	Id  [16]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCampaignCancelled is a free log retrieval operation binding the contract event 0x88511e0182a386da19fdcadcde158de4729b63240d9a034ce75f3cb0dc5288a5.
//
// Solidity: event CampaignCancelled(bytes16 indexed id)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) FilterCampaignCancelled(opts *bind.FilterOpts, id [][16]byte) (*IHedgeyAirdropCampaignCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.FilterLogs(opts, "CampaignCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropCampaignCancelledIterator{contract: _IHedgeyAirdrop.contract, event: "CampaignCancelled", logs: logs, sub: sub}, nil
}

// WatchCampaignCancelled is a free log subscription operation binding the contract event 0x88511e0182a386da19fdcadcde158de4729b63240d9a034ce75f3cb0dc5288a5.
//
// Solidity: event CampaignCancelled(bytes16 indexed id)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) WatchCampaignCancelled(opts *bind.WatchOpts, sink chan<- *IHedgeyAirdropCampaignCancelled, id [][16]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.WatchLogs(opts, "CampaignCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyAirdropCampaignCancelled)
				if err := _IHedgeyAirdrop.contract.UnpackLog(event, "CampaignCancelled", log); err != nil {
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

// ParseCampaignCancelled is a log parse operation binding the contract event 0x88511e0182a386da19fdcadcde158de4729b63240d9a034ce75f3cb0dc5288a5.
//
// Solidity: event CampaignCancelled(bytes16 indexed id)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) ParseCampaignCancelled(log types.Log) (*IHedgeyAirdropCampaignCancelled, error) {
	event := new(IHedgeyAirdropCampaignCancelled)
	if err := _IHedgeyAirdrop.contract.UnpackLog(event, "CampaignCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyAirdropCampaignStartedIterator is returned from FilterCampaignStarted and is used to iterate over the raw logs and unpacked data for CampaignStarted events raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropCampaignStartedIterator struct {
	Event *IHedgeyAirdropCampaignStarted // Event containing the contract specifics and raw log

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
func (it *IHedgeyAirdropCampaignStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyAirdropCampaignStarted)
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
		it.Event = new(IHedgeyAirdropCampaignStarted)
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
func (it *IHedgeyAirdropCampaignStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyAirdropCampaignStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyAirdropCampaignStarted represents a CampaignStarted event raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropCampaignStarted struct {
	Id            [16]byte
	Campaign      IHedgeyAirdropCampaign
	TotalClaimers *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCampaignStarted is a free log retrieval operation binding the contract event 0xc1d41f85682afff3c8291bb0c88f077fa6a77781d61c4941bd0344a6b780db9c.
//
// Solidity: event CampaignStarted(bytes16 indexed id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, uint256 totalClaimers)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) FilterCampaignStarted(opts *bind.FilterOpts, id [][16]byte) (*IHedgeyAirdropCampaignStartedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.FilterLogs(opts, "CampaignStarted", idRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropCampaignStartedIterator{contract: _IHedgeyAirdrop.contract, event: "CampaignStarted", logs: logs, sub: sub}, nil
}

// WatchCampaignStarted is a free log subscription operation binding the contract event 0xc1d41f85682afff3c8291bb0c88f077fa6a77781d61c4941bd0344a6b780db9c.
//
// Solidity: event CampaignStarted(bytes16 indexed id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, uint256 totalClaimers)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) WatchCampaignStarted(opts *bind.WatchOpts, sink chan<- *IHedgeyAirdropCampaignStarted, id [][16]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.WatchLogs(opts, "CampaignStarted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyAirdropCampaignStarted)
				if err := _IHedgeyAirdrop.contract.UnpackLog(event, "CampaignStarted", log); err != nil {
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

// ParseCampaignStarted is a log parse operation binding the contract event 0xc1d41f85682afff3c8291bb0c88f077fa6a77781d61c4941bd0344a6b780db9c.
//
// Solidity: event CampaignStarted(bytes16 indexed id, (address,address,uint256,uint256,uint256,uint8,bytes32,bool) campaign, uint256 totalClaimers)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) ParseCampaignStarted(log types.Log) (*IHedgeyAirdropCampaignStarted, error) {
	event := new(IHedgeyAirdropCampaignStarted)
	if err := _IHedgeyAirdrop.contract.UnpackLog(event, "CampaignStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyAirdropClaimLockupCreatedIterator is returned from FilterClaimLockupCreated and is used to iterate over the raw logs and unpacked data for ClaimLockupCreated events raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropClaimLockupCreatedIterator struct {
	Event *IHedgeyAirdropClaimLockupCreated // Event containing the contract specifics and raw log

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
func (it *IHedgeyAirdropClaimLockupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyAirdropClaimLockupCreated)
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
		it.Event = new(IHedgeyAirdropClaimLockupCreated)
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
func (it *IHedgeyAirdropClaimLockupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyAirdropClaimLockupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyAirdropClaimLockupCreated represents a ClaimLockupCreated event raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropClaimLockupCreated struct {
	Id          [16]byte
	ClaimLockup IHedgeyAirdropClaimLockup
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimLockupCreated is a free log retrieval operation binding the contract event 0xce04f3b9b2dfc43b1d1e78376db74a3b32c0a78daaf34efc81334908359448fc.
//
// Solidity: event ClaimLockupCreated(bytes16 indexed id, (address,uint256,uint256,uint256,uint256) claimLockup)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) FilterClaimLockupCreated(opts *bind.FilterOpts, id [][16]byte) (*IHedgeyAirdropClaimLockupCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.FilterLogs(opts, "ClaimLockupCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropClaimLockupCreatedIterator{contract: _IHedgeyAirdrop.contract, event: "ClaimLockupCreated", logs: logs, sub: sub}, nil
}

// WatchClaimLockupCreated is a free log subscription operation binding the contract event 0xce04f3b9b2dfc43b1d1e78376db74a3b32c0a78daaf34efc81334908359448fc.
//
// Solidity: event ClaimLockupCreated(bytes16 indexed id, (address,uint256,uint256,uint256,uint256) claimLockup)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) WatchClaimLockupCreated(opts *bind.WatchOpts, sink chan<- *IHedgeyAirdropClaimLockupCreated, id [][16]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.WatchLogs(opts, "ClaimLockupCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyAirdropClaimLockupCreated)
				if err := _IHedgeyAirdrop.contract.UnpackLog(event, "ClaimLockupCreated", log); err != nil {
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

// ParseClaimLockupCreated is a log parse operation binding the contract event 0xce04f3b9b2dfc43b1d1e78376db74a3b32c0a78daaf34efc81334908359448fc.
//
// Solidity: event ClaimLockupCreated(bytes16 indexed id, (address,uint256,uint256,uint256,uint256) claimLockup)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) ParseClaimLockupCreated(log types.Log) (*IHedgeyAirdropClaimLockupCreated, error) {
	event := new(IHedgeyAirdropClaimLockupCreated)
	if err := _IHedgeyAirdrop.contract.UnpackLog(event, "ClaimLockupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyAirdropClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropClaimedIterator struct {
	Event *IHedgeyAirdropClaimed // Event containing the contract specifics and raw log

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
func (it *IHedgeyAirdropClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyAirdropClaimed)
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
		it.Event = new(IHedgeyAirdropClaimed)
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
func (it *IHedgeyAirdropClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyAirdropClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyAirdropClaimed represents a Claimed event raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropClaimed struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed recipient, uint256 indexed amount)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) FilterClaimed(opts *bind.FilterOpts, recipient []common.Address, amount []*big.Int) (*IHedgeyAirdropClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.FilterLogs(opts, "Claimed", recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropClaimedIterator{contract: _IHedgeyAirdrop.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed recipient, uint256 indexed amount)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *IHedgeyAirdropClaimed, recipient []common.Address, amount []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.WatchLogs(opts, "Claimed", recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyAirdropClaimed)
				if err := _IHedgeyAirdrop.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed recipient, uint256 indexed amount)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) ParseClaimed(log types.Log) (*IHedgeyAirdropClaimed, error) {
	event := new(IHedgeyAirdropClaimed)
	if err := _IHedgeyAirdrop.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyAirdropLockedTokensClaimedIterator is returned from FilterLockedTokensClaimed and is used to iterate over the raw logs and unpacked data for LockedTokensClaimed events raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropLockedTokensClaimedIterator struct {
	Event *IHedgeyAirdropLockedTokensClaimed // Event containing the contract specifics and raw log

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
func (it *IHedgeyAirdropLockedTokensClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyAirdropLockedTokensClaimed)
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
		it.Event = new(IHedgeyAirdropLockedTokensClaimed)
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
func (it *IHedgeyAirdropLockedTokensClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyAirdropLockedTokensClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyAirdropLockedTokensClaimed represents a LockedTokensClaimed event raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropLockedTokensClaimed struct {
	Id              [16]byte
	Claimer         common.Address
	TokenId         *big.Int
	AmountClaimed   *big.Int
	AmountRemaining *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLockedTokensClaimed is a free log retrieval operation binding the contract event 0xb73fa52621a0eab1e12dfeda1e46cb8d3faf8f59e22cf6c2610c680eced3ebc3.
//
// Solidity: event LockedTokensClaimed(bytes16 indexed id, address indexed claimer, uint256 indexed tokenId, uint256 amountClaimed, uint256 amountRemaining)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) FilterLockedTokensClaimed(opts *bind.FilterOpts, id [][16]byte, claimer []common.Address, tokenId []*big.Int) (*IHedgeyAirdropLockedTokensClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.FilterLogs(opts, "LockedTokensClaimed", idRule, claimerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropLockedTokensClaimedIterator{contract: _IHedgeyAirdrop.contract, event: "LockedTokensClaimed", logs: logs, sub: sub}, nil
}

// WatchLockedTokensClaimed is a free log subscription operation binding the contract event 0xb73fa52621a0eab1e12dfeda1e46cb8d3faf8f59e22cf6c2610c680eced3ebc3.
//
// Solidity: event LockedTokensClaimed(bytes16 indexed id, address indexed claimer, uint256 indexed tokenId, uint256 amountClaimed, uint256 amountRemaining)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) WatchLockedTokensClaimed(opts *bind.WatchOpts, sink chan<- *IHedgeyAirdropLockedTokensClaimed, id [][16]byte, claimer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.WatchLogs(opts, "LockedTokensClaimed", idRule, claimerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyAirdropLockedTokensClaimed)
				if err := _IHedgeyAirdrop.contract.UnpackLog(event, "LockedTokensClaimed", log); err != nil {
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

// ParseLockedTokensClaimed is a log parse operation binding the contract event 0xb73fa52621a0eab1e12dfeda1e46cb8d3faf8f59e22cf6c2610c680eced3ebc3.
//
// Solidity: event LockedTokensClaimed(bytes16 indexed id, address indexed claimer, uint256 indexed tokenId, uint256 amountClaimed, uint256 amountRemaining)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) ParseLockedTokensClaimed(log types.Log) (*IHedgeyAirdropLockedTokensClaimed, error) {
	event := new(IHedgeyAirdropLockedTokensClaimed)
	if err := _IHedgeyAirdrop.contract.UnpackLog(event, "LockedTokensClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyAirdropUnlockedTokensClaimedIterator is returned from FilterUnlockedTokensClaimed and is used to iterate over the raw logs and unpacked data for UnlockedTokensClaimed events raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropUnlockedTokensClaimedIterator struct {
	Event *IHedgeyAirdropUnlockedTokensClaimed // Event containing the contract specifics and raw log

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
func (it *IHedgeyAirdropUnlockedTokensClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyAirdropUnlockedTokensClaimed)
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
		it.Event = new(IHedgeyAirdropUnlockedTokensClaimed)
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
func (it *IHedgeyAirdropUnlockedTokensClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyAirdropUnlockedTokensClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyAirdropUnlockedTokensClaimed represents a UnlockedTokensClaimed event raised by the IHedgeyAirdrop contract.
type IHedgeyAirdropUnlockedTokensClaimed struct {
	Id              [16]byte
	Claimer         common.Address
	AmountClaimed   *big.Int
	AmountRemaining *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlockedTokensClaimed is a free log retrieval operation binding the contract event 0xeda9c5842f9de63a524e4bafcab717969b8cfdb8bbd1533e39497520dbc11275.
//
// Solidity: event UnlockedTokensClaimed(bytes16 indexed id, address indexed claimer, uint256 amountClaimed, uint256 amountRemaining)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) FilterUnlockedTokensClaimed(opts *bind.FilterOpts, id [][16]byte, claimer []common.Address) (*IHedgeyAirdropUnlockedTokensClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.FilterLogs(opts, "UnlockedTokensClaimed", idRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyAirdropUnlockedTokensClaimedIterator{contract: _IHedgeyAirdrop.contract, event: "UnlockedTokensClaimed", logs: logs, sub: sub}, nil
}

// WatchUnlockedTokensClaimed is a free log subscription operation binding the contract event 0xeda9c5842f9de63a524e4bafcab717969b8cfdb8bbd1533e39497520dbc11275.
//
// Solidity: event UnlockedTokensClaimed(bytes16 indexed id, address indexed claimer, uint256 amountClaimed, uint256 amountRemaining)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) WatchUnlockedTokensClaimed(opts *bind.WatchOpts, sink chan<- *IHedgeyAirdropUnlockedTokensClaimed, id [][16]byte, claimer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _IHedgeyAirdrop.contract.WatchLogs(opts, "UnlockedTokensClaimed", idRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyAirdropUnlockedTokensClaimed)
				if err := _IHedgeyAirdrop.contract.UnpackLog(event, "UnlockedTokensClaimed", log); err != nil {
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

// ParseUnlockedTokensClaimed is a log parse operation binding the contract event 0xeda9c5842f9de63a524e4bafcab717969b8cfdb8bbd1533e39497520dbc11275.
//
// Solidity: event UnlockedTokensClaimed(bytes16 indexed id, address indexed claimer, uint256 amountClaimed, uint256 amountRemaining)
func (_IHedgeyAirdrop *IHedgeyAirdropFilterer) ParseUnlockedTokensClaimed(log types.Log) (*IHedgeyAirdropUnlockedTokensClaimed, error) {
	event := new(IHedgeyAirdropUnlockedTokensClaimed)
	if err := _IHedgeyAirdrop.contract.UnpackLog(event, "UnlockedTokensClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
