.PHONY: lint setup

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
