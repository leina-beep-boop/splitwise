# splitwise
Create a primitive splitwise grpc application. 
Only supports:
- adding transactions
- getting expenses
- resetting all expenses

## Generate protos
`protoc --go_out=protos --go-grpc_out=protos protos/ledger.proto`

## Run server
- `go run server.go`
- test server locally: `grpcurl -plaintext -d '{"expense":{"id":"1"}}' localhost:8080 ledger.LedgerService/AddExpense`
- local server commands: `grpcurl -plaintext localhost:8080 describe`
- alternatively, run client: `go run client.go`

## Future plans
- split costs for two people
- split costs for >2 people
- implement proper ID maintenance for the expenses
- connect to a real database
- add unit testing
- handle input errors i.e. improper data formats
- add middleware logging
- create a more CLI friendly application
- build and deploy on the internet, preferably kubernetes