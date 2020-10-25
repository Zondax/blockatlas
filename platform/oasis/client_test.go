package oasis

import (
	"os"
	"testing"

	"github.com/hasura/go-graphql-client"
)

const (
	EnvNameGraphQlURL = "GRAPHQL_URL"
)

const (
	inputPk           = ""
	expectedAddress   = ""
	graphqlDefaultURL = "http://127.0.0.1:8080/v1/graphql"
)

func TestGetCurrentBlock(t *testing.T) {

	url := os.Getenv(EnvNameGraphQlURL)
	if url == "" {
		url = graphqlDefaultURL
	}

	client := Client{graphqlClient: graphql.NewClient(url, nil)}

	height, err := client.GetCurrentBlock()

	if err != nil {
		t.Errorf("The block %d was not fetched: %v \n", height, err)
	}

	t.Logf("The fetched height block is %d", height)

}

func TestGetBlockByNumber(t *testing.T) {

	url := os.Getenv(EnvNameGraphQlURL)
	if url == "" {
		url = graphqlDefaultURL
	}

	height := int64(1)
	client := Client{graphqlClient: graphql.NewClient(url, nil)}

	trxs, err := client.GetBlockByNumber(height)

	if err != nil {
		t.Errorf("The block %d was not fetched: %v \n", height, err)
	}

	for i, trx := range *trxs {
		if int64(trx.Height) != height {
			t.Errorf("The trx %d fetched is not the correct one.", i)
		}
	}

}

func TestGetTrxOfAddress(t *testing.T) {

	url := os.Getenv(EnvNameGraphQlURL)
	if url == "" {
		url = graphqlDefaultURL
	}

	address := "oasis1qzp84num6xgspdst65yv7yqegln6ndcxmuuq8s9w"
	client := Client{graphqlClient: graphql.NewClient(url, nil)}

	trxs, err := client.GetTrxOfAddress(address)

	if err != nil {
		t.Errorf("The transactions with address %s were not fetched: %v \n", address, err)
	}

	for _, trx := range *trxs {
		if trx.From != graphql.String(address) && trx.To != graphql.String(address) {
			t.Error("The block fetched is not the correct one\n")
		}
	}

}

func TestGetAddressesFromXpub(t *testing.T) {

	client := Client{}
	addresses, err1 := client.GetAddressesFromXpub(inputPk)

	if err1 != nil {
		t.Errorf("The address value was not generated: %v \n", err1)
	}

	if addresses[0] != expectedAddress {
		t.Errorf("The address value [%v] is not equal to the expected one [%v] \n", addresses[0], expectedAddress)
	}
}
