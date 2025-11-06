package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lmittmann/tint"
	_ "github.com/mattn/go-sqlite3"
)

func migratePostgres(ctx context.Context) {
	// Get PostgreSQL connection string from environment
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		slog.ErrorContext(ctx, "DB_URL environment variable not set for PostgreSQL migration")
		return
	}

	// Setup PostgreSQL connection
	postgresDb, err := sql.Open("postgres", dbURL)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create PostgreSQL connection", slog.String("error", err.Error()))
		return
	}
	defer func() {
		postgresDb.Close()
		slog.InfoContext(ctx, "PostgreSQL database connection closed")
	}()

	// Test connection
	if err := postgresDb.Ping(); err != nil {
		slog.ErrorContext(ctx, "Failed to ping PostgreSQL database", slog.String("error", err.Error()))
		return
	}

	// Run migrations
	driver, err := postgres.WithInstance(postgresDb, &postgres.Config{})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create PostgreSQL migration driver", slog.String("error", err.Error()))
		return
	}

	// Use relative path from project root
	migrationsPath := "file://migrations/postgresql"
	slog.InfoContext(ctx, "Using PostgreSQL migrations path", slog.String("path", migrationsPath))

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create PostgreSQL migrator", slog.String("error", err.Error()))
		return
	}

	migrationErr := m.Up()
	if migrationErr != nil && !errors.Is(migrationErr, migrate.ErrNoChange) {
		slog.ErrorContext(ctx, "Failed to run PostgreSQL migrations", slog.String("error", migrationErr.Error()))
		return
	}

	if errors.Is(migrationErr, migrate.ErrNoChange) {
		slog.InfoContext(ctx, "PostgreSQL migrations already up to date")
	} else {
		slog.InfoContext(ctx, "PostgreSQL migrations applied successfully")
	}
}

func migrateSQLite(ctx context.Context) {
	// Ensure database directory exists
	if err := os.MkdirAll("database", 0755); err != nil {
		slog.ErrorContext(ctx, "Failed to create database directory", slog.String("error", err.Error()))
		return
	}

	// Setup SQLite connection
	dbPath := "database/data.sqlite"
	sqliteDb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create SQLite connection", slog.String("error", err.Error()))
		return
	}
	defer func() {
		sqliteDb.Close()
		slog.InfoContext(ctx, "SQLite database connection closed")
	}()

	// Test connection
	if err := sqliteDb.Ping(); err != nil {
		slog.ErrorContext(ctx, "Failed to ping SQLite database", slog.String("error", err.Error()))
		return
	}

	// Run migrations
	driver, err := sqlite3.WithInstance(sqliteDb, &sqlite3.Config{})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create SQLite migration driver", slog.String("error", err.Error()))
		return
	}

	// Use relative path from project root
	migrationsPath := "file://migrations/sqlite"
	slog.InfoContext(ctx, "Using migrations path", slog.String("path", migrationsPath))

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"sqlite3", driver)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create SQLite migrator", slog.String("error", err.Error()))
		return
	}

	migrationErr := m.Up()
	if migrationErr != nil && !errors.Is(migrationErr, migrate.ErrNoChange) {
		slog.ErrorContext(ctx, "Failed to run SQLite migrations", slog.String("error", migrationErr.Error()))
		return
	}

	if errors.Is(migrationErr, migrate.ErrNoChange) {
		slog.InfoContext(ctx, "SQLite migrations already up to date")
	} else {
		slog.InfoContext(ctx, "SQLite migrations applied successfully")
	}
}

func initLogger() {
	h := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
	})
	slog.SetDefault(slog.New(h))
}

func main() {
	initLogger()

	var dbType string

	// Parse command line arguments for database type
	args := os.Args[1:]
	if len(args) == 0 {
		slog.Error("Database type not specified. Use [postgres|sqlite]")
		return
	}
	dbType = args[0]

	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// Don't fail if .env file doesn't exist, just log it
		slog.Info("No .env file found, continuing with system environment variables")
	} else {
		slog.Info("Successfully loaded .env file")
	}

	ctx := context.Background()

	// Determine database type if not specified
	if dbType == "" {
		slog.ErrorContext(ctx, "invalid database type")
		slog.InfoContext(ctx, "Supported types: postgres, sqlite")
		return
	}

	// Run migrations based on database type
	switch dbType {
	case "postgres", "postgresql":
		slog.InfoContext(ctx, "Starting PostgreSQL database migration")
		migratePostgres(ctx)
	case "sqlite":
		slog.InfoContext(ctx, "Starting SQLite database migration")
		migrateSQLite(ctx)
	default:
		slog.ErrorContext(ctx, "Unsupported database type", slog.String("type", dbType))
		slog.InfoContext(ctx, "Supported types: postgres, sqlite")
	}
}
