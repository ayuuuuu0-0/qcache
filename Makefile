run: build
	@./bin/qcache

build:
	@go build -o bin/qcache .