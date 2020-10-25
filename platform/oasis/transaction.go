package oasis

import "github.com/trustwallet/blockatlas/pkg/blockatlas"

// GetTxsByAddress complete here
func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	trxs, err := p.client.GetTrxOfAddress(address)
	if err != nil {
		return nil, err
	}

	return p.NormalizeTxs(*trxs), nil
}
