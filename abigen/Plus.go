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

// TierInfo is an auto generated low-level Go binding around an user-defined struct.
type TierInfo struct {
	CashBackConversionRate *big.Int
	TokenLockAmount        *big.Int
	DebtToLiquidationValue *big.Int
}

// PlusMetaData contains all meta data concerning the Plus contract.
var PlusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existingTokenId\",\"type\":\"uint256\"}],\"name\":\"AgentAlreadyHasToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AlreadyActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notBeneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiaryIsNotAnAgent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"BeneficiaryOwnerIsNotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"reasons\",\"type\":\"uint8\"}],\"name\":\"InvalidCredential\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidDowngrade\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumTier\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidUpgrade\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"MaxCashBackPercentExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoCashBackToClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"notOwner\",\"type\":\"address\"}],\"name\":\"NotCardOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPool\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"faultySectors\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liveSectors\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTolerance\",\"type\":\"uint256\"}],\"name\":\"OverFaultySectorLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dtl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"OverLimitDtl\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenNotActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroCashBackConversionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroConversionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"CardActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"CardDowngraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumTier\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"CardUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CashBackClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FilVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GlfVaultFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"glfBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filEarned\",\"type\":\"uint256\"}],\"name\":\"PaymentProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCashBack\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCashBack\",\"type\":\"uint256\"}],\"name\":\"PersonalCashBackUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_VERIFIABLE_CREDENTIAL_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentFactory\",\"outputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"agentIdToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"changeOwnerForAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimCashBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credParser\",\"outputs\":[{\"internalType\":\"contractCredParser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credential\",\"type\":\"bytes32\"}],\"name\":\"credentialHasBeenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"deriveStructHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"_sc\",\"type\":\"tuple\"}],\"name\":\"downgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"filToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fundFilVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fundGlfVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"getTierInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"tierInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glfToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vcIssuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_filToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"},{\"internalType\":\"contractCredParser\",\"name\":\"_credParser\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_personalCashBackPercent\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_personalCashBackPercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintActivateAndFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_personalCashBackPercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"mintAndActivate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestAmount\",\"type\":\"uint256\"}],\"name\":\"onPaymentMade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolAddress\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sectorFaultyTolerancePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"}],\"name\":\"setAgentFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseConversionRate\",\"type\":\"uint256\"}],\"name\":\"setBaseConversionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCredParser\",\"name\":\"_credParser\",\"type\":\"address\"}],\"name\":\"setCredParser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_filToken\",\"type\":\"address\"}],\"name\":\"setFilToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_glfToken\",\"type\":\"address\"}],\"name\":\"setGlfToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxCashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setMaxCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cashBackPercent\",\"type\":\"uint256\"}],\"name\":\"setPersonalCashBackPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"setPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sectorFaultyTolerancePercent\",\"type\":\"uint256\"}],\"name\":\"setSectorFaultyTolerancePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"_tierInfo\",\"type\":\"tuple\"}],\"name\":\"setTierInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyFee\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierSwitchPenaltyWindow\",\"type\":\"uint256\"}],\"name\":\"setTierSwitchPenaltyWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vcIssuer\",\"type\":\"address\"}],\"name\":\"setVcIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierSwitchPenaltyWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"tierToTierInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cashBackConversionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenLockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToLiquidationValue\",\"type\":\"uint256\"}],\"internalType\":\"structTierInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToAgentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToFilCashbackEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToGlfVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToLastTierSwitchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToPersonalCashBackPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToTier\",\"outputs\":[{\"internalType\":\"enumTier\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFilCashbackVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumTier\",\"name\":\"_desiredTier\",\"type\":\"uint8\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowDtl\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"_vc\",\"type\":\"tuple\"}],\"name\":\"validateAgentLeverage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"validateCred\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vcIssuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawAllFilFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f5160206186fb5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161863490816100c7823960805181818161429601526144120152f35b6001600160401b0319166001600160401b039081175f5160206186fb5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60a06040526004361015610011575f80fd5b5f6080525f358060e01c9081630174174414615d1957816301ffc9a714615c295781630431a32514615bb657816306fdde0314615ab5578163081812fc14615a2e578163095ea7b3146159f75781630c1cd643146159995781630e6b4627146159195781631755ff21146158a957816318de0c6a146158425781631fac65d314615193575080631fe29f1814615139578063238b8786146150fd57806323b872dd146150c457806328d0e5a914614c3b5780632b85012b14614c035780632fe5c82014614b9c5780633727948814614b425780633f4ba83a14614a72578063403bb79d146149c857806340444c7c146149b057806340c10f191461496b57806342842e0e1461492257806343989f0a146148b65780634449ee091461484f578063463ad343146147235780634ead1527146146a85780634f1ef286146143a95780634fa56e3b1461433957806351c9749c1461430e57806352d1902d146142515780635c975abb146141f25780635df35481146140765780635e7277f31461400f57806361d027b314613f9f5780636352211e14613f45578063680af72414613ee75780636817c76c14613e8d5780636b12bee314613e3357806370a0823114613d62578063715018a614613c2257806373da692414613b44578063762b977d14613a6657806377c0f07d146139f657806379ba5097146139555780637caf9073146138775780637df107ea146138075780637dfe5b92146135045780637ed85ff314613486578063800ca2ae1461342a57806381eebd701461339257806383596c20146132055780638456cb591461313257806384b0196e14612f8d578063863de77014612ead5780638776d0d614612e3d5780638da5cb5b14612dcd5780638dd404d214612d735780638e18d64814612d035780638f68679f14612ca95780639202adfc14612a7457806395d89b411461291c5780639780b372146128b5578063a1e80e551461285b578063a22cb465146127f0578063a2f3c210146127d0578063abdf23df14612772578063ad3cb1cc146126f5578063aef0353c146121c3578063b2d4563814612169578063b77da533146120a4578063b88d4fde1461201a578063c87b56dd14611f9f578063cda43c1014611f1e578063d72422d814611e8a578063d7d2929e14611cba578063e30c397814611c48578063e985e9c514611b98578063e9e15b4f14611acb578063ea982a9d146107f3578063f0f44260146106e3578063f173e1fe14610678578063f2fde38b14610555578063f4a0a528146104f35763f9d20f53146103c7575f80fd5b346104ed5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed5760043560048110156104ed5760607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126104ed5761043581615e1a565b80158181156104d8575b61044891616718565b6104506178e4565b6024359081156104ac5761046381615e1a565b608051527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60205260406080512090815560443560018201556002606435910155608051608051f35b7f55dbcc3c00000000000000000000000000000000000000000000000000000000608051526004608051fd5b61044891506104e681617fc7565b915061043f565b60805180fd5b346104ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed5761052a6178e4565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713005560805180f35b346104ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed5773ffffffffffffffffffffffffffffffffffffffff6105a1615df7565b6105a96178e4565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700608051608051a360805180f35b346104ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed57600435608051527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671310602052602060406080512054604051908152f35b346104ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed5761071a615df7565b6107226178e4565b73ffffffffffffffffffffffffffffffffffffffff8116156107c7576107c19073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130455565b60805180f35b7fd92e233d00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346104ed576101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed5760043567ffffffffffffffff81116104ed576108439036906004016161ac565b60243567ffffffffffffffff81116104ed576108639036906004016161ac565b60443567ffffffffffffffff81116104ed576108839036906004016161ac565b906064359173ffffffffffffffffffffffffffffffffffffffff831692838103611ac7576084359473ffffffffffffffffffffffffffffffffffffffff861693848703611ac75760c4359073ffffffffffffffffffffffffffffffffffffffff82168203611ac75760e4359573ffffffffffffffffffffffffffffffffffffffff87168703611ac757610104359173ffffffffffffffffffffffffffffffffffffffff83168303611ac757610124359473ffffffffffffffffffffffffffffffffffffffff86168603611ac757610144359773ffffffffffffffffffffffffffffffffffffffff89168903611ac7577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549a8b67ffffffffffffffff811680159182611ab8575b506001149081611aae575b159081611aa5575b50611a795760017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008d16177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005560ff8c60401c1615611a24575b156107c757610a24618430565b610a2c618430565b815167ffffffffffffffff811161156057610a677fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102546161f0565b601f81116119b9575b508060206001601f8311146118d957608051916118ce575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102555b80519067ffffffffffffffff8211611560578190610b177fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103546161f0565b601f8111611858575b506020906001601f841114611776576080519261176b575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103555b6080517fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100556080517fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10155610bdc618430565b610be4618430565b80519067ffffffffffffffff8211611560578190610c227f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300546161f0565b601f811161169d575b506020906001601f84111461159c5760805192611591575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b80519067ffffffffffffffff8211611560578190610cd37f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301546161f0565b601f81116114ea575b506020906001601f84111461140857608051926113fd575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b610d4e618430565b610d56618430565b610d5e618430565b156113cb57610f356110b296610eb6610fb494610e1261103398610d8360ff9e617eb8565b610d8b618430565b610d93618430565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713065416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130655565b60a4357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713005573ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130155565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130255565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130455565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130355565b73ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713075416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130755565b60017feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671308556101f47feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e5568056bc75e2d631000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713115566038d7ea4c680007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713145560405161115d81615f4c565b608051815260026020820160805181526040830190670a688906bd8b00008252608051608051527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60205260406080512093518455516001840155519101556040516111c881615f4c565b68056bc75e2d631000008152600260208201690a968163f0a57b40000081526040830190670b1a2bc2ec50000082526001608051527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a602052604060805120935184555160018401555191015560405161124181615f4c565b680595698248593c000081526002602082016934f086f3b33b6840000081526040830190670bbe2470a1549d8a825282608051527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60205260406080512093518455516001840155519101556040516112b981615f4c565b6805f68e8131ecf80000815260026020820169d3c21bcecceda100000081526040830190670c249fdd3277800082526003608051527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a602052604060805120935184555160018401555191015560401c16156113355760805180f35b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a16107c1565b7f1e4fbdf700000000000000000000000000000000000000000000000000000000608051526080516004526024608051fd5b015190508b80610cf4565b608080517f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301905251828120937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016905b8181106114d2575090846001959493921061149b575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930155610d46565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558b808061146e565b92936020600181928786015181550195019301611458565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930160805152602060805120601f840160051c81019160208510611556575b90601f859493920160051c01905b8181106115465750610cdc565b6080518155849350600101611539565b909150819061152b565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b015190508c80610c43565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300608051527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81926080515b818110611685575090846001959493921061164e575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055610c95565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558c8080611621565b9293602060018192878601518155019501930161160b565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930060805152601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410611743575b90601f8493920160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b8181106117335750610c2b565b6080518155849350600101611726565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8191506116f8565b015190508d80610b38565b608080517fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103905251828120937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016905b8181106118405750908460019594939210611809575b505050811b017fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10355610b8a565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806117dc565b929360206001819287860151815501950193016117c6565b9091507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10360805152602060805120601f840160051c810191602085106118c4575b90601f859493920160051c01905b8181106118b45750610b20565b60805181558493506001016118a7565b9091508190611899565b90508301518e610a88565b608080517fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102905251818120927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016905b8181106119a15750908360019493921061196a575b5050811b017fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10255610ad9565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558e8061193e565b9192602060018192868a015181550194019201611929565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10260805152602060805120601f830160051c81019160208410611a1a575b601f0160051c01905b818110611a0d5750610a70565b6080518155600101611a00565b90915081906119f7565b680100000000000000017fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000008d16177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055610a17565b7ff92ee8a900000000000000000000000000000000000000000000000000000000608051526004608051fd5b9050158e6109bd565b303b1591506109b5565b60401c60ff161591508d6109aa565b5f80fd5b346104ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed5760043573ffffffffffffffffffffffffffffffffffffffff81168091036104ed57611b236178e4565b80156107c7577fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130555608051608051f35b346104ed5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed57611bcf615df7565b73ffffffffffffffffffffffffffffffffffffffff611bec615eb7565b91165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205273ffffffffffffffffffffffffffffffffffffffff60405f2091165f52602052602060ff60405f2054166040519015158152f35b346104ed576080517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed57602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b346104ed5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed57611cf1615df7565b611cf9615e94565b60643560048110156104ed57611d16916084359360243590616f49565b611d1e617c63565b8115611e5e577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526080519093916020918591606491839173ffffffffffffffffffffffffffffffffffffffff165af1928315611e5157602093611e26575b5081608051527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713158352604060805120611df0828254616edf565b9055604051908152817f615eb0eb079a1bab414b72e513f67a7286540e31499244785534e19b2c35c1eb843393a3604051908152f35b611e4590843d8611611e4a575b611e3d8183615f95565b8101906167e4565b611db6565b503d611e33565b6040513d608051823e3d90fd5b7f1f2a200500000000000000000000000000000000000000000000000000000000608051526004608051fd5b346104ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104ed57600435611ec46178e4565b8015611ef2577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713115560805180f35b7f247af9ce00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346104ed57611f2c36616178565b73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130554163303611f73576107c191617474565b7f4b60273500000000000000000000000000000000000000000000000000000000608051526004608051fd5b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757611fd9600435617924565b505f604051611fe9602082615f95565b525f608052612016604051611fff602082615f95565b5f8152604051918291602083526020830190615e51565b0390f35b34611ac75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757612051615df7565b5061205a615eb7565b5060643567ffffffffffffffff8111611ac75761207b9036906004016161ac565b507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004356004811015611ac7576120e7616f2b565b506120f181615e1a565b5f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60205261201660405f2060026040519161212d83615f4c565b80548352600181015460208401520154604082015260405191829182919091604080606083019480518452602081015160208501520151910152565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131154604051908152f35b34611ac75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576121fa615df7565b602435604435916004831015611ac757612212617c63565b61223f61221e83617924565b83339173ffffffffffffffffffffffffffffffffffffffff339116146166c9565b6122518361224c81617fc7565b616718565b61225a82617cb6565b6126c95761226783615e1a565b5f8381527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60209081526040808320600101547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015491517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810191909152928391606491839173ffffffffffffffffffffffffffffffffffffffff165af18015612623576126ac575b50815f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020524260405f2055815f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713096020526123828360405f2061680f565b73ffffffffffffffffffffffffffffffffffffffff811690816123d8575050604051916123ae81615e1a565b82527f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a3005b73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303949294541690604051907f1ffbb064000000000000000000000000000000000000000000000000000000008252856004830152602082602481865afa801561262357612461925f9161268d575b506168aa565b6040517f8da5cb5b000000000000000000000000000000000000000000000000000000008152602081600481885afa908115612623575f9161265e575b5073ffffffffffffffffffffffffffffffffffffffff3391160361262e576020602491604051928380927ffd66091e0000000000000000000000000000000000000000000000000000000082528860048301525afa908115612623575f916125f1575b50805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b60205260405f2054806125c257507f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c591816020925f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b83528460405f2055845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f835260405f2055604051906125bd81615e1a565b8152a3005b907f04e662d6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b90506020813d60201161261b575b8161260c60209383615f95565b81010312611ac7575184612501565b3d91506125ff565b6040513d5f823e3d90fd5b837f98f08d3a000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b612680915060203d602011612686575b6126788183615f95565b8101906168f4565b8561249e565b503d61266e565b6126a6915060203d602011611e4a57611e3d8183615f95565b8761245b565b6126c49060203d602011611e4a57611e3d8183615f95565b612320565b507f11d5a84c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757612016604051612734604082615f95565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190615e51565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576127a96178e4565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d55005b34611ac75760206127e86127e336616108565b6173cf565b604051908152f35b34611ac75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757612827615df7565b5060243580151503611ac7577f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d54604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671312602052602060405f2054604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154612979816161f0565b8084529060018116908115612a3257506001146129b5575b612016836129a181850382615f95565b604051918291602083526020830190615e51565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f9081527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b808210612a18575090915081016020016129a1612991565b919260018160209254838588010152019101909291612a00565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506129a19050612991565b34611ac75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757600435612aae615eb7565b90612ab7617c63565b612ae4612ac382617924565b82339173ffffffffffffffffffffffffffffffffffffffff339116146166c9565b73ffffffffffffffffffffffffffffffffffffffff8216918215612c8157815f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131660205260405f2054908115612c555790602081612bfb93855f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131683525f604081205573ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416905f6040518097819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115612623577f05443b5afb985a2c322e52876dd6461bc12f6aeac978d70dbb9d56f0d0bc470092602092612c3a575b50604051908152a3005b612c5090833d8511611e4a57611e3d8183615f95565b612c30565b827fa370a06b000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130854604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e54604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713065416604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757612ee4615df7565b612eec6178e4565b73ffffffffffffffffffffffffffffffffffffffff811615612c8157612f8b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713065416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130655565b005b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100541580613109575b156130ab5761304f612ff4616241565b612ffc616354565b602061305d6040519261300f8385615f95565b5f84525f3681376040519586957f0f00000000000000000000000000000000000000000000000000000000000000875260e08588015260e0870190615e51565b908582036040870152615e51565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b82811061309457505050500390f35b835185528695509381019392810192600101613085565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015415612fe4565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576131686178e4565b613170617c63565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760043561323f617c63565b801561336a577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671302546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905290602090829060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af180156126235761334d575b506132fc817feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131754616edf565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671317556040519081525f7f1012f0dfa5e4f0d93faaf4d2ff8e29506ed9a55ad1a06ead05e3e019bb3d607460203393a3005b6133659060203d602011611e4a57611e3d8183615f95565b6132d0565b7f1f2a2005000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004356133cc6178e4565b6101f481116133f9577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e55005b7f37829ee6000000000000000000000000000000000000000000000000000000005f526004526101f460245260445ffd5b34611ac75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757613461615df7565b613469615e94565b906064356004811015611ac7576020926127e89260243590616f49565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004356004811015611ac7576134c9616f2b565b506134d381615e1a565b80158181156134ef575b6134e691616718565b6120f181615e1a565b6134e691506134fd81617fc7565b91506134dd565b34611ac75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757600435602435906004821015611ac75761354b617c63565b613557612ac382617924565b61356081617cb6565b156137dc576135728261224c81617fc7565b805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260ff60405f205416916135ab83615e1a565b6135b481615e1a565b8281111561379b575f906135c784615e1a565b8382527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60205260206136d361363a600160408620015461360785615e1a565b8486527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a845260016040872001546167aa565b73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416906040519586809481937f23b872dd00000000000000000000000000000000000000000000000000000000835230336004850173ffffffffffffffffffffffffffffffffffffffff6040929594938160608401971683521660208201520152565b03925af1908115612623577f7aaafbadb3c6b6240d73322a3e69086965b7789111f0360a209bc84dca1104dd926137799261377e575b50835f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205261373e8160405f2061680f565b835f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020524260405f2055604051918291339683616754565b0390a3005b6137969060203d602011611e4a57611e3d8183615f95565b613709565b827fa0f244d1000000000000000000000000000000000000000000000000000000005f526137c881615e1a565b6004526137d481615e1a565b60245260445ffd5b7fed5a980a000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576138ae615df7565b6138b66178e4565b73ffffffffffffffffffffffffffffffffffffffff811615612c8157612f8b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130355565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416036139ca57612f8b33617eb8565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713075416604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757613a9d615df7565b613aa56178e4565b73ffffffffffffffffffffffffffffffffffffffff811615612c8157612f8b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130155565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757613b7b615df7565b613b836178e4565b73ffffffffffffffffffffffffffffffffffffffff811615612c8157612f8b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713025416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130255565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757613c586178e4565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00555f73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757613d99615df7565b73ffffffffffffffffffffffffffffffffffffffff811615613e0757613dfe60209173ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b54604051908152f35b7f89c62b64000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131754604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130054604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757613f1e6178e4565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131455005b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576020613f81600435617924565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713045416604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671315602052602060405f2054604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576140ad615df7565b6140b5617e64565b6140bd6178e4565b73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671302541690604051907f70a08231000000000000000000000000000000000000000000000000000000008252306004830152602082602481865afa918215612623575f926141bc575b506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024810191909152906020908290815f81604481015b03925af18015612623576141a457005b612f8b9060203d602011611e4a57611e3d8183615f95565b91506020823d6020116141ea575b816141d760209383615f95565b81010312611ac75790519061419461413a565b3d91506141ca565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036142e65760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611ac757612f8b61431f36616178565b90614328617c63565b614334612ac382617924565b617d7e565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416604051908152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576143db615df7565b60243567ffffffffffffffff8111611ac7576143fb9036906004016161ac565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115614666575b506142e65761444a6178e4565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181614632575b506144ca57837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8592036146075750813b156145dc57807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28151156145ab575f80836020612f8b95519101845af46145a5618487565b916184b6565b5050346145b457005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d60201161465e575b8161464e60209383615f95565b81010312611ac757519085614499565b3d9150614641565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc541614158361443d565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760043567ffffffffffffffff8111611ac75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112611ac757613f81602091600401616eec565b34611ac75761473136616178565b9061473a617c63565b811561336a577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905290602090829060649082905f9073ffffffffffffffffffffffffffffffffffffffff165af1801561262357614832575b50805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131560205260405f20614802838254616edf565b90556040519182527f615eb0eb079a1bab414b72e513f67a7286540e31499244785534e19b2c35c1eb60203393a3005b61484a9060203d602011611e4a57611e3d8183615f95565b6147cb565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f602052602060405f2054604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671313602052602060ff60405f2054166040519015158152f35b34611ac75761493036615eda565b5050505f604051614942602082615f95565b527f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611ac75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760206127e86149a7615df7565b60243590616ccb565b34611ac75760206127e86149c336616108565b616b84565b34611ac75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576024357fffffffff0000000000000000000000000000000000000000000000000000000081168103611ac7576044359067ffffffffffffffff8211611ac75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8336030112611ac757612f8b916004019060043561696e565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757614aa86178e4565b614ab0617e64565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131454604051908152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671316602052602060405f2054604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760206127e8616010565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757614c72615df7565b614c7a617c63565b73ffffffffffffffffffffffffffffffffffffffff8116908115612c8157602490602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713035416604051938480927f1ffbb0640000000000000000000000000000000000000000000000000000000082528760048301525afa801561262357614d1d925f916150a557506168aa565b604051907f8da5cb5b000000000000000000000000000000000000000000000000000000008252602082600481845afa918215612623575f92615084575b5073ffffffffffffffffffffffffffffffffffffffff8216908115612c81576020600491604051928380927faf640d0f0000000000000000000000000000000000000000000000000000000082525afa908115612623575f91615052575b505f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b60205260405f20549173ffffffffffffffffffffffffffffffffffffffff614e0484617924565b169082821461502a5760209160405191614e1e8484615f95565b5f835273ffffffffffffffffffffffffffffffffffffffff614e4087836180c3565b1680614e7257867f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b908183889303614ff757503b614e8457005b8391614ed560405194859384937f150b7a0200000000000000000000000000000000000000000000000000000000855233600486015260248501526044840152608060648401526084830190615e51565b03815f875af15f9181614f9f575b50614f285750614ef1618487565b80519182614f2557837f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b01fd5b7f150b7a020000000000000000000000000000000000000000000000000000000091507fffffffff000000000000000000000000000000000000000000000000000000001603614f7457005b7f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091508281813d8311614ff0575b614fb78183615f95565b81010312611ac757517fffffffff0000000000000000000000000000000000000000000000000000000081168103611ac7579084614ee3565b503d614fad565b82847f64283d7b000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b7f2a63c7cc000000000000000000000000000000000000000000000000000000005f5260045ffd5b90506020813d60201161507c575b8161506d60209383615f95565b81010312611ac7575183614db9565b3d9150615060565b61509e91925060203d602011612686576126788183615f95565b9082614d5b565b6150be915060203d602011611e4a57611e3d8183615f95565b8461245b565b34611ac7576150d236615eda565b5050507f8cd22d19000000000000000000000000000000000000000000000000000000005f5260045ffd5b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576020613f81600435616846565b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760207feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c54604051908152f35b34611ac75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004356024356004811015611ac7576044359267ffffffffffffffff8411611ac75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc856004019536030112611ac75761521c617c63565b61524961522884617924565b84339173ffffffffffffffffffffffffffffffffffffffff339116146166c9565b61525283617cb6565b156158165761526082615e1a565b8115828115615801575b61527391616718565b825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260ff60405f205416936152ac85615e1a565b6152b583615e1a565b848310156157d357835f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f60205260405f205490816156f3575b5050506152fc83615e1a565b825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60205261536c600160405f20015461533783615e1a565b825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a602052600160405f200154906167aa565b90825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020526153a260405f2054426167aa565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c541115615622576154d260206154086127106154007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130d54876167fc565b0480956167aa565b9373ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301541673ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130454165f6040518096819582947fa9059cbb000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af18015612623575f9360209261557d92615607575b5073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015416906040519586809481937fa9059cbb00000000000000000000000000000000000000000000000000000000835233600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115612623577f65e7b85a0b1fe94cdd38e2a9dbe4ac14dda1ceeb998856160d0ea99b804b105d92613779926155e857505b835f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205261373e8160405f2061680f565b6156009060203d602011611e4a57611e3d8183615f95565b5085613709565b61561d90843d8611611e4a57611e3d8183615f95565b6154ea565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081523360048201526024810193909352602090839060449082905f9073ffffffffffffffffffffffffffffffffffffffff165af1908115612623577f65e7b85a0b1fe94cdd38e2a9dbe4ac14dda1ceeb998856160d0ea99b804b105d92613779926156d4575b506155b4565b6156ec9060203d602011611e4a57611e3d8183615f95565b50856156ce565b6157726127e39261572a837fffffffff0000000000000000000000000000000000000000000000000000000061577897168361696e565b61573386615e1a565b855f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a602052600260405f20015461576c8480616777565b91616513565b80616777565b5f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131360205260405f2060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790558380806152f0565b82857f253471cf000000000000000000000000000000000000000000000000000000005f526137c881615e1a565b615273915061580f81617fc7565b915061526a565b827fed5a980a000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b602052602060405f2054604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416604051908152f35b34611ac75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75760443567ffffffffffffffff8111611ac7576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112611ac757612f8b90600401602435600435616513565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576159d06178e4565b6004357feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130c55005b34611ac75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac75761207b615df7565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757600435615a6981617924565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079304602052602073ffffffffffffffffffffffffffffffffffffffff60405f205416604051908152f35b34611ac7575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054615b12816161f0565b8084529060018116908115612a325750600114615b3957612016836129a181850382615f95565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b808210615b9c575090915081016020016129a1612991565b919260018160209254838588010152019101909291615b84565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004355f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671309602052602060ff60405f20541660405190615c2581615e1a565b8152f35b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac7576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103611ac757807f80ac58cd0000000000000000000000000000000000000000000000000000000060209214908115615cef575b8115615cc5575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501482615cba565b7f5b5e139f0000000000000000000000000000000000000000000000000000000081149150615cb3565b34611ac75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611ac757615d50615df7565b615d586178e4565b73ffffffffffffffffffffffffffffffffffffffff811615612c8157612f8b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713075416177feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130755565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203611ac757565b60041115615e2457565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6044359073ffffffffffffffffffffffffffffffffffffffff82168203611ac757565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203611ac757565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6060910112611ac75760043573ffffffffffffffffffffffffffffffffffffffff81168103611ac7579060243573ffffffffffffffffffffffffffffffffffffffff81168103611ac7579060443590565b6060810190811067ffffffffffffffff821117615f6857604052565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117615f6857604052565b67ffffffffffffffff8111615f6857601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60405161601e60c082615f95565b609681527f3634207461726765742c627974657320636c61696d290000000000000000000060a060208301927f56657269666961626c6543726564656e7469616c28616464726573732069737384527f7565722c75696e74323536207375626a6563742c75696e743235362065706f6360408201527f684973737565642c75696e743235362065706f636856616c6964556e74696c2c60608201527f75696e743235362076616c75652c62797465733420616374696f6e2c75696e746080820152015260405160966020820192835e5f60b68201526096815261610260b682615f95565b51902090565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112611ac7576004359067ffffffffffffffff8211611ac7577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8261010092030112611ac75760040190565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc6040910112611ac7576004359060243590565b81601f82011215611ac7576020813591016161c682615fd6565b926161d46040519485615f95565b82845282820111611ac757815f92602092838601378301015290565b90600182811c92168015616237575b602083101461620a57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f16916161ff565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025491616273836161f0565b80835292600181169081156163175750600114616299575b61629792500383615f95565b565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1025f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b8183106162fb5750509060206162979282010161628b565b60209193508060019154838589010152019101909184926162e3565b602092506162979491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b82010161628b565b604051905f827fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035491616386836161f0565b808352926001811690811561631757506001146163a95761629792500383615f95565b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1035f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b81831061640b5750509060206162979282010161628b565b60209193508060019154838589010152019101909184926163f3565b919061010083820312611ac75760405190610100820182811067ffffffffffffffff821117615f68576040528193803573ffffffffffffffffffffffffffffffffffffffff81168103611ac75783526020810135602084015260408101356040840152606081013560608401526080810135608084015260a08101357fffffffff0000000000000000000000000000000000000000000000000000000081168103611ac75760a084015260c081013567ffffffffffffffff81168103611ac75760c084015260e08101359167ffffffffffffffff8311611ac75760e09261650e92016161ac565b910152565b90929173ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671307541691602073ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713055416916024604051809481937f9f4326d700000000000000000000000000000000000000000000000000000000835260048301525afa908115612623575f91616697575b506165e0836165da3685616427565b836179b8565b9590818111616655575050506165f7929350617aa9565b917feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671314549261662457505050565b7f4c1b3697000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b6084935086604051937f4ec0c37d0000000000000000000000000000000000000000000000000000000085526004850152602484015260448301526064820152fd5b90506020813d6020116166c1575b816166b260209383615f95565b81010312611ac757515f6165cb565b3d91506166a5565b156166d2575050565b73ffffffffffffffffffffffffffffffffffffffff92507f3a0302a9000000000000000000000000000000000000000000000000000000005f526004521660245260445ffd5b156167205750565b7fbca1a956000000000000000000000000000000000000000000000000000000005f5261674c81615e1a565b60045260245ffd5b604081019392916020919061676881615e1a565b815261677383615e1a565b0152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215611ac7570190565b919082039182116167b757565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b90816020910312611ac757518015158103611ac75790565b818102929181159184041417156167b757565b9061681981615e1a565b60ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b8015616882575f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b60205261687f60405f2054617924565b90565b7fac8fb3c1000000000000000000000000000000000000000000000000000000005f5260045ffd5b156168b25750565b73ffffffffffffffffffffffffffffffffffffffff907f3a35a1b9000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b90816020910312611ac7575173ffffffffffffffffffffffffffffffffffffffff81168103611ac75790565b3573ffffffffffffffffffffffffffffffffffffffff81168103611ac75790565b357fffffffff0000000000000000000000000000000000000000000000000000000081168103611ac75790565b60ff929161697b83616eec565b9061698583616846565b9260206169928680616777565b013514936169a36127e38280616777565b5f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713136020528560405f20541615159373ffffffffffffffffffffffffffffffffffffffff806169fb6169f68580616777565b616920565b1694169384149384616b40575b5073ffffffffffffffffffffffffffffffffffffffff163381149081616b36575b506040616a368380616777565b01354210159182616b11575b60a0616a5182616a5793616777565b01616941565b935f9615616b08575b15616afe575b15616af4575b15616ae9575b7fffffffff00000000000000000000000000000000000000000000000000000000809116911603616adf575b616ad7575b1680616aac5750565b7f591b94f2000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b602017616aa3565b9060101790616a9e565b600890931792616a72565b9360041793616a6c565b9460021794616a66565b60019650616a60565b9150616a5760a0616a516060616b278680616777565b01354211159492505050616a42565b905015155f616a29565b73ffffffffffffffffffffffffffffffffffffffff809295507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713065416149390616a08565b616b8c616010565b90616b9681616920565b90616ba360a08201616941565b60e08201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe183360301811215611ac75782019081359167ffffffffffffffff8311611ac757602001918036038313611ac7577fffffffff0000000000000000000000000000000000000000000000000000000092608091616c60602060405183819483830196873781015f8382015203017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282615f95565b5190209373ffffffffffffffffffffffffffffffffffffffff604051966020880198895216604087015260208101356060870152604081013582870152606081013560a0870152013560c08501521660e0830152610100820152610100815261610261012082615f95565b90616cd4617c63565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671304547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671300546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9283166024820152604481019190915291602091839160649183915f91165af1801561262357616ec2575b507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130854917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83146167b757616e2090600184017feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130855617d1a565b73ffffffffffffffffffffffffffffffffffffffff811615616e9657616e5b8373ffffffffffffffffffffffffffffffffffffffff926180c3565b16616e6a5761687f9082617d7e565b7f73c6ac6e000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b7f64a0ae92000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b616eda9060203d602011611e4a57611e3d8183615f95565b616da6565b919082018092116167b757565b616ef96127e38280616777565b9060208101359060ff82168203611ac75761687f92616f229260406060840135930135916182da565b90929192618369565b60405190616f3882615f4c565b5f6040838281528260208201520152565b90616f5391616ccb565b91616f5c617c63565b616f6861522884617924565b616f758161224c81617fc7565b616f7e83617cb6565b6173a357616f8b81615e1a565b5f8181527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60209081526040808320600101547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713015491517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810191909152928391606491839173ffffffffffffffffffffffffffffffffffffffff165af1801561262357617386575b50825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713106020524260405f2055825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713096020526170a68160405f2061680f565b73ffffffffffffffffffffffffffffffffffffffff821691826170fe575060405191506170d281615e1a565b8152817f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c560203393a390565b73ffffffffffffffffffffffffffffffffffffffff7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671303541690604051907f1ffbb064000000000000000000000000000000000000000000000000000000008252846004830152602082602481865afa801561262357617183925f9161736757506168aa565b6040517f8da5cb5b000000000000000000000000000000000000000000000000000000008152602081600481875afa908115612623575f91617348575b5073ffffffffffffffffffffffffffffffffffffffff33911603617318576020602491604051928380927ffd66091e0000000000000000000000000000000000000000000000000000000082528760048301525afa908115612623575f916172e6575b50805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b60205260405f2054806125c2575060208492827f62133dedf0f5ee80c5427304aa3d2cd79fdda28731dbe60859a2d7072ea9c7c5935f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b83528460405f2055845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130f835260405f2055604051906172e081615e1a565b8152a390565b90506020813d602011617310575b8161730160209383615f95565b81010312611ac757515f617223565b3d91506172f4565b827f98f08d3a000000000000000000000000000000000000000000000000000000005f526004523360245260445ffd5b617361915060203d602011612686576126788183615f95565b5f6171c0565b617380915060203d602011611e4a57611e3d8183615f95565b5f61245b565b61739e9060203d602011611e4a57611e3d8183615f95565b617044565b827f11d5a84c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b6173da604291616b84565b6173e261854f565b6173ea6185b9565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261743b60c082615f95565b51902090604051917f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201522090565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166178e05781156178e057805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130b60205260405f20549182156178ae576174df83617cb6565b156178ae57825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131260205260405f2054156178ae57825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131260205260405f205480156178b3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0481116178ae57825f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260ff60405f2054166175a781615e1a565b5f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130a60205260405f206002604051916175e083615f4c565b805483526001810154602084015201546040820152835f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131260205261271061762d60405f2054846167fc565b0491845f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131560205260405f20549182156178a6575180670de0b6b3a76400008502049283811061788f575b507feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713175490848210617877575b50508115617870577feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671301547feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671304546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810185905291602091839160449183915f91165af19081617853575b5061775e575050505050565b7f299d374ad79b9f2ede438f8b72a6c71310be84f532c6b95eed6b61784724b9d692606092865f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131560205260405f206177b98282546167aa565b90556177e6827feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd2671317546167aa565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131755865f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131660205260405f2061783e838254616edf565b905560405192835260208301526040820152a3565b61786b9060203d602011611e4a57611e3d8183615f95565b617752565b5050505050565b909350670de0b6b3a764000084020491505f806176a5565b670de0b6b3a764000082820204945092505f617679565b505050505050565b505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5050565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633036139ca57565b61796c815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff82161561798d575090565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b90617a1673ffffffffffffffffffffffffffffffffffffffff949360e060209301516040519687809481937f1b2b5fad0000000000000000000000000000000000000000000000000000000083528760048401526024830190615e51565b0392165afa928315612623575f93617a75575b50828115617a6d5750828015617a4657617a429161800c565b9190565b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9190565b9250505f9190565b9092506020813d602011617aa1575b81617a9160209383615f95565b81010312611ac75751915f617a29565b3d9150617a84565b617b109291602060e073ffffffffffffffffffffffffffffffffffffffff617ad13686616427565b9316920151604051809681927f07124c060000000000000000000000000000000000000000000000000000000083528460048401526024830190615e51565b0381845afa938415612623575f94617c2e575b506020617b769160e0617b3887953690616427565b01519060405180809581947f402ed8cb0000000000000000000000000000000000000000000000000000000083528660048401526024830190615e51565b03915afa908115612623575f91617bfc575b508093811580617bf4575b617bec575050821580617be3575b617bdc57617baf838261800c565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713145410617bdc5791905f90565b9190600190565b50801515617ba1565b9350915f9150565b508015617b93565b90506020813d602011617c26575b81617c1760209383615f95565b81010312611ac757515f617b88565b3d9150617c0a565b9093506020813d602011617c5b575b81617c4a60209383615f95565b81010312611ac75751926020617b23565b3d9150617c3d565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416617c8e57565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130960205260ff60405f205416617ced81615e1a565b80159081159182617cfd57505090565b617d0682615e1a565b91617d10575b5090565b61687f9150617fc7565b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014617d6a575b15617d0c57617d5c9061804a565b90617d65575090565b905090565b505067ffffffffffffffff81166001617d4e565b7feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267130e54808311617e355750805f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd267131260205260405f205491815f527feb5f85dd766f96b2f6b2179489d6c8d314c6878705da165eaca0af3bd26713126020528060405f205560405192835260208301527fdefeab33850afef6813a2e79837ea2d353a9bd7a416d49cc188966c3d1cd90ca60403393a3565b827f37829ee6000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300541615617e9057565b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b617fd081615e1a565b60018114908115617ff7575b8115617fe6575090565b60039150617ff381615e1a565b1490565b905061800281615e1a565b6002811490617fdc565b7812725dd1d243aba0e75fe645cc4873f9e65afe688c928e1f22811082021561803d57670de0b6b3a7640000020490565b637c5f487d5f526004601cfd5b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff16036180b8575b821580156180ad575b6180a557565b5f9250829150565b5060163d141561809f565b5f9250829150618096565b9061810c815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9173ffffffffffffffffffffffffffffffffffffffff831680618217575b73ffffffffffffffffffffffffffffffffffffffff821691826181c3575b50825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f20827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a490565b61820a9073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190555f618148565b825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff000000000000000000000000000000000000000081541690556182af8473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff815401905561812a565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161835e579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612623575f5173ffffffffffffffffffffffffffffffffffffffff81161561835457905f905f90565b505f906001905f90565b5050505f9160039190565b61837281615e1a565b8061837b575050565b61838481615e1a565b600181036183b4577ff645eedf000000000000000000000000000000000000000000000000000000005f5260045ffd5b6183bd81615e1a565b600281036183f157507ffce698f7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b6003906183fd81615e1a565b146184055750565b7fd78bce0c000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561845f57565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b3d156184b1573d9061849882615fd6565b916184a66040519384615f95565b82523d5f602084013e565b606090565b906184f357508051156184cb57805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580618546575b618504575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b156184fc565b618557616241565b8051908115618567576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005480156185945790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b6185c1616354565b80519081156185d1576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101548015618594579056fea2646970667358221220cc66b494a088ac639a83d0d451b62f6a0c30d1662673af0bb10561779ab9358564736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_Plus *PlusCaller) VERIFIABLECREDENTIALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "_VERIFIABLE_CREDENTIAL_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_Plus *PlusSession) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _Plus.Contract.VERIFIABLECREDENTIALTYPEHASH(&_Plus.CallOpts)
}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_Plus *PlusCallerSession) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _Plus.Contract.VERIFIABLECREDENTIALTYPEHASH(&_Plus.CallOpts)
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

