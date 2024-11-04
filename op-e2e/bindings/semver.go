// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SemverMetaData contains all meta data concerning the Semver contract.
var SemverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_major\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_patch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60e060405234801561001057600080fd5b5060405161051738038061051783398101604081905261002f91610040565b60809290925260a05260c05261006e565b60008060006060848603121561005557600080fd5b8351925060208401519150604084015190509250925092565b60805160a05160c05161047d61009a600039600060a701526000607e015260006055015261047d6000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806354fd4d5014610030575b600080fd5b61003861004e565b604051610045919061025e565b60405180910390f35b60606100797f00000000000000000000000000000000000000000000000000000000000000006100f1565b6100a27f00000000000000000000000000000000000000000000000000000000000000006100f1565b6100cb7f00000000000000000000000000000000000000000000000000000000000000006100f1565b6040516020016100dd939291906102af565b604051602081830303815290604052905090565b60608160000361013457505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561015e578061014881610354565b91506101579050600a836103bb565b9150610138565b60008167ffffffffffffffff811115610179576101796103cf565b6040519080825280601f01601f1916602001820160405280156101a3576020820181803683370190505b5090505b8415610226576101b86001836103fe565b91506101c5600a86610415565b6101d0906030610429565b60f81b8183815181106101e5576101e5610441565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061021f600a866103bb565b94506101a7565b949350505050565b60005b83811015610249578181015183820152602001610231565b83811115610258576000848401525b50505050565b602081526000825180602084015261027d81604085016020870161022e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600084516102c181846020890161022e565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516102fd816001850160208a0161022e565b6001920191820152835161031881600284016020880161022e565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361038557610385610325565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826103ca576103ca61038c565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008282101561041057610410610325565b500390565b6000826104245761042461038c565b500690565b6000821982111561043c5761043c610325565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
}

// SemverABI is the input ABI used to generate the binding from.
// Deprecated: Use SemverMetaData.ABI instead.
var SemverABI = SemverMetaData.ABI

// SemverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SemverMetaData.Bin instead.
var SemverBin = SemverMetaData.Bin

// DeploySemver deploys a new Ethereum contract, binding an instance of Semver to it.
func DeploySemver(auth *bind.TransactOpts, backend bind.ContractBackend, _major *big.Int, _minor *big.Int, _patch *big.Int) (common.Address, *types.Transaction, *Semver, error) {
	parsed, err := SemverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SemverBin), backend, _major, _minor, _patch)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Semver{SemverCaller: SemverCaller{contract: contract}, SemverTransactor: SemverTransactor{contract: contract}, SemverFilterer: SemverFilterer{contract: contract}}, nil
}

// Semver is an auto generated Go binding around an Ethereum contract.
type Semver struct {
	SemverCaller     // Read-only binding to the contract
	SemverTransactor // Write-only binding to the contract
	SemverFilterer   // Log filterer for contract events
}

// SemverCaller is an auto generated read-only Go binding around an Ethereum contract.
type SemverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SemverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SemverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SemverSession struct {
	Contract     *Semver           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SemverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SemverCallerSession struct {
	Contract *SemverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SemverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SemverTransactorSession struct {
	Contract     *SemverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SemverRaw is an auto generated low-level Go binding around an Ethereum contract.
type SemverRaw struct {
	Contract *Semver // Generic contract binding to access the raw methods on
}

// SemverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SemverCallerRaw struct {
	Contract *SemverCaller // Generic read-only contract binding to access the raw methods on
}

// SemverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SemverTransactorRaw struct {
	Contract *SemverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSemver creates a new instance of Semver, bound to a specific deployed contract.
func NewSemver(address common.Address, backend bind.ContractBackend) (*Semver, error) {
	contract, err := bindSemver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Semver{SemverCaller: SemverCaller{contract: contract}, SemverTransactor: SemverTransactor{contract: contract}, SemverFilterer: SemverFilterer{contract: contract}}, nil
}

// NewSemverCaller creates a new read-only instance of Semver, bound to a specific deployed contract.
func NewSemverCaller(address common.Address, caller bind.ContractCaller) (*SemverCaller, error) {
	contract, err := bindSemver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SemverCaller{contract: contract}, nil
}

// NewSemverTransactor creates a new write-only instance of Semver, bound to a specific deployed contract.
func NewSemverTransactor(address common.Address, transactor bind.ContractTransactor) (*SemverTransactor, error) {
	contract, err := bindSemver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SemverTransactor{contract: contract}, nil
}

// NewSemverFilterer creates a new log filterer instance of Semver, bound to a specific deployed contract.
func NewSemverFilterer(address common.Address, filterer bind.ContractFilterer) (*SemverFilterer, error) {
	contract, err := bindSemver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SemverFilterer{contract: contract}, nil
}

// bindSemver binds a generic wrapper to an already deployed contract.
func bindSemver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SemverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Semver *SemverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Semver.Contract.SemverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Semver *SemverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Semver.Contract.SemverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Semver *SemverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Semver.Contract.SemverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Semver *SemverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Semver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Semver *SemverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Semver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Semver *SemverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Semver.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Semver *SemverCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Semver.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Semver *SemverSession) Version() (string, error) {
	return _Semver.Contract.Version(&_Semver.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Semver *SemverCallerSession) Version() (string, error) {
	return _Semver.Contract.Version(&_Semver.CallOpts)
}
