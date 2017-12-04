# Uhuchain-code

This reposirory includes the chaincode sources for the uhuchain ledger.

Prior to using this repository as source for a chaincode deployment, you need to run `make dependencies`. This uses [`dep`](https://github.com/golang/dep) to copy all required dependencies into the `vendor` subfolder.

Note: For some reasons there seem to be versioning issues if the `hyperledger fabric` dependencies are included. Therefor they are removed by the `make` target. Instead, they should be included directly from the regular `$GOPATH`.