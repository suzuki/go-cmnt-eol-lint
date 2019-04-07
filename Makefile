NAME := go-cmnt-eol-lint

BIN_DIR := ./bin
SRC_DIR := ./src

all: vet test build

.PHONY: build
build:
	GO111MODULE=on go build -o $(BIN_DIR)/$(NAME) $(SRC_DIR)/main.go

.PHONY: install
install:
	GO111MODULE=on go build -i -o $(GOPATH)/bin/$(NAME) $(SRC_DIR)/main.go

.PHONY: test
test:
	GO111MODULE=on go test $(SRC_DIR)/...

.PHONY: vet
vet:
	GO111MODULE=on go vet $(SRC_DIR)/...

.PHONY: clean-testcache
clean-testcache:
	go clean -testcache

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)/*
