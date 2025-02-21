#!/usr/bin/env bash

# This script starts a local tanenbaum setup using Docker Compose. We have to use
# this more complicated Bash script rather than Compose's native orchestration
# tooling because we need to start each service in a specific order, and specify
# their configuration along the way. The order is:
#
# 1. Start L1 (outside this script).
# 2. Compile contracts (outside this script).
# 3. Deploy the contracts to L1 if necessary (outside this script).
# 4. Start L2, inserting the compiled contract artifacts into the genesis (outside this script).
# 5. Get the genesis hashes and timestamps from L1/L2 (outside this script).
# 6. Generate the rollup driver's config using the genesis hashes and the
#    timestamps recovered in step 4 as well as the address of the OptimismPortal
#    contract deployed in step 3 (outside this script).
# 7. Start the rollup driver.
# 8. Start the L2 output submitter.
#
# The timestamps are critically important here, since the rollup driver will fill in
# empty blocks if the tip of L1 lags behind the current timestamp. This can lead to
# a perceived infinite loop. To get around this, we set the timestamp to the current
# time in this script.
#
# This script is safe to run multiple times. It stores state in `.devnet`, and
# contracts-bedrock/deployments/devnet.
#
# Don't run this script directly. Run it using the makefile, e.g. `make tanenbaum-up`.
# To clean up your devnet, run `make tanenbaum-clean`.

set -eu

L1_URL="https://rpc.tanenbaum.io"
L2_URL="http://localhost:9545"
OP_NODE="$PWD/op-node"
CONTRACTS_BEDROCK="$PWD/packages/contracts-bedrock"
CONTRACTS_GOVERNANCE="$PWD/packages/contracts-governance"
NETWORK=tanenbaum
DEVNET="$PWD/.nebula"
TESTNET=1
TAG="nebula-v1.0.0"
# Helper method that waits for a given URL to be up. Can't use
# cURL's built-in retry logic because connection reset errors
# are ignored unless you're using a very recent version of cURL
function wait_up {
  echo -n "Waiting for $1 to come up..."
  i=0
  until curl -s -f -o /dev/null "$1"
  do
    echo -n .
    sleep 0.25

    ((i=i+1))
    if [ "$i" -eq 300 ]; then
      echo " Timeout!" >&2
      exit 1
    fi
  done
  echo "Done!"
}



# Bring up L2.
(
  export TAG=$TAG
  cd ops-bedrock
  echo "Bringing up L2..."
  docker-compose up -d l2
  wait_up $L2_URL
)

L2OO_ADDRESS="$(cat $DEVNET/rollup.json | jq -r '.output_oracle_address')"
# Bring up everything else.
(
  export TAG=$TAG
  cd ops-bedrock
  echo "Bringing up L2 services..."
  L2OO_ADDRESS="$L2OO_ADDRESS" \
      docker-compose up -d op-proposer op-batcher

  echo "Bringing up stateviz webserver..."
  docker-compose up -d stateviz
)

echo "L2 ready."

