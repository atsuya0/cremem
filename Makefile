.PHONY: build install format

build: format
	@go build

install: format
	@go install

format:
	@goimports -w cmd main.go
