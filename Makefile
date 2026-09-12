.PHONY: build-api run-api dev-api

API_BIN := bin/api

build-api:
	@go build -o $(API_BIN) ./cmd/api

run-api: build-api
	@./$(API_BIN)

dev-api:
	@go run ./cmd/api
