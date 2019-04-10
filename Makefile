NAME := go-cmnt-eol-lint

BIN_DIR := ./bin
PKG_DIR := ./pkg

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

.PHONY: pkg
pkg: pkg-linux-amd64 pkg-darwin-amd64

.PHONY: pkg-linux-amd64
pkg-linux-amd64:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o $(PKG_DIR)/$(NAME)_linux_amd64 main.go

.PHONY: pkg-darwin-amd64
pkg-darwin-amd64:
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o $(PKG_DIR)/$(NAME)_darwin_amd64 main.go
