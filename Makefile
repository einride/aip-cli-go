SHELL := /usr/bin/env bash

.PHONY: all
all: \
	buf-generate

cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

include tools/buf/rules.mk

.PHONY: clean
clean:
	$(info [$@] removing generated files...)
	@rm -rf build gen tools/*/*/


.PHONY: buildprotoc-gen-einridectl
protoc-gen-einridectl:
	go build -o build/$@ ./cmd/protoc-gen-einridectl

protoc_gen_einridectl := ./bin/protoc-gen-einridectl
export PATH := $(dir $(abspath $(protoc_gen_einridectl))):$(PATH)

.PHONY: $(protoc_gen_einridectl)
$(protoc_gen_einridectl):
	$(info [$@] building protoc-gen-einridectl...)
	@go build -o $@ ./cmd/protoc-gen-einridectl


.PHONY: buf-generate
buf-generate: $(buf) $(protoc_gen_einridectl)
	$(info [$@] generating protobuf...)
	@rm -rf gen
	@$(buf) generate image.bin
