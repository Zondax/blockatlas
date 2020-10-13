package oasis

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// CurrentBlockNumber Method
func (p *Platform) CurrentBlockNumber() (int64, error) {
	block, err := p.client.GetCurrentBlock()
	if err != nil {
		return 0, err.Err
	}
	return block, nil
}

// GetBlockByNumber Method
func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcBlock, err1 := p.client.GetBlockByNumber(num)
	if err1 != nil {
		return nil, err1.Err
	}

	txs, err2 := p.NormalizeTransactions(srcBlock)
	if err2 != nil {
		return nil, err2
	}
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
