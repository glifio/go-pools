package sdk

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"errors"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/filecoin-project/go-address"
// 	"github.com/glif-confidential/cli/rpc"
// 	"github.com/glif-confidential/cli/util"
// 	"github.com/glifio/go-pools/abigen"
// 	"github.com/spf13/viper"
// )

// func (a *fevmActions) AgentCreate(ctx context.Context, deployerPk *ecdsa.PrivateKey, owner common.Address, operator common.Address, request common.Address) (*types.Transaction, error) {
// 	client, err := a.nodes.ConnectEthClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()

// 	agentFactoryTransactor, err := abigen.NewAgentFactoryTransactor(a.queries.AgentFactory(), client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	args := []interface{}{owner, operator, request}

// 	return WriteTx(ctx, deployerPk, client, common.Big0, args, agentFactoryTransactor.Create, "Agent Create")
// }

// // AgentPullFunds pulls funds from the agent to a miner
// func (q *fevmQueries) AgentPullFunds(
// 	ctx context.Context,
// 	agentAddr common.Address,
// 	amount *big.Int,
// 	miner address.Address,
// 	pk *ecdsa.PrivateKey,
// ) (*types.Transaction, error) {
// 	client, err := q.nodes.ConnectEthClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()

// 	as := util.AgentStore()
// 	agentIDStr, err := as.Get("id")
// 	if err != nil {
// 		return nil, err
// 	}

// 	agentID, ok := new(big.Int).SetString(agentIDStr, 10)
// 	if !ok {
// 		return nil, errors.New("could not convert agent id to big.Int")
// 	}

// 	minerRegistryCaller, err := abigen.NewMinerRegistryCaller(c.MinerRegistryAddr, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	minerU64, err := address.IDFromAddress(miner)
// 	if err != nil {
// 		return nil, err
// 	}

// 	registered, err := minerRegistryCaller.MinerRegistered(nil, agentID, minerU64)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !registered {
// 		return nil, errors.New("Miner not registered with agent. Be sure to call `agent add-miner` first before pulling funds.")
// 	}

// 	closer, err := rpc.NewADOClient(ctx, viper.GetString("ado.address"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer closer()

// 	sc, err := rpc.ADOClient.PullFunds(ctx, agentAddr, amount, miner)
// 	if err != nil {
// 		return nil, err
// 	}

// 	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	args := []interface{}{sc}

// 	return WriteTx(ctx, pk, client, common.Big0, args, agentTransactor.PullFunds, "Agent Pull Funds")
// }

// // AgentPushFunds pushes funds from the agent to a miner
// func (q *fevmQueries) AgentPushFunds(
// 	ctx context.Context,
// 	agentAddr common.Address,
// 	amount *big.Int,
// 	miner address.Address,
// 	pk *ecdsa.PrivateKey,
// ) (*types.Transaction, error) {
// 	client, err := q.nodes.ConnectEthClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()

// 	as := util.AgentStore()
// 	agentIDStr, err := as.Get("id")
// 	if err != nil {
// 		return nil, err
// 	}

// 	agentID, ok := new(big.Int).SetString(agentIDStr, 10)
// 	if !ok {
// 		return nil, errors.New("could not convert agent id to big.Int")
// 	}

// 	minerRegistryCaller, err := abigen.NewMinerRegistryCaller(c.MinerRegistryAddr, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	minerU64, err := address.IDFromAddress(miner)
// 	if err != nil {
// 		return nil, err
// 	}

// 	registered, err := minerRegistryCaller.MinerRegistered(nil, agentID, minerU64)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !registered {
// 		return nil, errors.New("Miner not registered with agent. Be sure to call `agent add-miner` first before pushing funds.")
// 	}

// 	closer, err := rpc.NewADOClient(ctx, viper.GetString("ado.address"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer closer()

// 	sc, err := rpc.ADOClient.PushFunds(ctx, agentAddr, amount, miner)
// 	if err != nil {
// 		return nil, err
// 	}

// 	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	args := []interface{}{sc}

// 	return WriteTx(ctx, pk, client, common.Big0, args, agentTransactor.PushFunds, "Agent Push Funds")
// }

// func (q *fevmQueries) AgentBorrow(
// 	ctx context.Context,
// 	agentAddr common.Address,
// 	poolID *big.Int,
// 	amount *big.Int,
// 	pk *ecdsa.PrivateKey,
// ) (*types.Transaction, error) {
// 	client, err := q.nodes.ConnectEthClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()

// 	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	closer, err := rpc.NewADOClient(ctx, viper.GetString("ado.address"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer closer()

// 	sc, err := rpc.ADOClient.Borrow(ctx, agentAddr, amount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	args := []interface{}{poolID, sc}

// 	return WriteTx(ctx, pk, client, common.Big0, args, agentTransactor.Borrow, "Agent Borrow")
// }

// func (q *fevmQueries) AgentPay(
// 	ctx context.Context,
// 	agentAddr common.Address,
// 	poolID *big.Int,
// 	amount *big.Int,
// 	pk *ecdsa.PrivateKey,
// ) (*types.Transaction, error) {
// 	client, err := q.nodes.ConnectEthClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()

// 	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	closer, err := rpc.NewADOClient(ctx, viper.GetString("ado.address"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer closer()

// 	sc, err := rpc.ADOClient.Pay(ctx, agentAddr, amount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	args := []interface{}{poolID, sc}

// 	return WriteTx(ctx, pk, client, common.Big0, args, agentTransactor.Pay, "Agent Pay")
// }

// func (q *fevmQueries) AgentWithdraw(
// 	ctx context.Context,
// 	agentAddr common.Address,
// 	receiver common.Address,
// 	amount *big.Int,
// 	pk *ecdsa.PrivateKey,
// ) (*types.Transaction, error) {
// 	client, err := q.nodes.ConnectEthClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()

// 	agentTransactor, err := abigen.NewAgentTransactor(agentAddr, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	closer, err := rpc.NewADOClient(ctx, viper.GetString("ado.address"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer closer()

// 	sc, err := rpc.ADOClient.Withdraw(ctx, agentAddr, amount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	args := []interface{}{receiver, sc}

// 	return WriteTx(ctx, pk, client, common.Big0, args, agentTransactor.Withdraw, "Agent Withdraw")
// }
