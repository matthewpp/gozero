# Go REST API Server with slog

A REST API server for User CRUD operations built with structured logging using Go's `slog` package.

## Your task

............................................

Create Plan CRUD feature

............................................

## Features

- **Layered Architecture**: API, Service, Repository pattern with interfaces
- **Structured Logging**: Context-aware logging with `slog` throughout all layers
- **Multiple Database Support**: PostgreSQL with pgx driver OR SQLite with go-sqlite3 driver
- **HTTP Framework**: Gin web framework
- **Testing**: Comprehensive unit tests with testify assertions and uber-go/mock generated mocks
- **Environment Configuration**: Automatic .env file loading with godotenv
- **Global Logger**: Centralized structured logging using slog
- **Database Pool**: Configurable connection pooling for PostgreSQL workloads

## Project Structure

```
20_server/
├── internal/
│   ├── api/                        # HTTP handlers with structured logging
│   │   ├── user.go                 # User API handlers
│   │   └── user_test.go            # API layer tests with testify assertions
│   ├── service/                    # Business logic layer with logging
│   │   ├── user.go                 # User service interface & implementation
│   │   ├── user_test.go            # Service layer tests
│   │   └── mock_services/          # Service mocks for API testing
│   │       └── user.go             # Generated service mock
│   ├── repository/                 # Data access layer with logging
│   │   ├── user.go                 # User repository interface
│   │   ├── user_postgres.go        # Postgresql user repository implementation
│   │   ├── user_postgres_test.go   # Postgresql Repository layer tests
│   │   ├── user_sqlite.go          # SQLite user repository implementation
│   │   ├── user_sqlite_test.go     # SQLite repository tests with testify
│   │   └── mock_repository/        # Repository mocks for service testing
│   │       └── user.go             # Generated repository mock
│   └── model/                      # Domain models
│       └── user.go                 # User model
├── migrations/                     # Database migrations
│   └── 001_create_users.sql
├── database/                       # Database Docker configuration
│   └── compose.yml                 # PostgreSQL container setup
├── main.go                         # Application entry point with .env loading
├── Makefile                        # Build and development tasks
├── .env.example                    # Environment variables template
├── .gitignore                      # Git ignore file
└── README.md                       # Project documentation
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

- Go 1.24+
- PostgreSQL OR SQLite
- Docker (optional, for PostgreSQL container)
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

3. Set up database and configure environment:

   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

   **Note**: The application automatically loads environment variables from the `.env` file if it exists using godotenv. If no `.env` file is found, it will fall back to system environment variables.

   **Database Options:**

   **Option A: PostgreSQL (Production)**

   ```bash
   # Start PostgreSQL with Docker
   make db-up
   make migrate-postgres
   ```

   **Option B: SQLite (Development/Testing)**

   ```bash
   # Run migrations, SQLite creates database file automatically
   make migrate-sqlite
   ```

4. Start the server:
   ```bash
   make run
   ```

### Development Commands

```bash
# Run the server
make run

# Run tests with testify assertions
make test

# Generate mocks using uber-go/mock
make mockgen

# Database commands (PostgreSQL)
make db-up         # Start PostgreSQL container
make db-down       # Stop PostgreSQL container
make db-restart    # Restart PostgreSQL container
make db-logs       # View PostgreSQL logs
make migrate       # Run database migrations

# Lint code
make lint

# Build project
make build

# Clean
make clean
```

### Testing

The project includes comprehensive unit tests with testify assertions and structured logging:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific test suites
go test ./internal/repository -v    # Repository tests (PostgreSQL + SQLite)
go test ./internal/api -v           # API tests with HTTP mocking
go test ./internal/service -v       # Service tests

# Run individual SQLite repository tests
go test ./internal/repository -v -run TestUserSQLiteRepository_Create
go test ./internal/repository -v -run TestUserSQLiteRepository_Update
```

**Test Features:**

- **Testify Assertions**: Professional assertions with clear error messages
- **Mock Services**: Complete isolation using uber-go/mock
- **HTTP Testing**: Full request/response validation
- **Database Testing**: Both PostgreSQL and SQLite repository testing
- **Error Scenarios**: Comprehensive error handling validation

### Mock Generation

This project uses [uber-go/mock](https://github.com/uber-go/mock) for generating mocks:

```bash
# Install mockgen tool (one time setup)
go install go.uber.org/mock/mockgen@latest

# Generate mocks
make mockgen
```

The mocks are organized by layer:

- `internal/repository/mock_repository/` - Repository mocks for service layer testing
- `internal/service/mock_services/` - Service mocks for API layer testing

### Environment Variables

**PostgreSQL Configuration:**

```env
# Database Connection
DB_URL=postgres://gozero_user:gozero_password@localhost:5432/gozero_db?sslmode=disable

# Enterprise Database Pool Configuration (optional)
DB_MAX_CONNS=25                    # Maximum connections in pool
DB_MIN_CONNS=5                     # Minimum connections to maintain
DB_MAX_CONN_IDLE_TIME=30m          # Max idle time before closing
DB_MAX_CONN_LIFETIME=1h            # Max connection lifetime
DB_CONNECT_TIMEOUT=10s             # Connection timeout

# Server Configuration
PORT=8080
GIN_MODE=release

# Logging Configuration
LOG_LEVEL=info
LOG_JSON=true
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

Import `gozero.postman_collection.json` to test the API endpoints with example requests.
