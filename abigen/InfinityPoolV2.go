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

// RateUpdate is an auto generated low-level Go binding around an user-defined struct.
type RateUpdate struct {
	TotalAccountsAtUpdate *big.Int
	TotalAccountsClosed   *big.Int
	StartEpoch            *big.Int
	NewRate               *big.Int
	InProcess             bool
}

// RewardAccrual is an auto generated low-level Go binding around an user-defined struct.
type RewardAccrual struct {
	Accrued *big.Int
	Paid    *big.Int
	Lost    *big.Int
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_duplicate struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// InfinityPoolV2MetaData contains all meta data concerning the InfinityPoolV2 contract.
var InfinityPoolV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidStakingToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumLiquidity_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountDNE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDefaulted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolShuttingDown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEarnings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iFILPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iFILSupply\",\"type\":\"uint256\"}],\"name\":\"UpdateAccounting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lostFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lostTFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"}],\"name\":\"WriteOff\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"continueRateUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credParser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"decommissionPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAbsMinLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentInterestOwed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"harvestAmount\",\"type\":\"uint256\"}],\"name\":\"harvestFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShuttingDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountPrincipal\",\"type\":\"uint256\"}],\"name\":\"jumpStartAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"jumpStartTotalBorrowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccountingUpdateEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingToken\",\"outputs\":[{\"internalType\":\"contractIPoolToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lm\",\"outputs\":[{\"internalType\":\"contractILiquidityMineSP\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accrued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structRewardAccrual\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateUpdate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalAccountsAtUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAccountsClosed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"inProcess\",\"type\":\"bool\"}],\"internalType\":\"structRateUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"recoverFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeemF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeemF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityMine\",\"type\":\"address\"}],\"name\":\"setLiquidityMine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAccountsToUpdatePerBatch_\",\"type\":\"uint256\"}],\"name\":\"setMaxAccountsToUpdatePerBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumLiquidity\",\"type\":\"uint256\"}],\"name\":\"setMinimumLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treasuryFeeRate\",\"type\":\"uint256\"}],\"name\":\"setTreasuryFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rentalFeesOwedPerEpoch_\",\"type\":\"uint256\"}],\"name\":\"startRateUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowableAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFeesOwed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accrued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structRewardAccrual\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAccounting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCredParser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"}],\"name\":\"writeOff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61010060405267016345785d8a00006005555f600681905560078190556008556009805460ff19169055620000816e1ce3922c2af443414b57dc000000006200004c601e610e1062000544565b6200005990601862000564565b620000679061016d62000564565b6200007b90670de0b6b3a764000062000564565b62000232565b600a5560fa60165534801562000095575f80fd5b506040516200607438038062006074833981016040819052620000b891620005a4565b84620000cd6001600160a01b03821662000251565b90506001600160a01b038116620000f757604051635435b28960e11b815260040160405180910390fd5b62000102816200029b565b506001805460ff60a01b191690556001600160a01b038416608081905260408051808201825260118152702927aaaa22a92faba324a62faa27a5a2a760791b60209091015251630d37324f60e11b815263aee0a13560e01b6004820152631a6e649e90602401602060405180830381865afa15801562000184573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001aa9190620005fd565b6001600160a01b0390811660c052600783905560a0829052831660e052620001d284620002e1565b600280546001600160a01b0319166001600160a01b0392909216919091179055620001fd8462000386565b600480546001600160a01b0319166001600160a01b039290921691909117905562000227620003e3565b505050505062000619565b5f6200024883670de0b6b3a76400008462000446565b90505b92915050565b5f80806200025f846200046a565b91509150816200027157509192915050565b5f806200027e836200049c565b915091508162000292575093949350505050565b95945050505050565b600180546001600160a01b0319169055620002bf6001600160a01b03821662000251565b5f80546001600160a01b0319166001600160a01b039290921691909117905550565b604080518082018252601481527f524f555445525f4147454e545f464143544f525900000000000000000000000060209091015251630d37324f60e11b81526314fb0b6d60e11b60048201525f906001600160a01b03831690631a6e649e906024015b602060405180830381865afa15801562000360573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200024b9190620005fd565b60408051808201825260128152712927aaaa22a92fa1a922a22fa820a929a2a960711b60209091015251630d37324f60e11b815263d26df3b760e01b60048201525f906001600160a01b03831690631a6e649e9060240162000344565b620003ed6200050f565b6001805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258620004293390565b6040516001600160a01b03909116815260200160405180910390a1565b5f825f1904841183021582026200045b575f80fd5b50910281810615159190040190565b5f80600160401b600160a01b03831660ff60981b81036200049657600192506001600160401b03841691505b50915091565b5f80825f526016600a60205f73fe000000000000000000000000000000000000025afa91505f51806001600160a01b031691508060a01c61ffff16905061040a8114620004ea575f92505f91505b50811580620004fa57503d601614155b156200050a57505f928392509050565b915091565b62000523600154600160a01b900460ff1690565b15620005425760405163d93c066560e01b815260040160405180910390fd5b565b5f826200055f57634e487b7160e01b5f52601260045260245ffd5b500490565b80820281158282048414176200024b57634e487b7160e01b5f52601160045260245ffd5b80516001600160a01b03811681146200059f575f80fd5b919050565b5f805f805f60a08688031215620005b9575f80fd5b620005c48662000588565b9450620005d46020870162000588565b9350620005e46040870162000588565b6060870151608090970151959894975095949392505050565b5f602082840312156200060e575f80fd5b620002488262000588565b60805160a05160c05160e05161587c620007f85f395f818161066501528181610ed2015281816112db015281816113ac01528181611d9e01528181613022015281816133dd01528181613a8301528181613b6301528181613bfc01528181613f080152818161435301528181614b76015281816150d6015261517601525f81816104a3015281816106e301528181610e2801528181610fe00152818161175101528181611b5901528181611d1a015281816121b3015281816123370152818161243b015281816126810152818161294301528181612b7601528181612cc701528181612f700152818161369701528181613764015281816137a701528181613d45015281816142cf01528181614ac40152818161521e01526152d501525f8181610b410152818161168401528181611aa701528181612251015281816128d6015281816132fa0152818161397701528181613ce601528181613ee2015281816141be015261473201525f8181611874015281816118d501528181611a8101528181611ea5015281816120930152818161222f0152818161246501528181612582015281816128b0015281816132d80152818161395501528181613cc401528181613ec00152818161419c015281816147100152614e41015261587c5ff3fe608060405260043610610483575f3560e01c806379ba50971161025d578063b3d7f6b911610142578063d905777e116100ba578063f086ce6011610089578063f2fde38b1161006e578063f2fde38b14610d0d578063f340fa0114610d2c578063f4765e9e14610d3f576104d5565b8063f086ce6014610cda578063f1ccc3df14610cf9576104d5565b8063d905777e14610c6a578063e30c397814610c89578063ef8b30f714610ca8578063efde4e6414610cc7576104d5565b8063c564f77211610111578063c6e6f592116100f6578063c6e6f59214610c0d578063ce43303c14610c2c578063ce96cb7714610c4b576104d5565b8063c564f77214610bee578063c63d75b614610719576104d5565b8063b3d7f6b914610b7c578063b460af9414610b9b578063b56cf01114610bba578063ba08765214610bcf576104d5565b80639050c3d7116101d55780639f40a7b3116101a4578063a318c1a411610189578063a318c1a414610b11578063af640d0f14610b30578063b2bcb00214610b63576104d5565b80639f40a7b314610ad35780639f4326d714610af2576104d5565b80639050c3d714610a6157806390fd6bc614610a8057806394bf804d14610a955780639a33bc7c14610ab4576104d5565b80638456cb591161022c578063870f77cf11610211578063870f77cf14610a05578063886f039a14610a245780638da5cb5b14610a43576104d5565b80638456cb59146109dd57806386a2c988146109f1576104d5565b806379ba50971461094c5780637d694730146109605780637edc8fe71461099f578063803ba97e146109be576104d5565b806341d1de97116103835780635fc2faf7116102fb5780636b285af1116102ca5780636e553f65116102af5780636e553f65146108fa578063773fcf641461091957806377c0f07d1461092d576104d5565b80636b285af1146108c75780636e3ac566146108db576104d5565b80635fc2faf71461086b578063666010321461088a578063679aefce1461089f5780636af1a7e8146108b3576104d5565b80634cdad5061161035257806359221e381161033757806359221e38146107fd5780635c975abb1461081c5780635d66b00a14610857576104d5565b80634cdad506146107bf578063549c5877146107de576104d5565b806341d1de971461075857806343fd5dfe14610777578063494347e7146107965780634c19386c146107aa576104d5565b806310b9e58311610416578063338891eb116103e55780633f4ba83a116103ca5780633f4ba83a14610705578063402d267d14610719578063415e819d14610739576104d5565b8063338891eb146106be57806338d52e0f146106d2576104d5565b806310b9e58314610621578063124dfd66146106355780631cbe8df614610654578063282567b41461069f576104d5565b80630c462c5f116104525780630c462c5f146105565780630c7290f2146105755780630c762141146105cc5780630e0a702314610602576104d5565b806301e1d114146104dd5780630521a2411461050457806307a2d13a146105185780630a28a47714610537576104d5565b366104d557610490610d5e565b610498610d9d565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104d3576104d133610df2565b505b005b610490610d5e565b3480156104e8575f80fd5b506104f1610fad565b6040519081526020015b60405180910390f35b34801561050f575f80fd5b506104d36111a2565b348015610523575f80fd5b506104f16105323660046154da565b6112d7565b348015610542575f80fd5b506104f16105513660046154da565b611382565b348015610561575f80fd5b506104f1610570366004615505565b61144b565b348015610580575f80fd5b506105896114f4565b6040516104fb91905f60a0820190508251825260208301516020830152604083015160408301526060830151606083015260808301511515608083015292915050565b3480156105d7575f80fd5b506105e061155d565b60408051825181526020808401519082015291810151908201526060016104fb565b34801561060d575f80fd5b506104f161061c366004615505565b6115bc565b34801561062c575f80fd5b506104d361165a565b348015610640575f80fd5b506104f161064f36600461554a565b611679565b34801561065f575f80fd5b506106877f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016104fb565b3480156106aa575f80fd5b506104d36106b93660046154da565b61185a565b3480156106c9575f80fd5b506104d361186f565b3480156106dd575f80fd5b506106877f000000000000000000000000000000000000000000000000000000000000000081565b348015610710575f80fd5b506104d3611933565b348015610724575f80fd5b506104f161073336600461554a565b505f1990565b348015610744575f80fd5b506104d3610753366004615565565b611943565b348015610763575f80fd5b506106876107723660046154da565b503090565b348015610782575f80fd5b506104d361079136600461554a565b611bd1565b3480156107a1575f80fd5b506104d3611c13565b3480156107b5575f80fd5b506104f160065481565b3480156107ca575f80fd5b506104f16107d93660046154da565b611e50565b3480156107e9575f80fd5b506104d36107f83660046154da565b611e82565b348015610808575f80fd5b506104d36108173660046154da565b611e97565b348015610827575f80fd5b5060015474010000000000000000000000000000000000000000900460ff165b60405190151581526020016104fb565b348015610862575f80fd5b506104f161219c565b348015610876575f80fd5b506104f16108853660046154da565b612229565b348015610895575f80fd5b506104f160055481565b3480156108aa575f80fd5b50600a546104f1565b3480156108be575f80fd5b506105e061227f565b3480156108d2575f80fd5b506104f16122d8565b3480156108e6575f80fd5b506104d36108f53660046154da565b61231a565b348015610905575f80fd5b506104f161091436600461559d565b612513565b348015610924575f80fd5b506104d3612575565b348015610938575f80fd5b50600454610687906001600160a01b031681565b348015610957575f80fd5b506104d36125a6565b34801561096b575f80fd5b5061097f61097a366004615565565b6125d9565b6040805194855260208501939093529183015260608201526080016104fb565b3480156109aa575f80fd5b506104d36109b936600461554a565b612af4565b3480156109c9575f80fd5b50600354610687906001600160a01b031681565b3480156109e8575f80fd5b506104d3612b18565b3480156109fc575f80fd5b506104f1612b28565b348015610a10575f80fd5b506104f1610a1f3660046155cb565b612c15565b348015610a2f575f80fd5b506104d3610a3e36600461560a565b612cbd565b348015610a4e575f80fd5b505f54610687906001600160a01b031681565b348015610a6c575f80fd5b506104f1610a7b3660046155cb565b612e26565b348015610a8b575f80fd5b506104f160085481565b348015610aa0575f80fd5b506104f1610aaf36600461559d565b612ec4565b348015610abf575f80fd5b506104d3610ace3660046154da565b6130da565b348015610ade575f80fd5b506104f1610aed366004615505565b613232565b348015610afd575f80fd5b506104f1610b0c3660046154da565b6132cf565b348015610b1c575f80fd5b506104f1610b2b366004615505565b61332c565b348015610b3b575f80fd5b506104f17f000000000000000000000000000000000000000000000000000000000000000081565b348015610b6e575f80fd5b506009546108479060ff1681565b348015610b87575f80fd5b506104f1610b963660046154da565b6133c9565b348015610ba6575f80fd5b506104f1610bb53660046155cb565b613476565b348015610bc5575f80fd5b506104f160075481565b348015610bda575f80fd5b506104f1610be93660046155cb565b613513565b348015610bf9575f80fd5b506104d3610c08366004615636565b6135b0565b348015610c18575f80fd5b506104f1610c273660046154da565b613a7f565b348015610c37575f80fd5b506104d3610c463660046154da565b613b1d565b348015610c56575f80fd5b506104f1610c6536600461554a565b613b3a565b348015610c75575f80fd5b506104f1610c8436600461554a565b613bdb565b348015610c94575f80fd5b50600154610687906001600160a01b031681565b348015610cb3575f80fd5b506104f1610cc23660046154da565b613ca0565b348015610cd2575f80fd5b5060016104f1565b348015610ce5575f80fd5b506104f1610cf43660046154da565b613cba565b348015610d04575f80fd5b506104f1613d1a565b348015610d18575f80fd5b506104d3610d2736600461554a565b613dbe565b6104f1610d3a36600461554a565b613e13565b348015610d4a575f80fd5b506104d3610d59366004615656565b613e6a565b60095460ff1615610d9b576040517f8afc0b7700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60015474010000000000000000000000000000000000000000900460ff1615610d9b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610dfb611c13565b345f03610e1b57604051635435b28960e11b815260040160405180910390fd5b610e2434613ca0565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004015f604051808303818588803b158015610e7f575f80fd5b505af1158015610e91573d5f803e3d5ffd5b50506040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152602482018690527f00000000000000000000000000000000000000000000000000000000000000001693506340c10f19925060440190506020604051808303815f875af1158015610f1c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f409190615697565b50610f53826001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b03167fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d73484604051610fa0929190918252602082015260400190565b60405180910390a3919050565b5f805f610fb8614024565b6006546040516370a0823160e01b81523060048201529294509092505f916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611025573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061104991906156b0565b61105391906156f4565b60408051606081018252600b548152600c546020820152600d54918101919091529091505f9061108390856140ab565b60408051606081018252600e548152600f546020820152601054918101919091529091505f906110b390856140ab565b90505f82604001518360200151845f01516110ce9190615707565b6110d89190615707565b90505f811280156110f157506110ed81615726565b8410155b15611110576110ff81615726565b611109908561575c565b935061112d565b5f811215611120575f935061112d565b61112a81856156f4565b93505b6040820151602083015183515f929161114591615707565b61114f9190615707565b90505f8112156111785761116281615726565b61116c90866156f4565b98975050505050505050565b80851161118c575f97505050505050505090565b611196818661575c565b97505050505050505090565b6111aa6140e7565b60155460ff166111cd5760405163baf3f0f760e01b815260040160405180910390fd5b6012546011546016545f906111e39084906156f4565b90505f6111f08284614110565b6014549091505f90611211906112078760016156f4565b601354859061411e565b905080156112225761122281614201565b83820361128857611231614404565b601454600a5561123f611c13565b6040805160a0810182525f80825260208201819052918101829052606081018290526080018190526011819055601281905560138190556014556015805460ff191690556112d0565b6040805160a0810182528581526020810185905260135491810191909152601454606082015260016080909101819052601185905560128490556015805460ff191690911790555b5050505050565b5f807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611335573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061135991906156b0565b905080156113795761137461136c610fad565b849083614474565b61137b565b825b9392505050565b5f61138b610d5e565b611393610d9d565b61139b61219c565b8211156113a957505f919050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611406573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061142a91906156b0565b90508015611379576113748161143e610fad565b85919061448f565b919050565b5f611454610d5e565b61145c610d9d565b836001600160a01b038116158061147b57506001600160a01b03811630145b1561149957604051631e4ec46b60e01b815260040160405180910390fd5b836114ac816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b0316146114dc576040516282b42960e81b815260040160405180910390fd5b6114e987878760016144b2565b979650505050505050565b6115236040518060a001604052805f81526020015f81526020015f81526020015f81526020015f151581525090565b506040805160a0810182526011548152601254602082015260135491810191909152601454606082015260155460ff161515608082015290565b61157e60405180606001604052805f81526020015f81526020015f81525090565b5f611587614024565b5060408051606081018252600b548152600c546020820152600d54918101919091529091506115b690826140ab565b91505090565b5f6115c5610d5e565b6115cd610d9d565b836001600160a01b03811615806115ec57506001600160a01b03811630145b1561160a57604051631e4ec46b60e01b815260040160405180910390fd5b8361161d816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b03161461164d576040516282b42960e81b815260040160405180910390fd5b6114e987878760016144d3565b6116626140e7565b61166a611c13565b6009805460ff19166001179055565b5f6116826140e7565b7f0000000000000000000000000000000000000000000000000000000000000000826001600160a01b031663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116df573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061170391906156b0565b141580611713575060095460ff16155b156117315760405163baf3f0f760e01b815260040160405180910390fd5b61173c6108f56122d8565b6040516370a0823160e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a9059cbb90849083906370a0823190602401602060405180830381865afa1580156117a8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117cc91906156b0565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801561182c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118509190615697565b5050600654919050565b6118626140e7565b61186a611c13565b600755565b6118987f00000000000000000000000000000000000000000000000000000000000000006144f4565b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556118f97f00000000000000000000000000000000000000000000000000000000000000006145af565b600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b61193b6140e7565b610d9b614404565b61194b610d5e565b611953610d9d565b8061195d8161462f565b611965611c13565b670de0b6b3a76400008260800135101561199257604051635435b28960e11b815260040160405180910390fd5b816080013561199f612b28565b10156119d7576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6119e583602001356146e2565b905080604001515f03611a1f57438082526040820181905260208201805160808601359190611a159083906156f4565b905250611a7b9050565b5f611a2c82600a54614756565b509050836080013582602001818151611a4591906156f4565b915081815250505f80611a5c84600a54854361476e565b6040860182905290925090508015611a7757611a7781614201565b5050505b611acb817f000000000000000000000000000000000000000000000000000000000000000060208601357f0000000000000000000000000000000000000000000000000000000000000000614832565b826080013560065f828254611ae091906156f4565b909155505060405160808401358152602080850135917f044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b910160405180910390a26040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152608084013560248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a9059cbb906044016020604051808303815f875af1158015611ba7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bcb9190615697565b50505050565b611bd96140e7565b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60085443118015611c3f575060015474010000000000000000000000000000000000000000900460ff16155b15610d9b575f80611c4e614024565b60408051606081018252600e548152600f546020820152601054918101919091529193509150611c7e90826140ab565b8051600e55602080820151600f556040918201516010558151606081018352600b548152600c5491810191909152600d5491810191909152611cc090836140ab565b8051600b556020810151600c5560400151600d55436008557fa44f29ae5036974cc3cc64b097368e60ba4210a0fef291753e670437b9587bb282611d02610fad565b6006546040516370a0823160e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015611d67573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d8b91906156b0565b611d9c670de0b6b3a76400006112d7565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611df8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e1c91906156b0565b604080519687526020870195909552938501929092526060840152608083015260a082015260c00160405180910390a15050565b5f611e59610d5e565b611e61610d9d565b611e6a826112d7565b9050611e7461219c565b81111561144657505f919050565b611e8a6140e7565b611e92611c13565b601655565b611e9f611c13565b33611ec97f000000000000000000000000000000000000000000000000000000000000000061483e565b6001600160a01b031614611eef576040516282b42960e81b815260040160405180910390fd5b60065415611f105760405163baf3f0f760e01b815260040160405180910390fd5b436008556006819055600254604080517fb7dc128400000000000000000000000000000000000000000000000000000000815290515f926001600160a01b03169163b7dc12849160048083019260209291908290030181865afa158015611f79573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f9d91906156b0565b9050611fc860405180608001604052805f81526020015f81526020015f81526020015f151581525090565b5f8060015b84811161201c57611fdd816146e2565b9350611fe98451151590565b1561200a57611ffa84600a54614756565b50925061200783836156f4565b91505b806120148161576f565b915050611fcd565b50801561202c5761202c81614201565b604080518082018252601481527f524f555445525f504f4f4c5f555047524144455200000000000000000000000060209091015251630d37324f60e11b81527f792d10c60000000000000000000000000000000000000000000000000000000060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690631a6e649e90602401602060405180830381865afa1580156120e0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121049190615787565b6001600160a01b0316638287aa2d826040518263ffffffff1660e01b815260040161213191815260200190565b6020604051808303815f875af115801561214d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121719190615697565b156121835761217e614404565b6112d0565b60405163baf3f0f760e01b815260040160405180910390fd5b6040516370a0823160e01b81523060048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015612200573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061222491906156b0565b905090565b5f6122757f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006148be565b6020015192915050565b6122a060405180606001604052805f81526020015f81526020015f81525090565b5f6122a9614024565b60408051606081018252600e548152600f546020820152601054918101919091529092506115b69150826140ab565b5f806122e2614024565b60408051606081018252600e548152600f546020820152601054918101919091529092506115b6915061231590836140ab565b61496e565b612322611c13565b6040516370a0823160e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015612384573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123a891906156b0565b8111806123bb57506123b86122d8565b81115b156123f2576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051606081018252600e548152600f5460208201526010549181019190915261241d90826149b2565b8051600e556020810151600f55604001516010556001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb6124897f00000000000000000000000000000000000000000000000000000000000000006149e5565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602481018490526044016020604051808303815f875af11580156124eb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061250f9190615697565b5050565b5f61251c610d5e565b612524610d9d565b816001600160a01b038116158061254357506001600160a01b03811630145b1561256157604051631e4ec46b60e01b815260040160405180910390fd5b61256b8484614a65565b91505b5092915050565b61257d6140e7565b6118f97f00000000000000000000000000000000000000000000000000000000000000006145af565b6001546001600160a01b031633146125d0576040516282b42960e81b815260040160405180910390fd5b610d9b33614c4e565b5f805f80846125e78161462f565b6125ef610d9d565b6125f7611c13565b5f61260587602001356146e2565b90506126118151151590565b612647576040517f260d436000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805f836060015115612752576040516323b872dd60e01b815233600482015230602482015260808b013560448201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303815f875af11580156126c7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126eb9190615697565b506040805160808c01358082526020808301919091525f828401529151918c0135917f0dfb9710cc5cf2de9c8b20b6cc77867afbe42c7ff1408df657d97a904ccfdc9c9181900360600190a2600a5484604001515f80985098509850985050505050612aec565b43846040015110156127a55761276a84600a54614756565b909350915081612786670de0b6b3a764000060808d01356157a2565b10156127a557604051635435b28960e11b815260040160405180910390fd5b828a60800135101561280f575f6127c060808c013584614cc2565b90505f6127cd8285614cd6565b60808d013593509050828110156127f2579150816127ef8160808e013561575c565b97505b818660400181815161280491906156f4565b9052506128aa915050565b5f61281e8460808d013561575c565b90508391508460200151811061287157846020015197508760065f828254612846919061575c565b9091555050602085015161285a908261575c565b5f80875260208701819052604087015296506128a8565b8097508060065f828254612885919061575c565b92505081905550808560200181815161289e919061575c565b9052504360408601525b505b6128fa847f000000000000000000000000000000000000000000000000000000000000000060208d01357f0000000000000000000000000000000000000000000000000000000000000000614832565b60408051606081018252600b548152600c546020820152600d549181019190915261292590826149b2565b8051600b556020810151600c5560400151600d556001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166323b872dd33306129748b866156f4565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064016020604051808303815f875af11580156129dd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a019190615697565b506003546001600160a01b031615612a90576003546040517fcda43c1000000000000000000000000000000000000000000000000000000000815260208c01356004820152602481018390526001600160a01b039091169063cda43c10906044015f604051808303815f87803b158015612a79575f80fd5b505af1158015612a8b573d5f803e3d5ffd5b505050505b6040805160808c013581526020818101849052918101899052908b0135907f0dfb9710cc5cf2de9c8b20b6cc77867afbe42c7ff1408df657d97a904ccfdc9c9060600160405180910390a2600a54846040015198509850505050505b509193509193565b612afc6140e7565b4715612b1557612b156001600160a01b03821647614cea565b50565b612b206140e7565b610d9b614dad565b6009545f9060ff1680612b55575060015474010000000000000000000000000000000000000000900460ff165b15612b5f57505f90565b6040516370a0823160e01b81523060048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015612bc3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612be791906156b0565b90505f612bf2613d1a565b905080821015612c04575f9250505090565b612c0e818361575c565b9250505090565b5f612c1e610d5e565b612c26610d9d565b826001600160a01b0381161580612c4557506001600160a01b03811630145b15612c6357604051631e4ec46b60e01b815260040160405180910390fd5b82612c76816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b031614612ca6576040516282b42960e81b815260040160405180910390fd5b612cb386868660016144b2565b9695505050505050565b612cc56140e7565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603612d16576040516282b42960e81b815260040160405180910390fd5b806001600160a01b031663a9059cbb612d37846001600160a01b0316613fe0565b6040516370a0823160e01b81523060048201526001600160a01b038516906370a0823190602401602060405180830381865afa158015612d79573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d9d91906156b0565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015612dfd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e219190615697565b505050565b5f612e2f610d5e565b612e37610d9d565b826001600160a01b0381161580612e5657506001600160a01b03811630145b15612e7457604051631e4ec46b60e01b815260040160405180910390fd5b82612e87816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b031614612eb7576040516282b42960e81b815260040160405180910390fd5b612cb386868660016144d3565b5f612ecd610d5e565b612ed5610d9d565b816001600160a01b0381161580612ef457506001600160a01b03811630145b15612f1257604051631e4ec46b60e01b815260040160405180910390fd5b612f1a611c13565b612f23846133c9565b9150811580612f30575083155b15612f4e57604051635435b28960e11b815260040160405180910390fd5b6040516323b872dd60e01b8152336004820152306024820152604481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303815f875af1158015612fbe573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fe29190615697565b506040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018690527f000000000000000000000000000000000000000000000000000000000000000016906340c10f19906044016020604051808303815f875af1158015613068573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061308c9190615697565b5060408051838152602081018690526001600160a01b0385169133917fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7910160405180910390a35092915050565b6130e26140e7565b6130ea611c13565b60155460ff161561310e5760405163baf3f0f760e01b815260040160405180910390fd5b613116614dad565b600254604080517fb7dc128400000000000000000000000000000000000000000000000000000000815290515f926001600160a01b03169163b7dc12849160048083019260209291908290030181865afa158015613176573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061319a91906156b0565b90505f6131b58360016131af60165486614110565b4361411e565b905080156131c6576131c681614201565b81601654106131de576131d7614404565b5050600a55565b6040805160a08101825283815260165460208201819052439282018390526060820186905260016080909201829052601185905560125560139190915560148490556015805460ff19169091179055505050565b5f61323b610d5e565b613243610d9d565b836001600160a01b038116158061326257506001600160a01b03811630145b1561328057604051631e4ec46b60e01b815260040160405180910390fd5b83613293816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b0316146132c3576040516282b42960e81b815260040160405180910390fd5b6114e98787875f6144d3565b5f61332661331e7f0000000000000000000000000000000000000000000000000000000000000000847f00000000000000000000000000000000000000000000000000000000000000006148be565b600a54614e1c565b92915050565b5f613335610d5e565b61333d610d9d565b836001600160a01b038116158061335c57506001600160a01b03811630145b1561337a57604051631e4ec46b60e01b815260040160405180910390fd5b8361338d816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b0316146133bd576040516282b42960e81b815260040160405180910390fd5b6114e98787875f6144b2565b5f6133d2610d5e565b6133da610d9d565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613437573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061345b91906156b0565b905080156113795761137461346e610fad565b84908361448f565b5f61347f610d5e565b613487610d9d565b826001600160a01b03811615806134a657506001600160a01b03811630145b156134c457604051631e4ec46b60e01b815260040160405180910390fd5b826134d7816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b031614613507576040516282b42960e81b815260040160405180910390fd5b612cb38686865f6144b2565b5f61351c610d5e565b613524610d9d565b826001600160a01b038116158061354357506001600160a01b03811630145b1561356157604051631e4ec46b60e01b815260040160405180910390fd5b82613574816001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b0316146135a4576040516282b42960e81b815260040160405180910390fd5b612cb38686865f6144d3565b6135b8614e3b565b6135c0611c13565b5f6135ca836146e2565b9050806060015115613608576040517ff29cb5e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001606082015260408101515f9081904311156136405761362b83600a54614756565b5060055490925061363d908390614e8b565b90505b5f83602001518361365191906156f4565b90505f805f808689116137105788935061366b848861575c565b60208901516040516323b872dd60e01b8152336004820152306024820152604481018c905291935091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064015b6020604051808303815f875af11580156136e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061370a9190615697565b50613890565b61371a868661575c565b891161379b5786935061372d848a61575c565b925082886020015161373f919061575c565b6040516323b872dd60e01b8152336004820152306024820152604481018b90529091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016136ca565b869350876020015192507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166323b872dd3330866137f7600554670de0b6b3a76400006137f0919061575c565b8a90614cd6565b61380191906156f4565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064016020604051808303815f875af115801561386a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061388e9190615697565b505b876020015160065f8282546138a5919061575c565b909155505060408051606081018252600b548152600c546020820152600d54918101919091526138e19083906138db90876149b2565b90614e9f565b8051600b556020810151600c5560400151600d556005545f90613905908990614e8b565b60408051606081018252600e548152600f546020820152601054918101919091529091506139339082614e9f565b8051600e55602080820151600f556040909101516010555f908a015261399b897f00000000000000000000000000000000000000000000000000000000000000008d7f0000000000000000000000000000000000000000000000000000000000000000614832565b6003546001600160a01b031615613a1f576003546040517f8b446431000000000000000000000000000000000000000000000000000000008152600481018d90526001600160a01b0390911690638b446431906024015f604051808303815f87803b158015613a08575f80fd5b505af1158015613a1a573d5f803e3d5ffd5b505050505b8a7ff02514bceabf2b74024e8529fdc9f423bebda69624841ff163ebd712c188f4a78b613a4c85876156f4565b60408051928352602083019190915281018490526060810188905260800160405180910390a25050505050505050505050565b5f807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613add573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b0191906156b0565b905080156113795761137481613b15610fad565b859190614474565b613b256140e7565b613b2d610d9d565b613b35611c13565b600555565b6040516370a0823160e01b81526001600160a01b0382811660048301525f9161332691613bce917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015613baa573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061053291906156b0565b613bd661219c565b614110565b6040516370a0823160e01b81526001600160a01b0382811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015613c43573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c6791906156b0565b90505f613c73826112d7565b90505f613c7e61219c565b905080821115613c9957613c9181611e50565b949350505050565b5050919050565b5f613ca9610d5e565b613cb1610d9d565b61332682613a7f565b5f80613d12613d0a7f0000000000000000000000000000000000000000000000000000000000000000857f00000000000000000000000000000000000000000000000000000000000000006148be565b600a54614756565b509392505050565b6007546006546040516370a0823160e01b81523060048201525f926122249290916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015613d8a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613dae91906156b0565b613db891906156f4565b90614cd6565b613dc66140e7565b613dd8816001600160a01b0316613fe0565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905550565b5f613e1c610d5e565b613e24610d9d565b816001600160a01b0381161580613e4357506001600160a01b03811630145b15613e6157604051631e4ec46b60e01b815260040160405180910390fd5b61137b83610df2565b613e726140e7565b613e7a611c13565b5f613e84836146e2565b905080602001515f14613eaa5760405163baf3f0f760e01b815260040160405180910390fd5b60208101829052438082526040820152613f06817f0000000000000000000000000000000000000000000000000000000000000000857f0000000000000000000000000000000000000000000000000000000000000000614832565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166340c10f1985613f3f85613a7f565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015613f9f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613fc39190615697565b508160065f828254613fd591906156f4565b909155505050505050565b5f805f613fec84614ed2565b9150915081613ffd57509192915050565b5f8061400883614f1e565b915091508161401b575093949350505050565b95945050505050565b5f8060085443118015614052575060015474010000000000000000000000000000000000000000900460ff16155b156140a35761408b60405180608001604052806008548152602001600654815260200160085481526020015f1515815250600a54614756565b5060055490925061409d908390614e8b565b90509091565b505f91829150565b6140cc60405180606001604052805f81526020015f81526020015f81525090565b81835f018181516140dd91906156f4565b9052509192915050565b5f546001600160a01b03163314610d9b576040516282b42960e81b815260040160405180910390fd5b5f818310611379578161137b565b604080516080810182525f80825260208201819052918101829052606081018290525f8080875b8781116141f457614155816146e2565b6020810151909550156141e25761416f85600a5489614f8e565b50935061417e858b868a61476e565b909350915061418d82876156f4565b6040860184905295506141e2857f0000000000000000000000000000000000000000000000000000000000000000837f0000000000000000000000000000000000000000000000000000000000000000614832565b806141ec8161576f565b915050614145565b5050505050949350505050565b60408051606081018252600b548152600c546020820152600d549181019190915261422c90826140ab565b8051600b556020810151600c5560400151600d5560055461427990614252908390614e8b565b60408051606081018252600e548152600f54602082015260105491810191909152906140ab565b8051600e556020810151600f55604001516010557fa44f29ae5036974cc3cc64b097368e60ba4210a0fef291753e670437b9587bb2816142b7610fad565b6006546040516370a0823160e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa15801561431c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061434091906156b0565b614351670de0b6b3a76400006112d7565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156143ad573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906143d191906156b0565b604080519687526020870195909552938501929092526060840152608083015260a082015260c00160405180910390a150565b61440c614fc4565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b5f825f190484118302158202614488575f80fd5b5091020490565b5f825f1904841183021582026144a3575f80fd5b50910281810615159190040190565b5f6144bb611c13565b6144c485611382565b9050613c918385838886615018565b5f6144dc611c13565b6144e585611e50565b9050613c918385878486615018565b604080518082018252601481527f524f555445525f4147454e545f464143544f525900000000000000000000000060209091015251630d37324f60e11b81527f29f616da0000000000000000000000000000000000000000000000000000000060048201525f906001600160a01b03831690631a6e649e906024015b602060405180830381865afa15801561458b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133269190615787565b604080518082018252601281527f524f555445525f435245445f504152534552000000000000000000000000000060209091015251630d37324f60e11b81527fd26df3b70000000000000000000000000000000000000000000000000000000060048201525f906001600160a01b03831690631a6e649e90602401614570565b602081013515806146c557506002546040517ffd66091e0000000000000000000000000000000000000000000000000000000081523360048201526020830135916001600160a01b03169063fd66091e90602401602060405180830381865afa15801561469e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906146c291906156b0565b14155b15612b15576040516282b42960e81b815260040160405180910390fd5b61470b60405180608001604052805f81526020015f81526020015f81526020015f151581525090565b6133267f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006148be565b5f80614763848443614f8e565b915091509250929050565b5f805f61477b8787615393565b90505f851180156147965750614792816001614cd6565b8511155b156147c6576147a660018561575c565b856147b2836001614cd6565b6147bc919061575c565b9250925050614829565b841561481f576147d685826153a4565b6147e0908561575c565b6040880181905292505f6147f48888614756565b5090508581146148145783614809878361575c565b935093505050614829565b505f91506148299050565b50505060408401515f5b94509492505050565b611bcb838383876153b8565b604080518082018252601481527f524f555445525f504f4f4c5f524547495354525900000000000000000000000060209091015251630d37324f60e11b81527f390c2e820000000000000000000000000000000000000000000000000000000060048201525f906001600160a01b03831690631a6e649e90602401614570565b6148e760405180608001604052805f81526020015f81526020015f81526020015f151581525090565b6040517f6361f6de00000000000000000000000000000000000000000000000000000000815260048101849052602481018390526001600160a01b03851690636361f6de90604401608060405180830381865afa15801561494a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c9191906157b9565b8051604082015160208301515f9291614986916156f4565b1061499257505f919050565b6040820151602083015183516149a8919061575c565b613326919061575c565b6149d360405180606001604052805f81526020015f81526020015f81525090565b81836020018181516140dd91906156f4565b604080518082018252600f81527f524f555445525f5452454153555259000000000000000000000000000000000060209091015251630d37324f60e11b81527fa1858d5f0000000000000000000000000000000000000000000000000000000060048201525f906001600160a01b03831690631a6e649e90602401614570565b5f614a6e611c13565b614a7783613ca0565b9050821580614a84575080155b15614aa257604051635435b28960e11b815260040160405180910390fd5b6040516323b872dd60e01b8152336004820152306024820152604481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303815f875af1158015614b12573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614b369190615697565b506040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018390527f000000000000000000000000000000000000000000000000000000000000000016906340c10f19906044016020604051808303815f875af1158015614bbc573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614be09190615697565b50614bf3826001600160a01b0316613fe0565b6001600160a01b0316336001600160a01b03167fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d78584604051614c40929190918252602082015260400190565b60405180910390a392915050565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055614c886001600160a01b038216613fe0565b5f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905550565b5f61137b83670de0b6b3a764000084614474565b5f61137b8383670de0b6b3a764000061448f565b80471015614d24576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114614d6d576040519150601f19603f3d011682016040523d82523d5f602084013e614d72565b606091505b5050905080612e21576040517f3204506f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614db5610d9d565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586144573390565b5f80614e288484614756565b50905080846020015161256b91906156f4565b33614e657f000000000000000000000000000000000000000000000000000000000000000061545a565b6001600160a01b031614610d9b576040516282b42960e81b815260040160405180910390fd5b5f61137b8383670de0b6b3a7640000614474565b614ec060405180606001604052805f81526020015f81526020015f81525090565b81836040018181516140dd91906156f4565b5f8073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103614f18576001925067ffffffffffffffff841691505b50915091565b5f80825f526016600a60205f73fe000000000000000000000000000000000000025afa91505f51806001600160a01b031691508060a01c61ffff16905061040a8114614f6b575f92505f91505b50811580614f7a57503d601614155b15614f8957505f928392509050565b915091565b5f805f856040015184614fa1919061575c565b9050614fad8686615393565b9150614fb98282614cd6565b925050935093915050565b60015474010000000000000000000000000000000000000000900460ff16610d9b576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b821580615023575081155b1561504157604051635435b28960e11b815260040160405180910390fd5b615053846001600160a01b0316613fe0565b9350615067856001600160a01b0316613fe0565b945061507161219c565b8211156150aa576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516323b872dd60e01b81526001600160a01b038681166004830152306024830152604482018590527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303815f875af115801561511c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906151409190615697565b506040517f9dc29fac000000000000000000000000000000000000000000000000000000008152306004820152602481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639dc29fac906044016020604051808303815f875af11580156151c4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906151e89190615697565b508015615296576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d906024015f604051808303815f87803b158015615267575f80fd5b505af1158015615279573d5f803e3d5ffd5b50615291925050506001600160a01b03851683614cea565b615341565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018490527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303815f875af115801561531b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061533f9190615697565b505b60408051838152602081018590526001600160a01b03808816929087169183917ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db910160405180910390a45050505050565b60208201515f9061137b9083614cd6565b5f61137b83670de0b6b3a76400008461448f565b604080517fc7afd04b000000000000000000000000000000000000000000000000000000008152600481018590526024810184905282516044820152602083015160648201529082015160848201526060820151151560a48201526001600160a01b0385169063c7afd04b9060c4015f604051808303815f87803b15801561543e575f80fd5b505af1158015615450573d5f803e3d5ffd5b5050505050505050565b604080518082018252601381527f524f555445525f4147454e545f504f4c4943450000000000000000000000000060209091015251630d37324f60e11b81527f6ea276a30000000000000000000000000000000000000000000000000000000060048201525f906001600160a01b03831690631a6e649e90602401614570565b5f602082840312156154ea575f80fd5b5035919050565b6001600160a01b0381168114612b15575f80fd5b5f805f8060808587031215615518575f80fd5b84359350602085013561552a816154f1565b9250604085013561553a816154f1565b9396929550929360600135925050565b5f6020828403121561555a575f80fd5b813561137b816154f1565b5f60208284031215615575575f80fd5b813567ffffffffffffffff81111561558b575f80fd5b8201610100818503121561137b575f80fd5b5f80604083850312156155ae575f80fd5b8235915060208301356155c0816154f1565b809150509250929050565b5f805f606084860312156155dd575f80fd5b8335925060208401356155ef816154f1565b915060408401356155ff816154f1565b809150509250925092565b5f806040838503121561561b575f80fd5b8235615626816154f1565b915060208301356155c0816154f1565b5f8060408385031215615647575f80fd5b50508035926020909101359150565b5f805f60608486031215615668575f80fd5b8335615673816154f1565b95602085013595506040909401359392505050565b80518015158114611446575f80fd5b5f602082840312156156a7575f80fd5b61137b82615688565b5f602082840312156156c0575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115613326576133266156c7565b8181035f83128015838313168383128216171561256e5761256e6156c7565b5f7f80000000000000000000000000000000000000000000000000000000000000008203615756576157566156c7565b505f0390565b81810381811115613326576133266156c7565b5f5f198203615780576157806156c7565b5060010190565b5f60208284031215615797575f80fd5b815161137b816154f1565b8082028115828204841417613326576133266156c7565b5f608082840312156157c9575f80fd5b6040516080810181811067ffffffffffffffff82111715615811577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b806040525082518152602083015160208201526040830151604082015261583a60608401615688565b6060820152939250505056fea2646970667358221220eb0d2fecb1b872502e43b88df5e2b17c07a567c1f0ead651796b90b79710bdb564736f6c63430008150033",
}

