PROTO_DIR := ../proto/cow
PROTO_FILES := $(notdir $(wildcard $(PROTO_DIR)/*.proto))

gen-proto: $(PROTO_FILES)
	@protoc --proto_path=$(PROTO_DIR) --go_out=./protoc-gen  $(PROTO_FILES)

GRPC_DIR := ../proto/grpc

grpc:
	@protoc --proto_path=$(GRPC_DIR) --go_out=./protoc-gen --go_grpc_out=./protoc-gen kill.proto

run:
	@go run main.go

build:
	@go build -o out/iot-server main.go
