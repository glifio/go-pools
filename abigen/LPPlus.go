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
	MinimumYBTBalance           *big.Int
	MinimumRWTBalance           *big.Int
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"}],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"CardAlreadyMinted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"FilReservesDepleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientRWTBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientVaultBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_second\",\"type\":\"uint256\"}],\"name\":\"InvalidLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"name\":\"RWTFutureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenNotActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"providedWindowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentWindowId\",\"type\":\"uint256\"}],\"name\":\"WindowIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroCard\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroConversionRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"}],\"name\":\"ZeroMerkleRoot\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rwtFutureTokenId\",\"type\":\"uint256\"}],\"name\":\"BonusApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referrerTokenId\",\"type\":\"uint256\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FilVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"referrerTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rwtFutureTokenId\",\"type\":\"uint256\"}],\"name\":\"NewReferral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"RWTDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"senderTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rwtFutureTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rwtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filAmount\",\"type\":\"uint256\"}],\"name\":\"RWTExchangedForFIL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filAmount\",\"type\":\"uint256\"}],\"name\":\"RWTFutureClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"RWTWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"}],\"name\":\"WindowCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"YBTVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"YBTVaultWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_RWTAmount\",\"type\":\"uint256\"}],\"name\":\"addRWT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_YBTAmount\",\"type\":\"uint256\"}],\"name\":\"addYBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_windowIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"}],\"internalType\":\"structLpPlusMerkleHelper.UserSnapshot[]\",\"name\":\"_userSnapshots\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_proofs\",\"type\":\"bytes32[][]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_windowIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"}],\"internalType\":\"structLpPlusMerkleHelper.UserSnapshot[]\",\"name\":\"_userSnapshots\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_proofs\",\"type\":\"bytes32[][]\"}],\"name\":\"claimAndExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumYBTBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumRWTBalance\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"_stakingSnapshot\",\"type\":\"tuple\"}],\"name\":\"commitWindow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_rwtFutureTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundFilVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureValidityDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getUserActionsBatched\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ybtBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rwtBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structUserAction[][]\",\"name\":\"actionsPerToken\",\"type\":\"tuple[][]\"},{\"internalType\":\"uint256\",\"name\":\"newStartTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newOffset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rwtToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_ybtToken\",\"type\":\"address\"},{\"internalType\":\"contractIRWTFuture\",\"name\":\"_rwtFuture\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_windowCommitter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interpolationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumRWTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumYBTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_referrerTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ybtToStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rwtToStake\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBonus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingTotalYbt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalFilToAllocate\",\"type\":\"uint256\"}],\"name\":\"previewStakingSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumYBTBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumRWTBalance\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referrerBonus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_RWTAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeRWT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_YBTAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeYBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwtFuture\",\"outputs\":[{\"internalType\":\"contractIRWTFuture\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_futureValidityDuration\",\"type\":\"uint256\"}],\"name\":\"setFutureValidityDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"_interpolationParams\",\"type\":\"tuple\"}],\"name\":\"setInterpolationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumRWTBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumRWTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumYBTBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumYBTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"_mintBonus\",\"type\":\"tuple\"}],\"name\":\"setMintBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRWTFuture\",\"name\":\"_RWTFuture\",\"type\":\"address\"}],\"name\":\"setRWTFuture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_RWTToken\",\"type\":\"address\"}],\"name\":\"setRWTToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"_referrerBonus\",\"type\":\"tuple\"}],\"name\":\"setReferrerBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_windowCommitter\",\"type\":\"address\"}],\"name\":\"setWindowCommitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_YBTToken\",\"type\":\"address\"}],\"name\":\"setYBTToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToRWTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToWindowIdToClaimedFlag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToYBTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRwtStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalYbtStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"windowCommitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"windowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"windowIdToMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"windowIdToStakingSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumYBTBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumRWTBalance\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFilFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ybtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f516020617ebd5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051617df690816100c7823960805181818161459201526149bb0152f35b6001600160401b0319166001600160401b039081175f516020617ebd5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a714615b745750806306fdde0314615a30578063081812fc146159a9578063095ea7b3146159495780630a6b5699146158515780631418f711146157ef57806318160ddd1461579557806323b872dd1461575c57806329f2014d146156f55780632bf2ba03146156935780632e58c0c7146156395780632eacc696146155c75780632f745c59146154f5578063300e3c4f1461541b5780633263edee146153c1578063344bb6b1146152f0578063385b5f85146150ac5780633a20853e146150305780633f4ba83a14614f3757806342842e0e14614eee57806346b3241e14614c935780634bae2ef114614c515780634f1ef286146149525780634f6ccce7146148ae578063510402cd1461484f578063519473611461460a57806352d1902d1461454d578063548d93561461447457806354fd4d501461443b57806358f28a98146143e15780635c975abb146143825780635e29e06b1461427057806361d027b3146142005780636352211e146141a65780636817c76c1461414c578063694072a614613daf57806370a0823114613d6757806370af96aa14613d2c578063715018a614613bea57806379ba509714613b48578063812bf90a14613ad75780638456cb5914613a025780638da5cb5b146139915780638edd7d4e146139405780638f68679f146138e55780638ffd6d2d1461388557806391293473146137f657806395d89b41146136995780639af1d03e1461363e5780639b0b92ef146135f7578063a22cb4651461358b578063a647e8ec14612635578063a818b1ee146125da578063ac9650d81461238d578063ad3cb1cc1461230e578063ae7e2d97146122a7578063b05dcd5a14612240578063b5ea015c146121e5578063b88d4fde1461215b578063bfb1909b14611e62578063c74c23d514611e02578063c87b56dd14611d8d578063cf7cdbb814611abf578063d262f444146119c5578063d452154814610b12578063d5d2ae7214610aa1578063da357c0f14610a05578063e080e4bd146109a5578063e0ea4a1b14610936578063e30c3978146108c5578063e985e9c514610814578063ee2df9fd146107a3578063f0f4426014610696578063f2fde38b14610571578063f4a0a52814610511578063f721dc76146104a05763fcd0907d1461036d575f80fd5b3461049d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d576040516103a881615ddb565b600435815260208101602435815260408201906044358252606083019260643584526103d2616998565b8151845111156104755782518151111561047557517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1055517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1155517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1255517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d135580f35b6004857f1f3b85d3000000000000000000000000000000000000000000000000000000008152fd5b80fd5b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416604051908152f35b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57610549616998565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d005580f35b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5773ffffffffffffffffffffffffffffffffffffffff6105c66105c1615cd7565b616a04565b6105ce616998565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d576106d96106d1615cd7565b6105c1616998565b73ffffffffffffffffffffffffffffffffffffffff81161561077b576107789073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0555565b80f35b6004827fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416604051908152f35b503461049d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5761084c615cd7565b73ffffffffffffffffffffffffffffffffffffffff610869615cfa565b91165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5761096d616577565b5060806109786165de565b6109a36040518092606080918051845260208101516020850152604081015160408501520151910152565bf35b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d576109dd616998565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085580f35b503461049d57610a26610a17366160c0565b969095916105c1959395616b20565b9673ffffffffffffffffffffffffffffffffffffffff881615610a7957610a75610a6189610a588a8a8a8a8a8a61739d565b92919091617a2e565b604051918291602083526020830190616199565b0390f35b807fd92e233d0000000000000000000000000000000000000000000000000000000060049252fd5b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416604051908152f35b503461049d576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760043567ffffffffffffffff81116119c157610b6390369060040161604b565b60243567ffffffffffffffff81116119bd57610b8390369060040161604b565b60443573ffffffffffffffffffffffffffffffffffffffff811681036119b95760643560843573ffffffffffffffffffffffffffffffffffffffff81168091036119b95760a4359073ffffffffffffffffffffffffffffffffffffffff82168092036119b95760c4359273ffffffffffffffffffffffffffffffffffffffff84168094036119b95760e4359473ffffffffffffffffffffffffffffffffffffffff86168087036119b957610104359773ffffffffffffffffffffffffffffffffffffffff891692838a036119b9577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549a60ff8c60401c16159b8c67ffffffffffffffff8216801591826119b1575b5060011490816119a7575b15908161199e575b50611976578c60017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055611921575b50610d01617d69565b610d09617d69565b610d11617d69565b80519067ffffffffffffffff82116118f4578d8291610d507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300546161cc565b601f8111611821575b50602091601f8411600114611726579261171b575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b815167ffffffffffffffff81116116ee57610dfb7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301546161cc565b8d601f821161166b575b90505060208d601f83116001146115875790610e7b948361157c575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b6105c1617d69565b610e83617d69565b610e8b617d69565b73ffffffffffffffffffffffffffffffffffffffff81161561155057610eb0906177fc565b610eb8617d69565b610ec0617d69565b821561152857831561150057841561150057851561150057156114d857156114b057936105c173ffffffffffffffffffffffffffffffffffffffff6110456111c398966105c183610fc56105c1996105c183610f456111449f9c6110c59d7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0055616a04565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0155565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0255565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0355565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0555565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0655565b60017f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d045562278d007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d075567016345785d8a00007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085567016345785d8a00007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095581602060405161127281615dbf565b8281520152817f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c55817f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d556212750060206040516112cf81615dbf565b674563918244f4000081520152674563918244f400007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e55621275007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f556703782dace9d90000606060405161134481615ddb565b670de0b6b3a76400008152662386f26fc100006020820152662386f26fc1000060408201520152670de0b6b3a76400007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1055662386f26fc100007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1155662386f26fc100007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d12556703782dace9d900007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d135561141c5780f35b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6004887fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b6004897fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b60048a7fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b60048a7f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b60248b7f1e4fbdf700000000000000000000000000000000000000000000000000000000815280600452fd5b015190505f80610e21565b9192937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084167f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793018452828420935b8181106116535750916001939185610e7b9796941061161c575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930155610e73565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f80806115ef565b929360206001819287860151815501950193016115d5565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930190527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f830160051c810191602084106116e4575b601f0160051c01905b8181106116d957508d610e05565b5f81556001016116cb565b90915081906116c2565b60248d7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b015190505f80610d6e565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0167f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930084527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81935b81811061180957509084600195949392106117d2575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055610dc0565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f80806117a5565b9293602060018192878601518155019501930161178f565b917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930091935052601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf810190602084106118cc575b918f91601f85940160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b8181106118b65750610d59565b90925060019193505f815501918f9184936118a9565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81915061187a565b60248e7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00555f610cf8565b60048e7ff92ee8a9000000000000000000000000000000000000000000000000000000008152fd5b9050155f610ca5565b303b159150610c9d565b91508e610c92565b5f80fd5b8280fd5b5080fd5b503461049d576119d436615d1d565b6119dc61659b565b508315611a9757610a7594507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0754927f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854927f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d09549460405196611a5e88615d59565b8752611a686165de565b602088015260408701526060860152608085015260a084015260c083015260e082015260405191829182615f0c565b6004857f247af9ce000000000000000000000000000000000000000000000000000000008152fd5b506101a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602435600435610160367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc01126119bd57604051611b2881615d59565b604435815260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c3601126119b95760405190611b6482615ddb565b6064358252608435602083015260a435604083015260c4356060830152602081019182526040810160e43581526060820191610104358352608081019361012435855260a082019261014435845260c0830161016435815260e0840191610184358352611bcf616cba565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d145498611bfb8a61628b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1455898103611d5d575060809694927f38686cd6969b96c0a5fd318ead3a4b994a396a7d64575f882573e3bde8c6b2b398969492600a92606060408f8f8f90611c66828215156168fc565b8183527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a6020528383205581527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1960205220958751875551805160018801556020810151600288015560408101516003880155015160048601555160058501558651600685015587516007850155855160088501555160098401555191015534611d2c575b5192519151905191604051938452602084015260408301526060820152a380f35b6040513481527f94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a60203392a2611d0b565b8b8a6044927f076edd14000000000000000000000000000000000000000000000000000000008352600452602452fd5b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57611dc860043561692f565b5080604051611dd8602082615e13565b5250610a75604051611deb602082615e13565b5f8152604051918291602083526020830190615c94565b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57611e3a616998565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d075580f35b503461049d57611e7136615fbc565b91611e7a616b20565b611ea7611e868261692f565b82339173ffffffffffffffffffffffffffffffffffffffff33911614616d6e565b811561213357611eb683616a04565b9273ffffffffffffffffffffffffffffffffffffffff84161561210b5761203e908286527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16602052602084611f1e8160408a2054611f1981838a82821115616703565b61621d565b968589527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16835260408920611f5483825461621d565b9055611f81827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5461621d565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a55611fac86616b73565b73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d02541690896040518097819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115612100577fc4032b9f060fe681645ab7a2b55a30a15304792f8669533901de6fb6e0dea1129273ffffffffffffffffffffffffffffffffffffffff926120d3575b506120cd7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5460405193849316973397846040919493926060820195825260208201520152565b0390a480f35b6120f49060203d6020116120f9575b6120ec8183615e13565b8101906166eb565b612086565b503d6120e2565b6040513d88823e3d90fd5b6004857fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b6004847f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b503461049d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57612193615cd7565b5061219c615cfa565b5060643567ffffffffffffffff81116119c157906121bf6004923690840161604b565b507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0754604051908152f35b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57604060209160043581527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a83522054604051908152f35b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57604060209160043581527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1683522054604051908152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5750610a7560405161234f604082615e13565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190615c94565b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760043567ffffffffffffffff81116119c1576123dd90369060040161608f565b9060206040516123ed8282615e13565b848152818101917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810136843761242385616257565b936124316040519586615e13565b85855261243d86616257565b9061246d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084880193018361626f565b87917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301925b88811015612561578060051b8301358481121561255d5783019081359167ffffffffffffffff83116125595786018b833603821361049d578061253d928a6125296001978c8f6040519483869484860198893784019083820190898252519283915e0101858152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282615e13565b5190305af4612536616a6c565b9030617cd0565b612547828b616318565b52612552818a616318565b5001612496565b8b80fd5b8a80fd5b6040805186815289518188018190528c92600582901b830181019186918a9085015b8287106125905785850386f35b9091929382806125ca837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a600196030186528851615c94565b9601920196019592919092612583565b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54604051908152f35b503461049d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5761266d615cd7565b602435906044359060643590612681616b20565b73ffffffffffffffffffffffffffffffffffffffff61269f82616a04565b161561210b576126ae90616a04565b916126c06126bb84616a04565b6172d3565b613549577f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0454936126f08561628b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d04557f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d01547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d05547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d00546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9283166024820152604481019190915291602091839160649183918c91165af1801561353e57613521575b5073ffffffffffffffffffffffffffffffffffffffff841680156134f557848161284c885f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b73ffffffffffffffffffffffffffffffffffffffff8116908a8a84847fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8115946128da86159a8b6133f45773ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190558481527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260408120847fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905580a4156132ef575090507f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254888a527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060408b2055680100000000000000008110156132c257906129cf8260018694017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025561680a565b81549060031b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c831b921b19161790555b036131f7575b6131cb57908591604051828152837f98d1d1eb3af2e9f5c522e0a03cec7374e965eb2a56911d0010ec933ede85affd60203393a47f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d145490604095865191612a6f8884615e13565b600183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe088019384366020860137806131395750895b612aaf8461630b565b52612ab8616756565b81151580613127575b8061311d575b80613110575b613074575b5050612adc6167b0565b87612ae68461630b565b5115158061306a575b8061305d575b612fc6575b50505050505080158015612d63575b505080158015612b1e575b6020848451908152f35b612b26616b20565b61213357612b938373ffffffffffffffffffffffffffffffffffffffff612b8b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b1615156166b8565b8284527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052612bc781838620546162b8565b8385527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d156020528083862055612c1e827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b546162b8565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b55612c4984616b73565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015483517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905290602090829060649082908a9073ffffffffffffffffffffffffffffffffffffffff165af18015612d5957917f8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a98491869360209850612d3e575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b548551948552602085019190915260408401523392606090a35f80612b14565b612d5490883d8a116120f9576120ec8183615e13565b612cfc565b84513d88823e3d90fd5b612d6b616b20565b612f9e57612dd08473ffffffffffffffffffffffffffffffffffffffff612b8b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b8385527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16602052612e0481848720546162b8565b8486527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d166020528084872055612e5b827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a546162b8565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a55612e8685616b73565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025484517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905290602090829060649082908b9073ffffffffffffffffffffffffffffffffffffffff165af18015612f9457917fade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb4793790918793612f77575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a548651948552602085019190915260408401523392606090a35f80612b09565b612f8f9060203d6020116120f9576120ec8183615e13565b612f35565b85513d89823e3d90fd5b6004857f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b7fb36e56ba881ab3b5829611a8b94b9adc2fcdd0e62dc5ef2f97e97d002844ceb09460209461304b9461303787613046968651956130048888615e13565b6001875280368489013761301a88519889615e13565b600188523683890137805161302e8761630b565b520151426162b8565b6130408561630b565b52617a2e565b61630b565b518651908152a2835f80808087612afa565b5060208201511515612af5565b5081511515612aef565b9091929394506130838261790b565b156130e45790887f2d2cd31c49980e0e1f89dbce712aa496e9dafd154af111d0be34d74d9858f6a160206130d66130468c8a8a879c9b9a99613037876130c88c61692f565b928651956130048888615e13565b518b51908152a35f80612ad2565b60248a837fed5a980a000000000000000000000000000000000000000000000000000000008252600452fd5b5060208101511515612acd565b5080511515612ac7565b506131318461630b565b511515612ac1565b9091929394507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161319e579088949392918a527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d19602052878a2054612aa6565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b6024877f73c6ac6e00000000000000000000000000000000000000000000000000000000815280600452fd5b6132036126bb87616a04565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111613295578289527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260408920818a526020528760408a20558789527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020526040892055612a08565b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8282036132fd575b50612a02565b6126bb61330991616a04565b898b527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260408b205490828c527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260408c20918181036133a9575b508a8c527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020528b60408120558b526020528960408120555f6132f7565b818d528260205260408d818120549181848493528660205220558d527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260408d20555f61336a565b8683527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052604083207fffffffffffffffffffffffff0000000000000000000000000000000000000000815416905561348c8973ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff815401905573ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b6024877f64a0ae9200000000000000000000000000000000000000000000000000000000815280600452fd5b6135399060203d6020116120f9576120ec8183615e13565b6127e4565b6040513d89823e3d90fd5b60248573ffffffffffffffffffffffffffffffffffffffff857faa0d86c800000000000000000000000000000000000000000000000000000000835216600452fd5b503461049d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d576135c3615cd7565b506024358015150361049d57807f8cd22d190000000000000000000000000000000000000000000000000000000060049252fd5b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602061363460043561790b565b6040519015158152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1454604051908152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760405190807f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154906136f9826161cc565b80855291600181169081156137b15750600114613735575b610a758461372181860382615e13565b604051918291602083526020830190615c94565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930181527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b8082106137975750909150810160200161372182613711565b91926001816020925483858801015201910190929161377e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b850190920192506137219150839050613711565b50807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57613828616b20565b341561385d576040513481527f94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a60203392a280f35b807f1f2a20050000000000000000000000000000000000000000000000000000000060049252fd5b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d576138bd616998565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095580f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d5760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0454604051908152f35b503461049d5760ff604060209261395636615f88565b9082527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d17855282822090825284522054166040519015158152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57613a39616cba565b613a41616b20565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416604051908152f35b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00541603613bbe57610778336177fc565b807f118cdaa7000000000000000000000000000000000000000000000000000000006024925233600452fd5b503461049d57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d57613c21616998565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00558073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b503461049d57610778613d59613d62613d44366160c0565b9390929197948096613d54616b20565b61739d565b50929091616860565b616dbd565b503461049d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049d576020613da76126bb6105c1615cd7565b604051908152f35b50346119b95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95760043567ffffffffffffffff81116119b957613dff90369060040161608f565b9190613e09615cfa565b90613e12616b20565b831561412457613e2184616860565b613e2a85616860565b915f73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416905b878110613f3d5750803b156119b9575f60405180927fb80f55c900000000000000000000000000000000000000000000000000000000825260206004830152818381613eb78d60248301908a6168bf565b03925af18015613f3257613f1d575b50613ed086616257565b90613ede6040519283615e13565b868252602082019660051b810190368211613f1957965b818810613f09575050610778949550616dbd565b8735815260209788019701613ef5565b8680fd5b613f2a9195505f90615e13565b5f935f613ec6565b6040513d5f823e3d90fd5b613f488189856168af565b356040517f6352211e000000000000000000000000000000000000000000000000000000008152816004820152602081602481875afa908115613f32575f916140d6575b5073ffffffffffffffffffffffffffffffffffffffff339116036140a757604051907f74acd209000000000000000000000000000000000000000000000000000000008252806004830152606082602481875afa918215613f32575f92614052575b5060408201519081421161402457505090602082600193516140108489616318565b52015161401d8288616318565b5201613e66565b7f431e9b4d000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b9091506060813d821161409f575b8161406d60609383615e13565b810103126119b957604080519161408383615da3565b805183526020810151602084015201516040820152905f613fee565b3d9150614060565b7f6d3d1858000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b90506020813d821161411c575b816140f060209383615e13565b810103126119b9575173ffffffffffffffffffffffffffffffffffffffff811681036119b9575f613f8c565b3d91506140e3565b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0054604051908152f35b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95760206141e260043561692f565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416604051908152f35b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95773ffffffffffffffffffffffffffffffffffffffff6142ce816142c0615cd7565b6142c8616998565b16616a04565b16801561435a576143589073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0255565b005b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954604051908152f35b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957602060405160018152f35b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b9576144ae6106d1615cd7565b73ffffffffffffffffffffffffffffffffffffffff81161561435a576143589073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0655565b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036145e25760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b346119b95761461836615f88565b90614621616b20565b8115614124576146888173ffffffffffffffffffffffffffffffffffffffff612b8b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d166020526146bd8260405f20546162b8565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d166020528060405f2055614715837f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a546162b8565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5561474082616b73565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d02546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529190602090839060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af1918215613f32577fade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb479379092614832575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5460408051958652602086019290925290840152339280606081015b0390a3005b61484a9060203d6020116120f9576120ec8183615e13565b6147ef565b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95761488561673e565b50610a756148916167b0565b604051918291829190916020806040830194805184520151910152565b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b9576004357f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02548110156149235761491460209161680a565b90549060031b1c604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f525f60045260245260445ffd5b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957614984615cd7565b60243567ffffffffffffffff81116119b9576149a490369060040161604b565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115614c0f575b506145e2576149f3616998565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181614bdb575b50614a7357837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc859203614bb05750813b15614b8557807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115614b54575f8083602061435895519101845af4614b4e616a6c565b91617cd0565b505034614b5d57005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011614c07575b81614bf760209383615e13565b810103126119b957519085614a42565b3d9150614bea565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54161415836149e6565b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957614c8761673e565b50610a75614891616756565b346119b957614ca136615fbc565b91614caa616b20565b614cb6611e868261692f565b811561412457614cc583616a04565b9273ffffffffffffffffffffffffffffffffffffffff84161561435a57614e3d90825f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052602084614d288160405f2054611f1981838a82821115616703565b96855f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1583528760405f2055614d80827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5461621d565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b55614dab86616b73565b73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416905f6040518097819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115613f32577f1e9c12c534d62150c4f5b916807e2bd6619095393fda9c833471a1feeefad6939273ffffffffffffffffffffffffffffffffffffffff92614ed1575b50614ecc7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5460405193849316973397846040919493926060820195825260208201520152565b0390a4005b614ee99060203d6020116120f9576120ec8183615e13565b614e85565b346119b957614efc36615e9a565b5050505f604051614f0e602082615e13565b527f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957614f6d616cba565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff811615615008577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95773ffffffffffffffffffffffffffffffffffffffff61507c615cd7565b615084616998565b16801561435a5773ffffffffffffffffffffffffffffffffffffffff61104561435892616a04565b346119b9576150ba36615f88565b906150c3616b20565b81156141245761512a8173ffffffffffffffffffffffffffffffffffffffff612b8b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1560205261515f8260405f20546162b8565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d156020528060405f20556151b7837f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b546162b8565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b556151e282616b73565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d01546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529190602090839060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af1918215613f32577f8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a984926152d3575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54604080519586526020860192909252908401523392806060810161482d565b6152eb9060203d6020116120f9576120ec8183615e13565b615291565b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95761532761659b565b506004355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d19602052610a7560405f20600a6040519161536783615d59565b8054835261537760018201616686565b6020840152600581015460408401526006810154606084015260078101546080840152600881015460a0840152600981015460c0840152015460e082015260405191829182615f0c565b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854604051908152f35b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95773ffffffffffffffffffffffffffffffffffffffff61546b816142c0615cd7565b16801561435a576143589073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0155565b346119b95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95761552c615cd7565b73ffffffffffffffffffffffffffffffffffffffff602435916155516126bb82616a04565b83101561559857165f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20905f52602052602060405f2054604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f521660045260245260445ffd5b346119b95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b9576155fe615cd7565b73ffffffffffffffffffffffffffffffffffffffff615622602435926105c1616998565b1690811561435a5780156141245761435891616a9b565b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54604051908152f35b346119b95760206156a336615e54565b6156ab616998565b80517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5501517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d55005b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b9576004355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052602060405f2054604051908152f35b346119b95761576a36615e9a565b5050507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b95760207f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254604051908152f35b346119b95760206157ff36615e54565b615807616998565b80517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5501517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f55005b346119b95761586b61586236615d1d565b9291909161632c565b909160405191606083019360608452825180955260808401602060808760051b8701019401905f905b8782106158af57505050839450602084015260408301520390f35b9091947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80878203018252855190602080835192838152019201905f905b808210615909575050506020806001929701920192019091615894565b909192602060a0600192608087518051835284810151858401526040810151604084015260608101516060840152015160808201520194019201906158ec565b346119b95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957615980615cd7565b507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b9576004356159e48161692f565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b346119b9575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b9576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054615a8d816161cc565b8084529060018116908115615b325750600114615ab5575b610a758361372181850382615e13565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b808210615b1857509091508101602001613721615aa5565b919260018160209254838588010152019101909291615b00565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506137219050615aa5565b346119b95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126119b957600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036119b957817f780e9d630000000000000000000000000000000000000000000000000000000060209314908115615c06575b5015158152f35b7f80ac58cd00000000000000000000000000000000000000000000000000000000811491508115615c6a575b8115615c40575b5083615bff565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483615c39565b7f5b5e139f0000000000000000000000000000000000000000000000000000000081149150615c32565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036119b957565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036119b957565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60809101126119b95760043590602435906044359060643590565b610100810190811067ffffffffffffffff821117615d7657604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6060810190811067ffffffffffffffff821117615d7657604052565b6040810190811067ffffffffffffffff821117615d7657604052565b6080810190811067ffffffffffffffff821117615d7657604052565b60a0810190811067ffffffffffffffff821117615d7657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117615d7657604052565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60409101126119b957604051615e8a81615dbf565b6004358152602435602082015290565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60609101126119b95760043573ffffffffffffffffffffffffffffffffffffffff811681036119b9579060243573ffffffffffffffffffffffffffffffffffffffff811681036119b9579060443590565b91909161014060e061016083019480518452615f4e60208201516020860190606080918051845260208101516020850152604081015160408501520151910152565b604081015160a0850152606081015160c085015260808101518285015260a081015161010085015260c08101516101208501520151910152565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60409101126119b9576004359060243590565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60609101126119b957600435906024359060443573ffffffffffffffffffffffffffffffffffffffff811681036119b95790565b67ffffffffffffffff8111615d7657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f820112156119b95760208135910161606582616011565b926160736040519485615e13565b828452828201116119b957815f92602092838601378301015290565b9181601f840112156119b95782359167ffffffffffffffff83116119b9576020808501948460051b0101116119b957565b9060807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126119b95760043567ffffffffffffffff81116119b9578261610a9160040161608f565b929092916024359167ffffffffffffffff83116119b957806023840112156119b95782600401359267ffffffffffffffff84116119b95781602460608602830101116119b957602401929160443573ffffffffffffffffffffffffffffffffffffffff811681036119b957916064359067ffffffffffffffff82116119b9576161959160040161608f565b9091565b90602080835192838152019201905f5b8181106161b65750505090565b82518452602093840193909201916001016161a9565b90600182811c92168015616213575b60208310146161e657565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f16916161db565b9190820391821161622a57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff8111615d765760051b60200190565b5f5b82811061627d57505050565b606082820152602001616271565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461622a5760010190565b9190820180921161622a57565b80548210156162de575f52600560205f20910201905f90565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8051156162de5760200190565b80518210156162de5760209160051b010190565b929391909361633b848661621d565b61638b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061638161636b84616257565b936163796040519586615e13565b808552616257565b016020830161626f565b9184905b8682108061656e575b1561656257815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1860205260405f209081549384821015616554576163de81836162b8565b9285841161654c575b6163f1838561621d565b926163fb84616257565b916164096040519384615e13565b8483527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061643686616257565b015f5b8181106165175750505f5b8581106164ad575050509061647361647993926164618b8861621d565b9061646c828b616318565b5288616318565b5061621d565b92831590816164a3575b5061649957506164935f9161628b565b9061638f565b9450945050929190565b905081105f616483565b806164c36164bd600193866162b8565b846162c5565b506004604051916164d383615df7565b805483528481015460208401526002810154604084015260038101546060840152015460808201526165058287616318565b526165108186616318565b5001616444565b60209060405161652681615df7565b5f81525f838201525f60408201525f60608201525f608082015282828801015201616439565b8593506163e7565b935050506164935f9161628b565b50929350509250905f90565b50821515616398565b6040519061658482615ddb565b5f6060838281528260208201528260408201520152565b604051906165a882615d59565b5f60e0838281526165b7616577565b60208201528260408201528260608201528260808201528260a08201528260c08201520152565b604051906165eb82615ddb565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d105482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d115460208301527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d125460408301527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d13546060830152565b9060405161669381615ddb565b6060600382948054845260018101546020850152600281015460408501520154910152565b156166c05750565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b908160209103126119b9575180151581036119b95790565b1561670d57505050565b7fda0a959a000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b6040519061674b82615dbf565b5f6020838281520152565b6040519061676382615dbf565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f546020830152565b604051906167bd82615dbf565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d546020830152565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02548110156162de577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025f5260205f2001905f90565b9061686a82616257565b6168776040519182615e13565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06168a58294616257565b0190602036910137565b91908110156162de5760051b0190565b90918281527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83116119b95760209260051b809284830137010190565b156169045750565b7f01707643000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b616977815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff8216156166c0575090565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633036169d857565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014616a58575b15616a5457616a4690617c57565b90616a4f575090565b905090565b5090565b505067ffffffffffffffff81166001616a38565b3d15616a96573d90616a7d82616011565b91616a8b6040519384615e13565b82523d5f602084013e565b606090565b814710616af8575f80809373ffffffffffffffffffffffffffffffffffffffff8294165af1616ac8616a6c565b5015616ad057565b7f3204506f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f356680b7000000000000000000000000000000000000000000000000000000005f5260045ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416616b4b57565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1860205260405f2090805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1660205260405f205491815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1560205260405f20547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d14549060405193616c2885615df7565b84526020840194855260408401908152606084019182526080840192428452805468010000000000000000811015615d7657616c69916001820181556162c5565b959095616c8e5760049451865551600186015551600285015551600384015551910155565b7f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015616d2e575b15616d0257565b7f8e4a23d6000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b5073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0654163314616cfb565b15616d77575050565b73ffffffffffffffffffffffffffffffffffffffff92507f3a0302a9000000000000000000000000000000000000000000000000000000005f526004521660245260445ffd5b919073ffffffffffffffffffffffffffffffffffffffff616ddd85616a04565b161561435a57616def6126bb33616a04565b156172ab57616e006126bb33616a04565b15617265579173ffffffffffffffffffffffffffffffffffffffff33165f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f205f805260205260405f2054915f915f948051915f935b83851061718c575050505050815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1560205260405f205481811061715a57506020616ff091835f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15825260405f20616ed882825461621d565b9055616f05817f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5461621d565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5573ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d01541673ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0554165f6040518096819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af18015613f325761713d575b505f8181527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d18602090815260408083207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d168352818420547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1590935292819020547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1454915190969294929391929091906170b783615df7565b82526020820194855260408201968752606082019283526080820190428252805468010000000000000000811015615d76576170f8916001820181556162c5565b939093616c8e5761713b9773ffffffffffffffffffffffffffffffffffffffff966004945186555160018601555160028501555160038401555191015516616a9b565b565b6171559060203d6020116120f9576120ec8183615e13565b616fff565b917f0cbe2bfe000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b90919293969461719c8883616318565b51976171a88186616318565b51987812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f2281108a021561725857886001937faefa52a7278ff7d9b2c4d4297bfdd0f92e58138de21c4689a25095556c839647604073ffffffffffffffffffffffffffffffffffffffff8f9e8661722e61723492670de0b6b3a7640000830290808204910615150180976162b8565b9e6162b8565b9e61723f888c616318565b5196835195865260208601521692a40193929190616e5d565b637c5f487d5f526004601cfd5b7fa57d13dc000000000000000000000000000000000000000000000000000000005f5273ffffffffffffffffffffffffffffffffffffffff33166004525f60245260445ffd5b7f60852b78000000000000000000000000000000000000000000000000000000005f5260045ffd5b73ffffffffffffffffffffffffffffffffffffffff81161561733a576173369073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b5490565b7f89c62b64000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b1561736f575050565b7f7c8345c2000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b95929193959490948515614124576173b88684818114617366565b6173c58683818114617366565b6173ce86616860565b966173d887616860565b966173e281616860565b965f5b8281106173f55750505050505050565b868110156162de5760608102820161740e8285886168af565b359080359061743e33833373ffffffffffffffffffffffffffffffffffffffff6174378361692f565b1614616d6e565b602081013590811561412457604081013590811561412457835f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1760205260405f20855f5260205260ff60405f2054166177cb57845f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a602052606060405f2054916174ce878415156168fc565b3603126119b9576040516174e181615da3565b84815260208101848152604082018481526040519160208301938985525160408401525160608301525160808201526080815261751f60a082615e13565b519020604051602081019182526020815261753b604082615e13565b519020908b8710156162de578660051b8a01357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18b3603018112156119b9578a0180359067ffffffffffffffff82116119b957602001918160051b360383136119b957835f5b83811061779157500361773b57505050835f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1960205261764560405f20604051906175ec82615d59565b805482526175fc60018201616686565b602083015260058101549182604082015260e0600a60068401549384606085015260078101546080850152600881015460a0850152600981015460c085015201549101526162b8565b80421161770c578f928f928f6176fa897f06541d4525fcfe2db6b9cf209f6bbca0b065559f65c43061f627fdaa9383bd5396856176f483809d9c9b60019f9a61769160409c8b9a616318565b528c5f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d176020528d8b5f20905f526020528a5f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055616318565b52616318565b5282519182526020820152a3016173e5565b7f431e9b4d000000000000000000000000000000000000000000000000000000005f525f60045260245260445ffd5b91617787879287946040519586957f10f7f5f9000000000000000000000000000000000000000000000000000000008752600487015260248601526080604486015260848501916168bf565b9060648301520390fd5b9061779d8285876168af565b3590818110156177ba575f52602052600160405f205b91016175a1565b905f52602052600160405f206177b3565b5050507f6bd4745f000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1660205260405f2054905f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1560205260405f20547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d08548210159182617a02575b8261799957505090565b811592509082156179a957505090565b9091507812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f22821081021561725857670de0b6b3a76400007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1054920204111590565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954821015925061798f565b909283515f905f905b808210617c0a575050478111904791617bdc5750509173ffffffffffffffffffffffffffffffffffffffff5f617ad99593617b398296617b09857f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416966040519a8b998a9889977f1d05159d000000000000000000000000000000000000000000000000000000008952166004880152608060248801526084870190616199565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc868303016044870152616199565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016064850152616199565b03925af1908115613f32575f91617b4e575090565b90503d805f833e617b5f8183615e13565b8101906020818303126119b95780519067ffffffffffffffff82116119b957019080601f830112156119b9578151617b9681616257565b92617ba46040519485615e13565b81845260208085019260051b8201019283116119b957602001905b828210617bcc5750505090565b8151815260209182019101617bbf565b7f7b0e080a000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b9091617c168388616318565b5190617c228488616318565b5191617c2e8587616318565b5190156141245782156141245780421161770c5750600191617c4f916162b8565b920190617a37565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff1603617cc5575b82158015617cba575b617cb257565b5f9250829150565b5060163d1415617cac565b5f9250829150617ca3565b90617d0d5750805115617ce557805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580617d60575b617d1e575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15617d16565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615617d9857565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffdfea264697066735822122001e4068037af3334806c645854da329889db66ae4d28d58de80fcea93871c90164736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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

