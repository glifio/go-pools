package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
)

func (q *fevmQueries) AgentID(ctx context.Context, address common.Address) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return nil, err
	}

	agentID, err := agentCaller.Id(nil)
	if err != nil {
		return nil, err
	}

	return agentID, nil
}

func (q *fevmQueries) AgentOwner(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	return agentCaller.Owner(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentAdministrator(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	return agentCaller.Administration(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentDefaulted(ctx context.Context, address common.Address) (bool, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return false, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return false, err
	}

	return agentCaller.Defaulted(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentOperator(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	return agentCaller.Operator(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentRequester(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	return agentCaller.AdoRequestKey(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentIsValid(ctx context.Context, address common.Address) (bool, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return false, err
	}
	defer client.Close()

	agentFactoryCaller, err := abigen.NewAgentFactoryCaller(q.agentFactory, client)
	if err != nil {
		return false, err
	}

	isAgent, err := agentFactoryCaller.IsAgent(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return false, err
	}

	return isAgent, nil
}

func (q *fevmQueries) AgentMiners(
	ctx context.Context,
	agentAddr common.Address,
	blockNumber *big.Int,
) ([]address.Address, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	id, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	minerRegCaller, err := abigen.NewMinerRegistryCaller(q.minerRegistry, ethClient)
	if err != nil {
		return nil, err
	}

	minerCount, err := minerRegCaller.MinersCount(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, id)
	if err != nil {
		return nil, err
	}

	miners := []address.Address{}

	for i := big.NewInt(0); i.Cmp(minerCount) == -1; i.Add(i, big.NewInt(1)) {
		miner, err := minerRegCaller.GetMiner(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, id, i)
		if err != nil {
			return nil, err
		}

		minerAddr, err := address.NewIDAddress(miner)
		if err != nil {
			return nil, err
		}

		miners = append(miners, minerAddr)
	}
	return miners, nil
}

func (q *fevmQueries) AgentLiquidAssets(ctx context.Context, address common.Address, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return nil, err
	}

	assets, err := agentCaller.LiquidAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (q *fevmQueries) AgentPrincipal(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	agentID, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	poolCaller, err := abigen.NewInfinityPoolV2(q.infinityPool, ethClient)
	if err != nil {
		return nil, err
	}

	return poolCaller.GetAgentBorrowed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, agentID)
}

func (q *fevmQueries) AgentAccount(ctx context.Context, agentAddr common.Address, poolID *big.Int, blockNumber *big.Int) (abigen.Account, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return abigen.Account{}, err
	}
	defer ethClient.Close()

	agentID, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return abigen.Account{}, err
	}

	routerCaller, err := abigen.NewRouterCaller(q.router, ethClient)
	if err != nil {
		return abigen.Account{}, err
	}

	return routerCaller.GetAccount(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, agentID, poolID)
}

func (q *fevmQueries) AgentAddrIDFromRcpt(ctx context.Context, receipt *types.Receipt) (common.Address, *big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, nil, err
	}
	defer client.Close()

	agentFactoryABI, err := abigen.AgentFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, err
	}

	agentFactoryFilterer, err := abigen.NewAgentFactoryFilterer(q.agentFactory, client)
	if err != nil {
		return common.Address{}, nil, err
	}

	var agentID *big.Int
	var agentAddr common.Address

	for _, l := range receipt.Logs {
		event, err := agentFactoryABI.EventByID(l.Topics[0])
		if err != nil {
			return common.Address{}, nil, err
		}
		if event.Name == "CreateAgent" {
			createAgentEvent, err := agentFactoryFilterer.ParseCreateAgent(*l)
			if err != nil {
				return common.Address{}, nil, err
			}
			agentAddr = createAgentEvent.Agent
			agentID = createAgentEvent.AgentID
		}
	}

	return agentAddr, agentID, nil
}

func (q *fevmQueries) AgentInterestOwed(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}

	pool, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, ethClient)
	if err != nil {
		return nil, err
	}

	id, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	return pool.GetAgentInterestOwed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, id)
}

func (q *fevmQueries) AgentDebt(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}

	pool, err := abigen.NewInfinityPoolV2Caller(q.infinityPool, ethClient)
	if err != nil {
		return nil, err
	}

	id, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	return pool.GetAgentDebt(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, id)
}

func (q *fevmQueries) AgentFaultyEpochStart(ctx context.Context, agentAddr common.Address) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(agentAddr, client)
	if err != nil {
		return nil, err
	}

	return agentCaller.FaultySectorStartEpoch(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentVersion(ctx context.Context, agentAddr common.Address) (uint8, uint8, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return 0, 0, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(agentAddr, client)
	if err != nil {
		return 0, 0, err
	}

	agentVersion, err := agentCaller.Version(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, 0, err
	}

	agDeployer, err := q.RouterGetRoute(ctx, constants.RouteAgentDeployer)
	if err != nil {
		return 0, 0, err
	}

	agDeployerCaller, err := abigen.NewAgentDeployerCaller(agDeployer, client)
	if err != nil {
		return 0, 0, err
	}

	deployerVersion, err := agDeployerCaller.Version(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, 0, err
	}

	return agentVersion, deployerVersion, nil
}
