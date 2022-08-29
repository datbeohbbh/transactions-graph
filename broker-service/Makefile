DIR=$(shell pwd)

ADDR_DIR=$(DIR)/address-manager

gen-address-manager: 
	protoc \
	--proto_path=./protos \
	--go_out=./address-manager --go_opt=paths=source_relative \
	--go-grpc_out=./address-manager --go-grpc_opt=paths=source_relative \
	address.proto

gen-graph: 
	protoc \
	--proto_path=./protos \
	--go_out=./graph-client --go_opt=paths=source_relative \
	--go-grpc_out=./graph-client --go-grpc_opt=paths=source_relative \
	graph.proto

gen-proto: gen-address-manager gen-graph

format:
	go fmt ./...