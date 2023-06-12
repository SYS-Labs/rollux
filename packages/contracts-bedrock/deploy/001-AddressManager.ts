import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  doOwnershipTransfer,
  deploy
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await deploy({
    hre,
    name: 'Lib_AddressManager',
    contract: 'AddressManager',
    args: [],
    postDeployAction: async (contract) => {
      // Owner is temporarily set to the deployer.
      await assertContractVariable(contract, 'owner', deployer)
    },
  })

  const addressManager = await getDeploymentAddress(hre, 'AddressManager')

  // Transfer the address manager to the final system owner.
  if ((await addressManager.owner()) !== hre.deployConfig.finalSystemOwner) {
    await doOwnershipTransfer({
      false,
      proxy: addressManager,
      name: 'AddressManager',
      transferFunc: 'transferOwnership',
      dictator: hre.deployConfig.finalSystemOwner,
    })

    await awaitCondition(
      async () => {
        return (
          (await addressManager.callStatic.owner()) === hre.deployConfig.finalSystemOwner
        )
      },
      30000,
      1000
    )
  }
}

deployFn.tags = ['AddressManager', 'setup', 'l1']

export default deployFn
