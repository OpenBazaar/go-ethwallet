package wallet

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const invalidInfuraKey = "IAMNOTREAL"
const validSourceAddress = "0xc0B4ef9E9d2806F643be94d2434e5C3d5cEcd255"
const validDestinationAddress = "0xcecb952de5b23950b15bfd49302d1bdd25f9ee67"

const validTxn1 = "0xae818b782ce2d5ef8160de1d022440fdaf92cf91d4cd444eb23c6b6a55240c5b"

var client *EthClient
var err error

var validInfuraKey string
var validurlsTest map[string]bool
var invalidurlsTest map[string]bool
var logicallyinvalidurlsTest map[string]bool
var plainurlsTest map[string]bool

var n uint32
var bal *big.Int
var ropstenURL string

var validRopstenTxn types.Transaction

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	validInfuraKey = InfuraAPIKey // os.Getenv("INFURA_KEY")
	fmt.Println("valid infura key is : ", validInfuraKey)

	ropstenURL = fmt.Sprintf("https://ropsten.infura.io/%s", validInfuraKey)

	validRopstenTxn = *types.NewTransaction(0, common.HexToAddress(validDestinationAddress),
		big.NewInt(3344556677), 53000, big.NewInt(1000000000),
		[]byte("f86780843b9aca0082cf0894cecb952de5b23950b15bfd49302d1bdd25f9ee6784c759e285801ca056dfa0ed4e028d2f2307c421bbe0ebe7516e5fdd140ff080091cb03137a885f8a03197486094c53edb2aaefb6ba5059dd6c82fc709134c93c0c422cb671a139352"))
}

func setup() {

	validurlsTest = map[string]bool{
		fmt.Sprintf("https://ropsten.infura.io/%s", validInfuraKey):   true,
		fmt.Sprintf("https://rinkeby.infura.io/%s", validInfuraKey):   true,
		fmt.Sprintf("https://mainnet.infura.io/%s", validInfuraKey):   true,
		fmt.Sprintf("https://ropsten.infura.io/%s", invalidInfuraKey): false,
		fmt.Sprintf("https://rinkeby.infura.io/%s", invalidInfuraKey): false,
		fmt.Sprintf("https://mainnet.infura.io/%s", invalidInfuraKey): false,
	}

	invalidurlsTest = map[string]bool{
		fmt.Sprintf("innet.infura.io/%s", invalidInfuraKey): false,
		"innet.infura.io/":                                  false,
	}

	logicallyinvalidurlsTest = map[string]bool{
		fmt.Sprintf("https://ropstenFTW.infura.io/%s", validInfuraKey): true,
	}

	plainurlsTest = map[string]bool{
		"https://ropsten.infura.io/": false,
		"https://rinkeby.infura.io/": false,
		"https://mainnet.infura.io/": false,
	}

	n = 0

}

func TestNewClient(t *testing.T) {
	setup()
	t.Parallel()

	if validInfuraKey == "" {
		t.Error("no infura key specified")
	}

	for baseURL := range validurlsTest {
		_, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
	}

	for baseURL := range plainurlsTest {
		_, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
	}

	for baseURL := range logicallyinvalidurlsTest {
		_, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
	}

	for baseURL := range invalidurlsTest {
		_, err = NewEthClient(baseURL)
		if err == nil {
			t.Errorf(baseURL + " client should not have initialized")
		}
	}
}

func TestGetLatestBlock(t *testing.T) {
	setup()
	t.Parallel()

	for baseURL := range validurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		n, _, err = client.GetLatestBlock()
		if err != nil || n <= 0 {
			t.Errorf("client should have fetched block number")
		}
	}

	for baseURL := range plainurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		n, _, err = client.GetLatestBlock()
		if err != nil || n <= 0 {
			t.Errorf("client should have fetched block number")
		}
	}

	for baseURL := range logicallyinvalidurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		n, _, err = client.GetLatestBlock()
		if err == nil || n > 0 {
			t.Errorf("client should not have fetched block number")
		}
	}
}

func TestGetBalance(t *testing.T) {
	setup()
	t.Parallel()

	addr := common.HexToAddress(validSourceAddress)

	for baseURL := range validurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.GetBalance(addr)
		if err != nil || bal == nil {
			t.Errorf("client should have fetched balance")
		}
	}

	for baseURL := range plainurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.GetBalance(addr)
		if err != nil || bal == nil {
			t.Errorf("client should have fetched balance")
		}
	}

	for baseURL := range logicallyinvalidurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.GetBalance(addr)
		if err == nil && bal != nil {
			t.Errorf("client should not have fetched balance")
		}
	}
}

func TestGetUnconfirmedBalance(t *testing.T) {
	setup()
	t.Parallel()

	addr := common.HexToAddress(validSourceAddress)

	for baseURL := range validurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.GetUnconfirmedBalance(addr)
		if err != nil || bal == nil {
			t.Errorf("client should have fetched unconfirmed balance")
		}
	}

	for baseURL := range plainurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.GetUnconfirmedBalance(addr)
		if err != nil || bal == nil {
			t.Errorf("client should have fetched unconfirmed balance")
		}
	}

	for baseURL := range logicallyinvalidurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.GetUnconfirmedBalance(addr)
		if err == nil && bal != nil {
			t.Errorf("client should not have fetched unconfirmed balance")
		}
	}
}

