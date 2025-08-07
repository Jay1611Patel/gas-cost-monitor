// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MyContractMetaData contains all meta data concerning the MyContract contract.
var MyContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialValue\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"LogMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"addAddressToArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_keys\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"addMultipleValuesToMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_shouldRun\",\"type\":\"bool\"}],\"name\":\"conditionalExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getAddressArrayElement\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressArrayLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"}],\"name\":\"getMappingValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUintValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"myMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myUintValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"performComplexCalculation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLastAddressFromArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"setUintValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"triggerEventAndReturn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MyContractMetaData.ABI instead.
var MyContractABI = MyContractMetaData.ABI

// MyContract is an auto generated Go binding around an Ethereum contract.
type MyContract struct {
	MyContractCaller     // Read-only binding to the contract
	MyContractTransactor // Write-only binding to the contract
	MyContractFilterer   // Log filterer for contract events
}

// MyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyContractSession struct {
	Contract     *MyContract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyContractCallerSession struct {
	Contract *MyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyContractTransactorSession struct {
	Contract     *MyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyContractRaw struct {
	Contract *MyContract // Generic contract binding to access the raw methods on
}

// MyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyContractCallerRaw struct {
	Contract *MyContractCaller // Generic read-only contract binding to access the raw methods on
}

// MyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyContractTransactorRaw struct {
	Contract *MyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyContract creates a new instance of MyContract, bound to a specific deployed contract.
func NewMyContract(address common.Address, backend bind.ContractBackend) (*MyContract, error) {
	contract, err := bindMyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// NewMyContractCaller creates a new read-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractCaller(address common.Address, caller bind.ContractCaller) (*MyContractCaller, error) {
	contract, err := bindMyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractCaller{contract: contract}, nil
}

// NewMyContractTransactor creates a new write-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MyContractTransactor, error) {
	contract, err := bindMyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractTransactor{contract: contract}, nil
}

// NewMyContractFilterer creates a new log filterer instance of MyContract, bound to a specific deployed contract.
func NewMyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MyContractFilterer, error) {
	contract, err := bindMyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyContractFilterer{contract: contract}, nil
}

// bindMyContract binds a generic wrapper to an already deployed contract.
func bindMyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.MyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transact(opts, method, params...)
}

// AddressArray is a free data retrieval call binding the contract method 0x0f96cf34.
//
// Solidity: function addressArray(uint256 ) view returns(address)
func (_MyContract *MyContractCaller) AddressArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "addressArray", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressArray is a free data retrieval call binding the contract method 0x0f96cf34.
//
// Solidity: function addressArray(uint256 ) view returns(address)
func (_MyContract *MyContractSession) AddressArray(arg0 *big.Int) (common.Address, error) {
	return _MyContract.Contract.AddressArray(&_MyContract.CallOpts, arg0)
}

// AddressArray is a free data retrieval call binding the contract method 0x0f96cf34.
//
// Solidity: function addressArray(uint256 ) view returns(address)
func (_MyContract *MyContractCallerSession) AddressArray(arg0 *big.Int) (common.Address, error) {
	return _MyContract.Contract.AddressArray(&_MyContract.CallOpts, arg0)
}