// BaseConversionRate is a free data retrieval call binding the contract method 0xb2d45638.
//
// Solidity: function baseConversionRate() view returns(uint256)
func (_Plus *PlusCaller) BaseConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "baseConversionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseConversionRate is a free data retrieval call binding the contract method 0xb2d45638.
//
// Solidity: function baseConversionRate() view returns(uint256)
func (_Plus *PlusSession) BaseConversionRate() (*big.Int, error) {
	return _Plus.Contract.BaseConversionRate(&_Plus.CallOpts)
}

// BaseConversionRate is a free data retrieval call binding the contract method 0xb2d45638.
//
// Solidity: function baseConversionRate() view returns(uint256)
func (_Plus *PlusCallerSession) BaseConversionRate() (*big.Int, error) {
	return _Plus.Contract.BaseConversionRate(&_Plus.CallOpts)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_Plus *PlusCaller) CredParser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "credParser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_Plus *PlusSession) CredParser() (common.Address, error) {
	return _Plus.Contract.CredParser(&_Plus.CallOpts)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_Plus *PlusCallerSession) CredParser() (common.Address, error) {
	return _Plus.Contract.CredParser(&_Plus.CallOpts)
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

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_Plus *PlusCaller) DeriveStructHash(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "deriveStructHash", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_Plus *PlusSession) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _Plus.Contract.DeriveStructHash(&_Plus.CallOpts, vc)
}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_Plus *PlusCallerSession) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _Plus.Contract.DeriveStructHash(&_Plus.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_Plus *PlusCaller) Digest(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "digest", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_Plus *PlusSession) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _Plus.Contract.Digest(&_Plus.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_Plus *PlusCallerSession) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _Plus.Contract.Digest(&_Plus.CallOpts, vc)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Plus *PlusCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Plus *PlusSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Plus.Contract.Eip712Domain(&_Plus.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Plus *PlusCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Plus.Contract.Eip712Domain(&_Plus.CallOpts)
}

