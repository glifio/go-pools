package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
	"github.com/jimpick/go-ethereum/common"
)

func WFilDepositEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.WFILDeposit, error) {

	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.WFILDeposit{}, err
	}

	filterer, err := abigen.NewWFILFilterer(sdk.Query().WFIL(), ethclient)
	if err != nil {
		return []*abigen.WFILDeposit{}, err
	}

	var events []*abigen.WFILDeposit

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterDeposit(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil)
		if err != nil {
			return []*abigen.WFILDeposit{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}

func WFilTransferEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int, from []common.Address, to []common.Address) ([]*abigen.WFILTransfer, error) {

	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.WFILTransfer{}, err
	}

	filterer, err := abigen.NewWFILFilterer(sdk.Query().WFIL(), ethclient)
	if err != nil {
		return []*abigen.WFILTransfer{}, err
	}

	var events []*abigen.WFILTransfer

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterTransfer(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), from, to) // from, to []common.Address{sdk.Query().InfinityPool()}
		if err != nil {
			return []*abigen.WFILTransfer{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}
