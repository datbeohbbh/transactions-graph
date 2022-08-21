DIR=$(shell pwd)

LISTENER_DIR=$(DIR)/listener

DAO_DIR=$(DIR)/dao

gen-proto-listener: 
	cd $(LISTENER_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./listener.proto

gen-proto-dao: 
	cd $(DAO_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./dao.proto

format:
	go fmt ./...

note:
	echo "all address must be the string obtain from common.Address.Hex()"