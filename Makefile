SHELL := /bin/bash

# Executable's name
TARGET = go-manage-event

SRC = $(shell find . -type f -name '*.go')

ENV_LOCAL=\
		   DB_USERNAME=root\
		   DB_PASSWORD=services\
		   ENV=dev

ENV_PROD=\
		   DB_USERNAME=root\
		   DB_PASSWORD=services\
		   ENV=prod\
		   GIN_MODE=release

build: 	
	go build -o $(TARGET) cmd/main.go

clean:
	rm -f $(TARGET)

fmt:
	gofmt -l -w $(SRC)

simplify:
	gofmt -s -l -w $(SRC)

run: clean build
	${ENV_LOCAL} ./$(TARGET)

deploy: clean build
	${ENV_PROD} ./$(TARGET)