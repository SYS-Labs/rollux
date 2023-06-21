# How to Run a P2P Node

## Step 1: Add P2P Flag to Docker Compose File

First, you need to add the following flag to your `p2p-docker-compose.yml` as a flag of op-node service:

``
--p2p.static=/ip4/<ip>/tcp/9003/p2p/<peer_ids>
``


Replace <peer_ids> with a comma-separated list of peer IDs that you want to connect to. This flag specifies the
IP address, port, and peer ID of the nodes you want to connect to. You can reach out to the syscoin team for peer id
and ip address information


## Step 2: Add sequencer rpc and disable miner
You'll need to set the sequencer rpc which will be provided by SYS team by adding it into `ops-bedrock/envs/p2p-node.env`
as the variable `SEQUENCER_RELAY_RPC` . This is important if you want your node to broadcast transactions to the network,
this part can be skipped if you will not be broadcasting transactions. Then also make sure mine flag `--mine` is set to
false in entrypoint.sh or take out flag completely to default to false


## Step 3: Generate a 32-byte hexadecimal string

To generate your peer ID for your node, you need a p2p-node-key which is a 32 byte hex.

You can generate this using openssl:

```
openssl rand -hex 32
```

## Step 4: Add node key

Next, you need to put your node key in a file called `p2p-node-key.txt` inside of `ops-bedrock`

## Step 5: Start node

Finally, you can run `make p2p-rollux-up` for mainnet and `make p2p-tanenbaum-up` for testnet to start your P2P node. This command will start the Docker
container and connect to the nodes specified in your `--p2p-static` flag. You will find your peer id in the startup
logs of op-node service. On first time run on a fresh machine, op-node will exit, this is expected because it requires
l1 to be synced up to the genesis block in .devnet/rollup.json, but ideally would be better to allow l1 sync up completely
before rerunning the command.

## Participating in the P2P Network

If you're interested in participating in the P2P network, you can reach out to the Syscoin team

