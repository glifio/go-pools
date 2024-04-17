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
type Account_duplicate_Infinity struct {
	StartEpoch *big.Int
	Principal  *big.Int
	EpochsPaid *big.Int
	Defaulted  bool
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_InfinityPool struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// InfinityPoolMetaData contains all meta data concerning the InfinityPool contract.
var InfinityPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rateModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidStakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_preStake\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountDNE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDefaulted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayUp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolShuttingDown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lostFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"}],\"name\":\"WriteOff\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"decommissionPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesCollected\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAbsMinLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"harvestAmount\",\"type\":\"uint256\"}],\"name\":\"harvestFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestToRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"defaulted\",\"type\":\"bool\"}],\"internalType\":\"structAccount\",\"name\":\"account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShuttingDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountPrincipal\",\"type\":\"uint256\"}],\"name\":\"jumpStartAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"jumpStartTotalBorrowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingToken\",\"outputs\":[{\"internalType\":\"contractIPoolToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEpochsOwedTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preStake\",\"outputs\":[{\"internalType\":\"contractIPreStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ramp\",\"outputs\":[{\"internalType\":\"contractIOffRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateModule\",\"outputs\":[{\"internalType\":\"contractIRateModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"recoverFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxEpochsOwedTolerance\",\"type\":\"uint256\"}],\"name\":\"setMaxEpochsOwedTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumLiquidity\",\"type\":\"uint256\"}],\"name\":\"setMinimumLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOffRamp\",\"name\":\"_ramp\",\"type\":\"address\"}],\"name\":\"setRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRateModule\",\"name\":\"_rateModule\",\"type\":\"address\"}],\"name\":\"setRateModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowableAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_preStake\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFromPreStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"}],\"name\":\"writeOff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalOwed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// InfinityPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use InfinityPoolMetaData.ABI instead.
var InfinityPoolABI = InfinityPoolMetaData.ABI

// InfinityPool is an auto generated Go binding around an Ethereum contract.
type InfinityPool struct {
	InfinityPoolCaller     // Read-only binding to the contract
	InfinityPoolTransactor // Write-only binding to the contract
	InfinityPoolFilterer   // Log filterer for contract events
}

// InfinityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type InfinityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InfinityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfinityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfinityPoolSession struct {
	Contract     *InfinityPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfinityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfinityPoolCallerSession struct {
	Contract *InfinityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InfinityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfinityPoolTransactorSession struct {
	Contract     *InfinityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InfinityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type InfinityPoolRaw struct {
	Contract *InfinityPool // Generic contract binding to access the raw methods on
}

// InfinityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfinityPoolCallerRaw struct {
	Contract *InfinityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// InfinityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfinityPoolTransactorRaw struct {
	Contract *InfinityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInfinityPool creates a new instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPool(address common.Address, backend bind.ContractBackend) (*InfinityPool, error) {
	contract, err := bindInfinityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfinityPool{InfinityPoolCaller: InfinityPoolCaller{contract: contract}, InfinityPoolTransactor: InfinityPoolTransactor{contract: contract}, InfinityPoolFilterer: InfinityPoolFilterer{contract: contract}}, nil
}

// NewInfinityPoolCaller creates a new read-only instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPoolCaller(address common.Address, caller bind.ContractCaller) (*InfinityPoolCaller, error) {
	contract, err := bindInfinityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolCaller{contract: contract}, nil
}

// NewInfinityPoolTransactor creates a new write-only instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*InfinityPoolTransactor, error) {
	contract, err := bindInfinityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolTransactor{contract: contract}, nil
}

// NewInfinityPoolFilterer creates a new log filterer instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*InfinityPoolFilterer, error) {
	contract, err := bindInfinityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolFilterer{contract: contract}, nil
}

