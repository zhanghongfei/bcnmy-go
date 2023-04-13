// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package demo

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
)

// TransferHandlerPermitOptions is an auto generated low-level Go binding around an user-defined struct.
type TransferHandlerPermitOptions struct {
	Value    *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Allowed  bool
	V        uint8
	R        [32]byte
	S        [32]byte
}

// TransferDemoMetaData contains all meta data concerning the TransferDemo contract.
var TransferDemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structTransferHandler.PermitOptions\",\"name\":\"options\",\"type\":\"tuple\"}],\"name\":\"permitDAIAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structTransferHandler.PermitOptions\",\"name\":\"options\",\"type\":\"tuple\"}],\"name\":\"permitEIP2612AndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TransferDemoABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferDemoMetaData.ABI instead.
var TransferDemoABI = TransferDemoMetaData.ABI

// TransferDemo is an auto generated Go binding around an Ethereum contract.
type TransferDemo struct {
	TransferDemoCaller     // Read-only binding to the contract
	TransferDemoTransactor // Write-only binding to the contract
	TransferDemoFilterer   // Log filterer for contract events
}

// TransferDemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferDemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferDemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferDemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferDemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferDemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferDemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferDemoSession struct {
	Contract     *TransferDemo     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferDemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferDemoCallerSession struct {
	Contract *TransferDemoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TransferDemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferDemoTransactorSession struct {
	Contract     *TransferDemoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TransferDemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferDemoRaw struct {
	Contract *TransferDemo // Generic contract binding to access the raw methods on
}

// TransferDemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferDemoCallerRaw struct {
	Contract *TransferDemoCaller // Generic read-only contract binding to access the raw methods on
}

// TransferDemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferDemoTransactorRaw struct {
	Contract *TransferDemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferDemo creates a new instance of TransferDemo, bound to a specific deployed contract.
func NewTransferDemo(address common.Address, backend bind.ContractBackend) (*TransferDemo, error) {
	contract, err := bindTransferDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferDemo{TransferDemoCaller: TransferDemoCaller{contract: contract}, TransferDemoTransactor: TransferDemoTransactor{contract: contract}, TransferDemoFilterer: TransferDemoFilterer{contract: contract}}, nil
}

// NewTransferDemoCaller creates a new read-only instance of TransferDemo, bound to a specific deployed contract.
func NewTransferDemoCaller(address common.Address, caller bind.ContractCaller) (*TransferDemoCaller, error) {
	contract, err := bindTransferDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferDemoCaller{contract: contract}, nil
}

// NewTransferDemoTransactor creates a new write-only instance of TransferDemo, bound to a specific deployed contract.
func NewTransferDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferDemoTransactor, error) {
	contract, err := bindTransferDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferDemoTransactor{contract: contract}, nil
}

// NewTransferDemoFilterer creates a new log filterer instance of TransferDemo, bound to a specific deployed contract.
func NewTransferDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferDemoFilterer, error) {
	contract, err := bindTransferDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferDemoFilterer{contract: contract}, nil
}

