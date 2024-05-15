<div align="center">
  <br />
  <br />
  <a href="https://rollux.com"><img alt="Rollux" src="https://raw.githubusercontent.com/SYS-Labs/brand-kits/main/rollux/SVG/rollux_logo.svg" width=300></a>
  <br />
  <h3><a href="https://rollux.com">Rollux.com</a></h3>
  <br />
</div>


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [What is Rollux?](#what-is-optimism)
- [Documentation](#documentation)
- [Specification](#specification)
- [Community](#community)
- [Contributing](#contributing)
- [Security Policy and Vulnerability Reporting](#security-policy-and-vulnerability-reporting)
- [Directory Structure](#directory-structure)
- [Development and Release Process](#development-and-release-process)
  - [Overview](#overview)
  - [Production Releases](#production-releases)
  - [Development branch](#development-branch)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## What is Rollux?

[Rollux](https://rollux.com), built by [SYS Labs](https://syslabs.com), is a project dedicated to scaling blockchain technology and expanding its ability to coordinate people from across the world to build effective decentralized economies and governance systems. More specifically, Rollux is the EVM-equivalent Layer 2 optimistic rollup that introduces some key scaling technologies and security characteristics to Optimism's [OP Stack](https://stack.optimism.io). Built on [Syscoin's](https://syscoin.org) holistically modular Layer 1, Rollux inherits the security of Bitcoin’s mining network through merged-mining, combined with decentralized multi-quorum finality. Rollux fully factors Syscoin's efficient data availability (Proof of Data Availability - PoDA) and data fee market rather than relying on EVM calldata or danksharding. Rollux is secure, scalable, and offers very low fees. Additionally, Rollux provides native Layer 2 data availability which makes it ideal for supporting Layer 3 and fractal scaling.

In all, Rollux is a unique and powerful alternative to other rollup-based stacks and blockchain scaling technologies in general.

In this repository, you'll find numerous core components of Rollux, the decentralized software stack maintained by SYS Labs, and much of the upstream [OP Stack](https://stack.optimism.io) which is maintained by the Optimism Collective. We encourage you to explore, modify, extend, and test the code as needed. By collaborating on free, open software and shared standards, SYS Labs, Syscoin Foundation and the Optimism Collective aim to prevent siloed software development and rapidly accelerate the development of blockchain ecosystems. Come contribute, build the future, and redefine power, together.

**NOTE: It is important to understand that this repository became public relatively recently. As such, some READMEs might be incomplete. We appreciate your patience while we work quickly to expand technical information about Rollux and refactor existing content! Should you have questions about this repo, feel free to chat with the Rollux community at the links below!**

## Documentation

- If you want to build on top of Rollux Mainnet, refer to the [Rollux Developers Docs](https://rollux.com/developers)
- If you want to contribute to Rollux, check out the [Protocol Specs](./specs)
- If you want to build your own OP Stack based blockchain, refer to the [OP Stack docs](https://stack.optimism.io)
- If you want to build on top of OP Mainnet, refer to the [Optimism Documentation](https://docs.optimism.io)
- If you want to build your own OP Stack based blockchain, refer to the [OP Stack Guide](https://docs.optimism.io/stack/getting-started), and make sure to understand this repository's [Development and Release Process](#development-and-release-process)

## Specification

If you're interested in the technical details of how Optimism works, refer to the [Optimism Protocol Specification](https://github.com/ethereum-optimism/specs).

## Community

General dev discussion happens most frequently on the [Rollux discord](https://discord.gg/rollux) in the `#🔨│builder-general` channel.
<!--- Governance discussion can also be found on the [Optimism Governance Forum](https://gov.optimism.io/). --->

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of the contributing process for this repository.
Use the [Developer Quick Start](./CONTRIBUTING.md#development-quick-start) to get your development environment set up to start working on the Rollux Monorepo.
Then check out the list of [Good First Issues](https://github.com/sys-labs/rollux/contribute) to find something fun to work on!
Use the [Developer Quick Start](./CONTRIBUTING.md#development-quick-start) to get your development environment set up to start working on the Optimism Monorepo.
Then check out the list of [Good First Issues](https://github.com/ethereum-optimism/optimism/issues?q=is:open+is:issue+label:D-good-first-issue) to find something fun to work on!
Typo fixes are welcome; however, please create a single commit with all of the typo fixes & batch as many fixes together in a PR as possible. Spammy PRs will be closed.

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
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Rollux easier
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   └── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Rollux
├── <a href="./docs">docs</a>: A collection of documents including audits and post-mortems
├── <a href="./op-batcher">op-batcher</a>: L2-Batch Submitter, submits bundles of batches to L1
├── <a href="./op-bindings">op-bindings</a>: Go bindings for Bedrock smart contracts.
├── <a href="./op-bootnode">op-bootnode</a>: Standalone op-node discovery bootnode
├── <a href="./op-chain-ops">op-chain-ops</a>: State surgery utilities
├── <a href="./op-challenger">op-challenger</a>: Dispute game challenge agent
├── <a href="./op-e2e">op-e2e</a>: End-to-End testing of all bedrock components in Go
├── <a href="./op-heartbeat">op-heartbeat</a>: Heartbeat monitor service
├── <a href="./op-node">op-node</a>: rollup consensus-layer client
├── <a href="./op-preimage">op-preimage</a>: Go bindings for Preimage Oracle
├── <a href="./op-program">op-program</a>: Fault proof program
├── <a href="./op-proposer">op-proposer</a>: L2-Output Submitter, submits proposals to L1
├── <a href="./op-service">op-service</a>: Common codebase utilities
├── <a href="./op-ufm">op-ufm</a>: Simulations for monitoring end-to-end transaction latency
├── <a href="./op-wheel">op-wheel</a>: Database utilities
├── <a href="./ops">ops</a>: Various operational packages
├── <a href="./ops-bedrock">ops-bedrock</a>: Bedrock devnet work
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts-bedrock">contracts-bedrock</a>: Bedrock smart contracts
│   ├── <a href="./packages/contracts-ts">contracts-ts</a>: ABI and Address constants
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Optimism easier
│   ├── <a href="./packages/fee-estimation">fee-estimation</a>: Tools for estimating gas on OP chains
│   ├── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Optimism
│   └── <a href="./packages/web3js-plugin">web3js-plugin</a>: Adds functions to estimate L1 and L2 gas
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
├── <a href="./specs">specs</a>: Specs of the rollup starting at the Bedrock upgrade
└── <a href="./ufm-test-services">ufm-test-services</a>: Runs a set of tasks to generate metrics
</pre>

## Development and Release Process

### Overview

Please read this section if you're planning to fork this repository, or make frequent PRs into this repository.

### Production Releases

Production releases are always tags, versioned as `<component-name>/v<semver>`.
For example, an `op-node` release might be versioned as `op-node/v1.1.2`, and  smart contract releases might be versioned as `op-contracts/v1.0.0`.
Release candidates are versioned in the format `op-node/v1.1.2-rc.1`.
We always start with `rc.1` rather than `rc`.

For contract releases, refer to the GitHub release notes for a given release, which will list the specific contracts being released—not all contracts are considered production ready within a release, and many are under active development.

Tags of the form `v<semver>`, such as `v1.1.4`, indicate releases of all Go code only, and **DO NOT** include smart contracts.
This naming scheme is required by Golang.
In the above list, this means these `v<semver` releases contain all `op-*` components, and exclude all `contracts-*` components.

`op-geth` embeds upstream geth’s version inside it’s own version as follows: `vMAJOR.GETH_MAJOR GETH_MINOR GETH_PATCH.PATCH`.
Basically, geth’s version is our minor version.
For example if geth is at `v1.12.0`, the corresponding op-geth version would be `v1.101200.0`.
Note that we pad out to three characters for the geth minor version and two characters for the geth patch version.
Since we cannot left-pad with zeroes, the geth major version is not padded.

See the [Node Software Releases](https://docs.optimism.io/builders/node-operators/releases) page of the documentation for more information about releases for the latest node components.
The full set of components that have releases are:

- `chain-mon`
- `ci-builder`
- `ci-builder`
- `indexer`
- `op-batcher`
- `op-contracts`
- `op-challenger`
- `op-heartbeat`
- `op-node`
- `op-proposer`
- `op-ufm`
- `proxyd`
- `ufm-metamask`

All other components and packages should be considered development components only and do not have releases.

### Development branch

The primary development branch is [`develop`](https://github.com/sys-labs/rollux/tree/develop/).
`develop` contains the most up-to-date software that remains backwards compatible with the latest testnet [network deployments](https://rollux.com/developers/docs/useful-tools/networks/).
If you're making a backwards compatible change, please direct your pull request towards `develop`.

**Changes to contracts within `packages/contracts-bedrock/src` are usually NOT considered backwards compatible.**
Some exceptions to this rule exist for cases in which we absolutely must deploy some new contract after a tag has already been fully deployed.
If you're changing or adding a contract and you're unsure about which branch to make a PR into, default to using a feature branch.
Feature branches are typically used when there are conflicts between 2 projects touching the same code, to avoid conflicts from merging both into `develop`.

## License

All other files within this repository are licensed under the [MIT License](https://github.com/sys-labs/rollux/blob/master/LICENSE) unless stated otherwise.
