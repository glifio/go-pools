package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
)

func (q *fevmQueries) PlusTokenIDFromRcpt(ctx context.Context, receipt *types.Receipt) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plusABI, err := abigen.PlusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	plusFilterer, err := abigen.NewPlusFilterer(q.plus, client)
	if err != nil {
		return nil, err
	}

	var tokenID *big.Int

	topicHash := plusABI.Events["Transfer"].ID

	for _, l := range receipt.Logs {
		event, err := plusABI.EventByID(topicHash)
		if err != nil {
			return nil, err
		}
		if event.Name == "Transfer" && len(l.Topics) == 4 {
			// Check length to match NFT Transfer event, not ERC20
			transferEvent, err := plusFilterer.ParseTransfer(*l)
			if err != nil {
				return nil, err
			}
			tokenID = transferEvent.TokenId
			break
		}
	}

	return tokenID, nil
}
