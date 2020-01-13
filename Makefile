.PHONY: build debug help clean

SHELL := /usr/bin/env bash

build:
	go build -tags=debug

debug: build
	docker-compose -f debug/docker-compose.yml up -d
	@echo -e "--> Execute \`\e[34msource completion/dcp\e[0m\`"

help: build
	./dcp --help

clean:
	rm dcp
	docker-compose -f debug/docker-compose.yml down
