package oasis

import (
	"github.com/coinbase/rosetta-sdk-go/fetcher"
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
		client: Client{fetcherClient: fetcher.New(api)},
	}
	return p
}

// Coin Required method
func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ROSE]
}
