// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedBaseTokenABI is the input ABI used to generate the binding from.
const IdFedBaseTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdFedBaseTokenFuncSigs maps the 4-byte function signature to its string representation.
var IdFedBaseTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"9dc29fac": "burn(address,uint256)",
	"313ce567": "decimals()",
	"6e553f65": "deposit(uint256,address)",
	"40c10f19": "mint(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IdFedBaseToken is an auto generated Go binding around an Ethereum contract.
type IdFedBaseToken struct {
	IdFedBaseTokenCaller     // Read-only binding to the contract
	IdFedBaseTokenTransactor // Write-only binding to the contract
	IdFedBaseTokenFilterer   // Log filterer for contract events
}

// IdFedBaseTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdFedBaseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedBaseTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdFedBaseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedBaseTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdFedBaseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedBaseTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdFedBaseTokenSession struct {
	Contract     *IdFedBaseToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdFedBaseTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdFedBaseTokenCallerSession struct {
	Contract *IdFedBaseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IdFedBaseTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdFedBaseTokenTransactorSession struct {
	Contract     *IdFedBaseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IdFedBaseTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdFedBaseTokenRaw struct {
	Contract *IdFedBaseToken // Generic contract binding to access the raw methods on
}

// IdFedBaseTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdFedBaseTokenCallerRaw struct {
	Contract *IdFedBaseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IdFedBaseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdFedBaseTokenTransactorRaw struct {
	Contract *IdFedBaseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdFedBaseToken creates a new instance of IdFedBaseToken, bound to a specific deployed contract.
func NewIdFedBaseToken(address common.Address, backend bind.ContractBackend) (*IdFedBaseToken, error) {
	contract, err := bindIdFedBaseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdFedBaseToken{IdFedBaseTokenCaller: IdFedBaseTokenCaller{contract: contract}, IdFedBaseTokenTransactor: IdFedBaseTokenTransactor{contract: contract}, IdFedBaseTokenFilterer: IdFedBaseTokenFilterer{contract: contract}}, nil
}

// NewIdFedBaseTokenCaller creates a new read-only instance of IdFedBaseToken, bound to a specific deployed contract.
func NewIdFedBaseTokenCaller(address common.Address, caller bind.ContractCaller) (*IdFedBaseTokenCaller, error) {
	contract, err := bindIdFedBaseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedBaseTokenCaller{contract: contract}, nil
}

// NewIdFedBaseTokenTransactor creates a new write-only instance of IdFedBaseToken, bound to a specific deployed contract.
func NewIdFedBaseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IdFedBaseTokenTransactor, error) {
	contract, err := bindIdFedBaseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedBaseTokenTransactor{contract: contract}, nil
}

// NewIdFedBaseTokenFilterer creates a new log filterer instance of IdFedBaseToken, bound to a specific deployed contract.
func NewIdFedBaseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IdFedBaseTokenFilterer, error) {
	contract, err := bindIdFedBaseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdFedBaseTokenFilterer{contract: contract}, nil
}

// bindIdFedBaseToken binds a generic wrapper to an already deployed contract.
func bindIdFedBaseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdFedBaseTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedBaseToken *IdFedBaseTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedBaseToken.Contract.IdFedBaseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedBaseToken *IdFedBaseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.IdFedBaseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedBaseToken *IdFedBaseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.IdFedBaseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedBaseToken *IdFedBaseTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedBaseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedBaseToken *IdFedBaseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedBaseToken *IdFedBaseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedBaseToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IdFedBaseToken.Contract.Allowance(&_IdFedBaseToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IdFedBaseToken.Contract.Allowance(&_IdFedBaseToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedBaseToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdFedBaseToken.Contract.BalanceOf(&_IdFedBaseToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdFedBaseToken.Contract.BalanceOf(&_IdFedBaseToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IdFedBaseToken *IdFedBaseTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IdFedBaseToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IdFedBaseToken *IdFedBaseTokenSession) Decimals() (uint8, error) {
	return _IdFedBaseToken.Contract.Decimals(&_IdFedBaseToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IdFedBaseToken *IdFedBaseTokenCallerSession) Decimals() (uint8, error) {
	return _IdFedBaseToken.Contract.Decimals(&_IdFedBaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdFedBaseToken *IdFedBaseTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IdFedBaseToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdFedBaseToken *IdFedBaseTokenSession) Name() (string, error) {
	return _IdFedBaseToken.Contract.Name(&_IdFedBaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdFedBaseToken *IdFedBaseTokenCallerSession) Name() (string, error) {
	return _IdFedBaseToken.Contract.Name(&_IdFedBaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IdFedBaseToken *IdFedBaseTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IdFedBaseToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IdFedBaseToken *IdFedBaseTokenSession) Symbol() (string, error) {
	return _IdFedBaseToken.Contract.Symbol(&_IdFedBaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IdFedBaseToken *IdFedBaseTokenCallerSession) Symbol() (string, error) {
	return _IdFedBaseToken.Contract.Symbol(&_IdFedBaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedBaseToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenSession) TotalSupply() (*big.Int, error) {
	return _IdFedBaseToken.Contract.TotalSupply(&_IdFedBaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedBaseToken *IdFedBaseTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IdFedBaseToken.Contract.TotalSupply(&_IdFedBaseToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Approve(&_IdFedBaseToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Approve(&_IdFedBaseToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_IdFedBaseToken *IdFedBaseTokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.contract.Transact(opts, "burn", from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_IdFedBaseToken *IdFedBaseTokenSession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Burn(&_IdFedBaseToken.TransactOpts, from, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (_IdFedBaseToken *IdFedBaseTokenTransactorSession) Burn(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Burn(&_IdFedBaseToken.TransactOpts, from, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address _to) returns()
func (_IdFedBaseToken *IdFedBaseTokenTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IdFedBaseToken.contract.Transact(opts, "deposit", amount, _to)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address _to) returns()
func (_IdFedBaseToken *IdFedBaseTokenSession) Deposit(amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Deposit(&_IdFedBaseToken.TransactOpts, amount, _to)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address _to) returns()
func (_IdFedBaseToken *IdFedBaseTokenTransactorSession) Deposit(amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Deposit(&_IdFedBaseToken.TransactOpts, amount, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_IdFedBaseToken *IdFedBaseTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_IdFedBaseToken *IdFedBaseTokenSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Mint(&_IdFedBaseToken.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_IdFedBaseToken *IdFedBaseTokenTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Mint(&_IdFedBaseToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Transfer(&_IdFedBaseToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.Transfer(&_IdFedBaseToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.TransferFrom(&_IdFedBaseToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedBaseToken *IdFedBaseTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedBaseToken.Contract.TransferFrom(&_IdFedBaseToken.TransactOpts, from, to, value)
}

// IdFedBaseTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IdFedBaseToken contract.
type IdFedBaseTokenApprovalIterator struct {
	Event *IdFedBaseTokenApproval // Event containing the contract specifics and raw log

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
func (it *IdFedBaseTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedBaseTokenApproval)
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
		it.Event = new(IdFedBaseTokenApproval)
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
func (it *IdFedBaseTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedBaseTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedBaseTokenApproval represents a Approval event raised by the IdFedBaseToken contract.
type IdFedBaseTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedBaseToken *IdFedBaseTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IdFedBaseTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IdFedBaseToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IdFedBaseTokenApprovalIterator{contract: _IdFedBaseToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedBaseToken *IdFedBaseTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IdFedBaseTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IdFedBaseToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedBaseTokenApproval)
				if err := _IdFedBaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedBaseToken *IdFedBaseTokenFilterer) ParseApproval(log types.Log) (*IdFedBaseTokenApproval, error) {
	event := new(IdFedBaseTokenApproval)
	if err := _IdFedBaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedBaseTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IdFedBaseToken contract.
type IdFedBaseTokenTransferIterator struct {
	Event *IdFedBaseTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IdFedBaseTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedBaseTokenTransfer)
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
		it.Event = new(IdFedBaseTokenTransfer)
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
func (it *IdFedBaseTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedBaseTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedBaseTokenTransfer represents a Transfer event raised by the IdFedBaseToken contract.
type IdFedBaseTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedBaseToken *IdFedBaseTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IdFedBaseTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedBaseToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IdFedBaseTokenTransferIterator{contract: _IdFedBaseToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedBaseToken *IdFedBaseTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IdFedBaseTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedBaseToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedBaseTokenTransfer)
				if err := _IdFedBaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedBaseToken *IdFedBaseTokenFilterer) ParseTransfer(log types.Log) (*IdFedBaseTokenTransfer, error) {
	event := new(IdFedBaseTokenTransfer)
	if err := _IdFedBaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedCreatorABI is the input ABI used to generate the binding from.
const IdFedCreatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdFedCreatorFuncSigs maps the 4-byte function signature to its string representation.
var IdFedCreatorFuncSigs = map[string]string{
	"c9c65396": "createPair(address,address)",
}

// IdFedCreator is an auto generated Go binding around an Ethereum contract.
type IdFedCreator struct {
	IdFedCreatorCaller     // Read-only binding to the contract
	IdFedCreatorTransactor // Write-only binding to the contract
	IdFedCreatorFilterer   // Log filterer for contract events
}

// IdFedCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdFedCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdFedCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdFedCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdFedCreatorSession struct {
	Contract     *IdFedCreator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdFedCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdFedCreatorCallerSession struct {
	Contract *IdFedCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IdFedCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdFedCreatorTransactorSession struct {
	Contract     *IdFedCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IdFedCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdFedCreatorRaw struct {
	Contract *IdFedCreator // Generic contract binding to access the raw methods on
}

// IdFedCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdFedCreatorCallerRaw struct {
	Contract *IdFedCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// IdFedCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdFedCreatorTransactorRaw struct {
	Contract *IdFedCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdFedCreator creates a new instance of IdFedCreator, bound to a specific deployed contract.
func NewIdFedCreator(address common.Address, backend bind.ContractBackend) (*IdFedCreator, error) {
	contract, err := bindIdFedCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdFedCreator{IdFedCreatorCaller: IdFedCreatorCaller{contract: contract}, IdFedCreatorTransactor: IdFedCreatorTransactor{contract: contract}, IdFedCreatorFilterer: IdFedCreatorFilterer{contract: contract}}, nil
}

// NewIdFedCreatorCaller creates a new read-only instance of IdFedCreator, bound to a specific deployed contract.
func NewIdFedCreatorCaller(address common.Address, caller bind.ContractCaller) (*IdFedCreatorCaller, error) {
	contract, err := bindIdFedCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedCreatorCaller{contract: contract}, nil
}

// NewIdFedCreatorTransactor creates a new write-only instance of IdFedCreator, bound to a specific deployed contract.
func NewIdFedCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IdFedCreatorTransactor, error) {
	contract, err := bindIdFedCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedCreatorTransactor{contract: contract}, nil
}

// NewIdFedCreatorFilterer creates a new log filterer instance of IdFedCreator, bound to a specific deployed contract.
func NewIdFedCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IdFedCreatorFilterer, error) {
	contract, err := bindIdFedCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdFedCreatorFilterer{contract: contract}, nil
}

// bindIdFedCreator binds a generic wrapper to an already deployed contract.
func bindIdFedCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdFedCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedCreator *IdFedCreatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedCreator.Contract.IdFedCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedCreator *IdFedCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedCreator.Contract.IdFedCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedCreator *IdFedCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedCreator.Contract.IdFedCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedCreator *IdFedCreatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedCreator *IdFedCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedCreator *IdFedCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedCreator.Contract.contract.Transact(opts, method, params...)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address token, address baseToken) returns(address)
func (_IdFedCreator *IdFedCreatorTransactor) CreatePair(opts *bind.TransactOpts, token common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _IdFedCreator.contract.Transact(opts, "createPair", token, baseToken)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address token, address baseToken) returns(address)
func (_IdFedCreator *IdFedCreatorSession) CreatePair(token common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _IdFedCreator.Contract.CreatePair(&_IdFedCreator.TransactOpts, token, baseToken)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address token, address baseToken) returns(address)
func (_IdFedCreator *IdFedCreatorTransactorSession) CreatePair(token common.Address, baseToken common.Address) (*types.Transaction, error) {
	return _IdFedCreator.Contract.CreatePair(&_IdFedCreator.TransactOpts, token, baseToken)
}

// IdFedPairABI is the input ABI used to generate the binding from.
const IdFedPairABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToken0Amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToken1Amount\",\"type\":\"uint256\"}],\"name\":\"DebtUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetNum\",\"type\":\"uint256\"}],\"name\":\"Mortgage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getDebtIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRepayByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mortgage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdFedPairFuncSigs maps the 4-byte function signature to its string representation.
var IdFedPairFuncSigs = map[string]string{
	"3644e515": "DOMAIN_SEPARATOR()",
	"30adf81f": "PERMIT_TYPEHASH()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"89afcb44": "burn(address)",
	"21466eb5": "buyToken(uint256,uint256,address)",
	"313ce567": "decimals()",
	"3c33a245": "getDebtIndexById(uint256)",
	"ea52a0be": "getRepayByIndex(uint256)",
	"0902f1ac": "getReserves()",
	"485cc955": "initialize(address,address)",
	"156e29f6": "mint(address,uint256,uint256)",
	"58687e63": "mortgage(uint256,uint256,address)",
	"06fdde03": "name()",
	"7ecebe00": "nonces(address)",
	"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	"371fd8e6": "repay(uint256)",
	"88036ac5": "sellToken(uint256,uint256,address)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IdFedPair is an auto generated Go binding around an Ethereum contract.
type IdFedPair struct {
	IdFedPairCaller     // Read-only binding to the contract
	IdFedPairTransactor // Write-only binding to the contract
	IdFedPairFilterer   // Log filterer for contract events
}

// IdFedPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdFedPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdFedPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdFedPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdFedPairSession struct {
	Contract     *IdFedPair        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdFedPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdFedPairCallerSession struct {
	Contract *IdFedPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IdFedPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdFedPairTransactorSession struct {
	Contract     *IdFedPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IdFedPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdFedPairRaw struct {
	Contract *IdFedPair // Generic contract binding to access the raw methods on
}

// IdFedPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdFedPairCallerRaw struct {
	Contract *IdFedPairCaller // Generic read-only contract binding to access the raw methods on
}

// IdFedPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdFedPairTransactorRaw struct {
	Contract *IdFedPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdFedPair creates a new instance of IdFedPair, bound to a specific deployed contract.
func NewIdFedPair(address common.Address, backend bind.ContractBackend) (*IdFedPair, error) {
	contract, err := bindIdFedPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdFedPair{IdFedPairCaller: IdFedPairCaller{contract: contract}, IdFedPairTransactor: IdFedPairTransactor{contract: contract}, IdFedPairFilterer: IdFedPairFilterer{contract: contract}}, nil
}

// NewIdFedPairCaller creates a new read-only instance of IdFedPair, bound to a specific deployed contract.
func NewIdFedPairCaller(address common.Address, caller bind.ContractCaller) (*IdFedPairCaller, error) {
	contract, err := bindIdFedPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedPairCaller{contract: contract}, nil
}

// NewIdFedPairTransactor creates a new write-only instance of IdFedPair, bound to a specific deployed contract.
func NewIdFedPairTransactor(address common.Address, transactor bind.ContractTransactor) (*IdFedPairTransactor, error) {
	contract, err := bindIdFedPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedPairTransactor{contract: contract}, nil
}

// NewIdFedPairFilterer creates a new log filterer instance of IdFedPair, bound to a specific deployed contract.
func NewIdFedPairFilterer(address common.Address, filterer bind.ContractFilterer) (*IdFedPairFilterer, error) {
	contract, err := bindIdFedPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdFedPairFilterer{contract: contract}, nil
}

// bindIdFedPair binds a generic wrapper to an already deployed contract.
func bindIdFedPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdFedPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedPair *IdFedPairRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedPair.Contract.IdFedPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedPair *IdFedPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedPair.Contract.IdFedPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedPair *IdFedPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedPair.Contract.IdFedPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedPair *IdFedPairCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedPair *IdFedPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedPair *IdFedPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedPair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IdFedPair *IdFedPairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IdFedPair *IdFedPairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IdFedPair.Contract.DOMAINSEPARATOR(&_IdFedPair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IdFedPair *IdFedPairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IdFedPair.Contract.DOMAINSEPARATOR(&_IdFedPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IdFedPair *IdFedPairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IdFedPair *IdFedPairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IdFedPair.Contract.PERMITTYPEHASH(&_IdFedPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IdFedPair *IdFedPairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IdFedPair.Contract.PERMITTYPEHASH(&_IdFedPair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedPair *IdFedPairCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedPair *IdFedPairSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IdFedPair.Contract.Allowance(&_IdFedPair.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedPair *IdFedPairCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IdFedPair.Contract.Allowance(&_IdFedPair.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedPair *IdFedPairCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedPair *IdFedPairSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdFedPair.Contract.BalanceOf(&_IdFedPair.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedPair *IdFedPairCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdFedPair.Contract.BalanceOf(&_IdFedPair.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IdFedPair *IdFedPairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IdFedPair *IdFedPairSession) Decimals() (uint8, error) {
	return _IdFedPair.Contract.Decimals(&_IdFedPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IdFedPair *IdFedPairCallerSession) Decimals() (uint8, error) {
	return _IdFedPair.Contract.Decimals(&_IdFedPair.CallOpts)
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 id) view returns(uint256 index)
func (_IdFedPair *IdFedPairCaller) GetDebtIndexById(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "getDebtIndexById", id)
	return *ret0, err
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 id) view returns(uint256 index)
func (_IdFedPair *IdFedPairSession) GetDebtIndexById(id *big.Int) (*big.Int, error) {
	return _IdFedPair.Contract.GetDebtIndexById(&_IdFedPair.CallOpts, id)
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 id) view returns(uint256 index)
func (_IdFedPair *IdFedPairCallerSession) GetDebtIndexById(id *big.Int) (*big.Int, error) {
	return _IdFedPair.Contract.GetDebtIndexById(&_IdFedPair.CallOpts, id)
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 index) view returns(uint256 repayAmount)
func (_IdFedPair *IdFedPairCaller) GetRepayByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "getRepayByIndex", index)
	return *ret0, err
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 index) view returns(uint256 repayAmount)
func (_IdFedPair *IdFedPairSession) GetRepayByIndex(index *big.Int) (*big.Int, error) {
	return _IdFedPair.Contract.GetRepayByIndex(&_IdFedPair.CallOpts, index)
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 index) view returns(uint256 repayAmount)
func (_IdFedPair *IdFedPairCallerSession) GetRepayByIndex(index *big.Int) (*big.Int, error) {
	return _IdFedPair.Contract.GetRepayByIndex(&_IdFedPair.CallOpts, index)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserve0, uint256 reserve1)
func (_IdFedPair *IdFedPairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	ret := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	out := ret
	err := _IdFedPair.contract.Call(opts, out, "getReserves")
	return *ret, err
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserve0, uint256 reserve1)
func (_IdFedPair *IdFedPairSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _IdFedPair.Contract.GetReserves(&_IdFedPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserve0, uint256 reserve1)
func (_IdFedPair *IdFedPairCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _IdFedPair.Contract.GetReserves(&_IdFedPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IdFedPair *IdFedPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IdFedPair *IdFedPairSession) Name() (string, error) {
	return _IdFedPair.Contract.Name(&_IdFedPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IdFedPair *IdFedPairCallerSession) Name() (string, error) {
	return _IdFedPair.Contract.Name(&_IdFedPair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdFedPair *IdFedPairCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "nonces", owner)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdFedPair *IdFedPairSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IdFedPair.Contract.Nonces(&_IdFedPair.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdFedPair *IdFedPairCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IdFedPair.Contract.Nonces(&_IdFedPair.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IdFedPair *IdFedPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IdFedPair *IdFedPairSession) Symbol() (string, error) {
	return _IdFedPair.Contract.Symbol(&_IdFedPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IdFedPair *IdFedPairCallerSession) Symbol() (string, error) {
	return _IdFedPair.Contract.Symbol(&_IdFedPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedPair *IdFedPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedPair.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedPair *IdFedPairSession) TotalSupply() (*big.Int, error) {
	return _IdFedPair.Contract.TotalSupply(&_IdFedPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedPair *IdFedPairCallerSession) TotalSupply() (*big.Int, error) {
	return _IdFedPair.Contract.TotalSupply(&_IdFedPair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Approve(&_IdFedPair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Approve(&_IdFedPair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IdFedPair *IdFedPairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IdFedPair *IdFedPairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.Burn(&_IdFedPair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_IdFedPair *IdFedPairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.Burn(&_IdFedPair.TransactOpts, to)
}

// BuyToken is a paid mutator transaction binding the contract method 0x21466eb5.
//
// Solidity: function buyToken(uint256 amountIn, uint256 amountOutMin, address to) returns()
func (_IdFedPair *IdFedPairTransactor) BuyToken(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "buyToken", amountIn, amountOutMin, to)
}

// BuyToken is a paid mutator transaction binding the contract method 0x21466eb5.
//
// Solidity: function buyToken(uint256 amountIn, uint256 amountOutMin, address to) returns()
func (_IdFedPair *IdFedPairSession) BuyToken(amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.BuyToken(&_IdFedPair.TransactOpts, amountIn, amountOutMin, to)
}

// BuyToken is a paid mutator transaction binding the contract method 0x21466eb5.
//
// Solidity: function buyToken(uint256 amountIn, uint256 amountOutMin, address to) returns()
func (_IdFedPair *IdFedPairTransactorSession) BuyToken(amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.BuyToken(&_IdFedPair.TransactOpts, amountIn, amountOutMin, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IdFedPair *IdFedPairTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "initialize", arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IdFedPair *IdFedPairSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.Initialize(&_IdFedPair.TransactOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_IdFedPair *IdFedPairTransactorSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.Initialize(&_IdFedPair.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount0, uint256 amount1) returns(uint256 liquidity)
func (_IdFedPair *IdFedPairTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "mint", to, amount0, amount1)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount0, uint256 amount1) returns(uint256 liquidity)
func (_IdFedPair *IdFedPairSession) Mint(to common.Address, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Mint(&_IdFedPair.TransactOpts, to, amount0, amount1)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount0, uint256 amount1) returns(uint256 liquidity)
func (_IdFedPair *IdFedPairTransactorSession) Mint(to common.Address, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Mint(&_IdFedPair.TransactOpts, to, amount0, amount1)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 pledgeAmount, uint256 targetNum, address to) returns(uint256 _debtId)
func (_IdFedPair *IdFedPairTransactor) Mortgage(opts *bind.TransactOpts, pledgeAmount *big.Int, targetNum *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "mortgage", pledgeAmount, targetNum, to)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 pledgeAmount, uint256 targetNum, address to) returns(uint256 _debtId)
func (_IdFedPair *IdFedPairSession) Mortgage(pledgeAmount *big.Int, targetNum *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.Mortgage(&_IdFedPair.TransactOpts, pledgeAmount, targetNum, to)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 pledgeAmount, uint256 targetNum, address to) returns(uint256 _debtId)
func (_IdFedPair *IdFedPairTransactorSession) Mortgage(pledgeAmount *big.Int, targetNum *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.Mortgage(&_IdFedPair.TransactOpts, pledgeAmount, targetNum, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IdFedPair *IdFedPairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IdFedPair *IdFedPairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdFedPair.Contract.Permit(&_IdFedPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IdFedPair *IdFedPairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdFedPair.Contract.Permit(&_IdFedPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 debtId) returns()
func (_IdFedPair *IdFedPairTransactor) Repay(opts *bind.TransactOpts, debtId *big.Int) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "repay", debtId)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 debtId) returns()
func (_IdFedPair *IdFedPairSession) Repay(debtId *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Repay(&_IdFedPair.TransactOpts, debtId)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 debtId) returns()
func (_IdFedPair *IdFedPairTransactorSession) Repay(debtId *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Repay(&_IdFedPair.TransactOpts, debtId)
}

// SellToken is a paid mutator transaction binding the contract method 0x88036ac5.
//
// Solidity: function sellToken(uint256 amountIn, uint256 amountOutMin, address to) returns()
func (_IdFedPair *IdFedPairTransactor) SellToken(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "sellToken", amountIn, amountOutMin, to)
}

// SellToken is a paid mutator transaction binding the contract method 0x88036ac5.
//
// Solidity: function sellToken(uint256 amountIn, uint256 amountOutMin, address to) returns()
func (_IdFedPair *IdFedPairSession) SellToken(amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.SellToken(&_IdFedPair.TransactOpts, amountIn, amountOutMin, to)
}

// SellToken is a paid mutator transaction binding the contract method 0x88036ac5.
//
// Solidity: function sellToken(uint256 amountIn, uint256 amountOutMin, address to) returns()
func (_IdFedPair *IdFedPairTransactorSession) SellToken(amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _IdFedPair.Contract.SellToken(&_IdFedPair.TransactOpts, amountIn, amountOutMin, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Transfer(&_IdFedPair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.Transfer(&_IdFedPair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.TransferFrom(&_IdFedPair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedPair *IdFedPairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedPair.Contract.TransferFrom(&_IdFedPair.TransactOpts, from, to, value)
}

// IdFedPairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IdFedPair contract.
type IdFedPairApprovalIterator struct {
	Event *IdFedPairApproval // Event containing the contract specifics and raw log

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
func (it *IdFedPairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedPairApproval)
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
		it.Event = new(IdFedPairApproval)
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
func (it *IdFedPairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedPairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedPairApproval represents a Approval event raised by the IdFedPair contract.
type IdFedPairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedPair *IdFedPairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IdFedPairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IdFedPair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IdFedPairApprovalIterator{contract: _IdFedPair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedPair *IdFedPairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IdFedPairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IdFedPair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedPairApproval)
				if err := _IdFedPair.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedPair *IdFedPairFilterer) ParseApproval(log types.Log) (*IdFedPairApproval, error) {
	event := new(IdFedPairApproval)
	if err := _IdFedPair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedPairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IdFedPair contract.
type IdFedPairBurnIterator struct {
	Event *IdFedPairBurn // Event containing the contract specifics and raw log

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
func (it *IdFedPairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedPairBurn)
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
		it.Event = new(IdFedPairBurn)
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
func (it *IdFedPairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedPairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedPairBurn represents a Burn event raised by the IdFedPair contract.
type IdFedPairBurn struct {
	Sender   common.Address
	LpAmount *big.Int
	Amount0  *big.Int
	Amount1  *big.Int
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 lpAmount, uint256 amount0, uint256 amount1, address indexed to)
func (_IdFedPair *IdFedPairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IdFedPairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedPair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IdFedPairBurnIterator{contract: _IdFedPair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 lpAmount, uint256 amount0, uint256 amount1, address indexed to)
func (_IdFedPair *IdFedPairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IdFedPairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedPair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedPairBurn)
				if err := _IdFedPair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 lpAmount, uint256 amount0, uint256 amount1, address indexed to)
func (_IdFedPair *IdFedPairFilterer) ParseBurn(log types.Log) (*IdFedPairBurn, error) {
	event := new(IdFedPairBurn)
	if err := _IdFedPair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedPairDebtUpdateIterator is returned from FilterDebtUpdate and is used to iterate over the raw logs and unpacked data for DebtUpdate events raised by the IdFedPair contract.
type IdFedPairDebtUpdateIterator struct {
	Event *IdFedPairDebtUpdate // Event containing the contract specifics and raw log

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
func (it *IdFedPairDebtUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedPairDebtUpdate)
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
		it.Event = new(IdFedPairDebtUpdate)
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
func (it *IdFedPairDebtUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedPairDebtUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedPairDebtUpdate represents a DebtUpdate event raised by the IdFedPair contract.
type IdFedPairDebtUpdate struct {
	Owner            common.Address
	DebtId           *big.Int
	PledgeAmount     *big.Int
	RepayAmount      *big.Int
	DebtToken0Amount *big.Int
	DebtToken1Amount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDebtUpdate is a free log retrieval operation binding the contract event 0xa0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869.
//
// Solidity: event DebtUpdate(address indexed owner, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount)
func (_IdFedPair *IdFedPairFilterer) FilterDebtUpdate(opts *bind.FilterOpts, owner []common.Address) (*IdFedPairDebtUpdateIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdFedPair.contract.FilterLogs(opts, "DebtUpdate", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IdFedPairDebtUpdateIterator{contract: _IdFedPair.contract, event: "DebtUpdate", logs: logs, sub: sub}, nil
}

// WatchDebtUpdate is a free log subscription operation binding the contract event 0xa0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869.
//
// Solidity: event DebtUpdate(address indexed owner, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount)
func (_IdFedPair *IdFedPairFilterer) WatchDebtUpdate(opts *bind.WatchOpts, sink chan<- *IdFedPairDebtUpdate, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdFedPair.contract.WatchLogs(opts, "DebtUpdate", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedPairDebtUpdate)
				if err := _IdFedPair.contract.UnpackLog(event, "DebtUpdate", log); err != nil {
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

// ParseDebtUpdate is a log parse operation binding the contract event 0xa0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869.
//
// Solidity: event DebtUpdate(address indexed owner, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount)
func (_IdFedPair *IdFedPairFilterer) ParseDebtUpdate(log types.Log) (*IdFedPairDebtUpdate, error) {
	event := new(IdFedPairDebtUpdate)
	if err := _IdFedPair.contract.UnpackLog(event, "DebtUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedPairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IdFedPair contract.
type IdFedPairMintIterator struct {
	Event *IdFedPairMint // Event containing the contract specifics and raw log

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
func (it *IdFedPairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedPairMint)
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
		it.Event = new(IdFedPairMint)
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
func (it *IdFedPairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedPairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedPairMint represents a Mint event raised by the IdFedPair contract.
type IdFedPairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IdFedPair *IdFedPairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*IdFedPairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IdFedPair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &IdFedPairMintIterator{contract: _IdFedPair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IdFedPair *IdFedPairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IdFedPairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IdFedPair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedPairMint)
				if err := _IdFedPair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_IdFedPair *IdFedPairFilterer) ParseMint(log types.Log) (*IdFedPairMint, error) {
	event := new(IdFedPairMint)
	if err := _IdFedPair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedPairMortgageIterator is returned from FilterMortgage and is used to iterate over the raw logs and unpacked data for Mortgage events raised by the IdFedPair contract.
type IdFedPairMortgageIterator struct {
	Event *IdFedPairMortgage // Event containing the contract specifics and raw log

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
func (it *IdFedPairMortgageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedPairMortgage)
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
		it.Event = new(IdFedPairMortgage)
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
func (it *IdFedPairMortgageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedPairMortgageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedPairMortgage represents a Mortgage event raised by the IdFedPair contract.
type IdFedPairMortgage struct {
	Owner        common.Address
	PledgeAmount *big.Int
	TargetNum    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMortgage is a free log retrieval operation binding the contract event 0x17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a4.
//
// Solidity: event Mortgage(address owner, uint256 pledgeAmount, uint256 targetNum)
func (_IdFedPair *IdFedPairFilterer) FilterMortgage(opts *bind.FilterOpts) (*IdFedPairMortgageIterator, error) {

	logs, sub, err := _IdFedPair.contract.FilterLogs(opts, "Mortgage")
	if err != nil {
		return nil, err
	}
	return &IdFedPairMortgageIterator{contract: _IdFedPair.contract, event: "Mortgage", logs: logs, sub: sub}, nil
}

// WatchMortgage is a free log subscription operation binding the contract event 0x17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a4.
//
// Solidity: event Mortgage(address owner, uint256 pledgeAmount, uint256 targetNum)
func (_IdFedPair *IdFedPairFilterer) WatchMortgage(opts *bind.WatchOpts, sink chan<- *IdFedPairMortgage) (event.Subscription, error) {

	logs, sub, err := _IdFedPair.contract.WatchLogs(opts, "Mortgage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedPairMortgage)
				if err := _IdFedPair.contract.UnpackLog(event, "Mortgage", log); err != nil {
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

// ParseMortgage is a log parse operation binding the contract event 0x17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a4.
//
// Solidity: event Mortgage(address owner, uint256 pledgeAmount, uint256 targetNum)
func (_IdFedPair *IdFedPairFilterer) ParseMortgage(log types.Log) (*IdFedPairMortgage, error) {
	event := new(IdFedPairMortgage)
	if err := _IdFedPair.contract.UnpackLog(event, "Mortgage", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedPairRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the IdFedPair contract.
type IdFedPairRepayIterator struct {
	Event *IdFedPairRepay // Event containing the contract specifics and raw log

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
func (it *IdFedPairRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedPairRepay)
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
		it.Event = new(IdFedPairRepay)
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
func (it *IdFedPairRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedPairRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedPairRepay represents a Repay event raised by the IdFedPair contract.
type IdFedPairRepay struct {
	DebtId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe46230.
//
// Solidity: event Repay(uint256 debtId)
func (_IdFedPair *IdFedPairFilterer) FilterRepay(opts *bind.FilterOpts) (*IdFedPairRepayIterator, error) {

	logs, sub, err := _IdFedPair.contract.FilterLogs(opts, "Repay")
	if err != nil {
		return nil, err
	}
	return &IdFedPairRepayIterator{contract: _IdFedPair.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe46230.
//
// Solidity: event Repay(uint256 debtId)
func (_IdFedPair *IdFedPairFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *IdFedPairRepay) (event.Subscription, error) {

	logs, sub, err := _IdFedPair.contract.WatchLogs(opts, "Repay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedPairRepay)
				if err := _IdFedPair.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe46230.
//
// Solidity: event Repay(uint256 debtId)
func (_IdFedPair *IdFedPairFilterer) ParseRepay(log types.Log) (*IdFedPairRepay, error) {
	event := new(IdFedPairRepay)
	if err := _IdFedPair.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedPairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IdFedPair contract.
type IdFedPairTransferIterator struct {
	Event *IdFedPairTransfer // Event containing the contract specifics and raw log

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
func (it *IdFedPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedPairTransfer)
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
		it.Event = new(IdFedPairTransfer)
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
func (it *IdFedPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedPairTransfer represents a Transfer event raised by the IdFedPair contract.
type IdFedPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedPair *IdFedPairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IdFedPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedPair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IdFedPairTransferIterator{contract: _IdFedPair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedPair *IdFedPairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IdFedPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedPair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedPairTransfer)
				if err := _IdFedPair.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedPair *IdFedPairFilterer) ParseTransfer(log types.Log) (*IdFedPairTransfer, error) {
	event := new(IdFedPairTransfer)
	if err := _IdFedPair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204148c8397a1a234dac2b69b8d5b8aca6df9d3ed34a161eb26934f4658f3ca9cb64736f6c63430007000033"

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// DFedFactoryABI is the input ABI used to generate the binding from.
const DFedFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accRewardPerUSDDGlobal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addTotalUSDDinLiquidityPoolGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardedBlockGlobal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mintBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mintFEDToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"removeTotalUSDDinLiquidityPoolGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"setBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"setCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_FEDToken\",\"type\":\"address\"}],\"name\":\"setFedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"setStartRewardBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUSDDinLiquidityPoolGlobal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DFedFactoryFuncSigs maps the 4-byte function signature to its string representation.
var DFedFactoryFuncSigs = map[string]string{
	"35ab4528": "FEDToken()",
	"00dab946": "accRewardPerUSDDGlobal()",
	"df6c1de1": "addTotalUSDDinLiquidityPoolGlobal(uint256)",
	"1e3dd18b": "allPairs(uint256)",
	"574f2ba3": "allPairsLength()",
	"c55dae63": "baseToken()",
	"0ed01e34": "burnBaseToken(address,uint256)",
	"9ccb0744": "createPair(address)",
	"02d05d3f": "creator()",
	"1a788a02": "getPair(address)",
	"01fb0a4f": "lastRewardedBlockGlobal()",
	"f8d1006a": "mintBaseToken(address,uint256)",
	"ccd072e5": "mintFEDToken(address,uint256)",
	"ac393aff": "pairExist(address)",
	"daaa88bd": "removeTotalUSDDinLiquidityPoolGlobal(uint256)",
	"16bb6c13": "setBaseToken(address)",
	"3f516018": "setCreator(address)",
	"59d17cdc": "setFedToken(address)",
	"62b9e098": "setStartRewardBlock(uint256)",
	"3f3108f7": "setter()",
	"d7ed714e": "startRewardBlock()",
	"53816b40": "totalUSDDinLiquidityPoolGlobal()",
	"b7c9252c": "updateInfo()",
}

// DFedFactoryBin is the compiled bytecode used for deploying new contracts.
var DFedFactoryBin = "0x608060405234801561001057600080fd5b50600680546001600160a01b031916331790556000196007819055600a55610f7a8061003d6000396000f3fe608060405234801561001057600080fd5b506004361061014c5760003560e01c8063574f2ba3116100c3578063c55dae631161007c578063c55dae631461031f578063ccd072e514610327578063d7ed714e14610353578063daaa88bd1461035b578063df6c1de114610378578063f8d1006a146103955761014c565b8063574f2ba31461026c57806359d17cdc1461027457806362b9e0981461029a5780639ccb0744146102b7578063ac393aff146102dd578063b7c9252c146103175761014c565b80631a788a02116101155780631a788a02146101eb5780631e3dd18b1461021157806335ab45281461022e5780633f3108f7146102365780633f5160181461023e57806353816b40146102645761014c565b8062dab9461461015157806301fb0a4f1461016b57806302d05d3f146101735780630ed01e341461019757806316bb6c13146101c5575b600080fd5b6101596103c1565b60408051918252519081900360200190f35b6101596103c7565b61017b6103cd565b604080516001600160a01b039092168252519081900360200190f35b6101c3600480360360408110156101ad57600080fd5b506001600160a01b0381351690602001356103dc565b005b6101c3600480360360208110156101db57600080fd5b50356001600160a01b0316610469565b61017b6004803603602081101561020157600080fd5b50356001600160a01b031661056d565b61017b6004803603602081101561022757600080fd5b5035610588565b61017b6105af565b61017b6105be565b6101c36004803603602081101561025457600080fd5b50356001600160a01b03166105cd565b61015961064e565b610159610654565b6101c36004803603602081101561028a57600080fd5b50356001600160a01b031661065a565b6101c3600480360360208110156102b057600080fd5b5035610753565b61017b600480360360208110156102cd57600080fd5b50356001600160a01b0316610774565b610303600480360360208110156102f357600080fd5b50356001600160a01b0316610bff565b604080519115158252519081900360200190f35b610159610c14565b61017b610ca2565b6101c36004803603604081101561033d57600080fd5b506001600160a01b038135169060200135610cb1565b610159610d26565b6101c36004803603602081101561037157600080fd5b5035610d2c565b6101c36004803603602081101561038e57600080fd5b5035610d5b565b6101c3600480360360408110156103ab57600080fd5b506001600160a01b038135169060200135610d84565b60085481565b60075481565b6005546001600160a01b031681565b3360009081526001602052604090205460ff166103f857600080fd5b60035460408051632770a7eb60e21b81526001600160a01b0385811660048301526024820185905291519190921691639dc29fac91604480830192600092919082900301818387803b15801561044d57600080fd5b505af1158015610461573d6000803e3d6000fd5b505050505050565b6006546001600160a01b0316331461048057600080fd5b6003546001600160a01b0316156104c85760405162461bcd60e51b8152600401808060200182810382526024815260200180610f216024913960400191505060405180910390fd5b600380546001600160a01b0319166001600160a01b0383811691909117918290556040805163313ce56760e01b81529051929091169163313ce56791600480820192602092909190829003018186803b15801561052457600080fd5b505afa158015610538573d6000803e3d6000fd5b505050506040513d602081101561054e57600080fd5b5051600b805460ff9092166101000261ff001990921691909117905550565b6000602081905290815260409020546001600160a01b031681565b6002818154811061059557fe5b6000918252602090912001546001600160a01b0316905081565b6004546001600160a01b031681565b6006546001600160a01b031681565b6006546001600160a01b031633146105e457600080fd5b6005546001600160a01b03161561062c5760405162461bcd60e51b8152600401808060200182810382526023815260200180610efe6023913960400191505060405180910390fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b60095481565b60025490565b6006546001600160a01b0316331461067157600080fd5b6004546001600160a01b0316156106b95760405162461bcd60e51b8152600401808060200182810382526023815260200180610efe6023913960400191505060405180910390fd5b600480546001600160a01b0319166001600160a01b03838116919091178083556040805163313ce56760e01b81529051919092169263313ce56792808201926020929091829003018186803b15801561071157600080fd5b505afa158015610725573d6000803e3d6000fd5b505050506040513d602081101561073b57600080fd5b5051600b805460ff191660ff90921691909117905550565b6006546001600160a01b0316331461076a57600080fd5b600a819055600755565b6003546000906001600160a01b03161580159061079957506001600160a01b03821615155b6107ea576040805162461bcd60e51b815260206004820152601960248201527f64466564466163746f72793a205a45524f5f4144445245535300000000000000604482015290519081900360640190fd5b6001600160a01b038281166000908152602081905260409020541615610857576040805162461bcd60e51b815260206004820152601860248201527f64466564466163746f72793a20504149525f4558495354530000000000000000604482015290519081900360640190fd5b600554600354604080516364e329cb60e11b81526001600160a01b03868116600483015292831660248201529051919092169163c9c653969160448083019260209291908290030181600087803b1580156108b157600080fd5b505af11580156108c5573d6000803e3d6000fd5b505050506040513d60208110156108db57600080fd5b50516001600160a01b038082166000818152600160208190526040808320805460ff1916909217909155600354815163485cc95560e01b81529085166004820152938716602485015251939450909263485cc95592604480820193929182900301818387803b15801561094d57600080fd5b505af1158015610961573d6000803e3d6000fd5b5050506001600160a01b0380841660008181526020819052604080822080549487166001600160a01b03199586168117909155600280546001810182559084527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180549095161790935582516395d89b4160e01b815292519193507ffa2893d46bf879832624d3a015fb2e4e4573d12a25841988e9e17c5c2ab8b1a39284926395d89b4192600480840193919291829003018186803b158015610a2457600080fd5b505afa158015610a38573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610a6157600080fd5b8101908080516040519392919084640100000000821115610a8157600080fd5b908301906020820185811115610a9657600080fd5b8251640100000000811182820188101715610ab057600080fd5b82525081516020918201929091019080838360005b83811015610add578181015183820152602001610ac5565b50505050905090810190601f168015610b0a5780820380516001836020036101000a031916815260200191505b50604052505050846001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015610b4a57600080fd5b505afa158015610b5e573d6000803e3d6000fd5b505050506040513d6020811015610b7457600080fd5b50516040805160ff8316602082810191909152828252845192820192909252835190918291606083019186019080838360005b83811015610bbf578181015183820152602001610ba7565b50505050905090810190601f168015610bec5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2919050565b60016020526000908152604090205460ff1681565b3360009081526001602052604081205460ff16610c3057600080fd5b600754431115610c9b5760095415610c9657600954600b54600754610c929291610c809160ff61010083048116600a90810a93610c7a939216900a90829060649082904390610df5565b90610e4b565b81610c8757fe5b600854919004610eae565b6008555b436007555b5060085490565b6003546001600160a01b031681565b3360009081526001602052604090205460ff16610ccd57600080fd5b60048054604080516340c10f1960e01b81526001600160a01b038681169482019490945260248101859052905192909116916340c10f199160448082019260009290919082900301818387803b15801561044d57600080fd5b600a5481565b3360009081526001602052604090205460ff16610d4857600080fd5b600954610d559082610df5565b60095550565b3360009081526001602052604090205460ff16610d7757600080fd5b600954610d559082610eae565b3360009081526001602052604090205460ff16610da057600080fd5b600354604080516340c10f1960e01b81526001600160a01b03858116600483015260248201859052915191909216916340c10f1991604480830192600092919082900301818387803b15801561044d57600080fd5b80820382811115610e45576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b92915050565b6000811580610e6657505080820282828281610e6357fe5b04145b610e45576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b80820182811015610e45576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fdfe64466564466163746f72793a204645445f544f4b454e5f4841535f4245454e5f53455464466564466163746f72793a20424153455f544f4b454e5f4841535f4245454e5f534554a26469706673582212208c11a156408630261061f33c77bc56cf394297da5d01ed1f943cc9232659cef164736f6c63430007000033"

// DeployDFedFactory deploys a new Ethereum contract, binding an instance of DFedFactory to it.
func DeployDFedFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DFedFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DFedFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DFedFactory{DFedFactoryCaller: DFedFactoryCaller{contract: contract}, DFedFactoryTransactor: DFedFactoryTransactor{contract: contract}, DFedFactoryFilterer: DFedFactoryFilterer{contract: contract}}, nil
}

// DFedFactory is an auto generated Go binding around an Ethereum contract.
type DFedFactory struct {
	DFedFactoryCaller     // Read-only binding to the contract
	DFedFactoryTransactor // Write-only binding to the contract
	DFedFactoryFilterer   // Log filterer for contract events
}

// DFedFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DFedFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DFedFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DFedFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DFedFactorySession struct {
	Contract     *DFedFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DFedFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DFedFactoryCallerSession struct {
	Contract *DFedFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DFedFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DFedFactoryTransactorSession struct {
	Contract     *DFedFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DFedFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DFedFactoryRaw struct {
	Contract *DFedFactory // Generic contract binding to access the raw methods on
}

// DFedFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DFedFactoryCallerRaw struct {
	Contract *DFedFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DFedFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DFedFactoryTransactorRaw struct {
	Contract *DFedFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDFedFactory creates a new instance of DFedFactory, bound to a specific deployed contract.
func NewDFedFactory(address common.Address, backend bind.ContractBackend) (*DFedFactory, error) {
	contract, err := bindDFedFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DFedFactory{DFedFactoryCaller: DFedFactoryCaller{contract: contract}, DFedFactoryTransactor: DFedFactoryTransactor{contract: contract}, DFedFactoryFilterer: DFedFactoryFilterer{contract: contract}}, nil
}

// NewDFedFactoryCaller creates a new read-only instance of DFedFactory, bound to a specific deployed contract.
func NewDFedFactoryCaller(address common.Address, caller bind.ContractCaller) (*DFedFactoryCaller, error) {
	contract, err := bindDFedFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DFedFactoryCaller{contract: contract}, nil
}

// NewDFedFactoryTransactor creates a new write-only instance of DFedFactory, bound to a specific deployed contract.
func NewDFedFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DFedFactoryTransactor, error) {
	contract, err := bindDFedFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DFedFactoryTransactor{contract: contract}, nil
}

// NewDFedFactoryFilterer creates a new log filterer instance of DFedFactory, bound to a specific deployed contract.
func NewDFedFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DFedFactoryFilterer, error) {
	contract, err := bindDFedFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DFedFactoryFilterer{contract: contract}, nil
}

// bindDFedFactory binds a generic wrapper to an already deployed contract.
func bindDFedFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedFactory *DFedFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedFactory.Contract.DFedFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedFactory *DFedFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedFactory.Contract.DFedFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedFactory *DFedFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedFactory.Contract.DFedFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedFactory *DFedFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedFactory *DFedFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedFactory *DFedFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedFactory.Contract.contract.Transact(opts, method, params...)
}

// FEDToken is a free data retrieval call binding the contract method 0x35ab4528.
//
// Solidity: function FEDToken() view returns(address)
func (_DFedFactory *DFedFactoryCaller) FEDToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "FEDToken")
	return *ret0, err
}

// FEDToken is a free data retrieval call binding the contract method 0x35ab4528.
//
// Solidity: function FEDToken() view returns(address)
func (_DFedFactory *DFedFactorySession) FEDToken() (common.Address, error) {
	return _DFedFactory.Contract.FEDToken(&_DFedFactory.CallOpts)
}

// FEDToken is a free data retrieval call binding the contract method 0x35ab4528.
//
// Solidity: function FEDToken() view returns(address)
func (_DFedFactory *DFedFactoryCallerSession) FEDToken() (common.Address, error) {
	return _DFedFactory.Contract.FEDToken(&_DFedFactory.CallOpts)
}

// AccRewardPerUSDDGlobal is a free data retrieval call binding the contract method 0x00dab946.
//
// Solidity: function accRewardPerUSDDGlobal() view returns(uint256)
func (_DFedFactory *DFedFactoryCaller) AccRewardPerUSDDGlobal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "accRewardPerUSDDGlobal")
	return *ret0, err
}

// AccRewardPerUSDDGlobal is a free data retrieval call binding the contract method 0x00dab946.
//
// Solidity: function accRewardPerUSDDGlobal() view returns(uint256)
func (_DFedFactory *DFedFactorySession) AccRewardPerUSDDGlobal() (*big.Int, error) {
	return _DFedFactory.Contract.AccRewardPerUSDDGlobal(&_DFedFactory.CallOpts)
}

// AccRewardPerUSDDGlobal is a free data retrieval call binding the contract method 0x00dab946.
//
// Solidity: function accRewardPerUSDDGlobal() view returns(uint256)
func (_DFedFactory *DFedFactoryCallerSession) AccRewardPerUSDDGlobal() (*big.Int, error) {
	return _DFedFactory.Contract.AccRewardPerUSDDGlobal(&_DFedFactory.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_DFedFactory *DFedFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "allPairs", arg0)
	return *ret0, err
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_DFedFactory *DFedFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _DFedFactory.Contract.AllPairs(&_DFedFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_DFedFactory *DFedFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _DFedFactory.Contract.AllPairs(&_DFedFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_DFedFactory *DFedFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "allPairsLength")
	return *ret0, err
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_DFedFactory *DFedFactorySession) AllPairsLength() (*big.Int, error) {
	return _DFedFactory.Contract.AllPairsLength(&_DFedFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_DFedFactory *DFedFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _DFedFactory.Contract.AllPairsLength(&_DFedFactory.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DFedFactory *DFedFactoryCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "baseToken")
	return *ret0, err
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DFedFactory *DFedFactorySession) BaseToken() (common.Address, error) {
	return _DFedFactory.Contract.BaseToken(&_DFedFactory.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DFedFactory *DFedFactoryCallerSession) BaseToken() (common.Address, error) {
	return _DFedFactory.Contract.BaseToken(&_DFedFactory.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_DFedFactory *DFedFactoryCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_DFedFactory *DFedFactorySession) Creator() (common.Address, error) {
	return _DFedFactory.Contract.Creator(&_DFedFactory.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_DFedFactory *DFedFactoryCallerSession) Creator() (common.Address, error) {
	return _DFedFactory.Contract.Creator(&_DFedFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address ) view returns(address)
func (_DFedFactory *DFedFactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "getPair", arg0)
	return *ret0, err
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address ) view returns(address)
func (_DFedFactory *DFedFactorySession) GetPair(arg0 common.Address) (common.Address, error) {
	return _DFedFactory.Contract.GetPair(&_DFedFactory.CallOpts, arg0)
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address ) view returns(address)
func (_DFedFactory *DFedFactoryCallerSession) GetPair(arg0 common.Address) (common.Address, error) {
	return _DFedFactory.Contract.GetPair(&_DFedFactory.CallOpts, arg0)
}

// LastRewardedBlockGlobal is a free data retrieval call binding the contract method 0x01fb0a4f.
//
// Solidity: function lastRewardedBlockGlobal() view returns(uint256)
func (_DFedFactory *DFedFactoryCaller) LastRewardedBlockGlobal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "lastRewardedBlockGlobal")
	return *ret0, err
}

// LastRewardedBlockGlobal is a free data retrieval call binding the contract method 0x01fb0a4f.
//
// Solidity: function lastRewardedBlockGlobal() view returns(uint256)
func (_DFedFactory *DFedFactorySession) LastRewardedBlockGlobal() (*big.Int, error) {
	return _DFedFactory.Contract.LastRewardedBlockGlobal(&_DFedFactory.CallOpts)
}

// LastRewardedBlockGlobal is a free data retrieval call binding the contract method 0x01fb0a4f.
//
// Solidity: function lastRewardedBlockGlobal() view returns(uint256)
func (_DFedFactory *DFedFactoryCallerSession) LastRewardedBlockGlobal() (*big.Int, error) {
	return _DFedFactory.Contract.LastRewardedBlockGlobal(&_DFedFactory.CallOpts)
}

// PairExist is a free data retrieval call binding the contract method 0xac393aff.
//
// Solidity: function pairExist(address ) view returns(bool)
func (_DFedFactory *DFedFactoryCaller) PairExist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "pairExist", arg0)
	return *ret0, err
}

// PairExist is a free data retrieval call binding the contract method 0xac393aff.
//
// Solidity: function pairExist(address ) view returns(bool)
func (_DFedFactory *DFedFactorySession) PairExist(arg0 common.Address) (bool, error) {
	return _DFedFactory.Contract.PairExist(&_DFedFactory.CallOpts, arg0)
}

// PairExist is a free data retrieval call binding the contract method 0xac393aff.
//
// Solidity: function pairExist(address ) view returns(bool)
func (_DFedFactory *DFedFactoryCallerSession) PairExist(arg0 common.Address) (bool, error) {
	return _DFedFactory.Contract.PairExist(&_DFedFactory.CallOpts, arg0)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(address)
func (_DFedFactory *DFedFactoryCaller) Setter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "setter")
	return *ret0, err
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(address)
func (_DFedFactory *DFedFactorySession) Setter() (common.Address, error) {
	return _DFedFactory.Contract.Setter(&_DFedFactory.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(address)
func (_DFedFactory *DFedFactoryCallerSession) Setter() (common.Address, error) {
	return _DFedFactory.Contract.Setter(&_DFedFactory.CallOpts)
}

// StartRewardBlock is a free data retrieval call binding the contract method 0xd7ed714e.
//
// Solidity: function startRewardBlock() view returns(uint256)
func (_DFedFactory *DFedFactoryCaller) StartRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "startRewardBlock")
	return *ret0, err
}

// StartRewardBlock is a free data retrieval call binding the contract method 0xd7ed714e.
//
// Solidity: function startRewardBlock() view returns(uint256)
func (_DFedFactory *DFedFactorySession) StartRewardBlock() (*big.Int, error) {
	return _DFedFactory.Contract.StartRewardBlock(&_DFedFactory.CallOpts)
}

// StartRewardBlock is a free data retrieval call binding the contract method 0xd7ed714e.
//
// Solidity: function startRewardBlock() view returns(uint256)
func (_DFedFactory *DFedFactoryCallerSession) StartRewardBlock() (*big.Int, error) {
	return _DFedFactory.Contract.StartRewardBlock(&_DFedFactory.CallOpts)
}

// TotalUSDDinLiquidityPoolGlobal is a free data retrieval call binding the contract method 0x53816b40.
//
// Solidity: function totalUSDDinLiquidityPoolGlobal() view returns(uint256)
func (_DFedFactory *DFedFactoryCaller) TotalUSDDinLiquidityPoolGlobal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedFactory.contract.Call(opts, out, "totalUSDDinLiquidityPoolGlobal")
	return *ret0, err
}

// TotalUSDDinLiquidityPoolGlobal is a free data retrieval call binding the contract method 0x53816b40.
//
// Solidity: function totalUSDDinLiquidityPoolGlobal() view returns(uint256)
func (_DFedFactory *DFedFactorySession) TotalUSDDinLiquidityPoolGlobal() (*big.Int, error) {
	return _DFedFactory.Contract.TotalUSDDinLiquidityPoolGlobal(&_DFedFactory.CallOpts)
}

// TotalUSDDinLiquidityPoolGlobal is a free data retrieval call binding the contract method 0x53816b40.
//
// Solidity: function totalUSDDinLiquidityPoolGlobal() view returns(uint256)
func (_DFedFactory *DFedFactoryCallerSession) TotalUSDDinLiquidityPoolGlobal() (*big.Int, error) {
	return _DFedFactory.Contract.TotalUSDDinLiquidityPoolGlobal(&_DFedFactory.CallOpts)
}

// AddTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdf6c1de1.
//
// Solidity: function addTotalUSDDinLiquidityPoolGlobal(uint256 _amount) returns()
func (_DFedFactory *DFedFactoryTransactor) AddTotalUSDDinLiquidityPoolGlobal(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "addTotalUSDDinLiquidityPoolGlobal", _amount)
}

// AddTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdf6c1de1.
//
// Solidity: function addTotalUSDDinLiquidityPoolGlobal(uint256 _amount) returns()
func (_DFedFactory *DFedFactorySession) AddTotalUSDDinLiquidityPoolGlobal(_amount *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.AddTotalUSDDinLiquidityPoolGlobal(&_DFedFactory.TransactOpts, _amount)
}

// AddTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdf6c1de1.
//
// Solidity: function addTotalUSDDinLiquidityPoolGlobal(uint256 _amount) returns()
func (_DFedFactory *DFedFactoryTransactorSession) AddTotalUSDDinLiquidityPoolGlobal(_amount *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.AddTotalUSDDinLiquidityPoolGlobal(&_DFedFactory.TransactOpts, _amount)
}

// BurnBaseToken is a paid mutator transaction binding the contract method 0x0ed01e34.
//
// Solidity: function burnBaseToken(address _from, uint256 _value) returns()
func (_DFedFactory *DFedFactoryTransactor) BurnBaseToken(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "burnBaseToken", _from, _value)
}

// BurnBaseToken is a paid mutator transaction binding the contract method 0x0ed01e34.
//
// Solidity: function burnBaseToken(address _from, uint256 _value) returns()
func (_DFedFactory *DFedFactorySession) BurnBaseToken(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.BurnBaseToken(&_DFedFactory.TransactOpts, _from, _value)
}

// BurnBaseToken is a paid mutator transaction binding the contract method 0x0ed01e34.
//
// Solidity: function burnBaseToken(address _from, uint256 _value) returns()
func (_DFedFactory *DFedFactoryTransactorSession) BurnBaseToken(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.BurnBaseToken(&_DFedFactory.TransactOpts, _from, _value)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9ccb0744.
//
// Solidity: function createPair(address _token) returns(address _pair)
func (_DFedFactory *DFedFactoryTransactor) CreatePair(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "createPair", _token)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9ccb0744.
//
// Solidity: function createPair(address _token) returns(address _pair)
func (_DFedFactory *DFedFactorySession) CreatePair(_token common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.CreatePair(&_DFedFactory.TransactOpts, _token)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9ccb0744.
//
// Solidity: function createPair(address _token) returns(address _pair)
func (_DFedFactory *DFedFactoryTransactorSession) CreatePair(_token common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.CreatePair(&_DFedFactory.TransactOpts, _token)
}

// MintBaseToken is a paid mutator transaction binding the contract method 0xf8d1006a.
//
// Solidity: function mintBaseToken(address _to, uint256 _value) returns()
func (_DFedFactory *DFedFactoryTransactor) MintBaseToken(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "mintBaseToken", _to, _value)
}

// MintBaseToken is a paid mutator transaction binding the contract method 0xf8d1006a.
//
// Solidity: function mintBaseToken(address _to, uint256 _value) returns()
func (_DFedFactory *DFedFactorySession) MintBaseToken(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.MintBaseToken(&_DFedFactory.TransactOpts, _to, _value)
}

// MintBaseToken is a paid mutator transaction binding the contract method 0xf8d1006a.
//
// Solidity: function mintBaseToken(address _to, uint256 _value) returns()
func (_DFedFactory *DFedFactoryTransactorSession) MintBaseToken(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.MintBaseToken(&_DFedFactory.TransactOpts, _to, _value)
}

// MintFEDToken is a paid mutator transaction binding the contract method 0xccd072e5.
//
// Solidity: function mintFEDToken(address _to, uint256 _value) returns()
func (_DFedFactory *DFedFactoryTransactor) MintFEDToken(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "mintFEDToken", _to, _value)
}

// MintFEDToken is a paid mutator transaction binding the contract method 0xccd072e5.
//
// Solidity: function mintFEDToken(address _to, uint256 _value) returns()
func (_DFedFactory *DFedFactorySession) MintFEDToken(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.MintFEDToken(&_DFedFactory.TransactOpts, _to, _value)
}

// MintFEDToken is a paid mutator transaction binding the contract method 0xccd072e5.
//
// Solidity: function mintFEDToken(address _to, uint256 _value) returns()
func (_DFedFactory *DFedFactoryTransactorSession) MintFEDToken(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.MintFEDToken(&_DFedFactory.TransactOpts, _to, _value)
}

// RemoveTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdaaa88bd.
//
// Solidity: function removeTotalUSDDinLiquidityPoolGlobal(uint256 _amount) returns()
func (_DFedFactory *DFedFactoryTransactor) RemoveTotalUSDDinLiquidityPoolGlobal(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "removeTotalUSDDinLiquidityPoolGlobal", _amount)
}

// RemoveTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdaaa88bd.
//
// Solidity: function removeTotalUSDDinLiquidityPoolGlobal(uint256 _amount) returns()
func (_DFedFactory *DFedFactorySession) RemoveTotalUSDDinLiquidityPoolGlobal(_amount *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.RemoveTotalUSDDinLiquidityPoolGlobal(&_DFedFactory.TransactOpts, _amount)
}

// RemoveTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdaaa88bd.
//
// Solidity: function removeTotalUSDDinLiquidityPoolGlobal(uint256 _amount) returns()
func (_DFedFactory *DFedFactoryTransactorSession) RemoveTotalUSDDinLiquidityPoolGlobal(_amount *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.RemoveTotalUSDDinLiquidityPoolGlobal(&_DFedFactory.TransactOpts, _amount)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address _baseToken) returns()
func (_DFedFactory *DFedFactoryTransactor) SetBaseToken(opts *bind.TransactOpts, _baseToken common.Address) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "setBaseToken", _baseToken)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address _baseToken) returns()
func (_DFedFactory *DFedFactorySession) SetBaseToken(_baseToken common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetBaseToken(&_DFedFactory.TransactOpts, _baseToken)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address _baseToken) returns()
func (_DFedFactory *DFedFactoryTransactorSession) SetBaseToken(_baseToken common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetBaseToken(&_DFedFactory.TransactOpts, _baseToken)
}

// SetCreator is a paid mutator transaction binding the contract method 0x3f516018.
//
// Solidity: function setCreator(address _creator) returns()
func (_DFedFactory *DFedFactoryTransactor) SetCreator(opts *bind.TransactOpts, _creator common.Address) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "setCreator", _creator)
}

// SetCreator is a paid mutator transaction binding the contract method 0x3f516018.
//
// Solidity: function setCreator(address _creator) returns()
func (_DFedFactory *DFedFactorySession) SetCreator(_creator common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetCreator(&_DFedFactory.TransactOpts, _creator)
}

// SetCreator is a paid mutator transaction binding the contract method 0x3f516018.
//
// Solidity: function setCreator(address _creator) returns()
func (_DFedFactory *DFedFactoryTransactorSession) SetCreator(_creator common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetCreator(&_DFedFactory.TransactOpts, _creator)
}

// SetFedToken is a paid mutator transaction binding the contract method 0x59d17cdc.
//
// Solidity: function setFedToken(address _FEDToken) returns()
func (_DFedFactory *DFedFactoryTransactor) SetFedToken(opts *bind.TransactOpts, _FEDToken common.Address) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "setFedToken", _FEDToken)
}

// SetFedToken is a paid mutator transaction binding the contract method 0x59d17cdc.
//
// Solidity: function setFedToken(address _FEDToken) returns()
func (_DFedFactory *DFedFactorySession) SetFedToken(_FEDToken common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetFedToken(&_DFedFactory.TransactOpts, _FEDToken)
}

// SetFedToken is a paid mutator transaction binding the contract method 0x59d17cdc.
//
// Solidity: function setFedToken(address _FEDToken) returns()
func (_DFedFactory *DFedFactoryTransactorSession) SetFedToken(_FEDToken common.Address) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetFedToken(&_DFedFactory.TransactOpts, _FEDToken)
}

// SetStartRewardBlock is a paid mutator transaction binding the contract method 0x62b9e098.
//
// Solidity: function setStartRewardBlock(uint256 _height) returns()
func (_DFedFactory *DFedFactoryTransactor) SetStartRewardBlock(opts *bind.TransactOpts, _height *big.Int) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "setStartRewardBlock", _height)
}

// SetStartRewardBlock is a paid mutator transaction binding the contract method 0x62b9e098.
//
// Solidity: function setStartRewardBlock(uint256 _height) returns()
func (_DFedFactory *DFedFactorySession) SetStartRewardBlock(_height *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetStartRewardBlock(&_DFedFactory.TransactOpts, _height)
}

// SetStartRewardBlock is a paid mutator transaction binding the contract method 0x62b9e098.
//
// Solidity: function setStartRewardBlock(uint256 _height) returns()
func (_DFedFactory *DFedFactoryTransactorSession) SetStartRewardBlock(_height *big.Int) (*types.Transaction, error) {
	return _DFedFactory.Contract.SetStartRewardBlock(&_DFedFactory.TransactOpts, _height)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xb7c9252c.
//
// Solidity: function updateInfo() returns(uint256)
func (_DFedFactory *DFedFactoryTransactor) UpdateInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedFactory.contract.Transact(opts, "updateInfo")
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xb7c9252c.
//
// Solidity: function updateInfo() returns(uint256)
func (_DFedFactory *DFedFactorySession) UpdateInfo() (*types.Transaction, error) {
	return _DFedFactory.Contract.UpdateInfo(&_DFedFactory.TransactOpts)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xb7c9252c.
//
// Solidity: function updateInfo() returns(uint256)
func (_DFedFactory *DFedFactoryTransactorSession) UpdateInfo() (*types.Transaction, error) {
	return _DFedFactory.Contract.UpdateInfo(&_DFedFactory.TransactOpts)
}

// DFedFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the DFedFactory contract.
type DFedFactoryPairCreatedIterator struct {
	Event *DFedFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *DFedFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedFactoryPairCreated)
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
		it.Event = new(DFedFactoryPairCreated)
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
func (it *DFedFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedFactoryPairCreated represents a PairCreated event raised by the DFedFactory contract.
type DFedFactoryPairCreated struct {
	Token    common.Address
	Symbol   string
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xfa2893d46bf879832624d3a015fb2e4e4573d12a25841988e9e17c5c2ab8b1a3.
//
// Solidity: event PairCreated(address indexed token, string symbol, uint8 decimals)
func (_DFedFactory *DFedFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token []common.Address) (*DFedFactoryPairCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DFedFactory.contract.FilterLogs(opts, "PairCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DFedFactoryPairCreatedIterator{contract: _DFedFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xfa2893d46bf879832624d3a015fb2e4e4573d12a25841988e9e17c5c2ab8b1a3.
//
// Solidity: event PairCreated(address indexed token, string symbol, uint8 decimals)
func (_DFedFactory *DFedFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *DFedFactoryPairCreated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DFedFactory.contract.WatchLogs(opts, "PairCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedFactoryPairCreated)
				if err := _DFedFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0xfa2893d46bf879832624d3a015fb2e4e4573d12a25841988e9e17c5c2ab8b1a3.
//
// Solidity: event PairCreated(address indexed token, string symbol, uint8 decimals)
func (_DFedFactory *DFedFactoryFilterer) ParsePairCreated(log types.Log) (*DFedFactoryPairCreated, error) {
	event := new(DFedFactoryPairCreated)
	if err := _DFedFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
