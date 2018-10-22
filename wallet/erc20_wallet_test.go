package wallet

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/OpenBazaar/multiwallet/config"
	wi "github.com/OpenBazaar/wallet-interface"
)

var validTokenWallet *ERC20Wallet
var destTokenWallet *ERC20Wallet

var tokenCfg config.CoinConfig

func setupTokenConfigRinkeby() {
	clientURL, _ := url.Parse("https://rinkeby.infura.io")
	tokenCfg.ClientAPI = *clientURL
	tokenCfg.CoinType = wi.Ethereum
	tokenCfg.Options = make(map[string]interface{})
	tokenCfg.Options["RegistryAddress"] = "0xab8dd0e05b73529b440d9c9df00b5f490c8596ff"
	tokenCfg.Options["Name"] = "OBToken"
	tokenCfg.Options["Symbol"] = "OBT"
	tokenCfg.Options["MainNetAddress"] = "0xe46ea07736e68df951df7b987dda453962ba7d5a"
}

func setupSourceErc20Wallet() {
	setupTokenConfigRinkeby()
	validTokenWallet, _ = NewERC20Wallet(tokenCfg, mnemonicStr, nil)
}

func TestNewErc20WalletWithValidCoinConfigValues(t *testing.T) {
	setupTokenConfigRinkeby()
	wallet, err := NewERC20Wallet(tokenCfg, mnemonicStr, nil)
	if err != nil || wallet == nil {
		t.Errorf("valid credentials should return a wallet")
	}
	fmt.Println(wallet.address.String())
	fmt.Println(validSourceAddress)
	if wallet.address.String() != mnemonicStrAddress {
		t.Errorf("valid credentials should return a wallet with proper initialization")
	}
}

func TestWalletGetTokenBalance(t *testing.T) {
	setupSourceErc20Wallet()

	if _, err := validTokenWallet.GetBalance(); err != nil {
		t.Errorf("valid wallet should return balance")
	}
}

func TestWalletGetTokenUnconfirmedBalance(t *testing.T) {
	setupSourceErc20Wallet()

	if _, err := validTokenWallet.GetUnconfirmedBalance(); err != nil {
		t.Errorf("valid wallet should return unconfirmed balance")
	}
}

func TestTokenWalletCurrencyCode(t *testing.T) {
	setupSourceErc20Wallet()

	if validSampleWallet.CurrencyCode() != "OBT" {
		t.Errorf("wallet should return proper currency code")
	}
}

func TestTokenWalletIsDust(t *testing.T) {
	setupSourceErc20Wallet()

	if validTokenWallet.IsDust(int64(10000 + 10000)) {
		t.Errorf("wallet should not indicate wrong dust")
	}

	if !validTokenWallet.IsDust(int64(10000 - 100)) {
		t.Errorf("wallet should not indicate wrong dust")
	}
}

func TestTokenWalletCurrentAddress(t *testing.T) {
	setupSourceErc20Wallet()

	addr := validTokenWallet.CurrentAddress(wi.EXTERNAL)

	if addr.String() != mnemonicStrAddress {
		t.Errorf("wallet should return correct current address")
	}
}

func TestTokenWalletNewAddress(t *testing.T) {
	setupSourceErc20Wallet()

	addr := validTokenWallet.NewAddress(wi.EXTERNAL)

	if addr.String() != mnemonicStrAddress {
		t.Errorf("wallet should return correct new address")
	}
}
