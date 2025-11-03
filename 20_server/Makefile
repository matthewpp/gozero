# Makefile for Go REST API Server

.PHONY: run test mockgen migrate

run:
	go run main.go

test:
	go test ./...

mockgen:
	mockgen -source=internal/repository/user.go -destination=internal/repository/mock/user_repository_mock.go -package=mock
	mockgen -source=internal/service/user.go -destination=internal/service/mock/user_service_mock.go -package=mock

migrate:
	# Example: use golang-migrate or psql for migration
	psql "$$DATABASE_URL" -f migrations/001_create_users.sql

lint:
	golangci-lint run

clean:
	rm -rf internal/repository/mock/*.go
	rm -rf internal/service/mock/*.go
	rm -rf internal/mock/
