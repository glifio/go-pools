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

// PlusMetaData contains all meta data concerning the Plus contract.
var PlusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existingTokenId\",\"type\":\"uint256\"}],\"name\":\"AgentAlreadyHasToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"AgentChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AlreadyActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notBeneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiaryIsNotAnAgent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"BeneficiaryOwnerIsNotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CredentialUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidDowngrade\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidUpgrade\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"MaxCashBackPercentExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoCashBackToClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPriceSetterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dtl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"OverLimitDtl\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenNotActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroCashBackPremium\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroConversionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"CardActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"CardDowngraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"CardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"CardUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CashBackClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FilVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"GlfVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"glfBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filEarned\",\"type\":\"uint256\"}],\"name\":\"PaymentProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCashBack\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCashBack\",\"type\":\"uint256\"}],\"name\":\"PersonalCashBackUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPriceSetter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPriceSetter\",\"type\":\"address\"}],\"name\":\"PriceSetterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentFactory\",\"outputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToCardOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentPoliceVcVerifier\",\"outputs\":[{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseConversionRateFILtoGLF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"changeOwnerForAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimCashBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credential\",\"type\":\"bytes32\"}],\"name\":\"credentialHasBeenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"_sc\",\"type\":\"tuple\"}],\"name\":\"downgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundFilVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fundGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"fundGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getTimeUntilPenaltyFreeDowngrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glfToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"},{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"_agentPoliceVcVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_personalCashBackPercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintActivateAndFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"mintAndActivate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"onDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestAmount\",\"type\":\"uint256\"}],\"name\":\"onPaymentMade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"}],\"name\":\"setAgentFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"_agentPoliceVcVerifier\",\"type\":\"address\"}],\"name\":\"setAgentPolice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseConversionRateFILtoGLF\",\"type\":\"uint256\"}],\"name\":\"setBaseConversionRateFILtoGLF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"}],\"name\":\"setGlfToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxCashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setMaxCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setPersonalCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"setPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceSetter\",\"type\":\"address\"}],\"name\":\"setPriceSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"_tierInfo\",\"type\":\"tuple\"}],\"name\":\"setTierInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyFee\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyWindow\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tierInfoToWithdrawableExtraLockedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawGlf\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"tierToTierInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackPremium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToAgentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToFilCashbackEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToGlfVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToLastTierSwitchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToPersonalCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToTier\",\"outputs\":[{\"internalType\":\"enumTier\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToTierLockAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFilCashbackVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawAllFilFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawExtraLockedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f516020617da75f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051617ce090816100c782396080518181816140ae015261422e0152f35b6001600160401b0319166001600160401b039081175f516020617da75f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f358060e01c91826301ffc9a714615e345750816302adafe114615dcd5781630431a32514615d5c57816306fdde0314615c18578163081812fc14615b91578163095ea7b314615b315781630c1cd64314615ad357816314b77f22146159795781631755ff211461590957816318de0c6a146158a25781631fac65d314614eda575080631fe29f1814614e7f57806323b872dd14614e4557806328d0e5a91461498c5780632c09bef71461491b5780632fe5c820146148b45780633f4ba83a146147e257806342842e0e1461479857806343989f0a1461472c5780634449ee09146146c5578063463ad343146144f65780634f1ef286146141c35780634fa56e3b1461415257806351c9749c1461412657806352d1902d146140685780635c975abb146140085780635df3548114613f7b5780635e7277f314613f1457806361d027b314613ea35780636352211e14613e6657806363791e3c14613d2b5780636817c76c14613cd05780636a62784214613c8e5780636b12bee314613c3357806370a0823114613b5d578063715018a614613a1b578063762b977d1461394057806379ba50971461389e5780637caf9073146137c35780637df107ea146137525780637dfe5b921461339c57806381eebd701461333c5780638456cb591461326757806386cc34371461318c5780638b446431146131595780638da5cb5b146130e85780638dd404d21461308d5780638f68679f146130325780639129347314612f555780639202adfc14612def57806395bda1d814612db257806395d89b4114612c555780639780b37214612bee578063a1e80e5514612b93578063a22cb46514612b27578063abdf23df14612a95578063ad3cb1cc14612a16578063ad7f4929146129bb578063aef0353c146123f8578063b367285614612344578063b3d0d7b714612307578063b77da53314612269578063b88d4fde146121df578063c87b56dd14612166578063cda43c10146120e6578063d0911b1e14612075578063d452154814610ff6578063d7d2929e14610d87578063d86e1a3514610c53578063e30c397814610be2578063e985e9c514610b31578063e9e15b4f14610a48578063ebca88fc146107df578063f0f44260146106d2578063f173e1fe1461066b578063f2fde38b14610546578063f4a0a528146104e6578063f9d20f53146103ea5763ffc2108914610385575f80fd5b346103e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e7576103bc615fba565b6103c4615fdd565b906044359260048410156103e75760206103df858585616d2f565b604051908152f35b80fd5b50346103e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760043560048110156104e25760607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126104e25780801580801592836104cd575b61046691616317565b61046e6172f5565b906104c2575b1561049a5761048290616350565b60243581556044356001820155600260643591015580f35b6004827f1db5b8a7000000000000000000000000000000000000000000000000000000008152fd5b506024351515610474565b61046691506104db816173e6565b915061045d565b5080fd5b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75761051e6172f5565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713005580f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75773ffffffffffffffffffffffffffffffffffffffff61059b610596615fba565b617422565b6105a36172f5565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e83522054604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75761071561070d615fba565b6105966172f5565b73ffffffffffffffffffffffffffffffffffffffff8116156107b7576107b49073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130355565b80f35b6004827fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b50346103e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602435600435604435610820617361565b8215610a20577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529060209082906064908290899073ffffffffffffffffffffffffffffffffffffffff165af18015610a15576109e8575b508184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313602052604084206108e8848254616812565b90558184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260408420548285527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313602052836040862054869380155f146109d9575014806109d1575b156109bc5750507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c546109888183617518565b60405192835260208301527f57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c660403393a380f35b156109885750826109cc8261748a565b610988565b508015610955565b93505050506109cc8183617518565b610a099060203d602011610a0e575b610a018183616133565b810190616292565b6108b1565b503d6109f7565b6040513d87823e3d90fd5b6004847f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75773ffffffffffffffffffffffffffffffffffffffff610aa781610a99615fba565b610aa16172f5565b16617422565b1680156107b7576107b49073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130455565b50346103e75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757610b69615fba565b73ffffffffffffffffffffffffffffffffffffffff610b86615fdd565b91165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760043573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015610d47575b15610d1f578015610cf7577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f5580f35b6004827f247af9ce000000000000000000000000000000000000000000000000000000008152fd5b6004827fd6d9d129000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131254163314610cc7565b50346103e75760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757610dbf615fba565b60243590610dcb616000565b6064356004811015610ff257610de49160843593616d2f565b91610ded617361565b8115610a20577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015230602482015260448101849052949060209086906064908290859073ffffffffffffffffffffffffffffffffffffffff165af1948515610fe557602095610fca575b508381527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313855260408120610eb9848254616812565b90558381527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671310855260408120548482527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713138652836040832054839480155f14610fb557501480610fad575b15610f91575050507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c54610f588184617518565b60405191825283820152817f57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c660403393a3604051908152f35b610f9c575b50610f58565b9050610fa78361748a565b5f610f96565b508015610f24565b945050505050610fc58184617518565b610f58565b610fe090863d8811610a0e57610a018183616133565b610e83565b50604051903d90823e3d90fd5b8480fd5b50346103e7576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760043567ffffffffffffffff81116104e2576110479036906004016161ae565b60243567ffffffffffffffff8111612071576110679036906004016161ae565b61106f616000565b9060643560843573ffffffffffffffffffffffffffffffffffffffff811680910361206d5760a4359073ffffffffffffffffffffffffffffffffffffffff82169283830361206d5760c4359373ffffffffffffffffffffffffffffffffffffffff851680950361206d5760e4359573ffffffffffffffffffffffffffffffffffffffff871680970361206d57610104359773ffffffffffffffffffffffffffffffffffffffff891680990361206d577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549960ff8b60401c16159a67ffffffffffffffff811680159081612065575b600114908161205b575b159081612052575b5061202a578b60017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055611fd5575b506111cc617bba565b6111d4617bba565b80519067ffffffffffffffff8211611fa8578c82916112137f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300546161f2565b601f8111611edb575b50602091601f8411600114611de05792611dd5575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b815167ffffffffffffffff8111611da8576112be7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301546161f2565b601f8111611d21575b5060208c601f8311600114611c3d579061133b9483611c32575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b610596617bba565b611343617bba565b61134b617bba565b73ffffffffffffffffffffffffffffffffffffffff811615611c065761137090617706565b611378617bba565b611380617bba565b8115611bde5715611bb6578115611bb6578315611bb6578415611bb6578515611bb65761158573ffffffffffffffffffffffffffffffffffffffff9561059687611505611605986105966114866116859d9a610596866114066105969d829d7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130055617422565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130155565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130355565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130255565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130555565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130455565b60017feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671306556101f47feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c55660e35fa931a00007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f556117b273ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713125416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131255565b62278d007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a556101f47feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b55604051611809816160ea565b8281526020808201848152670a688906bd8b000060408085019182525f80527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890935292517f98ececc9ef5447155a71900bf0a10a038a903deb6b29c18bf41c012bdacdfd7855517f98ececc9ef5447155a71900bf0a10a038a903deb6b29c18bf41c012bdacdfd795590517f98ececc9ef5447155a71900bf0a10a038a903deb6b29c18bf41c012bdacdfd7a55516118c1816160ea565b670de0b6b3a76400008152690a968163f0a57b4000006020808301918252670b1a2bc2ec500000604080850191825260015f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890925292517f2f7f69aa2d75e15f19ce52af93df9a7b1c1e6ef7db7d165311b61bd2c161868d5590517f2f7f69aa2d75e15f19ce52af93df9a7b1c1e6ef7db7d165311b61bd2c161868e5590517f2f7f69aa2d75e15f19ce52af93df9a7b1c1e6ef7db7d165311b61bd2c161868f555161198e816160ea565b670e4b4b8af6a7000081526934f086f3b33b684000006020808301918252670bbe2470a1549d8a604080850191825260025f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890925292517fcca1d77c34f404c66c45b233ec66db9f5c52d27866d7a7c46dbd420c1ff4d8cf5590517fcca1d77c34f404c66c45b233ec66db9f5c52d27866d7a7c46dbd420c1ff4d8d05590517fcca1d77c34f404c66c45b233ec66db9f5c52d27866d7a7c46dbd420c1ff4d8d15551611a5b816160ea565b670f43fc2c04ee0000815269d3c21bcecceda10000006020808301918252670c249fdd327780006040840190815260035f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890915291517fe6edec82ca249256f7285fb9fd1ebbc810fa32d98aed0f790e077004d6c9b18a55517fe6edec82ca249256f7285fb9fd1ebbc810fa32d98aed0f790e077004d6c9b18b55517fe6edec82ca249256f7285fb9fd1ebbc810fa32d98aed0f790e077004d6c9b18c55611b225780f35b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6004887fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b6004897f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b60248a7f1e4fbdf700000000000000000000000000000000000000000000000000000000815280600452fd5b015190505f806112e1565b9192937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084167f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793018452828420935b818110611d09575091600193918561133b97969410611cd2575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930155611333565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f8080611ca5565b92936020600181928786015181550195019301611c8b565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793018d527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f830160051c81019160208410611d9e575b601f0160051c01908d905b828110611d915750506112c7565b6001918155018d90611d83565b9091508190611d78565b60248c7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b015190505f80611231565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0167f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930084527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81935b818110611ec35750908460019594939210611e8c575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055611283565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690555f8080611e5f565b92936020600181928786015181550195019301611e49565b917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930091935052601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410611f80575b918e91601f85940160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b818110611f70575061121c565b9283558493508f92600101611f63565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf819150611f34565b60248d7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00555f6111c3565b60048d7ff92ee8a9000000000000000000000000000000000000000000000000000000008152fd5b9050155f611170565b303b159150611168565b8d915061115e565b5f80fd5b8280fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416604051908152f35b50346103e7576120f5366160b6565b9073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416330361213e57906107b491616927565b6004837f4b602735000000000000000000000000000000000000000000000000000000008152fd5b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e7576121a1600435617261565b50806040516121b1602082616133565b52506121db6040516121c4602082616133565b5f8152604051918291602083526020830190615f77565b0390f35b50346103e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757612217615fba565b50612220615fdd565b5060643567ffffffffffffffff81116104e25790612243600492369084016161ae565b507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760043560048110156104e2576122c990606092604080516122b9816160ea565b8281528260208201520152616350565b6040516122d5816160ea565b815491828252604060026001830154926020850193845201549201918252604051928352516020830152516040820152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760206103df6004356168ad565b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760043580156123d0576040826123b292602094527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130984522054617261565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b6004827fac8fb3c1000000000000000000000000000000000000000000000000000000008152fd5b50346103e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757612430615fba565b6024356044359160048310156129b75761248b9061244c617361565b61247961245884617261565b84339173ffffffffffffffffffffffffffffffffffffffff33911614616243565b61059684612486816173e6565b616317565b612494826173b4565b61298b5760016124a384616350565b01547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905291929190602090829060649082908a9073ffffffffffffffffffffffffffffffffffffffff165af1801561298057612963575b508285527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e60205260408520548386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d6020526040862054928487527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e6020524260408820558487527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526125ee8660408920616514565b8487527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671316602052604087205573ffffffffffffffffffffffffffffffffffffffff821691826126a65750505080612676575061264d6040518093615f6a565b7f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a380f35b836044917f0c6dc9c900000000000000000000000000000000000000000000000000000000825260045280602452fd5b90602492959391602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051958680927ffd66091e0000000000000000000000000000000000000000000000000000000082528b60048301525afa93841561292257889461292d575b506127349084151561654b565b6040517f8da5cb5b0000000000000000000000000000000000000000000000000000000081526020816004818a5afa9081156129225788916128f3575b5073ffffffffffffffffffffffffffffffffffffffff339116036128c3578287527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671309602052604087205480612893575015801561288a575b1561285c57507f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5918160209287527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130983528460408820558487527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d835260408720556128586040518092615f6a565ba380f35b7f0c6dc9c9000000000000000000000000000000000000000000000000000000008652600452602452604484fd5b508082146127c9565b87604491857f04e662d6000000000000000000000000000000000000000000000000000000008352600452602452fd5b604487877f98f08d3a00000000000000000000000000000000000000000000000000000000825260045233602452fd5b612915915060203d60201161291b575b61290d8183616133565b810190616595565b5f612771565b503d612903565b6040513d8a823e3d90fd5b9093506020813d60201161295b575b8161294960209383616133565b8101031261206d575192612734612727565b3d915061293c565b61297b9060203d602011610a0e57610a018183616133565b612533565b6040513d88823e3d90fd5b602484837f11d5a84c000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f54604051908152f35b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757506121db604051612a57604082616133565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190615f77565b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757600435612ad06172f5565b612710811015612aff577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b5580f35b6004827f58d620b3000000000000000000000000000000000000000000000000000000008152fd5b50346103e75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757612b5f615fba565b50602435801515036103e757807f8cd22d190000000000000000000000000000000000000000000000000000000060049252fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b54604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131083522054604051908152f35b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760405190807f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015490612cb5826161f2565b8085529160018116908115612d6d5750600114612cf1575b6121db84612cdd81860382616133565b604051918291602083526020830190615f77565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930181527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b808210612d5357509091508101602001612cdd82612ccd565b919260018160209254838588010152019101909291612d3a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b85019092019250612cdd9150839050612ccd565b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760206103df60043561681f565b50346103e75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760043573ffffffffffffffffffffffffffffffffffffffff612e63612e42615fdd565b612e4a617361565b610596612e5685617261565b8533918633911614616243565b16908115612f2d578083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131460205260408320548015612f015760207f05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700918386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713148252856040812055612ef78186617681565b604051908152a380f35b602484837fa370a06b000000000000000000000000000000000000000000000000000000008252600452fd5b6004837fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b50807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757612f87617361565b341561300a57612fb8347feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131554616812565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131555604051348152817f1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d607460203393a380f35b807f1f2a20050000000000000000000000000000000000000000000000000000000060049252fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130654604051908152f35b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c54604051908152f35b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75780f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75773ffffffffffffffffffffffffffffffffffffffff6131dd81610a99615fba565b1680156107b7576107b49073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130555565b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75761329e6172f5565b6132a6617361565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e7576133746172f5565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c5580f35b50346103e75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757600435602435906004821015612071576133e4617361565b6134116133f082617261565b82339173ffffffffffffffffffffffffffffffffffffffff33911614616243565b6134238161341e816173b4565b6162e4565b61343082612486816173e6565b8083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130760205260ff60408420541691600483101561372557828111156136eb5783908282527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671316602052604082205460016134aa83616350565b01549073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416908083115f1461364a576020916134ff61354892856162aa565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019190915295869283919082906064820190565b03925af1928315612980576136029361362b575b505b8386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526135928260408820616514565b8386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e6020524260408720558386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131660205260408620556135f86040518095615f6a565b6020840190615f6a565b7f7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd60403393a380f35b6136439060203d602011610a0e57610a018183616133565b505f61355c565b80831061365e575b5050613602925061355e565b60209161366e846136b1936162aa565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481019190915295869283919082906044820190565b03925af192831561298057613602936136cc575b8693613652565b6136e49060203d602011610a0e57610a018183616133565b505f6136c5565b8361372360449261371e867fa0f244d1000000000000000000000000000000000000000000000000000000008552615f4e565b615f5c565bfd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75773ffffffffffffffffffffffffffffffffffffffff61381481610a99615fba565b1680156107b7576107b49073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130255565b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e7573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00541603613914576107b433617706565b807f118cdaa7000000000000000000000000000000000000000000000000000000006024925233600452fd5b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75773ffffffffffffffffffffffffffffffffffffffff61399181610a99615fba565b1680156107b7576107b49073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130155565b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757613a526172f5565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00558073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757613b98610596615fba565b9073ffffffffffffffffffffffffffffffffffffffff821615613c07576020613bfe8373ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b54604051908152f35b807f89c62b64000000000000000000000000000000000000000000000000000000006024925280600452fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131554604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760206103df613ccb615fba565b6165c1565b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130054604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757613d6661070d615fba565b73ffffffffffffffffffffffffffffffffffffffff8116908115612f2d57613e3f73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131254169173ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713125416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131255565b7f379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa400798380a380f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760206123b2600435617261565b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131383522054604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75773ffffffffffffffffffffffffffffffffffffffff613fd3613fcb615fba565b61070d6175fe565b1680156107b7576107b4904790837feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131555617681565b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036140fe5760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b807fe07c8dba0000000000000000000000000000000000000000000000000000000060049252fd5b50346103e7576107b4614138366160b6565b90614141617361565b61414d6133f082617261565b617518565b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416604051908152f35b5060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e7576141f6615fba565b9060243567ffffffffffffffff81116104e2576142179036906004016161ae565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168030149081156144b4575b5061448c576142666172f5565b73ffffffffffffffffffffffffffffffffffffffff831690604051937f52d1902d000000000000000000000000000000000000000000000000000000008552602085600481865afa80958596614458575b506142e857602484847f4c9c8ce3000000000000000000000000000000000000000000000000000000008252600452fd5b9091847f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc810361442d5750813b1561440257807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8480a281518390156143cf57808360206143cb95519101845af46143c5617652565b91617c11565b5080f35b505050346143da5780f35b807fb398979f0000000000000000000000000000000000000000000000000000000060049252fd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000008452600452602483fd5b7faa1d49a4000000000000000000000000000000000000000000000000000000008552600452602484fd5b9095506020813d602011614484575b8161447460209383616133565b8101031261206d5751945f6142b7565b3d9150614467565b6004827fe07c8dba000000000000000000000000000000000000000000000000000000008152fd5b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f614259565b50346103e757614505366160b6565b9061450e617361565b614516617361565b811561469d577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529060209082906064908290889073ffffffffffffffffffffffffffffffffffffffff165af1801561469257614675575b508083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313602052604083206145de838254616812565b90558083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260408320548184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260408420549083859214806109d157156109bc5750507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c546109888183617518565b61468d9060203d602011610a0e57610a018183616133565b6145a7565b6040513d86823e3d90fd5b6004837f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d83522054604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760ff604060209260043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131184522054166040519015158152f35b50346103e7576004906147aa36616044565b505050806040516147bc602082616133565b527f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e7576148196172f5565b6148216175fe565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131483522054604051908152f35b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713125416604051908152f35b50346103e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e7576149cf6149c7615fba565b610596617361565b9073ffffffffffffffffffffffffffffffffffffffff82169182156107b7576024602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051928380927f1ffbb0640000000000000000000000000000000000000000000000000000000082528860048301525afa908115614def5790614a7692918491614e26575b5061654b565b604051917f8da5cb5b000000000000000000000000000000000000000000000000000000008352602083600481845afa928315614e1b578293614dfa575b5073ffffffffffffffffffffffffffffffffffffffff8316908115612f2d576020600491604051928380927faf640d0f0000000000000000000000000000000000000000000000000000000082525afa908115614def578391614dbd575b5082527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260408220549073ffffffffffffffffffffffffffffffffffffffff614b5d83617261565b1693818514614d955760209460405191614b778784616133565b85835273ffffffffffffffffffffffffffffffffffffffff614b9986836179a3565b1680614bcb57602487877f7e273289000000000000000000000000000000000000000000000000000000008252600452fd5b82819694959603614d6157503b614be0578480f35b908591614c3260405194859384937f150b7a0200000000000000000000000000000000000000000000000000000000855233600486015260248501526044840152608060648401526084830190615f77565b038186865af1839181614d09575b50614c885750614c4e617652565b80519384614c8257602484847f64a0ae92000000000000000000000000000000000000000000000000000000008252600452fd5b84925001fd5b7f150b7a0200000000000000000000000000000000000000000000000000000000919293507fffffffff000000000000000000000000000000000000000000000000000000001603614cde5750805f8080808480f35b7f64a0ae92000000000000000000000000000000000000000000000000000000008252600452602490fd5b9091508481813d8311614d5a575b614d218183616133565b810103126129b757517fffffffff00000000000000000000000000000000000000000000000000000000811681036129b757905f614c40565b503d614d17565b8660649185857f64283d7b000000000000000000000000000000000000000000000000000000008452600452602452604452fd5b6004847f2a63c7cc000000000000000000000000000000000000000000000000000000008152fd5b90506020813d602011614de7575b81614dd860209383616133565b8101031261206d57515f614b12565b3d9150614dcb565b6040513d85823e3d90fd5b614e1491935060203d60201161291b5761290d8183616133565b915f614ab4565b6040513d84823e3d90fd5b614e3f915060203d602011610a0e57610a018183616133565b5f614a70565b50346103e757600490614e5736616044565b5050507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b50346103e757807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103e75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a54604051908152f35b90503461206d5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d57600435602435600481101561206d576044359267ffffffffffffffff841161206d5783600401908436039160807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc84011261206d57614f68617361565b614f95614f7486617261565b86339173ffffffffffffffffffffffffffffffffffffffff33911614616243565b614fa28561341e816173b4565b831584811561588d575b614fb591616317565b845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130760205260ff60405f205416956004871015615860578685101561582b57908794939291865f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d60205260405f20549261503286616350565b9460405192615040846160ea565b865484526040600260018901549860208701998a52015494019384528561539d575b5050505050508383527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131660205260408320549051908484527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e6020526150cc6040852054426162aa565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015473ffffffffffffffffffffffffffffffffffffffff16908383101561518f575061511e906134ff602093856162aa565b03925af1928315612980576151669361362b57508386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526135928260408820616514565b7f65e7b85a0b1fe94cdd38e2a9dbe4ac14dda1ceeb998856160d0ea99b804b105d60403393a380f35b918381116151a5575b505050615166925061355e565b836151af916162aa565b917feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a54115f146153255761529e9061271061520b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b5485616501565b0460208173ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303541660405195869283927fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03818a865af190811561531a576020946152c19461366e936152ff575b506162aa565b03925af192831561298057615166936152e0575b505b85925f80615198565b6152f89060203d602011610a0e57610a018183616133565b505f6152d5565b61531590873d8911610a0e57610a018183616133565b6152bb565b6040513d89823e3d90fd5b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081523360048201526024810192909252909360209185916044918391905af1928315612980576151669361537e575b506152d7565b6153969060203d602011610a0e57610a018183616133565b505f615378565b6154239697985073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130554169160206153e98780616386565b604051809a81927fa2f3c21000000000000000000000000000000000000000000000000000000000835284600484015260248301906163ce565b0381865afa9788156157c4575f986157f7575b50875f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131160205260ff60405f2054166157cf57823b1561206d577fffffffff00000000000000000000000000000000000000000000000000000000604051947f403bb79d000000000000000000000000000000000000000000000000000000008652886004870152166024850152606060448501527ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffefd8635910181121561206d57615512906080606486015260048360e487019201016163ce565b9160248201359260ff841680940361206d575f858094926064829484986084850152604481013560a4850152013560c483015203925af180156157c4576157ad575b5061556190519180616386565b602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416936024604051809681937f9f4326d700000000000000000000000000000000000000000000000000000000835260048301525afa9283156157a257899361576e575b506101008136031261576a57604051610100810181811067ffffffffffffffff82111761573d5760405261561082616023565b81526020820135602082015260408201356040820152606082013560608201526080820135608082015261564660a08301615f21565b60a082015261565760c083016163b9565b60c082015260e082013567ffffffffffffffff81116157395761568b92615680913691016161ae565b60e082015283617815565b8282116156f857505050509085929183527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713116020526040832060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790555f8080808080615062565b60849450604051937f4ec0c37d0000000000000000000000000000000000000000000000000000000085526004850152602484015260448301526064820152fd5b8a80fd5b60248b7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8880fd5b9092506020813d60201161579a575b8161578a60209383616133565b8101031261206d5751915f6155dd565b3d915061577d565b6040513d8b823e3d90fd5b6157ba9199505f90616133565b5f97615561615554565b6040513d5f823e3d90fd5b7fae9e9f65000000000000000000000000000000000000000000000000000000005f5260045ffd5b9097506020813d602011615823575b8161581360209383616133565b8101031261206d5751965f615436565b3d9150615806565b61585b8561371e897f253471cf000000000000000000000000000000000000000000000000000000005f52615f4e565b60445ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b614fb5915061589b816173e6565b9150614fac565b3461206d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671309602052602060405f2054604051908152f35b3461206d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d57602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416604051908152f35b3461206d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d576004356159b3617361565b6159bf6133f082617261565b6159c8816168ad565b908115615aab577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081523360048201526024810184905290602090829060449082905f9073ffffffffffffffffffffffffffffffffffffffff165af180156157c457615a8e575b505f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671316602052615a8a60405f209182546162aa565b9055005b615aa69060203d602011610a0e57610a018183616133565b615a54565b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461206d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d57615b0a6172f5565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a55005b3461206d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d57615b68615fba565b507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461206d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d57600435615bcc81617261565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b3461206d575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054615c75816161f2565b8084529060018116908115615d1a5750600114615c9d575b6121db83612cdd81850382616133565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b808210615d0057509091508101602001612cdd615c8d565b919260018160209254838588010152019101909291615ce8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b84019091019150612cdd9050615c8d565b3461206d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671307602052602060ff60405f205416615dcb6040518092615f6a565bf35b3461206d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671316602052602060405f2054604051908152f35b3461206d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261206d57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361206d57817f80ac58cd0000000000000000000000000000000000000000000000000000000060209314908115615ef7575b8115615ecd575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483615ec6565b7f5b5e139f0000000000000000000000000000000000000000000000000000000081149150615ebf565b35907fffffffff000000000000000000000000000000000000000000000000000000008216820361206d57565b600481101561586057600452565b600481101561586057602452565b9060048210156158605752565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361206d57565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361206d57565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361206d57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361206d57565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc606091011261206d5760043573ffffffffffffffffffffffffffffffffffffffff8116810361206d579060243573ffffffffffffffffffffffffffffffffffffffff8116810361206d579060443590565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc604091011261206d576004359060243590565b6060810190811067ffffffffffffffff82111761610657604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761610657604052565b67ffffffffffffffff811161610657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561206d576020813591016161c882616174565b926161d66040519485616133565b8284528282011161206d57815f92602092838601378301015290565b90600182811c92168015616239575b602083101461620c57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691616201565b1561624c575050565b73ffffffffffffffffffffffffffffffffffffffff92507f3a0302a9000000000000000000000000000000000000000000000000000000005f526004521660245260445ffd5b9081602091031261206d5751801515810361206d5790565b919082039182116162b757565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b156162ec5750565b7fed5a980a000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b1561631f5750565b61634b907fbca1a956000000000000000000000000000000000000000000000000000000005f52615f4e565b60245ffd5b6004811015615860575f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130860205260405f2090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018136030182121561206d570190565b359067ffffffffffffffff8216820361206d57565b73ffffffffffffffffffffffffffffffffffffffff6163ec82616023565b168252602081013560208301526040810135604083015260608101356060830152608081013560808301527fffffffff0000000000000000000000000000000000000000000000000000000061644460a08301615f21565b1660a083015267ffffffffffffffff61645f60c083016163b9565b1660c083015260e08101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561206d5701906020823592019167ffffffffffffffff811161206d57803603831361206d57601f817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0926101209561010060e087015281610100870152868601375f8582860101520116010190565b818102929181159184041417156162b757565b9060048110156158605760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b156165535750565b73ffffffffffffffffffffffffffffffffffffffff907f3a35a1b9000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b9081602091031261206d575173ffffffffffffffffffffffffffffffffffffffff8116810361206d5790565b6165cd90610596617361565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671300546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9283166024820152604481019190915291602091839160649183915f91165af180156157c4576167f5575b507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82146162b757600182017feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713065573ffffffffffffffffffffffffffffffffffffffff8116156167c95773ffffffffffffffffffffffffffffffffffffffff61674b83836179a3565b1661679d576040805133815273ffffffffffffffffffffffffffffffffffffffff909216602083015282917f864b3215ee93b2a9291e17b45d665127a426c1259666cf483c30939c4a7103329190a290565b7f73c6ac6e000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b7f64a0ae92000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b61680d9060203d602011610a0e57610a018183616133565b61669f565b919082018092116162b757565b61682881617261565b50616832816173b4565b156168a8575f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e60205261686b60405f2054426162aa565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a5490818110156168a25761689f916162aa565b90565b50505f90565b505f90565b805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130760205260016168e660ff60405f205416616350565b0154905f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131660205260405f2054908082115f146168a25761689f916162aa565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416616d2b578115616d2b57805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260405f2054918215616d2657616992836173b4565b15616d2657825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260405f20547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c548082105f14616d1e5750905b8115616d1857835f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671307602052670de0b6b3a7640000612710616a68616a3d60ff60405f205416616350565b94600260405196616a4d886160ea565b80548852600181015460208901520154604087015284616501565b04927feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f549051020490818302670de0b6b3a7640000808204910615150191855f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260405f2054838110616d00575b507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713155490848210616cdf575b50508115616cd857616c7292855f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260405f20616b488482546162aa565b9055616b75817feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671315546162aa565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131555855f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131460205260405f20616bcd828254616812565b90557feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015260248101869052956020928792169082905f9082906044820190565b03925af19283156157c4577f299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d694606094616cbb575b5060405192835260208301526040820152a3565b616cd39060203d602011610a0e57610a018183616133565b616ca7565b5050505050565b9192508093509202670de0b6b3a76400008082049106151501905f80616b06565b9250809350670de0b6b3a7640000830204925f616ada565b50505050565b9050906169f1565b505050565b5050565b616d3e616d81919392936165c1565b92616d47617361565b616d74616d5385617261565b85339173ffffffffffffffffffffffffffffffffffffffff33911614616243565b61059683612486816173e6565b616d8a836173b4565b617235576001616d9983616350565b01547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905291929190602090829060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af180156157c457617218575b505f8481527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e6020908152604080832080547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d845282852054429092557feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671307909352922091939091616ebb908690616514565b855f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131660205260405f205573ffffffffffffffffffffffffffffffffffffffff82169182155f14616f765750505080616f475750616f1d6040518092615f6a565b817f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a390565b7f0c6dc9c9000000000000000000000000000000000000000000000000000000005f526004525f60245260445ffd5b90602492949391602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051958680927ffd66091e0000000000000000000000000000000000000000000000000000000082528a60048301525afa9384156157c4575f946171e2575b506170049084151561654b565b6040517f8da5cb5b000000000000000000000000000000000000000000000000000000008152602081600481895afa9081156157c4575f916171c3575b5073ffffffffffffffffffffffffffffffffffffffff3391160361719357825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260405f205480617164575015801561715b575b1561712d575060208492827f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5935f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130983528460405f2055845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d835260405f20556171296040518092615f6a565ba390565b7f0c6dc9c9000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b50808214617099565b837f04e662d6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b847f98f08d3a000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b6171dc915060203d60201161291b5761290d8183616133565b5f617041565b9093506020813d602011617210575b816171fe60209383616133565b8101031261206d575192617004616ff7565b3d91506171f1565b6172309060203d602011610a0e57610a018183616133565b616e29565b827f11d5a84c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b6172a9815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff8216156172ca575090565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416330361733557565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300541661738c57565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130760205261689f60ff60405f2054165b600481101590816158605760018114918215617415575b821561740857505090565b9091506158605760031490565b506002811491505f6173fd565b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014617476575b15617472576174649061792a565b9061746d575090565b905090565b5090565b505067ffffffffffffffff81166001617456565b805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260405f205490805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020525f60408120556040519182525f60208301527fdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca60403393a3565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c548083116175cf5750805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260405f205491815f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020528060405f205560405192835260208301527fdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca60403393a3565b827f37829ee6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054161561762a57565b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b3d1561767c573d9061766382616174565b916176716040519384616133565b82523d5f602084013e565b606090565b8147106176de575f80809373ffffffffffffffffffffffffffffffffffffffff8294165af16176ae617652565b50156176b657565b7f3204506f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f356680b7000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b919060e00151916101408380518101031261206d57604051610140810181811067ffffffffffffffff8211176161065760405260208401518152610120610140604086015195866020850152606081015160408501526080810151606085015260a0810151608085015260c081015160a085015260e081015160c085015261010081015160e0850152828101516101008501520151910152828115617922575082156178fc577812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f2281108302156178ef57670de0b6b3a7640000839102049190565b637c5f487d5f526004601cfd5b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9190565b9250505f9190565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff1603617998575b8215801561798d575b61798557565b5f9250829150565b5060163d141561797f565b5f9250829150617976565b906179ec815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9173ffffffffffffffffffffffffffffffffffffffff831680617af7575b73ffffffffffffffffffffffffffffffffffffffff82169182617aa3575b50825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f20827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a490565b617aea9073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190555f617a28565b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055617b8f8473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8154019055617a0a565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615617be957565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b90617c4e5750805115617c2657805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580617ca1575b617c5f575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15617c5756fea2646970667358221220c071d17e52e2a29ff768ccf924d492a761a1b9b473d1be851d140385ee09c79764736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// PlusABI is the input ABI used to generate the binding from.
// Deprecated: Use PlusMetaData.ABI instead.
var PlusABI = PlusMetaData.ABI

// PlusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PlusMetaData.Bin instead.
var PlusBin = PlusMetaData.Bin

// DeployPlus deploys a new Ethereum contract, binding an instance of Plus to it.
func DeployPlus(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Plus, error) {
	parsed, err := PlusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PlusBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Plus{PlusCaller: PlusCaller{contract: contract}, PlusTransactor: PlusTransactor{contract: contract}, PlusFilterer: PlusFilterer{contract: contract}}, nil
}

// Plus is an auto generated Go binding around an Ethereum contract.
type Plus struct {
	PlusCaller     // Read-only binding to the contract
	PlusTransactor // Write-only binding to the contract
	PlusFilterer   // Log filterer for contract events
}

// PlusCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlusSession struct {
	Contract     *Plus             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlusCallerSession struct {
	Contract *PlusCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PlusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlusTransactorSession struct {
	Contract     *PlusTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlusRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlusRaw struct {
	Contract *Plus // Generic contract binding to access the raw methods on
}

// PlusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlusCallerRaw struct {
	Contract *PlusCaller // Generic read-only contract binding to access the raw methods on
}

// PlusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlusTransactorRaw struct {
	Contract *PlusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlus creates a new instance of Plus, bound to a specific deployed contract.
func NewPlus(address common.Address, backend bind.ContractBackend) (*Plus, error) {
	contract, err := bindPlus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Plus{PlusCaller: PlusCaller{contract: contract}, PlusTransactor: PlusTransactor{contract: contract}, PlusFilterer: PlusFilterer{contract: contract}}, nil
}

// NewPlusCaller creates a new read-only instance of Plus, bound to a specific deployed contract.
func NewPlusCaller(address common.Address, caller bind.ContractCaller) (*PlusCaller, error) {
	contract, err := bindPlus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlusCaller{contract: contract}, nil
}

// NewPlusTransactor creates a new write-only instance of Plus, bound to a specific deployed contract.
func NewPlusTransactor(address common.Address, transactor bind.ContractTransactor) (*PlusTransactor, error) {
	contract, err := bindPlus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlusTransactor{contract: contract}, nil
}

// NewPlusFilterer creates a new log filterer instance of Plus, bound to a specific deployed contract.
func NewPlusFilterer(address common.Address, filterer bind.ContractFilterer) (*PlusFilterer, error) {
	contract, err := bindPlus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlusFilterer{contract: contract}, nil
}

// bindPlus binds a generic wrapper to an already deployed contract.
func bindPlus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plus *PlusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Plus.Contract.PlusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plus *PlusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plus.Contract.PlusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plus *PlusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plus.Contract.PlusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plus *PlusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Plus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plus *PlusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plus *PlusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plus.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Plus *PlusCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Plus *PlusSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Plus.Contract.UPGRADEINTERFACEVERSION(&_Plus.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Plus *PlusCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Plus.Contract.UPGRADEINTERFACEVERSION(&_Plus.CallOpts)
}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_Plus *PlusCaller) AgentFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "agentFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_Plus *PlusSession) AgentFactory() (common.Address, error) {
	return _Plus.Contract.AgentFactory(&_Plus.CallOpts)
}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_Plus *PlusCallerSession) AgentFactory() (common.Address, error) {
	return _Plus.Contract.AgentFactory(&_Plus.CallOpts)
}

