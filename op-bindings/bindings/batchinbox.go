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
	_ = abi.ConvertType
)

// BatchInboxMetaData contains all meta data concerning the BatchInbox contract.
var BatchInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_versionHashes\",\"type\":\"bytes32[]\"}],\"name\":\"appendSequencerBatch\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_versionHashes\",\"type\":\"bytes32[]\"}],\"name\":\"appendSequencerBatchToL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b50604051610c35380380610c3583398101604081905261003091610050565b6001608052600060a081905260c0526001600160a01b031660e052610080565b60006020828403121561006257600080fd5b81516001600160a01b038116811461007957600080fd5b9392505050565b60805160a05160c05160e051610b776100be60003960008181608901526102d20152600061013c015260006101130152600060ea0152610b776000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806354fd4d50146100515780635e2ff5981461006f578063927ede2d14610084578063a3a544c2146100d0575b600080fd5b6100596100e3565b6040516100669190610739565b60405180910390f35b61008261007d36600461079f565b610186565b005b6100ab7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610066565b6100826100de366004610869565b610348565b606061010e7f0000000000000000000000000000000000000000000000000000000000000000610582565b6101377f0000000000000000000000000000000000000000000000000000000000000000610582565b6101607f0000000000000000000000000000000000000000000000000000000000000000610582565b604051602001610172939291906108ab565b604051602081830303815290604052905090565b6040517fa3a544c2000000000000000000000000000000000000000000000000000000008152309063a3a544c2906101c49085908590600401610921565b60006040518083038186803b1580156101dc57600080fd5b505afa1580156101f0573d6000803e3d6000fd5b5050505060008484906102039190610976565b8383604051602401610216929190610921565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b9061030e9089908590620186a0906004016109be565b600060405180830381600087803b15801561032857600080fd5b505af115801561033c573d6000803e3d6000fd5b50505050505050505050565b806103da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4d757374207061737320696e2061746c65617374206f6e652076657273696f6e60448201527f20686173682e000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60005b8181101561057d5760008060636105788686868181106103ff576103ff610a03565b9050602002013560405160200161041891815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261045091610a32565b6000604051808303818686fa925050503d806000811461048c576040519150601f19603f3d011682016040523d82523d6000602084013e610491565b606091505b5091509150816104fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f53746174696363616c6c206661696c65642e000000000000000000000000000060448201526064016103d1565b6000815111610568576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f52657475726e2064617461206d757374206e6f7420626520656d7074792e000060448201526064016103d1565b5050808061057590610a7d565b9150506103dd565b505050565b6060816000036105c557505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156105ef57806105d981610a7d565b91506105e89050600a83610ae4565b91506105c9565b60008167ffffffffffffffff81111561060a5761060a610af8565b6040519080825280601f01601f191660200182016040528015610634576020820181803683370190505b5090505b84156106b757610649600183610b27565b9150610656600a86610b3e565b610661906030610b52565b60f81b81838151811061067657610676610a03565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506106b0600a86610ae4565b9450610638565b949350505050565b60005b838110156106da5781810151838201526020016106c2565b838111156106e9576000848401525b50505050565b600081518084526107078160208601602086016106bf565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061074c60208301846106ef565b9392505050565b60008083601f84011261076557600080fd5b50813567ffffffffffffffff81111561077d57600080fd5b6020830191508360208260051b850101111561079857600080fd5b9250929050565b6000806000806000606086880312156107b757600080fd5b853573ffffffffffffffffffffffffffffffffffffffff811681146107db57600080fd5b9450602086013567ffffffffffffffff808211156107f857600080fd5b818801915088601f83011261080c57600080fd5b81358181111561081b57600080fd5b89602082850101111561082d57600080fd5b60208301965080955050604088013591508082111561084b57600080fd5b5061085888828901610753565b969995985093965092949392505050565b6000806020838503121561087c57600080fd5b823567ffffffffffffffff81111561089357600080fd5b61089f85828601610753565b90969095509350505050565b600084516108bd8184602089016106bf565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516108f9816001850160208a016106bf565b600192019182015283516109148160028401602088016106bf565b0160020195945050505050565b6020815281602082015260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561095a57600080fd5b8260051b80856040850137600092016040019182525092915050565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156109b65780818660040360031b1b83161692505b505092915050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006109ed60608301856106ef565b905063ffffffff83166040830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008251610a448184602087016106bf565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610aae57610aae610a4e565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610af357610af3610ab5565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610b3957610b39610a4e565b500390565b600082610b4d57610b4d610ab5565b500690565b60008219821115610b6557610b65610a4e565b50019056fea164736f6c634300080f000a",
}

// BatchInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchInboxMetaData.ABI instead.
var BatchInboxABI = BatchInboxMetaData.ABI

// BatchInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchInboxMetaData.Bin instead.
var BatchInboxBin = BatchInboxMetaData.Bin

// DeployBatchInbox deploys a new Ethereum contract, binding an instance of BatchInbox to it.
func DeployBatchInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _messenger common.Address) (common.Address, *types.Transaction, *BatchInbox, error) {
	parsed, err := BatchInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchInboxBin), backend, _messenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchInbox{BatchInboxCaller: BatchInboxCaller{contract: contract}, BatchInboxTransactor: BatchInboxTransactor{contract: contract}, BatchInboxFilterer: BatchInboxFilterer{contract: contract}}, nil
}

// BatchInbox is an auto generated Go binding around an Ethereum contract.
type BatchInbox struct {
	BatchInboxCaller     // Read-only binding to the contract
	BatchInboxTransactor // Write-only binding to the contract
	BatchInboxFilterer   // Log filterer for contract events
}

// BatchInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchInboxSession struct {
	Contract     *BatchInbox       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchInboxCallerSession struct {
	Contract *BatchInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BatchInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchInboxTransactorSession struct {
	Contract     *BatchInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BatchInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchInboxRaw struct {
	Contract *BatchInbox // Generic contract binding to access the raw methods on
}

// BatchInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchInboxCallerRaw struct {
	Contract *BatchInboxCaller // Generic read-only contract binding to access the raw methods on
}

// BatchInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchInboxTransactorRaw struct {
	Contract *BatchInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchInbox creates a new instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInbox(address common.Address, backend bind.ContractBackend) (*BatchInbox, error) {
	contract, err := bindBatchInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchInbox{BatchInboxCaller: BatchInboxCaller{contract: contract}, BatchInboxTransactor: BatchInboxTransactor{contract: contract}, BatchInboxFilterer: BatchInboxFilterer{contract: contract}}, nil
}

// NewBatchInboxCaller creates a new read-only instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInboxCaller(address common.Address, caller bind.ContractCaller) (*BatchInboxCaller, error) {
	contract, err := bindBatchInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchInboxCaller{contract: contract}, nil
}

// NewBatchInboxTransactor creates a new write-only instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchInboxTransactor, error) {
	contract, err := bindBatchInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchInboxTransactor{contract: contract}, nil
}

// NewBatchInboxFilterer creates a new log filterer instance of BatchInbox, bound to a specific deployed contract.
func NewBatchInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchInboxFilterer, error) {
	contract, err := bindBatchInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchInboxFilterer{contract: contract}, nil
}

// bindBatchInbox binds a generic wrapper to an already deployed contract.
func bindBatchInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchInbox *BatchInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchInbox.Contract.BatchInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchInbox *BatchInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchInbox.Contract.BatchInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchInbox *BatchInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchInbox.Contract.BatchInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchInbox *BatchInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchInbox *BatchInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchInbox *BatchInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchInbox.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_BatchInbox *BatchInboxCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_BatchInbox *BatchInboxSession) MESSENGER() (common.Address, error) {
	return _BatchInbox.Contract.MESSENGER(&_BatchInbox.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_BatchInbox *BatchInboxCallerSession) MESSENGER() (common.Address, error) {
	return _BatchInbox.Contract.MESSENGER(&_BatchInbox.CallOpts)
}

// AppendSequencerBatch is a free data retrieval call binding the contract method 0xa3a544c2.
//
// Solidity: function appendSequencerBatch(bytes32[] _versionHashes) view returns()
func (_BatchInbox *BatchInboxCaller) AppendSequencerBatch(opts *bind.CallOpts, _versionHashes [][32]byte) error {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "appendSequencerBatch", _versionHashes)

	if err != nil {
		return err
	}

	return err

}

// AppendSequencerBatch is a free data retrieval call binding the contract method 0xa3a544c2.
//
// Solidity: function appendSequencerBatch(bytes32[] _versionHashes) view returns()
func (_BatchInbox *BatchInboxSession) AppendSequencerBatch(_versionHashes [][32]byte) error {
	return _BatchInbox.Contract.AppendSequencerBatch(&_BatchInbox.CallOpts, _versionHashes)
}

// AppendSequencerBatch is a free data retrieval call binding the contract method 0xa3a544c2.
//
// Solidity: function appendSequencerBatch(bytes32[] _versionHashes) view returns()
func (_BatchInbox *BatchInboxCallerSession) AppendSequencerBatch(_versionHashes [][32]byte) error {
	return _BatchInbox.Contract.AppendSequencerBatch(&_BatchInbox.CallOpts, _versionHashes)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchInbox *BatchInboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchInbox *BatchInboxSession) Version() (string, error) {
	return _BatchInbox.Contract.Version(&_BatchInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BatchInbox *BatchInboxCallerSession) Version() (string, error) {
	return _BatchInbox.Contract.Version(&_BatchInbox.CallOpts)
}

// AppendSequencerBatchToL2 is a paid mutator transaction binding the contract method 0x5e2ff598.
//
// Solidity: function appendSequencerBatchToL2(address _target, bytes _selector, bytes32[] _versionHashes) returns()
func (_BatchInbox *BatchInboxTransactor) AppendSequencerBatchToL2(opts *bind.TransactOpts, _target common.Address, _selector []byte, _versionHashes [][32]byte) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "appendSequencerBatchToL2", _target, _selector, _versionHashes)
}

// AppendSequencerBatchToL2 is a paid mutator transaction binding the contract method 0x5e2ff598.
//
// Solidity: function appendSequencerBatchToL2(address _target, bytes _selector, bytes32[] _versionHashes) returns()
func (_BatchInbox *BatchInboxSession) AppendSequencerBatchToL2(_target common.Address, _selector []byte, _versionHashes [][32]byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.AppendSequencerBatchToL2(&_BatchInbox.TransactOpts, _target, _selector, _versionHashes)
}

// AppendSequencerBatchToL2 is a paid mutator transaction binding the contract method 0x5e2ff598.
//
// Solidity: function appendSequencerBatchToL2(address _target, bytes _selector, bytes32[] _versionHashes) returns()
func (_BatchInbox *BatchInboxTransactorSession) AppendSequencerBatchToL2(_target common.Address, _selector []byte, _versionHashes [][32]byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.AppendSequencerBatchToL2(&_BatchInbox.TransactOpts, _target, _selector, _versionHashes)
}
