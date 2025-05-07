// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import { Deployer } from "./Deployer.sol";
import { L2OutputOracle } from "../contracts/L1/L2OutputOracle.sol";

contract Deploy is Script {
    // Modify these values as needed.
    uint256 internal constant SUBMISSION_INTERVAL = 1500;                 // e.g. ~ every 30 minutes
    uint256 internal constant L2_BLOCK_TIME = 2;                          // e.g. 2 seconds
    uint256 internal constant FINALIZATION_PERIOD_SECONDS = 3600;       // e.g. 10 days
    uint256 internal constant L2_OUTPUT_ORACLE_STARTING_BLOCK_NUMBER = 0;
    uint256  internal constant L2_OUTPUT_ORACLE_STARTING_TIMESTAMP = 1687365241;

    address internal constant L2_OUTPUT_ORACLE_PROPOSER   = 0x21821Dc24dd29D2D23fE7223dB2BB8D2a72cC95F;
    address internal constant L2_OUTPUT_ORACLE_CHALLENGER = 0xEF64E089f2e018C2d71213d6bB402EC14ce2c189;

    function run() external {
        // Start Foundry "broadcast" to perform on-chain deployments.
        vm.startBroadcast();

        /**
         * @dev Deploys the L2OutputOracle contract directly, passing the constructor
         *      parameters in the order expected by its constructor.
         */
        L2OutputOracle oracle = new L2OutputOracle(
            SUBMISSION_INTERVAL,
            L2_BLOCK_TIME,
            L2_OUTPUT_ORACLE_STARTING_BLOCK_NUMBER,
            L2_OUTPUT_ORACLE_STARTING_TIMESTAMP,
            L2_OUTPUT_ORACLE_PROPOSER,
            L2_OUTPUT_ORACLE_CHALLENGER,
            FINALIZATION_PERIOD_SECONDS
        );

        vm.stopBroadcast();

        // For debugging: log the address where the contract was deployed.
        console2.log("L2OutputOracle deployed at:", address(oracle));
    }
}