// InfinityPoolV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use InfinityPoolV2MetaData.ABI instead.
var InfinityPoolV2ABI = InfinityPoolV2MetaData.ABI

// InfinityPoolV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InfinityPoolV2MetaData.Bin instead.
var InfinityPoolV2Bin = InfinityPoolV2MetaData.Bin

// DeployInfinityPoolV2 deploys a new Ethereum contract, binding an instance of InfinityPoolV2 to it.
func DeployInfinityPoolV2(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, router_ common.Address, liquidStakingToken_ common.Address, minimumLiquidity_ *big.Int, id_ *big.Int) (common.Address, *types.Transaction, *InfinityPoolV2, error) {
	parsed, err := InfinityPoolV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InfinityPoolV2Bin), backend, owner_, router_, liquidStakingToken_, minimumLiquidity_, id_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InfinityPoolV2{InfinityPoolV2Caller: InfinityPoolV2Caller{contract: contract}, InfinityPoolV2Transactor: InfinityPoolV2Transactor{contract: contract}, InfinityPoolV2Filterer: InfinityPoolV2Filterer{contract: contract}}, nil
}

// InfinityPoolV2 is an auto generated Go binding around an Ethereum contract.
type InfinityPoolV2 struct {
	InfinityPoolV2Caller     // Read-only binding to the contract
	InfinityPoolV2Transactor // Write-only binding to the contract
	InfinityPoolV2Filterer   // Log filterer for contract events
}

