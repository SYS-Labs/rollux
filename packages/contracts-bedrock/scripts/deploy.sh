#!/usr/bin/env bash
set -euo pipefail


echo "> Deploying contracts"
forge script -vvv scripts/Deploy.s.sol:Deploy --rpc-url "$L1_RPC_URL" --broadcast --private-key "$PRIVATE_KEY" --verify --verifier blockscout --verifier-url https://tanenbaum.io/api

if [ -n "${DEPLOY_GENERATE_HARDHAT_ARTIFACTS:-}" ]; then
  echo "> Generating hardhat artifacts"
  forge script -vvv scripts/Deploy.s.sol:Deploy --sig 'sync()' --rpc-url "$L1_RPC_URL" --broadcast --private-key "$PRIVATE_KEY"
fi