// GetAddressArrayElement is a free data retrieval call binding the contract method 0xc8c7e90e.
//
// Solidity: function getAddressArrayElement(uint256 _index) view returns(address)
func (_MyContract *MyContractCaller) GetAddressArrayElement(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "getAddressArrayElement", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressArrayElement is a free data retrieval call binding the contract method 0xc8c7e90e.
//
// Solidity: function getAddressArrayElement(uint256 _index) view returns(address)
func (_MyContract *MyContractSession) GetAddressArrayElement(_index *big.Int) (common.Address, error) {
	return _MyContract.Contract.GetAddressArrayElement(&_MyContract.CallOpts, _index)
}

// GetAddressArrayElement is a free data retrieval call binding the contract method 0xc8c7e90e.
//
// Solidity: function getAddressArrayElement(uint256 _index) view returns(address)
func (_MyContract *MyContractCallerSession) GetAddressArrayElement(_index *big.Int) (common.Address, error) {
	return _MyContract.Contract.GetAddressArrayElement(&_MyContract.CallOpts, _index)
}

// GetAddressArrayLength is a free data retrieval call binding the contract method 0x994144f3.
//
// Solidity: function getAddressArrayLength() view returns(uint256)
func (_MyContract *MyContractCaller) GetAddressArrayLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "getAddressArrayLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressArrayLength is a free data retrieval call binding the contract method 0x994144f3.
//
// Solidity: function getAddressArrayLength() view returns(uint256)
func (_MyContract *MyContractSession) GetAddressArrayLength() (*big.Int, error) {
	return _MyContract.Contract.GetAddressArrayLength(&_MyContract.CallOpts)
}

// GetAddressArrayLength is a free data retrieval call binding the contract method 0x994144f3.
//
// Solidity: function getAddressArrayLength() view returns(uint256)
func (_MyContract *MyContractCallerSession) GetAddressArrayLength() (*big.Int, error) {
	return _MyContract.Contract.GetAddressArrayLength(&_MyContract.CallOpts)
}

// GetMappingValue is a free data retrieval call binding the contract method 0xbba50db2.
//
// Solidity: function getMappingValue(uint256 _key) view returns(uint256)
func (_MyContract *MyContractCaller) GetMappingValue(opts *bind.CallOpts, _key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "getMappingValue", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMappingValue is a free data retrieval call binding the contract method 0xbba50db2.
//
// Solidity: function getMappingValue(uint256 _key) view returns(uint256)
func (_MyContract *MyContractSession) GetMappingValue(_key *big.Int) (*big.Int, error) {
	return _MyContract.Contract.GetMappingValue(&_MyContract.CallOpts, _key)
}

// GetMappingValue is a free data retrieval call binding the contract method 0xbba50db2.
//
// Solidity: function getMappingValue(uint256 _key) view returns(uint256)
func (_MyContract *MyContractCallerSession) GetMappingValue(_key *big.Int) (*big.Int, error) {
	return _MyContract.Contract.GetMappingValue(&_MyContract.CallOpts, _key)
}

// GetUintValue is a free data retrieval call binding the contract method 0x55ec6354.
//
// Solidity: function getUintValue() view returns(uint256)
func (_MyContract *MyContractCaller) GetUintValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "getUintValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintValue is a free data retrieval call binding the contract method 0x55ec6354.
//
// Solidity: function getUintValue() view returns(uint256)
func (_MyContract *MyContractSession) GetUintValue() (*big.Int, error) {
	return _MyContract.Contract.GetUintValue(&_MyContract.CallOpts)
}

// GetUintValue is a free data retrieval call binding the contract method 0x55ec6354.
//
// Solidity: function getUintValue() view returns(uint256)
func (_MyContract *MyContractCallerSession) GetUintValue() (*big.Int, error) {
	return _MyContract.Contract.GetUintValue(&_MyContract.CallOpts)
}

// MyMapping is a free data retrieval call binding the contract method 0x7c668844.
//
// Solidity: function myMapping(uint256 ) view returns(uint256)
func (_MyContract *MyContractCaller) MyMapping(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "myMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyMapping is a free data retrieval call binding the contract method 0x7c668844.
//
// Solidity: function myMapping(uint256 ) view returns(uint256)
func (_MyContract *MyContractSession) MyMapping(arg0 *big.Int) (*big.Int, error) {
	return _MyContract.Contract.MyMapping(&_MyContract.CallOpts, arg0)
}

// MyMapping is a free data retrieval call binding the contract method 0x7c668844.
//
// Solidity: function myMapping(uint256 ) view returns(uint256)
func (_MyContract *MyContractCallerSession) MyMapping(arg0 *big.Int) (*big.Int, error) {
	return _MyContract.Contract.MyMapping(&_MyContract.CallOpts, arg0)
}

// MyUintValue is a free data retrieval call binding the contract method 0x99fa108d.
//
// Solidity: function myUintValue() view returns(uint256)
func (_MyContract *MyContractCaller) MyUintValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "myUintValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyUintValue is a free data retrieval call binding the contract method 0x99fa108d.
//
// Solidity: function myUintValue() view returns(uint256)
func (_MyContract *MyContractSession) MyUintValue() (*big.Int, error) {
	return _MyContract.Contract.MyUintValue(&_MyContract.CallOpts)
}

// MyUintValue is a free data retrieval call binding the contract method 0x99fa108d.
//
// Solidity: function myUintValue() view returns(uint256)
func (_MyContract *MyContractCallerSession) MyUintValue() (*big.Int, error) {
	return _MyContract.Contract.MyUintValue(&_MyContract.CallOpts)
}

// AddAddressToArray is a paid mutator transaction binding the contract method 0x8d645e54.
//
// Solidity: function addAddressToArray(address _newAddress) returns()
func (_MyContract *MyContractTransactor) AddAddressToArray(opts *bind.TransactOpts, _newAddress common.Address) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "addAddressToArray", _newAddress)
}

// AddAddressToArray is a paid mutator transaction binding the contract method 0x8d645e54.
//
// Solidity: function addAddressToArray(address _newAddress) returns()
func (_MyContract *MyContractSession) AddAddressToArray(_newAddress common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.AddAddressToArray(&_MyContract.TransactOpts, _newAddress)
}

// AddAddressToArray is a paid mutator transaction binding the contract method 0x8d645e54.
//
// Solidity: function addAddressToArray(address _newAddress) returns()
func (_MyContract *MyContractTransactorSession) AddAddressToArray(_newAddress common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.AddAddressToArray(&_MyContract.TransactOpts, _newAddress)
}

// AddMultipleValuesToMapping is a paid mutator transaction binding the contract method 0xb95ce087.
//
// Solidity: function addMultipleValuesToMapping(uint256[] _keys, uint256[] _values) returns()
func (_MyContract *MyContractTransactor) AddMultipleValuesToMapping(opts *bind.TransactOpts, _keys []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "addMultipleValuesToMapping", _keys, _values)
}

