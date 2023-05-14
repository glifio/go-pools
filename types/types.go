package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/glifio/go-pools/abigen"
)

type FEVMQueries interface {
	// agent methods
	AgentID(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	AgentAccount(ctx context.Context, agentAddr common.Address, poolID *big.Int) (abigen.Account, error)
	AgentAddrIDFromRcpt(ctx context.Context, rcpt *types.Receipt) (common.Address, *big.Int, error)
	AgentOwner(ctx context.Context, agentAddr common.Address) (common.Address, error)
	AgentIsValid(ctx context.Context, agentAddr common.Address) (bool, error)
	AgentMiners(ctx context.Context, agentAddr common.Address) ([]address.Address, error)
	AgentLiquidAssets(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	AgentPrincipal(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	// infinity pool methods
	InfPoolGetRate(ctx context.Context, cred abigen.VerifiableCredential) (*big.Int, error)
	InfPoolAgentAccount(ctx context.Context, agentAddr common.Address) (abigen.Account, error)
	// policing methods
	CredentialUsed(ctx context.Context, v uint8, r [32]byte, s [32]byte) (bool, error)
	CredentialValidityPeriod(ctx context.Context) (*big.Int, *big.Int, error)
	// chain methods
	ChainHeight(ctx context.Context) (*big.Int, error)
	ChainID() *big.Int
	// state methods
	StateWaitTx(ctx context.Context, txHash common.Hash, ch chan *types.Receipt)
	StateWaitReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// deployment addresses
	AgentPolice() common.Address
	MinerRegistry() common.Address
	Router() common.Address
	PoolRegistry() common.Address
	AgentFactory() common.Address
	IFIL() common.Address
	InfinityPool() common.Address
}

type FEVMActions interface {
	// agent actions
	// AgentCreate(ctx context.Context, owner common.Address, principal *big.Int, miners []address.Address) (common.Address, error)
}

type FEVMExtern interface {
	ConnectEthClient() (*ethclient.Client, error)
	ConnectLotusClient() (*api.FullNodeStruct, jsonrpc.ClientCloser, error)
	ConnectAdoClient(ctx context.Context) (jsonrpc.ClientCloser, error)
}

type FEVMConnection interface {
	Query() FEVMQueries
	Act() FEVMActions
	Extern() FEVMExtern
}

type AgentOwed struct {
	AgentAddr  common.Address
	AmountOwed *big.Int
	Gcred      *big.Int
}