// FilToken is a free data retrieval call binding the contract method 0x8e18d648.
//
// Solidity: function filToken() view returns(address)
func (_Plus *PlusCaller) FilToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "filToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FilToken is a free data retrieval call binding the contract method 0x8e18d648.
//
// Solidity: function filToken() view returns(address)
func (_Plus *PlusSession) FilToken() (common.Address, error) {
	return _Plus.Contract.FilToken(&_Plus.CallOpts)
}

// FilToken is a free data retrieval call binding the contract method 0x8e18d648.
//
// Solidity: function filToken() view returns(address)
func (_Plus *PlusCallerSession) FilToken() (common.Address, error) {
	return _Plus.Contract.FilToken(&_Plus.CallOpts)
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

// GetTierInfo is a free data retrieval call binding the contract method 0x7ed85ff3.
//
// Solidity: function getTierInfo(uint8 _tier) view returns((uint256,uint256,uint256) tierInfo)
func (_Plus *PlusCaller) GetTierInfo(opts *bind.CallOpts, _tier uint8) (TierInfo, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "getTierInfo", _tier)

	if err != nil {
		return *new(TierInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TierInfo)).(*TierInfo)

	return out0, err

}

// GetTierInfo is a free data retrieval call binding the contract method 0x7ed85ff3.
//
// Solidity: function getTierInfo(uint8 _tier) view returns((uint256,uint256,uint256) tierInfo)
func (_Plus *PlusSession) GetTierInfo(_tier uint8) (TierInfo, error) {
	return _Plus.Contract.GetTierInfo(&_Plus.CallOpts, _tier)
}