// bindTransferDemo binds a generic wrapper to an already deployed contract.
func bindTransferDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferDemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferDemo *TransferDemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferDemo.Contract.TransferDemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferDemo *TransferDemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferDemo.Contract.TransferDemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferDemo *TransferDemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferDemo.Contract.TransferDemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferDemo *TransferDemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferDemo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferDemo *TransferDemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferDemo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferDemo *TransferDemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferDemo.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TransferDemo *TransferDemoCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _TransferDemo.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TransferDemo *TransferDemoSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TransferDemo.Contract.IsTrustedForwarder(&_TransferDemo.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TransferDemo *TransferDemoCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TransferDemo.Contract.IsTrustedForwarder(&_TransferDemo.CallOpts, forwarder)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferDemo *TransferDemoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferDemo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferDemo *TransferDemoSession) Owner() (common.Address, error) {
	return _TransferDemo.Contract.Owner(&_TransferDemo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferDemo *TransferDemoCallerSession) Owner() (common.Address, error) {
	return _TransferDemo.Contract.Owner(&_TransferDemo.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TransferDemo *TransferDemoCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TransferDemo.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TransferDemo *TransferDemoSession) Paused() (bool, error) {
	return _TransferDemo.Contract.Paused(&_TransferDemo.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TransferDemo *TransferDemoCallerSession) Paused() (bool, error) {
	return _TransferDemo.Contract.Paused(&_TransferDemo.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_TransferDemo *TransferDemoCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TransferDemo.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_TransferDemo *TransferDemoSession) VersionRecipient() (string, error) {
	return _TransferDemo.Contract.VersionRecipient(&_TransferDemo.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_TransferDemo *TransferDemoCallerSession) VersionRecipient() (string, error) {
	return _TransferDemo.Contract.VersionRecipient(&_TransferDemo.CallOpts)
}

// PermitDAIAndTransfer is a paid mutator transaction binding the contract method 0xde7f1877.
//
// Solidity: function permitDAIAndTransfer(address token, address to, uint256 amount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) options) returns()
func (_TransferDemo *TransferDemoTransactor) PermitDAIAndTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, options TransferHandlerPermitOptions) (*types.Transaction, error) {
	return _TransferDemo.contract.Transact(opts, "permitDAIAndTransfer", token, to, amount, options)
}

// PermitDAIAndTransfer is a paid mutator transaction binding the contract method 0xde7f1877.
//
// Solidity: function permitDAIAndTransfer(address token, address to, uint256 amount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) options) returns()
func (_TransferDemo *TransferDemoSession) PermitDAIAndTransfer(token common.Address, to common.Address, amount *big.Int, options TransferHandlerPermitOptions) (*types.Transaction, error) {
	return _TransferDemo.Contract.PermitDAIAndTransfer(&_TransferDemo.TransactOpts, token, to, amount, options)
}

// PermitDAIAndTransfer is a paid mutator transaction binding the contract method 0xde7f1877.
//
// Solidity: function permitDAIAndTransfer(address token, address to, uint256 amount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) options) returns()
func (_TransferDemo *TransferDemoTransactorSession) PermitDAIAndTransfer(token common.Address, to common.Address, amount *big.Int, options TransferHandlerPermitOptions) (*types.Transaction, error) {
	return _TransferDemo.Contract.PermitDAIAndTransfer(&_TransferDemo.TransactOpts, token, to, amount, options)
}

// PermitEIP2612AndTransfer is a paid mutator transaction binding the contract method 0x71234eb0.
//
// Solidity: function permitEIP2612AndTransfer(address token, address to, uint256 amount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) options) returns()
func (_TransferDemo *TransferDemoTransactor) PermitEIP2612AndTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, options TransferHandlerPermitOptions) (*types.Transaction, error) {
	return _TransferDemo.contract.Transact(opts, "permitEIP2612AndTransfer", token, to, amount, options)
}

// PermitEIP2612AndTransfer is a paid mutator transaction binding the contract method 0x71234eb0.
//
// Solidity: function permitEIP2612AndTransfer(address token, address to, uint256 amount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) options) returns()
func (_TransferDemo *TransferDemoSession) PermitEIP2612AndTransfer(token common.Address, to common.Address, amount *big.Int, options TransferHandlerPermitOptions) (*types.Transaction, error) {
	return _TransferDemo.Contract.PermitEIP2612AndTransfer(&_TransferDemo.TransactOpts, token, to, amount, options)
}

// PermitEIP2612AndTransfer is a paid mutator transaction binding the contract method 0x71234eb0.
//
// Solidity: function permitEIP2612AndTransfer(address token, address to, uint256 amount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) options) returns()
func (_TransferDemo *TransferDemoTransactorSession) PermitEIP2612AndTransfer(token common.Address, to common.Address, amount *big.Int, options TransferHandlerPermitOptions) (*types.Transaction, error) {
	return _TransferDemo.Contract.PermitEIP2612AndTransfer(&_TransferDemo.TransactOpts, token, to, amount, options)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TransferDemo *TransferDemoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferDemo.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TransferDemo *TransferDemoSession) RenounceOwnership() (*types.Transaction, error) {
	return _TransferDemo.Contract.RenounceOwnership(&_TransferDemo.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TransferDemo *TransferDemoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TransferDemo.Contract.RenounceOwnership(&_TransferDemo.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_TransferDemo *TransferDemoTransactor) SetPause(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _TransferDemo.contract.Transact(opts, "setPause", val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_TransferDemo *TransferDemoSession) SetPause(val bool) (*types.Transaction, error) {
	return _TransferDemo.Contract.SetPause(&_TransferDemo.TransactOpts, val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_TransferDemo *TransferDemoTransactorSession) SetPause(val bool) (*types.Transaction, error) {
	return _TransferDemo.Contract.SetPause(&_TransferDemo.TransactOpts, val)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_TransferDemo *TransferDemoTransactor) Transfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TransferDemo.contract.Transact(opts, "transfer", token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_TransferDemo *TransferDemoSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TransferDemo.Contract.Transfer(&_TransferDemo.TransactOpts, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_TransferDemo *TransferDemoTransactorSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TransferDemo.Contract.Transfer(&_TransferDemo.TransactOpts, token, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferDemo *TransferDemoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TransferDemo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferDemo *TransferDemoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferDemo.Contract.TransferOwnership(&_TransferDemo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferDemo *TransferDemoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferDemo.Contract.TransferOwnership(&_TransferDemo.TransactOpts, newOwner)
}

// TransferDemoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TransferDemo contract.
type TransferDemoOwnershipTransferredIterator struct {
	Event *TransferDemoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TransferDemoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferDemoOwnershipTransferred)
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
		it.Event = new(TransferDemoOwnershipTransferred)
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
func (it *TransferDemoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferDemoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferDemoOwnershipTransferred represents a OwnershipTransferred event raised by the TransferDemo contract.
type TransferDemoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferDemo *TransferDemoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TransferDemoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferDemo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TransferDemoOwnershipTransferredIterator{contract: _TransferDemo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferDemo *TransferDemoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TransferDemoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferDemo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferDemoOwnershipTransferred)
				if err := _TransferDemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferDemo *TransferDemoFilterer) ParseOwnershipTransferred(log types.Log) (*TransferDemoOwnershipTransferred, error) {
	event := new(TransferDemoOwnershipTransferred)
	if err := _TransferDemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferDemoPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TransferDemo contract.
type TransferDemoPausedIterator struct {
	Event *TransferDemoPaused // Event containing the contract specifics and raw log

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
func (it *TransferDemoPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferDemoPaused)
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
		it.Event = new(TransferDemoPaused)
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
func (it *TransferDemoPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferDemoPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferDemoPaused represents a Paused event raised by the TransferDemo contract.
type TransferDemoPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TransferDemo *TransferDemoFilterer) FilterPaused(opts *bind.FilterOpts) (*TransferDemoPausedIterator, error) {

	logs, sub, err := _TransferDemo.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TransferDemoPausedIterator{contract: _TransferDemo.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TransferDemo *TransferDemoFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TransferDemoPaused) (event.Subscription, error) {

	logs, sub, err := _TransferDemo.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferDemoPaused)
				if err := _TransferDemo.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TransferDemo *TransferDemoFilterer) ParsePaused(log types.Log) (*TransferDemoPaused, error) {
	event := new(TransferDemoPaused)
	if err := _TransferDemo.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferDemoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TransferDemo contract.
type TransferDemoTransferIterator struct {
	Event *TransferDemoTransfer // Event containing the contract specifics and raw log

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
func (it *TransferDemoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferDemoTransfer)
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
		it.Event = new(TransferDemoTransfer)
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
func (it *TransferDemoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferDemoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferDemoTransfer represents a Transfer event raised by the TransferDemo contract.
type TransferDemoTransfer struct {
	Token  common.Address
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed token, address indexed from, address indexed to, uint256 amount)
func (_TransferDemo *TransferDemoFilterer) FilterTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*TransferDemoTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TransferDemo.contract.FilterLogs(opts, "Transfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TransferDemoTransferIterator{contract: _TransferDemo.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed token, address indexed from, address indexed to, uint256 amount)
func (_TransferDemo *TransferDemoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TransferDemoTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TransferDemo.contract.WatchLogs(opts, "Transfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferDemoTransfer)
				if err := _TransferDemo.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed token, address indexed from, address indexed to, uint256 amount)
func (_TransferDemo *TransferDemoFilterer) ParseTransfer(log types.Log) (*TransferDemoTransfer, error) {
	event := new(TransferDemoTransfer)
	if err := _TransferDemo.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferDemoUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TransferDemo contract.
type TransferDemoUnpausedIterator struct {
	Event *TransferDemoUnpaused // Event containing the contract specifics and raw log

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
func (it *TransferDemoUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferDemoUnpaused)
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
		it.Event = new(TransferDemoUnpaused)
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
func (it *TransferDemoUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferDemoUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferDemoUnpaused represents a Unpaused event raised by the TransferDemo contract.
type TransferDemoUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TransferDemo *TransferDemoFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TransferDemoUnpausedIterator, error) {

	logs, sub, err := _TransferDemo.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TransferDemoUnpausedIterator{contract: _TransferDemo.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TransferDemo *TransferDemoFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TransferDemoUnpaused) (event.Subscription, error) {

	logs, sub, err := _TransferDemo.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferDemoUnpaused)
				if err := _TransferDemo.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TransferDemo *TransferDemoFilterer) ParseUnpaused(log types.Log) (*TransferDemoUnpaused, error) {
	event := new(TransferDemoUnpaused)
	if err := _TransferDemo.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
