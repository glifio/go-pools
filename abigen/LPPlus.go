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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"}],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"CardAlreadyMinted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"FilReservesDepleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFILBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientRWTBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientVaultBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_second\",\"type\":\"uint256\"}],\"name\":\"InvalidLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"name\":\"RWTFutureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenNotActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroConversionRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"}],\"name\":\"ZeroMerkleRoot\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structRWTFutureBonus\",\"name\":\"bonus\",\"type\":\"tuple\"}],\"name\":\"BonusApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referrerTokenId\",\"type\":\"uint256\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FilVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"RWTDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rwtFutureTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rwtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filAmount\",\"type\":\"uint256\"}],\"name\":\"RWTExchangedForFIL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"RWTWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"YBTVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractTotal\",\"type\":\"uint256\"}],\"name\":\"YBTVaultWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_RWTAmount\",\"type\":\"uint256\"}],\"name\":\"addRWT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_YBTAmount\",\"type\":\"uint256\"}],\"name\":\"addYBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_windowIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"}],\"internalType\":\"structLpPlusMerkleHelper.UserSnapshot[]\",\"name\":\"_userSnapshots\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_proofs\",\"type\":\"bytes32[][]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_windowIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"}],\"internalType\":\"structLpPlusMerkleHelper.UserSnapshot[]\",\"name\":\"_userSnapshots\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_proofs\",\"type\":\"bytes32[][]\"}],\"name\":\"claimAndExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTotalYbt\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"_stakingSnapshot\",\"type\":\"tuple\"}],\"name\":\"commitWindow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_rwtFutureTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundFilVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureValidityDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getUserActionsBatched\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ybtBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rwtBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"windowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structUserAction[][]\",\"name\":\"actionsPerToken\",\"type\":\"tuple[][]\"},{\"internalType\":\"uint256\",\"name\":\"newStartTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newOffset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rwtToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_ybtToken\",\"type\":\"address\"},{\"internalType\":\"contractIRWTFuture\",\"name\":\"_rwtFuture\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_windowCommitter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interpolationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumRWTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumYBTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_referrerTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ybtToStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rwtToStake\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBonus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingTotalYbt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalFilToAllocate\",\"type\":\"uint256\"}],\"name\":\"previewStakingSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTotalYbt\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referrerBonus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_RWTAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeRWT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_YBTAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeYBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwtFuture\",\"outputs\":[{\"internalType\":\"contractIRWTFuture\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_futureValidityDuration\",\"type\":\"uint256\"}],\"name\":\"setFutureValidityDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"_interpolationParams\",\"type\":\"tuple\"}],\"name\":\"setInterpolationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumRWTBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumRWTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumYBTBalance\",\"type\":\"uint256\"}],\"name\":\"setMinimumYBTBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"_mintBonus\",\"type\":\"tuple\"}],\"name\":\"setMintBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRWTFuture\",\"name\":\"_RWTFuture\",\"type\":\"address\"}],\"name\":\"setRWTFuture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_RWTToken\",\"type\":\"address\"}],\"name\":\"setRWTToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"allocatedFilExtra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateExtra\",\"type\":\"uint256\"}],\"internalType\":\"structRWTFutureBonus\",\"name\":\"_referrerBonus\",\"type\":\"tuple\"}],\"name\":\"setReferrerBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_windowCommitter\",\"type\":\"address\"}],\"name\":\"setWindowCommitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_YBTToken\",\"type\":\"address\"}],\"name\":\"setYBTToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToRWTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToWindowIdToClaimedFlag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToYBTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRwtStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalYbtStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"windowCommitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"windowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"windowIdToMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_windowId\",\"type\":\"uint256\"}],\"name\":\"windowIdToStakingSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseConversionRateFilPerRwt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structRwtInterpolation\",\"name\":\"interpolationParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"futureValidityDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingYbtTwab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFilToAllocate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTotalYbt\",\"type\":\"uint256\"}],\"internalType\":\"structStakingSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFilFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ybtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f516020617eb95f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051617df290816100c782396080518181816145f40152614a4d0152f35b6001600160401b0319166001600160401b039081175f516020617eb95f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a714615fde5750806306fdde0314615e9a578063081812fc14615e13578063095ea7b314615db35780630a6b569914615cbb5780631418f71114615c5957806316290e1d146159a357806318160ddd1461594957806323b872dd1461591057806329f2014d146158a95780632bf2ba03146158475780632e58c0c7146157ed5780632eacc6961461573e5780632f745c591461566c578063300e3c4f146155925780633263edee14615538578063344bb6b114615471578063385b5f85146151fd5780633a20853e146151815780633f4ba83a14614fff57806342842e0e14614fb657806346b3241e14614d255780634bae2ef114614ce35780634f1ef286146149e45780634f6ccce714614940578063510402cd146148e1578063519473611461466c57806352d1902d146145af578063548d9356146144d657806354fd4d501461449d57806358f28a98146144435780635c975abb146143e45780635e29e06b146142d257806361d027b3146142625780636352211e146142085780636817c76c146141ae578063694072a614613e4457806370a0823114613dfc57806370af96aa14613dc1578063715018a614613c7f57806379ba509714613bdd578063812bf90a14613b6c5780638456cb5914613a0e5780638da5cb5b1461399d5780638edd7d4e1461394c5780638f68679f146138f15780638ffd6d2d14613891578063912934731461380257806395d89b41146136a55780639af1d03e1461364a578063a22cb465146135de578063a647e8ec146125b4578063a818b1ee14612559578063ac9650d81461230c578063ad3cb1cc1461228d578063ae7e2d9714612226578063b05dcd5a146121bf578063b5ea015c14612164578063b88d4fde146120da578063bfb1909b14611da3578063c74c23d514611d43578063c87b56dd14611cce578063d262f44414611bfa578063d452154814610d7f578063d5d2ae7214610d0e578063da357c0f146109f0578063e080e4bd14610990578063e0ea4a1b14610921578063e30c3978146108b0578063e985e9c5146107ff578063ee2df9fd1461078e578063f0f4426014610681578063f2fde38b1461055c578063f4a0a528146104fc578063f721dc761461048b5763fcd0907d14610362575f80fd5b346104885760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760405161039d81616244565b600435815260208101602435815260408201906044358252606083019260643584526103c7616ec4565b82518151111561046057517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1055517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1155517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1255517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d135580f35b6004857f1f3b85d3000000000000000000000000000000000000000000000000000000008152fd5b80fd5b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416604051908152f35b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857610534616ec4565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d005580f35b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885773ffffffffffffffffffffffffffffffffffffffff6105b16105ac616141565b616f30565b6105b9616ec4565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610488576106c46106bc616141565b6105ac616ec4565b73ffffffffffffffffffffffffffffffffffffffff811615610766576107639073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0555565b80f35b6004827fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416604051908152f35b50346104885760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857610837616141565b73ffffffffffffffffffffffffffffffffffffffff610854616164565b91165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857610958616a52565b506080610963616ab3565b61098e6040518092606080918051845260208101516020850152604081015160408501520151910152565bf35b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610488576109c8616ec4565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085580f35b503461048857610a2873ffffffffffffffffffffffffffffffffffffffff610a173661651e565b96916105ac9891969895939561704c565b16958615610ce657610a3b9697506175b0565b91909281515f905f905b808210610c42575050478111904791610c1457855f610ae387878388610b4373ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d03541694610b13604051998a98899788967f1d05159d00000000000000000000000000000000000000000000000000000000885260048801526080602488015260848701906165f7565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8683030160448701526165f7565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8483030160648501526165f7565b03925af18015610c09575f90610b72575b80610b6e91506040519182916020835260208301906165f7565b0390f35b503d805f833e610b82818361627c565b810190602081830312610c055780519067ffffffffffffffff8211610c0557019080601f83011215610c05578151610bb9816166b5565b92610bc7604051948561627c565b81845260208085019260051b820101928311610c0557602001905b828210610bf557505050610b6e90610b54565b8151815260209182019101610be2565b5f80fd5b6040513d5f823e3d90fd5b7f7b0e080a000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b9091610c4e8386616776565b5190610c5a8489616776565b5191610c668589616776565b519015610cbe578215610cbe57804211610c8f5750600191610c8791616716565b920190610a45565b7f431e9b4d000000000000000000000000000000000000000000000000000000005f525f60045260245260445ffd5b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b6004887fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416604051908152f35b5034610488576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760043567ffffffffffffffff8111611bf657610dd09036906004016164a9565b60243567ffffffffffffffff8111611bf257610df09036906004016164a9565b60443573ffffffffffffffffffffffffffffffffffffffff81168103610c055760643560843573ffffffffffffffffffffffffffffffffffffffff8116809103610c055760a4359073ffffffffffffffffffffffffffffffffffffffff8216809203610c055760c4359273ffffffffffffffffffffffffffffffffffffffff8416809403610c055760e4359473ffffffffffffffffffffffffffffffffffffffff8616808703610c0557610104359773ffffffffffffffffffffffffffffffffffffffff891692838a03610c05577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549a60ff8c60401c16159b8c67ffffffffffffffff821680159182611bea575b506001149081611be0575b159081611bd7575b50611baf578c60017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055611b5a575b50610f6e617d65565b610f76617d65565b610f7e617d65565b80519067ffffffffffffffff8211611b2d578d8291610fbd7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005461662a565b601f8111611a5a575b50602091601f841160011461195f5792611954575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b815167ffffffffffffffff8111611927576110687f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015461662a565b8d601f82116118a4575b90505060208d601f83116001146117c057906110e894836117b5575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b6105ac617d65565b6110f0617d65565b6110f8617d65565b73ffffffffffffffffffffffffffffffffffffffff8116156117895761111d906179a9565b611125617d65565b61112d617d65565b821561176157831561173957841561173957851561173957156117115715610ce657936105ac73ffffffffffffffffffffffffffffffffffffffff6112b261143098966105ac836112326105ac996105ac836111b26113b19f9c6113329d7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0055616f30565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0155565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0255565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0355565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0555565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0655565b60017f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d045562ed4e007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d075566038d7ea4c680007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085566038d7ea4c680007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095562ed4e0060206040516114e081616228565b662386f26fc1000081520152662386f26fc100007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5562ed4e007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d55662386f26fc100007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5562ed4e007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f5567016345785d8a000060606040516115a381616244565b670de0b6b3a76400008152666a94d74f43000060208201526706f05b59d3b2000060408201520152670de0b6b3a76400007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1055666a94d74f4300007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d11556706f05b59d3b200007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d125567016345785d8a00007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d135561167d5780f35b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6004897fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b60048a7fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b60048a7f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b60248b7f1e4fbdf700000000000000000000000000000000000000000000000000000000815280600452fd5b015190505f8061108e565b9192937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084167f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793018452828420935b81811061188c57509160019391856110e897969410611855575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301556110e0565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f8080611828565b9293602060018192878601518155019501930161180e565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930190527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f830160051c8101916020841061191d575b601f0160051c01905b81811061191257508d611072565b5f8155600101611904565b90915081906118fb565b60248d7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b015190505f80610fdb565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0167f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930084527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81935b818110611a425750908460019594939210611a0b575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005561102d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f80806119de565b929360206001819287860151815501950193016119c8565b917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930091935052601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410611b05575b918f91601f85940160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b818110611aef5750610fc6565b90925060019193505f815501918f918493611ae2565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf819150611ab3565b60248e7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00555f610f65565b60048e7ff92ee8a9000000000000000000000000000000000000000000000000000000008152fd5b9050155f610f12565b303b159150610f0a565b91508e610eff565b8280fd5b5080fd5b503461048857611c0936616187565b9091611c13616a76565b508315611ca657610b6e94507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0754927f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a549360405195611c72876161c3565b8652611c7c616ab3565b602087015260408601526060850152608084015260a083015260c082015260405191829182616375565b6004857f247af9ce000000000000000000000000000000000000000000000000000000008152fd5b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857611d09600435616e5b565b5080604051611d1960208261627c565b5250610b6e604051611d2c60208261627c565b5f81526040519182916020835260208301906160fe565b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857611d7b616ec4565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d075580f35b503461048857611db23661641a565b91611dbb61704c565b611de8611dc782616e5b565b82339173ffffffffffffffffffffffffffffffffffffffff33911614616c13565b81156120b257611df783616f30565b9273ffffffffffffffffffffffffffffffffffffffff84161561208a57611fb5908286527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16602052602084611e5f8160408a2054611e5a81838a82821115616c62565b61667b565b96611e967f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854898882821080159061208257616bc0565b8589527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16835260408920611ecb83825461667b565b9055611ef8827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5461667b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a55611f238661709f565b73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d02541690896040518097819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115612077577fc4032b9f060fe681645ab7a2b55a30a15304792f8669533901de6fb6e0dea1129273ffffffffffffffffffffffffffffffffffffffff9261204a575b506120447f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5460405193849316973397846040919493926060820195825260208201520152565b0390a480f35b61206b9060203d602011612070575b612063818361627c565b810190616bfb565b611ffd565b503d612059565b6040513d88823e3d90fd5b508115616bc0565b6004857fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b6004847f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b50346104885760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857612112616141565b5061211b616164565b5060643567ffffffffffffffff8111611bf6579061213e600492369084016164a9565b507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0754604051908152f35b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857604060209160043581527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a83522054604051908152f35b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857604060209160043581527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1683522054604051908152f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885750610b6e6040516122ce60408261627c565b600581527f352e302e3000000000000000000000000000000000000000000000000000000060208201526040519182916020835260208301906160fe565b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760043567ffffffffffffffff8111611bf65761235c9036906004016164ed565b90602060405161236c828261627c565b848152818101917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081013684376123a2856166b5565b936123b0604051958661627c565b8585526123bc866166b5565b906123ec7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08488019301836166cd565b87917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301925b888110156124e0578060051b830135848112156124dc5783019081359167ffffffffffffffff83116124d85786018b833603821361048857806124bc928a6124a86001978c8f6040519483869484860198893784019083820190898252519283915e0101858152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261627c565b5190305af46124b5616f98565b9030617ccc565b6124c6828b616776565b526124d1818a616776565b5001612415565b8b80fd5b8a80fd5b6040805186815289518188018190528c92600582901b830181019186918a9085015b82871061250f5785850386f35b909192938280612549837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a6001960301865288516160fe565b9601920196019592919092612502565b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54604051908152f35b50346104885760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610488576125ec616141565b6024356064356044356125fd61704c565b73ffffffffffffffffffffffffffffffffffffffff84169384156135b65761262490616f30565b61263561263082616f30565b6174e6565b613575577f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d045494612665866166e9565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d04557f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d01547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d05547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d00546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9283166024820152604481019190915291602091839160649183918d91165af1801561356a5761354d575b5073ffffffffffffffffffffffffffffffffffffffff82168015613521578088916127c2895f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b73ffffffffffffffffffffffffffffffffffffffff81169081158b84847fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8415998a61345e575b6128508d73ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190558481527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260408120847fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905580a41561335a575090507f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254898b527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060408c20556801000000000000000081101561332d57906129458260018594017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255616d69565b81549060031b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8d831b921b19161790555b0361328e575b5061326257604051858152867f98d1d1eb3af2e9f5c522e0a03cec7374e965eb2a56911d0010ec933ede85affd60203393a47f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d14546040948551906129e2878361627c565b600182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe087019283366020850137806131dc5750885b612a2283616769565b52612a2b616cb5565b90801515806131ca575b806131c0575b806131b3575b61304b575b5050612a50616d0f565b90612a5a81616769565b51151580613041575b80613034575b612f98575b5050505080158015612d05575b505080158015612a90575b6020848451908152f35b612a9861704c565b6120b257612b058373ffffffffffffffffffffffffffffffffffffffff612afd825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b161515616b8d565b8284527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052612b398183862054616716565b612b697f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954828682821015616bc0565b8385527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d156020528083862055612bc0827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54616716565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b55612beb8461709f565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015483517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905290602090829060649082908a9073ffffffffffffffffffffffffffffffffffffffff165af18015612cfb57917f8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a98491869360209850612ce0575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b548551948552602085019190915260408401523392606090a35f80612a86565b612cf690883d8a1161207057612063818361627c565b612c9e565b84513d88823e3d90fd5b612d0d61704c565b612f7057612d728473ffffffffffffffffffffffffffffffffffffffff612afd825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b8385527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16602052612da68184872054616716565b612dd67f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854828782821015616bc0565b8486527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d166020528084872055612e2d827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54616716565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a55612e588561709f565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025484517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905290602090829060649082908b9073ffffffffffffffffffffffffffffffffffffffff165af18015612f6657917fade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb4793790918793612f49575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a548651948552602085019190915260408401523392606090a35f80612a7b565b612f619060203d60201161207057612063818361627c565b612f07565b85513d89823e3d90fd5b6004857f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b8651825181526020808401519082015261302a949089907f08c20988b3a0d46fafe47abf0aa2ac3716335a2190a6acb020f12468259207ac90604090a261301b6020895194612fe78b8761627c565b600186528636838801378a5196612ffe8c8961627c565b600188523683890137805161301287616769565b52015142616716565b61302485616769565b52617ab8565b505f808080612a6e565b5060208201511515612a69565b5081511515612a63565b808a527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d16602052878a20547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d085411158061315e575b156131335761312b9183826130b58b94616e5b565b845184518152602080860151908201529091907f08c20988b3a0d46fafe47abf0aa2ac3716335a2190a6acb020f12468259207ac90604090a261301b6020855194613100878761627c565b600186528936838801376131168751978861627c565b60018752893683890137805161301287616769565b505f80612a46565b7fed5a980a000000000000000000000000000000000000000000000000000000008a52600452602489fd5b50808a527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052878a20547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095411156130a0565b5060208201511515612a41565b5081511515612a3b565b506131d483616769565b511515612a35565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019081116132355789527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1960205286892054612a19565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b6024877f73c6ac6e00000000000000000000000000000000000000000000000000000000815280600452fd5b61329a61263085616f30565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116132355789527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260408920818a526020528760408a20558789527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260408920555f61297e565b60248b7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b828203613368575b50612978565b61263061337491616f30565b8a8c527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260408c205490828d527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260408d20918d828203613415575b8d9150527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020528c60408120558c526020528a60408120555f613362565b808360409252846020528181205483825285602052808383205581527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205220555f8d6133d6565b8481527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052604081207fffffffffffffffffffffffff000000000000000000000000000000000000000081541690556134f68773ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8154019055612809565b6024887f64a0ae9200000000000000000000000000000000000000000000000000000000815280600452fd5b6135659060203d60201161207057612063818361627c565b612759565b6040513d8a823e3d90fd5b7faa0d86c800000000000000000000000000000000000000000000000000000000865273ffffffffffffffffffffffffffffffffffffffff16600452602485fd5b6004867fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b50346104885760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857613616616141565b506024358015150361048857807f8cd22d190000000000000000000000000000000000000000000000000000000060049252fd5b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1454604051908152f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760405190807f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154906137058261662a565b80855291600181169081156137bd5750600114613741575b610b6e8461372d8186038261627c565b6040519182916020835260208301906160fe565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930181527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b8082106137a35750909150810160200161372d8261371d565b91926001816020925483858801015201910190929161378a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b8501909201925061372d915083905061371d565b50807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885761383461704c565b3415613869576040513481527f94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a60203392a280f35b807f1f2a20050000000000000000000000000000000000000000000000000000000060049252fd5b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610488576138c9616ec4565b6004357f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d095580f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0454604051908152f35b50346104885760ff6040602092613962366163e6565b9082527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d17855282822090825284522054166040519015158152f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104885773ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015613b27575b613a899033906169d5565b613a9161704c565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b50613a8973ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d06541633149050613a7e565b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416604051908152f35b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610488573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00541603613c5357610763336179a9565b807f118cdaa7000000000000000000000000000000000000000000000000000000006024925233600452fd5b503461048857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048857613cb6616ec4565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00558073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b503461048857610763613dee613df7613dd93661651e565b9390929197948096613de961704c565b6175b0565b50929091616dbf565b6171e6565b50346104885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610488576020613e3c6126306105ac616141565b604051908152f35b5034610c055760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055760043567ffffffffffffffff8111610c0557613e949036906004016164ed565b9190613e9e616164565b90613ea761704c565b8315610cbe57613eb684616dbf565b613ebf85616dbf565b915f73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416905b878110613fc75750803b15610c05575f60405180927fb80f55c900000000000000000000000000000000000000000000000000000000825260206004830152818381613f4c8d60248301908a616e1e565b03925af18015610c0957613fb2575b50613f65866166b5565b90613f73604051928361627c565b868252602082019660051b810190368211613fae57965b818810613f9e5750506107639495506171e6565b8735815260209788019701613f8a565b8680fd5b613fbf9195505f9061627c565b5f935f613f5b565b613fd2818985616e0e565b356040517f6352211e000000000000000000000000000000000000000000000000000000008152816004820152602081602481875afa908115610c09575f91614160575b5073ffffffffffffffffffffffffffffffffffffffff3391160361413157604051907f74acd209000000000000000000000000000000000000000000000000000000008252806004830152606082602481875afa918215610c09575f926140dc575b506040820151908142116140ae575050906020826001935161409a8489616776565b5201516140a78288616776565b5201613efb565b7f431e9b4d000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b9091506060813d8211614129575b816140f76060938361627c565b81010312610c0557604080519161410d8361620c565b805183526020810151602084015201516040820152905f614078565b3d91506140ea565b7f6d3d1858000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b90506020813d82116141a6575b8161417a6020938361627c565b81010312610c05575173ffffffffffffffffffffffffffffffffffffffff81168103610c05575f614016565b3d915061416d565b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0054604051908152f35b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576020614244600435616e5b565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557602073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d055416604051908152f35b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055773ffffffffffffffffffffffffffffffffffffffff61433081614322616141565b61432a616ec4565b16616f30565b1680156143bc576143ba9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d025416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0255565b005b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954604051908152f35b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557602060405160018152f35b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576145106106bc616141565b73ffffffffffffffffffffffffffffffffffffffff8116156143bc576143ba9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d065416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0655565b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036146445760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c055761467a366163e6565b9061468361704c565b8115610cbe576146ea8173ffffffffffffffffffffffffffffffffffffffff612afd825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1660205261471f8260405f2054616716565b61474f7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854828482821015616bc0565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d166020528060405f20556147a7837f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a54616716565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a556147d28261709f565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d02546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529190602090839060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af1918215610c09577fade5466942956bce7111329d1a335dcee9b48002f571248262c46dffb4793790926148c4575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0a5460408051958652602086019290925290840152339280606081015b0390a3005b6148dc9060203d60201161207057612063818361627c565b614881565b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557614917616c9d565b50610b6e614923616d0f565b604051918291829190916020806040830194805184520151910152565b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576004357f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02548110156149b5576149a6602091616d69565b90549060031b1c604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f525f60045260245260445ffd5b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557614a16616141565b60243567ffffffffffffffff8111610c0557614a369036906004016164a9565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115614ca1575b5061464457614a85616ec4565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181614c6d575b50614b0557837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc859203614c425750813b15614c1757807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115614be6575f808360206143ba95519101845af4614be0616f98565b91617ccc565b505034614bef57005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011614c99575b81614c896020938361627c565b81010312610c0557519085614ad4565b3d9150614c7c565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141583614a78565b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557614d19616c9d565b50610b6e614923616cb5565b34610c0557614d333661641a565b91614d3c61704c565b614d48611dc782616e5b565b8115610cbe57614d5783616f30565b9273ffffffffffffffffffffffffffffffffffffffff8416156143bc57614f0590825f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052602084614dba8160405f2054611e5a81838a82821115616c62565b96614df17f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954898882821080159061208257616bc0565b855f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1583528760405f2055614e48827f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5461667b565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b55614e738661709f565b73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416905f6040518097819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115610c09577f1e9c12c534d62150c4f5b916807e2bd6619095393fda9c833471a1feeefad6939273ffffffffffffffffffffffffffffffffffffffff92614f99575b50614f947f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5460405193849316973397846040919493926060820195825260208201520152565b0390a4005b614fb19060203d60201161207057612063818361627c565b614f4d565b34610c0557614fc436616303565b5050505f604051614fd660208261627c565b527f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055773ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054163314801561513c575b6150799033906169d5565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff811615615114577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b5061507973ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0654163314905061506e565b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055773ffffffffffffffffffffffffffffffffffffffff6151cd616141565b6151d5616ec4565b1680156143bc5773ffffffffffffffffffffffffffffffffffffffff6112b26143ba92616f30565b34610c055761520b366163e6565b9061521461704c565b8115610cbe5761527b8173ffffffffffffffffffffffffffffffffffffffff612afd825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d156020526152b08260405f2054616716565b6152e07f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0954828482821015616bc0565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d156020528060405f2055615338837f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54616716565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b556153638261709f565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d01546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529190602090839060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af1918215610c09577f8eee9517ef29752d27e4ce269cc66c1b8b5b861cb46f1eb91cd5980a8268a98492615454575b507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b5460408051958652602086019290925290840152339280606081016148bf565b61546c9060203d60201161207057612063818361627c565b615412565b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576154a8616a76565b506004355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d19602052610b6e60405f206009604051916154e8836161c3565b805483526154f860018201616b5b565b6020840152600581015460408401526006810154606084015260078101546080840152600881015460a0840152015460c082015260405191829182616375565b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0854604051908152f35b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055773ffffffffffffffffffffffffffffffffffffffff6155e281614322616141565b1680156143bc576143ba9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d015416177f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0155565b34610c055760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576156a3616141565b73ffffffffffffffffffffffffffffffffffffffff602435916156c861263082616f30565b83101561570f57165f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20905f52602052602060405f2054604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f521660045260245260445ffd5b34610c055760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557615775616141565b73ffffffffffffffffffffffffffffffffffffffff615799602435926105ac616ec4565b169081156143bc578015610cbe5747918282116157ba576143ba9250616fc7565b507fda0a959a000000000000000000000000000000000000000000000000000000005f525f60045260245260445260645ffd5b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055760207f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0b54604051908152f35b34610c05576020615857366162bd565b61585f616ec4565b80517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5501517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d55005b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576004355f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d15602052602060405f2054604051908152f35b34610c055761591e36616303565b5050507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c055760207f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254604051908152f35b6101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576004356101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610c055760405190615a09826161c3565b602435825260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc360112610c0557600990604051615a4781616244565b60443581526064356020820152608435604082015260a4356060820152602084019081526040840160c43581526060850160e43581526080860191610104358352606060a088019461012435865260c089019661014435885273ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015615c14575b615aec9033906169d5565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d145490615b18826166e9565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1455615b4682821515616a1f565b815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a60205260405f20555f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1960205260405f209851895551805160018a0155602081015160028a0155604081015160038a0155015160048801555160058701555160068601555160078501555160088401555191015534615be657005b6040513481527f94fbd13b6130ebb7a0b79107760c885ffbfb3304dc77321ee295b01ac90ebe6a60203392a2005b50615aec73ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d06541633149050615ae1565b34610c05576020615c69366162bd565b615c71616ec4565b80517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5501517f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f55005b34610c0557615cd5615ccc36616187565b9291909161678a565b909160405191606083019360608452825180955260808401602060808760051b8701019401905f905b878210615d1957505050839450602084015260408301520390f35b9091947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80878203018252855190602080835192838152019201905f905b808210615d73575050506020806001929701920192019091615cfe565b909192602060a060019260808751805183528481015185840152604081015160408401526060810151606084015201516080820152019401920190615d56565b34610c055760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557615dea616141565b507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557600435615e4e81616e5b565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b34610c05575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c05576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054615ef78161662a565b8084529060018116908115615f9c5750600114615f1f575b610b6e8361372d8185038261627c565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b808210615f825750909150810160200161372d615f0f565b919260018160209254838588010152019101909291615f6a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b8401909101915061372d9050615f0f565b34610c055760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610c0557600435907fffffffff000000000000000000000000000000000000000000000000000000008216809203610c0557817f780e9d630000000000000000000000000000000000000000000000000000000060209314908115616070575b5015158152f35b7f80ac58cd000000000000000000000000000000000000000000000000000000008114915081156160d4575b81156160aa575b5083616069565b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836160a3565b7f5b5e139f000000000000000000000000000000000000000000000000000000008114915061609c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610c0557565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203610c0557565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6080910112610c055760043590602435906044359060643590565b60e0810190811067ffffffffffffffff8211176161df57604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6060810190811067ffffffffffffffff8211176161df57604052565b6040810190811067ffffffffffffffff8211176161df57604052565b6080810190811067ffffffffffffffff8211176161df57604052565b60a0810190811067ffffffffffffffff8211176161df57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176161df57604052565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112610c05576040516162f381616228565b6004358152602435602082015290565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6060910112610c055760043573ffffffffffffffffffffffffffffffffffffffff81168103610c05579060243573ffffffffffffffffffffffffffffffffffffffff81168103610c05579060443590565b91909161012060c0610140830194805184526163b760208201516020860190606080918051845260208101516020850152604081015160408501520151910152565b604081015160a0850152606081015182850152608081015160e085015260a08101516101008501520151910152565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112610c05576004359060243590565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6060910112610c0557600435906024359060443573ffffffffffffffffffffffffffffffffffffffff81168103610c055790565b67ffffffffffffffff81116161df57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f82011215610c05576020813591016164c38261646f565b926164d1604051948561627c565b82845282820111610c0557815f92602092838601378301015290565b9181601f84011215610c055782359167ffffffffffffffff8311610c05576020808501948460051b010111610c0557565b9060807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc830112610c055760043567ffffffffffffffff8111610c055782616568916004016164ed565b929092916024359167ffffffffffffffff8311610c055780602384011215610c055782600401359267ffffffffffffffff8411610c05578160246060860283010111610c0557602401929160443573ffffffffffffffffffffffffffffffffffffffff81168103610c0557916064359067ffffffffffffffff8211610c05576165f3916004016164ed565b9091565b90602080835192838152019201905f5b8181106166145750505090565b8251845260209384019390920191600101616607565b90600182811c92168015616671575b602083101461664457565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691616639565b9190820391821161668857565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff81116161df5760051b60200190565b5f5b8281106166db57505050565b6060828201526020016166cf565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146166885760010190565b9190820180921161668857565b805482101561673c575f52600560205f20910201905f90565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b80511561673c5760200190565b805182101561673c5760209160051b010190565b9293919093616799848661667b565b6167e97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06167df6167c9846166b5565b936167d7604051958661627c565b8085526166b5565b01602083016166cd565b9184905b868210806169cc575b156169c057815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1860205260405f2090815493848210156169b25761683c8183616716565b928584116169aa575b61684f838561667b565b92616859846166b5565b91616867604051938461627c565b8483527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0616894866166b5565b015f5b8181106169755750505f5b85811061690b57505050906168d16168d793926168bf8b8861667b565b906168ca828b616776565b5288616776565b5061667b565b9283159081616901575b506168f757506168f15f916166e9565b906167ed565b9450945050929190565b905081105f6168e1565b8061692161691b60019386616716565b84616723565b5060046040519161693183616260565b805483528481015460208401526002810154604084015260038101546060840152015460808201526169638287616776565b5261696e8186616776565b50016168a2565b60209060405161698481616260565b5f81525f838201525f60408201525f60608201525f608082015282828801015201616897565b859350616845565b935050506168f15f916166e9565b50929350509250905f90565b508215156167f6565b156169dd5750565b73ffffffffffffffffffffffffffffffffffffffff907f8e4a23d6000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b15616a275750565b7f01707643000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b60405190616a5f82616244565b5f6060838281528260208201528260408201520152565b60405190616a83826161c3565b5f60c083828152616a92616a52565b60208201528260408201528260608201528260808201528260a08201520152565b60405190616ac082616244565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d105482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d115460208301527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d125460408301527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d13546060830152565b90604051616b6881616244565b6060600382948054845260018101546020850152600281015460408501520154910152565b15616b955750565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b15616bca57505050565b7f0877c890000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b90816020910312610c0557518015158103610c055790565b15616c1c575050565b73ffffffffffffffffffffffffffffffffffffffff92507f3a0302a9000000000000000000000000000000000000000000000000000000005f526004521660245260445ffd5b15616c6c57505050565b7fda0a959a000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b60405190616caa82616228565b5f6020838281520152565b60405190616cc282616228565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0e5482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0f546020830152565b60405190616d1c82616228565b7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0c5482527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0d546020830152565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025481101561673c577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025f5260205f2001905f90565b90616dc9826166b5565b616dd6604051918261627c565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0616e0482946166b5565b0190602036910137565b919081101561673c5760051b0190565b90918281527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311610c055760209260051b809284830137010190565b616ea3815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff821615616b95575090565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054163303616f0457565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014616f84575b15616f8057616f7290617c53565b90616f7b575090565b905090565b5090565b505067ffffffffffffffff81166001616f64565b3d15616fc2573d90616fa98261646f565b91616fb7604051938461627c565b82523d5f602084013e565b606090565b814710617024575f80809373ffffffffffffffffffffffffffffffffffffffff8294165af1616ff4616f98565b5015616ffc57565b7f3204506f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f356680b7000000000000000000000000000000000000000000000000000000005f5260045ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300541661707757565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1860205260405f2090805f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1660205260405f205491815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1560205260405f20547f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1454906040519361715485616260565b845260208401948552604084019081526060840191825260808401924284528054680100000000000000008110156161df5761719591600182018155616723565b9590956171ba5760049451865551600186015551600285015551600384015551910155565b7f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b929173ffffffffffffffffffffffffffffffffffffffff61720684616f30565b16156143bc575f915f948151925f5b84811061741457505050505073ffffffffffffffffffffffffffffffffffffffff7f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d0154166040517f70a08231000000000000000000000000000000000000000000000000000000008152336004820152602081602481855afa908115610c09575f916173e2575b508281106173af5750834710479061738057507f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d05546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9091166024820152604481019290925290929190602090849060649082905f905af1908115610c09576173619373ffffffffffffffffffffffffffffffffffffffff92617363575b5016616fc7565b565b61737b9060203d60201161207057612063818361627c565b61735a565b847f1acd938c000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b827f7520e30b000000000000000000000000000000000000000000000000000000005f523360045260245260445260645ffd5b90506020813d60201161740c575b816173fd6020938361627c565b81010312610c0557515f61729c565b3d91506173f0565b6174218183999799616776565b519061742d8186616776565b51917812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f2281108302156174d95761747d81617477600195670de0b6b3a76400008302908082049106151501809d616716565b99616716565b996174888387616776565b5191604051918252602082015273ffffffffffffffffffffffffffffffffffffffff8a16907f4dfd983654825b97cb5a6edb928f7e2a2726bef374ec4be8b39e4c31660103bb60403392a401617215565b637c5f487d5f526004601cfd5b73ffffffffffffffffffffffffffffffffffffffff81161561754d576175499073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b5490565b7f89c62b64000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b15617582575050565b7f7c8345c2000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b959291939094958515610cbe578215610cbe578615610cbe576175d68684818114617579565b6175e38688818114617579565b6175ec86616dbf565b966175f687616dbf565b9661760081616dbf565b965f5b8281106176135750505050505050565b8681101561673c5760608102820161762c828588616e0e565b3581359061765b33833373ffffffffffffffffffffffffffffffffffffffff61765483616e5b565b1614616c13565b6020830135918215610cbe576040840135938415610cbe57815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1760205260405f20835f5260205260ff60405f20541661797a57825f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1a602052606060405f2054916176eb85841515616a1f565b360312610c05576040516176fe8161620c565b82815260208101858152604082018781526040519160208301938785525160408401525160608301525160808201526080815261773c60a08261627c565b519020604051602081019182526020815261775860408261627c565b519020908987101561673c578660051b8c01357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18d360301811215610c05578c0180359067ffffffffffffffff8211610c0557602001918160051b36038313610c0557835f5b8381106179405750036178eb57505050815f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1960205261785860405f2060405190617809826161c3565b8054825261781960018201616b5b565b602083015260058101549182604082015260c0600960068401549384606085015260078101546080850152600881015460a08501520154910152616716565b804211610c8f57948f819594938f93600198617877856178d997616776565b525f527f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d1760205260405f20905f5260205260405f20877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055616776565b526178e4828d616776565b5201617603565b617936929395506040519586957f10f7f5f900000000000000000000000000000000000000000000000000000000875260048701526024860152608060448601526084850191616e1e565b9060648301520390fd5b9061794c828587616e0e565b359081811015617969575f52602052600160405f205b91016177be565b905f52602052600160405f20617962565b507f6bd4745f000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b909283515f905f905b808210617c06575050478111904791610c145750509173ffffffffffffffffffffffffffffffffffffffff5f610ae39593617b638296610b13857f2e7f9cff060131e699363fd9158d6852fb046fdea4328ed8ea7b3f4e3f902d035416966040519a8b998a9889977f1d05159d0000000000000000000000000000000000000000000000000000000089521660048801526080602488015260848701906165f7565b03925af1908115610c09575f91617b78575090565b90503d805f833e617b89818361627c565b810190602081830312610c055780519067ffffffffffffffff8211610c0557019080601f83011215610c05578151617bc0816166b5565b92617bce604051948561627c565b81845260208085019260051b820101928311610c0557602001905b828210617bf65750505090565b8151815260209182019101617be9565b9091617c128388616776565b5190617c1e8488616776565b5191617c2a8587616776565b519015610cbe578215610cbe57804211610c8f5750600191617c4b91616716565b920190617ac1565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff1603617cc1575b82158015617cb6575b617cae57565b5f9250829150565b5060163d1415617ca8565b5f9250829150617c9f565b90617d095750805115617ce157805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580617d5c575b617d1a575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15617d12565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615617d9457565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffdfea2646970667358221220f4c82eead3e561489529d9995bef79e0bbf2954df1d144069aa18cfee665ead564736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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
