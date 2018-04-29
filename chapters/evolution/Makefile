.DEFAULT_GOAL := build
.PHONY: retool run setup

PWD = $(shell pwd)

BUILD_DEPS = server/proto/thing.pb.go client/proto/thing_pb.rb client/proto/thing_services_pb.rb
PROTO_TARGETS = server/proto/%.pb.go client/proto/%_pb.rb client/proto/%_services_pb.rb

retool:
	@if test -z $(shell which retool); then \
		go get github.com/twitchtv/retool; \
		retool add github.com/golang/dep/cmd/dep v0.4.1; \
	fi

setup: retool
	$(info Synching dev tools and dependencies...)
	@retool sync
	@retool do dep ensure

$(PROTO_TARGETS): proto/%.proto
	$(info Generating protos...)
	@docker run --rm -v $(PWD):/defs namely/protoc-all:1.9 -d proto -l go -o server/proto
	@docker run --rm -v $(PWD):/defs namely/protoc-all:1.9 -f proto/thing.proto -l ruby -o client

build: $(BUILD_DEPS)
	@go build -o bin/server ./server

run: build
	@bin/server
