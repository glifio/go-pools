package types

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	filtypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/vc"
)

//go:generate mockery --name FEVMQueries
type FEVMQueries interface {
	// agent methods
	AgentID(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	AgentAccount(ctx context.Context, agentAddr common.Address, poolID *big.Int, blockNumber *big.Int) (abigen.Account, error)
	AgentAddrIDFromRcpt(ctx context.Context, rcpt *types.Receipt) (common.Address, *big.Int, error)
	AgentOwner(ctx context.Context, agentAddr common.Address) (common.Address, error)
	AgentOperator(ctx context.Context, agentAddr common.Address) (common.Address, error)
	AgentRequester(ctx context.Context, agentAddr common.Address) (common.Address, error)
	AgentAdministrator(ctx context.Context, agentAddr common.Address) (common.Address, error)
	AgentDefaulted(ctx context.Context, agentAddr common.Address) (bool, error)
	AgentVersion(ctx context.Context, agentAddr common.Address) (uint8, uint8, error)
	AgentIsValid(ctx context.Context, agentAddr common.Address) (bool, error)
	AgentMiners(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) ([]address.Address, error)
	AgentLiquidAssets(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error)
	AgentPrincipal(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error)
	AgentInterestOwed(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error)
	AgentDebt(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error)
	// agent factory methods
	AgentFactoryAgentCount(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	// infinity pool methods
	InfPoolGetRate(ctx context.Context) (*big.Int, error)
	InfPoolGetAgentLvl(ctx context.Context, agentID *big.Int) (*big.Int, float64, error)
	InfPoolGetAccount(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (abigen.Account, error)
	InfPoolBorrowableLiquidity(ctx context.Context, blockNumber *big.Int) (*big.Float, error)
	InfPoolTotalAssets(ctx context.Context, blockNumber *big.Int) (*big.Float, error)
	InfPoolTotalBorrowed(ctx context.Context, blockNumber *big.Int) (*big.Float, error)
	InfPoolExitReserve(ctx context.Context, blockNumber *big.Int) (*big.Int, *big.Int, error)
	InfPoolAgentMaxBorrow(ctx context.Context, agentAddr common.Address, agentData *vc.AgentData) (*big.Int, error)
	InfPoolFeesAccrued(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	InfPoolApy(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	// pool registry methods
	TreasuryFeeRate(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	// ifil methods
	IFILBalanceOf(ctx context.Context, hodler common.Address) (*big.Float, error)
	IFILPrice(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	IFILSupply(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	IFILMinter(ctx context.Context) (common.Address, error)
	IFILBurner(ctx context.Context) (common.Address, error)
	// wfil methods
	WFILBalanceOf(ctx context.Context, hodler common.Address) (*big.Float, error)
	WFILAllowance(ctx context.Context, hodler common.Address, spender common.Address) (*big.Float, error)
	// policing methods
	CredentialUsed(ctx context.Context, vc abigen.VerifiableCredential, blockNumber *big.Int) (bool, error)
	CredentialValidityPeriod(ctx context.Context) (*big.Int, *big.Int, error)
	SectorFaultyTolerance(ctx context.Context) (*big.Int, error)
	// miner registry methods
	MinerRegistryAgentMinersCount(ctx context.Context, agentID *big.Int, blockNumber *big.Int) (*big.Int, error)
	MinerRegistryAgentMinersList(ctx context.Context, agentID *big.Int, blockNumber *big.Int) ([]address.Address, error)
	// chain methods
	ChainHeight(ctx context.Context) (*big.Int, error)
	ChainHead(ctx context.Context) (*filtypes.TipSet, error)
	ChainID() *big.Int
	ChainGetNonce(ctx context.Context, fromAddr common.Address) (*big.Int, error)
	// state methods
	StateWaitTx(ctx context.Context, txHash common.Hash, ch chan *types.Receipt)
	StateWaitReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	StateWaitNextTick(ctx context.Context, currentEpochHeight *big.Int) error
	// router methods
	RouterOwner(ctx context.Context) (common.Address, error)
	RouterGetRoute(ctx context.Context, route constants.Route) (common.Address, error)
	// deployment addresses
	AgentPolice() common.Address
	MinerRegistry() common.Address
	Router() common.Address
	AgentFactory() common.Address
	IFIL() common.Address
	WFIL() common.Address
	GLF() common.Address
	SPPlus() common.Address
	InfinityPool() common.Address

	// token related addresses
	TokenNFTWrapper() common.Address
	DelegatedClaimsCampaigns() common.Address
	Governor() common.Address

	// plus methods
	SPPlusTokenIDFromRcpt(ctx context.Context, receipt *types.Receipt) (*big.Int, error)
	SPPlusInfo(ctx context.Context, tokenID *big.Int, blockNumber *big.Int) (*SPPlusInfo, error)
	SPPlusTierInfo(ctx context.Context, blockNumber *big.Int) ([]abigen.TierInfo, error)
	SPPlusTierFromAgentAddress(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (uint8, error)
	SPPlusMintPrice(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	SPPlusTierSwitchPenaltyInfo(ctx context.Context, blockNumber *big.Int) (penaltyWindow *big.Int, penaltyFee *big.Int, err error)
	SPPlusAgentIdToTokenId(ctx context.Context, agentID *big.Int, blockNumber *big.Int) (*big.Int, error)
}

//go:generate mockery --name FEVMActions
type FEVMActions interface {
	// agent actions
	AgentCreate(ctx context.Context, auth *bind.TransactOpts, owner common.Address, operator common.Address, request common.Address) (*types.Transaction, error)
	AgentBorrow(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, poolID *big.Int, amount *big.Int, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPay(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, poolID *big.Int, amount *big.Int, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentAddMiner(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentRemoveMiner(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address, newOwnerAddr address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentChangeMinerWorker(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address, workerAddr address.Address, controlAddrs []address.Address) (*types.Transaction, error)
	AgentConfirmMinerWorkerChange(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address) (*types.Transaction, error)
	AgentPullFunds(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, amount *big.Int, miner address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPushFunds(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, amount *big.Int, miner address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentWithdraw(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, receiver common.Address, amount *big.Int, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentRefreshRoutes(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address) (*types.Transaction, error)
	AgentSetRecovered(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentTransferOwnership(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, newOwner common.Address) (*types.Transaction, error)
	AgentAcceptOwnership(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address) (*types.Transaction, error)
	AgentTransferOperator(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, newOperator common.Address) (*types.Transaction, error)
	AgentAcceptOperator(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address) (*types.Transaction, error)
	AgentChangeRequester(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, newRequester common.Address) (*types.Transaction, error)

	// infinity pool actions
	InfPoolDepositFIL(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, amount *big.Int) (*types.Transaction, error)
	InfPoolWithdraw(ctx context.Context, auth *bind.TransactOpts, assets *big.Int, sender common.Address, receiver common.Address) (*types.Transaction, error)
	InfPoolRedeem(ctx context.Context, auth *bind.TransactOpts, shares *big.Int, sender common.Address, receiver common.Address) (*types.Transaction, error)

	// iFIL actions
	IFILTransfer(ctx context.Context, auth *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error)
	IFILApprove(ctx context.Context, auth *bind.TransactOpts, spender common.Address, allowance *big.Int) (*types.Transaction, error)

	// plus actions
	SPPlusMint(ctx context.Context, auth *bind.TransactOpts) (*types.Transaction, error)
	SPPlusMintAndActivate(ctx context.Context, auth *bind.TransactOpts, beneficiary common.Address, tier uint8) (*types.Transaction, error)
	SPPlusMintActivateAndFund(ctx context.Context, auth *bind.TransactOpts, cashBackPercent *big.Int, beneficiary common.Address, tier uint8, amount *big.Int) (*types.Transaction, error)
	SPPlusActivate(ctx context.Context, auth *bind.TransactOpts, beneficiary common.Address, tokenID *big.Int, tier uint8) (*types.Transaction, error)
	SPPlusSetPersonalCashBackPercent(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, cashBackPercent *big.Int) (*types.Transaction, error)
	SPPlusFundGLFVault(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, amount *big.Int, cashBackPercent *big.Int) (*types.Transaction, error)
	SPPlusClaimCashBack(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, receiver common.Address) (*types.Transaction, error)
	SPPlusUpgrade(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, tier uint8) (*types.Transaction, error)
	SPPlusDowngrade(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, tier uint8, agentAddr common.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	SPPlusWithdrawExtraLockedFunds(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error)
	SPPlusWithdrawGlfVault(ctx context.Context, auth *bind.TransactOpts, tokenID *big.Int, amount *big.Int, receiver common.Address) (*types.Transaction, error)
}

//go:generate mockery --name FEVMExtern
type FEVMExtern interface {
	ConnectEthClient() (*ethclient.Client, error)
	ConnectLotusClient() (*api.FullNodeStruct, jsonrpc.ClientCloser, error)
	ConnectAdoClient(ctx context.Context) (jsonrpc.ClientCloser, error)
	GetEventsURL() string
}

//go:generate mockery --name PoolsSDK
type PoolsSDK interface {
	Query() FEVMQueries
	Act() FEVMActions
	Extern() FEVMExtern
}

type ProtocolMeta struct {
	AgentPolice              common.Address `json:"agentPolice"`
	MinerRegistry            common.Address `json:"minerRegistry"`
	Router                   common.Address `json:"router"`
	AgentFactory             common.Address `json:"agentFactory"`
	IFIL                     common.Address `json:"ifil"`
	WFIL                     common.Address `json:"wfil"`
	GLF                      common.Address `json:"glf"`
	SPPlus                   common.Address `json:"spPlus"`
	InfinityPool             common.Address `json:"infinityPool"`
	Governor                 common.Address `json:"governor"`
	TokenNFTWrapper          common.Address `json:"tokenNFTWrapper"`
	DelegatedClaimsCampaigns common.Address `json:"delegatedClaimsCampaigns"`
	ChainID                  *big.Int       `json:"chainID"`
}

type Extern struct {
	AdoAddr       string `json:"adoAddr"`
	LotusDialAddr string `json:"lotusDialAddr"`
	LotusToken    string `json:"lotusToken"`
	EventsURL     string `json:"eventsURL"`
}

type SPPlusInfo struct {
	AgentID                      *big.Int
	FilCashbackEarned            *big.Int
	GLFVaultBalance              *big.Int
	LastTierSwitchTimestamp      *big.Int
	PersonalCashBackPercent      *big.Int
	TierLockAmount               *big.Int
	WithdrawableExtraLockedFunds *big.Int
	BaseConversionRateFILtoGLF   *big.Int
	Tier                         uint8
}
