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
	CashBackConversionRate *big.Int
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existingTokenId\",\"type\":\"uint256\"}],\"name\":\"AgentAlreadyHasToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AlreadyActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notBeneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiaryIsNotAnAgent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"BeneficiaryOwnerIsNotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CredentialUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidDowngrade\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidUpgrade\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"MaxCashBackPercentExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoCashBackToClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPriceSetterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dtl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"OverLimitDtl\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenNotActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroCashBackConversionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroConversionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"CardActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"CardDowngraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"CardUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CashBackClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FilVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GlfVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"glfBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filEarned\",\"type\":\"uint256\"}],\"name\":\"PaymentProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCashBack\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCashBack\",\"type\":\"uint256\"}],\"name\":\"PersonalCashBackUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPriceSetter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPriceSetter\",\"type\":\"address\"}],\"name\":\"PriceSetterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentFactory\",\"outputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentPoliceVcVerifier\",\"outputs\":[{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseConversionRateFILtoGLF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"changeOwnerForAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimCashBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credential\",\"type\":\"bytes32\"}],\"name\":\"credentialHasBeenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"_sc\",\"type\":\"tuple\"}],\"name\":\"downgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundFilVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fundGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"fundGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getTimeUntilPenaltyFreeDowngrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glfToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"},{\"internalType\":\"contractIAgentPoliceVCVerifier\",\"name\":\"_agentPoliceVcVerifier\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_personalCashBackPercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintActivateAndFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"mintAndActivate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"onDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestAmount\",\"type\":\"uint256\"}],\"name\":\"onPaymentMade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"}],\"name\":\"setAgentFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseConversionRateFILtoGLF\",\"type\":\"uint256\"}],\"name\":\"setBaseConversionRateFILtoGLF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"}],\"name\":\"setGlfToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxCashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setMaxCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setPersonalCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"setPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceSetter\",\"type\":\"address\"}],\"name\":\"setPriceSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"_tierInfo\",\"type\":\"tuple\"}],\"name\":\"setTierInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyFee\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyWindow\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"tierToTierInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToAgentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToFilCashbackEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToGlfVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToLastTierSwitchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToPersonalCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToTier\",\"outputs\":[{\"internalType\":\"enumTier\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFilCashbackVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowDtl\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"_vc\",\"type\":\"tuple\"}],\"name\":\"validateAgentLeverage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawAllFilFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f5160206174415f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161737a90816100c78239608051818181612bd40152612d540152f35b6001600160401b0319166001600160401b039081175f5160206174415f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f358060e01c91826301ffc9a7146154f7575081630431a3251461548657816306fdde0314615342578163081812fc146152bb578163095ea7b31461525b5781630c1cd643146151fd5781630e6b46271461517b5781631755ff211461510b57816318de0c6a146150a45781631ac570891461420f5781631fac65d314613a8e575080631fe29f1814613a33578063238b87861461399d57806323b872dd1461396357806328d0e5a9146134ad5780632c09bef71461343c5780632fe5c820146133d55780633f4ba83a1461330357806342842e0e146132b957806343989f0a1461324d5780634449ee09146131e6578063463ad3431461301c5780634f1ef28614612ce95780634fa56e3b14612c7857806351c9749c14612c4c57806352d1902d14612b8e5780635c975abb14612b2e5780635df3548114612ac65780635e7277f314612a5f57806361d027b3146129ee5780636352211e1461299357806363791e3c146128535780636817c76c146127f85780636a627842146127b65780636b12bee31461275b57806370a0823114612688578063715018a614612546578063762b977d1461246757806379ba5097146123c55780637caf9073146122e65780637df107ea146122755780637dfe5b9214611fa057806381eebd7014611f405780638456cb5914611e6b5780638b44643114611e385780638da5cb5b14611dc75780638dd404d214611d6c5780638f68679f14611d115780639129347314611c345780639202adfc14611ad157806395bda1d814611a9457806395d89b41146119375780639780b372146118d0578063a1e80e5514611875578063a22cb46514611809578063abdf23df146117a9578063ad3cb1cc1461172a578063ad7f4929146116cf578063aef0353c146111f9578063b77da5331461115b578063b88d4fde146110d1578063c87b56dd14611058578063cda43c1014610fd8578063d0911b1e14610f67578063d7d2929e14610d20578063d86e1a3514610bec578063e30c397814610b7b578063e985e9c514610aca578063e9e15b4f14610a01578063ebca88fc146107a2578063f0f4426014610698578063f173e1fe14610631578063f2fde38b14610514578063f4a0a528146104b4578063f9d20f53146103c95763ffc2108914610364575f80fd5b346103c65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761039b61567d565b6103a36156c3565b906044359260048410156103c65760206103be858585616527565b604051908152f35b80fd5b50346103c65760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760043560048110156104b05760607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126104b057801581811561049b575b61044291615b1d565b61044a61698f565b6024359081156104735761045d90615cbc565b9081556044356001820155600260643591015580f35b6004837f55dbcc3c000000000000000000000000000000000000000000000000000000008152fd5b61044291506104a981616b95565b9150610439565b5080fd5b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576104ec61698f565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713005580f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65773ffffffffffffffffffffffffffffffffffffffff61056161567d565b61056961698f565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e83522054604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576106d061567d565b6106d861698f565b73ffffffffffffffffffffffffffffffffffffffff81161561077a576107779073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130355565b80f35b6004827fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b50346103c65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576024356004356044356107e3616b10565b82156109d9577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529060209082906064908290899073ffffffffffffffffffffffffffffffffffffffff165af180156109ce576109a1575b508184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313602052604084206108ab848254616010565b90558184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671310602052826040852054918386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260408620549080155f146109935750148061098b575b1561097857506109497feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c5482616c5f565b6040519182527f615eb0eb079a1bab414b72e513f67a7286540e31499244785534e19b2c35c1eb60203393a380f35b156109495761098681616bd1565b610949565b508015610918565b915050610986915082616c5f565b6109c29060203d6020116109c7575b6109ba8183615723565b810190615d2c565b610874565b503d6109b0565b6040513d87823e3d90fd5b6004847f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760043573ffffffffffffffffffffffffffffffffffffffff81168091036104b057610a5a61698f565b801561077a577fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045580f35b50346103c65760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657610b0261567d565b73ffffffffffffffffffffffffffffffffffffffff610b1f6156c3565b91165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760043573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015610ce0575b15610cb8578015610c90577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f5580f35b6004827f247af9ce000000000000000000000000000000000000000000000000000000008152fd5b6004827fd6d9d129000000000000000000000000000000000000000000000000000000008152fd5b5073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131254163314610c60565b50346103c65760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657610d5861567d565b60243590610d646156a0565b6064356004811015610f6357610d7d9160843593616527565b91610d86616b10565b81156109d9577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529060209082906064908290899073ffffffffffffffffffffffffffffffffffffffff165af180156109ce576020959284929091610f48575b508483527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313865260408320610e55838254616010565b90558483527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131086526040808420548685527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131388529320549080610f3a57501480610f32575b15610f1f5750610eeb7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c5483616c5f565b604051908152817f615eb0eb079a1bab414b72e513f67a7286540e31499244785534e19b2c35c1eb843393a3604051908152f35b15610eeb57610f2d82616bd1565b610eeb565b508015610eba565b915050610f2d915083616c5f565b610f5e90873d89116109c7576109ba8183615723565b610e1f565b8480fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416604051908152f35b50346103c657610fe736615854565b9073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671304541633036110305790610777916160ab565b6004837f4b602735000000000000000000000000000000000000000000000000000000008152fd5b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576110936004356168fb565b50806040516110a3602082615723565b52506110cd6040516110b6602082615723565b5f815260405191829160208352602083019061563a565b0390f35b50346103c65760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761110961567d565b506111126156c3565b5060643567ffffffffffffffff81116104b057906111356004923690840161579e565b507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760043560048110156104b0576111bb90606092604080516111ab81615707565b8281528260208201520152615cbc565b6040516111c781615707565b815491828252604060026001830154926020850193845201549201918252604051928352516020830152516040820152f35b50346103c65760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761123161567d565b6024356044359160048310156116cb57611249616b10565b611276611255836168fb565b83339173ffffffffffffffffffffffffffffffffffffffff33911614615a9b565b6112888361128381616b95565b615b1d565b61129182616b63565b61169f5783602061134160016112a687615cbc565b015473ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416906040519485809481937f23b872dd00000000000000000000000000000000000000000000000000000000835230336004850173ffffffffffffffffffffffffffffffffffffffff6040929594938160608401971683521660208201520152565b03925af180156109ce57611682575b508184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e6020524260408520558184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526113b28360408620615d57565b73ffffffffffffffffffffffffffffffffffffffff811690816114075750506113de604051809361562d565b7f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a380f35b602490602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713029694965416604051938480927ffd66091e0000000000000000000000000000000000000000000000000000000082528960048301525afa91821561163d578692611648575b5061149490821515615d8e565b6040517f8da5cb5b000000000000000000000000000000000000000000000000000000008152602081600481885afa90811561163d57869161160e575b5073ffffffffffffffffffffffffffffffffffffffff339116036115de578085527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713096020526040852054806115ae57507f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5918160209287527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130983528460408820558487527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d835260408720556115aa604051809261562d565ba380f35b85906044927f04e662d6000000000000000000000000000000000000000000000000000000008352600452602452fd5b604485857f98f08d3a00000000000000000000000000000000000000000000000000000000825260045233602452fd5b611630915060203d602011611636575b6116288183615723565b810190615dd8565b5f6114d1565b503d61161e565b6040513d88823e3d90fd5b9091506020813d60201161167a575b8161166460209383615723565b81010312611676575190611494611487565b5f80fd5b3d9150611657565b61169a9060203d6020116109c7576109ba8183615723565b611350565b602484837f11d5a84c000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f54604051908152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657506110cd60405161176b604082615723565b600581527f352e302e30000000000000000000000000000000000000000000000000000000602082015260405191829160208352602083019061563a565b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576117e161698f565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b5580f35b50346103c65760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761184161567d565b50602435801515036103c657807f8cd22d190000000000000000000000000000000000000000000000000000000060049252fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b54604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131083522054604051908152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760405190807f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301549061199782615888565b8085529160018116908115611a4f57506001146119d3575b6110cd846119bf81860382615723565b60405191829160208352602083019061563a565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930181527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b808210611a35575090915081016020016119bf826119af565b919260018160209254838588010152019101909291611a1c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b850190920192506119bf91508390506119af565b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760206103be60043561601d565b50346103c65760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760043573ffffffffffffffffffffffffffffffffffffffff611b216156c3565b611b29616b10565b611b42611b35846168fb565b8433918533911614615a9b565b16908115611c0c578083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131460205260408320548015611be05760207f05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc4700918386527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713148252856040812055611bd68186616dc8565b604051908152a380f35b602484837fa370a06b000000000000000000000000000000000000000000000000000000008252600452fd5b6004837fd92e233d000000000000000000000000000000000000000000000000000000008152fd5b50807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657611c66616b10565b3415611ce957611c97347feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131554616010565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131555604051348152817f1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d607460203393a380f35b807f1f2a20050000000000000000000000000000000000000000000000000000000060049252fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130654604051908152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c54604051908152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65780f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657611ea261698f565b611eaa616b10565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657611f7861698f565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c5580f35b50346103c65760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760043560243590600482101561227157611fe8616b10565b612015611ff4826168fb565b82339173ffffffffffffffffffffffffffffffffffffffff33911614615a9b565b6120278161202281616b63565b615aea565b6120348261128381616b95565b8083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130760205260ff604084205416916004831015612244578281111561220a578390602061213761209e600161208a88615cbc565b0154600161209786615cbc565b0154615cf2565b73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416906040519586809481937f23b872dd00000000000000000000000000000000000000000000000000000000835230336004850173ffffffffffffffffffffffffffffffffffffffff6040929594938160608401971683521660208201520152565b03925af19182156109ce576121c4926121ed575b508285527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526121808160408720615d57565b8285527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e6020524260408620556121ba604051809561562d565b602084019061562d565b7f7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd60403393a380f35b6122059060203d6020116109c7576109ba8183615723565b61214b565b8361224260449261223d867fa0f244d1000000000000000000000000000000000000000000000000000000008552615611565b61561f565bfd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b8280fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761231e61567d565b61232661698f565b73ffffffffffffffffffffffffffffffffffffffff81161561077a576107779073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130255565b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054160361243b5761077733616eb5565b807f118cdaa7000000000000000000000000000000000000000000000000000000006024925233600452fd5b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761249f61567d565b6124a761698f565b73ffffffffffffffffffffffffffffffffffffffff81161561077a576107779073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130155565b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761257d61698f565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00558073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576126c061567d565b9073ffffffffffffffffffffffffffffffffffffffff82161561272f5760206127268373ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b54604051908152f35b807f89c62b64000000000000000000000000000000000000000000000000000000006024925280600452fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131554604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760206103be6127f361567d565b615e04565b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130054604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761288b61567d565b61289361698f565b73ffffffffffffffffffffffffffffffffffffffff8116908115611c0c5761296c73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131254169173ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713125416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131255565b7f379e1448c87235649d75d5948244f5f2b41110976dce88c6b3687527daa400798380a380f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760206129d06004356168fb565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131383522054604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657610777612b0161567d565b612b09616d45565b612b1161698f565b73ffffffffffffffffffffffffffffffffffffffff479116616dc8565b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003612c245760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b807fe07c8dba0000000000000000000000000000000000000000000000000000000060049252fd5b50346103c657610777612c5e36615854565b90612c67616b10565b612c73611ff4826168fb565b616c5f565b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416604051908152f35b5060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657612d1c61567d565b9060243567ffffffffffffffff81116104b057612d3d90369060040161579e565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115612fda575b50612fb257612d8c61698f565b73ffffffffffffffffffffffffffffffffffffffff831690604051937f52d1902d000000000000000000000000000000000000000000000000000000008552602085600481865afa80958596612f7e575b50612e0e57602484847f4c9c8ce3000000000000000000000000000000000000000000000000000000008252600452fd5b9091847f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8103612f535750813b15612f2857807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8480a28151839015612ef55780836020612ef195519101845af4612eeb616d99565b916172ab565b5080f35b50505034612f005780f35b807fb398979f0000000000000000000000000000000000000000000000000000000060049252fd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000008452600452602483fd5b7faa1d49a4000000000000000000000000000000000000000000000000000000008552600452602484fd5b9095506020813d602011612faa575b81612f9a60209383615723565b810103126116765751945f612ddd565b3d9150612f8d565b6004827fe07c8dba000000000000000000000000000000000000000000000000000000008152fd5b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614155f612d7f565b50346103c65761302b36615854565b90613034616b10565b61303c616b10565b81156131be577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529060209082906064908290889073ffffffffffffffffffffffffffffffffffffffff165af180156131b357613196575b508083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260408320613104838254616010565b90558083527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260408320548184527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313602052826040852054148061098b571561097857506109497feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c5482616c5f565b6131ae9060203d6020116109c7576109ba8183615723565b6130cd565b6040513d86823e3d90fd5b6004837f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d83522054604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760ff604060209260043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131184522054166040519015158152f35b50346103c6576004906132cb366157e2565b505050806040516132dd602082615723565b527f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65761333a61698f565b613342616d45565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657604060209160043581527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131483522054604051908152f35b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c657602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713125416604051908152f35b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576134e561567d565b906134ee616b10565b73ffffffffffffffffffffffffffffffffffffffff821691821561077a576024602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051928380927f1ffbb0640000000000000000000000000000000000000000000000000000000082528860048301525afa90811561390d579061359492918491613944575b50615d8e565b604051917f8da5cb5b000000000000000000000000000000000000000000000000000000008352602083600481845afa928315613939578293613918575b5073ffffffffffffffffffffffffffffffffffffffff8316908115611c0c576020600491604051928380927faf640d0f0000000000000000000000000000000000000000000000000000000082525afa90811561390d5783916138db575b5082527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260408220549073ffffffffffffffffffffffffffffffffffffffff61367b836168fb565b16938185146138b357602094604051916136958784615723565b85835273ffffffffffffffffffffffffffffffffffffffff6136b78683617094565b16806136e957602487877f7e273289000000000000000000000000000000000000000000000000000000008252600452fd5b8281969495960361387f57503b6136fe578480f35b90859161375060405194859384937f150b7a020000000000000000000000000000000000000000000000000000000085523360048601526024850152604484015260806064840152608483019061563a565b038186865af1839181613827575b506137a6575061376c616d99565b805193846137a057602484847f64a0ae92000000000000000000000000000000000000000000000000000000008252600452fd5b84925001fd5b7f150b7a0200000000000000000000000000000000000000000000000000000000919293507fffffffff0000000000000000000000000000000000000000000000000000000016036137fc5750805f8080808480f35b7f64a0ae92000000000000000000000000000000000000000000000000000000008252600452602490fd5b9091508481813d8311613878575b61383f8183615723565b810103126116cb57517fffffffff00000000000000000000000000000000000000000000000000000000811681036116cb57905f61375e565b503d613835565b8660649185857f64283d7b000000000000000000000000000000000000000000000000000000008452600452602452604452fd5b6004847f2a63c7cc000000000000000000000000000000000000000000000000000000008152fd5b90506020813d602011613905575b816138f660209383615723565b8101031261167657515f613630565b3d91506138e9565b6040513d85823e3d90fd5b61393291935060203d602011611636576116288183615723565b915f6135d2565b6040513d84823e3d90fd5b61395d915060203d6020116109c7576109ba8183615723565b5f61358e565b50346103c657600490613975366157e2565b5050507f8cd22d19000000000000000000000000000000000000000000000000000000008152fd5b50346103c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c6576004358015613a0b576040826129d092602094527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671309845220546168fb565b6004827fac8fb3c1000000000000000000000000000000000000000000000000000000008152fd5b50346103c657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103c65760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a54604051908152f35b9050346116765760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126116765760043560243560048110156116765760443567ffffffffffffffff811161167657806004018136039460807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc87011261167657613b1a616b10565b613b47613b26866168fb565b86339173ffffffffffffffffffffffffffffffffffffffff33911614615a9b565b613b548561202281616b63565b83158481156141fa575b613b6791615b1d565b845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130760205260ff60405f2054169560048710156141cd578685101561419857908794939291865f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d60205260405f20549384613eda575b5050505050613c096001613bf486615cbc565b01546001613c0184615cbc565b015490615cf2565b8383527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e602052613c3e604084205442615cf2565b9073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416917feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a54115f14613e6357613d5a916020613ce0612710613cd87feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b5486615d44565b048094615cf2565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101949094529293849081906044820190565b038188855af19182156109ce57602093613db793613e48575b506040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481019190915294859283919082906044820190565b03925af19182156109ce57613e0092613e2957505b8285527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526121808160408720615d57565b7f65e7b85a0b1fe94cdd38e2a9dbe4ac14dda1ceeb998856160d0ea99b804b105d60403393a380f35b613e419060203d6020116109c7576109ba8183615723565b505f61214b565b613e5e90853d87116109c7576109ba8183615723565b613d73565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815233600482015260248101919091529260209184916044918391905af19182156109ce57613e0092613ebb575b50613dcc565b613ed39060203d6020116109c7576109ba8183615723565b505f613eb5565b613f5f95965073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416916020613f258680615b56565b604051809981927fa2f3c2100000000000000000000000000000000000000000000000000000000083528460048401526024830190615b89565b0381865afa968715614131575f97614164575b50865f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131160205260ff60405f20541661413c57823b15611676577fffffffff00000000000000000000000000000000000000000000000000000000604051947f403bb79d000000000000000000000000000000000000000000000000000000008652876004870152166024850152606060448501527ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffefd853591018112156116765761404e906080606486015260048360e48701920101615b89565b9160248201359260ff8416809403611676575f858094926064829484986084850152604481013560a4850152013560c483015203925af1801561413157614113575b508392916140b6916140b060026140a78b98615cbc565b01549180615b56565b916158ee565b82527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713116020526040822060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790555f80808080613be1565b6141239194939297505f90615723565b6140b65f9691929390614090565b6040513d5f823e3d90fd5b7fae9e9f65000000000000000000000000000000000000000000000000000000005f5260045ffd5b9096506020813d602011614190575b8161418060209383615723565b810103126116765751955f613f72565b3d9150614173565b6141c88561223d897f253471cf000000000000000000000000000000000000000000000000000000005f52615611565b60445ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b613b67915061420881616b95565b9150613b5e565b34611676576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126116765760043567ffffffffffffffff81116116765761425f90369060040161579e565b60243567ffffffffffffffff81116116765761427f90369060040161579e565b6142876156a0565b60843573ffffffffffffffffffffffffffffffffffffffff811681036116765760a4359073ffffffffffffffffffffffffffffffffffffffff821682036116765760c4359173ffffffffffffffffffffffffffffffffffffffff831683036116765760e4359473ffffffffffffffffffffffffffffffffffffffff8616809603611676577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549660ff8860401c16159767ffffffffffffffff81168015908161509c575b6001149081615092575b159081615089575b50615061578860017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005561500c575b506143b9616fc4565b6143c1616fc4565b80519067ffffffffffffffff8211614e0d5781906143ff7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054615888565b601f8111614f42575b50602090601f8311600114614e45575f92614e3a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b80519067ffffffffffffffff8211614e0d5781906144ae7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154615888565b601f8111614d80575b50602090601f8311600114614ca1575f92614c96575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b614527616fc4565b61452f616fc4565b614537616fc4565b73ffffffffffffffffffffffffffffffffffffffff841615614c6a5761461a6146999261456661471896616eb5565b61456e616fc4565b614576616fc4565b6064357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713005573ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130155565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130355565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130255565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055560017feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671306556101f47feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c5568056bc75e2d631000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f556148ae73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713125416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131255565b6040516148ba81615707565b5f8082526020808301828152670a688906bd8b000060408086019182529380527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890925292517f98ececc9ef5447155a71900bf0a10a038a903deb6b29c18bf41c012bdacdfd785591517f98ececc9ef5447155a71900bf0a10a038a903deb6b29c18bf41c012bdacdfd795590517f98ececc9ef5447155a71900bf0a10a038a903deb6b29c18bf41c012bdacdfd7a555161497481615707565b68056bc75e2d631000008152690a968163f0a57b4000006020808301918252670b1a2bc2ec500000604080850191825260015f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890925292517f2f7f69aa2d75e15f19ce52af93df9a7b1c1e6ef7db7d165311b61bd2c161868d5590517f2f7f69aa2d75e15f19ce52af93df9a7b1c1e6ef7db7d165311b61bd2c161868e5590517f2f7f69aa2d75e15f19ce52af93df9a7b1c1e6ef7db7d165311b61bd2c161868f5551614a4281615707565b680595698248593c000081526934f086f3b33b684000006020808301918252670bbe2470a1549d8a604080850191825260025f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890925292517fcca1d77c34f404c66c45b233ec66db9f5c52d27866d7a7c46dbd420c1ff4d8cf5590517fcca1d77c34f404c66c45b233ec66db9f5c52d27866d7a7c46dbd420c1ff4d8d05590517fcca1d77c34f404c66c45b233ec66db9f5c52d27866d7a7c46dbd420c1ff4d8d15551614b1081615707565b6805f68e8131ecf80000815269d3c21bcecceda10000006020808301918252670c249fdd327780006040840190815260035f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130890915291517fe6edec82ca249256f7285fb9fd1ebbc810fa32d98aed0f790e077004d6c9b18a55517fe6edec82ca249256f7285fb9fd1ebbc810fa32d98aed0f790e077004d6c9b18b55517fe6edec82ca249256f7285fb9fd1ebbc810fa32d98aed0f790e077004d6c9b18c55614bd757005b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b7f1e4fbdf7000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b0151905088806144cd565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f52815f20925f5b818110614d685750908460019594939210614d31575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015561451f565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055888080614d04565b92936020600181928786015181550195019301614cee565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f529091507ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f840160051c81019160208510614e03575b90601f859493920160051c01905b818110614df557506144b7565b5f8155849350600101614de8565b9091508190614dda565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b01519050898061441e565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81925f5b818110614f2a5750908460019594939210614ef3575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055614470565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080614ec6565b92936020600181928786015181550195019301614eb0565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f52601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410614fe4575b90601f8493920160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b818110614fd65750614408565b5f8155849350600101614fc9565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf819150614f9b565b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055886143b0565b7ff92ee8a9000000000000000000000000000000000000000000000000000000005f5260045ffd5b9050158a61435d565b303b159150614355565b8a915061434b565b346116765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611676576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671309602052602060405f2054604051908152f35b34611676575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261167657602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416604051908152f35b346116765760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126116765760443567ffffffffffffffff8111611676576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112611676576151fb906004016024356004356158ee565b005b346116765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126116765761523461698f565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a55005b346116765760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126116765761529261567d565b507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b346116765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611676576004356152f6816168fb565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b34611676575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611676576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005461539f81615888565b808452906001811690811561544457506001146153c7575b6110cd836119bf81850382615723565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b80821061542a575090915081016020016119bf6153b7565b919260018160209254838588010152019101909291615412565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506119bf90506153b7565b346116765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611676576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671307602052602060ff60405f2054166154f5604051809261562d565bf35b346116765760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261167657600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361167657817f80ac58cd00000000000000000000000000000000000000000000000000000000602093149081156155ba575b8115615590575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483615589565b7f5b5e139f0000000000000000000000000000000000000000000000000000000081149150615582565b35907fffffffff000000000000000000000000000000000000000000000000000000008216820361167657565b60048110156141cd57600452565b60048110156141cd57602452565b9060048210156141cd5752565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361167657565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361167657565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361167657565b359073ffffffffffffffffffffffffffffffffffffffff8216820361167657565b6060810190811067ffffffffffffffff821117614e0d57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117614e0d57604052565b67ffffffffffffffff8111614e0d57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f82011215611676576020813591016157b882615764565b926157c66040519485615723565b8284528282011161167657815f92602092838601378301015290565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60609101126116765760043573ffffffffffffffffffffffffffffffffffffffff81168103611676579060243573ffffffffffffffffffffffffffffffffffffffff81168103611676579060443590565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112611676576004359060243590565b90600182811c921680156158cf575b60208310146158a257565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691615897565b359067ffffffffffffffff8216820361167657565b91602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416936024604051809681937f9f4326d700000000000000000000000000000000000000000000000000000000835260048301525afa928315614131575f93615a67575b506101008136031261167657604051610100810181811067ffffffffffffffff821117614e0d5760405261599e826156e6565b8152602082013560208201526040820135604082015260608201356060820152608082013560808201526159d460a083016155e4565b60a08201526159e560c083016158d9565b60c082015260e082013567ffffffffffffffff811161167657615a1992615a0e9136910161579e565b60e0820152836169fb565b828211615a265750505050565b60849450604051937f4ec0c37d0000000000000000000000000000000000000000000000000000000085526004850152602484015260448301526064820152fd5b9092506020813d602011615a93575b81615a8360209383615723565b810103126116765751915f61596b565b3d9150615a76565b15615aa4575050565b73ffffffffffffffffffffffffffffffffffffffff92507f3a0302a9000000000000000000000000000000000000000000000000000000005f526004521660245260445ffd5b15615af25750565b7fed5a980a000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b15615b255750565b615b51907fbca1a956000000000000000000000000000000000000000000000000000000005f52615611565b60245ffd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215611676570190565b73ffffffffffffffffffffffffffffffffffffffff615ba7826156e6565b168252602081013560208301526040810135604083015260608101356060830152608081013560808301527fffffffff00000000000000000000000000000000000000000000000000000000615bff60a083016155e4565b1660a083015267ffffffffffffffff615c1a60c083016158d9565b1660c083015260e08101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156116765701906020823592019167ffffffffffffffff811161167657803603831361167657601f817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0926101209561010060e087015281610100870152868601375f8582860101520116010190565b60048110156141cd575f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130860205260405f2090565b91908203918211615cff57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b90816020910312611676575180151581036116765790565b81810292918115918404141715615cff57565b9060048110156141cd5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b15615d965750565b73ffffffffffffffffffffffffffffffffffffffff907f3a35a1b9000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b90816020910312611676575173ffffffffffffffffffffffffffffffffffffffff811681036116765790565b615e0c616b10565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671300546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9283166024820152604481019190915291602091839160649183915f91165af1801561413157615ff3575b507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214615cff57615f5890600183017feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130655616e4d565b73ffffffffffffffffffffffffffffffffffffffff811615615fc757615f938273ffffffffffffffffffffffffffffffffffffffff92617094565b16615f9b5790565b7f73c6ac6e000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b7f64a0ae92000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b61600b9060203d6020116109c7576109ba8183615723565b615ede565b91908201809211615cff57565b616026816168fb565b5061603081616b63565b156160a6575f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e60205261606960405f205442615cf2565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a5490818110156160a05761609d91615cf2565b90565b50505f90565b505f90565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300541661652357811561652357805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260405f20549182156164f15761611683616b63565b156164f157825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260405f2054156164f157825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260405f205480156164f6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0481116164f157825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526161dd60ff60405f205416615cbc565b6002604051916161ec83615707565b805483526001810154602084015201546040820152835f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205261271061623960405f205484615d44565b0491845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260405f20549182156164e9577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f549051670de0b6b3a764000085810283900402819004938481106164cd575b507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131554918583106164ad575b50505081156164a6577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810185905291602091839160449183915f91165af19081616489575b50616394575050505050565b7f299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d692606092865f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260405f206163ef828254615cf2565b905561641c827feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131554615cf2565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131555865f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131460205260405f20616474838254616010565b905560405192835260208301526040820152a3565b6164a19060203d6020116109c7576109ba8183615723565b616388565b5050505050565b90929350819450670de0b6b3a764000080929502040204905f80806162da565b93509350670de0b6b3a76400008481838602040204935f6162ae565b505050505050565b505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5050565b61653090615e04565b91616539616b10565b616566616545846168fb565b84339173ffffffffffffffffffffffffffffffffffffffff33911614615a9b565b6165738161128381616b95565b61657c83616b63565b6168cf575f602061659160016112a685615cbc565b03925af18015614131576168b2575b50825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e6020524260405f2055825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713076020526166028160405f20615d57565b73ffffffffffffffffffffffffffffffffffffffff821691826166595750905061662f604051809261562d565b817f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a390565b602490602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051938480927ffd66091e0000000000000000000000000000000000000000000000000000000082528860048301525afa918215614131575f9261687c575b506166e390821515615d8e565b6040517f8da5cb5b000000000000000000000000000000000000000000000000000000008152602081600481875afa908115614131575f9161685d575b5073ffffffffffffffffffffffffffffffffffffffff3391160361682d57805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260405f2054806167fe575060208492827f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5935f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130983528460405f2055845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d835260405f20556167fa604051809261562d565ba390565b907f04e662d6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b827f98f08d3a000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b616876915060203d602011611636576116288183615723565b5f616720565b9091506020813d6020116168aa575b8161689860209383615723565b810103126116765751906166e36166d6565b3d915061688b565b6168ca9060203d6020116109c7576109ba8183615723565b6165a0565b827f11d5a84c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b616943815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff821615616964575090565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633036169cf57565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b919060e00151916101408380518101031261167657604051610140810181811067ffffffffffffffff821117614e0d5760405260208401518152610120610140604086015195866020850152606081015160408501526080810151606085015260a0810151608085015260c081015160a085015260e081015160c085015261010081015160e0850152828101516101008501520151910152828115616b0857508215616ae2577812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f228110830215616ad557670de0b6b3a7640000839102049190565b637c5f487d5f526004601cfd5b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9190565b9250505f9190565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416616b3b57565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130760205261609d60ff60405f2054165b600481101590816141cd5760018114918215616bc4575b8215616bb757505090565b9091506141cd5760031490565b506002811491505f616bac565b805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260405f205490805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020525f60408120556040519182525f60208301527fdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca60403393a3565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c54808311616d165750805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131060205260405f205491815f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020528060405f205560405192835260208301527fdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca60403393a3565b827f37829ee6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300541615616d7157565b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b3d15616dc3573d90616daa82615764565b91616db86040519384615723565b82523d5f602084013e565b606090565b814710616e25575f80809373ffffffffffffffffffffffffffffffffffffffff8294165af1616df5616d99565b5015616dfd57565b7f3204506f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f356680b7000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014616ea1575b15616e9d57616e8f9061701b565b90616e98575090565b905090565b5090565b505067ffffffffffffffff81166001616e81565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615616ff357565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff1603617089575b8215801561707e575b61707657565b5f9250829150565b5060163d1415617070565b5f9250829150617067565b906170dd815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9173ffffffffffffffffffffffffffffffffffffffff8316806171e8575b73ffffffffffffffffffffffffffffffffffffffff82169182617194575b50825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f20827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a490565b6171db9073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190555f617119565b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff000000000000000000000000000000000000000081541690556172808473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190556170fb565b906172e857508051156172c057805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b8151158061733b575b6172f9575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b156172f156fea2646970667358221220d33f2561824cfec7f457a849dffaaececac9121a9bb6c3620aba7d2e3bd58b7564736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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

