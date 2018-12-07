package wallet

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/OpenBazaar/go-ethwallet/util"
)

const (
	validKeyFile       = "../test/UTC--2018-06-16T18-41-19.615987160Z--c0b4ef9e9d2806f643be94d2434e5c3d5cecd255"
	validPassword      = "hotpotato"
	invalidKeyFile     = "../test/UTC--IDontExist"
	invalidPassword    = "lookout"
	mnemonicStr        = "soup arch join universe table nasty fiber solve hotel luggage double clean tell oppose hurry weather isolate decline quick dune song enforce curious menu" // "wolf dragon lion stage rose snow sand snake kingdom hand daring flower foot walk sword"
	mnemonicStrAddress = "0x44Ae1C0955C7ad96700088Fb96906C72102c51E3"
)

var validEthAddr EthAddress

func setupAddr() {
	addr := common.HexToAddress(validSourceAddress)
	validEthAddr = EthAddress{
		address: &addr,
	}
}

func TestNewAccountWithValidCredentials(t *testing.T) {
	account, err := NewAccountFromKeyfile(validKeyFile, validPassword)
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
	account, err := NewAccountFromKeyfile(invalidKeyFile, validPassword)
	if err == nil {
		t.Errorf("invalid keyfile should not open properly")
	}
	if account != nil {
		t.Errorf("the account should not be returned")
	}
}

func TestNewAccountWithInValidPassword(t *testing.T) {
	account, err := NewAccountFromKeyfile(validKeyFile, invalidPassword)
	if err == nil {
		t.Errorf("invalid keyfile should not open properly")
	}
	if account != nil {
		t.Errorf("the account should not be returned")
	}
}

func TestNewAccountWithInValidCredentials(t *testing.T) {
	account, err := NewAccountFromKeyfile(invalidKeyFile, invalidPassword)
	if err == nil {
		t.Errorf("invalid keyfile should not open properly")
	}
	if account != nil {
		t.Errorf("the account should not be returned")
	}
}

func TestNewAccountWithMnemonic(t *testing.T) {
	account, err := NewAccountFromMnemonic(mnemonicStr, "")
	if err != nil {
		t.Errorf("failed to open account from mnemonic: %v", err)
	}
	//fmt.Println(account.key.PrivateKey)
	//fmt.Println(account.Address())
	//fmt.Println("account address : ", account.Address().String())
	if account.Address().String() != mnemonicStrAddress {
		t.Errorf("failed to gen correct address from mnemonic: %v", err)
	}
	if !util.IsValidAddress(account.Address()) {
		t.Errorf("failed to gen valid address from mnemonic: %v", err)
	}
}

func TestEthAddress(t *testing.T) {
	setupAddr()
	//fmt.Println(validEthAddr.String())
	if validEthAddr.String() != validSourceAddress[2:] {
		t.Errorf("address not initialized correctly")
	}
	if !bytes.Equal(validEthAddr.ScriptAddress(), []byte{192, 180, 239, 158, 157, 40, 6, 246, 67, 190, 148, 210, 67, 78, 92, 61, 92, 236, 210, 85}) {
		t.Errorf("address not initialized correctly")
	}
	if validEthAddr.EncodeAddress() != validSourceAddress[2:] {
		t.Errorf("address not initialized correctly")
	}
}