// InfinityPoolV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type InfinityPoolV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type InfinityPoolV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfinityPoolV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfinityPoolV2Session struct {
	Contract     *InfinityPoolV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfinityPoolV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfinityPoolV2CallerSession struct {
	Contract *InfinityPoolV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// InfinityPoolV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfinityPoolV2TransactorSession struct {
	Contract     *InfinityPoolV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InfinityPoolV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type InfinityPoolV2Raw struct {
	Contract *InfinityPoolV2 // Generic contract binding to access the raw methods on
}

// InfinityPoolV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfinityPoolV2CallerRaw struct {
	Contract *InfinityPoolV2Caller // Generic read-only contract binding to access the raw methods on
}

// InfinityPoolV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfinityPoolV2TransactorRaw struct {
	Contract *InfinityPoolV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewInfinityPoolV2 creates a new instance of InfinityPoolV2, bound to a specific deployed contract.
func NewInfinityPoolV2(address common.Address, backend bind.ContractBackend) (*InfinityPoolV2, error) {
	contract, err := bindInfinityPoolV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2{InfinityPoolV2Caller: InfinityPoolV2Caller{contract: contract}, InfinityPoolV2Transactor: InfinityPoolV2Transactor{contract: contract}, InfinityPoolV2Filterer: InfinityPoolV2Filterer{contract: contract}}, nil
}

