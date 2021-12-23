SHELL := /usr/bin/env bash

.PHONY: all
all: \
	clang-format \
	buf-lint \
	buf-generate \
	buf-generate-example

include tools/buf/rules.mk
include tools/clang-format/rules.mk

build/protoc-gen-go: ../go.mod
	$(info [$@] building binary...)
	@go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

build/protoc-gen-go-grpc: $(abspath $(lastword $(MAKEFILE_LIST)))
	$(info [$@] building binary...)
	@GOBIN=$(abspath $(dir $@)) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

.PHONY: build/protoc-gen-go-cli
build/protoc-gen-go-cli:
	$(info [$@] building binary...)
	@go build -o $@ ../

.PHONY: buf-lint
buf-lint: $(buf)
	$(info [$@] linting proto schemas...)
	@$(buf) lint

.PHONY: clang-format
clang-format: proto_files = $(shell find einride -type f -name '*.proto')
clang-format: $(clang_format)
	$(info [$@] formatting proto files...)
	@$(clang_format) \
		-i \
		--style='{BasedOnStyle: Google, ColumnLimit: 0, Language: Proto}' \
		$(proto_files)

.PHONY: buf-generate
buf-generate: $(buf) build/protoc-gen-go
	$(info [$@] generating protobuf stubs...)
	@rm -rf gen
	@$(buf) generate \
		--template buf.gen.yaml \
		--path einride

example_plugins := \
	build/protoc-gen-go \
	build/protoc-gen-go-grpc \
	build/protoc-gen-go-cli \

.PHONY: buf-generate-example
buf-generate-example: $(buf) $(example_plugins)
	$(info [$@] generating example CLI...)
	@rm -rf cmd
	@$(buf) generate \
		buf.build/einride/aip \
		--template buf.gen.example.yaml \
		--path einride/example/freight
