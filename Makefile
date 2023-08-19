.PHONY: all clean test coverage benchmark benchdata-gen benchdata-serve
all: coverage

# go list is the canonical utility to find go files
GOFILES := $(shell go list -f '{{ join .GoFiles "\n" }}' ./...)

test:
	go test -covermode=atomic -coverpkg=all ./...

benchmark:
	go test -bench=. ./...

benchdata-gen:
	go test -bench . -benchmem ./... | gobenchdata --append --json benchmarks.json

benchdata-serve: benchdata-gen
	gobenchdata web serve

coverage: .artifacts-stamp
	go-acc -o artifacts/c.out ./...
	go tool cover -html=artifacts/c.out -o artifacts/coverage.html

.artifacts-stamp:
	mkdir artifacts
	touch .artifacts-stamp

clean:
	rm -rf artifacts
	rm .artifacts-stamp
