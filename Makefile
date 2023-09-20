
deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run

rpc:
	protoc --go_out=pkg/message_broker --go_opt=paths=source_relative --go-grpc_out=pkg/message_broker --go-grpc_opt=paths=source_relative \
	protos/message_broker.proto
	protoc --go_out=pkg/registry --go_opt=paths=source_relative --go-grpc_out=pkg/registry --go-grpc_opt=paths=source_relative \
	protos/registry_service.proto
