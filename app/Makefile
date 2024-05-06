# Change these
MAIN_PACKAGE_PATH := ./cmd/server
BINARY_NAME := server

## help: print usage
help:
	@echo 'Usage: ' && sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: format code and tidy modfile
tidy:
	go fmt ./...

## build: build the app
build:
	go mod tidy
	go build -o=./bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## clean: clean build artifacts
clean:
	rm ./bin/${BINARY_NAME}

## run: run the app
run:
	godotenv go run ${MAIN_PACKAGE_PATH}

## run/live: run the app with reloading
run/live:
	reflex -r '\.go|tmpl' -R '_test\.go' -s godotenv go run ${MAIN_PACKAGE_PATH}

## test: run all tests
test:
	go test -v -race ./...

## test: run all tests with reloading
test/live:
	reflex -r '\.go|tmpl' -s go test -v -race -buildvcs ./...

## lint: run linters
lint:
	golangci-lint run ./...

## lint/fix: run linters and fix issues
lint/fix:
	golangci-lint run --fix ./...

## deps: install dev and prod deps
deps:
	go mod tidy
	go install github.com/cespare/reflex@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/amacneil/dbmate/v2@v2.14.0
	go install github.com/joho/godotenv/cmd/godotenv@latest

.PHONY: help tidy build clean run run/live test test/live lint lint/fix deps
