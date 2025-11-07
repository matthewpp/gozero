# Makefile for Go REST API Server
.PHONY: run build test mockgen migrate-sqlite migrate-postgres db-up db-down db-restart db-logs

run:
	@go run main.go

build:
	go build -o server main.go

test:
	go test ./...

lint:
	golangci-lint run

clean:
	go clean
	rm -f server
	rm -rf ./internal/repository/mock_repository
	rm -rf ./internal/service/mock_service
	
mockgen:
	go generate ./...

migrate-sqlite:
	go run ./migrations sqlite

migrate-postgres:
	go run ./migrations postgres


# Database Docker commands
db-up:
	docker compose -f database/compose.yml up -d

db-down:
	docker compose -f database/compose.yml down

db-restart:
	docker compose -f database/compose.yml restart gozero_db

db-logs:
	docker compose -f database/compose.yml logs -f gozero_db
