#!/bin/sh

mkdir -p /kvant
cd /kvant
ver=`wget -O - https://raw.githubusercontent.com/kvantdev-test/kvant-node/master/bin_name 2>/dev/null`
wget https://github.com/kvantdev-test/kvant-node/raw/master/release/$ver
chmod +x ./$ver
wget https://raw.githubusercontent.com/kvantdev-test/kvant-node/master/genesis/current/genesis.json
echo ./$ver node >> node_start.sh
chmod +x node_start.sh
./node_start.sh
