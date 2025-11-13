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

// LpPlusMerkleHelperUserSnapshot is an auto generated low-level Go binding around an user-defined struct.
type LpPlusMerkleHelperUserSnapshot struct {
	TokenId      *big.Int
	StrikePrice  *big.Int
	AllocatedFil *big.Int
}

// RWTFutureBonus is an auto generated low-level Go binding around an user-defined struct.
type RWTFutureBonus struct {
	AllocatedFilExtra   *big.Int
	ExpirationDateExtra *big.Int
}

// RwtInterpolation is an auto generated low-level Go binding around an user-defined struct.
type RwtInterpolation struct {
	LowRatio       *big.Int
	LowMultiplier  *big.Int
	HighRatio      *big.Int
	HighMultiplier *big.Int
}

// StakingSnapshot is an auto generated low-level Go binding around an user-defined struct.
type StakingSnapshot struct {
	BaseConversionRateFilPerRwt *big.Int
	InterpolationParams         RwtInterpolation
	FutureValidityDuration      *big.Int
	Timestamp                   *big.Int
	StakingYbtTwab              *big.Int
	TotalFilToAllocate          *big.Int
	ClosingTotalYbt             *big.Int
}

// UserAction is an auto generated low-level Go binding around an user-defined struct.
type UserAction struct {
	TokenId    *big.Int
	YbtBalance *big.Int
	RwtBalance *big.Int
	WindowId   *big.Int
	Timestamp  *big.Int
}

