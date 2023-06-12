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

  // If we have the key for the controller then we don't need to wait for external txns.
  // Set the DISABLE_LIVE_DEPLOYER=true in the env to ensure the script will pause to simulate scenarios
  // where the controller is not the deployer.
  const isLiveDeployer = await liveDeployer({
    hre,
    disabled: process.env.DISABLE_LIVE_DEPLOYER,
  })


  // Transfer ownership of the ProxyAdmin to the final system owner.
  if ((await proxyAdmin.owner()) !== hre.deployConfig.finalSystemOwner) {
    await doOwnershipTransfer({
      isLiveDeployer,
      proxy: proxyAdmin,
      name: 'ProxyAdmin',
      transferFunc: 'transferOwnership',
      dictator: hre.deployConfig.finalSystemOwner,
    })

    await awaitCondition(
      async () => {
        return (
          (await proxyAdmin.callStatic.owner()) === hre.deployConfig.finalSystemOwner
        )
      },
      30000,
      1000
    )
  }
}

deployFn.tags = ['ProxyAdmin', 'transferOwnership', 'l1']

export default deployFn
