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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IMetronionAccessoriesAccessories is an auto generated low-level Go binding around an user-defined struct.
type IMetronionAccessoriesAccessories struct {
	Name            string
	MaxSupply       *big.Int
	Minted          *big.Int
	Burnt           *big.Int
	AccessoriesType *big.Int
	Rarity          uint8
}

// MetronionAccessoriesABI is the input ABI used to generate the binding from.
const MetronionAccessoriesABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_accessoriesTypeName\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AccessoryBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accessoriesType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIMetronionAccessories.Rarity\",\"name\":\"rarity\",\"type\":\"uint8\"}],\"name\":\"AccessoryCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AccessoryMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"AccessoryReturn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"AccessoryStore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accessoriesTypeName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"}],\"name\":\"addAccessoriesTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessoriesType\",\"type\":\"uint256\"},{\"internalType\":\"enumIMetronionAccessories.Rarity\",\"name\":\"rarity\",\"type\":\"uint8\"}],\"name\":\"createAccessory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAccessory\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessoriesType\",\"type\":\"uint256\"},{\"internalType\":\"enumIMetronionAccessories.Rarity\",\"name\":\"rarity\",\"type\":\"uint8\"}],\"internalType\":\"structIMetronionAccessories.Accessories\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAccessoryType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accessoriesType\",\"type\":\"uint256\"}],\"name\":\"getAccessoryTypeName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAccessoryType\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getEquippedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestAccessoryId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBatchAmountSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"returnAccessories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"storeAccessories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"transferEquippedAccessories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MetronionAccessories is an auto generated Go binding around an Ethereum contract.
type MetronionAccessories struct {
	MetronionAccessoriesCaller     // Read-only binding to the contract
	MetronionAccessoriesTransactor // Write-only binding to the contract
	MetronionAccessoriesFilterer   // Log filterer for contract events
}

// MetronionAccessoriesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetronionAccessoriesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionAccessoriesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetronionAccessoriesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionAccessoriesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetronionAccessoriesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionAccessoriesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetronionAccessoriesSession struct {
	Contract     *MetronionAccessories // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MetronionAccessoriesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetronionAccessoriesCallerSession struct {
	Contract *MetronionAccessoriesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MetronionAccessoriesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetronionAccessoriesTransactorSession struct {
	Contract     *MetronionAccessoriesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MetronionAccessoriesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetronionAccessoriesRaw struct {
	Contract *MetronionAccessories // Generic contract binding to access the raw methods on
}

// MetronionAccessoriesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetronionAccessoriesCallerRaw struct {
	Contract *MetronionAccessoriesCaller // Generic read-only contract binding to access the raw methods on
}

// MetronionAccessoriesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetronionAccessoriesTransactorRaw struct {
	Contract *MetronionAccessoriesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetronionAccessories creates a new instance of MetronionAccessories, bound to a specific deployed contract.
