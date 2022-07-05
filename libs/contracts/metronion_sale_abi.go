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

// IMetronionSaleSaleConfig is an auto generated low-level Go binding around an user-defined struct.
type IMetronionSaleSaleConfig struct {
	PrivateTime uint64
	PublicTime  uint64
	EndTime     uint64
}

// IMetronionSaleSaleRecord is an auto generated low-level Go binding around an user-defined struct.
type IMetronionSaleSaleRecord struct {
	TotalSold   uint64
	PrivateSold uint64
	PublicSold  uint64
	OwnerBought uint64
}

// IMetronionSaleUserRecord is an auto generated low-level Go binding around an user-defined struct.
type IMetronionSaleUserRecord struct {
	PrivateBought uint64
	PublicBought  uint64
}

// MetronionSaleABI is the input ABI used to generate the binding from.
const MetronionSaleABI = "[{\"inputs\":[{\"internalType\":\"contractIMetronionNFT\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_versionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxWhitelistSize\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_privateTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_publicTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_endTime\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWeiPaid\",\"type\":\"uint256\"}],\"name\":\"OwnerBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWeiPaid\",\"type\":\"uint256\"}],\"name\":\"PrivateBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWeiPaid\",\"type\":\"uint256\"}],\"name\":\"PublicBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceiveETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"name\":\"UpdateWhitelistedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CAP_OWNER_INITIAL_MINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAP_PER_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAP_PER_PRIVATE_ADDRESS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SALE_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countWhitelistedGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSaleConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"privateTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"internalType\":\"structIMetronionSale.SaleConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSaleRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"totalSold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"privateSold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicSold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ownerBought\",\"type\":\"uint64\"}],\"internalType\":\"structIMetronionSale.SaleRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"privateBought\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicBought\",\"type\":\"uint64\"}],\"internalType\":\"structIMetronionSale.UserRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedGroup\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistedAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxWhitelistSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"contractIMetronionNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"name\":\"updateWhitelistedGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// MetronionSale is an auto generated Go binding around an Ethereum contract.
type MetronionSale struct {
	MetronionSaleCaller     // Read-only binding to the contract
	MetronionSaleTransactor // Write-only binding to the contract
	MetronionSaleFilterer   // Log filterer for contract events
}

// MetronionSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetronionSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetronionSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetronionSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetronionSaleSession struct {
	Contract     *MetronionSale    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetronionSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetronionSaleCallerSession struct {
	Contract *MetronionSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MetronionSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetronionSaleTransactorSession struct {
	Contract     *MetronionSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MetronionSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetronionSaleRaw struct {
	Contract *MetronionSale // Generic contract binding to access the raw methods on
}

// MetronionSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetronionSaleCallerRaw struct {
	Contract *MetronionSaleCaller // Generic read-only contract binding to access the raw methods on
}

// MetronionSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetronionSaleTransactorRaw struct {
	Contract *MetronionSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetronionSale creates a new instance of MetronionSale, bound to a specific deployed contract.
func NewMetronionSale(address common.Address, backend bind.ContractBackend) (*MetronionSale, error) {
	contract, err := bindMetronionSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetronionSale{MetronionSaleCaller: MetronionSaleCaller{contract: contract}, MetronionSaleTransactor: MetronionSaleTransactor{contract: contract}, MetronionSaleFilterer: MetronionSaleFilterer{contract: contract}}, nil
}

// NewMetronionSaleCaller creates a new read-only instance of MetronionSale, bound to a specific deployed contract.
func NewMetronionSaleCaller(address common.Address, caller bind.ContractCaller) (*MetronionSaleCaller, error) {
	contract, err := bindMetronionSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetronionSaleCaller{contract: contract}, nil
}

// NewMetronionSaleTransactor creates a new write-only instance of MetronionSale, bound to a specific deployed contract.
func NewMetronionSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*MetronionSaleTransactor, error) {
	contract, err := bindMetronionSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetronionSaleTransactor{contract: contract}, nil
}

// NewMetronionSaleFilterer creates a new log filterer instance of MetronionSale, bound to a specific deployed contract.
func NewMetronionSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*MetronionSaleFilterer, error) {
	contract, err := bindMetronionSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetronionSaleFilterer{contract: contract}, nil
}

// bindMetronionSale binds a generic wrapper to an already deployed contract.
func bindMetronionSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetronionSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetronionSale *MetronionSaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetronionSale.Contract.MetronionSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetronionSale *MetronionSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionSale.Contract.MetronionSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetronionSale *MetronionSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetronionSale.Contract.MetronionSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetronionSale *MetronionSaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetronionSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetronionSale *MetronionSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetronionSale *MetronionSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetronionSale.Contract.contract.Transact(opts, method, params...)
}

