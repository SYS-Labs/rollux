#!/bin/sh
set -exu

VERBOSITY=${GETH_VERBOSITY:-3}
ALLOW_UNPROTECTED_TXS=${ALLOW_UNPROTECTED_TXS:-true}
GETH_DATA_DIR=/db
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GETH_KEYSTORE_DIR="$GETH_DATA_DIR/keystore"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
CHAIN_ID=$(cat "$GENESIS_FILE_PATH" | jq -r .config.chainId)
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"
SEQUENCER_RELAY_RPC="$SEQUENCER_RELAY_RPC"
MINING_ENABLED="$MINING_ENABLED"
if [ ! -d "$GETH_KEYSTORE_DIR" ]; then
	echo "$GETH_KEYSTORE_DIR missing, running account import"
	echo -n "pwd" > "$GETH_DATA_DIR"/password
	echo -n "$BLOCK_SIGNER_PRIVATE_KEY" | sed 's/0x//' > "$GETH_DATA_DIR"/block-signer-key
	geth account import \
		--datadir="$GETH_DATA_DIR" \
		--password="$GETH_DATA_DIR"/password \
		"$GETH_DATA_DIR"/block-signer-key
else
	echo "$GETH_KEYSTORE_DIR exists."
fi

if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
	echo "$GETH_CHAINDATA_DIR missing, running init"
	echo "Initializing genesis."
	geth --verbosity="$VERBOSITY" init \
		--datadir="$GETH_DATA_DIR" \
		"$GENESIS_FILE_PATH"
else
	echo "$GETH_CHAINDATA_DIR exists."
fi
L1_URL="http://u:p@l1:8370"
function wait_up {
  echo -n "Waiting for $1 to come up inside entrypoint ..."
  i=0
  until curl -s --data-binary '{"jsonrpc": "2.0", "id":"curltest", "method": "getblockcount", "params": [] }' -H 'content-type: application/json;' "$L1_URL" 2>&1 | grep -c '"error":null'
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
wait_up $L1_URL
# Warning: Archive mode is required, otherwise old trie nodes will be
# pruned within minutes of starting the tanenbaum.

exec geth \
	--datadir="$GETH_DATA_DIR" \
	--verbosity="$VERBOSITY" \
	--http \
	--http.corsdomain="*" \
	--http.vhosts="*" \
	--http.addr=0.0.0.0 \
	--http.port="$RPC_PORT" \
	--http.api=web3,debug,eth,txpool,net,engine \
	--ws \
	--ws.addr=0.0.0.0 \
	--ws.port="$WS_PORT" \
	--ws.origins="*" \
	--ws.api=debug,eth,txpool,net,engine \
	--syncmode=full \
	--nodiscover \
	--maxpeers=0 \
	--networkid="$CHAIN_ID" \
	--rpc.allow-unprotected-txs \
	--authrpc.addr="0.0.0.0" \
	--authrpc.port="9551" \
	--authrpc.vhosts="*" \
	--authrpc.jwtsecret=/config/jwt-secret.txt \
	--gcmode=archive \
	--metrics \
	--metrics.addr=0.0.0.0 \
	--metrics.port=6060 \
	--rpc.allow-unprotected-txs=$ALLOW_UNPROTECTED_TXS \
	--rollup.disabletxpoolgossip=true \
	"$@" >> "$GETH_DATA_DIR"/xout-geth.log