func NewMetronionAccessories(address common.Address, backend bind.ContractBackend) (*MetronionAccessories, error) {
	contract, err := bindMetronionAccessories(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessories{MetronionAccessoriesCaller: MetronionAccessoriesCaller{contract: contract}, MetronionAccessoriesTransactor: MetronionAccessoriesTransactor{contract: contract}, MetronionAccessoriesFilterer: MetronionAccessoriesFilterer{contract: contract}}, nil
}

// NewMetronionAccessoriesCaller creates a new read-only instance of MetronionAccessories, bound to a specific deployed contract.
func NewMetronionAccessoriesCaller(address common.Address, caller bind.ContractCaller) (*MetronionAccessoriesCaller, error) {
	contract, err := bindMetronionAccessories(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesCaller{contract: contract}, nil
}

// NewMetronionAccessoriesTransactor creates a new write-only instance of MetronionAccessories, bound to a specific deployed contract.
func NewMetronionAccessoriesTransactor(address common.Address, transactor bind.ContractTransactor) (*MetronionAccessoriesTransactor, error) {
	contract, err := bindMetronionAccessories(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesTransactor{contract: contract}, nil
}

// NewMetronionAccessoriesFilterer creates a new log filterer instance of MetronionAccessories, bound to a specific deployed contract.
func NewMetronionAccessoriesFilterer(address common.Address, filterer bind.ContractFilterer) (*MetronionAccessoriesFilterer, error) {
	contract, err := bindMetronionAccessories(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesFilterer{contract: contract}, nil
}

// bindMetronionAccessories binds a generic wrapper to an already deployed contract.
func bindMetronionAccessories(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetronionAccessoriesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetronionAccessories *MetronionAccessoriesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetronionAccessories.Contract.MetronionAccessoriesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetronionAccessories *MetronionAccessoriesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.MetronionAccessoriesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetronionAccessories *MetronionAccessoriesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.MetronionAccessoriesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetronionAccessories *MetronionAccessoriesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetronionAccessories.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetronionAccessories *MetronionAccessoriesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetronionAccessories *MetronionAccessoriesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.contract.Transact(opts, method, params...)
}

// AccessoriesTypeName is a free data retrieval call binding the contract method 0x819408e6.
//
// Solidity: function accessoriesTypeName(uint256 ) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesCaller) AccessoriesTypeName(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "accessoriesTypeName", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccessoriesTypeName is a free data retrieval call binding the contract method 0x819408e6.
//
// Solidity: function accessoriesTypeName(uint256 ) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesSession) AccessoriesTypeName(arg0 *big.Int) (string, error) {
	return _MetronionAccessories.Contract.AccessoriesTypeName(&_MetronionAccessories.CallOpts, arg0)
}

// AccessoriesTypeName is a free data retrieval call binding the contract method 0x819408e6.
//
// Solidity: function accessoriesTypeName(uint256 ) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) AccessoriesTypeName(arg0 *big.Int) (string, error) {
	return _MetronionAccessories.Contract.AccessoriesTypeName(&_MetronionAccessories.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 assetId) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCaller) BalanceOf(opts *bind.CallOpts, account common.Address, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "balanceOf", account, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 assetId) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesSession) BalanceOf(account common.Address, assetId *big.Int) (*big.Int, error) {
	return _MetronionAccessories.Contract.BalanceOf(&_MetronionAccessories.CallOpts, account, assetId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 assetId) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) BalanceOf(account common.Address, assetId *big.Int) (*big.Int, error) {
	return _MetronionAccessories.Contract.BalanceOf(&_MetronionAccessories.CallOpts, account, assetId)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MetronionAccessories *MetronionAccessoriesCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MetronionAccessories *MetronionAccessoriesSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MetronionAccessories.Contract.BalanceOfBatch(&_MetronionAccessories.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MetronionAccessories *MetronionAccessoriesCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MetronionAccessories.Contract.BalanceOfBatch(&_MetronionAccessories.CallOpts, accounts, ids)
}

// GetAccessory is a free data retrieval call binding the contract method 0x01637ff4.
//
// Solidity: function getAccessory(uint256 id) view returns((string,uint256,uint256,uint256,uint256,uint8))
func (_MetronionAccessories *MetronionAccessoriesCaller) GetAccessory(opts *bind.CallOpts, id *big.Int) (IMetronionAccessoriesAccessories, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "getAccessory", id)

	if err != nil {
		return *new(IMetronionAccessoriesAccessories), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetronionAccessoriesAccessories)).(*IMetronionAccessoriesAccessories)

	return out0, err

}

// GetAccessory is a free data retrieval call binding the contract method 0x01637ff4.
//
// Solidity: function getAccessory(uint256 id) view returns((string,uint256,uint256,uint256,uint256,uint8))
func (_MetronionAccessories *MetronionAccessoriesSession) GetAccessory(id *big.Int) (IMetronionAccessoriesAccessories, error) {
	return _MetronionAccessories.Contract.GetAccessory(&_MetronionAccessories.CallOpts, id)
}

// GetAccessory is a free data retrieval call binding the contract method 0x01637ff4.
//
// Solidity: function getAccessory(uint256 id) view returns((string,uint256,uint256,uint256,uint256,uint8))
func (_MetronionAccessories *MetronionAccessoriesCallerSession) GetAccessory(id *big.Int) (IMetronionAccessoriesAccessories, error) {
	return _MetronionAccessories.Contract.GetAccessory(&_MetronionAccessories.CallOpts, id)
}

// GetAccessoryType is a free data retrieval call binding the contract method 0xd640813b.
//
// Solidity: function getAccessoryType(uint256 id) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCaller) GetAccessoryType(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "getAccessoryType", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccessoryType is a free data retrieval call binding the contract method 0xd640813b.
//
// Solidity: function getAccessoryType(uint256 id) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesSession) GetAccessoryType(id *big.Int) (*big.Int, error) {
	return _MetronionAccessories.Contract.GetAccessoryType(&_MetronionAccessories.CallOpts, id)
}

// GetAccessoryType is a free data retrieval call binding the contract method 0xd640813b.
//
// Solidity: function getAccessoryType(uint256 id) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) GetAccessoryType(id *big.Int) (*big.Int, error) {
	return _MetronionAccessories.Contract.GetAccessoryType(&_MetronionAccessories.CallOpts, id)
}

// GetAccessoryTypeName is a free data retrieval call binding the contract method 0x53d135f9.
//
// Solidity: function getAccessoryTypeName(uint256 accessoriesType) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesCaller) GetAccessoryTypeName(opts *bind.CallOpts, accessoriesType *big.Int) (string, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "getAccessoryTypeName", accessoriesType)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAccessoryTypeName is a free data retrieval call binding the contract method 0x53d135f9.
//
// Solidity: function getAccessoryTypeName(uint256 accessoriesType) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesSession) GetAccessoryTypeName(accessoriesType *big.Int) (string, error) {
	return _MetronionAccessories.Contract.GetAccessoryTypeName(&_MetronionAccessories.CallOpts, accessoriesType)
}