func TestEstimateTxnGas(t *testing.T) {
	setup()
	t.Parallel()

	addr := common.HexToAddress(validSourceAddress)
	dest := common.HexToAddress(validDestinationAddress)
	value := big.NewInt(200000)

	for baseURL := range validurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.EstimateTxnGas(addr, dest, value)
		if err != nil || bal == nil {
			t.Errorf("client should have estimated txn gas")
		}
	}

	for baseURL := range plainurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.EstimateTxnGas(addr, dest, value)
		if err != nil || bal == nil {
			t.Errorf("client should have estimated txn gas")
		}
	}

	for baseURL := range logicallyinvalidurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.EstimateTxnGas(addr, dest, value)
		if err == nil && bal != nil {
			t.Errorf("client should not have estimated txn gas")
		}
	}
}

func TestEstimateGasSpend(t *testing.T) {
	setup()
	t.Parallel()

	addr := common.HexToAddress(validSourceAddress)
	value := big.NewInt(200000)

	for baseURL := range validurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.EstimateGasSpend(addr, value)
		if err != nil || bal == nil {
			t.Errorf("client should have estimated txn gas")
		}
	}

	for baseURL := range plainurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.EstimateGasSpend(addr, value)
		if err != nil || bal == nil {
			t.Errorf("client should have estimated txn gas")
		}
	}

	for baseURL := range logicallyinvalidurlsTest {
		client, err = NewEthClient(baseURL)
		if err != nil {
			t.Errorf("client should have initialized")
		}
		bal, err = client.EstimateGasSpend(addr, value)
		if err == nil && bal != nil {
			t.Errorf("client should not have estimated txn gas")
		}
	}
}

func TestValidGetTransaction(t *testing.T) {
	t.Parallel()

	client, err := NewEthClient(ropstenURL)
	if err != nil {
		t.Errorf("client should have initialized")
	}
	txn, isPending, err := client.GetTransaction(common.HexToHash(validTxn1))
	spew.Println(txn)

	if err != nil {
		t.Errorf("txn should have been correctly fetched")
	}

	if isPending {
		t.Errorf("non pending txn should have been correctly fetched")
	}

	if validRopstenTxn.Nonce() != txn.Nonce() {
		t.Errorf("txn should have been correctly fetched")
	}

	if validRopstenTxn.Gas() != txn.Gas() {
		t.Errorf("txn should have been correctly fetched")
	}

	if validRopstenTxn.GasPrice().Cmp(txn.GasPrice()) != 0 {
		t.Errorf("txn should have been correctly fetched")
	}

}

func TestInvalidGetTransaction(t *testing.T) {
	t.Parallel()

	client, err := NewEthClient(ropstenURL)
	if err != nil {
		t.Errorf("client should have initialized")
	}
	txn, _, err := client.GetTransaction(common.HexToHash("ooommm"))
	if err == nil || txn != nil {
		t.Errorf("invalid txn should not have been fetched")
	}
}

func TestTransfer(t *testing.T) {
	//t.SkipNow()
	client, err := NewEthClient(ropstenURL)
	if err != nil {
		t.Errorf("client should have initialized")
	}
	addr := common.HexToAddress(validSourceAddress)
	account, err := NewAccountFromKeyfile("../test/UTC--2018-06-16T18-41-19.615987160Z--c0b4ef9e9d2806f643be94d2434e5c3d5cecd255", "hotpotato")
	if err != nil {
		t.Errorf("account should have initialized")
	}
	dest := common.HexToAddress(validDestinationAddress)
	value := big.NewInt(200000)

	var bal1, bal2, sbal1, dbal1, sbal2, dbal2, val *big.Int
	var txn *types.Transaction

	bal1, err = client.GetBalance(addr)
	if err != nil || bal1 == nil {
		t.Errorf("client should have fetched balance")
	}

	bal2, err = client.GetUnconfirmedBalance(addr)
	if err != nil || bal2 == nil {
		t.Errorf("client should have fetched balance")
	}

	sbal1 = big.NewInt(0)

	// get the source balance
	sbal1.Add(bal1, bal2)

	bal1, err = client.GetBalance(dest)
	if err != nil {
		t.Errorf("client should have fetched balance")
	}

	bal2, err = client.GetUnconfirmedBalance(dest)
	if err != nil {
		t.Errorf("client should have fetched balance")
	}

	dbal1 = big.NewInt(0)

	// get the dest balance
	dbal1.Add(bal1, bal2)

	hash, err := client.Transfer(account, dest, value)
	if err != nil {
		t.Errorf("client should transfer : %v", err)
	}

	txn, _, err = client.GetTransaction(hash)
	if err != nil {
		t.Errorf("txn should have been correctly fetched")
	}

	if txn.Value().Cmp(value) != 0 {
		t.Errorf("client should transfer correct amount")
	}

	bal1, err = client.GetBalance(addr)
	if err != nil {
		t.Errorf("client should have fetched balance")
	}

	bal2, err = client.GetUnconfirmedBalance(addr)
	if err != nil {
		t.Errorf("client should have fetched balance")
	}

	sbal2 = big.NewInt(0)

	// get the source balance
	sbal2.Add(bal1, bal2)

	bal1, err = client.GetBalance(dest)
	if err != nil {
		t.Errorf("client should have fetched balance")
	}

	bal2, err = client.GetUnconfirmedBalance(dest)
	if err != nil {
		t.Errorf("client should have fetched balance")
	}

	dbal2 = big.NewInt(0)
	val = big.NewInt(0)

	// get the dest balance
	dbal2.Add(bal1, bal2)

	val.Sub(dbal2, dbal1)

	if val.Cmp(value) != 0 {
		t.Errorf("client should have transferred balance")
	}
}
