DIR=$(shell pwd)

LISTENER_DIR=$(DIR)/listener

gen-proto: 
	cd $(LISTENER_DIR) && protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./listener.proto

format:
	go fmt ./...

log-in-mongo:
	cd ../project &&  \
	docker exec -it mongodb sh && \
	mongo --username nguyendat1211 --password password --host mongo --port 27017 --authenticationDatabase graphdb

note:
	echo "all address must be the string obtain from common.Address.Hex()"