# How to Run a P2P Node

## Step 1: Add P2P Flag to Docker Compose File

First, you need to add the following flag to your `p2p-docker-compose.yml` as a flag of op-node service:

``
--p2p.static=/ip4/<ip>/tcp/9003/p2p/<peer_ids>
``


Replace <peer_ids> with a comma-separated list of peer IDs that you want to connect to. This flag specifies the
IP address, port, and peer ID of the nodes you want to connect to. You can reach out to the syscoin team for peer id
and ip address information


## Step 2: Generate a 32-byte hexadecimal string

To generate your peer ID for your node, you need a p2p-node-key which is a 32 byte hex.

You can generate this using openssl:

```
openssl rand -hex 32
```

## Step 3: Add node key

Next, you need to put your node key in a file called `p2p-node-key.txt` inside of `ops-bedrock`

## Step 4: Start node

Finally, you can run the `make p2p-tanenbaum-up` command to start your P2P node. This command will start the Docker container and connect to the nodes specified in your `--p2p-static` flag. You will find your peer id in the startup logs of op-node service.

## Participating in the P2P Network

If you're interested in participating in the P2P network, you can reach out to the Syscoin team



