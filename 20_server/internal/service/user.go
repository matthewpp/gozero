package service

import (
	"context"
	"log/slog"
	"gozero/server/internal/model"
	"gozero/server/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context) ([]*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	slog.InfoContext(ctx, "Service: Creating user", "name", user.Name, "email", user.Email)
	
	err := s.repo.Create(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "Service: Failed to create user", "error", err, "name", user.Name, "email", user.Email)
		return err
	}
	
	slog.InfoContext(ctx, "Service: User created successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	return nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	slog.InfoContext(ctx, "Service: Getting user", "id", id)
	
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "Service: Failed to get user", "error", err, "id", id)
		return nil, err
	}
	
	slog.InfoContext(ctx, "Service: User retrieved successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user *model.User) error {
	slog.InfoContext(ctx, "Service: Updating user", "id", user.ID, "name", user.Name, "email", user.Email)
	
	err := s.repo.Update(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "Service: Failed to update user", "error", err, "id", user.ID)
		return err
	}
	
	slog.InfoContext(ctx, "Service: User updated successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	return nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	slog.InfoContext(ctx, "Service: Deleting user", "id", id)
	
	err := s.repo.Delete(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "Service: Failed to delete user", "error", err, "id", id)
		return err
	}
	
	slog.InfoContext(ctx, "Service: User deleted successfully", "id", id)
	return nil
}

func (s *userService) ListUsers(ctx context.Context) ([]*model.User, error) {
	slog.InfoContext(ctx, "Service: Listing users")
	
	users, err := s.repo.List(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Service: Failed to list users", "error", err)
		return nil, err
	}
	
	slog.InfoContext(ctx, "Service: Users listed successfully", "count", len(users))
	return users, nil
}
