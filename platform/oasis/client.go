package oasis

import (
	"context"

	"github.com/trustwallet/blockatlas/pkg/logger"

	"github.com/coinbase/rosetta-sdk-go/fetcher"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// Client base struct
type Client struct {
	fetcherClient *fetcher.Fetcher
}

// Init asdasd
func (c *Client) Init() (context.Context, *types.NetworkIdentifier, *types.NetworkStatusResponse, *fetcher.Error) {
	ctx := context.Background()
	primaryNetwork, networkStatus, err := c.fetcherClient.InitializeAsserter(ctx, nil)
	if err != nil {
		logger.Fatal(err)
		return nil, nil, nil, err
	}

	// Step 3: Print the primary network and network status
	logger.Debug("Primary Network: %s\n", types.PrettyPrintStruct(primaryNetwork))
	logger.Debug("Network Status: %s\n", types.PrettyPrintStruct(networkStatus))

	return ctx, primaryNetwork, networkStatus, err
}

// GetCurrentBlock method
func (c *Client) GetCurrentBlock() (int64, *fetcher.Error) {
	_, _, networkStatus, err := c.Init()

	if err != nil {
		return 0, err
	}

	return networkStatus.CurrentBlockIdentifier.Index, nil
}

// GetBlockByNumber method
func (c *Client) GetBlockByNumber(num int64) (*types.Block, *fetcher.Error) {
	ctx, primaryNetwork, _, err := c.Init()

	if err != nil {
		return nil, err
	}

	block, err := c.fetcherClient.BlockRetry(
		ctx,
		primaryNetwork,
		types.ConstructPartialBlockIdentifier(
			&types.BlockIdentifier{Index: num},
		),
	)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	// Step 5: Print the block
	logger.Debug("Current Block: %s\n", types.PrettyPrintStruct(block))

	return block, nil
}
