.PHONY: lint setup test

PWD=$(shell pwd)

setup:
	$(info Synching dev tools and dependencies...)
	@if test -z $(which retool); then go get github.com/twitchtv/retool; fi
	@retool sync
	@retool do dep ensure

lint:
	@docker run --rm \
		-v $(PWD):$(PWD) \
		-w $(PWD)/proto \
		gwihlidal/protoc swapi.proto -I. --lint_out=.

server/api/mocks/provider.go: server/db/provider.go
	$(info Generating mock provider...)
	@retool do mockery -name Provider -dir server/db -output ./server/api/mocks

server/proto/swapi.pb.go: proto/swapi.proto
	$(info Generating go and gRPC protos...)
	@retool do protoc -Iproto --go_out=plugins=grpc:server/proto proto/swapi.proto

rpc-server: server/main.go server/api/*.go server/proto/swapi.pb.go
	$(info Building RPC server...)
	@go build -o rpc-server ./server

run-server: rpc-server
	./rpc-server

test: server/proto/swapi.pb.go server/api/mocks/provider.go
	@go test -v -cover ./server/api ./server/db