// GetTierInfo is a free data retrieval call binding the contract method 0x7ed85ff3.
//
// Solidity: function getTierInfo(uint8 _tier) view returns((uint256,uint256,uint256) tierInfo)
func (_Plus *PlusCallerSession) GetTierInfo(_tier uint8) (TierInfo, error) {
	return _Plus.Contract.GetTierInfo(&_Plus.CallOpts, _tier)
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

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_Plus *PlusCaller) Recover(opts *bind.CallOpts, sc SignedCredential) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "recover", sc)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_Plus *PlusSession) Recover(sc SignedCredential) (common.Address, error) {
	return _Plus.Contract.Recover(&_Plus.CallOpts, sc)
}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_Plus *PlusCallerSession) Recover(sc SignedCredential) (common.Address, error) {
	return _Plus.Contract.Recover(&_Plus.CallOpts, sc)
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

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_Plus *PlusCaller) SectorFaultyTolerancePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "sectorFaultyTolerancePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_Plus *PlusSession) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _Plus.Contract.SectorFaultyTolerancePercent(&_Plus.CallOpts)
}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_Plus *PlusCallerSession) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _Plus.Contract.SectorFaultyTolerancePercent(&_Plus.CallOpts)
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

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agentId, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_Plus *PlusCaller) ValidateCred(opts *bind.CallOpts, agentId *big.Int, selector [4]byte, sc SignedCredential) error {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "validateCred", agentId, selector, sc)

	if err != nil {
		return err
	}

	return err

}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agentId, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_Plus *PlusSession) ValidateCred(agentId *big.Int, selector [4]byte, sc SignedCredential) error {
	return _Plus.Contract.ValidateCred(&_Plus.CallOpts, agentId, selector, sc)
}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agentId, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_Plus *PlusCallerSession) ValidateCred(agentId *big.Int, selector [4]byte, sc SignedCredential) error {
	return _Plus.Contract.ValidateCred(&_Plus.CallOpts, agentId, selector, sc)
}

