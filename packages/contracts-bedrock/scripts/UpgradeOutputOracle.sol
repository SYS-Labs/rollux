// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Script} from "forge-std/Script.sol";
import {ProxyAdmin} from "../contracts/universal/ProxyAdmin.sol";
import {L2OutputOracle} from "../contracts/L1/L2OutputOracle.sol";

contract UpgradeL2OutputOracle is Script {
    function run() external {
        address proxyAdminAddr = vm.envAddress("PROXY_ADMIN");
        address proxyAddr = vm.envAddress("L2_OUTPUT_ORACLE_PROXY");

        vm.startBroadcast();

        // Perform upgrade
        ProxyAdmin(proxyAdminAddr).upgrade(
            payable(proxyAddr),
            address(0x3F2F87e4F3584400f045c66D06c2377026D3dC84)
        );

        vm.stopBroadcast();
    }
}
