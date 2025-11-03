package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"gozero/server/internal/api"
	"gozero/server/internal/repository"
	"gozero/server/internal/service"
)

func main() {
	// Initialize structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	ctx := context.Background()
	
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		slog.ErrorContext(ctx, "DATABASE_URL environment variable not set")
		log.Fatal("DATABASE_URL not set")
	}

	slog.InfoContext(ctx, "Connecting to database", "dsn", dsn)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to connect to database", "error", err)
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	slog.InfoContext(ctx, "Successfully connected to database")

	repo := repository.NewUserRepository(pool)
	service := service.NewUserService(repo)
	handler := api.NewUserHandler(service)

	r := gin.Default()
	
	// Add logging middleware
	r.Use(func(c *gin.Context) {
		slog.InfoContext(c.Request.Context(), "HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"remote_addr", c.ClientIP(),
		)
		c.Next()
		slog.InfoContext(c.Request.Context(), "HTTP Response",
			"status", c.Writer.Status(),
		)
	})

	handler.RegisterRoutes(r)

	slog.InfoContext(ctx, "Starting server on :8080")
	if err := r.Run(); err != nil {
		slog.ErrorContext(ctx, "Server error", "error", err)
		log.Fatalf("server error: %v", err)
	}
}
