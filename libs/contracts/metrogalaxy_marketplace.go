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

// IMetroGalaxyMarketplaceAsset is an auto generated low-level Go binding around an user-defined struct.
type IMetroGalaxyMarketplaceAsset struct {
	PriceInWei *big.Int
	Amount     *big.Int
}

// MetroGalaxyMarketplaceABI is the input ABI used to generate the binding from.
const MetroGalaxyMarketplaceABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_acceptedToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AssetBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"AssetDelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AssetListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"AssetOfferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AssetOfferTaken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AssetOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MARKET_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"acceptedAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"delist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getListedAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIMetroGalaxyMarketplace.Asset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getOfferedAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIMetroGalaxyMarketplace.Asset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"}],\"name\":\"isAcceptedAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"listedAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketFeeInBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"offer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offeredAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"takeOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"updateAcceptedAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"updateMarketFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MetroGalaxyMarketplace is an auto generated Go binding around an Ethereum contract.
type MetroGalaxyMarketplace struct {
	MetroGalaxyMarketplaceCaller     // Read-only binding to the contract
	MetroGalaxyMarketplaceTransactor // Write-only binding to the contract
	MetroGalaxyMarketplaceFilterer   // Log filterer for contract events
}

// MetroGalaxyMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetroGalaxyMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetroGalaxyMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetroGalaxyMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetroGalaxyMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetroGalaxyMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetroGalaxyMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetroGalaxyMarketplaceSession struct {
	Contract     *MetroGalaxyMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MetroGalaxyMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetroGalaxyMarketplaceCallerSession struct {
	Contract *MetroGalaxyMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MetroGalaxyMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetroGalaxyMarketplaceTransactorSession struct {
	Contract     *MetroGalaxyMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MetroGalaxyMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetroGalaxyMarketplaceRaw struct {
	Contract *MetroGalaxyMarketplace // Generic contract binding to access the raw methods on
}

// MetroGalaxyMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetroGalaxyMarketplaceCallerRaw struct {
	Contract *MetroGalaxyMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MetroGalaxyMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetroGalaxyMarketplaceTransactorRaw struct {
	Contract *MetroGalaxyMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetroGalaxyMarketplace creates a new instance of MetroGalaxyMarketplace, bound to a specific deployed contract.
func NewMetroGalaxyMarketplace(address common.Address, backend bind.ContractBackend) (*MetroGalaxyMarketplace, error) {
	contract, err := bindMetroGalaxyMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplace{MetroGalaxyMarketplaceCaller: MetroGalaxyMarketplaceCaller{contract: contract}, MetroGalaxyMarketplaceTransactor: MetroGalaxyMarketplaceTransactor{contract: contract}, MetroGalaxyMarketplaceFilterer: MetroGalaxyMarketplaceFilterer{contract: contract}}, nil
}

// NewMetroGalaxyMarketplaceCaller creates a new read-only instance of MetroGalaxyMarketplace, bound to a specific deployed contract.
func NewMetroGalaxyMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MetroGalaxyMarketplaceCaller, error) {
	contract, err := bindMetroGalaxyMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceCaller{contract: contract}, nil
}

// NewMetroGalaxyMarketplaceTransactor creates a new write-only instance of MetroGalaxyMarketplace, bound to a specific deployed contract.
func NewMetroGalaxyMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MetroGalaxyMarketplaceTransactor, error) {
	contract, err := bindMetroGalaxyMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceTransactor{contract: contract}, nil
}

// NewMetroGalaxyMarketplaceFilterer creates a new log filterer instance of MetroGalaxyMarketplace, bound to a specific deployed contract.
func NewMetroGalaxyMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MetroGalaxyMarketplaceFilterer, error) {
	contract, err := bindMetroGalaxyMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceFilterer{contract: contract}, nil
}

// bindMetroGalaxyMarketplace binds a generic wrapper to an already deployed contract.
func bindMetroGalaxyMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetroGalaxyMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetroGalaxyMarketplace.Contract.MetroGalaxyMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.MetroGalaxyMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.MetroGalaxyMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetroGalaxyMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.contract.Transact(opts, method, params...)
}

// MAXMARKETFEE is a free data retrieval call binding the contract method 0xf3f0b6a9.
//
// Solidity: function MAX_MARKET_FEE() view returns(uint256)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) MAXMARKETFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "MAX_MARKET_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMARKETFEE is a free data retrieval call binding the contract method 0xf3f0b6a9.
//
// Solidity: function MAX_MARKET_FEE() view returns(uint256)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) MAXMARKETFEE() (*big.Int, error) {
	return _MetroGalaxyMarketplace.Contract.MAXMARKETFEE(&_MetroGalaxyMarketplace.CallOpts)
}

