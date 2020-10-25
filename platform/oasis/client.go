package oasis

import (
	"context"

	"github.com/hasura/go-graphql-client"
	"github.com/trustwallet/blockatlas/pkg/logger"

	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

// Client base struct
type Client struct {
	graphqlClient *graphql.Client
}

// GetCurrentBlock method
func (c *Client) GetCurrentBlock() (int64, error) {
	ctx := context.Background()

	var blk LastBlock
	err := c.graphqlClient.Query(ctx, &blk, nil)

	if err != nil {
		return 0, err
	}

	return int64(blk.Blocks[0].Height), nil
}

// GetBlockByNumber method
func (c *Client) GetBlockByNumber(num int64) (*[]Transaction, error) {
	ctx := context.Background()

	var trxs TransactionsByHeight
	variables := map[string]interface{}{
		"height": graphql.Int(num),
	}

	err := c.graphqlClient.Query(ctx, &trxs, variables)
	if err != nil {
		return nil, err
	}

	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return &trxs.Transactions, nil
}

// GetTrxOfAddress method
func (c *Client) GetTrxOfAddress(address string) (*[]Transaction, error) {
	ctx := context.Background()

	var result TransactionsByAddress
	variables := map[string]interface{}{
		"address": graphql.String(address),
	}

	err := c.graphqlClient.Query(ctx, &result, variables)
	if err != nil {
		return nil, err
	}

	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return &result.Transactions, nil
}

// GetAddressesFromXpub method
func (c *Client) GetAddressesFromXpub(xpub string) ([]string, error) {
	var xpubArray signature.PublicKey
	copy(xpubArray[:], xpub)

	address := api.NewAddress(xpubArray)
	addressString := string(address[:])

	addresses := make([]string, 0)
	addresses = append(addresses, addressString)

	return addresses, nil
}
