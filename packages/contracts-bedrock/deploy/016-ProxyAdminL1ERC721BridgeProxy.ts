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
  const proxy = await getDeploymentAddress(hre, 'L1ERC721BridgeProxy')

  await doOwnershipTransfer({
      false,
      proxy: proxy,
      name: 'L1ERC721BridgeProxy',
      transferFunc: 'changeAdmin',
      dictator: proxyAdmin,
    })

    await awaitCondition(
      async () => {
        return (
          (await proxy.callStatic.admin()) === proxyAdmin
        )
      },
      30000,
      1000
    )
}

deployFn.tags = ['ProxyAdmin', 'L1ERC721BridgeProxy', 'l1']

export default deployFn
