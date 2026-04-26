.PHONY: setup build test lint

setup:
	git config core.hooksPath .githooks
	@echo "Git hooks configured from .githooks/"

build:
	go build -o bin/forg ./cmd/forg

test:
	go test ./...

lint:
	gofmt -l .
	go vet ./...
