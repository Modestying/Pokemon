GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
PROTO_FILES=$(shell find proto -name *.proto)

# env
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# proto
.PHONY: proto
proto:
	protoc --proto_path=./proto \
 	    --go_out=paths=source_relative:./proto \
 	    --go-grpc_out=paths=source_relative:./proto \
	    $(PROTO_FILES)

# build
.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...
