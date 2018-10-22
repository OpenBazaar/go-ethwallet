package wallet

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	wi "github.com/OpenBazaar/wallet-interface"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hunterlong/tokenbalance"

	"github.com/OpenBazaar/go-ethwallet/util"
)

// EthClient represents the eth client
type EthClient struct {
	*ethclient.Client
	url string
}

var txns []wi.Txn

// NewEthClient returns a new eth client
func NewEthClient(url string) (*EthClient, error) {
	var conn *ethclient.Client
	var err error
	if conn, err = ethclient.Dial(url); err != nil {
		return nil, err
	}
	return &EthClient{
		Client: conn,
		url:    url,
	}, nil

}

// Transfer will transfer eth from this user account to dest address
func (client *EthClient) Transfer(from *Account, destAccount common.Address, value *big.Int) (common.Hash, error) {
	var err error
	fromAddress := from.Address()
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	msg := ethereum.CallMsg{From: fromAddress, Value: value}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	rawTx := types.NewTransaction(nonce, destAccount, value, gasLimit, gasPrice, nil)
	signedTx, err := from.SignTransaction(types.HomesteadSigner{}, rawTx)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	txns = append(txns, wi.Txn{
		Txid:      signedTx.Hash().Hex(),
		Value:     value.Int64(),
		Height:    0,
		Timestamp: time.Now(),
		WatchOnly: false,
		Bytes:     rawTx.Data()})

	// this for debug only
	fmt.Println("Txn ID : ", signedTx.Hash().Hex())

	return signedTx.Hash(), client.SendTransaction(context.Background(), signedTx)
}

// TransferToken will transfer erc20 token from this user account to dest address
func (client *EthClient) TransferToken(from *Account, toAddress common.Address, tokenAddress common.Address, value *big.Int) (common.Hash, error) {
	var err error
	fromAddress := from.Address()
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	fmt.Printf("Method ID: %s\n", hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("To address: %s\n", hexutil.Encode(paddedAddress))

	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
	fmt.Printf("Token amount: %s", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	fmt.Printf("Gas limit: %d", gasLimit)

	rawTx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := from.SignTransaction(types.HomesteadSigner{}, rawTx) //types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	txns = append(txns, wi.Txn{
		Txid:      signedTx.Hash().Hex(),
		Value:     value.Int64(),
		Height:    0,
		Timestamp: time.Now(),
		WatchOnly: false,
		Bytes:     rawTx.Data()})

	// this for debug only
	fmt.Println("Txn ID : ", signedTx.Hash().Hex())

	return signedTx.Hash(), client.SendTransaction(context.Background(), signedTx)
}

// GetBalance - returns the balance for this account
func (client *EthClient) GetBalance(destAccount common.Address) (*big.Int, error) {
	return client.BalanceAt(context.Background(), destAccount, nil)
}

// GetTokenBalance - returns the erc20 token balance for this account
func (client *EthClient) GetTokenBalance(destAccount, tokenAddress common.Address) (*big.Int, error) {
	configs := &tokenbalance.Config{
		GethLocation: client.url,
		Logs:         true,
	}
	configs.Connect()

	// insert a Token Contract address and Wallet address
	contract := tokenAddress.String()
	wallet := destAccount.String()

	// query the blockchain and wallet details
	token, err := tokenbalance.New(contract, wallet)
	return token.Balance, err
}

// GetUnconfirmedBalance - returns the unconfirmed balance for this account
func (client *EthClient) GetUnconfirmedBalance(destAccount common.Address) (*big.Int, error) {
	return client.PendingBalanceAt(context.Background(), destAccount)
}

// GetTransaction - returns a eth txn for the specified hash
func (client *EthClient) GetTransaction(hash common.Hash) (*types.Transaction, bool, error) {
	return client.TransactionByHash(context.Background(), hash)
}

// GetLatestBlock - returns the latest block
func (client *EthClient) GetLatestBlock() (uint32, string, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, "", err
	}
	return uint32(header.Number.Int64()), header.Hash().String(), nil
}

// EstimateTxnGas - returns estimated gas
func (client *EthClient) EstimateTxnGas(from, to common.Address, value *big.Int) (*big.Int, error) {
	gas := big.NewInt(0)
	if !(util.IsValidAddress(from.String()) && util.IsValidAddress(to.String())) {
		return gas, errors.New("invalid address")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return gas, err
	}
	msg := ethereum.CallMsg{From: from, To: &to, Value: value}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return gas, err
	}
	return gas.Mul(big.NewInt(int64(gasLimit)), gasPrice), nil
}

// EstimateGasSpend - returns estimated gas
func (client *EthClient) EstimateGasSpend(from common.Address, value *big.Int) (*big.Int, error) {
	gas := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return gas, err
	}
	msg := ethereum.CallMsg{From: from, Value: value}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return gas, err
	}
	return gas.Mul(big.NewInt(int64(gasLimit)), gasPrice), nil
}

/*
func getClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Info("error initializing client")
	}

	return client, err
}
*/

func init() {
	txns = []wi.Txn{}
}
