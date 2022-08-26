DIR=$(shell pwd)

ADDR_DIR=$(DIR)/address-manager

gen-proto-address: 
	cd $(ADDR_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./address.proto

GRAPH_DIR=$(DIR)/graph-client

gen-proto-graph: 
	cd $(GRAPH_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./graph.proto

format:
	go fmt ./...