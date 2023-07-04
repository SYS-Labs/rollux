import time
from web3 import Web3
from prometheus_client import start_http_server, Gauge
import threading
import os
import logging


logging.basicConfig(
  level=logging.INFO,
  format='%(asctime)s [%(levelname)s] %(message)s',
  handlers=[logging.StreamHandler()]
)
L2_URL = os.getenv("L1_URL")
PORT = 8000
CONTRACT_ADDRESS = os.getenv("CONTRACT_ADDRESS")
METRIC_NAME = os.getenv("METRIC_NAME")

w3 = Web3(Web3.HTTPProvider(L2_URL))
checksum_address = w3.to_checksum_address(CONTRACT_ADDRESS)

contract_health_metric = Gauge(METRIC_NAME, "Contract called in the past hour")


def check_contract_called_in_last_50_blocks():
    current_block_number = w3.eth.block_number

    # Determine the block number of the block that was 50 blocks ago
    fifty_blocks_ago_block_number = current_block_number - 50

    contract_call_events = get_contract_transactions(checksum_address, fifty_blocks_ago_block_number, current_block_number)
    logging.info(f"contract call events: {contract_call_events}")
    logging.info(f"number of contract call events is {len(contract_call_events)}")
    return len(contract_call_events) > 0


def update_metrics():
    logging.info("update_metrics called")
    while True:
        contract_called = check_contract_called_in_last_50_blocks()
        contract_health_metric.set(int(contract_called))
        time.sleep(3600)


def get_contract_transactions(contract_address, from_block, to_block):
    contract_transactions = []

    for block_number in range(from_block, to_block + 1):
        block = w3.eth.get_block(block_number, full_transactions=True)

        for transaction in block.get("transactions"):
            logging.info(f"Transaction {transaction}")
            tx = w3.eth.get_transaction(transaction.get("hash"))
            if tx.get('to') == contract_address:
                contract_transactions.append(transaction)

    return contract_transactions


if __name__ == "__main__":
    logging.info(f"Starting http server on port {PORT} ....")
    start_http_server(PORT)

    logging.info(f"Starting metrics update loop")
    metrics_thread = threading.Thread(target=update_metrics)
    metrics_thread.start()
