package service_test

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
	"gozero/server/internal/model"
	repositoryMock "gozero/server/internal/repository/mock"
	"gozero/server/internal/service"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repositoryMock.NewMockUserRepository(ctrl)
	svc := service.NewUserService(repo)

	user := &model.User{Name: "test", Email: "test@example.com"}
	repo.EXPECT().Create(gomock.Any(), user).Return(nil)

	err := svc.CreateUser(context.Background(), user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestUserService_GetUser_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repositoryMock.NewMockUserRepository(ctrl)
	svc := service.NewUserService(repo)

	repo.EXPECT().GetByID(gomock.Any(), int64(1)).Return(nil, errors.New("not found"))

	_, err := svc.GetUser(context.Background(), 1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
