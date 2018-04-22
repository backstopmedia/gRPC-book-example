.DEFAULT_GOAL := build
.PHONY: retool run setup

PWD = $(shell pwd)

BUILD_DEPS = server/proto/archiver.pb.go server/proto/database.pb.go server/proto/tokenizer.pb.go \
						 archiver/proto/archiver_pb2.py archiver/proto/archiver_pb2_grpc.py \
						 database/proto/database_pb.rb database/proto/database_services_pb.rb \
						 tokenizer/proto/tokenizer_pb.js tokenizer/proto/tokenizer_grpc_pb.js

PROTO_TARGETS = server/proto/%.pb.go \
								archiver/proto/%_pb2.py archiver/proto/%_grpc.py \
								database/proto/%_pb.rb database/proto/%_services_pb.rb \
								tokenizer/proto/%.pb.js

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
	@docker run --rm -v $(PWD):/defs namely/protoc-all:1.9 -f proto/archiver.proto -l python -o archiver
	@docker run --rm -v $(PWD):/defs namely/protoc-all:1.9 -f proto/database.proto -l ruby -o database
	@docker run --rm -v $(PWD):/defs namely/protoc-all:1.9 -f proto/tokenizer.proto -l node -o tokenizer

build: $(BUILD_DEPS)
	@go build -o bin/server ./server

run: build
	@bin/server
