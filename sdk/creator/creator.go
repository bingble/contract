// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creator

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

// IdFedFactoryABI is the input ABI used to generate the binding from.
const IdFedFactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accRewardPerUSDDGlobal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addTotalUSDDinLiquidityPoolGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardedBlockGlobal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mintBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mintFEDToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeTotalUSDDinLiquidityPoolGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"setBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUSDDinLiquidityPoolGlobal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdFedFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IdFedFactoryFuncSigs = map[string]string{
	"00dab946": "accRewardPerUSDDGlobal()",
	"df6c1de1": "addTotalUSDDinLiquidityPoolGlobal(uint256)",
	"1e3dd18b": "allPairs(uint256)",
	"574f2ba3": "allPairsLength()",
	"c55dae63": "baseToken()",
	"0ed01e34": "burnBaseToken(address,uint256)",
	"9ccb0744": "createPair(address)",
	"1a788a02": "getPair(address)",
	"01fb0a4f": "lastRewardedBlockGlobal()",
	"f8d1006a": "mintBaseToken(address,uint256)",
	"ccd072e5": "mintFEDToken(address,uint256)",
	"daaa88bd": "removeTotalUSDDinLiquidityPoolGlobal(uint256)",
	"16bb6c13": "setBaseToken(address)",
	"d7ed714e": "startRewardBlock()",
	"53816b40": "totalUSDDinLiquidityPoolGlobal()",
	"b7c9252c": "updateInfo()",
}

// IdFedFactory is an auto generated Go binding around an Ethereum contract.
type IdFedFactory struct {
	IdFedFactoryCaller     // Read-only binding to the contract
	IdFedFactoryTransactor // Write-only binding to the contract
	IdFedFactoryFilterer   // Log filterer for contract events
}

// IdFedFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdFedFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdFedFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdFedFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdFedFactorySession struct {
	Contract     *IdFedFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdFedFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdFedFactoryCallerSession struct {
	Contract *IdFedFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IdFedFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdFedFactoryTransactorSession struct {
	Contract     *IdFedFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IdFedFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdFedFactoryRaw struct {
	Contract *IdFedFactory // Generic contract binding to access the raw methods on
}

// IdFedFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdFedFactoryCallerRaw struct {
	Contract *IdFedFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IdFedFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdFedFactoryTransactorRaw struct {
	Contract *IdFedFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdFedFactory creates a new instance of IdFedFactory, bound to a specific deployed contract.
func NewIdFedFactory(address common.Address, backend bind.ContractBackend) (*IdFedFactory, error) {
	contract, err := bindIdFedFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdFedFactory{IdFedFactoryCaller: IdFedFactoryCaller{contract: contract}, IdFedFactoryTransactor: IdFedFactoryTransactor{contract: contract}, IdFedFactoryFilterer: IdFedFactoryFilterer{contract: contract}}, nil
}

// NewIdFedFactoryCaller creates a new read-only instance of IdFedFactory, bound to a specific deployed contract.
func NewIdFedFactoryCaller(address common.Address, caller bind.ContractCaller) (*IdFedFactoryCaller, error) {
	contract, err := bindIdFedFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedFactoryCaller{contract: contract}, nil
}

// NewIdFedFactoryTransactor creates a new write-only instance of IdFedFactory, bound to a specific deployed contract.
func NewIdFedFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdFedFactoryTransactor, error) {
	contract, err := bindIdFedFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedFactoryTransactor{contract: contract}, nil
}

// NewIdFedFactoryFilterer creates a new log filterer instance of IdFedFactory, bound to a specific deployed contract.
func NewIdFedFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdFedFactoryFilterer, error) {
	contract, err := bindIdFedFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdFedFactoryFilterer{contract: contract}, nil
}

// bindIdFedFactory binds a generic wrapper to an already deployed contract.
func bindIdFedFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdFedFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedFactory *IdFedFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedFactory.Contract.IdFedFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedFactory *IdFedFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedFactory.Contract.IdFedFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedFactory *IdFedFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedFactory.Contract.IdFedFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedFactory *IdFedFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedFactory *IdFedFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedFactory *IdFedFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedFactory.Contract.contract.Transact(opts, method, params...)
}

// AccRewardPerUSDDGlobal is a free data retrieval call binding the contract method 0x00dab946.
//
// Solidity: function accRewardPerUSDDGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCaller) AccRewardPerUSDDGlobal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "accRewardPerUSDDGlobal")
	return *ret0, err
}

// AccRewardPerUSDDGlobal is a free data retrieval call binding the contract method 0x00dab946.
//
// Solidity: function accRewardPerUSDDGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactorySession) AccRewardPerUSDDGlobal() (*big.Int, error) {
	return _IdFedFactory.Contract.AccRewardPerUSDDGlobal(&_IdFedFactory.CallOpts)
}

// AccRewardPerUSDDGlobal is a free data retrieval call binding the contract method 0x00dab946.
//
// Solidity: function accRewardPerUSDDGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCallerSession) AccRewardPerUSDDGlobal() (*big.Int, error) {
	return _IdFedFactory.Contract.AccRewardPerUSDDGlobal(&_IdFedFactory.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_IdFedFactory *IdFedFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "allPairs", arg0)
	return *ret0, err
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_IdFedFactory *IdFedFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IdFedFactory.Contract.AllPairs(&_IdFedFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_IdFedFactory *IdFedFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IdFedFactory.Contract.AllPairs(&_IdFedFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "allPairsLength")
	return *ret0, err
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IdFedFactory *IdFedFactorySession) AllPairsLength() (*big.Int, error) {
	return _IdFedFactory.Contract.AllPairsLength(&_IdFedFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _IdFedFactory.Contract.AllPairsLength(&_IdFedFactory.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_IdFedFactory *IdFedFactoryCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "baseToken")
	return *ret0, err
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_IdFedFactory *IdFedFactorySession) BaseToken() (common.Address, error) {
	return _IdFedFactory.Contract.BaseToken(&_IdFedFactory.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_IdFedFactory *IdFedFactoryCallerSession) BaseToken() (common.Address, error) {
	return _IdFedFactory.Contract.BaseToken(&_IdFedFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address token) view returns(address)
func (_IdFedFactory *IdFedFactoryCaller) GetPair(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "getPair", token)
	return *ret0, err
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address token) view returns(address)
func (_IdFedFactory *IdFedFactorySession) GetPair(token common.Address) (common.Address, error) {
	return _IdFedFactory.Contract.GetPair(&_IdFedFactory.CallOpts, token)
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address token) view returns(address)
func (_IdFedFactory *IdFedFactoryCallerSession) GetPair(token common.Address) (common.Address, error) {
	return _IdFedFactory.Contract.GetPair(&_IdFedFactory.CallOpts, token)
}

// LastRewardedBlockGlobal is a free data retrieval call binding the contract method 0x01fb0a4f.
//
// Solidity: function lastRewardedBlockGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCaller) LastRewardedBlockGlobal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "lastRewardedBlockGlobal")
	return *ret0, err
}

// LastRewardedBlockGlobal is a free data retrieval call binding the contract method 0x01fb0a4f.
//
// Solidity: function lastRewardedBlockGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactorySession) LastRewardedBlockGlobal() (*big.Int, error) {
	return _IdFedFactory.Contract.LastRewardedBlockGlobal(&_IdFedFactory.CallOpts)
}

// LastRewardedBlockGlobal is a free data retrieval call binding the contract method 0x01fb0a4f.
//
// Solidity: function lastRewardedBlockGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCallerSession) LastRewardedBlockGlobal() (*big.Int, error) {
	return _IdFedFactory.Contract.LastRewardedBlockGlobal(&_IdFedFactory.CallOpts)
}

// StartRewardBlock is a free data retrieval call binding the contract method 0xd7ed714e.
//
// Solidity: function startRewardBlock() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCaller) StartRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "startRewardBlock")
	return *ret0, err
}

// StartRewardBlock is a free data retrieval call binding the contract method 0xd7ed714e.
//
// Solidity: function startRewardBlock() view returns(uint256)
func (_IdFedFactory *IdFedFactorySession) StartRewardBlock() (*big.Int, error) {
	return _IdFedFactory.Contract.StartRewardBlock(&_IdFedFactory.CallOpts)
}

// StartRewardBlock is a free data retrieval call binding the contract method 0xd7ed714e.
//
// Solidity: function startRewardBlock() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCallerSession) StartRewardBlock() (*big.Int, error) {
	return _IdFedFactory.Contract.StartRewardBlock(&_IdFedFactory.CallOpts)
}

// TotalUSDDinLiquidityPoolGlobal is a free data retrieval call binding the contract method 0x53816b40.
//
// Solidity: function totalUSDDinLiquidityPoolGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCaller) TotalUSDDinLiquidityPoolGlobal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedFactory.contract.Call(opts, out, "totalUSDDinLiquidityPoolGlobal")
	return *ret0, err
}

// TotalUSDDinLiquidityPoolGlobal is a free data retrieval call binding the contract method 0x53816b40.
//
// Solidity: function totalUSDDinLiquidityPoolGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactorySession) TotalUSDDinLiquidityPoolGlobal() (*big.Int, error) {
	return _IdFedFactory.Contract.TotalUSDDinLiquidityPoolGlobal(&_IdFedFactory.CallOpts)
}

// TotalUSDDinLiquidityPoolGlobal is a free data retrieval call binding the contract method 0x53816b40.
//
// Solidity: function totalUSDDinLiquidityPoolGlobal() view returns(uint256)
func (_IdFedFactory *IdFedFactoryCallerSession) TotalUSDDinLiquidityPoolGlobal() (*big.Int, error) {
	return _IdFedFactory.Contract.TotalUSDDinLiquidityPoolGlobal(&_IdFedFactory.CallOpts)
}

// AddTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdf6c1de1.
//
// Solidity: function addTotalUSDDinLiquidityPoolGlobal(uint256 amount) returns()
func (_IdFedFactory *IdFedFactoryTransactor) AddTotalUSDDinLiquidityPoolGlobal(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "addTotalUSDDinLiquidityPoolGlobal", amount)
}

// AddTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdf6c1de1.
//
// Solidity: function addTotalUSDDinLiquidityPoolGlobal(uint256 amount) returns()
func (_IdFedFactory *IdFedFactorySession) AddTotalUSDDinLiquidityPoolGlobal(amount *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.AddTotalUSDDinLiquidityPoolGlobal(&_IdFedFactory.TransactOpts, amount)
}

// AddTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdf6c1de1.
//
// Solidity: function addTotalUSDDinLiquidityPoolGlobal(uint256 amount) returns()
func (_IdFedFactory *IdFedFactoryTransactorSession) AddTotalUSDDinLiquidityPoolGlobal(amount *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.AddTotalUSDDinLiquidityPoolGlobal(&_IdFedFactory.TransactOpts, amount)
}

// BurnBaseToken is a paid mutator transaction binding the contract method 0x0ed01e34.
//
// Solidity: function burnBaseToken(address from, uint256 value) returns()
func (_IdFedFactory *IdFedFactoryTransactor) BurnBaseToken(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "burnBaseToken", from, value)
}

// BurnBaseToken is a paid mutator transaction binding the contract method 0x0ed01e34.
//
// Solidity: function burnBaseToken(address from, uint256 value) returns()
func (_IdFedFactory *IdFedFactorySession) BurnBaseToken(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.BurnBaseToken(&_IdFedFactory.TransactOpts, from, value)
}

// BurnBaseToken is a paid mutator transaction binding the contract method 0x0ed01e34.
//
// Solidity: function burnBaseToken(address from, uint256 value) returns()
func (_IdFedFactory *IdFedFactoryTransactorSession) BurnBaseToken(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.BurnBaseToken(&_IdFedFactory.TransactOpts, from, value)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9ccb0744.
//
// Solidity: function createPair(address tokenA) returns(address)
func (_IdFedFactory *IdFedFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "createPair", tokenA)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9ccb0744.
//
// Solidity: function createPair(address tokenA) returns(address)
func (_IdFedFactory *IdFedFactorySession) CreatePair(tokenA common.Address) (*types.Transaction, error) {
	return _IdFedFactory.Contract.CreatePair(&_IdFedFactory.TransactOpts, tokenA)
}

// CreatePair is a paid mutator transaction binding the contract method 0x9ccb0744.
//
// Solidity: function createPair(address tokenA) returns(address)
func (_IdFedFactory *IdFedFactoryTransactorSession) CreatePair(tokenA common.Address) (*types.Transaction, error) {
	return _IdFedFactory.Contract.CreatePair(&_IdFedFactory.TransactOpts, tokenA)
}

// MintBaseToken is a paid mutator transaction binding the contract method 0xf8d1006a.
//
// Solidity: function mintBaseToken(address from, uint256 value) returns()
func (_IdFedFactory *IdFedFactoryTransactor) MintBaseToken(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "mintBaseToken", from, value)
}

// MintBaseToken is a paid mutator transaction binding the contract method 0xf8d1006a.
//
// Solidity: function mintBaseToken(address from, uint256 value) returns()
func (_IdFedFactory *IdFedFactorySession) MintBaseToken(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.MintBaseToken(&_IdFedFactory.TransactOpts, from, value)
}

// MintBaseToken is a paid mutator transaction binding the contract method 0xf8d1006a.
//
// Solidity: function mintBaseToken(address from, uint256 value) returns()
func (_IdFedFactory *IdFedFactoryTransactorSession) MintBaseToken(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.MintBaseToken(&_IdFedFactory.TransactOpts, from, value)
}

// MintFEDToken is a paid mutator transaction binding the contract method 0xccd072e5.
//
// Solidity: function mintFEDToken(address _to, uint256 _value) returns()
func (_IdFedFactory *IdFedFactoryTransactor) MintFEDToken(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "mintFEDToken", _to, _value)
}

// MintFEDToken is a paid mutator transaction binding the contract method 0xccd072e5.
//
// Solidity: function mintFEDToken(address _to, uint256 _value) returns()
func (_IdFedFactory *IdFedFactorySession) MintFEDToken(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.MintFEDToken(&_IdFedFactory.TransactOpts, _to, _value)
}

// MintFEDToken is a paid mutator transaction binding the contract method 0xccd072e5.
//
// Solidity: function mintFEDToken(address _to, uint256 _value) returns()
func (_IdFedFactory *IdFedFactoryTransactorSession) MintFEDToken(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.MintFEDToken(&_IdFedFactory.TransactOpts, _to, _value)
}

// RemoveTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdaaa88bd.
//
// Solidity: function removeTotalUSDDinLiquidityPoolGlobal(uint256 amount) returns()
func (_IdFedFactory *IdFedFactoryTransactor) RemoveTotalUSDDinLiquidityPoolGlobal(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "removeTotalUSDDinLiquidityPoolGlobal", amount)
}

// RemoveTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdaaa88bd.
//
// Solidity: function removeTotalUSDDinLiquidityPoolGlobal(uint256 amount) returns()
func (_IdFedFactory *IdFedFactorySession) RemoveTotalUSDDinLiquidityPoolGlobal(amount *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.RemoveTotalUSDDinLiquidityPoolGlobal(&_IdFedFactory.TransactOpts, amount)
}

// RemoveTotalUSDDinLiquidityPoolGlobal is a paid mutator transaction binding the contract method 0xdaaa88bd.
//
// Solidity: function removeTotalUSDDinLiquidityPoolGlobal(uint256 amount) returns()
func (_IdFedFactory *IdFedFactoryTransactorSession) RemoveTotalUSDDinLiquidityPoolGlobal(amount *big.Int) (*types.Transaction, error) {
	return _IdFedFactory.Contract.RemoveTotalUSDDinLiquidityPoolGlobal(&_IdFedFactory.TransactOpts, amount)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address _baseToken) returns()
func (_IdFedFactory *IdFedFactoryTransactor) SetBaseToken(opts *bind.TransactOpts, _baseToken common.Address) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "setBaseToken", _baseToken)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address _baseToken) returns()
func (_IdFedFactory *IdFedFactorySession) SetBaseToken(_baseToken common.Address) (*types.Transaction, error) {
	return _IdFedFactory.Contract.SetBaseToken(&_IdFedFactory.TransactOpts, _baseToken)
}

