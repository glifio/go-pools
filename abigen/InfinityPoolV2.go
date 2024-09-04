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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidStakingToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumLiquidity_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountDNE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDefaulted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolShuttingDown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsPaidCursor\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsPaidCursor\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEarnings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iFILPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iFILSupply\",\"type\":\"uint256\"}],\"name\":\"UpdateAccounting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lostFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lostTFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"}],\"name\":\"WriteOff\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"continueRateUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credParser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"decommissionPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAbsMinLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentInterestOwed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"harvestAmount\",\"type\":\"uint256\"}],\"name\":\"harvestFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShuttingDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountPrincipal\",\"type\":\"uint256\"}],\"name\":\"jumpStartAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"jumpStartTotalBorrowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccountingUpdateEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingToken\",\"outputs\":[{\"internalType\":\"contractIPoolToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lm\",\"outputs\":[{\"internalType\":\"contractILiquidityMineSP\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accrued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structRewardAccrual\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateUpdate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalAccountsAtUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAccountsClosed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"inProcess\",\"type\":\"bool\"}],\"internalType\":\"structRateUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"recoverFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeemF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeemF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityMine\",\"type\":\"address\"}],\"name\":\"setLiquidityMine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAccountsToUpdatePerBatch_\",\"type\":\"uint256\"}],\"name\":\"setMaxAccountsToUpdatePerBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumLiquidity\",\"type\":\"uint256\"}],\"name\":\"setMinimumLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treasuryFeeRate\",\"type\":\"uint256\"}],\"name\":\"setTreasuryFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rentalFeesOwedPerEpoch_\",\"type\":\"uint256\"}],\"name\":\"startRateUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowableAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFeesOwed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accrued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"internalType\":\"structRewardAccrual\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAccounting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCredParser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"}],\"name\":\"writeOff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "",
}

// InfinityPoolV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use InfinityPoolV2MetaData.ABI instead.
var InfinityPoolV2ABI = InfinityPoolV2MetaData.ABI

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
	Agent            *big.Int
	Amount           *big.Int
	EpochsPaidCursor *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xd67bffb11f634e8d9df020d50462c82383a8dab0e8025486778c08f303353432.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount, uint256 epochsPaidCursor)
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

// WatchBorrow is a free log subscription operation binding the contract event 0xd67bffb11f634e8d9df020d50462c82383a8dab0e8025486778c08f303353432.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount, uint256 epochsPaidCursor)
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

// ParseBorrow is a log parse operation binding the contract event 0xd67bffb11f634e8d9df020d50462c82383a8dab0e8025486778c08f303353432.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount, uint256 epochsPaidCursor)
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
	Agent            *big.Int
	Amount           *big.Int
	Interest         *big.Int
	Principal        *big.Int
	EpochsPaidCursor *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 amount, uint256 interest, uint256 principal, uint256 epochsPaidCursor)
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

// WatchPay is a free log subscription operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 amount, uint256 interest, uint256 principal, uint256 epochsPaidCursor)
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

// ParsePay is a log parse operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 amount, uint256 interest, uint256 principal, uint256 epochsPaidCursor)
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
