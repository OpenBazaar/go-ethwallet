package wallet

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

const (
	validKeyFile    = "../test/UTC--2018-06-16T18-41-19.615987160Z--c0b4ef9e9d2806f643be94d2434e5c3d5cecd255"
	validPassword   = "hotpotato"
	invalidKeyFile  = "../test/UTC--IDontExist"
	invalidPassword = "lookout"
)

var validEthAddr EthAddress

func setupAddr() {
	addr := common.HexToAddress(validSourceAddress)
	validEthAddr = EthAddress{
		address: &addr,
	}
}

func TestNewAccountWithValidCredentials(t *testing.T) {
	account, err := NewAccount(validKeyFile, validPassword)
	if err != nil {
		t.Errorf("valid keyfile should open properly")
	}
	if account.Address().String() == "" {
		t.Errorf("the account should have a valid address")
	}
	if account.Address().String() != validSourceAddress {
		t.Errorf("the account address is wrong")
	}
}

func TestNewAccountWithInValidKeyFile(t *testing.T) {
	account, err := NewAccount(invalidKeyFile, validPassword)
	if err == nil {
		t.Errorf("invalid keyfile should not open properly")
	}
	if account != nil {
		t.Errorf("the account should not be returned")
	}
}

func TestNewAccountWithInValidPassword(t *testing.T) {
	account, err := NewAccount(validKeyFile, invalidPassword)
	if err == nil {
		t.Errorf("invalid keyfile should not open properly")
	}
	if account != nil {
		t.Errorf("the account should not be returned")
	}
}

func TestNewAccountWithInValidCredentials(t *testing.T) {
	account, err := NewAccount(invalidKeyFile, invalidPassword)
	if err == nil {
		t.Errorf("invalid keyfile should not open properly")
	}
	if account != nil {
		t.Errorf("the account should not be returned")
	}
}

func TestEthAddress(t *testing.T) {
	setupAddr()
	if validEthAddr.String() != validSourceAddress {
		t.Errorf("address not initialized correctly")
	}
	if !bytes.Equal(validEthAddr.ScriptAddress(), []byte{192, 180, 239, 158, 157, 40, 6, 246, 67, 190, 148, 210, 67, 78, 92, 61, 92, 236, 210, 85}) {
		t.Errorf("address not initialized correctly")
	}
	if validEthAddr.EncodeAddress() != validSourceAddress {
		t.Errorf("address not initialized correctly")
	}
}
