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
  const chugSplashProxy = await getDeploymentAddress(hre, 'Proxy__OVM_L1StandardBridge')

  await doOwnershipTransfer({
      false,
      proxy: chugSplashProxy,
      name: 'L1ChugSplashProxy',
      transferFunc: 'setOwner',
      dictator: proxyAdmin,
    })

    await awaitCondition(
      async () => {
        return (
          (await chugSplashProxy.callStatic.owner()) === proxyAdmin
        )
      },
      30000,
      1000
    )
}

deployFn.tags = ['ProxyAdmin', 'L1StandardBridgeProxy', 'l1']

export default deployFn