// MAXMARKETFEE is a free data retrieval call binding the contract method 0xf3f0b6a9.
//
// Solidity: function MAX_MARKET_FEE() view returns(uint256)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) MAXMARKETFEE() (*big.Int, error) {
	return _MetroGalaxyMarketplace.Contract.MAXMARKETFEE(&_MetroGalaxyMarketplace.CallOpts)
}

// AcceptedAssets is a free data retrieval call binding the contract method 0x64275dae.
//
// Solidity: function acceptedAssets(address ) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) AcceptedAssets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "acceptedAssets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptedAssets is a free data retrieval call binding the contract method 0x64275dae.
//
// Solidity: function acceptedAssets(address ) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) AcceptedAssets(arg0 common.Address) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.AcceptedAssets(&_MetroGalaxyMarketplace.CallOpts, arg0)
}

// AcceptedAssets is a free data retrieval call binding the contract method 0x64275dae.
//
// Solidity: function acceptedAssets(address ) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) AcceptedAssets(arg0 common.Address) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.AcceptedAssets(&_MetroGalaxyMarketplace.CallOpts, arg0)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) AcceptedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "acceptedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) AcceptedToken() (common.Address, error) {
	return _MetroGalaxyMarketplace.Contract.AcceptedToken(&_MetroGalaxyMarketplace.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) AcceptedToken() (common.Address, error) {
	return _MetroGalaxyMarketplace.Contract.AcceptedToken(&_MetroGalaxyMarketplace.CallOpts)
}

// GetListedAsset is a free data retrieval call binding the contract method 0xed925aea.
//
// Solidity: function getListedAsset(address assetAddr, uint256 assetId, address account) view returns((uint256,uint256))
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) GetListedAsset(opts *bind.CallOpts, assetAddr common.Address, assetId *big.Int, account common.Address) (IMetroGalaxyMarketplaceAsset, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "getListedAsset", assetAddr, assetId, account)

	if err != nil {
		return *new(IMetroGalaxyMarketplaceAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetroGalaxyMarketplaceAsset)).(*IMetroGalaxyMarketplaceAsset)

	return out0, err

}

// GetListedAsset is a free data retrieval call binding the contract method 0xed925aea.
//
// Solidity: function getListedAsset(address assetAddr, uint256 assetId, address account) view returns((uint256,uint256))
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) GetListedAsset(assetAddr common.Address, assetId *big.Int, account common.Address) (IMetroGalaxyMarketplaceAsset, error) {
	return _MetroGalaxyMarketplace.Contract.GetListedAsset(&_MetroGalaxyMarketplace.CallOpts, assetAddr, assetId, account)
}

// GetListedAsset is a free data retrieval call binding the contract method 0xed925aea.
//
// Solidity: function getListedAsset(address assetAddr, uint256 assetId, address account) view returns((uint256,uint256))
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) GetListedAsset(assetAddr common.Address, assetId *big.Int, account common.Address) (IMetroGalaxyMarketplaceAsset, error) {
	return _MetroGalaxyMarketplace.Contract.GetListedAsset(&_MetroGalaxyMarketplace.CallOpts, assetAddr, assetId, account)
}

// GetOfferedAsset is a free data retrieval call binding the contract method 0xcd7c942b.
//
// Solidity: function getOfferedAsset(address assetAddr, uint256 assetId, address account) view returns((uint256,uint256))
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) GetOfferedAsset(opts *bind.CallOpts, assetAddr common.Address, assetId *big.Int, account common.Address) (IMetroGalaxyMarketplaceAsset, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "getOfferedAsset", assetAddr, assetId, account)

	if err != nil {
		return *new(IMetroGalaxyMarketplaceAsset), err
	}

	out0 := *abi.ConvertType(out[0], new(IMetroGalaxyMarketplaceAsset)).(*IMetroGalaxyMarketplaceAsset)

	return out0, err

}

// GetOfferedAsset is a free data retrieval call binding the contract method 0xcd7c942b.
//
// Solidity: function getOfferedAsset(address assetAddr, uint256 assetId, address account) view returns((uint256,uint256))
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) GetOfferedAsset(assetAddr common.Address, assetId *big.Int, account common.Address) (IMetroGalaxyMarketplaceAsset, error) {
	return _MetroGalaxyMarketplace.Contract.GetOfferedAsset(&_MetroGalaxyMarketplace.CallOpts, assetAddr, assetId, account)
}

// GetOfferedAsset is a free data retrieval call binding the contract method 0xcd7c942b.
//
// Solidity: function getOfferedAsset(address assetAddr, uint256 assetId, address account) view returns((uint256,uint256))
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) GetOfferedAsset(assetAddr common.Address, assetId *big.Int, account common.Address) (IMetroGalaxyMarketplaceAsset, error) {
	return _MetroGalaxyMarketplace.Contract.GetOfferedAsset(&_MetroGalaxyMarketplace.CallOpts, assetAddr, assetId, account)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) GetOperators() ([]common.Address, error) {
	return _MetroGalaxyMarketplace.Contract.GetOperators(&_MetroGalaxyMarketplace.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address[])
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) GetOperators() ([]common.Address, error) {
	return _MetroGalaxyMarketplace.Contract.GetOperators(&_MetroGalaxyMarketplace.CallOpts)
}

