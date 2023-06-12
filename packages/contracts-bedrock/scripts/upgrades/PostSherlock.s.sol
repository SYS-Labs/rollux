// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { console } from "forge-std/console.sol";
import { SafeBuilder } from "../universal/SafeBuilder.sol";
import { IMulticall3 } from "forge-std/interfaces/IMulticall3.sol";
import { IGnosisSafe, Enum } from "../interfaces/IGnosisSafe.sol";
import { LibSort } from "../libraries/LibSort.sol";
import { ProxyAdmin } from "../../contracts/universal/ProxyAdmin.sol";
import { Constants } from "../../contracts/libraries/Constants.sol";
import { SystemConfig } from "../../contracts/L1/SystemConfig.sol";
import { ResourceMetering } from "../../contracts/L1/ResourceMetering.sol";
import { Semver } from "../../contracts/universal/Semver.sol";

/**
 * @title PostSherlockL1
 * @notice Upgrade script for upgrading the L1 contracts after the sherlock audit.
 */
contract PostSherlockL1 is SafeBuilder {
    /**
     * @notice Address of the ProxyAdmin, passed in via constructor of `run`.
     */

    ProxyAdmin internal PROXY_ADMIN;

    /**
     * @notice Represents a set of L1 contracts. Used to represent a set of
     *         implementations and also a set of proxies.
     */
    struct ContractSet {
        address L1CrossDomainMessenger;
        address L1StandardBridge;
        address L2OutputOracle;
        address OptimismMintableERC20Factory;
        address OptimismPortal;
        address SystemConfig;
        address L1ERC721Bridge;
        // SYSCOIN
        address BatchInbox;
    }

    /**
     * @notice A mapping of chainid to a ContractSet of implementations.
     */
    mapping(uint256 => ContractSet) internal implementations;

    /**
     * @notice A mapping of chainid to ContractSet of proxy addresses.
     */
    mapping(uint256 => ContractSet) internal proxies;

    /**
     * @notice The expected versions for the contracts to be upgraded to.
     */
    string constant internal L1CrossDomainMessenger_Version = "1.4.0";
    string constant internal L1StandardBridge_Version = "1.1.0";
    string constant internal L2OutputOracle_Version = "1.3.0";
    string constant internal OptimismMintableERC20Factory_Version = "1.1.0";
    string constant internal OptimismPortal_Version = "1.6.0";
    string constant internal SystemConfig_Version = "1.3.0";
    string constant internal L1ERC721Bridge_Version = "1.1.1";
    string constant internal BatchInbox_Version = "1.0.0";

    /**
     * @notice Place the contract addresses in storage so they can be used when building calldata.
     */
    function setUp() external {
        implementations[GOERLI] = ContractSet({
            L1CrossDomainMessenger: 0x9b30CdC1aff7e7569E628834D00D2dd887F00174,
            L1StandardBridge: 0x75592Cb636e0fbE48F576C7b0A54e65C8945BA64,
            L2OutputOracle: 0x078A91d66fFc654C340093e472FEaC8156b98811,
            OptimismMintableERC20Factory: 0xbd0046FC69f969810267aC53f979b9325A6196f3,
            OptimismPortal: 0x6e8fd67c9E74918be4A6A983a8DD5aa82D775EDe,
            SystemConfig: 0x73703c5027FAA45fd66d592C61d22268B9730540,
            L1ERC721Bridge: 0xf7Fda8917c6B5589a514177F1878cc8ffE66f04a,
            BatchInbox: 0xf7Fda8917c6B5589a514177F1878cc8ffE66f04a
        });

        proxies[GOERLI] = ContractSet({
            L1CrossDomainMessenger: 0x46e963BE7CcF839b741f9DF0272d5241f22c2eA5,
            L1StandardBridge: 0xB806228Cd25620BBC55552632Bce419Aa403ba94,
            L2OutputOracle: 0x02dBDb985dC0fBa30De6715D1A34ee7179AC63Da,
            OptimismMintableERC20Factory: 0xc538309F438d52653A8f38290fB1da1e5f490395,
            OptimismPortal: 0xD251398404fD73E9f023dcfb66F913eecA4859F1,
            SystemConfig: 0x19CeD9B883cC0420F170DC0D1B270295699A5e8A,
            L1ERC721Bridge: 0x9365574Ee984442894a00aE25dFb72e68A567987,
            BatchInbox: 0x9365574Ee984442894a00aE25dFb72e68A567987
        });

        implementations[MAINNET] = ContractSet({
            L1CrossDomainMessenger: 0x007CD528731Ce3c429BF2AF871021ef0452c4C85,
            L1StandardBridge: 0xb3A84c1E60F12a60E9F1f0adb65A38391e2A248D,
            L2OutputOracle: 0x57A8B5eC2B4a0B77e4bf199Ba1978a42c69e1f6e,
            OptimismMintableERC20Factory: 0xaC3DD6A302fC9AFfFE6C78494923777D127f5579,
            OptimismPortal: 0x31c2490e4F3a6302Ff8d1614ED82c0da08907375,
            SystemConfig: 0xc45Ab831bc888601c2aEe68508c1D2a5388A9B97,
            L1ERC721Bridge: 0x4444b6c42404E259ed27c308d486ea1789210f46,
            BatchInbox: 0x7cdf9f61c28757e951B8DBB23c190Ac484116128
        });

        proxies[MAINNET] = ContractSet({
            L1CrossDomainMessenger: 0x0f81C3c3131Fa326a046d1C6103EBea5377843a0,
            L1StandardBridge: 0x67D271dA35cd98e591722475fba3BA28dbA3C58c,
            L2OutputOracle: 0xf1162A0b4ED6a980D63D59Ee4D8506Cd5BdD43fE,
            OptimismMintableERC20Factory: 0x98774d198baE88DbbAF3e4be29F51D6F572B86d4,
            OptimismPortal: 0xE4FBe4674610d4C201a037c032EB0b703d9d9377,
            SystemConfig: 0x489f4830f6a7c37D1F30E0491aaaB57DDe922f73,
            L1ERC721Bridge: 0x1F12C27a285EdcF22A341d3566E17BF7864A2735,
            BatchInbox: 0x32880C00EAD3979De1039A57A11Da00D8231FBcd
        });
    }

    /**
     * @notice Follow up assertions to ensure that the script ran to completion.
     */
    function _postCheck(ProxyAdmin _proxyAdmin) internal override view {
        ContractSet memory prox = getProxies();
        require(_versionHash(prox.L1CrossDomainMessenger) == keccak256(bytes(L1CrossDomainMessenger_Version)), "L1CrossDomainMessenger");
        require(_versionHash(prox.L1StandardBridge) == keccak256(bytes(L1StandardBridge_Version)), "L1StandardBridge");
        require(_versionHash(prox.L2OutputOracle) == keccak256(bytes(L2OutputOracle_Version)), "L2OutputOracle");
        require(_versionHash(prox.OptimismMintableERC20Factory) == keccak256(bytes(OptimismMintableERC20Factory_Version)), "OptimismMintableERC20Factory");
        require(_versionHash(prox.OptimismPortal) == keccak256(bytes(OptimismPortal_Version)), "OptimismPortal");
        require(_versionHash(prox.SystemConfig) == keccak256(bytes(SystemConfig_Version)), "SystemConfig");
        require(_versionHash(prox.L1ERC721Bridge) == keccak256(bytes(L1ERC721Bridge_Version)), "L1ERC721Bridge");
        require(_versionHash(prox.BatchInbox) == keccak256(bytes(BatchInbox_Version)), "BatchInbox");

        ResourceMetering.ResourceConfig memory rcfg = SystemConfig(prox.SystemConfig).resourceConfig();
        ResourceMetering.ResourceConfig memory dflt = Constants.DEFAULT_RESOURCE_CONFIG();
        require(keccak256(abi.encode(rcfg)) == keccak256(abi.encode(dflt)));

        // Check that the codehashes of all implementations match the proxies set implementations.
        ContractSet memory impl = getImplementations();
        require(_proxyAdmin.getProxyImplementation(prox.L1CrossDomainMessenger).codehash == impl.L1CrossDomainMessenger.codehash, "L1CrossDomainMessenger codehash");
        require(_proxyAdmin.getProxyImplementation(prox.L1StandardBridge).codehash == impl.L1StandardBridge.codehash, "L1StandardBridge codehash");
        require(_proxyAdmin.getProxyImplementation(prox.L2OutputOracle).codehash == impl.L2OutputOracle.codehash, "L2OutputOracle codehash");
        require(_proxyAdmin.getProxyImplementation(prox.OptimismMintableERC20Factory).codehash == impl.OptimismMintableERC20Factory.codehash, "OptimismMintableERC20Factory codehash");
        require(_proxyAdmin.getProxyImplementation(prox.OptimismPortal).codehash == impl.OptimismPortal.codehash, "OptimismPortal codehash");
        require(_proxyAdmin.getProxyImplementation(prox.SystemConfig).codehash == impl.SystemConfig.codehash, "SystemConfig codehash");
        require(_proxyAdmin.getProxyImplementation(prox.L1ERC721Bridge).codehash == impl.L1ERC721Bridge.codehash, "L1ERC721Bridge codehash");
        require(_proxyAdmin.getProxyImplementation(prox.BatchInbox).codehash == impl.BatchInbox.codehash, "BatchInbox codehash");
    }

    /**
     * @notice Test coverage of the logic. Should only run on goerli but other chains
     *         could be added.
     */
    function test_script_succeeds() skipWhenNotForking external {
        address _safe;
        address _proxyAdmin;

        if (block.chainid == MAINNET) {
            _safe = 0xA1307B87C87dbe4782C4C975e5Ba2326490DD720;
            _proxyAdmin = 0xE77924D4073642019EC2338f911ab1D16311A1B9;
        }

        require(_safe != address(0) && _proxyAdmin != address(0));

        address[] memory owners = IGnosisSafe(payable(_safe)).getOwners();

        for (uint256 i; i < owners.length; i++) {
            address owner = owners[i];
            vm.startBroadcast(owner);
            bool success = _run(_safe, _proxyAdmin);
            vm.stopBroadcast();

            if (success) {
                console.log("tx success");
                break;
            }
        }

        _postCheck(ProxyAdmin(_proxyAdmin));
    }

    /**
     * @notice Builds the calldata that the multisig needs to make for the upgrade to happen.
     *         A total of 8 calls are made, 7 upgrade implementations and 1 sets the resource
     *         config to the default value in the SystemConfig contract.
     */
    function buildCalldata(address _proxyAdmin) internal override view returns (bytes memory) {
        IMulticall3.Call3[] memory calls = new IMulticall3.Call3[](9);

        ContractSet memory impl = getImplementations();
        ContractSet memory prox = getProxies();

        // Upgrade the L1CrossDomainMessenger
        calls[0] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.L1CrossDomainMessenger), impl.L1CrossDomainMessenger)
            )
        });

        // Upgrade the L1StandardBridge
        calls[1] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.L1StandardBridge), impl.L1StandardBridge)
            )
        });

        // Upgrade the L2OutputOracle
        calls[2] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.L2OutputOracle), impl.L2OutputOracle)
            )
        });

        // Upgrade the OptimismMintableERC20Factory
        calls[3] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.OptimismMintableERC20Factory), impl.OptimismMintableERC20Factory)
            )
        });

        // Upgrade the OptimismPortal
        calls[4] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.OptimismPortal), impl.OptimismPortal)
            )
        });

        // Upgrade the SystemConfig
        calls[5] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.SystemConfig), impl.SystemConfig)
            )
        });

        // Upgrade the L1ERC721Bridge
        calls[6] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.L1ERC721Bridge), impl.L1ERC721Bridge)
            )
        });

        // Upgrade the BatchInbox
        calls[7] = IMulticall3.Call3({
            target: _proxyAdmin,
            allowFailure: false,
            callData: abi.encodeCall(
                ProxyAdmin.upgrade,
                (payable(prox.BatchInbox), impl.BatchInbox)
            )
        });

        // Set the default resource config
        ResourceMetering.ResourceConfig memory rcfg = Constants.DEFAULT_RESOURCE_CONFIG();
        calls[8] = IMulticall3.Call3({
            target: prox.SystemConfig,
            allowFailure: false,
            callData: abi.encodeCall(SystemConfig.setResourceConfig, (rcfg))
        });

        return abi.encodeCall(IMulticall3.aggregate3, (calls));
    }

    /**
     * @notice Returns the ContractSet that represents the implementations for a given network.
     */
    function getImplementations() internal view returns (ContractSet memory) {
        ContractSet memory set = implementations[block.chainid];
        require(set.L1CrossDomainMessenger != address(0), "no implementations for this network");
        return set;
    }

    /**
     * @notice Returns the ContractSet that represents the proxies for a given network.
     */
    function getProxies() internal view returns (ContractSet memory) {
        ContractSet memory set = proxies[block.chainid];
        require(set.L1CrossDomainMessenger != address(0), "no proxies for this network");
        return set;
    }
}
