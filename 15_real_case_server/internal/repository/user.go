package repository

import (
	"context"
	"gozero/server/internal/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=./user.go -destination=./mock_repository/user.go
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*model.User, error)
}
