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

// SignedCredential is an auto generated low-level Go binding around an user-defined struct.
type SignedCredential_duplicate4 struct {
	Vc VerifiableCredential
	V  uint8
	R  [32]byte
	S  [32]byte
}

// TierInfo is an auto generated low-level Go binding around an user-defined struct.
type TierInfo struct {
	CashBackPremium        *big.Int
	TokenLockAmount        *big.Int
	DebtToLiquidationValue *big.Int
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_duplicate4 struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// SPPlusMetaData contains all meta data concerning the SPPlus contract.
var SPPlusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existingTokenId\",\"type\":\"uint256\"}],\"name\":\"AgentAlreadyHasToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"AgentChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AlreadyActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notBeneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiaryIsNotAnAgent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"BeneficiaryOwnerIsNotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CredentialUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientVaultBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidDowngrade\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidUpgrade\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"MaxCashBackPercentExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoCashBackToClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPriceSetterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dtl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"OverLimitDtl\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenNotActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroCashBackPremium\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroConversionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"CardActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penaltyAmount\",\"type\":\"uint256\"}],\"name\":\"CardDowngraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"CardUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CashBackClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FilVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"GlfVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GlfVaultWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"glfBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filEarned\",\"type\":\"uint256\"}],\"name\":\"PaymentProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCashBack\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCashBack\",\"type\":\"uint256\"}],\"name\":\"PersonalCashBackUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPriceSetter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPriceSetter\",\"type\":\"address\"}],\"name\":\"PriceSetterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentFactory\",\"outputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToCardOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentPoliceVcVerifier\",\"outputs\":[{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseConversionRateFILtoGLF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"changeOwnerForAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimCashBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credential\",\"type\":\"bytes32\"}],\"name\":\"credentialHasBeenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"credentialUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"_sc\",\"type\":\"tuple\"}],\"name\":\"downgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundFilVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fundGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"fundGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getTimeUntilPenaltyFreeDowngrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glfToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"},{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"_agentPoliceVcVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"_poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceSetter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_personalCashBackPercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintActivateAndFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"mintAndActivate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"onDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestAmount\",\"type\":\"uint256\"}],\"name\":\"onPaymentMade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"}],\"name\":\"setAgentFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"_agentPoliceVcVerifier\",\"type\":\"address\"}],\"name\":\"setAgentPolice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseConversionRateFILtoGLF\",\"type\":\"uint256\"}],\"name\":\"setBaseConversionRateFILtoGLF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"}],\"name\":\"setGlfToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxCashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setMaxCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setPersonalCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"setPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceSetter\",\"type\":\"address\"}],\"name\":\"setPriceSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"_tierInfo\",\"type\":\"tuple\"}],\"name\":\"setTierInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyFee\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyWindow\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tierInfoToWithdrawableExtraLockedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawGlf\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"tierToTierInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToAgentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToFilCashbackEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToGlfVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToLastTierSwitchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToPersonalCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToTier\",\"outputs\":[{\"internalType\":\"enumTier\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToTierLockAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFilCashbackVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawExtraLockedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFilFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdrawGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f51602061847a5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b6040516183b390816100c7823960805181818161408e015261420e0152f35b6001600160401b0319166001600160401b039081175f51602061847a5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f358060e01c91826301ffc9a7146164215750816302adafe1146163ba5781630431a3251461634957816306fdde0314616205578163081812fc1461617e578163095ea7b31461611e5781630c1cd643146160c057816314b77f2214615f8e5781631755ff2114615f1e57816318de0c6a14615eb75781631eccd62814615c565781631f2218de14615af25781631fac65d3146150e0575080631fe29f181461508557806323b872dd1461504b57806328d0e5a914614b9d5780632c09bef714614b2c5780632eacc69614614a2e5780632fe5c820146149c75780633f4ba83a146148cc57806342842e0e1461488257806343989f0a146148195780634449ee09146147b2578063463ad343146144d65780634f1ef286146141a35780634fa56e3b1461413257806351c9749c1461410657806352d1902d1461404857806354fd4d501461400e5780635c975abb14613fae5780635e7277f314613f4757806361d027b314613ed65780636352211e14613e9957806363791e3c14613d5e5780636817c76c14613d035780636a62784214613cc15780636b12bee314613c6657806370a0823114613b90578063715018a614613a4e578063762b977d1461397357806379ba5097146138d15780637caf9073146137f65780637df107ea146137855780637dfe5b92146133d457806381eebd70146133745780638456cb591461329f57806386488a85146122f557806386cc34371461221a5780638b446431146121e75780638da5cb5b146121765780638dd404d21461211b5780638f68679f146120c0578063912934731461200b5780639202adfc14611ea557806395bda1d814611e6857806395d89b4114611d0b5780639780b37214611ca4578063a1e80e5514611c49578063a22cb46514611bdd578063abdf23df14611b4b578063ad3cb1cc14611acc578063ad7f492914611a71578063aef0353c146114aa578063b3672856146113f6578063b3d0d7b7146113b9578063b77da5331461131b578063b88d4fde14611291578063c87b56dd14611218578063cda43c1014611198578063d0911b1e14611127578063d86e1a3514610ff3578063e30c397814610f82578063e38a491114610f29578063e985e9c514610e78578063e9e15b4f14610d8f578063ebca88fc14610ab1578063f0f44260146109a4578063f13c2c1c1461068e578063f173e1fe14610627578063f2fde38b14610502578063f4a0a528146104a25763f9d20f53146103a6575f80fd5b3461049f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57600435600481101561049d5760607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261049d578080158080159283610488575b61042191616a4c565b610429617aaa565b9061047d575b156104555761043d90616a85565b60243581556044356001820155600260643591015580f35b6004827f1db5b8a7000000000000000000000000000000000000000000000000000000008152fd5b50602435151561042f565b610421915061049681617c03565b9150610418565b505b80fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576104da617aaa565b6004357fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d005580f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5773ffffffffffffffffffffffffffffffffffffffff6105576105526165a7565b617b69565b61055f617aaa565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57604060209160043581527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e83522054604051908152f35b503461049f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57600435906106ca6165ed565b60443560048110156109a0576106e390606435926174e8565b916106ec617b16565b6106fe836106f981617bd1565b616a19565b8115610978577fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015230602482015260448101849052949060209086906064908290869073ffffffffffffffffffffffffffffffffffffffff165af194851561096d57602095610942575b508382527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d138552604082206107ca848254616fcd565b90558382527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1085526040808320548584527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1387529220548290843373ffffffffffffffffffffffffffffffffffffffff61084389617a16565b16036108f857831591826108ee575b5050806108e6575b156108c7575050507fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c5461088e8184617cf3565b60405191825283820152817f57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c660403393a3604051908152f35b81036108d4575b5061088e565b90506108e08184617cf3565b5f6108ce565b50801561085a565b149050845f610852565b14915081610939575b501561088e57507fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c546109348184617cf3565b61088e565b9050155f610901565b61096190863d8811610966575b6109598183616720565b81019061687f565b610794565b503d61094f565b6040513d84823e3d90fd5b807f1f2a20050000000000000000000000000000000000000000000000000000000060049252fd5b8280fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576109e76109df6165a7565b610552617aaa565b73ffffffffffffffffffffffffffffffffffffffff811615610a8957610a869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d035416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0355565b80f35b6004827fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602435600435604435610af2617b16565b610aff826106f981617bd1565b8215610d67577fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529060209082906064908290899073ffffffffffffffffffffffffffffffffffffffff165af18015610d5c57610d3f575b508184527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1360205260408420610bc7848254616fcd565b90558184527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d106020526040842054908285527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d136020526040852054918092853373ffffffffffffffffffffffffffffffffffffffff610c4588617a16565b1603610cf55783159182610ceb575b505080610ce3575b15610cc4575050507fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c54610c908183617cf3565b60405192835260208301527f57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c660403393a380f35b8103610cd1575b50610c90565b9050610cdd8183617cf3565b5f610ccb565b508015610c5c565b149050855f610c54565b14915081610d36575b5015610c9057507fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c54610d318183617cf3565b610c90565b9050155f610cfe565b610d579060203d602011610966576109598183616720565b610b90565b6040513d87823e3d90fd5b6004847f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5773ffffffffffffffffffffffffffffffffffffffff610dee81610de06165a7565b610de8617aaa565b16617b69565b168015610a8957610a869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d045416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0455565b503461049f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57610eb06165a7565b73ffffffffffffffffffffffffffffffffffffffff610ecd6165ed565b91165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b503461049f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57610f616165a7565b60243591600483101561049f576020610f7a84846174e8565b604051908152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760043573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416331480156110e7575b156110bf578015611097577fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0f5580f35b6004827f247af9ce000000000000000000000000000000000000000000000000000000008152fd5b6004827fd6d9d129000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1254163314611067565b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d055416604051908152f35b503461049f576111a7366166a3565b9073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d04541633036111f05790610a86916170e2565b6004837f4b602735000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57611253600435617a16565b5080604051611263602082616720565b525061128d604051611276602082616720565b5f8152604051918291602083526020830190616564565b0390f35b503461049f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576112c96165a7565b506112d26165ed565b5060643567ffffffffffffffff811161049d57906112f56004923690840161679b565b507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57600435600481101561049d5761137b906060926040805161136b816166d7565b8281528260208201520152616a85565b604051611387816166d7565b815491828252604060026001830154926020850193845201549201918252604051928352516020830152516040820152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576020610f7a600435617068565b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760043580156114825760408261146492602094527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0984522054617a16565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b6004827fac8fb3c1000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576114e26165a7565b602435604435916004831015611a6d5761153d906114fe617b16565b61152b61150a84617a16565b84339173ffffffffffffffffffffffffffffffffffffffff33911614616830565b6105528461153881617c03565b616a4c565b61154682617bd1565b611a4157600161155584616a85565b01547fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905291929190602090829060649082908a9073ffffffffffffffffffffffffffffffffffffffff165af18015611a3657611a19575b508285527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e60205260408520548386527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0d6020526040862054928487527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e6020524260408820558487527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d076020526116a08660408920616b01565b8487527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d16602052604087205573ffffffffffffffffffffffffffffffffffffffff82169182611758575050508061172857506116ff6040518093616557565b7f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a380f35b836044917f0c6dc9c900000000000000000000000000000000000000000000000000000000825260045280602452fd5b90602492959391602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d025416604051958680927ffd66091e0000000000000000000000000000000000000000000000000000000082528b60048301525afa9384156119d45788946119df575b506117e690841515616b38565b6040517f8da5cb5b0000000000000000000000000000000000000000000000000000000081526020816004818a5afa9081156119d45788916119a5575b5073ffffffffffffffffffffffffffffffffffffffff33911603611975578287527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d09602052604087205480611945575015801561193c575b1561190e57507f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5918160209287527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0983528460408820558487527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0d8352604087205561190a6040518092616557565ba380f35b7f0c6dc9c9000000000000000000000000000000000000000000000000000000008652600452602452604484fd5b5080821461187b565b87604491857f04e662d6000000000000000000000000000000000000000000000000000000008352600452602452fd5b604487877f98f08d3a00000000000000000000000000000000000000000000000000000000825260045233602452fd5b6119c7915060203d6020116119cd575b6119bf8183616720565b810190616b82565b5f611823565b503d6119b5565b6040513d8a823e3d90fd5b9093506020813d602011611a11575b816119fb60209383616720565b81010312611a0d5751926117e66117d9565b5f80fd5b3d91506119ee565b611a319060203d602011610966576109598183616720565b6115e5565b6040513d88823e3d90fd5b602484837f11d5a84c000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760207fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0f54604051908152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f575061128d604051611b0d604082616720565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190616564565b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57600435611b86617aaa565b612710811015611bb5577fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0b5580f35b6004827f58d620b3000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57611c156165a7565b506024358015150361049f57807f8cd22d190000000000000000000000000000000000000000000000000000000060049252fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760207fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0b54604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57604060209160043581527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1083522054604051908152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760405190807f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015490611d6b826167df565b8085529160018116908115611e235750600114611da7575b61128d84611d9381860382616720565b604051918291602083526020830190616564565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930181527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b808210611e0957509091508101602001611d9382611d83565b919260018160209254838588010152019101909291611df0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b85019092019250611d939150839050611d83565b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576020610f7a600435616fda565b503461049f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760043573ffffffffffffffffffffffffffffffffffffffff611f19611ef86165ed565b611f00617b16565b610552611f0c85617a16565b8533918633911614616830565b16908115611fe3578083527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1460205260408320548015611fb75760207f05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700918386527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d148252856040812055611fad8186617c6e565b604051908152a380f35b602484837fa370a06b000000000000000000000000000000000000000000000000000000008252600452fd5b6004837fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b50807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5761203d617b16565b34156109785761206e347fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1554616fcd565b7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1555604051348152817f1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d607460203393a380f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760207fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0654604051908152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760207fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c54604051908152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5780f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5773ffffffffffffffffffffffffffffffffffffffff61226b81610de06165a7565b168015610a8957610a869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d055416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0555565b503461049f576101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760043567ffffffffffffffff811161049d5761234690369060040161679b565b9060243567ffffffffffffffff811161049d5761236790369060040161679b565b916123706165ca565b926084359373ffffffffffffffffffffffffffffffffffffffff85168503611a6d5760a43573ffffffffffffffffffffffffffffffffffffffff8116810361329b5760c43573ffffffffffffffffffffffffffffffffffffffff811681036132975760e4359173ffffffffffffffffffffffffffffffffffffffff8316830361329357610104359373ffffffffffffffffffffffffffffffffffffffff8516850361328f57610124359573ffffffffffffffffffffffffffffffffffffffff8716870361328b577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549767ffffffffffffffff89168015908161327b575b6001149081613271575b159081613268575b506132405760017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008a16177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005560ff8960401c16156131eb575b6124e2618076565b6124ea618076565b80519067ffffffffffffffff82116131be5781906125287f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300546167df565b601f81116130f6575b50602090601f8311600114612ffb578c92612ff0575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b80519067ffffffffffffffff8211612fc3576125d57f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301546167df565b601f8111612f41575b50602090601f8311600114612e5e576126559392918b9183612e53575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b610552618076565b61265d618076565b612665618076565b73ffffffffffffffffffffffffffffffffffffffff811615612e275761268a90617dd9565b612692618076565b61269a618076565b60643515612dff5773ffffffffffffffffffffffffffffffffffffffff811615612dd75773ffffffffffffffffffffffffffffffffffffffff881615612dd75773ffffffffffffffffffffffffffffffffffffffff821615612dd75773ffffffffffffffffffffffffffffffffffffffff831615612dd75773ffffffffffffffffffffffffffffffffffffffff841615612dd75773ffffffffffffffffffffffffffffffffffffffff851615612dd7576129be73ffffffffffffffffffffffffffffffffffffffff612a3e95610de88261293e81612abd9b9960ff9e9f826128be816105529d610de861283f839f610552856127bf81610de89b6064357fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d005516617b69565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d015416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0155565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d035416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0355565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d025416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0255565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d055416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0555565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d045416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0455565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d125416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1255565b60017fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d06556101f47fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c55661488e11ecd40007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0f556276a7007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0a556101f47fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0b55604051612b87816166d7565b8381526002602082018581526040830190670a688906bd8b000082528680527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d08602052604087209351845551600184015551910155604051612be8816166d7565b670e92596fd62900008152600260208201690a968163f0a57b40000081526040830190670b1a2bc2ec5000008252600187527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d08602052604087209351845551600184015551910155604051612c5c816166d7565b670f43fc2c04ee000081526002602082016934f086f3b33b6840000081526040830190670be4acf59c82800082528287527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d08602052604087209351845551600184015551910155604051612ccf816166d7565b671158e460913d00008152600260208201699ed194db19b238c0000081526040830190670c7d713b49da00008252600387527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0860205260408720935184555160018401555191015560401c1615612d435780f35b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6004877fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b6004877f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b6024887f1e4fbdf700000000000000000000000000000000000000000000000000000000815280600452fd5b015190505f806125fb565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793018b52818b2091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084168c5b818110612f29575091600193918561265597969410612ef2575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015561264d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f8080612ec5565b92936020600181928786015181550195019301612eab565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793018b527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f840160051c81019160208510612fb9575b601f0160051c01905b818110612fae57506125de565b8b8155600101612fa1565b9091508190612f98565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b015190505f80612547565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793008d527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168d5b8181106130de57509084600195949392106130a7575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055612599565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f808061307a565b92936020600181928786015181550195019301613064565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793008c52601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf810160208410613197575b908392915b8d601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101821061318a575050612531565b8155849350600101613153565b507f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8161314e565b60248b7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b680100000000000000017fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000008a16177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00556124da565b60048a7ff92ee8a9000000000000000000000000000000000000000000000000000000008152fd5b9050155f612480565b303b159150612478565b60408b901c60ff1615915061246e565b8880fd5b8780fd5b8680fd5b8580fd5b8480fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576132d6617aaa565b6132de617b16565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576133ac617aaa565b6004357fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c5580f35b503461049f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576004356024359060048210156109a05761341c617b16565b61344961342882617a16565b82339173ffffffffffffffffffffffffffffffffffffffff33911614616830565b613456816106f981617bd1565b6134638261153881617c03565b8083527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0760205260ff604084205416916004831015613758578281111561371e5783908282527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d16602052604082205460016134dd83616a85565b01549073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d015416908083115f1461367d5760209161353261357b9285616897565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019190915295869283919082906064820190565b03925af1928315611a36576136359361365e575b505b8386527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d076020526135c58260408820616b01565b8386527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e6020524260408720558386527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d16602052604086205561362b6040518095616557565b6020840190616557565b7f7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd60403393a380f35b6136769060203d602011610966576109598183616720565b505f61358f565b808310613691575b50506136359250613591565b6020916136a1846136e493616897565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481019190915295869283919082906044820190565b03925af1928315611a3657613635936136ff575b8693613685565b6137179060203d602011610966576109598183616720565b505f6136f8565b83613756604492613751867fa0f244d100000000000000000000000000000000000000000000000000000000855261653b565b616549565bfd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d025416604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5773ffffffffffffffffffffffffffffffffffffffff61384781610de06165a7565b168015610a8957610a869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d025416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0255565b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054160361394757610a8633617dd9565b807f118cdaa7000000000000000000000000000000000000000000000000000000006024925233600452fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5773ffffffffffffffffffffffffffffffffffffffff6139c481610de06165a7565b168015610a8957610a869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d015416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0155565b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57613a85617aaa565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00558073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57613bcb6105526165a7565b9073ffffffffffffffffffffffffffffffffffffffff821615613c3a576020613c318373ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b54604051908152f35b807f89c62b64000000000000000000000000000000000000000000000000000000006024925280600452fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760207fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1554604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576020610f7a613cfe6165a7565b616bae565b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760207fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0054604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57613d996109df6165a7565b73ffffffffffffffffffffffffffffffffffffffff8116908115611fe357613e7273ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1254169173ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d125416177fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1255565b7f379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa400798380a380f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576020611464600435617a16565b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d035416604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57604060209160043581527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1383522054604051908152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602060405160018152f35b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036140de5760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b807fe07c8dba0000000000000000000000000000000000000000000000000000000060049252fd5b503461049f57610a86614118366166a3565b90614121617b16565b61412d61342882617a16565b617cf3565b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d015416604051908152f35b5060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f576141d66165a7565b9060243567ffffffffffffffff811161049d576141f790369060040161679b565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115614494575b5061446c57614246617aaa565b73ffffffffffffffffffffffffffffffffffffffff831690604051937f52d1902d000000000000000000000000000000000000000000000000000000008552602085600481865afa80958596614438575b506142c857602484847f4c9c8ce3000000000000000000000000000000000000000000000000000000008252600452fd5b9091847f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc810361440d5750813b156143e257807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8480a281518390156143af57808360206143ab95519101845af46143a5617c3f565b916182e4565b5080f35b505050346143ba5780f35b807fb398979f0000000000000000000000000000000000000000000000000000000060049252fd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000008452600452602483fd5b7faa1d49a4000000000000000000000000000000000000000000000000000000008552600452602484fd5b9095506020813d602011614464575b8161445460209383616720565b81010312611a0d5751945f614297565b3d9150614447565b6004827fe07c8dba000000000000000000000000000000000000000000000000000000008152fd5b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f614239565b503461049f576144e5366166a3565b906144ee617b16565b6144f6617b16565b614503816106f981617bd1565b811561478a577fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529060209082906064908290889073ffffffffffffffffffffffffffffffffffffffff165af1801561477f57614762575b508083527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d13602052604083206145cb838254616fcd565b90558083527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1060205260408320548184527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d136020526040842054908381923373ffffffffffffffffffffffffffffffffffffffff61464887617a16565b1603614728571480614720575b156146885750507fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c54610c908183617cf3565b15610c905750828184527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1060205260408420548285527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d10602052846040812055604051908152846020820152827fdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca60403393a3610c90565b508015614655565b149081610d36575015610c9057507fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c54610d318183617cf3565b61477a9060203d602011610966576109598183616720565b614594565b6040513d86823e3d90fd5b6004837f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57604060209160043581527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0d83522054604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57604060209160043581527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d11835220541515604051908152f35b503461049f5760049061489436616631565b505050806040516148a6602082616720565b527f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57614903617aaa565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff81161561499f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b6004827f8dfc202b000000000000000000000000000000000000000000000000000000008152fd5b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57604060209160043581527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1483522054604051908152f35b503461049f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57614a666165a7565b9073ffffffffffffffffffffffffffffffffffffffff614a8b60243593610552617aaa565b16918215610a895747811561478a57808211614af95750610a869192614ad2827fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1554616897565b7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1555617c6e565b82917fda0a959a000000000000000000000000000000000000000000000000000000006064945282600452602452604452fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d125416604051908152f35b503461049f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f57614be0614bd86165a7565b610552617b16565b9073ffffffffffffffffffffffffffffffffffffffff8216918215610a89576024602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d025416604051928380927f1ffbb0640000000000000000000000000000000000000000000000000000000082528860048301525afa9081156150005790614c879291849161502c575b50616b38565b604051917f8da5cb5b000000000000000000000000000000000000000000000000000000008352602083600481845afa92831561096d57829361500b575b5073ffffffffffffffffffffffffffffffffffffffff8316908115611fe3576020600491604051928380927faf640d0f0000000000000000000000000000000000000000000000000000000082525afa908115615000578391614fce575b5082527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0960205260408220549073ffffffffffffffffffffffffffffffffffffffff614d6e83617a16565b1693818514614fa65760209460405191614d888784616720565b85835273ffffffffffffffffffffffffffffffffffffffff614daa86836180cd565b1680614ddc57602487877f7e273289000000000000000000000000000000000000000000000000000000008252600452fd5b82819694959603614f7257503b614df1578480f35b908591614e4360405194859384937f150b7a0200000000000000000000000000000000000000000000000000000000855233600486015260248501526044840152608060648401526084830190616564565b038186865af1839181614f1a575b50614e995750614e5f617c3f565b80519384614e9357602484847f64a0ae92000000000000000000000000000000000000000000000000000000008252600452fd5b84925001fd5b7f150b7a0200000000000000000000000000000000000000000000000000000000919293507fffffffff000000000000000000000000000000000000000000000000000000001603614eef5750805f8080808480f35b7f64a0ae92000000000000000000000000000000000000000000000000000000008252600452602490fd5b9091508481813d8311614f6b575b614f328183616720565b81010312611a6d57517fffffffff0000000000000000000000000000000000000000000000000000000081168103611a6d57905f614e51565b503d614f28565b8660649185857f64283d7b000000000000000000000000000000000000000000000000000000008452600452602452604452fd5b6004847f2a63c7cc000000000000000000000000000000000000000000000000000000008152fd5b90506020813d602011614ff8575b81614fe960209383616720565b81010312611a0d57515f614d23565b3d9150614fdc565b6040513d85823e3d90fd5b61502591935060203d6020116119cd576119bf8183616720565b915f614cc5565b615045915060203d602011610966576109598183616720565b5f614c81565b503461049f5760049061505d36616631565b5050507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b503461049f57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261049f5760207fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0a54604051908152f35b905034611a0d5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d57600435602435916004831015611a0d576044359267ffffffffffffffff8411611a0d5783600401938036039460807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc870112611a0d5761516f617b16565b61519c61517b86617a16565b86339173ffffffffffffffffffffffffffffffffffffffff33911614616830565b6151a9856106f981617bd1565b8215838115615add575b6151bc91616a4c565b845f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0760205260ff60405f205416956004871015615ab05786841015615a7b579083929161520b8995616a85565b9560405191615219836166d7565b875483526040600260018a015499602086019a8b5201549301928352885f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0d60205260405f2054948561564a575b5050505050508382527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d166020526040822054925182938584527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e6020526152d3604085205442616897565b7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d015473ffffffffffffffffffffffffffffffffffffffff1691908382101561543857506020916135326153269285616897565b03925af192831561542d576153e09361540e575b505b8487527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d076020526153708260408920616b01565b8487527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e6020524260408820558487527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1660205260408720556153d66040518096616557565b6020850190616557565b60408301527f89a80b3a2ac9f4ceb70a953f78c6572114c7041e11e2fae9f8f1e22425402c8b60603393a380f35b6154269060203d602011610966576109598183616720565b505f61533a565b6040513d89823e3d90fd5b9083819796939711615453575b5050506153e092935061533c565b836154619196929396616897565b917fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0a54115f146155d15750615552906127106154be7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0b5483616aee565b049560208773ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d03541660405195869283927fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b038189855af1928315611a365787615576936136a1926020966155b6575b50616897565b03925af192831561542d576153e093615597575b505b839250865f80615445565b6155af9060203d602011610966576109598183616720565b505f61558a565b6155cc90873d8911610966576109598183616720565b615570565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481019290925294909360209185916044918391905af192831561542d576153e09361562b575b5061558c565b6156439060203d602011610966576109598183616720565b505f615625565b6156cf96975073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0554169160206156958780616abb565b604051809a81927fa2f3c21000000000000000000000000000000000000000000000000000000000835284600484015260248301906168e6565b0381865afa978815615a14575f98615a47575b50875f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1160205260405f2054615a1f57823b15611a0d577fffffffff00000000000000000000000000000000000000000000000000000000604051947f403bb79d000000000000000000000000000000000000000000000000000000008652886004870152166024850152606060448501527ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffefd86359101811215611a0d576157bb906080606486015260048360e487019201016168e6565b9160248201359260ff8416809403611a0d575f858094926064829484986084850152604481013560a4850152013560c483015203925af18015615a14576159fd575b5061580a90519180616abb565b602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d045416936024604051809681937f9f4326d700000000000000000000000000000000000000000000000000000000835260048301525afa9283156159f25789936159be575b506101008136031261328b57604051610100810181811067ffffffffffffffff8211176131be576040526158b982616610565b8152602082013560208201526040820135604082015260608201356060820152608082013560808201526158ef60a0830161650e565b60a082015261590060c083016168d1565b60c082015260e082013567ffffffffffffffff81116159ba57615934926159299136910161679b565b60e082015283617f61565b828211615979575050505090859182527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d116020524360408320555f8080808080615268565b60849450604051937f4ec0c37d0000000000000000000000000000000000000000000000000000000085526004850152602484015260448301526064820152fd5b8a80fd5b9092506020813d6020116159ea575b816159da60209383616720565b81010312611a0d5751915f615886565b3d91506159cd565b6040513d8b823e3d90fd5b615a0a9199505f90616720565b5f9761580a6157fd565b6040513d5f823e3d90fd5b7fae9e9f65000000000000000000000000000000000000000000000000000000005f5260045ffd5b9097506020813d602011615a73575b81615a6360209383616720565b81010312611a0d5751965f6156e2565b3d9150615a56565b615aab84613751897f253471cf000000000000000000000000000000000000000000000000000000005f5261653b565b60445ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6151bc9150615aeb81617c03565b91506151b3565b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d5760043567ffffffffffffffff8111611a0d576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112611a0d576020615bdc9173ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d05541660405180809581947fa2f3c21000000000000000000000000000000000000000000000000000000000835286600484015260248301906004016168e6565b03915afa908115615a14575f91615c24575b505f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d11602052602060405f2054604051908152f35b90506020813d602011615c4e575b81615c3f60209383616720565b81010312611a0d575181615bee565b3d9150615c32565b34611a0d5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d57600435602435615c936165ca565b615c9b617b16565b615ca761150a84617a16565b8115615e8f57615cb690617b69565b9173ffffffffffffffffffffffffffffffffffffffff8316928315615e6757815f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1360205260405f2054808411615e345750602083615dd992845f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d13835260405f20615d45838254616897565b905573ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d015416905f6040518096819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af18015615a1457615e17575b506040519182527f36ae8a1f61e4ec367d806b7c73a9031431dd6fae6810392041e8791c1ddba34f60203393a4005b615e2f9060203d602011610966576109598183616720565b615de8565b83837fda0a959a000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d576004355f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d09602052602060405f2054604051908152f35b34611a0d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d57602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d045416604051908152f35b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d57600435615fc8617b16565b615fd461342882617a16565b615fdd81617068565b908115615e8f577fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081523360048201526024810184905290602090829060449082905f9073ffffffffffffffffffffffffffffffffffffffff165af18015615a14576160a3575b505f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1660205261609f60405f20918254616897565b9055005b6160bb9060203d602011610966576109598183616720565b616069565b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d576160f7617aaa565b6004357fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0a55005b34611a0d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d576161556165a7565b507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d576004356161b981617a16565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b34611a0d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054616262816167df565b8084529060018116908115616307575060011461628a575b61128d83611d9381850382616720565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b8082106162ed57509091508101602001611d9361627a565b9192600181602092548385880101520191019092916162d5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b84019091019150611d93905061627a565b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d576004355f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d07602052602060ff60405f2054166163b86040518092616557565bf35b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d576004355f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d16602052602060405f2054604051908152f35b34611a0d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611a0d57600435907fffffffff000000000000000000000000000000000000000000000000000000008216809203611a0d57817f80ac58cd00000000000000000000000000000000000000000000000000000000602093149081156164e4575b81156164ba575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836164b3565b7f5b5e139f00000000000000000000000000000000000000000000000000000000811491506164ac565b35907fffffffff0000000000000000000000000000000000000000000000000000000082168203611a0d57565b6004811015615ab057600452565b6004811015615ab057602452565b906004821015615ab05752565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203611a0d57565b6044359073ffffffffffffffffffffffffffffffffffffffff82168203611a0d57565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203611a0d57565b359073ffffffffffffffffffffffffffffffffffffffff82168203611a0d57565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6060910112611a0d5760043573ffffffffffffffffffffffffffffffffffffffff81168103611a0d579060243573ffffffffffffffffffffffffffffffffffffffff81168103611a0d579060443590565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112611a0d576004359060243590565b6060810190811067ffffffffffffffff8211176166f357604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176166f357604052565b67ffffffffffffffff81116166f357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f82011215611a0d576020813591016167b582616761565b926167c36040519485616720565b82845282820111611a0d57815f92602092838601378301015290565b90600182811c92168015616826575b60208310146167f957565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f16916167ee565b15616839575050565b73ffffffffffffffffffffffffffffffffffffffff92507f3a0302a9000000000000000000000000000000000000000000000000000000005f526004521660245260445ffd5b90816020910312611a0d57518015158103611a0d5790565b919082039182116168a457565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b359067ffffffffffffffff82168203611a0d57565b73ffffffffffffffffffffffffffffffffffffffff61690482616610565b168252602081013560208301526040810135604083015260608101356060830152608081013560808301527fffffffff0000000000000000000000000000000000000000000000000000000061695c60a0830161650e565b1660a083015267ffffffffffffffff61697760c083016168d1565b1660c083015260e08101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215611a0d5701906020823592019167ffffffffffffffff8111611a0d578036038313611a0d57601f817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0926101209561010060e087015281610100870152868601375f8582860101520116010190565b15616a215750565b7fed5a980a000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b15616a545750565b616a80907fbca1a956000000000000000000000000000000000000000000000000000000005f5261653b565b60245ffd5b6004811015615ab0575f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0860205260405f2090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215611a0d570190565b818102929181159184041417156168a457565b906004811015615ab05760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b15616b405750565b73ffffffffffffffffffffffffffffffffffffffff907f3a35a1b9000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b90816020910312611a0d575173ffffffffffffffffffffffffffffffffffffffff81168103611a0d5790565b616bba90610552617b16565b7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82146168a457600182017fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d06557fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01547fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d03547fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d00546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9283166024820152604481019190915291602091839160649183915f91165af18015615a1457616fb0575b506020604051616d0c8282616720565b5f815273ffffffffffffffffffffffffffffffffffffffff831615616f845773ffffffffffffffffffffffffffffffffffffffff616d4a85856180cd565b16616f5857823b616da6575b506040805133815273ffffffffffffffffffffffffffffffffffffffff9093169183019190915282917f864b3215ee93b2a9291e17b45d665127a426c1259666cf483c30939c4a7103329190a290565b9391616e2a8273ffffffffffffffffffffffffffffffffffffffff839694961696604051809381927f150b7a0200000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff331660048401525f6024840152876044840152608060648401526084830190616564565b03815f8a5af15f9181616f00575b50616e7e578585616e47617c3f565b80519182616e7b57837f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b01fd5b7fffffffff000000000000000000000000000000000000000000000000000000007f150b7a020000000000000000000000000000000000000000000000000000000091969294969593951603616ed5575082616d56565b7f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091508581813d8311616f51575b616f188183616720565b81010312611a0d57517fffffffff0000000000000000000000000000000000000000000000000000000081168103611a0d57905f616e38565b503d616f0e565b7f73c6ac6e000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b7f64a0ae92000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b616fc89060203d602011610966576109598183616720565b616cfc565b919082018092116168a457565b616fe381617a16565b50616fed81617bd1565b15617063575f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e60205261702660405f205442616897565b7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0a54908181101561705d5761705a91616897565b90565b50505f90565b505f90565b805f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0760205260016170a160ff60405f205416616a85565b0154905f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1660205260405f2054908082115f1461705d5761705a91616897565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166174e45781156174e457805f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0960205260405f20549182156174df5761714d83617bd1565b156174df57825f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1060205260405f20547fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c548082105f146174d75750905b81156174d157835f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d07602052670de0b6b3a76400006127106172236171f860ff60405f205416616a85565b94600260405196617208886166d7565b80548852600181015460208901520154604087015284616aee565b04927fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0f549051020490670de0b6b3a7640000830282808204910615150191855f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1360205260405f20548381106174ba575b507fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d15549084821061749a575b505081156174935761742d92855f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1360205260405f20617303848254616897565b9055617330817fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1554616897565b7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1555855f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1460205260405f20617388828254616fcd565b90557fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01547fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d03546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015260248101869052956020928792169082905f9082906044820190565b03925af1928315615a14577f299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d694606094617476575b5060405192835260208301526040820152a3565b61748e9060203d602011610966576109598183616720565b617462565b5050505050565b909350670de0b6b3a7640000840281810615159190040191505f806172c1565b670de0b6b3a764000082820204945092505f617295565b50505050565b9050906171ac565b505050565b5050565b906175366174f533616bae565b926174fe617b16565b61752933853373ffffffffffffffffffffffffffffffffffffffff61752283617a16565b1614616830565b6105528361153881617c03565b61753f83617bd1565b6179ea57600161754e83616a85565b01547fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d01546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905291929190602090829060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af18015615a14576179cd575b505f8481527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0e6020908152604080832080547fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0d845282852054429092557fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d07909352922091939091617670908690616b01565b855f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1660205260405f205573ffffffffffffffffffffffffffffffffffffffff82169182155f1461772b57505050806176fc57506176d26040518092616557565b817f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a390565b7f0c6dc9c9000000000000000000000000000000000000000000000000000000005f526004525f60245260445ffd5b90602492949391602073ffffffffffffffffffffffffffffffffffffffff7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d025416604051958680927ffd66091e0000000000000000000000000000000000000000000000000000000082528a60048301525afa938415615a14575f94617997575b506177b990841515616b38565b6040517f8da5cb5b000000000000000000000000000000000000000000000000000000008152602081600481895afa908115615a14575f91617978575b5073ffffffffffffffffffffffffffffffffffffffff3391160361794857825f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0960205260405f2054806179195750158015617910575b156178e2575060208492827f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5935f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0983528460405f2055845f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0d835260405f20556178de6040518092616557565ba390565b7f0c6dc9c9000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b5080821461784e565b837f04e662d6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b847f98f08d3a000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b617991915060203d6020116119cd576119bf8183616720565b5f6177f6565b9093506020813d6020116179c5575b816179b360209383616720565b81010312611a0d5751926177b96177ac565b3d91506179a6565b6179e59060203d602011610966576109598183616720565b6175de565b827f11d5a84c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b617a5e815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff821615617a7f575090565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054163303617aea57565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416617b4157565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014617bbd575b15617bb957617bab90617ee8565b90617bb4575090565b905090565b5090565b505067ffffffffffffffff81166001617b9d565b5f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0760205261705a60ff60405f2054165b60048110159081615ab05760018114918215617c32575b8215617c2557505090565b909150615ab05760031490565b506002811491505f617c1a565b3d15617c69573d90617c5082616761565b91617c5e6040519384616720565b82523d5f602084013e565b606090565b814710617ccb575f80809373ffffffffffffffffffffffffffffffffffffffff8294165af1617c9b617c3f565b5015617ca357565b7f3204506f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f356680b7000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d0c54808311617daa5750805f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d1060205260405f205491815f527fbb5d09f670b250da2058d07e0fbd5dd7910d75e9892adb7f53b40c0269ad6d106020528060405f205560405192835260208301527fdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca60403393a3565b827f37829ee6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff1603617f56575b82158015617f4b575b617f4357565b5f9250829150565b5060163d1415617f3d565b5f9250829150617f34565b919060e001519161014083805181010312611a0d57604051610140810181811067ffffffffffffffff8211176166f35760405260208401518152610120610140604086015195866020850152606081015160408501526080810151606085015260a0810151608085015260c081015160a085015260e081015160c085015261010081015160e085015282810151610100850152015191015282811561806e57508215618048577812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f22811083021561803b57670de0b6b3a7640000839102049190565b637c5f487d5f526004601cfd5b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9190565b9250505f9190565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c16156180a557565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b90618116815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9173ffffffffffffffffffffffffffffffffffffffff831680618221575b73ffffffffffffffffffffffffffffffffffffffff821691826181cd575b50825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f20827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a490565b6182149073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190555f618152565b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff000000000000000000000000000000000000000081541690556182b98473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8154019055618134565b9061832157508051156182f957805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580618374575b618332575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b1561832a56fea264697066735822122098a0eb5f24742192790a6831f00b89466e9b0eb4e5496c154fc2bd21133f5ef064736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// SPPlusABI is the input ABI used to generate the binding from.
// Deprecated: Use SPPlusMetaData.ABI instead.
var SPPlusABI = SPPlusMetaData.ABI

// SPPlusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SPPlusMetaData.Bin instead.
var SPPlusBin = SPPlusMetaData.Bin

// DeploySPPlus deploys a new Ethereum contract, binding an instance of SPPlus to it.
func DeploySPPlus(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SPPlus, error) {
	parsed, err := SPPlusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SPPlusBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SPPlus{SPPlusCaller: SPPlusCaller{contract: contract}, SPPlusTransactor: SPPlusTransactor{contract: contract}, SPPlusFilterer: SPPlusFilterer{contract: contract}}, nil
}

// SPPlus is an auto generated Go binding around an Ethereum contract.
type SPPlus struct {
	SPPlusCaller     // Read-only binding to the contract
	SPPlusTransactor // Write-only binding to the contract
	SPPlusFilterer   // Log filterer for contract events
}

// SPPlusCaller is an auto generated read-only Go binding around an Ethereum contract.
type SPPlusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SPPlusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SPPlusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SPPlusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SPPlusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SPPlusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SPPlusSession struct {
	Contract     *SPPlus           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SPPlusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SPPlusCallerSession struct {
	Contract *SPPlusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SPPlusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SPPlusTransactorSession struct {
	Contract     *SPPlusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SPPlusRaw is an auto generated low-level Go binding around an Ethereum contract.
type SPPlusRaw struct {
	Contract *SPPlus // Generic contract binding to access the raw methods on
}

// SPPlusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SPPlusCallerRaw struct {
	Contract *SPPlusCaller // Generic read-only contract binding to access the raw methods on
}

// SPPlusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SPPlusTransactorRaw struct {
	Contract *SPPlusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSPPlus creates a new instance of SPPlus, bound to a specific deployed contract.
func NewSPPlus(address common.Address, backend bind.ContractBackend) (*SPPlus, error) {
	contract, err := bindSPPlus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SPPlus{SPPlusCaller: SPPlusCaller{contract: contract}, SPPlusTransactor: SPPlusTransactor{contract: contract}, SPPlusFilterer: SPPlusFilterer{contract: contract}}, nil
}

// NewSPPlusCaller creates a new read-only instance of SPPlus, bound to a specific deployed contract.
func NewSPPlusCaller(address common.Address, caller bind.ContractCaller) (*SPPlusCaller, error) {
	contract, err := bindSPPlus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SPPlusCaller{contract: contract}, nil
}

// NewSPPlusTransactor creates a new write-only instance of SPPlus, bound to a specific deployed contract.
func NewSPPlusTransactor(address common.Address, transactor bind.ContractTransactor) (*SPPlusTransactor, error) {
	contract, err := bindSPPlus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SPPlusTransactor{contract: contract}, nil
}

// NewSPPlusFilterer creates a new log filterer instance of SPPlus, bound to a specific deployed contract.
func NewSPPlusFilterer(address common.Address, filterer bind.ContractFilterer) (*SPPlusFilterer, error) {
	contract, err := bindSPPlus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SPPlusFilterer{contract: contract}, nil
}

// bindSPPlus binds a generic wrapper to an already deployed contract.
func bindSPPlus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SPPlusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SPPlus *SPPlusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SPPlus.Contract.SPPlusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SPPlus *SPPlusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPPlus.Contract.SPPlusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SPPlus *SPPlusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SPPlus.Contract.SPPlusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SPPlus *SPPlusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SPPlus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SPPlus *SPPlusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPPlus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SPPlus *SPPlusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SPPlus.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SPPlus *SPPlusCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SPPlus *SPPlusSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SPPlus.Contract.UPGRADEINTERFACEVERSION(&_SPPlus.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SPPlus *SPPlusCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SPPlus.Contract.UPGRADEINTERFACEVERSION(&_SPPlus.CallOpts)
}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_SPPlus *SPPlusCaller) AgentFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "agentFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_SPPlus *SPPlusSession) AgentFactory() (common.Address, error) {
	return _SPPlus.Contract.AgentFactory(&_SPPlus.CallOpts)
}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_SPPlus *SPPlusCallerSession) AgentFactory() (common.Address, error) {
	return _SPPlus.Contract.AgentFactory(&_SPPlus.CallOpts)
}