// CAPOWNERINITIALMINT is a free data retrieval call binding the contract method 0xe7dcfc2b.
//
// Solidity: function CAP_OWNER_INITIAL_MINT() view returns(uint256)
func (_MetronionSale *MetronionSaleCaller) CAPOWNERINITIALMINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "CAP_OWNER_INITIAL_MINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CAPOWNERINITIALMINT is a free data retrieval call binding the contract method 0xe7dcfc2b.
//
// Solidity: function CAP_OWNER_INITIAL_MINT() view returns(uint256)
func (_MetronionSale *MetronionSaleSession) CAPOWNERINITIALMINT() (*big.Int, error) {
	return _MetronionSale.Contract.CAPOWNERINITIALMINT(&_MetronionSale.CallOpts)
}

// CAPOWNERINITIALMINT is a free data retrieval call binding the contract method 0xe7dcfc2b.
//
// Solidity: function CAP_OWNER_INITIAL_MINT() view returns(uint256)
func (_MetronionSale *MetronionSaleCallerSession) CAPOWNERINITIALMINT() (*big.Int, error) {
	return _MetronionSale.Contract.CAPOWNERINITIALMINT(&_MetronionSale.CallOpts)
}

// CAPPERADDRESS is a free data retrieval call binding the contract method 0xb13b76d7.
//
// Solidity: function CAP_PER_ADDRESS() view returns(uint256)
func (_MetronionSale *MetronionSaleCaller) CAPPERADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "CAP_PER_ADDRESS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CAPPERADDRESS is a free data retrieval call binding the contract method 0xb13b76d7.
//
// Solidity: function CAP_PER_ADDRESS() view returns(uint256)
func (_MetronionSale *MetronionSaleSession) CAPPERADDRESS() (*big.Int, error) {
	return _MetronionSale.Contract.CAPPERADDRESS(&_MetronionSale.CallOpts)
}

// CAPPERADDRESS is a free data retrieval call binding the contract method 0xb13b76d7.
//
// Solidity: function CAP_PER_ADDRESS() view returns(uint256)
func (_MetronionSale *MetronionSaleCallerSession) CAPPERADDRESS() (*big.Int, error) {
	return _MetronionSale.Contract.CAPPERADDRESS(&_MetronionSale.CallOpts)
}

// CAPPERPRIVATEADDRESS is a free data retrieval call binding the contract method 0xb579ccf0.
//
// Solidity: function CAP_PER_PRIVATE_ADDRESS() view returns(uint256)
func (_MetronionSale *MetronionSaleCaller) CAPPERPRIVATEADDRESS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "CAP_PER_PRIVATE_ADDRESS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CAPPERPRIVATEADDRESS is a free data retrieval call binding the contract method 0xb579ccf0.
//
// Solidity: function CAP_PER_PRIVATE_ADDRESS() view returns(uint256)
func (_MetronionSale *MetronionSaleSession) CAPPERPRIVATEADDRESS() (*big.Int, error) {
	return _MetronionSale.Contract.CAPPERPRIVATEADDRESS(&_MetronionSale.CallOpts)
}

// CAPPERPRIVATEADDRESS is a free data retrieval call binding the contract method 0xb579ccf0.
//
// Solidity: function CAP_PER_PRIVATE_ADDRESS() view returns(uint256)
func (_MetronionSale *MetronionSaleCallerSession) CAPPERPRIVATEADDRESS() (*big.Int, error) {
	return _MetronionSale.Contract.CAPPERPRIVATEADDRESS(&_MetronionSale.CallOpts)
}

// SALEPRICE is a free data retrieval call binding the contract method 0x7f205a74.
//
// Solidity: function SALE_PRICE() view returns(uint256)
func (_MetronionSale *MetronionSaleCaller) SALEPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "SALE_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SALEPRICE is a free data retrieval call binding the contract method 0x7f205a74.
//
// Solidity: function SALE_PRICE() view returns(uint256)
func (_MetronionSale *MetronionSaleSession) SALEPRICE() (*big.Int, error) {
	return _MetronionSale.Contract.SALEPRICE(&_MetronionSale.CallOpts)
}

// SALEPRICE is a free data retrieval call binding the contract method 0x7f205a74.
//
// Solidity: function SALE_PRICE() view returns(uint256)
func (_MetronionSale *MetronionSaleCallerSession) SALEPRICE() (*big.Int, error) {
	return _MetronionSale.Contract.SALEPRICE(&_MetronionSale.CallOpts)
}

// CountWhitelistedGroup is a free data retrieval call binding the contract method 0x80660bcf.
//
// Solidity: function countWhitelistedGroup() view returns(uint256)
func (_MetronionSale *MetronionSaleCaller) CountWhitelistedGroup(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "countWhitelistedGroup")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountWhitelistedGroup is a free data retrieval call binding the contract method 0x80660bcf.
//
// Solidity: function countWhitelistedGroup() view returns(uint256)
func (_MetronionSale *MetronionSaleSession) CountWhitelistedGroup() (*big.Int, error) {
	return _MetronionSale.Contract.CountWhitelistedGroup(&_MetronionSale.CallOpts)
}

