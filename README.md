# GO POOLS

An SDK for working with the GLIF Pools protocol

Check out the `/types` directory for the latest exports. 

```
type FEVMQueries interface {
	// agent methods
	AgentID(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	AgentAccount(ctx context.Context, agentAddr common.Address, poolID *big.Int) (abigen.Account, error)
	AgentAddrIDFromRcpt(ctx context.Context, rcpt *types.Receipt) (common.Address, *big.Int, error)
	AgentOwner(ctx context.Context, agentAddr common.Address) (common.Address, error)
	AgentVersion(ctx context.Context, agentAddr common.Address) (uint8, error)
	AgentIsValid(ctx context.Context, agentAddr common.Address) (bool, error)
	AgentMiners(ctx context.Context, agentAddr common.Address) ([]address.Address, error)
	AgentLiquidAssets(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	AgentPrincipal(ctx context.Context, agentAddr common.Address) (*big.Int, error)
	AgentOwes(ctx context.Context, agentAddr common.Address) (*big.Int, *big.Int, error)
	// infinity pool methods
	InfPoolGetRate(ctx context.Context, cred abigen.VerifiableCredential) (*big.Int, error)
	InfPoolGetAgentLvl(ctx context.Context, agentID *big.Int) (*big.Int, float64, error)
	InfPoolGetAccount(ctx context.Context, agentAddr common.Address) (abigen.Account, error)
	InfPoolBorrowableLiquidity(ctx context.Context) (*big.Float, error)
	// pool registry methods
	ListPools(ctx context.Context) ([]common.Address, error)
	// ifil methods
	IFILBalanceOf(ctx context.Context, hodler common.Address) (*big.Float, error)
	IFILPrice(ctx context.Context) (*big.Float, error)
	// policing methods
	CredentialUsed(ctx context.Context, v uint8, r [32]byte, s [32]byte) (bool, error)
	CredentialValidityPeriod(ctx context.Context) (*big.Int, *big.Int, error)
	DefaultEpoch(ctx context.Context) (*big.Int, error)
	// chain methods
	ChainHeight(ctx context.Context) (*big.Int, error)
	ChainID() *big.Int
	ChainGetNonce(ctx context.Context, fromAddr common.Address) (*big.Int, error)
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
	AgentCreate(ctx context.Context, owner common.Address, operator common.Address, request common.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentBorrow(ctx context.Context, agentAddr common.Address, poolID *big.Int, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPay(ctx context.Context, agentAddr common.Address, poolID *big.Int, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentAddMiner(ctx context.Context, agentAddr common.Address, minerAddr address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentRemoveMiner(ctx context.Context, agentAddr common.Address, minerAddr address.Address, newOwnerAddr address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentChangeMinerWorker(ctx context.Context, agentAddr common.Address, minerAddr address.Address, workerAddr address.Address, controlAddrs []address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPullFunds(ctx context.Context, agentAddr common.Address, amount *big.Int, miner address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentPushFunds(ctx context.Context, agentAddr common.Address, amount *big.Int, miner address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	AgentWithdraw(ctx context.Context, agentAddr common.Address, receiver common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)

	// infinity pool actions
	InfPoolDepositFIL(ctx context.Context, agentAddr common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)

	// iFIL actions
	IFILTransfer(ctx context.Context, receiver common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)
	IFILApprove(ctx context.Context, spender common.Address, allowance *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error)
}

// serve connections to external services
type FEVMExtern interface {
	ConnectEthClient() (*ethclient.Client, error)
	ConnectLotusClient() (*api.FullNodeStruct, jsonrpc.ClientCloser, error)
	ConnectAdoClient(ctx context.Context) (jsonrpc.ClientCloser, error)
}

type PoolsSDK interface {
	Query() FEVMQueries
	Act() FEVMActions
	Extern() FEVMExtern
}
```

