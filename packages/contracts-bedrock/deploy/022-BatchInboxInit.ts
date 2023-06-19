import { DeployFunction } from 'hardhat-deploy/dist/types'
import '@eth-optimism/hardhat-deploy-config'
import 'hardhat-deploy'

import { getContractsFromArtifacts } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()
  const [proxyAdmin, batchInboxProxy, batchInboxImpl] =
    await getContractsFromArtifacts(hre, [
      {
        name: 'ProxyAdmin',
        signerOrProvider: deployer,
      },
      {
        name: 'batchInboxProxy',
        iface: 'BatchInbox',
        signerOrProvider: deployer,
      },
      {
        name: 'BatchInbox',
      },
    ])

  try {
    const tx = await proxyAdmin.upgrade(
      batchInboxProxy.address,
      batchInboxImpl.address
    )
    await tx.wait()
  } catch (e) {
    console.log('BatchInbox already initialized')
  }

  const version = await batchInboxProxy.callStatic.version()
  console.log(`BatchInbox version: ${version}`)

  console.log('Upgraded BatchInbox')
}

deployFn.tags = ['BatchInboxInitialize', 'l1']

export default deployFn
