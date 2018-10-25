package wallet

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"
	"testing"
	"time"

	"github.com/OpenBazaar/multiwallet/config"
	wi "github.com/OpenBazaar/wallet-interface"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	hd "github.com/btcsuite/btcutil/hdkeychain"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"

	"github.com/OpenBazaar/go-ethwallet/util"
)

var validTokenWallet *ERC20Wallet
var destTokenWallet *ERC20Wallet

var tokenCfg config.CoinConfig
var tscript EthRedeemScript

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

func setupERCTokenRedeemScript(timeout time.Duration, threshold int) {

	chaincode := make([]byte, 32)
	_, err := rand.Read(chaincode)
	fmt.Println("chiancode : ", chaincode)
	if err != nil {
		fmt.Println(err)
		chaincode = []byte("423b5d4c32345ced77393b3530b1eed1")
	}
	tscript.TxnID = common.BytesToAddress(chaincode)
	tscript.Timeout = uint32(timeout.Hours())
	tscript.Threshold = uint8(threshold)
	tscript.Buyer = common.HexToAddress(mnemonicStrAddress)
	tscript.Seller = common.HexToAddress(validDestinationAddress)
	tscript.Moderator = common.BigToAddress(big.NewInt(0))
	tscript.MultisigAddress = common.HexToAddress("0x36e19e91DFFCA4251f4fB541f5c3a596252eA4BB")
	tscript.TokenAddress = common.HexToAddress("0xe46ea07736e68df951df7b987dda453962ba7d5a")

	//fmt.Println("in setup tscript: ")
	//spew.Dump(tscript)
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

	if validTokenWallet.CurrencyCode() != "OBT" {
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

func TestTokenWalletContractAddTransaction(t *testing.T) {
	setupSourceErc20Wallet()

	ver, err := validTokenWallet.registry.GetRecommendedVersion(nil, "escrow")
	if err != nil {
		t.Error("error fetching escrow from registry")
	}

	if util.IsZeroAddress(ver.Implementation) {
		log.Infof("escrow not available")
		return
	}

	d, _ := time.ParseDuration("1h")
	setupERCTokenRedeemScript(d, 1)

	tscript.MultisigAddress = ver.Implementation

	redeemScript, err := SerializeEthScript(tscript)
	if err != nil {
		t.Error("error serializing redeem script")
	}

	fmt.Println(redeemScript)

	spew.Dump(tscript)

	orderValue := big.NewInt(34567812347878)

	hash, err := validTokenWallet.callAddTokenTransaction(tscript, orderValue)

	fmt.Println("returned hash : ", hash)
	fmt.Println(err)

	chash, err := chainhash.NewHashFromStr(hash.String())

	fmt.Println("err : ", err)

	if err == nil {
		txn, err := validTokenWallet.GetTransaction(*chash)

		spew.Dump(txn)
		fmt.Println(err)
	}

	output := wi.TransactionOutput{
		Address: EthAddress{&tscript.Seller},
		Value:   orderValue.Int64(),
		Index:   1,
	}

	hkey := hd.NewExtendedKey([]byte{}, []byte{}, []byte{}, []byte{}, 0, 0, false)

	sig, err := validTokenWallet.CreateMultisigSignature([]wi.TransactionInput{}, []wi.TransactionOutput{output},
		hkey, redeemScript, 2000)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sig)

	time.Sleep(5 * time.Minute)

	txBytes, err := validTokenWallet.Multisign([]wi.TransactionInput{},
		[]wi.TransactionOutput{output},
		sig, []wi.Signature{wi.Signature{InputIndex: 1, Signature: []byte{}}}, redeemScript,
		20000, true)
	//fmt.Println("after multisign")
	//fmt.Println(txBytes)
	fmt.Println("err : ", err)

	mtx := &types.Transaction{}

	mtx.UnmarshalJSON(txBytes)

	spew.Dump(mtx)

	sshh, sshhstr, _ := GenTokenScriptHash(tscript)

	fmt.Println("script hash for ct : ", sshh)
	fmt.Println(sshhstr)

}