// NewInfinityPoolV2Caller creates a new read-only instance of InfinityPoolV2, bound to a specific deployed contract.
func NewInfinityPoolV2Caller(address common.Address, caller bind.ContractCaller) (*InfinityPoolV2Caller, error) {
	contract, err := bindInfinityPoolV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2Caller{contract: contract}, nil
}

// NewInfinityPoolV2Transactor creates a new write-only instance of InfinityPoolV2, bound to a specific deployed contract.
func NewInfinityPoolV2Transactor(address common.Address, transactor bind.ContractTransactor) (*InfinityPoolV2Transactor, error) {
	contract, err := bindInfinityPoolV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2Transactor{contract: contract}, nil
}

// NewInfinityPoolV2Filterer creates a new log filterer instance of InfinityPoolV2, bound to a specific deployed contract.
func NewInfinityPoolV2Filterer(address common.Address, filterer bind.ContractFilterer) (*InfinityPoolV2Filterer, error) {
	contract, err := bindInfinityPoolV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2Filterer{contract: contract}, nil
}

// bindInfinityPoolV2 binds a generic wrapper to an already deployed contract.
func bindInfinityPoolV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InfinityPoolV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfinityPoolV2 *InfinityPoolV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfinityPoolV2.Contract.InfinityPoolV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfinityPoolV2 *InfinityPoolV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.InfinityPoolV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfinityPoolV2 *InfinityPoolV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.InfinityPoolV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfinityPoolV2 *InfinityPoolV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfinityPoolV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfinityPoolV2 *InfinityPoolV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfinityPoolV2 *InfinityPoolV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Caller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "allPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Session) AllPools(arg0 *big.Int) (common.Address, error) {
	return _InfinityPoolV2.Contract.AllPools(&_InfinityPoolV2.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _InfinityPoolV2.Contract.AllPools(&_InfinityPoolV2.CallOpts, arg0)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) AllPoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "allPoolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) AllPoolsLength() (*big.Int, error) {
	return _InfinityPoolV2.Contract.AllPoolsLength(&_InfinityPoolV2.CallOpts)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) AllPoolsLength() (*big.Int, error) {
	return _InfinityPoolV2.Contract.AllPoolsLength(&_InfinityPoolV2.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Session) Asset() (common.Address, error) {
	return _InfinityPoolV2.Contract.Asset(&_InfinityPoolV2.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) Asset() (common.Address, error) {
	return _InfinityPoolV2.Contract.Asset(&_InfinityPoolV2.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.ConvertToAssets(&_InfinityPoolV2.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.ConvertToAssets(&_InfinityPoolV2.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.ConvertToShares(&_InfinityPoolV2.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.ConvertToShares(&_InfinityPoolV2.CallOpts, assets)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Caller) CredParser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "credParser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Session) CredParser() (common.Address, error) {
	return _InfinityPoolV2.Contract.CredParser(&_InfinityPoolV2.CallOpts)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) CredParser() (common.Address, error) {
	return _InfinityPoolV2.Contract.CredParser(&_InfinityPoolV2.CallOpts)
}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) GetAbsMinLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "getAbsMinLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) GetAbsMinLiquidity() (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAbsMinLiquidity(&_InfinityPoolV2.CallOpts)
}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) GetAbsMinLiquidity() (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAbsMinLiquidity(&_InfinityPoolV2.CallOpts)
}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) GetAgentBorrowed(opts *bind.CallOpts, agentID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "getAgentBorrowed", agentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) GetAgentBorrowed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAgentBorrowed(&_InfinityPoolV2.CallOpts, agentID)
}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) GetAgentBorrowed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAgentBorrowed(&_InfinityPoolV2.CallOpts, agentID)
}

