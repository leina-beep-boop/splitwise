# splitwise
 create a rudimentary splitwise application

## Generate protos
`protoc --go_out=protos/ledger --go-grpc_out=protos/ledger protos/ledger/ledger.proto`

## Run server
- `go run server.go`
- test server locally: `grpcurl -plaintext -d '{"expense":{"id":"1"}}' localhost:8080 ledger.LedgerService/AddExpense`
- alternatively, run client: `go run client.go`