# Meta Info
NAME := goexample
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

## Install Dependencies
.PHONY: deps
deps:
	go get -v -d

## Setup
.PHONY: deps
devel-deps: deps
	GO111MODULE=off go get \
		golang.org/x/lint/golint \
		github.com/motemen/gobump/cmd/gobump \
		github.com/Songmu/make2help/cmd/make2help

## Run test
.PHONY: test
test: deps
	go test -v -cover ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binaries
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

## build binary
build: bin/myprof

## SHOW help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)