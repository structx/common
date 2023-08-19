
deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run

rpc:
	protoc --go_out=pkg/pubsub --go_opt=paths=source_relative --go-grpc_out=pkg/pubcon --go-grpc_opt=paths=source_relative \
	protos/publisher_subscriber.proto
	protoc --go_out=pkg/pubcon --go_opt=paths=source_relative --go-grpc_out=pkg/pubcon --go-grpc_opt=paths=source_relative \
	protos/publisher_consumer.proto