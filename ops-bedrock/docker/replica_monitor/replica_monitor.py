import os
import time
import requests
from prometheus_client import start_http_server, Gauge

# RPC URLs, you will need to define these as environment variables in your Docker container
L2_RPC_URL = os.getenv('L2_RPC_URL')
SEQUENCER_RPC_URL = os.getenv('SEQUENCER_RPC_URL')

# Prometheus metrics
hash_match_gauge = Gauge('hash_match', 'Block hash match between L2 and Sequencer', ['block_number'])


def get_latest_block_number(rpc_url):
    payload = {
      "jsonrpc": "2.0",
      "method": "eth_blockNumber",
      "params": [],
      "id": 1
    }
    headers = {'Content-Type': 'application/json'}

    response = requests.post(rpc_url, json=payload, headers=headers)
    if response.ok:
        return int(response.json()['result'], 16)
    return None


def get_block_hash(rpc_url, block_number):
    payload = {
      "jsonrpc": "2.0",
      "method": "eth_getBlockByNumber",
      "params": [hex(block_number), False],
      "id": 1
    }
    headers = {'Content-Type': 'application/json'}

    response = requests.post(rpc_url, json=payload, headers=headers)
    if response.ok:
        block_data = response.json()
        return block_data['result']['hash']
    return None

# Start the Prometheus client
start_http_server(8090)

# Main polling loop
while True:
    try:
        latest_sequencer_block = get_latest_block_number(SEQUENCER_RPC_URL)
        if latest_sequencer_block is None:
            print('Could not fetch latest block number from Sequencer')
            time.sleep(5)
            continue


        latest_l2_block = get_latest_block_number(L2_RPC_URL)

        retries = 3
        while retries > 0 and (latest_l2_block is None or latest_l2_block < latest_sequencer_block):
            time.sleep(10)  # wait for 10 seconds before retrying
            latest_l2_block = get_latest_block_number(L2_RPC_URL)
            retries -= 1

        if latest_l2_block is None or latest_l2_block < latest_sequencer_block:
            hash_match_gauge.labels(block_number=latest_sequencer_block).set(0)
            continue

        l2_hash = get_block_hash(L2_RPC_URL, latest_l2_block)
        sequencer_hash = get_block_hash(SEQUENCER_RPC_URL, latest_sequencer_block)
        hash_match = 1 if l2_hash == sequencer_hash else 0
        hash_match_gauge.labels(block_number=latest_sequencer_block).set(hash_match)

        if hash_match == 0:
            print(f'Hash mismatch at block {latest_sequencer_block}: L2({l2_hash}) vs Sequencer({sequencer_hash})')

    except Exception as e:
        print(f'An error occurred: {e}')

    # Sleep before the next polling cycle
    time.sleep(5)