// IsAcceptedAsset is a free data retrieval call binding the contract method 0xb3628215.
//
// Solidity: function isAcceptedAsset(address assetAddr) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) IsAcceptedAsset(opts *bind.CallOpts, assetAddr common.Address) (bool, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "isAcceptedAsset", assetAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAcceptedAsset is a free data retrieval call binding the contract method 0xb3628215.
//
// Solidity: function isAcceptedAsset(address assetAddr) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) IsAcceptedAsset(assetAddr common.Address) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.IsAcceptedAsset(&_MetroGalaxyMarketplace.CallOpts, assetAddr)
}

// IsAcceptedAsset is a free data retrieval call binding the contract method 0xb3628215.
//
// Solidity: function isAcceptedAsset(address assetAddr) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) IsAcceptedAsset(assetAddr common.Address) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.IsAcceptedAsset(&_MetroGalaxyMarketplace.CallOpts, assetAddr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) IsOperator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "isOperator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) IsOperator(account common.Address) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.IsOperator(&_MetroGalaxyMarketplace.CallOpts, account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) IsOperator(account common.Address) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.IsOperator(&_MetroGalaxyMarketplace.CallOpts, account)
}

// ListedAssets is a free data retrieval call binding the contract method 0xebf1948a.
//
// Solidity: function listedAssets(bytes32 , address ) view returns(uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) ListedAssets(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	PriceInWei *big.Int
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "listedAssets", arg0, arg1)

	outstruct := new(struct {
		PriceInWei *big.Int
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceInWei = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ListedAssets is a free data retrieval call binding the contract method 0xebf1948a.
//
// Solidity: function listedAssets(bytes32 , address ) view returns(uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) ListedAssets(arg0 [32]byte, arg1 common.Address) (struct {
	PriceInWei *big.Int
	Amount     *big.Int
}, error) {
	return _MetroGalaxyMarketplace.Contract.ListedAssets(&_MetroGalaxyMarketplace.CallOpts, arg0, arg1)
}

// ListedAssets is a free data retrieval call binding the contract method 0xebf1948a.
//
// Solidity: function listedAssets(bytes32 , address ) view returns(uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) ListedAssets(arg0 [32]byte, arg1 common.Address) (struct {
	PriceInWei *big.Int
	Amount     *big.Int
}, error) {
	return _MetroGalaxyMarketplace.Contract.ListedAssets(&_MetroGalaxyMarketplace.CallOpts, arg0, arg1)
}

// MarketFeeInBps is a free data retrieval call binding the contract method 0xa51dabd0.
//
// Solidity: function marketFeeInBps() view returns(uint256)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) MarketFeeInBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "marketFeeInBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketFeeInBps is a free data retrieval call binding the contract method 0xa51dabd0.
//
// Solidity: function marketFeeInBps() view returns(uint256)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) MarketFeeInBps() (*big.Int, error) {
	return _MetroGalaxyMarketplace.Contract.MarketFeeInBps(&_MetroGalaxyMarketplace.CallOpts)
}

// MarketFeeInBps is a free data retrieval call binding the contract method 0xa51dabd0.
//
// Solidity: function marketFeeInBps() view returns(uint256)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) MarketFeeInBps() (*big.Int, error) {
	return _MetroGalaxyMarketplace.Contract.MarketFeeInBps(&_MetroGalaxyMarketplace.CallOpts)
}

// OfferedAssets is a free data retrieval call binding the contract method 0xabe57c73.
//
// Solidity: function offeredAssets(bytes32 , address ) view returns(uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) OfferedAssets(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	PriceInWei *big.Int
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "offeredAssets", arg0, arg1)

	outstruct := new(struct {
		PriceInWei *big.Int
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceInWei = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OfferedAssets is a free data retrieval call binding the contract method 0xabe57c73.
//
// Solidity: function offeredAssets(bytes32 , address ) view returns(uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) OfferedAssets(arg0 [32]byte, arg1 common.Address) (struct {
	PriceInWei *big.Int
	Amount     *big.Int
}, error) {
	return _MetroGalaxyMarketplace.Contract.OfferedAssets(&_MetroGalaxyMarketplace.CallOpts, arg0, arg1)
}

