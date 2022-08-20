DIR=$(shell pwd)

LISTENER_DIR=$(DIR)/listener

gen-proto: 
	cd $(LISTENER_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./listener.proto

format:
	go fmt ./...