// GetAccessoryTypeName is a free data retrieval call binding the contract method 0x53d135f9.
//
// Solidity: function getAccessoryTypeName(uint256 accessoriesType) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) GetAccessoryTypeName(accessoriesType *big.Int) (string, error) {
	return _MetronionAccessories.Contract.GetAccessoryTypeName(&_MetronionAccessories.CallOpts, accessoriesType)
}

// GetAllAccessoryType is a free data retrieval call binding the contract method 0x27e4aec4.
//
// Solidity: function getAllAccessoryType() view returns(string[])
func (_MetronionAccessories *MetronionAccessoriesCaller) GetAllAccessoryType(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "getAllAccessoryType")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllAccessoryType is a free data retrieval call binding the contract method 0x27e4aec4.
//
// Solidity: function getAllAccessoryType() view returns(string[])
func (_MetronionAccessories *MetronionAccessoriesSession) GetAllAccessoryType() ([]string, error) {
	return _MetronionAccessories.Contract.GetAllAccessoryType(&_MetronionAccessories.CallOpts)
}

// GetAllAccessoryType is a free data retrieval call binding the contract method 0x27e4aec4.
//
// Solidity: function getAllAccessoryType() view returns(string[])
func (_MetronionAccessories *MetronionAccessoriesCallerSession) GetAllAccessoryType() ([]string, error) {
	return _MetronionAccessories.Contract.GetAllAccessoryType(&_MetronionAccessories.CallOpts)
}

// GetEquippedBalance is a free data retrieval call binding the contract method 0x72ac8821.
//
// Solidity: function getEquippedBalance(address account, uint256 id) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCaller) GetEquippedBalance(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "getEquippedBalance", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEquippedBalance is a free data retrieval call binding the contract method 0x72ac8821.
//
// Solidity: function getEquippedBalance(address account, uint256 id) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesSession) GetEquippedBalance(account common.Address, id *big.Int) (*big.Int, error) {
	return _MetronionAccessories.Contract.GetEquippedBalance(&_MetronionAccessories.CallOpts, account, id)
}

// GetEquippedBalance is a free data retrieval call binding the contract method 0x72ac8821.
//
// Solidity: function getEquippedBalance(address account, uint256 id) view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) GetEquippedBalance(account common.Address, id *big.Int) (*big.Int, error) {
	return _MetronionAccessories.Contract.GetEquippedBalance(&_MetronionAccessories.CallOpts, account, id)
}

// GetLatestAccessoryId is a free data retrieval call binding the contract method 0x4b9ea7a0.
//
// Solidity: function getLatestAccessoryId() view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCaller) GetLatestAccessoryId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "getLatestAccessoryId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestAccessoryId is a free data retrieval call binding the contract method 0x4b9ea7a0.
//
// Solidity: function getLatestAccessoryId() view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesSession) GetLatestAccessoryId() (*big.Int, error) {
	return _MetronionAccessories.Contract.GetLatestAccessoryId(&_MetronionAccessories.CallOpts)
}

