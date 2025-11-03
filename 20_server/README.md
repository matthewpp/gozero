# Go REST API Server with slog

A REST API server for User CRUD operations built with structured logging using Go's `slog` package.

## Features

- **Layered Architecture**: API, Service, Repository pattern with interfaces
- **Structured Logging**: Context-aware logging with `slog` throughout all layers
- **Database**: PostgreSQL with pgx driver
- **HTTP Framework**: Gin
- **Testing**: Unit tests with uber-go/mock generated mocks
- **Global Logger**: Centralized structured logging using slog

## Project Structure

```
20_server/
├── internal/
│   ├── api/                  # HTTP handlers with structured logging
│   │   ├── user.go          # User API handlers
│   │   └── user_test.go     # API layer tests
│   ├── service/             # Business logic layer with logging
│   │   ├── user.go          # User service interface & implementation
│   │   ├── user_test.go     # Service layer tests
│   │   └── mock/            # Service mocks for API testing
│   │       └── user_service_mock.go
│   ├── repository/          # Data access layer with logging
│   │   ├── user.go          # User repository interface & implementation
│   │   ├── user_test.go     # Repository layer tests
│   │   └── mock/            # Repository mocks for service testing
│   │       └── user_repository_mock.go
│   └── model/               # Domain models
│       └── user.go          # User model
├── migrations/              # Database migrations
│   └── 001_create_users.sql
├── main.go                  # Application entry point
├── Makefile                 # Build and development tasks
├── .env                     # Environment variables
└── README.md                # Project documentation
```

## Logging Features

### Structured JSON Logging

- JSON format for production environments
- Contextual information with each log entry
- Request tracing through all layers
- Global logger setup in main.go with `slog.SetDefault()`

### Log Levels and Context

- **Info**: Normal operations (user creation, retrieval, etc.)
- **Error**: Error conditions with detailed context
- **Context Propagation**: Request context flows through all layers
- **Global Access**: All layers use `slog.InfoContext()` and `slog.ErrorContext()` directly

### Example Log Output

```json
{"time":"2025-11-03T10:30:45.123Z","level":"INFO","msg":"API: Creating user request received","method":"POST","path":"/users"}
{"time":"2025-11-03T10:30:45.124Z","level":"INFO","msg":"Service: Creating user","name":"John Doe","email":"john@example.com"}
{"time":"2025-11-03T10:30:45.125Z","level":"INFO","msg":"Creating user","name":"John Doe","email":"john@example.com"}
{"time":"2025-11-03T10:30:45.130Z","level":"INFO","msg":"User created successfully","id":1,"name":"John Doe","email":"john@example.com"}
```

## API Endpoints

- `POST /users` - Create a new user
- `GET /users/:id` - Get user by ID
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user
- `GET /users` - List all users

## Setup and Usage

### Prerequisites

- Go 1.21+
- PostgreSQL
- Make

### Installation

1. Clone and navigate to the project:

   ```bash
   cd 20_server
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up PostgreSQL and configure environment:

   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. Run migrations:

   ```bash
   make migrate
   ```

5. Start the server:
   ```bash
   make run
   ```

### Development Commands

```bash
# Run the server
make run

# Run tests with logging output
make test

# Generate mocks using uber-go/mock
make mockgen

# Run database migrations
make migrate

# Lint code
make lint

# Clean generated files
make clean
```

### Testing

The project includes unit tests with structured logging in test environments:

```bash
# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...
```

### Mock Generation

This project uses [uber-go/mock](https://github.com/uber-go/mock) for generating mocks:

```bash
# Install mockgen tool (one time setup)
go install go.uber.org/mock/mockgen@latest

# Generate mocks
make mockgen
```

The mocks are organized by layer:

- `internal/repository/mock/` - Repository mocks for service layer testing
- `internal/service/mock/` - Service mocks for API layer testing

### Environment Variables

```env
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

## Logging Configuration

The logger is configured in `main.go` with:

- JSON handler for structured output
- Info level logging
- Global logger setup using `slog.SetDefault()`
- Context propagation through HTTP middleware

Each layer uses the global logger directly:

- **API Layer**: HTTP request/response details using `slog.InfoContext()`
- **Service Layer**: Business operation context using `slog.InfoContext()`
- **Repository Layer**: Database operation details using `slog.InfoContext()`
- **No Dependency Injection**: All layers access the same global logger instance

## Postman Collection

Import `user_api.postman_collection.json` to test the API endpoints with example requests.
