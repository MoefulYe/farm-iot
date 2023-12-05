build:
	@go build -o target/http-server main.go

doc:
	@swag init


PROTO_DIR := ../proto/grpc
PROTO_FILE := kill.proto

grpc:
	@protoc --proto_path=$(PROTO_DIR) \
		--go_out=./protoc-gen \
		--go_grpc_out=./protoc-gen \
		$(PROTO_FILE)