// GetLatestAccessoryId is a free data retrieval call binding the contract method 0x4b9ea7a0.
//
// Solidity: function getLatestAccessoryId() view returns(uint256)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) GetLatestAccessoryId() (*big.Int, error) {
	return _MetronionAccessories.Contract.GetLatestAccessoryId(&_MetronionAccessories.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetronionAccessories *MetronionAccessoriesCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetronionAccessories *MetronionAccessoriesSession) GetOperators() ([]common.Address, error) {
	return _MetronionAccessories.Contract.GetOperators(&_MetronionAccessories.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetronionAccessories *MetronionAccessoriesCallerSession) GetOperators() ([]common.Address, error) {
	return _MetronionAccessories.Contract.GetOperators(&_MetronionAccessories.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MetronionAccessories.Contract.IsApprovedForAll(&_MetronionAccessories.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MetronionAccessories.Contract.IsApprovedForAll(&_MetronionAccessories.CallOpts, owner, operator)
}

// IsBatchAmountSupported is a free data retrieval call binding the contract method 0xde271441.
//
// Solidity: function isBatchAmountSupported() pure returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCaller) IsBatchAmountSupported(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "isBatchAmountSupported")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchAmountSupported is a free data retrieval call binding the contract method 0xde271441.
//
// Solidity: function isBatchAmountSupported() pure returns(bool)
func (_MetronionAccessories *MetronionAccessoriesSession) IsBatchAmountSupported() (bool, error) {
	return _MetronionAccessories.Contract.IsBatchAmountSupported(&_MetronionAccessories.CallOpts)
}

// IsBatchAmountSupported is a free data retrieval call binding the contract method 0xde271441.
//
// Solidity: function isBatchAmountSupported() pure returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) IsBatchAmountSupported() (bool, error) {
	return _MetronionAccessories.Contract.IsBatchAmountSupported(&_MetronionAccessories.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCaller) IsOperator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "isOperator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesSession) IsOperator(account common.Address) (bool, error) {
	return _MetronionAccessories.Contract.IsOperator(&_MetronionAccessories.CallOpts, account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) IsOperator(account common.Address) (bool, error) {
	return _MetronionAccessories.Contract.IsOperator(&_MetronionAccessories.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionAccessories *MetronionAccessoriesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionAccessories *MetronionAccessoriesSession) Owner() (common.Address, error) {
	return _MetronionAccessories.Contract.Owner(&_MetronionAccessories.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) Owner() (common.Address, error) {
	return _MetronionAccessories.Contract.Owner(&_MetronionAccessories.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetronionAccessories.Contract.SupportsInterface(&_MetronionAccessories.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetronionAccessories.Contract.SupportsInterface(&_MetronionAccessories.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _MetronionAccessories.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesSession) Uri(id *big.Int) (string, error) {
	return _MetronionAccessories.Contract.Uri(&_MetronionAccessories.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_MetronionAccessories *MetronionAccessoriesCallerSession) Uri(id *big.Int) (string, error) {
	return _MetronionAccessories.Contract.Uri(&_MetronionAccessories.CallOpts, id)
}

// AddAccessoriesTypes is a paid mutator transaction binding the contract method 0x8c4088bb.
//
// Solidity: function addAccessoriesTypes(string[] names) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) AddAccessoriesTypes(opts *bind.TransactOpts, names []string) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "addAccessoriesTypes", names)
}

// AddAccessoriesTypes is a paid mutator transaction binding the contract method 0x8c4088bb.
//
// Solidity: function addAccessoriesTypes(string[] names) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) AddAccessoriesTypes(names []string) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.AddAccessoriesTypes(&_MetronionAccessories.TransactOpts, names)
}

// AddAccessoriesTypes is a paid mutator transaction binding the contract method 0x8c4088bb.
//
// Solidity: function addAccessoriesTypes(string[] names) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) AddAccessoriesTypes(names []string) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.AddAccessoriesTypes(&_MetronionAccessories.TransactOpts, names)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "addOperator", operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.AddOperator(&_MetronionAccessories.TransactOpts, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.AddOperator(&_MetronionAccessories.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 amount) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "burn", account, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 amount) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) Burn(account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.Burn(&_MetronionAccessories.TransactOpts, account, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 amount) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) Burn(account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.Burn(&_MetronionAccessories.TransactOpts, account, id, amount)
}

// CreateAccessory is a paid mutator transaction binding the contract method 0x6000cd65.
//
// Solidity: function createAccessory(string name, uint256 maxSupply, uint256 accessoriesType, uint8 rarity) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) CreateAccessory(opts *bind.TransactOpts, name string, maxSupply *big.Int, accessoriesType *big.Int, rarity uint8) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "createAccessory", name, maxSupply, accessoriesType, rarity)
}

// CreateAccessory is a paid mutator transaction binding the contract method 0x6000cd65.
//
// Solidity: function createAccessory(string name, uint256 maxSupply, uint256 accessoriesType, uint8 rarity) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) CreateAccessory(name string, maxSupply *big.Int, accessoriesType *big.Int, rarity uint8) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.CreateAccessory(&_MetronionAccessories.TransactOpts, name, maxSupply, accessoriesType, rarity)
}

// CreateAccessory is a paid mutator transaction binding the contract method 0x6000cd65.
//
// Solidity: function createAccessory(string name, uint256 maxSupply, uint256 accessoriesType, uint8 rarity) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) CreateAccessory(name string, maxSupply *big.Int, accessoriesType *big.Int, rarity uint8) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.CreateAccessory(&_MetronionAccessories.TransactOpts, name, maxSupply, accessoriesType, rarity)
}

// Mint is a paid mutator transaction binding the contract method 0xe7d3fe6b.
//
// Solidity: function mint(uint256 id, uint256 amount, address to) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) Mint(opts *bind.TransactOpts, id *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "mint", id, amount, to)
}

// Mint is a paid mutator transaction binding the contract method 0xe7d3fe6b.
//
// Solidity: function mint(uint256 id, uint256 amount, address to) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) Mint(id *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.Mint(&_MetronionAccessories.TransactOpts, id, amount, to)
}

// Mint is a paid mutator transaction binding the contract method 0xe7d3fe6b.
//
// Solidity: function mint(uint256 id, uint256 amount, address to) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) Mint(id *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.Mint(&_MetronionAccessories.TransactOpts, id, amount, to)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.RemoveOperator(&_MetronionAccessories.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.RemoveOperator(&_MetronionAccessories.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionAccessories *MetronionAccessoriesSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetronionAccessories.Contract.RenounceOwnership(&_MetronionAccessories.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetronionAccessories.Contract.RenounceOwnership(&_MetronionAccessories.TransactOpts)
}

