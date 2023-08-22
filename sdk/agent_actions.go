package sdk

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	token "github.com/glifio/go-pools/jws"
	"github.com/glifio/go-pools/rpc"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) AgentCreate(ctx context.Context, auth *bind.TransactOpts, owner common.Address, operator common.Address, request common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	auth.Value = common.Big0

	agentFactory, err := abigen.NewAgentFactoryTransactor(a.queries.AgentFactory(), client)
	if err != nil {
		return nil, err
	}

	tx, err := agentFactory.Create(auth, owner, operator, request)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) AgentBorrow(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, poolID *big.Int, amount *big.Int, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, address.Undef, amount, constants.MethodBorrow, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	args := []interface{}{poolID, sc}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.Borrow, "Agent Borrow")
}

func (a *fevmActions) AgentPay(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, poolID *big.Int, amount *big.Int, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, address.Undef, amount, constants.MethodPay, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	args := []interface{}{poolID, sc}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.Pay, "Agent Pay")
}

func (a *fevmActions) AgentAddMiner(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, common.Big0, constants.MethodAddMiner, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	auth.Value = common.Big0

	agent, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	tx, err := agent.AddMiner(auth, sc)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) AgentRemoveMiner(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address, newOwnerAddr address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	if newOwnerAddr.Protocol() != address.ID {
		lapi, closer, err := a.extern.ConnectLotusClient()
		if err != nil {
			return nil, err
		}
		defer closer()

		idAddr, err := lapi.StateLookupID(context.Background(), newOwnerAddr, ltypes.EmptyTSK)
		if err != nil {
			return nil, err
		}
		newOwnerAddr = idAddr
	}

	newOwner64, err := address.IDFromAddress(newOwnerAddr)
	if err != nil {
		return nil, err
	}

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, common.Big0, constants.MethodRemoveMiner, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	args := []interface{}{newOwner64, sc}

	tx, err := util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.RemoveMiner, "Agent Remove Miner")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (a *fevmActions) AgentChangeMinerWorker(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address, workerAddr address.Address, controlAddrs []address.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	// convert miner address to ID address
	minerID, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return nil, err
	}

	// convert worker address to ID address
	workerID, err := address.IDFromAddress(workerAddr)
	if err != nil {
		return nil, err
	}

	// convert control addresses to ID addresses
	var controlIDs []uint64
	for _, controlAddr := range controlAddrs {
		controlID, err := address.IDFromAddress(controlAddr)
		if err != nil {
			return nil, err
		}
		controlIDs = append(controlIDs, controlID)
	}

	args := []interface{}{minerID, workerID, controlIDs}

	tx, err := util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.ChangeMinerWorker, "Agent Change Miner Worker")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (a *fevmActions) AgentConfirmMinerWorkerChange(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, minerAddr address.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	minerU64, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return nil, err
	}

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{minerU64}

	tx, err := util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.ConfirmChangeMinerWorker, "Agent Confirm Miner Worker Change")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// AgentPullFunds pulls funds from the agent to a miner
func (a *fevmActions) AgentPullFunds(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, amount *big.Int, minerAddr address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentID, err := a.queries.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	minerRegistryCaller, err := abigen.NewMinerRegistryCaller(a.queries.MinerRegistry(), client)
	if err != nil {
		return nil, err
	}

	minerU64, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return nil, err
	}

	registered, err := minerRegistryCaller.MinerRegistered(nil, agentID, minerU64)
	if err != nil {
		return nil, err
	}

	if !registered {
		return nil, errors.New("Miner not registered with agent. Be sure to call `agent add-miner` first before pulling funds.")
	}

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, amount, constants.MethodPullFunds, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{sc}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.PullFunds, "Agent Pull Funds")
}

// AgentPushFunds pushes funds from the agent to a miner
func (a *fevmActions) AgentPushFunds(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, amount *big.Int, minerAddr address.Address, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentID, err := a.queries.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	minerRegistryCaller, err := abigen.NewMinerRegistryCaller(a.queries.MinerRegistry(), client)
	if err != nil {
		return nil, err
	}

	minerU64, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return nil, err
	}

	registered, err := minerRegistryCaller.MinerRegistered(nil, agentID, minerU64)
	if err != nil {
		return nil, err
	}

	if !registered {
		return nil, errors.New("Miner not registered with agent. Be sure to call `agent add-miner` first before pushing funds.")
	}

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, amount, constants.MethodPushFunds, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{sc}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.PushFunds, "Agent Push Funds")
}

func (a *fevmActions) AgentWithdraw(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address, receiver common.Address, amount *big.Int, requesterKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, address.Undef, amount, constants.MethodWithdraw, requesterKey, a.queries)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return nil, err
	}

	args := []interface{}{receiver, sc}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, args, agentTransactor.Withdraw, "Agent Withdraw")
}

func (a *fevmActions) AgentRefreshRoutes(ctx context.Context, auth *bind.TransactOpts, agentAddr common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	return util.WriteTxStaging(ctx, auth, a.queries.ChainID(), common.Big0, nil, []interface{}{}, agentTransactor.RefreshRoutes, "Agent RefreshRoutes")
}
