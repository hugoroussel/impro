// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ImproABI is the input ABI used to generate the binding from.
const ImproABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"check\",\"outputs\":[{\"name\":\"exist_at\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ImproBin is the compiled bytecode used for deploying new contracts.
const ImproBin = `0x608060405234801561001057600080fd5b5061077a806100206000396000f3fe60806040526004361061003f5760003560e01c8063492cc76914610044578063b6f921ad146100ea578063ea87152b146101ad578063fbf58b3e14610274575b600080fd5b6100e86004803603602081101561005a57600080fd5b810190602081018135600160201b81111561007457600080fd5b82018360208201111561008657600080fd5b803590602001918460018302840111600160201b831117156100a757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610330945050505050565b005b3480156100f657600080fd5b5061019b6004803603602081101561010d57600080fd5b810190602081018135600160201b81111561012757600080fd5b82018360208201111561013957600080fd5b803590602001918460018302840111600160201b8311171561015a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103b8945050505050565b60408051918252519081900360200190f35b3480156101b957600080fd5b50610260600480360360408110156101d057600080fd5b810190602081018135600160201b8111156101ea57600080fd5b8201836020820111156101fc57600080fd5b803590602001918460018302840111600160201b8311171561021d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610531915050565b604080519115158252519081900360200190f35b34801561028057600080fd5b506100e86004803603604081101561029757600080fd5b810190602081018135600160201b8111156102b157600080fd5b8201836020820111156102c357600080fd5b803590602001918460018302840111600160201b831117156102e457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506106209050565b600061033b826103b8565b905080600019141561034957fe5b346000828154811061035757fe5b906000526020600020906004020160030154111561037157fe5b336000828154811061037f57fe5b906000526020600020906004020160020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505050565b6000600019815b60005481101561052a57836040516020018080602001828103825283818151815260200191508051906020019080838360005b8381101561040a5781810151838201526020016103f2565b50505050905090810190601f1680156104375780820380516001836020036101000a031916815260200191505b5092505050604051602081830303815290604052805190602001206000828154811061045f57fe5b906000526020600020906004020160010160405160200180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156104fc5780601f106104d1576101008083540402835291602001916104fc565b820191906000526020600020905b8154815290600101906020018083116104df57829003601f168201915b505092505050604051602081830303815290604052805190602001201415610522578091505b6001016103bf565b5092915050565b6000428161053e856103b8565b6000191415905080610617576040805160808101825283815260208082018881523393830193909352606082018790526000805460018101808355918052835160049091027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563810191825594518051929591936105e2937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e564909301929101906106b3565b5060408201516002820180546001600160a01b0319166001600160a01b03909216919091179055606090910151600390910155505b15949350505050565b600061062b836103b8565b905080600019141561063957fe5b6000818154811061064657fe5b60009182526020909120600490910201600201546001600160a01b0316331461066b57fe5b816000828154811061067957fe5b906000526020600020906004020160020160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106f457805160ff1916838001178555610721565b82800160010185558215610721579182015b82811115610721578251825591602001919060010190610706565b5061072d929150610731565b5090565b61074b91905b8082111561072d5760008155600101610737565b9056fea165627a7a723058204eaefbdbebaf63fc9869dc9e7132c360be05bb4b7fdd73ee9c616e996c4d3f8b0029`

