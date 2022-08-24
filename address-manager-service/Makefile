DIR=$(shell pwd)

ADDRESS_DIR=$(DIR)/address-manager

gen-proto-address: 
	cd $(ADDRESS_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./address.proto

format:
	go fmt ./...

note:
	echo "all address must be the string obtain from common.Address.Hex()"