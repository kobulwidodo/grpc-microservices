.PHONY: run-app
run-app:
	@go run ./src/cmd/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    src/pb/**.proto