// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
const ImproABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"},{\"name\":\"_price\",\"type\":\"int256\"}],\"name\":\"changePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"price\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ImproBin is the compiled bytecode used for deploying new contracts.
const ImproBin = `0x608060405234801561001057600080fd5b50610a86806100206000396000f3fe6080604052600436106100705760003560e01c8063492cc7691161004e578063492cc769146102b25780634aaf4a1214610356578063524f388914610423578063e2a1a99d146104d457610070565b8063261a323e146100755780632dee12841461013a57806332e5e595146101ef575b600080fd5b34801561008157600080fd5b506101266004803603602081101561009857600080fd5b810190602081018135600160201b8111156100b257600080fd5b8201836020820111156100c457600080fd5b803590602001918460018302840111600160201b831117156100e557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610587945050505050565b604080519115158252519081900360200190f35b34801561014657600080fd5b506101ed6004803603604081101561015d57600080fd5b810190602081018135600160201b81111561017757600080fd5b82018360208201111561018957600080fd5b803590602001918460018302840111600160201b831117156101aa57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506105f2915050565b005b3480156101fb57600080fd5b506102a06004803603602081101561021257600080fd5b810190602081018135600160201b81111561022c57600080fd5b82018360208201111561023e57600080fd5b803590602001918460018302840111600160201b8311171561025f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061068f945050505050565b60408051918252519081900360200190f35b6101ed600480360360208110156102c857600080fd5b810190602081018135600160201b8111156102e257600080fd5b8201836020820111156102f457600080fd5b803590602001918460018302840111600160201b8311171561031557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610706945050505050565b34801561036257600080fd5b506104076004803603602081101561037957600080fd5b810190602081018135600160201b81111561039357600080fd5b8201836020820111156103a557600080fd5b803590602001918460018302840111600160201b831117156103c657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107d0945050505050565b604080516001600160a01b039092168252519081900360200190f35b34801561042f57600080fd5b506102a06004803603602081101561044657600080fd5b810190602081018135600160201b81111561046057600080fd5b82018360208201111561047257600080fd5b803590602001918460018302840111600160201b8311171561049357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610853945050505050565b3480156104e057600080fd5b506101ed600480360360408110156104f757600080fd5b810190602081018135600160201b81111561051157600080fd5b82018360208201111561052357600080fd5b803590602001918460018302840111600160201b8311171561054457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506108cd915050565b60006001826040518082805190602001908083835b602083106105bb5780518252601f19909201916020918201910161059c565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600114949350505050565b6105fb82610587565b61060157fe5b61060a826107d0565b6001600160a01b0316336001600160a01b03161461062457fe5b806000836040518082805190602001908083835b602083106106575780518252601f199092019160209182019101610638565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206002019290925550505050565b600061069a82610587565b6106a057fe5b6000826040518082805190602001908083835b602083106106d25780518252601f1990920191602091820191016106b3565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054949350505050565b61070f81610587565b61071557fe5b61071e816107d0565b6001600160a01b0316336001600160a01b0316141561073957fe5b61074281610853565b34121561074b57fe5b336000826040518082805190602001908083835b6020831061077e5780518252601f19909201916020918201910161075f565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060010180546001600160a01b0319166001600160a01b039490941693909317909255505050565b60006107db82610587565b6107e157fe5b6000826040518082805190602001908083835b602083106108135780518252601f1990920191602091820191016107f4565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600101546001600160a01b0316949350505050565b600061085e82610587565b61086457fe5b6000826040518082805190602001908083835b602083106108965780518252601f199092019160209182019101610877565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060020154949350505050565b60004290506001836040518082805190602001908083835b602083106109045780518252601f1990920191602091820191016108e5565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205415915061093b905057fe5b6040518060600160405280828152602001336001600160a01b03168152602001838152506000846040518082805190602001908083835b602083106109915780518252601f199092019160209182019101610972565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185208651815586830151600180830180546001600160a01b0319166001600160a01b0390931692909217909155969091015160029091015587518594899450925082918401908083835b60208310610a245780518252601f199092019160209182019101610a05565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092909255505050505056fea165627a7a723058204f884a43f841a726838a854df8a815675f61a999654882ccc252b7f527beb43a0029`

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

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string _perceptualHash) constant returns(bool)
func (_Impro *ImproCaller) Exists(opts *bind.CallOpts, _perceptualHash string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Impro.contract.Call(opts, out, "exists", _perceptualHash)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string _perceptualHash) constant returns(bool)
func (_Impro *ImproSession) Exists(_perceptualHash string) (bool, error) {
	return _Impro.Contract.Exists(&_Impro.CallOpts, _perceptualHash)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string _perceptualHash) constant returns(bool)
