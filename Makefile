.PHONY: build clean

.EXPORT_ALL_VARIABLES:
LOG_LEVEL = debug

all: build

build:
	go build -mod vendor -o bin/server ./cmd/server/main.go

run:
	go run -mod vendor ./cmd/server/main.go
