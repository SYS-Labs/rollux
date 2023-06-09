// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { Semver } from "../universal/Semver.sol";
import { CrossDomainMessenger } from "../universal/CrossDomainMessenger.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000010
 * @title L2BatchInbox
 * @notice The L2BatchInbox receives L1 PoDA messages and stores them in podaMap, useful for L3's which need DA from L1
 */
contract L2BatchInbox is Semver {

    mapping(bytes32 => bool) public podaMap;

    /**
     * @notice Messenger contract on this domain.
     */
    CrossDomainMessenger public immutable MESSENGER;

    /**
     * @notice BatchInbox on L1 is the only one who can call this through CrossDomainMessenger
     */
    address public immutable OTHER_BRIDGE;
    /**
     * @custom:semver 1.0.0
     *
     * @param _otherBridge Address of the BatchInbox.
     */
    constructor(
        address payable _otherBridge
    ) Semver(1, 0, 0) {
        MESSENGER = CrossDomainMessenger(payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER));
        OTHER_BRIDGE = _otherBridge;
    }

    /**
     * @notice Ensures that the caller is a cross-chain message from the other bridge.
     */
    modifier onlyOtherBridge() {
        require(
            msg.sender == address(MESSENGER) &&
                MESSENGER.xDomainMessageSender() == OTHER_BRIDGE,
            "L2BatchInbox: function can only be called from the other bridge"
        );
        _;
    }

    /**
     * @notice appends an array of valid version hashes to the chain through calldata, each VH is checked via the VH precompile.
     * the calldata should be contingious set of 32 byte version hashes to check via precompile. Will consume memory for 1 hash and check that the a hash value was parrtoed back to indicate validity.
     *
     */
    function appendSequencerBatchFromL1(bytes calldata _extraData) external onlyOtherBridge {
        // Revert if the provided calldata does not consist of segments of 32 bytes.
        require((_extraData.length)%32 == 0);
        uint256 cursorPosition = 0;
        // Start loop. End once there is not sufficient remaining calldata to contain a 32 byte hash.
        while(cursorPosition <= (_extraData.length - 32)) {
            podaMap[bytes32(_extraData[cursorPosition])] = true;
            cursorPosition += 32;
        }
    }
}
