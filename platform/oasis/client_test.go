package oasis

import "testing"

const (
	inputPk         = ""
	expectedAddress = ""
)

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
