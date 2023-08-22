package chaincfg

import (
	"fmt"
	"strings"

	"github.com/ethereum-optimism/superchain-registry/superchain"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
)

//var Mainnet = rollup.Config{
//	Genesis: rollup.Genesis{
//		L1: eth.BlockID{
//			Hash:   common.HexToHash("0x438335a20d98863a4c0c97999eb2481921ccd28553eac6f913af7c12aec04108"),
//			Number: 17422590,
//		},
//		L2: eth.BlockID{
//			Hash:   common.HexToHash("0xdbf6a80fef073de06add9b0d14026d6e5a86c85f6d102c36d3d8e9cf89c2afd3"),
//			Number: 105235063,
//		},
//		L2Time: 1686068903,
//		SystemConfig: eth.SystemConfig{
//			BatcherAddr: common.HexToAddress("0x00D97b2A26Cb85252998fe7B4bd4eC2118bf6B6E"),
//			Overhead:    eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
//			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
//			GasLimit:    30_000_000,
//		},
//	},
//	BlockTime:              2,
//	MaxSequencerDrift:      1500,
//	SeqWindowSize:          288,
//	ChannelTimeout:         24,
//	L1ChainID:              big.NewInt(57),
//	L2ChainID:              big.NewInt(570),
//	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000420"),
//	DepositContractAddress: common.HexToAddress("0xD46Bf6354725bFd4409cd6A952695bFEb213aCB9"),
//	L1SystemConfigAddress:  common.HexToAddress("0x739d6e104C717566F65e4Ea711500CE81EF98D42"),
//	RegolithTime:           u64Ptr(0),
//}
//
//var Goerli = rollup.Config{
//	Genesis: rollup.Genesis{
//		L1: eth.BlockID{
//			Hash:   common.HexToHash("0xbaaa9a7834d9b5e928eeb36942b96eb64167701e16b9da02a7a5f3aa9c0a216c"),
//			Number: 247425,
//		},
//		L2: eth.BlockID{
//			Hash:   common.HexToHash("0x045514aee1f089c5acd01ee15995e39a406e92586495bfa4429aa93b9f6f1067"),
//			Number: 0,
//		},
//		L2Time: 1673550516,
//		SystemConfig: eth.SystemConfig{
//			BatcherAddr: common.HexToAddress("0x00d97b2a26cb85252998fe7b4bd4ec2118bf6b6e"),
//			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
//			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
//			GasLimit:    25_000_000,
//		},
//	},
//	BlockTime:              2,
//	MaxSequencerDrift:      1500,
//	SeqWindowSize:          288,
//	ChannelTimeout:         300,
//	L1ChainID:              big.NewInt(5700),
//	L2ChainID:              big.NewInt(57000),
//	BatchInboxAddress:      common.HexToAddress("0x678255ae6b5c4ba0e6206a8e70b59b874f20bc9c"),
//	DepositContractAddress: common.HexToAddress("0x61200b9fcbb421afd0bb5a732fe48ec98482e39c"),
//	L1SystemConfigAddress:  common.HexToAddress("0xd8daedc48ca71e20feb81cc3e51c9e3a89a3c84b"),
//	RegolithTime:           u64Ptr(1679079600),
//}
//
//var Sepolia = rollup.Config{
//	Genesis: rollup.Genesis{
//		L1: eth.BlockID{
//			Hash:   common.HexToHash("0x48f520cf4ddaf34c8336e6e490632ea3cf1e5e93b0b2bc6e917557e31845371b"),
//			Number: 4071408,
//		},
//		L2: eth.BlockID{
//			Hash:   common.HexToHash("0x102de6ffb001480cc9b8b548fd05c34cd4f46ae4aa91759393db90ea0409887d"),
//			Number: 0,
//		},
//		L2Time: 1691802540,
//		SystemConfig: eth.SystemConfig{
//			BatcherAddr: common.HexToAddress("0x8F23BB38F531600e5d8FDDaAEC41F13FaB46E98c"),
//			Overhead:    eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
//			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
//			GasLimit:    30000000,
//		},
//	},
//	BlockTime:              2,
//	MaxSequencerDrift:      600,
//	SeqWindowSize:          3600,
//	ChannelTimeout:         300,
//	L1ChainID:              big.NewInt(11155111),
//	L2ChainID:              big.NewInt(11155420),
//	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000011155420"),
//	DepositContractAddress: common.HexToAddress("0x16fc5058f25648194471939df75cf27a2fdc48bc"),
//	L1SystemConfigAddress:  common.HexToAddress("0x034edd2a225f7f429a63e0f1d2084b9e0a93b538"),
//	RegolithTime:           u64Ptr(0),
//}

var Mainnet, Goerli, Sepolia *rollup.Config

func init() {
	mustCfg := func(name string) *rollup.Config {
		cfg, err := GetRollupConfig(name)
		if err != nil {
			panic(fmt.Errorf("failed to load rollup config %q: %w", name, err))
		}
		return cfg
	}
	Mainnet = mustCfg("op-mainnet")
	Goerli = mustCfg("op-goerli")
	Sepolia = mustCfg("op-sepolia")
}

var L2ChainIDToNetworkDisplayName = func() map[string]string {
	out := make(map[string]string)
	for _, netCfg := range superchain.OPChains {
		out[fmt.Sprintf("%d", netCfg.ChainID)] = netCfg.Name
	}
	return out
}()

// AvailableNetworks returns the selection of network configurations that is available by default.
// Other configurations that are part of the superchain-registry can be used with the --beta.network flag.
func AvailableNetworks() []string {
	return []string{"op-mainnet", "op-goerli", "op-sepolia"}
}

// BetaAvailableNetworks returns all available network configurations in the superchain-registry.
// This set of configurations is experimental, and may change at any time.
func BetaAvailableNetworks() []string {
	var networks []string
	for _, cfg := range superchain.OPChains {
		networks = append(networks, cfg.Chain+"-"+cfg.Superchain)
	}
	return networks
}

func IsAvailableNetwork(name string, beta bool) bool {
	name = handleLegacyName(name)
	available := AvailableNetworks()
	if beta {
		available = BetaAvailableNetworks()
	}
	for _, v := range available {
		if v == name {
			return true
		}
	}
	return false
}

func handleLegacyName(name string) string {
	switch name {
	case "goerli":
		return "op-goerli"
	case "mainnet":
		return "op-mainnet"
	case "sepolia":
		return "op-sepolia"
	default:
		return name
	}
}

// ChainByName returns a chain, from known available configurations, by name.
// ChainByName returns nil when the chain name is unknown.
func ChainByName(name string) *superchain.ChainConfig {
	// Handle legacy name aliases
	name = handleLegacyName(name)
	for _, chainCfg := range superchain.OPChains {
		if strings.EqualFold(chainCfg.Chain+"-"+chainCfg.Superchain, name) {
			return chainCfg
		}
	}
	return nil
}

func GetRollupConfig(name string) (*rollup.Config, error) {
	chainCfg := ChainByName(name)
	if chainCfg == nil {
		return nil, fmt.Errorf("invalid network %s", name)
	}
	rollupCfg, err := rollup.LoadOPStackRollupConfig(chainCfg.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to load rollup config: %w", err)
	}
	return rollupCfg, nil
}
