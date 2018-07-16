package wallet

import (
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// EthAddress implements the WalletAddress interface
type EthAddress struct {
	address *common.Address
}

// String representation of eth address
func (addr EthAddress) String() string {
	return addr.address.String()
}

// EncodeAddress returns hex representation of the address
func (addr EthAddress) EncodeAddress() string {
	return addr.address.Hex()
}

// ScriptAddress returns byte representation of address
func (addr EthAddress) ScriptAddress() []byte {
	return addr.address.Bytes()
}

// Account represents ethereum keystore
type Account struct {
	key *keystore.Key
}

// NewAccount returns the account imported
func NewAccount(keyFile, password string) (*Account, error) {
	key, err := importKey(keyFile, password)
	if err != nil {
		return nil, err
	}

	return &Account{
		key: key,
	}, nil
}

func importKey(keyFile, password string) (*keystore.Key, error) {
	f, err := os.Open(keyFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	json, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return keystore.DecryptKey(json, password)
}

// Address returns the eth address
func (account *Account) Address() common.Address {
	return account.key.Address
}

// SignTransaction will sign the txn
func (account *Account) SignTransaction(signer types.Signer, tx *types.Transaction) (*types.Transaction, error) {
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), account.key.PrivateKey)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(signer, signature)
}
