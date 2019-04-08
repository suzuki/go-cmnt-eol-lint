NAME := go-cmnt-eol-lint

BIN_DIR := ./bin

all: vet test build

.PHONY: build
build:
	GO111MODULE=on go build -o $(BIN_DIR)/$(NAME) main.go

.PHONY: install
install:
	GO111MODULE=on go build -i -o $(GOPATH)/bin/$(NAME) main.go

.PHONY: test
test:
	GO111MODULE=on go test ./...

.PHONY: vet
vet:
	GO111MODULE=on go vet ./...

.PHONY: clean-testcache
clean-testcache:
	go clean -testcache

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)/*
