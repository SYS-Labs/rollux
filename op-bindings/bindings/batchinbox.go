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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"batchData\",\"type\":\"bytes\"}],\"name\":\"appendSequencerBatch\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"batchData\",\"type\":\"bytes\"}],\"name\":\"appendSequencerBatchToL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b50604051610aa2380380610aa283398101604081905261003091610050565b6001608052600060a081905260c0526001600160a01b031660e052610080565b60006020828403121561006257600080fd5b81516001600160a01b038116811461007957600080fd5b9392505050565b60805160a05160c05160e0516109e36100bf600039600081816089015261022c015260006102fa015260006102d1015260006102a801526109e36000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063465b6fcd1461005157806354fd4d5014610066578063927ede2d1461008457806398f05bb1146100d0575b600080fd5b61006461005f36600461056a565b6100e3565b005b61006e6102a1565b60405161007b9190610706565b60405180910390f35b6100ab7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161007b565b6100646100de366004610720565b610344565b6040517f98f05bb100000000000000000000000000000000000000000000000000000000815230906398f05bb1906101219085908590600401610762565b60006040518083038186803b15801561013957600080fd5b505afa15801561014d573d6000803e3d6000fd5b5050505060008361015d906107af565b8383604051602401610170929190610762565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b906102689088908590620186a0906004016107ff565b600060405180830381600087803b15801561028257600080fd5b505af1158015610296573d6000803e3d6000fd5b505050505050505050565b60606102cc7f00000000000000000000000000000000000000000000000000000000000000006103b5565b6102f57f00000000000000000000000000000000000000000000000000000000000000006103b5565b61031e7f00000000000000000000000000000000000000000000000000000000000000006103b5565b60405160200161033093929190610844565b604051602081830303815290604052905090565b61034f6020826108e9565b1561035957600080fd5b60005b61036760208361092c565b81116103b0576040516020828237602081016040526000806020836063610578fa905080801561004c573d61039b57600080fd5b506103a99050602082610943565b905061035c565b505050565b6060816000036103f857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610422578061040c8161095b565b915061041b9050600a83610993565b91506103fc565b60008167ffffffffffffffff81111561043d5761043d6104f2565b6040519080825280601f01601f191660200182016040528015610467576020820181803683370190505b5090505b84156104ea5761047c60018361092c565b9150610489600a866108e9565b610494906030610943565b60f81b8183815181106104a9576104a96109a7565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506104e3600a86610993565b945061046b565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008083601f84011261053357600080fd5b50813567ffffffffffffffff81111561054b57600080fd5b60208301915083602082850101111561056357600080fd5b9250929050565b6000806000806060858703121561058057600080fd5b843573ffffffffffffffffffffffffffffffffffffffff811681146105a457600080fd5b9350602085013567ffffffffffffffff808211156105c157600080fd5b818701915087601f8301126105d557600080fd5b8135818111156105e7576105e76104f2565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561062d5761062d6104f2565b816040528281528a602084870101111561064657600080fd5b82602086016020830137600060208483010152809750505050604087013591508082111561067357600080fd5b5061068087828801610521565b95989497509550505050565b60005b838110156106a757818101518382015260200161068f565b838111156106b6576000848401525b50505050565b600081518084526106d481602086016020860161068c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061071960208301846106bc565b9392505050565b6000806020838503121561073357600080fd5b823567ffffffffffffffff81111561074a57600080fd5b61075685828601610521565b90969095509350505050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156107f75780818460040360031b1b83161693505b505050919050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061082e60608301856106bc565b905063ffffffff83166040830152949350505050565b6000845161085681846020890161068c565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610892816001850160208a0161068c565b600192019182015283516108ad81600284016020880161068c565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826108f8576108f86108ba565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561093e5761093e6108fd565b500390565b60008219821115610956576109566108fd565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361098c5761098c6108fd565b5060010190565b6000826109a2576109a26108ba565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
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

// AppendSequencerBatch is a free data retrieval call binding the contract method 0x98f05bb1.
//
// Solidity: function appendSequencerBatch(bytes batchData) view returns()
func (_BatchInbox *BatchInboxCaller) AppendSequencerBatch(opts *bind.CallOpts, batchData []byte) error {
	var out []interface{}
	err := _BatchInbox.contract.Call(opts, &out, "appendSequencerBatch", batchData)

	if err != nil {
		return err
	}

	return err

}

// AppendSequencerBatch is a free data retrieval call binding the contract method 0x98f05bb1.
//
// Solidity: function appendSequencerBatch(bytes batchData) view returns()
func (_BatchInbox *BatchInboxSession) AppendSequencerBatch(batchData []byte) error {
	return _BatchInbox.Contract.AppendSequencerBatch(&_BatchInbox.CallOpts, batchData)
}

// AppendSequencerBatch is a free data retrieval call binding the contract method 0x98f05bb1.
//
// Solidity: function appendSequencerBatch(bytes batchData) view returns()
func (_BatchInbox *BatchInboxCallerSession) AppendSequencerBatch(batchData []byte) error {
	return _BatchInbox.Contract.AppendSequencerBatch(&_BatchInbox.CallOpts, batchData)
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

// AppendSequencerBatchToL2 is a paid mutator transaction binding the contract method 0x465b6fcd.
//
// Solidity: function appendSequencerBatchToL2(address _target, bytes _selector, bytes batchData) returns()
func (_BatchInbox *BatchInboxTransactor) AppendSequencerBatchToL2(opts *bind.TransactOpts, _target common.Address, _selector []byte, batchData []byte) (*types.Transaction, error) {
	return _BatchInbox.contract.Transact(opts, "appendSequencerBatchToL2", _target, _selector, batchData)
}

// AppendSequencerBatchToL2 is a paid mutator transaction binding the contract method 0x465b6fcd.
//
// Solidity: function appendSequencerBatchToL2(address _target, bytes _selector, bytes batchData) returns()
func (_BatchInbox *BatchInboxSession) AppendSequencerBatchToL2(_target common.Address, _selector []byte, batchData []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.AppendSequencerBatchToL2(&_BatchInbox.TransactOpts, _target, _selector, batchData)
}

// AppendSequencerBatchToL2 is a paid mutator transaction binding the contract method 0x465b6fcd.
//
// Solidity: function appendSequencerBatchToL2(address _target, bytes _selector, bytes batchData) returns()
func (_BatchInbox *BatchInboxTransactorSession) AppendSequencerBatchToL2(_target common.Address, _selector []byte, batchData []byte) (*types.Transaction, error) {
	return _BatchInbox.Contract.AppendSequencerBatchToL2(&_BatchInbox.TransactOpts, _target, _selector, batchData)
}
