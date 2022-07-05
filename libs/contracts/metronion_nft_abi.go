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

// IMetronionNFTMetronion is an auto generated low-level Go binding around an user-defined struct.
type IMetronionNFTMetronion struct {
	Name    string
	Version *big.Int
}

// IMetronionNFTVersion is an auto generated low-level Go binding around an user-defined struct.
type IMetronionNFTVersion struct {
	StartingIndex *big.Int
	CurrentSupply *big.Int
	MaxSupply     *big.Int
	Provenance    string
}

// MetronionNFTABI is the input ABI used to generate the binding from.
const MetronionNFTABI = "[{\"inputs\":[{\"internalType\":\"contractIMetronionAccessories\",\"name\":\"_accessoriesContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_baseUri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_provenance\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"AccessoriesEquipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"AccessoriesUnequipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"MetronionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"}],\"name\":\"NewVersionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingIndex\",\"type\":\"uint256\"}],\"name\":\"StartingIndexFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"UpdateBaseURI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CAP_PER_MINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IMetronionNFT_Interface_Id\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessoriesContract\",\"outputs\":[{\"internalType\":\"contractIMetronionAccessories\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"provenance\",\"type\":\"string\"}],\"name\":\"addNewVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"changeMetronionName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"equipAccessories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"}],\"name\":\"finalizeStartingIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"}],\"name\":\"getMetronion\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"internalType\":\"structIMetronionNFT.Metronion\",\"name\":\"metronion\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"}],\"name\":\"getMetronionAccessories\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessoriesType\",\"type\":\"uint256\"}],\"name\":\"getMetronionAccessory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accessoryId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBatchAmountSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintMetronion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metronionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"accessoryIds\",\"type\":\"uint256[]\"}],\"name\":\"removeAccessories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"reservedNames\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyWithVersionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"versionId\",\"type\":\"uint256\"}],\"name\":\"versionById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startingIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"provenance\",\"type\":\"string\"}],\"internalType\":\"structIMetronionNFT.Version\",\"name\":\"version\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MetronionNFT is an auto generated Go binding around an Ethereum contract.
type MetronionNFT struct {
	MetronionNFTCaller     // Read-only binding to the contract
	MetronionNFTTransactor // Write-only binding to the contract
	MetronionNFTFilterer   // Log filterer for contract events
}

// MetronionNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetronionNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetronionNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetronionNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetronionNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetronionNFTSession struct {
	Contract     *MetronionNFT     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetronionNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetronionNFTCallerSession struct {
	Contract *MetronionNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MetronionNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetronionNFTTransactorSession struct {
	Contract     *MetronionNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MetronionNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetronionNFTRaw struct {
	Contract *MetronionNFT // Generic contract binding to access the raw methods on
}

// MetronionNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetronionNFTCallerRaw struct {
	Contract *MetronionNFTCaller // Generic read-only contract binding to access the raw methods on
}

// MetronionNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetronionNFTTransactorRaw struct {
	Contract *MetronionNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetronionNFT creates a new instance of MetronionNFT, bound to a specific deployed contract.
func NewMetronionNFT(address common.Address, backend bind.ContractBackend) (*MetronionNFT, error) {
	contract, err := bindMetronionNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetronionNFT{MetronionNFTCaller: MetronionNFTCaller{contract: contract}, MetronionNFTTransactor: MetronionNFTTransactor{contract: contract}, MetronionNFTFilterer: MetronionNFTFilterer{contract: contract}}, nil
}

// NewMetronionNFTCaller creates a new read-only instance of MetronionNFT, bound to a specific deployed contract.
func NewMetronionNFTCaller(address common.Address, caller bind.ContractCaller) (*MetronionNFTCaller, error) {
	contract, err := bindMetronionNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTCaller{contract: contract}, nil
}

// NewMetronionNFTTransactor creates a new write-only instance of MetronionNFT, bound to a specific deployed contract.
func NewMetronionNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*MetronionNFTTransactor, error) {
	contract, err := bindMetronionNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTTransactor{contract: contract}, nil
}

// NewMetronionNFTFilterer creates a new log filterer instance of MetronionNFT, bound to a specific deployed contract.
func NewMetronionNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*MetronionNFTFilterer, error) {
	contract, err := bindMetronionNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTFilterer{contract: contract}, nil
}

// bindMetronionNFT binds a generic wrapper to an already deployed contract.
func bindMetronionNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetronionNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetronionNFT *MetronionNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetronionNFT.Contract.MetronionNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetronionNFT *MetronionNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionNFT.Contract.MetronionNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetronionNFT *MetronionNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetronionNFT.Contract.MetronionNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetronionNFT *MetronionNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetronionNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetronionNFT *MetronionNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetronionNFT *MetronionNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetronionNFT.Contract.contract.Transact(opts, method, params...)
}

// CAPPERMINT is a free data retrieval call binding the contract method 0xdc548a9a.
//
// Solidity: function CAP_PER_MINT() view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) CAPPERMINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "CAP_PER_MINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CAPPERMINT is a free data retrieval call binding the contract method 0xdc548a9a.
//
// Solidity: function CAP_PER_MINT() view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) CAPPERMINT() (*big.Int, error) {
	return _MetronionNFT.Contract.CAPPERMINT(&_MetronionNFT.CallOpts)
}

// CAPPERMINT is a free data retrieval call binding the contract method 0xdc548a9a.
//
// Solidity: function CAP_PER_MINT() view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) CAPPERMINT() (*big.Int, error) {
	return _MetronionNFT.Contract.CAPPERMINT(&_MetronionNFT.CallOpts)
}

// IMetronionNFTInterfaceId is a free data retrieval call binding the contract method 0xaad60726.
//
// Solidity: function IMetronionNFT_Interface_Id() view returns(bytes4)
func (_MetronionNFT *MetronionNFTCaller) IMetronionNFTInterfaceId(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "IMetronionNFT_Interface_Id")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IMetronionNFTInterfaceId is a free data retrieval call binding the contract method 0xaad60726.
//
// Solidity: function IMetronionNFT_Interface_Id() view returns(bytes4)
func (_MetronionNFT *MetronionNFTSession) IMetronionNFTInterfaceId() ([4]byte, error) {
	return _MetronionNFT.Contract.IMetronionNFTInterfaceId(&_MetronionNFT.CallOpts)
}

