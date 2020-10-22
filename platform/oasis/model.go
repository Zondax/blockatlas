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
	// Amount is an optional fee that the sender commits to pay to execute this
	// transaction.
	Amount graphql.Int
	// Method is the method that should be called.
	Method graphql.String
	// Gas is the maximum gas that a transaction can use.
	Gas graphql.Int
}
