#!/bin/bash

sudo apt-get update

sudo apt-get install -y curl unzip

curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash -
sudo apt-get install -y nodejs

npm install -g bip39

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

unzip awscliv2.zip

sudo ./aws/install

rm -rf awscliv2.zip aws

aws --version

echo "Configuring the AWS CLI..."
aws configure