// IMetronionNFTInterfaceId is a free data retrieval call binding the contract method 0xaad60726.
//
// Solidity: function IMetronionNFT_Interface_Id() view returns(bytes4)
func (_MetronionNFT *MetronionNFTCallerSession) IMetronionNFTInterfaceId() ([4]byte, error) {
	return _MetronionNFT.Contract.IMetronionNFTInterfaceId(&_MetronionNFT.CallOpts)
}

// AccessoriesContract is a free data retrieval call binding the contract method 0xa4b83625.
//
// Solidity: function accessoriesContract() view returns(address)
func (_MetronionNFT *MetronionNFTCaller) AccessoriesContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "accessoriesContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessoriesContract is a free data retrieval call binding the contract method 0xa4b83625.
//
// Solidity: function accessoriesContract() view returns(address)
func (_MetronionNFT *MetronionNFTSession) AccessoriesContract() (common.Address, error) {
	return _MetronionNFT.Contract.AccessoriesContract(&_MetronionNFT.CallOpts)
}

// AccessoriesContract is a free data retrieval call binding the contract method 0xa4b83625.
//
// Solidity: function accessoriesContract() view returns(address)
func (_MetronionNFT *MetronionNFTCallerSession) AccessoriesContract() (common.Address, error) {
	return _MetronionNFT.Contract.AccessoriesContract(&_MetronionNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 assetId) view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) BalanceOf(opts *bind.CallOpts, account common.Address, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "balanceOf", account, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 assetId) view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) BalanceOf(account common.Address, assetId *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.BalanceOf(&_MetronionNFT.CallOpts, account, assetId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 assetId) view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) BalanceOf(account common.Address, assetId *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.BalanceOf(&_MetronionNFT.CallOpts, account, assetId)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) BalanceOf0(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "balanceOf0", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) BalanceOf0(owner common.Address) (*big.Int, error) {
	return _MetronionNFT.Contract.BalanceOf0(&_MetronionNFT.CallOpts, owner)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) BalanceOf0(owner common.Address) (*big.Int, error) {
	return _MetronionNFT.Contract.BalanceOf0(&_MetronionNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MetronionNFT *MetronionNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MetronionNFT *MetronionNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MetronionNFT.Contract.GetApproved(&_MetronionNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MetronionNFT *MetronionNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MetronionNFT.Contract.GetApproved(&_MetronionNFT.CallOpts, tokenId)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) GetLatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "getLatestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) GetLatestVersion() (*big.Int, error) {
	return _MetronionNFT.Contract.GetLatestVersion(&_MetronionNFT.CallOpts)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) GetLatestVersion() (*big.Int, error) {
	return _MetronionNFT.Contract.GetLatestVersion(&_MetronionNFT.CallOpts)
}

// GetMetronion is a free data retrieval call binding the contract method 0x3f90c66a.
//
// Solidity: function getMetronion(uint256 metronionId) view returns((string,uint256) metronion)
func (_MetronionNFT *MetronionNFTCaller) GetMetronion(opts *bind.CallOpts, metronionId *big.Int) (IMetronionNFTMetronion, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "getMetronion", metronionId)

	if err != nil {
		return *new(IMetronionNFTMetronion), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetronionNFTMetronion)).(*IMetronionNFTMetronion)

	return out0, err

}

// GetMetronion is a free data retrieval call binding the contract method 0x3f90c66a.
//
// Solidity: function getMetronion(uint256 metronionId) view returns((string,uint256) metronion)
func (_MetronionNFT *MetronionNFTSession) GetMetronion(metronionId *big.Int) (IMetronionNFTMetronion, error) {
	return _MetronionNFT.Contract.GetMetronion(&_MetronionNFT.CallOpts, metronionId)
}

// GetMetronion is a free data retrieval call binding the contract method 0x3f90c66a.
//
// Solidity: function getMetronion(uint256 metronionId) view returns((string,uint256) metronion)
func (_MetronionNFT *MetronionNFTCallerSession) GetMetronion(metronionId *big.Int) (IMetronionNFTMetronion, error) {
	return _MetronionNFT.Contract.GetMetronion(&_MetronionNFT.CallOpts, metronionId)
}

// GetMetronionAccessories is a free data retrieval call binding the contract method 0x591214f3.
//
// Solidity: function getMetronionAccessories(uint256 metronionId) view returns(uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTCaller) GetMetronionAccessories(opts *bind.CallOpts, metronionId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "getMetronionAccessories", metronionId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetMetronionAccessories is a free data retrieval call binding the contract method 0x591214f3.
//
// Solidity: function getMetronionAccessories(uint256 metronionId) view returns(uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTSession) GetMetronionAccessories(metronionId *big.Int) ([]*big.Int, error) {
	return _MetronionNFT.Contract.GetMetronionAccessories(&_MetronionNFT.CallOpts, metronionId)
}

// GetMetronionAccessories is a free data retrieval call binding the contract method 0x591214f3.
//
// Solidity: function getMetronionAccessories(uint256 metronionId) view returns(uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTCallerSession) GetMetronionAccessories(metronionId *big.Int) ([]*big.Int, error) {
	return _MetronionNFT.Contract.GetMetronionAccessories(&_MetronionNFT.CallOpts, metronionId)
}

