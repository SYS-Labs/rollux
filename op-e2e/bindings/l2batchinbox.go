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

// L2BatchInboxMetaData contains all meta data concerning the L2BatchInbox contract.
var L2BatchInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_otherBridge\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCrossDomainMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appendSequencerBatchFromL1\",\"inputs\":[{\"name\":\"_versionHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"podaMap\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x61012060405234801561001157600080fd5b5060405161094e38038061094e83398101604081905261003091610069565b6001608052600060a081905260c05273420000000000000000000000000000000000000760e0526001600160a01b031661010052610099565b60006020828403121561007b57600080fd5b81516001600160a01b038116811461009257600080fd5b9392505050565b60805160a05160c05160e0516101005161085d6100f16000396000818160d701526101870152600081816101230152818161015d01526101be015260006103bc015260006103930152600061036a015261085d6000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806366c0a4b41161005057806366c0a4b41461009f5780637f46ddb2146100d2578063927ede2d1461011e57600080fd5b80631daf41b01461006c57806354fd4d5014610081575b600080fd5b61007f61007a366004610543565b610145565b005b610089610363565b60405161009691906105e8565b60405180910390f35b6100c26100ad366004610639565b60006020819052908152604090205460ff1681565b6040519015158152602001610096565b6100f97f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610096565b6100f97f000000000000000000000000000000000000000000000000000000000000000081565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561026357507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610227573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024b9190610652565b73ffffffffffffffffffffffffffffffffffffffff16145b6102f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4c324261746368496e626f783a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f746865722062726964676500606482015260840160405180910390fd5b806102fd57600080fd5b60005b8181101561035e57600160008085858581811061031f5761031f61068f565b90506020020135815260200190815260200160002060006101000a81548160ff0219169083151502179055508080610356906106ed565b915050610300565b505050565b606061038e7f0000000000000000000000000000000000000000000000000000000000000000610406565b6103b77f0000000000000000000000000000000000000000000000000000000000000000610406565b6103e07f0000000000000000000000000000000000000000000000000000000000000000610406565b6040516020016103f293929190610725565b604051602081830303815290604052905090565b60608160000361044957505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610473578061045d816106ed565b915061046c9050600a836107ca565b915061044d565b60008167ffffffffffffffff81111561048e5761048e6107de565b6040519080825280601f01601f1916602001820160405280156104b8576020820181803683370190505b5090505b841561053b576104cd60018361080d565b91506104da600a86610824565b6104e5906030610838565b60f81b8183815181106104fa576104fa61068f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610534600a866107ca565b94506104bc565b949350505050565b6000806020838503121561055657600080fd5b823567ffffffffffffffff8082111561056e57600080fd5b818501915085601f83011261058257600080fd5b81358181111561059157600080fd5b8660208260051b85010111156105a657600080fd5b60209290920196919550909350505050565b60005b838110156105d35781810151838201526020016105bb565b838111156105e2576000848401525b50505050565b60208152600082518060208401526106078160408501602087016105b8565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561064b57600080fd5b5035919050565b60006020828403121561066457600080fd5b815173ffffffffffffffffffffffffffffffffffffffff8116811461068857600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361071e5761071e6106be565b5060010190565b600084516107378184602089016105b8565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610773816001850160208a016105b8565b6001920191820152835161078e8160028401602088016105b8565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826107d9576107d961079b565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008282101561081f5761081f6106be565b500390565b6000826108335761083361079b565b500690565b6000821982111561084b5761084b6106be565b50019056fea164736f6c634300080f000a",
}

// L2BatchInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use L2BatchInboxMetaData.ABI instead.
var L2BatchInboxABI = L2BatchInboxMetaData.ABI

// L2BatchInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2BatchInboxMetaData.Bin instead.
var L2BatchInboxBin = L2BatchInboxMetaData.Bin

// DeployL2BatchInbox deploys a new Ethereum contract, binding an instance of L2BatchInbox to it.
func DeployL2BatchInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _otherBridge common.Address) (common.Address, *types.Transaction, *L2BatchInbox, error) {
	parsed, err := L2BatchInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2BatchInboxBin), backend, _otherBridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2BatchInbox{L2BatchInboxCaller: L2BatchInboxCaller{contract: contract}, L2BatchInboxTransactor: L2BatchInboxTransactor{contract: contract}, L2BatchInboxFilterer: L2BatchInboxFilterer{contract: contract}}, nil
}

// L2BatchInbox is an auto generated Go binding around an Ethereum contract.
type L2BatchInbox struct {
	L2BatchInboxCaller     // Read-only binding to the contract
	L2BatchInboxTransactor // Write-only binding to the contract
	L2BatchInboxFilterer   // Log filterer for contract events
}

// L2BatchInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2BatchInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BatchInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2BatchInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BatchInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2BatchInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BatchInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2BatchInboxSession struct {
	Contract     *L2BatchInbox     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2BatchInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2BatchInboxCallerSession struct {
	Contract *L2BatchInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// L2BatchInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2BatchInboxTransactorSession struct {
	Contract     *L2BatchInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2BatchInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2BatchInboxRaw struct {
	Contract *L2BatchInbox // Generic contract binding to access the raw methods on
}

// L2BatchInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2BatchInboxCallerRaw struct {
	Contract *L2BatchInboxCaller // Generic read-only contract binding to access the raw methods on
}

// L2BatchInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2BatchInboxTransactorRaw struct {
	Contract *L2BatchInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2BatchInbox creates a new instance of L2BatchInbox, bound to a specific deployed contract.