// AgentIdToCardOwner is a free data retrieval call binding the contract method 0xb3672856.
//
// Solidity: function agentIdToCardOwner(uint256 _agentId) view returns(address owner)
func (_SPPlus *SPPlusCaller) AgentIdToCardOwner(opts *bind.CallOpts, _agentId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "agentIdToCardOwner", _agentId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentIdToCardOwner is a free data retrieval call binding the contract method 0xb3672856.
//
// Solidity: function agentIdToCardOwner(uint256 _agentId) view returns(address owner)
func (_SPPlus *SPPlusSession) AgentIdToCardOwner(_agentId *big.Int) (common.Address, error) {
	return _SPPlus.Contract.AgentIdToCardOwner(&_SPPlus.CallOpts, _agentId)
}

// AgentIdToCardOwner is a free data retrieval call binding the contract method 0xb3672856.
//
// Solidity: function agentIdToCardOwner(uint256 _agentId) view returns(address owner)
func (_SPPlus *SPPlusCallerSession) AgentIdToCardOwner(_agentId *big.Int) (common.Address, error) {
	return _SPPlus.Contract.AgentIdToCardOwner(&_SPPlus.CallOpts, _agentId)
}

// AgentIdToTokenId is a free data retrieval call binding the contract method 0x18de0c6a.
//
// Solidity: function agentIdToTokenId(uint256 agentId) view returns(uint256)
func (_SPPlus *SPPlusCaller) AgentIdToTokenId(opts *bind.CallOpts, agentId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "agentIdToTokenId", agentId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AgentIdToTokenId is a free data retrieval call binding the contract method 0x18de0c6a.
//
// Solidity: function agentIdToTokenId(uint256 agentId) view returns(uint256)
func (_SPPlus *SPPlusSession) AgentIdToTokenId(agentId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.AgentIdToTokenId(&_SPPlus.CallOpts, agentId)
}

// AgentIdToTokenId is a free data retrieval call binding the contract method 0x18de0c6a.
//
// Solidity: function agentIdToTokenId(uint256 agentId) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) AgentIdToTokenId(agentId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.AgentIdToTokenId(&_SPPlus.CallOpts, agentId)
}

// AgentPoliceVcVerifier is a free data retrieval call binding the contract method 0xd0911b1e.
//
// Solidity: function agentPoliceVcVerifier() view returns(address)
func (_SPPlus *SPPlusCaller) AgentPoliceVcVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "agentPoliceVcVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentPoliceVcVerifier is a free data retrieval call binding the contract method 0xd0911b1e.
//
// Solidity: function agentPoliceVcVerifier() view returns(address)
func (_SPPlus *SPPlusSession) AgentPoliceVcVerifier() (common.Address, error) {
	return _SPPlus.Contract.AgentPoliceVcVerifier(&_SPPlus.CallOpts)
}

// AgentPoliceVcVerifier is a free data retrieval call binding the contract method 0xd0911b1e.
//
// Solidity: function agentPoliceVcVerifier() view returns(address)
func (_SPPlus *SPPlusCallerSession) AgentPoliceVcVerifier() (common.Address, error) {
	return _SPPlus.Contract.AgentPoliceVcVerifier(&_SPPlus.CallOpts)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_SPPlus *SPPlusCaller) Approve(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) error {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "approve", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_SPPlus *SPPlusSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _SPPlus.Contract.Approve(&_SPPlus.CallOpts, arg0, arg1)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_SPPlus *SPPlusCallerSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _SPPlus.Contract.Approve(&_SPPlus.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SPPlus *SPPlusCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SPPlus *SPPlusSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SPPlus.Contract.BalanceOf(&_SPPlus.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SPPlus.Contract.BalanceOf(&_SPPlus.CallOpts, owner)
}

// BaseConversionRateFILtoGLF is a free data retrieval call binding the contract method 0xad7f4929.
//
// Solidity: function baseConversionRateFILtoGLF() view returns(uint256)
func (_SPPlus *SPPlusCaller) BaseConversionRateFILtoGLF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "baseConversionRateFILtoGLF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseConversionRateFILtoGLF is a free data retrieval call binding the contract method 0xad7f4929.
//
// Solidity: function baseConversionRateFILtoGLF() view returns(uint256)
func (_SPPlus *SPPlusSession) BaseConversionRateFILtoGLF() (*big.Int, error) {
	return _SPPlus.Contract.BaseConversionRateFILtoGLF(&_SPPlus.CallOpts)
}

// BaseConversionRateFILtoGLF is a free data retrieval call binding the contract method 0xad7f4929.
//
// Solidity: function baseConversionRateFILtoGLF() view returns(uint256)
func (_SPPlus *SPPlusCallerSession) BaseConversionRateFILtoGLF() (*big.Int, error) {
	return _SPPlus.Contract.BaseConversionRateFILtoGLF(&_SPPlus.CallOpts)
}

// CredentialHasBeenUsed is a free data retrieval call binding the contract method 0x43989f0a.
//
// Solidity: function credentialHasBeenUsed(bytes32 credential) view returns(bool)
func (_SPPlus *SPPlusCaller) CredentialHasBeenUsed(opts *bind.CallOpts, credential [32]byte) (bool, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "credentialHasBeenUsed", credential)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CredentialHasBeenUsed is a free data retrieval call binding the contract method 0x43989f0a.
//
// Solidity: function credentialHasBeenUsed(bytes32 credential) view returns(bool)
func (_SPPlus *SPPlusSession) CredentialHasBeenUsed(credential [32]byte) (bool, error) {
	return _SPPlus.Contract.CredentialHasBeenUsed(&_SPPlus.CallOpts, credential)
}

// CredentialHasBeenUsed is a free data retrieval call binding the contract method 0x43989f0a.
//
// Solidity: function credentialHasBeenUsed(bytes32 credential) view returns(bool)
func (_SPPlus *SPPlusCallerSession) CredentialHasBeenUsed(credential [32]byte) (bool, error) {
	return _SPPlus.Contract.CredentialHasBeenUsed(&_SPPlus.CallOpts, credential)
}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256)
func (_SPPlus *SPPlusCaller) CredentialUsed(opts *bind.CallOpts, vc VerifiableCredential) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "credentialUsed", vc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256)
func (_SPPlus *SPPlusSession) CredentialUsed(vc VerifiableCredential) (*big.Int, error) {
	return _SPPlus.Contract.CredentialUsed(&_SPPlus.CallOpts, vc)
}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) CredentialUsed(vc VerifiableCredential) (*big.Int, error) {
	return _SPPlus.Contract.CredentialUsed(&_SPPlus.CallOpts, vc)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SPPlus *SPPlusCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SPPlus *SPPlusSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SPPlus.Contract.GetApproved(&_SPPlus.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SPPlus *SPPlusCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SPPlus.Contract.GetApproved(&_SPPlus.CallOpts, tokenId)
}

// GetTimeUntilPenaltyFreeDowngrade is a free data retrieval call binding the contract method 0x95bda1d8.
//
// Solidity: function getTimeUntilPenaltyFreeDowngrade(uint256 _tokenId) view returns(uint256 timeLeft)
func (_SPPlus *SPPlusCaller) GetTimeUntilPenaltyFreeDowngrade(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "getTimeUntilPenaltyFreeDowngrade", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeUntilPenaltyFreeDowngrade is a free data retrieval call binding the contract method 0x95bda1d8.
//
// Solidity: function getTimeUntilPenaltyFreeDowngrade(uint256 _tokenId) view returns(uint256 timeLeft)
func (_SPPlus *SPPlusSession) GetTimeUntilPenaltyFreeDowngrade(_tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.GetTimeUntilPenaltyFreeDowngrade(&_SPPlus.CallOpts, _tokenId)
}

// GetTimeUntilPenaltyFreeDowngrade is a free data retrieval call binding the contract method 0x95bda1d8.
//
// Solidity: function getTimeUntilPenaltyFreeDowngrade(uint256 _tokenId) view returns(uint256 timeLeft)
func (_SPPlus *SPPlusCallerSession) GetTimeUntilPenaltyFreeDowngrade(_tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.GetTimeUntilPenaltyFreeDowngrade(&_SPPlus.CallOpts, _tokenId)
}

// GlfToken is a free data retrieval call binding the contract method 0x4fa56e3b.
//
// Solidity: function glfToken() view returns(address)
func (_SPPlus *SPPlusCaller) GlfToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "glfToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlfToken is a free data retrieval call binding the contract method 0x4fa56e3b.
//
// Solidity: function glfToken() view returns(address)
func (_SPPlus *SPPlusSession) GlfToken() (common.Address, error) {
	return _SPPlus.Contract.GlfToken(&_SPPlus.CallOpts)
}

// GlfToken is a free data retrieval call binding the contract method 0x4fa56e3b.
//
// Solidity: function glfToken() view returns(address)
func (_SPPlus *SPPlusCallerSession) GlfToken() (common.Address, error) {
	return _SPPlus.Contract.GlfToken(&_SPPlus.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SPPlus *SPPlusCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SPPlus *SPPlusSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SPPlus.Contract.IsApprovedForAll(&_SPPlus.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SPPlus *SPPlusCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SPPlus.Contract.IsApprovedForAll(&_SPPlus.CallOpts, owner, operator)
}

// MaxCashBackPercent is a free data retrieval call binding the contract method 0x8dd404d2.
//
// Solidity: function maxCashBackPercent() view returns(uint256)
func (_SPPlus *SPPlusCaller) MaxCashBackPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "maxCashBackPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCashBackPercent is a free data retrieval call binding the contract method 0x8dd404d2.
//
// Solidity: function maxCashBackPercent() view returns(uint256)
func (_SPPlus *SPPlusSession) MaxCashBackPercent() (*big.Int, error) {
	return _SPPlus.Contract.MaxCashBackPercent(&_SPPlus.CallOpts)
}

// MaxCashBackPercent is a free data retrieval call binding the contract method 0x8dd404d2.
//
// Solidity: function maxCashBackPercent() view returns(uint256)
func (_SPPlus *SPPlusCallerSession) MaxCashBackPercent() (*big.Int, error) {
	return _SPPlus.Contract.MaxCashBackPercent(&_SPPlus.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SPPlus *SPPlusCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SPPlus *SPPlusSession) MintPrice() (*big.Int, error) {
	return _SPPlus.Contract.MintPrice(&_SPPlus.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SPPlus *SPPlusCallerSession) MintPrice() (*big.Int, error) {
	return _SPPlus.Contract.MintPrice(&_SPPlus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SPPlus *SPPlusCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SPPlus *SPPlusSession) Name() (string, error) {
	return _SPPlus.Contract.Name(&_SPPlus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SPPlus *SPPlusCallerSession) Name() (string, error) {
	return _SPPlus.Contract.Name(&_SPPlus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SPPlus *SPPlusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SPPlus *SPPlusSession) Owner() (common.Address, error) {
	return _SPPlus.Contract.Owner(&_SPPlus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SPPlus *SPPlusCallerSession) Owner() (common.Address, error) {
	return _SPPlus.Contract.Owner(&_SPPlus.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SPPlus *SPPlusCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SPPlus *SPPlusSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SPPlus.Contract.OwnerOf(&_SPPlus.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SPPlus *SPPlusCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SPPlus.Contract.OwnerOf(&_SPPlus.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SPPlus *SPPlusCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SPPlus *SPPlusSession) Paused() (bool, error) {
	return _SPPlus.Contract.Paused(&_SPPlus.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SPPlus *SPPlusCallerSession) Paused() (bool, error) {
	return _SPPlus.Contract.Paused(&_SPPlus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SPPlus *SPPlusCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SPPlus *SPPlusSession) PendingOwner() (common.Address, error) {
	return _SPPlus.Contract.PendingOwner(&_SPPlus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SPPlus *SPPlusCallerSession) PendingOwner() (common.Address, error) {
	return _SPPlus.Contract.PendingOwner(&_SPPlus.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_SPPlus *SPPlusCaller) PoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "poolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_SPPlus *SPPlusSession) PoolAddress() (common.Address, error) {
	return _SPPlus.Contract.PoolAddress(&_SPPlus.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_SPPlus *SPPlusCallerSession) PoolAddress() (common.Address, error) {
	return _SPPlus.Contract.PoolAddress(&_SPPlus.CallOpts)
}

// PriceSetter is a free data retrieval call binding the contract method 0x2c09bef7.
//
// Solidity: function priceSetter() view returns(address)
func (_SPPlus *SPPlusCaller) PriceSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "priceSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceSetter is a free data retrieval call binding the contract method 0x2c09bef7.
//
// Solidity: function priceSetter() view returns(address)
func (_SPPlus *SPPlusSession) PriceSetter() (common.Address, error) {
	return _SPPlus.Contract.PriceSetter(&_SPPlus.CallOpts)
}

// PriceSetter is a free data retrieval call binding the contract method 0x2c09bef7.
//
// Solidity: function priceSetter() view returns(address)
func (_SPPlus *SPPlusCallerSession) PriceSetter() (common.Address, error) {
	return _SPPlus.Contract.PriceSetter(&_SPPlus.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SPPlus *SPPlusCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SPPlus *SPPlusSession) ProxiableUUID() ([32]byte, error) {
	return _SPPlus.Contract.ProxiableUUID(&_SPPlus.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SPPlus *SPPlusCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SPPlus.Contract.ProxiableUUID(&_SPPlus.CallOpts)
}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_SPPlus *SPPlusCaller) SafeTransferFrom0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "safeTransferFrom0", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_SPPlus *SPPlusSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _SPPlus.Contract.SafeTransferFrom0(&_SPPlus.CallOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_SPPlus *SPPlusCallerSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _SPPlus.Contract.SafeTransferFrom0(&_SPPlus.CallOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_SPPlus *SPPlusCaller) SetApprovalForAll(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "setApprovalForAll", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_SPPlus *SPPlusSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _SPPlus.Contract.SetApprovalForAll(&_SPPlus.CallOpts, arg0, arg1)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_SPPlus *SPPlusCallerSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _SPPlus.Contract.SetApprovalForAll(&_SPPlus.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SPPlus *SPPlusCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SPPlus *SPPlusSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SPPlus.Contract.SupportsInterface(&_SPPlus.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SPPlus *SPPlusCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SPPlus.Contract.SupportsInterface(&_SPPlus.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SPPlus *SPPlusCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SPPlus *SPPlusSession) Symbol() (string, error) {
	return _SPPlus.Contract.Symbol(&_SPPlus.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SPPlus *SPPlusCallerSession) Symbol() (string, error) {
	return _SPPlus.Contract.Symbol(&_SPPlus.CallOpts)
}

// TierInfoToWithdrawableExtraLockedFunds is a free data retrieval call binding the contract method 0xb3d0d7b7.
//
// Solidity: function tierInfoToWithdrawableExtraLockedFunds(uint256 _tokenId) view returns(uint256 withdrawGlf)
func (_SPPlus *SPPlusCaller) TierInfoToWithdrawableExtraLockedFunds(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tierInfoToWithdrawableExtraLockedFunds", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierInfoToWithdrawableExtraLockedFunds is a free data retrieval call binding the contract method 0xb3d0d7b7.
//
// Solidity: function tierInfoToWithdrawableExtraLockedFunds(uint256 _tokenId) view returns(uint256 withdrawGlf)
func (_SPPlus *SPPlusSession) TierInfoToWithdrawableExtraLockedFunds(_tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TierInfoToWithdrawableExtraLockedFunds(&_SPPlus.CallOpts, _tokenId)
}

// TierInfoToWithdrawableExtraLockedFunds is a free data retrieval call binding the contract method 0xb3d0d7b7.
//
// Solidity: function tierInfoToWithdrawableExtraLockedFunds(uint256 _tokenId) view returns(uint256 withdrawGlf)
func (_SPPlus *SPPlusCallerSession) TierInfoToWithdrawableExtraLockedFunds(_tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TierInfoToWithdrawableExtraLockedFunds(&_SPPlus.CallOpts, _tokenId)
}

// TierSwitchPenaltyFee is a free data retrieval call binding the contract method 0xa1e80e55.
//
// Solidity: function tierSwitchPenaltyFee() view returns(uint256)
func (_SPPlus *SPPlusCaller) TierSwitchPenaltyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tierSwitchPenaltyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierSwitchPenaltyFee is a free data retrieval call binding the contract method 0xa1e80e55.
//
// Solidity: function tierSwitchPenaltyFee() view returns(uint256)
func (_SPPlus *SPPlusSession) TierSwitchPenaltyFee() (*big.Int, error) {
	return _SPPlus.Contract.TierSwitchPenaltyFee(&_SPPlus.CallOpts)
}

// TierSwitchPenaltyFee is a free data retrieval call binding the contract method 0xa1e80e55.
//
// Solidity: function tierSwitchPenaltyFee() view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TierSwitchPenaltyFee() (*big.Int, error) {
	return _SPPlus.Contract.TierSwitchPenaltyFee(&_SPPlus.CallOpts)
}

// TierSwitchPenaltyWindow is a free data retrieval call binding the contract method 0x1fe29f18.
//
// Solidity: function tierSwitchPenaltyWindow() view returns(uint256)
func (_SPPlus *SPPlusCaller) TierSwitchPenaltyWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tierSwitchPenaltyWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierSwitchPenaltyWindow is a free data retrieval call binding the contract method 0x1fe29f18.
//
// Solidity: function tierSwitchPenaltyWindow() view returns(uint256)
func (_SPPlus *SPPlusSession) TierSwitchPenaltyWindow() (*big.Int, error) {
	return _SPPlus.Contract.TierSwitchPenaltyWindow(&_SPPlus.CallOpts)
}

// TierSwitchPenaltyWindow is a free data retrieval call binding the contract method 0x1fe29f18.
//
// Solidity: function tierSwitchPenaltyWindow() view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TierSwitchPenaltyWindow() (*big.Int, error) {
	return _SPPlus.Contract.TierSwitchPenaltyWindow(&_SPPlus.CallOpts)
}

// TierToTierInfo is a free data retrieval call binding the contract method 0xb77da533.
//
// Solidity: function tierToTierInfo(uint8 tier) view returns((uint256,uint256,uint256))
func (_SPPlus *SPPlusCaller) TierToTierInfo(opts *bind.CallOpts, tier uint8) (TierInfo, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tierToTierInfo", tier)

	if err != nil {
		return *new(TierInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TierInfo)).(*TierInfo)

	return out0, err

}

// TierToTierInfo is a free data retrieval call binding the contract method 0xb77da533.
//
// Solidity: function tierToTierInfo(uint8 tier) view returns((uint256,uint256,uint256))
func (_SPPlus *SPPlusSession) TierToTierInfo(tier uint8) (TierInfo, error) {
	return _SPPlus.Contract.TierToTierInfo(&_SPPlus.CallOpts, tier)
}

// TierToTierInfo is a free data retrieval call binding the contract method 0xb77da533.
//
// Solidity: function tierToTierInfo(uint8 tier) view returns((uint256,uint256,uint256))
func (_SPPlus *SPPlusCallerSession) TierToTierInfo(tier uint8) (TierInfo, error) {
	return _SPPlus.Contract.TierToTierInfo(&_SPPlus.CallOpts, tier)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_SPPlus *SPPlusCaller) TokenIdGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdGenerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_SPPlus *SPPlusSession) TokenIdGenerator() (*big.Int, error) {
	return _SPPlus.Contract.TokenIdGenerator(&_SPPlus.CallOpts)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TokenIdGenerator() (*big.Int, error) {
	return _SPPlus.Contract.TokenIdGenerator(&_SPPlus.CallOpts)
}

// TokenIdToAgentId is a free data retrieval call binding the contract method 0x4449ee09.
//
// Solidity: function tokenIdToAgentId(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCaller) TokenIdToAgentId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdToAgentId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToAgentId is a free data retrieval call binding the contract method 0x4449ee09.
//
// Solidity: function tokenIdToAgentId(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusSession) TokenIdToAgentId(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToAgentId(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToAgentId is a free data retrieval call binding the contract method 0x4449ee09.
//
// Solidity: function tokenIdToAgentId(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TokenIdToAgentId(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToAgentId(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToFilCashbackEarned is a free data retrieval call binding the contract method 0x2fe5c820.
//
// Solidity: function tokenIdToFilCashbackEarned(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCaller) TokenIdToFilCashbackEarned(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdToFilCashbackEarned", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToFilCashbackEarned is a free data retrieval call binding the contract method 0x2fe5c820.
//
// Solidity: function tokenIdToFilCashbackEarned(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusSession) TokenIdToFilCashbackEarned(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToFilCashbackEarned(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToFilCashbackEarned is a free data retrieval call binding the contract method 0x2fe5c820.
//
// Solidity: function tokenIdToFilCashbackEarned(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TokenIdToFilCashbackEarned(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToFilCashbackEarned(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToGlfVaultBalance is a free data retrieval call binding the contract method 0x5e7277f3.
//
// Solidity: function tokenIdToGlfVaultBalance(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCaller) TokenIdToGlfVaultBalance(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdToGlfVaultBalance", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToGlfVaultBalance is a free data retrieval call binding the contract method 0x5e7277f3.
//
// Solidity: function tokenIdToGlfVaultBalance(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusSession) TokenIdToGlfVaultBalance(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToGlfVaultBalance(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToGlfVaultBalance is a free data retrieval call binding the contract method 0x5e7277f3.
//
// Solidity: function tokenIdToGlfVaultBalance(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TokenIdToGlfVaultBalance(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToGlfVaultBalance(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToLastTierSwitchTimestamp is a free data retrieval call binding the contract method 0xf173e1fe.
//
// Solidity: function tokenIdToLastTierSwitchTimestamp(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCaller) TokenIdToLastTierSwitchTimestamp(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdToLastTierSwitchTimestamp", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToLastTierSwitchTimestamp is a free data retrieval call binding the contract method 0xf173e1fe.
//
// Solidity: function tokenIdToLastTierSwitchTimestamp(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusSession) TokenIdToLastTierSwitchTimestamp(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToLastTierSwitchTimestamp(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToLastTierSwitchTimestamp is a free data retrieval call binding the contract method 0xf173e1fe.
//
// Solidity: function tokenIdToLastTierSwitchTimestamp(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TokenIdToLastTierSwitchTimestamp(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToLastTierSwitchTimestamp(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToPersonalCashBackPercent is a free data retrieval call binding the contract method 0x9780b372.
//
// Solidity: function tokenIdToPersonalCashBackPercent(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCaller) TokenIdToPersonalCashBackPercent(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdToPersonalCashBackPercent", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToPersonalCashBackPercent is a free data retrieval call binding the contract method 0x9780b372.
//
// Solidity: function tokenIdToPersonalCashBackPercent(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusSession) TokenIdToPersonalCashBackPercent(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToPersonalCashBackPercent(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToPersonalCashBackPercent is a free data retrieval call binding the contract method 0x9780b372.
//
// Solidity: function tokenIdToPersonalCashBackPercent(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TokenIdToPersonalCashBackPercent(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToPersonalCashBackPercent(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToTier is a free data retrieval call binding the contract method 0x0431a325.
//
// Solidity: function tokenIdToTier(uint256 tokenId) view returns(uint8)
func (_SPPlus *SPPlusCaller) TokenIdToTier(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdToTier", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenIdToTier is a free data retrieval call binding the contract method 0x0431a325.
//
// Solidity: function tokenIdToTier(uint256 tokenId) view returns(uint8)
func (_SPPlus *SPPlusSession) TokenIdToTier(tokenId *big.Int) (uint8, error) {
	return _SPPlus.Contract.TokenIdToTier(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToTier is a free data retrieval call binding the contract method 0x0431a325.
//
// Solidity: function tokenIdToTier(uint256 tokenId) view returns(uint8)
func (_SPPlus *SPPlusCallerSession) TokenIdToTier(tokenId *big.Int) (uint8, error) {
	return _SPPlus.Contract.TokenIdToTier(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToTierLockAmount is a free data retrieval call binding the contract method 0x02adafe1.
//
// Solidity: function tokenIdToTierLockAmount(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCaller) TokenIdToTierLockAmount(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenIdToTierLockAmount", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToTierLockAmount is a free data retrieval call binding the contract method 0x02adafe1.
//
// Solidity: function tokenIdToTierLockAmount(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusSession) TokenIdToTierLockAmount(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToTierLockAmount(&_SPPlus.CallOpts, tokenId)
}

// TokenIdToTierLockAmount is a free data retrieval call binding the contract method 0x02adafe1.
//
// Solidity: function tokenIdToTierLockAmount(uint256 tokenId) view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TokenIdToTierLockAmount(tokenId *big.Int) (*big.Int, error) {
	return _SPPlus.Contract.TokenIdToTierLockAmount(&_SPPlus.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SPPlus *SPPlusCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SPPlus *SPPlusSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SPPlus.Contract.TokenURI(&_SPPlus.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SPPlus *SPPlusCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SPPlus.Contract.TokenURI(&_SPPlus.CallOpts, tokenId)
}

// TotalFilCashbackVaultBalance is a free data retrieval call binding the contract method 0x6b12bee3.
//
// Solidity: function totalFilCashbackVaultBalance() view returns(uint256)
func (_SPPlus *SPPlusCaller) TotalFilCashbackVaultBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "totalFilCashbackVaultBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFilCashbackVaultBalance is a free data retrieval call binding the contract method 0x6b12bee3.
//
// Solidity: function totalFilCashbackVaultBalance() view returns(uint256)
func (_SPPlus *SPPlusSession) TotalFilCashbackVaultBalance() (*big.Int, error) {
	return _SPPlus.Contract.TotalFilCashbackVaultBalance(&_SPPlus.CallOpts)
}

// TotalFilCashbackVaultBalance is a free data retrieval call binding the contract method 0x6b12bee3.
//
// Solidity: function totalFilCashbackVaultBalance() view returns(uint256)
func (_SPPlus *SPPlusCallerSession) TotalFilCashbackVaultBalance() (*big.Int, error) {
	return _SPPlus.Contract.TotalFilCashbackVaultBalance(&_SPPlus.CallOpts)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_SPPlus *SPPlusCaller) TransferFrom(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "transferFrom", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_SPPlus *SPPlusSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _SPPlus.Contract.TransferFrom(&_SPPlus.CallOpts, arg0, arg1, arg2)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_SPPlus *SPPlusCallerSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _SPPlus.Contract.TransferFrom(&_SPPlus.CallOpts, arg0, arg1, arg2)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_SPPlus *SPPlusCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_SPPlus *SPPlusSession) Treasury() (common.Address, error) {
	return _SPPlus.Contract.Treasury(&_SPPlus.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_SPPlus *SPPlusCallerSession) Treasury() (common.Address, error) {
	return _SPPlus.Contract.Treasury(&_SPPlus.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_SPPlus *SPPlusCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SPPlus.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_SPPlus *SPPlusSession) Version() (*big.Int, error) {
	return _SPPlus.Contract.Version(&_SPPlus.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_SPPlus *SPPlusCallerSession) Version() (*big.Int, error) {
	return _SPPlus.Contract.Version(&_SPPlus.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SPPlus *SPPlusTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SPPlus *SPPlusSession) AcceptOwnership() (*types.Transaction, error) {
	return _SPPlus.Contract.AcceptOwnership(&_SPPlus.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SPPlus *SPPlusTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SPPlus.Contract.AcceptOwnership(&_SPPlus.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xaef0353c.
//
// Solidity: function activate(address _beneficiary, uint256 _tokenId, uint8 _tier) returns()
func (_SPPlus *SPPlusTransactor) Activate(opts *bind.TransactOpts, _beneficiary common.Address, _tokenId *big.Int, _tier uint8) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "activate", _beneficiary, _tokenId, _tier)
}

// Activate is a paid mutator transaction binding the contract method 0xaef0353c.
//
// Solidity: function activate(address _beneficiary, uint256 _tokenId, uint8 _tier) returns()
func (_SPPlus *SPPlusSession) Activate(_beneficiary common.Address, _tokenId *big.Int, _tier uint8) (*types.Transaction, error) {
	return _SPPlus.Contract.Activate(&_SPPlus.TransactOpts, _beneficiary, _tokenId, _tier)
}

// Activate is a paid mutator transaction binding the contract method 0xaef0353c.
//
// Solidity: function activate(address _beneficiary, uint256 _tokenId, uint8 _tier) returns()
func (_SPPlus *SPPlusTransactorSession) Activate(_beneficiary common.Address, _tokenId *big.Int, _tier uint8) (*types.Transaction, error) {
	return _SPPlus.Contract.Activate(&_SPPlus.TransactOpts, _beneficiary, _tokenId, _tier)
}

// ChangeOwnerForAgent is a paid mutator transaction binding the contract method 0x28d0e5a9.
//
// Solidity: function changeOwnerForAgent(address _agent) returns()
func (_SPPlus *SPPlusTransactor) ChangeOwnerForAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "changeOwnerForAgent", _agent)
}

// ChangeOwnerForAgent is a paid mutator transaction binding the contract method 0x28d0e5a9.
//
// Solidity: function changeOwnerForAgent(address _agent) returns()
func (_SPPlus *SPPlusSession) ChangeOwnerForAgent(_agent common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.ChangeOwnerForAgent(&_SPPlus.TransactOpts, _agent)
}

// ChangeOwnerForAgent is a paid mutator transaction binding the contract method 0x28d0e5a9.
//
// Solidity: function changeOwnerForAgent(address _agent) returns()
func (_SPPlus *SPPlusTransactorSession) ChangeOwnerForAgent(_agent common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.ChangeOwnerForAgent(&_SPPlus.TransactOpts, _agent)
}

// ClaimCashBack is a paid mutator transaction binding the contract method 0x9202adfc.
//
// Solidity: function claimCashBack(uint256 _tokenId, address _receiver) returns()
func (_SPPlus *SPPlusTransactor) ClaimCashBack(opts *bind.TransactOpts, _tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "claimCashBack", _tokenId, _receiver)
}

// ClaimCashBack is a paid mutator transaction binding the contract method 0x9202adfc.
//
// Solidity: function claimCashBack(uint256 _tokenId, address _receiver) returns()
func (_SPPlus *SPPlusSession) ClaimCashBack(_tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.ClaimCashBack(&_SPPlus.TransactOpts, _tokenId, _receiver)
}

// ClaimCashBack is a paid mutator transaction binding the contract method 0x9202adfc.
//
// Solidity: function claimCashBack(uint256 _tokenId, address _receiver) returns()
func (_SPPlus *SPPlusTransactorSession) ClaimCashBack(_tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.ClaimCashBack(&_SPPlus.TransactOpts, _tokenId, _receiver)
}

// Downgrade is a paid mutator transaction binding the contract method 0x1fac65d3.
//
// Solidity: function downgrade(uint256 _tokenId, uint8 _desiredTier, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) _sc) returns()
func (_SPPlus *SPPlusTransactor) Downgrade(opts *bind.TransactOpts, _tokenId *big.Int, _desiredTier uint8, _sc SignedCredential) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "downgrade", _tokenId, _desiredTier, _sc)
}

// Downgrade is a paid mutator transaction binding the contract method 0x1fac65d3.
//
// Solidity: function downgrade(uint256 _tokenId, uint8 _desiredTier, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) _sc) returns()
func (_SPPlus *SPPlusSession) Downgrade(_tokenId *big.Int, _desiredTier uint8, _sc SignedCredential) (*types.Transaction, error) {
	return _SPPlus.Contract.Downgrade(&_SPPlus.TransactOpts, _tokenId, _desiredTier, _sc)
}

// Downgrade is a paid mutator transaction binding the contract method 0x1fac65d3.
//
// Solidity: function downgrade(uint256 _tokenId, uint8 _desiredTier, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) _sc) returns()
func (_SPPlus *SPPlusTransactorSession) Downgrade(_tokenId *big.Int, _desiredTier uint8, _sc SignedCredential) (*types.Transaction, error) {
	return _SPPlus.Contract.Downgrade(&_SPPlus.TransactOpts, _tokenId, _desiredTier, _sc)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_SPPlus *SPPlusTransactor) FundFilVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "fundFilVault")
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_SPPlus *SPPlusSession) FundFilVault() (*types.Transaction, error) {
	return _SPPlus.Contract.FundFilVault(&_SPPlus.TransactOpts)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_SPPlus *SPPlusTransactorSession) FundFilVault() (*types.Transaction, error) {
	return _SPPlus.Contract.FundFilVault(&_SPPlus.TransactOpts)
}

// FundGlfVault is a paid mutator transaction binding the contract method 0x463ad343.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount) returns()
func (_SPPlus *SPPlusTransactor) FundGlfVault(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "fundGlfVault", _tokenId, _amount)
}

// FundGlfVault is a paid mutator transaction binding the contract method 0x463ad343.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount) returns()
func (_SPPlus *SPPlusSession) FundGlfVault(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.FundGlfVault(&_SPPlus.TransactOpts, _tokenId, _amount)
}

// FundGlfVault is a paid mutator transaction binding the contract method 0x463ad343.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount) returns()
func (_SPPlus *SPPlusTransactorSession) FundGlfVault(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.FundGlfVault(&_SPPlus.TransactOpts, _tokenId, _amount)
}

// FundGlfVault0 is a paid mutator transaction binding the contract method 0xebca88fc.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount, uint256 _cashBackPercent) returns()
func (_SPPlus *SPPlusTransactor) FundGlfVault0(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "fundGlfVault0", _tokenId, _amount, _cashBackPercent)
}

// FundGlfVault0 is a paid mutator transaction binding the contract method 0xebca88fc.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount, uint256 _cashBackPercent) returns()
func (_SPPlus *SPPlusSession) FundGlfVault0(_tokenId *big.Int, _amount *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.FundGlfVault0(&_SPPlus.TransactOpts, _tokenId, _amount, _cashBackPercent)
}

// FundGlfVault0 is a paid mutator transaction binding the contract method 0xebca88fc.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount, uint256 _cashBackPercent) returns()
func (_SPPlus *SPPlusTransactorSession) FundGlfVault0(_tokenId *big.Int, _amount *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.FundGlfVault0(&_SPPlus.TransactOpts, _tokenId, _amount, _cashBackPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x86488a85.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier, address _poolAddress, address _priceSetter) returns()
func (_SPPlus *SPPlusTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address, _poolAddress common.Address, _priceSetter common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "initialize", _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier, _poolAddress, _priceSetter)
}

// Initialize is a paid mutator transaction binding the contract method 0x86488a85.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier, address _poolAddress, address _priceSetter) returns()
func (_SPPlus *SPPlusSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address, _poolAddress common.Address, _priceSetter common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.Initialize(&_SPPlus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier, _poolAddress, _priceSetter)
}

// Initialize is a paid mutator transaction binding the contract method 0x86488a85.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier, address _poolAddress, address _priceSetter) returns()
func (_SPPlus *SPPlusTransactorSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address, _poolAddress common.Address, _priceSetter common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.Initialize(&_SPPlus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier, _poolAddress, _priceSetter)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 tokenId)
func (_SPPlus *SPPlusTransactor) Mint(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "mint", _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 tokenId)
func (_SPPlus *SPPlusSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.Mint(&_SPPlus.TransactOpts, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 tokenId)
func (_SPPlus *SPPlusTransactorSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.Mint(&_SPPlus.TransactOpts, _to)
}

// MintActivateAndFund is a paid mutator transaction binding the contract method 0xf13c2c1c.
//
// Solidity: function mintActivateAndFund(uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier, uint256 _amount) returns(uint256 tokenId)
func (_SPPlus *SPPlusTransactor) MintActivateAndFund(opts *bind.TransactOpts, _personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "mintActivateAndFund", _personalCashBackPercent, _beneficiary, _tier, _amount)
}

// MintActivateAndFund is a paid mutator transaction binding the contract method 0xf13c2c1c.
//
// Solidity: function mintActivateAndFund(uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier, uint256 _amount) returns(uint256 tokenId)
func (_SPPlus *SPPlusSession) MintActivateAndFund(_personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.MintActivateAndFund(&_SPPlus.TransactOpts, _personalCashBackPercent, _beneficiary, _tier, _amount)
}

// MintActivateAndFund is a paid mutator transaction binding the contract method 0xf13c2c1c.
//
// Solidity: function mintActivateAndFund(uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier, uint256 _amount) returns(uint256 tokenId)
func (_SPPlus *SPPlusTransactorSession) MintActivateAndFund(_personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.MintActivateAndFund(&_SPPlus.TransactOpts, _personalCashBackPercent, _beneficiary, _tier, _amount)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0xe38a4911.
//
// Solidity: function mintAndActivate(address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_SPPlus *SPPlusTransactor) MintAndActivate(opts *bind.TransactOpts, _beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "mintAndActivate", _beneficiary, _tier)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0xe38a4911.
//
// Solidity: function mintAndActivate(address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_SPPlus *SPPlusSession) MintAndActivate(_beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _SPPlus.Contract.MintAndActivate(&_SPPlus.TransactOpts, _beneficiary, _tier)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0xe38a4911.
//
// Solidity: function mintAndActivate(address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_SPPlus *SPPlusTransactorSession) MintAndActivate(_beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _SPPlus.Contract.MintAndActivate(&_SPPlus.TransactOpts, _beneficiary, _tier)
}

// OnDefault is a paid mutator transaction binding the contract method 0x8b446431.
//
// Solidity: function onDefault(uint256 _agentId) returns()
func (_SPPlus *SPPlusTransactor) OnDefault(opts *bind.TransactOpts, _agentId *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "onDefault", _agentId)
}

// OnDefault is a paid mutator transaction binding the contract method 0x8b446431.
//
// Solidity: function onDefault(uint256 _agentId) returns()
func (_SPPlus *SPPlusSession) OnDefault(_agentId *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.OnDefault(&_SPPlus.TransactOpts, _agentId)
}

// OnDefault is a paid mutator transaction binding the contract method 0x8b446431.
//
// Solidity: function onDefault(uint256 _agentId) returns()
func (_SPPlus *SPPlusTransactorSession) OnDefault(_agentId *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.OnDefault(&_SPPlus.TransactOpts, _agentId)
}

// OnPaymentMade is a paid mutator transaction binding the contract method 0xcda43c10.
//
// Solidity: function onPaymentMade(uint256 _agentId, uint256 _interestAmount) returns()
func (_SPPlus *SPPlusTransactor) OnPaymentMade(opts *bind.TransactOpts, _agentId *big.Int, _interestAmount *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "onPaymentMade", _agentId, _interestAmount)
}

// OnPaymentMade is a paid mutator transaction binding the contract method 0xcda43c10.
//
// Solidity: function onPaymentMade(uint256 _agentId, uint256 _interestAmount) returns()
func (_SPPlus *SPPlusSession) OnPaymentMade(_agentId *big.Int, _interestAmount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.OnPaymentMade(&_SPPlus.TransactOpts, _agentId, _interestAmount)
}

// OnPaymentMade is a paid mutator transaction binding the contract method 0xcda43c10.
//
// Solidity: function onPaymentMade(uint256 _agentId, uint256 _interestAmount) returns()
func (_SPPlus *SPPlusTransactorSession) OnPaymentMade(_agentId *big.Int, _interestAmount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.OnPaymentMade(&_SPPlus.TransactOpts, _agentId, _interestAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SPPlus *SPPlusTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SPPlus *SPPlusSession) Pause() (*types.Transaction, error) {
	return _SPPlus.Contract.Pause(&_SPPlus.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SPPlus *SPPlusTransactorSession) Pause() (*types.Transaction, error) {
	return _SPPlus.Contract.Pause(&_SPPlus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SPPlus *SPPlusTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SPPlus *SPPlusSession) RenounceOwnership() (*types.Transaction, error) {
	return _SPPlus.Contract.RenounceOwnership(&_SPPlus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SPPlus *SPPlusTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SPPlus.Contract.RenounceOwnership(&_SPPlus.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SPPlus *SPPlusTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SPPlus *SPPlusSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SafeTransferFrom(&_SPPlus.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SPPlus *SPPlusTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SafeTransferFrom(&_SPPlus.TransactOpts, from, to, tokenId)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address _agentFactory) returns()
func (_SPPlus *SPPlusTransactor) SetAgentFactory(opts *bind.TransactOpts, _agentFactory common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setAgentFactory", _agentFactory)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address _agentFactory) returns()
func (_SPPlus *SPPlusSession) SetAgentFactory(_agentFactory common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetAgentFactory(&_SPPlus.TransactOpts, _agentFactory)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address _agentFactory) returns()
func (_SPPlus *SPPlusTransactorSession) SetAgentFactory(_agentFactory common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetAgentFactory(&_SPPlus.TransactOpts, _agentFactory)
}

// SetAgentPolice is a paid mutator transaction binding the contract method 0x86cc3437.
//
// Solidity: function setAgentPolice(address _agentPoliceVcVerifier) returns()
func (_SPPlus *SPPlusTransactor) SetAgentPolice(opts *bind.TransactOpts, _agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setAgentPolice", _agentPoliceVcVerifier)
}

// SetAgentPolice is a paid mutator transaction binding the contract method 0x86cc3437.
//
// Solidity: function setAgentPolice(address _agentPoliceVcVerifier) returns()
func (_SPPlus *SPPlusSession) SetAgentPolice(_agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetAgentPolice(&_SPPlus.TransactOpts, _agentPoliceVcVerifier)
}

// SetAgentPolice is a paid mutator transaction binding the contract method 0x86cc3437.
//
// Solidity: function setAgentPolice(address _agentPoliceVcVerifier) returns()
func (_SPPlus *SPPlusTransactorSession) SetAgentPolice(_agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetAgentPolice(&_SPPlus.TransactOpts, _agentPoliceVcVerifier)
}

// SetBaseConversionRateFILtoGLF is a paid mutator transaction binding the contract method 0xd86e1a35.
//
// Solidity: function setBaseConversionRateFILtoGLF(uint256 _baseConversionRateFILtoGLF) returns()
func (_SPPlus *SPPlusTransactor) SetBaseConversionRateFILtoGLF(opts *bind.TransactOpts, _baseConversionRateFILtoGLF *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setBaseConversionRateFILtoGLF", _baseConversionRateFILtoGLF)
}

// SetBaseConversionRateFILtoGLF is a paid mutator transaction binding the contract method 0xd86e1a35.
//
// Solidity: function setBaseConversionRateFILtoGLF(uint256 _baseConversionRateFILtoGLF) returns()
func (_SPPlus *SPPlusSession) SetBaseConversionRateFILtoGLF(_baseConversionRateFILtoGLF *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetBaseConversionRateFILtoGLF(&_SPPlus.TransactOpts, _baseConversionRateFILtoGLF)
}

// SetBaseConversionRateFILtoGLF is a paid mutator transaction binding the contract method 0xd86e1a35.
//
// Solidity: function setBaseConversionRateFILtoGLF(uint256 _baseConversionRateFILtoGLF) returns()
func (_SPPlus *SPPlusTransactorSession) SetBaseConversionRateFILtoGLF(_baseConversionRateFILtoGLF *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetBaseConversionRateFILtoGLF(&_SPPlus.TransactOpts, _baseConversionRateFILtoGLF)
}

// SetGlfToken is a paid mutator transaction binding the contract method 0x762b977d.
//
// Solidity: function setGlfToken(address _glfToken) returns()
func (_SPPlus *SPPlusTransactor) SetGlfToken(opts *bind.TransactOpts, _glfToken common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setGlfToken", _glfToken)
}

// SetGlfToken is a paid mutator transaction binding the contract method 0x762b977d.
//
// Solidity: function setGlfToken(address _glfToken) returns()
func (_SPPlus *SPPlusSession) SetGlfToken(_glfToken common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetGlfToken(&_SPPlus.TransactOpts, _glfToken)
}

// SetGlfToken is a paid mutator transaction binding the contract method 0x762b977d.
//
// Solidity: function setGlfToken(address _glfToken) returns()
func (_SPPlus *SPPlusTransactorSession) SetGlfToken(_glfToken common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetGlfToken(&_SPPlus.TransactOpts, _glfToken)
}

// SetMaxCashBackPercent is a paid mutator transaction binding the contract method 0x81eebd70.
//
// Solidity: function setMaxCashBackPercent(uint256 _maxCashBackPercent) returns()
func (_SPPlus *SPPlusTransactor) SetMaxCashBackPercent(opts *bind.TransactOpts, _maxCashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setMaxCashBackPercent", _maxCashBackPercent)
}

// SetMaxCashBackPercent is a paid mutator transaction binding the contract method 0x81eebd70.
//
// Solidity: function setMaxCashBackPercent(uint256 _maxCashBackPercent) returns()
func (_SPPlus *SPPlusSession) SetMaxCashBackPercent(_maxCashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetMaxCashBackPercent(&_SPPlus.TransactOpts, _maxCashBackPercent)
}

// SetMaxCashBackPercent is a paid mutator transaction binding the contract method 0x81eebd70.
//
// Solidity: function setMaxCashBackPercent(uint256 _maxCashBackPercent) returns()
func (_SPPlus *SPPlusTransactorSession) SetMaxCashBackPercent(_maxCashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetMaxCashBackPercent(&_SPPlus.TransactOpts, _maxCashBackPercent)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_SPPlus *SPPlusTransactor) SetMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setMintPrice", _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_SPPlus *SPPlusSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetMintPrice(&_SPPlus.TransactOpts, _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_SPPlus *SPPlusTransactorSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetMintPrice(&_SPPlus.TransactOpts, _mintPrice)
}

// SetPersonalCashBackPercent is a paid mutator transaction binding the contract method 0x51c9749c.
//
// Solidity: function setPersonalCashBackPercent(uint256 _tokenId, uint256 _cashBackPercent) returns()
func (_SPPlus *SPPlusTransactor) SetPersonalCashBackPercent(opts *bind.TransactOpts, _tokenId *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setPersonalCashBackPercent", _tokenId, _cashBackPercent)
}

// SetPersonalCashBackPercent is a paid mutator transaction binding the contract method 0x51c9749c.
//
// Solidity: function setPersonalCashBackPercent(uint256 _tokenId, uint256 _cashBackPercent) returns()
func (_SPPlus *SPPlusSession) SetPersonalCashBackPercent(_tokenId *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetPersonalCashBackPercent(&_SPPlus.TransactOpts, _tokenId, _cashBackPercent)
}

// SetPersonalCashBackPercent is a paid mutator transaction binding the contract method 0x51c9749c.
//
// Solidity: function setPersonalCashBackPercent(uint256 _tokenId, uint256 _cashBackPercent) returns()
func (_SPPlus *SPPlusTransactorSession) SetPersonalCashBackPercent(_tokenId *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetPersonalCashBackPercent(&_SPPlus.TransactOpts, _tokenId, _cashBackPercent)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xe9e15b4f.
//
// Solidity: function setPoolAddress(address _poolAddress) returns()
func (_SPPlus *SPPlusTransactor) SetPoolAddress(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setPoolAddress", _poolAddress)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xe9e15b4f.
//
// Solidity: function setPoolAddress(address _poolAddress) returns()
func (_SPPlus *SPPlusSession) SetPoolAddress(_poolAddress common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetPoolAddress(&_SPPlus.TransactOpts, _poolAddress)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xe9e15b4f.
//
// Solidity: function setPoolAddress(address _poolAddress) returns()
func (_SPPlus *SPPlusTransactorSession) SetPoolAddress(_poolAddress common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetPoolAddress(&_SPPlus.TransactOpts, _poolAddress)
}

// SetPriceSetter is a paid mutator transaction binding the contract method 0x63791e3c.
//
// Solidity: function setPriceSetter(address _priceSetter) returns()
func (_SPPlus *SPPlusTransactor) SetPriceSetter(opts *bind.TransactOpts, _priceSetter common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setPriceSetter", _priceSetter)
}

// SetPriceSetter is a paid mutator transaction binding the contract method 0x63791e3c.
//
// Solidity: function setPriceSetter(address _priceSetter) returns()
func (_SPPlus *SPPlusSession) SetPriceSetter(_priceSetter common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetPriceSetter(&_SPPlus.TransactOpts, _priceSetter)
}

// SetPriceSetter is a paid mutator transaction binding the contract method 0x63791e3c.
//
// Solidity: function setPriceSetter(address _priceSetter) returns()
func (_SPPlus *SPPlusTransactorSession) SetPriceSetter(_priceSetter common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetPriceSetter(&_SPPlus.TransactOpts, _priceSetter)
}

// SetTierInfo is a paid mutator transaction binding the contract method 0xf9d20f53.
//
// Solidity: function setTierInfo(uint8 _tier, (uint256,uint256,uint256) _tierInfo) returns()
func (_SPPlus *SPPlusTransactor) SetTierInfo(opts *bind.TransactOpts, _tier uint8, _tierInfo TierInfo) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setTierInfo", _tier, _tierInfo)
}

// SetTierInfo is a paid mutator transaction binding the contract method 0xf9d20f53.
//
// Solidity: function setTierInfo(uint8 _tier, (uint256,uint256,uint256) _tierInfo) returns()
func (_SPPlus *SPPlusSession) SetTierInfo(_tier uint8, _tierInfo TierInfo) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTierInfo(&_SPPlus.TransactOpts, _tier, _tierInfo)
}

// SetTierInfo is a paid mutator transaction binding the contract method 0xf9d20f53.
//
// Solidity: function setTierInfo(uint8 _tier, (uint256,uint256,uint256) _tierInfo) returns()
func (_SPPlus *SPPlusTransactorSession) SetTierInfo(_tier uint8, _tierInfo TierInfo) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTierInfo(&_SPPlus.TransactOpts, _tier, _tierInfo)
}

// SetTierSwitchPenaltyFee is a paid mutator transaction binding the contract method 0xabdf23df.
//
// Solidity: function setTierSwitchPenaltyFee(uint256 _tierSwitchPenaltyFee) returns()
func (_SPPlus *SPPlusTransactor) SetTierSwitchPenaltyFee(opts *bind.TransactOpts, _tierSwitchPenaltyFee *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setTierSwitchPenaltyFee", _tierSwitchPenaltyFee)
}

// SetTierSwitchPenaltyFee is a paid mutator transaction binding the contract method 0xabdf23df.
//
// Solidity: function setTierSwitchPenaltyFee(uint256 _tierSwitchPenaltyFee) returns()
func (_SPPlus *SPPlusSession) SetTierSwitchPenaltyFee(_tierSwitchPenaltyFee *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTierSwitchPenaltyFee(&_SPPlus.TransactOpts, _tierSwitchPenaltyFee)
}

// SetTierSwitchPenaltyFee is a paid mutator transaction binding the contract method 0xabdf23df.
//
// Solidity: function setTierSwitchPenaltyFee(uint256 _tierSwitchPenaltyFee) returns()
func (_SPPlus *SPPlusTransactorSession) SetTierSwitchPenaltyFee(_tierSwitchPenaltyFee *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTierSwitchPenaltyFee(&_SPPlus.TransactOpts, _tierSwitchPenaltyFee)
}

// SetTierSwitchPenaltyWindow is a paid mutator transaction binding the contract method 0x0c1cd643.
//
// Solidity: function setTierSwitchPenaltyWindow(uint256 _tierSwitchPenaltyWindow) returns()
func (_SPPlus *SPPlusTransactor) SetTierSwitchPenaltyWindow(opts *bind.TransactOpts, _tierSwitchPenaltyWindow *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setTierSwitchPenaltyWindow", _tierSwitchPenaltyWindow)
}

// SetTierSwitchPenaltyWindow is a paid mutator transaction binding the contract method 0x0c1cd643.
//
// Solidity: function setTierSwitchPenaltyWindow(uint256 _tierSwitchPenaltyWindow) returns()
func (_SPPlus *SPPlusSession) SetTierSwitchPenaltyWindow(_tierSwitchPenaltyWindow *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTierSwitchPenaltyWindow(&_SPPlus.TransactOpts, _tierSwitchPenaltyWindow)
}

// SetTierSwitchPenaltyWindow is a paid mutator transaction binding the contract method 0x0c1cd643.
//
// Solidity: function setTierSwitchPenaltyWindow(uint256 _tierSwitchPenaltyWindow) returns()
func (_SPPlus *SPPlusTransactorSession) SetTierSwitchPenaltyWindow(_tierSwitchPenaltyWindow *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTierSwitchPenaltyWindow(&_SPPlus.TransactOpts, _tierSwitchPenaltyWindow)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_SPPlus *SPPlusTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_SPPlus *SPPlusSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTreasury(&_SPPlus.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_SPPlus *SPPlusTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.SetTreasury(&_SPPlus.TransactOpts, _treasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SPPlus *SPPlusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SPPlus *SPPlusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.TransferOwnership(&_SPPlus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SPPlus *SPPlusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.TransferOwnership(&_SPPlus.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SPPlus *SPPlusTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SPPlus *SPPlusSession) Unpause() (*types.Transaction, error) {
	return _SPPlus.Contract.Unpause(&_SPPlus.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SPPlus *SPPlusTransactorSession) Unpause() (*types.Transaction, error) {
	return _SPPlus.Contract.Unpause(&_SPPlus.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x7dfe5b92.
//
// Solidity: function upgrade(uint256 _tokenId, uint8 _desiredTier) returns()
func (_SPPlus *SPPlusTransactor) Upgrade(opts *bind.TransactOpts, _tokenId *big.Int, _desiredTier uint8) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "upgrade", _tokenId, _desiredTier)
}

// Upgrade is a paid mutator transaction binding the contract method 0x7dfe5b92.
//
// Solidity: function upgrade(uint256 _tokenId, uint8 _desiredTier) returns()
func (_SPPlus *SPPlusSession) Upgrade(_tokenId *big.Int, _desiredTier uint8) (*types.Transaction, error) {
	return _SPPlus.Contract.Upgrade(&_SPPlus.TransactOpts, _tokenId, _desiredTier)
}

// Upgrade is a paid mutator transaction binding the contract method 0x7dfe5b92.
//
// Solidity: function upgrade(uint256 _tokenId, uint8 _desiredTier) returns()
func (_SPPlus *SPPlusTransactorSession) Upgrade(_tokenId *big.Int, _desiredTier uint8) (*types.Transaction, error) {
	return _SPPlus.Contract.Upgrade(&_SPPlus.TransactOpts, _tokenId, _desiredTier)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SPPlus *SPPlusTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SPPlus *SPPlusSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SPPlus.Contract.UpgradeToAndCall(&_SPPlus.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SPPlus *SPPlusTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SPPlus.Contract.UpgradeToAndCall(&_SPPlus.TransactOpts, newImplementation, data)
}

// WithdrawExtraLockedFunds is a paid mutator transaction binding the contract method 0x14b77f22.
//
// Solidity: function withdrawExtraLockedFunds(uint256 _tokenId) returns()
func (_SPPlus *SPPlusTransactor) WithdrawExtraLockedFunds(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "withdrawExtraLockedFunds", _tokenId)
}

// WithdrawExtraLockedFunds is a paid mutator transaction binding the contract method 0x14b77f22.
//
// Solidity: function withdrawExtraLockedFunds(uint256 _tokenId) returns()
func (_SPPlus *SPPlusSession) WithdrawExtraLockedFunds(_tokenId *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.WithdrawExtraLockedFunds(&_SPPlus.TransactOpts, _tokenId)
}

// WithdrawExtraLockedFunds is a paid mutator transaction binding the contract method 0x14b77f22.
//
// Solidity: function withdrawExtraLockedFunds(uint256 _tokenId) returns()
func (_SPPlus *SPPlusTransactorSession) WithdrawExtraLockedFunds(_tokenId *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.WithdrawExtraLockedFunds(&_SPPlus.TransactOpts, _tokenId)
}

// WithdrawFilFunds is a paid mutator transaction binding the contract method 0x2eacc696.
//
// Solidity: function withdrawFilFunds(address _to, uint256 _amount) returns()
func (_SPPlus *SPPlusTransactor) WithdrawFilFunds(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "withdrawFilFunds", _to, _amount)
}

// WithdrawFilFunds is a paid mutator transaction binding the contract method 0x2eacc696.
//
// Solidity: function withdrawFilFunds(address _to, uint256 _amount) returns()
func (_SPPlus *SPPlusSession) WithdrawFilFunds(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.WithdrawFilFunds(&_SPPlus.TransactOpts, _to, _amount)
}

// WithdrawFilFunds is a paid mutator transaction binding the contract method 0x2eacc696.
//
// Solidity: function withdrawFilFunds(address _to, uint256 _amount) returns()
func (_SPPlus *SPPlusTransactorSession) WithdrawFilFunds(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SPPlus.Contract.WithdrawFilFunds(&_SPPlus.TransactOpts, _to, _amount)
}

// WithdrawGlfVault is a paid mutator transaction binding the contract method 0x1eccd628.
//
// Solidity: function withdrawGlfVault(uint256 _tokenId, uint256 _amount, address _receiver) returns()
func (_SPPlus *SPPlusTransactor) WithdrawGlfVault(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SPPlus.contract.Transact(opts, "withdrawGlfVault", _tokenId, _amount, _receiver)
}

// WithdrawGlfVault is a paid mutator transaction binding the contract method 0x1eccd628.
//
// Solidity: function withdrawGlfVault(uint256 _tokenId, uint256 _amount, address _receiver) returns()
func (_SPPlus *SPPlusSession) WithdrawGlfVault(_tokenId *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.WithdrawGlfVault(&_SPPlus.TransactOpts, _tokenId, _amount, _receiver)
}

// WithdrawGlfVault is a paid mutator transaction binding the contract method 0x1eccd628.
//
// Solidity: function withdrawGlfVault(uint256 _tokenId, uint256 _amount, address _receiver) returns()
func (_SPPlus *SPPlusTransactorSession) WithdrawGlfVault(_tokenId *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SPPlus.Contract.WithdrawGlfVault(&_SPPlus.TransactOpts, _tokenId, _amount, _receiver)
}

// SPPlusApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SPPlus contract.
type SPPlusApprovalIterator struct {
	Event *SPPlusApproval // Event containing the contract specifics and raw log

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
func (it *SPPlusApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusApproval)
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
		it.Event = new(SPPlusApproval)
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
func (it *SPPlusApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusApproval represents a Approval event raised by the SPPlus contract.
type SPPlusApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SPPlus *SPPlusFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SPPlusApprovalIterator, error) {

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

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusApprovalIterator{contract: _SPPlus.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SPPlus *SPPlusFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SPPlusApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusApproval)
				if err := _SPPlus.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseApproval(log types.Log) (*SPPlusApproval, error) {
	event := new(SPPlusApproval)
	if err := _SPPlus.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SPPlus contract.
type SPPlusApprovalForAllIterator struct {
	Event *SPPlusApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SPPlusApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusApprovalForAll)
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
		it.Event = new(SPPlusApprovalForAll)
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
func (it *SPPlusApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusApprovalForAll represents a ApprovalForAll event raised by the SPPlus contract.
type SPPlusApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SPPlus *SPPlusFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SPPlusApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusApprovalForAllIterator{contract: _SPPlus.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SPPlus *SPPlusFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SPPlusApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusApprovalForAll)
				if err := _SPPlus.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseApprovalForAll(log types.Log) (*SPPlusApprovalForAll, error) {
	event := new(SPPlusApprovalForAll)
	if err := _SPPlus.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusCardActivatedIterator is returned from FilterCardActivated and is used to iterate over the raw logs and unpacked data for CardActivated events raised by the SPPlus contract.
type SPPlusCardActivatedIterator struct {
	Event *SPPlusCardActivated // Event containing the contract specifics and raw log

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
func (it *SPPlusCardActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusCardActivated)
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
		it.Event = new(SPPlusCardActivated)
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
func (it *SPPlusCardActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusCardActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusCardActivated represents a CardActivated event raised by the SPPlus contract.
type SPPlusCardActivated struct {
	TokenId     *big.Int
	Beneficiary common.Address
	Tier        uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCardActivated is a free log retrieval operation binding the contract event 0x62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5.
//
// Solidity: event CardActivated(uint256 indexed tokenId, address indexed beneficiary, uint8 tier)
func (_SPPlus *SPPlusFilterer) FilterCardActivated(opts *bind.FilterOpts, tokenId []*big.Int, beneficiary []common.Address) (*SPPlusCardActivatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "CardActivated", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusCardActivatedIterator{contract: _SPPlus.contract, event: "CardActivated", logs: logs, sub: sub}, nil
}

// WatchCardActivated is a free log subscription operation binding the contract event 0x62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5.
//
// Solidity: event CardActivated(uint256 indexed tokenId, address indexed beneficiary, uint8 tier)
func (_SPPlus *SPPlusFilterer) WatchCardActivated(opts *bind.WatchOpts, sink chan<- *SPPlusCardActivated, tokenId []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "CardActivated", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusCardActivated)
				if err := _SPPlus.contract.UnpackLog(event, "CardActivated", log); err != nil {
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

// ParseCardActivated is a log parse operation binding the contract event 0x62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5.
//
// Solidity: event CardActivated(uint256 indexed tokenId, address indexed beneficiary, uint8 tier)
func (_SPPlus *SPPlusFilterer) ParseCardActivated(log types.Log) (*SPPlusCardActivated, error) {
	event := new(SPPlusCardActivated)
	if err := _SPPlus.contract.UnpackLog(event, "CardActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusCardDowngradedIterator is returned from FilterCardDowngraded and is used to iterate over the raw logs and unpacked data for CardDowngraded events raised by the SPPlus contract.
type SPPlusCardDowngradedIterator struct {
	Event *SPPlusCardDowngraded // Event containing the contract specifics and raw log

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
func (it *SPPlusCardDowngradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusCardDowngraded)
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
		it.Event = new(SPPlusCardDowngraded)
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
func (it *SPPlusCardDowngradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusCardDowngradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusCardDowngraded represents a CardDowngraded event raised by the SPPlus contract.
type SPPlusCardDowngraded struct {
	TokenId       *big.Int
	Beneficiary   common.Address
	OldTier       uint8
	NewTier       uint8
	PenaltyAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCardDowngraded is a free log retrieval operation binding the contract event 0x89a80b3a2ac9f4ceb70a953f78c6572114c7041e11e2fae9f8f1e22425402c8b.
//
// Solidity: event CardDowngraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier, uint256 penaltyAmount)
func (_SPPlus *SPPlusFilterer) FilterCardDowngraded(opts *bind.FilterOpts, tokenId []*big.Int, beneficiary []common.Address) (*SPPlusCardDowngradedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "CardDowngraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusCardDowngradedIterator{contract: _SPPlus.contract, event: "CardDowngraded", logs: logs, sub: sub}, nil
}

// WatchCardDowngraded is a free log subscription operation binding the contract event 0x89a80b3a2ac9f4ceb70a953f78c6572114c7041e11e2fae9f8f1e22425402c8b.
//
// Solidity: event CardDowngraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier, uint256 penaltyAmount)
func (_SPPlus *SPPlusFilterer) WatchCardDowngraded(opts *bind.WatchOpts, sink chan<- *SPPlusCardDowngraded, tokenId []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "CardDowngraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusCardDowngraded)
				if err := _SPPlus.contract.UnpackLog(event, "CardDowngraded", log); err != nil {
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

// ParseCardDowngraded is a log parse operation binding the contract event 0x89a80b3a2ac9f4ceb70a953f78c6572114c7041e11e2fae9f8f1e22425402c8b.
//
// Solidity: event CardDowngraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier, uint256 penaltyAmount)
func (_SPPlus *SPPlusFilterer) ParseCardDowngraded(log types.Log) (*SPPlusCardDowngraded, error) {
	event := new(SPPlusCardDowngraded)
	if err := _SPPlus.contract.UnpackLog(event, "CardDowngraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusCardMintedIterator is returned from FilterCardMinted and is used to iterate over the raw logs and unpacked data for CardMinted events raised by the SPPlus contract.
type SPPlusCardMintedIterator struct {
	Event *SPPlusCardMinted // Event containing the contract specifics and raw log

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
func (it *SPPlusCardMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusCardMinted)
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
		it.Event = new(SPPlusCardMinted)
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
func (it *SPPlusCardMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusCardMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusCardMinted represents a CardMinted event raised by the SPPlus contract.
type SPPlusCardMinted struct {
	TokenId  *big.Int
	Minter   common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCardMinted is a free log retrieval operation binding the contract event 0x864b3215ee93b2a9291e17b45d665127a426c1259666cf483c30939c4a710332.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address minter, address receiver)
func (_SPPlus *SPPlusFilterer) FilterCardMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*SPPlusCardMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "CardMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusCardMintedIterator{contract: _SPPlus.contract, event: "CardMinted", logs: logs, sub: sub}, nil
}

// WatchCardMinted is a free log subscription operation binding the contract event 0x864b3215ee93b2a9291e17b45d665127a426c1259666cf483c30939c4a710332.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address minter, address receiver)
func (_SPPlus *SPPlusFilterer) WatchCardMinted(opts *bind.WatchOpts, sink chan<- *SPPlusCardMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "CardMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusCardMinted)
				if err := _SPPlus.contract.UnpackLog(event, "CardMinted", log); err != nil {
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

// ParseCardMinted is a log parse operation binding the contract event 0x864b3215ee93b2a9291e17b45d665127a426c1259666cf483c30939c4a710332.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address minter, address receiver)
func (_SPPlus *SPPlusFilterer) ParseCardMinted(log types.Log) (*SPPlusCardMinted, error) {
	event := new(SPPlusCardMinted)
	if err := _SPPlus.contract.UnpackLog(event, "CardMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusCardUpgradedIterator is returned from FilterCardUpgraded and is used to iterate over the raw logs and unpacked data for CardUpgraded events raised by the SPPlus contract.
type SPPlusCardUpgradedIterator struct {
	Event *SPPlusCardUpgraded // Event containing the contract specifics and raw log

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
func (it *SPPlusCardUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusCardUpgraded)
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
		it.Event = new(SPPlusCardUpgraded)
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
func (it *SPPlusCardUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusCardUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusCardUpgraded represents a CardUpgraded event raised by the SPPlus contract.
type SPPlusCardUpgraded struct {
	TokenId     *big.Int
	Beneficiary common.Address
	OldTier     uint8
	NewTier     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCardUpgraded is a free log retrieval operation binding the contract event 0x7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd.
//
// Solidity: event CardUpgraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_SPPlus *SPPlusFilterer) FilterCardUpgraded(opts *bind.FilterOpts, tokenId []*big.Int, beneficiary []common.Address) (*SPPlusCardUpgradedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "CardUpgraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusCardUpgradedIterator{contract: _SPPlus.contract, event: "CardUpgraded", logs: logs, sub: sub}, nil
}

// WatchCardUpgraded is a free log subscription operation binding the contract event 0x7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd.
//
// Solidity: event CardUpgraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_SPPlus *SPPlusFilterer) WatchCardUpgraded(opts *bind.WatchOpts, sink chan<- *SPPlusCardUpgraded, tokenId []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "CardUpgraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusCardUpgraded)
				if err := _SPPlus.contract.UnpackLog(event, "CardUpgraded", log); err != nil {
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

// ParseCardUpgraded is a log parse operation binding the contract event 0x7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd.
//
// Solidity: event CardUpgraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_SPPlus *SPPlusFilterer) ParseCardUpgraded(log types.Log) (*SPPlusCardUpgraded, error) {
	event := new(SPPlusCardUpgraded)
	if err := _SPPlus.contract.UnpackLog(event, "CardUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusCashBackClaimedIterator is returned from FilterCashBackClaimed and is used to iterate over the raw logs and unpacked data for CashBackClaimed events raised by the SPPlus contract.
type SPPlusCashBackClaimedIterator struct {
	Event *SPPlusCashBackClaimed // Event containing the contract specifics and raw log

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
func (it *SPPlusCashBackClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusCashBackClaimed)
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
		it.Event = new(SPPlusCashBackClaimed)
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
func (it *SPPlusCashBackClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusCashBackClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusCashBackClaimed represents a CashBackClaimed event raised by the SPPlus contract.
type SPPlusCashBackClaimed struct {
	TokenId  *big.Int
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCashBackClaimed is a free log retrieval operation binding the contract event 0x05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700.
//
// Solidity: event CashBackClaimed(uint256 indexed tokenId, address indexed receiver, uint256 amount)
func (_SPPlus *SPPlusFilterer) FilterCashBackClaimed(opts *bind.FilterOpts, tokenId []*big.Int, receiver []common.Address) (*SPPlusCashBackClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "CashBackClaimed", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusCashBackClaimedIterator{contract: _SPPlus.contract, event: "CashBackClaimed", logs: logs, sub: sub}, nil
}

// WatchCashBackClaimed is a free log subscription operation binding the contract event 0x05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700.
//
// Solidity: event CashBackClaimed(uint256 indexed tokenId, address indexed receiver, uint256 amount)
func (_SPPlus *SPPlusFilterer) WatchCashBackClaimed(opts *bind.WatchOpts, sink chan<- *SPPlusCashBackClaimed, tokenId []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "CashBackClaimed", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusCashBackClaimed)
				if err := _SPPlus.contract.UnpackLog(event, "CashBackClaimed", log); err != nil {
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

// ParseCashBackClaimed is a log parse operation binding the contract event 0x05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700.
//
// Solidity: event CashBackClaimed(uint256 indexed tokenId, address indexed receiver, uint256 amount)
func (_SPPlus *SPPlusFilterer) ParseCashBackClaimed(log types.Log) (*SPPlusCashBackClaimed, error) {
	event := new(SPPlusCashBackClaimed)
	if err := _SPPlus.contract.UnpackLog(event, "CashBackClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusFilVaultFundedIterator is returned from FilterFilVaultFunded and is used to iterate over the raw logs and unpacked data for FilVaultFunded events raised by the SPPlus contract.
type SPPlusFilVaultFundedIterator struct {
	Event *SPPlusFilVaultFunded // Event containing the contract specifics and raw log

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
func (it *SPPlusFilVaultFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusFilVaultFunded)
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
		it.Event = new(SPPlusFilVaultFunded)
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
func (it *SPPlusFilVaultFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusFilVaultFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusFilVaultFunded represents a FilVaultFunded event raised by the SPPlus contract.
type SPPlusFilVaultFunded struct {
	TokenId *big.Int
	Owner   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFilVaultFunded is a free log retrieval operation binding the contract event 0x1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d6074.
//
// Solidity: event FilVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
func (_SPPlus *SPPlusFilterer) FilterFilVaultFunded(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*SPPlusFilVaultFundedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "FilVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusFilVaultFundedIterator{contract: _SPPlus.contract, event: "FilVaultFunded", logs: logs, sub: sub}, nil
}

// WatchFilVaultFunded is a free log subscription operation binding the contract event 0x1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d6074.
//
// Solidity: event FilVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
func (_SPPlus *SPPlusFilterer) WatchFilVaultFunded(opts *bind.WatchOpts, sink chan<- *SPPlusFilVaultFunded, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "FilVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusFilVaultFunded)
				if err := _SPPlus.contract.UnpackLog(event, "FilVaultFunded", log); err != nil {
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

// ParseFilVaultFunded is a log parse operation binding the contract event 0x1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d6074.
//
// Solidity: event FilVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
func (_SPPlus *SPPlusFilterer) ParseFilVaultFunded(log types.Log) (*SPPlusFilVaultFunded, error) {
	event := new(SPPlusFilVaultFunded)
	if err := _SPPlus.contract.UnpackLog(event, "FilVaultFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusGlfVaultFundedIterator is returned from FilterGlfVaultFunded and is used to iterate over the raw logs and unpacked data for GlfVaultFunded events raised by the SPPlus contract.
type SPPlusGlfVaultFundedIterator struct {
	Event *SPPlusGlfVaultFunded // Event containing the contract specifics and raw log

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
func (it *SPPlusGlfVaultFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusGlfVaultFunded)
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
		it.Event = new(SPPlusGlfVaultFunded)
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
func (it *SPPlusGlfVaultFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusGlfVaultFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusGlfVaultFunded represents a GlfVaultFunded event raised by the SPPlus contract.
type SPPlusGlfVaultFunded struct {
	TokenId         *big.Int
	Owner           common.Address
	Amount          *big.Int
	CashBackPercent *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGlfVaultFunded is a free log retrieval operation binding the contract event 0x57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c6.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 cashBackPercent)
func (_SPPlus *SPPlusFilterer) FilterGlfVaultFunded(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*SPPlusGlfVaultFundedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "GlfVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusGlfVaultFundedIterator{contract: _SPPlus.contract, event: "GlfVaultFunded", logs: logs, sub: sub}, nil
}

// WatchGlfVaultFunded is a free log subscription operation binding the contract event 0x57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c6.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 cashBackPercent)
func (_SPPlus *SPPlusFilterer) WatchGlfVaultFunded(opts *bind.WatchOpts, sink chan<- *SPPlusGlfVaultFunded, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "GlfVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusGlfVaultFunded)
				if err := _SPPlus.contract.UnpackLog(event, "GlfVaultFunded", log); err != nil {
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

// ParseGlfVaultFunded is a log parse operation binding the contract event 0x57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c6.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 cashBackPercent)
func (_SPPlus *SPPlusFilterer) ParseGlfVaultFunded(log types.Log) (*SPPlusGlfVaultFunded, error) {
	event := new(SPPlusGlfVaultFunded)
	if err := _SPPlus.contract.UnpackLog(event, "GlfVaultFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusGlfVaultWithdrawnIterator is returned from FilterGlfVaultWithdrawn and is used to iterate over the raw logs and unpacked data for GlfVaultWithdrawn events raised by the SPPlus contract.
type SPPlusGlfVaultWithdrawnIterator struct {
	Event *SPPlusGlfVaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *SPPlusGlfVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusGlfVaultWithdrawn)
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
		it.Event = new(SPPlusGlfVaultWithdrawn)
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
func (it *SPPlusGlfVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusGlfVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusGlfVaultWithdrawn represents a GlfVaultWithdrawn event raised by the SPPlus contract.
type SPPlusGlfVaultWithdrawn struct {
	TokenId  *big.Int
	Owner    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGlfVaultWithdrawn is a free log retrieval operation binding the contract event 0x36ae8a1f61e4ec367d806b7c73a9031431dd6fae6810392041e8791c1ddba34f.
//
// Solidity: event GlfVaultWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount)
func (_SPPlus *SPPlusFilterer) FilterGlfVaultWithdrawn(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address, receiver []common.Address) (*SPPlusGlfVaultWithdrawnIterator, error) {

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

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "GlfVaultWithdrawn", tokenIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusGlfVaultWithdrawnIterator{contract: _SPPlus.contract, event: "GlfVaultWithdrawn", logs: logs, sub: sub}, nil
}

// WatchGlfVaultWithdrawn is a free log subscription operation binding the contract event 0x36ae8a1f61e4ec367d806b7c73a9031431dd6fae6810392041e8791c1ddba34f.
//
// Solidity: event GlfVaultWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount)
func (_SPPlus *SPPlusFilterer) WatchGlfVaultWithdrawn(opts *bind.WatchOpts, sink chan<- *SPPlusGlfVaultWithdrawn, tokenId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "GlfVaultWithdrawn", tokenIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusGlfVaultWithdrawn)
				if err := _SPPlus.contract.UnpackLog(event, "GlfVaultWithdrawn", log); err != nil {
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

// ParseGlfVaultWithdrawn is a log parse operation binding the contract event 0x36ae8a1f61e4ec367d806b7c73a9031431dd6fae6810392041e8791c1ddba34f.
//
// Solidity: event GlfVaultWithdrawn(uint256 indexed tokenId, address indexed owner, address indexed receiver, uint256 amount)
func (_SPPlus *SPPlusFilterer) ParseGlfVaultWithdrawn(log types.Log) (*SPPlusGlfVaultWithdrawn, error) {
	event := new(SPPlusGlfVaultWithdrawn)
	if err := _SPPlus.contract.UnpackLog(event, "GlfVaultWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SPPlus contract.
type SPPlusInitializedIterator struct {
	Event *SPPlusInitialized // Event containing the contract specifics and raw log

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
func (it *SPPlusInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusInitialized)
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
		it.Event = new(SPPlusInitialized)
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
func (it *SPPlusInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusInitialized represents a Initialized event raised by the SPPlus contract.
type SPPlusInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SPPlus *SPPlusFilterer) FilterInitialized(opts *bind.FilterOpts) (*SPPlusInitializedIterator, error) {

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SPPlusInitializedIterator{contract: _SPPlus.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SPPlus *SPPlusFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SPPlusInitialized) (event.Subscription, error) {

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusInitialized)
				if err := _SPPlus.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseInitialized(log types.Log) (*SPPlusInitialized, error) {
	event := new(SPPlusInitialized)
	if err := _SPPlus.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the SPPlus contract.
type SPPlusOwnershipTransferStartedIterator struct {
	Event *SPPlusOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *SPPlusOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusOwnershipTransferStarted)
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
		it.Event = new(SPPlusOwnershipTransferStarted)
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
func (it *SPPlusOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the SPPlus contract.
type SPPlusOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SPPlus *SPPlusFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SPPlusOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusOwnershipTransferStartedIterator{contract: _SPPlus.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SPPlus *SPPlusFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SPPlusOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusOwnershipTransferStarted)
				if err := _SPPlus.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseOwnershipTransferStarted(log types.Log) (*SPPlusOwnershipTransferStarted, error) {
	event := new(SPPlusOwnershipTransferStarted)
	if err := _SPPlus.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SPPlus contract.
type SPPlusOwnershipTransferredIterator struct {
	Event *SPPlusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SPPlusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusOwnershipTransferred)
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
		it.Event = new(SPPlusOwnershipTransferred)
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
func (it *SPPlusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusOwnershipTransferred represents a OwnershipTransferred event raised by the SPPlus contract.
type SPPlusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SPPlus *SPPlusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SPPlusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusOwnershipTransferredIterator{contract: _SPPlus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SPPlus *SPPlusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SPPlusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusOwnershipTransferred)
				if err := _SPPlus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseOwnershipTransferred(log types.Log) (*SPPlusOwnershipTransferred, error) {
	event := new(SPPlusOwnershipTransferred)
	if err := _SPPlus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SPPlus contract.
type SPPlusPausedIterator struct {
	Event *SPPlusPaused // Event containing the contract specifics and raw log

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
func (it *SPPlusPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusPaused)
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
		it.Event = new(SPPlusPaused)
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
func (it *SPPlusPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusPaused represents a Paused event raised by the SPPlus contract.
type SPPlusPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SPPlus *SPPlusFilterer) FilterPaused(opts *bind.FilterOpts) (*SPPlusPausedIterator, error) {

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SPPlusPausedIterator{contract: _SPPlus.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SPPlus *SPPlusFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SPPlusPaused) (event.Subscription, error) {

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusPaused)
				if err := _SPPlus.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParsePaused(log types.Log) (*SPPlusPaused, error) {
	event := new(SPPlusPaused)
	if err := _SPPlus.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusPaymentProcessedIterator is returned from FilterPaymentProcessed and is used to iterate over the raw logs and unpacked data for PaymentProcessed events raised by the SPPlus contract.
type SPPlusPaymentProcessedIterator struct {
	Event *SPPlusPaymentProcessed // Event containing the contract specifics and raw log

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
func (it *SPPlusPaymentProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusPaymentProcessed)
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
		it.Event = new(SPPlusPaymentProcessed)
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
func (it *SPPlusPaymentProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusPaymentProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusPaymentProcessed represents a PaymentProcessed event raised by the SPPlus contract.
type SPPlusPaymentProcessed struct {
	AgentId        *big.Int
	TokenId        *big.Int
	InterestAmount *big.Int
	GlfBurned      *big.Int
	FilEarned      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPaymentProcessed is a free log retrieval operation binding the contract event 0x299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d6.
//
// Solidity: event PaymentProcessed(uint256 indexed agentId, uint256 indexed tokenId, uint256 interestAmount, uint256 glfBurned, uint256 filEarned)
func (_SPPlus *SPPlusFilterer) FilterPaymentProcessed(opts *bind.FilterOpts, agentId []*big.Int, tokenId []*big.Int) (*SPPlusPaymentProcessedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "PaymentProcessed", agentIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusPaymentProcessedIterator{contract: _SPPlus.contract, event: "PaymentProcessed", logs: logs, sub: sub}, nil
}

// WatchPaymentProcessed is a free log subscription operation binding the contract event 0x299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d6.
//
// Solidity: event PaymentProcessed(uint256 indexed agentId, uint256 indexed tokenId, uint256 interestAmount, uint256 glfBurned, uint256 filEarned)
func (_SPPlus *SPPlusFilterer) WatchPaymentProcessed(opts *bind.WatchOpts, sink chan<- *SPPlusPaymentProcessed, agentId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "PaymentProcessed", agentIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusPaymentProcessed)
				if err := _SPPlus.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
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

// ParsePaymentProcessed is a log parse operation binding the contract event 0x299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d6.
//
// Solidity: event PaymentProcessed(uint256 indexed agentId, uint256 indexed tokenId, uint256 interestAmount, uint256 glfBurned, uint256 filEarned)
func (_SPPlus *SPPlusFilterer) ParsePaymentProcessed(log types.Log) (*SPPlusPaymentProcessed, error) {
	event := new(SPPlusPaymentProcessed)
	if err := _SPPlus.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusPersonalCashBackUpdatedIterator is returned from FilterPersonalCashBackUpdated and is used to iterate over the raw logs and unpacked data for PersonalCashBackUpdated events raised by the SPPlus contract.
type SPPlusPersonalCashBackUpdatedIterator struct {
	Event *SPPlusPersonalCashBackUpdated // Event containing the contract specifics and raw log

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
func (it *SPPlusPersonalCashBackUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusPersonalCashBackUpdated)
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
		it.Event = new(SPPlusPersonalCashBackUpdated)
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
func (it *SPPlusPersonalCashBackUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusPersonalCashBackUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusPersonalCashBackUpdated represents a PersonalCashBackUpdated event raised by the SPPlus contract.
type SPPlusPersonalCashBackUpdated struct {
	TokenId     *big.Int
	Owner       common.Address
	OldCashBack *big.Int
	NewCashBack *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPersonalCashBackUpdated is a free log retrieval operation binding the contract event 0xdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca.
//
// Solidity: event PersonalCashBackUpdated(uint256 indexed tokenId, address indexed owner, uint256 oldCashBack, uint256 newCashBack)
func (_SPPlus *SPPlusFilterer) FilterPersonalCashBackUpdated(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*SPPlusPersonalCashBackUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "PersonalCashBackUpdated", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusPersonalCashBackUpdatedIterator{contract: _SPPlus.contract, event: "PersonalCashBackUpdated", logs: logs, sub: sub}, nil
}

// WatchPersonalCashBackUpdated is a free log subscription operation binding the contract event 0xdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca.
//
// Solidity: event PersonalCashBackUpdated(uint256 indexed tokenId, address indexed owner, uint256 oldCashBack, uint256 newCashBack)
func (_SPPlus *SPPlusFilterer) WatchPersonalCashBackUpdated(opts *bind.WatchOpts, sink chan<- *SPPlusPersonalCashBackUpdated, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "PersonalCashBackUpdated", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusPersonalCashBackUpdated)
				if err := _SPPlus.contract.UnpackLog(event, "PersonalCashBackUpdated", log); err != nil {
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

// ParsePersonalCashBackUpdated is a log parse operation binding the contract event 0xdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca.
//
// Solidity: event PersonalCashBackUpdated(uint256 indexed tokenId, address indexed owner, uint256 oldCashBack, uint256 newCashBack)
func (_SPPlus *SPPlusFilterer) ParsePersonalCashBackUpdated(log types.Log) (*SPPlusPersonalCashBackUpdated, error) {
	event := new(SPPlusPersonalCashBackUpdated)
	if err := _SPPlus.contract.UnpackLog(event, "PersonalCashBackUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusPriceSetterUpdatedIterator is returned from FilterPriceSetterUpdated and is used to iterate over the raw logs and unpacked data for PriceSetterUpdated events raised by the SPPlus contract.
type SPPlusPriceSetterUpdatedIterator struct {
	Event *SPPlusPriceSetterUpdated // Event containing the contract specifics and raw log

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
func (it *SPPlusPriceSetterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusPriceSetterUpdated)
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
		it.Event = new(SPPlusPriceSetterUpdated)
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
func (it *SPPlusPriceSetterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusPriceSetterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusPriceSetterUpdated represents a PriceSetterUpdated event raised by the SPPlus contract.
type SPPlusPriceSetterUpdated struct {
	OldPriceSetter common.Address
	NewPriceSetter common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceSetterUpdated is a free log retrieval operation binding the contract event 0x379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa40079.
//
// Solidity: event PriceSetterUpdated(address indexed oldPriceSetter, address indexed newPriceSetter)
func (_SPPlus *SPPlusFilterer) FilterPriceSetterUpdated(opts *bind.FilterOpts, oldPriceSetter []common.Address, newPriceSetter []common.Address) (*SPPlusPriceSetterUpdatedIterator, error) {

	var oldPriceSetterRule []interface{}
	for _, oldPriceSetterItem := range oldPriceSetter {
		oldPriceSetterRule = append(oldPriceSetterRule, oldPriceSetterItem)
	}
	var newPriceSetterRule []interface{}
	for _, newPriceSetterItem := range newPriceSetter {
		newPriceSetterRule = append(newPriceSetterRule, newPriceSetterItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "PriceSetterUpdated", oldPriceSetterRule, newPriceSetterRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusPriceSetterUpdatedIterator{contract: _SPPlus.contract, event: "PriceSetterUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceSetterUpdated is a free log subscription operation binding the contract event 0x379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa40079.
//
// Solidity: event PriceSetterUpdated(address indexed oldPriceSetter, address indexed newPriceSetter)
func (_SPPlus *SPPlusFilterer) WatchPriceSetterUpdated(opts *bind.WatchOpts, sink chan<- *SPPlusPriceSetterUpdated, oldPriceSetter []common.Address, newPriceSetter []common.Address) (event.Subscription, error) {

	var oldPriceSetterRule []interface{}
	for _, oldPriceSetterItem := range oldPriceSetter {
		oldPriceSetterRule = append(oldPriceSetterRule, oldPriceSetterItem)
	}
	var newPriceSetterRule []interface{}
	for _, newPriceSetterItem := range newPriceSetter {
		newPriceSetterRule = append(newPriceSetterRule, newPriceSetterItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "PriceSetterUpdated", oldPriceSetterRule, newPriceSetterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusPriceSetterUpdated)
				if err := _SPPlus.contract.UnpackLog(event, "PriceSetterUpdated", log); err != nil {
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

// ParsePriceSetterUpdated is a log parse operation binding the contract event 0x379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa40079.
//
// Solidity: event PriceSetterUpdated(address indexed oldPriceSetter, address indexed newPriceSetter)
func (_SPPlus *SPPlusFilterer) ParsePriceSetterUpdated(log types.Log) (*SPPlusPriceSetterUpdated, error) {
	event := new(SPPlusPriceSetterUpdated)
	if err := _SPPlus.contract.UnpackLog(event, "PriceSetterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SPPlus contract.
type SPPlusTransferIterator struct {
	Event *SPPlusTransfer // Event containing the contract specifics and raw log

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
func (it *SPPlusTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusTransfer)
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
		it.Event = new(SPPlusTransfer)
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
func (it *SPPlusTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusTransfer represents a Transfer event raised by the SPPlus contract.
type SPPlusTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SPPlus *SPPlusFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SPPlusTransferIterator, error) {

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

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusTransferIterator{contract: _SPPlus.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SPPlus *SPPlusFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SPPlusTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusTransfer)
				if err := _SPPlus.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseTransfer(log types.Log) (*SPPlusTransfer, error) {
	event := new(SPPlusTransfer)
	if err := _SPPlus.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SPPlus contract.
type SPPlusUnpausedIterator struct {
	Event *SPPlusUnpaused // Event containing the contract specifics and raw log

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
func (it *SPPlusUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusUnpaused)
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
		it.Event = new(SPPlusUnpaused)
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
func (it *SPPlusUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusUnpaused represents a Unpaused event raised by the SPPlus contract.
type SPPlusUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SPPlus *SPPlusFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SPPlusUnpausedIterator, error) {

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SPPlusUnpausedIterator{contract: _SPPlus.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SPPlus *SPPlusFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SPPlusUnpaused) (event.Subscription, error) {

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusUnpaused)
				if err := _SPPlus.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseUnpaused(log types.Log) (*SPPlusUnpaused, error) {
	event := new(SPPlusUnpaused)
	if err := _SPPlus.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SPPlusUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SPPlus contract.
type SPPlusUpgradedIterator struct {
	Event *SPPlusUpgraded // Event containing the contract specifics and raw log

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
func (it *SPPlusUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SPPlusUpgraded)
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
		it.Event = new(SPPlusUpgraded)
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
func (it *SPPlusUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SPPlusUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SPPlusUpgraded represents a Upgraded event raised by the SPPlus contract.
type SPPlusUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SPPlus *SPPlusFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SPPlusUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SPPlus.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SPPlusUpgradedIterator{contract: _SPPlus.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SPPlus *SPPlusFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SPPlusUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SPPlus.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SPPlusUpgraded)
				if err := _SPPlus.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SPPlus *SPPlusFilterer) ParseUpgraded(log types.Log) (*SPPlusUpgraded, error) {
	event := new(SPPlusUpgraded)
	if err := _SPPlus.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
