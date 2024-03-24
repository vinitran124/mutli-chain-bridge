// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeRouter

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

// BridgeRouterMetaData contains all meta data concerning the BridgeRouter contract.
var BridgeRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"}],\"name\":\"CreatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"liquidityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allBridgePools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgePools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"createBridgePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getAmountTokenInPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityId\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"_rewardRate\",\"type\":\"uint256[4]\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeRouterMetaData.ABI instead.
var BridgeRouterABI = BridgeRouterMetaData.ABI

// BridgeRouter is an auto generated Go binding around an Ethereum contract.
type BridgeRouter struct {
	BridgeRouterCaller     // Read-only binding to the contract
	BridgeRouterTransactor // Write-only binding to the contract
	BridgeRouterFilterer   // Log filterer for contract events
}

// BridgeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeRouterSession struct {
	Contract     *BridgeRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeRouterCallerSession struct {
	Contract *BridgeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeRouterTransactorSession struct {
	Contract     *BridgeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRouterRaw struct {
	Contract *BridgeRouter // Generic contract binding to access the raw methods on
}

// BridgeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeRouterCallerRaw struct {
	Contract *BridgeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeRouterTransactorRaw struct {
	Contract *BridgeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeRouter creates a new instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouter(address common.Address, backend bind.ContractBackend) (*BridgeRouter, error) {
	contract, err := bindBridgeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRouter{BridgeRouterCaller: BridgeRouterCaller{contract: contract}, BridgeRouterTransactor: BridgeRouterTransactor{contract: contract}, BridgeRouterFilterer: BridgeRouterFilterer{contract: contract}}, nil
}

// NewBridgeRouterCaller creates a new read-only instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouterCaller(address common.Address, caller bind.ContractCaller) (*BridgeRouterCaller, error) {
	contract, err := bindBridgeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterCaller{contract: contract}, nil
}

// NewBridgeRouterTransactor creates a new write-only instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeRouterTransactor, error) {
	contract, err := bindBridgeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterTransactor{contract: contract}, nil
}

// NewBridgeRouterFilterer creates a new log filterer instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeRouterFilterer, error) {
	contract, err := bindBridgeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterFilterer{contract: contract}, nil
}

