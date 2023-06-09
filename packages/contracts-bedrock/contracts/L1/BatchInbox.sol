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
     * @notice appends an array of valid version hashes to the chain through calldata, each VH is checked via the VH precompile.
     * the calldata should be contingious set of 32 byte version hashes to check via precompile. Will consume memory for 1 hash and check that the a hash value was parrtoed back to indicate validity.
     *
     */
    function appendSequencerBatch(bytes calldata batchData) external view {
        // Revert if the provided calldata does not consist of segments of 32 bytes.
        require((batchData.length%32) == 0);
        uint256 cursorPosition = 0;
        // Start loop. End once there is not sufficient remaining calldata to contain a 32 byte hash.
        while(cursorPosition <= (batchData.length - 32)) {
            assembly{
                // Allocate memory for VH
                let memPtr := mload(0x40)
                // load 32 bytes from cursorPosition in calldata to memPtr location in memory
                calldatacopy(memPtr, cursorPosition, 0x20)
                // Set free pointer before function call.
                mstore(0x40, add(memPtr, 0x20))
                let result := staticcall(1400, 0x63, memPtr, 0x20, 0, 0)
                // check the RESULT does not indicate an error.
                switch result
                // Revert if precompile RESULT indicates an error.
                case 0 { revert(0, 0) }
                // Otherwise check the RETURNDATA
                default {
                    if eq(returndatasize(), 0) {
                        revert(0, 0)
                    }
                }
            }
            cursorPosition += 32;
        }
    }

    /**
     * @notice appends an array of valid version hashes to the chain through calldata and sends message to L2BatchInbox, each VH is checked via the VH precompile.
     * the calldata should be contingious set of 32 byte version hashes to check via precompile. Will consume memory for 1 hash and check that the a hash value was parrtoed back to indicate validity.
     *
     * @param _target L2 contract where message will be received with PoDA hashes.
     * @param _selector The function selector inside the L2 contract
     * @param batchData Array of keccak256 version hashes in segments of 32 bytes each
     */
    function appendSequencerBatchToL2(address _target, bytes memory _selector, bytes calldata batchData) external {
        this.appendSequencerBatch(batchData);
        // Construct calldata
        bytes memory message = abi.encodeWithSelector(
            bytes4(_selector),
            batchData
        );
        // Send calldata into L2
        MESSENGER.sendMessage(_target, message, RECEIVE_DEFAULT_GAS_LIMIT);
    }
}