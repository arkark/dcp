.PHONY: build debug complete help clean

build:
	go build -tags=debug

debug: build
	docker-compose -f debug/docker-compose.yml up -d

help: build
	./dcp --help

clean:
	rm dcp
	docker-compose -f debug/docker-compose.yml down
