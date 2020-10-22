package oasis

import (
	"github.com/hasura/go-graphql-client"
	"github.com/trustwallet/blockatlas/coin"
)

// Platform Required method
type Platform struct {
	client    Client
	CoinIndex uint
}

// Init Required method
func Init(coin uint, api string) *Platform {
	p := &Platform{
		client: Client{graphqlClient: graphql.NewClient(api, nil)},
	}
	return p
}

// Coin Required method
func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ROSE]
}
