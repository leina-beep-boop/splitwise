# splitwise
 create a primitive splitwise grpc application

## Generate protos
`protoc --go_out=protos/ledger --go-grpc_out=protos/ledger protos/ledger/ledger.proto`

`protoc --go_out=protos/splitter --go-grpc_out=protos/splitter protos/splitter/splitter.proto`

## Run server
- `go run server.go`
- test server locally: `grpcurl -plaintext -d '{"expense":{"id":"1"}}' localhost:8080 ledger.LedgerService/AddExpense`
- local server commands: `grpcurl -plaintext localhost:8080 describe`
- alternatively, run client: `go run client.go`