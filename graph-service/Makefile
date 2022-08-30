DIR=$(shell pwd)

GRAPH_DIR=$(DIR)/graph

gen-graph:
	protoc \
	--proto_path=./protos \
	--go_out=./graph --go_opt=paths=source_relative \
	--go-grpc_out=./graph --go-grpc_opt=paths=source_relative \
	graph.proto

gen-graph-render-data:
	protoc \
	--proto_path=./protos \
	--go_out=./graphAlgo --go_opt=paths=source_relative \
	graph-render-data.proto

format-go:
	go fmt ./... 
	
format-proto:
	clang-format -i ./protos/*.proto

format: format-go format-proto

note:
	echo "all address must be the string obtain from common.Address.Hex()"

gen-proto: gen-graph gen-graph-render-data