// GetAgentDebt is a free data retrieval call binding the contract method 0x9f4326d7.
//
// Solidity: function getAgentDebt(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) GetAgentDebt(opts *bind.CallOpts, agentID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "getAgentDebt", agentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentDebt is a free data retrieval call binding the contract method 0x9f4326d7.
//
// Solidity: function getAgentDebt(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) GetAgentDebt(agentID *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAgentDebt(&_InfinityPoolV2.CallOpts, agentID)
}

// GetAgentDebt is a free data retrieval call binding the contract method 0x9f4326d7.
//
// Solidity: function getAgentDebt(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) GetAgentDebt(agentID *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAgentDebt(&_InfinityPoolV2.CallOpts, agentID)
}

// GetAgentInterestOwed is a free data retrieval call binding the contract method 0xf086ce60.
//
// Solidity: function getAgentInterestOwed(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) GetAgentInterestOwed(opts *bind.CallOpts, agentID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "getAgentInterestOwed", agentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentInterestOwed is a free data retrieval call binding the contract method 0xf086ce60.
//
// Solidity: function getAgentInterestOwed(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) GetAgentInterestOwed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAgentInterestOwed(&_InfinityPoolV2.CallOpts, agentID)
}

// GetAgentInterestOwed is a free data retrieval call binding the contract method 0xf086ce60.
//
// Solidity: function getAgentInterestOwed(uint256 agentID) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) GetAgentInterestOwed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetAgentInterestOwed(&_InfinityPoolV2.CallOpts, agentID)
}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) GetLiquidAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "getLiquidAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) GetLiquidAssets() (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetLiquidAssets(&_InfinityPoolV2.CallOpts)
}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) GetLiquidAssets() (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetLiquidAssets(&_InfinityPoolV2.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256 rate)
func (_InfinityPoolV2 *InfinityPoolV2Caller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256 rate)
func (_InfinityPoolV2 *InfinityPoolV2Session) GetRate() (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetRate(&_InfinityPoolV2.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256 rate)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) GetRate() (*big.Int, error) {
	return _InfinityPoolV2.Contract.GetRate(&_InfinityPoolV2.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) Id() (*big.Int, error) {
	return _InfinityPoolV2.Contract.Id(&_InfinityPoolV2.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) Id() (*big.Int, error) {
	return _InfinityPoolV2.Contract.Id(&_InfinityPoolV2.CallOpts)
}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPoolV2 *InfinityPoolV2Caller) IsShuttingDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "isShuttingDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPoolV2 *InfinityPoolV2Session) IsShuttingDown() (bool, error) {
	return _InfinityPoolV2.Contract.IsShuttingDown(&_InfinityPoolV2.CallOpts)
}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) IsShuttingDown() (bool, error) {
	return _InfinityPoolV2.Contract.IsShuttingDown(&_InfinityPoolV2.CallOpts)
}

// LastAccountingUpdateEpoch is a free data retrieval call binding the contract method 0x90fd6bc6.
//
// Solidity: function lastAccountingUpdateEpoch() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) LastAccountingUpdateEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "lastAccountingUpdateEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAccountingUpdateEpoch is a free data retrieval call binding the contract method 0x90fd6bc6.
//
// Solidity: function lastAccountingUpdateEpoch() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) LastAccountingUpdateEpoch() (*big.Int, error) {
	return _InfinityPoolV2.Contract.LastAccountingUpdateEpoch(&_InfinityPoolV2.CallOpts)
}

// LastAccountingUpdateEpoch is a free data retrieval call binding the contract method 0x90fd6bc6.
//
// Solidity: function lastAccountingUpdateEpoch() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) LastAccountingUpdateEpoch() (*big.Int, error) {
	return _InfinityPoolV2.Contract.LastAccountingUpdateEpoch(&_InfinityPoolV2.CallOpts)
}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Caller) LiquidStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "liquidStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Session) LiquidStakingToken() (common.Address, error) {
	return _InfinityPoolV2.Contract.LiquidStakingToken(&_InfinityPoolV2.CallOpts)
}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) LiquidStakingToken() (common.Address, error) {
	return _InfinityPoolV2.Contract.LiquidStakingToken(&_InfinityPoolV2.CallOpts)
}

// Lm is a free data retrieval call binding the contract method 0x803ba97e.
//
// Solidity: function lm() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Caller) Lm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "lm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lm is a free data retrieval call binding the contract method 0x803ba97e.
//
// Solidity: function lm() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Session) Lm() (common.Address, error) {
	return _InfinityPoolV2.Contract.Lm(&_InfinityPoolV2.CallOpts)
}

// Lm is a free data retrieval call binding the contract method 0x803ba97e.
//
// Solidity: function lm() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) Lm() (common.Address, error) {
	return _InfinityPoolV2.Contract.Lm(&_InfinityPoolV2.CallOpts)
}