// bindBridgeRouter binds a generic wrapper to an already deployed contract.
func bindBridgeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRouter *BridgeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeRouter.Contract.BridgeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRouter *BridgeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRouter *BridgeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRouter *BridgeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRouter *BridgeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRouter *BridgeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRouter.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeRouter *BridgeRouterSession) ADMINROLE() ([32]byte, error) {
	return _BridgeRouter.Contract.ADMINROLE(&_BridgeRouter.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCallerSession) ADMINROLE() ([32]byte, error) {
	return _BridgeRouter.Contract.ADMINROLE(&_BridgeRouter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeRouter *BridgeRouterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeRouter.Contract.DEFAULTADMINROLE(&_BridgeRouter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeRouter.Contract.DEFAULTADMINROLE(&_BridgeRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeRouter *BridgeRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeRouter *BridgeRouterSession) WETH() (common.Address, error) {
	return _BridgeRouter.Contract.WETH(&_BridgeRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeRouter *BridgeRouterCallerSession) WETH() (common.Address, error) {
	return _BridgeRouter.Contract.WETH(&_BridgeRouter.CallOpts)
}

// AllBridgePools is a free data retrieval call binding the contract method 0x493ecb33.
//
// Solidity: function allBridgePools(uint256 ) view returns(address)
func (_BridgeRouter *BridgeRouterCaller) AllBridgePools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "allBridgePools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllBridgePools is a free data retrieval call binding the contract method 0x493ecb33.
//
// Solidity: function allBridgePools(uint256 ) view returns(address)
func (_BridgeRouter *BridgeRouterSession) AllBridgePools(arg0 *big.Int) (common.Address, error) {
	return _BridgeRouter.Contract.AllBridgePools(&_BridgeRouter.CallOpts, arg0)
}

// AllBridgePools is a free data retrieval call binding the contract method 0x493ecb33.
//
// Solidity: function allBridgePools(uint256 ) view returns(address)
func (_BridgeRouter *BridgeRouterCallerSession) AllBridgePools(arg0 *big.Int) (common.Address, error) {
	return _BridgeRouter.Contract.AllBridgePools(&_BridgeRouter.CallOpts, arg0)
}

// BridgePools is a free data retrieval call binding the contract method 0x17b7e432.
//
// Solidity: function bridgePools(address ) view returns(address)
func (_BridgeRouter *BridgeRouterCaller) BridgePools(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "bridgePools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgePools is a free data retrieval call binding the contract method 0x17b7e432.
//
// Solidity: function bridgePools(address ) view returns(address)
func (_BridgeRouter *BridgeRouterSession) BridgePools(arg0 common.Address) (common.Address, error) {
	return _BridgeRouter.Contract.BridgePools(&_BridgeRouter.CallOpts, arg0)
}

// BridgePools is a free data retrieval call binding the contract method 0x17b7e432.
//
// Solidity: function bridgePools(address ) view returns(address)
func (_BridgeRouter *BridgeRouterCallerSession) BridgePools(arg0 common.Address) (common.Address, error) {
	return _BridgeRouter.Contract.BridgePools(&_BridgeRouter.CallOpts, arg0)
}

// GetAmountTokenInPool is a free data retrieval call binding the contract method 0xb20789f9.
//
// Solidity: function getAmountTokenInPool(address _token) view returns(uint256)
func (_BridgeRouter *BridgeRouterCaller) GetAmountTokenInPool(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "getAmountTokenInPool", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountTokenInPool is a free data retrieval call binding the contract method 0xb20789f9.
//
// Solidity: function getAmountTokenInPool(address _token) view returns(uint256)
func (_BridgeRouter *BridgeRouterSession) GetAmountTokenInPool(_token common.Address) (*big.Int, error) {
	return _BridgeRouter.Contract.GetAmountTokenInPool(&_BridgeRouter.CallOpts, _token)
}

// GetAmountTokenInPool is a free data retrieval call binding the contract method 0xb20789f9.
//
// Solidity: function getAmountTokenInPool(address _token) view returns(uint256)
func (_BridgeRouter *BridgeRouterCallerSession) GetAmountTokenInPool(_token common.Address) (*big.Int, error) {
	return _BridgeRouter.Contract.GetAmountTokenInPool(&_BridgeRouter.CallOpts, _token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeRouter *BridgeRouterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeRouter *BridgeRouterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeRouter.Contract.GetRoleAdmin(&_BridgeRouter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeRouter *BridgeRouterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeRouter.Contract.GetRoleAdmin(&_BridgeRouter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeRouter *BridgeRouterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeRouter *BridgeRouterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeRouter.Contract.HasRole(&_BridgeRouter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeRouter *BridgeRouterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeRouter.Contract.HasRole(&_BridgeRouter.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeRouter *BridgeRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeRouter *BridgeRouterSession) Owner() (common.Address, error) {
	return _BridgeRouter.Contract.Owner(&_BridgeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeRouter *BridgeRouterCallerSession) Owner() (common.Address, error) {
	return _BridgeRouter.Contract.Owner(&_BridgeRouter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeRouter *BridgeRouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeRouter *BridgeRouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeRouter.Contract.SupportsInterface(&_BridgeRouter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeRouter *BridgeRouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeRouter.Contract.SupportsInterface(&_BridgeRouter.CallOpts, interfaceId)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) payable returns()
func (_BridgeRouter *BridgeRouterTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "addLiquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) payable returns()
func (_BridgeRouter *BridgeRouterSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.AddLiquidity(&_BridgeRouter.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) payable returns()
func (_BridgeRouter *BridgeRouterTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.AddLiquidity(&_BridgeRouter.TransactOpts, _token, _amount)
}

// CreateBridgePool is a paid mutator transaction binding the contract method 0x36f682c8.
//
// Solidity: function createBridgePool(address _token) returns()
func (_BridgeRouter *BridgeRouterTransactor) CreateBridgePool(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "createBridgePool", _token)
}

// CreateBridgePool is a paid mutator transaction binding the contract method 0x36f682c8.
//
// Solidity: function createBridgePool(address _token) returns()
func (_BridgeRouter *BridgeRouterSession) CreateBridgePool(_token common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.CreateBridgePool(&_BridgeRouter.TransactOpts, _token)
}

// CreateBridgePool is a paid mutator transaction binding the contract method 0x36f682c8.
//
// Solidity: function createBridgePool(address _token) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) CreateBridgePool(_token common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.CreateBridgePool(&_BridgeRouter.TransactOpts, _token)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) payable returns()
func (_BridgeRouter *BridgeRouterTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) payable returns()
func (_BridgeRouter *BridgeRouterSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.Deposit(&_BridgeRouter.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) payable returns()
func (_BridgeRouter *BridgeRouterTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.Deposit(&_BridgeRouter.TransactOpts, _token, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.GrantRole(&_BridgeRouter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.GrantRole(&_BridgeRouter.TransactOpts, role, account)
}

// GrantRole0 is a paid mutator transaction binding the contract method 0x38977686.
//
// Solidity: function grantRole(address _admin) returns()
func (_BridgeRouter *BridgeRouterTransactor) GrantRole0(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "grantRole0", _admin)
}

// GrantRole0 is a paid mutator transaction binding the contract method 0x38977686.
//
// Solidity: function grantRole(address _admin) returns()
func (_BridgeRouter *BridgeRouterSession) GrantRole0(_admin common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.GrantRole0(&_BridgeRouter.TransactOpts, _admin)
}

// GrantRole0 is a paid mutator transaction binding the contract method 0x38977686.
//
// Solidity: function grantRole(address _admin) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) GrantRole0(_admin common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.GrantRole0(&_BridgeRouter.TransactOpts, _admin)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe38192e3.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount, uint256 _liquidityId) returns()
func (_BridgeRouter *BridgeRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _liquidityId *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "removeLiquidity", _token, _amount, _liquidityId)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe38192e3.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount, uint256 _liquidityId) returns()
func (_BridgeRouter *BridgeRouterSession) RemoveLiquidity(_token common.Address, _amount *big.Int, _liquidityId *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RemoveLiquidity(&_BridgeRouter.TransactOpts, _token, _amount, _liquidityId)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe38192e3.
//
// Solidity: function removeLiquidity(address _token, uint256 _amount, uint256 _liquidityId) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) RemoveLiquidity(_token common.Address, _amount *big.Int, _liquidityId *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RemoveLiquidity(&_BridgeRouter.TransactOpts, _token, _amount, _liquidityId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeRouter *BridgeRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeRouter *BridgeRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeRouter.Contract.RenounceOwnership(&_BridgeRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeRouter *BridgeRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeRouter.Contract.RenounceOwnership(&_BridgeRouter.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RenounceRole(&_BridgeRouter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RenounceRole(&_BridgeRouter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x80e52e3f.
//
// Solidity: function revokeRole(address _admin) returns()
func (_BridgeRouter *BridgeRouterTransactor) RevokeRole(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "revokeRole", _admin)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x80e52e3f.
//
// Solidity: function revokeRole(address _admin) returns()
func (_BridgeRouter *BridgeRouterSession) RevokeRole(_admin common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RevokeRole(&_BridgeRouter.TransactOpts, _admin)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x80e52e3f.
//
// Solidity: function revokeRole(address _admin) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) RevokeRole(_admin common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RevokeRole(&_BridgeRouter.TransactOpts, _admin)
}

// RevokeRole0 is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterTransactor) RevokeRole0(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "revokeRole0", role, account)
}

// RevokeRole0 is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterSession) RevokeRole0(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RevokeRole0(&_BridgeRouter.TransactOpts, role, account)
}

// RevokeRole0 is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) RevokeRole0(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.RevokeRole0(&_BridgeRouter.TransactOpts, role, account)
}

// SetRate is a paid mutator transaction binding the contract method 0x5b76d18d.
//
// Solidity: function setRate(address _token, uint256[4] _rewardRate) returns()
func (_BridgeRouter *BridgeRouterTransactor) SetRate(opts *bind.TransactOpts, _token common.Address, _rewardRate [4]*big.Int) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "setRate", _token, _rewardRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x5b76d18d.
//
// Solidity: function setRate(address _token, uint256[4] _rewardRate) returns()
func (_BridgeRouter *BridgeRouterSession) SetRate(_token common.Address, _rewardRate [4]*big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.SetRate(&_BridgeRouter.TransactOpts, _token, _rewardRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x5b76d18d.
//
// Solidity: function setRate(address _token, uint256[4] _rewardRate) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) SetRate(_token common.Address, _rewardRate [4]*big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.SetRate(&_BridgeRouter.TransactOpts, _token, _rewardRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeRouter *BridgeRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeRouter *BridgeRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.TransferOwnership(&_BridgeRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.TransferOwnership(&_BridgeRouter.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _user, uint256 _amount) returns()
func (_BridgeRouter *BridgeRouterTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "withdraw", _token, _user, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _user, uint256 _amount) returns()
func (_BridgeRouter *BridgeRouterSession) Withdraw(_token common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.Withdraw(&_BridgeRouter.TransactOpts, _token, _user, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _user, uint256 _amount) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) Withdraw(_token common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeRouter.Contract.Withdraw(&_BridgeRouter.TransactOpts, _token, _user, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeRouter *BridgeRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeRouter *BridgeRouterSession) Receive() (*types.Transaction, error) {
	return _BridgeRouter.Contract.Receive(&_BridgeRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeRouter *BridgeRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeRouter.Contract.Receive(&_BridgeRouter.TransactOpts)
}

// BridgeRouterAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the BridgeRouter contract.
type BridgeRouterAddLiquidityIterator struct {
	Event *BridgeRouterAddLiquidity // Event containing the contract specifics and raw log

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
func (it *BridgeRouterAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterAddLiquidity)
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
		it.Event = new(BridgeRouterAddLiquidity)
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
func (it *BridgeRouterAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterAddLiquidity represents a AddLiquidity event raised by the BridgeRouter contract.
type BridgeRouterAddLiquidity struct {
	Sender common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x668256213e6a9a0247adc238fcbf44cc6b98921642fca93479c5dc3873660837.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) FilterAddLiquidity(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*BridgeRouterAddLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "AddLiquidity", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterAddLiquidityIterator{contract: _BridgeRouter.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x668256213e6a9a0247adc238fcbf44cc6b98921642fca93479c5dc3873660837.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *BridgeRouterAddLiquidity, sender []common.Address, token []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "AddLiquidity", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterAddLiquidity)
				if err := _BridgeRouter.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x668256213e6a9a0247adc238fcbf44cc6b98921642fca93479c5dc3873660837.
//
// Solidity: event AddLiquidity(address indexed sender, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) ParseAddLiquidity(log types.Log) (*BridgeRouterAddLiquidity, error) {
	event := new(BridgeRouterAddLiquidity)
	if err := _BridgeRouter.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterCreatePoolIterator is returned from FilterCreatePool and is used to iterate over the raw logs and unpacked data for CreatePool events raised by the BridgeRouter contract.
type BridgeRouterCreatePoolIterator struct {
	Event *BridgeRouterCreatePool // Event containing the contract specifics and raw log

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
func (it *BridgeRouterCreatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterCreatePool)
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
		it.Event = new(BridgeRouterCreatePool)
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
func (it *BridgeRouterCreatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterCreatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterCreatePool represents a CreatePool event raised by the BridgeRouter contract.
type BridgeRouterCreatePool struct {
	Token    common.Address
	PoolAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreatePool is a free log retrieval operation binding the contract event 0xdbd2a1ea6808362e6adbec4db4969cbc11e3b0b28fb6c74cb342defaaf1daada.
//
// Solidity: event CreatePool(address indexed token, address indexed poolAddr)
func (_BridgeRouter *BridgeRouterFilterer) FilterCreatePool(opts *bind.FilterOpts, token []common.Address, poolAddr []common.Address) (*BridgeRouterCreatePoolIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "CreatePool", tokenRule, poolAddrRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterCreatePoolIterator{contract: _BridgeRouter.contract, event: "CreatePool", logs: logs, sub: sub}, nil
}

// WatchCreatePool is a free log subscription operation binding the contract event 0xdbd2a1ea6808362e6adbec4db4969cbc11e3b0b28fb6c74cb342defaaf1daada.
//
// Solidity: event CreatePool(address indexed token, address indexed poolAddr)
func (_BridgeRouter *BridgeRouterFilterer) WatchCreatePool(opts *bind.WatchOpts, sink chan<- *BridgeRouterCreatePool, token []common.Address, poolAddr []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "CreatePool", tokenRule, poolAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterCreatePool)
				if err := _BridgeRouter.contract.UnpackLog(event, "CreatePool", log); err != nil {
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

// ParseCreatePool is a log parse operation binding the contract event 0xdbd2a1ea6808362e6adbec4db4969cbc11e3b0b28fb6c74cb342defaaf1daada.
//
// Solidity: event CreatePool(address indexed token, address indexed poolAddr)
func (_BridgeRouter *BridgeRouterFilterer) ParseCreatePool(log types.Log) (*BridgeRouterCreatePool, error) {
	event := new(BridgeRouterCreatePool)
	if err := _BridgeRouter.contract.UnpackLog(event, "CreatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BridgeRouter contract.
type BridgeRouterDepositIterator struct {
	Event *BridgeRouterDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeRouterDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterDeposit)
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
		it.Event = new(BridgeRouterDeposit)
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
func (it *BridgeRouterDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterDeposit represents a Deposit event raised by the BridgeRouter contract.
type BridgeRouterDeposit struct {
	Sender common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed sender, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, token []common.Address) (*BridgeRouterDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "Deposit", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterDepositIterator{contract: _BridgeRouter.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed sender, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeRouterDeposit, sender []common.Address, token []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "Deposit", senderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterDeposit)
				if err := _BridgeRouter.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed sender, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) ParseDeposit(log types.Log) (*BridgeRouterDeposit, error) {
	event := new(BridgeRouterDeposit)
	if err := _BridgeRouter.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeRouter contract.
type BridgeRouterOwnershipTransferredIterator struct {
	Event *BridgeRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterOwnershipTransferred)
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
		it.Event = new(BridgeRouterOwnershipTransferred)
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
func (it *BridgeRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeRouter contract.
type BridgeRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeRouter *BridgeRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterOwnershipTransferredIterator{contract: _BridgeRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeRouter *BridgeRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterOwnershipTransferred)
				if err := _BridgeRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeRouter *BridgeRouterFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeRouterOwnershipTransferred, error) {
	event := new(BridgeRouterOwnershipTransferred)
	if err := _BridgeRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the BridgeRouter contract.
type BridgeRouterRemoveLiquidityIterator struct {
	Event *BridgeRouterRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *BridgeRouterRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterRemoveLiquidity)
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
		it.Event = new(BridgeRouterRemoveLiquidity)
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
func (it *BridgeRouterRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterRemoveLiquidity represents a RemoveLiquidity event raised by the BridgeRouter contract.
type BridgeRouterRemoveLiquidity struct {
	Sender      common.Address
	Token       common.Address
	LiquidityId *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x34ef8e86237e7385b43618862e895c6ce827b2b7d6107ad415d54336c1dd2dd6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed token, uint256 indexed liquidityId, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, sender []common.Address, token []common.Address, liquidityId []*big.Int) (*BridgeRouterRemoveLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var liquidityIdRule []interface{}
	for _, liquidityIdItem := range liquidityId {
		liquidityIdRule = append(liquidityIdRule, liquidityIdItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "RemoveLiquidity", senderRule, tokenRule, liquidityIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterRemoveLiquidityIterator{contract: _BridgeRouter.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x34ef8e86237e7385b43618862e895c6ce827b2b7d6107ad415d54336c1dd2dd6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed token, uint256 indexed liquidityId, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *BridgeRouterRemoveLiquidity, sender []common.Address, token []common.Address, liquidityId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var liquidityIdRule []interface{}
	for _, liquidityIdItem := range liquidityId {
		liquidityIdRule = append(liquidityIdRule, liquidityIdItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "RemoveLiquidity", senderRule, tokenRule, liquidityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterRemoveLiquidity)
				if err := _BridgeRouter.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x34ef8e86237e7385b43618862e895c6ce827b2b7d6107ad415d54336c1dd2dd6.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed token, uint256 indexed liquidityId, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) ParseRemoveLiquidity(log types.Log) (*BridgeRouterRemoveLiquidity, error) {
	event := new(BridgeRouterRemoveLiquidity)
	if err := _BridgeRouter.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeRouter contract.
type BridgeRouterRoleAdminChangedIterator struct {
	Event *BridgeRouterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeRouterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterRoleAdminChanged)
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
		it.Event = new(BridgeRouterRoleAdminChanged)
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
func (it *BridgeRouterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeRouter contract.
type BridgeRouterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeRouter *BridgeRouterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeRouterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterRoleAdminChangedIterator{contract: _BridgeRouter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeRouter *BridgeRouterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeRouterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterRoleAdminChanged)
				if err := _BridgeRouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeRouter *BridgeRouterFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeRouterRoleAdminChanged, error) {
	event := new(BridgeRouterRoleAdminChanged)
	if err := _BridgeRouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeRouter contract.
type BridgeRouterRoleGrantedIterator struct {
	Event *BridgeRouterRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeRouterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterRoleGranted)
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
		it.Event = new(BridgeRouterRoleGranted)
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
func (it *BridgeRouterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterRoleGranted represents a RoleGranted event raised by the BridgeRouter contract.
type BridgeRouterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeRouter *BridgeRouterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRouterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterRoleGrantedIterator{contract: _BridgeRouter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeRouter *BridgeRouterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRouterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterRoleGranted)
				if err := _BridgeRouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeRouter *BridgeRouterFilterer) ParseRoleGranted(log types.Log) (*BridgeRouterRoleGranted, error) {
	event := new(BridgeRouterRoleGranted)
	if err := _BridgeRouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeRouter contract.
type BridgeRouterRoleRevokedIterator struct {
	Event *BridgeRouterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeRouterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterRoleRevoked)
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
		it.Event = new(BridgeRouterRoleRevoked)
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
func (it *BridgeRouterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterRoleRevoked represents a RoleRevoked event raised by the BridgeRouter contract.
type BridgeRouterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeRouter *BridgeRouterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRouterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterRoleRevokedIterator{contract: _BridgeRouter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeRouter *BridgeRouterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRouterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterRoleRevoked)
				if err := _BridgeRouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeRouter *BridgeRouterFilterer) ParseRoleRevoked(log types.Log) (*BridgeRouterRoleRevoked, error) {
	event := new(BridgeRouterRoleRevoked)
	if err := _BridgeRouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BridgeRouter contract.
type BridgeRouterWithdrawIterator struct {
	Event *BridgeRouterWithdraw // Event containing the contract specifics and raw log

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
func (it *BridgeRouterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouterWithdraw)
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
		it.Event = new(BridgeRouterWithdraw)
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
func (it *BridgeRouterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouterWithdraw represents a Withdraw event raised by the BridgeRouter contract.
type BridgeRouterWithdraw struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*BridgeRouterWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeRouter.contract.FilterLogs(opts, "Withdraw", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterWithdrawIterator{contract: _BridgeRouter.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeRouterWithdraw, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeRouter.contract.WatchLogs(opts, "Withdraw", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouterWithdraw)
				if err := _BridgeRouter.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed token, uint256 amount)
func (_BridgeRouter *BridgeRouterFilterer) ParseWithdraw(log types.Log) (*BridgeRouterWithdraw, error) {
	event := new(BridgeRouterWithdraw)
	if err := _BridgeRouter.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