// bindInfinityPool binds a generic wrapper to an already deployed contract.
func bindInfinityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InfinityPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfinityPool *InfinityPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfinityPool.Contract.InfinityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfinityPool *InfinityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.Contract.InfinityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfinityPool *InfinityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfinityPool.Contract.InfinityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfinityPool *InfinityPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfinityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfinityPool *InfinityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfinityPool *InfinityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfinityPool.Contract.contract.Transact(opts, method, params...)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPool *InfinityPoolCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPool *InfinityPoolSession) Asset() (common.Address, error) {
	return _InfinityPool.Contract.Asset(&_InfinityPool.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) Asset() (common.Address, error) {
	return _InfinityPool.Contract.Asset(&_InfinityPool.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToAssets(&_InfinityPool.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToAssets(&_InfinityPool.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToShares(&_InfinityPool.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToShares(&_InfinityPool.CallOpts, assets)
}

// FeesCollected is a free data retrieval call binding the contract method 0xf071db5a.
//
// Solidity: function feesCollected() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) FeesCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "feesCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeesCollected is a free data retrieval call binding the contract method 0xf071db5a.
//
// Solidity: function feesCollected() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) FeesCollected() (*big.Int, error) {
	return _InfinityPool.Contract.FeesCollected(&_InfinityPool.CallOpts)
}

// FeesCollected is a free data retrieval call binding the contract method 0xf071db5a.
//
// Solidity: function feesCollected() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) FeesCollected() (*big.Int, error) {
	return _InfinityPool.Contract.FeesCollected(&_InfinityPool.CallOpts)
}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) GetAbsMinLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getAbsMinLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) GetAbsMinLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.GetAbsMinLiquidity(&_InfinityPool.CallOpts)
}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) GetAbsMinLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.GetAbsMinLiquidity(&_InfinityPool.CallOpts)
}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) GetAgentBorrowed(opts *bind.CallOpts, agentID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getAgentBorrowed", agentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) GetAgentBorrowed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.GetAgentBorrowed(&_InfinityPool.CallOpts, agentID)
}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) GetAgentBorrowed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.GetAgentBorrowed(&_InfinityPool.CallOpts, agentID)
}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) GetLiquidAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getLiquidAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) GetLiquidAssets() (*big.Int, error) {
	return _InfinityPool.Contract.GetLiquidAssets(&_InfinityPool.CallOpts)
}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) GetLiquidAssets() (*big.Int, error) {
	return _InfinityPool.Contract.GetLiquidAssets(&_InfinityPool.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_InfinityPool *InfinityPoolCaller) GetRate(opts *bind.CallOpts, vc VerifiableCredential) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getRate", vc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_InfinityPool *InfinityPoolSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _InfinityPool.Contract.GetRate(&_InfinityPool.CallOpts, vc)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_InfinityPool *InfinityPoolCallerSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _InfinityPool.Contract.GetRate(&_InfinityPool.CallOpts, vc)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) Id() (*big.Int, error) {
	return _InfinityPool.Contract.Id(&_InfinityPool.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) Id() (*big.Int, error) {
	return _InfinityPool.Contract.Id(&_InfinityPool.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool approved)
func (_InfinityPool *InfinityPoolCaller) IsApproved(opts *bind.CallOpts, account Account, vc VerifiableCredential) (bool, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "isApproved", account, vc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool approved)
func (_InfinityPool *InfinityPoolSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _InfinityPool.Contract.IsApproved(&_InfinityPool.CallOpts, account, vc)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool approved)
func (_InfinityPool *InfinityPoolCallerSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _InfinityPool.Contract.IsApproved(&_InfinityPool.CallOpts, account, vc)
}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPool *InfinityPoolCaller) IsShuttingDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "isShuttingDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPool *InfinityPoolSession) IsShuttingDown() (bool, error) {
	return _InfinityPool.Contract.IsShuttingDown(&_InfinityPool.CallOpts)
}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPool *InfinityPoolCallerSession) IsShuttingDown() (bool, error) {
	return _InfinityPool.Contract.IsShuttingDown(&_InfinityPool.CallOpts)
}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPool *InfinityPoolCaller) LiquidStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "liquidStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPool *InfinityPoolSession) LiquidStakingToken() (common.Address, error) {
	return _InfinityPool.Contract.LiquidStakingToken(&_InfinityPool.CallOpts)
}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) LiquidStakingToken() (common.Address, error) {
	return _InfinityPool.Contract.LiquidStakingToken(&_InfinityPool.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxDeposit(&_InfinityPool.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxDeposit(&_InfinityPool.CallOpts, arg0)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxEpochsOwedTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxEpochsOwedTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _InfinityPool.Contract.MaxEpochsOwedTolerance(&_InfinityPool.CallOpts)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _InfinityPool.Contract.MaxEpochsOwedTolerance(&_InfinityPool.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxMint(&_InfinityPool.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxMint(&_InfinityPool.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxRedeem(&_InfinityPool.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxRedeem(&_InfinityPool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxWithdraw(&_InfinityPool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxWithdraw(&_InfinityPool.CallOpts, owner)
}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MinimumLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "minimumLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MinimumLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.MinimumLiquidity(&_InfinityPool.CallOpts)
}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MinimumLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.MinimumLiquidity(&_InfinityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPool *InfinityPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPool *InfinityPoolSession) Owner() (common.Address, error) {
	return _InfinityPool.Contract.Owner(&_InfinityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) Owner() (common.Address, error) {
	return _InfinityPool.Contract.Owner(&_InfinityPool.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPool *InfinityPoolCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPool *InfinityPoolSession) PendingOwner() (common.Address, error) {
	return _InfinityPool.Contract.PendingOwner(&_InfinityPool.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) PendingOwner() (common.Address, error) {
	return _InfinityPool.Contract.PendingOwner(&_InfinityPool.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_InfinityPool *InfinityPoolCaller) PreStake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "preStake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_InfinityPool *InfinityPoolSession) PreStake() (common.Address, error) {
	return _InfinityPool.Contract.PreStake(&_InfinityPool.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) PreStake() (common.Address, error) {
	return _InfinityPool.Contract.PreStake(&_InfinityPool.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewDeposit(&_InfinityPool.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewDeposit(&_InfinityPool.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewMint(&_InfinityPool.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewMint(&_InfinityPool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewRedeem(&_InfinityPool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewRedeem(&_InfinityPool.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewWithdraw(&_InfinityPool.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewWithdraw(&_InfinityPool.CallOpts, assets)
}

// Ramp is a free data retrieval call binding the contract method 0x15d276e1.
//
// Solidity: function ramp() view returns(address)
func (_InfinityPool *InfinityPoolCaller) Ramp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "ramp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ramp is a free data retrieval call binding the contract method 0x15d276e1.
//
// Solidity: function ramp() view returns(address)
func (_InfinityPool *InfinityPoolSession) Ramp() (common.Address, error) {
	return _InfinityPool.Contract.Ramp(&_InfinityPool.CallOpts)
}

// Ramp is a free data retrieval call binding the contract method 0x15d276e1.
//
// Solidity: function ramp() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) Ramp() (common.Address, error) {
	return _InfinityPool.Contract.Ramp(&_InfinityPool.CallOpts)
}

// RateModule is a free data retrieval call binding the contract method 0x336d391e.
//
// Solidity: function rateModule() view returns(address)
func (_InfinityPool *InfinityPoolCaller) RateModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "rateModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateModule is a free data retrieval call binding the contract method 0x336d391e.
//
// Solidity: function rateModule() view returns(address)
func (_InfinityPool *InfinityPoolSession) RateModule() (common.Address, error) {
	return _InfinityPool.Contract.RateModule(&_InfinityPool.CallOpts)
}

// RateModule is a free data retrieval call binding the contract method 0x336d391e.
//
// Solidity: function rateModule() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) RateModule() (common.Address, error) {
	return _InfinityPool.Contract.RateModule(&_InfinityPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) TotalAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalAssets(&_InfinityPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) TotalAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalAssets(&_InfinityPool.CallOpts)
}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) TotalBorrowableAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "totalBorrowableAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) TotalBorrowableAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowableAssets(&_InfinityPool.CallOpts)
}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) TotalBorrowableAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowableAssets(&_InfinityPool.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) TotalBorrowed() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowed(&_InfinityPool.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) TotalBorrowed() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowed(&_InfinityPool.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPool *InfinityPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPool *InfinityPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _InfinityPool.Contract.AcceptOwnership(&_InfinityPool.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPool *InfinityPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _InfinityPool.Contract.AcceptOwnership(&_InfinityPool.TransactOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPool *InfinityPoolTransactor) Borrow(opts *bind.TransactOpts, vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "borrow", vc)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPool *InfinityPoolSession) Borrow(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Borrow(&_InfinityPool.TransactOpts, vc)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPool *InfinityPoolTransactorSession) Borrow(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Borrow(&_InfinityPool.TransactOpts, vc)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPool *InfinityPoolTransactor) DecommissionPool(opts *bind.TransactOpts, newPool common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "decommissionPool", newPool)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPool *InfinityPoolSession) DecommissionPool(newPool common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.DecommissionPool(&_InfinityPool.TransactOpts, newPool)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPool *InfinityPoolTransactorSession) DecommissionPool(newPool common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.DecommissionPool(&_InfinityPool.TransactOpts, newPool)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPool *InfinityPoolSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit(&_InfinityPool.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit(&_InfinityPool.TransactOpts, assets, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactor) Deposit0(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "deposit0", receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPool *InfinityPoolSession) Deposit0(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit0(&_InfinityPool.TransactOpts, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactorSession) Deposit0(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit0(&_InfinityPool.TransactOpts, receiver)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPool *InfinityPoolTransactor) HarvestFees(opts *bind.TransactOpts, harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "harvestFees", harvestAmount)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPool *InfinityPoolSession) HarvestFees(harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestFees(&_InfinityPool.TransactOpts, harvestAmount)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPool *InfinityPoolTransactorSession) HarvestFees(harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestFees(&_InfinityPool.TransactOpts, harvestAmount)
}

// HarvestToRamp is a paid mutator transaction binding the contract method 0x38dff49c.
//
// Solidity: function harvestToRamp() returns()
func (_InfinityPool *InfinityPoolTransactor) HarvestToRamp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "harvestToRamp")
}

// HarvestToRamp is a paid mutator transaction binding the contract method 0x38dff49c.
//
// Solidity: function harvestToRamp() returns()
func (_InfinityPool *InfinityPoolSession) HarvestToRamp() (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestToRamp(&_InfinityPool.TransactOpts)
}

// HarvestToRamp is a paid mutator transaction binding the contract method 0x38dff49c.
//
// Solidity: function harvestToRamp() returns()
func (_InfinityPool *InfinityPoolTransactorSession) HarvestToRamp() (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestToRamp(&_InfinityPool.TransactOpts)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPool *InfinityPoolTransactor) JumpStartAccount(opts *bind.TransactOpts, receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "jumpStartAccount", receiver, agentID, accountPrincipal)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPool *InfinityPoolSession) JumpStartAccount(receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartAccount(&_InfinityPool.TransactOpts, receiver, agentID, accountPrincipal)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPool *InfinityPoolTransactorSession) JumpStartAccount(receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartAccount(&_InfinityPool.TransactOpts, receiver, agentID, accountPrincipal)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPool *InfinityPoolTransactor) JumpStartTotalBorrowed(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "jumpStartTotalBorrowed", amount)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPool *InfinityPoolSession) JumpStartTotalBorrowed(amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartTotalBorrowed(&_InfinityPool.TransactOpts, amount)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPool *InfinityPoolTransactorSession) JumpStartTotalBorrowed(amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartTotalBorrowed(&_InfinityPool.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPool *InfinityPoolSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Mint(&_InfinityPool.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Mint(&_InfinityPool.TransactOpts, shares, receiver)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolTransactor) Pay(opts *bind.TransactOpts, vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "pay", vc)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolSession) Pay(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Pay(&_InfinityPool.TransactOpts, vc)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolTransactorSession) Pay(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Pay(&_InfinityPool.TransactOpts, vc)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPool *InfinityPoolTransactor) RecoverERC20(opts *bind.TransactOpts, receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "recoverERC20", receiver, token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPool *InfinityPoolSession) RecoverERC20(receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverERC20(&_InfinityPool.TransactOpts, receiver, token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPool *InfinityPoolTransactorSession) RecoverERC20(receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverERC20(&_InfinityPool.TransactOpts, receiver, token)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPool *InfinityPoolTransactor) RecoverFIL(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "recoverFIL", receiver)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPool *InfinityPoolSession) RecoverFIL(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverFIL(&_InfinityPool.TransactOpts, receiver)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPool *InfinityPoolTransactorSession) RecoverFIL(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverFIL(&_InfinityPool.TransactOpts, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPool *InfinityPoolSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Redeem(&_InfinityPool.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Redeem(&_InfinityPool.TransactOpts, shares, receiver, owner)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPool *InfinityPoolTransactor) RefreshRoutes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "refreshRoutes")
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPool *InfinityPoolSession) RefreshRoutes() (*types.Transaction, error) {
	return _InfinityPool.Contract.RefreshRoutes(&_InfinityPool.TransactOpts)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPool *InfinityPoolTransactorSession) RefreshRoutes() (*types.Transaction, error) {
	return _InfinityPool.Contract.RefreshRoutes(&_InfinityPool.TransactOpts)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_InfinityPool *InfinityPoolTransactor) SetMaxEpochsOwedTolerance(opts *bind.TransactOpts, _maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setMaxEpochsOwedTolerance", _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_InfinityPool *InfinityPoolSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMaxEpochsOwedTolerance(&_InfinityPool.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMaxEpochsOwedTolerance(&_InfinityPool.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPool *InfinityPoolTransactor) SetMinimumLiquidity(opts *bind.TransactOpts, _minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setMinimumLiquidity", _minimumLiquidity)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPool *InfinityPoolSession) SetMinimumLiquidity(_minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMinimumLiquidity(&_InfinityPool.TransactOpts, _minimumLiquidity)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetMinimumLiquidity(_minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMinimumLiquidity(&_InfinityPool.TransactOpts, _minimumLiquidity)
}

// SetRamp is a paid mutator transaction binding the contract method 0x3d9c3c17.
//
// Solidity: function setRamp(address _ramp) returns()
func (_InfinityPool *InfinityPoolTransactor) SetRamp(opts *bind.TransactOpts, _ramp common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setRamp", _ramp)
}

// SetRamp is a paid mutator transaction binding the contract method 0x3d9c3c17.
//
// Solidity: function setRamp(address _ramp) returns()
func (_InfinityPool *InfinityPoolSession) SetRamp(_ramp common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRamp(&_InfinityPool.TransactOpts, _ramp)
}

// SetRamp is a paid mutator transaction binding the contract method 0x3d9c3c17.
//
// Solidity: function setRamp(address _ramp) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetRamp(_ramp common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRamp(&_InfinityPool.TransactOpts, _ramp)
}

// SetRateModule is a paid mutator transaction binding the contract method 0x31ff92fa.
//
// Solidity: function setRateModule(address _rateModule) returns()
func (_InfinityPool *InfinityPoolTransactor) SetRateModule(opts *bind.TransactOpts, _rateModule common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setRateModule", _rateModule)
}

// SetRateModule is a paid mutator transaction binding the contract method 0x31ff92fa.
//
// Solidity: function setRateModule(address _rateModule) returns()
func (_InfinityPool *InfinityPoolSession) SetRateModule(_rateModule common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRateModule(&_InfinityPool.TransactOpts, _rateModule)
}

// SetRateModule is a paid mutator transaction binding the contract method 0x31ff92fa.
//
// Solidity: function setRateModule(address _rateModule) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetRateModule(_rateModule common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRateModule(&_InfinityPool.TransactOpts, _rateModule)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPool *InfinityPoolTransactor) ShutDown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "shutDown")
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPool *InfinityPoolSession) ShutDown() (*types.Transaction, error) {
	return _InfinityPool.Contract.ShutDown(&_InfinityPool.TransactOpts)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPool *InfinityPoolTransactorSession) ShutDown() (*types.Transaction, error) {
	return _InfinityPool.Contract.ShutDown(&_InfinityPool.TransactOpts)
}

// TransferFromPreStake is a paid mutator transaction binding the contract method 0x00dcfc20.
//
// Solidity: function transferFromPreStake(address _preStake, uint256 _amount) returns()
func (_InfinityPool *InfinityPoolTransactor) TransferFromPreStake(opts *bind.TransactOpts, _preStake common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "transferFromPreStake", _preStake, _amount)
}

// TransferFromPreStake is a paid mutator transaction binding the contract method 0x00dcfc20.
//
// Solidity: function transferFromPreStake(address _preStake, uint256 _amount) returns()
func (_InfinityPool *InfinityPoolSession) TransferFromPreStake(_preStake common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferFromPreStake(&_InfinityPool.TransactOpts, _preStake, _amount)
}

// TransferFromPreStake is a paid mutator transaction binding the contract method 0x00dcfc20.
//
// Solidity: function transferFromPreStake(address _preStake, uint256 _amount) returns()
func (_InfinityPool *InfinityPoolTransactorSession) TransferFromPreStake(_preStake common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferFromPreStake(&_InfinityPool.TransactOpts, _preStake, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPool *InfinityPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPool *InfinityPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferOwnership(&_InfinityPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPool *InfinityPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferOwnership(&_InfinityPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPool *InfinityPoolSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Withdraw(&_InfinityPool.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Withdraw(&_InfinityPool.TransactOpts, assets, receiver, owner)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns(uint256 totalOwed)
func (_InfinityPool *InfinityPoolTransactor) WriteOff(opts *bind.TransactOpts, agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "writeOff", agentID, recoveredFunds)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns(uint256 totalOwed)
func (_InfinityPool *InfinityPoolSession) WriteOff(agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.WriteOff(&_InfinityPool.TransactOpts, agentID, recoveredFunds)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns(uint256 totalOwed)
func (_InfinityPool *InfinityPoolTransactorSession) WriteOff(agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.WriteOff(&_InfinityPool.TransactOpts, agentID, recoveredFunds)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPool *InfinityPoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _InfinityPool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPool *InfinityPoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InfinityPool.Contract.Fallback(&_InfinityPool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPool *InfinityPoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InfinityPool.Contract.Fallback(&_InfinityPool.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPool *InfinityPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPool *InfinityPoolSession) Receive() (*types.Transaction, error) {
	return _InfinityPool.Contract.Receive(&_InfinityPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPool *InfinityPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _InfinityPool.Contract.Receive(&_InfinityPool.TransactOpts)
}

// InfinityPoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the InfinityPool contract.
type InfinityPoolBorrowIterator struct {
	Event *InfinityPoolBorrow // Event containing the contract specifics and raw log

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
func (it *InfinityPoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolBorrow)
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
		it.Event = new(InfinityPoolBorrow)
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
func (it *InfinityPoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolBorrow represents a Borrow event raised by the InfinityPool contract.
type InfinityPoolBorrow struct {
	Agent  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPool *InfinityPoolFilterer) FilterBorrow(opts *bind.FilterOpts, agent []*big.Int) (*InfinityPoolBorrowIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Borrow", agentRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolBorrowIterator{contract: _InfinityPool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPool *InfinityPoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *InfinityPoolBorrow, agent []*big.Int) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Borrow", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolBorrow)
				if err := _InfinityPool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPool *InfinityPoolFilterer) ParseBorrow(log types.Log) (*InfinityPoolBorrow, error) {
	event := new(InfinityPoolBorrow)
	if err := _InfinityPool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the InfinityPool contract.
type InfinityPoolDepositIterator struct {
	Event *InfinityPoolDeposit // Event containing the contract specifics and raw log

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
func (it *InfinityPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolDeposit)
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
		it.Event = new(InfinityPoolDeposit)
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
func (it *InfinityPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolDeposit represents a Deposit event raised by the InfinityPool contract.
type InfinityPoolDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*InfinityPoolDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolDepositIterator{contract: _InfinityPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *InfinityPoolDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolDeposit)
				if err := _InfinityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) ParseDeposit(log types.Log) (*InfinityPoolDeposit, error) {
	event := new(InfinityPoolDeposit)
	if err := _InfinityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the InfinityPool contract.
type InfinityPoolPayIterator struct {
	Event *InfinityPoolPay // Event containing the contract specifics and raw log

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
func (it *InfinityPoolPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolPay)
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
		it.Event = new(InfinityPoolPay)
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
func (it *InfinityPoolPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolPay represents a Pay event raised by the InfinityPool contract.
type InfinityPoolPay struct {
	Agent         *big.Int
	Rate          *big.Int
	EpochsPaid    *big.Int
	PrincipalPaid *big.Int
	Refund        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolFilterer) FilterPay(opts *bind.FilterOpts, agent []*big.Int) (*InfinityPoolPayIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Pay", agentRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolPayIterator{contract: _InfinityPool.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *InfinityPoolPay, agent []*big.Int) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Pay", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolPay)
				if err := _InfinityPool.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolFilterer) ParsePay(log types.Log) (*InfinityPoolPay, error) {
	event := new(InfinityPoolPay)
	if err := _InfinityPool.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the InfinityPool contract.
type InfinityPoolWithdrawIterator struct {
	Event *InfinityPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *InfinityPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolWithdraw)
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
		it.Event = new(InfinityPoolWithdraw)
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
func (it *InfinityPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolWithdraw represents a Withdraw event raised by the InfinityPool contract.
type InfinityPoolWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*InfinityPoolWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolWithdrawIterator{contract: _InfinityPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *InfinityPoolWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolWithdraw)
				if err := _InfinityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) ParseWithdraw(log types.Log) (*InfinityPoolWithdraw, error) {
	event := new(InfinityPoolWithdraw)
	if err := _InfinityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolWriteOffIterator is returned from FilterWriteOff and is used to iterate over the raw logs and unpacked data for WriteOff events raised by the InfinityPool contract.
type InfinityPoolWriteOffIterator struct {
	Event *InfinityPoolWriteOff // Event containing the contract specifics and raw log

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
func (it *InfinityPoolWriteOffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolWriteOff)
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
		it.Event = new(InfinityPoolWriteOff)
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
func (it *InfinityPoolWriteOffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolWriteOffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolWriteOff represents a WriteOff event raised by the InfinityPool contract.
type InfinityPoolWriteOff struct {
	AgentID        *big.Int
	RecoveredFunds *big.Int
	LostFunds      *big.Int
	InterestPaid   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWriteOff is a free log retrieval operation binding the contract event 0x31d6d43b5ab97f9739e3e78ab72e97d7575b3233f23e8e929160920032436aa5.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 interestPaid)
func (_InfinityPool *InfinityPoolFilterer) FilterWriteOff(opts *bind.FilterOpts, agentID []*big.Int) (*InfinityPoolWriteOffIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "WriteOff", agentIDRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolWriteOffIterator{contract: _InfinityPool.contract, event: "WriteOff", logs: logs, sub: sub}, nil
}

// WatchWriteOff is a free log subscription operation binding the contract event 0x31d6d43b5ab97f9739e3e78ab72e97d7575b3233f23e8e929160920032436aa5.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 interestPaid)
func (_InfinityPool *InfinityPoolFilterer) WatchWriteOff(opts *bind.WatchOpts, sink chan<- *InfinityPoolWriteOff, agentID []*big.Int) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "WriteOff", agentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolWriteOff)
				if err := _InfinityPool.contract.UnpackLog(event, "WriteOff", log); err != nil {
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

// ParseWriteOff is a log parse operation binding the contract event 0x31d6d43b5ab97f9739e3e78ab72e97d7575b3233f23e8e929160920032436aa5.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 interestPaid)
func (_InfinityPool *InfinityPoolFilterer) ParseWriteOff(log types.Log) (*InfinityPoolWriteOff, error) {
	event := new(InfinityPoolWriteOff)
	if err := _InfinityPool.contract.UnpackLog(event, "WriteOff", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