// DeployImpro deploys a new Ethereum contract, binding an instance of Impro to it.
func DeployImpro(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Impro, error) {
	parsed, err := abi.JSON(strings.NewReader(ImproABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ImproBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Impro{ImproCaller: ImproCaller{contract: contract}, ImproTransactor: ImproTransactor{contract: contract}, ImproFilterer: ImproFilterer{contract: contract}}, nil
}

// Impro is an auto generated Go binding around an Ethereum contract.
type Impro struct {
	ImproCaller     // Read-only binding to the contract
	ImproTransactor // Write-only binding to the contract
	ImproFilterer   // Log filterer for contract events
}

// ImproCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImproCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImproTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImproTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImproFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImproFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImproSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImproSession struct {
	Contract     *Impro            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImproCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImproCallerSession struct {
	Contract *ImproCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ImproTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImproTransactorSession struct {
	Contract     *ImproTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImproRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImproRaw struct {
	Contract *Impro // Generic contract binding to access the raw methods on
}

// ImproCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImproCallerRaw struct {
	Contract *ImproCaller // Generic read-only contract binding to access the raw methods on
}

// ImproTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImproTransactorRaw struct {
	Contract *ImproTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImpro creates a new instance of Impro, bound to a specific deployed contract.
func NewImpro(address common.Address, backend bind.ContractBackend) (*Impro, error) {
	contract, err := bindImpro(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Impro{ImproCaller: ImproCaller{contract: contract}, ImproTransactor: ImproTransactor{contract: contract}, ImproFilterer: ImproFilterer{contract: contract}}, nil
}

// NewImproCaller creates a new read-only instance of Impro, bound to a specific deployed contract.
func NewImproCaller(address common.Address, caller bind.ContractCaller) (*ImproCaller, error) {
	contract, err := bindImpro(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImproCaller{contract: contract}, nil
}

// NewImproTransactor creates a new write-only instance of Impro, bound to a specific deployed contract.
func NewImproTransactor(address common.Address, transactor bind.ContractTransactor) (*ImproTransactor, error) {
	contract, err := bindImpro(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImproTransactor{contract: contract}, nil
}

// NewImproFilterer creates a new log filterer instance of Impro, bound to a specific deployed contract.
func NewImproFilterer(address common.Address, filterer bind.ContractFilterer) (*ImproFilterer, error) {
	contract, err := bindImpro(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImproFilterer{contract: contract}, nil
}

// bindImpro binds a generic wrapper to an already deployed contract.
func bindImpro(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ImproABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Impro *ImproRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Impro.Contract.ImproCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Impro *ImproRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impro.Contract.ImproTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Impro *ImproRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Impro.Contract.ImproTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Impro *ImproCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Impro.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Impro *ImproTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Impro.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Impro *ImproTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Impro.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0xb6f921ad.
//
// Solidity: function check(string _perceptualHash) constant returns(int256 exist_at)
func (_Impro *ImproCaller) Check(opts *bind.CallOpts, _perceptualHash string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Impro.contract.Call(opts, out, "check", _perceptualHash)
	return *ret0, err
}

// Check is a free data retrieval call binding the contract method 0xb6f921ad.
//
// Solidity: function check(string _perceptualHash) constant returns(int256 exist_at)
func (_Impro *ImproSession) Check(_perceptualHash string) (*big.Int, error) {
	return _Impro.Contract.Check(&_Impro.CallOpts, _perceptualHash)
}

// Check is a free data retrieval call binding the contract method 0xb6f921ad.
//
// Solidity: function check(string _perceptualHash) constant returns(int256 exist_at)
func (_Impro *ImproCallerSession) Check(_perceptualHash string) (*big.Int, error) {
	return _Impro.Contract.Check(&_Impro.CallOpts, _perceptualHash)
}

// Buy is a paid mutator transaction binding the contract method 0x492cc769.
//
// Solidity: function buy(string _perceptualHash) returns()
func (_Impro *ImproTransactor) Buy(opts *bind.TransactOpts, _perceptualHash string) (*types.Transaction, error) {
	return _Impro.contract.Transact(opts, "buy", _perceptualHash)
}

// Buy is a paid mutator transaction binding the contract method 0x492cc769.
//
// Solidity: function buy(string _perceptualHash) returns()
func (_Impro *ImproSession) Buy(_perceptualHash string) (*types.Transaction, error) {
	return _Impro.Contract.Buy(&_Impro.TransactOpts, _perceptualHash)
}

// Buy is a paid mutator transaction binding the contract method 0x492cc769.
//
// Solidity: function buy(string _perceptualHash) returns()
func (_Impro *ImproTransactorSession) Buy(_perceptualHash string) (*types.Transaction, error) {
	return _Impro.Contract.Buy(&_Impro.TransactOpts, _perceptualHash)
}

// Register is a paid mutator transaction binding the contract method 0xea87152b.
//
// Solidity: function register(string _perceptualHash, uint256 price) returns(bool success)
func (_Impro *ImproTransactor) Register(opts *bind.TransactOpts, _perceptualHash string, price *big.Int) (*types.Transaction, error) {
	return _Impro.contract.Transact(opts, "register", _perceptualHash, price)
}

// Register is a paid mutator transaction binding the contract method 0xea87152b.
//
// Solidity: function register(string _perceptualHash, uint256 price) returns(bool success)
func (_Impro *ImproSession) Register(_perceptualHash string, price *big.Int) (*types.Transaction, error) {
	return _Impro.Contract.Register(&_Impro.TransactOpts, _perceptualHash, price)
}

// Register is a paid mutator transaction binding the contract method 0xea87152b.
//
// Solidity: function register(string _perceptualHash, uint256 price) returns(bool success)
func (_Impro *ImproTransactorSession) Register(_perceptualHash string, price *big.Int) (*types.Transaction, error) {
	return _Impro.Contract.Register(&_Impro.TransactOpts, _perceptualHash, price)
}

// Transfer is a paid mutator transaction binding the contract method 0xfbf58b3e.
//
// Solidity: function transfer(string _perceptualHash, address _newOwner) returns()
func (_Impro *ImproTransactor) Transfer(opts *bind.TransactOpts, _perceptualHash string, _newOwner common.Address) (*types.Transaction, error) {
	return _Impro.contract.Transact(opts, "transfer", _perceptualHash, _newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0xfbf58b3e.
//
// Solidity: function transfer(string _perceptualHash, address _newOwner) returns()
func (_Impro *ImproSession) Transfer(_perceptualHash string, _newOwner common.Address) (*types.Transaction, error) {
	return _Impro.Contract.Transfer(&_Impro.TransactOpts, _perceptualHash, _newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0xfbf58b3e.
//
// Solidity: function transfer(string _perceptualHash, address _newOwner) returns()
func (_Impro *ImproTransactorSession) Transfer(_perceptualHash string, _newOwner common.Address) (*types.Transaction, error) {
	return _Impro.Contract.Transfer(&_Impro.TransactOpts, _perceptualHash, _newOwner)
}