// LpRewards is a free data retrieval call binding the contract method 0x0c762141.
//
// Solidity: function lpRewards() view returns((uint256,uint256,uint256))
func (_InfinityPoolV2 *InfinityPoolV2Caller) LpRewards(opts *bind.CallOpts) (RewardAccrual, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "lpRewards")

	if err != nil {
		return *new(RewardAccrual), err
	}

	out0 := *abi.ConvertType(out[0], new(RewardAccrual)).(*RewardAccrual)

	return out0, err

}

// LpRewards is a free data retrieval call binding the contract method 0x0c762141.
//
// Solidity: function lpRewards() view returns((uint256,uint256,uint256))
func (_InfinityPoolV2 *InfinityPoolV2Session) LpRewards() (RewardAccrual, error) {
	return _InfinityPoolV2.Contract.LpRewards(&_InfinityPoolV2.CallOpts)
}

// LpRewards is a free data retrieval call binding the contract method 0x0c762141.
//
// Solidity: function lpRewards() view returns((uint256,uint256,uint256))
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) LpRewards() (RewardAccrual, error) {
	return _InfinityPoolV2.Contract.LpRewards(&_InfinityPoolV2.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxDeposit(&_InfinityPoolV2.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxDeposit(&_InfinityPoolV2.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxMint(&_InfinityPoolV2.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxMint(&_InfinityPoolV2.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Caller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxRedeem(&_InfinityPoolV2.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxRedeem(&_InfinityPoolV2.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxWithdraw(&_InfinityPoolV2.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _InfinityPoolV2.Contract.MaxWithdraw(&_InfinityPoolV2.CallOpts, owner)
}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) MinimumLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "minimumLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) MinimumLiquidity() (*big.Int, error) {
	return _InfinityPoolV2.Contract.MinimumLiquidity(&_InfinityPoolV2.CallOpts)
}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) MinimumLiquidity() (*big.Int, error) {
	return _InfinityPoolV2.Contract.MinimumLiquidity(&_InfinityPoolV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Session) Owner() (common.Address, error) {
	return _InfinityPoolV2.Contract.Owner(&_InfinityPoolV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) Owner() (common.Address, error) {
	return _InfinityPoolV2.Contract.Owner(&_InfinityPoolV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfinityPoolV2 *InfinityPoolV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfinityPoolV2 *InfinityPoolV2Session) Paused() (bool, error) {
	return _InfinityPoolV2.Contract.Paused(&_InfinityPoolV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) Paused() (bool, error) {
	return _InfinityPoolV2.Contract.Paused(&_InfinityPoolV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2Session) PendingOwner() (common.Address, error) {
	return _InfinityPoolV2.Contract.PendingOwner(&_InfinityPoolV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) PendingOwner() (common.Address, error) {
	return _InfinityPoolV2.Contract.PendingOwner(&_InfinityPoolV2.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewDeposit(&_InfinityPoolV2.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewDeposit(&_InfinityPoolV2.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewMint(&_InfinityPoolV2.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewMint(&_InfinityPoolV2.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Caller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Session) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewRedeem(&_InfinityPoolV2.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewRedeem(&_InfinityPoolV2.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Caller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewWithdraw(&_InfinityPoolV2.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _InfinityPoolV2.Contract.PreviewWithdraw(&_InfinityPoolV2.CallOpts, assets)
}

// RateUpdate is a free data retrieval call binding the contract method 0x0c7290f2.
//
// Solidity: function rateUpdate() view returns((uint256,uint256,uint256,uint256,bool))
func (_InfinityPoolV2 *InfinityPoolV2Caller) RateUpdate(opts *bind.CallOpts) (RateUpdate, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "rateUpdate")

	if err != nil {
		return *new(RateUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(RateUpdate)).(*RateUpdate)

	return out0, err

}

// RateUpdate is a free data retrieval call binding the contract method 0x0c7290f2.
//
// Solidity: function rateUpdate() view returns((uint256,uint256,uint256,uint256,bool))
func (_InfinityPoolV2 *InfinityPoolV2Session) RateUpdate() (RateUpdate, error) {
	return _InfinityPoolV2.Contract.RateUpdate(&_InfinityPoolV2.CallOpts)
}

// RateUpdate is a free data retrieval call binding the contract method 0x0c7290f2.
//
// Solidity: function rateUpdate() view returns((uint256,uint256,uint256,uint256,bool))
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) RateUpdate() (RateUpdate, error) {
	return _InfinityPoolV2.Contract.RateUpdate(&_InfinityPoolV2.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) TotalAssets() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TotalAssets(&_InfinityPoolV2.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) TotalAssets() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TotalAssets(&_InfinityPoolV2.CallOpts)
}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) TotalBorrowableAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "totalBorrowableAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) TotalBorrowableAssets() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TotalBorrowableAssets(&_InfinityPoolV2.CallOpts)
}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) TotalBorrowableAssets() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TotalBorrowableAssets(&_InfinityPoolV2.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) TotalBorrowed() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TotalBorrowed(&_InfinityPoolV2.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) TotalBorrowed() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TotalBorrowed(&_InfinityPoolV2.CallOpts)
}

// TreasuryFeeRate is a free data retrieval call binding the contract method 0x66601032.
//
// Solidity: function treasuryFeeRate() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) TreasuryFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "treasuryFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryFeeRate is a free data retrieval call binding the contract method 0x66601032.
//
// Solidity: function treasuryFeeRate() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) TreasuryFeeRate() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TreasuryFeeRate(&_InfinityPoolV2.CallOpts)
}

// TreasuryFeeRate is a free data retrieval call binding the contract method 0x66601032.
//
// Solidity: function treasuryFeeRate() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) TreasuryFeeRate() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TreasuryFeeRate(&_InfinityPoolV2.CallOpts)
}

// TreasuryFeesOwed is a free data retrieval call binding the contract method 0x6b285af1.
//
// Solidity: function treasuryFeesOwed() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Caller) TreasuryFeesOwed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "treasuryFeesOwed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryFeesOwed is a free data retrieval call binding the contract method 0x6b285af1.
//
// Solidity: function treasuryFeesOwed() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2Session) TreasuryFeesOwed() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TreasuryFeesOwed(&_InfinityPoolV2.CallOpts)
}

// TreasuryFeesOwed is a free data retrieval call binding the contract method 0x6b285af1.
//
// Solidity: function treasuryFeesOwed() view returns(uint256)
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) TreasuryFeesOwed() (*big.Int, error) {
	return _InfinityPoolV2.Contract.TreasuryFeesOwed(&_InfinityPoolV2.CallOpts)
}

// TreasuryRewards is a free data retrieval call binding the contract method 0x6af1a7e8.
//
// Solidity: function treasuryRewards() view returns((uint256,uint256,uint256))
func (_InfinityPoolV2 *InfinityPoolV2Caller) TreasuryRewards(opts *bind.CallOpts) (RewardAccrual, error) {
	var out []interface{}
	err := _InfinityPoolV2.contract.Call(opts, &out, "treasuryRewards")

	if err != nil {
		return *new(RewardAccrual), err
	}

	out0 := *abi.ConvertType(out[0], new(RewardAccrual)).(*RewardAccrual)

	return out0, err

}

// TreasuryRewards is a free data retrieval call binding the contract method 0x6af1a7e8.
//
// Solidity: function treasuryRewards() view returns((uint256,uint256,uint256))
func (_InfinityPoolV2 *InfinityPoolV2Session) TreasuryRewards() (RewardAccrual, error) {
	return _InfinityPoolV2.Contract.TreasuryRewards(&_InfinityPoolV2.CallOpts)
}

// TreasuryRewards is a free data retrieval call binding the contract method 0x6af1a7e8.
//
// Solidity: function treasuryRewards() view returns((uint256,uint256,uint256))
func (_InfinityPoolV2 *InfinityPoolV2CallerSession) TreasuryRewards() (RewardAccrual, error) {
	return _InfinityPoolV2.Contract.TreasuryRewards(&_InfinityPoolV2.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.AcceptOwnership(&_InfinityPoolV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.AcceptOwnership(&_InfinityPoolV2.TransactOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Borrow(opts *bind.TransactOpts, vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "borrow", vc)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) Borrow(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Borrow(&_InfinityPoolV2.TransactOpts, vc)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Borrow(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Borrow(&_InfinityPoolV2.TransactOpts, vc)
}

// ContinueRateUpdate is a paid mutator transaction binding the contract method 0x0521a241.
//
// Solidity: function continueRateUpdate() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) ContinueRateUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "continueRateUpdate")
}

// ContinueRateUpdate is a paid mutator transaction binding the contract method 0x0521a241.
//
// Solidity: function continueRateUpdate() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) ContinueRateUpdate() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.ContinueRateUpdate(&_InfinityPoolV2.TransactOpts)
}

// ContinueRateUpdate is a paid mutator transaction binding the contract method 0x0521a241.
//
// Solidity: function continueRateUpdate() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) ContinueRateUpdate() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.ContinueRateUpdate(&_InfinityPoolV2.TransactOpts)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) DecommissionPool(opts *bind.TransactOpts, newPool common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "decommissionPool", newPool)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPoolV2 *InfinityPoolV2Session) DecommissionPool(newPool common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.DecommissionPool(&_InfinityPoolV2.TransactOpts, newPool)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) DecommissionPool(newPool common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.DecommissionPool(&_InfinityPoolV2.TransactOpts, newPool)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Deposit(&_InfinityPoolV2.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Deposit(&_InfinityPoolV2.TransactOpts, assets, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Deposit0(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "deposit0", receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) Deposit0(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Deposit0(&_InfinityPoolV2.TransactOpts, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Deposit0(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Deposit0(&_InfinityPoolV2.TransactOpts, receiver)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) HarvestFees(opts *bind.TransactOpts, harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "harvestFees", harvestAmount)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) HarvestFees(harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.HarvestFees(&_InfinityPoolV2.TransactOpts, harvestAmount)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) HarvestFees(harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.HarvestFees(&_InfinityPoolV2.TransactOpts, harvestAmount)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) JumpStartAccount(opts *bind.TransactOpts, receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "jumpStartAccount", receiver, agentID, accountPrincipal)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) JumpStartAccount(receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.JumpStartAccount(&_InfinityPoolV2.TransactOpts, receiver, agentID, accountPrincipal)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) JumpStartAccount(receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.JumpStartAccount(&_InfinityPoolV2.TransactOpts, receiver, agentID, accountPrincipal)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) JumpStartTotalBorrowed(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "jumpStartTotalBorrowed", amount)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) JumpStartTotalBorrowed(amount *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.JumpStartTotalBorrowed(&_InfinityPoolV2.TransactOpts, amount)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) JumpStartTotalBorrowed(amount *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.JumpStartTotalBorrowed(&_InfinityPoolV2.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Session) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Mint(&_InfinityPoolV2.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Mint(&_InfinityPoolV2.TransactOpts, shares, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) Pause() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Pause(&_InfinityPoolV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Pause() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Pause(&_InfinityPoolV2.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Pay(opts *bind.TransactOpts, vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "pay", vc)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPoolV2 *InfinityPoolV2Session) Pay(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Pay(&_InfinityPoolV2.TransactOpts, vc)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Pay(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Pay(&_InfinityPoolV2.TransactOpts, vc)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) RecoverERC20(opts *bind.TransactOpts, receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "recoverERC20", receiver, token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) RecoverERC20(receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RecoverERC20(&_InfinityPoolV2.TransactOpts, receiver, token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) RecoverERC20(receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RecoverERC20(&_InfinityPoolV2.TransactOpts, receiver, token)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) RecoverFIL(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "recoverFIL", receiver)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) RecoverFIL(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RecoverFIL(&_InfinityPoolV2.TransactOpts, receiver)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) RecoverFIL(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RecoverFIL(&_InfinityPoolV2.TransactOpts, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "redeem", shares, receiver, owner, arg3)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Session) Redeem(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Redeem(&_InfinityPoolV2.TransactOpts, shares, receiver, owner, arg3)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Redeem(&_InfinityPoolV2.TransactOpts, shares, receiver, owner, arg3)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "redeem0", shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Session) Redeem0(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Redeem0(&_InfinityPoolV2.TransactOpts, shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Redeem0(&_InfinityPoolV2.TransactOpts, shares, receiver, owner)
}

// RedeemF is a paid mutator transaction binding the contract method 0x0e0a7023.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) RedeemF(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "redeemF", shares, receiver, owner, arg3)
}

// RedeemF is a paid mutator transaction binding the contract method 0x0e0a7023.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Session) RedeemF(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RedeemF(&_InfinityPoolV2.TransactOpts, shares, receiver, owner, arg3)
}

// RedeemF is a paid mutator transaction binding the contract method 0x0e0a7023.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) RedeemF(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RedeemF(&_InfinityPoolV2.TransactOpts, shares, receiver, owner, arg3)
}

// RedeemF0 is a paid mutator transaction binding the contract method 0x9050c3d7.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) RedeemF0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "redeemF0", shares, receiver, owner)
}

// RedeemF0 is a paid mutator transaction binding the contract method 0x9050c3d7.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2Session) RedeemF0(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RedeemF0(&_InfinityPoolV2.TransactOpts, shares, receiver, owner)
}

// RedeemF0 is a paid mutator transaction binding the contract method 0x9050c3d7.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) RedeemF0(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RedeemF0(&_InfinityPoolV2.TransactOpts, shares, receiver, owner)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) RefreshRoutes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "refreshRoutes")
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) RefreshRoutes() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RefreshRoutes(&_InfinityPoolV2.TransactOpts)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) RefreshRoutes() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.RefreshRoutes(&_InfinityPoolV2.TransactOpts)
}

// SetLiquidityMine is a paid mutator transaction binding the contract method 0x43fd5dfe.
//
// Solidity: function setLiquidityMine(address _liquidityMine) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) SetLiquidityMine(opts *bind.TransactOpts, _liquidityMine common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "setLiquidityMine", _liquidityMine)
}

// SetLiquidityMine is a paid mutator transaction binding the contract method 0x43fd5dfe.
//
// Solidity: function setLiquidityMine(address _liquidityMine) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) SetLiquidityMine(_liquidityMine common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetLiquidityMine(&_InfinityPoolV2.TransactOpts, _liquidityMine)
}

