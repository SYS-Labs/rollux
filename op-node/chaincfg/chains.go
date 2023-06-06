package chaincfg

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
)

var Mainnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x438335a20d98863a4c0c97999eb2481921ccd28553eac6f913af7c12aec04108"),
			Number: 17422590,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xdbf6a80fef073de06add9b0d14026d6e5a86c85f6d102c36d3d8e9cf89c2afd3"),
			Number: 105235063,
		},
		L2Time: 1686068903,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x00D97b2A26Cb85252998fe7B4bd4eC2118bf6B6E"),
			Overhead:    eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
			GasLimit:    30_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      1500,
	SeqWindowSize:          288,
	ChannelTimeout:         24,
	L1ChainID:              big.NewInt(57),
	L2ChainID:              big.NewInt(570),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000420"),
	DepositContractAddress: common.HexToAddress("0xD46Bf6354725bFd4409cd6A952695bFEb213aCB9"),
	L1SystemConfigAddress:  common.HexToAddress("0x739d6e104C717566F65e4Ea711500CE81EF98D42"),
	RegolithTime:           u64Ptr(0),
}

var Goerli = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0xbaaa9a7834d9b5e928eeb36942b96eb64167701e16b9da02a7a5f3aa9c0a216c"),
			Number: 247425,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x045514aee1f089c5acd01ee15995e39a406e92586495bfa4429aa93b9f6f1067"),
			Number: 0,
		},
		L2Time: 1673550516,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x00d97b2a26cb85252998fe7b4bd4ec2118bf6b6e"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    25_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      1500,
	SeqWindowSize:          288,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(5700),
	L2ChainID:              big.NewInt(57000),
	BatchInboxAddress:      common.HexToAddress("0x678255ae6b5c4ba0e6206a8e70b59b874f20bc9c"),
	DepositContractAddress: common.HexToAddress("0x61200b9fcbb421afd0bb5a732fe48ec98482e39c"),
	L1SystemConfigAddress:  common.HexToAddress("0xd8daedc48ca71e20feb81cc3e51c9e3a89a3c84b"),
	RegolithTime:           u64Ptr(1679079600),
}

var NetworksByName = map[string]rollup.Config{
	"goerli":  Goerli,
	"mainnet": Mainnet,
}

var L2ChainIDToNetworkName = func() map[string]string {
	out := make(map[string]string)
	for name, netCfg := range NetworksByName {
		out[netCfg.L2ChainID.String()] = name
	}
	return out
}()

func AvailableNetworks() []string {
	var networks []string
	for name := range NetworksByName {
		networks = append(networks, name)
	}
	return networks
}

func GetRollupConfig(name string) (rollup.Config, error) {
	network, ok := NetworksByName[name]
	if !ok {
		return rollup.Config{}, fmt.Errorf("invalid network %s", name)
	}

	return network, nil
}

func u64Ptr(v uint64) *uint64 {
	return &v
}
