DIR=$(shell pwd)

GRAPH_DIR=$(DIR)/graph-data

gen-graph:
	protoc \
	--proto_path=./protos \
	--go_out=./graph-data --go_opt=paths=source_relative \
	--go-grpc_out=./graph-data --go-grpc_opt=paths=source_relative \
	graph.proto

format:
	go fmt ./...

note:
	echo "all address must be the string obtain from common.Address.Hex()"

gen-proto: gen-graph