// SetLiquidityMine is a paid mutator transaction binding the contract method 0x43fd5dfe.
//
// Solidity: function setLiquidityMine(address _liquidityMine) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) SetLiquidityMine(_liquidityMine common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetLiquidityMine(&_InfinityPoolV2.TransactOpts, _liquidityMine)
}

// SetMaxAccountsToUpdatePerBatch is a paid mutator transaction binding the contract method 0x549c5877.
//
// Solidity: function setMaxAccountsToUpdatePerBatch(uint256 maxAccountsToUpdatePerBatch_) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) SetMaxAccountsToUpdatePerBatch(opts *bind.TransactOpts, maxAccountsToUpdatePerBatch_ *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "setMaxAccountsToUpdatePerBatch", maxAccountsToUpdatePerBatch_)
}

// SetMaxAccountsToUpdatePerBatch is a paid mutator transaction binding the contract method 0x549c5877.
//
// Solidity: function setMaxAccountsToUpdatePerBatch(uint256 maxAccountsToUpdatePerBatch_) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) SetMaxAccountsToUpdatePerBatch(maxAccountsToUpdatePerBatch_ *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetMaxAccountsToUpdatePerBatch(&_InfinityPoolV2.TransactOpts, maxAccountsToUpdatePerBatch_)
}

// SetMaxAccountsToUpdatePerBatch is a paid mutator transaction binding the contract method 0x549c5877.
//
// Solidity: function setMaxAccountsToUpdatePerBatch(uint256 maxAccountsToUpdatePerBatch_) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) SetMaxAccountsToUpdatePerBatch(maxAccountsToUpdatePerBatch_ *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetMaxAccountsToUpdatePerBatch(&_InfinityPoolV2.TransactOpts, maxAccountsToUpdatePerBatch_)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) SetMinimumLiquidity(opts *bind.TransactOpts, _minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "setMinimumLiquidity", _minimumLiquidity)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) SetMinimumLiquidity(_minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetMinimumLiquidity(&_InfinityPoolV2.TransactOpts, _minimumLiquidity)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) SetMinimumLiquidity(_minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetMinimumLiquidity(&_InfinityPoolV2.TransactOpts, _minimumLiquidity)
}

// SetTreasuryFeeRate is a paid mutator transaction binding the contract method 0xce43303c.
//
// Solidity: function setTreasuryFeeRate(uint256 _treasuryFeeRate) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) SetTreasuryFeeRate(opts *bind.TransactOpts, _treasuryFeeRate *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "setTreasuryFeeRate", _treasuryFeeRate)
}

// SetTreasuryFeeRate is a paid mutator transaction binding the contract method 0xce43303c.
//
// Solidity: function setTreasuryFeeRate(uint256 _treasuryFeeRate) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) SetTreasuryFeeRate(_treasuryFeeRate *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetTreasuryFeeRate(&_InfinityPoolV2.TransactOpts, _treasuryFeeRate)
}

// SetTreasuryFeeRate is a paid mutator transaction binding the contract method 0xce43303c.
//
// Solidity: function setTreasuryFeeRate(uint256 _treasuryFeeRate) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) SetTreasuryFeeRate(_treasuryFeeRate *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.SetTreasuryFeeRate(&_InfinityPoolV2.TransactOpts, _treasuryFeeRate)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) ShutDown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "shutDown")
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) ShutDown() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.ShutDown(&_InfinityPoolV2.TransactOpts)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) ShutDown() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.ShutDown(&_InfinityPoolV2.TransactOpts)
}

// StartRateUpdate is a paid mutator transaction binding the contract method 0x9a33bc7c.
//
// Solidity: function startRateUpdate(uint256 rentalFeesOwedPerEpoch_) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) StartRateUpdate(opts *bind.TransactOpts, rentalFeesOwedPerEpoch_ *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "startRateUpdate", rentalFeesOwedPerEpoch_)
}

// StartRateUpdate is a paid mutator transaction binding the contract method 0x9a33bc7c.
//
// Solidity: function startRateUpdate(uint256 rentalFeesOwedPerEpoch_) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) StartRateUpdate(rentalFeesOwedPerEpoch_ *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.StartRateUpdate(&_InfinityPoolV2.TransactOpts, rentalFeesOwedPerEpoch_)
}

// StartRateUpdate is a paid mutator transaction binding the contract method 0x9a33bc7c.
//
// Solidity: function startRateUpdate(uint256 rentalFeesOwedPerEpoch_) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) StartRateUpdate(rentalFeesOwedPerEpoch_ *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.StartRateUpdate(&_InfinityPoolV2.TransactOpts, rentalFeesOwedPerEpoch_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.TransferOwnership(&_InfinityPoolV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.TransferOwnership(&_InfinityPoolV2.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) Unpause() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Unpause(&_InfinityPoolV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Unpause(&_InfinityPoolV2.TransactOpts)
}

// UpdateAccounting is a paid mutator transaction binding the contract method 0x494347e7.
//
// Solidity: function updateAccounting() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) UpdateAccounting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "updateAccounting")
}

// UpdateAccounting is a paid mutator transaction binding the contract method 0x494347e7.
//
// Solidity: function updateAccounting() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) UpdateAccounting() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.UpdateAccounting(&_InfinityPoolV2.TransactOpts)
}

// UpdateAccounting is a paid mutator transaction binding the contract method 0x494347e7.
//
// Solidity: function updateAccounting() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) UpdateAccounting() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.UpdateAccounting(&_InfinityPoolV2.TransactOpts)
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) UpdateCredParser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "updateCredParser")
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) UpdateCredParser() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.UpdateCredParser(&_InfinityPoolV2.TransactOpts)
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) UpdateCredParser() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.UpdateCredParser(&_InfinityPoolV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "withdraw", assets, receiver, owner, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Withdraw(&_InfinityPoolV2.TransactOpts, assets, receiver, owner, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Withdraw(&_InfinityPoolV2.TransactOpts, assets, receiver, owner, arg3)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "withdraw0", assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Withdraw0(&_InfinityPoolV2.TransactOpts, assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Withdraw0(&_InfinityPoolV2.TransactOpts, assets, receiver, owner)
}

// WithdrawF is a paid mutator transaction binding the contract method 0x0c462c5f.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) WithdrawF(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "withdrawF", assets, receiver, owner, arg3)
}

// WithdrawF is a paid mutator transaction binding the contract method 0x0c462c5f.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) WithdrawF(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.WithdrawF(&_InfinityPoolV2.TransactOpts, assets, receiver, owner, arg3)
}

// WithdrawF is a paid mutator transaction binding the contract method 0x0c462c5f.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) WithdrawF(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.WithdrawF(&_InfinityPoolV2.TransactOpts, assets, receiver, owner, arg3)
}

// WithdrawF0 is a paid mutator transaction binding the contract method 0x870f77cf.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Transactor) WithdrawF0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "withdrawF0", assets, receiver, owner)
}

// WithdrawF0 is a paid mutator transaction binding the contract method 0x870f77cf.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Session) WithdrawF0(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.WithdrawF0(&_InfinityPoolV2.TransactOpts, assets, receiver, owner)
}

