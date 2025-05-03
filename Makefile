run: build
	@./bin/qcache --listenAddr :5001

build:
	@go build -o bin/qcache .