func NewL2BatchInbox(address common.Address, backend bind.ContractBackend) (*L2BatchInbox, error) {
	contract, err := bindL2BatchInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2BatchInbox{L2BatchInboxCaller: L2BatchInboxCaller{contract: contract}, L2BatchInboxTransactor: L2BatchInboxTransactor{contract: contract}, L2BatchInboxFilterer: L2BatchInboxFilterer{contract: contract}}, nil
}

// NewL2BatchInboxCaller creates a new read-only instance of L2BatchInbox, bound to a specific deployed contract.
func NewL2BatchInboxCaller(address common.Address, caller bind.ContractCaller) (*L2BatchInboxCaller, error) {
	contract, err := bindL2BatchInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2BatchInboxCaller{contract: contract}, nil
}

// NewL2BatchInboxTransactor creates a new write-only instance of L2BatchInbox, bound to a specific deployed contract.
func NewL2BatchInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*L2BatchInboxTransactor, error) {
	contract, err := bindL2BatchInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2BatchInboxTransactor{contract: contract}, nil
}

// NewL2BatchInboxFilterer creates a new log filterer instance of L2BatchInbox, bound to a specific deployed contract.
func NewL2BatchInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*L2BatchInboxFilterer, error) {
	contract, err := bindL2BatchInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2BatchInboxFilterer{contract: contract}, nil
}

// bindL2BatchInbox binds a generic wrapper to an already deployed contract.
func bindL2BatchInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2BatchInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2BatchInbox *L2BatchInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2BatchInbox.Contract.L2BatchInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2BatchInbox *L2BatchInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2BatchInbox.Contract.L2BatchInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2BatchInbox *L2BatchInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2BatchInbox.Contract.L2BatchInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2BatchInbox *L2BatchInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2BatchInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2BatchInbox *L2BatchInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2BatchInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2BatchInbox *L2BatchInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2BatchInbox.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2BatchInbox *L2BatchInboxCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2BatchInbox.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2BatchInbox *L2BatchInboxSession) MESSENGER() (common.Address, error) {
	return _L2BatchInbox.Contract.MESSENGER(&_L2BatchInbox.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2BatchInbox *L2BatchInboxCallerSession) MESSENGER() (common.Address, error) {
	return _L2BatchInbox.Contract.MESSENGER(&_L2BatchInbox.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2BatchInbox *L2BatchInboxCaller) OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2BatchInbox.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2BatchInbox *L2BatchInboxSession) OTHERBRIDGE() (common.Address, error) {
	return _L2BatchInbox.Contract.OTHERBRIDGE(&_L2BatchInbox.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2BatchInbox *L2BatchInboxCallerSession) OTHERBRIDGE() (common.Address, error) {
	return _L2BatchInbox.Contract.OTHERBRIDGE(&_L2BatchInbox.CallOpts)
}

// PodaMap is a free data retrieval call binding the contract method 0x66c0a4b4.
//
// Solidity: function podaMap(bytes32 ) view returns(bool)
func (_L2BatchInbox *L2BatchInboxCaller) PodaMap(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2BatchInbox.contract.Call(opts, &out, "podaMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PodaMap is a free data retrieval call binding the contract method 0x66c0a4b4.
//
// Solidity: function podaMap(bytes32 ) view returns(bool)
func (_L2BatchInbox *L2BatchInboxSession) PodaMap(arg0 [32]byte) (bool, error) {
	return _L2BatchInbox.Contract.PodaMap(&_L2BatchInbox.CallOpts, arg0)
}

// PodaMap is a free data retrieval call binding the contract method 0x66c0a4b4.
//
// Solidity: function podaMap(bytes32 ) view returns(bool)
func (_L2BatchInbox *L2BatchInboxCallerSession) PodaMap(arg0 [32]byte) (bool, error) {
	return _L2BatchInbox.Contract.PodaMap(&_L2BatchInbox.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2BatchInbox *L2BatchInboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2BatchInbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2BatchInbox *L2BatchInboxSession) Version() (string, error) {
	return _L2BatchInbox.Contract.Version(&_L2BatchInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2BatchInbox *L2BatchInboxCallerSession) Version() (string, error) {
	return _L2BatchInbox.Contract.Version(&_L2BatchInbox.CallOpts)
}

// AppendSequencerBatchFromL1 is a paid mutator transaction binding the contract method 0x1daf41b0.
//
// Solidity: function appendSequencerBatchFromL1(bytes32[] _versionHashes) returns()
func (_L2BatchInbox *L2BatchInboxTransactor) AppendSequencerBatchFromL1(opts *bind.TransactOpts, _versionHashes [][32]byte) (*types.Transaction, error) {
	return _L2BatchInbox.contract.Transact(opts, "appendSequencerBatchFromL1", _versionHashes)
}

// AppendSequencerBatchFromL1 is a paid mutator transaction binding the contract method 0x1daf41b0.
//
// Solidity: function appendSequencerBatchFromL1(bytes32[] _versionHashes) returns()
func (_L2BatchInbox *L2BatchInboxSession) AppendSequencerBatchFromL1(_versionHashes [][32]byte) (*types.Transaction, error) {
	return _L2BatchInbox.Contract.AppendSequencerBatchFromL1(&_L2BatchInbox.TransactOpts, _versionHashes)
}

// AppendSequencerBatchFromL1 is a paid mutator transaction binding the contract method 0x1daf41b0.
//
// Solidity: function appendSequencerBatchFromL1(bytes32[] _versionHashes) returns()
func (_L2BatchInbox *L2BatchInboxTransactorSession) AppendSequencerBatchFromL1(_versionHashes [][32]byte) (*types.Transaction, error) {
	return _L2BatchInbox.Contract.AppendSequencerBatchFromL1(&_L2BatchInbox.TransactOpts, _versionHashes)
}
