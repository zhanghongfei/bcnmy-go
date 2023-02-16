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

// ISwapRouterV3ExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterV3ExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// UniswapDemoMetaData contains all meta data concerning the UniswapDemo contract.
var UniswapDemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouterV3.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapDemoABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapDemoMetaData.ABI instead.
var UniswapDemoABI = UniswapDemoMetaData.ABI

// UniswapDemo is an auto generated Go binding around an Ethereum contract.
type UniswapDemo struct {
	UniswapDemoCaller     // Read-only binding to the contract
	UniswapDemoTransactor // Write-only binding to the contract
	UniswapDemoFilterer   // Log filterer for contract events
}

// UniswapDemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapDemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapDemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapDemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapDemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapDemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapDemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapDemoSession struct {
	Contract     *UniswapDemo      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapDemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapDemoCallerSession struct {
	Contract *UniswapDemoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UniswapDemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapDemoTransactorSession struct {
	Contract     *UniswapDemoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UniswapDemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapDemoRaw struct {
	Contract *UniswapDemo // Generic contract binding to access the raw methods on
}

// UniswapDemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapDemoCallerRaw struct {
	Contract *UniswapDemoCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapDemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapDemoTransactorRaw struct {
	Contract *UniswapDemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapDemo creates a new instance of UniswapDemo, bound to a specific deployed contract.
func NewUniswapDemo(address common.Address, backend bind.ContractBackend) (*UniswapDemo, error) {
	contract, err := bindUniswapDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapDemo{UniswapDemoCaller: UniswapDemoCaller{contract: contract}, UniswapDemoTransactor: UniswapDemoTransactor{contract: contract}, UniswapDemoFilterer: UniswapDemoFilterer{contract: contract}}, nil
}

// NewUniswapDemoCaller creates a new read-only instance of UniswapDemo, bound to a specific deployed contract.
func NewUniswapDemoCaller(address common.Address, caller bind.ContractCaller) (*UniswapDemoCaller, error) {
	contract, err := bindUniswapDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapDemoCaller{contract: contract}, nil
}

// NewUniswapDemoTransactor creates a new write-only instance of UniswapDemo, bound to a specific deployed contract.
func NewUniswapDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapDemoTransactor, error) {
	contract, err := bindUniswapDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapDemoTransactor{contract: contract}, nil
}

// NewUniswapDemoFilterer creates a new log filterer instance of UniswapDemo, bound to a specific deployed contract.
func NewUniswapDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapDemoFilterer, error) {
	contract, err := bindUniswapDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapDemoFilterer{contract: contract}, nil
}

// bindUniswapDemo binds a generic wrapper to an already deployed contract.
func bindUniswapDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapDemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapDemo *UniswapDemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapDemo.Contract.UniswapDemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapDemo *UniswapDemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapDemo.Contract.UniswapDemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapDemo *UniswapDemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapDemo.Contract.UniswapDemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapDemo *UniswapDemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapDemo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapDemo *UniswapDemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapDemo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapDemo *UniswapDemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapDemo.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_UniswapDemo *UniswapDemoCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _UniswapDemo.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_UniswapDemo *UniswapDemoSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _UniswapDemo.Contract.IsTrustedForwarder(&_UniswapDemo.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_UniswapDemo *UniswapDemoCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _UniswapDemo.Contract.IsTrustedForwarder(&_UniswapDemo.CallOpts, forwarder)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x1f02f4f2.
//
// Solidity: function exactInputSingle(address sender, (address,address,uint24,address,uint256,uint256,uint160) params) payable returns(bytes)
func (_UniswapDemo *UniswapDemoTransactor) ExactInputSingle(opts *bind.TransactOpts, sender common.Address, params ISwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapDemo.contract.Transact(opts, "exactInputSingle", sender, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x1f02f4f2.
//
// Solidity: function exactInputSingle(address sender, (address,address,uint24,address,uint256,uint256,uint160) params) payable returns(bytes)
func (_UniswapDemo *UniswapDemoSession) ExactInputSingle(sender common.Address, params ISwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapDemo.Contract.ExactInputSingle(&_UniswapDemo.TransactOpts, sender, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x1f02f4f2.
//
// Solidity: function exactInputSingle(address sender, (address,address,uint24,address,uint256,uint256,uint160) params) payable returns(bytes)
func (_UniswapDemo *UniswapDemoTransactorSession) ExactInputSingle(sender common.Address, params ISwapRouterV3ExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapDemo.Contract.ExactInputSingle(&_UniswapDemo.TransactOpts, sender, params)
}
