// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallet

import (
	ethereum "go-ethereum"
	"go-ethereum/accounts/abi"
	"go-ethereum/accounts/abi/bind"
	"go-ethereum/common"
	"go-ethereum/core/types"
	"go-ethereum/event"
	"math/big"
	"strings"
)

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"scriptHash\",\"type\":\"bytes32\"}],\"name\":\"addFundsToTransaction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"name\":\"buyer\",\"type\":\"address\"},{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"lastFunded\",\"type\":\"uint256\"},{\"name\":\"timeoutHours\",\"type\":\"uint32\"},{\"name\":\"threshold\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"},{\"name\":\"_seller\",\"type\":\"address\"},{\"name\":\"_moderators\",\"type\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\"},{\"name\":\"timeoutHours\",\"type\":\"uint32\"},{\"name\":\"scriptHash\",\"type\":\"bytes32\"}],\"name\":\"addTransaction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_partyAddress\",\"type\":\"address\"}],\"name\":\"getAllTransactionsForParty\",\"outputs\":[{\"name\":\"scriptHashes\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"name\":\"sigS\",\"type\":\"bytes32[]\"},{\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"name\":\"uniqueId\",\"type\":\"bytes20\"},{\"name\":\"destinations\",\"type\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"destinations\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"valueAdded\",\"type\":\"uint256\"}],\"name\":\"FundAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Funded\",\"type\":\"event\"}]=======openzeppelin-solidity/contracts/math/SafeMath.sol:SafeMath=======[]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `6080604052600060015534801561001557600080fd5b5061219d806100256000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632d9ef96e1461007d578063642f2eaf146100a1578063b77bf60014610203578063b8310cac1461022e578063be84ceaf146102f2578063c8331de61461038a575b600080fd5b61009f6004803603810190808035600019169060200190929190505050610523565b005b3480156100ad57600080fd5b506100d0600480360381019080803560001916906020019092919050505061071b565b604051808a600019166000191681526020018973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186600181111561015857fe5b60ff168152602001806020018581526020018463ffffffff1663ffffffff1681526020018360ff1660ff168152602001828103825286818151815260200191508051906020019080838360005b838110156101c05780820151818401526020810190506101a5565b50505050905090810190601f1680156101ed5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b34801561020f57600080fd5b5061021861086b565b6040518082815260200191505060405180910390f35b6102f0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803560ff169060200190929190803563ffffffff1690602001909291908035600019169060200190929190505050610871565b005b3480156102fe57600080fd5b50610333600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611076565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561037657808201518184015260208101905061035b565b505050509050019250505060405180910390f35b34801561039657600080fd5b50610521600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803560001916906020019092919080356bffffffffffffffffffffffff191690602001909291908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050611111565b005b60008160008060008360001916600019168152602001908152602001600020600401541415151561055357600080fd5b826000600181111561056157fe5b600080836000191660001916815260200190815260200160002060050160009054906101000a900460ff16600181111561059757fe5b141515610632576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f5472616e73616374696f6e2069732065697468657220696e206469737075746581526020017f206f722072656c6561736564207374617465000000000000000000000000000081525060400191505060405180910390fd5b34925060008311151561064457600080fd5b6106748360008087600019166000191681526020019081526020016000206004015461199490919063ffffffff16565b600080866000191660001916815260200190815260200160002060040181905550426000808660001916600019168152602001908152602001600020600701819055503373ffffffffffffffffffffffffffffffffffffffff167ff66fd2ae9e24a6a24b02e1b5b7512ffde5149a4176085fc0298ae228c9b9d72985856040518083600019166000191681526020018281526020019250505060405180910390a250505050565b60006020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050160009054906101000a900460ff1690806006018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108325780601f1061080757610100808354040283529160200191610832565b820191906000526020600020905b81548152906001019060200180831161081557829003601f168201915b5050505050908060070154908060080160009054906101000a900463ffffffff16908060080160049054906101000a900460ff16905089565b60015481565b6000808260008060008360001916600019168152602001908152602001600020600401541415156108a157600080fd5b34925060008311151561091c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f56616c756520706173736564206973203000000000000000000000000000000081525060200191505060405180910390fd5b600087511115156109bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001807f54686572652073686f756c642062652061746c656173742031206d6f6465726181526020017f746f72000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60028751018660ff1611151515610a60576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f5468726573686f6c642069732067726561746572207468616e20746f74616c2081526020017f6f776e657273000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b61014060405190810160405280856000191681526020018a73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff16815260200188815260200184815260200160006001811115610ac857fe5b8152602001602060405190810160405280600081525081526020014281526020018663ffffffff1681526020018760ff1681525060008086600019166000191681526020019081526020016000206000820151816000019060001916905560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610bd0929190611fff565b506080820151816004015560a08201518160050160006101000a81548160ff02191690836001811115610bff57fe5b021790555060c0820151816006019080519060200190610c20929190612089565b5060e082015181600701556101008201518160080160006101000a81548163ffffffff021916908363ffffffff1602179055506101208201518160080160046101000a81548160ff021916908360ff1602179055509050506001600080866000191660001916815260200190815260200160002060090160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600080866000191660001916815260200190815260200160002060090160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600091505b86518260ff161015610f005760008085600019166000191681526020019081526020016000206009016000888460ff16815181101515610d9f57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515610e66576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d6f64657261746f7220697320626569676e207265706561746564000000000081525060200191505060405180910390fd5b600160008086600019166000191681526020019081526020016000206009016000898560ff16815181101515610e9857fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508180600101925050610d63565b600160008154809291906001019190505550600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020849080600181540180825580915050906001820390600052602060002001600090919290919091509060001916905550600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208490806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055503373ffffffffffffffffffffffffffffffffffffffff167fce7089d0668849fb9ca29778c0cbf1e764d9efb048d81fd71fb34c94f26db368600080876000191660001916815260200190815260200160002060000154856040518083600019166000191681526020018281526020019250505060405180910390a2505050505050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561110557602002820191906000526020600020905b815460001916815260200190600101908083116110ed575b50505050509050919050565b6000806000806000808960008060008360001916600019168152602001908152602001600020600401541415151561114857600080fd5b8a6000600181111561115657fe5b600080836000191660001916815260200190815260200160002060050160009054906101000a900460ff16600181111561118c57fe5b141515611227576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f5472616e73616374696f6e2069732065697468657220696e206469737075746581526020017f206f722072656c6561736564207374617465000000000000000000000000000081525060400191505060405180910390fd5b60008a51118015611239575088518a51145b151561124457600080fd5b6000808d6000191660001916815260200190815260200160002097508a8860080160049054906101000a900460ff168960080160009054906101000a900463ffffffff168a60010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b60020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168c60030160008154811015156112e457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200180876bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014018660ff1660ff167f01000000000000000000000000000000000000000000000000000000000000000281526001018563ffffffff1663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140196505050505050506040516020818303038152906040526040518082805190602001908083835b6020831015156114a85780518252602082019150602081019050602083039250611483565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020965086600019168c6000191614151561157c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260388152602001807f43616c63756c6174656420736372697074206861736820646f6573206e6f742081526020017f6d6174636820706173736564207363726970742068617368000000000000000081525060400191505060405180910390fd5b61158a8f8f8f8f8e8e6119b0565b95506115ae8860080160009054906101000a900463ffffffff168960070154611fb8565b94508760080160049054906101000a900460ff1660ff168f511080156115d2575084155b156115dc57600080fd5b60018f511480156115ea5750845b801561164657508760020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614155b1561165057600080fd5b8760080160049054906101000a900460ff1660ff168f51101561167257600080fd5b60016000808e6000191660001916815260200190815260200160002060050160006101000a81548160ff021916908360018111156116ac57fe5b021790555060009350600092505b89518360ff16101561188657600073ffffffffffffffffffffffffffffffffffffffff168a8460ff168151811015156116ef57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff161415801561179b57506000808d6000191660001916815260200190815260200160002060090160008b8560ff1681518110151561174857fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b15156117a657600080fd5b6000898460ff168151811015156117b957fe5b906020019060200201511115156117cf57600080fd5b6117fc898460ff168151811015156117e357fe5b906020019060200201518561199490919063ffffffff16565b9350898360ff1681518110151561180f57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff166108fc8a8560ff1681518110151561184357fe5b906020019060200201519081150290604051600060405180830381858888f19350505050158015611878573d6000803e3d6000fd5b5082806001019350506116ba565b6000808d600019166000191681526020019081526020016000206004015484111515156118b257600080fd5b7f688e2a1b34445bcd47b0e11ba2a9c8c4d850a1831b64199b59d1c70e297015458c8b8b6040518084600019166000191681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561192b578082015181840152602081019050611910565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561196d578082015181840152602081019050611952565b505050509050019550505050505060405180910390a1505050505050505050505050505050565b600081830190508281101515156119a757fe5b80905092915050565b600080600080875189511480156119c8575089518951145b15156119d357600080fd5b60197f01000000000000000000000000000000000000000000000000000000000000000260007f0100000000000000000000000000000000000000000000000000000000000000023088888b60405160200180877effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101867effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401848051906020019060200280838360005b83811015611b27578082015181840152602081019050611b0c565b50505050905001838051906020019060200280838360005b83811015611b5a578082015181840152602081019050611b3f565b50505050905001826000191660001916815260200196505050505050506040516020818303038152906040526040518082805190602001908083835b602083101515611bbb5780518252602082019150602081019050602083039250611b96565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c0182600019166000191681526020019150506040516020818303038152906040526040518082805190602001908083835b602083101515611c6c5780518252602082019150602081019050602083039250611c47565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250600091505b8851821015611fab576001838b84815181101515611cba57fe5b906020019060200201518b85815181101515611cd257fe5b906020019060200201518b86815181101515611cea57fe5b90602001906020020151604051600081526020016040526040518085600019166000191681526020018460ff1660ff168152602001836000191660001916815260200182600019166000191681526020019450505050506020604051602081039080840390855afa158015611d63573d6000803e3d6000fd5b505050602060405103519050600080886000191660001916815260200190815260200160002060090160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611e4b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f496e76616c6964207369676e617475726500000000000000000000000000000081525060200191505060405180910390fd5b6000808860001916600019168152602001908152602001600020600a0160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515611f28576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f53616d65207369676e61747572652073656e742074776963650000000000000081525060200191505060405180910390fd5b60016000808960001916600019168152602001908152602001600020600a0160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508093508180600101925050611ca0565b5050509695505050505050565b600080611fce8342611fe690919063ffffffff16565b9050610e10840263ffffffff16811191505092915050565b6000828211151515611ff457fe5b818303905092915050565b828054828255906000526020600020908101928215612078579160200282015b828111156120775782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061201f565b5b5090506120859190612109565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106120ca57805160ff19168380011785556120f8565b828001600101855582156120f8579182015b828111156120f75782518255916020019190600101906120dc565b5b509050612105919061214c565b5090565b61214991905b8082111561214557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010161210f565b5090565b90565b61216e91905b8082111561216a576000816000905550600101612152565b5090565b905600a165627a7a7230582028c6a45a5a30ac8978c682848de6ff562f486e5a22b8e9e882164ab016e1f6480029

======= openzeppelin-solidity/contracts/math/SafeMath.sol:SafeMath =======
604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058201f11c3b9646e4f7d728669bbc871ca0f5ca20db4a8a9ff4bb95c01848b9cc2620029`

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// GetAllTransactionsForParty is a free data retrieval call binding the contract method 0xbe84ceaf.
//
// Solidity: function getAllTransactionsForParty(_partyAddress address) constant returns(scriptHashes bytes32[])
func (_Wallet *WalletCaller) GetAllTransactionsForParty(opts *bind.CallOpts, _partyAddress common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "getAllTransactionsForParty", _partyAddress)
	return *ret0, err
}

// GetAllTransactionsForParty is a free data retrieval call binding the contract method 0xbe84ceaf.
//
// Solidity: function getAllTransactionsForParty(_partyAddress address) constant returns(scriptHashes bytes32[])
func (_Wallet *WalletSession) GetAllTransactionsForParty(_partyAddress common.Address) ([][32]byte, error) {
	return _Wallet.Contract.GetAllTransactionsForParty(&_Wallet.CallOpts, _partyAddress)
}

// GetAllTransactionsForParty is a free data retrieval call binding the contract method 0xbe84ceaf.
//
// Solidity: function getAllTransactionsForParty(_partyAddress address) constant returns(scriptHashes bytes32[])
func (_Wallet *WalletCallerSession) GetAllTransactionsForParty(_partyAddress common.Address) ([][32]byte, error) {
	return _Wallet.Contract.GetAllTransactionsForParty(&_Wallet.CallOpts, _partyAddress)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_Wallet *WalletCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_Wallet *WalletSession) TransactionCount() (*big.Int, error) {
	return _Wallet.Contract.TransactionCount(&_Wallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_Wallet *WalletCallerSession) TransactionCount() (*big.Int, error) {
	return _Wallet.Contract.TransactionCount(&_Wallet.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(scriptHash bytes32, buyer address, seller address, value uint256, status uint8, ipfsHash string, lastFunded uint256, timeoutHours uint32, threshold uint8)
func (_Wallet *WalletCaller) Transactions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ScriptHash   [32]byte
	Buyer        common.Address
	Seller       common.Address
	Value        *big.Int
	Status       uint8
	IpfsHash     string
	LastFunded   *big.Int
	TimeoutHours uint32
	Threshold    uint8
}, error) {
	ret := new(struct {
		ScriptHash   [32]byte
		Buyer        common.Address
		Seller       common.Address
		Value        *big.Int
		Status       uint8
		IpfsHash     string
		LastFunded   *big.Int
		TimeoutHours uint32
		Threshold    uint8
	})
	out := ret
	err := _Wallet.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(scriptHash bytes32, buyer address, seller address, value uint256, status uint8, ipfsHash string, lastFunded uint256, timeoutHours uint32, threshold uint8)
func (_Wallet *WalletSession) Transactions(arg0 [32]byte) (struct {
	ScriptHash   [32]byte
	Buyer        common.Address
	Seller       common.Address
	Value        *big.Int
	Status       uint8
	IpfsHash     string
	LastFunded   *big.Int
	TimeoutHours uint32
	Threshold    uint8
}, error) {
	return _Wallet.Contract.Transactions(&_Wallet.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(scriptHash bytes32, buyer address, seller address, value uint256, status uint8, ipfsHash string, lastFunded uint256, timeoutHours uint32, threshold uint8)
func (_Wallet *WalletCallerSession) Transactions(arg0 [32]byte) (struct {
	ScriptHash   [32]byte
	Buyer        common.Address
	Seller       common.Address
	Value        *big.Int
	Status       uint8
	IpfsHash     string
	LastFunded   *big.Int
	TimeoutHours uint32
	Threshold    uint8
}, error) {
	return _Wallet.Contract.Transactions(&_Wallet.CallOpts, arg0)
}

// AddFundsToTransaction is a paid mutator transaction binding the contract method 0x2d9ef96e.
//
// Solidity: function addFundsToTransaction(scriptHash bytes32) returns()
func (_Wallet *WalletTransactor) AddFundsToTransaction(opts *bind.TransactOpts, scriptHash [32]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addFundsToTransaction", scriptHash)
}

// AddFundsToTransaction is a paid mutator transaction binding the contract method 0x2d9ef96e.
//
// Solidity: function addFundsToTransaction(scriptHash bytes32) returns()
func (_Wallet *WalletSession) AddFundsToTransaction(scriptHash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.AddFundsToTransaction(&_Wallet.TransactOpts, scriptHash)
}

// AddFundsToTransaction is a paid mutator transaction binding the contract method 0x2d9ef96e.
//
// Solidity: function addFundsToTransaction(scriptHash bytes32) returns()
func (_Wallet *WalletTransactorSession) AddFundsToTransaction(scriptHash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.AddFundsToTransaction(&_Wallet.TransactOpts, scriptHash)
}

// AddTransaction is a paid mutator transaction binding the contract method 0xb8310cac.
//
// Solidity: function addTransaction(_buyer address, _seller address, _moderators address[], threshold uint8, timeoutHours uint32, scriptHash bytes32) returns()
func (_Wallet *WalletTransactor) AddTransaction(opts *bind.TransactOpts, _buyer common.Address, _seller common.Address, _moderators []common.Address, threshold uint8, timeoutHours uint32, scriptHash [32]byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addTransaction", _buyer, _seller, _moderators, threshold, timeoutHours, scriptHash)
}

// AddTransaction is a paid mutator transaction binding the contract method 0xb8310cac.
//
// Solidity: function addTransaction(_buyer address, _seller address, _moderators address[], threshold uint8, timeoutHours uint32, scriptHash bytes32) returns()
func (_Wallet *WalletSession) AddTransaction(_buyer common.Address, _seller common.Address, _moderators []common.Address, threshold uint8, timeoutHours uint32, scriptHash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.AddTransaction(&_Wallet.TransactOpts, _buyer, _seller, _moderators, threshold, timeoutHours, scriptHash)
}

// AddTransaction is a paid mutator transaction binding the contract method 0xb8310cac.
//
// Solidity: function addTransaction(_buyer address, _seller address, _moderators address[], threshold uint8, timeoutHours uint32, scriptHash bytes32) returns()
func (_Wallet *WalletTransactorSession) AddTransaction(_buyer common.Address, _seller common.Address, _moderators []common.Address, threshold uint8, timeoutHours uint32, scriptHash [32]byte) (*types.Transaction, error) {
	return _Wallet.Contract.AddTransaction(&_Wallet.TransactOpts, _buyer, _seller, _moderators, threshold, timeoutHours, scriptHash)
}

// Execute is a paid mutator transaction binding the contract method 0xc8331de6.
//
// Solidity: function execute(sigV uint8[], sigR bytes32[], sigS bytes32[], scriptHash bytes32, uniqueId bytes20, destinations address[], amounts uint256[]) returns()
func (_Wallet *WalletTransactor) Execute(opts *bind.TransactOpts, sigV []uint8, sigR [][32]byte, sigS [][32]byte, scriptHash [32]byte, uniqueId [20]byte, destinations []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "execute", sigV, sigR, sigS, scriptHash, uniqueId, destinations, amounts)
}

// Execute is a paid mutator transaction binding the contract method 0xc8331de6.
//
// Solidity: function execute(sigV uint8[], sigR bytes32[], sigS bytes32[], scriptHash bytes32, uniqueId bytes20, destinations address[], amounts uint256[]) returns()
func (_Wallet *WalletSession) Execute(sigV []uint8, sigR [][32]byte, sigS [][32]byte, scriptHash [32]byte, uniqueId [20]byte, destinations []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Execute(&_Wallet.TransactOpts, sigV, sigR, sigS, scriptHash, uniqueId, destinations, amounts)
}

// Execute is a paid mutator transaction binding the contract method 0xc8331de6.
//
// Solidity: function execute(sigV uint8[], sigR bytes32[], sigS bytes32[], scriptHash bytes32, uniqueId bytes20, destinations address[], amounts uint256[]) returns()
func (_Wallet *WalletTransactorSession) Execute(sigV []uint8, sigR [][32]byte, sigS [][32]byte, scriptHash [32]byte, uniqueId [20]byte, destinations []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Execute(&_Wallet.TransactOpts, sigV, sigR, sigS, scriptHash, uniqueId, destinations, amounts)
}

// WalletExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Wallet contract.
type WalletExecutedIterator struct {
	Event *WalletExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletExecuted represents a Executed event raised by the Wallet contract.
type WalletExecuted struct {
	ScriptHash   [32]byte
	Destinations []common.Address
	Amounts      []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x688e2a1b34445bcd47b0e11ba2a9c8c4d850a1831b64199b59d1c70e29701545.
//
// Solidity: e Executed(scriptHash bytes32, destinations address[], amounts uint256[])
func (_Wallet *WalletFilterer) FilterExecuted(opts *bind.FilterOpts) (*WalletExecutedIterator, error) {

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &WalletExecutedIterator{contract: _Wallet.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x688e2a1b34445bcd47b0e11ba2a9c8c4d850a1831b64199b59d1c70e29701545.
//
// Solidity: e Executed(scriptHash bytes32, destinations address[], amounts uint256[])
func (_Wallet *WalletFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *WalletExecuted) (event.Subscription, error) {

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletExecuted)
				if err := _Wallet.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletFundAddedIterator is returned from FilterFundAdded and is used to iterate over the raw logs and unpacked data for FundAdded events raised by the Wallet contract.
type WalletFundAddedIterator struct {
	Event *WalletFundAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletFundAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFundAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletFundAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletFundAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFundAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFundAdded represents a FundAdded event raised by the Wallet contract.
type WalletFundAdded struct {
	ScriptHash [32]byte
	From       common.Address
	ValueAdded *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundAdded is a free log retrieval operation binding the contract event 0xf66fd2ae9e24a6a24b02e1b5b7512ffde5149a4176085fc0298ae228c9b9d729.
//
// Solidity: e FundAdded(scriptHash bytes32, from indexed address, valueAdded uint256)
func (_Wallet *WalletFilterer) FilterFundAdded(opts *bind.FilterOpts, from []common.Address) (*WalletFundAddedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "FundAdded", fromRule)
	if err != nil {
		return nil, err
	}
	return &WalletFundAddedIterator{contract: _Wallet.contract, event: "FundAdded", logs: logs, sub: sub}, nil
}

// WatchFundAdded is a free log subscription operation binding the contract event 0xf66fd2ae9e24a6a24b02e1b5b7512ffde5149a4176085fc0298ae228c9b9d729.
//
// Solidity: e FundAdded(scriptHash bytes32, from indexed address, valueAdded uint256)
func (_Wallet *WalletFilterer) WatchFundAdded(opts *bind.WatchOpts, sink chan<- *WalletFundAdded, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "FundAdded", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFundAdded)
				if err := _Wallet.contract.UnpackLog(event, "FundAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WalletFundedIterator is returned from FilterFunded and is used to iterate over the raw logs and unpacked data for Funded events raised by the Wallet contract.
type WalletFundedIterator struct {
	Event *WalletFunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFunded represents a Funded event raised by the Wallet contract.
type WalletFunded struct {
	ScriptHash [32]byte
	From       common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFunded is a free log retrieval operation binding the contract event 0xce7089d0668849fb9ca29778c0cbf1e764d9efb048d81fd71fb34c94f26db368.
//
// Solidity: e Funded(scriptHash bytes32, from indexed address, value uint256)
func (_Wallet *WalletFilterer) FilterFunded(opts *bind.FilterOpts, from []common.Address) (*WalletFundedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "Funded", fromRule)
	if err != nil {
		return nil, err
	}
	return &WalletFundedIterator{contract: _Wallet.contract, event: "Funded", logs: logs, sub: sub}, nil
}

// WatchFunded is a free log subscription operation binding the contract event 0xce7089d0668849fb9ca29778c0cbf1e764d9efb048d81fd71fb34c94f26db368.
//
// Solidity: e Funded(scriptHash bytes32, from indexed address, value uint256)
func (_Wallet *WalletFilterer) WatchFunded(opts *bind.WatchOpts, sink chan<- *WalletFunded, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "Funded", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFunded)
				if err := _Wallet.contract.UnpackLog(event, "Funded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
