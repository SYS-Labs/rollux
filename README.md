
<div align="center">
  <br />
  <br />
  <a href="https://rollux.com"><img alt="Rollux" src="https://raw.githubusercontent.com/SYS-Labs/brand-kits/main/rollux/SVG/rollux_logo.svg" width=300></a>
  <br />
  <h3><a href="https://rollux.com">Rollux.com</a></h3>
  <br />
</div>

## What is Rollux?

[Rollux](https://rollux.com), built by [SYS Labs](https://syslabs.com), is a project dedicated to scaling blockchain technology and expanding its ability to coordinate people from across the world to build effective decentralized economies and governance systems. More specifically, Rollux is the EVM-equivalent Layer 2 optimistic rollup that introduces some key scaling technologies and security characteristics to Optimism's [OP Stack](https://stack.optimism.io). Built on [Syscoin's](https://syscoin.org) holistically modular Layer 1, Rollux inherits the security of Bitcoin’s mining network through merged-mining, combined with decentralized multi-quorum finality. Rollux fully factors Syscoin's efficient data availability (Proof of Data Availability - PoDA) and data fee market rather than relying on EVM calldata or danksharding. Rollux is secure, scalable, and offers very low fees. Additionally, Rollux provides native Layer 2 data availability which makes it ideal for supporting Layer 3 and fractal scaling.

In all, Rollux is a unique and powerful alternative to other rollup-based stacks and blockchain scaling technologies in general.

In this repository, you'll find numerous core components of Rollux, the decentralized software stack maintained by SYS Labs, and much of the upstream [OP Stack](https://stack.optimism.io) which is maintained by the Optimism Collective. We encourage you to explore, modify, extend, and test the code as needed. By collaborating on free, open software and shared standards, SYS Labs, Syscoin Foundation and the Optimism Collective aim to prevent siloed software development and rapidly accelerate the development of blockchain ecosystems. Come contribute, build the future, and redefine power, together.

**NOTE: It is important to understand that this repository became public relatively recently. As such, some READMEs might be incomplete. We appreciate your patience while we work quickly to expand technical information about Rollux and refactor existing content! Should you have questions about this repo, feel free to chat with the Rollux community at the links below!**

## Documentation

- If you want to build on top of Rollux Mainnet, refer to the [Rollux Developers Docs](https://rollux.com/developers)
- If you want to contribute to Rollux, check out the [Protocol Specs](./specs)
- If you want to build your own OP Stack based blockchain, refer to the [OP Stack docs](https://stack.optimism.io)

## Community

General dev discussion happens most frequently on the [Rollux discord](https://discord.gg/rollux) in the `#🔨│builder-general` channel.
<!--- Governance discussion can also be found on the [Optimism Governance Forum](https://gov.optimism.io/). --->

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of the contributing process for this repository.
Use the [Developer Quick Start](./CONTRIBUTING.md#development-quick-start) to get your development environment set up to start working on the Rollux Monorepo.
Then check out the list of [Good First Issues](https://github.com/sys-labs/rollux/contribute) to find something fun to work on!

## Security Policy and Vulnerability Reporting

If you are reporting any vulnerabilites exclusive to the Rollux codebase, you should follow the common sense "How to" guidelines echoed in Optimism's canonical [Security Policy](https://github.com/ethereum-optimism/.github/blob/master/SECURITY.md).
Bounty hunters are encouraged to check out [the Optimism Immunefi bug bounty program](https://immunefi.com/bounty/optimism/).
While this does not apply to any Rollux-specific discoveries, the Optimism Immunefi program offers up to $2,000,042 for in-scope critical vulnerabilities in the Optimism codebase.
For vulnerabilities in any Rollux or SYS Labs websites, email servers or other non-critical infrastructure, please email SYS Labs at contact@syslabs.com. We appreciate detailed instructions for confirming the vulnerability.



## Bedrock-based

Rollux is based upon Optimism Bedrock!
You can find detailed specifications for the Bedrock upgrade within the [specs folder](./specs) in this repository.

## Directory Structure

<pre>
~~ Production ~~
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts-bedrock">contracts-bedrock</a>: Rollux Bedrock smart contracts.
│   ├── <a href="./packages/contracts-periphery">contracts-periphery</a>: Peripheral contracts for Rollux
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Optimism easier
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   ├── <a href="./packages/fault-detector">fault-detector</a>: Service for detecting Sequencer faults
│   ├── <a href="./packages/replica-healthcheck">replica-healthcheck</a>: Service for monitoring the health of a replica node
│   └── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Rollux
├── <a href="./op-bindings">op-bindings</a>: Go bindings for Bedrock smart contracts.
├── <a href="./op-batcher">op-batcher</a>: L2-Batch Submitter, submits bundles of batches to L1
├── <a href="./op-bootnode">op-bootnode</a>: Standalone op-node discovery bootnode
├── <a href="./op-chain-ops">op-chain-ops</a>: State surgery utilities
├── <a href="./op-challenger">op-challenger</a>: Dispute game challenge agent
├── <a href="./op-e2e">op-e2e</a>: End-to-End testing of all bedrock components in Go
├── <a href="./op-exporter">op-exporter</a>: Prometheus exporter client
├── <a href="./op-heartbeat">op-heartbeat</a>: Heartbeat monitor service
├── <a href="./op-node">op-node</a>: rollup consensus-layer client
├── <a href="./op-program">op-program</a>: Fault proof program
├── <a href="./op-proposer">op-proposer</a>: L2-Output Submitter, submits proposals to L1
├── <a href="./op-service">op-service</a>: Common codebase utilities
├── <a href="./op-signer">op-signer</a>: Client signer
├── <a href="./op-wheel">op-wheel</a>: Database utilities
├── <a href="./ops-bedrock">ops-bedrock</a>: Bedrock devnet work
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
└── <a href="./specs">specs</a>: Specs of the rollup starting at the Bedrock upgrade

~~ Pre-BEDROCK ~~
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts-periphery">contracts-periphery</a>: Peripheral contracts for Rollux
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Rollux easier
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   ├── <a href="./packages/fault-detector">fault-detector</a>: Service for detecting Sequencer faults
│   ├── <a href="./packages/replica-healthcheck">replica-healthcheck</a>: Service for monitoring the health of a replica node
│   └── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Optimism
├── <a href="./indexer">indexer</a>: indexes and syncs transactions
├── <a href="./op-exporter">op-exporter</a>: A prometheus exporter to collect/serve metrics from a Rollux node
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
└── <a href="./technical-documents">technical-documents</a>: audits and post-mortem documents
</pre>

## Branching Model

### Active Branches

| Branch          | Status                                                                           |
| --------------- | -------------------------------------------------------------------------------- |
| [master](https://github.com/sys-labs/rollux/tree/master/)                   | Accepts PRs from `develop` when intending to deploy to production.                  |
| [develop](https://github.com/sys-labs/rollux/tree/develop/)                 | Accepts PRs that are compatible with `master` OR from `release/X.X.X` branches.                    |
| release/X.X.X                                                                          | Accepts PRs for all changes, particularly those not backwards compatible with `develop` and `master`. |

### Overview

This repository generally follows [this Git branching model](https://nvie.com/posts/a-successful-git-branching-model/).
Please read the linked post if you're planning to make frequent PRs into this repository.

### Production branch

The production branch is `master`.
The `master` branch contains the code for latest "stable" releases.
Updates from `master` **always** come from the `develop` branch.

### Development branch

The primary development branch is [`develop`](https://github.com/sys-labs/rollux/tree/develop/).
`develop` contains the most up-to-date software that remains backwards compatible with the latest testnet [network deployments](https://rollux.com/developers/docs/useful-tools/networks/).
If you're making a backwards compatible change, please direct your pull request towards `develop`.

**Changes to contracts within `packages/contracts-bedrock/contracts` are usually NOT considered backwards compatible and SHOULD be made against a release candidate branch**.
Some exceptions to this rule exist for cases in which we absolutely must deploy some new contract after a release candidate branch has already been fully deployed.
If you're changing or adding a contract and you're unsure about which branch to make a PR into, default to using the latest release candidate branch.
See below for info about release candidate branches.

### Release candidate branches

Branches marked `release/X.X.X` are **release candidate branches**.
Changes that are not backwards compatible and all changes to contracts within `packages/contracts-bedrock/contracts` MUST be directed towards a release candidate branch.
Release candidates are merged into `develop` and then into `master` once they've been fully deployed.
We may sometimes have more than one active `release/X.X.X` branch if we're in the middle of a deployment.
See table in the **Active Branches** section above to find the right branch to target.

## Releases

### Changesets

We use [changesets](https://github.com/changesets/changesets) to mark packages for new releases.
When merging commits to the `develop` branch you MUST include a changeset file if your change would require that a new version of a package be released.

To add a changeset, run the command `yarn changeset` in the root of this monorepo.
You will be presented with a small prompt to select the packages to be released, the scope of the release (major, minor, or patch), and the reason for the release.
Comments within changeset files will be automatically included in the changelog of the package.

### Triggering Releases

Releases can be triggered using the following process:

1. Create a PR that merges the `develop` branch into the `master` branch.
2. Wait for the auto-generated `Version Packages` PR to be opened (may take several minutes).
3. Change the base branch of the auto-generated `Version Packages` PR from `master` to `develop` and merge into `develop`.
4. Create a second PR to merge the `develop` branch into the `master` branch.

After merging the second PR into the `master` branch, packages will be automatically released to their respective locations according to the set of changeset files in the `develop` branch at the start of the process.
Please carry this process out exactly as listed to avoid `develop` and `master` falling out of sync.

**NOTE**: PRs containing changeset files merged into `develop` during the release process can cause issues with changesets that can require manual intervention to fix.
It's strongly recommended to avoid merging PRs into develop during an active release.

## License

All other files within this repository are licensed under the [MIT License](https://github.com/sys-labs/rollux/blob/master/LICENSE) unless stated otherwise.
