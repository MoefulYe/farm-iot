PROTO_DIR := ../proto/cow
PROTO_FILES := $(notdir $(wildcard $(PROTO_DIR)/*.proto))

gen-proto:
	@protoc --proto_path=$(PROTO_DIR) --go_out=./protoc-gen  $(PROTO_FILES)

run:
	@go run main.go

build:
	@go build -o out/iot-server main.go