// ReturnAccessories is a paid mutator transaction binding the contract method 0x8957a512.
//
// Solidity: function returnAccessories(address account, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) ReturnAccessories(opts *bind.TransactOpts, account common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "returnAccessories", account, accessoryIds)
}

// ReturnAccessories is a paid mutator transaction binding the contract method 0x8957a512.
//
// Solidity: function returnAccessories(address account, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) ReturnAccessories(account common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.ReturnAccessories(&_MetronionAccessories.TransactOpts, account, accessoryIds)
}

// ReturnAccessories is a paid mutator transaction binding the contract method 0x8957a512.
//
// Solidity: function returnAccessories(address account, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) ReturnAccessories(account common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.ReturnAccessories(&_MetronionAccessories.TransactOpts, account, accessoryIds)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SafeBatchTransferFrom(&_MetronionAccessories.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SafeBatchTransferFrom(&_MetronionAccessories.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 assetId, uint256 amount) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, assetId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "safeTransferFrom", from, to, assetId, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 assetId, uint256 amount) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) SafeTransferFrom(from common.Address, to common.Address, assetId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SafeTransferFrom(&_MetronionAccessories.TransactOpts, from, to, assetId, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 assetId, uint256 amount) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) SafeTransferFrom(from common.Address, to common.Address, assetId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SafeTransferFrom(&_MetronionAccessories.TransactOpts, from, to, assetId, amount)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "safeTransferFrom0", from, to, id, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SafeTransferFrom0(&_MetronionAccessories.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SafeTransferFrom0(&_MetronionAccessories.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SetApprovalForAll(&_MetronionAccessories.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SetApprovalForAll(&_MetronionAccessories.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newuri) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) SetBaseURI(opts *bind.TransactOpts, newuri string) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "setBaseURI", newuri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newuri) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) SetBaseURI(newuri string) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SetBaseURI(&_MetronionAccessories.TransactOpts, newuri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newuri) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) SetBaseURI(newuri string) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.SetBaseURI(&_MetronionAccessories.TransactOpts, newuri)
}

// StoreAccessories is a paid mutator transaction binding the contract method 0x1cea7c61.
//
// Solidity: function storeAccessories(address account, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) StoreAccessories(opts *bind.TransactOpts, account common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "storeAccessories", account, accessoryIds)
}

// StoreAccessories is a paid mutator transaction binding the contract method 0x1cea7c61.
//
// Solidity: function storeAccessories(address account, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) StoreAccessories(account common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.StoreAccessories(&_MetronionAccessories.TransactOpts, account, accessoryIds)
}

// StoreAccessories is a paid mutator transaction binding the contract method 0x1cea7c61.
//
// Solidity: function storeAccessories(address account, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) StoreAccessories(account common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.StoreAccessories(&_MetronionAccessories.TransactOpts, account, accessoryIds)
}

// TransferEquippedAccessories is a paid mutator transaction binding the contract method 0x1bf064d0.
//
// Solidity: function transferEquippedAccessories(address from, address to, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) TransferEquippedAccessories(opts *bind.TransactOpts, from common.Address, to common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "transferEquippedAccessories", from, to, accessoryIds)
}

// TransferEquippedAccessories is a paid mutator transaction binding the contract method 0x1bf064d0.
//
// Solidity: function transferEquippedAccessories(address from, address to, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) TransferEquippedAccessories(from common.Address, to common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.TransferEquippedAccessories(&_MetronionAccessories.TransactOpts, from, to, accessoryIds)
}

// TransferEquippedAccessories is a paid mutator transaction binding the contract method 0x1bf064d0.
//
// Solidity: function transferEquippedAccessories(address from, address to, uint256[] accessoryIds) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) TransferEquippedAccessories(from common.Address, to common.Address, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.TransferEquippedAccessories(&_MetronionAccessories.TransactOpts, from, to, accessoryIds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionAccessories *MetronionAccessoriesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.TransferOwnership(&_MetronionAccessories.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionAccessories *MetronionAccessoriesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetronionAccessories.Contract.TransferOwnership(&_MetronionAccessories.TransactOpts, newOwner)
}

