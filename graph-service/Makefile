DIR=$(shell pwd)

GRAPH_DIR=$(DIR)/graph-data

gen-proto-graph: 
	cd $(GRAPH_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./graph.proto

format:
	go fmt ./...

note:
	echo "all address must be the string obtain from common.Address.Hex()"