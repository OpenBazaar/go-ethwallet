package wallet

import (
	"context"
	"crypto/rand"
	"encoding/binary"
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
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
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
	tokenCfg.Options["RegistryAddress"] = "0x403d907982474cdd51687b09a8968346159378f3"
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
	//tscript.Moderator = common.HexToAddress("0xa6Ac51BB2593e833C629A3352c4c432267714385")
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

	orderValue := big.NewInt(345678123478789123)

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

func TestTokenWalletContractApproveEvent(t *testing.T) {
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

	fromAddress := validTokenWallet.account.Address()
	nonce, err := validTokenWallet.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := validTokenWallet.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(validTokenWallet.account.privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 4000000    // in units
	auth.GasPrice = gasPrice

	header, err := validTokenWallet.client.HeaderByNumber(context.Background(), nil)
	fmt.Println("current header no : ", header.Number.Int64())

	var tx *types.Transaction

	tx, err = validTokenWallet.token.Approve(auth, script.MultisigAddress, orderValue)

	if err != nil {
		log.Error(err)
		return
	}

	spew.Dump(tx)

	tclient, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Error(err)
	}

	//tchan := make(chan *TokenApproval)

	validTokenWallet.client, _ = NewEthClient("wss://rinkeby.infura.io/ws") // &EthClient{tclient, ""}

	//var startBlock *uint64
	startBlock := new(uint64) //(*uint64)(unsafe.Pointer(&uint64(0)))

	*startBlock = header.Number.Uint64()

	/*
		wopts := &bind.WatchOpts{
			Start:   startBlock,
			Context: context.Background(),
		}

		sub, err := validTokenWallet.token.WatchApproval(wopts, tchan,
			[]common.Address{validTokenWallet.account.Address()},
			[]common.Address{script.MultisigAddress})
		if err != nil {
			fmt.Println("cannot watch")
			log.Error(err)
			return
		}

		for {
			select {
			case err := <-sub.Err():
				log.Error(err)
				break
			case tlog := <-tchan:
				fmt.Println("yyyyyyyyyyy")
				fmt.Println(tlog)
			}
		}
	*/

	validTokenWallet.token.FilterApproval(nil,
		[]common.Address{validTokenWallet.account.Address()},
		[]common.Address{script.MultisigAddress})

	query := ethereum.FilterQuery{
		Addresses: []common.Address{tscript.TokenAddress},
		FromBlock: header.Number,
		Topics:    [][]common.Hash{{common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")}},
	}
	logs := make(chan types.Log)
	sub1, err := tclient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub1.Unsubscribe()
	flag := false
	for !flag {
		select {
		case err := <-sub1.Err():
			log.Fatal(err)
		case vLog := <-logs:
			//fmt.Println(vLog) // pointer to event log
			//spew.Dump(vLog)
			//fmt.Println(vLog.Topics[0])
			fmt.Println(vLog.Address.String())
			if tx.Hash() == vLog.TxHash {
				fmt.Println("we have found the approval ...")
				spew.Dump(vLog)
				flag = true
				break
			}
		}
	}

}

func TestTokenWalletContractScriptHash(t *testing.T) {
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

	chaincode := []byte("423b5d4c32345ced77393b3530b1eed1")

	tscript.TxnID = common.BytesToAddress(chaincode)

	tscript.MultisigAddress = ver.Implementation

	/*
		fmt.Println("buyer : ", script.Buyer)
		fmt.Println("seller : ", script.Seller)
		fmt.Println("moderator : ", script.Moderator)
		fmt.Println("threshold : ", script.Threshold)
		fmt.Println("timeout : ", script.Timeout)
		fmt.Println("scrptHash : ", shash)
	*/

	spew.Dump(tscript)

	fmt.Println("escrow address : ", ver.Implementation.String())

	smtct, err := NewEscrow(ver.Implementation, validTokenWallet.client)
	if err != nil {
		t.Errorf("error initilaizing contract failed: %s", err.Error())
	}

	retHash, err := smtct.CalculateRedeemScriptHash(nil, tscript.TxnID, tscript.Threshold,
		tscript.Timeout, tscript.Buyer, tscript.Seller, tscript.Moderator, tscript.TokenAddress)

	fmt.Println(err)
	fmt.Println("from smtct : ", retHash)

	rethash1Str := hexutil.Encode(retHash[:])
	fmt.Println("rethash1Str : ", rethash1Str)

	ahash := sha3.NewKeccak256()
	a := make([]byte, 4)
	binary.BigEndian.PutUint32(a, tscript.Timeout)
	arr := append(tscript.TxnID.Bytes(), append([]byte{tscript.Threshold},
		append(a[:], append(tscript.Buyer.Bytes(),
			append(tscript.Seller.Bytes(), append(tscript.Moderator.Bytes(),
				append(tscript.MultisigAddress.Bytes(),
					append(tscript.TokenAddress.Bytes())...)...)...)...)...)...)...)
	ahash.Write(arr)
	ahashStr := hexutil.Encode(ahash.Sum(nil)[:])

	fmt.Println("computed : ", ahashStr)

	if rethash1Str == ahashStr {
		fmt.Println("yay!!!!!!!!!!!!")
	}

	fmt.Println("priv key : ", validTokenWallet.account.privateKey)

	b := []byte{161, 162, 209, 139, 227, 101, 186, 196, 93, 247, 64, 186, 79, 166, 235, 225, 191, 123, 139, 89, 247, 48, 49, 71, 46, 130, 125, 221, 137, 35, 41, 51}

	fmt.Println(hexutil.Encode(b))

	privateKeyBytes := crypto.FromECDSA(validTokenWallet.account.privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	fmt.Println("dest : ", tscript.MultisigAddress.String()[2:])
	fmt.Println("dest : ", string(tscript.MultisigAddress.Bytes()))
	fmt.Println("dest : ", tscript.MultisigAddress.Hex())
	fmt.Println("dest : ", []byte(tscript.MultisigAddress.String())[2:])

	a1, b1, c1 := GenTokenScriptHash(tscript)
	fmt.Println("scrpt hash : ", a1, "   ", b1[2:], "    ", c1)
}
