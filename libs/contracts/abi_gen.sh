
#!/usr/bin/env bash

# set -euo pipefail

# readonly src_dir=/tmp/kyber-api/src/github.com/ethereum
# export GOPATH=/tmp/kyber-api
# export PATH=$GOPATH/bin:$PATH

# mkdir -p "$src_dir"
# cd "$src_dir"
# [[ -d go-ethereum ]] || git clone https://github.com/ethereum/go-ethereum.git
# go install -v github.com/ethereum/go-ethereum/cmd/abigen

# Guarantee out files will be in the same folder as this script
SCRIPT_PATH=${0%/*}
if [ "$0" != "$SCRIPT_PATH" ] && [ "$SCRIPT_PATH" != "" ]; then
    cd $SCRIPT_PATH
fi

abigen -abi ./metronion_accessories.abi -pkg contracts -type MetronionAccessories -out ./metronion_accessories_abi.go
abigen -abi ./metronion_nft.abi -pkg contracts -type MetronionNFT -out ./metronion_nft_abi.go
abigen -abi ./metronion_sale.abi -pkg contracts -type MetronionSale -out ./metronion_sale_abi.go
abigen -abi ./metrogalaxy_marketplace.abi -pkg contracts -type MetroGalaxyMarketplace -out ./metrogalaxy_marketplace.go
