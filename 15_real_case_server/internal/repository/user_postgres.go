package repository

import (
	"context"
	"database/sql"
	"errors"
	"gozero/server/internal/errs"
	"gozero/server/internal/model"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userPostgresqlRepository struct {
	db *pgxpool.Pool
}

func NewUserPostgresRepository(db *pgxpool.Pool) UserRepository {
	return &userPostgresqlRepository{
		db: db,
	}
}

func (r *userPostgresqlRepository) Create(ctx context.Context, user *model.User) error {
	slog.InfoContext(ctx, "Creating user", "name", user.Name, "email", user.Email)

	row := r.db.QueryRow(ctx, "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email)
	err := row.Scan(&user.ID)

	if err != nil {
		slog.ErrorContext(ctx, "Failed to create user", "error", err, "name", user.Name, "email", user.Email)
		return err
	}

	slog.InfoContext(ctx, "User created successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	return nil
}

func (r *userPostgresqlRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	slog.InfoContext(ctx, "Getting user by ID", "id", id)

	var user model.User
	err := r.db.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.InfoContext(ctx, "User not found in PostgreSQL", "id", id)
			return nil, errors.Join(errs.ErrInvalidUserID, err)
		}

		slog.ErrorContext(ctx, "Failed to get user by ID", "error", err, "id", id)
		return nil, err
	}

	slog.InfoContext(ctx, "User retrieved successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	return &user, nil
}

func (r *userPostgresqlRepository) Update(ctx context.Context, user *model.User) error {
	slog.InfoContext(ctx, "Updating user", "id", user.ID, "name", user.Name, "email", user.Email)

	_, err := r.db.Exec(ctx, "UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to update user", "error", err, "id", user.ID)
		return err
	}

	slog.InfoContext(ctx, "User updated successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	return nil
}

func (r *userPostgresqlRepository) Delete(ctx context.Context, id int64) error {
	slog.InfoContext(ctx, "Deleting user", "id", id)

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete user", "error", err, "id", id)
		return err
	}

	slog.InfoContext(ctx, "User deleted successfully", "id", id)
	return nil
}

func (r *userPostgresqlRepository) List(ctx context.Context) ([]*model.User, error) {
	slog.InfoContext(ctx, "Listing all users")

	rows, err := r.db.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		slog.ErrorContext(ctx, "Failed to list users", "error", err)
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			slog.ErrorContext(ctx, "Failed to scan user row", "error", err)
			return nil, err
		}
		users = append(users, &user)
	}

	slog.InfoContext(ctx, "Users listed successfully", "count", len(users))
	return users, nil
}
