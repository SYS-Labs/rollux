# Simple Setup
Please make sure to only run the latest release rather than directly from develop branch. You can do that with
`git fetch --tags` and `git checkout tags/<tag-name>`. Latest release as at today is `rollux-1.2.5` tag.
Testnet files are under `testnet/` directory, for testnet use files under testnet.
## Step 1: Set all environment variables in .env file

Most environment variables required in docker compose can be found in `.env` file, set and load these variables to
whatever environment you are deploying to, some of these have been set already and should not be changed.

## Step 3: Generate a 32-byte hexadecimal string

To generate your peer ID for your node, you need a p2p-node-key which is a 32 byte hex.

You can generate this using openssl:

```
openssl rand -hex 32
```

## Step 4: Add node key

Next, you need to put your node key into `p2p-node-key.txt`

## Step 5: Start node
Finally, just run the docker-compose file


