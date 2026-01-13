BIN_DIR := bin
BIN := $(BIN_DIR)/swctl

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN) .

install:
	go install .

test:
	go test ./...
