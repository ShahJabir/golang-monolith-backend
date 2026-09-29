.PHONY: build-api run-api dev-api migrate-build migrate-up-run migrate-up-dev migrate-down-run migrate-down-dev

API_BIN := bin/api
MIGRATE_BIN := bin/migrate

build-api:
	@go build -o $(API_BIN) ./cmd/api

run-api: build-api
	@./$(API_BIN)

dev-api:
	@go run ./cmd/api

migrate-build:
	@go build -o $(MIGRATE_BIN) ./cmd/migrate

migrate-up-run: migrate-build
	@./$(MIGRATE_BIN) up

migrate-up-dev:
	@go run ./cmd/migrate up

migrate-down-run: migrate-build
	@./$(MIGRATE_BIN) down

migrate-down-dev:
	@go run ./cmd/migrate down
