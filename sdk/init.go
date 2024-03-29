package sdk

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/deploy"
	"github.com/glifio/go-pools/types"
)

func InitFEVMConnection(
	agentPolice common.Address,
	minerRegistry common.Address,
	router common.Address,
	poolRegistry common.Address,
	agentFactory common.Address,
	iFIL common.Address,
	wFIL common.Address,
	infinityPool common.Address,
	simpleRamp common.Address,
	adoAddr string,
	adoNamespace string,
	dialAddr string,
	token string,
	chainID *big.Int,
) *fevmConnection {
	extern := &fevmExtern{
		dialAddr:     dialAddr,
		token:        token,
		adoAddr:      adoAddr,
		adoNamespace: adoNamespace,
	}

	fevmQueries := &fevmQueries{
		router:        router,
		iFIL:          iFIL,
		wFIL:          wFIL,
		infinityPool:  infinityPool,
		simpleRamp:    simpleRamp,
		agentFactory:  agentFactory,
		poolRegistry:  poolRegistry,
		minerRegistry: minerRegistry,
		agentPolice:   agentPolice,
		chainID:       chainID,
		extern:        extern,
	}

	fevmActions := &fevmActions{extern: extern, queries: fevmQueries}

	return &fevmConnection{
		query:  fevmQueries,
		act:    fevmActions,
		extern: extern,
	}
}

func getRoutes(ctx context.Context, routes []constants.Route, routerCaller *abigen.RouterCaller) (map[constants.Route]common.Address, error) {
	type routeResult struct {
		Route   constants.Route
		Address common.Address
	}

	var wg sync.WaitGroup
	results := make([]routeResult, len(routes))
	errc := make(chan error, 1)

	for i, route := range routes {
		wg.Add(1)
		go func(i int, route constants.Route) {
			defer wg.Done()

			address, err := routerCaller.GetRoute0(&bind.CallOpts{Context: ctx}, string(route))
			if err != nil {
				select {
				case errc <- err:
				default:
					// Don't block if an error has already been sent.
				}
				return
			}

			results[i] = routeResult{
				Route:   route,
				Address: address,
			}
		}(i, route)
	}

	wg.Wait()

	select {
	case err := <-errc:
		// Return any error that was sent on the channel.
		return nil, err
	default:
		// If no error was sent on the channel, return the results.
		routeMap := make(map[constants.Route]common.Address)
		for _, result := range results {
			routeMap[result.Route] = result.Address
		}
		return routeMap, nil
	}
}

func LazyInit(
	ctx context.Context,
	sdk *types.PoolsSDK,
	router common.Address,
	adoAddr string,
	adoNamespace string,
	dialAddr string,
	token string,
) error {
	client, err := ethclient.Dial(dialAddr)
	if err != nil {
		return err
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	routerCaller, err := abigen.NewRouterCaller(router, client)
	if err != nil {
		return err
	}

	fetchRoutes := []constants.Route{
		constants.RouteAgentPolice,
		constants.RouteAgentFactory,
		constants.RoutePoolRegistry,
		constants.RouteMinerRegistry,
		constants.RouteWFIL,
	}

	routes, err := getRoutes(ctx, fetchRoutes, routerCaller)
	if err != nil {
		return err
	}

	poolRegCaller, err := abigen.NewPoolRegistryCaller(routes[constants.RoutePoolRegistry], client)
	if err != nil {
		return err
	}

	// infpool is poolID 0
	infpool, err := poolRegCaller.AllPools(&bind.CallOpts{Context: ctx}, big.NewInt(0))
	if err != nil {
		return err
	}

	infpoolCaller, err := abigen.NewInfinityPoolCaller(infpool, client)
	if err != nil {
		return err
	}

	iFIL, err := infpoolCaller.LiquidStakingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	simpleRamp, err := infpoolCaller.Ramp(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	*sdk = InitFEVMConnection(
		routes[constants.RouteAgentPolice],
		routes[constants.RouteMinerRegistry],
		router,
		routes[constants.RoutePoolRegistry],
		routes[constants.RouteAgentFactory],
		iFIL,
		routes[constants.RouteWFIL],
		infpool,
		simpleRamp,
		adoAddr,
		adoNamespace,
		dialAddr,
		token,
		chainID,
	)

	return nil
}

func New(
	ctx context.Context,
	chainID *big.Int,
	extern types.Extern,
) (types.PoolsSDK, error) {
	var sdk types.PoolsSDK
	ethClient, err := connectEthClient(extern.LotusDialAddr, extern.LotusToken)
	if err != nil {
		return nil, err
	}

	id, err := ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	var protoMeta types.ProtocolMeta
	switch id.Int64() {
	case constants.MainnetChainID:
		protoMeta = deploy.ProtoMeta
	case constants.CalibnetChainID:
		protoMeta = deploy.TestProtoMeta
	default:
		return nil, fmt.Errorf("unsupported chain id: %d", id.Int64())
	}

	if protoMeta.ChainID.Cmp(chainID) != 0 {
		return nil, fmt.Errorf("chain id mismatch: %d != %d", protoMeta.ChainID, chainID)
	}

	var namespace = "Mock"
	if chainID.Int64() == constants.MainnetChainID {
		namespace = "ADO"
	}

	sdk = InitFEVMConnection(
		protoMeta.AgentPolice,
		protoMeta.MinerRegistry,
		protoMeta.Router,
		protoMeta.PoolRegistry,
		protoMeta.AgentFactory,
		protoMeta.IFIL,
		protoMeta.WFIL,
		protoMeta.InfinityPool,
		protoMeta.SimpleRamp,
		extern.AdoAddr,
		namespace,
		extern.LotusDialAddr,
		extern.LotusToken,
		chainID,
	)

	return sdk, nil
}

func Init(
	ctx context.Context,
	sdk *types.PoolsSDK,
	agentPolice common.Address,
	minerRegistry common.Address,
	router common.Address,
	poolRegistry common.Address,
	agentFactory common.Address,
	iFIL common.Address,
	wFIL common.Address,
	infinityPool common.Address,
	simpleRamp common.Address,
	adoAddr string,
	adoNamespace string,
	dialAddr string,
	token string,
) error {
	client, err := connectEthClient(dialAddr, token)
	if err != nil {
		return err
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	*sdk = InitFEVMConnection(
		agentPolice,
		minerRegistry,
		router,
		poolRegistry,
		agentFactory,
		iFIL,
		wFIL,
		infinityPool,
		simpleRamp,
		adoAddr,
		adoNamespace,
		dialAddr,
		token,
		chainID,
	)

	return nil
}

// Implement the FEVMConnection interface
func (c *fevmConnection) Query() types.FEVMQueries {
	return c.query
}

func (c *fevmConnection) Act() types.FEVMActions {
	return c.act
}

func (c *fevmConnection) Extern() types.FEVMExtern {
	return c.extern
}