// AddMultipleValuesToMapping is a paid mutator transaction binding the contract method 0xb95ce087.
//
// Solidity: function addMultipleValuesToMapping(uint256[] _keys, uint256[] _values) returns()
func (_MyContract *MyContractSession) AddMultipleValuesToMapping(_keys []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.AddMultipleValuesToMapping(&_MyContract.TransactOpts, _keys, _values)
}

// AddMultipleValuesToMapping is a paid mutator transaction binding the contract method 0xb95ce087.
//
// Solidity: function addMultipleValuesToMapping(uint256[] _keys, uint256[] _values) returns()
func (_MyContract *MyContractTransactorSession) AddMultipleValuesToMapping(_keys []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.AddMultipleValuesToMapping(&_MyContract.TransactOpts, _keys, _values)
}

// ConditionalExecution is a paid mutator transaction binding the contract method 0xa75a0da9.
//
// Solidity: function conditionalExecution(bool _shouldRun) returns()
func (_MyContract *MyContractTransactor) ConditionalExecution(opts *bind.TransactOpts, _shouldRun bool) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "conditionalExecution", _shouldRun)
}

// ConditionalExecution is a paid mutator transaction binding the contract method 0xa75a0da9.
//
// Solidity: function conditionalExecution(bool _shouldRun) returns()
func (_MyContract *MyContractSession) ConditionalExecution(_shouldRun bool) (*types.Transaction, error) {
	return _MyContract.Contract.ConditionalExecution(&_MyContract.TransactOpts, _shouldRun)
}

// ConditionalExecution is a paid mutator transaction binding the contract method 0xa75a0da9.
//
// Solidity: function conditionalExecution(bool _shouldRun) returns()
func (_MyContract *MyContractTransactorSession) ConditionalExecution(_shouldRun bool) (*types.Transaction, error) {
	return _MyContract.Contract.ConditionalExecution(&_MyContract.TransactOpts, _shouldRun)
}

// PerformComplexCalculation is a paid mutator transaction binding the contract method 0x9cf74b42.
//
// Solidity: function performComplexCalculation(uint256 _a, uint256 _b) returns()
func (_MyContract *MyContractTransactor) PerformComplexCalculation(opts *bind.TransactOpts, _a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "performComplexCalculation", _a, _b)
}

// PerformComplexCalculation is a paid mutator transaction binding the contract method 0x9cf74b42.
//
// Solidity: function performComplexCalculation(uint256 _a, uint256 _b) returns()
func (_MyContract *MyContractSession) PerformComplexCalculation(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.PerformComplexCalculation(&_MyContract.TransactOpts, _a, _b)
}

// PerformComplexCalculation is a paid mutator transaction binding the contract method 0x9cf74b42.
//
// Solidity: function performComplexCalculation(uint256 _a, uint256 _b) returns()
func (_MyContract *MyContractTransactorSession) PerformComplexCalculation(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.PerformComplexCalculation(&_MyContract.TransactOpts, _a, _b)
}

// RemoveLastAddressFromArray is a paid mutator transaction binding the contract method 0x686e543c.
//
// Solidity: function removeLastAddressFromArray() returns()
func (_MyContract *MyContractTransactor) RemoveLastAddressFromArray(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "removeLastAddressFromArray")
}

