// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import { L2OutputOracle } from "../contracts/L1/L2OutputOracle.sol";

contract Deploy is Script {
    // Modify these values as needed.
    uint256 internal constant SUBMISSION_INTERVAL = 1800;                 // e.g. ~ every 30 minutes
    uint256 internal constant L2_BLOCK_TIME = 2;                          // e.g. 2 seconds
    uint256 internal constant FINALIZATION_PERIOD_SECONDS = 864000;       // e.g. 10 days
    uint256 internal constant L2_OUTPUT_ORACLE_STARTING_BLOCK_NUMBER = 0;
    uint256  internal constant L2_OUTPUT_ORACLE_STARTING_TIMESTAMP = 0;

    address internal constant L2_OUTPUT_ORACLE_PROPOSER   = 0xF92BeC130533d9A791166902C29A6a7B6766BF2D;
    address internal constant L2_OUTPUT_ORACLE_CHALLENGER = 0x5925c588F6ef60052Df686Eea67284968a54762F;

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
