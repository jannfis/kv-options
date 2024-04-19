all: build test

./test/out:
	mkdir -p ./test/out

.PHONY: build
build:
	CGO_ENABLED=0 go build -v ./...

.PHONY: test
test: ./test/out
	go test -coverprofile=./test/out/coverage.txt -race ./...

.PHONY: clean
clean:
	rm -rf ./test/out