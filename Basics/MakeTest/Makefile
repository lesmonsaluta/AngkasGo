.DEFAULT_GOAL := all


	
fmt:
	go fmt ./...
.PHONY:fmt
lint: fmt
	golint ./...
.PHONY:lint
vet: fmt
	go vet ./...
.PHONY:vet
build: vet
	go build hello.go
.PHONY:build
run:
	go run hello.go
all: fmt lint vet build run