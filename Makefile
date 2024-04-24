.PHONY: install proto

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get -u
	go mod tidy

proto:
	protoc -I=./proto --go_out=./crawler --go_opt=paths=source_relative ./proto/gtfs-realtime.proto