// RemoveLastAddressFromArray is a paid mutator transaction binding the contract method 0x686e543c.
//
// Solidity: function removeLastAddressFromArray() returns()
func (_MyContract *MyContractSession) RemoveLastAddressFromArray() (*types.Transaction, error) {
	return _MyContract.Contract.RemoveLastAddressFromArray(&_MyContract.TransactOpts)
}

// RemoveLastAddressFromArray is a paid mutator transaction binding the contract method 0x686e543c.
//
// Solidity: function removeLastAddressFromArray() returns()
func (_MyContract *MyContractTransactorSession) RemoveLastAddressFromArray() (*types.Transaction, error) {
	return _MyContract.Contract.RemoveLastAddressFromArray(&_MyContract.TransactOpts)
}

// SetUintValue is a paid mutator transaction binding the contract method 0x2f3b21a2.
//
// Solidity: function setUintValue(uint256 _newValue) returns()
func (_MyContract *MyContractTransactor) SetUintValue(opts *bind.TransactOpts, _newValue *big.Int) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "setUintValue", _newValue)
}

// SetUintValue is a paid mutator transaction binding the contract method 0x2f3b21a2.
//
// Solidity: function setUintValue(uint256 _newValue) returns()
func (_MyContract *MyContractSession) SetUintValue(_newValue *big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.SetUintValue(&_MyContract.TransactOpts, _newValue)
}

// SetUintValue is a paid mutator transaction binding the contract method 0x2f3b21a2.
//
// Solidity: function setUintValue(uint256 _newValue) returns()
func (_MyContract *MyContractTransactorSession) SetUintValue(_newValue *big.Int) (*types.Transaction, error) {
	return _MyContract.Contract.SetUintValue(&_MyContract.TransactOpts, _newValue)
}

// TriggerEventAndReturn is a paid mutator transaction binding the contract method 0x1b64400e.
//
// Solidity: function triggerEventAndReturn(string _message) returns(bool)
func (_MyContract *MyContractTransactor) TriggerEventAndReturn(opts *bind.TransactOpts, _message string) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "triggerEventAndReturn", _message)
}

// TriggerEventAndReturn is a paid mutator transaction binding the contract method 0x1b64400e.
//
// Solidity: function triggerEventAndReturn(string _message) returns(bool)
func (_MyContract *MyContractSession) TriggerEventAndReturn(_message string) (*types.Transaction, error) {
	return _MyContract.Contract.TriggerEventAndReturn(&_MyContract.TransactOpts, _message)
}

// TriggerEventAndReturn is a paid mutator transaction binding the contract method 0x1b64400e.
//
// Solidity: function triggerEventAndReturn(string _message) returns(bool)
func (_MyContract *MyContractTransactorSession) TriggerEventAndReturn(_message string) (*types.Transaction, error) {
	return _MyContract.Contract.TriggerEventAndReturn(&_MyContract.TransactOpts, _message)
}

// MyContractLogMessageIterator is returned from FilterLogMessage and is used to iterate over the raw logs and unpacked data for LogMessage events raised by the MyContract contract.
type MyContractLogMessageIterator struct {
	Event *MyContractLogMessage // Event containing the contract specifics and raw log

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
func (it *MyContractLogMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractLogMessage)
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
		it.Event = new(MyContractLogMessage)
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
func (it *MyContractLogMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractLogMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractLogMessage represents a LogMessage event raised by the MyContract contract.
type MyContractLogMessage struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogMessage is a free log retrieval operation binding the contract event 0x96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480.
//
// Solidity: event LogMessage(string message)
func (_MyContract *MyContractFilterer) FilterLogMessage(opts *bind.FilterOpts) (*MyContractLogMessageIterator, error) {

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "LogMessage")
	if err != nil {
		return nil, err
	}
	return &MyContractLogMessageIterator{contract: _MyContract.contract, event: "LogMessage", logs: logs, sub: sub}, nil
}

// WatchLogMessage is a free log subscription operation binding the contract event 0x96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480.
//
// Solidity: event LogMessage(string message)
func (_MyContract *MyContractFilterer) WatchLogMessage(opts *bind.WatchOpts, sink chan<- *MyContractLogMessage) (event.Subscription, error) {

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "LogMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractLogMessage)
				if err := _MyContract.contract.UnpackLog(event, "LogMessage", log); err != nil {
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

// ParseLogMessage is a log parse operation binding the contract event 0x96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480.
//
// Solidity: event LogMessage(string message)
func (_MyContract *MyContractFilterer) ParseLogMessage(log types.Log) (*MyContractLogMessage, error) {
	event := new(MyContractLogMessage)
	if err := _MyContract.contract.UnpackLog(event, "LogMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
