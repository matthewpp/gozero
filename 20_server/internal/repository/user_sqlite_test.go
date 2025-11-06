package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"gozero/server/internal/model"
	"gozero/server/internal/repository"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// setupSQLiteTestDB creates a test database connection and schema
func setupSQLiteTestDB(t *testing.T) (*sql.DB, func()) {
	dbFile := fmt.Sprintf("test_users_%s.db", t.Name())

	// Create database connection
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		t.Fatalf("Failed to open SQLite database: %v", err)
	}

	// Create the users table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);
	`
	if _, err := db.Exec(createTableSQL); err != nil {
		t.Fatalf("Failed to create users table: %v", err)
	}

	// Return cleanup function
	cleanup := func() {
		db.Close()
		os.Remove(dbFile)
	}

	return db, cleanup
}

func TestUserSQLiteRepository_Create(t *testing.T) {
	db, cleanup := setupSQLiteTestDB(t)
	defer cleanup()

	repo := repository.NewUserSQLiteRepository(db)
	ctx := context.Background()

	t.Run("successful creation", func(t *testing.T) {
		user := &model.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		err := repo.Create(ctx, user)
		assert.NoError(t, err, "Failed to create user")
		assert.NotZero(t, user.ID, "Expected user ID to be set after creation")
	})

	t.Run("duplicate email constraint", func(t *testing.T) {
		user1 := &model.User{
			Name:  "User One",
			Email: "unique@example.com",
		}
		err := repo.Create(ctx, user1)
		assert.NoError(t, err, "Failed to create first user")

		user2 := &model.User{
			Name:  "User Two",
			Email: "unique@example.com", // Same email
		}
		err = repo.Create(ctx, user2)
		assert.Error(t, err, "Expected error when creating user with duplicate email")
		assert.Contains(t, err.Error(), "UNIQUE constraint failed", "Error should mention unique constraint")
	})
}

func TestUserSQLiteRepository_GetByID(t *testing.T) {
	db, cleanup := setupSQLiteTestDB(t)
	defer cleanup()

	repo := repository.NewUserSQLiteRepository(db)
	ctx := context.Background()

	t.Run("successful retrieval", func(t *testing.T) {
		// First create a user
		user := &model.User{
			Name:  "Jane Doe",
			Email: "jane@example.com",
		}
		err := repo.Create(ctx, user)
		assert.NoError(t, err, "Failed to create user")

		// Then retrieve it
		retrieved, err := repo.GetByID(ctx, user.ID)
		assert.NoError(t, err, "Failed to get user by ID")
		assert.NotNil(t, retrieved, "Retrieved user should not be nil")
		assert.Equal(t, user.ID, retrieved.ID, "User IDs should match")
		assert.Equal(t, user.Name, retrieved.Name, "User names should match")
		assert.Equal(t, user.Email, retrieved.Email, "User emails should match")
	})

	t.Run("non-existent user", func(t *testing.T) {
		retrieved, err := repo.GetByID(ctx, 99999)
		assert.Error(t, err, "Expected error for non-existent user")
		assert.Equal(t, sql.ErrNoRows, err, "Expected sql.ErrNoRows")
		assert.Nil(t, retrieved, "Retrieved user should be nil")
	})
}

func TestUserSQLiteRepository_Update(t *testing.T) {
	db, cleanup := setupSQLiteTestDB(t)
	defer cleanup()

	repo := repository.NewUserSQLiteRepository(db)
	ctx := context.Background()

	t.Run("successful update", func(t *testing.T) {
		// First create a user
		user := &model.User{
			Name:  "Bob Smith",
			Email: "bob@example.com",
		}
		err := repo.Create(ctx, user)
		assert.NoError(t, err, "Failed to create user")

		// Update the user
		originalID := user.ID
		user.Name = "Bob Johnson"
		user.Email = "bob.johnson@example.com"
		err = repo.Update(ctx, user)
		assert.NoError(t, err, "Failed to update user")

		// Verify the update
		updated, err := repo.GetByID(ctx, user.ID)
		assert.NoError(t, err, "Failed to get updated user")
		assert.NotNil(t, updated, "Updated user should not be nil")
		assert.Equal(t, originalID, updated.ID, "User ID should remain unchanged")
		assert.Equal(t, "Bob Johnson", updated.Name, "User name should be updated")
		assert.Equal(t, "bob.johnson@example.com", updated.Email, "User email should be updated")
	})

	t.Run("non-existent user", func(t *testing.T) {
		user := &model.User{
			ID:    99999,
			Name:  "Non Existent",
			Email: "nonexistent@example.com",
		}
		err := repo.Update(ctx, user)
		assert.Error(t, err, "Expected error for non-existent user")
		assert.Equal(t, sql.ErrNoRows, err, "Expected sql.ErrNoRows")
	})
}

func TestUserSQLiteRepository_Delete(t *testing.T) {
	db, cleanup := setupSQLiteTestDB(t)
	defer cleanup()

	repo := repository.NewUserSQLiteRepository(db)
	ctx := context.Background()

	t.Run("successful deletion", func(t *testing.T) {
		// First create a user
		user := &model.User{
			Name:  "To Delete",
			Email: "delete@example.com",
		}
		err := repo.Create(ctx, user)
		assert.NoError(t, err, "Failed to create user")

		// Delete the user
		err = repo.Delete(ctx, user.ID)
		assert.NoError(t, err, "Failed to delete user")

		// Verify deletion
		deleted, err := repo.GetByID(ctx, user.ID)
		assert.Error(t, err, "Expected error when trying to get deleted user")
		assert.Equal(t, sql.ErrNoRows, err, "Expected sql.ErrNoRows")
		assert.Nil(t, deleted, "Deleted user should be nil")
	})

	t.Run("non-existent user", func(t *testing.T) {
		err := repo.Delete(ctx, 99999)
		assert.Error(t, err, "Expected error for non-existent user")
		assert.Equal(t, sql.ErrNoRows, err, "Expected sql.ErrNoRows")
	})
}

func TestUserSQLiteRepository_List(t *testing.T) {
	db, cleanup := setupSQLiteTestDB(t)
	defer cleanup()

	repo := repository.NewUserSQLiteRepository(db)
	ctx := context.Background()

	t.Run("empty list", func(t *testing.T) {
		users, err := repo.List(ctx)
		assert.NoError(t, err, "Failed to list users")
		assert.Empty(t, users, "Expected empty list")
		assert.Len(t, users, 0, "Expected empty list with 0 length")
	})

	t.Run("list multiple users", func(t *testing.T) {
		// Create multiple users
		testUsers := []*model.User{
			{Name: "Alice", Email: "alice@example.com"},
			{Name: "Charlie", Email: "charlie@example.com"},
			{Name: "David", Email: "david@example.com"},
		}

		for _, user := range testUsers {
			err := repo.Create(ctx, user)
			assert.NoError(t, err, "Failed to create user %s", user.Name)
		}

		// List all users
		allUsers, err := repo.List(ctx)
		assert.NoError(t, err, "Failed to list users")
		assert.NotNil(t, allUsers, "Users list should not be nil")
		assert.Len(t, allUsers, len(testUsers), "Expected %d users", len(testUsers))

		// Verify users are ordered by ID
		for i := 1; i < len(allUsers); i++ {
			assert.Greater(t, allUsers[i].ID, allUsers[i-1].ID, "Users should be ordered by ID (ascending)")
		}

		// Verify all created users are in the list
		expectedNames := []string{"Alice", "Charlie", "David"}
		actualNames := make([]string, len(allUsers))
		for i, user := range allUsers {
			actualNames[i] = user.Name
		}
		assert.ElementsMatch(t, expectedNames, actualNames, "All created users should be in the list")
	})
}