// MetronionAccessoriesAccessoryBurntIterator is returned from FilterAccessoryBurnt and is used to iterate over the raw logs and unpacked data for AccessoryBurnt events raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryBurntIterator struct {
	Event *MetronionAccessoriesAccessoryBurnt // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesAccessoryBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesAccessoryBurnt)
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
		it.Event = new(MetronionAccessoriesAccessoryBurnt)
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
func (it *MetronionAccessoriesAccessoryBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesAccessoryBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesAccessoryBurnt represents a AccessoryBurnt event raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryBurnt struct {
	From   common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAccessoryBurnt is a free log retrieval operation binding the contract event 0xf4959fdb0f4d475a80f818bb0c31600326dae6c6c03ff67cdaef5675e90e3e21.
//
// Solidity: event AccessoryBurnt(address from, uint256 id, uint256 amount)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterAccessoryBurnt(opts *bind.FilterOpts) (*MetronionAccessoriesAccessoryBurntIterator, error) {

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "AccessoryBurnt")
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesAccessoryBurntIterator{contract: _MetronionAccessories.contract, event: "AccessoryBurnt", logs: logs, sub: sub}, nil
}

// WatchAccessoryBurnt is a free log subscription operation binding the contract event 0xf4959fdb0f4d475a80f818bb0c31600326dae6c6c03ff67cdaef5675e90e3e21.
//
// Solidity: event AccessoryBurnt(address from, uint256 id, uint256 amount)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchAccessoryBurnt(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesAccessoryBurnt) (event.Subscription, error) {

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "AccessoryBurnt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesAccessoryBurnt)
				if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryBurnt", log); err != nil {
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

