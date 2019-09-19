SHELL := /bin/sh

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 0.0.1
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKGS_BY_PATH = $(shell go list ./... | grep -v /vendor/)

.PHONY: all build clean install uninstall fmt test coverage run help
.PHONY: tools lint doc tidy


all: version tools lint install

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help               Show this help screen.'
	@echo '    clean              Remove binaries, artifacts and releases.'
	@echo '    doc                Start Go documentation server on port 8080.'
	@echo '    tools              Install tools needed by the project.'
	@echo '    lint               Runs go golangci-lint run ./..
	@echo '    test               Run unit tests, and check.'
	@echo '    testverbose        Run unit tests in verbose mode, and check.'
	@echo '    coverage           Report code tests coverage, and check.'
	@echo '    build              Build project for current platform.'
	@echo '    tidy 							runs go mod tidy'
	@echo '    fmt                Run go fmt.'
	@echo '    install            Run go install'
	@echo '    uninstall          Force removes the artifact'
	@echo '    version 	          Checks the go version'
	@echo ''
	@echo 'Targets run by default are: checkall, install'
	@echo ''

build: $(TARGET)
	@true

clean: tidy
	@rm -f $(TARGET)

tidy:
	@go mod tidy

doc:
	godoc -http=:8080 -index

install:
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

# to be installed prior to running
tools:
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	@golangci-lint run ./...

test: lint
	@go test $(PKGS_BY_PATH)

testverbose: lint
	@go test -v $(PKGS_BY_PATH)

coverage: lint
	@go test -cover $(PKGS_BY_PATH)

version:
	@go version
