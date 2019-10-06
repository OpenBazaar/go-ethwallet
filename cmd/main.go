package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"ethwallet/util"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/btcsuite/btcd/btcec"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/davecgh/go-spew/spew"
	bip39 "github.com/tyler-smith/go-bip39"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

var (
	password string
	keyDir   string
	keyFile  string
)

const ethToWei = 1 << 17

// InfuraRopstenBase is the base URL for Infura Ropsten network
const InfuraRopstenBase string = "https://ropsten.infura.io/"

func init() {
	flag.StringVar(&password, "p", "", "password for keystore")
	flag.StringVar(&keyDir, "d", "", "key dir to generate key")
	flag.StringVar(&keyFile, "f", "", "key file path")
}

func main() {
	fmt.Println(os.Getenv("INFURA_KEY"))

	//ropstenURL := InfuraRopstenBase + os.Getenv("INFURA_KEY")

	flag.Parse()

	fmt.Println("Password is : ", password)
	fmt.Println("keydir is: ", keyDir)
	fmt.Println("keyfile is : ", keyFile)

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client

	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Println(address.Hex())
	// 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(address.Hash().Hex())
	// 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes())

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// Get the balance at a particular instance of time expressed as block number
	blockNumber := big.NewInt(5532993)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	//wallet.GenWallet()

	//wallet.GenDefaultKeyStore(password)
	//var myAccount *wallet.Account
	//myAccount, err = wallet.NewAccountFromKeyfile(keyFile, password)
	//if err != nil {
	//	log.Fatalf("key file validation failed: %v", err.Error())
	//}
	//fmt.Println(myAccount.Address().String())

	// create the source wallet obj for Infura Ropsten
	//myWallet := wallet.NewEthereumWalletWithKeyfile(ropstenURL, keyFile, password)
	//fmt.Println(myWallet.GetBalance())

	// create dest account
	//wallet.GenDefaultKeyStore(password)
	//var destAccount *wallet.Account
	//destKeyFile := "./UTC--2018-06-16T20-09-33.726552102Z--cecb952de5b23950b15bfd49302d1bdd25f9ee67"
	//destAccount, err = wallet.NewAccountFromKeyfile(destKeyFile, password)
	//if err != nil {
	//	log.Fatalf("key file validation failed: %s", err.Error())
	//}
	//fmt.Println(destAccount.Address().String())

	// create the destination wallet obj for Infura Ropsten
	//destWallet := wallet.NewEthereumWalletWithKeyfile(ropstenURL, destKeyFile, password)
	//fmt.Println(destWallet.GetBalance())

	// lets transfer
	//err = myWallet.Transfer(destAccount.Address().String(), big.NewInt(3344556677))
	//if err != nil {
	//	fmt.Println("what happened here : ", err)
	//}
	//fmt.Println("after transfer : ")
	//fmt.Println("Dest balance ")
	//fmt.Println(destWallet.GetBalance())
	//fmt.Println(destWallet.Balance())
	//fmt.Println("Source balance ")
	//fmt.Println(myWallet.GetBalance())
	//fmt.Println(myWallet.Balance())

	//fmt.Println(myWallet.CreateAddress())

	mnemonic := "soup arch join universe table nasty fiber solve hotel luggage double clean tell oppose hurry weather isolate decline quick dune song enforce curious menu" // "wolf dragon lion stage rose snow sand snake kingdom hand daring flower foot walk sword"
	//mnemonicStrAddress := "0x44Ae1C0955C7ad96700088Fb96906C72102c51E3"
	password := ""

	seed := bip39.NewSeed(mnemonic, password)

	var params chaincfg.Params

	params = chaincfg.MainNetParams

	mPrivKey, err := hdkeychain.NewMaster(seed, &params)
	if err != nil {
		fmt.Println("err   : ", err)
	}

	fmt.Println("priv key ....", mPrivKey)

	spew.Dump(mPrivKey)

	mPubKey, err := mPrivKey.Neuter()
	if err != nil {
		fmt.Println("err   : ", err)
	}

	fmt.Println("pub key ....", mPubKey)

	//spew.Dump(mPubKey)

	ecPubKey, err := mPrivKey.ECPubKey()
	if err != nil {
		fmt.Println("err   : ", err)
	}
	fmt.Println("ecPubkey  : ", ecPubKey)

	keysBitcoin := ecPubKey.SerializeCompressed()

	fmt.Println("keysBitcoin :", hex.EncodeToString(keysBitcoin))

	//spew.Dump(ecPubKey)

	pubKey := ecPubKey.ToECDSA()

	fmt.Println("btcec pubkey : ", pubKey)

	addrBTC := crypto.PubkeyToAddress(*pubKey)

	fmt.Println("addrBTC : ", addrBTC.String())

	fmt.Println("is valid address : ", util.IsValidAddress(addrBTC))

	keysBitcoinStr := hex.EncodeToString(keysBitcoin)

	keysBitcoinBytes, err := hex.DecodeString(keysBitcoinStr)
	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Println("are they equal : ", bytes.Equal(keysBitcoinBytes, keysBitcoin))

	dKey, err := btcec.ParsePubKey(keysBitcoinBytes, btcec.S256())
	if err != nil {
		fmt.Println("err : ", err)
	}

	dpubKey := dKey.ToECDSA()

	fmt.Println("btcec dpubkey : ", dpubKey)

	addrBTCdd := crypto.PubkeyToAddress(*dpubKey)

	fmt.Println("addrBTCdd : ", addrBTCdd.String())

	fmt.Println("is valid address dd : ", util.IsValidAddress(addrBTCdd))

	privateKeyECDSA, err := crypto.ToECDSA(seed[:32])
	if err != nil {
		fmt.Println("err   : ", err)
	}

	fmt.Println("ecdsa prv key : ", privateKeyECDSA)

	//spew.Dump(privateKeyECDSA)

	epubkey := privateKeyECDSA.PublicKey

	fmt.Println("epubkey  : ", epubkey)

	addr := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

	fmt.Println("addr : ", addr.String())

	spew.Dump("fin")

	var masterKey = []byte("Bitcoin seed")
	HDPrivateKeyID := [4]byte{0x04, 0x88, 0xad, 0xe4}

	// First take the HMAC-SHA512 of the master key and the seed data:
	//   I = HMAC-SHA512(Key = "Bitcoin seed", Data = S)
	hmac512 := hmac.New(sha512.New, masterKey)
	hmac512.Write(seed)
	lr := hmac512.Sum(nil)

	// Split "I" into two 32-byte sequences Il and Ir where:
	//   Il = master secret key
	//   Ir = master chain code
	secretKey := lr[:len(lr)/2]
	chainCode := lr[len(lr)/2:]

	parentFP := []byte{0x00, 0x00, 0x00, 0x00}
	exKeyEth := hdkeychain.NewExtendedKey(HDPrivateKeyID[:], secretKey, chainCode,
		parentFP, 0, 0, true)

	fmt.Println(exKeyEth)

	exPrivKey, err := exKeyEth.ECPrivKey()
	if err != nil {
		fmt.Println("err   : ", err)
	}

	ethPrivkey := exPrivKey.ToECDSA()

	fmt.Println("ethPrivkey : ", ethPrivkey)

	ethPubkey := ethPrivkey.PublicKey

	fmt.Println("eth pubkey : ", ethPubkey)

	addrEth := crypto.PubkeyToAddress(ethPubkey)

	fmt.Println("addrEth : ", addrEth.String())

	b1 := addrEth.Bytes()
	b2 := []byte(addrEth.String())

	fmt.Println(bytes.Equal(b1, b2))

	fmt.Println("b1 : ", b1)
	fmt.Println("b2 : ", b2)
	fmt.Println("addr len : ", common.AddressLength)

	b3 := []byte("this is a aasimpleers string which is great")
	fmt.Println(b3, "  ", len(b3))

	b4 := []byte("this is a aasimpleer")
	fmt.Println(b4, "  ", len(b4))

	fmt.Println(bytes.Equal(b4, b3[:common.AddressLength]))
	fmt.Println("key : ", b3[:common.AddressLength])
	fmt.Println("rem : ", b3[common.AddressLength:])

	hash := common.HexToHash("0x22365e53ce13cb11329fe042f37effa8d041e9a26ab8f9aa6ff02520097cd957")

	fmt.Println(hash.String())
	fmt.Println(hash.Hex())
	fmt.Println(hash.Bytes())
	fmt.Println(string(hash.Bytes()))
	fmt.Println(hexutil.Encode(hash.Bytes()))

}