// OfferedAssets is a free data retrieval call binding the contract method 0xabe57c73.
//
// Solidity: function offeredAssets(bytes32 , address ) view returns(uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) OfferedAssets(arg0 [32]byte, arg1 common.Address) (struct {
	PriceInWei *big.Int
	Amount     *big.Int
}, error) {
	return _MetroGalaxyMarketplace.Contract.OfferedAssets(&_MetroGalaxyMarketplace.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) Owner() (common.Address, error) {
	return _MetroGalaxyMarketplace.Contract.Owner(&_MetroGalaxyMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) Owner() (common.Address, error) {
	return _MetroGalaxyMarketplace.Contract.Owner(&_MetroGalaxyMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) Paused() (bool, error) {
	return _MetroGalaxyMarketplace.Contract.Paused(&_MetroGalaxyMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) Paused() (bool, error) {
	return _MetroGalaxyMarketplace.Contract.Paused(&_MetroGalaxyMarketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MetroGalaxyMarketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.SupportsInterface(&_MetroGalaxyMarketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetroGalaxyMarketplace.Contract.SupportsInterface(&_MetroGalaxyMarketplace.CallOpts, interfaceId)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "addOperator", operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.AddOperator(&_MetroGalaxyMarketplace.TransactOpts, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.AddOperator(&_MetroGalaxyMarketplace.TransactOpts, operator)
}

// Buy is a paid mutator transaction binding the contract method 0x800fe083.
//
// Solidity: function buy(address assetAddr, uint256 assetId, address seller, uint256 priceInWei, uint256 amount) payable returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) Buy(opts *bind.TransactOpts, assetAddr common.Address, assetId *big.Int, seller common.Address, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "buy", assetAddr, assetId, seller, priceInWei, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x800fe083.
//
// Solidity: function buy(address assetAddr, uint256 assetId, address seller, uint256 priceInWei, uint256 amount) payable returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) Buy(assetAddr common.Address, assetId *big.Int, seller common.Address, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Buy(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, seller, priceInWei, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x800fe083.
//
// Solidity: function buy(address assetAddr, uint256 assetId, address seller, uint256 priceInWei, uint256 amount) payable returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) Buy(assetAddr common.Address, assetId *big.Int, seller common.Address, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Buy(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, seller, priceInWei, amount)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address assetAddr, uint256 assetId) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) CancelOffer(opts *bind.TransactOpts, assetAddr common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "cancelOffer", assetAddr, assetId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address assetAddr, uint256 assetId) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) CancelOffer(assetAddr common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.CancelOffer(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address assetAddr, uint256 assetId) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) CancelOffer(assetAddr common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.CancelOffer(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId)
}

// Delist is a paid mutator transaction binding the contract method 0xf074258e.
//
// Solidity: function delist(address assetAddr, uint256 assetId) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) Delist(opts *bind.TransactOpts, assetAddr common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "delist", assetAddr, assetId)
}

// Delist is a paid mutator transaction binding the contract method 0xf074258e.
//
// Solidity: function delist(address assetAddr, uint256 assetId) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) Delist(assetAddr common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Delist(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId)
}

// Delist is a paid mutator transaction binding the contract method 0xf074258e.
//
// Solidity: function delist(address assetAddr, uint256 assetId) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) Delist(assetAddr common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Delist(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId)
}

// List is a paid mutator transaction binding the contract method 0xbb74a1c2.
//
// Solidity: function list(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) List(opts *bind.TransactOpts, assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "list", assetAddr, assetId, priceInWei, amount)
}

// List is a paid mutator transaction binding the contract method 0xbb74a1c2.
//
// Solidity: function list(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) List(assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.List(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, priceInWei, amount)
}

// List is a paid mutator transaction binding the contract method 0xbb74a1c2.
//
// Solidity: function list(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) List(assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.List(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, priceInWei, amount)
}

// Offer is a paid mutator transaction binding the contract method 0xf3e962c1.
//
// Solidity: function offer(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount) payable returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) Offer(opts *bind.TransactOpts, assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "offer", assetAddr, assetId, priceInWei, amount)
}

// Offer is a paid mutator transaction binding the contract method 0xf3e962c1.
//
// Solidity: function offer(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount) payable returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) Offer(assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Offer(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, priceInWei, amount)
}

// Offer is a paid mutator transaction binding the contract method 0xf3e962c1.
//
// Solidity: function offer(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount) payable returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) Offer(assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Offer(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, priceInWei, amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.OnERC1155BatchReceived(&_MetroGalaxyMarketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.OnERC1155BatchReceived(&_MetroGalaxyMarketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.OnERC1155Received(&_MetroGalaxyMarketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.OnERC1155Received(&_MetroGalaxyMarketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.OnERC721Received(&_MetroGalaxyMarketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.OnERC721Received(&_MetroGalaxyMarketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) Pause() (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Pause(&_MetroGalaxyMarketplace.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) Pause() (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Pause(&_MetroGalaxyMarketplace.TransactOpts)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.RemoveOperator(&_MetroGalaxyMarketplace.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.RemoveOperator(&_MetroGalaxyMarketplace.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.RenounceOwnership(&_MetroGalaxyMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.RenounceOwnership(&_MetroGalaxyMarketplace.TransactOpts)
}

