[![Maintained by Spydra.app](https://img.shields.io/badge/maintained%20by-spydra.app-blueviolet)](https://www.spydra.app/?utm_source=github&utm_medium=fabric_contract)
[![Go Reference](https://pkg.go.dev/badge/github.com/spydra-tech/fabric-contract-go.svg)](https://pkg.go.dev/github.com/spydra-tech/fabric-contract-go)
# Spydra Hyperledger Fabric Base Contract

The Spydra Hyperledger Fabric Base Contract has various utility methods which provide additional functionalities out of the box to any Chaincode that is deployed in a [Spydra](https://www.spydra.app/?utm_source=github&utm_medium=fabric_contract) blockchain network. Simply extend your own Chaincode Smart Contract from the Spydra Base Contract to get these additional features.

- Query the world state using Graph QL. For more details, refer to [Spydra Graph QL](https://docs.spydra.app/products-overview/graphql).

## Quick Start
1. Run `go get`.
    ```sh
    $ go get -u github.com/spydra-tech/fabric-contract-go/spydracontract
    ```
2. Extend your Smart Contract from the Spydra Base Contract.
    ```go
    type MyCustomContract struct {
	    spydracontract.SpydraContract
    }
    ```
3. [Deploy](https://docs.spydra.app/how-to/apps/deploy-app) the Chaincode on a Spydra Blockchain network and start querying on any attribute using [Graph QL](https://docs.spydra.app/products-overview/graphql).