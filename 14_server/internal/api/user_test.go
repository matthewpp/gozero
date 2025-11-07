package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gozero/server/internal/api"
	"gozero/server/internal/model"
	serviceMock "gozero/server/internal/service/mock_services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserHandler_CreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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
		mockService.EXPECT().CreateUser(gomock.Any(), user).DoAndReturn(
			func(ctx context.Context, u *model.User) error {
				u.ID = 1 // Simulate DB assignment
				return nil
			},
		)

		// Create request
		jsonData, err := json.Marshal(user)
		assert.NoError(t, err, "Failed to marshal user JSON")

		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusCreated, w.Code, "Expected HTTP 201 Created status")

		var response model.User
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Equal(t, int64(1), response.ID, "Expected user ID to be 1")
		assert.Equal(t, "John Doe", response.Name, "Expected user name to match")
		assert.Equal(t, "john@example.com", response.Email, "Expected user email to match")
	})

	t.Run("invalid json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Create request with invalid JSON
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code, "Expected HTTP 400 Bad Request status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response, "error", "Response should contain error field")
	})

	t.Run("service error", func(t *testing.T) {
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

		// Mock expectation - service returns error
		mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errors.New("database connection failed"))

		// Create request
		jsonData, err := json.Marshal(user)
		assert.NoError(t, err, "Failed to marshal user JSON")

		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 Internal Server Error status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response["error"], "database connection failed", "Response should contain service error message")
	})
}

func TestUserHandler_GetUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Expected user
		expectedUser := &model.User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}

		// Mock expectation
		mockService.EXPECT().GetUser(gomock.Any(), int64(1)).Return(expectedUser, nil)

		// Create request
		req, err := http.NewRequest("GET", "/users/1", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response model.User
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Equal(t, expectedUser.ID, response.ID, "User ID should match")
		assert.Equal(t, expectedUser.Name, response.Name, "User name should match")
		assert.Equal(t, expectedUser.Email, response.Email, "User email should match")
	})

	t.Run("not found", func(t *testing.T) {
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
		req, err := http.NewRequest("GET", "/users/999", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, w.Code, "Expected HTTP 404 Not Found status")

		// Check response body contains error message
		responseBody := w.Body.String()
		assert.Contains(t, responseBody, "error", "Response should contain error field")
	})

	t.Run("invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Create request with invalid ID
		req, err := http.NewRequest("GET", "/users/invalid", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code, "Expected HTTP 400 Bad Request status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response["error"], "invalid id", "Response should contain invalid id error")
	})
}

func TestUserHandler_UpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		updateData := &model.User{Name: "John Smith", Email: "johnsmith@example.com"}

		// Mock expectation
		mockService.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, u *model.User) error {
				assert.Equal(t, int64(1), u.ID, "User ID should be set from URL parameter")
				return nil
			},
		)

		// Create request
		jsonData, err := json.Marshal(updateData)
		assert.NoError(t, err, "Failed to marshal update data JSON")

		req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response model.User
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Equal(t, int64(1), response.ID, "Expected user ID to be 1")
		assert.Equal(t, updateData.Name, response.Name, "Expected updated name")
		assert.Equal(t, updateData.Email, response.Email, "Expected updated email")
	})

	t.Run("invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		updateData := &model.User{Name: "John Smith", Email: "johnsmith@example.com"}

		// Create request with invalid ID
		jsonData, err := json.Marshal(updateData)
		assert.NoError(t, err, "Failed to marshal update data JSON")

		req, err := http.NewRequest("PUT", "/users/invalid", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code, "Expected HTTP 400 Bad Request status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response["error"], "invalid id", "Response should contain invalid id error")
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		updateData := &model.User{Name: "John Smith", Email: "johnsmith@example.com"}

		// Mock expectation - service returns error
		mockService.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("update failed"))

		// Create request
		jsonData, err := json.Marshal(updateData)
		assert.NoError(t, err, "Failed to marshal update data JSON")

		req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 Internal Server Error status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response["error"], "update failed", "Response should contain service error message")
	})
}

func TestUserHandler_DeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Mock expectation
		mockService.EXPECT().DeleteUser(gomock.Any(), int64(1)).Return(nil)

		// Create request
		req, err := http.NewRequest("DELETE", "/users/1", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusNoContent, w.Code, "Expected HTTP 204 No Content status")
		assert.Empty(t, w.Body.String(), "Response body should be empty for delete operation")
	})

	t.Run("invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Create request with invalid ID
		req, err := http.NewRequest("DELETE", "/users/invalid", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code, "Expected HTTP 400 Bad Request status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response["error"], "invalid id", "Response should contain invalid id error")
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Mock expectation - service returns error
		mockService.EXPECT().DeleteUser(gomock.Any(), int64(1)).Return(errors.New("delete failed"))

		// Create request
		req, err := http.NewRequest("DELETE", "/users/1", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 Internal Server Error status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response["error"], "delete failed", "Response should contain service error message")
	})
}

func TestUserHandler_ListUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Expected users
		expectedUsers := []*model.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		}

		// Mock expectation
		mockService.EXPECT().ListUsers(gomock.Any()).Return(expectedUsers, nil)

		// Create request
		req, err := http.NewRequest("GET", "/users", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response []*model.User
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Len(t, response, 2, "Expected 2 users in response")
		assert.Equal(t, expectedUsers[0].ID, response[0].ID, "First user ID should match")
		assert.Equal(t, expectedUsers[0].Name, response[0].Name, "First user name should match")
		assert.Equal(t, expectedUsers[1].ID, response[1].ID, "Second user ID should match")
		assert.Equal(t, expectedUsers[1].Name, response[1].Name, "Second user name should match")
	})

	t.Run("empty list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Mock expectation - empty list
		mockService.EXPECT().ListUsers(gomock.Any()).Return([]*model.User{}, nil)

		// Create request
		req, err := http.NewRequest("GET", "/users", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response []*model.User
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Len(t, response, 0, "Expected empty list in response")
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockUserService(ctrl)
		handler := api.NewUserHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Mock expectation - service returns error
		mockService.EXPECT().ListUsers(gomock.Any()).Return(nil, errors.New("database connection failed"))

		// Create request
		req, err := http.NewRequest("GET", "/users", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 Internal Server Error status")

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Contains(t, response["error"], "database connection failed", "Response should contain service error message")
	})
}
