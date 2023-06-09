package types

import (
	"context"
	"crypto/ecdsa"
	"math/big"

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
	AgentOwes(ctx context.Context, agentAddr common.Address) (*big.Int, *big.Int, error)
	AgentFaultyEpochStart(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	// agent factory methods
	AgentFactoryAgentCount(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	// infinity pool methods
	InfPoolGetRate(ctx context.Context, cred abigen.VerifiableCredential) (*big.Int, error)
	InfPoolRateFromGCRED(ctx context.Context, gcred *big.Int) (*big.Float, error)
	InfPoolGetAgentLvl(ctx context.Context, agentID *big.Int) (*big.Int, float64, error)
	InfPoolGetAccount(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (abigen.Account, error)
	InfPoolBorrowableLiquidity(ctx context.Context, blockNumber *big.Int) (*big.Float, error)
	InfPoolTotalAssets(ctx context.Context, blockNumber *big.Int) (*big.Float, error)
	InfPoolTotalBorrowed(ctx context.Context, blockNumber *big.Int) (*big.Float, error)
	InfPoolIsApprovedWithReason(ctx context.Context, agentAddr common.Address, agentData *vc.AgentData) (bool, RejectionReason, error)
	InfPoolAgentMaxBorrow(ctx context.Context, agentAddr common.Address, agentData *vc.AgentData) (*big.Int, error)
	InfPoolMaxEpochsOwedTolerance(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	InfPoolFeesAccrued(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	// pool registry methods
	ListPools(ctx context.Context) ([]common.Address, error)
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
	CredentialUsed(ctx context.Context, v uint8, r [32]byte, s [32]byte) (bool, error)
	CredentialValidityPeriod(ctx context.Context) (*big.Int, *big.Int, error)
	DefaultEpoch(ctx context.Context) (*big.Int, error)
	MaxConsecutiveFaultEpochs(ctx context.Context) (*big.Int, error)
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
	PoolRegistry() common.Address
	AgentFactory() common.Address
	IFIL() common.Address
	WFIL() common.Address
	InfinityPool() common.Address
	// RateModule gets fetched from InfinityPool
	RateModule() (common.Address, error)
}

//go:generate mockery --name FEVMActions
type FEVMActions interface {
	// agent actions
	AgentCreate(ctx context.Context, owner common.Address, operator common.Address, request common.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentBorrow(ctx context.Context, agentAddr common.Address, poolID *big.Int, amount *big.Int, senderKey *ecdsa.PrivateKey, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPay(ctx context.Context, agentAddr common.Address, poolID *big.Int, amount *big.Int, senderKey *ecdsa.PrivateKey, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentAddMiner(ctx context.Context, agentAddr common.Address, minerAddr address.Address, senderKey *ecdsa.PrivateKey, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentRemoveMiner(ctx context.Context, agentAddr common.Address, minerAddr address.Address, newOwnerAddr address.Address, senderKey *ecdsa.PrivateKey, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentChangeMinerWorker(ctx context.Context, agentAddr common.Address, minerAddr address.Address, workerAddr address.Address, controlAddrs []address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentConfirmMinerWorkerChange(ctx context.Context, agentAddr common.Address, minerAddr address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPullFunds(ctx context.Context, agentAddr common.Address, amount *big.Int, miner address.Address, senderKey *ecdsa.PrivateKey, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPushFunds(ctx context.Context, agentAddr common.Address, amount *big.Int, miner address.Address, senderKey *ecdsa.PrivateKey, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentWithdraw(ctx context.Context, agentAddr common.Address, receiver common.Address, amount *big.Int, senderKey *ecdsa.PrivateKey, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentRefreshRoutes(ctx context.Context, agentAddr common.Address, senderKey *ecdsa.PrivateKey) (*types.Transaction, error)

	// infinity pool actions
	InfPoolDepositFIL(ctx context.Context, agentAddr common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)

	// iFIL actions
	IFILTransfer(ctx context.Context, receiver common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	IFILApprove(ctx context.Context, spender common.Address, allowance *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)
}

//go:generate mockery --name FEVMExtern
type FEVMExtern interface {
	ConnectEthClient() (*ethclient.Client, error)
	ConnectLotusClient() (*api.FullNodeStruct, jsonrpc.ClientCloser, error)
	ConnectAdoClient(ctx context.Context) (jsonrpc.ClientCloser, error)
}

//go:generate mockery --name PoolsSDK
type PoolsSDK interface {
	Query() FEVMQueries
	Act() FEVMActions
	Extern() FEVMExtern
}

type AgentOwed struct {
	AgentAddr  common.Address `json:"agentAddr"`
	AmountOwed *big.Int       `json:"amountOwed"`
	Gcred      *big.Int       `json:"gcred"`
}

type ProtocolMeta struct {
	AgentPolice   common.Address `json:"agentPolice"`
	MinerRegistry common.Address `json:"minerRegistry"`
	Router        common.Address `json:"router"`
	PoolRegistry  common.Address `json:"poolRegistry"`
	AgentFactory  common.Address `json:"agentFactory"`
	IFIL          common.Address `json:"ifil"`
	WFIL          common.Address `json:"wfil"`
	InfinityPool  common.Address `json:"infinityPool"`
	ChainID       *big.Int       `json:"chainID"`
}

type Extern struct {
	AdoAddr       string `json:"adoAddr"`
	LotusDialAddr string `json:"lotusDialAddr"`
	LotusToken    string `json:"lotusToken"`
}