// CountWhitelistedGroup is a free data retrieval call binding the contract method 0x80660bcf.
//
// Solidity: function countWhitelistedGroup() view returns(uint256)
func (_MetronionSale *MetronionSaleCallerSession) CountWhitelistedGroup() (*big.Int, error) {
	return _MetronionSale.Contract.CountWhitelistedGroup(&_MetronionSale.CallOpts)
}

// GetSaleConfig is a free data retrieval call binding the contract method 0xcea943ee.
//
// Solidity: function getSaleConfig() view returns((uint64,uint64,uint64))
func (_MetronionSale *MetronionSaleCaller) GetSaleConfig(opts *bind.CallOpts) (IMetronionSaleSaleConfig, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "getSaleConfig")

	if err != nil {
		return *new(IMetronionSaleSaleConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetronionSaleSaleConfig)).(*IMetronionSaleSaleConfig)

	return out0, err

}

// GetSaleConfig is a free data retrieval call binding the contract method 0xcea943ee.
//
// Solidity: function getSaleConfig() view returns((uint64,uint64,uint64))
func (_MetronionSale *MetronionSaleSession) GetSaleConfig() (IMetronionSaleSaleConfig, error) {
	return _MetronionSale.Contract.GetSaleConfig(&_MetronionSale.CallOpts)
}

// GetSaleConfig is a free data retrieval call binding the contract method 0xcea943ee.
//
// Solidity: function getSaleConfig() view returns((uint64,uint64,uint64))
func (_MetronionSale *MetronionSaleCallerSession) GetSaleConfig() (IMetronionSaleSaleConfig, error) {
	return _MetronionSale.Contract.GetSaleConfig(&_MetronionSale.CallOpts)
}

// GetSaleRecord is a free data retrieval call binding the contract method 0xd2c99a65.
//
// Solidity: function getSaleRecord() view returns((uint64,uint64,uint64,uint64))
func (_MetronionSale *MetronionSaleCaller) GetSaleRecord(opts *bind.CallOpts) (IMetronionSaleSaleRecord, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "getSaleRecord")

	if err != nil {
		return *new(IMetronionSaleSaleRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetronionSaleSaleRecord)).(*IMetronionSaleSaleRecord)

	return out0, err

}

// GetSaleRecord is a free data retrieval call binding the contract method 0xd2c99a65.
//
// Solidity: function getSaleRecord() view returns((uint64,uint64,uint64,uint64))
func (_MetronionSale *MetronionSaleSession) GetSaleRecord() (IMetronionSaleSaleRecord, error) {
	return _MetronionSale.Contract.GetSaleRecord(&_MetronionSale.CallOpts)
}

// GetSaleRecord is a free data retrieval call binding the contract method 0xd2c99a65.
//
// Solidity: function getSaleRecord() view returns((uint64,uint64,uint64,uint64))
func (_MetronionSale *MetronionSaleCallerSession) GetSaleRecord() (IMetronionSaleSaleRecord, error) {
	return _MetronionSale.Contract.GetSaleRecord(&_MetronionSale.CallOpts)
}

// GetUserRecord is a free data retrieval call binding the contract method 0x67e2efc6.
//
// Solidity: function getUserRecord(address account) view returns((uint64,uint64))
func (_MetronionSale *MetronionSaleCaller) GetUserRecord(opts *bind.CallOpts, account common.Address) (IMetronionSaleUserRecord, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "getUserRecord", account)

	if err != nil {
		return *new(IMetronionSaleUserRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetronionSaleUserRecord)).(*IMetronionSaleUserRecord)

	return out0, err

}

// GetUserRecord is a free data retrieval call binding the contract method 0x67e2efc6.
//
// Solidity: function getUserRecord(address account) view returns((uint64,uint64))
func (_MetronionSale *MetronionSaleSession) GetUserRecord(account common.Address) (IMetronionSaleUserRecord, error) {
	return _MetronionSale.Contract.GetUserRecord(&_MetronionSale.CallOpts, account)
}

// GetUserRecord is a free data retrieval call binding the contract method 0x67e2efc6.
//
// Solidity: function getUserRecord(address account) view returns((uint64,uint64))
func (_MetronionSale *MetronionSaleCallerSession) GetUserRecord(account common.Address) (IMetronionSaleUserRecord, error) {
	return _MetronionSale.Contract.GetUserRecord(&_MetronionSale.CallOpts, account)
}

