// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package wallet

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"math/big"
	"strings"
)

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"scriptHash\",\"type\":\"bytes32\"}],\"name\":\"addFundsToTransaction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"name\":\"buyer\",\"type\":\"address\"},{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"ipfsHash\",\"type\":\"string\"},{\"name\":\"lastFunded\",\"type\":\"uint256\"},{\"name\":\"timeoutHours\",\"type\":\"uint32\"},{\"name\":\"threshold\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"},{\"name\":\"_seller\",\"type\":\"address\"},{\"name\":\"_moderators\",\"type\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\"},{\"name\":\"timeoutHours\",\"type\":\"uint32\"},{\"name\":\"scriptHash\",\"type\":\"bytes32\"}],\"name\":\"addTransaction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_partyAddress\",\"type\":\"address\"}],\"name\":\"getAllTransactionsForParty\",\"outputs\":[{\"name\":\"scriptHashes\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"name\":\"sigS\",\"type\":\"bytes32[]\"},{\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"name\":\"uniqueId\",\"type\":\"bytes20\"},{\"name\":\"destinations\",\"type\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"partyVsTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"destinations\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"valueAdded\",\"type\":\"uint256\"}],\"name\":\"FundAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Funded\",\"type\":\"event\"}]=======openzeppelin-solidity/contracts/math/SafeMath.sol:SafeMath=======[]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `6080604052600060015534801561001557600080fd5b50612756806100256000396000f300608060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632d9ef96e14610088578063642f2eaf146100ac578063b77bf6001461020e578063b8310cac14610239578063be84ceaf146102fd578063c8331de614610395578063f5dbe09c1461052e575b600080fd5b6100aa6004803603810190808035600019169060200190929190505050610597565b005b3480156100b857600080fd5b506100db60048036038101908080356000191690602001909291905050506107f8565b604051808a600019166000191681526020018973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186600181111561016357fe5b60ff168152602001806020018581526020018463ffffffff1663ffffffff1681526020018360ff1660ff168152602001828103825286818151815260200191508051906020019080838360005b838110156101cb5780820151818401526020810190506101b0565b50505050905090810190601f1680156101f85780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b34801561021a57600080fd5b50610223610948565b6040518082815260200191505060405180910390f35b6102fb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803560ff169060200190929190803563ffffffff169060200190929190803560001916906020019092919050505061094e565b005b34801561030957600080fd5b5061033e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113b1565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610381578082015181840152602081019050610366565b505050509050019250505060405180910390f35b3480156103a157600080fd5b5061052c600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290803560001916906020019092919080356bffffffffffffffffffffffff19169060200190929190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929050505061144c565b005b34801561053a57600080fd5b50610579600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611ebf565b60405180826000191660001916815260200191505060405180910390f35b600081600080600083600019166000191681526020019081526020016000206004015414151515610630576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f5472616e73616374696f6e20646f6573206e6f7420657869737473000000000081525060200191505060405180910390fd5b826000600181111561063e57fe5b600080836000191660001916815260200190815260200160002060050160009054906101000a900460ff16600181111561067457fe5b14151561070f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f5472616e73616374696f6e2069732065697468657220696e206469737075746581526020017f206f722072656c6561736564207374617465000000000000000000000000000081525060400191505060405180910390fd5b34925060008311151561072157600080fd5b61075183600080876000191660001916815260200190815260200160002060040154611eef90919063ffffffff16565b600080866000191660001916815260200190815260200160002060040181905550426000808660001916600019168152602001908152602001600020600701819055503373ffffffffffffffffffffffffffffffffffffffff167ff66fd2ae9e24a6a24b02e1b5b7512ffde5149a4176085fc0298ae228c9b9d72985856040518083600019166000191681526020018281526020019250505060405180910390a250505050565b60006020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050160009054906101000a900460ff1690806006018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561090f5780601f106108e45761010080835404028352916020019161090f565b820191906000526020600020905b8154815290600101906020018083116108f257829003601f168201915b5050505050908060070154908060080160009054906101000a900463ffffffff16908060080160049054906101000a900460ff16905089565b60015481565b6000808260008060008360001916600019168152602001908152602001600020600401541415156109e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f5472616e73616374696f6e20657869737473000000000000000000000000000081525060200191505060405180910390fd5b88600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610a8d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5a65726f2061646472657373207061737365640000000000000000000000000081525060200191505060405180910390fd5b88600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610b33576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5a65726f2061646472657373207061737365640000000000000000000000000081525060200191505060405180910390fd5b3494508973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff1614151515610bda576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f427579657220616e642073656c6c6572206172652073616d650000000000000081525060200191505060405180910390fd5b600085111515610c52576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f56616c756520706173736564206973203000000000000000000000000000000081525060200191505060405180910390fd5b60028951018860ff1611151515610cf7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f5468726573686f6c642069732067726561746572207468616e20746f74616c2081526020017f6f776e657273000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b61014060405190810160405280876000191681526020018c73ffffffffffffffffffffffffffffffffffffffff1681526020018b73ffffffffffffffffffffffffffffffffffffffff1681526020018a815260200186815260200160006001811115610d5f57fe5b8152602001602060405190810160405280600081525081526020014281526020018863ffffffff1681526020018960ff1681525060008088600019166000191681526020019081526020016000206000820151816000019060001916905560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610e679291906125b8565b506080820151816004015560a08201518160050160006101000a81548160ff02191690836001811115610e9657fe5b021790555060c0820151816006019080519060200190610eb7929190612642565b5060e082015181600701556101008201518160080160006101000a81548163ffffffff021916908363ffffffff1602179055506101208201518160080160046101000a81548160ff021916908360ff1602179055509050506001600080886000191660001916815260200190815260200160002060090160008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600080886000191660001916815260200190815260200160002060090160008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600093505b88518460ff16101561125657600073ffffffffffffffffffffffffffffffffffffffff16898560ff1681518110151561102f57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff16141515156110c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5a65726f2061646472657373207061737365640000000000000000000000000081525060200191505060405180910390fd5b600080876000191660001916815260200190815260200160002060090160008a8660ff168151811015156110f557fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515156111bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d6f64657261746f7220697320626569676e207265706561746564000000000081525060200191505060405180910390fd5b6001600080886000191660001916815260200190815260200160002060090160008b8760ff168151811015156111ee57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508380600101945050610ffa565b600160008154809291906001019190505550600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020869080600181540180825580915050906001820390600052602060002001600090919290919091509060001916905550600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208690806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055503373ffffffffffffffffffffffffffffffffffffffff167fce7089d0668849fb9ca29778c0cbf1e764d9efb048d81fd71fb34c94f26db36887876040518083600019166000191681526020018281526020019250505060405180910390a25050505050505050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561144057602002820191906000526020600020905b81546000191681526020019060010190808311611428575b50505050509050919050565b600080600080600080896000806000836000191660001916815260200190815260200160002060040154141515156114ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f5472616e73616374696f6e20646f6573206e6f7420657869737473000000000081525060200191505060405180910390fd5b8a600060018111156114fa57fe5b600080836000191660001916815260200190815260200160002060050160009054906101000a900460ff16600181111561153057fe5b1415156115cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f5472616e73616374696f6e2069732065697468657220696e206469737075746581526020017f206f722072656c6561736564207374617465000000000000000000000000000081525060400191505060405180910390fd5b60008a511180156115dd575088518a51145b15156115e857600080fd5b6000808d6000191660001916815260200190815260200160002097508a8860080160049054906101000a900460ff168960080160009054906101000a900463ffffffff168a60010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b60020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168c600301600081548110151561168857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200180876bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014018660ff1660ff167f01000000000000000000000000000000000000000000000000000000000000000281526001018563ffffffff1663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140196505050505050506040516020818303038152906040526040518082805190602001908083835b60208310151561184c5780518252602082019150602081019050602083039250611827565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020965086600019168c60001916141515611920576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260388152602001807f43616c63756c6174656420736372697074206861736820646f6573206e6f742081526020017f6d6174636820706173736564207363726970742068617368000000000000000081525060400191505060405180910390fd5b61192e8f8f8f8f8e8e611f0b565b95506119528860080160009054906101000a900463ffffffff168960070154612513565b94508760080160049054906101000a900460ff1660ff168f51108015611976575084155b1561198057600080fd5b60018f5114801561198e5750845b80156119ea57508760020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614155b156119f457600080fd5b8760080160049054906101000a900460ff1660ff168f511015611a1657600080fd5b60016000808e6000191660001916815260200190815260200160002060050160006101000a81548160ff02191690836001811115611a5057fe5b021790555060009350600092505b89518360ff161015611d2257600073ffffffffffffffffffffffffffffffffffffffff168a8460ff16815181101515611a9357fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1614158015611b3f57506000808d6000191660001916815260200190815260200160002060090160008b8560ff16815181101515611aec57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b1515611bb3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4e6f7420612076616c69642064657374696e6174696f6e00000000000000000081525060200191505060405180910390fd5b6000898460ff16815181101515611bc657fe5b90602001906020020151111515611c6b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001807f416d6f756e7420746f2062652073656e742073686f756c64206265206772656181526020017f746572207468616e20300000000000000000000000000000000000000000000081525060400191505060405180910390fd5b611c98898460ff16815181101515611c7f57fe5b9060200190602002015185611eef90919063ffffffff16565b9350898360ff16815181101515611cab57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff166108fc8a8560ff16815181101515611cdf57fe5b906020019060200201519081150290604051600060405180830381858888f19350505050158015611d14573d6000803e3d6000fd5b508280600101935050611a5e565b6000808d60001916600019168152602001908152602001600020600401548411151515611ddd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603c8152602001807f546f74616c2076616c756520746f2062652073656e742069732067726561746581526020017f72207468616e20746865207472616e73616374696f6e2076616c75650000000081525060400191505060405180910390fd5b7f688e2a1b34445bcd47b0e11ba2a9c8c4d850a1831b64199b59d1c70e297015458c8b8b6040518084600019166000191681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611e56578082015181840152602081019050611e3b565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611e98578082015181840152602081019050611e7d565b505050509050019550505050505060405180910390a1505050505050505050505050505050565b600260205281600052604060002081815481101515611eda57fe5b90600052602060002001600091509150505481565b60008183019050828110151515611f0257fe5b80905092915050565b60008060008087518951148015611f23575089518951145b1515611f2e57600080fd5b60197f01000000000000000000000000000000000000000000000000000000000000000260007f0100000000000000000000000000000000000000000000000000000000000000023088888b60405160200180877effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101867effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401848051906020019060200280838360005b83811015612082578082015181840152602081019050612067565b50505050905001838051906020019060200280838360005b838110156120b557808201518184015260208101905061209a565b50505050905001826000191660001916815260200196505050505050506040516020818303038152906040526040518082805190602001908083835b60208310151561211657805182526020820191506020810190506020830392506120f1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c0182600019166000191681526020019150506040516020818303038152906040526040518082805190602001908083835b6020831015156121c757805182526020820191506020810190506020830392506121a2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250600091505b8851821015612506576001838b8481518110151561221557fe5b906020019060200201518b8581518110151561222d57fe5b906020019060200201518b8681518110151561224557fe5b90602001906020020151604051600081526020016040526040518085600019166000191681526020018460ff1660ff168152602001836000191660001916815260200182600019166000191681526020019450505050506020604051602081039080840390855afa1580156122be573d6000803e3d6000fd5b505050602060405103519050600080886000191660001916815260200190815260200160002060090160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156123a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f496e76616c6964207369676e617475726500000000000000000000000000000081525060200191505060405180910390fd5b6000808860001916600019168152602001908152602001600020600a0160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515612483576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f53616d65207369676e61747572652073656e742074776963650000000000000081525060200191505060405180910390fd5b60016000808960001916600019168152602001908152602001600020600a0160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080935081806001019250506121fb565b5050509695505050505050565b600080612529834261256790919063ffffffff16565b905060008463ffffffff161461255b57612554610e108563ffffffff1661258090919063ffffffff16565b811161255e565b60005b91505092915050565b600082821115151561257557fe5b818303905092915050565b60008083141561259357600090506125b2565b81830290508183828115156125a457fe5b041415156125ae57fe5b8090505b92915050565b828054828255906000526020600020908101928215612631579160200282015b828111156126305782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906125d8565b5b50905061263e91906126c2565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061268357805160ff19168380011785556126b1565b828001600101855582156126b1579182015b828111156126b0578251825591602001919060010190612695565b5b5090506126be9190612705565b5090565b61270291905b808211156126fe57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016126c8565b5090565b90565b61272791905b8082111561272357600081600090555060010161270b565b5090565b905600a165627a7a7230582038061cf85e6099ec8664293a74a0ba05752c0ffd988639900a6fff0f2e2a72850029

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
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
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
	contract, err := bindWallet(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// PartyVsTransaction is a free data retrieval call binding the contract method 0xf5dbe09c.
//
// Solidity: function partyVsTransaction( address,  uint256) constant returns(bytes32)
func (_Wallet *WalletCaller) PartyVsTransaction(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "partyVsTransaction", arg0, arg1)
	return *ret0, err
}

// PartyVsTransaction is a free data retrieval call binding the contract method 0xf5dbe09c.
//
// Solidity: function partyVsTransaction( address,  uint256) constant returns(bytes32)
func (_Wallet *WalletSession) PartyVsTransaction(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Wallet.Contract.PartyVsTransaction(&_Wallet.CallOpts, arg0, arg1)
}

// PartyVsTransaction is a free data retrieval call binding the contract method 0xf5dbe09c.
//
// Solidity: function partyVsTransaction( address,  uint256) constant returns(bytes32)
func (_Wallet *WalletCallerSession) PartyVsTransaction(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Wallet.Contract.PartyVsTransaction(&_Wallet.CallOpts, arg0, arg1)
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
