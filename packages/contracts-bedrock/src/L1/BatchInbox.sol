// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;


import { Semver } from "../universal/Semver.sol";
import { CrossDomainMessenger } from "../universal/CrossDomainMessenger.sol";
/**
 * @custom:proxied
 * @title BatchInbox
 * @notice Calldata entries of version hashes which are checked against the precompile of blobs to verify they exist
 */
// slither-disable-next-line locked-ether
contract BatchInbox is Semver {
    uint32 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 100_000;
    address internal constant PODA_PRECOMPILE_ADDRESS = address(0x63);
    uint16 internal constant PODA_PRECOMPILE_COST = 1400;
    /**
     * @notice Messenger contract on this domain.
     */
    CrossDomainMessenger public immutable MESSENGER;
    /**
     * @custom:semver 1.0.0
     *
     * @param _messenger The address of the messenger on this domain.
     */
    constructor(
        address payable _messenger
    ) Semver(1, 0, 0) {
        MESSENGER = CrossDomainMessenger(_messenger);
    }

    /**
     * @notice appends an array of valid version hashes to the chain, each VH is checked via the VH precompile.
     *
     */
    function appendSequencerBatch(bytes32[] calldata _versionHashes) external view {
        require(_versionHashes.length > 0, "Must pass in atleast one version hash.");
        for (uint256 i = 0; i < _versionHashes.length; i++) {
            (bool success, bytes memory result) = PODA_PRECOMPILE_ADDRESS.staticcall{gas: PODA_PRECOMPILE_COST}(abi.encode(_versionHashes[i]));

            require(success, "Staticcall failed.");
            require(result.length > 0, "Return data must not be empty.");
        }
    }

    /**
     * @notice appends an array of valid version hashes to the chain and sends message to L2BatchInbox, each VH is checked via the VH precompile.
     *
     * @param _target L2 contract where message will be received with PoDA hashes.
     * @param _selector The function selector inside the L2 contract
     * @param _versionHashes Array of keccak256 version hashes in segments of 32 bytes each
     */
    function appendSequencerBatchToL2(address _target, bytes calldata _selector, bytes32[] calldata _versionHashes) external {
        this.appendSequencerBatch(_versionHashes);
        // Construct calldata
        bytes memory message = abi.encodeWithSelector(
            bytes4(_selector),
            _versionHashes
        );
        // Send calldata into L2
        MESSENGER.sendMessage(_target, message, RECEIVE_DEFAULT_GAS_LIMIT);
    }
}