// AgentIdToCardOwner is a free data retrieval call binding the contract method 0xb3672856.
//
// Solidity: function agentIdToCardOwner(uint256 _agentId) view returns(address owner)
func (_Plus *PlusCaller) AgentIdToCardOwner(opts *bind.CallOpts, _agentId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "agentIdToCardOwner", _agentId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentIdToCardOwner is a free data retrieval call binding the contract method 0xb3672856.
//
// Solidity: function agentIdToCardOwner(uint256 _agentId) view returns(address owner)
func (_Plus *PlusSession) AgentIdToCardOwner(_agentId *big.Int) (common.Address, error) {
	return _Plus.Contract.AgentIdToCardOwner(&_Plus.CallOpts, _agentId)
}

// AgentIdToCardOwner is a free data retrieval call binding the contract method 0xb3672856.
//
// Solidity: function agentIdToCardOwner(uint256 _agentId) view returns(address owner)
func (_Plus *PlusCallerSession) AgentIdToCardOwner(_agentId *big.Int) (common.Address, error) {
	return _Plus.Contract.AgentIdToCardOwner(&_Plus.CallOpts, _agentId)
}

// AgentIdToTokenId is a free data retrieval call binding the contract method 0x18de0c6a.
//
// Solidity: function agentIdToTokenId(uint256 agentId) view returns(uint256)
func (_Plus *PlusCaller) AgentIdToTokenId(opts *bind.CallOpts, agentId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "agentIdToTokenId", agentId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AgentIdToTokenId is a free data retrieval call binding the contract method 0x18de0c6a.
//
// Solidity: function agentIdToTokenId(uint256 agentId) view returns(uint256)
func (_Plus *PlusSession) AgentIdToTokenId(agentId *big.Int) (*big.Int, error) {
	return _Plus.Contract.AgentIdToTokenId(&_Plus.CallOpts, agentId)
}

// AgentIdToTokenId is a free data retrieval call binding the contract method 0x18de0c6a.
//
// Solidity: function agentIdToTokenId(uint256 agentId) view returns(uint256)
func (_Plus *PlusCallerSession) AgentIdToTokenId(agentId *big.Int) (*big.Int, error) {
	return _Plus.Contract.AgentIdToTokenId(&_Plus.CallOpts, agentId)
}

// AgentPoliceVcVerifier is a free data retrieval call binding the contract method 0xd0911b1e.
//
// Solidity: function agentPoliceVcVerifier() view returns(address)
func (_Plus *PlusCaller) AgentPoliceVcVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "agentPoliceVcVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentPoliceVcVerifier is a free data retrieval call binding the contract method 0xd0911b1e.
//
// Solidity: function agentPoliceVcVerifier() view returns(address)
func (_Plus *PlusSession) AgentPoliceVcVerifier() (common.Address, error) {
	return _Plus.Contract.AgentPoliceVcVerifier(&_Plus.CallOpts)
}

// AgentPoliceVcVerifier is a free data retrieval call binding the contract method 0xd0911b1e.
//
// Solidity: function agentPoliceVcVerifier() view returns(address)
func (_Plus *PlusCallerSession) AgentPoliceVcVerifier() (common.Address, error) {
	return _Plus.Contract.AgentPoliceVcVerifier(&_Plus.CallOpts)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_Plus *PlusCaller) Approve(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) error {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "approve", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_Plus *PlusSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _Plus.Contract.Approve(&_Plus.CallOpts, arg0, arg1)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns()
func (_Plus *PlusCallerSession) Approve(arg0 common.Address, arg1 *big.Int) error {
	return _Plus.Contract.Approve(&_Plus.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Plus *PlusCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Plus *PlusSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Plus.Contract.BalanceOf(&_Plus.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Plus *PlusCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Plus.Contract.BalanceOf(&_Plus.CallOpts, owner)
}

// BaseConversionRateFILtoGLF is a free data retrieval call binding the contract method 0xad7f4929.
//
// Solidity: function baseConversionRateFILtoGLF() view returns(uint256)
func (_Plus *PlusCaller) BaseConversionRateFILtoGLF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "baseConversionRateFILtoGLF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseConversionRateFILtoGLF is a free data retrieval call binding the contract method 0xad7f4929.
//
// Solidity: function baseConversionRateFILtoGLF() view returns(uint256)
func (_Plus *PlusSession) BaseConversionRateFILtoGLF() (*big.Int, error) {
	return _Plus.Contract.BaseConversionRateFILtoGLF(&_Plus.CallOpts)
}

// BaseConversionRateFILtoGLF is a free data retrieval call binding the contract method 0xad7f4929.
//
// Solidity: function baseConversionRateFILtoGLF() view returns(uint256)
func (_Plus *PlusCallerSession) BaseConversionRateFILtoGLF() (*big.Int, error) {
	return _Plus.Contract.BaseConversionRateFILtoGLF(&_Plus.CallOpts)
}

// CredentialHasBeenUsed is a free data retrieval call binding the contract method 0x43989f0a.
//
// Solidity: function credentialHasBeenUsed(bytes32 credential) view returns(bool)
func (_Plus *PlusCaller) CredentialHasBeenUsed(opts *bind.CallOpts, credential [32]byte) (bool, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "credentialHasBeenUsed", credential)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CredentialHasBeenUsed is a free data retrieval call binding the contract method 0x43989f0a.
//
// Solidity: function credentialHasBeenUsed(bytes32 credential) view returns(bool)
func (_Plus *PlusSession) CredentialHasBeenUsed(credential [32]byte) (bool, error) {
	return _Plus.Contract.CredentialHasBeenUsed(&_Plus.CallOpts, credential)
}

// CredentialHasBeenUsed is a free data retrieval call binding the contract method 0x43989f0a.
//
// Solidity: function credentialHasBeenUsed(bytes32 credential) view returns(bool)
func (_Plus *PlusCallerSession) CredentialHasBeenUsed(credential [32]byte) (bool, error) {
	return _Plus.Contract.CredentialHasBeenUsed(&_Plus.CallOpts, credential)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Plus *PlusCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Plus *PlusSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Plus.Contract.GetApproved(&_Plus.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Plus *PlusCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Plus.Contract.GetApproved(&_Plus.CallOpts, tokenId)
}

// GetTimeUntilPenaltyFreeDowngrade is a free data retrieval call binding the contract method 0x95bda1d8.
//
// Solidity: function getTimeUntilPenaltyFreeDowngrade(uint256 _tokenId) view returns(uint256 timeLeft)
func (_Plus *PlusCaller) GetTimeUntilPenaltyFreeDowngrade(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "getTimeUntilPenaltyFreeDowngrade", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeUntilPenaltyFreeDowngrade is a free data retrieval call binding the contract method 0x95bda1d8.
//
// Solidity: function getTimeUntilPenaltyFreeDowngrade(uint256 _tokenId) view returns(uint256 timeLeft)
func (_Plus *PlusSession) GetTimeUntilPenaltyFreeDowngrade(_tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.GetTimeUntilPenaltyFreeDowngrade(&_Plus.CallOpts, _tokenId)
}

// GetTimeUntilPenaltyFreeDowngrade is a free data retrieval call binding the contract method 0x95bda1d8.
//
// Solidity: function getTimeUntilPenaltyFreeDowngrade(uint256 _tokenId) view returns(uint256 timeLeft)
func (_Plus *PlusCallerSession) GetTimeUntilPenaltyFreeDowngrade(_tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.GetTimeUntilPenaltyFreeDowngrade(&_Plus.CallOpts, _tokenId)
}

// GlfToken is a free data retrieval call binding the contract method 0x4fa56e3b.
//
// Solidity: function glfToken() view returns(address)
func (_Plus *PlusCaller) GlfToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "glfToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlfToken is a free data retrieval call binding the contract method 0x4fa56e3b.
//
// Solidity: function glfToken() view returns(address)
func (_Plus *PlusSession) GlfToken() (common.Address, error) {
	return _Plus.Contract.GlfToken(&_Plus.CallOpts)
}

// GlfToken is a free data retrieval call binding the contract method 0x4fa56e3b.
//
// Solidity: function glfToken() view returns(address)
func (_Plus *PlusCallerSession) GlfToken() (common.Address, error) {
	return _Plus.Contract.GlfToken(&_Plus.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Plus *PlusCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Plus *PlusSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Plus.Contract.IsApprovedForAll(&_Plus.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Plus *PlusCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Plus.Contract.IsApprovedForAll(&_Plus.CallOpts, owner, operator)
}

// MaxCashBackPercent is a free data retrieval call binding the contract method 0x8dd404d2.
//
// Solidity: function maxCashBackPercent() view returns(uint256)
func (_Plus *PlusCaller) MaxCashBackPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "maxCashBackPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCashBackPercent is a free data retrieval call binding the contract method 0x8dd404d2.
//
// Solidity: function maxCashBackPercent() view returns(uint256)
func (_Plus *PlusSession) MaxCashBackPercent() (*big.Int, error) {
	return _Plus.Contract.MaxCashBackPercent(&_Plus.CallOpts)
}

// MaxCashBackPercent is a free data retrieval call binding the contract method 0x8dd404d2.
//
// Solidity: function maxCashBackPercent() view returns(uint256)
func (_Plus *PlusCallerSession) MaxCashBackPercent() (*big.Int, error) {
	return _Plus.Contract.MaxCashBackPercent(&_Plus.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Plus *PlusCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Plus *PlusSession) MintPrice() (*big.Int, error) {
	return _Plus.Contract.MintPrice(&_Plus.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Plus *PlusCallerSession) MintPrice() (*big.Int, error) {
	return _Plus.Contract.MintPrice(&_Plus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Plus *PlusCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Plus *PlusSession) Name() (string, error) {
	return _Plus.Contract.Name(&_Plus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Plus *PlusCallerSession) Name() (string, error) {
	return _Plus.Contract.Name(&_Plus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Plus *PlusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Plus *PlusSession) Owner() (common.Address, error) {
	return _Plus.Contract.Owner(&_Plus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Plus *PlusCallerSession) Owner() (common.Address, error) {
	return _Plus.Contract.Owner(&_Plus.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Plus *PlusCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Plus *PlusSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Plus.Contract.OwnerOf(&_Plus.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Plus *PlusCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Plus.Contract.OwnerOf(&_Plus.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Plus *PlusCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Plus *PlusSession) Paused() (bool, error) {
	return _Plus.Contract.Paused(&_Plus.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Plus *PlusCallerSession) Paused() (bool, error) {
	return _Plus.Contract.Paused(&_Plus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Plus *PlusCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Plus *PlusSession) PendingOwner() (common.Address, error) {
	return _Plus.Contract.PendingOwner(&_Plus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Plus *PlusCallerSession) PendingOwner() (common.Address, error) {
	return _Plus.Contract.PendingOwner(&_Plus.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Plus *PlusCaller) PoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "poolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Plus *PlusSession) PoolAddress() (common.Address, error) {
	return _Plus.Contract.PoolAddress(&_Plus.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_Plus *PlusCallerSession) PoolAddress() (common.Address, error) {
	return _Plus.Contract.PoolAddress(&_Plus.CallOpts)
}

// PriceSetter is a free data retrieval call binding the contract method 0x2c09bef7.
//
// Solidity: function priceSetter() view returns(address)
func (_Plus *PlusCaller) PriceSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "priceSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceSetter is a free data retrieval call binding the contract method 0x2c09bef7.
//
// Solidity: function priceSetter() view returns(address)
func (_Plus *PlusSession) PriceSetter() (common.Address, error) {
	return _Plus.Contract.PriceSetter(&_Plus.CallOpts)
}

// PriceSetter is a free data retrieval call binding the contract method 0x2c09bef7.
//
// Solidity: function priceSetter() view returns(address)
func (_Plus *PlusCallerSession) PriceSetter() (common.Address, error) {
	return _Plus.Contract.PriceSetter(&_Plus.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Plus *PlusCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Plus *PlusSession) ProxiableUUID() ([32]byte, error) {
	return _Plus.Contract.ProxiableUUID(&_Plus.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Plus *PlusCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Plus.Contract.ProxiableUUID(&_Plus.CallOpts)
}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_Plus *PlusCaller) SafeTransferFrom0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "safeTransferFrom0", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_Plus *PlusSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _Plus.Contract.SafeTransferFrom0(&_Plus.CallOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a free data retrieval call binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) pure returns()
func (_Plus *PlusCallerSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) error {
	return _Plus.Contract.SafeTransferFrom0(&_Plus.CallOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_Plus *PlusCaller) SetApprovalForAll(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "setApprovalForAll", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_Plus *PlusSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _Plus.Contract.SetApprovalForAll(&_Plus.CallOpts, arg0, arg1)
}

// SetApprovalForAll is a free data retrieval call binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) pure returns()
func (_Plus *PlusCallerSession) SetApprovalForAll(arg0 common.Address, arg1 bool) error {
	return _Plus.Contract.SetApprovalForAll(&_Plus.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Plus *PlusCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Plus *PlusSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Plus.Contract.SupportsInterface(&_Plus.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Plus *PlusCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Plus.Contract.SupportsInterface(&_Plus.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Plus *PlusCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Plus *PlusSession) Symbol() (string, error) {
	return _Plus.Contract.Symbol(&_Plus.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Plus *PlusCallerSession) Symbol() (string, error) {
	return _Plus.Contract.Symbol(&_Plus.CallOpts)
}

// TierInfoToWithdrawableExtraLockedFunds is a free data retrieval call binding the contract method 0xb3d0d7b7.
//
// Solidity: function tierInfoToWithdrawableExtraLockedFunds(uint256 _tokenId) view returns(uint256 withdrawGlf)
func (_Plus *PlusCaller) TierInfoToWithdrawableExtraLockedFunds(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tierInfoToWithdrawableExtraLockedFunds", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierInfoToWithdrawableExtraLockedFunds is a free data retrieval call binding the contract method 0xb3d0d7b7.
//
// Solidity: function tierInfoToWithdrawableExtraLockedFunds(uint256 _tokenId) view returns(uint256 withdrawGlf)
func (_Plus *PlusSession) TierInfoToWithdrawableExtraLockedFunds(_tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TierInfoToWithdrawableExtraLockedFunds(&_Plus.CallOpts, _tokenId)
}

// TierInfoToWithdrawableExtraLockedFunds is a free data retrieval call binding the contract method 0xb3d0d7b7.
//
// Solidity: function tierInfoToWithdrawableExtraLockedFunds(uint256 _tokenId) view returns(uint256 withdrawGlf)
func (_Plus *PlusCallerSession) TierInfoToWithdrawableExtraLockedFunds(_tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TierInfoToWithdrawableExtraLockedFunds(&_Plus.CallOpts, _tokenId)
}

// TierSwitchPenaltyFee is a free data retrieval call binding the contract method 0xa1e80e55.
//
// Solidity: function tierSwitchPenaltyFee() view returns(uint256)
func (_Plus *PlusCaller) TierSwitchPenaltyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tierSwitchPenaltyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierSwitchPenaltyFee is a free data retrieval call binding the contract method 0xa1e80e55.
//
// Solidity: function tierSwitchPenaltyFee() view returns(uint256)
func (_Plus *PlusSession) TierSwitchPenaltyFee() (*big.Int, error) {
	return _Plus.Contract.TierSwitchPenaltyFee(&_Plus.CallOpts)
}

// TierSwitchPenaltyFee is a free data retrieval call binding the contract method 0xa1e80e55.
//
// Solidity: function tierSwitchPenaltyFee() view returns(uint256)
func (_Plus *PlusCallerSession) TierSwitchPenaltyFee() (*big.Int, error) {
	return _Plus.Contract.TierSwitchPenaltyFee(&_Plus.CallOpts)
}

// TierSwitchPenaltyWindow is a free data retrieval call binding the contract method 0x1fe29f18.
//
// Solidity: function tierSwitchPenaltyWindow() view returns(uint256)
func (_Plus *PlusCaller) TierSwitchPenaltyWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tierSwitchPenaltyWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierSwitchPenaltyWindow is a free data retrieval call binding the contract method 0x1fe29f18.
//
// Solidity: function tierSwitchPenaltyWindow() view returns(uint256)
func (_Plus *PlusSession) TierSwitchPenaltyWindow() (*big.Int, error) {
	return _Plus.Contract.TierSwitchPenaltyWindow(&_Plus.CallOpts)
}

// TierSwitchPenaltyWindow is a free data retrieval call binding the contract method 0x1fe29f18.
//
// Solidity: function tierSwitchPenaltyWindow() view returns(uint256)
func (_Plus *PlusCallerSession) TierSwitchPenaltyWindow() (*big.Int, error) {
	return _Plus.Contract.TierSwitchPenaltyWindow(&_Plus.CallOpts)
}

// TierToTierInfo is a free data retrieval call binding the contract method 0xb77da533.
//
// Solidity: function tierToTierInfo(uint8 tier) view returns((uint256,uint256,uint256))
func (_Plus *PlusCaller) TierToTierInfo(opts *bind.CallOpts, tier uint8) (TierInfo, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tierToTierInfo", tier)

	if err != nil {
		return *new(TierInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TierInfo)).(*TierInfo)

	return out0, err

}

// TierToTierInfo is a free data retrieval call binding the contract method 0xb77da533.
//
// Solidity: function tierToTierInfo(uint8 tier) view returns((uint256,uint256,uint256))
func (_Plus *PlusSession) TierToTierInfo(tier uint8) (TierInfo, error) {
	return _Plus.Contract.TierToTierInfo(&_Plus.CallOpts, tier)
}

// TierToTierInfo is a free data retrieval call binding the contract method 0xb77da533.
//
// Solidity: function tierToTierInfo(uint8 tier) view returns((uint256,uint256,uint256))
func (_Plus *PlusCallerSession) TierToTierInfo(tier uint8) (TierInfo, error) {
	return _Plus.Contract.TierToTierInfo(&_Plus.CallOpts, tier)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_Plus *PlusCaller) TokenIdGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdGenerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_Plus *PlusSession) TokenIdGenerator() (*big.Int, error) {
	return _Plus.Contract.TokenIdGenerator(&_Plus.CallOpts)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_Plus *PlusCallerSession) TokenIdGenerator() (*big.Int, error) {
	return _Plus.Contract.TokenIdGenerator(&_Plus.CallOpts)
}

// TokenIdToAgentId is a free data retrieval call binding the contract method 0x4449ee09.
//
// Solidity: function tokenIdToAgentId(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCaller) TokenIdToAgentId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdToAgentId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToAgentId is a free data retrieval call binding the contract method 0x4449ee09.
//
// Solidity: function tokenIdToAgentId(uint256 tokenId) view returns(uint256)
func (_Plus *PlusSession) TokenIdToAgentId(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToAgentId(&_Plus.CallOpts, tokenId)
}

// TokenIdToAgentId is a free data retrieval call binding the contract method 0x4449ee09.
//
// Solidity: function tokenIdToAgentId(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCallerSession) TokenIdToAgentId(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToAgentId(&_Plus.CallOpts, tokenId)
}

// TokenIdToFilCashbackEarned is a free data retrieval call binding the contract method 0x2fe5c820.
//
// Solidity: function tokenIdToFilCashbackEarned(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCaller) TokenIdToFilCashbackEarned(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdToFilCashbackEarned", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToFilCashbackEarned is a free data retrieval call binding the contract method 0x2fe5c820.
//
// Solidity: function tokenIdToFilCashbackEarned(uint256 tokenId) view returns(uint256)
func (_Plus *PlusSession) TokenIdToFilCashbackEarned(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToFilCashbackEarned(&_Plus.CallOpts, tokenId)
}

// TokenIdToFilCashbackEarned is a free data retrieval call binding the contract method 0x2fe5c820.
//
// Solidity: function tokenIdToFilCashbackEarned(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCallerSession) TokenIdToFilCashbackEarned(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToFilCashbackEarned(&_Plus.CallOpts, tokenId)
}

// TokenIdToGlfVaultBalance is a free data retrieval call binding the contract method 0x5e7277f3.
//
// Solidity: function tokenIdToGlfVaultBalance(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCaller) TokenIdToGlfVaultBalance(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdToGlfVaultBalance", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToGlfVaultBalance is a free data retrieval call binding the contract method 0x5e7277f3.
//
// Solidity: function tokenIdToGlfVaultBalance(uint256 tokenId) view returns(uint256)
func (_Plus *PlusSession) TokenIdToGlfVaultBalance(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToGlfVaultBalance(&_Plus.CallOpts, tokenId)
}

// TokenIdToGlfVaultBalance is a free data retrieval call binding the contract method 0x5e7277f3.
//
// Solidity: function tokenIdToGlfVaultBalance(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCallerSession) TokenIdToGlfVaultBalance(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToGlfVaultBalance(&_Plus.CallOpts, tokenId)
}

// TokenIdToLastTierSwitchTimestamp is a free data retrieval call binding the contract method 0xf173e1fe.
//
// Solidity: function tokenIdToLastTierSwitchTimestamp(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCaller) TokenIdToLastTierSwitchTimestamp(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdToLastTierSwitchTimestamp", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToLastTierSwitchTimestamp is a free data retrieval call binding the contract method 0xf173e1fe.
//
// Solidity: function tokenIdToLastTierSwitchTimestamp(uint256 tokenId) view returns(uint256)
func (_Plus *PlusSession) TokenIdToLastTierSwitchTimestamp(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToLastTierSwitchTimestamp(&_Plus.CallOpts, tokenId)
}

// TokenIdToLastTierSwitchTimestamp is a free data retrieval call binding the contract method 0xf173e1fe.
//
// Solidity: function tokenIdToLastTierSwitchTimestamp(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCallerSession) TokenIdToLastTierSwitchTimestamp(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToLastTierSwitchTimestamp(&_Plus.CallOpts, tokenId)
}

// TokenIdToPersonalCashBackPercent is a free data retrieval call binding the contract method 0x9780b372.
//
// Solidity: function tokenIdToPersonalCashBackPercent(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCaller) TokenIdToPersonalCashBackPercent(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdToPersonalCashBackPercent", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToPersonalCashBackPercent is a free data retrieval call binding the contract method 0x9780b372.
//
// Solidity: function tokenIdToPersonalCashBackPercent(uint256 tokenId) view returns(uint256)
func (_Plus *PlusSession) TokenIdToPersonalCashBackPercent(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToPersonalCashBackPercent(&_Plus.CallOpts, tokenId)
}

// TokenIdToPersonalCashBackPercent is a free data retrieval call binding the contract method 0x9780b372.
//
// Solidity: function tokenIdToPersonalCashBackPercent(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCallerSession) TokenIdToPersonalCashBackPercent(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToPersonalCashBackPercent(&_Plus.CallOpts, tokenId)
}

// TokenIdToTier is a free data retrieval call binding the contract method 0x0431a325.
//
// Solidity: function tokenIdToTier(uint256 tokenId) view returns(uint8)
func (_Plus *PlusCaller) TokenIdToTier(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdToTier", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenIdToTier is a free data retrieval call binding the contract method 0x0431a325.
//
// Solidity: function tokenIdToTier(uint256 tokenId) view returns(uint8)
func (_Plus *PlusSession) TokenIdToTier(tokenId *big.Int) (uint8, error) {
	return _Plus.Contract.TokenIdToTier(&_Plus.CallOpts, tokenId)
}

// TokenIdToTier is a free data retrieval call binding the contract method 0x0431a325.
//
// Solidity: function tokenIdToTier(uint256 tokenId) view returns(uint8)
func (_Plus *PlusCallerSession) TokenIdToTier(tokenId *big.Int) (uint8, error) {
	return _Plus.Contract.TokenIdToTier(&_Plus.CallOpts, tokenId)
}

// TokenIdToTierLockAmount is a free data retrieval call binding the contract method 0x02adafe1.
//
// Solidity: function tokenIdToTierLockAmount(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCaller) TokenIdToTierLockAmount(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenIdToTierLockAmount", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdToTierLockAmount is a free data retrieval call binding the contract method 0x02adafe1.
//
// Solidity: function tokenIdToTierLockAmount(uint256 tokenId) view returns(uint256)
func (_Plus *PlusSession) TokenIdToTierLockAmount(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToTierLockAmount(&_Plus.CallOpts, tokenId)
}

// TokenIdToTierLockAmount is a free data retrieval call binding the contract method 0x02adafe1.
//
// Solidity: function tokenIdToTierLockAmount(uint256 tokenId) view returns(uint256)
func (_Plus *PlusCallerSession) TokenIdToTierLockAmount(tokenId *big.Int) (*big.Int, error) {
	return _Plus.Contract.TokenIdToTierLockAmount(&_Plus.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Plus *PlusCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Plus *PlusSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Plus.Contract.TokenURI(&_Plus.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Plus *PlusCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Plus.Contract.TokenURI(&_Plus.CallOpts, tokenId)
}

// TotalFilCashbackVaultBalance is a free data retrieval call binding the contract method 0x6b12bee3.
//
// Solidity: function totalFilCashbackVaultBalance() view returns(uint256)
func (_Plus *PlusCaller) TotalFilCashbackVaultBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "totalFilCashbackVaultBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFilCashbackVaultBalance is a free data retrieval call binding the contract method 0x6b12bee3.
//
// Solidity: function totalFilCashbackVaultBalance() view returns(uint256)
func (_Plus *PlusSession) TotalFilCashbackVaultBalance() (*big.Int, error) {
	return _Plus.Contract.TotalFilCashbackVaultBalance(&_Plus.CallOpts)
}

// TotalFilCashbackVaultBalance is a free data retrieval call binding the contract method 0x6b12bee3.
//
// Solidity: function totalFilCashbackVaultBalance() view returns(uint256)
func (_Plus *PlusCallerSession) TotalFilCashbackVaultBalance() (*big.Int, error) {
	return _Plus.Contract.TotalFilCashbackVaultBalance(&_Plus.CallOpts)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_Plus *PlusCaller) TransferFrom(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "transferFrom", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_Plus *PlusSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _Plus.Contract.TransferFrom(&_Plus.CallOpts, arg0, arg1, arg2)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns()
func (_Plus *PlusCallerSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) error {
	return _Plus.Contract.TransferFrom(&_Plus.CallOpts, arg0, arg1, arg2)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Plus *PlusCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Plus *PlusSession) Treasury() (common.Address, error) {
	return _Plus.Contract.Treasury(&_Plus.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Plus *PlusCallerSession) Treasury() (common.Address, error) {
	return _Plus.Contract.Treasury(&_Plus.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Plus *PlusTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Plus *PlusSession) AcceptOwnership() (*types.Transaction, error) {
	return _Plus.Contract.AcceptOwnership(&_Plus.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Plus *PlusTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Plus.Contract.AcceptOwnership(&_Plus.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xaef0353c.
//
// Solidity: function activate(address _beneficiary, uint256 _tokenId, uint8 _tier) returns()
func (_Plus *PlusTransactor) Activate(opts *bind.TransactOpts, _beneficiary common.Address, _tokenId *big.Int, _tier uint8) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "activate", _beneficiary, _tokenId, _tier)
}

// Activate is a paid mutator transaction binding the contract method 0xaef0353c.
//
// Solidity: function activate(address _beneficiary, uint256 _tokenId, uint8 _tier) returns()
func (_Plus *PlusSession) Activate(_beneficiary common.Address, _tokenId *big.Int, _tier uint8) (*types.Transaction, error) {
	return _Plus.Contract.Activate(&_Plus.TransactOpts, _beneficiary, _tokenId, _tier)
}

// Activate is a paid mutator transaction binding the contract method 0xaef0353c.
//
// Solidity: function activate(address _beneficiary, uint256 _tokenId, uint8 _tier) returns()
func (_Plus *PlusTransactorSession) Activate(_beneficiary common.Address, _tokenId *big.Int, _tier uint8) (*types.Transaction, error) {
	return _Plus.Contract.Activate(&_Plus.TransactOpts, _beneficiary, _tokenId, _tier)
}

// ChangeOwnerForAgent is a paid mutator transaction binding the contract method 0x28d0e5a9.
//
// Solidity: function changeOwnerForAgent(address _agent) returns()
func (_Plus *PlusTransactor) ChangeOwnerForAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "changeOwnerForAgent", _agent)
}

// ChangeOwnerForAgent is a paid mutator transaction binding the contract method 0x28d0e5a9.
//
// Solidity: function changeOwnerForAgent(address _agent) returns()
func (_Plus *PlusSession) ChangeOwnerForAgent(_agent common.Address) (*types.Transaction, error) {
	return _Plus.Contract.ChangeOwnerForAgent(&_Plus.TransactOpts, _agent)
}

// ChangeOwnerForAgent is a paid mutator transaction binding the contract method 0x28d0e5a9.
//
// Solidity: function changeOwnerForAgent(address _agent) returns()
func (_Plus *PlusTransactorSession) ChangeOwnerForAgent(_agent common.Address) (*types.Transaction, error) {
	return _Plus.Contract.ChangeOwnerForAgent(&_Plus.TransactOpts, _agent)
}

// ClaimCashBack is a paid mutator transaction binding the contract method 0x9202adfc.
//
// Solidity: function claimCashBack(uint256 _tokenId, address _receiver) returns()
func (_Plus *PlusTransactor) ClaimCashBack(opts *bind.TransactOpts, _tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "claimCashBack", _tokenId, _receiver)
}

// ClaimCashBack is a paid mutator transaction binding the contract method 0x9202adfc.
//
// Solidity: function claimCashBack(uint256 _tokenId, address _receiver) returns()
func (_Plus *PlusSession) ClaimCashBack(_tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Plus.Contract.ClaimCashBack(&_Plus.TransactOpts, _tokenId, _receiver)
}

// ClaimCashBack is a paid mutator transaction binding the contract method 0x9202adfc.
//
// Solidity: function claimCashBack(uint256 _tokenId, address _receiver) returns()
func (_Plus *PlusTransactorSession) ClaimCashBack(_tokenId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Plus.Contract.ClaimCashBack(&_Plus.TransactOpts, _tokenId, _receiver)
}

// Downgrade is a paid mutator transaction binding the contract method 0x1fac65d3.
//
// Solidity: function downgrade(uint256 _tokenId, uint8 _desiredTier, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) _sc) returns()
func (_Plus *PlusTransactor) Downgrade(opts *bind.TransactOpts, _tokenId *big.Int, _desiredTier uint8, _sc SignedCredential) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "downgrade", _tokenId, _desiredTier, _sc)
}

// Downgrade is a paid mutator transaction binding the contract method 0x1fac65d3.
//
// Solidity: function downgrade(uint256 _tokenId, uint8 _desiredTier, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) _sc) returns()
func (_Plus *PlusSession) Downgrade(_tokenId *big.Int, _desiredTier uint8, _sc SignedCredential) (*types.Transaction, error) {
	return _Plus.Contract.Downgrade(&_Plus.TransactOpts, _tokenId, _desiredTier, _sc)
}

// Downgrade is a paid mutator transaction binding the contract method 0x1fac65d3.
//
// Solidity: function downgrade(uint256 _tokenId, uint8 _desiredTier, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) _sc) returns()
func (_Plus *PlusTransactorSession) Downgrade(_tokenId *big.Int, _desiredTier uint8, _sc SignedCredential) (*types.Transaction, error) {
	return _Plus.Contract.Downgrade(&_Plus.TransactOpts, _tokenId, _desiredTier, _sc)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_Plus *PlusTransactor) FundFilVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "fundFilVault")
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_Plus *PlusSession) FundFilVault() (*types.Transaction, error) {
	return _Plus.Contract.FundFilVault(&_Plus.TransactOpts)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x91293473.
//
// Solidity: function fundFilVault() payable returns()
func (_Plus *PlusTransactorSession) FundFilVault() (*types.Transaction, error) {
	return _Plus.Contract.FundFilVault(&_Plus.TransactOpts)
}

// FundGlfVault is a paid mutator transaction binding the contract method 0x463ad343.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount) returns()
func (_Plus *PlusTransactor) FundGlfVault(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "fundGlfVault", _tokenId, _amount)
}

// FundGlfVault is a paid mutator transaction binding the contract method 0x463ad343.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount) returns()
func (_Plus *PlusSession) FundGlfVault(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.FundGlfVault(&_Plus.TransactOpts, _tokenId, _amount)
}

// FundGlfVault is a paid mutator transaction binding the contract method 0x463ad343.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount) returns()
func (_Plus *PlusTransactorSession) FundGlfVault(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.FundGlfVault(&_Plus.TransactOpts, _tokenId, _amount)
}

// FundGlfVault0 is a paid mutator transaction binding the contract method 0xebca88fc.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount, uint256 _cashBackPercent) returns()
func (_Plus *PlusTransactor) FundGlfVault0(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "fundGlfVault0", _tokenId, _amount, _cashBackPercent)
}

// FundGlfVault0 is a paid mutator transaction binding the contract method 0xebca88fc.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount, uint256 _cashBackPercent) returns()
func (_Plus *PlusSession) FundGlfVault0(_tokenId *big.Int, _amount *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.FundGlfVault0(&_Plus.TransactOpts, _tokenId, _amount, _cashBackPercent)
}

// FundGlfVault0 is a paid mutator transaction binding the contract method 0xebca88fc.
//
// Solidity: function fundGlfVault(uint256 _tokenId, uint256 _amount, uint256 _cashBackPercent) returns()
func (_Plus *PlusTransactorSession) FundGlfVault0(_tokenId *big.Int, _amount *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.FundGlfVault0(&_Plus.TransactOpts, _tokenId, _amount, _cashBackPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4521548.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier, address _poolAddress) returns()
func (_Plus *PlusTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "initialize", _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier, _poolAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4521548.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier, address _poolAddress) returns()
func (_Plus *PlusSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Initialize(&_Plus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier, _poolAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4521548.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier, address _poolAddress) returns()
func (_Plus *PlusTransactorSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address, _poolAddress common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Initialize(&_Plus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier, _poolAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 tokenId)
func (_Plus *PlusTransactor) Mint(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "mint", _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 tokenId)
func (_Plus *PlusSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Mint(&_Plus.TransactOpts, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 tokenId)
func (_Plus *PlusTransactorSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Mint(&_Plus.TransactOpts, _to)
}

// MintActivateAndFund is a paid mutator transaction binding the contract method 0xd7d2929e.
//
// Solidity: function mintActivateAndFund(address _to, uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier, uint256 _amount) returns(uint256 tokenId)
func (_Plus *PlusTransactor) MintActivateAndFund(opts *bind.TransactOpts, _to common.Address, _personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8, _amount *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "mintActivateAndFund", _to, _personalCashBackPercent, _beneficiary, _tier, _amount)
}

// MintActivateAndFund is a paid mutator transaction binding the contract method 0xd7d2929e.
//
// Solidity: function mintActivateAndFund(address _to, uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier, uint256 _amount) returns(uint256 tokenId)
func (_Plus *PlusSession) MintActivateAndFund(_to common.Address, _personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8, _amount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.MintActivateAndFund(&_Plus.TransactOpts, _to, _personalCashBackPercent, _beneficiary, _tier, _amount)
}

// MintActivateAndFund is a paid mutator transaction binding the contract method 0xd7d2929e.
//
// Solidity: function mintActivateAndFund(address _to, uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier, uint256 _amount) returns(uint256 tokenId)
func (_Plus *PlusTransactorSession) MintActivateAndFund(_to common.Address, _personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8, _amount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.MintActivateAndFund(&_Plus.TransactOpts, _to, _personalCashBackPercent, _beneficiary, _tier, _amount)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0xffc21089.
//
// Solidity: function mintAndActivate(address _to, address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_Plus *PlusTransactor) MintAndActivate(opts *bind.TransactOpts, _to common.Address, _beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "mintAndActivate", _to, _beneficiary, _tier)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0xffc21089.
//
// Solidity: function mintAndActivate(address _to, address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_Plus *PlusSession) MintAndActivate(_to common.Address, _beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _Plus.Contract.MintAndActivate(&_Plus.TransactOpts, _to, _beneficiary, _tier)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0xffc21089.
//
// Solidity: function mintAndActivate(address _to, address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_Plus *PlusTransactorSession) MintAndActivate(_to common.Address, _beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _Plus.Contract.MintAndActivate(&_Plus.TransactOpts, _to, _beneficiary, _tier)
}

// OnDefault is a paid mutator transaction binding the contract method 0x8b446431.
//
// Solidity: function onDefault(uint256 _agentId) returns()
func (_Plus *PlusTransactor) OnDefault(opts *bind.TransactOpts, _agentId *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "onDefault", _agentId)
}

// OnDefault is a paid mutator transaction binding the contract method 0x8b446431.
//
// Solidity: function onDefault(uint256 _agentId) returns()
func (_Plus *PlusSession) OnDefault(_agentId *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.OnDefault(&_Plus.TransactOpts, _agentId)
}

// OnDefault is a paid mutator transaction binding the contract method 0x8b446431.
//
// Solidity: function onDefault(uint256 _agentId) returns()
func (_Plus *PlusTransactorSession) OnDefault(_agentId *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.OnDefault(&_Plus.TransactOpts, _agentId)
}

// OnPaymentMade is a paid mutator transaction binding the contract method 0xcda43c10.
//
// Solidity: function onPaymentMade(uint256 _agentId, uint256 _interestAmount) returns()
func (_Plus *PlusTransactor) OnPaymentMade(opts *bind.TransactOpts, _agentId *big.Int, _interestAmount *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "onPaymentMade", _agentId, _interestAmount)
}

// OnPaymentMade is a paid mutator transaction binding the contract method 0xcda43c10.
//
// Solidity: function onPaymentMade(uint256 _agentId, uint256 _interestAmount) returns()
func (_Plus *PlusSession) OnPaymentMade(_agentId *big.Int, _interestAmount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.OnPaymentMade(&_Plus.TransactOpts, _agentId, _interestAmount)
}

// OnPaymentMade is a paid mutator transaction binding the contract method 0xcda43c10.
//
// Solidity: function onPaymentMade(uint256 _agentId, uint256 _interestAmount) returns()
func (_Plus *PlusTransactorSession) OnPaymentMade(_agentId *big.Int, _interestAmount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.OnPaymentMade(&_Plus.TransactOpts, _agentId, _interestAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Plus *PlusTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Plus *PlusSession) Pause() (*types.Transaction, error) {
	return _Plus.Contract.Pause(&_Plus.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Plus *PlusTransactorSession) Pause() (*types.Transaction, error) {
	return _Plus.Contract.Pause(&_Plus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Plus *PlusTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Plus *PlusSession) RenounceOwnership() (*types.Transaction, error) {
	return _Plus.Contract.RenounceOwnership(&_Plus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Plus *PlusTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Plus.Contract.RenounceOwnership(&_Plus.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Plus *PlusTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Plus *PlusSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SafeTransferFrom(&_Plus.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Plus *PlusTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SafeTransferFrom(&_Plus.TransactOpts, from, to, tokenId)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address _agentFactory) returns()
func (_Plus *PlusTransactor) SetAgentFactory(opts *bind.TransactOpts, _agentFactory common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setAgentFactory", _agentFactory)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address _agentFactory) returns()
func (_Plus *PlusSession) SetAgentFactory(_agentFactory common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetAgentFactory(&_Plus.TransactOpts, _agentFactory)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address _agentFactory) returns()
func (_Plus *PlusTransactorSession) SetAgentFactory(_agentFactory common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetAgentFactory(&_Plus.TransactOpts, _agentFactory)
}

// SetAgentPolice is a paid mutator transaction binding the contract method 0x86cc3437.
//
// Solidity: function setAgentPolice(address _agentPoliceVcVerifier) returns()
func (_Plus *PlusTransactor) SetAgentPolice(opts *bind.TransactOpts, _agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setAgentPolice", _agentPoliceVcVerifier)
}

// SetAgentPolice is a paid mutator transaction binding the contract method 0x86cc3437.
//
// Solidity: function setAgentPolice(address _agentPoliceVcVerifier) returns()
func (_Plus *PlusSession) SetAgentPolice(_agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetAgentPolice(&_Plus.TransactOpts, _agentPoliceVcVerifier)
}

// SetAgentPolice is a paid mutator transaction binding the contract method 0x86cc3437.
//
// Solidity: function setAgentPolice(address _agentPoliceVcVerifier) returns()
func (_Plus *PlusTransactorSession) SetAgentPolice(_agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetAgentPolice(&_Plus.TransactOpts, _agentPoliceVcVerifier)
}

// SetBaseConversionRateFILtoGLF is a paid mutator transaction binding the contract method 0xd86e1a35.
//
// Solidity: function setBaseConversionRateFILtoGLF(uint256 _baseConversionRateFILtoGLF) returns()
func (_Plus *PlusTransactor) SetBaseConversionRateFILtoGLF(opts *bind.TransactOpts, _baseConversionRateFILtoGLF *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setBaseConversionRateFILtoGLF", _baseConversionRateFILtoGLF)
}

// SetBaseConversionRateFILtoGLF is a paid mutator transaction binding the contract method 0xd86e1a35.
//
// Solidity: function setBaseConversionRateFILtoGLF(uint256 _baseConversionRateFILtoGLF) returns()
func (_Plus *PlusSession) SetBaseConversionRateFILtoGLF(_baseConversionRateFILtoGLF *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetBaseConversionRateFILtoGLF(&_Plus.TransactOpts, _baseConversionRateFILtoGLF)
}

// SetBaseConversionRateFILtoGLF is a paid mutator transaction binding the contract method 0xd86e1a35.
//
// Solidity: function setBaseConversionRateFILtoGLF(uint256 _baseConversionRateFILtoGLF) returns()
func (_Plus *PlusTransactorSession) SetBaseConversionRateFILtoGLF(_baseConversionRateFILtoGLF *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetBaseConversionRateFILtoGLF(&_Plus.TransactOpts, _baseConversionRateFILtoGLF)
}

// SetGlfToken is a paid mutator transaction binding the contract method 0x762b977d.
//
// Solidity: function setGlfToken(address _glfToken) returns()
func (_Plus *PlusTransactor) SetGlfToken(opts *bind.TransactOpts, _glfToken common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setGlfToken", _glfToken)
}

// SetGlfToken is a paid mutator transaction binding the contract method 0x762b977d.
//
// Solidity: function setGlfToken(address _glfToken) returns()
func (_Plus *PlusSession) SetGlfToken(_glfToken common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetGlfToken(&_Plus.TransactOpts, _glfToken)
}

// SetGlfToken is a paid mutator transaction binding the contract method 0x762b977d.
//
// Solidity: function setGlfToken(address _glfToken) returns()
func (_Plus *PlusTransactorSession) SetGlfToken(_glfToken common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetGlfToken(&_Plus.TransactOpts, _glfToken)
}

// SetMaxCashBackPercent is a paid mutator transaction binding the contract method 0x81eebd70.
//
// Solidity: function setMaxCashBackPercent(uint256 _maxCashBackPercent) returns()
func (_Plus *PlusTransactor) SetMaxCashBackPercent(opts *bind.TransactOpts, _maxCashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setMaxCashBackPercent", _maxCashBackPercent)
}

// SetMaxCashBackPercent is a paid mutator transaction binding the contract method 0x81eebd70.
//
// Solidity: function setMaxCashBackPercent(uint256 _maxCashBackPercent) returns()
func (_Plus *PlusSession) SetMaxCashBackPercent(_maxCashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetMaxCashBackPercent(&_Plus.TransactOpts, _maxCashBackPercent)
}

// SetMaxCashBackPercent is a paid mutator transaction binding the contract method 0x81eebd70.
//
// Solidity: function setMaxCashBackPercent(uint256 _maxCashBackPercent) returns()
func (_Plus *PlusTransactorSession) SetMaxCashBackPercent(_maxCashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetMaxCashBackPercent(&_Plus.TransactOpts, _maxCashBackPercent)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_Plus *PlusTransactor) SetMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setMintPrice", _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_Plus *PlusSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetMintPrice(&_Plus.TransactOpts, _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_Plus *PlusTransactorSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetMintPrice(&_Plus.TransactOpts, _mintPrice)
}

// SetPersonalCashBackPercent is a paid mutator transaction binding the contract method 0x51c9749c.
//
// Solidity: function setPersonalCashBackPercent(uint256 _tokenId, uint256 _cashBackPercent) returns()
func (_Plus *PlusTransactor) SetPersonalCashBackPercent(opts *bind.TransactOpts, _tokenId *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setPersonalCashBackPercent", _tokenId, _cashBackPercent)
}

// SetPersonalCashBackPercent is a paid mutator transaction binding the contract method 0x51c9749c.
//
// Solidity: function setPersonalCashBackPercent(uint256 _tokenId, uint256 _cashBackPercent) returns()
func (_Plus *PlusSession) SetPersonalCashBackPercent(_tokenId *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetPersonalCashBackPercent(&_Plus.TransactOpts, _tokenId, _cashBackPercent)
}

// SetPersonalCashBackPercent is a paid mutator transaction binding the contract method 0x51c9749c.
//
// Solidity: function setPersonalCashBackPercent(uint256 _tokenId, uint256 _cashBackPercent) returns()
func (_Plus *PlusTransactorSession) SetPersonalCashBackPercent(_tokenId *big.Int, _cashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetPersonalCashBackPercent(&_Plus.TransactOpts, _tokenId, _cashBackPercent)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xe9e15b4f.
//
// Solidity: function setPoolAddress(address _poolAddress) returns()
func (_Plus *PlusTransactor) SetPoolAddress(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setPoolAddress", _poolAddress)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xe9e15b4f.
//
// Solidity: function setPoolAddress(address _poolAddress) returns()
func (_Plus *PlusSession) SetPoolAddress(_poolAddress common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetPoolAddress(&_Plus.TransactOpts, _poolAddress)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xe9e15b4f.
//
// Solidity: function setPoolAddress(address _poolAddress) returns()
func (_Plus *PlusTransactorSession) SetPoolAddress(_poolAddress common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetPoolAddress(&_Plus.TransactOpts, _poolAddress)
}

// SetPriceSetter is a paid mutator transaction binding the contract method 0x63791e3c.
//
// Solidity: function setPriceSetter(address _priceSetter) returns()
func (_Plus *PlusTransactor) SetPriceSetter(opts *bind.TransactOpts, _priceSetter common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setPriceSetter", _priceSetter)
}

// SetPriceSetter is a paid mutator transaction binding the contract method 0x63791e3c.
//
// Solidity: function setPriceSetter(address _priceSetter) returns()
func (_Plus *PlusSession) SetPriceSetter(_priceSetter common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetPriceSetter(&_Plus.TransactOpts, _priceSetter)
}

// SetPriceSetter is a paid mutator transaction binding the contract method 0x63791e3c.
//
// Solidity: function setPriceSetter(address _priceSetter) returns()
func (_Plus *PlusTransactorSession) SetPriceSetter(_priceSetter common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetPriceSetter(&_Plus.TransactOpts, _priceSetter)
}

// SetTierInfo is a paid mutator transaction binding the contract method 0xf9d20f53.
//
// Solidity: function setTierInfo(uint8 _tier, (uint256,uint256,uint256) _tierInfo) returns()
func (_Plus *PlusTransactor) SetTierInfo(opts *bind.TransactOpts, _tier uint8, _tierInfo TierInfo) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setTierInfo", _tier, _tierInfo)
}

// SetTierInfo is a paid mutator transaction binding the contract method 0xf9d20f53.
//
// Solidity: function setTierInfo(uint8 _tier, (uint256,uint256,uint256) _tierInfo) returns()
func (_Plus *PlusSession) SetTierInfo(_tier uint8, _tierInfo TierInfo) (*types.Transaction, error) {
	return _Plus.Contract.SetTierInfo(&_Plus.TransactOpts, _tier, _tierInfo)
}

// SetTierInfo is a paid mutator transaction binding the contract method 0xf9d20f53.
//
// Solidity: function setTierInfo(uint8 _tier, (uint256,uint256,uint256) _tierInfo) returns()
func (_Plus *PlusTransactorSession) SetTierInfo(_tier uint8, _tierInfo TierInfo) (*types.Transaction, error) {
	return _Plus.Contract.SetTierInfo(&_Plus.TransactOpts, _tier, _tierInfo)
}

// SetTierSwitchPenaltyFee is a paid mutator transaction binding the contract method 0xabdf23df.
//
// Solidity: function setTierSwitchPenaltyFee(uint256 _tierSwitchPenaltyFee) returns()
func (_Plus *PlusTransactor) SetTierSwitchPenaltyFee(opts *bind.TransactOpts, _tierSwitchPenaltyFee *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setTierSwitchPenaltyFee", _tierSwitchPenaltyFee)
}

// SetTierSwitchPenaltyFee is a paid mutator transaction binding the contract method 0xabdf23df.
//
// Solidity: function setTierSwitchPenaltyFee(uint256 _tierSwitchPenaltyFee) returns()
func (_Plus *PlusSession) SetTierSwitchPenaltyFee(_tierSwitchPenaltyFee *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetTierSwitchPenaltyFee(&_Plus.TransactOpts, _tierSwitchPenaltyFee)
}

// SetTierSwitchPenaltyFee is a paid mutator transaction binding the contract method 0xabdf23df.
//
// Solidity: function setTierSwitchPenaltyFee(uint256 _tierSwitchPenaltyFee) returns()
func (_Plus *PlusTransactorSession) SetTierSwitchPenaltyFee(_tierSwitchPenaltyFee *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetTierSwitchPenaltyFee(&_Plus.TransactOpts, _tierSwitchPenaltyFee)
}

// SetTierSwitchPenaltyWindow is a paid mutator transaction binding the contract method 0x0c1cd643.
//
// Solidity: function setTierSwitchPenaltyWindow(uint256 _tierSwitchPenaltyWindow) returns()
func (_Plus *PlusTransactor) SetTierSwitchPenaltyWindow(opts *bind.TransactOpts, _tierSwitchPenaltyWindow *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setTierSwitchPenaltyWindow", _tierSwitchPenaltyWindow)
}

// SetTierSwitchPenaltyWindow is a paid mutator transaction binding the contract method 0x0c1cd643.
//
// Solidity: function setTierSwitchPenaltyWindow(uint256 _tierSwitchPenaltyWindow) returns()
func (_Plus *PlusSession) SetTierSwitchPenaltyWindow(_tierSwitchPenaltyWindow *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetTierSwitchPenaltyWindow(&_Plus.TransactOpts, _tierSwitchPenaltyWindow)
}

// SetTierSwitchPenaltyWindow is a paid mutator transaction binding the contract method 0x0c1cd643.
//
// Solidity: function setTierSwitchPenaltyWindow(uint256 _tierSwitchPenaltyWindow) returns()
func (_Plus *PlusTransactorSession) SetTierSwitchPenaltyWindow(_tierSwitchPenaltyWindow *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetTierSwitchPenaltyWindow(&_Plus.TransactOpts, _tierSwitchPenaltyWindow)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Plus *PlusTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Plus *PlusSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetTreasury(&_Plus.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Plus *PlusTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetTreasury(&_Plus.TransactOpts, _treasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Plus *PlusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Plus *PlusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Plus.Contract.TransferOwnership(&_Plus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Plus *PlusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Plus.Contract.TransferOwnership(&_Plus.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Plus *PlusTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Plus *PlusSession) Unpause() (*types.Transaction, error) {
	return _Plus.Contract.Unpause(&_Plus.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Plus *PlusTransactorSession) Unpause() (*types.Transaction, error) {
	return _Plus.Contract.Unpause(&_Plus.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x7dfe5b92.
//
// Solidity: function upgrade(uint256 _tokenId, uint8 _desiredTier) returns()
func (_Plus *PlusTransactor) Upgrade(opts *bind.TransactOpts, _tokenId *big.Int, _desiredTier uint8) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "upgrade", _tokenId, _desiredTier)
}

// Upgrade is a paid mutator transaction binding the contract method 0x7dfe5b92.
//
// Solidity: function upgrade(uint256 _tokenId, uint8 _desiredTier) returns()
func (_Plus *PlusSession) Upgrade(_tokenId *big.Int, _desiredTier uint8) (*types.Transaction, error) {
	return _Plus.Contract.Upgrade(&_Plus.TransactOpts, _tokenId, _desiredTier)
}

// Upgrade is a paid mutator transaction binding the contract method 0x7dfe5b92.
//
// Solidity: function upgrade(uint256 _tokenId, uint8 _desiredTier) returns()
func (_Plus *PlusTransactorSession) Upgrade(_tokenId *big.Int, _desiredTier uint8) (*types.Transaction, error) {
	return _Plus.Contract.Upgrade(&_Plus.TransactOpts, _tokenId, _desiredTier)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Plus *PlusTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Plus *PlusSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Plus.Contract.UpgradeToAndCall(&_Plus.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Plus *PlusTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Plus.Contract.UpgradeToAndCall(&_Plus.TransactOpts, newImplementation, data)
}

// WithdrawAllFilFunds is a paid mutator transaction binding the contract method 0x5df35481.
//
// Solidity: function withdrawAllFilFunds(address _to) returns()
func (_Plus *PlusTransactor) WithdrawAllFilFunds(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "withdrawAllFilFunds", _to)
}

// WithdrawAllFilFunds is a paid mutator transaction binding the contract method 0x5df35481.
//
// Solidity: function withdrawAllFilFunds(address _to) returns()
func (_Plus *PlusSession) WithdrawAllFilFunds(_to common.Address) (*types.Transaction, error) {
	return _Plus.Contract.WithdrawAllFilFunds(&_Plus.TransactOpts, _to)
}

// WithdrawAllFilFunds is a paid mutator transaction binding the contract method 0x5df35481.
//
// Solidity: function withdrawAllFilFunds(address _to) returns()
func (_Plus *PlusTransactorSession) WithdrawAllFilFunds(_to common.Address) (*types.Transaction, error) {
	return _Plus.Contract.WithdrawAllFilFunds(&_Plus.TransactOpts, _to)
}

// WithdrawExtraLockedFunds is a paid mutator transaction binding the contract method 0x14b77f22.
//
// Solidity: function withdrawExtraLockedFunds(uint256 _tokenId) returns()
func (_Plus *PlusTransactor) WithdrawExtraLockedFunds(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "withdrawExtraLockedFunds", _tokenId)
}

// WithdrawExtraLockedFunds is a paid mutator transaction binding the contract method 0x14b77f22.
//
// Solidity: function withdrawExtraLockedFunds(uint256 _tokenId) returns()
func (_Plus *PlusSession) WithdrawExtraLockedFunds(_tokenId *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.WithdrawExtraLockedFunds(&_Plus.TransactOpts, _tokenId)
}

// WithdrawExtraLockedFunds is a paid mutator transaction binding the contract method 0x14b77f22.
//
// Solidity: function withdrawExtraLockedFunds(uint256 _tokenId) returns()
func (_Plus *PlusTransactorSession) WithdrawExtraLockedFunds(_tokenId *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.WithdrawExtraLockedFunds(&_Plus.TransactOpts, _tokenId)
}

// PlusApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Plus contract.
type PlusApprovalIterator struct {
	Event *PlusApproval // Event containing the contract specifics and raw log

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
func (it *PlusApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusApproval)
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
		it.Event = new(PlusApproval)
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
func (it *PlusApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusApproval represents a Approval event raised by the Plus contract.
type PlusApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Plus *PlusFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PlusApprovalIterator, error) {

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

	logs, sub, err := _Plus.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlusApprovalIterator{contract: _Plus.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Plus *PlusFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PlusApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Plus.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusApproval)
				if err := _Plus.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Plus *PlusFilterer) ParseApproval(log types.Log) (*PlusApproval, error) {
	event := new(PlusApproval)
	if err := _Plus.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Plus contract.
type PlusApprovalForAllIterator struct {
	Event *PlusApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PlusApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusApprovalForAll)
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
		it.Event = new(PlusApprovalForAll)
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
func (it *PlusApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusApprovalForAll represents a ApprovalForAll event raised by the Plus contract.
type PlusApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Plus *PlusFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PlusApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PlusApprovalForAllIterator{contract: _Plus.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Plus *PlusFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PlusApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusApprovalForAll)
				if err := _Plus.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Plus *PlusFilterer) ParseApprovalForAll(log types.Log) (*PlusApprovalForAll, error) {
	event := new(PlusApprovalForAll)
	if err := _Plus.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusCardActivatedIterator is returned from FilterCardActivated and is used to iterate over the raw logs and unpacked data for CardActivated events raised by the Plus contract.
type PlusCardActivatedIterator struct {
	Event *PlusCardActivated // Event containing the contract specifics and raw log

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
func (it *PlusCardActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusCardActivated)
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
		it.Event = new(PlusCardActivated)
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
func (it *PlusCardActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusCardActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusCardActivated represents a CardActivated event raised by the Plus contract.
type PlusCardActivated struct {
	TokenId     *big.Int
	Beneficiary common.Address
	Tier        uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCardActivated is a free log retrieval operation binding the contract event 0x62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5.
//
// Solidity: event CardActivated(uint256 indexed tokenId, address indexed beneficiary, uint8 tier)
func (_Plus *PlusFilterer) FilterCardActivated(opts *bind.FilterOpts, tokenId []*big.Int, beneficiary []common.Address) (*PlusCardActivatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "CardActivated", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &PlusCardActivatedIterator{contract: _Plus.contract, event: "CardActivated", logs: logs, sub: sub}, nil
}

// WatchCardActivated is a free log subscription operation binding the contract event 0x62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5.
//
// Solidity: event CardActivated(uint256 indexed tokenId, address indexed beneficiary, uint8 tier)
func (_Plus *PlusFilterer) WatchCardActivated(opts *bind.WatchOpts, sink chan<- *PlusCardActivated, tokenId []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "CardActivated", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusCardActivated)
				if err := _Plus.contract.UnpackLog(event, "CardActivated", log); err != nil {
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
func (_Plus *PlusFilterer) ParseCardActivated(log types.Log) (*PlusCardActivated, error) {
	event := new(PlusCardActivated)
	if err := _Plus.contract.UnpackLog(event, "CardActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusCardDowngradedIterator is returned from FilterCardDowngraded and is used to iterate over the raw logs and unpacked data for CardDowngraded events raised by the Plus contract.
type PlusCardDowngradedIterator struct {
	Event *PlusCardDowngraded // Event containing the contract specifics and raw log

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
func (it *PlusCardDowngradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusCardDowngraded)
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
		it.Event = new(PlusCardDowngraded)
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
func (it *PlusCardDowngradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusCardDowngradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusCardDowngraded represents a CardDowngraded event raised by the Plus contract.
type PlusCardDowngraded struct {
	TokenId     *big.Int
	Beneficiary common.Address
	OldTier     uint8
	NewTier     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCardDowngraded is a free log retrieval operation binding the contract event 0x65e7b85a0b1fe94cdd38e2a9dbe4ac14dda1ceeb998856160d0ea99b804b105d.
//
// Solidity: event CardDowngraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_Plus *PlusFilterer) FilterCardDowngraded(opts *bind.FilterOpts, tokenId []*big.Int, beneficiary []common.Address) (*PlusCardDowngradedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "CardDowngraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &PlusCardDowngradedIterator{contract: _Plus.contract, event: "CardDowngraded", logs: logs, sub: sub}, nil
}

// WatchCardDowngraded is a free log subscription operation binding the contract event 0x65e7b85a0b1fe94cdd38e2a9dbe4ac14dda1ceeb998856160d0ea99b804b105d.
//
// Solidity: event CardDowngraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_Plus *PlusFilterer) WatchCardDowngraded(opts *bind.WatchOpts, sink chan<- *PlusCardDowngraded, tokenId []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "CardDowngraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusCardDowngraded)
				if err := _Plus.contract.UnpackLog(event, "CardDowngraded", log); err != nil {
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

// ParseCardDowngraded is a log parse operation binding the contract event 0x65e7b85a0b1fe94cdd38e2a9dbe4ac14dda1ceeb998856160d0ea99b804b105d.
//
// Solidity: event CardDowngraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_Plus *PlusFilterer) ParseCardDowngraded(log types.Log) (*PlusCardDowngraded, error) {
	event := new(PlusCardDowngraded)
	if err := _Plus.contract.UnpackLog(event, "CardDowngraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusCardMintedIterator is returned from FilterCardMinted and is used to iterate over the raw logs and unpacked data for CardMinted events raised by the Plus contract.
type PlusCardMintedIterator struct {
	Event *PlusCardMinted // Event containing the contract specifics and raw log

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
func (it *PlusCardMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusCardMinted)
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
		it.Event = new(PlusCardMinted)
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
func (it *PlusCardMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusCardMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusCardMinted represents a CardMinted event raised by the Plus contract.
type PlusCardMinted struct {
	TokenId  *big.Int
	Minter   common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCardMinted is a free log retrieval operation binding the contract event 0x864b3215ee93b2a9291e17b45d665127a426c1259666cf483c30939c4a710332.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address minter, address receiver)
func (_Plus *PlusFilterer) FilterCardMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*PlusCardMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "CardMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlusCardMintedIterator{contract: _Plus.contract, event: "CardMinted", logs: logs, sub: sub}, nil
}

// WatchCardMinted is a free log subscription operation binding the contract event 0x864b3215ee93b2a9291e17b45d665127a426c1259666cf483c30939c4a710332.
//
// Solidity: event CardMinted(uint256 indexed tokenId, address minter, address receiver)
func (_Plus *PlusFilterer) WatchCardMinted(opts *bind.WatchOpts, sink chan<- *PlusCardMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "CardMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusCardMinted)
				if err := _Plus.contract.UnpackLog(event, "CardMinted", log); err != nil {
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
func (_Plus *PlusFilterer) ParseCardMinted(log types.Log) (*PlusCardMinted, error) {
	event := new(PlusCardMinted)
	if err := _Plus.contract.UnpackLog(event, "CardMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusCardUpgradedIterator is returned from FilterCardUpgraded and is used to iterate over the raw logs and unpacked data for CardUpgraded events raised by the Plus contract.
type PlusCardUpgradedIterator struct {
	Event *PlusCardUpgraded // Event containing the contract specifics and raw log

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
func (it *PlusCardUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusCardUpgraded)
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
		it.Event = new(PlusCardUpgraded)
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
func (it *PlusCardUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusCardUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusCardUpgraded represents a CardUpgraded event raised by the Plus contract.
type PlusCardUpgraded struct {
	TokenId     *big.Int
	Beneficiary common.Address
	OldTier     uint8
	NewTier     uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCardUpgraded is a free log retrieval operation binding the contract event 0x7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd.
//
// Solidity: event CardUpgraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_Plus *PlusFilterer) FilterCardUpgraded(opts *bind.FilterOpts, tokenId []*big.Int, beneficiary []common.Address) (*PlusCardUpgradedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "CardUpgraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &PlusCardUpgradedIterator{contract: _Plus.contract, event: "CardUpgraded", logs: logs, sub: sub}, nil
}

// WatchCardUpgraded is a free log subscription operation binding the contract event 0x7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd.
//
// Solidity: event CardUpgraded(uint256 indexed tokenId, address indexed beneficiary, uint8 oldTier, uint8 newTier)
func (_Plus *PlusFilterer) WatchCardUpgraded(opts *bind.WatchOpts, sink chan<- *PlusCardUpgraded, tokenId []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "CardUpgraded", tokenIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusCardUpgraded)
				if err := _Plus.contract.UnpackLog(event, "CardUpgraded", log); err != nil {
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
func (_Plus *PlusFilterer) ParseCardUpgraded(log types.Log) (*PlusCardUpgraded, error) {
	event := new(PlusCardUpgraded)
	if err := _Plus.contract.UnpackLog(event, "CardUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusCashBackClaimedIterator is returned from FilterCashBackClaimed and is used to iterate over the raw logs and unpacked data for CashBackClaimed events raised by the Plus contract.
type PlusCashBackClaimedIterator struct {
	Event *PlusCashBackClaimed // Event containing the contract specifics and raw log

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
func (it *PlusCashBackClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusCashBackClaimed)
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
		it.Event = new(PlusCashBackClaimed)
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
func (it *PlusCashBackClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusCashBackClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusCashBackClaimed represents a CashBackClaimed event raised by the Plus contract.
type PlusCashBackClaimed struct {
	TokenId  *big.Int
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCashBackClaimed is a free log retrieval operation binding the contract event 0x05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700.
//
// Solidity: event CashBackClaimed(uint256 indexed tokenId, address indexed receiver, uint256 amount)
func (_Plus *PlusFilterer) FilterCashBackClaimed(opts *bind.FilterOpts, tokenId []*big.Int, receiver []common.Address) (*PlusCashBackClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "CashBackClaimed", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PlusCashBackClaimedIterator{contract: _Plus.contract, event: "CashBackClaimed", logs: logs, sub: sub}, nil
}

// WatchCashBackClaimed is a free log subscription operation binding the contract event 0x05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700.
//
// Solidity: event CashBackClaimed(uint256 indexed tokenId, address indexed receiver, uint256 amount)
func (_Plus *PlusFilterer) WatchCashBackClaimed(opts *bind.WatchOpts, sink chan<- *PlusCashBackClaimed, tokenId []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "CashBackClaimed", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusCashBackClaimed)
				if err := _Plus.contract.UnpackLog(event, "CashBackClaimed", log); err != nil {
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
func (_Plus *PlusFilterer) ParseCashBackClaimed(log types.Log) (*PlusCashBackClaimed, error) {
	event := new(PlusCashBackClaimed)
	if err := _Plus.contract.UnpackLog(event, "CashBackClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusFilVaultFundedIterator is returned from FilterFilVaultFunded and is used to iterate over the raw logs and unpacked data for FilVaultFunded events raised by the Plus contract.
type PlusFilVaultFundedIterator struct {
	Event *PlusFilVaultFunded // Event containing the contract specifics and raw log

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
func (it *PlusFilVaultFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusFilVaultFunded)
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
		it.Event = new(PlusFilVaultFunded)
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
func (it *PlusFilVaultFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusFilVaultFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusFilVaultFunded represents a FilVaultFunded event raised by the Plus contract.
type PlusFilVaultFunded struct {
	TokenId *big.Int
	Owner   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFilVaultFunded is a free log retrieval operation binding the contract event 0x1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d6074.
//
// Solidity: event FilVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
func (_Plus *PlusFilterer) FilterFilVaultFunded(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*PlusFilVaultFundedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "FilVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PlusFilVaultFundedIterator{contract: _Plus.contract, event: "FilVaultFunded", logs: logs, sub: sub}, nil
}

// WatchFilVaultFunded is a free log subscription operation binding the contract event 0x1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d6074.
//
// Solidity: event FilVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
func (_Plus *PlusFilterer) WatchFilVaultFunded(opts *bind.WatchOpts, sink chan<- *PlusFilVaultFunded, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "FilVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusFilVaultFunded)
				if err := _Plus.contract.UnpackLog(event, "FilVaultFunded", log); err != nil {
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
func (_Plus *PlusFilterer) ParseFilVaultFunded(log types.Log) (*PlusFilVaultFunded, error) {
	event := new(PlusFilVaultFunded)
	if err := _Plus.contract.UnpackLog(event, "FilVaultFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusGlfVaultFundedIterator is returned from FilterGlfVaultFunded and is used to iterate over the raw logs and unpacked data for GlfVaultFunded events raised by the Plus contract.
type PlusGlfVaultFundedIterator struct {
	Event *PlusGlfVaultFunded // Event containing the contract specifics and raw log

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
func (it *PlusGlfVaultFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusGlfVaultFunded)
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
		it.Event = new(PlusGlfVaultFunded)
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
func (it *PlusGlfVaultFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusGlfVaultFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusGlfVaultFunded represents a GlfVaultFunded event raised by the Plus contract.
type PlusGlfVaultFunded struct {
	TokenId         *big.Int
	Owner           common.Address
	Amount          *big.Int
	CashBackPercent *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGlfVaultFunded is a free log retrieval operation binding the contract event 0x57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c6.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 cashBackPercent)
func (_Plus *PlusFilterer) FilterGlfVaultFunded(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*PlusGlfVaultFundedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "GlfVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PlusGlfVaultFundedIterator{contract: _Plus.contract, event: "GlfVaultFunded", logs: logs, sub: sub}, nil
}

// WatchGlfVaultFunded is a free log subscription operation binding the contract event 0x57ad62153adc9e3be6862aa6aa4671b89c163b973fd87c66436d3d89641596c6.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount, uint256 cashBackPercent)
func (_Plus *PlusFilterer) WatchGlfVaultFunded(opts *bind.WatchOpts, sink chan<- *PlusGlfVaultFunded, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "GlfVaultFunded", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusGlfVaultFunded)
				if err := _Plus.contract.UnpackLog(event, "GlfVaultFunded", log); err != nil {
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
func (_Plus *PlusFilterer) ParseGlfVaultFunded(log types.Log) (*PlusGlfVaultFunded, error) {
	event := new(PlusGlfVaultFunded)
	if err := _Plus.contract.UnpackLog(event, "GlfVaultFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Plus contract.
type PlusInitializedIterator struct {
	Event *PlusInitialized // Event containing the contract specifics and raw log

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
func (it *PlusInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusInitialized)
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
		it.Event = new(PlusInitialized)
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
func (it *PlusInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusInitialized represents a Initialized event raised by the Plus contract.
type PlusInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Plus *PlusFilterer) FilterInitialized(opts *bind.FilterOpts) (*PlusInitializedIterator, error) {

	logs, sub, err := _Plus.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PlusInitializedIterator{contract: _Plus.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Plus *PlusFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PlusInitialized) (event.Subscription, error) {

	logs, sub, err := _Plus.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusInitialized)
				if err := _Plus.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Plus *PlusFilterer) ParseInitialized(log types.Log) (*PlusInitialized, error) {
	event := new(PlusInitialized)
	if err := _Plus.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Plus contract.
type PlusOwnershipTransferStartedIterator struct {
	Event *PlusOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *PlusOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusOwnershipTransferStarted)
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
		it.Event = new(PlusOwnershipTransferStarted)
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
func (it *PlusOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Plus contract.
type PlusOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Plus *PlusFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlusOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlusOwnershipTransferStartedIterator{contract: _Plus.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Plus *PlusFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *PlusOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusOwnershipTransferStarted)
				if err := _Plus.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Plus *PlusFilterer) ParseOwnershipTransferStarted(log types.Log) (*PlusOwnershipTransferStarted, error) {
	event := new(PlusOwnershipTransferStarted)
	if err := _Plus.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Plus contract.
type PlusOwnershipTransferredIterator struct {
	Event *PlusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PlusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusOwnershipTransferred)
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
		it.Event = new(PlusOwnershipTransferred)
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
func (it *PlusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusOwnershipTransferred represents a OwnershipTransferred event raised by the Plus contract.
type PlusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Plus *PlusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlusOwnershipTransferredIterator{contract: _Plus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Plus *PlusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusOwnershipTransferred)
				if err := _Plus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Plus *PlusFilterer) ParseOwnershipTransferred(log types.Log) (*PlusOwnershipTransferred, error) {
	event := new(PlusOwnershipTransferred)
	if err := _Plus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Plus contract.
type PlusPausedIterator struct {
	Event *PlusPaused // Event containing the contract specifics and raw log

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
func (it *PlusPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusPaused)
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
		it.Event = new(PlusPaused)
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
func (it *PlusPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusPaused represents a Paused event raised by the Plus contract.
type PlusPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Plus *PlusFilterer) FilterPaused(opts *bind.FilterOpts) (*PlusPausedIterator, error) {

	logs, sub, err := _Plus.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PlusPausedIterator{contract: _Plus.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Plus *PlusFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PlusPaused) (event.Subscription, error) {

	logs, sub, err := _Plus.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusPaused)
				if err := _Plus.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Plus *PlusFilterer) ParsePaused(log types.Log) (*PlusPaused, error) {
	event := new(PlusPaused)
	if err := _Plus.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusPaymentProcessedIterator is returned from FilterPaymentProcessed and is used to iterate over the raw logs and unpacked data for PaymentProcessed events raised by the Plus contract.
type PlusPaymentProcessedIterator struct {
	Event *PlusPaymentProcessed // Event containing the contract specifics and raw log

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
func (it *PlusPaymentProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusPaymentProcessed)
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
		it.Event = new(PlusPaymentProcessed)
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
func (it *PlusPaymentProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusPaymentProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusPaymentProcessed represents a PaymentProcessed event raised by the Plus contract.
type PlusPaymentProcessed struct {
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
func (_Plus *PlusFilterer) FilterPaymentProcessed(opts *bind.FilterOpts, agentId []*big.Int, tokenId []*big.Int) (*PlusPaymentProcessedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "PaymentProcessed", agentIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlusPaymentProcessedIterator{contract: _Plus.contract, event: "PaymentProcessed", logs: logs, sub: sub}, nil
}

// WatchPaymentProcessed is a free log subscription operation binding the contract event 0x299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d6.
//
// Solidity: event PaymentProcessed(uint256 indexed agentId, uint256 indexed tokenId, uint256 interestAmount, uint256 glfBurned, uint256 filEarned)
func (_Plus *PlusFilterer) WatchPaymentProcessed(opts *bind.WatchOpts, sink chan<- *PlusPaymentProcessed, agentId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "PaymentProcessed", agentIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusPaymentProcessed)
				if err := _Plus.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
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
func (_Plus *PlusFilterer) ParsePaymentProcessed(log types.Log) (*PlusPaymentProcessed, error) {
	event := new(PlusPaymentProcessed)
	if err := _Plus.contract.UnpackLog(event, "PaymentProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusPersonalCashBackUpdatedIterator is returned from FilterPersonalCashBackUpdated and is used to iterate over the raw logs and unpacked data for PersonalCashBackUpdated events raised by the Plus contract.
type PlusPersonalCashBackUpdatedIterator struct {
	Event *PlusPersonalCashBackUpdated // Event containing the contract specifics and raw log

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
func (it *PlusPersonalCashBackUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusPersonalCashBackUpdated)
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
		it.Event = new(PlusPersonalCashBackUpdated)
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
func (it *PlusPersonalCashBackUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusPersonalCashBackUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusPersonalCashBackUpdated represents a PersonalCashBackUpdated event raised by the Plus contract.
type PlusPersonalCashBackUpdated struct {
	TokenId     *big.Int
	Owner       common.Address
	OldCashBack *big.Int
	NewCashBack *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPersonalCashBackUpdated is a free log retrieval operation binding the contract event 0xdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca.
//
// Solidity: event PersonalCashBackUpdated(uint256 indexed tokenId, address indexed owner, uint256 oldCashBack, uint256 newCashBack)
func (_Plus *PlusFilterer) FilterPersonalCashBackUpdated(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*PlusPersonalCashBackUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "PersonalCashBackUpdated", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &PlusPersonalCashBackUpdatedIterator{contract: _Plus.contract, event: "PersonalCashBackUpdated", logs: logs, sub: sub}, nil
}

// WatchPersonalCashBackUpdated is a free log subscription operation binding the contract event 0xdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca.
//
// Solidity: event PersonalCashBackUpdated(uint256 indexed tokenId, address indexed owner, uint256 oldCashBack, uint256 newCashBack)
func (_Plus *PlusFilterer) WatchPersonalCashBackUpdated(opts *bind.WatchOpts, sink chan<- *PlusPersonalCashBackUpdated, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "PersonalCashBackUpdated", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusPersonalCashBackUpdated)
				if err := _Plus.contract.UnpackLog(event, "PersonalCashBackUpdated", log); err != nil {
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
func (_Plus *PlusFilterer) ParsePersonalCashBackUpdated(log types.Log) (*PlusPersonalCashBackUpdated, error) {
	event := new(PlusPersonalCashBackUpdated)
	if err := _Plus.contract.UnpackLog(event, "PersonalCashBackUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusPriceSetterUpdatedIterator is returned from FilterPriceSetterUpdated and is used to iterate over the raw logs and unpacked data for PriceSetterUpdated events raised by the Plus contract.
type PlusPriceSetterUpdatedIterator struct {
	Event *PlusPriceSetterUpdated // Event containing the contract specifics and raw log

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
func (it *PlusPriceSetterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusPriceSetterUpdated)
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
		it.Event = new(PlusPriceSetterUpdated)
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
func (it *PlusPriceSetterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusPriceSetterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusPriceSetterUpdated represents a PriceSetterUpdated event raised by the Plus contract.
type PlusPriceSetterUpdated struct {
	OldPriceSetter common.Address
	NewPriceSetter common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceSetterUpdated is a free log retrieval operation binding the contract event 0x379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa40079.
//
// Solidity: event PriceSetterUpdated(address indexed oldPriceSetter, address indexed newPriceSetter)
func (_Plus *PlusFilterer) FilterPriceSetterUpdated(opts *bind.FilterOpts, oldPriceSetter []common.Address, newPriceSetter []common.Address) (*PlusPriceSetterUpdatedIterator, error) {

	var oldPriceSetterRule []interface{}
	for _, oldPriceSetterItem := range oldPriceSetter {
		oldPriceSetterRule = append(oldPriceSetterRule, oldPriceSetterItem)
	}
	var newPriceSetterRule []interface{}
	for _, newPriceSetterItem := range newPriceSetter {
		newPriceSetterRule = append(newPriceSetterRule, newPriceSetterItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "PriceSetterUpdated", oldPriceSetterRule, newPriceSetterRule)
	if err != nil {
		return nil, err
	}
	return &PlusPriceSetterUpdatedIterator{contract: _Plus.contract, event: "PriceSetterUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceSetterUpdated is a free log subscription operation binding the contract event 0x379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa40079.
//
// Solidity: event PriceSetterUpdated(address indexed oldPriceSetter, address indexed newPriceSetter)
func (_Plus *PlusFilterer) WatchPriceSetterUpdated(opts *bind.WatchOpts, sink chan<- *PlusPriceSetterUpdated, oldPriceSetter []common.Address, newPriceSetter []common.Address) (event.Subscription, error) {

	var oldPriceSetterRule []interface{}
	for _, oldPriceSetterItem := range oldPriceSetter {
		oldPriceSetterRule = append(oldPriceSetterRule, oldPriceSetterItem)
	}
	var newPriceSetterRule []interface{}
	for _, newPriceSetterItem := range newPriceSetter {
		newPriceSetterRule = append(newPriceSetterRule, newPriceSetterItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "PriceSetterUpdated", oldPriceSetterRule, newPriceSetterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusPriceSetterUpdated)
				if err := _Plus.contract.UnpackLog(event, "PriceSetterUpdated", log); err != nil {
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
func (_Plus *PlusFilterer) ParsePriceSetterUpdated(log types.Log) (*PlusPriceSetterUpdated, error) {
	event := new(PlusPriceSetterUpdated)
	if err := _Plus.contract.UnpackLog(event, "PriceSetterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Plus contract.
type PlusTransferIterator struct {
	Event *PlusTransfer // Event containing the contract specifics and raw log

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
func (it *PlusTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusTransfer)
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
		it.Event = new(PlusTransfer)
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
func (it *PlusTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusTransfer represents a Transfer event raised by the Plus contract.
type PlusTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Plus *PlusFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PlusTransferIterator, error) {

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

	logs, sub, err := _Plus.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlusTransferIterator{contract: _Plus.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Plus *PlusFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PlusTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Plus.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusTransfer)
				if err := _Plus.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Plus *PlusFilterer) ParseTransfer(log types.Log) (*PlusTransfer, error) {
	event := new(PlusTransfer)
	if err := _Plus.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Plus contract.
type PlusUnpausedIterator struct {
	Event *PlusUnpaused // Event containing the contract specifics and raw log

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
func (it *PlusUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusUnpaused)
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
		it.Event = new(PlusUnpaused)
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
func (it *PlusUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusUnpaused represents a Unpaused event raised by the Plus contract.
type PlusUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Plus *PlusFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PlusUnpausedIterator, error) {

	logs, sub, err := _Plus.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PlusUnpausedIterator{contract: _Plus.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Plus *PlusFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PlusUnpaused) (event.Subscription, error) {

	logs, sub, err := _Plus.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusUnpaused)
				if err := _Plus.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Plus *PlusFilterer) ParseUnpaused(log types.Log) (*PlusUnpaused, error) {
	event := new(PlusUnpaused)
	if err := _Plus.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlusUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Plus contract.
type PlusUpgradedIterator struct {
	Event *PlusUpgraded // Event containing the contract specifics and raw log

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
func (it *PlusUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusUpgraded)
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
		it.Event = new(PlusUpgraded)
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
func (it *PlusUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusUpgraded represents a Upgraded event raised by the Plus contract.
type PlusUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Plus *PlusFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PlusUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Plus.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PlusUpgradedIterator{contract: _Plus.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Plus *PlusFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PlusUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Plus.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusUpgraded)
				if err := _Plus.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Plus *PlusFilterer) ParseUpgraded(log types.Log) (*PlusUpgraded, error) {
	event := new(PlusUpgraded)
	if err := _Plus.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