// VcIssuer is a free data retrieval call binding the contract method 0x8776d0d6.
//
// Solidity: function vcIssuer() view returns(address)
func (_Plus *PlusCaller) VcIssuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plus.contract.Call(opts, &out, "vcIssuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VcIssuer is a free data retrieval call binding the contract method 0x8776d0d6.
//
// Solidity: function vcIssuer() view returns(address)
func (_Plus *PlusSession) VcIssuer() (common.Address, error) {
	return _Plus.Contract.VcIssuer(&_Plus.CallOpts)
}

// VcIssuer is a free data retrieval call binding the contract method 0x8776d0d6.
//
// Solidity: function vcIssuer() view returns(address)
func (_Plus *PlusCallerSession) VcIssuer() (common.Address, error) {
	return _Plus.Contract.VcIssuer(&_Plus.CallOpts)
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

// FundFilVault is a paid mutator transaction binding the contract method 0x83596c20.
//
// Solidity: function fundFilVault(uint256 _amount) returns()
func (_Plus *PlusTransactor) FundFilVault(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "fundFilVault", _amount)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x83596c20.
//
// Solidity: function fundFilVault(uint256 _amount) returns()
func (_Plus *PlusSession) FundFilVault(_amount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.FundFilVault(&_Plus.TransactOpts, _amount)
}

// FundFilVault is a paid mutator transaction binding the contract method 0x83596c20.
//
// Solidity: function fundFilVault(uint256 _amount) returns()
func (_Plus *PlusTransactorSession) FundFilVault(_amount *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.FundFilVault(&_Plus.TransactOpts, _amount)
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

// Initialize is a paid mutator transaction binding the contract method 0xea982a9d.
//
// Solidity: function initialize(string _name, string _symbol, string _version, address _vcIssuer, address _initialOwner, uint256 _mintPrice, address _glfToken, address _filToken, address _treasury, address _agentFactory, address _credParser) returns()
func (_Plus *PlusTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _version string, _vcIssuer common.Address, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _filToken common.Address, _treasury common.Address, _agentFactory common.Address, _credParser common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "initialize", _name, _symbol, _version, _vcIssuer, _initialOwner, _mintPrice, _glfToken, _filToken, _treasury, _agentFactory, _credParser)
}

// Initialize is a paid mutator transaction binding the contract method 0xea982a9d.
//
// Solidity: function initialize(string _name, string _symbol, string _version, address _vcIssuer, address _initialOwner, uint256 _mintPrice, address _glfToken, address _filToken, address _treasury, address _agentFactory, address _credParser) returns()
func (_Plus *PlusSession) Initialize(_name string, _symbol string, _version string, _vcIssuer common.Address, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _filToken common.Address, _treasury common.Address, _agentFactory common.Address, _credParser common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Initialize(&_Plus.TransactOpts, _name, _symbol, _version, _vcIssuer, _initialOwner, _mintPrice, _glfToken, _filToken, _treasury, _agentFactory, _credParser)
}

// Initialize is a paid mutator transaction binding the contract method 0xea982a9d.
//
// Solidity: function initialize(string _name, string _symbol, string _version, address _vcIssuer, address _initialOwner, uint256 _mintPrice, address _glfToken, address _filToken, address _treasury, address _agentFactory, address _credParser) returns()
func (_Plus *PlusTransactorSession) Initialize(_name string, _symbol string, _version string, _vcIssuer common.Address, _initialOwner common.Address, _mintPrice *big.Int, _glfToken common.Address, _filToken common.Address, _treasury common.Address, _agentFactory common.Address, _credParser common.Address) (*types.Transaction, error) {
	return _Plus.Contract.Initialize(&_Plus.TransactOpts, _name, _symbol, _version, _vcIssuer, _initialOwner, _mintPrice, _glfToken, _filToken, _treasury, _agentFactory, _credParser)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _personalCashBackPercent) returns(uint256 tokenId)
func (_Plus *PlusTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _personalCashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "mint", _to, _personalCashBackPercent)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _personalCashBackPercent) returns(uint256 tokenId)
func (_Plus *PlusSession) Mint(_to common.Address, _personalCashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.Mint(&_Plus.TransactOpts, _to, _personalCashBackPercent)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _personalCashBackPercent) returns(uint256 tokenId)
func (_Plus *PlusTransactorSession) Mint(_to common.Address, _personalCashBackPercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.Mint(&_Plus.TransactOpts, _to, _personalCashBackPercent)
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

// MintAndActivate is a paid mutator transaction binding the contract method 0x800ca2ae.
//
// Solidity: function mintAndActivate(address _to, uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_Plus *PlusTransactor) MintAndActivate(opts *bind.TransactOpts, _to common.Address, _personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "mintAndActivate", _to, _personalCashBackPercent, _beneficiary, _tier)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0x800ca2ae.
//
// Solidity: function mintAndActivate(address _to, uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_Plus *PlusSession) MintAndActivate(_to common.Address, _personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _Plus.Contract.MintAndActivate(&_Plus.TransactOpts, _to, _personalCashBackPercent, _beneficiary, _tier)
}

// MintAndActivate is a paid mutator transaction binding the contract method 0x800ca2ae.
//
// Solidity: function mintAndActivate(address _to, uint256 _personalCashBackPercent, address _beneficiary, uint8 _tier) returns(uint256 tokenId)
func (_Plus *PlusTransactorSession) MintAndActivate(_to common.Address, _personalCashBackPercent *big.Int, _beneficiary common.Address, _tier uint8) (*types.Transaction, error) {
	return _Plus.Contract.MintAndActivate(&_Plus.TransactOpts, _to, _personalCashBackPercent, _beneficiary, _tier)
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

// SetBaseConversionRate is a paid mutator transaction binding the contract method 0xd72422d8.
//
// Solidity: function setBaseConversionRate(uint256 _baseConversionRate) returns()
func (_Plus *PlusTransactor) SetBaseConversionRate(opts *bind.TransactOpts, _baseConversionRate *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setBaseConversionRate", _baseConversionRate)
}

// SetBaseConversionRate is a paid mutator transaction binding the contract method 0xd72422d8.
//
// Solidity: function setBaseConversionRate(uint256 _baseConversionRate) returns()
func (_Plus *PlusSession) SetBaseConversionRate(_baseConversionRate *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetBaseConversionRate(&_Plus.TransactOpts, _baseConversionRate)
}

// SetBaseConversionRate is a paid mutator transaction binding the contract method 0xd72422d8.
//
// Solidity: function setBaseConversionRate(uint256 _baseConversionRate) returns()
func (_Plus *PlusTransactorSession) SetBaseConversionRate(_baseConversionRate *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetBaseConversionRate(&_Plus.TransactOpts, _baseConversionRate)
}

// SetCredParser is a paid mutator transaction binding the contract method 0x01741744.
//
// Solidity: function setCredParser(address _credParser) returns()
func (_Plus *PlusTransactor) SetCredParser(opts *bind.TransactOpts, _credParser common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setCredParser", _credParser)
}

// SetCredParser is a paid mutator transaction binding the contract method 0x01741744.
//
// Solidity: function setCredParser(address _credParser) returns()
func (_Plus *PlusSession) SetCredParser(_credParser common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetCredParser(&_Plus.TransactOpts, _credParser)
}

// SetCredParser is a paid mutator transaction binding the contract method 0x01741744.
//
// Solidity: function setCredParser(address _credParser) returns()
func (_Plus *PlusTransactorSession) SetCredParser(_credParser common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetCredParser(&_Plus.TransactOpts, _credParser)
}

// SetFilToken is a paid mutator transaction binding the contract method 0x73da6924.
//
// Solidity: function setFilToken(address _filToken) returns()
func (_Plus *PlusTransactor) SetFilToken(opts *bind.TransactOpts, _filToken common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setFilToken", _filToken)
}

// SetFilToken is a paid mutator transaction binding the contract method 0x73da6924.
//
// Solidity: function setFilToken(address _filToken) returns()
func (_Plus *PlusSession) SetFilToken(_filToken common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetFilToken(&_Plus.TransactOpts, _filToken)
}

// SetFilToken is a paid mutator transaction binding the contract method 0x73da6924.
//
// Solidity: function setFilToken(address _filToken) returns()
func (_Plus *PlusTransactorSession) SetFilToken(_filToken common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetFilToken(&_Plus.TransactOpts, _filToken)
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

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_Plus *PlusTransactor) SetSectorFaultyTolerancePercent(opts *bind.TransactOpts, _sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setSectorFaultyTolerancePercent", _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_Plus *PlusSession) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetSectorFaultyTolerancePercent(&_Plus.TransactOpts, _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_Plus *PlusTransactorSession) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _Plus.Contract.SetSectorFaultyTolerancePercent(&_Plus.TransactOpts, _sectorFaultyTolerancePercent)
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

// SetVcIssuer is a paid mutator transaction binding the contract method 0x863de770.
//
// Solidity: function setVcIssuer(address _vcIssuer) returns()
func (_Plus *PlusTransactor) SetVcIssuer(opts *bind.TransactOpts, _vcIssuer common.Address) (*types.Transaction, error) {
	return _Plus.contract.Transact(opts, "setVcIssuer", _vcIssuer)
}

// SetVcIssuer is a paid mutator transaction binding the contract method 0x863de770.
//
// Solidity: function setVcIssuer(address _vcIssuer) returns()
func (_Plus *PlusSession) SetVcIssuer(_vcIssuer common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetVcIssuer(&_Plus.TransactOpts, _vcIssuer)
}

// SetVcIssuer is a paid mutator transaction binding the contract method 0x863de770.
//
// Solidity: function setVcIssuer(address _vcIssuer) returns()
func (_Plus *PlusTransactorSession) SetVcIssuer(_vcIssuer common.Address) (*types.Transaction, error) {
	return _Plus.Contract.SetVcIssuer(&_Plus.TransactOpts, _vcIssuer)
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

// PlusEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Plus contract.
type PlusEIP712DomainChangedIterator struct {
	Event *PlusEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *PlusEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlusEIP712DomainChanged)
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
		it.Event = new(PlusEIP712DomainChanged)
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
func (it *PlusEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlusEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlusEIP712DomainChanged represents a EIP712DomainChanged event raised by the Plus contract.
type PlusEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Plus *PlusFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*PlusEIP712DomainChangedIterator, error) {

	logs, sub, err := _Plus.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &PlusEIP712DomainChangedIterator{contract: _Plus.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Plus *PlusFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *PlusEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Plus.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlusEIP712DomainChanged)
				if err := _Plus.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Plus *PlusFilterer) ParseEIP712DomainChanged(log types.Log) (*PlusEIP712DomainChanged, error) {
	event := new(PlusEIP712DomainChanged)
	if err := _Plus.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