func (_Impro *ImproCallerSession) Exists(_perceptualHash string) (bool, error) {
	return _Impro.Contract.Exists(&_Impro.CallOpts, _perceptualHash)
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string _perceptualHash) constant returns(address owner)
func (_Impro *ImproCaller) GetOwner(opts *bind.CallOpts, _perceptualHash string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Impro.contract.Call(opts, out, "getOwner", _perceptualHash)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string _perceptualHash) constant returns(address owner)
func (_Impro *ImproSession) GetOwner(_perceptualHash string) (common.Address, error) {
	return _Impro.Contract.GetOwner(&_Impro.CallOpts, _perceptualHash)
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string _perceptualHash) constant returns(address owner)
func (_Impro *ImproCallerSession) GetOwner(_perceptualHash string) (common.Address, error) {
	return _Impro.Contract.GetOwner(&_Impro.CallOpts, _perceptualHash)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _perceptualHash) constant returns(int256 price)
func (_Impro *ImproCaller) GetPrice(opts *bind.CallOpts, _perceptualHash string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Impro.contract.Call(opts, out, "getPrice", _perceptualHash)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _perceptualHash) constant returns(int256 price)
func (_Impro *ImproSession) GetPrice(_perceptualHash string) (*big.Int, error) {
	return _Impro.Contract.GetPrice(&_Impro.CallOpts, _perceptualHash)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _perceptualHash) constant returns(int256 price)
func (_Impro *ImproCallerSession) GetPrice(_perceptualHash string) (*big.Int, error) {
	return _Impro.Contract.GetPrice(&_Impro.CallOpts, _perceptualHash)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x32e5e595.
//
// Solidity: function getTimestamp(string _perceptualHash) constant returns(uint256 price)
func (_Impro *ImproCaller) GetTimestamp(opts *bind.CallOpts, _perceptualHash string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Impro.contract.Call(opts, out, "getTimestamp", _perceptualHash)
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0x32e5e595.
//
// Solidity: function getTimestamp(string _perceptualHash) constant returns(uint256 price)
func (_Impro *ImproSession) GetTimestamp(_perceptualHash string) (*big.Int, error) {
	return _Impro.Contract.GetTimestamp(&_Impro.CallOpts, _perceptualHash)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x32e5e595.
//
// Solidity: function getTimestamp(string _perceptualHash) constant returns(uint256 price)
func (_Impro *ImproCallerSession) GetTimestamp(_perceptualHash string) (*big.Int, error) {
	return _Impro.Contract.GetTimestamp(&_Impro.CallOpts, _perceptualHash)
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

// ChangePrice is a paid mutator transaction binding the contract method 0x2dee1284.
//
// Solidity: function changePrice(string _perceptualHash, int256 _price) returns()
func (_Impro *ImproTransactor) ChangePrice(opts *bind.TransactOpts, _perceptualHash string, _price *big.Int) (*types.Transaction, error) {
	return _Impro.contract.Transact(opts, "changePrice", _perceptualHash, _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x2dee1284.
//
// Solidity: function changePrice(string _perceptualHash, int256 _price) returns()
func (_Impro *ImproSession) ChangePrice(_perceptualHash string, _price *big.Int) (*types.Transaction, error) {
	return _Impro.Contract.ChangePrice(&_Impro.TransactOpts, _perceptualHash, _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x2dee1284.
//
// Solidity: function changePrice(string _perceptualHash, int256 _price) returns()
func (_Impro *ImproTransactorSession) ChangePrice(_perceptualHash string, _price *big.Int) (*types.Transaction, error) {
	return _Impro.Contract.ChangePrice(&_Impro.TransactOpts, _perceptualHash, _price)
}

// Register is a paid mutator transaction binding the contract method 0xe2a1a99d.
//
// Solidity: function register(string _perceptualHash, int256 price) returns()
func (_Impro *ImproTransactor) Register(opts *bind.TransactOpts, _perceptualHash string, price *big.Int) (*types.Transaction, error) {
	return _Impro.contract.Transact(opts, "register", _perceptualHash, price)
}

// Register is a paid mutator transaction binding the contract method 0xe2a1a99d.
//
// Solidity: function register(string _perceptualHash, int256 price) returns()
func (_Impro *ImproSession) Register(_perceptualHash string, price *big.Int) (*types.Transaction, error) {
	return _Impro.Contract.Register(&_Impro.TransactOpts, _perceptualHash, price)
}

// Register is a paid mutator transaction binding the contract method 0xe2a1a99d.
//
// Solidity: function register(string _perceptualHash, int256 price) returns()
func (_Impro *ImproTransactorSession) Register(_perceptualHash string, price *big.Int) (*types.Transaction, error) {
	return _Impro.Contract.Register(&_Impro.TransactOpts, _perceptualHash, price)
}
