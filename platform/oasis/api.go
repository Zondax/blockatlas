package oasis

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// NormalizeTransactions method
func (p *Platform) NormalizeTransactions(block *Block) ([]blockatlas.Tx, error) {

	/*txs := make([]blockatlas.Tx, 0)
	for _, srcTx := range block.Transactions {
		tx := p.NormalizeTransaction(srcTx)
		if tx != nil {
			txs = append(txs, *tx)
		}
	}
	return txs, nil*/
	return nil, nil
}

// NormalizeTransaction method
func (p *Platform) NormalizeTransaction(trx *types.Transaction) *blockatlas.Tx {

	/*var status blockatlas.Status
	if !srcTx.Success {
		status = blockatlas.StatusError
	} else {
		status = blockatlas.StatusCompleted
	}

	result := blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     p.Coin().ID,
		Date:     int64(srcTx.Timestamp),
		Block:    srcTx.BlockNumber,
		Status:   status,
		Sequence: srcTx.Nonce,
	}*/

	return &blockatlas.Tx{}
}