// WithdrawF0 is a paid mutator transaction binding the contract method 0x870f77cf.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) WithdrawF0(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.WithdrawF0(&_InfinityPoolV2.TransactOpts, assets, receiver, owner)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) WriteOff(opts *bind.TransactOpts, agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.Transact(opts, "writeOff", agentID, recoveredFunds)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) WriteOff(agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.WriteOff(&_InfinityPoolV2.TransactOpts, agentID, recoveredFunds)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) WriteOff(agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.WriteOff(&_InfinityPoolV2.TransactOpts, agentID, recoveredFunds)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Fallback(&_InfinityPoolV2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Fallback(&_InfinityPoolV2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPoolV2 *InfinityPoolV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPoolV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPoolV2 *InfinityPoolV2Session) Receive() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Receive(&_InfinityPoolV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPoolV2 *InfinityPoolV2TransactorSession) Receive() (*types.Transaction, error) {
	return _InfinityPoolV2.Contract.Receive(&_InfinityPoolV2.TransactOpts)
}

// InfinityPoolV2BorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the InfinityPoolV2 contract.
type InfinityPoolV2BorrowIterator struct {
	Event *InfinityPoolV2Borrow // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2BorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2Borrow)
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
		it.Event = new(InfinityPoolV2Borrow)
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
func (it *InfinityPoolV2BorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2BorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2Borrow represents a Borrow event raised by the InfinityPoolV2 contract.
type InfinityPoolV2Borrow struct {
	Agent  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterBorrow(opts *bind.FilterOpts, agent []*big.Int) (*InfinityPoolV2BorrowIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "Borrow", agentRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2BorrowIterator{contract: _InfinityPoolV2.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2Borrow, agent []*big.Int) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "Borrow", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2Borrow)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParseBorrow(log types.Log) (*InfinityPoolV2Borrow, error) {
	event := new(InfinityPoolV2Borrow)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolV2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the InfinityPoolV2 contract.
type InfinityPoolV2DepositIterator struct {
	Event *InfinityPoolV2Deposit // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2Deposit)
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
		it.Event = new(InfinityPoolV2Deposit)
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
func (it *InfinityPoolV2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2Deposit represents a Deposit event raised by the InfinityPoolV2 contract.
type InfinityPoolV2Deposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*InfinityPoolV2DepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2DepositIterator{contract: _InfinityPoolV2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2Deposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2Deposit)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParseDeposit(log types.Log) (*InfinityPoolV2Deposit, error) {
	event := new(InfinityPoolV2Deposit)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the InfinityPoolV2 contract.
type InfinityPoolV2PausedIterator struct {
	Event *InfinityPoolV2Paused // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2Paused)
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
		it.Event = new(InfinityPoolV2Paused)
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
func (it *InfinityPoolV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2Paused represents a Paused event raised by the InfinityPoolV2 contract.
type InfinityPoolV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterPaused(opts *bind.FilterOpts) (*InfinityPoolV2PausedIterator, error) {

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2PausedIterator{contract: _InfinityPoolV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2Paused) (event.Subscription, error) {

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2Paused)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParsePaused(log types.Log) (*InfinityPoolV2Paused, error) {
	event := new(InfinityPoolV2Paused)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolV2PayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the InfinityPoolV2 contract.
type InfinityPoolV2PayIterator struct {
	Event *InfinityPoolV2Pay // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2PayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2Pay)
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
		it.Event = new(InfinityPoolV2Pay)
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
func (it *InfinityPoolV2PayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2PayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2Pay represents a Pay event raised by the InfinityPoolV2 contract.
type InfinityPoolV2Pay struct {
	Agent     *big.Int
	Amount    *big.Int
	Interest  *big.Int
	Principal *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x0dfb9710cc5cf2de9c8b20b6cc77867afbe42c7ff1408df657d97a904ccfdc9c.
//
// Solidity: event Pay(uint256 indexed agent, uint256 amount, uint256 interest, uint256 principal)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterPay(opts *bind.FilterOpts, agent []*big.Int) (*InfinityPoolV2PayIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "Pay", agentRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2PayIterator{contract: _InfinityPoolV2.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x0dfb9710cc5cf2de9c8b20b6cc77867afbe42c7ff1408df657d97a904ccfdc9c.
//
// Solidity: event Pay(uint256 indexed agent, uint256 amount, uint256 interest, uint256 principal)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchPay(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2Pay, agent []*big.Int) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "Pay", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2Pay)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0x0dfb9710cc5cf2de9c8b20b6cc77867afbe42c7ff1408df657d97a904ccfdc9c.
//
// Solidity: event Pay(uint256 indexed agent, uint256 amount, uint256 interest, uint256 principal)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParsePay(log types.Log) (*InfinityPoolV2Pay, error) {
	event := new(InfinityPoolV2Pay)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the InfinityPoolV2 contract.
type InfinityPoolV2UnpausedIterator struct {
	Event *InfinityPoolV2Unpaused // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2Unpaused)
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
		it.Event = new(InfinityPoolV2Unpaused)
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
func (it *InfinityPoolV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2Unpaused represents a Unpaused event raised by the InfinityPoolV2 contract.
type InfinityPoolV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*InfinityPoolV2UnpausedIterator, error) {

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2UnpausedIterator{contract: _InfinityPoolV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2Unpaused)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParseUnpaused(log types.Log) (*InfinityPoolV2Unpaused, error) {
	event := new(InfinityPoolV2Unpaused)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolV2UpdateAccountingIterator is returned from FilterUpdateAccounting and is used to iterate over the raw logs and unpacked data for UpdateAccounting events raised by the InfinityPoolV2 contract.
type InfinityPoolV2UpdateAccountingIterator struct {
	Event *InfinityPoolV2UpdateAccounting // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2UpdateAccountingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2UpdateAccounting)
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
		it.Event = new(InfinityPoolV2UpdateAccounting)
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
func (it *InfinityPoolV2UpdateAccountingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2UpdateAccountingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2UpdateAccounting represents a UpdateAccounting event raised by the InfinityPoolV2 contract.
type InfinityPoolV2UpdateAccounting struct {
	NewEarnings   *big.Int
	TotalAssets   *big.Int
	TotalBorrowed *big.Int
	AssetBalance  *big.Int
	IFILPrice     *big.Int
	IFILSupply    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateAccounting is a free log retrieval operation binding the contract event 0xa44f29ae5036974cc3cc64b097368e60ba4210a0fef291753e670437b9587bb2.
//
// Solidity: event UpdateAccounting(uint256 newEarnings, uint256 totalAssets, uint256 totalBorrowed, uint256 assetBalance, uint256 iFILPrice, uint256 iFILSupply)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterUpdateAccounting(opts *bind.FilterOpts) (*InfinityPoolV2UpdateAccountingIterator, error) {

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "UpdateAccounting")
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2UpdateAccountingIterator{contract: _InfinityPoolV2.contract, event: "UpdateAccounting", logs: logs, sub: sub}, nil
}

// WatchUpdateAccounting is a free log subscription operation binding the contract event 0xa44f29ae5036974cc3cc64b097368e60ba4210a0fef291753e670437b9587bb2.
//
// Solidity: event UpdateAccounting(uint256 newEarnings, uint256 totalAssets, uint256 totalBorrowed, uint256 assetBalance, uint256 iFILPrice, uint256 iFILSupply)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchUpdateAccounting(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2UpdateAccounting) (event.Subscription, error) {

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "UpdateAccounting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2UpdateAccounting)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "UpdateAccounting", log); err != nil {
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

// ParseUpdateAccounting is a log parse operation binding the contract event 0xa44f29ae5036974cc3cc64b097368e60ba4210a0fef291753e670437b9587bb2.
//
// Solidity: event UpdateAccounting(uint256 newEarnings, uint256 totalAssets, uint256 totalBorrowed, uint256 assetBalance, uint256 iFILPrice, uint256 iFILSupply)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParseUpdateAccounting(log types.Log) (*InfinityPoolV2UpdateAccounting, error) {
	event := new(InfinityPoolV2UpdateAccounting)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "UpdateAccounting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolV2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the InfinityPoolV2 contract.
type InfinityPoolV2WithdrawIterator struct {
	Event *InfinityPoolV2Withdraw // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2Withdraw)
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
		it.Event = new(InfinityPoolV2Withdraw)
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
func (it *InfinityPoolV2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2Withdraw represents a Withdraw event raised by the InfinityPoolV2 contract.
type InfinityPoolV2Withdraw struct {
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
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*InfinityPoolV2WithdrawIterator, error) {

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

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2WithdrawIterator{contract: _InfinityPoolV2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2Withdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2Withdraw)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParseWithdraw(log types.Log) (*InfinityPoolV2Withdraw, error) {
	event := new(InfinityPoolV2Withdraw)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolV2WriteOffIterator is returned from FilterWriteOff and is used to iterate over the raw logs and unpacked data for WriteOff events raised by the InfinityPoolV2 contract.
type InfinityPoolV2WriteOffIterator struct {
	Event *InfinityPoolV2WriteOff // Event containing the contract specifics and raw log

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
func (it *InfinityPoolV2WriteOffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolV2WriteOff)
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
		it.Event = new(InfinityPoolV2WriteOff)
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
func (it *InfinityPoolV2WriteOffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolV2WriteOffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolV2WriteOff represents a WriteOff event raised by the InfinityPoolV2 contract.
type InfinityPoolV2WriteOff struct {
	AgentID        *big.Int
	RecoveredFunds *big.Int
	LostFunds      *big.Int
	LostTFees      *big.Int
	InterestPaid   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWriteOff is a free log retrieval operation binding the contract event 0xf02514bceabf2b74024e8529fdc9f423bebda69624841ff163ebd712c188f4a7.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 lostTFees, uint256 interestPaid)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) FilterWriteOff(opts *bind.FilterOpts, agentID []*big.Int) (*InfinityPoolV2WriteOffIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.FilterLogs(opts, "WriteOff", agentIDRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolV2WriteOffIterator{contract: _InfinityPoolV2.contract, event: "WriteOff", logs: logs, sub: sub}, nil
}

// WatchWriteOff is a free log subscription operation binding the contract event 0xf02514bceabf2b74024e8529fdc9f423bebda69624841ff163ebd712c188f4a7.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 lostTFees, uint256 interestPaid)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) WatchWriteOff(opts *bind.WatchOpts, sink chan<- *InfinityPoolV2WriteOff, agentID []*big.Int) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _InfinityPoolV2.contract.WatchLogs(opts, "WriteOff", agentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolV2WriteOff)
				if err := _InfinityPoolV2.contract.UnpackLog(event, "WriteOff", log); err != nil {
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

// ParseWriteOff is a log parse operation binding the contract event 0xf02514bceabf2b74024e8529fdc9f423bebda69624841ff163ebd712c188f4a7.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 lostTFees, uint256 interestPaid)
func (_InfinityPoolV2 *InfinityPoolV2Filterer) ParseWriteOff(log types.Log) (*InfinityPoolV2WriteOff, error) {
	event := new(InfinityPoolV2WriteOff)
	if err := _InfinityPoolV2.contract.UnpackLog(event, "WriteOff", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
