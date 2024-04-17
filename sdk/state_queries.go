package sdk

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/jimpick/go-ethereum/common"
	"github.com/jimpick/go-ethereum/core/types"
)

func (q *fevmQueries) StateWaitTx(ctx context.Context, hash common.Hash, ch chan *types.Receipt) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return
	}
	defer client.Close()

	for {
		time.Sleep(time.Millisecond * 5000)

		tx, err := client.TransactionReceipt(ctx, hash)
		if err == nil && tx != nil {
			ch <- tx
			return
		}
	}
}

func (q *fevmQueries) StateWaitReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	eapi, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer eapi.Close()

	// wait for confirmations (default 5) before getting the receipt
	confirmations := os.Getenv("GLIF_CONFIRMATIONS")
	c, _ := strconv.ParseUint(confirmations, 10, 64)
	if c == 0 {
		c = 5
	}
	err = q.StateWaitForConfirmations(ctx, c)
	if err != nil {
		return nil, err
	}

	ch := make(chan *types.Receipt)
	go q.StateWaitTx(ctx, hash, ch)

	var receipt *types.Receipt
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timed out waiting for transaction")
	case receipt = <-ch:
	}

	if receipt == nil {
		return nil, fmt.Errorf("transaction receipt not found")
	}

	if receipt.Status != 1 {
		return nil, fmt.Errorf("transaction failed: %v", receipt.Status)
	}

	return receipt, nil
}

func (q *fevmQueries) StateWaitNextTick(ctx context.Context, currentHeight *big.Int) error {
	eapi, err := q.extern.ConnectEthClient()
	if err != nil {
		return err
	}
	defer eapi.Close()

	target := currentHeight.Uint64() + 1
	for {
		time.Sleep(time.Millisecond * 5000)

		b, err := eapi.BlockNumber(ctx)
		if err != nil {
			return err
		}

		if b >= target {
			return nil
		}
	}
}

func (q *fevmQueries) StateWaitForConfirmations(ctx context.Context, confirmations uint64) error {
	eapi, err := q.extern.ConnectEthClient()
	if err != nil {
		return err
	}
	defer eapi.Close()

	b, err := eapi.BlockNumber(ctx)
	if err != nil {
		return err
	}

	target := b + confirmations
	for {
		time.Sleep(time.Millisecond * 5000)

		b, err := eapi.BlockNumber(ctx)
		if err != nil {
			return err
		}

		if b >= target {
			return nil
		}
	}
}
