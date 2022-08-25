DIR=$(shell pwd)

ADDR_DIR=$(DIR)/address-manager

gen-proto-address: 
	cd $(ADDR_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./address.proto

format:
	go fmt ./...