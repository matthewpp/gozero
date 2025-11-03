package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"gozero/server/internal/api"
	"gozero/server/internal/model"
	serviceMock "gozero/server/internal/service/mock"
)

func TestUserHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := serviceMock.NewMockUserService(ctrl)
	handler := api.NewUserHandler(mockService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler.RegisterRoutes(router)

	// Test data
	user := &model.User{Name: "John Doe", Email: "john@example.com"}

	// Mock expectation
	mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, u *model.User) error {
			u.ID = 1 // Simulate DB assignment
			return nil
		},
	)

	// Create request
	jsonData, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var response model.User
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.ID != 1 {
		t.Errorf("expected ID 1, got %d", response.ID)
	}
}

func TestUserHandler_GetUser_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := serviceMock.NewMockUserService(ctrl)
	handler := api.NewUserHandler(mockService)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler.RegisterRoutes(router)

	// Mock expectation
	mockService.EXPECT().GetUser(gomock.Any(), int64(999)).Return(nil, errors.New("user not found"))

	// Create request
	req, _ := http.NewRequest("GET", "/users/999", nil)

	// Record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}