package service_test

import (
	"context"
	"errors"
	"testing"

	"gozero/server/internal/model"
	repositoryMock "gozero/server/internal/repository/mock_repository"
	"gozero/server/internal/service"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		user := &model.User{Name: "test", Email: "test@example.com"}
		repo.EXPECT().Create(gomock.Any(), user).Return(nil)

		err := svc.CreateUser(context.Background(), user)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		user := &model.User{Name: "test", Email: "test@example.com"}
		expectedError := errors.New("database error")
		repo.EXPECT().Create(gomock.Any(), user).Return(expectedError)

		err := svc.CreateUser(context.Background(), user)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestUserService_GetUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		expectedUser := &model.User{ID: 1, Name: "test", Email: "test@example.com"}
		repo.EXPECT().GetByID(gomock.Any(), int64(1)).Return(expectedUser, nil)

		user, err := svc.GetUser(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		expectedError := errors.New("not found")
		repo.EXPECT().GetByID(gomock.Any(), int64(1)).Return(nil, expectedError)

		user, err := svc.GetUser(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		user := &model.User{ID: 1, Name: "updated", Email: "updated@example.com"}
		repo.EXPECT().Update(gomock.Any(), user).Return(nil)

		err := svc.UpdateUser(context.Background(), user)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		user := &model.User{ID: 1, Name: "updated", Email: "updated@example.com"}
		expectedError := errors.New("update failed")
		repo.EXPECT().Update(gomock.Any(), user).Return(expectedError)

		err := svc.UpdateUser(context.Background(), user)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestUserService_DeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		repo.EXPECT().Delete(gomock.Any(), int64(1)).Return(nil)

		err := svc.DeleteUser(context.Background(), 1)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		expectedError := errors.New("delete failed")
		repo.EXPECT().Delete(gomock.Any(), int64(1)).Return(expectedError)

		err := svc.DeleteUser(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestUserService_ListUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		expectedUsers := []*model.User{
			{ID: 1, Name: "user1", Email: "user1@example.com"},
			{ID: 2, Name: "user2", Email: "user2@example.com"},
		}
		repo.EXPECT().List(gomock.Any()).Return(expectedUsers, nil)

		users, err := svc.ListUsers(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)
		assert.Len(t, users, 2)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := repositoryMock.NewMockUserRepository(ctrl)
		svc := service.NewUserService(repo)

		expectedError := errors.New("list failed")
		repo.EXPECT().List(gomock.Any()).Return(nil, expectedError)

		users, err := svc.ListUsers(context.Background())
		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, expectedError, err)
	})
}
