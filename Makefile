.PHONY: install proto

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

proto:
	protoc -I=./proto --go_out=./realtime --go_opt=paths=source_relative ./proto/gtfs-realtime.proto
