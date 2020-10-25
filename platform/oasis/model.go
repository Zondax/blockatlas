package oasis

import "github.com/hasura/go-graphql-client"

// Block complete
type Block struct {
	Height graphql.Int
	Hash   graphql.String
	Time   graphql.String
}

// LastBlock complete
type LastBlock struct {
	Blocks []Block `graphql:"blocks(limit: 1, order_by: {height: desc})"`
}

// BlockByHeight complete
type BlockByHeight struct {
	Blocks []Block `graphql:"blocks(where: {height: {_eq: $height}})"`
}

// Transaction complete here
type Transaction struct {

	// Height contains the block height.
	Height graphql.Int
	// Nonce is a nonce to prevent replay.
	Nonce graphql.Int
	// Timestamp is the timestamp of trx
	Timestamp graphql.Int
	// Amount is an optional fee that the sender commits to pay to execute this
	// transaction.
	Amount graphql.Int
	// Method is the method that should be called.
	Method graphql.String
	// Gas is the maximum gas that a transaction can use.
	Gas graphql.Int
	// From complete
	From graphql.String
	// To complete
	To graphql.String
	// Tx_Hash complete
	Tx_Hash graphql.String
}

// TransactionsByAddress complete
type TransactionsByAddress struct {
	Transactions []Transaction `graphql:"transactions(where: { _or: {from: {_eq: \"$address\"}, to: {_eq: \"$address\"} } } )"`
}

// TransactionsByHeight complete
type TransactionsByHeight struct {
	Transactions []Transaction `graphql:"transactions(where: { height: {_eq: $height} } )"`
}
