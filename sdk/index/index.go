// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package index

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

// IdFedERC20ABI is the input ABI used to generate the binding from.
const IdFedERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdFedERC20FuncSigs maps the 4-byte function signature to its string representation.
var IdFedERC20FuncSigs = map[string]string{
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

// IdFedERC20 is an auto generated Go binding around an Ethereum contract.
type IdFedERC20 struct {
	IdFedERC20Caller     // Read-only binding to the contract
	IdFedERC20Transactor // Write-only binding to the contract
	IdFedERC20Filterer   // Log filterer for contract events
}

// IdFedERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IdFedERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IdFedERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdFedERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdFedERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdFedERC20Session struct {
	Contract     *IdFedERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdFedERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdFedERC20CallerSession struct {
	Contract *IdFedERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IdFedERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdFedERC20TransactorSession struct {
	Contract     *IdFedERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IdFedERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IdFedERC20Raw struct {
	Contract *IdFedERC20 // Generic contract binding to access the raw methods on
}

// IdFedERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdFedERC20CallerRaw struct {
	Contract *IdFedERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IdFedERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdFedERC20TransactorRaw struct {
	Contract *IdFedERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIdFedERC20 creates a new instance of IdFedERC20, bound to a specific deployed contract.
func NewIdFedERC20(address common.Address, backend bind.ContractBackend) (*IdFedERC20, error) {
	contract, err := bindIdFedERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdFedERC20{IdFedERC20Caller: IdFedERC20Caller{contract: contract}, IdFedERC20Transactor: IdFedERC20Transactor{contract: contract}, IdFedERC20Filterer: IdFedERC20Filterer{contract: contract}}, nil
}

// NewIdFedERC20Caller creates a new read-only instance of IdFedERC20, bound to a specific deployed contract.
func NewIdFedERC20Caller(address common.Address, caller bind.ContractCaller) (*IdFedERC20Caller, error) {
	contract, err := bindIdFedERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedERC20Caller{contract: contract}, nil
}

// NewIdFedERC20Transactor creates a new write-only instance of IdFedERC20, bound to a specific deployed contract.
func NewIdFedERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IdFedERC20Transactor, error) {
	contract, err := bindIdFedERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdFedERC20Transactor{contract: contract}, nil
}

// NewIdFedERC20Filterer creates a new log filterer instance of IdFedERC20, bound to a specific deployed contract.
func NewIdFedERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IdFedERC20Filterer, error) {
	contract, err := bindIdFedERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdFedERC20Filterer{contract: contract}, nil
}

// bindIdFedERC20 binds a generic wrapper to an already deployed contract.
func bindIdFedERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdFedERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedERC20 *IdFedERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedERC20.Contract.IdFedERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedERC20 *IdFedERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedERC20.Contract.IdFedERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedERC20 *IdFedERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedERC20.Contract.IdFedERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdFedERC20 *IdFedERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdFedERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdFedERC20 *IdFedERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdFedERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdFedERC20 *IdFedERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdFedERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IdFedERC20 *IdFedERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IdFedERC20 *IdFedERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _IdFedERC20.Contract.DOMAINSEPARATOR(&_IdFedERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IdFedERC20 *IdFedERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IdFedERC20.Contract.DOMAINSEPARATOR(&_IdFedERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IdFedERC20 *IdFedERC20Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IdFedERC20 *IdFedERC20Session) PERMITTYPEHASH() ([32]byte, error) {
	return _IdFedERC20.Contract.PERMITTYPEHASH(&_IdFedERC20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IdFedERC20 *IdFedERC20CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IdFedERC20.Contract.PERMITTYPEHASH(&_IdFedERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedERC20 *IdFedERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedERC20 *IdFedERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IdFedERC20.Contract.Allowance(&_IdFedERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IdFedERC20 *IdFedERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IdFedERC20.Contract.Allowance(&_IdFedERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedERC20 *IdFedERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedERC20 *IdFedERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdFedERC20.Contract.BalanceOf(&_IdFedERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdFedERC20 *IdFedERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdFedERC20.Contract.BalanceOf(&_IdFedERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IdFedERC20 *IdFedERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IdFedERC20 *IdFedERC20Session) Decimals() (uint8, error) {
	return _IdFedERC20.Contract.Decimals(&_IdFedERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IdFedERC20 *IdFedERC20CallerSession) Decimals() (uint8, error) {
	return _IdFedERC20.Contract.Decimals(&_IdFedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IdFedERC20 *IdFedERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IdFedERC20 *IdFedERC20Session) Name() (string, error) {
	return _IdFedERC20.Contract.Name(&_IdFedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IdFedERC20 *IdFedERC20CallerSession) Name() (string, error) {
	return _IdFedERC20.Contract.Name(&_IdFedERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdFedERC20 *IdFedERC20Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "nonces", owner)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdFedERC20 *IdFedERC20Session) Nonces(owner common.Address) (*big.Int, error) {
	return _IdFedERC20.Contract.Nonces(&_IdFedERC20.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IdFedERC20 *IdFedERC20CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IdFedERC20.Contract.Nonces(&_IdFedERC20.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IdFedERC20 *IdFedERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IdFedERC20 *IdFedERC20Session) Symbol() (string, error) {
	return _IdFedERC20.Contract.Symbol(&_IdFedERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IdFedERC20 *IdFedERC20CallerSession) Symbol() (string, error) {
	return _IdFedERC20.Contract.Symbol(&_IdFedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedERC20 *IdFedERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdFedERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedERC20 *IdFedERC20Session) TotalSupply() (*big.Int, error) {
	return _IdFedERC20.Contract.TotalSupply(&_IdFedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IdFedERC20 *IdFedERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IdFedERC20.Contract.TotalSupply(&_IdFedERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.Contract.Approve(&_IdFedERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.Contract.Approve(&_IdFedERC20.TransactOpts, spender, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IdFedERC20 *IdFedERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdFedERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IdFedERC20 *IdFedERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdFedERC20.Contract.Permit(&_IdFedERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IdFedERC20 *IdFedERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IdFedERC20.Contract.Permit(&_IdFedERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.Contract.Transfer(&_IdFedERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.Contract.Transfer(&_IdFedERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.Contract.TransferFrom(&_IdFedERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IdFedERC20 *IdFedERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IdFedERC20.Contract.TransferFrom(&_IdFedERC20.TransactOpts, from, to, value)
}

// IdFedERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IdFedERC20 contract.
type IdFedERC20ApprovalIterator struct {
	Event *IdFedERC20Approval // Event containing the contract specifics and raw log

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
func (it *IdFedERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedERC20Approval)
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
		it.Event = new(IdFedERC20Approval)
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
func (it *IdFedERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedERC20Approval represents a Approval event raised by the IdFedERC20 contract.
type IdFedERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedERC20 *IdFedERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IdFedERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IdFedERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IdFedERC20ApprovalIterator{contract: _IdFedERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IdFedERC20 *IdFedERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IdFedERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IdFedERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedERC20Approval)
				if err := _IdFedERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IdFedERC20 *IdFedERC20Filterer) ParseApproval(log types.Log) (*IdFedERC20Approval, error) {
	event := new(IdFedERC20Approval)
	if err := _IdFedERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IdFedERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IdFedERC20 contract.
type IdFedERC20TransferIterator struct {
	Event *IdFedERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IdFedERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdFedERC20Transfer)
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
		it.Event = new(IdFedERC20Transfer)
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
func (it *IdFedERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdFedERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdFedERC20Transfer represents a Transfer event raised by the IdFedERC20 contract.
type IdFedERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedERC20 *IdFedERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IdFedERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IdFedERC20TransferIterator{contract: _IdFedERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IdFedERC20 *IdFedERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IdFedERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IdFedERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdFedERC20Transfer)
				if err := _IdFedERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IdFedERC20 *IdFedERC20Filterer) ParseTransfer(log types.Log) (*IdFedERC20Transfer, error) {
	event := new(IdFedERC20Transfer)
	if err := _IdFedERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// DFedIndexABI is the input ABI used to generate the binding from.
const DFedIndexABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_baseTokenAmountDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmountDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseTokenAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_baseTokenAmountDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmountDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseTokenAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmountMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"addLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"buyTokenWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"mortgage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_debtId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"repayWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DFedIndexFuncSigs maps the 4-byte function signature to its string representation.
var DFedIndexFuncSigs = map[string]string{
	"9187aaca": "addLiquidity(address,uint256,uint256,uint256,uint256,address,uint256)",
	"d9a3fbd3": "addLiquidityWithPermit(address,uint256,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"c55dae63": "baseToken()",
	"b880f2cd": "buyToken(address,uint256,uint256,address,uint256)",
	"705fb603": "buyTokenWithPermit(address,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"c45a0155": "factory()",
	"1a788a02": "getPair(address)",
	"ce256bc0": "mortgage(address,uint256,uint256,address,uint256)",
	"c6217733": "removeLiquidity(address,uint256,address,uint256,uint256,uint256)",
	"104870fd": "removeLiquidityWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
	"8cd2e0c7": "repay(address,uint256,uint256)",
	"d7187ec3": "repayWithPermit(address,uint256,uint256,uint256,uint8,bytes32,bytes32)",
	"ee91d954": "sellToken(address,uint256,uint256,address,uint256)",
}

// DFedIndexBin is the compiled bytecode used for deploying new contracts.
var DFedIndexBin = "0x608060405234801561001057600080fd5b506040516117413803806117418339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556116c78061007a6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063c45a01551161008c578063ce256bc011610066578063ce256bc014610322578063d7187ec314610364578063d9a3fbd3146103b1578063ee91d95414610420576100cf565b8063c45a0155146102cc578063c55dae63146102d4578063c6217733146102dc576100cf565b8063104870fd146100d45780631a788a0214610153578063705fb603146101955780638cd2e0c7146101f85780639187aaca1461022a578063b880f2cd1461028a575b600080fd5b61013a60048036036101408110156100eb57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a08101359060c081013515159060ff60e08201351690610100810135906101200135610462565b6040805192835260208301919091528051918290030190f35b6101796004803603602081101561016957600080fd5b50356001600160a01b0316610534565b604080516001600160a01b039092168252519081900360200190f35b6101f660048036036101208110156101ac57600080fd5b506001600160a01b03813581169160208101359160408201359160608101359091169060808101359060a081013515159060ff60c0820135169060e08101359061010001356105fa565b005b6101f66004803603606081101561020e57600080fd5b506001600160a01b038135169060208101359060400135610702565b610278600480360360e081101561024057600080fd5b506001600160a01b03813581169160208101359160408201359160608101359160808201359160a08101359091169060c00135610916565b60408051918252519081900360200190f35b6101f6600480360360a08110156102a057600080fd5b506001600160a01b03813581169160208101359160408201359160608101359091169060800135610cbd565b610179610da1565b610179610db0565b61013a600480360360c08110156102f257600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060808101359060a00135610dbf565b610278600480360360a081101561033857600080fd5b506001600160a01b03813581169160208101359160408201359160608101359091169060800135610f98565b6101f6600480360360e081101561037a57600080fd5b506001600160a01b038135169060208101359060408101359060608101359060ff6080820135169060a08101359060c00135611095565b61027860048036036101608110156103c857600080fd5b506001600160a01b03813581169160208101359160408201359160608101359160808201359160a08101359091169060c08101359060e081013515159060ff6101008201351690610120810135906101400135611184565b6101f6600480360360a081101561043657600080fd5b506001600160a01b03813581169160208101359160408201359160608101359091169060800135611295565b60008060006104708d610534565b905060008761047f578c610483565b6000195b6040805163d505accf60e01b815233600482015230602482015260448101839052606481018c905260ff8a16608482015260a4810189905260c4810188905290519192506001600160a01b0384169163d505accf9160e48082019260009290919082900301818387803b1580156104f957600080fd5b505af115801561050d573d6000803e3d6000fd5b5050505061051f828e8c8f8f8e610dbf565b909f909e509c50505050505050505050505050565b6000805460408051630d3c450160e11b81526001600160a01b03858116600483015291519190921691631a788a02916024808301926020929190829003018186803b15801561058257600080fd5b505afa158015610596573d6000803e3d6000fd5b505050506040513d60208110156105ac57600080fd5b505190506001600160a01b0381166105f55760405162461bcd60e51b81526004018080602001828103825260248152602001806116246024913960400191505060405180910390fd5b919050565b8442811015610645576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b6000856106525789610656565b6000195b6001546040805163d505accf60e01b815233600482015230602482015260448101849052606481018b905260ff8916608482015260a4810188905260c4810187905290519293506001600160a01b039091169163d505accf9160e48082019260009290919082900301818387803b1580156106d057600080fd5b505af11580156106e4573d6000803e3d6000fd5b505050506106f58b8b8b8b8b610cbd565b5050505050505050505050565b804281101561074d576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b600061075885610534565b90506000816001600160a01b031663ea52a0be836001600160a01b0316633c33a245886040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156107af57600080fd5b505afa1580156107c3573d6000803e3d6000fd5b505050506040513d60208110156107d957600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926020929190829003018186803b15801561081857600080fd5b505afa15801561082c573d6000803e3d6000fd5b505050506040513d602081101561084257600080fd5b5051905080610898576040805162461bcd60e51b815260206004820152601f60248201527f64466564496e6465783a2052455041595f414d4f554e545f494e56414c494400604482015290519081900360640190fd5b6001546108b0906001600160a01b0316338484611358565b816001600160a01b031663371fd8e6866040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156108f657600080fd5b505af115801561090a573d6000803e3d6000fd5b50505050505050505050565b60008142811015610963576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b6000805460408051630d3c450160e11b81526001600160a01b038d8116600483015291519190921691631a788a02916024808301926020929190829003018186803b1580156109b157600080fd5b505afa1580156109c5573d6000803e3d6000fd5b505050506040513d60208110156109db57600080fd5b505190506001600160a01b038116610a7e5760008054906101000a90046001600160a01b03166001600160a01b0316639ccb07448b6040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050602060405180830381600087803b158015610a4f57600080fd5b505af1158015610a63573d6000803e3d6000fd5b505050506040513d6020811015610a7957600080fd5b505190505b600080600080846001600160a01b0316630902f1ac6040518163ffffffff1660e01b8152600401604080518083038186803b158015610abc57600080fd5b505afa158015610ad0573d6000803e3d6000fd5b505050506040513d6040811015610ae657600080fd5b508051602090910151909250905081158015610b00575080155b15610b10578c93508b9250610bfc565b6000610b1d8e84846114b5565b90508c8111610b86578a811015610b7b576040805162461bcd60e51b815260206004820181905260248201527f64466564496e6465783a20494e53554646494349454e545f425f414d4f554e54604482015290519081900360640190fd5b8d9450925082610bfa565b6000610b938e84866114b5565b90508e811115610b9f57fe5b8c811015610bf4576040805162461bcd60e51b815260206004820181905260248201527f64466564496e6465783a20494e53554646494349454e545f415f414d4f554e54604482015290519081900360640190fd5b94508c93505b505b600154610c14906001600160a01b0316338787611358565b610c208e338786611358565b846001600160a01b031663156e29f68a86866040518463ffffffff1660e01b815260040180846001600160a01b031681526020018381526020018281526020019350505050602060405180830381600087803b158015610c7f57600080fd5b505af1158015610c93573d6000803e3d6000fd5b505050506040513d6020811015610ca957600080fd5b50519e9d5050505050505050505050505050565b8042811015610d08576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b6000610d1387610534565b600154909150610d2e906001600160a01b0316338389611358565b806001600160a01b03166321466eb58787876040518463ffffffff1660e01b815260040180848152602001838152602001826001600160a01b031681526020019350505050600060405180830381600087803b158015610d8d57600080fd5b505af11580156106f5573d6000803e3d6000fd5b6000546001600160a01b031681565b6001546001600160a01b031681565b6000808242811015610e0d576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b604080516323b872dd60e01b81523360048201526001600160a01b038b1660248201819052604482018b905291516323b872dd916064808201926020929091908290030181600087803b158015610e6357600080fd5b505af1158015610e77573d6000803e3d6000fd5b505050506040513d6020811015610e8d57600080fd5b50506040805163226bf2d160e21b81526001600160a01b0389811660048301528251908c16926389afcb4492602480820193918290030181600087803b158015610ed657600080fd5b505af1158015610eea573d6000803e3d6000fd5b505050506040513d6040811015610f0057600080fd5b508051602090910151909350915085831015610f4d5760405162461bcd60e51b81526004018080602001828103825260268152602001806115fe6026913960400191505060405180910390fd5b84821015610f8c5760405162461bcd60e51b815260040180806020018281038252602681526020018061166c6026913960400191505060405180910390fd5b50965096945050505050565b60008142811015610fe5576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b6000610ff088610534565b9050610ffe8833838a611358565b806001600160a01b03166358687e638888886040518463ffffffff1660e01b815260040180848152602001838152602001826001600160a01b031681526020019350505050602060405180830381600087803b15801561105d57600080fd5b505af1158015611071573d6000803e3d6000fd5b505050506040513d602081101561108757600080fd5b505198975050505050505050565b83428110156110e0576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b6001546040805163d505accf60e01b8152336004820152306024820152604481018990526064810188905260ff8716608482015260a4810186905260c4810185905290516001600160a01b039092169163d505accf9160e48082019260009290919082900301818387803b15801561115757600080fd5b505af115801561116b573d6000803e3d6000fd5b5050505061117a888887610702565b5050505050505050565b600085428110156111d1576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b6000866111de578c6111e2565b6000195b6001546040805163d505accf60e01b815233600482015230602482015260448101849052606481018c905260ff8a16608482015260a4810189905260c4810188905290519293506001600160a01b039091169163d505accf9160e48082019260009290919082900301818387803b15801561125c57600080fd5b505af1158015611270573d6000803e3d6000fd5b505050506112838e8e8e8e8e8e8e610916565b9e9d5050505050505050505050505050565b80428110156112e0576040805162461bcd60e51b815260206004820152601260248201527119119959125b99195e0e881156141254915160721b604482015290519081900360640190fd5b60006112eb87610534565b90506112f987338389611358565b806001600160a01b03166388036ac58787876040518463ffffffff1660e01b815260040180848152602001838152602001826001600160a01b031681526020019350505050600060405180830381600087803b158015610d8d57600080fd5b604080516001600160a01b0385811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180516001600160e01b03166323b872dd60e01b17815292518251600094606094938a169392918291908083835b602083106113dd5780518252601f1990920191602091820191016113be565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461143f576040519150601f19603f3d011682016040523d82523d6000602084013e611444565b606091505b5091509150818015611472575080511580611472575080806020019051602081101561146f57600080fd5b50515b6114ad5760405162461bcd60e51b81526004018080602001828103825260248152602001806116486024913960400191505060405180910390fd5b505050505050565b600080841161150b576040805162461bcd60e51b815260206004820181905260248201527f644665644c6962726172793a20494e53554646494349454e545f414d4f554e54604482015290519081900360640190fd5b60008311801561151b5750600082115b6115565760405162461bcd60e51b81526004018080602001828103825260238152602001806115db6023913960400191505060405180910390fd5b826115618584611571565b8161156857fe5b04949350505050565b600081158061158c5750508082028282828161158957fe5b04145b6115d4576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b9291505056fe644665644c6962726172793a20494e53554646494349454e545f4c495155494449545964466564496e6465783a20494e53554646494349454e545f4f55545055545f414d4f554e543064466564496e6465783a20544f4b454e5f504149525f444f45535f4e4f545f45584953545472616e7366657248656c7065723a205452414e534645525f46524f4d5f4641494c454464466564496e6465783a20494e53554646494349454e545f4f55545055545f414d4f554e5431a26469706673582212202622281264a8c7c5d9ee56c98ef4cd2ad0481d9f8d5472727c6edf0e3d1a520f64736f6c63430007000033"

// DeployDFedIndex deploys a new Ethereum contract, binding an instance of DFedIndex to it.
func DeployDFedIndex(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _baseToken common.Address) (common.Address, *types.Transaction, *DFedIndex, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedIndexABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DFedIndexBin), backend, _factory, _baseToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DFedIndex{DFedIndexCaller: DFedIndexCaller{contract: contract}, DFedIndexTransactor: DFedIndexTransactor{contract: contract}, DFedIndexFilterer: DFedIndexFilterer{contract: contract}}, nil
}

// DFedIndex is an auto generated Go binding around an Ethereum contract.
type DFedIndex struct {
	DFedIndexCaller     // Read-only binding to the contract
	DFedIndexTransactor // Write-only binding to the contract
	DFedIndexFilterer   // Log filterer for contract events
}

// DFedIndexCaller is an auto generated read-only Go binding around an Ethereum contract.
type DFedIndexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedIndexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DFedIndexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedIndexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DFedIndexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DFedIndexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DFedIndexSession struct {
	Contract     *DFedIndex        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DFedIndexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DFedIndexCallerSession struct {
	Contract *DFedIndexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DFedIndexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DFedIndexTransactorSession struct {
	Contract     *DFedIndexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DFedIndexRaw is an auto generated low-level Go binding around an Ethereum contract.
type DFedIndexRaw struct {
	Contract *DFedIndex // Generic contract binding to access the raw methods on
}

// DFedIndexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DFedIndexCallerRaw struct {
	Contract *DFedIndexCaller // Generic read-only contract binding to access the raw methods on
}

// DFedIndexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DFedIndexTransactorRaw struct {
	Contract *DFedIndexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDFedIndex creates a new instance of DFedIndex, bound to a specific deployed contract.
func NewDFedIndex(address common.Address, backend bind.ContractBackend) (*DFedIndex, error) {
	contract, err := bindDFedIndex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DFedIndex{DFedIndexCaller: DFedIndexCaller{contract: contract}, DFedIndexTransactor: DFedIndexTransactor{contract: contract}, DFedIndexFilterer: DFedIndexFilterer{contract: contract}}, nil
}

// NewDFedIndexCaller creates a new read-only instance of DFedIndex, bound to a specific deployed contract.
func NewDFedIndexCaller(address common.Address, caller bind.ContractCaller) (*DFedIndexCaller, error) {
	contract, err := bindDFedIndex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DFedIndexCaller{contract: contract}, nil
}

// NewDFedIndexTransactor creates a new write-only instance of DFedIndex, bound to a specific deployed contract.
func NewDFedIndexTransactor(address common.Address, transactor bind.ContractTransactor) (*DFedIndexTransactor, error) {
	contract, err := bindDFedIndex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DFedIndexTransactor{contract: contract}, nil
}

// NewDFedIndexFilterer creates a new log filterer instance of DFedIndex, bound to a specific deployed contract.
func NewDFedIndexFilterer(address common.Address, filterer bind.ContractFilterer) (*DFedIndexFilterer, error) {
	contract, err := bindDFedIndex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DFedIndexFilterer{contract: contract}, nil
}

// bindDFedIndex binds a generic wrapper to an already deployed contract.
func bindDFedIndex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DFedIndexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedIndex *DFedIndexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedIndex.Contract.DFedIndexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedIndex *DFedIndexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedIndex.Contract.DFedIndexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedIndex *DFedIndexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedIndex.Contract.DFedIndexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DFedIndex *DFedIndexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DFedIndex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DFedIndex *DFedIndexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DFedIndex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DFedIndex *DFedIndexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DFedIndex.Contract.contract.Transact(opts, method, params...)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DFedIndex *DFedIndexCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedIndex.contract.Call(opts, out, "baseToken")
	return *ret0, err
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DFedIndex *DFedIndexSession) BaseToken() (common.Address, error) {
	return _DFedIndex.Contract.BaseToken(&_DFedIndex.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DFedIndex *DFedIndexCallerSession) BaseToken() (common.Address, error) {
	return _DFedIndex.Contract.BaseToken(&_DFedIndex.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedIndex *DFedIndexCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedIndex.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedIndex *DFedIndexSession) Factory() (common.Address, error) {
	return _DFedIndex.Contract.Factory(&_DFedIndex.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_DFedIndex *DFedIndexCallerSession) Factory() (common.Address, error) {
	return _DFedIndex.Contract.Factory(&_DFedIndex.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address _token) view returns(address _pair)
func (_DFedIndex *DFedIndexCaller) GetPair(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DFedIndex.contract.Call(opts, out, "getPair", _token)
	return *ret0, err
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address _token) view returns(address _pair)
func (_DFedIndex *DFedIndexSession) GetPair(_token common.Address) (common.Address, error) {
	return _DFedIndex.Contract.GetPair(&_DFedIndex.CallOpts, _token)
}

// GetPair is a free data retrieval call binding the contract method 0x1a788a02.
//
// Solidity: function getPair(address _token) view returns(address _pair)
func (_DFedIndex *DFedIndexCallerSession) GetPair(_token common.Address) (common.Address, error) {
	return _DFedIndex.Contract.GetPair(&_DFedIndex.CallOpts, _token)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9187aaca.
//
// Solidity: function addLiquidity(address _token, uint256 _baseTokenAmountDesired, uint256 _tokenAmountDesired, uint256 _baseTokenAmountMin, uint256 _tokenAmountMin, address _to, uint256 _deadline) returns(uint256 _liquidity)
func (_DFedIndex *DFedIndexTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _baseTokenAmountDesired *big.Int, _tokenAmountDesired *big.Int, _baseTokenAmountMin *big.Int, _tokenAmountMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "addLiquidity", _token, _baseTokenAmountDesired, _tokenAmountDesired, _baseTokenAmountMin, _tokenAmountMin, _to, _deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9187aaca.
//
// Solidity: function addLiquidity(address _token, uint256 _baseTokenAmountDesired, uint256 _tokenAmountDesired, uint256 _baseTokenAmountMin, uint256 _tokenAmountMin, address _to, uint256 _deadline) returns(uint256 _liquidity)
func (_DFedIndex *DFedIndexSession) AddLiquidity(_token common.Address, _baseTokenAmountDesired *big.Int, _tokenAmountDesired *big.Int, _baseTokenAmountMin *big.Int, _tokenAmountMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.AddLiquidity(&_DFedIndex.TransactOpts, _token, _baseTokenAmountDesired, _tokenAmountDesired, _baseTokenAmountMin, _tokenAmountMin, _to, _deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9187aaca.
//
// Solidity: function addLiquidity(address _token, uint256 _baseTokenAmountDesired, uint256 _tokenAmountDesired, uint256 _baseTokenAmountMin, uint256 _tokenAmountMin, address _to, uint256 _deadline) returns(uint256 _liquidity)
func (_DFedIndex *DFedIndexTransactorSession) AddLiquidity(_token common.Address, _baseTokenAmountDesired *big.Int, _tokenAmountDesired *big.Int, _baseTokenAmountMin *big.Int, _tokenAmountMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.AddLiquidity(&_DFedIndex.TransactOpts, _token, _baseTokenAmountDesired, _tokenAmountDesired, _baseTokenAmountMin, _tokenAmountMin, _to, _deadline)
}

// AddLiquidityWithPermit is a paid mutator transaction binding the contract method 0xd9a3fbd3.
//
// Solidity: function addLiquidityWithPermit(address _token, uint256 _baseTokenAmountDesired, uint256 _tokenAmountDesired, uint256 _baseTokenAmountMin, uint256 _tokenAmountMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _liquidity)
func (_DFedIndex *DFedIndexTransactor) AddLiquidityWithPermit(opts *bind.TransactOpts, _token common.Address, _baseTokenAmountDesired *big.Int, _tokenAmountDesired *big.Int, _baseTokenAmountMin *big.Int, _tokenAmountMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "addLiquidityWithPermit", _token, _baseTokenAmountDesired, _tokenAmountDesired, _baseTokenAmountMin, _tokenAmountMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// AddLiquidityWithPermit is a paid mutator transaction binding the contract method 0xd9a3fbd3.
//
// Solidity: function addLiquidityWithPermit(address _token, uint256 _baseTokenAmountDesired, uint256 _tokenAmountDesired, uint256 _baseTokenAmountMin, uint256 _tokenAmountMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _liquidity)
func (_DFedIndex *DFedIndexSession) AddLiquidityWithPermit(_token common.Address, _baseTokenAmountDesired *big.Int, _tokenAmountDesired *big.Int, _baseTokenAmountMin *big.Int, _tokenAmountMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.AddLiquidityWithPermit(&_DFedIndex.TransactOpts, _token, _baseTokenAmountDesired, _tokenAmountDesired, _baseTokenAmountMin, _tokenAmountMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// AddLiquidityWithPermit is a paid mutator transaction binding the contract method 0xd9a3fbd3.
//
// Solidity: function addLiquidityWithPermit(address _token, uint256 _baseTokenAmountDesired, uint256 _tokenAmountDesired, uint256 _baseTokenAmountMin, uint256 _tokenAmountMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _liquidity)
func (_DFedIndex *DFedIndexTransactorSession) AddLiquidityWithPermit(_token common.Address, _baseTokenAmountDesired *big.Int, _tokenAmountDesired *big.Int, _baseTokenAmountMin *big.Int, _tokenAmountMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.AddLiquidityWithPermit(&_DFedIndex.TransactOpts, _token, _baseTokenAmountDesired, _tokenAmountDesired, _baseTokenAmountMin, _tokenAmountMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// BuyToken is a paid mutator transaction binding the contract method 0xb880f2cd.
//
// Solidity: function buyToken(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexTransactor) BuyToken(opts *bind.TransactOpts, _token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "buyToken", _token, _amountIn, _amountOutMin, _to, _deadline)
}

// BuyToken is a paid mutator transaction binding the contract method 0xb880f2cd.
//
// Solidity: function buyToken(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexSession) BuyToken(_token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.BuyToken(&_DFedIndex.TransactOpts, _token, _amountIn, _amountOutMin, _to, _deadline)
}

// BuyToken is a paid mutator transaction binding the contract method 0xb880f2cd.
//
// Solidity: function buyToken(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexTransactorSession) BuyToken(_token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.BuyToken(&_DFedIndex.TransactOpts, _token, _amountIn, _amountOutMin, _to, _deadline)
}

// BuyTokenWithPermit is a paid mutator transaction binding the contract method 0x705fb603.
//
// Solidity: function buyTokenWithPermit(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedIndex *DFedIndexTransactor) BuyTokenWithPermit(opts *bind.TransactOpts, _token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "buyTokenWithPermit", _token, _amountIn, _amountOutMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// BuyTokenWithPermit is a paid mutator transaction binding the contract method 0x705fb603.
//
// Solidity: function buyTokenWithPermit(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedIndex *DFedIndexSession) BuyTokenWithPermit(_token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.BuyTokenWithPermit(&_DFedIndex.TransactOpts, _token, _amountIn, _amountOutMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// BuyTokenWithPermit is a paid mutator transaction binding the contract method 0x705fb603.
//
// Solidity: function buyTokenWithPermit(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedIndex *DFedIndexTransactorSession) BuyTokenWithPermit(_token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.BuyTokenWithPermit(&_DFedIndex.TransactOpts, _token, _amountIn, _amountOutMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// Mortgage is a paid mutator transaction binding the contract method 0xce256bc0.
//
// Solidity: function mortgage(address _token, uint256 _pledgeAmount, uint256 _targetNum, address _to, uint256 _deadline) returns(uint256 _debtId)
func (_DFedIndex *DFedIndexTransactor) Mortgage(opts *bind.TransactOpts, _token common.Address, _pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "mortgage", _token, _pledgeAmount, _targetNum, _to, _deadline)
}

// Mortgage is a paid mutator transaction binding the contract method 0xce256bc0.
//
// Solidity: function mortgage(address _token, uint256 _pledgeAmount, uint256 _targetNum, address _to, uint256 _deadline) returns(uint256 _debtId)
func (_DFedIndex *DFedIndexSession) Mortgage(_token common.Address, _pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.Mortgage(&_DFedIndex.TransactOpts, _token, _pledgeAmount, _targetNum, _to, _deadline)
}

// Mortgage is a paid mutator transaction binding the contract method 0xce256bc0.
//
// Solidity: function mortgage(address _token, uint256 _pledgeAmount, uint256 _targetNum, address _to, uint256 _deadline) returns(uint256 _debtId)
func (_DFedIndex *DFedIndexTransactorSession) Mortgage(_token common.Address, _pledgeAmount *big.Int, _targetNum *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.Mortgage(&_DFedIndex.TransactOpts, _token, _pledgeAmount, _targetNum, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc6217733.
//
// Solidity: function removeLiquidity(address _pair, uint256 _liquidity, address _to, uint256 _amount0Min, uint256 _amount1Min, uint256 _deadline) returns(uint256 _amount0, uint256 _amount1)
func (_DFedIndex *DFedIndexTransactor) RemoveLiquidity(opts *bind.TransactOpts, _pair common.Address, _liquidity *big.Int, _to common.Address, _amount0Min *big.Int, _amount1Min *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "removeLiquidity", _pair, _liquidity, _to, _amount0Min, _amount1Min, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc6217733.
//
// Solidity: function removeLiquidity(address _pair, uint256 _liquidity, address _to, uint256 _amount0Min, uint256 _amount1Min, uint256 _deadline) returns(uint256 _amount0, uint256 _amount1)
func (_DFedIndex *DFedIndexSession) RemoveLiquidity(_pair common.Address, _liquidity *big.Int, _to common.Address, _amount0Min *big.Int, _amount1Min *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.RemoveLiquidity(&_DFedIndex.TransactOpts, _pair, _liquidity, _to, _amount0Min, _amount1Min, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc6217733.
//
// Solidity: function removeLiquidity(address _pair, uint256 _liquidity, address _to, uint256 _amount0Min, uint256 _amount1Min, uint256 _deadline) returns(uint256 _amount0, uint256 _amount1)
func (_DFedIndex *DFedIndexTransactorSession) RemoveLiquidity(_pair common.Address, _liquidity *big.Int, _to common.Address, _amount0Min *big.Int, _amount1Min *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.RemoveLiquidity(&_DFedIndex.TransactOpts, _pair, _liquidity, _to, _amount0Min, _amount1Min, _deadline)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x104870fd.
//
// Solidity: function removeLiquidityWithPermit(address _token, uint256 _liquidity, uint256 _amount0Min, uint256 _amount1Min, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amount0, uint256 _amount1)
func (_DFedIndex *DFedIndexTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, _token common.Address, _liquidity *big.Int, _amount0Min *big.Int, _amount1Min *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "removeLiquidityWithPermit", _token, _liquidity, _amount0Min, _amount1Min, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x104870fd.
//
// Solidity: function removeLiquidityWithPermit(address _token, uint256 _liquidity, uint256 _amount0Min, uint256 _amount1Min, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amount0, uint256 _amount1)
func (_DFedIndex *DFedIndexSession) RemoveLiquidityWithPermit(_token common.Address, _liquidity *big.Int, _amount0Min *big.Int, _amount1Min *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.RemoveLiquidityWithPermit(&_DFedIndex.TransactOpts, _token, _liquidity, _amount0Min, _amount1Min, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x104870fd.
//
// Solidity: function removeLiquidityWithPermit(address _token, uint256 _liquidity, uint256 _amount0Min, uint256 _amount1Min, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amount0, uint256 _amount1)
func (_DFedIndex *DFedIndexTransactorSession) RemoveLiquidityWithPermit(_token common.Address, _liquidity *big.Int, _amount0Min *big.Int, _amount1Min *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.RemoveLiquidityWithPermit(&_DFedIndex.TransactOpts, _token, _liquidity, _amount0Min, _amount1Min, _to, _deadline, _approveMax, _v, _r, _s)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address _token, uint256 _debtId, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexTransactor) Repay(opts *bind.TransactOpts, _token common.Address, _debtId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "repay", _token, _debtId, _deadline)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address _token, uint256 _debtId, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexSession) Repay(_token common.Address, _debtId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.Repay(&_DFedIndex.TransactOpts, _token, _debtId, _deadline)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address _token, uint256 _debtId, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexTransactorSession) Repay(_token common.Address, _debtId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.Repay(&_DFedIndex.TransactOpts, _token, _debtId, _deadline)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xd7187ec3.
//
// Solidity: function repayWithPermit(address _token, uint256 _debtId, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedIndex *DFedIndexTransactor) RepayWithPermit(opts *bind.TransactOpts, _token common.Address, _debtId *big.Int, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "repayWithPermit", _token, _debtId, _value, _deadline, _v, _r, _s)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xd7187ec3.
//
// Solidity: function repayWithPermit(address _token, uint256 _debtId, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedIndex *DFedIndexSession) RepayWithPermit(_token common.Address, _debtId *big.Int, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.RepayWithPermit(&_DFedIndex.TransactOpts, _token, _debtId, _value, _deadline, _v, _r, _s)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xd7187ec3.
//
// Solidity: function repayWithPermit(address _token, uint256 _debtId, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_DFedIndex *DFedIndexTransactorSession) RepayWithPermit(_token common.Address, _debtId *big.Int, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _DFedIndex.Contract.RepayWithPermit(&_DFedIndex.TransactOpts, _token, _debtId, _value, _deadline, _v, _r, _s)
}

// SellToken is a paid mutator transaction binding the contract method 0xee91d954.
//
// Solidity: function sellToken(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexTransactor) SellToken(opts *bind.TransactOpts, _token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.contract.Transact(opts, "sellToken", _token, _amountIn, _amountOutMin, _to, _deadline)
}

// SellToken is a paid mutator transaction binding the contract method 0xee91d954.
//
// Solidity: function sellToken(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexSession) SellToken(_token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.SellToken(&_DFedIndex.TransactOpts, _token, _amountIn, _amountOutMin, _to, _deadline)
}

// SellToken is a paid mutator transaction binding the contract method 0xee91d954.
//
// Solidity: function sellToken(address _token, uint256 _amountIn, uint256 _amountOutMin, address _to, uint256 _deadline) returns()
func (_DFedIndex *DFedIndexTransactorSession) SellToken(_token common.Address, _amountIn *big.Int, _amountOutMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _DFedIndex.Contract.SellToken(&_DFedIndex.TransactOpts, _token, _amountIn, _amountOutMin, _to, _deadline)
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
