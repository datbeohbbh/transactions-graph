DIR=$(shell pwd)

format-go:
	go fmt ./... 
	
format-proto:
	clang-format -i ./protos/*.proto

format: format-go

note:
	echo "all address must be the string obtain from common.Address.Hex()"