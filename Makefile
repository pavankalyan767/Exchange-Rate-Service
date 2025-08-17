MAIN_FILE := main.go
SERVICE_NAME := rate-exchange-service

.PHONY: build
build:
	mkdir -p ./bin
	go build -tags=viper_bind_struct -o=./bin/$(SERVICE_NAME) ./$(MAIN_FILE)

.PHONY: run
run: build
	./bin/$(SERVICE_NAME)

.PHONY: serve
serve:
	BINARY_NAME=$(SERVICE_NAME) air -c .air.toml

.PHONY: clean
clean:
	rm -rf ./bin ./tmp

.PHONY: rebuild
rebuild: clean build