// TakeOffer is a paid mutator transaction binding the contract method 0xa393ffcb.
//
// Solidity: function takeOffer(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount, address buyer) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) TakeOffer(opts *bind.TransactOpts, assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "takeOffer", assetAddr, assetId, priceInWei, amount, buyer)
}

// TakeOffer is a paid mutator transaction binding the contract method 0xa393ffcb.
//
// Solidity: function takeOffer(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount, address buyer) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) TakeOffer(assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.TakeOffer(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, priceInWei, amount, buyer)
}

// TakeOffer is a paid mutator transaction binding the contract method 0xa393ffcb.
//
// Solidity: function takeOffer(address assetAddr, uint256 assetId, uint256 priceInWei, uint256 amount, address buyer) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) TakeOffer(assetAddr common.Address, assetId *big.Int, priceInWei *big.Int, amount *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.TakeOffer(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, assetId, priceInWei, amount, buyer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.TransferOwnership(&_MetroGalaxyMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.TransferOwnership(&_MetroGalaxyMarketplace.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) Unpause() (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Unpause(&_MetroGalaxyMarketplace.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) Unpause() (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.Unpause(&_MetroGalaxyMarketplace.TransactOpts)
}

// UpdateAcceptedAsset is a paid mutator transaction binding the contract method 0x353b74d2.
//
// Solidity: function updateAcceptedAsset(address assetAddr, bool isSupported) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) UpdateAcceptedAsset(opts *bind.TransactOpts, assetAddr common.Address, isSupported bool) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "updateAcceptedAsset", assetAddr, isSupported)
}

// UpdateAcceptedAsset is a paid mutator transaction binding the contract method 0x353b74d2.
//
// Solidity: function updateAcceptedAsset(address assetAddr, bool isSupported) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) UpdateAcceptedAsset(assetAddr common.Address, isSupported bool) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.UpdateAcceptedAsset(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, isSupported)
}

// UpdateAcceptedAsset is a paid mutator transaction binding the contract method 0x353b74d2.
//
// Solidity: function updateAcceptedAsset(address assetAddr, bool isSupported) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) UpdateAcceptedAsset(assetAddr common.Address, isSupported bool) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.UpdateAcceptedAsset(&_MetroGalaxyMarketplace.TransactOpts, assetAddr, isSupported)
}

// UpdateMarketFeeBps is a paid mutator transaction binding the contract method 0xb51cdec4.
//
// Solidity: function updateMarketFeeBps(uint256 feeBps) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactor) UpdateMarketFeeBps(opts *bind.TransactOpts, feeBps *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.contract.Transact(opts, "updateMarketFeeBps", feeBps)
}

// UpdateMarketFeeBps is a paid mutator transaction binding the contract method 0xb51cdec4.
//
// Solidity: function updateMarketFeeBps(uint256 feeBps) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceSession) UpdateMarketFeeBps(feeBps *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.UpdateMarketFeeBps(&_MetroGalaxyMarketplace.TransactOpts, feeBps)
}

// UpdateMarketFeeBps is a paid mutator transaction binding the contract method 0xb51cdec4.
//
// Solidity: function updateMarketFeeBps(uint256 feeBps) returns()
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceTransactorSession) UpdateMarketFeeBps(feeBps *big.Int) (*types.Transaction, error) {
	return _MetroGalaxyMarketplace.Contract.UpdateMarketFeeBps(&_MetroGalaxyMarketplace.TransactOpts, feeBps)
}