// ParseAccessoryBurnt is a log parse operation binding the contract event 0xf4959fdb0f4d475a80f818bb0c31600326dae6c6c03ff67cdaef5675e90e3e21.
//
// Solidity: event AccessoryBurnt(address from, uint256 id, uint256 amount)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseAccessoryBurnt(log types.Log) (*MetronionAccessoriesAccessoryBurnt, error) {
	event := new(MetronionAccessoriesAccessoryBurnt)
	if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesAccessoryCreatedIterator is returned from FilterAccessoryCreated and is used to iterate over the raw logs and unpacked data for AccessoryCreated events raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryCreatedIterator struct {
	Event *MetronionAccessoriesAccessoryCreated // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesAccessoryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesAccessoryCreated)
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
		it.Event = new(MetronionAccessoriesAccessoryCreated)
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
func (it *MetronionAccessoriesAccessoryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesAccessoryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesAccessoryCreated represents a AccessoryCreated event raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryCreated struct {
	Id              *big.Int
	Name            string
	MaxSupply       *big.Int
	AccessoriesType *big.Int
	Rarity          uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccessoryCreated is a free log retrieval operation binding the contract event 0xc79076d67f624e5121684465a119da0f2a17dbcd6b1f187d942df263d157572a.
//
// Solidity: event AccessoryCreated(uint256 id, string name, uint256 maxSupply, uint256 accessoriesType, uint8 rarity)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterAccessoryCreated(opts *bind.FilterOpts) (*MetronionAccessoriesAccessoryCreatedIterator, error) {

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "AccessoryCreated")
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesAccessoryCreatedIterator{contract: _MetronionAccessories.contract, event: "AccessoryCreated", logs: logs, sub: sub}, nil
}

// WatchAccessoryCreated is a free log subscription operation binding the contract event 0xc79076d67f624e5121684465a119da0f2a17dbcd6b1f187d942df263d157572a.
//
// Solidity: event AccessoryCreated(uint256 id, string name, uint256 maxSupply, uint256 accessoriesType, uint8 rarity)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchAccessoryCreated(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesAccessoryCreated) (event.Subscription, error) {

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "AccessoryCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesAccessoryCreated)
				if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryCreated", log); err != nil {
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

// ParseAccessoryCreated is a log parse operation binding the contract event 0xc79076d67f624e5121684465a119da0f2a17dbcd6b1f187d942df263d157572a.
//
// Solidity: event AccessoryCreated(uint256 id, string name, uint256 maxSupply, uint256 accessoriesType, uint8 rarity)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseAccessoryCreated(log types.Log) (*MetronionAccessoriesAccessoryCreated, error) {
	event := new(MetronionAccessoriesAccessoryCreated)
	if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesAccessoryMintIterator is returned from FilterAccessoryMint and is used to iterate over the raw logs and unpacked data for AccessoryMint events raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryMintIterator struct {
	Event *MetronionAccessoriesAccessoryMint // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesAccessoryMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesAccessoryMint)
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
		it.Event = new(MetronionAccessoriesAccessoryMint)
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
func (it *MetronionAccessoriesAccessoryMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesAccessoryMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesAccessoryMint represents a AccessoryMint event raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryMint struct {
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAccessoryMint is a free log retrieval operation binding the contract event 0xcdae6c9905b08a2a5832545f5c0937498ab62deb40ee875033cc7cc1b8337c8f.
//
// Solidity: event AccessoryMint(address to, uint256 id, uint256 amount)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterAccessoryMint(opts *bind.FilterOpts) (*MetronionAccessoriesAccessoryMintIterator, error) {

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "AccessoryMint")
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesAccessoryMintIterator{contract: _MetronionAccessories.contract, event: "AccessoryMint", logs: logs, sub: sub}, nil
}

// WatchAccessoryMint is a free log subscription operation binding the contract event 0xcdae6c9905b08a2a5832545f5c0937498ab62deb40ee875033cc7cc1b8337c8f.
//
// Solidity: event AccessoryMint(address to, uint256 id, uint256 amount)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchAccessoryMint(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesAccessoryMint) (event.Subscription, error) {

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "AccessoryMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesAccessoryMint)
				if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryMint", log); err != nil {
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

// ParseAccessoryMint is a log parse operation binding the contract event 0xcdae6c9905b08a2a5832545f5c0937498ab62deb40ee875033cc7cc1b8337c8f.
//
// Solidity: event AccessoryMint(address to, uint256 id, uint256 amount)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseAccessoryMint(log types.Log) (*MetronionAccessoriesAccessoryMint, error) {
	event := new(MetronionAccessoriesAccessoryMint)
	if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesAccessoryReturnIterator is returned from FilterAccessoryReturn and is used to iterate over the raw logs and unpacked data for AccessoryReturn events raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryReturnIterator struct {
	Event *MetronionAccessoriesAccessoryReturn // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesAccessoryReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesAccessoryReturn)
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
		it.Event = new(MetronionAccessoriesAccessoryReturn)
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
func (it *MetronionAccessoriesAccessoryReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesAccessoryReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesAccessoryReturn represents a AccessoryReturn event raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryReturn struct {
	To           common.Address
	AccessoryIds []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccessoryReturn is a free log retrieval operation binding the contract event 0x49e945017a02ceb99778dc346687e218bd59643d337341ea3b08d21b37d6a5dd.
//
// Solidity: event AccessoryReturn(address to, uint256[] accessoryIds)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterAccessoryReturn(opts *bind.FilterOpts) (*MetronionAccessoriesAccessoryReturnIterator, error) {

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "AccessoryReturn")
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesAccessoryReturnIterator{contract: _MetronionAccessories.contract, event: "AccessoryReturn", logs: logs, sub: sub}, nil
}

// WatchAccessoryReturn is a free log subscription operation binding the contract event 0x49e945017a02ceb99778dc346687e218bd59643d337341ea3b08d21b37d6a5dd.
//
// Solidity: event AccessoryReturn(address to, uint256[] accessoryIds)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchAccessoryReturn(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesAccessoryReturn) (event.Subscription, error) {

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "AccessoryReturn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesAccessoryReturn)
				if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryReturn", log); err != nil {
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

// ParseAccessoryReturn is a log parse operation binding the contract event 0x49e945017a02ceb99778dc346687e218bd59643d337341ea3b08d21b37d6a5dd.
//
// Solidity: event AccessoryReturn(address to, uint256[] accessoryIds)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseAccessoryReturn(log types.Log) (*MetronionAccessoriesAccessoryReturn, error) {
	event := new(MetronionAccessoriesAccessoryReturn)
	if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesAccessoryStoreIterator is returned from FilterAccessoryStore and is used to iterate over the raw logs and unpacked data for AccessoryStore events raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryStoreIterator struct {
	Event *MetronionAccessoriesAccessoryStore // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesAccessoryStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesAccessoryStore)
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
		it.Event = new(MetronionAccessoriesAccessoryStore)
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
func (it *MetronionAccessoriesAccessoryStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesAccessoryStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesAccessoryStore represents a AccessoryStore event raised by the MetronionAccessories contract.
type MetronionAccessoriesAccessoryStore struct {
	From         common.Address
	AccessoryIds []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccessoryStore is a free log retrieval operation binding the contract event 0xd6d3869c0d65cf60b54d62ef8a802863ae55b673187a5be15ee286a16be09157.
//
// Solidity: event AccessoryStore(address from, uint256[] accessoryIds)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterAccessoryStore(opts *bind.FilterOpts) (*MetronionAccessoriesAccessoryStoreIterator, error) {

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "AccessoryStore")
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesAccessoryStoreIterator{contract: _MetronionAccessories.contract, event: "AccessoryStore", logs: logs, sub: sub}, nil
}

// WatchAccessoryStore is a free log subscription operation binding the contract event 0xd6d3869c0d65cf60b54d62ef8a802863ae55b673187a5be15ee286a16be09157.
//
// Solidity: event AccessoryStore(address from, uint256[] accessoryIds)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchAccessoryStore(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesAccessoryStore) (event.Subscription, error) {

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "AccessoryStore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesAccessoryStore)
				if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryStore", log); err != nil {
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

// ParseAccessoryStore is a log parse operation binding the contract event 0xd6d3869c0d65cf60b54d62ef8a802863ae55b673187a5be15ee286a16be09157.
//
// Solidity: event AccessoryStore(address from, uint256[] accessoryIds)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseAccessoryStore(log types.Log) (*MetronionAccessoriesAccessoryStore, error) {
	event := new(MetronionAccessoriesAccessoryStore)
	if err := _MetronionAccessories.contract.UnpackLog(event, "AccessoryStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MetronionAccessories contract.
type MetronionAccessoriesApprovalForAllIterator struct {
	Event *MetronionAccessoriesApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesApprovalForAll)
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
		it.Event = new(MetronionAccessoriesApprovalForAll)
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
func (it *MetronionAccessoriesApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesApprovalForAll represents a ApprovalForAll event raised by the MetronionAccessories contract.
type MetronionAccessoriesApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*MetronionAccessoriesApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesApprovalForAllIterator{contract: _MetronionAccessories.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesApprovalForAll)
				if err := _MetronionAccessories.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseApprovalForAll(log types.Log) (*MetronionAccessoriesApprovalForAll, error) {
	event := new(MetronionAccessoriesApprovalForAll)
	if err := _MetronionAccessories.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetronionAccessories contract.
type MetronionAccessoriesOwnershipTransferredIterator struct {
	Event *MetronionAccessoriesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesOwnershipTransferred)
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
		it.Event = new(MetronionAccessoriesOwnershipTransferred)
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
func (it *MetronionAccessoriesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesOwnershipTransferred represents a OwnershipTransferred event raised by the MetronionAccessories contract.
type MetronionAccessoriesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetronionAccessoriesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesOwnershipTransferredIterator{contract: _MetronionAccessories.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesOwnershipTransferred)
				if err := _MetronionAccessories.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseOwnershipTransferred(log types.Log) (*MetronionAccessoriesOwnershipTransferred, error) {
	event := new(MetronionAccessoriesOwnershipTransferred)
	if err := _MetronionAccessories.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the MetronionAccessories contract.
type MetronionAccessoriesTransferBatchIterator struct {
	Event *MetronionAccessoriesTransferBatch // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesTransferBatch)
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
		it.Event = new(MetronionAccessoriesTransferBatch)
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
func (it *MetronionAccessoriesTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesTransferBatch represents a TransferBatch event raised by the MetronionAccessories contract.
type MetronionAccessoriesTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MetronionAccessoriesTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesTransferBatchIterator{contract: _MetronionAccessories.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesTransferBatch)
				if err := _MetronionAccessories.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseTransferBatch(log types.Log) (*MetronionAccessoriesTransferBatch, error) {
	event := new(MetronionAccessoriesTransferBatch)
	if err := _MetronionAccessories.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the MetronionAccessories contract.
type MetronionAccessoriesTransferSingleIterator struct {
	Event *MetronionAccessoriesTransferSingle // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesTransferSingle)
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
		it.Event = new(MetronionAccessoriesTransferSingle)
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
func (it *MetronionAccessoriesTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesTransferSingle represents a TransferSingle event raised by the MetronionAccessories contract.
type MetronionAccessoriesTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MetronionAccessoriesTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesTransferSingleIterator{contract: _MetronionAccessories.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesTransferSingle)
				if err := _MetronionAccessories.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseTransferSingle(log types.Log) (*MetronionAccessoriesTransferSingle, error) {
	event := new(MetronionAccessoriesTransferSingle)
	if err := _MetronionAccessories.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionAccessoriesURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the MetronionAccessories contract.
type MetronionAccessoriesURIIterator struct {
	Event *MetronionAccessoriesURI // Event containing the contract specifics and raw log

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
func (it *MetronionAccessoriesURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionAccessoriesURI)
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
		it.Event = new(MetronionAccessoriesURI)
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
func (it *MetronionAccessoriesURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionAccessoriesURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionAccessoriesURI represents a URI event raised by the MetronionAccessories contract.
type MetronionAccessoriesURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MetronionAccessories *MetronionAccessoriesFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*MetronionAccessoriesURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetronionAccessories.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &MetronionAccessoriesURIIterator{contract: _MetronionAccessories.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MetronionAccessories *MetronionAccessoriesFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *MetronionAccessoriesURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetronionAccessories.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionAccessoriesURI)
				if err := _MetronionAccessories.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MetronionAccessories *MetronionAccessoriesFilterer) ParseURI(log types.Log) (*MetronionAccessoriesURI, error) {
	event := new(MetronionAccessoriesURI)
	if err := _MetronionAccessories.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
