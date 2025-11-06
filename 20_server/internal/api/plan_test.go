package api_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gozero/server/internal/api"
	"gozero/server/internal/model"
	serviceMock "gozero/server/internal/service/mock_services"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPlanHandler_CreatePlan(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		plan := &model.Plan{
			Code:    "BASIC",
			Name:    "Basic Plan",
			Premium: decimal.NewFromFloat(99.99),
		}

		// Mock expectation
		mockService.EXPECT().CreatePlan(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, p *model.Plan) error {
				p.ID = 1 // Simulate DB assignment
				return nil
			},
		)

		// Create request
		jsonData, err := json.Marshal(plan)
		assert.NoError(t, err, "Failed to marshal plan JSON")

		req, err := http.NewRequest("POST", "/plans", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusCreated, w.Code, "Expected HTTP 201 Created status")

		var response model.Plan
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Equal(t, int64(1), response.ID, "Expected plan ID to be 1")
		assert.Equal(t, "BASIC", response.Code, "Expected plan code to match")
		assert.Equal(t, "Basic Plan", response.Name, "Expected plan name to match")
		assert.True(t, plan.Premium.Equal(response.Premium), "Expected plan premium to match")
	})

	t.Run("invalid json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Create request with invalid JSON
		req, err := http.NewRequest("POST", "/plans", bytes.NewBuffer([]byte("invalid json")))
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

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		plan := &model.Plan{
			Code:    "BASIC",
			Name:    "Basic Plan",
			Premium: decimal.NewFromFloat(99.99),
		}

		// Mock expectation - service returns error
		mockService.EXPECT().CreatePlan(gomock.Any(), gomock.Any()).Return(errors.New("database connection failed"))

		// Create request
		jsonData, err := json.Marshal(plan)
		assert.NoError(t, err, "Failed to marshal plan JSON")

		req, err := http.NewRequest("POST", "/plans", bytes.NewBuffer(jsonData))
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

func TestPlanHandler_GetPlan(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Expected plan
		expectedPlan := &model.Plan{
			ID:      1,
			Code:    "BASIC",
			Name:    "Basic Plan",
			Premium: decimal.NewFromFloat(99.99),
		}

		// Mock expectation
		mockService.EXPECT().GetPlan(gomock.Any(), int64(1)).Return(expectedPlan, nil)

		// Create request
		req, err := http.NewRequest("GET", "/plans/1", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response model.Plan
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Equal(t, expectedPlan.ID, response.ID, "Plan ID should match")
		assert.Equal(t, expectedPlan.Code, response.Code, "Plan code should match")
		assert.Equal(t, expectedPlan.Name, response.Name, "Plan name should match")
		assert.True(t, expectedPlan.Premium.Equal(response.Premium), "Plan premium should match")
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Mock expectation
		mockService.EXPECT().GetPlan(gomock.Any(), int64(999)).Return(nil, sql.ErrNoRows)

		// Create request
		req, err := http.NewRequest("GET", "/plans/999", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, w.Code, "Expected HTTP 404 Not Found status")

		// Check response body contains error message
		responseBody := w.Body.String()
		assert.Contains(t, responseBody, "error", "Response should contain error field")
		assert.Contains(t, responseBody, "plan not found", "Response should contain plan not found message")
	})

	t.Run("invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Create request with invalid ID
		req, err := http.NewRequest("GET", "/plans/invalid", nil)
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

func TestPlanHandler_UpdatePlan(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		updateData := &model.Plan{
			Code:    "PREMIUM",
			Name:    "Premium Plan",
			Premium: decimal.NewFromFloat(199.99),
		}

		// Mock expectation
		mockService.EXPECT().UpdatePlan(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, p *model.Plan) error {
				assert.Equal(t, int64(1), p.ID, "Plan ID should be set from URL parameter")
				return nil
			},
		)

		// Create request
		jsonData, err := json.Marshal(updateData)
		assert.NoError(t, err, "Failed to marshal update data JSON")

		req, err := http.NewRequest("PUT", "/plans/1", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create HTTP request")
		req.Header.Set("Content-Type", "application/json")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response model.Plan
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Equal(t, int64(1), response.ID, "Expected plan ID to be 1")
		assert.Equal(t, updateData.Code, response.Code, "Expected updated code")
		assert.Equal(t, updateData.Name, response.Name, "Expected updated name")
		assert.True(t, updateData.Premium.Equal(response.Premium), "Expected updated premium")
	})

	t.Run("invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		updateData := &model.Plan{
			Code:    "PREMIUM",
			Name:    "Premium Plan",
			Premium: decimal.NewFromFloat(199.99),
		}

		// Create request with invalid ID
		jsonData, err := json.Marshal(updateData)
		assert.NoError(t, err, "Failed to marshal update data JSON")

		req, err := http.NewRequest("PUT", "/plans/invalid", bytes.NewBuffer(jsonData))
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

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Test data
		updateData := &model.Plan{
			Code:    "PREMIUM",
			Name:    "Premium Plan",
			Premium: decimal.NewFromFloat(199.99),
		}

		// Mock expectation - service returns error
		mockService.EXPECT().UpdatePlan(gomock.Any(), gomock.Any()).Return(errors.New("update failed"))

		// Create request
		jsonData, err := json.Marshal(updateData)
		assert.NoError(t, err, "Failed to marshal update data JSON")

		req, err := http.NewRequest("PUT", "/plans/1", bytes.NewBuffer(jsonData))
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

func TestPlanHandler_ListPlans(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Expected plans
		expectedPlans := []*model.Plan{
			{ID: 1, Code: "BASIC", Name: "Basic Plan", Premium: decimal.NewFromFloat(99.99)},
			{ID: 2, Code: "PREMIUM", Name: "Premium Plan", Premium: decimal.NewFromFloat(199.99)},
		}

		// Mock expectation
		mockService.EXPECT().ListPlans(gomock.Any()).Return(expectedPlans, nil)

		// Create request
		req, err := http.NewRequest("GET", "/plans", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response []*model.Plan
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Len(t, response, 2, "Expected 2 plans in response")
		assert.Equal(t, expectedPlans[0].ID, response[0].ID, "First plan ID should match")
		assert.Equal(t, expectedPlans[0].Code, response[0].Code, "First plan code should match")
		assert.Equal(t, expectedPlans[1].ID, response[1].ID, "Second plan ID should match")
		assert.Equal(t, expectedPlans[1].Code, response[1].Code, "Second plan code should match")
	})

	t.Run("empty list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Mock expectation - empty list
		mockService.EXPECT().ListPlans(gomock.Any()).Return([]*model.Plan{}, nil)

		// Create request
		req, err := http.NewRequest("GET", "/plans", nil)
		assert.NoError(t, err, "Failed to create HTTP request")

		// Record response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK status")

		var response []*model.Plan
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err, "Failed to unmarshal response JSON")
		assert.Len(t, response, 0, "Expected empty list in response")
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := serviceMock.NewMockPlanService(ctrl)
		handler := api.NewPlanHandler(mockService)

		// Setup Gin router
		gin.SetMode(gin.TestMode)
		router := gin.New()
		handler.RegisterRoutes(router)

		// Mock expectation - service returns error
		mockService.EXPECT().ListPlans(gomock.Any()).Return(nil, errors.New("database connection failed"))

		// Create request
		req, err := http.NewRequest("GET", "/plans", nil)
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
