PROTO_DIR := ../proto
GRPC_DIR := $(PROTO_DIR)/grpc
GRPC_FILES := $(notdir $(wildcard $(GRPC_DIR)/*.proto))

run:
	@swag init
	@go run main.go

grpc-gen:
	@protoc --proto_path=$(GRPC_DIR) --go_out=protoc-gen --go-grpc_out=protoc-gen $(GRPC_FILES)
