# foldmarket

## generate protobuf code
`protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    market/market.proto`

## db
### queries
`sqlc generate`

### migrations

#### create
`migrate create -ext sql -dir migrations -seq create_users_table`

#### run
`migrate -database postgres://postgres:postgres@localhost:5432/marketdb -path migrations up`

