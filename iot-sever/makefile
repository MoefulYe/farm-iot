PROTO_DIR := ../proto
COW_DIR := $(PROTO_DIR)/cow
PROTO_COW_FILES := $(notdir $(wildcard $(COW_DIR)/*.proto))
GRPC_DIR := $(PROTO_DIR)/grpc
GRPC_FILES := $(notdir $(wildcard $(GRPC_DIR)/*.proto))

proto-cow:
	@protoc --proto_path=$(COW_DIR) --go_out=protoc-gen  $(PROTO_COW_FILES)

grpc-gen:
	@protoc --proto_path=$(GRPC_DIR) --go_out=protoc-gen --go-grpc_out=protoc-gen $(GRPC_FILES)

run:
	@go run main.go

build:
	@go build -o out/iot-server main.go