// AgentIdToOwner is a free data retrieval call binding the contract method 0x238b8786.
//
// Solidity: function agentIdToOwner(uint256 _agentId) view returns(address owner)
func (_Plus *PlusCaller) AgentIdToOwner(opts *bind.CallOpts, _agentId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "agentIdToOwner", _agentId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentIdToOwner is a free data retrieval call binding the contract method 0x238b8786.
//
// Solidity: function agentIdToOwner(uint256 _agentId) view returns(address owner)
func (_Plus *PlusSession) AgentIdToOwner(_agentId *big.Int) (common.Address, error) {
	return _Plus.Contract.AgentIdToOwner(&_Plus.CallOpts, _agentId)
}

// AgentIdToOwner is a free data retrieval call binding the contract method 0x238b8786.
//
// Solidity: function agentIdToOwner(uint256 _agentId) view returns(address owner)
func (_Plus *PlusCallerSession) AgentIdToOwner(_agentId *big.Int) (common.Address, error) {
	return _Plus.Contract.AgentIdToOwner(&_Plus.CallOpts, _agentId)
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

// ValidateAgentLeverage is a free data retrieval call binding the contract method 0x0e6b4627.
//
// Solidity: function validateAgentLeverage(uint256 _agentId, uint256 borrowDtl, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) _vc) view returns()
func (_Plus *PlusCaller) ValidateAgentLeverage(opts *bind.CallOpts, _agentId *big.Int, borrowDtl *big.Int, _vc VerifiableCredential) error {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "validateAgentLeverage", _agentId, borrowDtl, _vc)

	if err != nil {
		return err
	}

	return err

}

// ValidateAgentLeverage is a free data retrieval call binding the contract method 0x0e6b4627.
//
// Solidity: function validateAgentLeverage(uint256 _agentId, uint256 borrowDtl, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) _vc) view returns()
func (_Plus *PlusSession) ValidateAgentLeverage(_agentId *big.Int, borrowDtl *big.Int, _vc VerifiableCredential) error {
	return _Plus.Contract.ValidateAgentLeverage(&_Plus.CallOpts, _agentId, borrowDtl, _vc)
}

// ValidateAgentLeverage is a free data retrieval call binding the contract method 0x0e6b4627.
//
// Solidity: function validateAgentLeverage(uint256 _agentId, uint256 borrowDtl, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) _vc) view returns()
func (_Plus *PlusCallerSession) ValidateAgentLeverage(_agentId *big.Int, borrowDtl *big.Int, _vc VerifiableCredential) error {
	return _Plus.Contract.ValidateAgentLeverage(&_Plus.CallOpts, _agentId, borrowDtl, _vc)
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

// Initialize is a paid mutator transaction binding the contract method 0x1ac57089.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier) returns()
func (_Plus *PlusTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "initialize", _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1ac57089.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier) returns()
func (_Plus *PlusSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Initialize(&_Plus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1ac57089.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, uint256 _mintPrice, address _glfToken, address _treasury, address _agentFactory, address _agentPoliceVcVerifier) returns()
func (_Plus *PlusTransactorSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _treasury common.Address, _agentFactory common.Address, _agentPoliceVcVerifier common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Initialize(&_Plus.TransactOpts, _name, _symbol, _initialOwner, _mintPrice, _glfToken, _treasury, _agentFactory, _agentPoliceVcVerifier)
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
	TokenId *big.Int
	Owner   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGlfVaultFunded is a free log retrieval operation binding the contract event 0x615eb0eb079a1bab414b72e513f67a7286540e31499244785534e19b2c35c1eb.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
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

// WatchGlfVaultFunded is a free log subscription operation binding the contract event 0x615eb0eb079a1bab414b72e513f67a7286540e31499244785534e19b2c35c1eb.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
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

// ParseGlfVaultFunded is a log parse operation binding the contract event 0x615eb0eb079a1bab414b72e513f67a7286540e31499244785534e19b2c35c1eb.
//
// Solidity: event GlfVaultFunded(uint256 indexed tokenId, address indexed owner, uint256 amount)
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
