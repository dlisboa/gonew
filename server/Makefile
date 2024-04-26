# Change these
MAIN_PACKAGE_PATH := ./cmd/server
BINARY_NAME := server

## help: print usage
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: format code and tidy modfile
tidy:
	go fmt ./...
	go mod tidy -v

## build: build the app
build:
	go build -o=./bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## clean: clean build artifacts
clean:
	rm ./bin/${BINARY_NAME}

## run: run the app
run: build
	./bin/${BINARY_NAME}

## run/live: run the app with reloading
run/live:
	go run github.com/cespare/reflex@latest \
    -r '\.go|tmpl' -R '_test\.go' -s go run ${MAIN_PACKAGE_PATH}

## test: run all tests
test:
	go test -v -race -buildvcs ./...

## test: run all tests with reloading
test/live:
	go run github.com/cespare/reflex@latest \
    -r '\.go|tmpl' -s go test -v -race -buildvcs ./...

.PHONY: help tidy build clean run run/live test test/live
