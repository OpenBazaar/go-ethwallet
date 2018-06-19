package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	log "github.com/sirupsen/logrus"

	"github.com/OpenBazaar/go-ethwallet/util"
)

// EthereumWallet is the wallet implementation for ethereum
type EthereumWallet struct {
	client  *EthClient
	account *Account
	address *EthAddress
}

// NewEthereumWallet will return a reference to the Eth Wallet
func NewEthereumWallet(url, keyFile, passwd string) *EthereumWallet {
	client, err := NewEthClient(url)
	if err != nil {
		log.Fatalf("error initializing wallet: %v", err)
	}
	var myAccount *Account
	myAccount, err = NewAccount(keyFile, passwd)
	if err != nil {
		log.Fatalf("key file validation failed: %s", err.Error())
	}
	addr := myAccount.Address()

	return &EthereumWallet{client, myAccount, &EthAddress{&addr}}
}

// GetBalance returns the balance for the wallet
func (wallet *EthereumWallet) GetBalance() (*big.Int, error) {
	return wallet.client.GetBalance(wallet.account.Address())
}

// Transfer will transfer the amount from this wallet to the spec address
func (wallet *EthereumWallet) Transfer(to string, value *big.Int) error {
	toAddress := common.HexToAddress(to)
	return wallet.client.Transfer(wallet.account, toAddress, value)
}

// Start will start the wallet daemon
func (wallet *EthereumWallet) Start() {
	// daemonize the wallet
}

// CurrencyCode returns ETH
func (wallet *EthereumWallet) CurrencyCode() string {
	return "ETH"
}

// IsDust Check if this amount is considered dust - 10000 wei
func (wallet *EthereumWallet) IsDust(amount int64) bool {
	return amount < 10000
}

// CurrentAddress - Get the current address for the given purpose
func (wallet *EthereumWallet) CurrentAddress(purpose KeyPurpose) EthAddress {
	return *wallet.address
}

// NewAddress - Returns a fresh address that has never been returned by this function
func (wallet *EthereumWallet) NewAddress(purpose KeyPurpose) EthAddress {
	return *wallet.address
}

// DecodeAddress - Parse the address string and return an address interface
func (wallet *EthereumWallet) DecodeAddress(addr string) (EthAddress, error) {
	ethAddr := common.HexToAddress(addr)
	if wallet.HasKey(EthAddress{&ethAddr}) {
		return *wallet.address, nil
	}
	return EthAddress{}, errors.New("invalid or unknown address")
}

// HasKey - Returns if the wallet has the key for the given address
func (wallet *EthereumWallet) HasKey(addr EthAddress) bool {
	if !util.IsValidAddress(addr.address.String()) {
		return false
	}
	return wallet.account.Address().String() == addr.address.String()
}

// Balance - Get the confirmed and unconfirmed balances
func (wallet *EthereumWallet) Balance() (confirmed, unconfirmed int64) {
	var balance int64
	bal, err := wallet.GetBalance()
	if err == nil {
		balance = bal.Int64()
	}
	return balance, 0
}

// Close will stop the wallet daemon
func (wallet *EthereumWallet) Close() {
	// stop the wallet daemon
}

// GenWallet creates a wallet
func GenWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e

	fmt.Println(util.IsValidAddress(address))

}

// GenDefaultKeyStore will generate a default keystore
func GenDefaultKeyStore(passwd string) (*Account, error) {
	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(passwd)
	if err != nil {
		return nil, err
	}
	return NewAccount(account.URL.Path, passwd)
}
