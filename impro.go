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
const ImproABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"},{\"name\":\"_price\",\"type\":\"int256\"}],\"name\":\"changePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"price\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_perceptualHash\",\"type\":\"string\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ImproBin is the compiled bytecode used for deploying new contracts.
const ImproBin = `0x608060405234801561001057600080fd5b50610d4f806100206000396000f3fe60806040526004361061007b5760003560e01c80634aaf4a121161004e5780634aaf4a1214610361578063524f38891461042e578063e2a1a99d146104df578063fbf58b3e146105925761007b565b8063261a323e146100805780632dee12841461014557806332e5e595146101fa578063492cc769146102bd575b600080fd5b34801561008c57600080fd5b50610131600480360360208110156100a357600080fd5b810190602081018135600160201b8111156100bd57600080fd5b8201836020820111156100cf57600080fd5b803590602001918460018302840111600160201b831117156100f057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610641945050505050565b604080519115158252519081900360200190f35b34801561015157600080fd5b506101f86004803603604081101561016857600080fd5b810190602081018135600160201b81111561018257600080fd5b82018360208201111561019457600080fd5b803590602001918460018302840111600160201b831117156101b557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506106ac915050565b005b34801561020657600080fd5b506102ab6004803603602081101561021d57600080fd5b810190602081018135600160201b81111561023757600080fd5b82018360208201111561024957600080fd5b803590602001918460018302840111600160201b8311171561026a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610749945050505050565b60408051918252519081900360200190f35b6101f8600480360360208110156102d357600080fd5b810190602081018135600160201b8111156102ed57600080fd5b8201836020820111156102ff57600080fd5b803590602001918460018302840111600160201b8311171561032057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107c0945050505050565b34801561036d57600080fd5b506104126004803603602081101561038457600080fd5b810190602081018135600160201b81111561039e57600080fd5b8201836020820111156103b057600080fd5b803590602001918460018302840111600160201b831117156103d157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109a6945050505050565b604080516001600160a01b039092168252519081900360200190f35b34801561043a57600080fd5b506102ab6004803603602081101561045157600080fd5b810190602081018135600160201b81111561046b57600080fd5b82018360208201111561047d57600080fd5b803590602001918460018302840111600160201b8311171561049e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a29945050505050565b3480156104eb57600080fd5b506101f86004803603604081101561050257600080fd5b810190602081018135600160201b81111561051c57600080fd5b82018360208201111561052e57600080fd5b803590602001918460018302840111600160201b8311171561054f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610aa3915050565b6101f8600480360360408110156105a857600080fd5b810190602081018135600160201b8111156105c257600080fd5b8201836020820111156105d457600080fd5b803590602001918460018302840111600160201b831117156105f557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610c309050565b60006001826040518082805190602001908083835b602083106106755780518252601f199092019160209182019101610656565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600114949350505050565b6106b582610641565b6106bb57fe5b6106c4826109a6565b6001600160a01b0316336001600160a01b0316146106de57fe5b806000836040518082805190602001908083835b602083106107115780518252601f1990920191602091820191016106f2565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206002019290925550505050565b600061075482610641565b61075a57fe5b6000826040518082805190602001908083835b6020831061078c5780518252601f19909201916020918201910161076d565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054949350505050565b600080826040518082805190602001908083835b602083106107f35780518252601f1990920191602091820191016107d4565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600101546001600160a01b0316925033915061083b905083610641565b61084157fe5b61084a836109a6565b6001600160a01b0316336001600160a01b0316141561086557fe5b61086e83610a29565b34121561087757fe5b806000846040518082805190602001908083835b602083106108aa5780518252601f19909201916020918201910161088b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506000196000846040518082805190602001908083835b602083106109395780518252601f19909201916020918201910161091a565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301832060020193909355506001600160a01b03841691503480156108fc02916000818181858888f193505050501580156109a0573d6000803e3d6000fd5b50505050565b60006109b182610641565b6109b757fe5b6000826040518082805190602001908083835b602083106109e95780518252601f1990920191602091820191016109ca565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600101546001600160a01b0316949350505050565b6000610a3482610641565b610a3a57fe5b6000826040518082805190602001908083835b60208310610a6c5780518252601f199092019160209182019101610a4d565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060020154949350505050565b60004290506001836040518082805190602001908083835b60208310610ada5780518252601f199092019160209182019101610abb565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054159150610b11905057fe5b6040518060600160405280828152602001336001600160a01b03168152602001838152506000846040518082805190602001908083835b60208310610b675780518252601f199092019160209182019101610b48565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185208651815586830151600180830180546001600160a01b0319166001600160a01b0390931692909217909155969091015160029091015587518594899450925082918401908083835b60208310610bfa5780518252601f199092019160209182019101610bdb565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929092555050505050565b610c3982610641565b610c3f57fe5b610c48826109a6565b6001600160a01b0316336001600160a01b031614610c6257fe5b806000836040518082805190602001908083835b60208310610c955780518252601f199092019160209182019101610c76565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600019600083604051808280519060200190808383602083106107115780518252601f1990920191602091820191016106f256fea165627a7a7230582048b0191ac659956dfc72585474e2c02087f70b42a2b9b9b2e006b111ba7d1fa80029`

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
