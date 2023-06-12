import assert from 'assert'

import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { awaitCondition } from '@eth-optimism/core-utils'
import '@eth-optimism/hardhat-deploy-config'
import 'hardhat-deploy'

import {
  getDeploymentAddress,
  doOwnershipTransfer,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const proxyAdmin = await getDeploymentAddress(hre, 'ProxyAdmin')
  const l2OutputOracle = await getDeploymentAddress(hre, "L2OutputOracle")
  const l2OutputOracleProxy = await getDeploymentAddress(hre, 'L2OutputOracleProxy')

  await proxyAdmin.upgradeAndCall(
    l2OutputOracleProxy,
    l2OutputOracle,
    // TODO: how to abi encode here
  )

  // Upgrade and initialize the L2OutputOracle.
  // config.globalConfig.proxyAdmin.upgradeAndCall(
  //     payable(config.proxyAddressConfig.l2OutputOracleProxy),
  //     address(config.implementationAddressConfig.l2OutputOracleImpl),
  //     abi.encodeCall(
  //         L2OutputOracle.initialize,
  //         (
  //             l2OutputOracleDynamicConfig.l2OutputOracleStartingBlockNumber,
  //             l2OutputOracleDynamicConfig.l2OutputOracleStartingTimestamp
  //         )
  //     )
  // );

  await awaitCondition(
    async () => {
      return (
        (await proxyAdmin.callStatic.addressManager()) === addressManager
      )
    },
    30000,
    1000
  )
}

deployFn.tags = ['ProxyAdmin', 'L2OutputOracle', 'l1']

export default deployFn
