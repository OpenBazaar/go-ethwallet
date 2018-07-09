package wallet

import (
	"fmt"
	"math/big"
	"testing"

	wi "github.com/OpenBazaar/wallet-interface"
)

var validRopstenURL = fmt.Sprintf("https://ropsten.infura.io/%s", validInfuraKey)

var validRopstenWallet *EthereumWallet
var destWallet *EthereumWallet

func setupRopstenWallet() {
	validRopstenWallet = NewEthereumWallet(validRopstenURL, validKeyFile, validPassword)
}

func setupDestWallet() {
	destWallet = NewEthereumWallet(validRopstenURL,
		"../test/UTC--2018-06-16T20-09-33.726552102Z--cecb952de5b23950b15bfd49302d1bdd25f9ee67", validPassword)

}

func TestNewWalletWithValidValues(t *testing.T) {
	wallet := NewEthereumWallet(validRopstenURL, validKeyFile, validPassword)
	if wallet == nil {
		t.Errorf("valid credentials should return a wallet")
	}
	if wallet.address.String() != validSourceAddress {
		t.Errorf("valid credentials should return a wallet with proper initialization")
	}
}

func TestNewWalletWithInValidValues(t *testing.T) {
	t.SkipNow()
	wallet := NewEthereumWallet(validRopstenURL, validKeyFile, invalidPassword)
	if wallet != nil {
		t.Errorf("invalid credentials should return a wallet")
	}
}

func TestWalletGetBalance(t *testing.T) {
	setupRopstenWallet()

	if _, err := validRopstenWallet.GetBalance(); err != nil {
		t.Errorf("valid wallet should return balance")
	}
}

func TestWalletGetUnconfirmedBalance(t *testing.T) {
	setupRopstenWallet()

	if _, err := validRopstenWallet.GetUnconfirmedBalance(); err != nil {
		t.Errorf("valid wallet should return unconfirmed balance")
	}
}

func TestWalletTransfer(t *testing.T) {
	//t.SkipNow()
	setupRopstenWallet()
	setupDestWallet()

	value := big.NewInt(200000)

	sbal1 := big.NewInt(0)
	dbal1 := big.NewInt(0)

	cbal1, _ := validRopstenWallet.GetBalance()
	ucbal1, _ := validRopstenWallet.GetUnconfirmedBalance()

	cbal2, _ := destWallet.GetBalance()
	ucbal2, _ := destWallet.GetUnconfirmedBalance()

	sbal1.Add(cbal1, ucbal1)
	dbal1.Add(cbal2, ucbal2)

	_, err := validRopstenWallet.Transfer(validDestinationAddress, value)

	if err != nil {
		t.Errorf("valid wallet should allow transfer")
	}

	//_, err = chainhash.NewHashFromStr(hash.String())

	//if err != nil {
	//	t.Errorf("wallet should return a valid transaction")
	//}

	//txn, err := validRopstenWallet.GetTransaction(*chash)

	//if err != nil {
	//	t.Errorf("wallet should return a valid transaction")
	//}

	//if txn.Value != value.Int64() {
	//	t.Errorf("wallet is not forming the correct txn")
	//}

	sbal2 := big.NewInt(0)
	dbal2 := big.NewInt(0)

	cbal1, _ = validRopstenWallet.GetBalance()
	ucbal1, _ = validRopstenWallet.GetUnconfirmedBalance()

	cbal2, _ = destWallet.GetBalance()
	ucbal2, _ = destWallet.GetUnconfirmedBalance()

	sbal2.Add(cbal1, ucbal1)
	dbal2.Add(cbal2, ucbal2)

	val := big.NewInt(0)

	val.Sub(dbal2, dbal1)

	if val.Cmp(value) != 0 {
		t.Errorf("client should have transferred balance")
	}

}

func TestWalletCurrencyCode(t *testing.T) {
	setupRopstenWallet()

	if validRopstenWallet.CurrencyCode() != "ETH" {
		t.Errorf("wallet should return proper currency code")
	}
}

func TestWalletIsDust(t *testing.T) {
	setupRopstenWallet()

	if validRopstenWallet.IsDust(int64(10000 + 10000)) {
		t.Errorf("wallet should not indicate wrong dust")
	}

	if !validRopstenWallet.IsDust(int64(10000 - 100)) {
		t.Errorf("wallet should not indicate wrong dust")
	}
}

func TestWalletCurrentAddress(t *testing.T) {
	setupRopstenWallet()

	addr := validRopstenWallet.CurrentAddress(wi.EXTERNAL)

	if addr.String() != validSourceAddress {
		t.Errorf("wallet should return correct current address")
	}
}

func TestWalletNewAddress(t *testing.T) {
	setupRopstenWallet()

	addr := validRopstenWallet.NewAddress(wi.EXTERNAL)

	if addr.String() != validSourceAddress {
		t.Errorf("wallet should return correct new address")
	}
}
