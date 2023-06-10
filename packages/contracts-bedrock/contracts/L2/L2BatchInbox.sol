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
     * @notice appends an array of valid version hashes to the chain, each VH is checked via the VH precompile.
     *
     */
    function appendSequencerBatchFromL1(bytes32[] calldata _versionHashes) external onlyOtherBridge {
        require(_versionHashes.length > 0);
        for (uint256 i = 0; i < _versionHashes.length; i++) {
            podaMap[_versionHashes[i]] = true;
        }
    }
}
