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
  const addressManager = await getDeploymentAddress(hre, 'AddressManager')

  await proxyAdmin.setAddressManager(addressManager)

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

deployFn.tags = ['ProxyAdmin', 'setAddressManager', 'l1']

export default deployFn