// GetMetronionAccessory is a free data retrieval call binding the contract method 0x39d66639.
//
// Solidity: function getMetronionAccessory(uint256 metronionId, uint256 accessoriesType) view returns(uint256 accessoryId)
func (_MetronionNFT *MetronionNFTCaller) GetMetronionAccessory(opts *bind.CallOpts, metronionId *big.Int, accessoriesType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "getMetronionAccessory", metronionId, accessoriesType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMetronionAccessory is a free data retrieval call binding the contract method 0x39d66639.
//
// Solidity: function getMetronionAccessory(uint256 metronionId, uint256 accessoriesType) view returns(uint256 accessoryId)
func (_MetronionNFT *MetronionNFTSession) GetMetronionAccessory(metronionId *big.Int, accessoriesType *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.GetMetronionAccessory(&_MetronionNFT.CallOpts, metronionId, accessoriesType)
}

// GetMetronionAccessory is a free data retrieval call binding the contract method 0x39d66639.
//
// Solidity: function getMetronionAccessory(uint256 metronionId, uint256 accessoriesType) view returns(uint256 accessoryId)
func (_MetronionNFT *MetronionNFTCallerSession) GetMetronionAccessory(metronionId *big.Int, accessoriesType *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.GetMetronionAccessory(&_MetronionNFT.CallOpts, metronionId, accessoriesType)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetronionNFT *MetronionNFTCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetronionNFT *MetronionNFTSession) GetOperators() ([]common.Address, error) {
	return _MetronionNFT.Contract.GetOperators(&_MetronionNFT.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetronionNFT *MetronionNFTCallerSession) GetOperators() ([]common.Address, error) {
	return _MetronionNFT.Contract.GetOperators(&_MetronionNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetronionNFT *MetronionNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetronionNFT *MetronionNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MetronionNFT.Contract.IsApprovedForAll(&_MetronionNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MetronionNFT *MetronionNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MetronionNFT.Contract.IsApprovedForAll(&_MetronionNFT.CallOpts, owner, operator)
}

// IsBatchAmountSupported is a free data retrieval call binding the contract method 0xde271441.
//
// Solidity: function isBatchAmountSupported() pure returns(bool)
func (_MetronionNFT *MetronionNFTCaller) IsBatchAmountSupported(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "isBatchAmountSupported")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchAmountSupported is a free data retrieval call binding the contract method 0xde271441.
//
// Solidity: function isBatchAmountSupported() pure returns(bool)
func (_MetronionNFT *MetronionNFTSession) IsBatchAmountSupported() (bool, error) {
	return _MetronionNFT.Contract.IsBatchAmountSupported(&_MetronionNFT.CallOpts)
}

// IsBatchAmountSupported is a free data retrieval call binding the contract method 0xde271441.
//
// Solidity: function isBatchAmountSupported() pure returns(bool)
func (_MetronionNFT *MetronionNFTCallerSession) IsBatchAmountSupported() (bool, error) {
	return _MetronionNFT.Contract.IsBatchAmountSupported(&_MetronionNFT.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetronionNFT *MetronionNFTCaller) IsOperator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "isOperator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetronionNFT *MetronionNFTSession) IsOperator(account common.Address) (bool, error) {
	return _MetronionNFT.Contract.IsOperator(&_MetronionNFT.CallOpts, account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetronionNFT *MetronionNFTCallerSession) IsOperator(account common.Address) (bool, error) {
	return _MetronionNFT.Contract.IsOperator(&_MetronionNFT.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetronionNFT *MetronionNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetronionNFT *MetronionNFTSession) Name() (string, error) {
	return _MetronionNFT.Contract.Name(&_MetronionNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetronionNFT *MetronionNFTCallerSession) Name() (string, error) {
	return _MetronionNFT.Contract.Name(&_MetronionNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionNFT *MetronionNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionNFT *MetronionNFTSession) Owner() (common.Address, error) {
	return _MetronionNFT.Contract.Owner(&_MetronionNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetronionNFT *MetronionNFTCallerSession) Owner() (common.Address, error) {
	return _MetronionNFT.Contract.Owner(&_MetronionNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MetronionNFT *MetronionNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MetronionNFT *MetronionNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MetronionNFT.Contract.OwnerOf(&_MetronionNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MetronionNFT *MetronionNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MetronionNFT.Contract.OwnerOf(&_MetronionNFT.CallOpts, tokenId)
}

// ReservedNames is a free data retrieval call binding the contract method 0x68436f10.
//
// Solidity: function reservedNames(string ) view returns(bool)
func (_MetronionNFT *MetronionNFTCaller) ReservedNames(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "reservedNames", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReservedNames is a free data retrieval call binding the contract method 0x68436f10.
//
// Solidity: function reservedNames(string ) view returns(bool)
func (_MetronionNFT *MetronionNFTSession) ReservedNames(arg0 string) (bool, error) {
	return _MetronionNFT.Contract.ReservedNames(&_MetronionNFT.CallOpts, arg0)
}

// ReservedNames is a free data retrieval call binding the contract method 0x68436f10.
//
// Solidity: function reservedNames(string ) view returns(bool)
func (_MetronionNFT *MetronionNFTCallerSession) ReservedNames(arg0 string) (bool, error) {
	return _MetronionNFT.Contract.ReservedNames(&_MetronionNFT.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetronionNFT *MetronionNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetronionNFT *MetronionNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetronionNFT.Contract.SupportsInterface(&_MetronionNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetronionNFT *MetronionNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetronionNFT.Contract.SupportsInterface(&_MetronionNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetronionNFT *MetronionNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetronionNFT *MetronionNFTSession) Symbol() (string, error) {
	return _MetronionNFT.Contract.Symbol(&_MetronionNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetronionNFT *MetronionNFTCallerSession) Symbol() (string, error) {
	return _MetronionNFT.Contract.Symbol(&_MetronionNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.TokenByIndex(&_MetronionNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.TokenByIndex(&_MetronionNFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.TokenOfOwnerByIndex(&_MetronionNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.TokenOfOwnerByIndex(&_MetronionNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MetronionNFT *MetronionNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MetronionNFT *MetronionNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MetronionNFT.Contract.TokenURI(&_MetronionNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MetronionNFT *MetronionNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MetronionNFT.Contract.TokenURI(&_MetronionNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) TotalSupply() (*big.Int, error) {
	return _MetronionNFT.Contract.TotalSupply(&_MetronionNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _MetronionNFT.Contract.TotalSupply(&_MetronionNFT.CallOpts)
}

// TotalSupplyWithVersionId is a free data retrieval call binding the contract method 0xc1221f46.
//
// Solidity: function totalSupplyWithVersionId(uint256 versionId) view returns(uint256)
func (_MetronionNFT *MetronionNFTCaller) TotalSupplyWithVersionId(opts *bind.CallOpts, versionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "totalSupplyWithVersionId", versionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyWithVersionId is a free data retrieval call binding the contract method 0xc1221f46.
//
// Solidity: function totalSupplyWithVersionId(uint256 versionId) view returns(uint256)
func (_MetronionNFT *MetronionNFTSession) TotalSupplyWithVersionId(versionId *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.TotalSupplyWithVersionId(&_MetronionNFT.CallOpts, versionId)
}

// TotalSupplyWithVersionId is a free data retrieval call binding the contract method 0xc1221f46.
//
// Solidity: function totalSupplyWithVersionId(uint256 versionId) view returns(uint256)
func (_MetronionNFT *MetronionNFTCallerSession) TotalSupplyWithVersionId(versionId *big.Int) (*big.Int, error) {
	return _MetronionNFT.Contract.TotalSupplyWithVersionId(&_MetronionNFT.CallOpts, versionId)
}

// VersionById is a free data retrieval call binding the contract method 0x53a6c007.
//
// Solidity: function versionById(uint256 versionId) view returns((uint256,uint256,uint256,string) version)
func (_MetronionNFT *MetronionNFTCaller) VersionById(opts *bind.CallOpts, versionId *big.Int) (IMetronionNFTVersion, error) {
	var out []interface{}
	err := _MetronionNFT.contract.Call(opts, &out, "versionById", versionId)

	if err != nil {
		return *new(IMetronionNFTVersion), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetronionNFTVersion)).(*IMetronionNFTVersion)

	return out0, err

}

// VersionById is a free data retrieval call binding the contract method 0x53a6c007.
//
// Solidity: function versionById(uint256 versionId) view returns((uint256,uint256,uint256,string) version)
func (_MetronionNFT *MetronionNFTSession) VersionById(versionId *big.Int) (IMetronionNFTVersion, error) {
	return _MetronionNFT.Contract.VersionById(&_MetronionNFT.CallOpts, versionId)
}

// VersionById is a free data retrieval call binding the contract method 0x53a6c007.
//
// Solidity: function versionById(uint256 versionId) view returns((uint256,uint256,uint256,string) version)
func (_MetronionNFT *MetronionNFTCallerSession) VersionById(versionId *big.Int) (IMetronionNFTVersion, error) {
	return _MetronionNFT.Contract.VersionById(&_MetronionNFT.CallOpts, versionId)
}

// AddNewVersion is a paid mutator transaction binding the contract method 0x0f73167a.
//
// Solidity: function addNewVersion(uint256 maxSupply, string provenance) returns()
func (_MetronionNFT *MetronionNFTTransactor) AddNewVersion(opts *bind.TransactOpts, maxSupply *big.Int, provenance string) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "addNewVersion", maxSupply, provenance)
}

// AddNewVersion is a paid mutator transaction binding the contract method 0x0f73167a.
//
// Solidity: function addNewVersion(uint256 maxSupply, string provenance) returns()
func (_MetronionNFT *MetronionNFTSession) AddNewVersion(maxSupply *big.Int, provenance string) (*types.Transaction, error) {
	return _MetronionNFT.Contract.AddNewVersion(&_MetronionNFT.TransactOpts, maxSupply, provenance)
}

// AddNewVersion is a paid mutator transaction binding the contract method 0x0f73167a.
//
// Solidity: function addNewVersion(uint256 maxSupply, string provenance) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) AddNewVersion(maxSupply *big.Int, provenance string) (*types.Transaction, error) {
	return _MetronionNFT.Contract.AddNewVersion(&_MetronionNFT.TransactOpts, maxSupply, provenance)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetronionNFT *MetronionNFTTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "addOperator", operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetronionNFT *MetronionNFTSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.AddOperator(&_MetronionNFT.TransactOpts, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.AddOperator(&_MetronionNFT.TransactOpts, operator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.Approve(&_MetronionNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.Approve(&_MetronionNFT.TransactOpts, to, tokenId)
}

// ChangeMetronionName is a paid mutator transaction binding the contract method 0x73fdbb7f.
//
// Solidity: function changeMetronionName(uint256 metronionId, string newName) returns()
func (_MetronionNFT *MetronionNFTTransactor) ChangeMetronionName(opts *bind.TransactOpts, metronionId *big.Int, newName string) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "changeMetronionName", metronionId, newName)
}

// ChangeMetronionName is a paid mutator transaction binding the contract method 0x73fdbb7f.
//
// Solidity: function changeMetronionName(uint256 metronionId, string newName) returns()
func (_MetronionNFT *MetronionNFTSession) ChangeMetronionName(metronionId *big.Int, newName string) (*types.Transaction, error) {
	return _MetronionNFT.Contract.ChangeMetronionName(&_MetronionNFT.TransactOpts, metronionId, newName)
}

// ChangeMetronionName is a paid mutator transaction binding the contract method 0x73fdbb7f.
//
// Solidity: function changeMetronionName(uint256 metronionId, string newName) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) ChangeMetronionName(metronionId *big.Int, newName string) (*types.Transaction, error) {
	return _MetronionNFT.Contract.ChangeMetronionName(&_MetronionNFT.TransactOpts, metronionId, newName)
}

// EquipAccessories is a paid mutator transaction binding the contract method 0x1deaafc4.
//
// Solidity: function equipAccessories(uint256 metronionId, uint256[] accessoryIds) returns()
func (_MetronionNFT *MetronionNFTTransactor) EquipAccessories(opts *bind.TransactOpts, metronionId *big.Int, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "equipAccessories", metronionId, accessoryIds)
}

// EquipAccessories is a paid mutator transaction binding the contract method 0x1deaafc4.
//
// Solidity: function equipAccessories(uint256 metronionId, uint256[] accessoryIds) returns()
func (_MetronionNFT *MetronionNFTSession) EquipAccessories(metronionId *big.Int, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.EquipAccessories(&_MetronionNFT.TransactOpts, metronionId, accessoryIds)
}

// EquipAccessories is a paid mutator transaction binding the contract method 0x1deaafc4.
//
// Solidity: function equipAccessories(uint256 metronionId, uint256[] accessoryIds) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) EquipAccessories(metronionId *big.Int, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.EquipAccessories(&_MetronionNFT.TransactOpts, metronionId, accessoryIds)
}

// FinalizeStartingIndex is a paid mutator transaction binding the contract method 0x3add8fed.
//
// Solidity: function finalizeStartingIndex(uint256 versionId) returns()
func (_MetronionNFT *MetronionNFTTransactor) FinalizeStartingIndex(opts *bind.TransactOpts, versionId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "finalizeStartingIndex", versionId)
}

// FinalizeStartingIndex is a paid mutator transaction binding the contract method 0x3add8fed.
//
// Solidity: function finalizeStartingIndex(uint256 versionId) returns()
func (_MetronionNFT *MetronionNFTSession) FinalizeStartingIndex(versionId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.FinalizeStartingIndex(&_MetronionNFT.TransactOpts, versionId)
}

// FinalizeStartingIndex is a paid mutator transaction binding the contract method 0x3add8fed.
//
// Solidity: function finalizeStartingIndex(uint256 versionId) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) FinalizeStartingIndex(versionId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.FinalizeStartingIndex(&_MetronionNFT.TransactOpts, versionId)
}

// MintMetronion is a paid mutator transaction binding the contract method 0x0ed528fa.
//
// Solidity: function mintMetronion(uint256 versionId, uint256 amount, address to) returns()
func (_MetronionNFT *MetronionNFTTransactor) MintMetronion(opts *bind.TransactOpts, versionId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "mintMetronion", versionId, amount, to)
}

// MintMetronion is a paid mutator transaction binding the contract method 0x0ed528fa.
//
// Solidity: function mintMetronion(uint256 versionId, uint256 amount, address to) returns()
func (_MetronionNFT *MetronionNFTSession) MintMetronion(versionId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.MintMetronion(&_MetronionNFT.TransactOpts, versionId, amount, to)
}

// MintMetronion is a paid mutator transaction binding the contract method 0x0ed528fa.
//
// Solidity: function mintMetronion(uint256 versionId, uint256 amount, address to) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) MintMetronion(versionId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.MintMetronion(&_MetronionNFT.TransactOpts, versionId, amount, to)
}

// RemoveAccessories is a paid mutator transaction binding the contract method 0xeaccb03e.
//
// Solidity: function removeAccessories(uint256 metronionId, uint256[] accessoryIds) returns()
func (_MetronionNFT *MetronionNFTTransactor) RemoveAccessories(opts *bind.TransactOpts, metronionId *big.Int, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "removeAccessories", metronionId, accessoryIds)
}

// RemoveAccessories is a paid mutator transaction binding the contract method 0xeaccb03e.
//
// Solidity: function removeAccessories(uint256 metronionId, uint256[] accessoryIds) returns()
func (_MetronionNFT *MetronionNFTSession) RemoveAccessories(metronionId *big.Int, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.RemoveAccessories(&_MetronionNFT.TransactOpts, metronionId, accessoryIds)
}

// RemoveAccessories is a paid mutator transaction binding the contract method 0xeaccb03e.
//
// Solidity: function removeAccessories(uint256 metronionId, uint256[] accessoryIds) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) RemoveAccessories(metronionId *big.Int, accessoryIds []*big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.RemoveAccessories(&_MetronionNFT.TransactOpts, metronionId, accessoryIds)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetronionNFT *MetronionNFTTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetronionNFT *MetronionNFTSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.RemoveOperator(&_MetronionNFT.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.RemoveOperator(&_MetronionNFT.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionNFT *MetronionNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionNFT *MetronionNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetronionNFT.Contract.RenounceOwnership(&_MetronionNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetronionNFT *MetronionNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetronionNFT.Contract.RenounceOwnership(&_MetronionNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 assetId, uint256 ) returns()
func (_MetronionNFT *MetronionNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, assetId *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "safeTransferFrom", from, to, assetId, arg3)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 assetId, uint256 ) returns()
func (_MetronionNFT *MetronionNFTSession) SafeTransferFrom(from common.Address, to common.Address, assetId *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SafeTransferFrom(&_MetronionNFT.TransactOpts, from, to, assetId, arg3)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x0febdd49.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 assetId, uint256 ) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, assetId *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SafeTransferFrom(&_MetronionNFT.TransactOpts, from, to, assetId, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SafeTransferFrom0(&_MetronionNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SafeTransferFrom0(&_MetronionNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom1 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MetronionNFT *MetronionNFTTransactor) SafeTransferFrom1(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "safeTransferFrom1", from, to, tokenId, _data)
}

// SafeTransferFrom1 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MetronionNFT *MetronionNFTSession) SafeTransferFrom1(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SafeTransferFrom1(&_MetronionNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom1 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) SafeTransferFrom1(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SafeTransferFrom1(&_MetronionNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetronionNFT *MetronionNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetronionNFT *MetronionNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SetApprovalForAll(&_MetronionNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SetApprovalForAll(&_MetronionNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_MetronionNFT *MetronionNFTTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_MetronionNFT *MetronionNFTSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SetBaseURI(&_MetronionNFT.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _MetronionNFT.Contract.SetBaseURI(&_MetronionNFT.TransactOpts, baseURI)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.TransferFrom(&_MetronionNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetronionNFT.Contract.TransferFrom(&_MetronionNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionNFT *MetronionNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MetronionNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionNFT *MetronionNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.TransferOwnership(&_MetronionNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetronionNFT *MetronionNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetronionNFT.Contract.TransferOwnership(&_MetronionNFT.TransactOpts, newOwner)
}

// MetronionNFTAccessoriesEquippedIterator is returned from FilterAccessoriesEquipped and is used to iterate over the raw logs and unpacked data for AccessoriesEquipped events raised by the MetronionNFT contract.
type MetronionNFTAccessoriesEquippedIterator struct {
	Event *MetronionNFTAccessoriesEquipped // Event containing the contract specifics and raw log

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
func (it *MetronionNFTAccessoriesEquippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTAccessoriesEquipped)
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
		it.Event = new(MetronionNFTAccessoriesEquipped)
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
func (it *MetronionNFTAccessoriesEquippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTAccessoriesEquippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTAccessoriesEquipped represents a AccessoriesEquipped event raised by the MetronionNFT contract.
type MetronionNFTAccessoriesEquipped struct {
	MetronionId  *big.Int
	AccessoryIds []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccessoriesEquipped is a free log retrieval operation binding the contract event 0x8542bdcfcff20d1c1f4d62bad45e1074b5b02b65671cb0d5542bc2f3616c79a6.
//
// Solidity: event AccessoriesEquipped(uint256 indexed metronionId, uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTFilterer) FilterAccessoriesEquipped(opts *bind.FilterOpts, metronionId []*big.Int) (*MetronionNFTAccessoriesEquippedIterator, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "AccessoriesEquipped", metronionIdRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTAccessoriesEquippedIterator{contract: _MetronionNFT.contract, event: "AccessoriesEquipped", logs: logs, sub: sub}, nil
}

// WatchAccessoriesEquipped is a free log subscription operation binding the contract event 0x8542bdcfcff20d1c1f4d62bad45e1074b5b02b65671cb0d5542bc2f3616c79a6.
//
// Solidity: event AccessoriesEquipped(uint256 indexed metronionId, uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTFilterer) WatchAccessoriesEquipped(opts *bind.WatchOpts, sink chan<- *MetronionNFTAccessoriesEquipped, metronionId []*big.Int) (event.Subscription, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "AccessoriesEquipped", metronionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTAccessoriesEquipped)
				if err := _MetronionNFT.contract.UnpackLog(event, "AccessoriesEquipped", log); err != nil {
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

// ParseAccessoriesEquipped is a log parse operation binding the contract event 0x8542bdcfcff20d1c1f4d62bad45e1074b5b02b65671cb0d5542bc2f3616c79a6.
//
// Solidity: event AccessoriesEquipped(uint256 indexed metronionId, uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTFilterer) ParseAccessoriesEquipped(log types.Log) (*MetronionNFTAccessoriesEquipped, error) {
	event := new(MetronionNFTAccessoriesEquipped)
	if err := _MetronionNFT.contract.UnpackLog(event, "AccessoriesEquipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTAccessoriesUnequippedIterator is returned from FilterAccessoriesUnequipped and is used to iterate over the raw logs and unpacked data for AccessoriesUnequipped events raised by the MetronionNFT contract.
type MetronionNFTAccessoriesUnequippedIterator struct {
	Event *MetronionNFTAccessoriesUnequipped // Event containing the contract specifics and raw log

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
func (it *MetronionNFTAccessoriesUnequippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTAccessoriesUnequipped)
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
		it.Event = new(MetronionNFTAccessoriesUnequipped)
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
func (it *MetronionNFTAccessoriesUnequippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTAccessoriesUnequippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTAccessoriesUnequipped represents a AccessoriesUnequipped event raised by the MetronionNFT contract.
type MetronionNFTAccessoriesUnequipped struct {
	MetronionId  *big.Int
	AccessoryIds []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccessoriesUnequipped is a free log retrieval operation binding the contract event 0xc293cce285d493f6602a8b2420723f08acc358a693b3573ffa33ad90597350dc.
//
// Solidity: event AccessoriesUnequipped(uint256 indexed metronionId, uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTFilterer) FilterAccessoriesUnequipped(opts *bind.FilterOpts, metronionId []*big.Int) (*MetronionNFTAccessoriesUnequippedIterator, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "AccessoriesUnequipped", metronionIdRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTAccessoriesUnequippedIterator{contract: _MetronionNFT.contract, event: "AccessoriesUnequipped", logs: logs, sub: sub}, nil
}

// WatchAccessoriesUnequipped is a free log subscription operation binding the contract event 0xc293cce285d493f6602a8b2420723f08acc358a693b3573ffa33ad90597350dc.
//
// Solidity: event AccessoriesUnequipped(uint256 indexed metronionId, uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTFilterer) WatchAccessoriesUnequipped(opts *bind.WatchOpts, sink chan<- *MetronionNFTAccessoriesUnequipped, metronionId []*big.Int) (event.Subscription, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "AccessoriesUnequipped", metronionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTAccessoriesUnequipped)
				if err := _MetronionNFT.contract.UnpackLog(event, "AccessoriesUnequipped", log); err != nil {
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

// ParseAccessoriesUnequipped is a log parse operation binding the contract event 0xc293cce285d493f6602a8b2420723f08acc358a693b3573ffa33ad90597350dc.
//
// Solidity: event AccessoriesUnequipped(uint256 indexed metronionId, uint256[] accessoryIds)
func (_MetronionNFT *MetronionNFTFilterer) ParseAccessoriesUnequipped(log types.Log) (*MetronionNFTAccessoriesUnequipped, error) {
	event := new(MetronionNFTAccessoriesUnequipped)
	if err := _MetronionNFT.contract.UnpackLog(event, "AccessoriesUnequipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MetronionNFT contract.
type MetronionNFTApprovalIterator struct {
	Event *MetronionNFTApproval // Event containing the contract specifics and raw log

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
func (it *MetronionNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTApproval)
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
		it.Event = new(MetronionNFTApproval)
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
func (it *MetronionNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTApproval represents a Approval event raised by the MetronionNFT contract.
type MetronionNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MetronionNFT *MetronionNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MetronionNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTApprovalIterator{contract: _MetronionNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MetronionNFT *MetronionNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MetronionNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTApproval)
				if err := _MetronionNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MetronionNFT *MetronionNFTFilterer) ParseApproval(log types.Log) (*MetronionNFTApproval, error) {
	event := new(MetronionNFTApproval)
	if err := _MetronionNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MetronionNFT contract.
type MetronionNFTApprovalForAllIterator struct {
	Event *MetronionNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MetronionNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTApprovalForAll)
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
		it.Event = new(MetronionNFTApprovalForAll)
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
func (it *MetronionNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTApprovalForAll represents a ApprovalForAll event raised by the MetronionNFT contract.
type MetronionNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MetronionNFT *MetronionNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MetronionNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTApprovalForAllIterator{contract: _MetronionNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MetronionNFT *MetronionNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MetronionNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTApprovalForAll)
				if err := _MetronionNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MetronionNFT *MetronionNFTFilterer) ParseApprovalForAll(log types.Log) (*MetronionNFTApprovalForAll, error) {
	event := new(MetronionNFTApprovalForAll)
	if err := _MetronionNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTMetronionCreatedIterator is returned from FilterMetronionCreated and is used to iterate over the raw logs and unpacked data for MetronionCreated events raised by the MetronionNFT contract.
type MetronionNFTMetronionCreatedIterator struct {
	Event *MetronionNFTMetronionCreated // Event containing the contract specifics and raw log

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
func (it *MetronionNFTMetronionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTMetronionCreated)
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
		it.Event = new(MetronionNFTMetronionCreated)
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
func (it *MetronionNFTMetronionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTMetronionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTMetronionCreated represents a MetronionCreated event raised by the MetronionNFT contract.
type MetronionNFTMetronionCreated struct {
	MetronionId *big.Int
	VersionId   *big.Int
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMetronionCreated is a free log retrieval operation binding the contract event 0xc6ba76065f5132d4931c49972c276c27327502e3127e19e9499170265674da48.
//
// Solidity: event MetronionCreated(uint256 indexed metronionId, uint256 indexed versionId, address to)
func (_MetronionNFT *MetronionNFTFilterer) FilterMetronionCreated(opts *bind.FilterOpts, metronionId []*big.Int, versionId []*big.Int) (*MetronionNFTMetronionCreatedIterator, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}
	var versionIdRule []interface{}
	for _, versionIdItem := range versionId {
		versionIdRule = append(versionIdRule, versionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "MetronionCreated", metronionIdRule, versionIdRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTMetronionCreatedIterator{contract: _MetronionNFT.contract, event: "MetronionCreated", logs: logs, sub: sub}, nil
}

// WatchMetronionCreated is a free log subscription operation binding the contract event 0xc6ba76065f5132d4931c49972c276c27327502e3127e19e9499170265674da48.
//
// Solidity: event MetronionCreated(uint256 indexed metronionId, uint256 indexed versionId, address to)
func (_MetronionNFT *MetronionNFTFilterer) WatchMetronionCreated(opts *bind.WatchOpts, sink chan<- *MetronionNFTMetronionCreated, metronionId []*big.Int, versionId []*big.Int) (event.Subscription, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}
	var versionIdRule []interface{}
	for _, versionIdItem := range versionId {
		versionIdRule = append(versionIdRule, versionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "MetronionCreated", metronionIdRule, versionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTMetronionCreated)
				if err := _MetronionNFT.contract.UnpackLog(event, "MetronionCreated", log); err != nil {
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

// ParseMetronionCreated is a log parse operation binding the contract event 0xc6ba76065f5132d4931c49972c276c27327502e3127e19e9499170265674da48.
//
// Solidity: event MetronionCreated(uint256 indexed metronionId, uint256 indexed versionId, address to)
func (_MetronionNFT *MetronionNFTFilterer) ParseMetronionCreated(log types.Log) (*MetronionNFTMetronionCreated, error) {
	event := new(MetronionNFTMetronionCreated)
	if err := _MetronionNFT.contract.UnpackLog(event, "MetronionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the MetronionNFT contract.
type MetronionNFTNameChangedIterator struct {
	Event *MetronionNFTNameChanged // Event containing the contract specifics and raw log

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
func (it *MetronionNFTNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTNameChanged)
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
		it.Event = new(MetronionNFTNameChanged)
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
func (it *MetronionNFTNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTNameChanged represents a NameChanged event raised by the MetronionNFT contract.
type MetronionNFTNameChanged struct {
	MetronionId *big.Int
	NewName     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0x8edfa912e70e283a8ef6d6f52cd1faef9690ff989eff2f11a134e8478ba7b28b.
//
// Solidity: event NameChanged(uint256 indexed metronionId, string newName)
func (_MetronionNFT *MetronionNFTFilterer) FilterNameChanged(opts *bind.FilterOpts, metronionId []*big.Int) (*MetronionNFTNameChangedIterator, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "NameChanged", metronionIdRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTNameChangedIterator{contract: _MetronionNFT.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0x8edfa912e70e283a8ef6d6f52cd1faef9690ff989eff2f11a134e8478ba7b28b.
//
// Solidity: event NameChanged(uint256 indexed metronionId, string newName)
func (_MetronionNFT *MetronionNFTFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *MetronionNFTNameChanged, metronionId []*big.Int) (event.Subscription, error) {

	var metronionIdRule []interface{}
	for _, metronionIdItem := range metronionId {
		metronionIdRule = append(metronionIdRule, metronionIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "NameChanged", metronionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTNameChanged)
				if err := _MetronionNFT.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0x8edfa912e70e283a8ef6d6f52cd1faef9690ff989eff2f11a134e8478ba7b28b.
//
// Solidity: event NameChanged(uint256 indexed metronionId, string newName)
func (_MetronionNFT *MetronionNFTFilterer) ParseNameChanged(log types.Log) (*MetronionNFTNameChanged, error) {
	event := new(MetronionNFTNameChanged)
	if err := _MetronionNFT.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTNewVersionAddedIterator is returned from FilterNewVersionAdded and is used to iterate over the raw logs and unpacked data for NewVersionAdded events raised by the MetronionNFT contract.
type MetronionNFTNewVersionAddedIterator struct {
	Event *MetronionNFTNewVersionAdded // Event containing the contract specifics and raw log

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
func (it *MetronionNFTNewVersionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTNewVersionAdded)
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
		it.Event = new(MetronionNFTNewVersionAdded)
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
func (it *MetronionNFTNewVersionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTNewVersionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTNewVersionAdded represents a NewVersionAdded event raised by the MetronionNFT contract.
type MetronionNFTNewVersionAdded struct {
	VersionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewVersionAdded is a free log retrieval operation binding the contract event 0xf5177867b3cf937deaac25da8e7733f654828440429b272c4171f20e50129bf7.
//
// Solidity: event NewVersionAdded(uint256 versionId)
func (_MetronionNFT *MetronionNFTFilterer) FilterNewVersionAdded(opts *bind.FilterOpts) (*MetronionNFTNewVersionAddedIterator, error) {

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "NewVersionAdded")
	if err != nil {
		return nil, err
	}
	return &MetronionNFTNewVersionAddedIterator{contract: _MetronionNFT.contract, event: "NewVersionAdded", logs: logs, sub: sub}, nil
}

// WatchNewVersionAdded is a free log subscription operation binding the contract event 0xf5177867b3cf937deaac25da8e7733f654828440429b272c4171f20e50129bf7.
//
// Solidity: event NewVersionAdded(uint256 versionId)
func (_MetronionNFT *MetronionNFTFilterer) WatchNewVersionAdded(opts *bind.WatchOpts, sink chan<- *MetronionNFTNewVersionAdded) (event.Subscription, error) {

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "NewVersionAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTNewVersionAdded)
				if err := _MetronionNFT.contract.UnpackLog(event, "NewVersionAdded", log); err != nil {
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

// ParseNewVersionAdded is a log parse operation binding the contract event 0xf5177867b3cf937deaac25da8e7733f654828440429b272c4171f20e50129bf7.
//
// Solidity: event NewVersionAdded(uint256 versionId)
func (_MetronionNFT *MetronionNFTFilterer) ParseNewVersionAdded(log types.Log) (*MetronionNFTNewVersionAdded, error) {
	event := new(MetronionNFTNewVersionAdded)
	if err := _MetronionNFT.contract.UnpackLog(event, "NewVersionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetronionNFT contract.
type MetronionNFTOwnershipTransferredIterator struct {
	Event *MetronionNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MetronionNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTOwnershipTransferred)
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
		it.Event = new(MetronionNFTOwnershipTransferred)
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
func (it *MetronionNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTOwnershipTransferred represents a OwnershipTransferred event raised by the MetronionNFT contract.
type MetronionNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetronionNFT *MetronionNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetronionNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTOwnershipTransferredIterator{contract: _MetronionNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetronionNFT *MetronionNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetronionNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTOwnershipTransferred)
				if err := _MetronionNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MetronionNFT *MetronionNFTFilterer) ParseOwnershipTransferred(log types.Log) (*MetronionNFTOwnershipTransferred, error) {
	event := new(MetronionNFTOwnershipTransferred)
	if err := _MetronionNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTStartingIndexFinalizedIterator is returned from FilterStartingIndexFinalized and is used to iterate over the raw logs and unpacked data for StartingIndexFinalized events raised by the MetronionNFT contract.
type MetronionNFTStartingIndexFinalizedIterator struct {
	Event *MetronionNFTStartingIndexFinalized // Event containing the contract specifics and raw log

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
func (it *MetronionNFTStartingIndexFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTStartingIndexFinalized)
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
		it.Event = new(MetronionNFTStartingIndexFinalized)
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
func (it *MetronionNFTStartingIndexFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTStartingIndexFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTStartingIndexFinalized represents a StartingIndexFinalized event raised by the MetronionNFT contract.
type MetronionNFTStartingIndexFinalized struct {
	VersionId     *big.Int
	StartingIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStartingIndexFinalized is a free log retrieval operation binding the contract event 0x2cef2ac652db7d723fb9cffa5df9729234275ec9377763568b8337912851d9f7.
//
// Solidity: event StartingIndexFinalized(uint256 versionId, uint256 startingIndex)
func (_MetronionNFT *MetronionNFTFilterer) FilterStartingIndexFinalized(opts *bind.FilterOpts) (*MetronionNFTStartingIndexFinalizedIterator, error) {

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "StartingIndexFinalized")
	if err != nil {
		return nil, err
	}
	return &MetronionNFTStartingIndexFinalizedIterator{contract: _MetronionNFT.contract, event: "StartingIndexFinalized", logs: logs, sub: sub}, nil
}

// WatchStartingIndexFinalized is a free log subscription operation binding the contract event 0x2cef2ac652db7d723fb9cffa5df9729234275ec9377763568b8337912851d9f7.
//
// Solidity: event StartingIndexFinalized(uint256 versionId, uint256 startingIndex)
func (_MetronionNFT *MetronionNFTFilterer) WatchStartingIndexFinalized(opts *bind.WatchOpts, sink chan<- *MetronionNFTStartingIndexFinalized) (event.Subscription, error) {

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "StartingIndexFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTStartingIndexFinalized)
				if err := _MetronionNFT.contract.UnpackLog(event, "StartingIndexFinalized", log); err != nil {
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

// ParseStartingIndexFinalized is a log parse operation binding the contract event 0x2cef2ac652db7d723fb9cffa5df9729234275ec9377763568b8337912851d9f7.
//
// Solidity: event StartingIndexFinalized(uint256 versionId, uint256 startingIndex)
func (_MetronionNFT *MetronionNFTFilterer) ParseStartingIndexFinalized(log types.Log) (*MetronionNFTStartingIndexFinalized, error) {
	event := new(MetronionNFTStartingIndexFinalized)
	if err := _MetronionNFT.contract.UnpackLog(event, "StartingIndexFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MetronionNFT contract.
type MetronionNFTTransferIterator struct {
	Event *MetronionNFTTransfer // Event containing the contract specifics and raw log

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
func (it *MetronionNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTTransfer)
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
		it.Event = new(MetronionNFTTransfer)
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
func (it *MetronionNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTTransfer represents a Transfer event raised by the MetronionNFT contract.
type MetronionNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MetronionNFT *MetronionNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MetronionNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MetronionNFTTransferIterator{contract: _MetronionNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MetronionNFT *MetronionNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetronionNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTTransfer)
				if err := _MetronionNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MetronionNFT *MetronionNFTFilterer) ParseTransfer(log types.Log) (*MetronionNFTTransfer, error) {
	event := new(MetronionNFTTransfer)
	if err := _MetronionNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetronionNFTUpdateBaseURIIterator is returned from FilterUpdateBaseURI and is used to iterate over the raw logs and unpacked data for UpdateBaseURI events raised by the MetronionNFT contract.
type MetronionNFTUpdateBaseURIIterator struct {
	Event *MetronionNFTUpdateBaseURI // Event containing the contract specifics and raw log

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
func (it *MetronionNFTUpdateBaseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetronionNFTUpdateBaseURI)
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
		it.Event = new(MetronionNFTUpdateBaseURI)
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
func (it *MetronionNFTUpdateBaseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetronionNFTUpdateBaseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetronionNFTUpdateBaseURI represents a UpdateBaseURI event raised by the MetronionNFT contract.
type MetronionNFTUpdateBaseURI struct {
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateBaseURI is a free log retrieval operation binding the contract event 0x157d450c8fb1377294d9db75af1de2753efc52d8e5578551d70d2c7d9cd74df9.
//
// Solidity: event UpdateBaseURI(string uri)
func (_MetronionNFT *MetronionNFTFilterer) FilterUpdateBaseURI(opts *bind.FilterOpts) (*MetronionNFTUpdateBaseURIIterator, error) {

	logs, sub, err := _MetronionNFT.contract.FilterLogs(opts, "UpdateBaseURI")
	if err != nil {
		return nil, err
	}
	return &MetronionNFTUpdateBaseURIIterator{contract: _MetronionNFT.contract, event: "UpdateBaseURI", logs: logs, sub: sub}, nil
}

// WatchUpdateBaseURI is a free log subscription operation binding the contract event 0x157d450c8fb1377294d9db75af1de2753efc52d8e5578551d70d2c7d9cd74df9.
//
// Solidity: event UpdateBaseURI(string uri)
func (_MetronionNFT *MetronionNFTFilterer) WatchUpdateBaseURI(opts *bind.WatchOpts, sink chan<- *MetronionNFTUpdateBaseURI) (event.Subscription, error) {

	logs, sub, err := _MetronionNFT.contract.WatchLogs(opts, "UpdateBaseURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetronionNFTUpdateBaseURI)
				if err := _MetronionNFT.contract.UnpackLog(event, "UpdateBaseURI", log); err != nil {
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

// ParseUpdateBaseURI is a log parse operation binding the contract event 0x157d450c8fb1377294d9db75af1de2753efc52d8e5578551d70d2c7d9cd74df9.
//
// Solidity: event UpdateBaseURI(string uri)
func (_MetronionNFT *MetronionNFTFilterer) ParseUpdateBaseURI(log types.Log) (*MetronionNFTUpdateBaseURI, error) {
	event := new(MetronionNFTUpdateBaseURI)
	if err := _MetronionNFT.contract.UnpackLog(event, "UpdateBaseURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
