package oasis

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// NormalizeTransactions method
func (p *Platform) NormalizeTransactions(trxs *[]Transaction) ([]blockatlas.Tx, error) {

	txs := make([]blockatlas.Tx, 0)
	for _, srcTx := range *trxs {
		tx := p.NormalizeTransaction(srcTx)
		if tx != nil {
			txs = append(txs, *tx)
		}
	}
	return txs, nil
}

// NormalizeTransaction method
func (p *Platform) NormalizeTransaction(trx Transaction) *blockatlas.Tx {

	/*var status blockatlas.Status
	if !srcTx.Success {
		status = blockatlas.StatusError
	} else {
		status = blockatlas.StatusCompleted
	}*/

	result := blockatlas.Tx{
		ID:       string(trx.Tx_Hash),
		Coin:     p.Coin().ID,
		Date:     int64(trx.Timestamp),
		Block:    uint64(trx.Height),
		Status:   blockatlas.StatusCompleted,
		Sequence: uint64(trx.Nonce),
	}

	return &result
}

// NormalizeTxs complete here
func (p *Platform) NormalizeTxs(srcTxs []Transaction) blockatlas.TxPage {
	var txs blockatlas.TxPage
	for _, srcTx := range srcTxs {
		tx := p.NormalizeTx(&srcTx)
		txs = append(txs, tx)
	}
	return txs
}

// NormalizeTx complete here
func (p *Platform) NormalizeTx(srcTx *Transaction) blockatlas.Tx {
	bTx := blockatlas.Tx{
		ID:    string(srcTx.Tx_Hash),
		From:  string(srcTx.From),
		To:    string(srcTx.To),
		Coin:  coin.ROSE,
		Block: uint64(srcTx.Height),
		Fee:   blockatlas.Amount(srcTx.Amount),
		Date:  int64(srcTx.Timestamp),
	}

	return bTx
}
