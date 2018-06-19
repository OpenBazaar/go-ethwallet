package wallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient represents the eth client
type EthClient struct {
	*ethclient.Client
}

// NewEthClient returns a new eth client
func NewEthClient(url string) (*EthClient, error) {
	var conn *ethclient.Client
	var err error
	if conn, err = ethclient.Dial(url); err != nil {
		return nil, err
	}
	return &EthClient{
		Client: conn,
	}, nil

}

// Transfer will transfer eth from this user account to dest address
func (client *EthClient) Transfer(from *Account, destAccount common.Address, value *big.Int) error {
	var err error
	fromAddress := from.Address()
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{From: fromAddress, Value: value}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return err
	}

	rawTx := types.NewTransaction(nonce, destAccount, value, gasLimit, gasPrice, nil)
	signedTx, err := from.SignTransaction(types.HomesteadSigner{}, rawTx)
	if err != nil {
		return err
	}

	// this for debug only
	fmt.Println("Txn ID : ", signedTx.Hash().Hex())

	return client.SendTransaction(context.Background(), signedTx)
}

// GetBalance - returns the balance for this account
func (client *EthClient) GetBalance(destAccount common.Address) (*big.Int, error) {
	return client.BalanceAt(context.Background(), destAccount, nil)
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
