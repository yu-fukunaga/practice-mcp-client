.PHONY: env setup lint
GOPATH := $(shell go env GOPATH)
GOBIN := $(CURDIR)/bin
PATH := $(GOBIN):$(PATH)
SHELL := env PATH=$(PATH) bash

env:
	cp .env-sample .env

setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v2.0.1

lint:
	golangci-lint run
