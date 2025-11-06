package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gozero/server/internal/api"
	"gozero/server/internal/middleware"
	"gozero/server/internal/repository"
	"gozero/server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

// getEnvAsInt retrieves an environment variable as an integer with a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsDuration retrieves an environment variable as a duration with a default value
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// createDatabasePool creates and configures a pgx connection pool for enterprise use
func createDatabasePool(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		slog.ErrorContext(ctx, "DB_URL environment variable not set")
		return nil, errors.New("DB_URL environment variable not set")
	}

	// Parse the database URL to get the base configuration
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// Enterprise-grade connection pool configuration
	// Maximum number of connections in the pool
	config.MaxConns = int32(getEnvAsInt("DB_MAX_CONNS", 25))

	// Minimum number of connections to maintain in the pool
	config.MinConns = int32(getEnvAsInt("DB_MIN_CONNS", 5))

	// Maximum time a connection can remain idle before being closed
	config.MaxConnIdleTime = getEnvAsDuration("DB_MAX_CONN_IDLE_TIME", 30*time.Minute)

	// Maximum lifetime of a connection before it's closed and recreated
	config.MaxConnLifetime = getEnvAsDuration("DB_MAX_CONN_LIFETIME", 1*time.Hour)

	// Connection timeout
	config.ConnConfig.ConnectTimeout = getEnvAsDuration("DB_CONNECT_TIMEOUT", 10*time.Second)

	slog.InfoContext(ctx, "Database pool configuration",
		"max_conns", config.MaxConns,
		"min_conns", config.MinConns,
		"max_conn_idle_time", config.MaxConnIdleTime,
		"max_conn_lifetime", config.MaxConnLifetime,
		"connect_timeout", config.ConnConfig.ConnectTimeout,
	)

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func createSqliteConnection(ctx context.Context, dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func initLogger() {
	// Initialize structured logger
	var h slog.Handler
	if os.Getenv("LOG_JSON") == "true" {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		h = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.TimeOnly,
		})
	}
	logger := slog.New(h)
	slog.SetDefault(logger)
}

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, continuing with system environment variables")
	} else {
		slog.Info("Successfully loaded .env file")
	}

	initLogger()

	ctx := context.Background()

	// // Setup PostgreSQL connection pool
	// slog.InfoContext(ctx, "Starting database connection pool setup")
	// pgxPool, err := createDatabasePool(ctx)
	// if err != nil {
	// 	slog.ErrorContext(ctx, "Failed to create database pool", slog.String("error", err.Error()))
	// 	return
	// }
	// defer func() {
	// 	pgxPool.Close()
	// 	slog.InfoContext(ctx, "database connection pool closed")
	// }()
	// slog.InfoContext(ctx, "Successfully connected to database with connection pool")

	// Setup SQLite connection
	dbPath := "database/data.sqlite"
	sqliteDb, err := createSqliteConnection(ctx, dbPath)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create SQLite connection", slog.String("error", err.Error()))
		return
	}
	defer func() {
		sqliteDb.Close()
		slog.InfoContext(ctx, "SQLite database connection closed")
	}()

	// Initialize User feature : repositories, services, and handlers
	// userPostgresRepo := repository.NewUserPostgresRepository(pgxPool)
	userSqliteRepo := repository.NewUserSQLiteRepository(sqliteDb)
	userService := service.NewUserService(userSqliteRepo)
	userHandler := api.NewUserHandler(userService)

	// TODO : Initialize Plan feature : repositories, services, and handlers
	// your code here ...

	// Setup Gin router
	router := gin.Default()
	router.Use(middleware.AccessLog())
	router.Use(gin.Recovery())

	// Register routes
	userHandler.RegisterRoutes(router)
	// TODO : register Plan routes
	// your code here ...

	addr := net.JoinHostPort("localhost", os.Getenv("PORT"))
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		slog.InfoContext(ctx, fmt.Sprintf("Server start listening at %s", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.ErrorContext(ctx, "listen and start server error", slog.String("error", err.Error()))
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.InfoContext(ctx, "Shutting down server")

	// Perform any necessary cleanup here
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "Server Shutdown", slog.String("error", err.Error()))
		return
	}

	slog.InfoContext(ctx, "Server exiting")
}