// SetBaseToken is a paid mutator transaction binding the contract method 0x16bb6c13.
//
// Solidity: function setBaseToken(address _baseToken) returns()
func (_IdFedFactory *IdFedFactoryTransactorSession) SetBaseToken(_baseToken common.Address) (*types.Transaction, error) {
	return _IdFedFactory.Contract.SetBaseToken(&_IdFedFactory.TransactOpts, _baseToken)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xb7c9252c.
//
// Solidity: function updateInfo() returns(uint256)
func (_IdFedFactory *IdFedFactoryTransactor) UpdateInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedFactory.contract.Transact(opts, "updateInfo")
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xb7c9252c.
//
// Solidity: function updateInfo() returns(uint256)
func (_IdFedFactory *IdFedFactorySession) UpdateInfo() (*types.Transaction, error) {
	return _IdFedFactory.Contract.UpdateInfo(&_IdFedFactory.TransactOpts)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xb7c9252c.
//
// Solidity: function updateInfo() returns(uint256)
func (_IdFedFactory *IdFedFactoryTransactorSession) UpdateInfo() (*types.Transaction, error) {
	return _IdFedFactory.Contract.UpdateInfo(&_IdFedFactory.TransactOpts)
}

// IdFedFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the IdFedFactory contract.
type IdFedFactoryPairCreatedIterator struct {
	Event *IdFedFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *IdFedFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedFactoryPairCreated)
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
		it.Event = new(IdFedFactoryPairCreated)
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
func (it *IdFedFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedFactoryPairCreated represents a PairCreated event raised by the IdFedFactory contract.
type IdFedFactoryPairCreated struct {
	Token    common.Address
	Symbol   string
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xfa2893d46bf879832624d3a015fb2e4e4573d12a25841988e9e17c5c2ab8b1a3.
//
// Solidity: event PairCreated(address indexed token, string symbol, uint8 decimals)
func (_IdFedFactory *IdFedFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token []common.Address) (*IdFedFactoryPairCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IdFedFactory.contract.FilterLogs(opts, "PairCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IdFedFactoryPairCreatedIterator{contract: _IdFedFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xfa2893d46bf879832624d3a015fb2e4e4573d12a25841988e9e17c5c2ab8b1a3.
//
// Solidity: event PairCreated(address indexed token, string symbol, uint8 decimals)
func (_IdFedFactory *IdFedFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *IdFedFactoryPairCreated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IdFedFactory.contract.WatchLogs(opts, "PairCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedFactoryPairCreated)
				if err := _IdFedFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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
func (_IdFedFactory *IdFedFactoryFilterer) ParsePairCreated(log types.Log) (*IdFedFactoryPairCreated, error) {
	event := new(IdFedFactoryPairCreated)
	if err := _IdFedFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// StringsABI is the input ABI used to generate the binding from.
const StringsABI = "[]"

// StringsBin is the compiled bytecode used for deploying new contracts.
var StringsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122035fbd92cf8c3d5f721c073e37049b2d74994a00dee290e7763b4462aec31710164736f6c63430007000033"

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// TransferHelperABI is the input ABI used to generate the binding from.
const TransferHelperABI = "[]"

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
var TransferHelperBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220551ff6e8e532369e220dba76431298437616912a9c80eefd31784a59eb13bc3964736f6c63430007000033"

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}

// DFedBankABI is the input ABI used to generate the binding from.
const DFedBankABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToken0Amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToken1Amount\",\"type\":\"uint256\"}],\"name\":\"DebtUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetNum\",\"type\":\"uint256\"}],\"name\":\"Mortgage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"debtInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToken0Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToken1Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"name\":\"getDebtIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRepayByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetNum\",\"type\":\"uint256\"}],\"name\":\"isValidDebit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mortgage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"refundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRefundToken0Amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRefundToken1Amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRefundToken1EqualToken0Amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniqueDebtId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DFedBankFuncSigs maps the 4-byte function signature to its string representation.
var DFedBankFuncSigs = map[string]string{
	"3644e515": "DOMAIN_SEPARATOR()",
	"30adf81f": "PERMIT_TYPEHASH()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"e3c20fd5": "debtInfos(uint256)",
	"313ce567": "decimals()",
	"c45a0155": "factory()",
	"ddca3f43": "fee()",
	"3c33a245": "getDebtIndexById(uint256)",
	"ea52a0be": "getRepayByIndex(uint256)",
	"e991685f": "isValidDebit(uint256,uint256)",
	"58687e63": "mortgage(uint256,uint256,address)",
	"06fdde03": "name()",
	"7ecebe00": "nonces(address)",
	"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	"acde5d28": "refundAmount(address)",
	"371fd8e6": "repay(uint256)",
	"443cb4bc": "reserve0()",
	"5a76f25e": "reserve1()",
	"95d89b41": "symbol()",
	"0dfe1681": "token0()",
	"d21220a7": "token1()",
	"c21a43e4": "totalPledge()",
	"a058274f": "totalRefundToken0Amount()",
	"076195b1": "totalRefundToken1Amount()",
	"5b863ab9": "totalRefundToken1EqualToken0Amount()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"cfc6aa9e": "uniqueDebtId()",
}

// DFedBankBin is the compiled bytecode used for deploying new contracts.
var DFedBankBin = "0x60806040526008805460ff1916600190811790915560105534801561002357600080fd5b50611e33806100336000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806370a082311161010f578063cfc6aa9e116100a2578063ddca3f4311610071578063ddca3f4314610524578063e3c20fd51461052c578063e991685f14610594578063ea52a0be146105d0576101e5565b8063cfc6aa9e14610495578063d21220a71461049d578063d505accf146104a5578063dd62ed3e146104f6576101e5565b8063a9059cbb116100de578063a9059cbb14610433578063acde5d281461045f578063c21a43e414610485578063c45a01551461048d576101e5565b806370a08231146103d75780637ecebe00146103fd57806395d89b4114610423578063a058274f1461042b576101e5565b8063313ce56711610187578063443cb4bc11610156578063443cb4bc1461038d57806358687e63146103955780635a76f25e146103c75780635b863ab9146103cf576101e5565b8063313ce5671461032b5780633644e51514610349578063371fd8e6146103515780633c33a24514610370576101e5565b80630dfe1681116101c35780630dfe1681146102c157806318160ddd146102e557806323b872dd146102ed57806330adf81f14610323576101e5565b806306fdde03146101ea578063076195b114610267578063095ea7b314610281575b600080fd5b6101f26105ed565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022c578181015183820152602001610214565b50505050905090810190601f1680156102595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61026f61067b565b60408051918252519081900360200190f35b6102ad6004803603604081101561029757600080fd5b506001600160a01b038135169060200135610681565b604080519115158252519081900360200190f35b6102c9610698565b604080516001600160a01b039092168252519081900360200190f35b61026f6106a7565b6102ad6004803603606081101561030357600080fd5b506001600160a01b038135811691602081013590911690604001356106ad565b61026f610741565b610333610765565b6040805160ff9092168252519081900360200190f35b61026f61076a565b61036e6004803603602081101561036757600080fd5b5035610770565b005b61026f6004803603602081101561038657600080fd5b5035610b22565b61026f610bac565b61026f600480360360608110156103ab57600080fd5b50803590602081013590604001356001600160a01b0316610bb2565b61026f610e96565b61026f610e9c565b61026f600480360360208110156103ed57600080fd5b50356001600160a01b0316610ea2565b61026f6004803603602081101561041357600080fd5b50356001600160a01b0316610eb4565b6101f2610ec6565b61026f610f20565b6102ad6004803603604081101561044957600080fd5b506001600160a01b038135169060200135610f26565b61026f6004803603602081101561047557600080fd5b50356001600160a01b0316610f33565b61026f610f45565b6102c9610f4b565b61026f610f5a565b6102c9610f60565b61036e600480360360e08110156104bb57600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c00135610f6f565b61026f6004803603604081101561050c57600080fd5b506001600160a01b0381358116916020013516611171565b61026f61118e565b6105496004803603602081101561054257600080fd5b5035611194565b604080516001600160a01b0390991689526020890197909752878701959095526060870193909352608086019190915260a085015260c084015260e083015251908190036101000190f35b6105b7600480360360408110156105aa57600080fd5b50803590602001356111f1565b6040805192835290151560208301528051918290030190f35b61026f600480360360208110156105e657600080fd5b5035611351565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106735780601f1061064857610100808354040283529160200191610673565b820191906000526020600020905b81548152906001019060200180831161065657829003601f168201915b505050505081565b60125481565b600061068e3384846113b1565b5060015b92915050565b600a546001600160a01b031681565b60025481565b6001600160a01b03831660009081526004602090815260408083203384529091528120546000191461072c576001600160a01b03841660009081526004602090815260408083203384529091529020546107079083611413565b6001600160a01b03851660009081526004602090815260408083203384529091529020555b610737848484611463565b5060019392505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60055481565b60085460ff166107ba576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff1916905560006107cf82610b22565b90508061081f576040805162461bcd60e51b81526020600482015260196024820152781911995910985b9ace88111150925517d393d517d193d55391603a1b604482015290519081900360640190fd5b60006108706007838154811061083157fe5b9060005260206000209060080201600401546007848154811061085057fe5b90600052602060002090600802016003015461151190919063ffffffff16565b905060006108c36007848154811061088457fe5b906000526020600020906008020160050154600785815481106108a357fe5b90600052602060002090600802016002015461151190919063ffffffff16565b90506000600784815481106108d457fe5b600091825260209182902060089091020154601154600954600d54600a54604080516370a0823160e01b815230600482015290516001600160a01b0396871698508a976109829761097c9695879590949116926370a0823192602480840193919291829003018186803b15801561094a57600080fd5b505afa15801561095e573d6000803e3d6000fd5b505050506040513d602081101561097457600080fd5b505190611413565b90611413565b10156109bf5760405162461bcd60e51b8152600401808060200182810382526023815260200180611ddb6023913960400191505060405180910390fd5b6109f2600785815481106109cf57fe5b90600052602060002090600802016004015460115461151190919063ffffffff16565b6011556109fe84611560565b600c54604080516303b4078d60e21b81523060048201526024810186905290516001600160a01b0390921691630ed01e349160448082019260009290919082900301818387803b158015610a5157600080fd5b505af1158015610a65573d6000803e3d6000fd5b5050600b54610a8192506001600160a01b0316905082846116d0565b6040805186815260006020820181905281830181905260608201819052608082015290516001600160a01b038316917fa0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869919081900360a00190a26040805186815290517fa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe462309181900360200190a150506008805460ff19166001179055505050565b6000806007600081548110610b3357fe5b90600052602060002090600802016006015490505b8015610ba1578260078281548110610b5c57fe5b9060005260206000209060080201600101541415610b7b579050610ba7565b60078181548110610b8857fe5b9060005260206000209060080201600601549050610b48565b60009150505b919050565b600d5481565b60085460009060ff16610bff576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff19169055601254600f54600e54600b54604080516370a0823160e01b815230600482015290518995610c6d95909461097c949193859391926001600160a01b0316916370a0823191602480820192602092909190829003018186803b15801561094a57600080fd5b1015610caa5760405162461bcd60e51b8152600401808060200182810382526023815260200180611ddb6023913960400191505060405180910390fd5b60008311610cff576040805162461bcd60e51b815260206004820152601f60248201527f6446656442616e6b3a205441524745545f4e554d4245525f494e56414c494400604482015290519081900360640190fd5b600080610d0c86866111f1565b9150915080610d62576040805162461bcd60e51b815260206004820152601760248201527f6446656442616e6b3a20494e56414c49445f4445424954000000000000000000604482015290519081900360640190fd5b610d6e8287878761183a565b600c5460408051637c68803560e11b81526001600160a01b038881166004830152602482018a9052915193965091169163f8d1006a9160448082019260009290919082900301818387803b158015610dc557600080fd5b505af1158015610dd9573d6000803e3d6000fd5b505060408051868152602081018a9052808201899052600060608201819052608082015290516001600160a01b03881693507fa0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a3386992509081900360a00190a2604080516001600160a01b03861681526020810188905280820187905290517f17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a49181900360600190a150506008805460ff191660011790559392505050565b600e5481565b60135481565b60036020526000908152604090205481565b60066020526000908152604090205481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106735780601f1061064857610100808354040283529160200191610673565b60115481565b600061068e338484611463565b60146020526000908152604090205481565b600f5481565b600c546001600160a01b031681565b60105481565b600b546001600160a01b031681565b42841015610fb9576040805162461bcd60e51b815260206004820152601260248201527119119959115490cc8c0e881156141254915160721b604482015290519081900360640190fd5b6005546001600160a01b0380891660008181526006602090815260408083208054600180820190925582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98186015280840196909652958d166060860152608085018c905260a085019590955260c08085018b90528151808603909101815260e08501825280519083012061190160f01b6101008601526101028501969096526101228085019690965280518085039096018652610142840180825286519683019690962095839052610162840180825286905260ff89166101828501526101a284018890526101c28401879052519193926101e280820193601f1981019281900390910190855afa1580156110d4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381161580159061110a5750886001600160a01b0316816001600160a01b0316145b61115b576040805162461bcd60e51b815260206004820152601c60248201527f6446656445524332303a20494e56414c49445f5349474e415455524500000000604482015290519081900360640190fd5b6111668989896113b1565b505050505050505050565b600460209081526000928352604080842090915290825290205481565b60095481565b600781815481106111a157fe5b6000918252602090912060089091020180546001820154600283015460038401546004850154600586015460068701546007909701546001600160a01b0390961697509395929491939092909188565b60008060006112068585600d54600e54611b33565b9050600e5481101561121f57600080925092505061134a565b6000805b6007858154811061123057fe5b906000526020600020906008020160060154600014611342576007858154811061125657fe5b90600052602060002090600802016006015491506112b66007838154811061127a57fe5b9060005260206000209060080201600201546007848154811061129957fe5b906000526020600020906008020160030154600d54600e54611b33565b9050806112c38489611511565b116112d557506001925061134a915050565b611306600783815481106112e557fe5b9060005260206000209060080201600201548261151190919063ffffffff16565b83101561131c576000809450945050505061134a565b6007858154811061132957fe5b9060005260206000209060080201600601549450611223565b506001925050505b9250929050565b6000816113a1576040805162461bcd60e51b81526020600482015260196024820152781911995910985b9ace88111150925517d393d517d193d55391603a1b604482015290519081900360640190fd5b6106926007838154811061083157fe5b6001600160a01b03808416600081815260046020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b80820382811115610692576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6001600160a01b0383166000908152600360205260409020546114869082611413565b6001600160a01b0380851660009081526003602052604080822093909355908416815220546114b59082611511565b6001600160a01b0380841660008181526003602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b80820182811015610692576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fd5b6007818154811061156d57fe5b906000526020600020906008020160060154600780838154811061158d57fe5b906000526020600020906008020160070154815481106115a957fe5b906000526020600020906008020160060181905550600781815481106115cb57fe5b90600052602060002090600802016007015460078083815481106115eb57fe5b9060005260206000209060080201600601548154811061160757fe5b9060005260206000209060080201600701819055506116716007828154811061162c57fe5b90600052602060002090600802016005015461097c6007848154811061164e57fe5b906000526020600020906008020160020154600f5461141390919063ffffffff16565b600f55600780548290811061168257fe5b60009182526020822060089091020180546001600160a01b03191681556001810182905560028101829055600381018290556004810182905560058101829055600681018290556007015550565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b178152925182516000946060949389169392918291908083835b6020831061174d5780518252601f19909201916020918201910161172e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146117af576040519150601f19603f3d011682016040523d82523d6000602084013e6117b4565b606091505b50915091508180156117e25750805115806117e257508080602001905160208110156117df57600080fd5b50515b611833576040805162461bcd60e51b815260206004820152601f60248201527f5472616e7366657248656c7065723a205452414e534645525f4641494c454400604482015290519081900360640190fd5b5050505050565b600080611845611bb0565b60108054600181019091559250905061185c611d8c565b604051806101000160405280856001600160a01b031681526020018481526020018781526020018681526020016000815260200160008152602001600789815481106118a457fe5b90600052602060002090600802016006015481526020018881525090508160078260c00151815481106118d357fe5b90600052602060002090600802016007018190555081600788815481106118f657fe5b6000918252602090912060066008909202010155600754821415611a91576007805460018101825560009190915281517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688600890920291820180546001600160a01b0319166001600160a01b0390921691909117905560208201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68982015560408201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68a82015560608201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68b82015560808201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68c82015560a08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68d82015560c08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68e82015560e08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68f90910155611b18565b8060078381548110611a9f57fe5b600091825260209182902083516008929092020180546001600160a01b0319166001600160a01b0390921691909117815590820151600182015560408201516002820155606082015160038201556080820151600482015560a0820151600582015560c0820151600682015560e0909101516007909101555b600f54611b259087611511565b600f55509095945050505050565b600080611b408484611bf4565b9050858102600082888381611b5157fe5b0414611b6957611b62838989611c57565b9050611b76565b868281611b7257fe5b0490505b611ba46002890461097c611b9f6004611b8f8d80611bf4565b81611b9657fe5b86919004611511565b611d0e565b98975050505050505050565b60015b600754811015611bf15760078181548110611bca57fe5b90600052602060002090600802016001015460001415611be957611bf1565b600101611bb3565b90565b6000811580611c0f57505080820282828281611c0c57fe5b04145b610692576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b6000806000611c668686611d5f565b91509150838110611c7657600080fd5b60008480611c8057fe5b868809905082811115611c94576001820391505b918290039160008590038516808681611ca957fe5b049550808481611cb557fe5b049350808160000381611cc457fe5b046001019290920292909201600285810380870282030280870282030280870282030280870282030280870282030280870282030295860290039094029390930295945050505050565b60006003821115611d51575080600160028204015b81811015611d4b57809150600281828581611d3a57fe5b040181611d4357fe5b049050611d23565b50610ba7565b8115610ba757506001919050565b6000808060001984860990508385029250828103915082811015611d84576001820391505b509250929050565b60405180610100016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152509056fe6446656442616e6b3a20494e53554646494349454e545f494e5055545f414d4f554e54a264697066735822122060ee91b594f4f63b16a0a7752fbaddb60790db95c867e4bcdbf00bdccd0c550564736f6c63430007000033"

// DeployDFedBank deploys a new Ethereum contract, binding an instance of DFedBank to it.
func DeployDFedBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DFedBank, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedBankABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DFedBankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DFedBank{DFedBankCaller: DFedBankCaller{contract: contract}, DFedBankTransactor: DFedBankTransactor{contract: contract}, DFedBankFilterer: DFedBankFilterer{contract: contract}}, nil
}

// DFedBank is an auto generated Go binding around an Ethereum contract.
type DFedBank struct {
	DFedBankCaller     // Read-only binding to the contract
	DFedBankTransactor // Write-only binding to the contract
	DFedBankFilterer   // Log filterer for contract events
}

// DFedBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type DFedBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DFedBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DFedBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DFedBankSession struct {
	Contract     *DFedBank         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DFedBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DFedBankCallerSession struct {
	Contract *DFedBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DFedBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DFedBankTransactorSession struct {
	Contract     *DFedBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DFedBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type DFedBankRaw struct {
	Contract *DFedBank // Generic contract binding to access the raw methods on
}

// DFedBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DFedBankCallerRaw struct {
	Contract *DFedBankCaller // Generic read-only contract binding to access the raw methods on
}

// DFedBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DFedBankTransactorRaw struct {
	Contract *DFedBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDFedBank creates a new instance of DFedBank, bound to a specific deployed contract.
func NewDFedBank(address common.Address, backend bind.ContractBackend) (*DFedBank, error) {
	contract, err := bindDFedBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DFedBank{DFedBankCaller: DFedBankCaller{contract: contract}, DFedBankTransactor: DFedBankTransactor{contract: contract}, DFedBankFilterer: DFedBankFilterer{contract: contract}}, nil
}

// NewDFedBankCaller creates a new read-only instance of DFedBank, bound to a specific deployed contract.
func NewDFedBankCaller(address common.Address, caller bind.ContractCaller) (*DFedBankCaller, error) {
	contract, err := bindDFedBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DFedBankCaller{contract: contract}, nil
}

// NewDFedBankTransactor creates a new write-only instance of DFedBank, bound to a specific deployed contract.
func NewDFedBankTransactor(address common.Address, transactor bind.ContractTransactor) (*DFedBankTransactor, error) {
	contract, err := bindDFedBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DFedBankTransactor{contract: contract}, nil
}

// NewDFedBankFilterer creates a new log filterer instance of DFedBank, bound to a specific deployed contract.
func NewDFedBankFilterer(address common.Address, filterer bind.ContractFilterer) (*DFedBankFilterer, error) {
	contract, err := bindDFedBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DFedBankFilterer{contract: contract}, nil
}

// bindDFedBank binds a generic wrapper to an already deployed contract.
func bindDFedBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedBank *DFedBankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedBank.Contract.DFedBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedBank *DFedBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedBank.Contract.DFedBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedBank *DFedBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedBank.Contract.DFedBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedBank *DFedBankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedBank *DFedBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedBank *DFedBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedBank.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedBank *DFedBankCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedBank *DFedBankSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DFedBank.Contract.DOMAINSEPARATOR(&_DFedBank.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedBank *DFedBankCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DFedBank.Contract.DOMAINSEPARATOR(&_DFedBank.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedBank *DFedBankCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedBank *DFedBankSession) PERMITTYPEHASH() ([32]byte, error) {
	return _DFedBank.Contract.PERMITTYPEHASH(&_DFedBank.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedBank *DFedBankCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _DFedBank.Contract.PERMITTYPEHASH(&_DFedBank.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedBank *DFedBankCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedBank *DFedBankSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.Allowance(&_DFedBank.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedBank *DFedBankCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.Allowance(&_DFedBank.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedBank *DFedBankCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedBank *DFedBankSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.BalanceOf(&_DFedBank.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedBank *DFedBankCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.BalanceOf(&_DFedBank.CallOpts, arg0)
}

// DebtInfos is a free data retrieval call binding the contract method 0xe3c20fd5.
//
// Solidity: function debtInfos(uint256 ) view returns(address user, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount, uint256 nextIndex, uint256 lastIndex)
func (_DFedBank *DFedBankCaller) DebtInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User             common.Address
	DebtId           *big.Int
	PledgeAmount     *big.Int
	RepayAmount      *big.Int
	DebtToken0Amount *big.Int
	DebtToken1Amount *big.Int
	NextIndex        *big.Int
	LastIndex        *big.Int
}, error) {
	ret := new(struct {
		User             common.Address
		DebtId           *big.Int
		PledgeAmount     *big.Int
		RepayAmount      *big.Int
		DebtToken0Amount *big.Int
		DebtToken1Amount *big.Int
		NextIndex        *big.Int
		LastIndex        *big.Int
	})
	out := ret
	err := _DFedBank.contract.Call(opts, out, "debtInfos", arg0)
	return *ret, err
}

// DebtInfos is a free data retrieval call binding the contract method 0xe3c20fd5.
//
// Solidity: function debtInfos(uint256 ) view returns(address user, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount, uint256 nextIndex, uint256 lastIndex)
func (_DFedBank *DFedBankSession) DebtInfos(arg0 *big.Int) (struct {
	User             common.Address
	DebtId           *big.Int
	PledgeAmount     *big.Int
	RepayAmount      *big.Int
	DebtToken0Amount *big.Int
	DebtToken1Amount *big.Int
	NextIndex        *big.Int
	LastIndex        *big.Int
}, error) {
	return _DFedBank.Contract.DebtInfos(&_DFedBank.CallOpts, arg0)
}

// DebtInfos is a free data retrieval call binding the contract method 0xe3c20fd5.
//
// Solidity: function debtInfos(uint256 ) view returns(address user, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount, uint256 nextIndex, uint256 lastIndex)
func (_DFedBank *DFedBankCallerSession) DebtInfos(arg0 *big.Int) (struct {
	User             common.Address
	DebtId           *big.Int
	PledgeAmount     *big.Int
	RepayAmount      *big.Int
	DebtToken0Amount *big.Int
	DebtToken1Amount *big.Int
	NextIndex        *big.Int
	LastIndex        *big.Int
}, error) {
	return _DFedBank.Contract.DebtInfos(&_DFedBank.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedBank *DFedBankCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedBank *DFedBankSession) Decimals() (uint8, error) {
	return _DFedBank.Contract.Decimals(&_DFedBank.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedBank *DFedBankCallerSession) Decimals() (uint8, error) {
	return _DFedBank.Contract.Decimals(&_DFedBank.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedBank *DFedBankCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedBank *DFedBankSession) Factory() (common.Address, error) {
	return _DFedBank.Contract.Factory(&_DFedBank.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedBank *DFedBankCallerSession) Factory() (common.Address, error) {
	return _DFedBank.Contract.Factory(&_DFedBank.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DFedBank *DFedBankCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DFedBank *DFedBankSession) Fee() (*big.Int, error) {
	return _DFedBank.Contract.Fee(&_DFedBank.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) Fee() (*big.Int, error) {
	return _DFedBank.Contract.Fee(&_DFedBank.CallOpts)
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 _debtId) view returns(uint256 _index)
func (_DFedBank *DFedBankCaller) GetDebtIndexById(opts *bind.CallOpts, _debtId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "getDebtIndexById", _debtId)
	return *ret0, err
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 _debtId) view returns(uint256 _index)
func (_DFedBank *DFedBankSession) GetDebtIndexById(_debtId *big.Int) (*big.Int, error) {
	return _DFedBank.Contract.GetDebtIndexById(&_DFedBank.CallOpts, _debtId)
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 _debtId) view returns(uint256 _index)
func (_DFedBank *DFedBankCallerSession) GetDebtIndexById(_debtId *big.Int) (*big.Int, error) {
	return _DFedBank.Contract.GetDebtIndexById(&_DFedBank.CallOpts, _debtId)
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 _index) view returns(uint256 _repayAmount)
func (_DFedBank *DFedBankCaller) GetRepayByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "getRepayByIndex", _index)
	return *ret0, err
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 _index) view returns(uint256 _repayAmount)
func (_DFedBank *DFedBankSession) GetRepayByIndex(_index *big.Int) (*big.Int, error) {
	return _DFedBank.Contract.GetRepayByIndex(&_DFedBank.CallOpts, _index)
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 _index) view returns(uint256 _repayAmount)
func (_DFedBank *DFedBankCallerSession) GetRepayByIndex(_index *big.Int) (*big.Int, error) {
	return _DFedBank.Contract.GetRepayByIndex(&_DFedBank.CallOpts, _index)
}

// IsValidDebit is a free data retrieval call binding the contract method 0xe991685f.
//
// Solidity: function isValidDebit(uint256 _pledgeAmount, uint256 _targetNum) view returns(uint256 _index, bool _valid)
func (_DFedBank *DFedBankCaller) IsValidDebit(opts *bind.CallOpts, _pledgeAmount *big.Int, _targetNum *big.Int) (struct {
	Index *big.Int
	Valid bool
}, error) {
	ret := new(struct {
		Index *big.Int
		Valid bool
	})
	out := ret
	err := _DFedBank.contract.Call(opts, out, "isValidDebit", _pledgeAmount, _targetNum)
	return *ret, err
}

// IsValidDebit is a free data retrieval call binding the contract method 0xe991685f.
//
// Solidity: function isValidDebit(uint256 _pledgeAmount, uint256 _targetNum) view returns(uint256 _index, bool _valid)
func (_DFedBank *DFedBankSession) IsValidDebit(_pledgeAmount *big.Int, _targetNum *big.Int) (struct {
	Index *big.Int
	Valid bool
}, error) {
	return _DFedBank.Contract.IsValidDebit(&_DFedBank.CallOpts, _pledgeAmount, _targetNum)
}

// IsValidDebit is a free data retrieval call binding the contract method 0xe991685f.
//
// Solidity: function isValidDebit(uint256 _pledgeAmount, uint256 _targetNum) view returns(uint256 _index, bool _valid)
func (_DFedBank *DFedBankCallerSession) IsValidDebit(_pledgeAmount *big.Int, _targetNum *big.Int) (struct {
	Index *big.Int
	Valid bool
}, error) {
	return _DFedBank.Contract.IsValidDebit(&_DFedBank.CallOpts, _pledgeAmount, _targetNum)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedBank *DFedBankCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedBank *DFedBankSession) Name() (string, error) {
	return _DFedBank.Contract.Name(&_DFedBank.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedBank *DFedBankCallerSession) Name() (string, error) {
	return _DFedBank.Contract.Name(&_DFedBank.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedBank *DFedBankCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedBank *DFedBankSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.Nonces(&_DFedBank.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedBank *DFedBankCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.Nonces(&_DFedBank.CallOpts, arg0)
}

// RefundAmount is a free data retrieval call binding the contract method 0xacde5d28.
//
// Solidity: function refundAmount(address ) view returns(uint256)
func (_DFedBank *DFedBankCaller) RefundAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "refundAmount", arg0)
	return *ret0, err
}

// RefundAmount is a free data retrieval call binding the contract method 0xacde5d28.
//
// Solidity: function refundAmount(address ) view returns(uint256)
func (_DFedBank *DFedBankSession) RefundAmount(arg0 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.RefundAmount(&_DFedBank.CallOpts, arg0)
}

// RefundAmount is a free data retrieval call binding the contract method 0xacde5d28.
//
// Solidity: function refundAmount(address ) view returns(uint256)
func (_DFedBank *DFedBankCallerSession) RefundAmount(arg0 common.Address) (*big.Int, error) {
	return _DFedBank.Contract.RefundAmount(&_DFedBank.CallOpts, arg0)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_DFedBank *DFedBankCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "reserve0")
	return *ret0, err
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_DFedBank *DFedBankSession) Reserve0() (*big.Int, error) {
	return _DFedBank.Contract.Reserve0(&_DFedBank.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) Reserve0() (*big.Int, error) {
	return _DFedBank.Contract.Reserve0(&_DFedBank.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_DFedBank *DFedBankCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "reserve1")
	return *ret0, err
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_DFedBank *DFedBankSession) Reserve1() (*big.Int, error) {
	return _DFedBank.Contract.Reserve1(&_DFedBank.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) Reserve1() (*big.Int, error) {
	return _DFedBank.Contract.Reserve1(&_DFedBank.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedBank *DFedBankCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedBank *DFedBankSession) Symbol() (string, error) {
	return _DFedBank.Contract.Symbol(&_DFedBank.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedBank *DFedBankCallerSession) Symbol() (string, error) {
	return _DFedBank.Contract.Symbol(&_DFedBank.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_DFedBank *DFedBankCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "token0")
	return *ret0, err
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_DFedBank *DFedBankSession) Token0() (common.Address, error) {
	return _DFedBank.Contract.Token0(&_DFedBank.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_DFedBank *DFedBankCallerSession) Token0() (common.Address, error) {
	return _DFedBank.Contract.Token0(&_DFedBank.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_DFedBank *DFedBankCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "token1")
	return *ret0, err
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_DFedBank *DFedBankSession) Token1() (common.Address, error) {
	return _DFedBank.Contract.Token1(&_DFedBank.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_DFedBank *DFedBankCallerSession) Token1() (common.Address, error) {
	return _DFedBank.Contract.Token1(&_DFedBank.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_DFedBank *DFedBankCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "totalPledge")
	return *ret0, err
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_DFedBank *DFedBankSession) TotalPledge() (*big.Int, error) {
	return _DFedBank.Contract.TotalPledge(&_DFedBank.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) TotalPledge() (*big.Int, error) {
	return _DFedBank.Contract.TotalPledge(&_DFedBank.CallOpts)
}

// TotalRefundToken0Amount is a free data retrieval call binding the contract method 0xa058274f.
//
// Solidity: function totalRefundToken0Amount() view returns(uint256)
func (_DFedBank *DFedBankCaller) TotalRefundToken0Amount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "totalRefundToken0Amount")
	return *ret0, err
}

// TotalRefundToken0Amount is a free data retrieval call binding the contract method 0xa058274f.
//
// Solidity: function totalRefundToken0Amount() view returns(uint256)
func (_DFedBank *DFedBankSession) TotalRefundToken0Amount() (*big.Int, error) {
	return _DFedBank.Contract.TotalRefundToken0Amount(&_DFedBank.CallOpts)
}

// TotalRefundToken0Amount is a free data retrieval call binding the contract method 0xa058274f.
//
// Solidity: function totalRefundToken0Amount() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) TotalRefundToken0Amount() (*big.Int, error) {
	return _DFedBank.Contract.TotalRefundToken0Amount(&_DFedBank.CallOpts)
}

// TotalRefundToken1Amount is a free data retrieval call binding the contract method 0x076195b1.
//
// Solidity: function totalRefundToken1Amount() view returns(uint256)
func (_DFedBank *DFedBankCaller) TotalRefundToken1Amount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "totalRefundToken1Amount")
	return *ret0, err
}

// TotalRefundToken1Amount is a free data retrieval call binding the contract method 0x076195b1.
//
// Solidity: function totalRefundToken1Amount() view returns(uint256)
func (_DFedBank *DFedBankSession) TotalRefundToken1Amount() (*big.Int, error) {
	return _DFedBank.Contract.TotalRefundToken1Amount(&_DFedBank.CallOpts)
}

// TotalRefundToken1Amount is a free data retrieval call binding the contract method 0x076195b1.
//
// Solidity: function totalRefundToken1Amount() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) TotalRefundToken1Amount() (*big.Int, error) {
	return _DFedBank.Contract.TotalRefundToken1Amount(&_DFedBank.CallOpts)
}

// TotalRefundToken1EqualToken0Amount is a free data retrieval call binding the contract method 0x5b863ab9.
//
// Solidity: function totalRefundToken1EqualToken0Amount() view returns(uint256)
func (_DFedBank *DFedBankCaller) TotalRefundToken1EqualToken0Amount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "totalRefundToken1EqualToken0Amount")
	return *ret0, err
}

// TotalRefundToken1EqualToken0Amount is a free data retrieval call binding the contract method 0x5b863ab9.
//
// Solidity: function totalRefundToken1EqualToken0Amount() view returns(uint256)
func (_DFedBank *DFedBankSession) TotalRefundToken1EqualToken0Amount() (*big.Int, error) {
	return _DFedBank.Contract.TotalRefundToken1EqualToken0Amount(&_DFedBank.CallOpts)
}

// TotalRefundToken1EqualToken0Amount is a free data retrieval call binding the contract method 0x5b863ab9.
//
// Solidity: function totalRefundToken1EqualToken0Amount() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) TotalRefundToken1EqualToken0Amount() (*big.Int, error) {
	return _DFedBank.Contract.TotalRefundToken1EqualToken0Amount(&_DFedBank.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedBank *DFedBankCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedBank *DFedBankSession) TotalSupply() (*big.Int, error) {
	return _DFedBank.Contract.TotalSupply(&_DFedBank.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) TotalSupply() (*big.Int, error) {
	return _DFedBank.Contract.TotalSupply(&_DFedBank.CallOpts)
}

// UniqueDebtId is a free data retrieval call binding the contract method 0xcfc6aa9e.
//
// Solidity: function uniqueDebtId() view returns(uint256)
func (_DFedBank *DFedBankCaller) UniqueDebtId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedBank.contract.Call(opts, out, "uniqueDebtId")
	return *ret0, err
}

// UniqueDebtId is a free data retrieval call binding the contract method 0xcfc6aa9e.
//
// Solidity: function uniqueDebtId() view returns(uint256)
func (_DFedBank *DFedBankSession) UniqueDebtId() (*big.Int, error) {
	return _DFedBank.Contract.UniqueDebtId(&_DFedBank.CallOpts)
}

// UniqueDebtId is a free data retrieval call binding the contract method 0xcfc6aa9e.
//
// Solidity: function uniqueDebtId() view returns(uint256)
func (_DFedBank *DFedBankCallerSession) UniqueDebtId() (*big.Int, error) {
	return _DFedBank.Contract.UniqueDebtId(&_DFedBank.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedBank *DFedBankTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedBank *DFedBankSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.Approve(&_DFedBank.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedBank *DFedBankTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.Approve(&_DFedBank.TransactOpts, _spender, _value)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 _pledgeAmount, uint256 _targetNum, address _to) returns(uint256 _debtId)
func (_DFedBank *DFedBankTransactor) Mortgage(opts *bind.TransactOpts, _pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedBank.contract.Transact(opts, "mortgage", _pledgeAmount, _targetNum, _to)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 _pledgeAmount, uint256 _targetNum, address _to) returns(uint256 _debtId)
func (_DFedBank *DFedBankSession) Mortgage(_pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedBank.Contract.Mortgage(&_DFedBank.TransactOpts, _pledgeAmount, _targetNum, _to)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 _pledgeAmount, uint256 _targetNum, address _to) returns(uint256 _debtId)
func (_DFedBank *DFedBankTransactorSession) Mortgage(_pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedBank.Contract.Mortgage(&_DFedBank.TransactOpts, _pledgeAmount, _targetNum, _to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedBank *DFedBankTransactor) Permit(opts *bind.TransactOpts, _from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedBank.contract.Transact(opts, "permit", _from, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedBank *DFedBankSession) Permit(_from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedBank.Contract.Permit(&_DFedBank.TransactOpts, _from, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedBank *DFedBankTransactorSession) Permit(_from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedBank.Contract.Permit(&_DFedBank.TransactOpts, _from, _spender, _value, _deadline, _v, _r, _s)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _debtId) returns()
func (_DFedBank *DFedBankTransactor) Repay(opts *bind.TransactOpts, _debtId *big.Int) (*types.Transaction, error) {
	return _DFedBank.contract.Transact(opts, "repay", _debtId)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _debtId) returns()
func (_DFedBank *DFedBankSession) Repay(_debtId *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.Repay(&_DFedBank.TransactOpts, _debtId)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _debtId) returns()
func (_DFedBank *DFedBankTransactorSession) Repay(_debtId *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.Repay(&_DFedBank.TransactOpts, _debtId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedBank *DFedBankTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedBank *DFedBankSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.Transfer(&_DFedBank.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedBank *DFedBankTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.Transfer(&_DFedBank.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedBank *DFedBankTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedBank *DFedBankSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.TransferFrom(&_DFedBank.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedBank *DFedBankTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedBank.Contract.TransferFrom(&_DFedBank.TransactOpts, _from, _to, _value)
}

// DFedBankApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DFedBank contract.
type DFedBankApprovalIterator struct {
	Event *DFedBankApproval // Event containing the contract specifics and raw log

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
func (it *DFedBankApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedBankApproval)
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
		it.Event = new(DFedBankApproval)
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
func (it *DFedBankApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedBankApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedBankApproval represents a Approval event raised by the DFedBank contract.
type DFedBankApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DFedBank *DFedBankFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DFedBankApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DFedBank.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DFedBankApprovalIterator{contract: _DFedBank.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DFedBank *DFedBankFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DFedBankApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DFedBank.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedBankApproval)
				if err := _DFedBank.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DFedBank *DFedBankFilterer) ParseApproval(log types.Log) (*DFedBankApproval, error) {
	event := new(DFedBankApproval)
	if err := _DFedBank.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedBankDebtUpdateIterator is returned from FilterDebtUpdate and is used to iterate over the raw logs and unpacked data for DebtUpdate events raised by the DFedBank contract.
type DFedBankDebtUpdateIterator struct {
	Event *DFedBankDebtUpdate // Event containing the contract specifics and raw log

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
func (it *DFedBankDebtUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedBankDebtUpdate)
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
		it.Event = new(DFedBankDebtUpdate)
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
func (it *DFedBankDebtUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedBankDebtUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedBankDebtUpdate represents a DebtUpdate event raised by the DFedBank contract.
type DFedBankDebtUpdate struct {
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
func (_DFedBank *DFedBankFilterer) FilterDebtUpdate(opts *bind.FilterOpts, owner []common.Address) (*DFedBankDebtUpdateIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DFedBank.contract.FilterLogs(opts, "DebtUpdate", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DFedBankDebtUpdateIterator{contract: _DFedBank.contract, event: "DebtUpdate", logs: logs, sub: sub}, nil
}

// WatchDebtUpdate is a free log subscription operation binding the contract event 0xa0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869.
//
// Solidity: event DebtUpdate(address indexed owner, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount)
func (_DFedBank *DFedBankFilterer) WatchDebtUpdate(opts *bind.WatchOpts, sink chan<- *DFedBankDebtUpdate, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DFedBank.contract.WatchLogs(opts, "DebtUpdate", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedBankDebtUpdate)
				if err := _DFedBank.contract.UnpackLog(event, "DebtUpdate", log); err != nil {
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
func (_DFedBank *DFedBankFilterer) ParseDebtUpdate(log types.Log) (*DFedBankDebtUpdate, error) {
	event := new(DFedBankDebtUpdate)
	if err := _DFedBank.contract.UnpackLog(event, "DebtUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedBankMortgageIterator is returned from FilterMortgage and is used to iterate over the raw logs and unpacked data for Mortgage events raised by the DFedBank contract.
type DFedBankMortgageIterator struct {
	Event *DFedBankMortgage // Event containing the contract specifics and raw log

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
func (it *DFedBankMortgageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedBankMortgage)
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
		it.Event = new(DFedBankMortgage)
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
func (it *DFedBankMortgageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedBankMortgageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedBankMortgage represents a Mortgage event raised by the DFedBank contract.
type DFedBankMortgage struct {
	Owner        common.Address
	PledgeAmount *big.Int
	TargetNum    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMortgage is a free log retrieval operation binding the contract event 0x17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a4.
//
// Solidity: event Mortgage(address owner, uint256 pledgeAmount, uint256 targetNum)
func (_DFedBank *DFedBankFilterer) FilterMortgage(opts *bind.FilterOpts) (*DFedBankMortgageIterator, error) {

	logs, sub, err := _DFedBank.contract.FilterLogs(opts, "Mortgage")
	if err != nil {
		return nil, err
	}
	return &DFedBankMortgageIterator{contract: _DFedBank.contract, event: "Mortgage", logs: logs, sub: sub}, nil
}

// WatchMortgage is a free log subscription operation binding the contract event 0x17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a4.
//
// Solidity: event Mortgage(address owner, uint256 pledgeAmount, uint256 targetNum)
func (_DFedBank *DFedBankFilterer) WatchMortgage(opts *bind.WatchOpts, sink chan<- *DFedBankMortgage) (event.Subscription, error) {

	logs, sub, err := _DFedBank.contract.WatchLogs(opts, "Mortgage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedBankMortgage)
				if err := _DFedBank.contract.UnpackLog(event, "Mortgage", log); err != nil {
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
func (_DFedBank *DFedBankFilterer) ParseMortgage(log types.Log) (*DFedBankMortgage, error) {
	event := new(DFedBankMortgage)
	if err := _DFedBank.contract.UnpackLog(event, "Mortgage", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedBankRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the DFedBank contract.
type DFedBankRepayIterator struct {
	Event *DFedBankRepay // Event containing the contract specifics and raw log

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
func (it *DFedBankRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedBankRepay)
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
		it.Event = new(DFedBankRepay)
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
func (it *DFedBankRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedBankRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedBankRepay represents a Repay event raised by the DFedBank contract.
type DFedBankRepay struct {
	DebtId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe46230.
//
// Solidity: event Repay(uint256 debtId)
func (_DFedBank *DFedBankFilterer) FilterRepay(opts *bind.FilterOpts) (*DFedBankRepayIterator, error) {

	logs, sub, err := _DFedBank.contract.FilterLogs(opts, "Repay")
	if err != nil {
		return nil, err
	}
	return &DFedBankRepayIterator{contract: _DFedBank.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe46230.
//
// Solidity: event Repay(uint256 debtId)
func (_DFedBank *DFedBankFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *DFedBankRepay) (event.Subscription, error) {

	logs, sub, err := _DFedBank.contract.WatchLogs(opts, "Repay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedBankRepay)
				if err := _DFedBank.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_DFedBank *DFedBankFilterer) ParseRepay(log types.Log) (*DFedBankRepay, error) {
	event := new(DFedBankRepay)
	if err := _DFedBank.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedBankTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DFedBank contract.
type DFedBankTransferIterator struct {
	Event *DFedBankTransfer // Event containing the contract specifics and raw log

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
func (it *DFedBankTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedBankTransfer)
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
		it.Event = new(DFedBankTransfer)
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
func (it *DFedBankTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedBankTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedBankTransfer represents a Transfer event raised by the DFedBank contract.
type DFedBankTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DFedBank *DFedBankFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DFedBankTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedBank.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DFedBankTransferIterator{contract: _DFedBank.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DFedBank *DFedBankFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DFedBankTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedBank.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedBankTransfer)
				if err := _DFedBank.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DFedBank *DFedBankFilterer) ParseTransfer(log types.Log) (*DFedBankTransfer, error) {
	event := new(DFedBankTransfer)
	if err := _DFedBank.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedCreatorABI is the input ABI used to generate the binding from.
const DFedCreatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DFedCreatorFuncSigs maps the 4-byte function signature to its string representation.
var DFedCreatorFuncSigs = map[string]string{
	"c9c65396": "createPair(address,address)",
	"c45a0155": "factory()",
	"3c9de1b8": "getInitHash()",
}

// DFedCreatorBin is the compiled bytecode used for deploying new contracts.
var DFedCreatorBin = "0x608060405234801561001057600080fd5b5060405161546d38038061546d8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055615408806100656000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633c9de1b814610046578063c45a015514610060578063c9c6539614610084575b600080fd5b61004e6100b2565b60408051918252519081900360200190f35b6100686100e8565b604080516001600160a01b039092168252519081900360200190f35b6100686004803603604081101561009a57600080fd5b506001600160a01b03813581169160200135166100f7565b60006060604051806020016100c690610195565b601f1982820381018352601f9091011660405280516020919091012091505090565b6000546001600160a01b031681565b600080546001600160a01b0316331461010f57600080fd5b60606040518060200161012190610195565b6020820181038252601f19601f8201166040525090506000838560405160200180836001600160a01b031660601b8152601401826001600160a01b031660601b815260140192505050604051602081830303815290604052805190602001209050808251602084016000f595945050505050565b615230806101a38339019056fe608060405260088054600160ff19918216811790925560109190915560158054909116905534801561003057600080fd5b506040805161010081018252600080825260208201818152928201818152606083018281526080840183815260a0850184815260c0860185815260e08701868152600780546001810182559752965160089096027fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688810180546001600160a01b03989098166001600160a01b03199098169790971790965596517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68986015592517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68a85015590517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68b840155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68c830155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68d82015591517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68e830155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68f90910155615051806101df6000396000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c80637dea576e1161015c578063c21a43e4116100ce578063dd62ed3e11610087578063dd62ed3e14610733578063ddca3f4314610761578063e3c20fd514610769578063e60688cc146107d1578063e991685f146107fd578063ea52a0be146108395761027f565b8063c21a43e41461069c578063c45a0155146106a4578063cfc6aa9e146106ac578063d21220a7146106b4578063d505accf146106bc578063d6fedef81461070d5761027f565b8063a058274f11610120578063a058274f1461062a578063a1f0400114610632578063a48685fe1461063a578063a9059cbb14610642578063acde5d281461066e578063ba9a7a56146106945761027f565b80637dea576e1461059c5780637ecebe00146105a457806388036ac5146105ca57806389afcb44146105fc57806395d89b41146106225761027f565b8063313ce567116101f5578063485cc955116101b9578063485cc955146104da57806358687e63146105085780635a76f25e1461053a5780635b863ab91461054257806370a082311461054a5780637c3a2473146105705761027f565b8063313ce567146104725780633644e51514610490578063371fd8e6146104985780633c33a245146104b5578063443cb4bc146104d25761027f565b80630e5c011e116102475780630e5c011e146103a0578063156e29f6146103c857806318160ddd146103fa57806321466eb51461040257806323b872dd1461043457806330adf81f1461046a5761027f565b806306fdde0314610284578063076195b1146103015780630902f1ac1461031b578063095ea7b31461033c5780630dfe16811461037c575b600080fd5b61028c610856565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102c65781810151838201526020016102ae565b50505050905090810190601f1680156102f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103096108e4565b60408051918252519081900360200190f35b6103236108ea565b6040805192835260208301919091528051918290030190f35b6103686004803603604081101561035257600080fd5b506001600160a01b0381351690602001356108f4565b604080519115158252519081900360200190f35b61038461090b565b604080516001600160a01b039092168252519081900360200190f35b6103c6600480360360208110156103b657600080fd5b50356001600160a01b031661091a565b005b610309600480360360608110156103de57600080fd5b506001600160a01b038135169060208101359060400135610a0e565b610309610f1a565b6103c66004803603606081101561041857600080fd5b50803590602081013590604001356001600160a01b0316610f20565b6103686004803603606081101561044a57600080fd5b506001600160a01b03813581169160208101359091169060400135611210565b6103096112a5565b61047a6112c9565b6040805160ff9092168252519081900360200190f35b6103096112ce565b6103c6600480360360208110156104ae57600080fd5b50356112d4565b610309600480360360208110156104cb57600080fd5b503561163c565b6103096116c6565b6103c6600480360360408110156104f057600080fd5b506001600160a01b03813581169160200135166116cc565b6103096004803603606081101561051e57600080fd5b50803590602081013590604001356001600160a01b031661199d565b610309611c6f565b610309611c75565b6103096004803603602081101561056057600080fd5b50356001600160a01b0316611c7b565b6103c66004803603604081101561058657600080fd5b506001600160a01b038135169060200135611c8d565b610309611d31565b610309600480360360208110156105ba57600080fd5b50356001600160a01b0316611d37565b6103c6600480360360608110156105e057600080fd5b50803590602081013590604001356001600160a01b0316611d49565b6103236004803603602081101561061257600080fd5b50356001600160a01b03166120b1565b61028c61241a565b610309612474565b61030961247a565b61047a612480565b6103686004803603604081101561065857600080fd5b506001600160a01b038135169060200135612489565b6103096004803603602081101561068457600080fd5b50356001600160a01b0316612496565b6103096124a8565b6103096124ae565b6103846124b4565b6103096124c3565b6103846124c9565b6103c6600480360360e08110156106d257600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c001356124d8565b6103096004803603602081101561072357600080fd5b50356001600160a01b03166126da565b6103096004803603604081101561074957600080fd5b506001600160a01b03813581169160200135166126ec565b610309612709565b6107866004803603602081101561077f57600080fd5b503561270f565b604080516001600160a01b0390991689526020890197909752878701959095526060870193909352608086019190915260a085015260c084015260e083015251908190036101000190f35b6103c6600480360360408110156107e757600080fd5b506001600160a01b03813516906020013561276c565b6108206004803603604081101561081357600080fd5b50803590602001356127d9565b6040805192835290151560208301528051918290030190f35b6103096004803603602081101561084f57600080fd5b5035612939565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108dc5780601f106108b1576101008083540402835291602001916108dc565b820191906000526020600020905b8154815290600101906020018083116108bf57829003601f168201915b505050505081565b60125481565b600d54600e549091565b6000610901338484612999565b5060015b92915050565b600a546001600160a01b031681565b6109226129fb565b600c546001600160a01b03828116600090815260196020526040902054601654919092169163ccd072e5918491670de0b6b3a764000091610986916109679190612b67565b6001600160a01b03871660009081526003602052604090205490612bb7565b8161098d57fe5b046040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156109d457600080fd5b505af11580156109e8573d6000803e3d6000fd5b50506016546001600160a01b039093166000908152601960205260409020929092555050565b60085460009060ff16610a5b576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff19169055600954600d54600a54604080516370a0823160e01b81523060048201529051610af99493610af39390926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610ac157600080fd5b505afa158015610ad5573d6000803e3d6000fd5b505050506040513d6020811015610aeb57600080fd5b505190612b67565b90612b67565b831115610b4d576040805162461bcd60e51b815260206004820152601f60248201527f64466564506169723a20494e53554646494349454e545f414d4f554e545f3000604482015290519081900360640190fd5b610bc1601254610af3600f54610af3600e54600b60009054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610ac157600080fd5b821115610c15576040805162461bcd60e51b815260206004820152601f60248201527f64466564506169723a20494e53554646494349454e545f414d4f554e545f3100604482015290519081900360640190fd5b600254610c206129fb565b80610c5157610c3d6103e8610af3610c388787612bb7565b612c1a565b9150610c4c60006103e8612c6b565b610c90565b600d54610c8d90610c628684612bb7565b81610c6957fe5b04600e54610c808487612bb790919063ffffffff16565b81610c8757fe5b04612cf6565b91505b60008211610ccf5760405162461bcd60e51b8152600401808060200182810382526027815260200180614fcf6027913960400191505060405180910390fd5b6001600160a01b038516600090815260036020526040902054610d0d576016546001600160a01b038616600090815260196020526040902055610da1565b6001600160a01b038516600090815260036020526040902054610d309083612d0c565b610d7c610d4860165485612bb790919063ffffffff16565b6001600160a01b038816600090815260196020908152604080832054600390925290912054610d7691612bb7565b90612d0c565b81610d8357fe5b6001600160a01b038716600090815260196020526040902091900490555b610dab8583612c6b565b610dcc600d54610dc685600e54612d0c90919063ffffffff16565b90612bb7565b610de7600e54610dc687600d54612d0c90919063ffffffff16565b1015610e3a576040805162461bcd60e51b815260206004820152601860248201527f64466564506169723a20494e56414c49445f5550444154450000000000000000604482015290519081900360640190fd5b600c546040805163df6c1de160e01b81526004810187905290516001600160a01b039092169163df6c1de19160248082019260009290919082900301818387803b158015610e8757600080fd5b505af1158015610e9b573d6000803e3d6000fd5b5050600d54610ec29250610eb0915086612d0c565b600e54610ebd9086612d0c565b612d5b565b604080518581526020810185905281516001600160a01b038816927f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f928290030190a2506008805460ff191660011790559392505050565b60025481565b60085460ff16610f6a576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff1916905582610fc0576040805162461bcd60e51b81526020600482015260176024820152761911995914185a5c8e881253959053125117d253941555604a1b604482015290519081900360640190fd5b82611035601154610af3600954610af3600d54600a60009054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610ac157600080fd5b10156110725760405162461bcd60e51b8152600401808060200182810382526023815260200180614f856023913960400191505060405180910390fd5b6015546000906127109061108a90869060ff16612bb7565b8161109157fe5b04905060006110a08583612b67565b905060006110b382600d54600e54612d66565b9050848110156110f45760405162461bcd60e51b8152600401808060200182810382526024815260200180614efb6024913960400191505060405180910390fd5b6110fc6129fb565b801561111957600b54611119906001600160a01b03168583612e21565b82156111305760095461112c9084612d0c565b6009555b600d5461114e906111419084612d0c565b600e54610ebd9084612b67565b600c546040805163df6c1de160e01b81526004810185905290516001600160a01b039092169163df6c1de19160248082019260009290919082900301818387803b15801561119b57600080fd5b505af11580156111af573d6000803e3d6000fd5b5050604080518981526020810185905281516001600160a01b03891694503393507ff2fc1627c485a28da79658364df5619455857b1d9d2fcc5787b19b45555665f3929181900390910190a350506008805460ff1916600117905550505050565b6001600160a01b03831660009081526004602090815260408083203384529091528120546000191461128f576001600160a01b038416600090815260046020908152604080832033845290915290205461126a9083612b67565b6001600160a01b03851660009081526004602090815260408083203384529091529020555b61129a848484612f8b565b5060015b9392505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60055481565b60085460ff1661131e576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff1916905560006113338261163c565b905080611383576040805162461bcd60e51b81526020600482015260196024820152781911995910985b9ace88111150925517d393d517d193d55391603a1b604482015290519081900360640190fd5b60006113d46007838154811061139557fe5b906000526020600020906008020160040154600784815481106113b457fe5b906000526020600020906008020160030154612d0c90919063ffffffff16565b90506000611427600784815481106113e857fe5b9060005260206000209060080201600501546007858154811061140757fe5b906000526020600020906008020160020154612d0c90919063ffffffff16565b905060006007848154811061143857fe5b600091825260209182902060089091020154601154600954600d54600a54604080516370a0823160e01b815230600482015290516001600160a01b0396871698508a976114ae97610af39695879590949116926370a0823192602480840193919291829003018186803b158015610ac157600080fd5b10156114eb5760405162461bcd60e51b8152600401808060200182810382526023815260200180614f1f6023913960400191505060405180910390fd5b61151e600785815481106114fb57fe5b906000526020600020906008020160040154601154612d0c90919063ffffffff16565b60115561152a84613039565b600c54604080516303b4078d60e21b81523060048201526024810186905290516001600160a01b0390921691630ed01e349160448082019260009290919082900301818387803b15801561157d57600080fd5b505af1158015611591573d6000803e3d6000fd5b5050600b546115ad92506001600160a01b031690508284612e21565b6040805186815260006020820181905281830181905260608201819052608082015290516001600160a01b03831691600080516020614f42833981519152919081900360a00190a26040805186815290517fa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe462309181900360200190a150506008805460ff19166001179055505050565b600080600760008154811061164d57fe5b90600052602060002090600802016006015490505b80156116bb57826007828154811061167657fe5b90600052602060002090600802016001015414156116955790506116c1565b600781815481106116a257fe5b9060005260206000209060080201600601549050611662565b60009150505b919050565b600d5481565b600c546001600160a01b031615611720576040805162461bcd60e51b8152602060048201526013602482015272322332b22830b4b91d102327a92124a22222a760691b604482015290519081900360640190fd5b600c8054336001600160a01b031991821617909155600a805482166001600160a01b03858116918217909255600b8054909316918416919091179091556040805163313ce56760e01b8152905163313ce56791600480820192602092909190829003018186803b15801561179357600080fd5b505afa1580156117a7573d6000803e3d6000fd5b505050506040513d60208110156117bd57600080fd5b50516018805460ff191660ff9092169190911790556117db816131a9565b80516117ef91600091602090910190614d8a565b506000805461181291600191600260001961010083861615020190911604614e08565b50600c60009054906101000a90046001600160a01b03166001600160a01b031663b7c9252c6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561186357600080fd5b505af1158015611877573d6000803e3d6000fd5b505050506040513d602081101561188d57600080fd5b50516017556040516000805446927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f9291819083906002600019610100600184161502019091160480156119185780601f106118f6576101008083540402835291820191611918565b820191906000526020600020905b815481529060010190602001808311611904575b505060408051918290038220828201825260018352603160f81b602093840152815180840196909652858201527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606086015260808501959095523060a0808601919091528551808603909101815260c0909401909452505080519101206005555050565b60085460009060ff166119ea576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff19169055601254600f54600e54600b54604080516370a0823160e01b815230600482015290518995611a58959094610af3949193859391926001600160a01b0316916370a0823191602480820192602092909190829003018186803b158015610ac157600080fd5b1015611a955760405162461bcd60e51b8152600401808060200182810382526023815260200180614f1f6023913960400191505060405180910390fd5b60008311611aea576040805162461bcd60e51b815260206004820152601f60248201527f6446656442616e6b3a205441524745545f4e554d4245525f494e56414c494400604482015290519081900360640190fd5b600080611af786866127d9565b9150915080611b4d576040805162461bcd60e51b815260206004820152601760248201527f6446656442616e6b3a20494e56414c49445f4445424954000000000000000000604482015290519081900360640190fd5b611b5982878787613335565b600c5460408051637c68803560e11b81526001600160a01b038881166004830152602482018a9052915193965091169163f8d1006a9160448082019260009290919082900301818387803b158015611bb057600080fd5b505af1158015611bc4573d6000803e3d6000fd5b505060408051868152602081018a9052808201899052600060608201819052608082015290516001600160a01b0388169350600080516020614f4283398151915292509081900360a00190a2604080516001600160a01b03861681526020810188905280820187905290517f17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a49181900360600190a150506008805460ff191660011790559392505050565b600e5481565b60135481565b60036020526000908152604090205481565b6001600160a01b038216600090815260146020526040902054611cb09082612b67565b6001600160a01b038316600090815260146020526040812091909155601354601254611cdd908490612bb7565b81611ce457fe5b049050611cfc82601354612b6790919063ffffffff16565b601355601254611d0c9082612b67565b6012558015611d2c57600b54611d2c906001600160a01b03168483612e21565b505050565b60165481565b60066020526000908152604090205481565b60085460ff16611d93576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff1916905582611de9576040805162461bcd60e51b81526020600482015260176024820152761911995914185a5c8e881253959053125117d253941555604a1b604482015290519081900360640190fd5b82611e5e601254610af3600f54610af3600e54600b60009054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610ac157600080fd5b1015611e9b5760405162461bcd60e51b8152600401808060200182810382526023815260200180614f856023913960400191505060405180910390fd5b600d548210611ef1576040805162461bcd60e51b815260206004820181905260248201527f64466564506169723a20494e53554646494349454e545f4c4951554944495459604482015290519081900360640190fd5b600080611efd8561362e565b915091506000611f1283600e54600d54612d66565b90506000611f208284612b67565b60155490915060009061271090611f3b90849060ff16612bb7565b81611f4257fe5b0490506000611f518383612b67565b905087811015611f925760405162461bcd60e51b8152600401808060200182810382526024815260200180614efb6024913960400191505060405180910390fd5b611f9a6129fb565b8015611fb757600a54611fb7906001600160a01b03168883612e21565b8115611fce57600954611fca9083612d0c565b6009555b600d54611fec90611fdf9086612b67565b600e54610ebd9089612d0c565b600c546040805163daaa88bd60e01b81526004810187905290516001600160a01b039092169163daaa88bd9160248082019260009290919082900301818387803b15801561203957600080fd5b505af115801561204d573d6000803e3d6000fd5b5050604080518c81526020810185905281516001600160a01b038c1694503393507f1b38ce2719d9be70e0c768bd2a89d778eaef1cd41e3a4c02fa2317d415c3d348929181900390910190a350506008805460ff1916600117905550505050505050565b600854600090819060ff16612100576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff191690556121126129fb565b30600090815260036020526040902054600254600d548190612135908490612bb7565b8161213c57fe5b04935080612155600e5484612bb790919063ffffffff16565b8161215c57fe5b04925060008411801561216f5750600083115b6121aa5760405162461bcd60e51b8152600401808060200182810382526027815260200180614fa86027913960400191505060405180910390fd5b6121b430836139bf565b60008060006121e36121d188600d54612b6790919063ffffffff16565b600e546121de9089612b67565b613a50565b6001600160a01b038b16600090815260146020526040902054929550909350915061220e9082612d0c565b6001600160a01b038916600090815260146020526040902055600d54612245906122389089612b67565b600e54610ebd9089612b67565b61224f8684612d0c565b955061225b8783612b67565b600c546040805163daaa88bd60e01b81526004810184905290519299506001600160a01b039091169163daaa88bd9160248082019260009290919082900301818387803b1580156122ab57600080fd5b505af11580156122bf573d6000803e3d6000fd5b5050600c546001600160a01b038b811660009081526019602052604090205460165491909216935063ccd072e592508b91670de0b6b3a76400009161230f916123089190612b67565b8a90612bb7565b8161231657fe5b046040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561235d57600080fd5b505af1158015612371573d6000803e3d6000fd5b50505050600086111561239557600b54612395906001600160a01b03168988612e21565b86156123b257600a546123b2906001600160a01b03168989612e21565b604080518681526020810189905280820188905290516001600160a01b038a169133917fd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c8833639181900360600190a350506008805460ff1916600117905550929491935090915050565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108dc5780601f106108b1576101008083540402835291602001916108dc565b60115481565b60175481565b60155460ff1681565b6000610901338484612f8b565b60146020526000908152604090205481565b6103e881565b600f5481565b600c546001600160a01b031681565b60105481565b600b546001600160a01b031681565b42841015612522576040805162461bcd60e51b815260206004820152601260248201527119119959115490cc8c0e881156141254915160721b604482015290519081900360640190fd5b6005546001600160a01b0380891660008181526006602090815260408083208054600180820190925582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98186015280840196909652958d166060860152608085018c905260a085019590955260c08085018b90528151808603909101815260e08501825280519083012061190160f01b6101008601526101028501969096526101228085019690965280518085039096018652610142840180825286519683019690962095839052610162840180825286905260ff89166101828501526101a284018890526101c28401879052519193926101e280820193601f1981019281900390910190855afa15801561263d573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116158015906126735750886001600160a01b0316816001600160a01b0316145b6126c4576040805162461bcd60e51b815260206004820152601c60248201527f6446656445524332303a20494e56414c49445f5349474e415455524500000000604482015290519081900360640190fd5b6126cf898989612999565b505050505050505050565b60196020526000908152604090205481565b600460209081526000928352604080842090915290825290205481565b60095481565b6007818154811061271c57fe5b6000918252602090912060089091020180546001820154600283015460038401546004850154600586015460068701546007909701546001600160a01b0390961697509395929491939092909188565b6001600160a01b03821660009081526014602052604090205461278f9082612b67565b6001600160a01b0383166000908152601460205260409020556011546127b59082612b67565b60115580156127d557600a546127d5906001600160a01b03168383612e21565b5050565b60008060006127ee8585600d54600e54614716565b9050600e54811015612807576000809250925050612932565b6000805b6007858154811061281857fe5b90600052602060002090600802016006015460001461292a576007858154811061283e57fe5b906000526020600020906008020160060154915061289e6007838154811061286257fe5b9060005260206000209060080201600201546007848154811061288157fe5b906000526020600020906008020160030154600d54600e54614716565b9050806128ab8489612d0c565b116128bd575060019250612932915050565b6128ee600783815481106128cd57fe5b90600052602060002090600802016002015482612d0c90919063ffffffff16565b8310156129045760008094509450505050612932565b6007858154811061291157fe5b906000526020600020906008020160060154945061280b565b506001925050505b9250929050565b600081612989576040805162461bcd60e51b81526020600482015260196024820152781911995910985b9ace88111150925517d393d517d193d55391603a1b604482015290519081900360640190fd5b6109056007838154811061139557fe5b6001600160a01b03808416600081815260046020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b600c60009054906101000a90046001600160a01b03166001600160a01b031663d7ed714e6040518163ffffffff1660e01b815260040160206040518083038186803b158015612a4957600080fd5b505afa158015612a5d573d6000803e3d6000fd5b505050506040513d6020811015612a7357600080fd5b50514311612a8057612b65565b600c5460408051632df2494b60e21b815290516000926001600160a01b03169163b7c9252c91600480830192602092919082900301818787803b158015612ac657600080fd5b505af1158015612ada573d6000803e3d6000fd5b505050506040513d6020811015612af057600080fd5b505160175490915081118015612b0857506000600254115b15612b6357600254601854600d54601754612b5a939260ff16600a0a91612b4091670de0b6b3a764000091610dc69182908990612b67565b81612b4757fe5b0481612b4f57fe5b601654919004612d0c565b60165560178190555b505b565b80820382811115610905576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6000811580612bd257505080820282828281612bcf57fe5b04145b610905576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b60006003821115612c5d575080600160028204015b81811015612c5757809150600281828581612c4657fe5b040181612c4f57fe5b049050612c2f565b506116c1565b81156116c157506001919050565b600254612c789082612d0c565b6002556001600160a01b038216600090815260036020526040902054612c9e9082612d0c565b6001600160a01b03831660008181526003602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6000818310612d05578161129e565b5090919050565b80820182811015610905576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fd5b600d91909155600e55565b6000808411612da65760405162461bcd60e51b8152600401808060200182810382526026815260200180614ff66026913960400191505060405180910390fd5b600083118015612db65750600082115b612df15760405162461bcd60e51b8152600401808060200182810382526023815260200180614f626023913960400191505060405180910390fd5b6000612dfd8584612bb7565b90506000612e0b8587612d0c565b9050808281612e1657fe5b049695505050505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b178152925182516000946060949389169392918291908083835b60208310612e9e5780518252601f199092019160209182019101612e7f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612f00576040519150601f19603f3d011682016040523d82523d6000602084013e612f05565b606091505b5091509150818015612f33575080511580612f335750808060200190516020811015612f3057600080fd5b50515b612f84576040805162461bcd60e51b815260206004820152601f60248201527f5472616e7366657248656c7065723a205452414e534645525f4641494c454400604482015290519081900360640190fd5b5050505050565b6001600160a01b038316600090815260036020526040902054612fae9082612b67565b6001600160a01b038085166000908152600360205260408082209390935590841681522054612fdd9082612d0c565b6001600160a01b0380841660008181526003602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6007818154811061304657fe5b906000526020600020906008020160060154600780838154811061306657fe5b9060005260206000209060080201600701548154811061308257fe5b906000526020600020906008020160060181905550600781815481106130a457fe5b90600052602060002090600802016007015460078083815481106130c457fe5b906000526020600020906008020160060154815481106130e057fe5b90600052602060002090600802016007018190555061314a6007828154811061310557fe5b906000526020600020906008020160050154610af36007848154811061312757fe5b906000526020600020906008020160020154600f54612b6790919063ffffffff16565b600f55600780548290811061315b57fe5b60009182526020822060089091020180546001600160a01b03191681556001810182905560028101829055600381018290556004810182905560058101829055600681018290556007015550565b60606109056132dc836001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156131ea57600080fd5b505afa1580156131fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561322757600080fd5b810190808051604051939291908464010000000082111561324757600080fd5b90830190602082018581111561325c57600080fd5b825164010000000081118282018810171561327657600080fd5b82525081516020918201929091019080838360005b838110156132a357818101518382015260200161328b565b50505050905090810190601f1680156132d05780820380516001836020036101000a031916815260200191505b5060405250505061478e565b61332a613330613304604051806040016040528060018152602001602d60f81b81525061478e565b61332a6040518060400160405280600581526020016404645444c560dc1b81525061478e565b906147b3565b61478e565b60008061334061483a565b601080546001810190915592509050613357614e7d565b604051806101000160405280856001600160a01b0316815260200184815260200187815260200186815260200160008152602001600081526020016007898154811061339f57fe5b90600052602060002090600802016006015481526020018881525090508160078260c00151815481106133ce57fe5b90600052602060002090600802016007018190555081600788815481106133f157fe5b600091825260209091206006600890920201015560075482141561358c576007805460018101825560009190915281517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688600890920291820180546001600160a01b0319166001600160a01b0390921691909117905560208201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68982015560408201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68a82015560608201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68b82015560808201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68c82015560a08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68d82015560c08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68e82015560e08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68f90910155613613565b806007838154811061359a57fe5b600091825260209182902083516008929092020180546001600160a01b0319166001600160a01b0390921691909117815590820151600182015560408201516002820155606082015160038201556080820151600482015560a0820151600582015560c0820151600682015560e0909101516007909101555b600f546136209087612d0c565b600f55509095945050505050565b6000808291506000600760008154811061364457fe5b90600052602060002090600802016006015490505b80156139b9576007818154811061366c57fe5b90600052602060002090600802016002015460001480156136ab57506007818154811061369557fe5b9060005260206000209060080201600301546000145b1561384957600e546136bd9084612d0c565b613709600783815481106136cd57fe5b906000526020600020906008020160050154600784815481106136ec57fe5b906000526020600020906008020160040154600d54600e54614716565b10156138445760006007828154811061371e57fe5b90600052602060002090600802016006015490506137656007838154811061374257fe5b906000526020600020906008020160050154601254612d0c90919063ffffffff16565b60128190555061379e6007838154811061377b57fe5b906000526020600020906008020160040154601354612d0c90919063ffffffff16565b60135560078054839081106137af57fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f428339815191529190859081106137eb57fe5b906000526020600020906008020160010154600080600080604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a261383d82613039565b9050613659565b6139b9565b600e546138569084612d0c565b6138666007838154811061286257fe5b10156138445760006007828154811061387b57fe5b90600052602060002090600802016006015490506138c06007838154811061389f57fe5b90600052602060002090600802016002015485612d0c90919063ffffffff16565b93506138f3600783815481106138d257fe5b90600052602060002090600802016003015484612d0c90919063ffffffff16565b92506007828154811061390257fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908590811061393e57fe5b906000526020600020906008020160010154600080600080604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a26139976007838154811061374257fe5b6012819055506139ad6007838154811061377b57fe5b60135561383d82613039565b50915091565b6001600160a01b0382166000908152600360205260409020546139e29082612b67565b6001600160a01b038316600090815260036020526040902055600254613a089082612b67565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600080600080600019905060008060008060006007600081548110613a7157fe5b90600052602060002090600802016007015490505b80156147095760078181548110613a9957fe5b9060005260206000209060080201600201546000148015613ad8575060078181548110613ac257fe5b9060005260206000209060080201600301546000145b15613b035760078181548110613aea57fe5b9060005260206000209060080201600701549050614704565b613b4b60078281548110613b1357fe5b90600052602060002090600802016002015460078381548110613b3257fe5b9060005260206000209060080201600301548d8d614716565b94508985101561418c5789613b8760078381548110613b6657fe5b90600052602060002090600802016002015487612d0c90919063ffffffff16565b11613e3557613bbd60078281548110613b9c57fe5b9060005260206000209060080201600201548a612d0c90919063ffffffff16565b9850613bf060078281548110613bcf57fe5b90600052602060002090600802016003015489612d0c90919063ffffffff16565b975060078181548110613bff57fe5b906000526020600020906008020160070154915060078181548110613c2057fe5b9060005260206000209060080201600401546000148015613c5f575060078181548110613c4957fe5b9060005260206000209060080201600501546000145b15613d045760078181548110613c7157fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f42833981519152919084908110613cad57fe5b906000526020600020906008020160010154600080600080604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a2613cff81613039565b613e2e565b613d146007828154811061312757fe5b600f556007805482908110613d2557fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f42833981519152919084908110613d6157fe5b90600052602060002090600802016001015460008060078681548110613d8357fe5b90600052602060002090600802016004015460078781548110613da257fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a2600060078281548110613df457fe5b906000526020600020906008020160020181905550600060078281548110613e1857fe5b9060005260206000209060080201600301819055505b5080614704565b613e3f8a86612b67565b925084613e4c8c85612bb7565b81613e5357fe5b04935060078181548110613e6357fe5b906000526020600020906008020160030154841180613e9f575060078181548110613e8a57fe5b90600052602060002090600802016002015483115b15613ee75760078181548110613eb157fe5b906000526020600020906008020160030154935060078181548110613ed257fe5b90600052602060002090600802016002015492505b613ef18984612d0c565b9850613efd8885612d0c565b9750613f308360078381548110613f1057fe5b906000526020600020906008020160020154612b6790919063ffffffff16565b60078281548110613f3d57fe5b906000526020600020906008020160020181905550613f838460078381548110613f6357fe5b906000526020600020906008020160030154612b6790919063ffffffff16565b60078281548110613f9057fe5b6000918252602090912060036008909202010155600f54613fb19084612b67565b600f55600780548b965082908110613fc557fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908490811061400157fe5b9060005260206000209060080201600101546007848154811061402057fe5b9060005260206000209060080201600201546007858154811061403f57fe5b9060005260206000209060080201600301546007868154811061405e57fe5b9060005260206000209060080201600401546007878154811061407d57fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a2600781815481106140cd57fe5b9060005260206000209060080201600701549150600781815481106140ee57fe5b906000526020600020906008020160030154600014801561412d57506007818154811061411757fe5b9060005260206000209060080201600201546000145b1561418c576007818154811061413f57fe5b906000526020600020906008020160040154600014801561417e57506007818154811061416857fe5b9060005260206000209060080201600501546000145b15613e2e57613e2e81613039565b8585106143e1576141a360078281548110613bcf57fe5b97506141d6600782815481106141b557fe5b90600052602060002090600802016003015488612d0c90919063ffffffff16565b9650614227600782815481106141e857fe5b9060005260206000209060080201600301546007838154811061420757fe5b906000526020600020906008020160040154612d0c90919063ffffffff16565b6007828154811061423457fe5b9060005260206000209060080201600401819055506142986007828154811061425957fe5b9060005260206000209060080201600201546007838154811061427857fe5b906000526020600020906008020160050154612d0c90919063ffffffff16565b600782815481106142a557fe5b906000526020600020906008020160050181905550600781815481106142c757fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908490811061430357fe5b9060005260206000209060080201600101546000806007868154811061432557fe5b9060005260206000209060080201600401546007878154811061434457fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a260006007828154811061439657fe5b9060005260206000209060080201600201819055506000600782815481106143ba57fe5b9060005260206000209060080201600301819055506143da818c8c61487e565b9050614704565b856143f260078381548110613b6657fe5b11156146df5761442d86610af36007848154811061440c57fe5b90600052602060002090600802016002015488612d0c90919063ffffffff16565b925061444961444260078381548110613b6657fe5b8790612bb7565b61445784610dc68e8e612bb7565b8161445e57fe5b0493506007818154811061446e57fe5b9060005260206000209060080201600301548411806144aa57506007818154811061449557fe5b90600052602060002090600802016002015483115b156144f257600781815481106144bc57fe5b9060005260206000209060080201600301549350600781815481106144dd57fe5b90600052602060002090600802016002015492505b6144fc8885612d0c565b97506145088785612d0c565b965061451b846007838154811061420757fe5b6007828154811061452857fe5b90600052602060002090600802016004018190555061454e836007838154811061427857fe5b6007828154811061455b57fe5b9060005260206000209060080201600501819055506145818360078381548110613f1057fe5b6007828154811061458e57fe5b9060005260206000209060080201600201819055506145b48460078381548110613f6357fe5b600782815481106145c157fe5b906000526020600020906008020160030181905550600781815481106145e357fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908490811061461f57fe5b9060005260206000209060080201600101546007848154811061463e57fe5b9060005260206000209060080201600201546007858154811061465d57fe5b9060005260206000209060080201600301546007868154811061467c57fe5b9060005260206000209060080201600401546007878154811061469b57fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a25b849550600781815481106146ef57fe5b90600052602060002090600802016007015490505b613a86565b5050505050509250925092565b6000806147238484612bb7565b905085810260008288838161473457fe5b041461474c57614745838989614c68565b9050614759565b86828161475557fe5b0490505b61478260028904610af3610c3860046147728d80612bb7565b8161477957fe5b86919004612d0c565b98975050505050505050565b614796614ecb565b506040805180820190915281518152602082810190820152919050565b8051825160609182910167ffffffffffffffff811180156147d357600080fd5b506040519080825280601f01601f1916602001820160405280156147fe576020820181803683370190505b509050600060208201905061481c8186602001518760000151614d1f565b8451602085015185516148329284019190614d1f565b509392505050565b60015b60075481101561487b576007818154811061485457fe5b906000526020600020906008020160010154600014156148735761487b565b60010161483d565b90565b6000806007858154811061488e57fe5b9060005260206000209060080201600701549050600785815481106148af57fe5b90600052602060002090600802016006015460078087815481106148cf57fe5b906000526020600020906008020160070154815481106148eb57fe5b90600052602060002090600802016006018190555080600780878154811061490f57fe5b9060005260206000209060080201600601548154811061492b57fe5b906000526020600020906008020160070181905550600061498a6007878154811061495257fe5b9060005260206000209060080201600501546007888154811061497157fe5b9060005260206000209060080201600401548787614716565b9050600080600760008154811061499d57fe5b90600052602060002090600802016006015490505b8015614bcf57600781815481106149c557fe5b9060005260206000209060080201600201546000148015614a045750600781815481106149ee57fe5b9060005260206000209060080201600301546000145b15614a5857614a5160078281548110614a1957fe5b90600052602060002090600802016005015460078381548110614a3857fe5b9060005260206000209060080201600401548989614716565b9150614aa3565b614aa060078281548110614a6857fe5b90600052602060002090600802016002015460078381548110614a8757fe5b9060005260206000209060080201600301548989614716565b91505b818311614b7d5760078181548110614ab757fe5b90600052602060002090600802016007015460078981548110614ad657fe5b906000526020600020906008020160070181905550876007808381548110614afa57fe5b90600052602060002090600802016007015481548110614b1657fe5b9060005260206000209060080201600601819055508060078981548110614b3957fe5b9060005260206000209060080201600601819055508760078281548110614b5c57fe5b9060005260206000209060080201600701819055508394505050505061129e565b60078181548110614b8a57fe5b90600052602060002090600802016006015460001415614ba957614bcf565b60078181548110614bb657fe5b90600052602060002090600802016006015490506149b2565b8060078981548110614bdd57fe5b9060005260206000209060080201600701819055508760078281548110614c0057fe5b906000526020600020906008020160060181905550600060078981548110614c2457fe5b906000526020600020906008020160060181905550876007600081548110614c4857fe5b600091825260209091206008909102016007015550919695505050505050565b6000806000614c778686614d5d565b91509150838110614c8757600080fd5b60008480614c9157fe5b868809905082811115614ca5576001820391505b918290039160008590038516808681614cba57fe5b049550808481614cc657fe5b049350808160000381614cd557fe5b046001019290920292909201600285810380870282030280870282030280870282030280870282030280870282030280870282030295860290039094029390930295945050505050565b5b60208110614d3f578151835260209283019290910190601f1901614d20565b905182516020929092036101000a6000190180199091169116179052565b6000808060001984860990508385029250828103915082811015614d82576001820391505b509250929050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614dcb57805160ff1916838001178555614df8565b82800160010185558215614df8579182015b82811115614df8578251825591602001919060010190614ddd565b50614e04929150614ee5565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614e415780548555614df8565b82800160010185558215614df857600052602060002091601f016020900482015b82811115614df8578254825591600101919060010190614e62565b60405180610100016040528060006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604051806040016040528060008152602001600081525090565b5b80821115614e045760008155600101614ee656fe64466564506169723a20494e53554646494349454e545f4f55545055545f414d4f554e546446656442616e6b3a20494e53554646494349454e545f494e5055545f414d4f554e54a0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869644665644c6962726172793a20494e53554646494349454e545f4c495155494449545964466564506169723a20494e53554646494349454e545f494e5055545f414d4f554e5464466564506169723a20494e53554646494349454e545f4c49515549444954595f4255524e454464466564506169723a20494e53554646494349454e545f4c49515549444954595f4d494e544544644665644c6962726172793a20494e53554646494349454e545f494e5055545f414d4f554e54a2646970667358221220af73e875362f7f20e3c9b3e8303ddea008ac026d22568c63cba01f015a68f3d764736f6c63430007000033a26469706673582212207bb58d6154d4cfa497d4fe63ab16f8be05a446150b5c1acb3f8dd5646d201abb64736f6c63430007000033"

// DeployDFedCreator deploys a new Ethereum contract, binding an instance of DFedCreator to it.
func DeployDFedCreator(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address) (common.Address, *types.Transaction, *DFedCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DFedCreatorBin), backend, _factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DFedCreator{DFedCreatorCaller: DFedCreatorCaller{contract: contract}, DFedCreatorTransactor: DFedCreatorTransactor{contract: contract}, DFedCreatorFilterer: DFedCreatorFilterer{contract: contract}}, nil
}

// DFedCreator is an auto generated Go binding around an Ethereum contract.
type DFedCreator struct {
	DFedCreatorCaller     // Read-only binding to the contract
	DFedCreatorTransactor // Write-only binding to the contract
	DFedCreatorFilterer   // Log filterer for contract events
}

// DFedCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DFedCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DFedCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DFedCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DFedCreatorSession struct {
	Contract     *DFedCreator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DFedCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DFedCreatorCallerSession struct {
	Contract *DFedCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DFedCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DFedCreatorTransactorSession struct {
	Contract     *DFedCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DFedCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DFedCreatorRaw struct {
	Contract *DFedCreator // Generic contract binding to access the raw methods on
}

// DFedCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DFedCreatorCallerRaw struct {
	Contract *DFedCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// DFedCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DFedCreatorTransactorRaw struct {
	Contract *DFedCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDFedCreator creates a new instance of DFedCreator, bound to a specific deployed contract.
func NewDFedCreator(address common.Address, backend bind.ContractBackend) (*DFedCreator, error) {
	contract, err := bindDFedCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DFedCreator{DFedCreatorCaller: DFedCreatorCaller{contract: contract}, DFedCreatorTransactor: DFedCreatorTransactor{contract: contract}, DFedCreatorFilterer: DFedCreatorFilterer{contract: contract}}, nil
}

// NewDFedCreatorCaller creates a new read-only instance of DFedCreator, bound to a specific deployed contract.
func NewDFedCreatorCaller(address common.Address, caller bind.ContractCaller) (*DFedCreatorCaller, error) {
	contract, err := bindDFedCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DFedCreatorCaller{contract: contract}, nil
}

// NewDFedCreatorTransactor creates a new write-only instance of DFedCreator, bound to a specific deployed contract.
func NewDFedCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DFedCreatorTransactor, error) {
	contract, err := bindDFedCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DFedCreatorTransactor{contract: contract}, nil
}

// NewDFedCreatorFilterer creates a new log filterer instance of DFedCreator, bound to a specific deployed contract.
func NewDFedCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DFedCreatorFilterer, error) {
	contract, err := bindDFedCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DFedCreatorFilterer{contract: contract}, nil
}

// bindDFedCreator binds a generic wrapper to an already deployed contract.
func bindDFedCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedCreator *DFedCreatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedCreator.Contract.DFedCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedCreator *DFedCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedCreator.Contract.DFedCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedCreator *DFedCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedCreator.Contract.DFedCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedCreator *DFedCreatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedCreator *DFedCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedCreator *DFedCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedCreator.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedCreator *DFedCreatorCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedCreator.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedCreator *DFedCreatorSession) Factory() (common.Address, error) {
	return _DFedCreator.Contract.Factory(&_DFedCreator.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedCreator *DFedCreatorCallerSession) Factory() (common.Address, error) {
	return _DFedCreator.Contract.Factory(&_DFedCreator.CallOpts)
}

// GetInitHash is a free data retrieval call binding the contract method 0x3c9de1b8.
//
// Solidity: function getInitHash() view returns(bytes32)
func (_DFedCreator *DFedCreatorCaller) GetInitHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DFedCreator.contract.Call(opts, out, "getInitHash")
	return *ret0, err
}

// GetInitHash is a free data retrieval call binding the contract method 0x3c9de1b8.
//
// Solidity: function getInitHash() view returns(bytes32)
func (_DFedCreator *DFedCreatorSession) GetInitHash() ([32]byte, error) {
	return _DFedCreator.Contract.GetInitHash(&_DFedCreator.CallOpts)
}

// GetInitHash is a free data retrieval call binding the contract method 0x3c9de1b8.
//
// Solidity: function getInitHash() view returns(bytes32)
func (_DFedCreator *DFedCreatorCallerSession) GetInitHash() ([32]byte, error) {
	return _DFedCreator.Contract.GetInitHash(&_DFedCreator.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _token, address _baseToken) returns(address _pair)
func (_DFedCreator *DFedCreatorTransactor) CreatePair(opts *bind.TransactOpts, _token common.Address, _baseToken common.Address) (*types.Transaction, error) {
	return _DFedCreator.contract.Transact(opts, "createPair", _token, _baseToken)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _token, address _baseToken) returns(address _pair)
func (_DFedCreator *DFedCreatorSession) CreatePair(_token common.Address, _baseToken common.Address) (*types.Transaction, error) {
	return _DFedCreator.Contract.CreatePair(&_DFedCreator.TransactOpts, _token, _baseToken)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address _token, address _baseToken) returns(address _pair)
func (_DFedCreator *DFedCreatorTransactorSession) CreatePair(_token common.Address, _baseToken common.Address) (*types.Transaction, error) {
	return _DFedCreator.Contract.CreatePair(&_DFedCreator.TransactOpts, _token, _baseToken)
}

// DFedERC20ABI is the input ABI used to generate the binding from.
const DFedERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DFedERC20FuncSigs maps the 4-byte function signature to its string representation.
var DFedERC20FuncSigs = map[string]string{
	"3644e515": "DOMAIN_SEPARATOR()",
	"30adf81f": "PERMIT_TYPEHASH()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"7ecebe00": "nonces(address)",
	"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// DFedERC20Bin is the compiled bytecode used for deploying new contracts.
var DFedERC20Bin = "0x608060405234801561001057600080fd5b5061090d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80633644e5151161008c57806395d89b411161006657806395d89b411461025b578063a9059cbb14610263578063d505accf1461028f578063dd62ed3e146102e2576100cf565b80633644e5151461020757806370a082311461020f5780637ecebe0014610235576100cf565b806306fdde03146100d4578063095ea7b31461015157806318160ddd1461019157806323b872dd146101ab57806330adf81f146101e1578063313ce567146101e9575b600080fd5b6100dc610310565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101165781810151838201526020016100fe565b50505050905090810190601f1680156101435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61017d6004803603604081101561016757600080fd5b506001600160a01b03813516906020013561039e565b604080519115158252519081900360200190f35b6101996103b5565b60408051918252519081900360200190f35b61017d600480360360608110156101c157600080fd5b506001600160a01b038135811691602081013590911690604001356103bb565b61019961044f565b6101f1610473565b6040805160ff9092168252519081900360200190f35b610199610478565b6101996004803603602081101561022557600080fd5b50356001600160a01b031661047e565b6101996004803603602081101561024b57600080fd5b50356001600160a01b0316610490565b6100dc6104a2565b61017d6004803603604081101561027957600080fd5b506001600160a01b0381351690602001356104fc565b6102e0600480360360e08110156102a557600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c00135610509565b005b610199600480360360408110156102f857600080fd5b506001600160a01b038135811691602001351661070b565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103965780601f1061036b57610100808354040283529160200191610396565b820191906000526020600020905b81548152906001019060200180831161037957829003601f168201915b505050505081565b60006103ab338484610728565b5060015b92915050565b60025481565b6001600160a01b03831660009081526004602090815260408083203384529091528120546000191461043a576001600160a01b0384166000908152600460209081526040808320338452909152902054610415908361078a565b6001600160a01b03851660009081526004602090815260408083203384529091529020555b6104458484846107da565b5060019392505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60055481565b60036020526000908152604090205481565b60066020526000908152604090205481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103965780601f1061036b57610100808354040283529160200191610396565b60006103ab3384846107da565b42841015610553576040805162461bcd60e51b815260206004820152601260248201527119119959115490cc8c0e881156141254915160721b604482015290519081900360640190fd5b6005546001600160a01b0380891660008181526006602090815260408083208054600180820190925582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98186015280840196909652958d166060860152608085018c905260a085019590955260c08085018b90528151808603909101815260e08501825280519083012061190160f01b6101008601526101028501969096526101228085019690965280518085039096018652610142840180825286519683019690962095839052610162840180825286905260ff89166101828501526101a284018890526101c28401879052519193926101e280820193601f1981019281900390910190855afa15801561066e573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116158015906106a45750886001600160a01b0316816001600160a01b0316145b6106f5576040805162461bcd60e51b815260206004820152601c60248201527f6446656445524332303a20494e56414c49445f5349474e415455524500000000604482015290519081900360640190fd5b610700898989610728565b505050505050505050565b600460209081526000928352604080842090915290825290205481565b6001600160a01b03808416600081815260046020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b808203828111156103af576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6001600160a01b0383166000908152600360205260409020546107fd908261078a565b6001600160a01b03808516600090815260036020526040808220939093559084168152205461082c9082610888565b6001600160a01b0380841660008181526003602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b808201828110156103af576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fdfea264697066735822122002c567ec3a100c2b720786f8e82e6b81aa26bec72699318f81158803611180e764736f6c63430007000033"

// DeployDFedERC20 deploys a new Ethereum contract, binding an instance of DFedERC20 to it.
func DeployDFedERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DFedERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DFedERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DFedERC20{DFedERC20Caller: DFedERC20Caller{contract: contract}, DFedERC20Transactor: DFedERC20Transactor{contract: contract}, DFedERC20Filterer: DFedERC20Filterer{contract: contract}}, nil
}

// DFedERC20 is an auto generated Go binding around an Ethereum contract.
type DFedERC20 struct {
	DFedERC20Caller     // Read-only binding to the contract
	DFedERC20Transactor // Write-only binding to the contract
	DFedERC20Filterer   // Log filterer for contract events
}

// DFedERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type DFedERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DFedERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DFedERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DFedERC20Session struct {
	Contract     *DFedERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DFedERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DFedERC20CallerSession struct {
	Contract *DFedERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DFedERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DFedERC20TransactorSession struct {
	Contract     *DFedERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DFedERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type DFedERC20Raw struct {
	Contract *DFedERC20 // Generic contract binding to access the raw methods on
}

// DFedERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DFedERC20CallerRaw struct {
	Contract *DFedERC20Caller // Generic read-only contract binding to access the raw methods on
}

// DFedERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DFedERC20TransactorRaw struct {
	Contract *DFedERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDFedERC20 creates a new instance of DFedERC20, bound to a specific deployed contract.
func NewDFedERC20(address common.Address, backend bind.ContractBackend) (*DFedERC20, error) {
	contract, err := bindDFedERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DFedERC20{DFedERC20Caller: DFedERC20Caller{contract: contract}, DFedERC20Transactor: DFedERC20Transactor{contract: contract}, DFedERC20Filterer: DFedERC20Filterer{contract: contract}}, nil
}

// NewDFedERC20Caller creates a new read-only instance of DFedERC20, bound to a specific deployed contract.
func NewDFedERC20Caller(address common.Address, caller bind.ContractCaller) (*DFedERC20Caller, error) {
	contract, err := bindDFedERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DFedERC20Caller{contract: contract}, nil
}

// NewDFedERC20Transactor creates a new write-only instance of DFedERC20, bound to a specific deployed contract.
func NewDFedERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*DFedERC20Transactor, error) {
	contract, err := bindDFedERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DFedERC20Transactor{contract: contract}, nil
}

// NewDFedERC20Filterer creates a new log filterer instance of DFedERC20, bound to a specific deployed contract.
func NewDFedERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*DFedERC20Filterer, error) {
	contract, err := bindDFedERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DFedERC20Filterer{contract: contract}, nil
}

// bindDFedERC20 binds a generic wrapper to an already deployed contract.
func bindDFedERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedERC20 *DFedERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedERC20.Contract.DFedERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedERC20 *DFedERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedERC20.Contract.DFedERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedERC20 *DFedERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedERC20.Contract.DFedERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedERC20 *DFedERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedERC20 *DFedERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedERC20 *DFedERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedERC20 *DFedERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedERC20 *DFedERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _DFedERC20.Contract.DOMAINSEPARATOR(&_DFedERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedERC20 *DFedERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DFedERC20.Contract.DOMAINSEPARATOR(&_DFedERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedERC20 *DFedERC20Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedERC20 *DFedERC20Session) PERMITTYPEHASH() ([32]byte, error) {
	return _DFedERC20.Contract.PERMITTYPEHASH(&_DFedERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedERC20 *DFedERC20CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _DFedERC20.Contract.PERMITTYPEHASH(&_DFedERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedERC20 *DFedERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedERC20 *DFedERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DFedERC20.Contract.Allowance(&_DFedERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedERC20 *DFedERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DFedERC20.Contract.Allowance(&_DFedERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedERC20 *DFedERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedERC20 *DFedERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DFedERC20.Contract.BalanceOf(&_DFedERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedERC20 *DFedERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DFedERC20.Contract.BalanceOf(&_DFedERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedERC20 *DFedERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedERC20 *DFedERC20Session) Decimals() (uint8, error) {
	return _DFedERC20.Contract.Decimals(&_DFedERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedERC20 *DFedERC20CallerSession) Decimals() (uint8, error) {
	return _DFedERC20.Contract.Decimals(&_DFedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedERC20 *DFedERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedERC20 *DFedERC20Session) Name() (string, error) {
	return _DFedERC20.Contract.Name(&_DFedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedERC20 *DFedERC20CallerSession) Name() (string, error) {
	return _DFedERC20.Contract.Name(&_DFedERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedERC20 *DFedERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedERC20 *DFedERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DFedERC20.Contract.Nonces(&_DFedERC20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedERC20 *DFedERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DFedERC20.Contract.Nonces(&_DFedERC20.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedERC20 *DFedERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedERC20 *DFedERC20Session) Symbol() (string, error) {
	return _DFedERC20.Contract.Symbol(&_DFedERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedERC20 *DFedERC20CallerSession) Symbol() (string, error) {
	return _DFedERC20.Contract.Symbol(&_DFedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedERC20 *DFedERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedERC20 *DFedERC20Session) TotalSupply() (*big.Int, error) {
	return _DFedERC20.Contract.TotalSupply(&_DFedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedERC20 *DFedERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _DFedERC20.Contract.TotalSupply(&_DFedERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.Contract.Approve(&_DFedERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.Contract.Approve(&_DFedERC20.TransactOpts, _spender, _value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedERC20 *DFedERC20Transactor) Permit(opts *bind.TransactOpts, _from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedERC20.contract.Transact(opts, "permit", _from, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedERC20 *DFedERC20Session) Permit(_from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedERC20.Contract.Permit(&_DFedERC20.TransactOpts, _from, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedERC20 *DFedERC20TransactorSession) Permit(_from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedERC20.Contract.Permit(&_DFedERC20.TransactOpts, _from, _spender, _value, _deadline, _v, _r, _s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.Contract.Transfer(&_DFedERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.Contract.Transfer(&_DFedERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.Contract.TransferFrom(&_DFedERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedERC20 *DFedERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedERC20.Contract.TransferFrom(&_DFedERC20.TransactOpts, _from, _to, _value)
}

// DFedERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DFedERC20 contract.
type DFedERC20ApprovalIterator struct {
	Event *DFedERC20Approval // Event containing the contract specifics and raw log

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
func (it *DFedERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedERC20Approval)
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
		it.Event = new(DFedERC20Approval)
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
func (it *DFedERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedERC20Approval represents a Approval event raised by the DFedERC20 contract.
type DFedERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DFedERC20 *DFedERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DFedERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DFedERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DFedERC20ApprovalIterator{contract: _DFedERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DFedERC20 *DFedERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DFedERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DFedERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedERC20Approval)
				if err := _DFedERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DFedERC20 *DFedERC20Filterer) ParseApproval(log types.Log) (*DFedERC20Approval, error) {
	event := new(DFedERC20Approval)
	if err := _DFedERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DFedERC20 contract.
type DFedERC20TransferIterator struct {
	Event *DFedERC20Transfer // Event containing the contract specifics and raw log

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
func (it *DFedERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedERC20Transfer)
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
		it.Event = new(DFedERC20Transfer)
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
func (it *DFedERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedERC20Transfer represents a Transfer event raised by the DFedERC20 contract.
type DFedERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DFedERC20 *DFedERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DFedERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DFedERC20TransferIterator{contract: _DFedERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DFedERC20 *DFedERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DFedERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedERC20Transfer)
				if err := _DFedERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DFedERC20 *DFedERC20Filterer) ParseTransfer(log types.Log) (*DFedERC20Transfer, error) {
	event := new(DFedERC20Transfer)
	if err := _DFedERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedLibraryABI is the input ABI used to generate the binding from.
const DFedLibraryABI = "[]"

// DFedLibraryBin is the compiled bytecode used for deploying new contracts.
var DFedLibraryBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203cf290d73fa3171bd14fb36a47562523f9700a9fc846ad6a1b8dc5968b02fb6364736f6c63430007000033"

// DeployDFedLibrary deploys a new Ethereum contract, binding an instance of DFedLibrary to it.
func DeployDFedLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DFedLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DFedLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DFedLibrary{DFedLibraryCaller: DFedLibraryCaller{contract: contract}, DFedLibraryTransactor: DFedLibraryTransactor{contract: contract}, DFedLibraryFilterer: DFedLibraryFilterer{contract: contract}}, nil
}

// DFedLibrary is an auto generated Go binding around an Ethereum contract.
type DFedLibrary struct {
	DFedLibraryCaller     // Read-only binding to the contract
	DFedLibraryTransactor // Write-only binding to the contract
	DFedLibraryFilterer   // Log filterer for contract events
}

// DFedLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DFedLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DFedLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DFedLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DFedLibrarySession struct {
	Contract     *DFedLibrary      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DFedLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DFedLibraryCallerSession struct {
	Contract *DFedLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DFedLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DFedLibraryTransactorSession struct {
	Contract     *DFedLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DFedLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DFedLibraryRaw struct {
	Contract *DFedLibrary // Generic contract binding to access the raw methods on
}

// DFedLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DFedLibraryCallerRaw struct {
	Contract *DFedLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// DFedLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DFedLibraryTransactorRaw struct {
	Contract *DFedLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDFedLibrary creates a new instance of DFedLibrary, bound to a specific deployed contract.
func NewDFedLibrary(address common.Address, backend bind.ContractBackend) (*DFedLibrary, error) {
	contract, err := bindDFedLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DFedLibrary{DFedLibraryCaller: DFedLibraryCaller{contract: contract}, DFedLibraryTransactor: DFedLibraryTransactor{contract: contract}, DFedLibraryFilterer: DFedLibraryFilterer{contract: contract}}, nil
}

// NewDFedLibraryCaller creates a new read-only instance of DFedLibrary, bound to a specific deployed contract.
func NewDFedLibraryCaller(address common.Address, caller bind.ContractCaller) (*DFedLibraryCaller, error) {
	contract, err := bindDFedLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DFedLibraryCaller{contract: contract}, nil
}

// NewDFedLibraryTransactor creates a new write-only instance of DFedLibrary, bound to a specific deployed contract.
func NewDFedLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*DFedLibraryTransactor, error) {
	contract, err := bindDFedLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DFedLibraryTransactor{contract: contract}, nil
}

// NewDFedLibraryFilterer creates a new log filterer instance of DFedLibrary, bound to a specific deployed contract.
func NewDFedLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*DFedLibraryFilterer, error) {
	contract, err := bindDFedLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DFedLibraryFilterer{contract: contract}, nil
}

// bindDFedLibrary binds a generic wrapper to an already deployed contract.
func bindDFedLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedLibrary *DFedLibraryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedLibrary.Contract.DFedLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedLibrary *DFedLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedLibrary.Contract.DFedLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedLibrary *DFedLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedLibrary.Contract.DFedLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedLibrary *DFedLibraryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedLibrary *DFedLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedLibrary *DFedLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedLibrary.Contract.contract.Transact(opts, method, params...)
}

// DFedPairABI is the input ABI used to generate the binding from.
const DFedPairABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BuyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToken0Amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToken1Amount\",\"type\":\"uint256\"}],\"name\":\"DebtUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetNum\",\"type\":\"uint256\"}],\"name\":\"Mortgage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SellToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accRewardPerFEDlpOfPair\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"debtInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToken0Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToken1Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRatePer10K\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"name\":\"getDebtIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRepayByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetNum\",\"type\":\"uint256\"}],\"name\":\"isValidDebit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReceivedRewardPerUSDDOfPair\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mortgage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"refundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRefundToken0Amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRefundToken1Amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRefundToken1EqualToken0Amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniqueDebtId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardDebtPerFEDlp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DFedPairFuncSigs maps the 4-byte function signature to its string representation.
var DFedPairFuncSigs = map[string]string{
	"3644e515": "DOMAIN_SEPARATOR()",
	"ba9a7a56": "MINIMUM_LIQUIDITY()",
	"30adf81f": "PERMIT_TYPEHASH()",
	"7dea576e": "accRewardPerFEDlpOfPair()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"89afcb44": "burn(address)",
	"21466eb5": "buyToken(uint256,uint256,address)",
	"e3c20fd5": "debtInfos(uint256)",
	"313ce567": "decimals()",
	"c45a0155": "factory()",
	"ddca3f43": "fee()",
	"a48685fe": "feeRatePer10K()",
	"3c33a245": "getDebtIndexById(uint256)",
	"ea52a0be": "getRepayByIndex(uint256)",
	"0902f1ac": "getReserves()",
	"0e5c011e": "harvest(address)",
	"485cc955": "initialize(address,address)",
	"e991685f": "isValidDebit(uint256,uint256)",
	"a1f04001": "lastReceivedRewardPerUSDDOfPair()",
	"156e29f6": "mint(address,uint256,uint256)",
	"58687e63": "mortgage(uint256,uint256,address)",
	"06fdde03": "name()",
	"7ecebe00": "nonces(address)",
	"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	"acde5d28": "refundAmount(address)",
	"371fd8e6": "repay(uint256)",
	"443cb4bc": "reserve0()",
	"5a76f25e": "reserve1()",
	"88036ac5": "sellToken(uint256,uint256,address)",
	"95d89b41": "symbol()",
	"0dfe1681": "token0()",
	"d21220a7": "token1()",
	"c21a43e4": "totalPledge()",
	"a058274f": "totalRefundToken0Amount()",
	"076195b1": "totalRefundToken1Amount()",
	"5b863ab9": "totalRefundToken1EqualToken0Amount()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"cfc6aa9e": "uniqueDebtId()",
	"d6fedef8": "userRewardDebtPerFEDlp(address)",
	"e60688cc": "withdrawToken0(address,uint256)",
	"7c3a2473": "withdrawToken1(address,uint256)",
}

// DFedPairBin is the compiled bytecode used for deploying new contracts.
var DFedPairBin = "0x608060405260088054600160ff19918216811790925560109190915560158054909116905534801561003057600080fd5b506040805161010081018252600080825260208201818152928201818152606083018281526080840183815260a0850184815260c0860185815260e08701868152600780546001810182559752965160089096027fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688810180546001600160a01b03989098166001600160a01b03199098169790971790965596517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68986015592517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68a85015590517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68b840155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68c830155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68d82015591517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68e830155517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68f90910155615051806101df6000396000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c80637dea576e1161015c578063c21a43e4116100ce578063dd62ed3e11610087578063dd62ed3e14610733578063ddca3f4314610761578063e3c20fd514610769578063e60688cc146107d1578063e991685f146107fd578063ea52a0be146108395761027f565b8063c21a43e41461069c578063c45a0155146106a4578063cfc6aa9e146106ac578063d21220a7146106b4578063d505accf146106bc578063d6fedef81461070d5761027f565b8063a058274f11610120578063a058274f1461062a578063a1f0400114610632578063a48685fe1461063a578063a9059cbb14610642578063acde5d281461066e578063ba9a7a56146106945761027f565b80637dea576e1461059c5780637ecebe00146105a457806388036ac5146105ca57806389afcb44146105fc57806395d89b41146106225761027f565b8063313ce567116101f5578063485cc955116101b9578063485cc955146104da57806358687e63146105085780635a76f25e1461053a5780635b863ab91461054257806370a082311461054a5780637c3a2473146105705761027f565b8063313ce567146104725780633644e51514610490578063371fd8e6146104985780633c33a245146104b5578063443cb4bc146104d25761027f565b80630e5c011e116102475780630e5c011e146103a0578063156e29f6146103c857806318160ddd146103fa57806321466eb51461040257806323b872dd1461043457806330adf81f1461046a5761027f565b806306fdde0314610284578063076195b1146103015780630902f1ac1461031b578063095ea7b31461033c5780630dfe16811461037c575b600080fd5b61028c610856565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102c65781810151838201526020016102ae565b50505050905090810190601f1680156102f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103096108e4565b60408051918252519081900360200190f35b6103236108ea565b6040805192835260208301919091528051918290030190f35b6103686004803603604081101561035257600080fd5b506001600160a01b0381351690602001356108f4565b604080519115158252519081900360200190f35b61038461090b565b604080516001600160a01b039092168252519081900360200190f35b6103c6600480360360208110156103b657600080fd5b50356001600160a01b031661091a565b005b610309600480360360608110156103de57600080fd5b506001600160a01b038135169060208101359060400135610a0e565b610309610f1a565b6103c66004803603606081101561041857600080fd5b50803590602081013590604001356001600160a01b0316610f20565b6103686004803603606081101561044a57600080fd5b506001600160a01b03813581169160208101359091169060400135611210565b6103096112a5565b61047a6112c9565b6040805160ff9092168252519081900360200190f35b6103096112ce565b6103c6600480360360208110156104ae57600080fd5b50356112d4565b610309600480360360208110156104cb57600080fd5b503561163c565b6103096116c6565b6103c6600480360360408110156104f057600080fd5b506001600160a01b03813581169160200135166116cc565b6103096004803603606081101561051e57600080fd5b50803590602081013590604001356001600160a01b031661199d565b610309611c6f565b610309611c75565b6103096004803603602081101561056057600080fd5b50356001600160a01b0316611c7b565b6103c66004803603604081101561058657600080fd5b506001600160a01b038135169060200135611c8d565b610309611d31565b610309600480360360208110156105ba57600080fd5b50356001600160a01b0316611d37565b6103c6600480360360608110156105e057600080fd5b50803590602081013590604001356001600160a01b0316611d49565b6103236004803603602081101561061257600080fd5b50356001600160a01b03166120b1565b61028c61241a565b610309612474565b61030961247a565b61047a612480565b6103686004803603604081101561065857600080fd5b506001600160a01b038135169060200135612489565b6103096004803603602081101561068457600080fd5b50356001600160a01b0316612496565b6103096124a8565b6103096124ae565b6103846124b4565b6103096124c3565b6103846124c9565b6103c6600480360360e08110156106d257600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c001356124d8565b6103096004803603602081101561072357600080fd5b50356001600160a01b03166126da565b6103096004803603604081101561074957600080fd5b506001600160a01b03813581169160200135166126ec565b610309612709565b6107866004803603602081101561077f57600080fd5b503561270f565b604080516001600160a01b0390991689526020890197909752878701959095526060870193909352608086019190915260a085015260c084015260e083015251908190036101000190f35b6103c6600480360360408110156107e757600080fd5b506001600160a01b03813516906020013561276c565b6108206004803603604081101561081357600080fd5b50803590602001356127d9565b6040805192835290151560208301528051918290030190f35b6103096004803603602081101561084f57600080fd5b5035612939565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108dc5780601f106108b1576101008083540402835291602001916108dc565b820191906000526020600020905b8154815290600101906020018083116108bf57829003601f168201915b505050505081565b60125481565b600d54600e549091565b6000610901338484612999565b5060015b92915050565b600a546001600160a01b031681565b6109226129fb565b600c546001600160a01b03828116600090815260196020526040902054601654919092169163ccd072e5918491670de0b6b3a764000091610986916109679190612b67565b6001600160a01b03871660009081526003602052604090205490612bb7565b8161098d57fe5b046040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156109d457600080fd5b505af11580156109e8573d6000803e3d6000fd5b50506016546001600160a01b039093166000908152601960205260409020929092555050565b60085460009060ff16610a5b576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff19169055600954600d54600a54604080516370a0823160e01b81523060048201529051610af99493610af39390926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b158015610ac157600080fd5b505afa158015610ad5573d6000803e3d6000fd5b505050506040513d6020811015610aeb57600080fd5b505190612b67565b90612b67565b831115610b4d576040805162461bcd60e51b815260206004820152601f60248201527f64466564506169723a20494e53554646494349454e545f414d4f554e545f3000604482015290519081900360640190fd5b610bc1601254610af3600f54610af3600e54600b60009054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610ac157600080fd5b821115610c15576040805162461bcd60e51b815260206004820152601f60248201527f64466564506169723a20494e53554646494349454e545f414d4f554e545f3100604482015290519081900360640190fd5b600254610c206129fb565b80610c5157610c3d6103e8610af3610c388787612bb7565b612c1a565b9150610c4c60006103e8612c6b565b610c90565b600d54610c8d90610c628684612bb7565b81610c6957fe5b04600e54610c808487612bb790919063ffffffff16565b81610c8757fe5b04612cf6565b91505b60008211610ccf5760405162461bcd60e51b8152600401808060200182810382526027815260200180614fcf6027913960400191505060405180910390fd5b6001600160a01b038516600090815260036020526040902054610d0d576016546001600160a01b038616600090815260196020526040902055610da1565b6001600160a01b038516600090815260036020526040902054610d309083612d0c565b610d7c610d4860165485612bb790919063ffffffff16565b6001600160a01b038816600090815260196020908152604080832054600390925290912054610d7691612bb7565b90612d0c565b81610d8357fe5b6001600160a01b038716600090815260196020526040902091900490555b610dab8583612c6b565b610dcc600d54610dc685600e54612d0c90919063ffffffff16565b90612bb7565b610de7600e54610dc687600d54612d0c90919063ffffffff16565b1015610e3a576040805162461bcd60e51b815260206004820152601860248201527f64466564506169723a20494e56414c49445f5550444154450000000000000000604482015290519081900360640190fd5b600c546040805163df6c1de160e01b81526004810187905290516001600160a01b039092169163df6c1de19160248082019260009290919082900301818387803b158015610e8757600080fd5b505af1158015610e9b573d6000803e3d6000fd5b5050600d54610ec29250610eb0915086612d0c565b600e54610ebd9086612d0c565b612d5b565b604080518581526020810185905281516001600160a01b038816927f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f928290030190a2506008805460ff191660011790559392505050565b60025481565b60085460ff16610f6a576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff1916905582610fc0576040805162461bcd60e51b81526020600482015260176024820152761911995914185a5c8e881253959053125117d253941555604a1b604482015290519081900360640190fd5b82611035601154610af3600954610af3600d54600a60009054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610ac157600080fd5b10156110725760405162461bcd60e51b8152600401808060200182810382526023815260200180614f856023913960400191505060405180910390fd5b6015546000906127109061108a90869060ff16612bb7565b8161109157fe5b04905060006110a08583612b67565b905060006110b382600d54600e54612d66565b9050848110156110f45760405162461bcd60e51b8152600401808060200182810382526024815260200180614efb6024913960400191505060405180910390fd5b6110fc6129fb565b801561111957600b54611119906001600160a01b03168583612e21565b82156111305760095461112c9084612d0c565b6009555b600d5461114e906111419084612d0c565b600e54610ebd9084612b67565b600c546040805163df6c1de160e01b81526004810185905290516001600160a01b039092169163df6c1de19160248082019260009290919082900301818387803b15801561119b57600080fd5b505af11580156111af573d6000803e3d6000fd5b5050604080518981526020810185905281516001600160a01b03891694503393507ff2fc1627c485a28da79658364df5619455857b1d9d2fcc5787b19b45555665f3929181900390910190a350506008805460ff1916600117905550505050565b6001600160a01b03831660009081526004602090815260408083203384529091528120546000191461128f576001600160a01b038416600090815260046020908152604080832033845290915290205461126a9083612b67565b6001600160a01b03851660009081526004602090815260408083203384529091529020555b61129a848484612f8b565b5060015b9392505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60055481565b60085460ff1661131e576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff1916905560006113338261163c565b905080611383576040805162461bcd60e51b81526020600482015260196024820152781911995910985b9ace88111150925517d393d517d193d55391603a1b604482015290519081900360640190fd5b60006113d46007838154811061139557fe5b906000526020600020906008020160040154600784815481106113b457fe5b906000526020600020906008020160030154612d0c90919063ffffffff16565b90506000611427600784815481106113e857fe5b9060005260206000209060080201600501546007858154811061140757fe5b906000526020600020906008020160020154612d0c90919063ffffffff16565b905060006007848154811061143857fe5b600091825260209182902060089091020154601154600954600d54600a54604080516370a0823160e01b815230600482015290516001600160a01b0396871698508a976114ae97610af39695879590949116926370a0823192602480840193919291829003018186803b158015610ac157600080fd5b10156114eb5760405162461bcd60e51b8152600401808060200182810382526023815260200180614f1f6023913960400191505060405180910390fd5b61151e600785815481106114fb57fe5b906000526020600020906008020160040154601154612d0c90919063ffffffff16565b60115561152a84613039565b600c54604080516303b4078d60e21b81523060048201526024810186905290516001600160a01b0390921691630ed01e349160448082019260009290919082900301818387803b15801561157d57600080fd5b505af1158015611591573d6000803e3d6000fd5b5050600b546115ad92506001600160a01b031690508284612e21565b6040805186815260006020820181905281830181905260608201819052608082015290516001600160a01b03831691600080516020614f42833981519152919081900360a00190a26040805186815290517fa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe462309181900360200190a150506008805460ff19166001179055505050565b600080600760008154811061164d57fe5b90600052602060002090600802016006015490505b80156116bb57826007828154811061167657fe5b90600052602060002090600802016001015414156116955790506116c1565b600781815481106116a257fe5b9060005260206000209060080201600601549050611662565b60009150505b919050565b600d5481565b600c546001600160a01b031615611720576040805162461bcd60e51b8152602060048201526013602482015272322332b22830b4b91d102327a92124a22222a760691b604482015290519081900360640190fd5b600c8054336001600160a01b031991821617909155600a805482166001600160a01b03858116918217909255600b8054909316918416919091179091556040805163313ce56760e01b8152905163313ce56791600480820192602092909190829003018186803b15801561179357600080fd5b505afa1580156117a7573d6000803e3d6000fd5b505050506040513d60208110156117bd57600080fd5b50516018805460ff191660ff9092169190911790556117db816131a9565b80516117ef91600091602090910190614d8a565b506000805461181291600191600260001961010083861615020190911604614e08565b50600c60009054906101000a90046001600160a01b03166001600160a01b031663b7c9252c6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561186357600080fd5b505af1158015611877573d6000803e3d6000fd5b505050506040513d602081101561188d57600080fd5b50516017556040516000805446927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f9291819083906002600019610100600184161502019091160480156119185780601f106118f6576101008083540402835291820191611918565b820191906000526020600020905b815481529060010190602001808311611904575b505060408051918290038220828201825260018352603160f81b602093840152815180840196909652858201527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606086015260808501959095523060a0808601919091528551808603909101815260c0909401909452505080519101206005555050565b60085460009060ff166119ea576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff19169055601254600f54600e54600b54604080516370a0823160e01b815230600482015290518995611a58959094610af3949193859391926001600160a01b0316916370a0823191602480820192602092909190829003018186803b158015610ac157600080fd5b1015611a955760405162461bcd60e51b8152600401808060200182810382526023815260200180614f1f6023913960400191505060405180910390fd5b60008311611aea576040805162461bcd60e51b815260206004820152601f60248201527f6446656442616e6b3a205441524745545f4e554d4245525f494e56414c494400604482015290519081900360640190fd5b600080611af786866127d9565b9150915080611b4d576040805162461bcd60e51b815260206004820152601760248201527f6446656442616e6b3a20494e56414c49445f4445424954000000000000000000604482015290519081900360640190fd5b611b5982878787613335565b600c5460408051637c68803560e11b81526001600160a01b038881166004830152602482018a9052915193965091169163f8d1006a9160448082019260009290919082900301818387803b158015611bb057600080fd5b505af1158015611bc4573d6000803e3d6000fd5b505060408051868152602081018a9052808201899052600060608201819052608082015290516001600160a01b0388169350600080516020614f4283398151915292509081900360a00190a2604080516001600160a01b03861681526020810188905280820187905290517f17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a49181900360600190a150506008805460ff191660011790559392505050565b600e5481565b60135481565b60036020526000908152604090205481565b6001600160a01b038216600090815260146020526040902054611cb09082612b67565b6001600160a01b038316600090815260146020526040812091909155601354601254611cdd908490612bb7565b81611ce457fe5b049050611cfc82601354612b6790919063ffffffff16565b601355601254611d0c9082612b67565b6012558015611d2c57600b54611d2c906001600160a01b03168483612e21565b505050565b60165481565b60066020526000908152604090205481565b60085460ff16611d93576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff1916905582611de9576040805162461bcd60e51b81526020600482015260176024820152761911995914185a5c8e881253959053125117d253941555604a1b604482015290519081900360640190fd5b82611e5e601254610af3600f54610af3600e54600b60009054906101000a90046001600160a01b03166001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610ac157600080fd5b1015611e9b5760405162461bcd60e51b8152600401808060200182810382526023815260200180614f856023913960400191505060405180910390fd5b600d548210611ef1576040805162461bcd60e51b815260206004820181905260248201527f64466564506169723a20494e53554646494349454e545f4c4951554944495459604482015290519081900360640190fd5b600080611efd8561362e565b915091506000611f1283600e54600d54612d66565b90506000611f208284612b67565b60155490915060009061271090611f3b90849060ff16612bb7565b81611f4257fe5b0490506000611f518383612b67565b905087811015611f925760405162461bcd60e51b8152600401808060200182810382526024815260200180614efb6024913960400191505060405180910390fd5b611f9a6129fb565b8015611fb757600a54611fb7906001600160a01b03168883612e21565b8115611fce57600954611fca9083612d0c565b6009555b600d54611fec90611fdf9086612b67565b600e54610ebd9089612d0c565b600c546040805163daaa88bd60e01b81526004810187905290516001600160a01b039092169163daaa88bd9160248082019260009290919082900301818387803b15801561203957600080fd5b505af115801561204d573d6000803e3d6000fd5b5050604080518c81526020810185905281516001600160a01b038c1694503393507f1b38ce2719d9be70e0c768bd2a89d778eaef1cd41e3a4c02fa2317d415c3d348929181900390910190a350506008805460ff1916600117905550505050505050565b600854600090819060ff16612100576040805162461bcd60e51b815260206004820152601060248201526f1911995910985b9ace881313d0d2d15160821b604482015290519081900360640190fd5b6008805460ff191690556121126129fb565b30600090815260036020526040902054600254600d548190612135908490612bb7565b8161213c57fe5b04935080612155600e5484612bb790919063ffffffff16565b8161215c57fe5b04925060008411801561216f5750600083115b6121aa5760405162461bcd60e51b8152600401808060200182810382526027815260200180614fa86027913960400191505060405180910390fd5b6121b430836139bf565b60008060006121e36121d188600d54612b6790919063ffffffff16565b600e546121de9089612b67565b613a50565b6001600160a01b038b16600090815260146020526040902054929550909350915061220e9082612d0c565b6001600160a01b038916600090815260146020526040902055600d54612245906122389089612b67565b600e54610ebd9089612b67565b61224f8684612d0c565b955061225b8783612b67565b600c546040805163daaa88bd60e01b81526004810184905290519299506001600160a01b039091169163daaa88bd9160248082019260009290919082900301818387803b1580156122ab57600080fd5b505af11580156122bf573d6000803e3d6000fd5b5050600c546001600160a01b038b811660009081526019602052604090205460165491909216935063ccd072e592508b91670de0b6b3a76400009161230f916123089190612b67565b8a90612bb7565b8161231657fe5b046040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561235d57600080fd5b505af1158015612371573d6000803e3d6000fd5b50505050600086111561239557600b54612395906001600160a01b03168988612e21565b86156123b257600a546123b2906001600160a01b03168989612e21565b604080518681526020810189905280820188905290516001600160a01b038a169133917fd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c8833639181900360600190a350506008805460ff1916600117905550929491935090915050565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108dc5780601f106108b1576101008083540402835291602001916108dc565b60115481565b60175481565b60155460ff1681565b6000610901338484612f8b565b60146020526000908152604090205481565b6103e881565b600f5481565b600c546001600160a01b031681565b60105481565b600b546001600160a01b031681565b42841015612522576040805162461bcd60e51b815260206004820152601260248201527119119959115490cc8c0e881156141254915160721b604482015290519081900360640190fd5b6005546001600160a01b0380891660008181526006602090815260408083208054600180820190925582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98186015280840196909652958d166060860152608085018c905260a085019590955260c08085018b90528151808603909101815260e08501825280519083012061190160f01b6101008601526101028501969096526101228085019690965280518085039096018652610142840180825286519683019690962095839052610162840180825286905260ff89166101828501526101a284018890526101c28401879052519193926101e280820193601f1981019281900390910190855afa15801561263d573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116158015906126735750886001600160a01b0316816001600160a01b0316145b6126c4576040805162461bcd60e51b815260206004820152601c60248201527f6446656445524332303a20494e56414c49445f5349474e415455524500000000604482015290519081900360640190fd5b6126cf898989612999565b505050505050505050565b60196020526000908152604090205481565b600460209081526000928352604080842090915290825290205481565b60095481565b6007818154811061271c57fe5b6000918252602090912060089091020180546001820154600283015460038401546004850154600586015460068701546007909701546001600160a01b0390961697509395929491939092909188565b6001600160a01b03821660009081526014602052604090205461278f9082612b67565b6001600160a01b0383166000908152601460205260409020556011546127b59082612b67565b60115580156127d557600a546127d5906001600160a01b03168383612e21565b5050565b60008060006127ee8585600d54600e54614716565b9050600e54811015612807576000809250925050612932565b6000805b6007858154811061281857fe5b90600052602060002090600802016006015460001461292a576007858154811061283e57fe5b906000526020600020906008020160060154915061289e6007838154811061286257fe5b9060005260206000209060080201600201546007848154811061288157fe5b906000526020600020906008020160030154600d54600e54614716565b9050806128ab8489612d0c565b116128bd575060019250612932915050565b6128ee600783815481106128cd57fe5b90600052602060002090600802016002015482612d0c90919063ffffffff16565b8310156129045760008094509450505050612932565b6007858154811061291157fe5b906000526020600020906008020160060154945061280b565b506001925050505b9250929050565b600081612989576040805162461bcd60e51b81526020600482015260196024820152781911995910985b9ace88111150925517d393d517d193d55391603a1b604482015290519081900360640190fd5b6109056007838154811061139557fe5b6001600160a01b03808416600081815260046020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b600c60009054906101000a90046001600160a01b03166001600160a01b031663d7ed714e6040518163ffffffff1660e01b815260040160206040518083038186803b158015612a4957600080fd5b505afa158015612a5d573d6000803e3d6000fd5b505050506040513d6020811015612a7357600080fd5b50514311612a8057612b65565b600c5460408051632df2494b60e21b815290516000926001600160a01b03169163b7c9252c91600480830192602092919082900301818787803b158015612ac657600080fd5b505af1158015612ada573d6000803e3d6000fd5b505050506040513d6020811015612af057600080fd5b505160175490915081118015612b0857506000600254115b15612b6357600254601854600d54601754612b5a939260ff16600a0a91612b4091670de0b6b3a764000091610dc69182908990612b67565b81612b4757fe5b0481612b4f57fe5b601654919004612d0c565b60165560178190555b505b565b80820382811115610905576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6000811580612bd257505080820282828281612bcf57fe5b04145b610905576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b60006003821115612c5d575080600160028204015b81811015612c5757809150600281828581612c4657fe5b040181612c4f57fe5b049050612c2f565b506116c1565b81156116c157506001919050565b600254612c789082612d0c565b6002556001600160a01b038216600090815260036020526040902054612c9e9082612d0c565b6001600160a01b03831660008181526003602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6000818310612d05578161129e565b5090919050565b80820182811015610905576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fd5b600d91909155600e55565b6000808411612da65760405162461bcd60e51b8152600401808060200182810382526026815260200180614ff66026913960400191505060405180910390fd5b600083118015612db65750600082115b612df15760405162461bcd60e51b8152600401808060200182810382526023815260200180614f626023913960400191505060405180910390fd5b6000612dfd8584612bb7565b90506000612e0b8587612d0c565b9050808281612e1657fe5b049695505050505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b178152925182516000946060949389169392918291908083835b60208310612e9e5780518252601f199092019160209182019101612e7f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612f00576040519150601f19603f3d011682016040523d82523d6000602084013e612f05565b606091505b5091509150818015612f33575080511580612f335750808060200190516020811015612f3057600080fd5b50515b612f84576040805162461bcd60e51b815260206004820152601f60248201527f5472616e7366657248656c7065723a205452414e534645525f4641494c454400604482015290519081900360640190fd5b5050505050565b6001600160a01b038316600090815260036020526040902054612fae9082612b67565b6001600160a01b038085166000908152600360205260408082209390935590841681522054612fdd9082612d0c565b6001600160a01b0380841660008181526003602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6007818154811061304657fe5b906000526020600020906008020160060154600780838154811061306657fe5b9060005260206000209060080201600701548154811061308257fe5b906000526020600020906008020160060181905550600781815481106130a457fe5b90600052602060002090600802016007015460078083815481106130c457fe5b906000526020600020906008020160060154815481106130e057fe5b90600052602060002090600802016007018190555061314a6007828154811061310557fe5b906000526020600020906008020160050154610af36007848154811061312757fe5b906000526020600020906008020160020154600f54612b6790919063ffffffff16565b600f55600780548290811061315b57fe5b60009182526020822060089091020180546001600160a01b03191681556001810182905560028101829055600381018290556004810182905560058101829055600681018290556007015550565b60606109056132dc836001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156131ea57600080fd5b505afa1580156131fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561322757600080fd5b810190808051604051939291908464010000000082111561324757600080fd5b90830190602082018581111561325c57600080fd5b825164010000000081118282018810171561327657600080fd5b82525081516020918201929091019080838360005b838110156132a357818101518382015260200161328b565b50505050905090810190601f1680156132d05780820380516001836020036101000a031916815260200191505b5060405250505061478e565b61332a613330613304604051806040016040528060018152602001602d60f81b81525061478e565b61332a6040518060400160405280600581526020016404645444c560dc1b81525061478e565b906147b3565b61478e565b60008061334061483a565b601080546001810190915592509050613357614e7d565b604051806101000160405280856001600160a01b0316815260200184815260200187815260200186815260200160008152602001600081526020016007898154811061339f57fe5b90600052602060002090600802016006015481526020018881525090508160078260c00151815481106133ce57fe5b90600052602060002090600802016007018190555081600788815481106133f157fe5b600091825260209091206006600890920201015560075482141561358c576007805460018101825560009190915281517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688600890920291820180546001600160a01b0319166001600160a01b0390921691909117905560208201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68982015560408201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68a82015560608201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68b82015560808201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68c82015560a08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68d82015560c08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68e82015560e08201517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68f90910155613613565b806007838154811061359a57fe5b600091825260209182902083516008929092020180546001600160a01b0319166001600160a01b0390921691909117815590820151600182015560408201516002820155606082015160038201556080820151600482015560a0820151600582015560c0820151600682015560e0909101516007909101555b600f546136209087612d0c565b600f55509095945050505050565b6000808291506000600760008154811061364457fe5b90600052602060002090600802016006015490505b80156139b9576007818154811061366c57fe5b90600052602060002090600802016002015460001480156136ab57506007818154811061369557fe5b9060005260206000209060080201600301546000145b1561384957600e546136bd9084612d0c565b613709600783815481106136cd57fe5b906000526020600020906008020160050154600784815481106136ec57fe5b906000526020600020906008020160040154600d54600e54614716565b10156138445760006007828154811061371e57fe5b90600052602060002090600802016006015490506137656007838154811061374257fe5b906000526020600020906008020160050154601254612d0c90919063ffffffff16565b60128190555061379e6007838154811061377b57fe5b906000526020600020906008020160040154601354612d0c90919063ffffffff16565b60135560078054839081106137af57fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f428339815191529190859081106137eb57fe5b906000526020600020906008020160010154600080600080604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a261383d82613039565b9050613659565b6139b9565b600e546138569084612d0c565b6138666007838154811061286257fe5b10156138445760006007828154811061387b57fe5b90600052602060002090600802016006015490506138c06007838154811061389f57fe5b90600052602060002090600802016002015485612d0c90919063ffffffff16565b93506138f3600783815481106138d257fe5b90600052602060002090600802016003015484612d0c90919063ffffffff16565b92506007828154811061390257fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908590811061393e57fe5b906000526020600020906008020160010154600080600080604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a26139976007838154811061374257fe5b6012819055506139ad6007838154811061377b57fe5b60135561383d82613039565b50915091565b6001600160a01b0382166000908152600360205260409020546139e29082612b67565b6001600160a01b038316600090815260036020526040902055600254613a089082612b67565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600080600080600019905060008060008060006007600081548110613a7157fe5b90600052602060002090600802016007015490505b80156147095760078181548110613a9957fe5b9060005260206000209060080201600201546000148015613ad8575060078181548110613ac257fe5b9060005260206000209060080201600301546000145b15613b035760078181548110613aea57fe5b9060005260206000209060080201600701549050614704565b613b4b60078281548110613b1357fe5b90600052602060002090600802016002015460078381548110613b3257fe5b9060005260206000209060080201600301548d8d614716565b94508985101561418c5789613b8760078381548110613b6657fe5b90600052602060002090600802016002015487612d0c90919063ffffffff16565b11613e3557613bbd60078281548110613b9c57fe5b9060005260206000209060080201600201548a612d0c90919063ffffffff16565b9850613bf060078281548110613bcf57fe5b90600052602060002090600802016003015489612d0c90919063ffffffff16565b975060078181548110613bff57fe5b906000526020600020906008020160070154915060078181548110613c2057fe5b9060005260206000209060080201600401546000148015613c5f575060078181548110613c4957fe5b9060005260206000209060080201600501546000145b15613d045760078181548110613c7157fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f42833981519152919084908110613cad57fe5b906000526020600020906008020160010154600080600080604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a2613cff81613039565b613e2e565b613d146007828154811061312757fe5b600f556007805482908110613d2557fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f42833981519152919084908110613d6157fe5b90600052602060002090600802016001015460008060078681548110613d8357fe5b90600052602060002090600802016004015460078781548110613da257fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a2600060078281548110613df457fe5b906000526020600020906008020160020181905550600060078281548110613e1857fe5b9060005260206000209060080201600301819055505b5080614704565b613e3f8a86612b67565b925084613e4c8c85612bb7565b81613e5357fe5b04935060078181548110613e6357fe5b906000526020600020906008020160030154841180613e9f575060078181548110613e8a57fe5b90600052602060002090600802016002015483115b15613ee75760078181548110613eb157fe5b906000526020600020906008020160030154935060078181548110613ed257fe5b90600052602060002090600802016002015492505b613ef18984612d0c565b9850613efd8885612d0c565b9750613f308360078381548110613f1057fe5b906000526020600020906008020160020154612b6790919063ffffffff16565b60078281548110613f3d57fe5b906000526020600020906008020160020181905550613f838460078381548110613f6357fe5b906000526020600020906008020160030154612b6790919063ffffffff16565b60078281548110613f9057fe5b6000918252602090912060036008909202010155600f54613fb19084612b67565b600f55600780548b965082908110613fc557fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908490811061400157fe5b9060005260206000209060080201600101546007848154811061402057fe5b9060005260206000209060080201600201546007858154811061403f57fe5b9060005260206000209060080201600301546007868154811061405e57fe5b9060005260206000209060080201600401546007878154811061407d57fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a2600781815481106140cd57fe5b9060005260206000209060080201600701549150600781815481106140ee57fe5b906000526020600020906008020160030154600014801561412d57506007818154811061411757fe5b9060005260206000209060080201600201546000145b1561418c576007818154811061413f57fe5b906000526020600020906008020160040154600014801561417e57506007818154811061416857fe5b9060005260206000209060080201600501546000145b15613e2e57613e2e81613039565b8585106143e1576141a360078281548110613bcf57fe5b97506141d6600782815481106141b557fe5b90600052602060002090600802016003015488612d0c90919063ffffffff16565b9650614227600782815481106141e857fe5b9060005260206000209060080201600301546007838154811061420757fe5b906000526020600020906008020160040154612d0c90919063ffffffff16565b6007828154811061423457fe5b9060005260206000209060080201600401819055506142986007828154811061425957fe5b9060005260206000209060080201600201546007838154811061427857fe5b906000526020600020906008020160050154612d0c90919063ffffffff16565b600782815481106142a557fe5b906000526020600020906008020160050181905550600781815481106142c757fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908490811061430357fe5b9060005260206000209060080201600101546000806007868154811061432557fe5b9060005260206000209060080201600401546007878154811061434457fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a260006007828154811061439657fe5b9060005260206000209060080201600201819055506000600782815481106143ba57fe5b9060005260206000209060080201600301819055506143da818c8c61487e565b9050614704565b856143f260078381548110613b6657fe5b11156146df5761442d86610af36007848154811061440c57fe5b90600052602060002090600802016002015488612d0c90919063ffffffff16565b925061444961444260078381548110613b6657fe5b8790612bb7565b61445784610dc68e8e612bb7565b8161445e57fe5b0493506007818154811061446e57fe5b9060005260206000209060080201600301548411806144aa57506007818154811061449557fe5b90600052602060002090600802016002015483115b156144f257600781815481106144bc57fe5b9060005260206000209060080201600301549350600781815481106144dd57fe5b90600052602060002090600802016002015492505b6144fc8885612d0c565b97506145088785612d0c565b965061451b846007838154811061420757fe5b6007828154811061452857fe5b90600052602060002090600802016004018190555061454e836007838154811061427857fe5b6007828154811061455b57fe5b9060005260206000209060080201600501819055506145818360078381548110613f1057fe5b6007828154811061458e57fe5b9060005260206000209060080201600201819055506145b48460078381548110613f6357fe5b600782815481106145c157fe5b906000526020600020906008020160030181905550600781815481106145e357fe5b6000918252602090912060089091020154600780546001600160a01b0390921691600080516020614f4283398151915291908490811061461f57fe5b9060005260206000209060080201600101546007848154811061463e57fe5b9060005260206000209060080201600201546007858154811061465d57fe5b9060005260206000209060080201600301546007868154811061467c57fe5b9060005260206000209060080201600401546007878154811061469b57fe5b906000526020600020906008020160050154604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a25b849550600781815481106146ef57fe5b90600052602060002090600802016007015490505b613a86565b5050505050509250925092565b6000806147238484612bb7565b905085810260008288838161473457fe5b041461474c57614745838989614c68565b9050614759565b86828161475557fe5b0490505b61478260028904610af3610c3860046147728d80612bb7565b8161477957fe5b86919004612d0c565b98975050505050505050565b614796614ecb565b506040805180820190915281518152602082810190820152919050565b8051825160609182910167ffffffffffffffff811180156147d357600080fd5b506040519080825280601f01601f1916602001820160405280156147fe576020820181803683370190505b509050600060208201905061481c8186602001518760000151614d1f565b8451602085015185516148329284019190614d1f565b509392505050565b60015b60075481101561487b576007818154811061485457fe5b906000526020600020906008020160010154600014156148735761487b565b60010161483d565b90565b6000806007858154811061488e57fe5b9060005260206000209060080201600701549050600785815481106148af57fe5b90600052602060002090600802016006015460078087815481106148cf57fe5b906000526020600020906008020160070154815481106148eb57fe5b90600052602060002090600802016006018190555080600780878154811061490f57fe5b9060005260206000209060080201600601548154811061492b57fe5b906000526020600020906008020160070181905550600061498a6007878154811061495257fe5b9060005260206000209060080201600501546007888154811061497157fe5b9060005260206000209060080201600401548787614716565b9050600080600760008154811061499d57fe5b90600052602060002090600802016006015490505b8015614bcf57600781815481106149c557fe5b9060005260206000209060080201600201546000148015614a045750600781815481106149ee57fe5b9060005260206000209060080201600301546000145b15614a5857614a5160078281548110614a1957fe5b90600052602060002090600802016005015460078381548110614a3857fe5b9060005260206000209060080201600401548989614716565b9150614aa3565b614aa060078281548110614a6857fe5b90600052602060002090600802016002015460078381548110614a8757fe5b9060005260206000209060080201600301548989614716565b91505b818311614b7d5760078181548110614ab757fe5b90600052602060002090600802016007015460078981548110614ad657fe5b906000526020600020906008020160070181905550876007808381548110614afa57fe5b90600052602060002090600802016007015481548110614b1657fe5b9060005260206000209060080201600601819055508060078981548110614b3957fe5b9060005260206000209060080201600601819055508760078281548110614b5c57fe5b9060005260206000209060080201600701819055508394505050505061129e565b60078181548110614b8a57fe5b90600052602060002090600802016006015460001415614ba957614bcf565b60078181548110614bb657fe5b90600052602060002090600802016006015490506149b2565b8060078981548110614bdd57fe5b9060005260206000209060080201600701819055508760078281548110614c0057fe5b906000526020600020906008020160060181905550600060078981548110614c2457fe5b906000526020600020906008020160060181905550876007600081548110614c4857fe5b600091825260209091206008909102016007015550919695505050505050565b6000806000614c778686614d5d565b91509150838110614c8757600080fd5b60008480614c9157fe5b868809905082811115614ca5576001820391505b918290039160008590038516808681614cba57fe5b049550808481614cc657fe5b049350808160000381614cd557fe5b046001019290920292909201600285810380870282030280870282030280870282030280870282030280870282030280870282030295860290039094029390930295945050505050565b5b60208110614d3f578151835260209283019290910190601f1901614d20565b905182516020929092036101000a6000190180199091169116179052565b6000808060001984860990508385029250828103915082811015614d82576001820391505b509250929050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614dcb57805160ff1916838001178555614df8565b82800160010185558215614df8579182015b82811115614df8578251825591602001919060010190614ddd565b50614e04929150614ee5565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614e415780548555614df8565b82800160010185558215614df857600052602060002091601f016020900482015b82811115614df8578254825591600101919060010190614e62565b60405180610100016040528060006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604051806040016040528060008152602001600081525090565b5b80821115614e045760008155600101614ee656fe64466564506169723a20494e53554646494349454e545f4f55545055545f414d4f554e546446656442616e6b3a20494e53554646494349454e545f494e5055545f414d4f554e54a0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869644665644c6962726172793a20494e53554646494349454e545f4c495155494449545964466564506169723a20494e53554646494349454e545f494e5055545f414d4f554e5464466564506169723a20494e53554646494349454e545f4c49515549444954595f4255524e454464466564506169723a20494e53554646494349454e545f4c49515549444954595f4d494e544544644665644c6962726172793a20494e53554646494349454e545f494e5055545f414d4f554e54a2646970667358221220af73e875362f7f20e3c9b3e8303ddea008ac026d22568c63cba01f015a68f3d764736f6c63430007000033"

// DeployDFedPair deploys a new Ethereum contract, binding an instance of DFedPair to it.
func DeployDFedPair(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DFedPair, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedPairABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DFedPairBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DFedPair{DFedPairCaller: DFedPairCaller{contract: contract}, DFedPairTransactor: DFedPairTransactor{contract: contract}, DFedPairFilterer: DFedPairFilterer{contract: contract}}, nil
}

// DFedPair is an auto generated Go binding around an Ethereum contract.
type DFedPair struct {
	DFedPairCaller     // Read-only binding to the contract
	DFedPairTransactor // Write-only binding to the contract
	DFedPairFilterer   // Log filterer for contract events
}

// DFedPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type DFedPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DFedPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DFedPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DFedPairSession struct {
	Contract     *DFedPair         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DFedPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DFedPairCallerSession struct {
	Contract *DFedPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DFedPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DFedPairTransactorSession struct {
	Contract     *DFedPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DFedPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type DFedPairRaw struct {
	Contract *DFedPair // Generic contract binding to access the raw methods on
}

// DFedPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DFedPairCallerRaw struct {
	Contract *DFedPairCaller // Generic read-only contract binding to access the raw methods on
}

// DFedPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DFedPairTransactorRaw struct {
	Contract *DFedPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDFedPair creates a new instance of DFedPair, bound to a specific deployed contract.
func NewDFedPair(address common.Address, backend bind.ContractBackend) (*DFedPair, error) {
	contract, err := bindDFedPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DFedPair{DFedPairCaller: DFedPairCaller{contract: contract}, DFedPairTransactor: DFedPairTransactor{contract: contract}, DFedPairFilterer: DFedPairFilterer{contract: contract}}, nil
}

// NewDFedPairCaller creates a new read-only instance of DFedPair, bound to a specific deployed contract.
func NewDFedPairCaller(address common.Address, caller bind.ContractCaller) (*DFedPairCaller, error) {
	contract, err := bindDFedPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DFedPairCaller{contract: contract}, nil
}

// NewDFedPairTransactor creates a new write-only instance of DFedPair, bound to a specific deployed contract.
func NewDFedPairTransactor(address common.Address, transactor bind.ContractTransactor) (*DFedPairTransactor, error) {
	contract, err := bindDFedPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DFedPairTransactor{contract: contract}, nil
}

// NewDFedPairFilterer creates a new log filterer instance of DFedPair, bound to a specific deployed contract.
func NewDFedPairFilterer(address common.Address, filterer bind.ContractFilterer) (*DFedPairFilterer, error) {
	contract, err := bindDFedPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DFedPairFilterer{contract: contract}, nil
}

// bindDFedPair binds a generic wrapper to an already deployed contract.
func bindDFedPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedPair *DFedPairRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedPair.Contract.DFedPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedPair *DFedPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedPair.Contract.DFedPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedPair *DFedPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedPair.Contract.DFedPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedPair *DFedPairCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedPair *DFedPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedPair *DFedPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedPair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedPair *DFedPairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedPair *DFedPairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DFedPair.Contract.DOMAINSEPARATOR(&_DFedPair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DFedPair *DFedPairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DFedPair.Contract.DOMAINSEPARATOR(&_DFedPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_DFedPair *DFedPairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "MINIMUM_LIQUIDITY")
	return *ret0, err
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_DFedPair *DFedPairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _DFedPair.Contract.MINIMUMLIQUIDITY(&_DFedPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _DFedPair.Contract.MINIMUMLIQUIDITY(&_DFedPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedPair *DFedPairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedPair *DFedPairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _DFedPair.Contract.PERMITTYPEHASH(&_DFedPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_DFedPair *DFedPairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _DFedPair.Contract.PERMITTYPEHASH(&_DFedPair.CallOpts)
}

// AccRewardPerFEDlpOfPair is a free data retrieval call binding the contract method 0x7dea576e.
//
// Solidity: function accRewardPerFEDlpOfPair() view returns(uint256)
func (_DFedPair *DFedPairCaller) AccRewardPerFEDlpOfPair(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "accRewardPerFEDlpOfPair")
	return *ret0, err
}

// AccRewardPerFEDlpOfPair is a free data retrieval call binding the contract method 0x7dea576e.
//
// Solidity: function accRewardPerFEDlpOfPair() view returns(uint256)
func (_DFedPair *DFedPairSession) AccRewardPerFEDlpOfPair() (*big.Int, error) {
	return _DFedPair.Contract.AccRewardPerFEDlpOfPair(&_DFedPair.CallOpts)
}

// AccRewardPerFEDlpOfPair is a free data retrieval call binding the contract method 0x7dea576e.
//
// Solidity: function accRewardPerFEDlpOfPair() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) AccRewardPerFEDlpOfPair() (*big.Int, error) {
	return _DFedPair.Contract.AccRewardPerFEDlpOfPair(&_DFedPair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedPair *DFedPairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedPair *DFedPairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.Allowance(&_DFedPair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DFedPair *DFedPairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.Allowance(&_DFedPair.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedPair *DFedPairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedPair *DFedPairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.BalanceOf(&_DFedPair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DFedPair *DFedPairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.BalanceOf(&_DFedPair.CallOpts, arg0)
}

// DebtInfos is a free data retrieval call binding the contract method 0xe3c20fd5.
//
// Solidity: function debtInfos(uint256 ) view returns(address user, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount, uint256 nextIndex, uint256 lastIndex)
func (_DFedPair *DFedPairCaller) DebtInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User             common.Address
	DebtId           *big.Int
	PledgeAmount     *big.Int
	RepayAmount      *big.Int
	DebtToken0Amount *big.Int
	DebtToken1Amount *big.Int
	NextIndex        *big.Int
	LastIndex        *big.Int
}, error) {
	ret := new(struct {
		User             common.Address
		DebtId           *big.Int
		PledgeAmount     *big.Int
		RepayAmount      *big.Int
		DebtToken0Amount *big.Int
		DebtToken1Amount *big.Int
		NextIndex        *big.Int
		LastIndex        *big.Int
	})
	out := ret
	err := _DFedPair.contract.Call(opts, out, "debtInfos", arg0)
	return *ret, err
}

// DebtInfos is a free data retrieval call binding the contract method 0xe3c20fd5.
//
// Solidity: function debtInfos(uint256 ) view returns(address user, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount, uint256 nextIndex, uint256 lastIndex)
func (_DFedPair *DFedPairSession) DebtInfos(arg0 *big.Int) (struct {
	User             common.Address
	DebtId           *big.Int
	PledgeAmount     *big.Int
	RepayAmount      *big.Int
	DebtToken0Amount *big.Int
	DebtToken1Amount *big.Int
	NextIndex        *big.Int
	LastIndex        *big.Int
}, error) {
	return _DFedPair.Contract.DebtInfos(&_DFedPair.CallOpts, arg0)
}

// DebtInfos is a free data retrieval call binding the contract method 0xe3c20fd5.
//
// Solidity: function debtInfos(uint256 ) view returns(address user, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount, uint256 nextIndex, uint256 lastIndex)
func (_DFedPair *DFedPairCallerSession) DebtInfos(arg0 *big.Int) (struct {
	User             common.Address
	DebtId           *big.Int
	PledgeAmount     *big.Int
	RepayAmount      *big.Int
	DebtToken0Amount *big.Int
	DebtToken1Amount *big.Int
	NextIndex        *big.Int
	LastIndex        *big.Int
}, error) {
	return _DFedPair.Contract.DebtInfos(&_DFedPair.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedPair *DFedPairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedPair *DFedPairSession) Decimals() (uint8, error) {
	return _DFedPair.Contract.Decimals(&_DFedPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DFedPair *DFedPairCallerSession) Decimals() (uint8, error) {
	return _DFedPair.Contract.Decimals(&_DFedPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedPair *DFedPairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedPair *DFedPairSession) Factory() (common.Address, error) {
	return _DFedPair.Contract.Factory(&_DFedPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedPair *DFedPairCallerSession) Factory() (common.Address, error) {
	return _DFedPair.Contract.Factory(&_DFedPair.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DFedPair *DFedPairCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DFedPair *DFedPairSession) Fee() (*big.Int, error) {
	return _DFedPair.Contract.Fee(&_DFedPair.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) Fee() (*big.Int, error) {
	return _DFedPair.Contract.Fee(&_DFedPair.CallOpts)
}

// FeeRatePer10K is a free data retrieval call binding the contract method 0xa48685fe.
//
// Solidity: function feeRatePer10K() view returns(uint8)
func (_DFedPair *DFedPairCaller) FeeRatePer10K(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "feeRatePer10K")
	return *ret0, err
}

// FeeRatePer10K is a free data retrieval call binding the contract method 0xa48685fe.
//
// Solidity: function feeRatePer10K() view returns(uint8)
func (_DFedPair *DFedPairSession) FeeRatePer10K() (uint8, error) {
	return _DFedPair.Contract.FeeRatePer10K(&_DFedPair.CallOpts)
}

// FeeRatePer10K is a free data retrieval call binding the contract method 0xa48685fe.
//
// Solidity: function feeRatePer10K() view returns(uint8)
func (_DFedPair *DFedPairCallerSession) FeeRatePer10K() (uint8, error) {
	return _DFedPair.Contract.FeeRatePer10K(&_DFedPair.CallOpts)
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 _debtId) view returns(uint256 _index)
func (_DFedPair *DFedPairCaller) GetDebtIndexById(opts *bind.CallOpts, _debtId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "getDebtIndexById", _debtId)
	return *ret0, err
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 _debtId) view returns(uint256 _index)
func (_DFedPair *DFedPairSession) GetDebtIndexById(_debtId *big.Int) (*big.Int, error) {
	return _DFedPair.Contract.GetDebtIndexById(&_DFedPair.CallOpts, _debtId)
}

// GetDebtIndexById is a free data retrieval call binding the contract method 0x3c33a245.
//
// Solidity: function getDebtIndexById(uint256 _debtId) view returns(uint256 _index)
func (_DFedPair *DFedPairCallerSession) GetDebtIndexById(_debtId *big.Int) (*big.Int, error) {
	return _DFedPair.Contract.GetDebtIndexById(&_DFedPair.CallOpts, _debtId)
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 _index) view returns(uint256 _repayAmount)
func (_DFedPair *DFedPairCaller) GetRepayByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "getRepayByIndex", _index)
	return *ret0, err
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 _index) view returns(uint256 _repayAmount)
func (_DFedPair *DFedPairSession) GetRepayByIndex(_index *big.Int) (*big.Int, error) {
	return _DFedPair.Contract.GetRepayByIndex(&_DFedPair.CallOpts, _index)
}

// GetRepayByIndex is a free data retrieval call binding the contract method 0xea52a0be.
//
// Solidity: function getRepayByIndex(uint256 _index) view returns(uint256 _repayAmount)
func (_DFedPair *DFedPairCallerSession) GetRepayByIndex(_index *big.Int) (*big.Int, error) {
	return _DFedPair.Contract.GetRepayByIndex(&_DFedPair.CallOpts, _index)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_DFedPair *DFedPairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	ret := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	out := ret
	err := _DFedPair.contract.Call(opts, out, "getReserves")
	return *ret, err
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_DFedPair *DFedPairSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _DFedPair.Contract.GetReserves(&_DFedPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_DFedPair *DFedPairCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _DFedPair.Contract.GetReserves(&_DFedPair.CallOpts)
}

// IsValidDebit is a free data retrieval call binding the contract method 0xe991685f.
//
// Solidity: function isValidDebit(uint256 _pledgeAmount, uint256 _targetNum) view returns(uint256 _index, bool _valid)
func (_DFedPair *DFedPairCaller) IsValidDebit(opts *bind.CallOpts, _pledgeAmount *big.Int, _targetNum *big.Int) (struct {
	Index *big.Int
	Valid bool
}, error) {
	ret := new(struct {
		Index *big.Int
		Valid bool
	})
	out := ret
	err := _DFedPair.contract.Call(opts, out, "isValidDebit", _pledgeAmount, _targetNum)
	return *ret, err
}

// IsValidDebit is a free data retrieval call binding the contract method 0xe991685f.
//
// Solidity: function isValidDebit(uint256 _pledgeAmount, uint256 _targetNum) view returns(uint256 _index, bool _valid)
func (_DFedPair *DFedPairSession) IsValidDebit(_pledgeAmount *big.Int, _targetNum *big.Int) (struct {
	Index *big.Int
	Valid bool
}, error) {
	return _DFedPair.Contract.IsValidDebit(&_DFedPair.CallOpts, _pledgeAmount, _targetNum)
}

// IsValidDebit is a free data retrieval call binding the contract method 0xe991685f.
//
// Solidity: function isValidDebit(uint256 _pledgeAmount, uint256 _targetNum) view returns(uint256 _index, bool _valid)
func (_DFedPair *DFedPairCallerSession) IsValidDebit(_pledgeAmount *big.Int, _targetNum *big.Int) (struct {
	Index *big.Int
	Valid bool
}, error) {
	return _DFedPair.Contract.IsValidDebit(&_DFedPair.CallOpts, _pledgeAmount, _targetNum)
}

// LastReceivedRewardPerUSDDOfPair is a free data retrieval call binding the contract method 0xa1f04001.
//
// Solidity: function lastReceivedRewardPerUSDDOfPair() view returns(uint256)
func (_DFedPair *DFedPairCaller) LastReceivedRewardPerUSDDOfPair(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "lastReceivedRewardPerUSDDOfPair")
	return *ret0, err
}

// LastReceivedRewardPerUSDDOfPair is a free data retrieval call binding the contract method 0xa1f04001.
//
// Solidity: function lastReceivedRewardPerUSDDOfPair() view returns(uint256)
func (_DFedPair *DFedPairSession) LastReceivedRewardPerUSDDOfPair() (*big.Int, error) {
	return _DFedPair.Contract.LastReceivedRewardPerUSDDOfPair(&_DFedPair.CallOpts)
}

// LastReceivedRewardPerUSDDOfPair is a free data retrieval call binding the contract method 0xa1f04001.
//
// Solidity: function lastReceivedRewardPerUSDDOfPair() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) LastReceivedRewardPerUSDDOfPair() (*big.Int, error) {
	return _DFedPair.Contract.LastReceivedRewardPerUSDDOfPair(&_DFedPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedPair *DFedPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedPair *DFedPairSession) Name() (string, error) {
	return _DFedPair.Contract.Name(&_DFedPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DFedPair *DFedPairCallerSession) Name() (string, error) {
	return _DFedPair.Contract.Name(&_DFedPair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedPair *DFedPairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedPair *DFedPairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.Nonces(&_DFedPair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DFedPair *DFedPairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.Nonces(&_DFedPair.CallOpts, arg0)
}

// RefundAmount is a free data retrieval call binding the contract method 0xacde5d28.
//
// Solidity: function refundAmount(address ) view returns(uint256)
func (_DFedPair *DFedPairCaller) RefundAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "refundAmount", arg0)
	return *ret0, err
}

// RefundAmount is a free data retrieval call binding the contract method 0xacde5d28.
//
// Solidity: function refundAmount(address ) view returns(uint256)
func (_DFedPair *DFedPairSession) RefundAmount(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.RefundAmount(&_DFedPair.CallOpts, arg0)
}

// RefundAmount is a free data retrieval call binding the contract method 0xacde5d28.
//
// Solidity: function refundAmount(address ) view returns(uint256)
func (_DFedPair *DFedPairCallerSession) RefundAmount(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.RefundAmount(&_DFedPair.CallOpts, arg0)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_DFedPair *DFedPairCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "reserve0")
	return *ret0, err
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_DFedPair *DFedPairSession) Reserve0() (*big.Int, error) {
	return _DFedPair.Contract.Reserve0(&_DFedPair.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) Reserve0() (*big.Int, error) {
	return _DFedPair.Contract.Reserve0(&_DFedPair.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_DFedPair *DFedPairCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "reserve1")
	return *ret0, err
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_DFedPair *DFedPairSession) Reserve1() (*big.Int, error) {
	return _DFedPair.Contract.Reserve1(&_DFedPair.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) Reserve1() (*big.Int, error) {
	return _DFedPair.Contract.Reserve1(&_DFedPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedPair *DFedPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedPair *DFedPairSession) Symbol() (string, error) {
	return _DFedPair.Contract.Symbol(&_DFedPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DFedPair *DFedPairCallerSession) Symbol() (string, error) {
	return _DFedPair.Contract.Symbol(&_DFedPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_DFedPair *DFedPairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "token0")
	return *ret0, err
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_DFedPair *DFedPairSession) Token0() (common.Address, error) {
	return _DFedPair.Contract.Token0(&_DFedPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_DFedPair *DFedPairCallerSession) Token0() (common.Address, error) {
	return _DFedPair.Contract.Token0(&_DFedPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_DFedPair *DFedPairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "token1")
	return *ret0, err
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_DFedPair *DFedPairSession) Token1() (common.Address, error) {
	return _DFedPair.Contract.Token1(&_DFedPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_DFedPair *DFedPairCallerSession) Token1() (common.Address, error) {
	return _DFedPair.Contract.Token1(&_DFedPair.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_DFedPair *DFedPairCaller) TotalPledge(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "totalPledge")
	return *ret0, err
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_DFedPair *DFedPairSession) TotalPledge() (*big.Int, error) {
	return _DFedPair.Contract.TotalPledge(&_DFedPair.CallOpts)
}

// TotalPledge is a free data retrieval call binding the contract method 0xc21a43e4.
//
// Solidity: function totalPledge() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) TotalPledge() (*big.Int, error) {
	return _DFedPair.Contract.TotalPledge(&_DFedPair.CallOpts)
}

// TotalRefundToken0Amount is a free data retrieval call binding the contract method 0xa058274f.
//
// Solidity: function totalRefundToken0Amount() view returns(uint256)
func (_DFedPair *DFedPairCaller) TotalRefundToken0Amount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "totalRefundToken0Amount")
	return *ret0, err
}

// TotalRefundToken0Amount is a free data retrieval call binding the contract method 0xa058274f.
//
// Solidity: function totalRefundToken0Amount() view returns(uint256)
func (_DFedPair *DFedPairSession) TotalRefundToken0Amount() (*big.Int, error) {
	return _DFedPair.Contract.TotalRefundToken0Amount(&_DFedPair.CallOpts)
}

// TotalRefundToken0Amount is a free data retrieval call binding the contract method 0xa058274f.
//
// Solidity: function totalRefundToken0Amount() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) TotalRefundToken0Amount() (*big.Int, error) {
	return _DFedPair.Contract.TotalRefundToken0Amount(&_DFedPair.CallOpts)
}

// TotalRefundToken1Amount is a free data retrieval call binding the contract method 0x076195b1.
//
// Solidity: function totalRefundToken1Amount() view returns(uint256)
func (_DFedPair *DFedPairCaller) TotalRefundToken1Amount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "totalRefundToken1Amount")
	return *ret0, err
}

// TotalRefundToken1Amount is a free data retrieval call binding the contract method 0x076195b1.
//
// Solidity: function totalRefundToken1Amount() view returns(uint256)
func (_DFedPair *DFedPairSession) TotalRefundToken1Amount() (*big.Int, error) {
	return _DFedPair.Contract.TotalRefundToken1Amount(&_DFedPair.CallOpts)
}

// TotalRefundToken1Amount is a free data retrieval call binding the contract method 0x076195b1.
//
// Solidity: function totalRefundToken1Amount() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) TotalRefundToken1Amount() (*big.Int, error) {
	return _DFedPair.Contract.TotalRefundToken1Amount(&_DFedPair.CallOpts)
}

// TotalRefundToken1EqualToken0Amount is a free data retrieval call binding the contract method 0x5b863ab9.
//
// Solidity: function totalRefundToken1EqualToken0Amount() view returns(uint256)
func (_DFedPair *DFedPairCaller) TotalRefundToken1EqualToken0Amount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "totalRefundToken1EqualToken0Amount")
	return *ret0, err
}

// TotalRefundToken1EqualToken0Amount is a free data retrieval call binding the contract method 0x5b863ab9.
//
// Solidity: function totalRefundToken1EqualToken0Amount() view returns(uint256)
func (_DFedPair *DFedPairSession) TotalRefundToken1EqualToken0Amount() (*big.Int, error) {
	return _DFedPair.Contract.TotalRefundToken1EqualToken0Amount(&_DFedPair.CallOpts)
}

// TotalRefundToken1EqualToken0Amount is a free data retrieval call binding the contract method 0x5b863ab9.
//
// Solidity: function totalRefundToken1EqualToken0Amount() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) TotalRefundToken1EqualToken0Amount() (*big.Int, error) {
	return _DFedPair.Contract.TotalRefundToken1EqualToken0Amount(&_DFedPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedPair *DFedPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedPair *DFedPairSession) TotalSupply() (*big.Int, error) {
	return _DFedPair.Contract.TotalSupply(&_DFedPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) TotalSupply() (*big.Int, error) {
	return _DFedPair.Contract.TotalSupply(&_DFedPair.CallOpts)
}

// UniqueDebtId is a free data retrieval call binding the contract method 0xcfc6aa9e.
//
// Solidity: function uniqueDebtId() view returns(uint256)
func (_DFedPair *DFedPairCaller) UniqueDebtId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "uniqueDebtId")
	return *ret0, err
}

// UniqueDebtId is a free data retrieval call binding the contract method 0xcfc6aa9e.
//
// Solidity: function uniqueDebtId() view returns(uint256)
func (_DFedPair *DFedPairSession) UniqueDebtId() (*big.Int, error) {
	return _DFedPair.Contract.UniqueDebtId(&_DFedPair.CallOpts)
}

// UniqueDebtId is a free data retrieval call binding the contract method 0xcfc6aa9e.
//
// Solidity: function uniqueDebtId() view returns(uint256)
func (_DFedPair *DFedPairCallerSession) UniqueDebtId() (*big.Int, error) {
	return _DFedPair.Contract.UniqueDebtId(&_DFedPair.CallOpts)
}

// UserRewardDebtPerFEDlp is a free data retrieval call binding the contract method 0xd6fedef8.
//
// Solidity: function userRewardDebtPerFEDlp(address ) view returns(uint256)
func (_DFedPair *DFedPairCaller) UserRewardDebtPerFEDlp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DFedPair.contract.Call(opts, out, "userRewardDebtPerFEDlp", arg0)
	return *ret0, err
}

// UserRewardDebtPerFEDlp is a free data retrieval call binding the contract method 0xd6fedef8.
//
// Solidity: function userRewardDebtPerFEDlp(address ) view returns(uint256)
func (_DFedPair *DFedPairSession) UserRewardDebtPerFEDlp(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.UserRewardDebtPerFEDlp(&_DFedPair.CallOpts, arg0)
}

// UserRewardDebtPerFEDlp is a free data retrieval call binding the contract method 0xd6fedef8.
//
// Solidity: function userRewardDebtPerFEDlp(address ) view returns(uint256)
func (_DFedPair *DFedPairCallerSession) UserRewardDebtPerFEDlp(arg0 common.Address) (*big.Int, error) {
	return _DFedPair.Contract.UserRewardDebtPerFEDlp(&_DFedPair.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedPair *DFedPairTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedPair *DFedPairSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Approve(&_DFedPair.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DFedPair *DFedPairTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Approve(&_DFedPair.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _to) returns(uint256 _amount0, uint256 _amount1)
func (_DFedPair *DFedPairTransactor) Burn(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "burn", _to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _to) returns(uint256 _amount0, uint256 _amount1)
func (_DFedPair *DFedPairSession) Burn(_to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Burn(&_DFedPair.TransactOpts, _to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _to) returns(uint256 _amount0, uint256 _amount1)
func (_DFedPair *DFedPairTransactorSession) Burn(_to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Burn(&_DFedPair.TransactOpts, _to)
}

// BuyToken is a paid mutator transaction binding the contract method 0x21466eb5.
//
// Solidity: function buyToken(uint256 _amountIn, uint256 _amountOutMin, address _to) returns()
func (_DFedPair *DFedPairTransactor) BuyToken(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "buyToken", _amountIn, _amountOutMin, _to)
}

// BuyToken is a paid mutator transaction binding the contract method 0x21466eb5.
//
// Solidity: function buyToken(uint256 _amountIn, uint256 _amountOutMin, address _to) returns()
func (_DFedPair *DFedPairSession) BuyToken(_amountIn *big.Int, _amountOutMin *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.BuyToken(&_DFedPair.TransactOpts, _amountIn, _amountOutMin, _to)
}

// BuyToken is a paid mutator transaction binding the contract method 0x21466eb5.
//
// Solidity: function buyToken(uint256 _amountIn, uint256 _amountOutMin, address _to) returns()
func (_DFedPair *DFedPairTransactorSession) BuyToken(_amountIn *big.Int, _amountOutMin *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.BuyToken(&_DFedPair.TransactOpts, _amountIn, _amountOutMin, _to)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address _to) returns()
func (_DFedPair *DFedPairTransactor) Harvest(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "harvest", _to)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address _to) returns()
func (_DFedPair *DFedPairSession) Harvest(_to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Harvest(&_DFedPair.TransactOpts, _to)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address _to) returns()
func (_DFedPair *DFedPairTransactorSession) Harvest(_to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Harvest(&_DFedPair.TransactOpts, _to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _baseToken, address _token) returns()
func (_DFedPair *DFedPairTransactor) Initialize(opts *bind.TransactOpts, _baseToken common.Address, _token common.Address) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "initialize", _baseToken, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _baseToken, address _token) returns()
func (_DFedPair *DFedPairSession) Initialize(_baseToken common.Address, _token common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Initialize(&_DFedPair.TransactOpts, _baseToken, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _baseToken, address _token) returns()
func (_DFedPair *DFedPairTransactorSession) Initialize(_baseToken common.Address, _token common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Initialize(&_DFedPair.TransactOpts, _baseToken, _token)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _to, uint256 _amount0, uint256 _amount1) returns(uint256 _liquidity)
func (_DFedPair *DFedPairTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount0 *big.Int, _amount1 *big.Int) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "mint", _to, _amount0, _amount1)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _to, uint256 _amount0, uint256 _amount1) returns(uint256 _liquidity)
func (_DFedPair *DFedPairSession) Mint(_to common.Address, _amount0 *big.Int, _amount1 *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Mint(&_DFedPair.TransactOpts, _to, _amount0, _amount1)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _to, uint256 _amount0, uint256 _amount1) returns(uint256 _liquidity)
func (_DFedPair *DFedPairTransactorSession) Mint(_to common.Address, _amount0 *big.Int, _amount1 *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Mint(&_DFedPair.TransactOpts, _to, _amount0, _amount1)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 _pledgeAmount, uint256 _targetNum, address _to) returns(uint256 _debtId)
func (_DFedPair *DFedPairTransactor) Mortgage(opts *bind.TransactOpts, _pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "mortgage", _pledgeAmount, _targetNum, _to)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 _pledgeAmount, uint256 _targetNum, address _to) returns(uint256 _debtId)
func (_DFedPair *DFedPairSession) Mortgage(_pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Mortgage(&_DFedPair.TransactOpts, _pledgeAmount, _targetNum, _to)
}

// Mortgage is a paid mutator transaction binding the contract method 0x58687e63.
//
// Solidity: function mortgage(uint256 _pledgeAmount, uint256 _targetNum, address _to) returns(uint256 _debtId)
func (_DFedPair *DFedPairTransactorSession) Mortgage(_pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.Mortgage(&_DFedPair.TransactOpts, _pledgeAmount, _targetNum, _to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedPair *DFedPairTransactor) Permit(opts *bind.TransactOpts, _from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "permit", _from, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedPair *DFedPairSession) Permit(_from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedPair.Contract.Permit(&_DFedPair.TransactOpts, _from, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _from, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedPair *DFedPairTransactorSession) Permit(_from common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedPair.Contract.Permit(&_DFedPair.TransactOpts, _from, _spender, _value, _deadline, _v, _r, _s)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _debtId) returns()
func (_DFedPair *DFedPairTransactor) Repay(opts *bind.TransactOpts, _debtId *big.Int) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "repay", _debtId)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _debtId) returns()
func (_DFedPair *DFedPairSession) Repay(_debtId *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Repay(&_DFedPair.TransactOpts, _debtId)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _debtId) returns()
func (_DFedPair *DFedPairTransactorSession) Repay(_debtId *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Repay(&_DFedPair.TransactOpts, _debtId)
}

// SellToken is a paid mutator transaction binding the contract method 0x88036ac5.
//
// Solidity: function sellToken(uint256 _amountIn, uint256 _amountOutMin, address _to) returns()
func (_DFedPair *DFedPairTransactor) SellToken(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "sellToken", _amountIn, _amountOutMin, _to)
}

// SellToken is a paid mutator transaction binding the contract method 0x88036ac5.
//
// Solidity: function sellToken(uint256 _amountIn, uint256 _amountOutMin, address _to) returns()
func (_DFedPair *DFedPairSession) SellToken(_amountIn *big.Int, _amountOutMin *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.SellToken(&_DFedPair.TransactOpts, _amountIn, _amountOutMin, _to)
}

// SellToken is a paid mutator transaction binding the contract method 0x88036ac5.
//
// Solidity: function sellToken(uint256 _amountIn, uint256 _amountOutMin, address _to) returns()
func (_DFedPair *DFedPairTransactorSession) SellToken(_amountIn *big.Int, _amountOutMin *big.Int, _to common.Address) (*types.Transaction, error) {
	return _DFedPair.Contract.SellToken(&_DFedPair.TransactOpts, _amountIn, _amountOutMin, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedPair *DFedPairTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedPair *DFedPairSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Transfer(&_DFedPair.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DFedPair *DFedPairTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.Transfer(&_DFedPair.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedPair *DFedPairTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedPair *DFedPairSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.TransferFrom(&_DFedPair.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DFedPair *DFedPairTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.TransferFrom(&_DFedPair.TransactOpts, _from, _to, _value)
}

// WithdrawToken0 is a paid mutator transaction binding the contract method 0xe60688cc.
//
// Solidity: function withdrawToken0(address _to, uint256 _amount) returns()
func (_DFedPair *DFedPairTransactor) WithdrawToken0(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "withdrawToken0", _to, _amount)
}

// WithdrawToken0 is a paid mutator transaction binding the contract method 0xe60688cc.
//
// Solidity: function withdrawToken0(address _to, uint256 _amount) returns()
func (_DFedPair *DFedPairSession) WithdrawToken0(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.WithdrawToken0(&_DFedPair.TransactOpts, _to, _amount)
}

// WithdrawToken0 is a paid mutator transaction binding the contract method 0xe60688cc.
//
// Solidity: function withdrawToken0(address _to, uint256 _amount) returns()
func (_DFedPair *DFedPairTransactorSession) WithdrawToken0(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.WithdrawToken0(&_DFedPair.TransactOpts, _to, _amount)
}

// WithdrawToken1 is a paid mutator transaction binding the contract method 0x7c3a2473.
//
// Solidity: function withdrawToken1(address _to, uint256 _amount) returns()
func (_DFedPair *DFedPairTransactor) WithdrawToken1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DFedPair.contract.Transact(opts, "withdrawToken1", _to, _amount)
}

// WithdrawToken1 is a paid mutator transaction binding the contract method 0x7c3a2473.
//
// Solidity: function withdrawToken1(address _to, uint256 _amount) returns()
func (_DFedPair *DFedPairSession) WithdrawToken1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.WithdrawToken1(&_DFedPair.TransactOpts, _to, _amount)
}

// WithdrawToken1 is a paid mutator transaction binding the contract method 0x7c3a2473.
//
// Solidity: function withdrawToken1(address _to, uint256 _amount) returns()
func (_DFedPair *DFedPairTransactorSession) WithdrawToken1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DFedPair.Contract.WithdrawToken1(&_DFedPair.TransactOpts, _to, _amount)
}

// DFedPairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DFedPair contract.
type DFedPairApprovalIterator struct {
	Event *DFedPairApproval // Event containing the contract specifics and raw log

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
func (it *DFedPairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairApproval)
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
		it.Event = new(DFedPairApproval)
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
func (it *DFedPairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairApproval represents a Approval event raised by the DFedPair contract.
type DFedPairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DFedPair *DFedPairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DFedPairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DFedPairApprovalIterator{contract: _DFedPair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DFedPair *DFedPairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DFedPairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairApproval)
				if err := _DFedPair.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DFedPair *DFedPairFilterer) ParseApproval(log types.Log) (*DFedPairApproval, error) {
	event := new(DFedPairApproval)
	if err := _DFedPair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the DFedPair contract.
type DFedPairBurnIterator struct {
	Event *DFedPairBurn // Event containing the contract specifics and raw log

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
func (it *DFedPairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairBurn)
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
		it.Event = new(DFedPairBurn)
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
func (it *DFedPairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairBurn represents a Burn event raised by the DFedPair contract.
type DFedPairBurn struct {
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
func (_DFedPair *DFedPairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*DFedPairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DFedPairBurnIterator{contract: _DFedPair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 lpAmount, uint256 amount0, uint256 amount1, address indexed to)
func (_DFedPair *DFedPairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DFedPairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairBurn)
				if err := _DFedPair.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_DFedPair *DFedPairFilterer) ParseBurn(log types.Log) (*DFedPairBurn, error) {
	event := new(DFedPairBurn)
	if err := _DFedPair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairBuyTokenIterator is returned from FilterBuyToken and is used to iterate over the raw logs and unpacked data for BuyToken events raised by the DFedPair contract.
type DFedPairBuyTokenIterator struct {
	Event *DFedPairBuyToken // Event containing the contract specifics and raw log

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
func (it *DFedPairBuyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairBuyToken)
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
		it.Event = new(DFedPairBuyToken)
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
func (it *DFedPairBuyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairBuyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairBuyToken represents a BuyToken event raised by the DFedPair contract.
type DFedPairBuyToken struct {
	Sender    common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBuyToken is a free log retrieval operation binding the contract event 0xf2fc1627c485a28da79658364df5619455857b1d9d2fcc5787b19b45555665f3.
//
// Solidity: event BuyToken(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed to)
func (_DFedPair *DFedPairFilterer) FilterBuyToken(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*DFedPairBuyTokenIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "BuyToken", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DFedPairBuyTokenIterator{contract: _DFedPair.contract, event: "BuyToken", logs: logs, sub: sub}, nil
}

// WatchBuyToken is a free log subscription operation binding the contract event 0xf2fc1627c485a28da79658364df5619455857b1d9d2fcc5787b19b45555665f3.
//
// Solidity: event BuyToken(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed to)
func (_DFedPair *DFedPairFilterer) WatchBuyToken(opts *bind.WatchOpts, sink chan<- *DFedPairBuyToken, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "BuyToken", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairBuyToken)
				if err := _DFedPair.contract.UnpackLog(event, "BuyToken", log); err != nil {
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

// ParseBuyToken is a log parse operation binding the contract event 0xf2fc1627c485a28da79658364df5619455857b1d9d2fcc5787b19b45555665f3.
//
// Solidity: event BuyToken(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed to)
func (_DFedPair *DFedPairFilterer) ParseBuyToken(log types.Log) (*DFedPairBuyToken, error) {
	event := new(DFedPairBuyToken)
	if err := _DFedPair.contract.UnpackLog(event, "BuyToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairDebtUpdateIterator is returned from FilterDebtUpdate and is used to iterate over the raw logs and unpacked data for DebtUpdate events raised by the DFedPair contract.
type DFedPairDebtUpdateIterator struct {
	Event *DFedPairDebtUpdate // Event containing the contract specifics and raw log

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
func (it *DFedPairDebtUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairDebtUpdate)
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
		it.Event = new(DFedPairDebtUpdate)
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
func (it *DFedPairDebtUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairDebtUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairDebtUpdate represents a DebtUpdate event raised by the DFedPair contract.
type DFedPairDebtUpdate struct {
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
func (_DFedPair *DFedPairFilterer) FilterDebtUpdate(opts *bind.FilterOpts, owner []common.Address) (*DFedPairDebtUpdateIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "DebtUpdate", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DFedPairDebtUpdateIterator{contract: _DFedPair.contract, event: "DebtUpdate", logs: logs, sub: sub}, nil
}

// WatchDebtUpdate is a free log subscription operation binding the contract event 0xa0980114e19bf18cfe2cba04908650c32aafcfdf5e5edb32ad16b072a7a33869.
//
// Solidity: event DebtUpdate(address indexed owner, uint256 debtId, uint256 pledgeAmount, uint256 repayAmount, uint256 debtToken0Amount, uint256 debtToken1Amount)
func (_DFedPair *DFedPairFilterer) WatchDebtUpdate(opts *bind.WatchOpts, sink chan<- *DFedPairDebtUpdate, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "DebtUpdate", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairDebtUpdate)
				if err := _DFedPair.contract.UnpackLog(event, "DebtUpdate", log); err != nil {
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
func (_DFedPair *DFedPairFilterer) ParseDebtUpdate(log types.Log) (*DFedPairDebtUpdate, error) {
	event := new(DFedPairDebtUpdate)
	if err := _DFedPair.contract.UnpackLog(event, "DebtUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DFedPair contract.
type DFedPairMintIterator struct {
	Event *DFedPairMint // Event containing the contract specifics and raw log

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
func (it *DFedPairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairMint)
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
		it.Event = new(DFedPairMint)
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
func (it *DFedPairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairMint represents a Mint event raised by the DFedPair contract.
type DFedPairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_DFedPair *DFedPairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*DFedPairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &DFedPairMintIterator{contract: _DFedPair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_DFedPair *DFedPairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DFedPairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairMint)
				if err := _DFedPair.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_DFedPair *DFedPairFilterer) ParseMint(log types.Log) (*DFedPairMint, error) {
	event := new(DFedPairMint)
	if err := _DFedPair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairMortgageIterator is returned from FilterMortgage and is used to iterate over the raw logs and unpacked data for Mortgage events raised by the DFedPair contract.
type DFedPairMortgageIterator struct {
	Event *DFedPairMortgage // Event containing the contract specifics and raw log

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
func (it *DFedPairMortgageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairMortgage)
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
		it.Event = new(DFedPairMortgage)
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
func (it *DFedPairMortgageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairMortgageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairMortgage represents a Mortgage event raised by the DFedPair contract.
type DFedPairMortgage struct {
	Owner        common.Address
	PledgeAmount *big.Int
	TargetNum    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMortgage is a free log retrieval operation binding the contract event 0x17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a4.
//
// Solidity: event Mortgage(address owner, uint256 pledgeAmount, uint256 targetNum)
func (_DFedPair *DFedPairFilterer) FilterMortgage(opts *bind.FilterOpts) (*DFedPairMortgageIterator, error) {

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "Mortgage")
	if err != nil {
		return nil, err
	}
	return &DFedPairMortgageIterator{contract: _DFedPair.contract, event: "Mortgage", logs: logs, sub: sub}, nil
}

// WatchMortgage is a free log subscription operation binding the contract event 0x17835e5a0c1ae820969d0dffccd9593a6c761f84cc747293a6ed8a25052a82a4.
//
// Solidity: event Mortgage(address owner, uint256 pledgeAmount, uint256 targetNum)
func (_DFedPair *DFedPairFilterer) WatchMortgage(opts *bind.WatchOpts, sink chan<- *DFedPairMortgage) (event.Subscription, error) {

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "Mortgage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairMortgage)
				if err := _DFedPair.contract.UnpackLog(event, "Mortgage", log); err != nil {
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
func (_DFedPair *DFedPairFilterer) ParseMortgage(log types.Log) (*DFedPairMortgage, error) {
	event := new(DFedPairMortgage)
	if err := _DFedPair.contract.UnpackLog(event, "Mortgage", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the DFedPair contract.
type DFedPairRepayIterator struct {
	Event *DFedPairRepay // Event containing the contract specifics and raw log

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
func (it *DFedPairRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairRepay)
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
		it.Event = new(DFedPairRepay)
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
func (it *DFedPairRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairRepay represents a Repay event raised by the DFedPair contract.
type DFedPairRepay struct {
	DebtId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe46230.
//
// Solidity: event Repay(uint256 debtId)
func (_DFedPair *DFedPairFilterer) FilterRepay(opts *bind.FilterOpts) (*DFedPairRepayIterator, error) {

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "Repay")
	if err != nil {
		return nil, err
	}
	return &DFedPairRepayIterator{contract: _DFedPair.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xa6ffc78a660e4971a47a0f916a0abae483804e6f42c9292ed06aa64f8fe46230.
//
// Solidity: event Repay(uint256 debtId)
func (_DFedPair *DFedPairFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *DFedPairRepay) (event.Subscription, error) {

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "Repay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairRepay)
				if err := _DFedPair.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_DFedPair *DFedPairFilterer) ParseRepay(log types.Log) (*DFedPairRepay, error) {
	event := new(DFedPairRepay)
	if err := _DFedPair.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairSellTokenIterator is returned from FilterSellToken and is used to iterate over the raw logs and unpacked data for SellToken events raised by the DFedPair contract.
type DFedPairSellTokenIterator struct {
	Event *DFedPairSellToken // Event containing the contract specifics and raw log

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
func (it *DFedPairSellTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairSellToken)
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
		it.Event = new(DFedPairSellToken)
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
func (it *DFedPairSellTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairSellTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairSellToken represents a SellToken event raised by the DFedPair contract.
type DFedPairSellToken struct {
	Sender    common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSellToken is a free log retrieval operation binding the contract event 0x1b38ce2719d9be70e0c768bd2a89d778eaef1cd41e3a4c02fa2317d415c3d348.
//
// Solidity: event SellToken(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed to)
func (_DFedPair *DFedPairFilterer) FilterSellToken(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*DFedPairSellTokenIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "SellToken", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DFedPairSellTokenIterator{contract: _DFedPair.contract, event: "SellToken", logs: logs, sub: sub}, nil
}

// WatchSellToken is a free log subscription operation binding the contract event 0x1b38ce2719d9be70e0c768bd2a89d778eaef1cd41e3a4c02fa2317d415c3d348.
//
// Solidity: event SellToken(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed to)
func (_DFedPair *DFedPairFilterer) WatchSellToken(opts *bind.WatchOpts, sink chan<- *DFedPairSellToken, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "SellToken", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairSellToken)
				if err := _DFedPair.contract.UnpackLog(event, "SellToken", log); err != nil {
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

// ParseSellToken is a log parse operation binding the contract event 0x1b38ce2719d9be70e0c768bd2a89d778eaef1cd41e3a4c02fa2317d415c3d348.
//
// Solidity: event SellToken(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed to)
func (_DFedPair *DFedPairFilterer) ParseSellToken(log types.Log) (*DFedPairSellToken, error) {
	event := new(DFedPairSellToken)
	if err := _DFedPair.contract.UnpackLog(event, "SellToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DFedPairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DFedPair contract.
type DFedPairTransferIterator struct {
	Event *DFedPairTransfer // Event containing the contract specifics and raw log

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
func (it *DFedPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DFedPairTransfer)
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
		it.Event = new(DFedPairTransfer)
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
func (it *DFedPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DFedPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DFedPairTransfer represents a Transfer event raised by the DFedPair contract.
type DFedPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DFedPair *DFedPairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DFedPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DFedPairTransferIterator{contract: _DFedPair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DFedPair *DFedPairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DFedPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DFedPair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DFedPairTransfer)
				if err := _DFedPair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DFedPair *DFedPairFilterer) ParseTransfer(log types.Log) (*DFedPairTransfer, error) {
	event := new(DFedPairTransfer)
	if err := _DFedPair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
