
# Code Structure
```
/services => business logic
/libs => contracts, abi gen
/mocks => mock test
/document => api doc gen
/cmd => entry points
/scripts => sql scripts
/api => well-structured api server
/inject => dependecy injection
/configs => chain-specific configuration
```

# ABI Contract Code Generate
```bash
# Install abigen
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
# Run abigen
./libs/contracts/abi_gen.sh
```

# Run
`go run main.go serve`

# Build
`go build -v ./...`

# Generate docs using swaggo
`swag init`
References: https://github.com/swaggo/swag

## CICD