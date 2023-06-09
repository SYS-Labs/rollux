import { DeployFunction } from 'hardhat-deploy/dist/types'
import { predeploys } from '../src'
import { ethers } from 'ethers'
import '@eth-optimism/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const Artifact__BatchInbox = await hre.companionNetworks[
    'l1'
  ].deployments.get('BatchInboxProxy')

  await deploy({
    hre,
    name: 'L2BatchInbox',
    args: [Artifact__BatchInbox.address],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'MESSENGER',
        ethers.utils.getAddress(predeploys.L2CrossDomainMessenger)
      )
      await assertContractVariable(
        contract,
        'OTHER_BRIDGE',
        ethers.utils.getAddress(Artifact__BatchInbox.address)
      )
    },
  })
}

deployFn.tags = ['L2BatchInboxImpl', 'l2']

export default deployFn
