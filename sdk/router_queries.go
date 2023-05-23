package sdk

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
)

func (q *fevmQueries) RouterOwner(ctx context.Context) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	routerCaller, err := abigen.NewRouterCaller(q.router, client)
	if err != nil {
		return common.Address{}, err
	}

	return routerCaller.Owner(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) RouterGetRoute(ctx context.Context, route constants.Route) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	routerCaller, err := abigen.NewRouterCaller(q.router, client)
	if err != nil {
		return common.Address{}, err
	}

	return routerCaller.GetRoute0(&bind.CallOpts{Context: ctx}, string(route))
}