// GetWhitelistedGroup is a free data retrieval call binding the contract method 0x297d9924.
//
// Solidity: function getWhitelistedGroup() view returns(address[])
func (_MetronionSale *MetronionSaleCaller) GetWhitelistedGroup(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "getWhitelistedGroup")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedGroup is a free data retrieval call binding the contract method 0x297d9924.
//
// Solidity: function getWhitelistedGroup() view returns(address[])
func (_MetronionSale *MetronionSaleSession) GetWhitelistedGroup() ([]common.Address, error) {
	return _MetronionSale.Contract.GetWhitelistedGroup(&_MetronionSale.CallOpts)
}

// GetWhitelistedGroup is a free data retrieval call binding the contract method 0x297d9924.
//
// Solidity: function getWhitelistedGroup() view returns(address[])
func (_MetronionSale *MetronionSaleCallerSession) GetWhitelistedGroup() ([]common.Address, error) {
	return _MetronionSale.Contract.GetWhitelistedGroup(&_MetronionSale.CallOpts)
}

// IsWhitelistedAddress is a free data retrieval call binding the contract method 0x5fae0576.
//
// Solidity: function isWhitelistedAddress(address account) view returns(bool)
func (_MetronionSale *MetronionSaleCaller) IsWhitelistedAddress(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "isWhitelistedAddress", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedAddress is a free data retrieval call binding the contract method 0x5fae0576.
//
// Solidity: function isWhitelistedAddress(address account) view returns(bool)
func (_MetronionSale *MetronionSaleSession) IsWhitelistedAddress(account common.Address) (bool, error) {
	return _MetronionSale.Contract.IsWhitelistedAddress(&_MetronionSale.CallOpts, account)
}

// IsWhitelistedAddress is a free data retrieval call binding the contract method 0x5fae0576.
//
// Solidity: function isWhitelistedAddress(address account) view returns(bool)
func (_MetronionSale *MetronionSaleCallerSession) IsWhitelistedAddress(account common.Address) (bool, error) {
	return _MetronionSale.Contract.IsWhitelistedAddress(&_MetronionSale.CallOpts, account)
}

// MaxWhitelistSize is a free data retrieval call binding the contract method 0x5f22eb6d.
//
// Solidity: function maxWhitelistSize() view returns(uint256)
func (_MetronionSale *MetronionSaleCaller) MaxWhitelistSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "maxWhitelistSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWhitelistSize is a free data retrieval call binding the contract method 0x5f22eb6d.
//
// Solidity: function maxWhitelistSize() view returns(uint256)
func (_MetronionSale *MetronionSaleSession) MaxWhitelistSize() (*big.Int, error) {
	return _MetronionSale.Contract.MaxWhitelistSize(&_MetronionSale.CallOpts)
}

// MaxWhitelistSize is a free data retrieval call binding the contract method 0x5f22eb6d.
//
// Solidity: function maxWhitelistSize() view returns(uint256)
func (_MetronionSale *MetronionSaleCallerSession) MaxWhitelistSize() (*big.Int, error) {
	return _MetronionSale.Contract.MaxWhitelistSize(&_MetronionSale.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_MetronionSale *MetronionSaleCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_MetronionSale *MetronionSaleSession) NftContract() (common.Address, error) {
	return _MetronionSale.Contract.NftContract(&_MetronionSale.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_MetronionSale *MetronionSaleCallerSession) NftContract() (common.Address, error) {
	return _MetronionSale.Contract.NftContract(&_MetronionSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionSale *MetronionSaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionSale *MetronionSaleSession) Owner() (common.Address, error) {
	return _MetronionSale.Contract.Owner(&_MetronionSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionSale *MetronionSaleCallerSession) Owner() (common.Address, error) {
	return _MetronionSale.Contract.Owner(&_MetronionSale.CallOpts)
}

// VersionId is a free data retrieval call binding the contract method 0xac0d925c.
//
// Solidity: function versionId() view returns(uint256)
func (_MetronionSale *MetronionSaleCaller) VersionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionSale.contract.Call(opts, &out, "versionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VersionId is a free data retrieval call binding the contract method 0xac0d925c.
//
// Solidity: function versionId() view returns(uint256)
func (_MetronionSale *MetronionSaleSession) VersionId() (*big.Int, error) {
	return _MetronionSale.Contract.VersionId(&_MetronionSale.CallOpts)
}

// VersionId is a free data retrieval call binding the contract method 0xac0d925c.
//
// Solidity: function versionId() view returns(uint256)
func (_MetronionSale *MetronionSaleCallerSession) VersionId() (*big.Int, error) {
	return _MetronionSale.Contract.VersionId(&_MetronionSale.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x215adc4a.
//
// Solidity: function buy(uint64 amount) payable returns()
func (_MetronionSale *MetronionSaleTransactor) Buy(opts *bind.TransactOpts, amount uint64) (*types.Transaction, error) {
	return _MetronionSale.contract.Transact(opts, "buy", amount)
}

// Buy is a paid mutator transaction binding the contract method 0x215adc4a.
//
// Solidity: function buy(uint64 amount) payable returns()
func (_MetronionSale *MetronionSaleSession) Buy(amount uint64) (*types.Transaction, error) {
	return _MetronionSale.Contract.Buy(&_MetronionSale.TransactOpts, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x215adc4a.
//
// Solidity: function buy(uint64 amount) payable returns()
func (_MetronionSale *MetronionSaleTransactorSession) Buy(amount uint64) (*types.Transaction, error) {
	return _MetronionSale.Contract.Buy(&_MetronionSale.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionSale *MetronionSaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionSale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionSale *MetronionSaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetronionSale.Contract.RenounceOwnership(&_MetronionSale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionSale *MetronionSaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetronionSale.Contract.RenounceOwnership(&_MetronionSale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionSale *MetronionSaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MetronionSale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionSale *MetronionSaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetronionSale.Contract.TransferOwnership(&_MetronionSale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionSale *MetronionSaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetronionSale.Contract.TransferOwnership(&_MetronionSale.TransactOpts, newOwner)
}

// UpdateWhitelistedGroup is a paid mutator transaction binding the contract method 0xadbf2149.
//
// Solidity: function updateWhitelistedGroup(address[] accounts, bool isWhitelisted) returns()
func (_MetronionSale *MetronionSaleTransactor) UpdateWhitelistedGroup(opts *bind.TransactOpts, accounts []common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _MetronionSale.contract.Transact(opts, "updateWhitelistedGroup", accounts, isWhitelisted)
}

// UpdateWhitelistedGroup is a paid mutator transaction binding the contract method 0xadbf2149.
//
// Solidity: function updateWhitelistedGroup(address[] accounts, bool isWhitelisted) returns()
func (_MetronionSale *MetronionSaleSession) UpdateWhitelistedGroup(accounts []common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _MetronionSale.Contract.UpdateWhitelistedGroup(&_MetronionSale.TransactOpts, accounts, isWhitelisted)
}

// UpdateWhitelistedGroup is a paid mutator transaction binding the contract method 0xadbf2149.
//
// Solidity: function updateWhitelistedGroup(address[] accounts, bool isWhitelisted) returns()
func (_MetronionSale *MetronionSaleTransactorSession) UpdateWhitelistedGroup(accounts []common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _MetronionSale.Contract.UpdateWhitelistedGroup(&_MetronionSale.TransactOpts, accounts, isWhitelisted)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_MetronionSale *MetronionSaleTransactor) WithdrawETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetronionSale.contract.Transact(opts, "withdrawETH", recipient, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_MetronionSale *MetronionSaleSession) WithdrawETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetronionSale.Contract.WithdrawETH(&_MetronionSale.TransactOpts, recipient, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_MetronionSale *MetronionSaleTransactorSession) WithdrawETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetronionSale.Contract.WithdrawETH(&_MetronionSale.TransactOpts, recipient, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address recipient, uint256 amount) returns()
func (_MetronionSale *MetronionSaleTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetronionSale.contract.Transact(opts, "withdrawToken", token, recipient, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address recipient, uint256 amount) returns()
func (_MetronionSale *MetronionSaleSession) WithdrawToken(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetronionSale.Contract.WithdrawToken(&_MetronionSale.TransactOpts, token, recipient, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address recipient, uint256 amount) returns()
func (_MetronionSale *MetronionSaleTransactorSession) WithdrawToken(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetronionSale.Contract.WithdrawToken(&_MetronionSale.TransactOpts, token, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetronionSale *MetronionSaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionSale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetronionSale *MetronionSaleSession) Receive() (*types.Transaction, error) {
	return _MetronionSale.Contract.Receive(&_MetronionSale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetronionSale *MetronionSaleTransactorSession) Receive() (*types.Transaction, error) {
	return _MetronionSale.Contract.Receive(&_MetronionSale.TransactOpts)
}

// MetronionSaleOwnerBoughtIterator is returned from FilterOwnerBought and is used to iterate over the raw logs and unpacked data for OwnerBought events raised by the MetronionSale contract.
type MetronionSaleOwnerBoughtIterator struct {
	Event *MetronionSaleOwnerBought // Event containing the contract specifics and raw log

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
func (it *MetronionSaleOwnerBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSaleOwnerBought)
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
		it.Event = new(MetronionSaleOwnerBought)
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
func (it *MetronionSaleOwnerBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSaleOwnerBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSaleOwnerBought represents a OwnerBought event raised by the MetronionSale contract.
type MetronionSaleOwnerBought struct {
	Buyer        common.Address
	VersionId    *big.Int
	TotalWeiPaid *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnerBought is a free log retrieval operation binding the contract event 0x6572890071cbac932382b30df01c6e6a38527876965e3f76818d6a297e5d99ad.
//
// Solidity: event OwnerBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) FilterOwnerBought(opts *bind.FilterOpts, buyer []common.Address) (*MetronionSaleOwnerBoughtIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "OwnerBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return &MetronionSaleOwnerBoughtIterator{contract: _MetronionSale.contract, event: "OwnerBought", logs: logs, sub: sub}, nil
}

// WatchOwnerBought is a free log subscription operation binding the contract event 0x6572890071cbac932382b30df01c6e6a38527876965e3f76818d6a297e5d99ad.
//
// Solidity: event OwnerBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) WatchOwnerBought(opts *bind.WatchOpts, sink chan<- *MetronionSaleOwnerBought, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "OwnerBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSaleOwnerBought)
				if err := _MetronionSale.contract.UnpackLog(event, "OwnerBought", log); err != nil {
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

// ParseOwnerBought is a log parse operation binding the contract event 0x6572890071cbac932382b30df01c6e6a38527876965e3f76818d6a297e5d99ad.
//
// Solidity: event OwnerBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) ParseOwnerBought(log types.Log) (*MetronionSaleOwnerBought, error) {
	event := new(MetronionSaleOwnerBought)
	if err := _MetronionSale.contract.UnpackLog(event, "OwnerBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionSaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetronionSale contract.
type MetronionSaleOwnershipTransferredIterator struct {
	Event *MetronionSaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MetronionSaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSaleOwnershipTransferred)
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
		it.Event = new(MetronionSaleOwnershipTransferred)
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
func (it *MetronionSaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSaleOwnershipTransferred represents a OwnershipTransferred event raised by the MetronionSale contract.
type MetronionSaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetronionSale *MetronionSaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetronionSaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetronionSaleOwnershipTransferredIterator{contract: _MetronionSale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetronionSale *MetronionSaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetronionSaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSaleOwnershipTransferred)
				if err := _MetronionSale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MetronionSale *MetronionSaleFilterer) ParseOwnershipTransferred(log types.Log) (*MetronionSaleOwnershipTransferred, error) {
	event := new(MetronionSaleOwnershipTransferred)
	if err := _MetronionSale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionSalePrivateBoughtIterator is returned from FilterPrivateBought and is used to iterate over the raw logs and unpacked data for PrivateBought events raised by the MetronionSale contract.
type MetronionSalePrivateBoughtIterator struct {
	Event *MetronionSalePrivateBought // Event containing the contract specifics and raw log

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
func (it *MetronionSalePrivateBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSalePrivateBought)
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
		it.Event = new(MetronionSalePrivateBought)
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
func (it *MetronionSalePrivateBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSalePrivateBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSalePrivateBought represents a PrivateBought event raised by the MetronionSale contract.
type MetronionSalePrivateBought struct {
	Buyer        common.Address
	VersionId    *big.Int
	TotalWeiPaid *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPrivateBought is a free log retrieval operation binding the contract event 0xa5e205e2aedf69a8cc205039b2416f6c4de2c274d202762643fd0c1655ef2b07.
//
// Solidity: event PrivateBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) FilterPrivateBought(opts *bind.FilterOpts, buyer []common.Address) (*MetronionSalePrivateBoughtIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "PrivateBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return &MetronionSalePrivateBoughtIterator{contract: _MetronionSale.contract, event: "PrivateBought", logs: logs, sub: sub}, nil
}

// WatchPrivateBought is a free log subscription operation binding the contract event 0xa5e205e2aedf69a8cc205039b2416f6c4de2c274d202762643fd0c1655ef2b07.
//
// Solidity: event PrivateBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) WatchPrivateBought(opts *bind.WatchOpts, sink chan<- *MetronionSalePrivateBought, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "PrivateBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSalePrivateBought)
				if err := _MetronionSale.contract.UnpackLog(event, "PrivateBought", log); err != nil {
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

// ParsePrivateBought is a log parse operation binding the contract event 0xa5e205e2aedf69a8cc205039b2416f6c4de2c274d202762643fd0c1655ef2b07.
//
// Solidity: event PrivateBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) ParsePrivateBought(log types.Log) (*MetronionSalePrivateBought, error) {
	event := new(MetronionSalePrivateBought)
	if err := _MetronionSale.contract.UnpackLog(event, "PrivateBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionSalePublicBoughtIterator is returned from FilterPublicBought and is used to iterate over the raw logs and unpacked data for PublicBought events raised by the MetronionSale contract.
type MetronionSalePublicBoughtIterator struct {
	Event *MetronionSalePublicBought // Event containing the contract specifics and raw log

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
func (it *MetronionSalePublicBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSalePublicBought)
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
		it.Event = new(MetronionSalePublicBought)
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
func (it *MetronionSalePublicBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSalePublicBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSalePublicBought represents a PublicBought event raised by the MetronionSale contract.
type MetronionSalePublicBought struct {
	Buyer        common.Address
	VersionId    *big.Int
	TotalWeiPaid *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPublicBought is a free log retrieval operation binding the contract event 0xbb6fccc04239b090cbe501ab994213815a96218e14a89d381646d0d4de434db5.
//
// Solidity: event PublicBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) FilterPublicBought(opts *bind.FilterOpts, buyer []common.Address) (*MetronionSalePublicBoughtIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "PublicBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return &MetronionSalePublicBoughtIterator{contract: _MetronionSale.contract, event: "PublicBought", logs: logs, sub: sub}, nil
}

// WatchPublicBought is a free log subscription operation binding the contract event 0xbb6fccc04239b090cbe501ab994213815a96218e14a89d381646d0d4de434db5.
//
// Solidity: event PublicBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) WatchPublicBought(opts *bind.WatchOpts, sink chan<- *MetronionSalePublicBought, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "PublicBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSalePublicBought)
				if err := _MetronionSale.contract.UnpackLog(event, "PublicBought", log); err != nil {
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

// ParsePublicBought is a log parse operation binding the contract event 0xbb6fccc04239b090cbe501ab994213815a96218e14a89d381646d0d4de434db5.
//
// Solidity: event PublicBought(address indexed buyer, uint256 versionId, uint256 totalWeiPaid)
func (_MetronionSale *MetronionSaleFilterer) ParsePublicBought(log types.Log) (*MetronionSalePublicBought, error) {
	event := new(MetronionSalePublicBought)
	if err := _MetronionSale.contract.UnpackLog(event, "PublicBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionSaleReceiveETHIterator is returned from FilterReceiveETH and is used to iterate over the raw logs and unpacked data for ReceiveETH events raised by the MetronionSale contract.
type MetronionSaleReceiveETHIterator struct {
	Event *MetronionSaleReceiveETH // Event containing the contract specifics and raw log

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
func (it *MetronionSaleReceiveETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSaleReceiveETH)
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
		it.Event = new(MetronionSaleReceiveETH)
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
func (it *MetronionSaleReceiveETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSaleReceiveETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSaleReceiveETH represents a ReceiveETH event raised by the MetronionSale contract.
type MetronionSaleReceiveETH struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceiveETH is a free log retrieval operation binding the contract event 0x830d2d700a97af574b186c80d40429385d24241565b08a7c559ba283a964d9b1.
//
// Solidity: event ReceiveETH(address from, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) FilterReceiveETH(opts *bind.FilterOpts) (*MetronionSaleReceiveETHIterator, error) {

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "ReceiveETH")
	if err != nil {
		return nil, err
	}
	return &MetronionSaleReceiveETHIterator{contract: _MetronionSale.contract, event: "ReceiveETH", logs: logs, sub: sub}, nil
}

// WatchReceiveETH is a free log subscription operation binding the contract event 0x830d2d700a97af574b186c80d40429385d24241565b08a7c559ba283a964d9b1.
//
// Solidity: event ReceiveETH(address from, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) WatchReceiveETH(opts *bind.WatchOpts, sink chan<- *MetronionSaleReceiveETH) (event.Subscription, error) {

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "ReceiveETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSaleReceiveETH)
				if err := _MetronionSale.contract.UnpackLog(event, "ReceiveETH", log); err != nil {
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

// ParseReceiveETH is a log parse operation binding the contract event 0x830d2d700a97af574b186c80d40429385d24241565b08a7c559ba283a964d9b1.
//
// Solidity: event ReceiveETH(address from, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) ParseReceiveETH(log types.Log) (*MetronionSaleReceiveETH, error) {
	event := new(MetronionSaleReceiveETH)
	if err := _MetronionSale.contract.UnpackLog(event, "ReceiveETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionSaleUpdateWhitelistedAddressIterator is returned from FilterUpdateWhitelistedAddress and is used to iterate over the raw logs and unpacked data for UpdateWhitelistedAddress events raised by the MetronionSale contract.
type MetronionSaleUpdateWhitelistedAddressIterator struct {
	Event *MetronionSaleUpdateWhitelistedAddress // Event containing the contract specifics and raw log

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
func (it *MetronionSaleUpdateWhitelistedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSaleUpdateWhitelistedAddress)
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
		it.Event = new(MetronionSaleUpdateWhitelistedAddress)
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
func (it *MetronionSaleUpdateWhitelistedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSaleUpdateWhitelistedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSaleUpdateWhitelistedAddress represents a UpdateWhitelistedAddress event raised by the MetronionSale contract.
type MetronionSaleUpdateWhitelistedAddress struct {
	Account       common.Address
	IsWhitelisted bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateWhitelistedAddress is a free log retrieval operation binding the contract event 0xf52869d259cfc36356be764e3ef9a4fc41b9858b06ca0c97577f362871b85b5c.
//
// Solidity: event UpdateWhitelistedAddress(address account, bool isWhitelisted)
func (_MetronionSale *MetronionSaleFilterer) FilterUpdateWhitelistedAddress(opts *bind.FilterOpts) (*MetronionSaleUpdateWhitelistedAddressIterator, error) {

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "UpdateWhitelistedAddress")
	if err != nil {
		return nil, err
	}
	return &MetronionSaleUpdateWhitelistedAddressIterator{contract: _MetronionSale.contract, event: "UpdateWhitelistedAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateWhitelistedAddress is a free log subscription operation binding the contract event 0xf52869d259cfc36356be764e3ef9a4fc41b9858b06ca0c97577f362871b85b5c.
//
// Solidity: event UpdateWhitelistedAddress(address account, bool isWhitelisted)
func (_MetronionSale *MetronionSaleFilterer) WatchUpdateWhitelistedAddress(opts *bind.WatchOpts, sink chan<- *MetronionSaleUpdateWhitelistedAddress) (event.Subscription, error) {

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "UpdateWhitelistedAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSaleUpdateWhitelistedAddress)
				if err := _MetronionSale.contract.UnpackLog(event, "UpdateWhitelistedAddress", log); err != nil {
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

// ParseUpdateWhitelistedAddress is a log parse operation binding the contract event 0xf52869d259cfc36356be764e3ef9a4fc41b9858b06ca0c97577f362871b85b5c.
//
// Solidity: event UpdateWhitelistedAddress(address account, bool isWhitelisted)
func (_MetronionSale *MetronionSaleFilterer) ParseUpdateWhitelistedAddress(log types.Log) (*MetronionSaleUpdateWhitelistedAddress, error) {
	event := new(MetronionSaleUpdateWhitelistedAddress)
	if err := _MetronionSale.contract.UnpackLog(event, "UpdateWhitelistedAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionSaleWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the MetronionSale contract.
type MetronionSaleWithdrawETHIterator struct {
	Event *MetronionSaleWithdrawETH // Event containing the contract specifics and raw log

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
func (it *MetronionSaleWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSaleWithdrawETH)
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
		it.Event = new(MetronionSaleWithdrawETH)
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
func (it *MetronionSaleWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSaleWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSaleWithdrawETH represents a WithdrawETH event raised by the MetronionSale contract.
type MetronionSaleWithdrawETH struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address recipient, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) FilterWithdrawETH(opts *bind.FilterOpts) (*MetronionSaleWithdrawETHIterator, error) {

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "WithdrawETH")
	if err != nil {
		return nil, err
	}
	return &MetronionSaleWithdrawETHIterator{contract: _MetronionSale.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address recipient, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *MetronionSaleWithdrawETH) (event.Subscription, error) {

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "WithdrawETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSaleWithdrawETH)
				if err := _MetronionSale.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address recipient, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) ParseWithdrawETH(log types.Log) (*MetronionSaleWithdrawETH, error) {
	event := new(MetronionSaleWithdrawETH)
	if err := _MetronionSale.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionSaleWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the MetronionSale contract.
type MetronionSaleWithdrawTokenIterator struct {
	Event *MetronionSaleWithdrawToken // Event containing the contract specifics and raw log

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
func (it *MetronionSaleWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionSaleWithdrawToken)
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
		it.Event = new(MetronionSaleWithdrawToken)
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
func (it *MetronionSaleWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionSaleWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionSaleWithdrawToken represents a WithdrawToken event raised by the MetronionSale contract.
type MetronionSaleWithdrawToken struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x037238854fe57fbf51f09946f854fc3916fe83938d6521f09bd05463839f1304.
//
// Solidity: event WithdrawToken(address token, address recipient, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) FilterWithdrawToken(opts *bind.FilterOpts) (*MetronionSaleWithdrawTokenIterator, error) {

	logs, sub, err := _MetronionSale.contract.FilterLogs(opts, "WithdrawToken")
	if err != nil {
		return nil, err
	}
	return &MetronionSaleWithdrawTokenIterator{contract: _MetronionSale.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x037238854fe57fbf51f09946f854fc3916fe83938d6521f09bd05463839f1304.
//
// Solidity: event WithdrawToken(address token, address recipient, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *MetronionSaleWithdrawToken) (event.Subscription, error) {

	logs, sub, err := _MetronionSale.contract.WatchLogs(opts, "WithdrawToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionSaleWithdrawToken)
				if err := _MetronionSale.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x037238854fe57fbf51f09946f854fc3916fe83938d6521f09bd05463839f1304.
//
// Solidity: event WithdrawToken(address token, address recipient, uint256 amount)
func (_MetronionSale *MetronionSaleFilterer) ParseWithdrawToken(log types.Log) (*MetronionSaleWithdrawToken, error) {
	event := new(MetronionSaleWithdrawToken)
	if err := _MetronionSale.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
