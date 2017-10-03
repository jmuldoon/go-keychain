SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 0.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

TESTS = $(shell find . -type d -not -path "./.git*" -not -path "./.vscode*")

.PHONY: all build clean install uninstall fmt simplify check test coverage run help

all: check install

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help               Show this help screen.'
	@echo '    clean              Remove binaries, artifacts and releases.'
	@echo '    check              Runs go fmt, lint, vet'
	@echo '    test               Run unit tests.'
	@echo '    coverage           Report code tests coverage.'
	@echo '    build              Build project for current platform.'
	@echo '    fmt                Run go fmt.'
	@echo '    simplify           Run go fmt with -s.'
	@echo '    install            Run go install'
	@echo '    uninstall          Force removes the artifact'
	@echo '    run                Run executes install, then the binary is called'
	@echo ''
	@echo 'Targets run by default are: check, install'
	@echo ''

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

test:
	@go test -v $(TESTS)

coverage:
	@go test -cover $(TESTS)


run: install
	@$(TARGET)