export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export BIN=$(ROOT)/bin
export GOBIN?=$(BIN)
export GO=$(shell which go)
export BUILD=cd $(ROOT) && $(GO) install -v -ldflags "-s"
export DEBUG=cd $(ROOT) && $(GO) install -gcflags "all=-N -l"

build:
	$(BUILD) ./cmd/...
debug:
	$(DEBUG) ./cmd/...