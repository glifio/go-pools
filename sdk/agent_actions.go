package sdk

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/rpc"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) AgentCreate(ctx context.Context, owner common.Address, operator common.Address, request common.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentFactoryTransactor, err := abigen.NewAgentFactoryTransactor(a.queries.AgentFactory(), client)
	if err != nil {
		return nil, err
	}

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{owner, operator, request}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentFactoryTransactor.Create, "Agent Create")
}

func (a *fevmActions) AgentBorrow(ctx context.Context, agentAddr common.Address, poolID *big.Int, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	sc, err := rpc.ADOClient.Borrow(ctx, agentAddr, amount)
	if err != nil {
		log.Fatal(err)
	}

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{poolID, sc}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.Borrow, "Agent Borrow")
}

func (a *fevmActions) AgentPay(ctx context.Context, agentAddr common.Address, poolID *big.Int, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.Pay(ctx, agentAddr, amount)
	if err != nil {
		log.Fatal(err)
	}

	args := []interface{}{poolID, sc}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.Pay, "Agent Pay")
}

func (a *fevmActions) AgentAddMiner(ctx context.Context, agentAddr common.Address, minerAddr address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.AddMiner(ctx, agentAddr, minerAddr)
	if err != nil {
		return nil, err
	}

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{sc}

	tx, err := util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.AddMiner, "Agent Add Miner")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (a *fevmActions) AgentRemoveMiner(ctx context.Context, agentAddr common.Address, minerAddr address.Address, newOwnerAddr address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.RemoveMiner(ctx, agentAddr, minerAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{newOwner64, sc}

	tx, err := util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.RemoveMiner, "Agent Remove Miner")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (a *fevmActions) AgentChangeMinerWorker(ctx context.Context, agentAddr common.Address, minerAddr address.Address, workerAddr address.Address, controlAddrs []address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{minerID, workerID, controlIDs}

	tx, err := util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.ChangeMinerWorker, "Agent Change Miner Worker")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (a *fevmActions) AgentConfirmMinerWorkerChange(ctx context.Context, agentAddr common.Address, minerAddr address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{minerU64}

	tx, err := util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.ConfirmChangeMinerWorker, "Agent Confirm Miner Worker Change")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// AgentPullFunds pulls funds from the agent to a miner
func (a *fevmActions) AgentPullFunds(ctx context.Context, agentAddr common.Address, amount *big.Int, miner address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	minerU64, err := address.IDFromAddress(miner)
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

	sc, err := rpc.ADOClient.PullFunds(ctx, agentAddr, amount, miner)
	if err != nil {
		return nil, err
	}

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{sc}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.PullFunds, "Agent Pull Funds")
}

// AgentPushFunds pushes funds from the agent to a miner
func (a *fevmActions) AgentPushFunds(ctx context.Context, agentAddr common.Address, amount *big.Int, miner address.Address, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	minerU64, err := address.IDFromAddress(miner)
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

	sc, err := rpc.ADOClient.PushFunds(ctx, agentAddr, amount, miner)
	if err != nil {
		return nil, err
	}

	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
	if err != nil {
		return nil, err
	}

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{sc}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.PushFunds, "Agent Push Funds")
}

func (a *fevmActions) AgentWithdraw(ctx context.Context, agentAddr common.Address, receiver common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
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

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	sc, err := rpc.ADOClient.Withdraw(ctx, agentAddr, amount)
	if err != nil {
		log.Fatal(err)
	}

	args := []interface{}{receiver, sc}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, agentTransactor.Withdraw, "Agent Withdraw")
}
