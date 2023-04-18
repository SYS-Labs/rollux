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
BATCH_INBOX_ADDRESS = os.getenv("CONTRACT_ADDRESS")
METRIC_NAME = os.getenv("METRIC_NAME")

w3 = Web3(Web3.HTTPProvider(L2_URL))
contract_address = w3.to_checksum_address(BATCH_INBOX_ADDRESS)

contract_health_metric = Gauge(METRIC_NAME, "Contract called in the past hour", ["called_in_past_hour"])


def check_contract_called_in_past_hour():
    current_block_number = w3.eth.block_number
    current_block = w3.eth.get_block(current_block_number)
    current_timestamp = current_block['timestamp']

    # Iterate back through blocks to find the one that is approximately two hours ago
    two_hours_ago_block_number = current_block_number
    one_hour_ago_timestamp = current_timestamp - 7200

    while True:
        logging.info(f"Entering while loop for checking all blocks in the past 2hrs")
        two_hours_ago_block_number -= 1
        block = w3.eth.get_block(two_hours_ago_block_number)
        if block['timestamp'] <= one_hour_ago_timestamp:
            logging.info(f"Exiting while loop that checks all blocks in the past 2hrs")
            break

    filter_params = {
      'fromBlock': two_hours_ago_block_number,
      'toBlock': current_block_number,
      'address': contract_address
    }
    logging.info(f"filter params: {filter_params}")
    contract_call_events = w3.eth.get_filter_logs(filter_params)
    logging.info(f"number of contract call events is {len(contract_call_events)}")
    return len(contract_call_events) > 0


def update_metrics():
    logging.info("update_metrics called")
    while True:
        contract_called = check_contract_called_in_past_hour()
        contract_health_metric.labels(called_in_past_hour=int(contract_called)).set(int(contract_called))
        time.sleep(3600)


if __name__ == "__main__":
    logging.info(f"Starting http server on port {PORT} ....")
    start_http_server(PORT)

    logging.info(f"Starting metrics update loop")
    metrics_thread = threading.Thread(target=update_metrics)
    metrics_thread.start()