// IsTokenActive is a free data retrieval call binding the contract method 0x9b0b92ef.
//
// Solidity: function isTokenActive(uint256 _tokenId) view returns(bool)
func (_LPPlus *LPPlusCaller) IsTokenActive(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "isTokenActive", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenActive is a free data retrieval call binding the contract method 0x9b0b92ef.
//
// Solidity: function isTokenActive(uint256 _tokenId) view returns(bool)
func (_LPPlus *LPPlusSession) IsTokenActive(_tokenId *big.Int) (bool, error) {
	return _LPPlus.Contract.IsTokenActive(&_LPPlus.CallOpts, _tokenId)
}

// IsTokenActive is a free data retrieval call binding the contract method 0x9b0b92ef.
//
// Solidity: function isTokenActive(uint256 _tokenId) view returns(bool)
func (_LPPlus *LPPlusCallerSession) IsTokenActive(_tokenId *big.Int) (bool, error) {
	return _LPPlus.Contract.IsTokenActive(&_LPPlus.CallOpts, _tokenId)
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
// Solidity: function previewStakingSnapshot(uint256 _baseConversionRateFilPerRwt, uint256 _stakingTotalYbt, uint256 _timestamp, uint256 _totalFilToAllocate) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256))
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
// Solidity: function previewStakingSnapshot(uint256 _baseConversionRateFilPerRwt, uint256 _stakingTotalYbt, uint256 _timestamp, uint256 _totalFilToAllocate) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusSession) PreviewStakingSnapshot(_baseConversionRateFilPerRwt *big.Int, _stakingTotalYbt *big.Int, _timestamp *big.Int, _totalFilToAllocate *big.Int) (StakingSnapshot, error) {
	return _LPPlus.Contract.PreviewStakingSnapshot(&_LPPlus.CallOpts, _baseConversionRateFilPerRwt, _stakingTotalYbt, _timestamp, _totalFilToAllocate)
}

// PreviewStakingSnapshot is a free data retrieval call binding the contract method 0xd262f444.
//
// Solidity: function previewStakingSnapshot(uint256 _baseConversionRateFilPerRwt, uint256 _stakingTotalYbt, uint256 _timestamp, uint256 _totalFilToAllocate) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256))
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

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LPPlus *LPPlusCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LPPlus *LPPlusSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenByIndex(&_LPPlus.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenByIndex(&_LPPlus.CallOpts, index)
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

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LPPlus *LPPlusCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LPPlus *LPPlusSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenOfOwnerByIndex(&_LPPlus.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LPPlus.Contract.TokenOfOwnerByIndex(&_LPPlus.CallOpts, owner, index)
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

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LPPlus *LPPlusCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPPlus.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LPPlus *LPPlusSession) TotalSupply() (*big.Int, error) {
	return _LPPlus.Contract.TotalSupply(&_LPPlus.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LPPlus *LPPlusCallerSession) TotalSupply() (*big.Int, error) {
	return _LPPlus.Contract.TotalSupply(&_LPPlus.CallOpts)
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
// Solidity: function windowIdToStakingSnapshot(uint256 _windowId) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256))
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
// Solidity: function windowIdToStakingSnapshot(uint256 _windowId) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256))
func (_LPPlus *LPPlusSession) WindowIdToStakingSnapshot(_windowId *big.Int) (StakingSnapshot, error) {
	return _LPPlus.Contract.WindowIdToStakingSnapshot(&_LPPlus.CallOpts, _windowId)
}

// WindowIdToStakingSnapshot is a free data retrieval call binding the contract method 0x344bb6b1.
//
// Solidity: function windowIdToStakingSnapshot(uint256 _windowId) view returns((uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256))
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

// CommitWindow is a paid mutator transaction binding the contract method 0xcf7cdbb8.
//
// Solidity: function commitWindow(uint256 _windowId, bytes32 _root, (uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256) _stakingSnapshot) payable returns()
func (_LPPlus *LPPlusTransactor) CommitWindow(opts *bind.TransactOpts, _windowId *big.Int, _root [32]byte, _stakingSnapshot StakingSnapshot) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "commitWindow", _windowId, _root, _stakingSnapshot)
}

// CommitWindow is a paid mutator transaction binding the contract method 0xcf7cdbb8.
//
// Solidity: function commitWindow(uint256 _windowId, bytes32 _root, (uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256) _stakingSnapshot) payable returns()
func (_LPPlus *LPPlusSession) CommitWindow(_windowId *big.Int, _root [32]byte, _stakingSnapshot StakingSnapshot) (*types.Transaction, error) {
	return _LPPlus.Contract.CommitWindow(&_LPPlus.TransactOpts, _windowId, _root, _stakingSnapshot)
}

// CommitWindow is a paid mutator transaction binding the contract method 0xcf7cdbb8.
//
// Solidity: function commitWindow(uint256 _windowId, bytes32 _root, (uint256,(uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256) _stakingSnapshot) payable returns()
func (_LPPlus *LPPlusTransactorSession) CommitWindow(_windowId *big.Int, _root [32]byte, _stakingSnapshot StakingSnapshot) (*types.Transaction, error) {
	return _LPPlus.Contract.CommitWindow(&_LPPlus.TransactOpts, _windowId, _root, _stakingSnapshot)
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

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_LPPlus *LPPlusTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _LPPlus.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_LPPlus *LPPlusSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _LPPlus.Contract.Multicall(&_LPPlus.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_LPPlus *LPPlusTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _LPPlus.Contract.Multicall(&_LPPlus.TransactOpts, data)
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
	TokenId          *big.Int
	RwtFutureTokenId *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBonusApplied is a free log retrieval operation binding the contract event 0xb36e56ba881ab3b5829611a8b94b9adc2fcdd0e62dc5ef2f97e97d002844ceb0.
//
// Solidity: event BonusApplied(uint256 indexed tokenId, uint256 rwtFutureTokenId)
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

// WatchBonusApplied is a free log subscription operation binding the contract event 0xb36e56ba881ab3b5829611a8b94b9adc2fcdd0e62dc5ef2f97e97d002844ceb0.
//
// Solidity: event BonusApplied(uint256 indexed tokenId, uint256 rwtFutureTokenId)
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

// ParseBonusApplied is a log parse operation binding the contract event 0xb36e56ba881ab3b5829611a8b94b9adc2fcdd0e62dc5ef2f97e97d002844ceb0.
//
// Solidity: event BonusApplied(uint256 indexed tokenId, uint256 rwtFutureTokenId)
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

// LPPlusNewReferralIterator is returned from FilterNewReferral and is used to iterate over the raw logs and unpacked data for NewReferral events raised by the LPPlus contract.
type LPPlusNewReferralIterator struct {
	Event *LPPlusNewReferral // Event containing the contract specifics and raw log

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
func (it *LPPlusNewReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusNewReferral)
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
		it.Event = new(LPPlusNewReferral)
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
func (it *LPPlusNewReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusNewReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusNewReferral represents a NewReferral event raised by the LPPlus contract.
type LPPlusNewReferral struct {
	NewTokenId       *big.Int
	ReferrerTokenId  *big.Int
	RwtFutureTokenId *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewReferral is a free log retrieval operation binding the contract event 0x2d2cd31c49980e0e1f89dbce712aa496e9dafd154af111d0be34d74d9858f6a1.
//
// Solidity: event NewReferral(uint256 indexed newTokenId, uint256 indexed referrerTokenId, uint256 rwtFutureTokenId)
func (_LPPlus *LPPlusFilterer) FilterNewReferral(opts *bind.FilterOpts, newTokenId []*big.Int, referrerTokenId []*big.Int) (*LPPlusNewReferralIterator, error) {

	var newTokenIdRule []interface{}
	for _, newTokenIdItem := range newTokenId {
		newTokenIdRule = append(newTokenIdRule, newTokenIdItem)
	}
	var referrerTokenIdRule []interface{}
	for _, referrerTokenIdItem := range referrerTokenId {
		referrerTokenIdRule = append(referrerTokenIdRule, referrerTokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "NewReferral", newTokenIdRule, referrerTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusNewReferralIterator{contract: _LPPlus.contract, event: "NewReferral", logs: logs, sub: sub}, nil
}

// WatchNewReferral is a free log subscription operation binding the contract event 0x2d2cd31c49980e0e1f89dbce712aa496e9dafd154af111d0be34d74d9858f6a1.
//
// Solidity: event NewReferral(uint256 indexed newTokenId, uint256 indexed referrerTokenId, uint256 rwtFutureTokenId)
func (_LPPlus *LPPlusFilterer) WatchNewReferral(opts *bind.WatchOpts, sink chan<- *LPPlusNewReferral, newTokenId []*big.Int, referrerTokenId []*big.Int) (event.Subscription, error) {

	var newTokenIdRule []interface{}
	for _, newTokenIdItem := range newTokenId {
		newTokenIdRule = append(newTokenIdRule, newTokenIdItem)
	}
	var referrerTokenIdRule []interface{}
	for _, referrerTokenIdItem := range referrerTokenId {
		referrerTokenIdRule = append(referrerTokenIdRule, referrerTokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "NewReferral", newTokenIdRule, referrerTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusNewReferral)
				if err := _LPPlus.contract.UnpackLog(event, "NewReferral", log); err != nil {
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

// ParseNewReferral is a log parse operation binding the contract event 0x2d2cd31c49980e0e1f89dbce712aa496e9dafd154af111d0be34d74d9858f6a1.
//
// Solidity: event NewReferral(uint256 indexed newTokenId, uint256 indexed referrerTokenId, uint256 rwtFutureTokenId)
func (_LPPlus *LPPlusFilterer) ParseNewReferral(log types.Log) (*LPPlusNewReferral, error) {
	event := new(LPPlusNewReferral)
	if err := _LPPlus.contract.UnpackLog(event, "NewReferral", log); err != nil {
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
	Receiver         common.Address
	SenderTokenId    *big.Int
	RwtFutureTokenId *big.Int
	RwtAmount        *big.Int
	FilAmount        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRWTExchangedForFIL is a free log retrieval operation binding the contract event 0xaefa52a7278ff7d9b2c4d4297bfdd0f92e58138de21c4689a25095556c839647.
//
// Solidity: event RWTExchangedForFIL(address indexed receiver, uint256 indexed senderTokenId, uint256 indexed rwtFutureTokenId, uint256 rwtAmount, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) FilterRWTExchangedForFIL(opts *bind.FilterOpts, receiver []common.Address, senderTokenId []*big.Int, rwtFutureTokenId []*big.Int) (*LPPlusRWTExchangedForFILIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var senderTokenIdRule []interface{}
	for _, senderTokenIdItem := range senderTokenId {
		senderTokenIdRule = append(senderTokenIdRule, senderTokenIdItem)
	}
	var rwtFutureTokenIdRule []interface{}
	for _, rwtFutureTokenIdItem := range rwtFutureTokenId {
		rwtFutureTokenIdRule = append(rwtFutureTokenIdRule, rwtFutureTokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "RWTExchangedForFIL", receiverRule, senderTokenIdRule, rwtFutureTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusRWTExchangedForFILIterator{contract: _LPPlus.contract, event: "RWTExchangedForFIL", logs: logs, sub: sub}, nil
}

// WatchRWTExchangedForFIL is a free log subscription operation binding the contract event 0xaefa52a7278ff7d9b2c4d4297bfdd0f92e58138de21c4689a25095556c839647.
//
// Solidity: event RWTExchangedForFIL(address indexed receiver, uint256 indexed senderTokenId, uint256 indexed rwtFutureTokenId, uint256 rwtAmount, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) WatchRWTExchangedForFIL(opts *bind.WatchOpts, sink chan<- *LPPlusRWTExchangedForFIL, receiver []common.Address, senderTokenId []*big.Int, rwtFutureTokenId []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var senderTokenIdRule []interface{}
	for _, senderTokenIdItem := range senderTokenId {
		senderTokenIdRule = append(senderTokenIdRule, senderTokenIdItem)
	}
	var rwtFutureTokenIdRule []interface{}
	for _, rwtFutureTokenIdItem := range rwtFutureTokenId {
		rwtFutureTokenIdRule = append(rwtFutureTokenIdRule, rwtFutureTokenIdItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "RWTExchangedForFIL", receiverRule, senderTokenIdRule, rwtFutureTokenIdRule)
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

// ParseRWTExchangedForFIL is a log parse operation binding the contract event 0xaefa52a7278ff7d9b2c4d4297bfdd0f92e58138de21c4689a25095556c839647.
//
// Solidity: event RWTExchangedForFIL(address indexed receiver, uint256 indexed senderTokenId, uint256 indexed rwtFutureTokenId, uint256 rwtAmount, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) ParseRWTExchangedForFIL(log types.Log) (*LPPlusRWTExchangedForFIL, error) {
	event := new(LPPlusRWTExchangedForFIL)
	if err := _LPPlus.contract.UnpackLog(event, "RWTExchangedForFIL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPPlusRWTFutureClaimedIterator is returned from FilterRWTFutureClaimed and is used to iterate over the raw logs and unpacked data for RWTFutureClaimed events raised by the LPPlus contract.
type LPPlusRWTFutureClaimedIterator struct {
	Event *LPPlusRWTFutureClaimed // Event containing the contract specifics and raw log

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
func (it *LPPlusRWTFutureClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusRWTFutureClaimed)
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
		it.Event = new(LPPlusRWTFutureClaimed)
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
func (it *LPPlusRWTFutureClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusRWTFutureClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusRWTFutureClaimed represents a RWTFutureClaimed event raised by the LPPlus contract.
type LPPlusRWTFutureClaimed struct {
	TokenId     *big.Int
	WindowId    *big.Int
	StrikePrice *big.Int
	FilAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRWTFutureClaimed is a free log retrieval operation binding the contract event 0x06541d4525fcfe2db6b9cf209f6bbca0b065559f65c43061f627fdaa9383bd53.
//
// Solidity: event RWTFutureClaimed(uint256 indexed tokenId, uint256 indexed windowId, uint256 strikePrice, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) FilterRWTFutureClaimed(opts *bind.FilterOpts, tokenId []*big.Int, windowId []*big.Int) (*LPPlusRWTFutureClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var windowIdRule []interface{}
	for _, windowIdItem := range windowId {
		windowIdRule = append(windowIdRule, windowIdItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "RWTFutureClaimed", tokenIdRule, windowIdRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusRWTFutureClaimedIterator{contract: _LPPlus.contract, event: "RWTFutureClaimed", logs: logs, sub: sub}, nil
}

// WatchRWTFutureClaimed is a free log subscription operation binding the contract event 0x06541d4525fcfe2db6b9cf209f6bbca0b065559f65c43061f627fdaa9383bd53.
//
// Solidity: event RWTFutureClaimed(uint256 indexed tokenId, uint256 indexed windowId, uint256 strikePrice, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) WatchRWTFutureClaimed(opts *bind.WatchOpts, sink chan<- *LPPlusRWTFutureClaimed, tokenId []*big.Int, windowId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var windowIdRule []interface{}
	for _, windowIdItem := range windowId {
		windowIdRule = append(windowIdRule, windowIdItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "RWTFutureClaimed", tokenIdRule, windowIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusRWTFutureClaimed)
				if err := _LPPlus.contract.UnpackLog(event, "RWTFutureClaimed", log); err != nil {
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

// ParseRWTFutureClaimed is a log parse operation binding the contract event 0x06541d4525fcfe2db6b9cf209f6bbca0b065559f65c43061f627fdaa9383bd53.
//
// Solidity: event RWTFutureClaimed(uint256 indexed tokenId, uint256 indexed windowId, uint256 strikePrice, uint256 filAmount)
func (_LPPlus *LPPlusFilterer) ParseRWTFutureClaimed(log types.Log) (*LPPlusRWTFutureClaimed, error) {
	event := new(LPPlusRWTFutureClaimed)
	if err := _LPPlus.contract.UnpackLog(event, "RWTFutureClaimed", log); err != nil {
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

// LPPlusWindowCommittedIterator is returned from FilterWindowCommitted and is used to iterate over the raw logs and unpacked data for WindowCommitted events raised by the LPPlus contract.
type LPPlusWindowCommittedIterator struct {
	Event *LPPlusWindowCommitted // Event containing the contract specifics and raw log

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
func (it *LPPlusWindowCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPPlusWindowCommitted)
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
		it.Event = new(LPPlusWindowCommitted)
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
func (it *LPPlusWindowCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPPlusWindowCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPPlusWindowCommitted represents a WindowCommitted event raised by the LPPlus contract.
type LPPlusWindowCommitted struct {
	WindowId                    *big.Int
	Root                        [32]byte
	BaseConversionRateFilPerRwt *big.Int
	StakingYbtTwab              *big.Int
	Timestamp                   *big.Int
	TotalFilToAllocate          *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterWindowCommitted is a free log retrieval operation binding the contract event 0x38686cd6969b96c0a5fd318ead3a4b994a396a7d64575f882573e3bde8c6b2b3.
//
// Solidity: event WindowCommitted(uint256 indexed windowId, bytes32 indexed root, uint256 baseConversionRateFilPerRwt, uint256 stakingYbtTwab, uint256 timestamp, uint256 totalFilToAllocate)
func (_LPPlus *LPPlusFilterer) FilterWindowCommitted(opts *bind.FilterOpts, windowId []*big.Int, root [][32]byte) (*LPPlusWindowCommittedIterator, error) {

	var windowIdRule []interface{}
	for _, windowIdItem := range windowId {
		windowIdRule = append(windowIdRule, windowIdItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _LPPlus.contract.FilterLogs(opts, "WindowCommitted", windowIdRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &LPPlusWindowCommittedIterator{contract: _LPPlus.contract, event: "WindowCommitted", logs: logs, sub: sub}, nil
}

// WatchWindowCommitted is a free log subscription operation binding the contract event 0x38686cd6969b96c0a5fd318ead3a4b994a396a7d64575f882573e3bde8c6b2b3.
//
// Solidity: event WindowCommitted(uint256 indexed windowId, bytes32 indexed root, uint256 baseConversionRateFilPerRwt, uint256 stakingYbtTwab, uint256 timestamp, uint256 totalFilToAllocate)
func (_LPPlus *LPPlusFilterer) WatchWindowCommitted(opts *bind.WatchOpts, sink chan<- *LPPlusWindowCommitted, windowId []*big.Int, root [][32]byte) (event.Subscription, error) {

	var windowIdRule []interface{}
	for _, windowIdItem := range windowId {
		windowIdRule = append(windowIdRule, windowIdItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _LPPlus.contract.WatchLogs(opts, "WindowCommitted", windowIdRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPPlusWindowCommitted)
				if err := _LPPlus.contract.UnpackLog(event, "WindowCommitted", log); err != nil {
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

// ParseWindowCommitted is a log parse operation binding the contract event 0x38686cd6969b96c0a5fd318ead3a4b994a396a7d64575f882573e3bde8c6b2b3.
//
// Solidity: event WindowCommitted(uint256 indexed windowId, bytes32 indexed root, uint256 baseConversionRateFilPerRwt, uint256 stakingYbtTwab, uint256 timestamp, uint256 totalFilToAllocate)
func (_LPPlus *LPPlusFilterer) ParseWindowCommitted(log types.Log) (*LPPlusWindowCommitted, error) {
	event := new(LPPlusWindowCommitted)
	if err := _LPPlus.contract.UnpackLog(event, "WindowCommitted", log); err != nil {
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
