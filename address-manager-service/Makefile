DIR=$(shell pwd)

ADDRESS_DIR=$(DIR)/address-manager

gen-address-manager:
	protoc \
	--proto_path=./protos \
	--go_out=./address-manager --go_opt=paths=source_relative \
	--go-grpc_out=./address-manager --go-grpc_opt=paths=source_relative \
	address.proto

format:
	go fmt ./...

note:
	echo "all address must be the string obtain from common.Address.Hex()"

gen-proto: gen-address-manager