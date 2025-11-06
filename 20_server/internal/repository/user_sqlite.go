package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"gozero/server/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

type userSQLiteRepository struct {
	db *sql.DB
}

func NewUserSQLiteRepository(db *sql.DB) UserRepository {
	return &userSQLiteRepository{
		db: db,
	}
}

func (r *userSQLiteRepository) Create(ctx context.Context, user *model.User) error {
	slog.InfoContext(ctx, "Creating user in SQLite", "name", user.Name, "email", user.Email)

	result, err := r.db.ExecContext(ctx, "INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create user in SQLite", "error", err, "name", user.Name, "email", user.Email)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get last insert ID", "error", err)
		return err
	}

	user.ID = id
	slog.InfoContext(ctx, "User created successfully in SQLite", "id", user.ID, "name", user.Name, "email", user.Email)
	return nil
}

func (r *userSQLiteRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	slog.InfoContext(ctx, "Getting user by ID from SQLite", "id", id)

	var user model.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.InfoContext(ctx, "User not found in SQLite", "id", id)
			return nil, err
		}
		slog.ErrorContext(ctx, "Failed to get user by ID from SQLite", "error", err, "id", id)
		return nil, err
	}

	slog.InfoContext(ctx, "User retrieved successfully from SQLite", "id", user.ID, "name", user.Name, "email", user.Email)
	return &user, nil
}

func (r *userSQLiteRepository) Update(ctx context.Context, user *model.User) error {
	slog.InfoContext(ctx, "Updating user in SQLite", "id", user.ID, "name", user.Name, "email", user.Email)

	result, err := r.db.ExecContext(ctx, "UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to update user in SQLite", "error", err, "id", user.ID)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get rows affected", "error", err)
		return err
	}

	if rowsAffected == 0 {
		slog.InfoContext(ctx, "No user found to update in SQLite", "id", user.ID)
		return sql.ErrNoRows
	}

	slog.InfoContext(ctx, "User updated successfully in SQLite", "id", user.ID, "name", user.Name, "email", user.Email)
	return nil
}

func (r *userSQLiteRepository) Delete(ctx context.Context, id int64) error {
	slog.InfoContext(ctx, "Deleting user from SQLite", "id", id)

	result, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete user from SQLite", "error", err, "id", id)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get rows affected", "error", err)
		return err
	}

	if rowsAffected == 0 {
		slog.InfoContext(ctx, "No user found to delete in SQLite", "id", id)
		return sql.ErrNoRows
	}

	slog.InfoContext(ctx, "User deleted successfully from SQLite", "id", id)
	return nil
}

func (r *userSQLiteRepository) List(ctx context.Context) ([]*model.User, error) {
	slog.InfoContext(ctx, "Listing all users from SQLite")

	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		slog.ErrorContext(ctx, "Failed to list users from SQLite", "error", err)
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			slog.ErrorContext(ctx, "Failed to scan user row from SQLite", "error", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		slog.ErrorContext(ctx, "Error occurred during row iteration", "error", err)
		return nil, err
	}

	slog.InfoContext(ctx, "Users listed successfully from SQLite", "count", len(users))
	return users, nil
}

func (r *userSQLiteRepository) Close() error {
	slog.Info("Closing SQLite database connection")
	return r.db.Close()
}