// LPPlusMetaData contains all meta data concerning the LPPlus contract.
var LPPlusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"}],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"FilReservesDepleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFILBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientRWTBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientVaultBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_second\",\"type\":\"uint256\"}],\"name\":\"InvalidLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"name\":\"RWTFutureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenNotActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroConversionRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"}],\"name\":\"ZeroMerkleRoot\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structRWTFutureBonus\",\"name\":\"bonus\",\"type\":\"tuple\"}],\"name\":\"BonusApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referrerTokenId\",\"type\":\"uint256\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FilVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"RWTDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rwtFutureTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rwtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filAmount\",\"type\":\"uint256\"}],\"name\":\"RWTExchangedForFIL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"RWTWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"YBTVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"YBTVaultWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_RWTAmount\",\"type\":\"uint256\"}],\"name\":\"addRWT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_YBTAmount\",\"type\":\"uint256\"}],\"name\":\"addYBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_windowIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"}],\"internalType\":\"structLpPlusMerkleHelper.UserSnapshot[]\",\"name\":\"_userSnapshots\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_proofs\",\"type\":\"bytes32[][]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_windowIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"}],\"internalType\":\"structLpPlusMerkleHelper.UserSnapshot[]\",\"name\":\"_userSnapshots\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_proofs\",\"type\":\"bytes32[][]\"}],\"name\":\"claimAndExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTotalYbt\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"_stakingSnapshot\",\"type\":\"tuple\"}],\"name\":\"commitWindow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_rwtFutureTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundFilVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureValidityDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getUserActionsBatched\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ybtBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rwtBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structUserAction[][]\",\"name\":\"actionsPerToken\",\"type\":\"tuple[][]\"},{\"internalType\":\"uint256\",\"name\":\"newStartTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newOffset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rwtToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_ybtToken\",\"type\":\"address\"},{\"internalType\":\"contractIRWTFuture\",\"name\":\"_rwtFuture\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_windowCommitter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interpolationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumRWTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumYBTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_referrerTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ybtToStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rwtToStake\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBonus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingTotalYbt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalFilToAllocate\",\"type\":\"uint256\"}],\"name\":\"previewStakingSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTotalYbt\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referrerBonus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_RWTAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeRWT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_YBTAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeYBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwtFuture\",\"outputs\":[{\"internalType\":\"contractIRWTFuture\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_futureValidityDuration\",\"type\":\"uint256\"}],\"name\":\"setFutureValidityDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"_interpolationParams\",\"type\":\"tuple\"}],\"name\":\"setInterpolationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumRWTBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumRWTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumYBTBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumYBTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"_mintBonus\",\"type\":\"tuple\"}],\"name\":\"setMintBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRWTFuture\",\"name\":\"_RWTFuture\",\"type\":\"address\"}],\"name\":\"setRWTFuture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_RWTToken\",\"type\":\"address\"}],\"name\":\"setRWTToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"_referrerBonus\",\"type\":\"tuple\"}],\"name\":\"setReferrerBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_windowCommitter\",\"type\":\"address\"}],\"name\":\"setWindowCommitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_YBTToken\",\"type\":\"address\"}],\"name\":\"setYBTToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToRWTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToWindowIdToClaimedFlag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToYBTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRwtStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalYbtStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"windowCommitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"windowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"windowIdToMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"windowIdToStakingSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTotalYbt\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFilFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ybtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f516020617ba95f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051617ae290816100c78239608051818181614a450152614dfa0152f35b6001600160401b0319166001600160401b039081175f516020617ba95f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60a0806040526004361015610012575f80fd5b5f6080525f3560e01c90816301ffc9a71461621c5750806306fdde031461611b578063081812fc14616094578063095ea7b3146160345780630a6b569914615f3c5780631418f71114615eda57806316290e1d14615c2457806323b872dd14615beb57806329f2014d14615b845780632bf2ba0314615b225780632e58c0c714615ac85780632eacc69614615a19578063300e3c4f1461593f5780633263edee146158e5578063344bb6b11461581e578063385b5f85146155aa5780633a20853e1461552e5780633f4ba83a146153ac57806342842e0e1461536357806346b3241e146150d25780634bae2ef1146150905780634f1ef28614614d91578063510402cd14614d325780635194736114614abd57806352d1902d14614a00578063548d93561461492757806354fd4d50146148ee57806358f28a98146148945780635c975abb146148355780635e29e06b1461472357806361d027b3146146b35780636352211e146146595780636817c76c146145ff578063694072a61461429e57806370a08231146141c457806370af96aa14613d2d578063715018a614613be057806379ba509714613b39578063812bf90a14613ac75780638456cb59146139655780638da5cb5b146138f35780638edd7d4e146138985780638f68679f1461383c5780638ffd6d2d146137da578063912934731461376f57806395d89b41146136295780639af1d03e146135cd578063a22cb4651461355e578063a647e8ec1461269f578063a818b1ee14612643578063ad3cb1cc146125c3578063ae7e2d9714612558578063b05dcd5a146124ed578063b5ea015c14612491578063b88d4fde14612403578063bfb1909b1461210d578063c74c23d5146120ab578063c87b56dd14612034578063d262f44414611f5d578063d452154814611165578063d5d2ae72146110f3578063da357c0f146109e3578063e080e4bd14610981578063e0ea4a1b14610911578063e30c39781461089f578063e985e9c5146107ef578063ee2df9fd1461077d578063f0f442601461066a578063f2fde38b1461053f578063f4a0a528146104dd578063f721dc761461046b5763fcd0907d14610338575f80fd5b346104655760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465576040516103738161644f565b6004358152602081016024358152604082019060443582526060830192606435845261039d617040565b82518151111561043957517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1055517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1155517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1255517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d135560805180f35b7f1f3b85d300000000000000000000000000000000000000000000000000000000608051526004608051fd5b60805180fd5b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416604051908152f35b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557610514617040565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d005560805180f35b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655773ffffffffffffffffffffffffffffffffffffffff61059361058e61634c565b6170ac565b61059b617040565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700608051608051a360805180f35b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465576106ac6106a461634c565b61058e617040565b73ffffffffffffffffffffffffffffffffffffffff8116156107515761074b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0555565b60805180f35b7fd92e233d00000000000000000000000000000000000000000000000000000000608051526004608051fd5b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416604051908152f35b346104655760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655761082661634c565b73ffffffffffffffffffffffffffffffffffffffff61084361636f565b91165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557610949616c24565b506080610954616c85565b61097f6040518092606080918051845260208101516020850152604081015160408501520151910152565bf35b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465576109b8617040565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085560805180f35b3461046557610a1a73ffffffffffffffffffffffffffffffffffffffff610a0936616729565b959293969161058e989195986171c8565b16928315610751578615610cf9578215610cf9578415610cf957610a418784818114617662565b610a4e8786818114617662565b610a5787616f3b565b94610a6188616f3b565b97610a6b81616f3b565b975f5b828110610d21578a8a8a8a81515f905f905b808210610c7d575050478111904791610c4f57855f610b1e87878388610b7e73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d03541694610b4e604051998a98899788967f1d05159d0000000000000000000000000000000000000000000000000000000088526004880152608060248801526084870190616802565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc868303016044870152616802565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016064850152616802565b03925af18015610c44575f90610bad575b80610ba99150604051918291602083526020830190616802565b0390f35b503d805f833e610bbd8183616487565b810190602081830312610c405780519067ffffffffffffffff8211610c4057019080601f83011215610c40578151610bf4816168c0565b92610c026040519485616487565b81845260208085019260051b820101928311610c4057602001905b828210610c3057505050610ba990610b8f565b8151815260209182019101610c1d565b5f80fd5b6040513d5f823e3d90fd5b7f7b0e080a000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b9091610c898386616938565b5190610c958489616938565b5191610ca18589616938565b519015610cf9578215610cf957804211610cca5750600191610cc291616905565b920190610a80565b7f431e9b4d000000000000000000000000000000000000000000000000000000005f525f60045260245260445ffd5b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b8681101561109657606081028201610d3a828589616f8a565b35610d673383358173ffffffffffffffffffffffffffffffffffffffff610d6083616fd7565b1614616de5565b602082013515610cf957604082013515610cf95781355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1760205260405f20815f5260205260ff60405f2054166110c357805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a60205260405f2054610df182821515616bf1565b606083360312610c4057604051610e0781616417565b833581526020810160208501358152604082016040860135815260405191602083019386855251604084015251606083015251608082015260808152610e4e60a082616487565b5190206040516020810191825260208152610e6a604082616487565b5190209087851015611096578460051b8901357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18a360301811215610c4057890180359067ffffffffffffffff8211610c4057602001918160051b36038313610c4057835f5b83811061105c57500361100857505050805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d19602052610f6a60405f2060405190610f1b826163ce565b80548252610f2b60018201616d2d565b602083015260058101549182604082015260c0600960068401549384606085015260078101546080850152600881015460a08501520154910152616905565b804211610cca578360408f93948f8f909661100195610f8c8660019b9a616938565b5282355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d17602052835f20905f52602052825f20887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055610ff884602084013592616938565b52013592616938565b5201610a6e565b90611052916040519586957f10f7f5f90000000000000000000000000000000000000000000000000000000087526004870152356024860152608060448601526084850191616f9a565b9060648301520390fd5b90611068828587616f8a565b359081811015611085575f52602052600160405f205b9101610ed0565b905f52602052600160405f2061107e565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b907f6bd4745f000000000000000000000000000000000000000000000000000000005f523560045260245260445ffd5b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416604051908152f35b34610465576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655760043567ffffffffffffffff8111610465576111b59036906004016166b4565b60243567ffffffffffffffff8111610465576111d59036906004016166b4565b60443573ffffffffffffffffffffffffffffffffffffffff81168103610c40576064359060843573ffffffffffffffffffffffffffffffffffffffff8116809103610c405760a4359373ffffffffffffffffffffffffffffffffffffffff8516809503610c405760c4359073ffffffffffffffffffffffffffffffffffffffff8216809203610c405760e4359073ffffffffffffffffffffffffffffffffffffffff8216808303610c4057610104359573ffffffffffffffffffffffffffffffffffffffff871692838803610c40577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549a60ff8c60401c16159b8c67ffffffffffffffff821680159182611f55575b506001149081611f4b575b159081611f42575b50611f16578c60017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055611ec1575b506113546179bc565b61135c6179bc565b80519067ffffffffffffffff8211611cb657819061139a7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054616835565b601f8111611df3575b506020906001601f841114611cf25760805192611ce7575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b80519067ffffffffffffffff8211611cb6576114497f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154616835565b601f8111611c4b575b506020906001601f841114611b61579180916114cc949360805192611b56575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b61058e6179bc565b6114d46179bc565b6114dc6179bc565b73ffffffffffffffffffffffffffffffffffffffff811615611b245761150190617699565b6115096179bc565b6115116179bc565b8615611af8578415610751578715610751578315610751571561075157156107515761058e73ffffffffffffffffffffffffffffffffffffffff6116946117149461058e836116146118129c61058e836115946117939f9d61058e9e7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d00556170ac565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0155565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0255565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0355565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0555565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0655565b60017f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d045562ed4e007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d075566038d7ea4c680007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085566038d7ea4c680007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095562ed4e0060206040516118c281616433565b662386f26fc1000081520152662386f26fc100007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5562ed4e007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d55662386f26fc100007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5562ed4e007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f5567016345785d8a000060606040516119858161644f565b670de0b6b3a76400008152666a94d74f43000060208201526706f05b59d3b2000060408201520152670de0b6b3a76400007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1055666a94d74f4300007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d11556706f05b59d3b200007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d125567016345785d8a00007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1355611a625760805180f35b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a161074b565b7f1f2a200500000000000000000000000000000000000000000000000000000000608051526004608051fd5b7f1e4fbdf700000000000000000000000000000000000000000000000000000000608051526080516004526024608051fd5b015190508d80611472565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301608051528160805120926080515b818110611c3357509160019391856114cc97969410611bfc575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301556114c4565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080611bcf565b92936020600181928786015181550195019301611bb5565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930160805152602060805120601f840160051c81019160208510611cac575b601f0160051c01905b818110611c9f5750611452565b6080518155600101611c92565b9091508190611c89565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b015190508d806113bb565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300608051527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81926080515b818110611ddb5750908460019594939210611da4575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005561140d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080611d77565b92936020600181928786015181550195019301611d61565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930060805152601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410611e99575b90601f8493920160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b818110611e8957506113a3565b6080518155849350600101611e7c565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf819150611e4e565b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00558c61134b565b7ff92ee8a900000000000000000000000000000000000000000000000000000000608051526004608051fd5b9050158e6112f8565b303b1591506112f0565b91508e6112e5565b3461046557611f6b36616392565b611f76939293616c48565b50821561200857610ba9937f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0754927f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a549360405195611fd4876163ce565b8652611fde616c85565b602087015260408601526060850152608084015260a083015260c082015260405191829182616580565b7f247af9ce00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655761206e600435616fd7565b5060405161207d602082616487565b6080519052604051610ba990612094602082616487565b5f8152604051918291602083526020830190616309565b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465576120e2617040565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d075560805180f35b346104655761211b36616625565b916121246171c8565b61215161213082616fd7565b82339173ffffffffffffffffffffffffffffffffffffffff33911614616de5565b8115611af857612160836170ac565b9273ffffffffffffffffffffffffffffffffffffffff841615610751576123259082608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d166020526020846121cc81604060805120546121c781838a82821115616e34565b616886565b966122037f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085489888282108015906123fb57616d92565b85608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16835260406080512061223c838254616886565b9055612269827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54616886565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a556122948661721b565b73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416906040518096819482937fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0391608051905af19081156123ee577fc4032b9f060fe681645ab7a2b55a30a15304792f8669533901de6fb6e0dea1129273ffffffffffffffffffffffffffffffffffffffff926123c1575b506123b87f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5460405193849316973397846040919493926060820195825260208201520152565b0390a460805180f35b6123e29060203d6020116123e7575b6123da8183616487565b810190616dcd565b612371565b503d6123d0565b6040513d608051823e3d90fd5b508115616d92565b346104655760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655761243a61634c565b5061244361636f565b5060643567ffffffffffffffff8111610465576124649036906004016166b4565b507f8cd22d1900000000000000000000000000000000000000000000000000000000608051526004608051fd5b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0754604051908152f35b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557600435608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a602052602060406080512054604051908152f35b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557600435608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16602052602060406080512054604051908152f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655760408051610ba9916126059082616487565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190616309565b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54604051908152f35b346104655760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465576126d661634c565b6064356044356024356126e76171c8565b6126f0846170ac565b927f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d04549261271d846168d8565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d04557f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d01547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d05547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d00546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff92831660248201526044810191909152608051909260209284926064928492165af180156123ee57613541575b50602094604051966128258789616487565b608051885273ffffffffffffffffffffffffffffffffffffffff821697881561350f57612890875f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff8216918215159081613448575b506128fb8573ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b6001815401905588608051527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793028a526040608051208b7fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055888b604051947fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef608051608051a461341657833b613283575b5050909192939495965073ffffffffffffffffffffffffffffffffffffffff604051918683521690867f98d1d1eb3af2e9f5c522e0a03cec7374e965eb2a56911d0010ec933ede85affd893393a47f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1454604094855190612a118783616487565b600182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe087019283368b850137806131f657506080515b612a528361692b565b52612a5b616e87565b90801515806131e4575b806131da575b806131ce575b61305d575b5050612a80616ee1565b90612a8a8161692b565b51151580613053575b80613047575b612fac575b5050505080158015612d39575b505080158015612abe575b505051908152f35b612ac66171c8565b611af857612b338373ffffffffffffffffffffffffffffffffffffffff612b2b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b161515616d5f565b82608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d158452612b6a81836080512054616905565b612b9a7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954828682821015616d92565b83608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15855280836080512055612bf4827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54616905565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b55612c1f8461721b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015483517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015230602482015260448101849052608051909187918391606491839173ffffffffffffffffffffffffffffffffffffffff165af18015612d2d5785927f8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a984929091612d12575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b548551948552602085019190915260408401523392606090a38380612ab6565b612d2890883d8a116123e7576123da8183616487565b612cd0565b84513d608051823e3d90fd5b612d416171c8565b611af857612da68473ffffffffffffffffffffffffffffffffffffffff612b2b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b83608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d168552612ddd81846080512054616905565b612e0d7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854828782821015616d92565b84608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16865280846080512055612e67827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54616905565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a55612e928561721b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025484517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015230602482015260448101849052608051909188918391606491839173ffffffffffffffffffffffffffffffffffffffff165af18015612fa05786927fade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb4793790929091612f85575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a548651948552602085019190915260408401523392606090a38480612aab565b612f9b90893d8b116123e7576123da8183616487565b612f43565b85513d608051823e3d90fd5b8651825181526020808401519082015261303d949089907f08c20988b3a0d46fafe47abf0aa2ac3716335a2190a6acb020f12468259207ac90604090a261302e8a895194612ffa8b87616487565b600186528636838801378a51966130118c89616487565b60018852368389013780516130258761692b565b52015142616905565b6130378561692b565b526177a8565b5085808080612a9e565b50888201511515612a99565b5081511515612a93565b80608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d168a528760805120547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854111580613176575b156131475761313f9183826130ca8b94616fd7565b845184518152602080860151908201529091907f08c20988b3a0d46fafe47abf0aa2ac3716335a2190a6acb020f12468259207ac90604090a261302e8d8551946131148787616487565b6001865289368388013761312a87519788616487565b6001875289368389013780516130258761692b565b508880612a76565b7fed5a980a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b5080608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d158a528760805120547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095411156130b5565b50898201511515612a71565b5081511515612a6b565b506131ee8361692b565b511515612a65565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161325257608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d198952866080512054612a49565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260116004526024608051fd5b81806132d28b937f150b7a0200000000000000000000000000000000000000000000000000000000835233600484015260805160248401528b6044840152608060648401526084830190616309565b03816080518d5af160805191816133be575b5061332e5788886132f3617114565b8051918261332b57837f64a0ae9200000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b01fd5b7fffffffff000000000000000000000000000000000000000000000000000000007f150b7a020000000000000000000000000000000000000000000000000000000091999293949596979899160361338f5750908695949392918880612991565b7f64a0ae9200000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b9091508881813d831161340f575b6133d68183616487565b8101031261046557517fffffffff000000000000000000000000000000000000000000000000000000008116810361046557908a6132e4565b503d6133cc565b7f73c6ac6e00000000000000000000000000000000000000000000000000000000608051526080516004526024608051fd5b6134e3908a608051527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793048c526040608051207fffffffffffffffffffffffff0000000000000000000000000000000000000000815416905573ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190558b6128b3565b7f64a0ae9200000000000000000000000000000000000000000000000000000000608051526080516004526024608051fd5b6135599060203d6020116123e7576123da8183616487565b612813565b346104655760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655761359561634c565b5060243580151503610465577f8cd22d1900000000000000000000000000000000000000000000000000000000608051526004608051fd5b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1454604051908152f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557608051506040516080517f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015461368e81616835565b808452906001811690811561372d57506001146136ca575b610ba9836136b681850382616487565b604051918291602083526020830190616309565b608080517f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930190525160208120939250905b808210613713575090915081016020016136b66136a6565b9192600181602092548385880101520191019092916136fb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506136b690506136a6565b6080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465576137a26171c8565b3415611af8576040513481527f94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a60203392a260805180f35b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557613811617040565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095560805180f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0454604051908152f35b34610465576138a6366165f1565b90608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d176020526040608051209060805152602052602060ff60406080512054166040519015158152f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104655773ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015613a82575b6139e1903390616ba7565b6139e96171c8565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a160805180f35b506139e173ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416331490506139d6565b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416604051908152f35b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00541603613bb05761074b33617699565b7f118cdaa70000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b34610465576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261046557613c18617040565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00557f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005573ffffffffffffffffffffffffffffffffffffffff60805191167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0608051608051a360805180f35b3461046557613d3b36616729565b93613d4b969295969391936171c8565b8615611af8578115611af8578415611af857613d6a8783818114617662565b613d778786818114617662565b613d8087616f3b565b93613d8a88616f3b565b95613d9489616f3b565b946080515b8a8110613db45761074b8a8a8a613daf8f616f3b565b617362565b8581101561415a57613dc7818c84616f8a565b35613df233606084028801358173ffffffffffffffffffffffffffffffffffffffff610d6083616fd7565b6020606083028701013515611af8576040606083028701013515611af85760608202860135608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d17602052604060805120816080515260205260ff604060805120541661418b5780608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a60205260406080512054613e9782821515616bf1565b6060808402880136031261046557604051613eb181616417565b606084028801358152602081016020606086028a0101358152604082016040606087028b010135815260405191602083019386855251604084015251606083015251608082015260808152613f0760a082616487565b5190206040516020810191825260208152613f23604082616487565b519020908584101561415a578360051b8701357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112156104655787019067ffffffffffffffff82351161046557813560051b3603602083011361046557826080515b833581106141145750036140c057505080608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d19602052613fd960406080512060405190610f1b826163ce565b80421161408b57908291613fef6001948b616938565b5260608202870135608051527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d176020526040608051209060805152602052604060805120827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556020606082028701013561406f828b616938565b5260406060820287010135614084828c616938565b5201613d99565b7f431e9b4d00000000000000000000000000000000000000000000000000000000608051526080516004526024526044608051fd5b611052908860606040519687967f10f7f5f900000000000000000000000000000000000000000000000000000000885260048801520201356024850152608060448501526084840190602081359101616f9a565b9061412482853560208701616f8a565b359081811015614145576080515260205260016040608051205b9101613f8a565b9060805152602052600160406080512061413e565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260326004526024608051fd5b856060837f6bd4745f00000000000000000000000000000000000000000000000000000000608051520201356004526024526044608051fd5b346104655760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610465576141fe61058e61634c565b73ffffffffffffffffffffffffffffffffffffffff81161561426c5761426360209173ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b54604051908152f35b7f89c62b6400000000000000000000000000000000000000000000000000000000608051526080516004526024608051fd5b34610c405760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405760043567ffffffffffffffff8111610c40576142ed9036906004016166f8565b906142f661636f565b906142ff6171c8565b8215610cf95761430e83616f3b565b61431784616f3b565b915f73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416905b8681106144185750803b15610c40575f60405180927fb80f55c9000000000000000000000000000000000000000000000000000000008252602060048301528183816143a3602482018d8a616f9a565b03925af18015610c4457614404575b506143bc856168c0565b906143ca6040519283616487565b858252602082019560051b81019036821161046557955b8187106143f457505061074b9450617362565b86358152602096870196016143e1565b5f61440e91616487565b5f608052856143b2565b614423818885616f8a565b356040517f6352211e000000000000000000000000000000000000000000000000000000008152816004820152602081602481875afa908115610c44575f916145b1575b5073ffffffffffffffffffffffffffffffffffffffff3391160361458257604051907f74acd209000000000000000000000000000000000000000000000000000000008252806004830152606082602481875afa918215610c44575f9261452d575b506040820151908142116144ff57505090602082600193516144eb8489616938565b5201516144f88288616938565b5201614353565b7f431e9b4d000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b9091506060813d821161457a575b8161454860609383616487565b81010312610c4057604080519161455e83616417565b80518352602081015160208401520151604082015290896144c9565b3d915061453b565b7f6d3d1858000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b90506020813d82116145f7575b816145cb60209383616487565b81010312610c40575173ffffffffffffffffffffffffffffffffffffffff81168103610c405789614467565b3d91506145be565b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0054604051908152f35b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c40576020614695600435616fd7565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416604051908152f35b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405773ffffffffffffffffffffffffffffffffffffffff6147818161477361634c565b61477b617040565b166170ac565b16801561480d5761480b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0255565b005b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954604051908152f35b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057602060405160018152f35b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c40576149616106a461634c565b73ffffffffffffffffffffffffffffffffffffffff81161561480d5761480b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0655565b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003614a955760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c4057614acb366165f1565b90614ad46171c8565b8115610cf957614b3b8173ffffffffffffffffffffffffffffffffffffffff612b2b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16602052614b708260405f2054616905565b614ba07f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854828482821015616d92565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d166020528060405f2055614bf8837f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54616905565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a55614c238261721b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d02546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529190602090839060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af1918215610c44577fade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb479379092614d15575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5460408051958652602086019290925290840152339280606081015b0390a3005b614d2d9060203d6020116123e7576123da8183616487565b614cd2565b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057614d68616e6f565b50610ba9614d74616ee1565b604051918291829190916020806040830194805184520151910152565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057614dc361634c565b60243567ffffffffffffffff8111610c4057614de39036906004016166b4565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680301490811561504e575b50614a9557614e32617040565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f918161501a575b50614eb257837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc859203614fef5750813b15614fc457807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115614f93575f8083602061480b95519101845af4614f8d617114565b91617a13565b505034614f9c57005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011615046575b8161503660209383616487565b81010312610c4057519085614e81565b3d9150615029565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141583614e25565b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c40576150c6616e6f565b50610ba9614d74616e87565b34610c40576150e036616625565b916150e96171c8565b6150f561213082616fd7565b8115610cf957615104836170ac565b9273ffffffffffffffffffffffffffffffffffffffff84161561480d576152b290825f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d156020526020846151678160405f20546121c781838a82821115616e34565b9661519e7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095489888282108015906123fb57616d92565b855f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1583528760405f20556151f5827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54616886565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b556152208661721b565b73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416905f6040518097819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115610c44577f1e9c12c534d62150c4f5b916807e2bd6619095393fda9c833471a1feeefad6939273ffffffffffffffffffffffffffffffffffffffff92615346575b506153417f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5460405193849316973397846040919493926060820195825260208201520152565b0390a4005b61535e9060203d6020116123e7576123da8183616487565b6152fa565b34610c40576153713661650e565b5050505f604051615383602082616487565b527f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405773ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416331480156154e9575b615426903390616ba7565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff8116156154c1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b5061542673ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0654163314905061541b565b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405773ffffffffffffffffffffffffffffffffffffffff61557a61634c565b615582617040565b16801561480d5773ffffffffffffffffffffffffffffffffffffffff61169461480b926170ac565b34610c40576155b8366165f1565b906155c16171c8565b8115610cf9576156288173ffffffffffffffffffffffffffffffffffffffff612b2b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1560205261565d8260405f2054616905565b61568d7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954828482821015616d92565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d156020528060405f20556156e5837f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54616905565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b556157108261721b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d01546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529190602090839060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af1918215610c44577f8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a98492615801575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b546040805195865260208601929092529084015233928060608101614d10565b6158199060203d6020116123e7576123da8183616487565b6157bf565b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057615855616c48565b506004355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d19602052610ba960405f20600960405191615895836163ce565b805483526158a560018201616d2d565b6020840152600581015460408401526006810154606084015260078101546080840152600881015460a0840152015460c082015260405191829182616580565b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854604051908152f35b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405773ffffffffffffffffffffffffffffffffffffffff61598f8161477361634c565b16801561480d5761480b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0155565b34610c405760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057615a5061634c565b73ffffffffffffffffffffffffffffffffffffffff615a746024359261058e617040565b1690811561480d578015610cf9574791828211615a955761480b9250617143565b507fda0a959a000000000000000000000000000000000000000000000000000000005f525f60045260245260445260645ffd5b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54604051908152f35b34610c40576020615b32366164c8565b615b3a617040565b80517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5501517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d55005b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c40576004355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052602060405f2054604051908152f35b34610c4057615bf93661650e565b5050507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b6101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c40576004356101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610c405760405190615c8a826163ce565b602435825260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc360112610c4057600990604051615cc88161644f565b60443581526064356020820152608435604082015260a4356060820152602084019081526040840160c43581526060850160e43581526080860191610104358352606060a088019461012435865260c089019661014435885273ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015615e95575b615d6d903390616ba7565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d145490615d99826168d8565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1455615dc782821515616bf1565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a60205260405f20555f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1960205260405f209851895551805160018a0155602081015160028a0155604081015160038a0155015160048801555160058701555160068601555160078501555160088401555191015534615e6757005b6040513481527f94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a60203392a2005b50615d6d73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d06541633149050615d62565b34610c40576020615eea366164c8565b615ef2617040565b80517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5501517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f55005b34610c4057615f56615f4d36616392565b9291909161694c565b909160405191606083019360608452825180955260808401602060808760051b8701019401905f905b878210615f9a57505050839450602084015260408301520390f35b9091947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80878203018252855190602080835192838152019201905f905b808210615ff4575050506020806001929701920192019091615f7f565b909192602060a060019260808751805183528481015185840152604081015160408401526060810151606084015201516080820152019401920190615fd7565b34610c405760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c405761606b61634c565b507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c40576004356160cf81616fd7565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b34610c40575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c40576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005461617881616835565b808452906001811690811561372d575060011461619f57610ba9836136b681850382616487565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b808210616202575090915081016020016136b66136a6565b9192600181602092548385880101520191019092916161ea565b34610c405760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c4057600435907fffffffff000000000000000000000000000000000000000000000000000000008216809203610c4057817f80ac58cd00000000000000000000000000000000000000000000000000000000602093149081156162df575b81156162b5575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836162ae565b7f5b5e139f00000000000000000000000000000000000000000000000000000000811491506162a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610c4057565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203610c4057565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6080910112610c405760043590602435906044359060643590565b60e0810190811067ffffffffffffffff8211176163ea57604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6060810190811067ffffffffffffffff8211176163ea57604052565b6040810190811067ffffffffffffffff8211176163ea57604052565b6080810190811067ffffffffffffffff8211176163ea57604052565b60a0810190811067ffffffffffffffff8211176163ea57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176163ea57604052565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112610c40576040516164fe81616433565b6004358152602435602082015290565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6060910112610c405760043573ffffffffffffffffffffffffffffffffffffffff81168103610c40579060243573ffffffffffffffffffffffffffffffffffffffff81168103610c40579060443590565b91909161012060c0610140830194805184526165c260208201516020860190606080918051845260208101516020850152604081015160408501520151910152565b604081015160a0850152606081015182850152608081015160e085015260a08101516101008501520151910152565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112610c40576004359060243590565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6060910112610c4057600435906024359060443573ffffffffffffffffffffffffffffffffffffffff81168103610c405790565b67ffffffffffffffff81116163ea57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f82011215610c40576020813591016166ce8261667a565b926166dc6040519485616487565b82845282820111610c4057815f92602092838601378301015290565b9181601f84011215610c405782359167ffffffffffffffff8311610c40576020808501948460051b010111610c4057565b9060807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc830112610c405760043567ffffffffffffffff8111610c405782616773916004016166f8565b929092916024359167ffffffffffffffff8311610c405780602384011215610c405782600401359267ffffffffffffffff8411610c40578160246060860283010111610c4057602401929160443573ffffffffffffffffffffffffffffffffffffffff81168103610c4057916064359067ffffffffffffffff8211610c40576167fe916004016166f8565b9091565b90602080835192838152019201905f5b81811061681f5750505090565b8251845260209384019390920191600101616812565b90600182811c9216801561687c575b602083101461684f57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691616844565b9190820391821161689357565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff81116163ea5760051b60200190565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146168935760010190565b9190820180921161689357565b8054821015611096575f52600560205f20910201905f90565b8051156110965760200190565b80518210156110965760209160051b010190565b929391909361695b8486616886565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061699e616988836168c0565b926169966040519485616487565b8084526168c0565b015f5b818110616b965750509184905b86821080616b8d575b15616b8157815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1860205260405f209081549384821015616b73576169fd8183616905565b92858411616b6b575b616a108385616886565b92616a1a846168c0565b91616a286040519384616487565b8483527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0616a55866168c0565b015f5b818110616b365750505f5b858110616acc5750505090616a92616a989392616a808b88616886565b90616a8b828b616938565b5288616938565b50616886565b9283159081616ac2575b50616ab85750616ab25f916168d8565b906169ae565b9450945050929190565b905081105f616aa2565b80616ae2616adc60019386616905565b84616912565b50600460405191616af28361646b565b80548352848101546020840152600281015460408401526003810154606084015201546080820152616b248287616938565b52616b2f8186616938565b5001616a63565b602090604051616b458161646b565b5f81525f838201525f60408201525f60608201525f608082015282828801015201616a58565b859350616a06565b93505050616ab25f916168d8565b50929350509250905f90565b508215156169b7565b8060606020809386010152016169a1565b15616baf5750565b73ffffffffffffffffffffffffffffffffffffffff907f8e4a23d6000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b15616bf95750565b7f01707643000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b60405190616c318261644f565b5f6060838281528260208201528260408201520152565b60405190616c55826163ce565b5f60c083828152616c64616c24565b60208201528260408201528260608201528260808201528260a08201520152565b60405190616c928261644f565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d105482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d115460208301527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d125460408301527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d13546060830152565b90604051616d3a8161644f565b6060600382948054845260018101546020850152600281015460408501520154910152565b15616d675750565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b15616d9c57505050565b7f0877c890000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b90816020910312610c4057518015158103610c405790565b15616dee575050565b73ffffffffffffffffffffffffffffffffffffffff92507f3a0302a9000000000000000000000000000000000000000000000000000000005f526004521660245260445ffd5b15616e3e57505050565b7fda0a959a000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b60405190616e7c82616433565b5f6020838281520152565b60405190616e9482616433565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f546020830152565b60405190616eee82616433565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d546020830152565b90616f45826168c0565b616f526040519182616487565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0616f8082946168c0565b0190602036910137565b91908110156110965760051b0190565b90918281527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311610c405760209260051b809284830137010190565b61701f815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff821615616d67575090565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416330361708057565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014617100575b156170fc576170ee90617943565b906170f7575090565b905090565b5090565b505067ffffffffffffffff811660016170e0565b3d1561713e573d906171258261667a565b916171336040519384616487565b82523d5f602084013e565b606090565b8147106171a0575f80809373ffffffffffffffffffffffffffffffffffffffff8294165af1617170617114565b501561717857565b7f3204506f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f356680b7000000000000000000000000000000000000000000000000000000005f5260045ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166171f357565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1860205260405f2090805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1660205260405f205491815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1560205260405f20547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d145490604051936172d08561646b565b845260208401948552604084019081526060840191825260808401924284528054680100000000000000008110156163ea5761731191600182018155616912565b9590956173365760049451865551600186015551600285015551600384015551910155565b7f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b929173ffffffffffffffffffffffffffffffffffffffff617382846170ac565b161561480d575f915f948151925f5b84811061759057505050505073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0154166040517f70a08231000000000000000000000000000000000000000000000000000000008152336004820152602081602481855afa908115610c44575f9161755e575b5082811061752b575083471047906174fc57507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d05546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9091166024820152604481019290925290929190602090849060649082905f905af1908115610c44576174dd9373ffffffffffffffffffffffffffffffffffffffff926174df575b5016617143565b565b6174f79060203d6020116123e7576123da8183616487565b6174d6565b847f1acd938c000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b827f7520e30b000000000000000000000000000000000000000000000000000000005f523360045260245260445260645ffd5b90506020813d602011617588575b8161757960209383616487565b81010312610c4057515f617418565b3d915061756c565b61759d8183999799616938565b51906175a98186616938565b51917812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f228110830215617655576175f9816175f3600195670de0b6b3a76400008302908082049106151501809d616905565b99616905565b996176048387616938565b5191604051918252602082015273ffffffffffffffffffffffffffffffffffffffff8a16907f4dfd983654825b97cb5a6edb928f7e2a2726bef374ec4be8b39e4c31660103bb60403392a401617391565b637c5f487d5f526004601cfd5b1561766b575050565b7f7c8345c2000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b909283515f905f905b8082106178f6575050478111904791610c4f5750509173ffffffffffffffffffffffffffffffffffffffff5f610b1e95936178538296610b4e857f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416966040519a8b998a9889977f1d05159d000000000000000000000000000000000000000000000000000000008952166004880152608060248801526084870190616802565b03925af1908115610c44575f91617868575090565b90503d805f833e6178798183616487565b810190602081830312610c405780519067ffffffffffffffff8211610c4057019080601f83011215610c405781516178b0816168c0565b926178be6040519485616487565b81845260208085019260051b820101928311610c4057602001905b8282106178e65750505090565b81518152602091820191016178d9565b90916179028388616938565b519061790e8488616938565b519161791a8587616938565b519015610cf9578215610cf957804211610cca575060019161793b91616905565b9201906177b1565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff16036179b1575b821580156179a6575b61799e57565b5f9250829150565b5060163d1415617998565b5f925082915061798f565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c16156179eb57565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b90617a505750805115617a2857805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580617aa3575b617a61575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15617a5956fea2646970667358221220d4a5b58f059d4c6cdbe226162cae44b428f0cae58463b079415395e25ee3879c64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// LPPlusABI is the input ABI used to generate the binding from.
// Deprecated: Use LPPlusMetaData.ABI instead.
var LPPlusABI = LPPlusMetaData.ABI

// LPPlusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LPPlusMetaData.Bin instead.
var LPPlusBin = LPPlusMetaData.Bin

// DeployLPPlus deploys a new Ethereum contract, binding an instance of LPPlus to it.
func DeployLPPlus(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LPPlus, error) {
	parsed, err := LPPlusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LPPlusBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LPPlus{LPPlusCaller: LPPlusCaller{contract: contract}, LPPlusTransactor: LPPlusTransactor{contract: contract}, LPPlusFilterer: LPPlusFilterer{contract: contract}}, nil
}

// LPPlus is an auto generated Go binding around an Ethereum contract.
type LPPlus struct {
	LPPlusCaller     // Read-only binding to the contract
	LPPlusTransactor // Write-only binding to the contract
	LPPlusFilterer   // Log filterer for contract events
}

// LPPlusCaller is an auto generated read-only Go binding around an Ethereum contract.
type LPPlusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPPlusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LPPlusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPPlusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LPPlusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPPlusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LPPlusSession struct {
	Contract     *LPPlus           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LPPlusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LPPlusCallerSession struct {
	Contract *LPPlusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LPPlusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LPPlusTransactorSession struct {
	Contract     *LPPlusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LPPlusRaw is an auto generated low-level Go binding around an Ethereum contract.
type LPPlusRaw struct {
	Contract *LPPlus // Generic contract binding to access the raw methods on
}

// LPPlusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LPPlusCallerRaw struct {
	Contract *LPPlusCaller // Generic read-only contract binding to access the raw methods on
}

// LPPlusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LPPlusTransactorRaw struct {
	Contract *LPPlusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLPPlus creates a new instance of LPPlus, bound to a specific deployed contract.
func NewLPPlus(address common.Address, backend bind.ContractBackend) (*LPPlus, error) {
	contract, err := bindLPPlus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LPPlus{LPPlusCaller: LPPlusCaller{contract: contract}, LPPlusTransactor: LPPlusTransactor{contract: contract}, LPPlusFilterer: LPPlusFilterer{contract: contract}}, nil
}

// NewLPPlusCaller creates a new read-only instance of LPPlus, bound to a specific deployed contract.
func NewLPPlusCaller(address common.Address, caller bind.ContractCaller) (*LPPlusCaller, error) {
	contract, err := bindLPPlus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LPPlusCaller{contract: contract}, nil
}

// NewLPPlusTransactor creates a new write-only instance of LPPlus, bound to a specific deployed contract.
func NewLPPlusTransactor(address common.Address, transactor bind.ContractTransactor) (*LPPlusTransactor, error) {
	contract, err := bindLPPlus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LPPlusTransactor{contract: contract}, nil
}

// NewLPPlusFilterer creates a new log filterer instance of LPPlus, bound to a specific deployed contract.
func NewLPPlusFilterer(address common.Address, filterer bind.ContractFilterer) (*LPPlusFilterer, error) {
	contract, err := bindLPPlus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LPPlusFilterer{contract: contract}, nil
}

// bindLPPlus binds a generic wrapper to an already deployed contract.
func bindLPPlus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LPPlusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LPPlus *LPPlusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LPPlus.Contract.LPPlusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LPPlus *LPPlusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPPlus.Contract.LPPlusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LPPlus *LPPlusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LPPlus.Contract.LPPlusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LPPlus *LPPlusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LPPlus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LPPlus *LPPlusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPPlus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LPPlus *LPPlusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LPPlus.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LPPlus *LPPlusCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LPPlus *LPPlusSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LPPlus.Contract.UPGRADEINTERFACEVERSION(&_LPPlus.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LPPlus *LPPlusCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LPPlus.Contract.UPGRADEINTERFACEVERSION(&_LPPlus.CallOpts)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_LPPlus *LPPlusCaller) Approve(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) error {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "approve", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_LPPlus *LPPlusSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _LPPlus.Contract.Approve(&_LPPlus.CallOpts, arg0, arg1)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_LPPlus *LPPlusCallerSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _LPPlus.Contract.Approve(&_LPPlus.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LPPlus *LPPlusCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LPPlus *LPPlusSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LPPlus.Contract.BalanceOf(&_LPPlus.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LPPlus *LPPlusCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LPPlus.Contract.BalanceOf(&_LPPlus.CallOpts, owner)
}

// FutureValidityDuration is a free data retrieval call binding the contract method 0xb5ea015c.
//
// Solidity: function futureValidityDuration() view returns(uint256)
func (_LPPlus *LPPlusCaller) FutureValidityDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "futureValidityDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureValidityDuration is a free data retrieval call binding the contract method 0xb5ea015c.
//
// Solidity: function futureValidityDuration() view returns(uint256)
func (_LPPlus *LPPlusSession) FutureValidityDuration() (*big.Int, error) {
	return _LPPlus.Contract.FutureValidityDuration(&_LPPlus.CallOpts)
}

// FutureValidityDuration is a free data retrieval call binding the contract method 0xb5ea015c.
//
// Solidity: function futureValidityDuration() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) FutureValidityDuration() (*big.Int, error) {
	return _LPPlus.Contract.FutureValidityDuration(&_LPPlus.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LPPlus *LPPlusCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LPPlus *LPPlusSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LPPlus.Contract.GetApproved(&_LPPlus.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LPPlus *LPPlusCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LPPlus.Contract.GetApproved(&_LPPlus.CallOpts, tokenId)
}

// GetUserActionsBatched is a free data retrieval call binding the contract method 0x0a6b5699.
//
// Solidity: function getUserActionsBatched(uint256 _startTokenId, uint256 _endTokenId, uint256 _offset, uint256 _limit) view returns((uint256,uint256,uint256,uint256,uint256)[][] actionsPerToken, uint256 newStartTokenId, uint256 newOffset)
func (_LPPlus *LPPlusCaller) GetUserActionsBatched(opts *bind.CallOpts, _startTokenId *big.Int, _endTokenId *big.Int, _offset *big.Int, _limit *big.Int) (struct {
	ActionsPerToken [][]UserAction
	NewStartTokenId *big.Int
	NewOffset       *big.Int
}, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "getUserActionsBatched", _startTokenId, _endTokenId, _offset, _limit)

	outstruct := new(struct {
		ActionsPerToken [][]UserAction
		NewStartTokenId *big.Int
		NewOffset       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ActionsPerToken = *abi.ConvertType(out[0], new([][]UserAction)).(*[][]UserAction)
	outstruct.NewStartTokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewOffset = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserActionsBatched is a free data retrieval call binding the contract method 0x0a6b5699.
//
// Solidity: function getUserActionsBatched(uint256 _startTokenId, uint256 _endTokenId, uint256 _offset, uint256 _limit) view returns((uint256,uint256,uint256,uint256,uint256)[][] actionsPerToken, uint256 newStartTokenId, uint256 newOffset)
func (_LPPlus *LPPlusSession) GetUserActionsBatched(_startTokenId *big.Int, _endTokenId *big.Int, _offset *big.Int, _limit *big.Int) (struct {
	ActionsPerToken [][]UserAction
	NewStartTokenId *big.Int
	NewOffset       *big.Int
}, error) {
	return _LPPlus.Contract.GetUserActionsBatched(&_LPPlus.CallOpts, _startTokenId, _endTokenId, _offset, _limit)
}

// GetUserActionsBatched is a free data retrieval call binding the contract method 0x0a6b5699.
//
// Solidity: function getUserActionsBatched(uint256 _startTokenId, uint256 _endTokenId, uint256 _offset, uint256 _limit) view returns((uint256,uint256,uint256,uint256,uint256)[][] actionsPerToken, uint256 newStartTokenId, uint256 newOffset)
func (_LPPlus *LPPlusCallerSession) GetUserActionsBatched(_startTokenId *big.Int, _endTokenId *big.Int, _offset *big.Int, _limit *big.Int) (struct {
	ActionsPerToken [][]UserAction
	NewStartTokenId *big.Int
	NewOffset       *big.Int
}, error) {
	return _LPPlus.Contract.GetUserActionsBatched(&_LPPlus.CallOpts, _startTokenId, _endTokenId, _offset, _limit)
}

// InterpolationParams is a free data retrieval call binding the contract method 0xe0ea4a1b.
//
// Solidity: function interpolationParams() view returns((uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusCaller) InterpolationParams(opts *bind.CallOpts) (RwtInterpolation, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "interpolationParams")

	if err != nil {
		return *new(RwtInterpolation), err
	}

	out0 := *abi.ConvertType(out[0], new(RwtInterpolation)).(*RwtInterpolation)

	return out0, err

}

// InterpolationParams is a free data retrieval call binding the contract method 0xe0ea4a1b.
//
// Solidity: function interpolationParams() view returns((uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusSession) InterpolationParams() (RwtInterpolation, error) {
	return _LPPlus.Contract.InterpolationParams(&_LPPlus.CallOpts)
}

// InterpolationParams is a free data retrieval call binding the contract method 0xe0ea4a1b.
//
// Solidity: function interpolationParams() view returns((uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusCallerSession) InterpolationParams() (RwtInterpolation, error) {
	return _LPPlus.Contract.InterpolationParams(&_LPPlus.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LPPlus *LPPlusCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LPPlus *LPPlusSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LPPlus.Contract.IsApprovedForAll(&_LPPlus.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LPPlus *LPPlusCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LPPlus.Contract.IsApprovedForAll(&_LPPlus.CallOpts, owner, operator)
}

// MinimumRWTBalance is a free data retrieval call binding the contract method 0x58f28a98.
//
// Solidity: function minimumRWTBalance() view returns(uint256)
func (_LPPlus *LPPlusCaller) MinimumRWTBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "minimumRWTBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumRWTBalance is a free data retrieval call binding the contract method 0x58f28a98.
//
// Solidity: function minimumRWTBalance() view returns(uint256)
func (_LPPlus *LPPlusSession) MinimumRWTBalance() (*big.Int, error) {
	return _LPPlus.Contract.MinimumRWTBalance(&_LPPlus.CallOpts)
}

// MinimumRWTBalance is a free data retrieval call binding the contract method 0x58f28a98.
//
// Solidity: function minimumRWTBalance() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) MinimumRWTBalance() (*big.Int, error) {
	return _LPPlus.Contract.MinimumRWTBalance(&_LPPlus.CallOpts)
}

// MinimumYBTBalance is a free data retrieval call binding the contract method 0x3263edee.
//
// Solidity: function minimumYBTBalance() view returns(uint256)
func (_LPPlus *LPPlusCaller) MinimumYBTBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "minimumYBTBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumYBTBalance is a free data retrieval call binding the contract method 0x3263edee.
//
// Solidity: function minimumYBTBalance() view returns(uint256)
func (_LPPlus *LPPlusSession) MinimumYBTBalance() (*big.Int, error) {
	return _LPPlus.Contract.MinimumYBTBalance(&_LPPlus.CallOpts)
}

// MinimumYBTBalance is a free data retrieval call binding the contract method 0x3263edee.
//
// Solidity: function minimumYBTBalance() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) MinimumYBTBalance() (*big.Int, error) {
	return _LPPlus.Contract.MinimumYBTBalance(&_LPPlus.CallOpts)
}

// MintBonus is a free data retrieval call binding the contract method 0x510402cd.
//
// Solidity: function mintBonus() view returns((uint256,uint256))
func (_LPPlus *LPPlusCaller) MintBonus(opts *bind.CallOpts) (RWTFutureBonus, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "mintBonus")

	if err != nil {
		return *new(RWTFutureBonus), err
	}

	out0 := *abi.ConvertType(out[0], new(RWTFutureBonus)).(*RWTFutureBonus)

	return out0, err

}

// MintBonus is a free data retrieval call binding the contract method 0x510402cd.
//
// Solidity: function mintBonus() view returns((uint256,uint256))
func (_LPPlus *LPPlusSession) MintBonus() (RWTFutureBonus, error) {
	return _LPPlus.Contract.MintBonus(&_LPPlus.CallOpts)
}

// MintBonus is a free data retrieval call binding the contract method 0x510402cd.
//
// Solidity: function mintBonus() view returns((uint256,uint256))
func (_LPPlus *LPPlusCallerSession) MintBonus() (RWTFutureBonus, error) {
	return _LPPlus.Contract.MintBonus(&_LPPlus.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_LPPlus *LPPlusCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_LPPlus *LPPlusSession) MintPrice() (*big.Int, error) {
	return _LPPlus.Contract.MintPrice(&_LPPlus.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) MintPrice() (*big.Int, error) {
	return _LPPlus.Contract.MintPrice(&_LPPlus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LPPlus *LPPlusCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LPPlus *LPPlusSession) Name() (string, error) {
	return _LPPlus.Contract.Name(&_LPPlus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LPPlus *LPPlusCallerSession) Name() (string, error) {
	return _LPPlus.Contract.Name(&_LPPlus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPPlus *LPPlusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPPlus *LPPlusSession) Owner() (common.Address, error) {
	return _LPPlus.Contract.Owner(&_LPPlus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPPlus *LPPlusCallerSession) Owner() (common.Address, error) {
	return _LPPlus.Contract.Owner(&_LPPlus.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LPPlus *LPPlusCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LPPlus *LPPlusSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LPPlus.Contract.OwnerOf(&_LPPlus.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LPPlus *LPPlusCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LPPlus.Contract.OwnerOf(&_LPPlus.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LPPlus *LPPlusCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LPPlus *LPPlusSession) Paused() (bool, error) {
	return _LPPlus.Contract.Paused(&_LPPlus.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LPPlus *LPPlusCallerSession) Paused() (bool, error) {
	return _LPPlus.Contract.Paused(&_LPPlus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LPPlus *LPPlusCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LPPlus *LPPlusSession) PendingOwner() (common.Address, error) {
	return _LPPlus.Contract.PendingOwner(&_LPPlus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LPPlus *LPPlusCallerSession) PendingOwner() (common.Address, error) {
	return _LPPlus.Contract.PendingOwner(&_LPPlus.CallOpts)
}

// PreviewStakingSnapshot is a free data retrieval call binding the contract method 0xd262f444.
//
// Solidity: function previewStakingSnapshot(uint256 _baseConversionRateFilPerRwt, uint256 _stakingTotalYbt, uint256 _timestamp, uint256 _totalFilToAllocate) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusCaller) PreviewStakingSnapshot(opts *bind.CallOpts, _baseConversionRateFilPerRwt *big.Int, _stakingTotalYbt *big.Int, _timestamp *big.Int, _totalFilToAllocate *big.Int) (StakingSnapshot, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "previewStakingSnapshot", _baseConversionRateFilPerRwt, _stakingTotalYbt, _timestamp, _totalFilToAllocate)

	if err != nil {
		return *new(StakingSnapshot), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingSnapshot)).(*StakingSnapshot)

	return out0, err

}

// PreviewStakingSnapshot is a free data retrieval call binding the contract method 0xd262f444.
//
// Solidity: function previewStakingSnapshot(uint256 _baseConversionRateFilPerRwt, uint256 _stakingTotalYbt, uint256 _timestamp, uint256 _totalFilToAllocate) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusSession) PreviewStakingSnapshot(_baseConversionRateFilPerRwt *big.Int, _stakingTotalYbt *big.Int, _timestamp *big.Int, _totalFilToAllocate *big.Int) (StakingSnapshot, error) {
	return _LPPlus.Contract.PreviewStakingSnapshot(&_LPPlus.CallOpts, _baseConversionRateFilPerRwt, _stakingTotalYbt, _timestamp, _totalFilToAllocate)
}

// PreviewStakingSnapshot is a free data retrieval call binding the contract method 0xd262f444.
//
// Solidity: function previewStakingSnapshot(uint256 _baseConversionRateFilPerRwt, uint256 _stakingTotalYbt, uint256 _timestamp, uint256 _totalFilToAllocate) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusCallerSession) PreviewStakingSnapshot(_baseConversionRateFilPerRwt *big.Int, _stakingTotalYbt *big.Int, _timestamp *big.Int, _totalFilToAllocate *big.Int) (StakingSnapshot, error) {
	return _LPPlus.Contract.PreviewStakingSnapshot(&_LPPlus.CallOpts, _baseConversionRateFilPerRwt, _stakingTotalYbt, _timestamp, _totalFilToAllocate)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LPPlus *LPPlusCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LPPlus *LPPlusSession) ProxiableUUID() ([32]byte, error) {
	return _LPPlus.Contract.ProxiableUUID(&_LPPlus.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LPPlus *LPPlusCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LPPlus.Contract.ProxiableUUID(&_LPPlus.CallOpts)
}

// ReferrerBonus is a free data retrieval call binding the contract method 0x4bae2ef1.
//
// Solidity: function referrerBonus() view returns((uint256,uint256))
func (_LPPlus *LPPlusCaller) ReferrerBonus(opts *bind.CallOpts) (RWTFutureBonus, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "referrerBonus")

	if err != nil {
		return *new(RWTFutureBonus), err
	}

	out0 := *abi.ConvertType(out[0], new(RWTFutureBonus)).(*RWTFutureBonus)

	return out0, err

}

// ReferrerBonus is a free data retrieval call binding the contract method 0x4bae2ef1.
//
// Solidity: function referrerBonus() view returns((uint256,uint256))
func (_LPPlus *LPPlusSession) ReferrerBonus() (RWTFutureBonus, error) {
	return _LPPlus.Contract.ReferrerBonus(&_LPPlus.CallOpts)
}

// ReferrerBonus is a free data retrieval call binding the contract method 0x4bae2ef1.
//
// Solidity: function referrerBonus() view returns((uint256,uint256))
func (_LPPlus *LPPlusCallerSession) ReferrerBonus() (RWTFutureBonus, error) {
	return _LPPlus.Contract.ReferrerBonus(&_LPPlus.CallOpts)
}

// RwtFuture is a free data retrieval call binding the contract method 0xd5d2ae72.
//
// Solidity: function rwtFuture() view returns(address)
func (_LPPlus *LPPlusCaller) RwtFuture(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "rwtFuture")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RwtFuture is a free data retrieval call binding the contract method 0xd5d2ae72.
//
// Solidity: function rwtFuture() view returns(address)
func (_LPPlus *LPPlusSession) RwtFuture() (common.Address, error) {
	return _LPPlus.Contract.RwtFuture(&_LPPlus.CallOpts)
}

// RwtFuture is a free data retrieval call binding the contract method 0xd5d2ae72.
//
// Solidity: function rwtFuture() view returns(address)
func (_LPPlus *LPPlusCallerSession) RwtFuture() (common.Address, error) {
	return _LPPlus.Contract.RwtFuture(&_LPPlus.CallOpts)
}

// RwtToken is a free data retrieval call binding the contract method 0x812bf90a.
//
// Solidity: function rwtToken() view returns(address)
func (_LPPlus *LPPlusCaller) RwtToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "rwtToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RwtToken is a free data retrieval call binding the contract method 0x812bf90a.
//
// Solidity: function rwtToken() view returns(address)
func (_LPPlus *LPPlusSession) RwtToken() (common.Address, error) {
	return _LPPlus.Contract.RwtToken(&_LPPlus.CallOpts)
}

// RwtToken is a free data retrieval call binding the contract method 0x812bf90a.
//
// Solidity: function rwtToken() view returns(address)
func (_LPPlus *LPPlusCallerSession) RwtToken() (common.Address, error) {
	return _LPPlus.Contract.RwtToken(&_LPPlus.CallOpts)
}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_LPPlus *LPPlusCaller) SafeTransferFrom0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "safeTransferFrom0", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_LPPlus *LPPlusSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _LPPlus.Contract.SafeTransferFrom0(&_LPPlus.CallOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_LPPlus *LPPlusCallerSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _LPPlus.Contract.SafeTransferFrom0(&_LPPlus.CallOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_LPPlus *LPPlusCaller) SetApprovalForAll(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "setApprovalForAll", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_LPPlus *LPPlusSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _LPPlus.Contract.SetApprovalForAll(&_LPPlus.CallOpts, arg0, arg1)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_LPPlus *LPPlusCallerSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _LPPlus.Contract.SetApprovalForAll(&_LPPlus.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LPPlus *LPPlusCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LPPlus *LPPlusSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LPPlus.Contract.SupportsInterface(&_LPPlus.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LPPlus *LPPlusCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LPPlus.Contract.SupportsInterface(&_LPPlus.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LPPlus *LPPlusCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LPPlus *LPPlusSession) Symbol() (string, error) {
	return _LPPlus.Contract.Symbol(&_LPPlus.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LPPlus *LPPlusCallerSession) Symbol() (string, error) {
	return _LPPlus.Contract.Symbol(&_LPPlus.CallOpts)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_LPPlus *LPPlusCaller) TokenIdGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "tokenIdGenerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_LPPlus *LPPlusSession) TokenIdGenerator() (*big.Int, error) {
	return _LPPlus.Contract.TokenIdGenerator(&_LPPlus.CallOpts)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TokenIdGenerator() (*big.Int, error) {
	return _LPPlus.Contract.TokenIdGenerator(&_LPPlus.CallOpts)
}

// TokenIdToRWTBalance is a free data retrieval call binding the contract method 0x29f2014d.
//
// Solidity: function tokenIdToRWTBalance(uint256 _tokenId) view returns(uint256)
func (_LPPlus *LPPlusCaller) TokenIdToRWTBalance(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "tokenIdToRWTBalance", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToRWTBalance is a free data retrieval call binding the contract method 0x29f2014d.
//
// Solidity: function tokenIdToRWTBalance(uint256 _tokenId) view returns(uint256)
func (_LPPlus *LPPlusSession) TokenIdToRWTBalance(_tokenId *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenIdToRWTBalance(&_LPPlus.CallOpts, _tokenId)
}

// TokenIdToRWTBalance is a free data retrieval call binding the contract method 0x29f2014d.
//
// Solidity: function tokenIdToRWTBalance(uint256 _tokenId) view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TokenIdToRWTBalance(_tokenId *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenIdToRWTBalance(&_LPPlus.CallOpts, _tokenId)
}

// TokenIdToWindowIdToClaimedFlag is a free data retrieval call binding the contract method 0x8edd7d4e.
//
// Solidity: function tokenIdToWindowIdToClaimedFlag(uint256 _tokenId, uint256 _windowId) view returns(bool)
func (_LPPlus *LPPlusCaller) TokenIdToWindowIdToClaimedFlag(opts *bind.CallOpts, _tokenId *big.Int, _windowId *big.Int) (bool, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "tokenIdToWindowIdToClaimedFlag", _tokenId, _windowId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenIdToWindowIdToClaimedFlag is a free data retrieval call binding the contract method 0x8edd7d4e.
//
// Solidity: function tokenIdToWindowIdToClaimedFlag(uint256 _tokenId, uint256 _windowId) view returns(bool)
func (_LPPlus *LPPlusSession) TokenIdToWindowIdToClaimedFlag(_tokenId *big.Int, _windowId *big.Int) (bool, error) {
	return _LPPlus.Contract.TokenIdToWindowIdToClaimedFlag(&_LPPlus.CallOpts, _tokenId, _windowId)
}

// TokenIdToWindowIdToClaimedFlag is a free data retrieval call binding the contract method 0x8edd7d4e.
//
// Solidity: function tokenIdToWindowIdToClaimedFlag(uint256 _tokenId, uint256 _windowId) view returns(bool)
func (_LPPlus *LPPlusCallerSession) TokenIdToWindowIdToClaimedFlag(_tokenId *big.Int, _windowId *big.Int) (bool, error) {
	return _LPPlus.Contract.TokenIdToWindowIdToClaimedFlag(&_LPPlus.CallOpts, _tokenId, _windowId)
}

// TokenIdToYBTBalance is a free data retrieval call binding the contract method 0xae7e2d97.
//
// Solidity: function tokenIdToYBTBalance(uint256 _tokenId) view returns(uint256)
func (_LPPlus *LPPlusCaller) TokenIdToYBTBalance(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "tokenIdToYBTBalance", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToYBTBalance is a free data retrieval call binding the contract method 0xae7e2d97.
//
// Solidity: function tokenIdToYBTBalance(uint256 _tokenId) view returns(uint256)
func (_LPPlus *LPPlusSession) TokenIdToYBTBalance(_tokenId *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenIdToYBTBalance(&_LPPlus.CallOpts, _tokenId)
}

// TokenIdToYBTBalance is a free data retrieval call binding the contract method 0xae7e2d97.
//
// Solidity: function tokenIdToYBTBalance(uint256 _tokenId) view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TokenIdToYBTBalance(_tokenId *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenIdToYBTBalance(&_LPPlus.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LPPlus *LPPlusCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LPPlus *LPPlusSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LPPlus.Contract.TokenURI(&_LPPlus.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LPPlus *LPPlusCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LPPlus.Contract.TokenURI(&_LPPlus.CallOpts, tokenId)
}

// TotalRwtStaked is a free data retrieval call binding the contract method 0x2e58c0c7.
//
// Solidity: function totalRwtStaked() view returns(uint256)
func (_LPPlus *LPPlusCaller) TotalRwtStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "totalRwtStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRwtStaked is a free data retrieval call binding the contract method 0x2e58c0c7.
//
// Solidity: function totalRwtStaked() view returns(uint256)
func (_LPPlus *LPPlusSession) TotalRwtStaked() (*big.Int, error) {
	return _LPPlus.Contract.TotalRwtStaked(&_LPPlus.CallOpts)
}

// TotalRwtStaked is a free data retrieval call binding the contract method 0x2e58c0c7.
//
// Solidity: function totalRwtStaked() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TotalRwtStaked() (*big.Int, error) {
	return _LPPlus.Contract.TotalRwtStaked(&_LPPlus.CallOpts)
}

// TotalYbtStaked is a free data retrieval call binding the contract method 0xa818b1ee.
//
// Solidity: function totalYbtStaked() view returns(uint256)
func (_LPPlus *LPPlusCaller) TotalYbtStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "totalYbtStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalYbtStaked is a free data retrieval call binding the contract method 0xa818b1ee.
//
// Solidity: function totalYbtStaked() view returns(uint256)
func (_LPPlus *LPPlusSession) TotalYbtStaked() (*big.Int, error) {
	return _LPPlus.Contract.TotalYbtStaked(&_LPPlus.CallOpts)
}

// TotalYbtStaked is a free data retrieval call binding the contract method 0xa818b1ee.
//
// Solidity: function totalYbtStaked() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TotalYbtStaked() (*big.Int, error) {
	return _LPPlus.Contract.TotalYbtStaked(&_LPPlus.CallOpts)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_LPPlus *LPPlusCaller) TransferFrom(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "transferFrom", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_LPPlus *LPPlusSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _LPPlus.Contract.TransferFrom(&_LPPlus.CallOpts, arg0, arg1, arg2)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_LPPlus *LPPlusCallerSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _LPPlus.Contract.TransferFrom(&_LPPlus.CallOpts, arg0, arg1, arg2)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_LPPlus *LPPlusCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_LPPlus *LPPlusSession) Treasury() (common.Address, error) {
	return _LPPlus.Contract.Treasury(&_LPPlus.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_LPPlus *LPPlusCallerSession) Treasury() (common.Address, error) {
	return _LPPlus.Contract.Treasury(&_LPPlus.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_LPPlus *LPPlusCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_LPPlus *LPPlusSession) Version() (*big.Int, error) {
	return _LPPlus.Contract.Version(&_LPPlus.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_LPPlus *LPPlusCallerSession) Version() (*big.Int, error) {
	return _LPPlus.Contract.Version(&_LPPlus.CallOpts)
}

// WindowCommitter is a free data retrieval call binding the contract method 0xf721dc76.
//
// Solidity: function windowCommitter() view returns(address)
func (_LPPlus *LPPlusCaller) WindowCommitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "windowCommitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WindowCommitter is a free data retrieval call binding the contract method 0xf721dc76.
//
// Solidity: function windowCommitter() view returns(address)
func (_LPPlus *LPPlusSession) WindowCommitter() (common.Address, error) {
	return _LPPlus.Contract.WindowCommitter(&_LPPlus.CallOpts)
}

// WindowCommitter is a free data retrieval call binding the contract method 0xf721dc76.
//
// Solidity: function windowCommitter() view returns(address)
func (_LPPlus *LPPlusCallerSession) WindowCommitter() (common.Address, error) {
	return _LPPlus.Contract.WindowCommitter(&_LPPlus.CallOpts)
}

// WindowId is a free data retrieval call binding the contract method 0x9af1d03e.
//
// Solidity: function windowId() view returns(uint256)
func (_LPPlus *LPPlusCaller) WindowId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "windowId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WindowId is a free data retrieval call binding the contract method 0x9af1d03e.
//
// Solidity: function windowId() view returns(uint256)
func (_LPPlus *LPPlusSession) WindowId() (*big.Int, error) {
	return _LPPlus.Contract.WindowId(&_LPPlus.CallOpts)
}

// WindowId is a free data retrieval call binding the contract method 0x9af1d03e.
//
// Solidity: function windowId() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) WindowId() (*big.Int, error) {
	return _LPPlus.Contract.WindowId(&_LPPlus.CallOpts)
}

// WindowIdToMerkleRoot is a free data retrieval call binding the contract method 0xb05dcd5a.
//
// Solidity: function windowIdToMerkleRoot(uint256 _windowId) view returns(bytes32)
func (_LPPlus *LPPlusCaller) WindowIdToMerkleRoot(opts *bind.CallOpts, _windowId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "windowIdToMerkleRoot", _windowId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WindowIdToMerkleRoot is a free data retrieval call binding the contract method 0xb05dcd5a.
//
// Solidity: function windowIdToMerkleRoot(uint256 _windowId) view returns(bytes32)
func (_LPPlus *LPPlusSession) WindowIdToMerkleRoot(_windowId *big.Int) ([32]byte, error) {
	return _LPPlus.Contract.WindowIdToMerkleRoot(&_LPPlus.CallOpts, _windowId)
}

// WindowIdToMerkleRoot is a free data retrieval call binding the contract method 0xb05dcd5a.
//
// Solidity: function windowIdToMerkleRoot(uint256 _windowId) view returns(bytes32)
func (_LPPlus *LPPlusCallerSession) WindowIdToMerkleRoot(_windowId *big.Int) ([32]byte, error) {
	return _LPPlus.Contract.WindowIdToMerkleRoot(&_LPPlus.CallOpts, _windowId)
}

// WindowIdToStakingSnapshot is a free data retrieval call binding the contract method 0x344bb6b1.
//
// Solidity: function windowIdToStakingSnapshot(uint256 _windowId) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusCaller) WindowIdToStakingSnapshot(opts *bind.CallOpts, _windowId *big.Int) (StakingSnapshot, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "windowIdToStakingSnapshot", _windowId)

	if err != nil {
		return *new(StakingSnapshot), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingSnapshot)).(*StakingSnapshot)

	return out0, err

}

// WindowIdToStakingSnapshot is a free data retrieval call binding the contract method 0x344bb6b1.
//
// Solidity: function windowIdToStakingSnapshot(uint256 _windowId) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusSession) WindowIdToStakingSnapshot(_windowId *big.Int) (StakingSnapshot, error) {
	return _LPPlus.Contract.WindowIdToStakingSnapshot(&_LPPlus.CallOpts, _windowId)
}

// WindowIdToStakingSnapshot is a free data retrieval call binding the contract method 0x344bb6b1.
//
// Solidity: function windowIdToStakingSnapshot(uint256 _windowId) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusCallerSession) WindowIdToStakingSnapshot(_windowId *big.Int) (StakingSnapshot, error) {
	return _LPPlus.Contract.WindowIdToStakingSnapshot(&_LPPlus.CallOpts, _windowId)
}

// YbtToken is a free data retrieval call binding the contract method 0xee2df9fd.
//
// Solidity: function ybtToken() view returns(address)
func (_LPPlus *LPPlusCaller) YbtToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "ybtToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YbtToken is a free data retrieval call binding the contract method 0xee2df9fd.
//
// Solidity: function ybtToken() view returns(address)
func (_LPPlus *LPPlusSession) YbtToken() (common.Address, error) {
	return _LPPlus.Contract.YbtToken(&_LPPlus.CallOpts)
}

// YbtToken is a free data retrieval call binding the contract method 0xee2df9fd.
//
// Solidity: function ybtToken() view returns(address)
func (_LPPlus *LPPlusCallerSession) YbtToken() (common.Address, error) {
	return _LPPlus.Contract.YbtToken(&_LPPlus.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LPPlus *LPPlusTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LPPlus *LPPlusSession) AcceptOwnership() (*types.Transaction, error) {
	return _LPPlus.Contract.AcceptOwnership(&_LPPlus.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LPPlus *LPPlusTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LPPlus.Contract.AcceptOwnership(&_LPPlus.TransactOpts)
}

// AddRWT is a paid mutator transaction binding the contract method 0x385b5f85.
//
// Solidity: function addRWT(uint256 _tokenId, uint256 _RWTAmount) returns()
func (_LPPlus *LPPlusTransactor) AddRWT(opts *bind.TransactOpts, _tokenId *big.Int, _RWTAmount *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "addRWT", _tokenId, _RWTAmount)
}

// AddRWT is a paid mutator transaction binding the contract method 0x385b5f85.
//
// Solidity: function addRWT(uint256 _tokenId, uint256 _RWTAmount) returns()
func (_LPPlus *LPPlusSession) AddRWT(_tokenId *big.Int, _RWTAmount *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.AddRWT(&_LPPlus.TransactOpts, _tokenId, _RWTAmount)
}

// AddRWT is a paid mutator transaction binding the contract method 0x385b5f85.
//
// Solidity: function addRWT(uint256 _tokenId, uint256 _RWTAmount) returns()
func (_LPPlus *LPPlusTransactorSession) AddRWT(_tokenId *big.Int, _RWTAmount *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.AddRWT(&_LPPlus.TransactOpts, _tokenId, _RWTAmount)
}

// AddYBT is a paid mutator transaction binding the contract method 0x51947361.
//
// Solidity: function addYBT(uint256 _tokenId, uint256 _YBTAmount) returns()
func (_LPPlus *LPPlusTransactor) AddYBT(opts *bind.TransactOpts, _tokenId *big.Int, _YBTAmount *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "addYBT", _tokenId, _YBTAmount)
}

// AddYBT is a paid mutator transaction binding the contract method 0x51947361.
//
// Solidity: function addYBT(uint256 _tokenId, uint256 _YBTAmount) returns()
func (_LPPlus *LPPlusSession) AddYBT(_tokenId *big.Int, _YBTAmount *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.AddYBT(&_LPPlus.TransactOpts, _tokenId, _YBTAmount)
}

// AddYBT is a paid mutator transaction binding the contract method 0x51947361.
//
// Solidity: function addYBT(uint256 _tokenId, uint256 _YBTAmount) returns()
func (_LPPlus *LPPlusTransactorSession) AddYBT(_tokenId *big.Int, _YBTAmount *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.AddYBT(&_LPPlus.TransactOpts, _tokenId, _YBTAmount)
}

// Claim is a paid mutator transaction binding the contract method 0xda357c0f.
//
// Solidity: function claim(uint256[] _windowIds, (uint256,uint256,uint256)[] _userSnapshots, address _receiver, bytes32[][] _proofs) returns(uint256[] tokenIds)
func (_LPPlus *LPPlusTransactor) Claim(opts *bind.TransactOpts, _windowIds []*big.Int, _userSnapshots []LpPlusMerkleHelperUserSnapshot, _receiver common.Address, _proofs [][][32]byte) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "claim", _windowIds, _userSnapshots, _receiver, _proofs)
}

// Claim is a paid mutator transaction binding the contract method 0xda357c0f.
//
// Solidity: function claim(uint256[] _windowIds, (uint256,uint256,uint256)[] _userSnapshots, address _receiver, bytes32[][] _proofs) returns(uint256[] tokenIds)
func (_LPPlus *LPPlusSession) Claim(_windowIds []*big.Int, _userSnapshots []LpPlusMerkleHelperUserSnapshot, _receiver common.Address, _proofs [][][32]byte) (*types.Transaction, error) {
	return _LPPlus.Contract.Claim(&_LPPlus.TransactOpts, _windowIds, _userSnapshots, _receiver, _proofs)
}

// Claim is a paid mutator transaction binding the contract method 0xda357c0f.
//
// Solidity: function claim(uint256[] _windowIds, (uint256,uint256,uint256)[] _userSnapshots, address _receiver, bytes32[][] _proofs) returns(uint256[] tokenIds)
func (_LPPlus *LPPlusTransactorSession) Claim(_windowIds []*big.Int, _userSnapshots []LpPlusMerkleHelperUserSnapshot, _receiver common.Address, _proofs [][][32]byte) (*types.Transaction, error) {
	return _LPPlus.Contract.Claim(&_LPPlus.TransactOpts, _windowIds, _userSnapshots, _receiver, _proofs)
}

// ClaimAndExecute is a paid mutator transaction binding the contract method 0x70af96aa.
//
// Solidity: function claimAndExecute(uint256[] _windowIds, (uint256,uint256,uint256)[] _userSnapshots, address _receiver, bytes32[][] _proofs) returns()
func (_LPPlus *LPPlusTransactor) ClaimAndExecute(opts *bind.TransactOpts, _windowIds []*big.Int, _userSnapshots []LpPlusMerkleHelperUserSnapshot, _receiver common.Address, _proofs [][][32]byte) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "claimAndExecute", _windowIds, _userSnapshots, _receiver, _proofs)
}

// ClaimAndExecute is a paid mutator transaction binding the contract method 0x70af96aa.
//
// Solidity: function claimAndExecute(uint256[] _windowIds, (uint256,uint256,uint256)[] _userSnapshots, address _receiver, bytes32[][] _proofs) returns()
func (_LPPlus *LPPlusSession) ClaimAndExecute(_windowIds []*big.Int, _userSnapshots []LpPlusMerkleHelperUserSnapshot, _receiver common.Address, _proofs [][][32]byte) (*types.Transaction, error) {
	return _LPPlus.Contract.ClaimAndExecute(&_LPPlus.TransactOpts, _windowIds, _userSnapshots, _receiver, _proofs)
}

// ClaimAndExecute is a paid mutator transaction binding the contract method 0x70af96aa.
//
// Solidity: function claimAndExecute(uint256[] _windowIds, (uint256,uint256,uint256)[] _userSnapshots, address _receiver, bytes32[][] _proofs) returns()
func (_LPPlus *LPPlusTransactorSession) ClaimAndExecute(_windowIds []*big.Int, _userSnapshots []LpPlusMerkleHelperUserSnapshot, _receiver common.Address, _proofs [][][32]byte) (*types.Transaction, error) {
	return _LPPlus.Contract.ClaimAndExecute(&_LPPlus.TransactOpts, _windowIds, _userSnapshots, _receiver, _proofs)
}

// CommitWindow is a paid mutator transaction binding the contract method 0x16290e1d.
//
// Solidity: function commitWindow(bytes32 _root, (uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) _stakingSnapshot) payable returns()
func (_LPPlus *LPPlusTransactor) CommitWindow(opts *bind.TransactOpts, _root [32]byte, _stakingSnapshot StakingSnapshot) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "commitWindow", _root, _stakingSnapshot)
}

// CommitWindow is a paid mutator transaction binding the contract method 0x16290e1d.
//
// Solidity: function commitWindow(bytes32 _root, (uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) _stakingSnapshot) payable returns()
func (_LPPlus *LPPlusSession) CommitWindow(_root [32]byte, _stakingSnapshot StakingSnapshot) (*types.Transaction, error) {
	return _LPPlus.Contract.CommitWindow(&_LPPlus.TransactOpts, _root, _stakingSnapshot)
}

// CommitWindow is a paid mutator transaction binding the contract method 0x16290e1d.
//
// Solidity: function commitWindow(bytes32 _root, (uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) _stakingSnapshot) payable returns()
func (_LPPlus *LPPlusTransactorSession) CommitWindow(_root [32]byte, _stakingSnapshot StakingSnapshot) (*types.Transaction, error) {
	return _LPPlus.Contract.CommitWindow(&_LPPlus.TransactOpts, _root, _stakingSnapshot)
}

// Execute is a paid mutator transaction binding the contract method 0x694072a6.
//
// Solidity: function execute(uint256[] _rwtFutureTokenIds, address _receiver) returns()
func (_LPPlus *LPPlusTransactor) Execute(opts *bind.TransactOpts, _rwtFutureTokenIds []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "execute", _rwtFutureTokenIds, _receiver)
}

// Execute is a paid mutator transaction binding the contract method 0x694072a6.
//
// Solidity: function execute(uint256[] _rwtFutureTokenIds, address _receiver) returns()
func (_LPPlus *LPPlusSession) Execute(_rwtFutureTokenIds []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.Execute(&_LPPlus.TransactOpts, _rwtFutureTokenIds, _receiver)
}

// Execute is a paid mutator transaction binding the contract method 0x694072a6.
//
// Solidity: function execute(uint256[] _rwtFutureTokenIds, address _receiver) returns()
func (_LPPlus *LPPlusTransactorSession) Execute(_rwtFutureTokenIds []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.Execute(&_LPPlus.TransactOpts, _rwtFutureTokenIds, _receiver)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_LPPlus *LPPlusTransactor) FundFilVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "fundFilVault")
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_LPPlus *LPPlusSession) FundFilVault() (*types.Transaction, error) {
	return _LPPlus.Contract.FundFilVault(&_LPPlus.TransactOpts)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_LPPlus *LPPlusTransactorSession) FundFilVault() (*types.Transaction, error) {
	return _LPPlus.Contract.FundFilVault(&_LPPlus.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4521548.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _rwtToken, address _ybtToken, address _rwtFuture, address _treasury, address _windowCommitter) returns()
func (_LPPlus *LPPlusTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _rwtToken common.Address, _ybtToken common.Address, _rwtFuture common.Address, _treasury common.Address, _windowCommitter common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "initialize", _name, _symbol, _initialOwner, _mintPrice, _rwtToken, _ybtToken, _rwtFuture, _treasury, _windowCommitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4521548.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _rwtToken, address _ybtToken, address _rwtFuture, address _treasury, address _windowCommitter) returns()
func (_LPPlus *LPPlusSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _rwtToken common.Address, _ybtToken common.Address, _rwtFuture common.Address, _treasury common.Address, _windowCommitter common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.Initialize(&_LPPlus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _rwtToken, _ybtToken, _rwtFuture, _treasury, _windowCommitter)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4521548.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _rwtToken, address _ybtToken, address _rwtFuture, address _treasury, address _windowCommitter) returns()
func (_LPPlus *LPPlusTransactorSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _rwtToken common.Address, _ybtToken common.Address, _rwtFuture common.Address, _treasury common.Address, _windowCommitter common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.Initialize(&_LPPlus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _rwtToken, _ybtToken, _rwtFuture, _treasury, _windowCommitter)
}

// Mint is a paid mutator transaction binding the contract method 0xa647e8ec.
//
// Solidity: function mint(address _to, uint256 _referrerTokenId, uint256 _ybtToStake, uint256 _rwtToStake) returns(uint256 tokenId)
func (_LPPlus *LPPlusTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _referrerTokenId *big.Int, _ybtToStake *big.Int, _rwtToStake *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "mint", _to, _referrerTokenId, _ybtToStake, _rwtToStake)
}

// Mint is a paid mutator transaction binding the contract method 0xa647e8ec.
//
// Solidity: function mint(address _to, uint256 _referrerTokenId, uint256 _ybtToStake, uint256 _rwtToStake) returns(uint256 tokenId)
func (_LPPlus *LPPlusSession) Mint(_to common.Address, _referrerTokenId *big.Int, _ybtToStake *big.Int, _rwtToStake *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.Mint(&_LPPlus.TransactOpts, _to, _referrerTokenId, _ybtToStake, _rwtToStake)
}

// Mint is a paid mutator transaction binding the contract method 0xa647e8ec.
//
// Solidity: function mint(address _to, uint256 _referrerTokenId, uint256 _ybtToStake, uint256 _rwtToStake) returns(uint256 tokenId)
func (_LPPlus *LPPlusTransactorSession) Mint(_to common.Address, _referrerTokenId *big.Int, _ybtToStake *big.Int, _rwtToStake *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.Mint(&_LPPlus.TransactOpts, _to, _referrerTokenId, _ybtToStake, _rwtToStake)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LPPlus *LPPlusTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LPPlus *LPPlusSession) Pause() (*types.Transaction, error) {
	return _LPPlus.Contract.Pause(&_LPPlus.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LPPlus *LPPlusTransactorSession) Pause() (*types.Transaction, error) {
	return _LPPlus.Contract.Pause(&_LPPlus.TransactOpts)
}

// RemoveRWT is a paid mutator transaction binding the contract method 0x46b3241e.
//
// Solidity: function removeRWT(uint256 _tokenId, uint256 _RWTAmount, address _receiver) returns()
func (_LPPlus *LPPlusTransactor) RemoveRWT(opts *bind.TransactOpts, _tokenId *big.Int, _RWTAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "removeRWT", _tokenId, _RWTAmount, _receiver)
}

// RemoveRWT is a paid mutator transaction binding the contract method 0x46b3241e.
//
// Solidity: function removeRWT(uint256 _tokenId, uint256 _RWTAmount, address _receiver) returns()
func (_LPPlus *LPPlusSession) RemoveRWT(_tokenId *big.Int, _RWTAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.RemoveRWT(&_LPPlus.TransactOpts, _tokenId, _RWTAmount, _receiver)
}

// RemoveRWT is a paid mutator transaction binding the contract method 0x46b3241e.
//
// Solidity: function removeRWT(uint256 _tokenId, uint256 _RWTAmount, address _receiver) returns()
func (_LPPlus *LPPlusTransactorSession) RemoveRWT(_tokenId *big.Int, _RWTAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.RemoveRWT(&_LPPlus.TransactOpts, _tokenId, _RWTAmount, _receiver)
}

// RemoveYBT is a paid mutator transaction binding the contract method 0xbfb1909b.
//
// Solidity: function removeYBT(uint256 _tokenId, uint256 _YBTAmount, address _receiver) returns()
func (_LPPlus *LPPlusTransactor) RemoveYBT(opts *bind.TransactOpts, _tokenId *big.Int, _YBTAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "removeYBT", _tokenId, _YBTAmount, _receiver)
}

// RemoveYBT is a paid mutator transaction binding the contract method 0xbfb1909b.
//
// Solidity: function removeYBT(uint256 _tokenId, uint256 _YBTAmount, address _receiver) returns()
func (_LPPlus *LPPlusSession) RemoveYBT(_tokenId *big.Int, _YBTAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.RemoveYBT(&_LPPlus.TransactOpts, _tokenId, _YBTAmount, _receiver)
}

// RemoveYBT is a paid mutator transaction binding the contract method 0xbfb1909b.
//
// Solidity: function removeYBT(uint256 _tokenId, uint256 _YBTAmount, address _receiver) returns()
func (_LPPlus *LPPlusTransactorSession) RemoveYBT(_tokenId *big.Int, _YBTAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.RemoveYBT(&_LPPlus.TransactOpts, _tokenId, _YBTAmount, _receiver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPPlus *LPPlusTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPPlus *LPPlusSession) RenounceOwnership() (*types.Transaction, error) {
	return _LPPlus.Contract.RenounceOwnership(&_LPPlus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPPlus *LPPlusTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LPPlus.Contract.RenounceOwnership(&_LPPlus.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LPPlus *LPPlusTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LPPlus *LPPlusSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SafeTransferFrom(&_LPPlus.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LPPlus *LPPlusTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SafeTransferFrom(&_LPPlus.TransactOpts, from, to, tokenId)
}

// SetFutureValidityDuration is a paid mutator transaction binding the contract method 0xc74c23d5.
//
// Solidity: function setFutureValidityDuration(uint256 _futureValidityDuration) returns()
func (_LPPlus *LPPlusTransactor) SetFutureValidityDuration(opts *bind.TransactOpts, _futureValidityDuration *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setFutureValidityDuration", _futureValidityDuration)
}

// SetFutureValidityDuration is a paid mutator transaction binding the contract method 0xc74c23d5.
//
// Solidity: function setFutureValidityDuration(uint256 _futureValidityDuration) returns()
func (_LPPlus *LPPlusSession) SetFutureValidityDuration(_futureValidityDuration *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetFutureValidityDuration(&_LPPlus.TransactOpts, _futureValidityDuration)
}

// SetFutureValidityDuration is a paid mutator transaction binding the contract method 0xc74c23d5.
//
// Solidity: function setFutureValidityDuration(uint256 _futureValidityDuration) returns()
func (_LPPlus *LPPlusTransactorSession) SetFutureValidityDuration(_futureValidityDuration *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetFutureValidityDuration(&_LPPlus.TransactOpts, _futureValidityDuration)
}

// SetInterpolationParams is a paid mutator transaction binding the contract method 0xfcd0907d.
//
// Solidity: function setInterpolationParams((uint256,uint256,uint256,uint256) _interpolationParams) returns()
func (_LPPlus *LPPlusTransactor) SetInterpolationParams(opts *bind.TransactOpts, _interpolationParams RwtInterpolation) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setInterpolationParams", _interpolationParams)
}

// SetInterpolationParams is a paid mutator transaction binding the contract method 0xfcd0907d.
//
// Solidity: function setInterpolationParams((uint256,uint256,uint256,uint256) _interpolationParams) returns()
func (_LPPlus *LPPlusSession) SetInterpolationParams(_interpolationParams RwtInterpolation) (*types.Transaction, error) {
	return _LPPlus.Contract.SetInterpolationParams(&_LPPlus.TransactOpts, _interpolationParams)
}

// SetInterpolationParams is a paid mutator transaction binding the contract method 0xfcd0907d.
//
// Solidity: function setInterpolationParams((uint256,uint256,uint256,uint256) _interpolationParams) returns()
func (_LPPlus *LPPlusTransactorSession) SetInterpolationParams(_interpolationParams RwtInterpolation) (*types.Transaction, error) {
	return _LPPlus.Contract.SetInterpolationParams(&_LPPlus.TransactOpts, _interpolationParams)
}

// SetMinimumRWTBalance is a paid mutator transaction binding the contract method 0x8ffd6d2d.
//
// Solidity: function setMinimumRWTBalance(uint256 _minimumRWTBalance) returns()
func (_LPPlus *LPPlusTransactor) SetMinimumRWTBalance(opts *bind.TransactOpts, _minimumRWTBalance *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setMinimumRWTBalance", _minimumRWTBalance)
}

// SetMinimumRWTBalance is a paid mutator transaction binding the contract method 0x8ffd6d2d.
//
// Solidity: function setMinimumRWTBalance(uint256 _minimumRWTBalance) returns()
func (_LPPlus *LPPlusSession) SetMinimumRWTBalance(_minimumRWTBalance *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMinimumRWTBalance(&_LPPlus.TransactOpts, _minimumRWTBalance)
}

// SetMinimumRWTBalance is a paid mutator transaction binding the contract method 0x8ffd6d2d.
//
// Solidity: function setMinimumRWTBalance(uint256 _minimumRWTBalance) returns()
func (_LPPlus *LPPlusTransactorSession) SetMinimumRWTBalance(_minimumRWTBalance *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMinimumRWTBalance(&_LPPlus.TransactOpts, _minimumRWTBalance)
}

// SetMinimumYBTBalance is a paid mutator transaction binding the contract method 0xe080e4bd.
//
// Solidity: function setMinimumYBTBalance(uint256 _minimumYBTBalance) returns()
func (_LPPlus *LPPlusTransactor) SetMinimumYBTBalance(opts *bind.TransactOpts, _minimumYBTBalance *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setMinimumYBTBalance", _minimumYBTBalance)
}

// SetMinimumYBTBalance is a paid mutator transaction binding the contract method 0xe080e4bd.
//
// Solidity: function setMinimumYBTBalance(uint256 _minimumYBTBalance) returns()
func (_LPPlus *LPPlusSession) SetMinimumYBTBalance(_minimumYBTBalance *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMinimumYBTBalance(&_LPPlus.TransactOpts, _minimumYBTBalance)
}

// SetMinimumYBTBalance is a paid mutator transaction binding the contract method 0xe080e4bd.
//
// Solidity: function setMinimumYBTBalance(uint256 _minimumYBTBalance) returns()
func (_LPPlus *LPPlusTransactorSession) SetMinimumYBTBalance(_minimumYBTBalance *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMinimumYBTBalance(&_LPPlus.TransactOpts, _minimumYBTBalance)
}

// SetMintBonus is a paid mutator transaction binding the contract method 0x2bf2ba03.
//
// Solidity: function setMintBonus((uint256,uint256) _mintBonus) returns()
func (_LPPlus *LPPlusTransactor) SetMintBonus(opts *bind.TransactOpts, _mintBonus RWTFutureBonus) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setMintBonus", _mintBonus)
}

// SetMintBonus is a paid mutator transaction binding the contract method 0x2bf2ba03.
//
// Solidity: function setMintBonus((uint256,uint256) _mintBonus) returns()
func (_LPPlus *LPPlusSession) SetMintBonus(_mintBonus RWTFutureBonus) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMintBonus(&_LPPlus.TransactOpts, _mintBonus)
}

// SetMintBonus is a paid mutator transaction binding the contract method 0x2bf2ba03.
//
// Solidity: function setMintBonus((uint256,uint256) _mintBonus) returns()
func (_LPPlus *LPPlusTransactorSession) SetMintBonus(_mintBonus RWTFutureBonus) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMintBonus(&_LPPlus.TransactOpts, _mintBonus)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_LPPlus *LPPlusTransactor) SetMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setMintPrice", _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_LPPlus *LPPlusSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMintPrice(&_LPPlus.TransactOpts, _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_LPPlus *LPPlusTransactorSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.SetMintPrice(&_LPPlus.TransactOpts, _mintPrice)
}

// SetRWTFuture is a paid mutator transaction binding the contract method 0x3a20853e.
//
// Solidity: function setRWTFuture(address _RWTFuture) returns()
func (_LPPlus *LPPlusTransactor) SetRWTFuture(opts *bind.TransactOpts, _RWTFuture common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setRWTFuture", _RWTFuture)
}

// SetRWTFuture is a paid mutator transaction binding the contract method 0x3a20853e.
//
// Solidity: function setRWTFuture(address _RWTFuture) returns()
func (_LPPlus *LPPlusSession) SetRWTFuture(_RWTFuture common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetRWTFuture(&_LPPlus.TransactOpts, _RWTFuture)
}

// SetRWTFuture is a paid mutator transaction binding the contract method 0x3a20853e.
//
// Solidity: function setRWTFuture(address _RWTFuture) returns()
func (_LPPlus *LPPlusTransactorSession) SetRWTFuture(_RWTFuture common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetRWTFuture(&_LPPlus.TransactOpts, _RWTFuture)
}

// SetRWTToken is a paid mutator transaction binding the contract method 0x300e3c4f.
//
// Solidity: function setRWTToken(address _RWTToken) returns()
func (_LPPlus *LPPlusTransactor) SetRWTToken(opts *bind.TransactOpts, _RWTToken common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setRWTToken", _RWTToken)
}

// SetRWTToken is a paid mutator transaction binding the contract method 0x300e3c4f.
//
// Solidity: function setRWTToken(address _RWTToken) returns()
func (_LPPlus *LPPlusSession) SetRWTToken(_RWTToken common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetRWTToken(&_LPPlus.TransactOpts, _RWTToken)
}

// SetRWTToken is a paid mutator transaction binding the contract method 0x300e3c4f.
//
// Solidity: function setRWTToken(address _RWTToken) returns()
func (_LPPlus *LPPlusTransactorSession) SetRWTToken(_RWTToken common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetRWTToken(&_LPPlus.TransactOpts, _RWTToken)
}

// SetReferrerBonus is a paid mutator transaction binding the contract method 0x1418f711.
//
// Solidity: function setReferrerBonus((uint256,uint256) _referrerBonus) returns()
func (_LPPlus *LPPlusTransactor) SetReferrerBonus(opts *bind.TransactOpts, _referrerBonus RWTFutureBonus) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setReferrerBonus", _referrerBonus)
}

// SetReferrerBonus is a paid mutator transaction binding the contract method 0x1418f711.
//
// Solidity: function setReferrerBonus((uint256,uint256) _referrerBonus) returns()
func (_LPPlus *LPPlusSession) SetReferrerBonus(_referrerBonus RWTFutureBonus) (*types.Transaction, error) {
	return _LPPlus.Contract.SetReferrerBonus(&_LPPlus.TransactOpts, _referrerBonus)
}

// SetReferrerBonus is a paid mutator transaction binding the contract method 0x1418f711.
//
// Solidity: function setReferrerBonus((uint256,uint256) _referrerBonus) returns()
func (_LPPlus *LPPlusTransactorSession) SetReferrerBonus(_referrerBonus RWTFutureBonus) (*types.Transaction, error) {
	return _LPPlus.Contract.SetReferrerBonus(&_LPPlus.TransactOpts, _referrerBonus)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_LPPlus *LPPlusTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_LPPlus *LPPlusSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetTreasury(&_LPPlus.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_LPPlus *LPPlusTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetTreasury(&_LPPlus.TransactOpts, _treasury)
}

// SetWindowCommitter is a paid mutator transaction binding the contract method 0x548d9356.
//
// Solidity: function setWindowCommitter(address _windowCommitter) returns()
func (_LPPlus *LPPlusTransactor) SetWindowCommitter(opts *bind.TransactOpts, _windowCommitter common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setWindowCommitter", _windowCommitter)
}

// SetWindowCommitter is a paid mutator transaction binding the contract method 0x548d9356.
//
// Solidity: function setWindowCommitter(address _windowCommitter) returns()
func (_LPPlus *LPPlusSession) SetWindowCommitter(_windowCommitter common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetWindowCommitter(&_LPPlus.TransactOpts, _windowCommitter)
}

// SetWindowCommitter is a paid mutator transaction binding the contract method 0x548d9356.
//
// Solidity: function setWindowCommitter(address _windowCommitter) returns()
func (_LPPlus *LPPlusTransactorSession) SetWindowCommitter(_windowCommitter common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetWindowCommitter(&_LPPlus.TransactOpts, _windowCommitter)
}

// SetYBTToken is a paid mutator transaction binding the contract method 0x5e29e06b.
//
// Solidity: function setYBTToken(address _YBTToken) returns()
func (_LPPlus *LPPlusTransactor) SetYBTToken(opts *bind.TransactOpts, _YBTToken common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "setYBTToken", _YBTToken)
}

// SetYBTToken is a paid mutator transaction binding the contract method 0x5e29e06b.
//
// Solidity: function setYBTToken(address _YBTToken) returns()
func (_LPPlus *LPPlusSession) SetYBTToken(_YBTToken common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetYBTToken(&_LPPlus.TransactOpts, _YBTToken)
}

// SetYBTToken is a paid mutator transaction binding the contract method 0x5e29e06b.
//
// Solidity: function setYBTToken(address _YBTToken) returns()
func (_LPPlus *LPPlusTransactorSession) SetYBTToken(_YBTToken common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.SetYBTToken(&_LPPlus.TransactOpts, _YBTToken)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPPlus *LPPlusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPPlus *LPPlusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.TransferOwnership(&_LPPlus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPPlus *LPPlusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LPPlus.Contract.TransferOwnership(&_LPPlus.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LPPlus *LPPlusTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LPPlus *LPPlusSession) Unpause() (*types.Transaction, error) {
	return _LPPlus.Contract.Unpause(&_LPPlus.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LPPlus *LPPlusTransactorSession) Unpause() (*types.Transaction, error) {
	return _LPPlus.Contract.Unpause(&_LPPlus.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LPPlus *LPPlusTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LPPlus *LPPlusSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LPPlus.Contract.UpgradeToAndCall(&_LPPlus.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LPPlus *LPPlusTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LPPlus.Contract.UpgradeToAndCall(&_LPPlus.TransactOpts, newImplementation, data)
}

// WithdrawFilFunds is a paid mutator transaction binding the contract method 0x2eacc696.
//
// Solidity: function withdrawFilFunds(address _to, uint256 _amount) returns()
func (_LPPlus *LPPlusTransactor) WithdrawFilFunds(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "withdrawFilFunds", _to, _amount)
}

// WithdrawFilFunds is a paid mutator transaction binding the contract method 0x2eacc696.
//
// Solidity: function withdrawFilFunds(address _to, uint256 _amount) returns()
func (_LPPlus *LPPlusSession) WithdrawFilFunds(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.WithdrawFilFunds(&_LPPlus.TransactOpts, _to, _amount)
}

// WithdrawFilFunds is a paid mutator transaction binding the contract method 0x2eacc696.
//
// Solidity: function withdrawFilFunds(address _to, uint256 _amount) returns()
func (_LPPlus *LPPlusTransactorSession) WithdrawFilFunds(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LPPlus.Contract.WithdrawFilFunds(&_LPPlus.TransactOpts, _to, _amount)
}

// LPPlusApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LPPlus contract.
type LPPlusApprovalIterator struct {
	Event *LPPlusApproval // Event containing the contract specifics and raw log

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
func (it *LPPlusApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusApproval)
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
		it.Event = new(LPPlusApproval)
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
func (it *LPPlusApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusApproval represents a Approval event raised by the LPPlus contract.
type LPPlusApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LPPlus *LPPlusFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LPPlusApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusApprovalIterator{contract: _LPPlus.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LPPlus *LPPlusFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LPPlusApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusApproval)
				if err := _LPPlus.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LPPlus *LPPlusFilterer) ParseApproval(log types.Log) (*LPPlusApproval, error) {
	event := new(LPPlusApproval)
	if err := _LPPlus.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LPPlus contract.
type LPPlusApprovalForAllIterator struct {
	Event *LPPlusApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LPPlusApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusApprovalForAll)
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
		it.Event = new(LPPlusApprovalForAll)
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
func (it *LPPlusApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusApprovalForAll represents a ApprovalForAll event raised by the LPPlus contract.
type LPPlusApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LPPlus *LPPlusFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LPPlusApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusApprovalForAllIterator{contract: _LPPlus.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LPPlus *LPPlusFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LPPlusApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusApprovalForAll)
				if err := _LPPlus.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LPPlus *LPPlusFilterer) ParseApprovalForAll(log types.Log) (*LPPlusApprovalForAll, error) {
	event := new(LPPlusApprovalForAll)
	if err := _LPPlus.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusBonusAppliedIterator is returned from FilterBonusApplied and is used to iterate over the raw logs and unpacked data for BonusApplied events raised by the LPPlus contract.
type LPPlusBonusAppliedIterator struct {
	Event *LPPlusBonusApplied // Event containing the contract specifics and raw log

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
func (it *LPPlusBonusAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusBonusApplied)
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
		it.Event = new(LPPlusBonusApplied)
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
func (it *LPPlusBonusAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusBonusAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusBonusApplied represents a BonusApplied event raised by the LPPlus contract.
type LPPlusBonusApplied struct {
	TokenId *big.Int
	Bonus   RWTFutureBonus
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBonusApplied is a free log retrieval operation binding the contract event 0x08c20988b3a0d46fafe47abf0aa2ac3716335a2190a6acb020f12468259207ac.
//
// Solidity: event BonusApplied(uint256 indexed tokenId, (uint256,uint256) bonus)
func (_LPPlus *LPPlusFilterer) FilterBonusApplied(opts *bind.FilterOpts, tokenId []*big.Int) (*LPPlusBonusAppliedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "BonusApplied", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusBonusAppliedIterator{contract: _LPPlus.contract, event: "BonusApplied", logs: logs, sub: sub}, nil
}

// WatchBonusApplied is a free log subscription operation binding the contract event 0x08c20988b3a0d46fafe47abf0aa2ac3716335a2190a6acb020f12468259207ac.
//
// Solidity: event BonusApplied(uint256 indexed tokenId, (uint256,uint256) bonus)
func (_LPPlus *LPPlusFilterer) WatchBonusApplied(opts *bind.WatchOpts, sink chan<- *LPPlusBonusApplied, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "BonusApplied", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusBonusApplied)
				if err := _LPPlus.contract.UnpackLog(event, "BonusApplied", log); err != nil {
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

// ParseBonusApplied is a log parse operation binding the contract event 0x08c20988b3a0d46fafe47abf0aa2ac3716335a2190a6acb020f12468259207ac.
//
// Solidity: event BonusApplied(uint256 indexed tokenId, (uint256,uint256) bonus)
func (_LPPlus *LPPlusFilterer) ParseBonusApplied(log types.Log) (*LPPlusBonusApplied, error) {
	event := new(LPPlusBonusApplied)
	if err := _LPPlus.contract.UnpackLog(event, "BonusApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusCardMintedIterator is returned from FilterCardMinted and is used to iterate over the raw logs and unpacked data for CardMinted events raised by the LPPlus contract.
type LPPlusCardMintedIterator struct {
	Event *LPPlusCardMinted // Event containing the contract specifics and raw log

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
func (it *LPPlusCardMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusCardMinted)
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
		it.Event = new(LPPlusCardMinted)
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
func (it *LPPlusCardMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusCardMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusCardMinted represents a CardMinted event raised by the LPPlus contract.
type LPPlusCardMinted struct {
	TokenId         *big.Int
	Minter          common.Address
	Receiver        common.Address
	ReferrerTokenId *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCardMinted is a free log retrieval operation binding the contract event 0x98d1d1eb3af2e9f5c522e0a03cec7374e965eb2a56911d0010ec933ede85affd.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address indexed minter, address indexed receiver, uint256 referrerTokenId)
func (_LPPlus *LPPlusFilterer) FilterCardMinted(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address, receiver []common.Address) (*LPPlusCardMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "CardMinted", tokenIdRule, minterRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusCardMintedIterator{contract: _LPPlus.contract, event: "CardMinted", logs: logs, sub: sub}, nil
}

// WatchCardMinted is a free log subscription operation binding the contract event 0x98d1d1eb3af2e9f5c522e0a03cec7374e965eb2a56911d0010ec933ede85affd.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address indexed minter, address indexed receiver, uint256 referrerTokenId)
func (_LPPlus *LPPlusFilterer) WatchCardMinted(opts *bind.WatchOpts, sink chan<- *LPPlusCardMinted, tokenId []*big.Int, minter []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "CardMinted", tokenIdRule, minterRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusCardMinted)
				if err := _LPPlus.contract.UnpackLog(event, "CardMinted", log); err != nil {
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

// ParseCardMinted is a log parse operation binding the contract event 0x98d1d1eb3af2e9f5c522e0a03cec7374e965eb2a56911d0010ec933ede85affd.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address indexed minter, address indexed receiver, uint256 referrerTokenId)
func (_LPPlus *LPPlusFilterer) ParseCardMinted(log types.Log) (*LPPlusCardMinted, error) {
	event := new(LPPlusCardMinted)
	if err := _LPPlus.contract.UnpackLog(event, "CardMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusFilVaultFundedIterator is returned from FilterFilVaultFunded and is used to iterate over the raw logs and unpacked data for FilVaultFunded events raised by the LPPlus contract.
type LPPlusFilVaultFundedIterator struct {
	Event *LPPlusFilVaultFunded // Event containing the contract specifics and raw log

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
func (it *LPPlusFilVaultFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusFilVaultFunded)
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
		it.Event = new(LPPlusFilVaultFunded)
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
func (it *LPPlusFilVaultFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusFilVaultFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusFilVaultFunded represents a FilVaultFunded event raised by the LPPlus contract.
type LPPlusFilVaultFunded struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFilVaultFunded is a free log retrieval operation binding the contract event 0x94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a.
//
// Solidity: event FilVaultFunded(address indexed owner, uint256 amount)
func (_LPPlus *LPPlusFilterer) FilterFilVaultFunded(opts *bind.FilterOpts, owner []common.Address) (*LPPlusFilVaultFundedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "FilVaultFunded", ownerRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusFilVaultFundedIterator{contract: _LPPlus.contract, event: "FilVaultFunded", logs: logs, sub: sub}, nil
}

// WatchFilVaultFunded is a free log subscription operation binding the contract event 0x94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a.
//
// Solidity: event FilVaultFunded(address indexed owner, uint256 amount)
func (_LPPlus *LPPlusFilterer) WatchFilVaultFunded(opts *bind.WatchOpts, sink chan<- *LPPlusFilVaultFunded, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "FilVaultFunded", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusFilVaultFunded)
				if err := _LPPlus.contract.UnpackLog(event, "FilVaultFunded", log); err != nil {
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

// ParseFilVaultFunded is a log parse operation binding the contract event 0x94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a.
//
// Solidity: event FilVaultFunded(address indexed owner, uint256 amount)
func (_LPPlus *LPPlusFilterer) ParseFilVaultFunded(log types.Log) (*LPPlusFilVaultFunded, error) {
	event := new(LPPlusFilVaultFunded)
	if err := _LPPlus.contract.UnpackLog(event, "FilVaultFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LPPlus contract.
type LPPlusInitializedIterator struct {
	Event *LPPlusInitialized // Event containing the contract specifics and raw log

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
func (it *LPPlusInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusInitialized)
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
		it.Event = new(LPPlusInitialized)
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
func (it *LPPlusInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusInitialized represents a Initialized event raised by the LPPlus contract.
type LPPlusInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LPPlus *LPPlusFilterer) FilterInitialized(opts *bind.FilterOpts) (*LPPlusInitializedIterator, error) {

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LPPlusInitializedIterator{contract: _LPPlus.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LPPlus *LPPlusFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LPPlusInitialized) (event.Subscription, error) {

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusInitialized)
				if err := _LPPlus.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LPPlus *LPPlusFilterer) ParseInitialized(log types.Log) (*LPPlusInitialized, error) {
	event := new(LPPlusInitialized)
	if err := _LPPlus.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the LPPlus contract.
type LPPlusOwnershipTransferStartedIterator struct {
	Event *LPPlusOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *LPPlusOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusOwnershipTransferStarted)
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
		it.Event = new(LPPlusOwnershipTransferStarted)
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
func (it *LPPlusOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the LPPlus contract.
type LPPlusOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_LPPlus *LPPlusFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LPPlusOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusOwnershipTransferStartedIterator{contract: _LPPlus.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_LPPlus *LPPlusFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *LPPlusOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusOwnershipTransferStarted)
				if err := _LPPlus.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_LPPlus *LPPlusFilterer) ParseOwnershipTransferStarted(log types.Log) (*LPPlusOwnershipTransferStarted, error) {
	event := new(LPPlusOwnershipTransferStarted)
	if err := _LPPlus.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LPPlus contract.
type LPPlusOwnershipTransferredIterator struct {
	Event *LPPlusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LPPlusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusOwnershipTransferred)
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
		it.Event = new(LPPlusOwnershipTransferred)
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
func (it *LPPlusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusOwnershipTransferred represents a OwnershipTransferred event raised by the LPPlus contract.
type LPPlusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPPlus *LPPlusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LPPlusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusOwnershipTransferredIterator{contract: _LPPlus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPPlus *LPPlusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LPPlusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusOwnershipTransferred)
				if err := _LPPlus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPPlus *LPPlusFilterer) ParseOwnershipTransferred(log types.Log) (*LPPlusOwnershipTransferred, error) {
	event := new(LPPlusOwnershipTransferred)
	if err := _LPPlus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LPPlus contract.
type LPPlusPausedIterator struct {
	Event *LPPlusPaused // Event containing the contract specifics and raw log

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
func (it *LPPlusPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusPaused)
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
		it.Event = new(LPPlusPaused)
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
func (it *LPPlusPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusPaused represents a Paused event raised by the LPPlus contract.
type LPPlusPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LPPlus *LPPlusFilterer) FilterPaused(opts *bind.FilterOpts) (*LPPlusPausedIterator, error) {

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LPPlusPausedIterator{contract: _LPPlus.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LPPlus *LPPlusFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LPPlusPaused) (event.Subscription, error) {

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusPaused)
				if err := _LPPlus.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LPPlus *LPPlusFilterer) ParsePaused(log types.Log) (*LPPlusPaused, error) {
	event := new(LPPlusPaused)
	if err := _LPPlus.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusRWTDepositedIterator is returned from FilterRWTDeposited and is used to iterate over the raw logs and unpacked data for RWTDeposited events raised by the LPPlus contract.
type LPPlusRWTDepositedIterator struct {
	Event *LPPlusRWTDeposited // Event containing the contract specifics and raw log

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
func (it *LPPlusRWTDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusRWTDeposited)
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
		it.Event = new(LPPlusRWTDeposited)
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
func (it *LPPlusRWTDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusRWTDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusRWTDeposited represents a RWTDeposited event raised by the LPPlus contract.
type LPPlusRWTDeposited struct {
	TokenId       *big.Int
	Owner         common.Address
	Amount        *big.Int
	UserTotal     *big.Int
	ContractTotal *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRWTDeposited is a free log retrieval operation binding the contract event 0x8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a984.
//
// Solidity: event RWTDeposited(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) FilterRWTDeposited(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*LPPlusRWTDepositedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "RWTDeposited", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusRWTDepositedIterator{contract: _LPPlus.contract, event: "RWTDeposited", logs: logs, sub: sub}, nil
}

// WatchRWTDeposited is a free log subscription operation binding the contract event 0x8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a984.
//
// Solidity: event RWTDeposited(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) WatchRWTDeposited(opts *bind.WatchOpts, sink chan<- *LPPlusRWTDeposited, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "RWTDeposited", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusRWTDeposited)
				if err := _LPPlus.contract.UnpackLog(event, "RWTDeposited", log); err != nil {
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

// ParseRWTDeposited is a log parse operation binding the contract event 0x8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a984.
//
// Solidity: event RWTDeposited(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) ParseRWTDeposited(log types.Log) (*LPPlusRWTDeposited, error) {
	event := new(LPPlusRWTDeposited)
	if err := _LPPlus.contract.UnpackLog(event, "RWTDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusRWTExchangedForFILIterator is returned from FilterRWTExchangedForFIL and is used to iterate over the raw logs and unpacked data for RWTExchangedForFIL events raised by the LPPlus contract.
type LPPlusRWTExchangedForFILIterator struct {
	Event *LPPlusRWTExchangedForFIL // Event containing the contract specifics and raw log

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
func (it *LPPlusRWTExchangedForFILIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusRWTExchangedForFIL)
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
		it.Event = new(LPPlusRWTExchangedForFIL)
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
func (it *LPPlusRWTExchangedForFILIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusRWTExchangedForFILIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusRWTExchangedForFIL represents a RWTExchangedForFIL event raised by the LPPlus contract.
type LPPlusRWTExchangedForFIL struct {
	Sender           common.Address
	Receiver         common.Address
	RwtFutureTokenId *big.Int
	RwtAmount        *big.Int
	FilAmount        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRWTExchangedForFIL is a free log retrieval operation binding the contract event 0x4dfd983654825b97cb5a6edb928f7e2a2726bef374ec4be8b39e4c31660103bb.
//
// Solidity: event RWTExchangedForFIL(address indexed sender, address indexed receiver, uint256 indexed rwtFutureTokenId, uint256 rwtAmount, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) FilterRWTExchangedForFIL(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, rwtFutureTokenId []*big.Int) (*LPPlusRWTExchangedForFILIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var rwtFutureTokenIdRule []interface{}
	for _, rwtFutureTokenIdItem := range rwtFutureTokenId {
		rwtFutureTokenIdRule = append(rwtFutureTokenIdRule, rwtFutureTokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "RWTExchangedForFIL", senderRule, receiverRule, rwtFutureTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusRWTExchangedForFILIterator{contract: _LPPlus.contract, event: "RWTExchangedForFIL", logs: logs, sub: sub}, nil
}

// WatchRWTExchangedForFIL is a free log subscription operation binding the contract event 0x4dfd983654825b97cb5a6edb928f7e2a2726bef374ec4be8b39e4c31660103bb.
//
// Solidity: event RWTExchangedForFIL(address indexed sender, address indexed receiver, uint256 indexed rwtFutureTokenId, uint256 rwtAmount, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) WatchRWTExchangedForFIL(opts *bind.WatchOpts, sink chan<- *LPPlusRWTExchangedForFIL, sender []common.Address, receiver []common.Address, rwtFutureTokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var rwtFutureTokenIdRule []interface{}
	for _, rwtFutureTokenIdItem := range rwtFutureTokenId {
		rwtFutureTokenIdRule = append(rwtFutureTokenIdRule, rwtFutureTokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "RWTExchangedForFIL", senderRule, receiverRule, rwtFutureTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusRWTExchangedForFIL)
				if err := _LPPlus.contract.UnpackLog(event, "RWTExchangedForFIL", log); err != nil {
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

// ParseRWTExchangedForFIL is a log parse operation binding the contract event 0x4dfd983654825b97cb5a6edb928f7e2a2726bef374ec4be8b39e4c31660103bb.
//
// Solidity: event RWTExchangedForFIL(address indexed sender, address indexed receiver, uint256 indexed rwtFutureTokenId, uint256 rwtAmount, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) ParseRWTExchangedForFIL(log types.Log) (*LPPlusRWTExchangedForFIL, error) {
	event := new(LPPlusRWTExchangedForFIL)
	if err := _LPPlus.contract.UnpackLog(event, "RWTExchangedForFIL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusRWTWithdrawnIterator is returned from FilterRWTWithdrawn and is used to iterate over the raw logs and unpacked data for RWTWithdrawn events raised by the LPPlus contract.
type LPPlusRWTWithdrawnIterator struct {
	Event *LPPlusRWTWithdrawn // Event containing the contract specifics and raw log

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
func (it *LPPlusRWTWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusRWTWithdrawn)
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
		it.Event = new(LPPlusRWTWithdrawn)
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
func (it *LPPlusRWTWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusRWTWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusRWTWithdrawn represents a RWTWithdrawn event raised by the LPPlus contract.
type LPPlusRWTWithdrawn struct {
	TokenId       *big.Int
	Owner         common.Address
	Receiver      common.Address
	Amount        *big.Int
	UserTotal     *big.Int
	ContractTotal *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRWTWithdrawn is a free log retrieval operation binding the contract event 0x1e9c12c534d62150c4f5b916807e2bd6619095393fda9c833471a1feeefad693.
//
// Solidity: event RWTWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) FilterRWTWithdrawn(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address, receiver []common.Address) (*LPPlusRWTWithdrawnIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "RWTWithdrawn", tokenIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusRWTWithdrawnIterator{contract: _LPPlus.contract, event: "RWTWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRWTWithdrawn is a free log subscription operation binding the contract event 0x1e9c12c534d62150c4f5b916807e2bd6619095393fda9c833471a1feeefad693.
//
// Solidity: event RWTWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) WatchRWTWithdrawn(opts *bind.WatchOpts, sink chan<- *LPPlusRWTWithdrawn, tokenId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "RWTWithdrawn", tokenIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusRWTWithdrawn)
				if err := _LPPlus.contract.UnpackLog(event, "RWTWithdrawn", log); err != nil {
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

// ParseRWTWithdrawn is a log parse operation binding the contract event 0x1e9c12c534d62150c4f5b916807e2bd6619095393fda9c833471a1feeefad693.
//
// Solidity: event RWTWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) ParseRWTWithdrawn(log types.Log) (*LPPlusRWTWithdrawn, error) {
	event := new(LPPlusRWTWithdrawn)
	if err := _LPPlus.contract.UnpackLog(event, "RWTWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LPPlus contract.
type LPPlusTransferIterator struct {
	Event *LPPlusTransfer // Event containing the contract specifics and raw log

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
func (it *LPPlusTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusTransfer)
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
		it.Event = new(LPPlusTransfer)
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
func (it *LPPlusTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusTransfer represents a Transfer event raised by the LPPlus contract.
type LPPlusTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LPPlus *LPPlusFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LPPlusTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusTransferIterator{contract: _LPPlus.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LPPlus *LPPlusFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LPPlusTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusTransfer)
				if err := _LPPlus.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LPPlus *LPPlusFilterer) ParseTransfer(log types.Log) (*LPPlusTransfer, error) {
	event := new(LPPlusTransfer)
	if err := _LPPlus.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LPPlus contract.
type LPPlusUnpausedIterator struct {
	Event *LPPlusUnpaused // Event containing the contract specifics and raw log

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
func (it *LPPlusUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusUnpaused)
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
		it.Event = new(LPPlusUnpaused)
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
func (it *LPPlusUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusUnpaused represents a Unpaused event raised by the LPPlus contract.
type LPPlusUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LPPlus *LPPlusFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LPPlusUnpausedIterator, error) {

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LPPlusUnpausedIterator{contract: _LPPlus.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LPPlus *LPPlusFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LPPlusUnpaused) (event.Subscription, error) {

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusUnpaused)
				if err := _LPPlus.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LPPlus *LPPlusFilterer) ParseUnpaused(log types.Log) (*LPPlusUnpaused, error) {
	event := new(LPPlusUnpaused)
	if err := _LPPlus.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LPPlus contract.
type LPPlusUpgradedIterator struct {
	Event *LPPlusUpgraded // Event containing the contract specifics and raw log

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
func (it *LPPlusUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusUpgraded)
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
		it.Event = new(LPPlusUpgraded)
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
func (it *LPPlusUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusUpgraded represents a Upgraded event raised by the LPPlus contract.
type LPPlusUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LPPlus *LPPlusFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LPPlusUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusUpgradedIterator{contract: _LPPlus.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LPPlus *LPPlusFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LPPlusUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusUpgraded)
				if err := _LPPlus.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LPPlus *LPPlusFilterer) ParseUpgraded(log types.Log) (*LPPlusUpgraded, error) {
	event := new(LPPlusUpgraded)
	if err := _LPPlus.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusYBTVaultFundedIterator is returned from FilterYBTVaultFunded and is used to iterate over the raw logs and unpacked data for YBTVaultFunded events raised by the LPPlus contract.
type LPPlusYBTVaultFundedIterator struct {
	Event *LPPlusYBTVaultFunded // Event containing the contract specifics and raw log

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
func (it *LPPlusYBTVaultFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusYBTVaultFunded)
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
		it.Event = new(LPPlusYBTVaultFunded)
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
func (it *LPPlusYBTVaultFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusYBTVaultFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusYBTVaultFunded represents a YBTVaultFunded event raised by the LPPlus contract.
type LPPlusYBTVaultFunded struct {
	TokenId       *big.Int
	Owner         common.Address
	Amount        *big.Int
	UserTotal     *big.Int
	ContractTotal *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterYBTVaultFunded is a free log retrieval operation binding the contract event 0xade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb4793790.
//
// Solidity: event YBTVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) FilterYBTVaultFunded(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*LPPlusYBTVaultFundedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "YBTVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusYBTVaultFundedIterator{contract: _LPPlus.contract, event: "YBTVaultFunded", logs: logs, sub: sub}, nil
}

// WatchYBTVaultFunded is a free log subscription operation binding the contract event 0xade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb4793790.
//
// Solidity: event YBTVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) WatchYBTVaultFunded(opts *bind.WatchOpts, sink chan<- *LPPlusYBTVaultFunded, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "YBTVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusYBTVaultFunded)
				if err := _LPPlus.contract.UnpackLog(event, "YBTVaultFunded", log); err != nil {
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

// ParseYBTVaultFunded is a log parse operation binding the contract event 0xade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb4793790.
//
// Solidity: event YBTVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) ParseYBTVaultFunded(log types.Log) (*LPPlusYBTVaultFunded, error) {
	event := new(LPPlusYBTVaultFunded)
	if err := _LPPlus.contract.UnpackLog(event, "YBTVaultFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusYBTVaultWithdrawnIterator is returned from FilterYBTVaultWithdrawn and is used to iterate over the raw logs and unpacked data for YBTVaultWithdrawn events raised by the LPPlus contract.
type LPPlusYBTVaultWithdrawnIterator struct {
	Event *LPPlusYBTVaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *LPPlusYBTVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusYBTVaultWithdrawn)
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
		it.Event = new(LPPlusYBTVaultWithdrawn)
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
func (it *LPPlusYBTVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusYBTVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusYBTVaultWithdrawn represents a YBTVaultWithdrawn event raised by the LPPlus contract.
type LPPlusYBTVaultWithdrawn struct {
	TokenId       *big.Int
	Owner         common.Address
	Receiver      common.Address
	Amount        *big.Int
	UserTotal     *big.Int
	ContractTotal *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterYBTVaultWithdrawn is a free log retrieval operation binding the contract event 0xc4032b9f060fe681645ab7a2b55a30a15304792f8669533901de6fb6e0dea112.
//
// Solidity: event YBTVaultWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) FilterYBTVaultWithdrawn(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address, receiver []common.Address) (*LPPlusYBTVaultWithdrawnIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "YBTVaultWithdrawn", tokenIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusYBTVaultWithdrawnIterator{contract: _LPPlus.contract, event: "YBTVaultWithdrawn", logs: logs, sub: sub}, nil
}

// WatchYBTVaultWithdrawn is a free log subscription operation binding the contract event 0xc4032b9f060fe681645ab7a2b55a30a15304792f8669533901de6fb6e0dea112.
//
// Solidity: event YBTVaultWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) WatchYBTVaultWithdrawn(opts *bind.WatchOpts, sink chan<- *LPPlusYBTVaultWithdrawn, tokenId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "YBTVaultWithdrawn", tokenIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusYBTVaultWithdrawn)
				if err := _LPPlus.contract.UnpackLog(event, "YBTVaultWithdrawn", log); err != nil {
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

// ParseYBTVaultWithdrawn is a log parse operation binding the contract event 0xc4032b9f060fe681645ab7a2b55a30a15304792f8669533901de6fb6e0dea112.
//
// Solidity: event YBTVaultWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount, uint256 userTotal, uint256 contractTotal)
func (_LPPlus *LPPlusFilterer) ParseYBTVaultWithdrawn(log types.Log) (*LPPlusYBTVaultWithdrawn, error) {
	event := new(LPPlusYBTVaultWithdrawn)
	if err := _LPPlus.contract.UnpackLog(event, "YBTVaultWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