// MetroGalaxyMarketplaceAssetBoughtIterator is returned from FilterAssetBought and is used to iterate over the raw logs and unpacked data for AssetBought events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetBoughtIterator struct {
	Event *MetroGalaxyMarketplaceAssetBought // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceAssetBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceAssetBought)
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
		it.Event = new(MetroGalaxyMarketplaceAssetBought)
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
func (it *MetroGalaxyMarketplaceAssetBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceAssetBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceAssetBought represents a AssetBought event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetBought struct {
	AssetAddr  common.Address
	AssetId    *big.Int
	Buyer      common.Address
	Seller     common.Address
	PriceInWei *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetBought is a free log retrieval operation binding the contract event 0xe9fd766ab2c0a20bea80fe8dbc2721ed1aaabf344405bb0402a1f6091ceb1568.
//
// Solidity: event AssetBought(address indexed assetAddr, uint256 indexed assetId, address buyer, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterAssetBought(opts *bind.FilterOpts, assetAddr []common.Address, assetId []*big.Int) (*MetroGalaxyMarketplaceAssetBoughtIterator, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "AssetBought", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceAssetBoughtIterator{contract: _MetroGalaxyMarketplace.contract, event: "AssetBought", logs: logs, sub: sub}, nil
}

// WatchAssetBought is a free log subscription operation binding the contract event 0xe9fd766ab2c0a20bea80fe8dbc2721ed1aaabf344405bb0402a1f6091ceb1568.
//
// Solidity: event AssetBought(address indexed assetAddr, uint256 indexed assetId, address buyer, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchAssetBought(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceAssetBought, assetAddr []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "AssetBought", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceAssetBought)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetBought", log); err != nil {
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

// ParseAssetBought is a log parse operation binding the contract event 0xe9fd766ab2c0a20bea80fe8dbc2721ed1aaabf344405bb0402a1f6091ceb1568.
//
// Solidity: event AssetBought(address indexed assetAddr, uint256 indexed assetId, address buyer, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseAssetBought(log types.Log) (*MetroGalaxyMarketplaceAssetBought, error) {
	event := new(MetroGalaxyMarketplaceAssetBought)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplaceAssetDelistedIterator is returned from FilterAssetDelisted and is used to iterate over the raw logs and unpacked data for AssetDelisted events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetDelistedIterator struct {
	Event *MetroGalaxyMarketplaceAssetDelisted // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceAssetDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceAssetDelisted)
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
		it.Event = new(MetroGalaxyMarketplaceAssetDelisted)
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
func (it *MetroGalaxyMarketplaceAssetDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceAssetDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceAssetDelisted represents a AssetDelisted event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetDelisted struct {
	AssetAddr common.Address
	AssetId   *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetDelisted is a free log retrieval operation binding the contract event 0xfdfc049199a375908fbeedd1e1c146402d80f4ee0989d0982b785c6c2118af1c.
//
// Solidity: event AssetDelisted(address indexed assetAddr, uint256 indexed assetId, address seller)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterAssetDelisted(opts *bind.FilterOpts, assetAddr []common.Address, assetId []*big.Int) (*MetroGalaxyMarketplaceAssetDelistedIterator, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "AssetDelisted", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceAssetDelistedIterator{contract: _MetroGalaxyMarketplace.contract, event: "AssetDelisted", logs: logs, sub: sub}, nil
}

// WatchAssetDelisted is a free log subscription operation binding the contract event 0xfdfc049199a375908fbeedd1e1c146402d80f4ee0989d0982b785c6c2118af1c.
//
// Solidity: event AssetDelisted(address indexed assetAddr, uint256 indexed assetId, address seller)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchAssetDelisted(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceAssetDelisted, assetAddr []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "AssetDelisted", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceAssetDelisted)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
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

// ParseAssetDelisted is a log parse operation binding the contract event 0xfdfc049199a375908fbeedd1e1c146402d80f4ee0989d0982b785c6c2118af1c.
//
// Solidity: event AssetDelisted(address indexed assetAddr, uint256 indexed assetId, address seller)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseAssetDelisted(log types.Log) (*MetroGalaxyMarketplaceAssetDelisted, error) {
	event := new(MetroGalaxyMarketplaceAssetDelisted)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplaceAssetListedIterator is returned from FilterAssetListed and is used to iterate over the raw logs and unpacked data for AssetListed events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetListedIterator struct {
	Event *MetroGalaxyMarketplaceAssetListed // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceAssetListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceAssetListed)
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
		it.Event = new(MetroGalaxyMarketplaceAssetListed)
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
func (it *MetroGalaxyMarketplaceAssetListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceAssetListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceAssetListed represents a AssetListed event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetListed struct {
	AssetAddr  common.Address
	AssetId    *big.Int
	Seller     common.Address
	PriceInWei *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetListed is a free log retrieval operation binding the contract event 0x9f54cb9eaa8a1fc4ed65eda03b51c9337a4fecf46621d4d9080cc4b32bef49ad.
//
// Solidity: event AssetListed(address indexed assetAddr, uint256 indexed assetId, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterAssetListed(opts *bind.FilterOpts, assetAddr []common.Address, assetId []*big.Int) (*MetroGalaxyMarketplaceAssetListedIterator, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "AssetListed", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceAssetListedIterator{contract: _MetroGalaxyMarketplace.contract, event: "AssetListed", logs: logs, sub: sub}, nil
}

// WatchAssetListed is a free log subscription operation binding the contract event 0x9f54cb9eaa8a1fc4ed65eda03b51c9337a4fecf46621d4d9080cc4b32bef49ad.
//
// Solidity: event AssetListed(address indexed assetAddr, uint256 indexed assetId, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchAssetListed(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceAssetListed, assetAddr []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "AssetListed", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceAssetListed)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetListed", log); err != nil {
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

// ParseAssetListed is a log parse operation binding the contract event 0x9f54cb9eaa8a1fc4ed65eda03b51c9337a4fecf46621d4d9080cc4b32bef49ad.
//
// Solidity: event AssetListed(address indexed assetAddr, uint256 indexed assetId, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseAssetListed(log types.Log) (*MetroGalaxyMarketplaceAssetListed, error) {
	event := new(MetroGalaxyMarketplaceAssetListed)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplaceAssetOfferCancelledIterator is returned from FilterAssetOfferCancelled and is used to iterate over the raw logs and unpacked data for AssetOfferCancelled events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetOfferCancelledIterator struct {
	Event *MetroGalaxyMarketplaceAssetOfferCancelled // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceAssetOfferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceAssetOfferCancelled)
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
		it.Event = new(MetroGalaxyMarketplaceAssetOfferCancelled)
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
func (it *MetroGalaxyMarketplaceAssetOfferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceAssetOfferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceAssetOfferCancelled represents a AssetOfferCancelled event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetOfferCancelled struct {
	AssetAddr common.Address
	AssetId   *big.Int
	Buyer     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetOfferCancelled is a free log retrieval operation binding the contract event 0x6a190932629fd6e818979ab2706bd297b02f469d7e24a9c48455eaea7c5121bb.
//
// Solidity: event AssetOfferCancelled(address indexed assetAddr, uint256 indexed assetId, address buyer)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterAssetOfferCancelled(opts *bind.FilterOpts, assetAddr []common.Address, assetId []*big.Int) (*MetroGalaxyMarketplaceAssetOfferCancelledIterator, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "AssetOfferCancelled", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceAssetOfferCancelledIterator{contract: _MetroGalaxyMarketplace.contract, event: "AssetOfferCancelled", logs: logs, sub: sub}, nil
}

// WatchAssetOfferCancelled is a free log subscription operation binding the contract event 0x6a190932629fd6e818979ab2706bd297b02f469d7e24a9c48455eaea7c5121bb.
//
// Solidity: event AssetOfferCancelled(address indexed assetAddr, uint256 indexed assetId, address buyer)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchAssetOfferCancelled(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceAssetOfferCancelled, assetAddr []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "AssetOfferCancelled", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceAssetOfferCancelled)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetOfferCancelled", log); err != nil {
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

// ParseAssetOfferCancelled is a log parse operation binding the contract event 0x6a190932629fd6e818979ab2706bd297b02f469d7e24a9c48455eaea7c5121bb.
//
// Solidity: event AssetOfferCancelled(address indexed assetAddr, uint256 indexed assetId, address buyer)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseAssetOfferCancelled(log types.Log) (*MetroGalaxyMarketplaceAssetOfferCancelled, error) {
	event := new(MetroGalaxyMarketplaceAssetOfferCancelled)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetOfferCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplaceAssetOfferTakenIterator is returned from FilterAssetOfferTaken and is used to iterate over the raw logs and unpacked data for AssetOfferTaken events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetOfferTakenIterator struct {
	Event *MetroGalaxyMarketplaceAssetOfferTaken // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceAssetOfferTakenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceAssetOfferTaken)
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
		it.Event = new(MetroGalaxyMarketplaceAssetOfferTaken)
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
func (it *MetroGalaxyMarketplaceAssetOfferTakenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceAssetOfferTakenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceAssetOfferTaken represents a AssetOfferTaken event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetOfferTaken struct {
	AssetAddr  common.Address
	AssetId    *big.Int
	Buyer      common.Address
	Seller     common.Address
	PriceInWei *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetOfferTaken is a free log retrieval operation binding the contract event 0xfa7e322ffa32fc74af97729c40c9b4a44431b78bdd7754daed01f3e0e6ba5336.
//
// Solidity: event AssetOfferTaken(address indexed assetAddr, uint256 indexed assetId, address buyer, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterAssetOfferTaken(opts *bind.FilterOpts, assetAddr []common.Address, assetId []*big.Int) (*MetroGalaxyMarketplaceAssetOfferTakenIterator, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "AssetOfferTaken", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceAssetOfferTakenIterator{contract: _MetroGalaxyMarketplace.contract, event: "AssetOfferTaken", logs: logs, sub: sub}, nil
}

// WatchAssetOfferTaken is a free log subscription operation binding the contract event 0xfa7e322ffa32fc74af97729c40c9b4a44431b78bdd7754daed01f3e0e6ba5336.
//
// Solidity: event AssetOfferTaken(address indexed assetAddr, uint256 indexed assetId, address buyer, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchAssetOfferTaken(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceAssetOfferTaken, assetAddr []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "AssetOfferTaken", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceAssetOfferTaken)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetOfferTaken", log); err != nil {
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

// ParseAssetOfferTaken is a log parse operation binding the contract event 0xfa7e322ffa32fc74af97729c40c9b4a44431b78bdd7754daed01f3e0e6ba5336.
//
// Solidity: event AssetOfferTaken(address indexed assetAddr, uint256 indexed assetId, address buyer, address seller, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseAssetOfferTaken(log types.Log) (*MetroGalaxyMarketplaceAssetOfferTaken, error) {
	event := new(MetroGalaxyMarketplaceAssetOfferTaken)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetOfferTaken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplaceAssetOfferedIterator is returned from FilterAssetOffered and is used to iterate over the raw logs and unpacked data for AssetOffered events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetOfferedIterator struct {
	Event *MetroGalaxyMarketplaceAssetOffered // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceAssetOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceAssetOffered)
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
		it.Event = new(MetroGalaxyMarketplaceAssetOffered)
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
func (it *MetroGalaxyMarketplaceAssetOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceAssetOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceAssetOffered represents a AssetOffered event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceAssetOffered struct {
	AssetAddr  common.Address
	AssetId    *big.Int
	Buyer      common.Address
	PriceInWei *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetOffered is a free log retrieval operation binding the contract event 0x49282849d6127d4ebc65a657f1eaa0a31e3c245addd5a2863e32c29732f2fa2e.
//
// Solidity: event AssetOffered(address indexed assetAddr, uint256 indexed assetId, address buyer, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterAssetOffered(opts *bind.FilterOpts, assetAddr []common.Address, assetId []*big.Int) (*MetroGalaxyMarketplaceAssetOfferedIterator, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "AssetOffered", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceAssetOfferedIterator{contract: _MetroGalaxyMarketplace.contract, event: "AssetOffered", logs: logs, sub: sub}, nil
}

// WatchAssetOffered is a free log subscription operation binding the contract event 0x49282849d6127d4ebc65a657f1eaa0a31e3c245addd5a2863e32c29732f2fa2e.
//
// Solidity: event AssetOffered(address indexed assetAddr, uint256 indexed assetId, address buyer, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchAssetOffered(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceAssetOffered, assetAddr []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var assetAddrRule []interface{}
	for _, assetAddrItem := range assetAddr {
		assetAddrRule = append(assetAddrRule, assetAddrItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "AssetOffered", assetAddrRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceAssetOffered)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetOffered", log); err != nil {
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

// ParseAssetOffered is a log parse operation binding the contract event 0x49282849d6127d4ebc65a657f1eaa0a31e3c245addd5a2863e32c29732f2fa2e.
//
// Solidity: event AssetOffered(address indexed assetAddr, uint256 indexed assetId, address buyer, uint256 priceInWei, uint256 amount)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseAssetOffered(log types.Log) (*MetroGalaxyMarketplaceAssetOffered, error) {
	event := new(MetroGalaxyMarketplaceAssetOffered)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "AssetOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceOwnershipTransferredIterator struct {
	Event *MetroGalaxyMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceOwnershipTransferred)
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
		it.Event = new(MetroGalaxyMarketplaceOwnershipTransferred)
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
func (it *MetroGalaxyMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetroGalaxyMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceOwnershipTransferredIterator{contract: _MetroGalaxyMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceOwnershipTransferred)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MetroGalaxyMarketplaceOwnershipTransferred, error) {
	event := new(MetroGalaxyMarketplaceOwnershipTransferred)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplacePausedIterator struct {
	Event *MetroGalaxyMarketplacePaused // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplacePaused)
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
		it.Event = new(MetroGalaxyMarketplacePaused)
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
func (it *MetroGalaxyMarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplacePaused represents a Paused event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*MetroGalaxyMarketplacePausedIterator, error) {

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplacePausedIterator{contract: _MetroGalaxyMarketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplacePaused)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParsePaused(log types.Log) (*MetroGalaxyMarketplacePaused, error) {
	event := new(MetroGalaxyMarketplacePaused)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetroGalaxyMarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceUnpausedIterator struct {
	Event *MetroGalaxyMarketplaceUnpaused // Event containing the contract specifics and raw log

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
func (it *MetroGalaxyMarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetroGalaxyMarketplaceUnpaused)
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
		it.Event = new(MetroGalaxyMarketplaceUnpaused)
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
func (it *MetroGalaxyMarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetroGalaxyMarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetroGalaxyMarketplaceUnpaused represents a Unpaused event raised by the MetroGalaxyMarketplace contract.
type MetroGalaxyMarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MetroGalaxyMarketplaceUnpausedIterator, error) {

	logs, sub, err := _MetroGalaxyMarketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MetroGalaxyMarketplaceUnpausedIterator{contract: _MetroGalaxyMarketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MetroGalaxyMarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _MetroGalaxyMarketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetroGalaxyMarketplaceUnpaused)
				if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MetroGalaxyMarketplace *MetroGalaxyMarketplaceFilterer) ParseUnpaused(log types.Log) (*MetroGalaxyMarketplaceUnpaused, error) {
	event := new(MetroGalaxyMarketplaceUnpaused)
	if err := _MetroGalaxyMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
