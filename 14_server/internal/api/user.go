package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"gozero/server/internal/model"
	"gozero/server/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET(":id", h.GetUser)
		users.PUT(":id", h.UpdateUser)
		users.DELETE(":id", h.DeleteUser)
		users.GET("", h.ListUsers)
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	slog.InfoContext(ctx, "API: Creating user request received")

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		slog.ErrorContext(ctx, "API: Invalid JSON in create user request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateUser(ctx, &user); err != nil {
		slog.ErrorContext(ctx, "API: Failed to create user", "error", err, "name", user.Name, "email", user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.InfoContext(ctx, "API: User created successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	slog.InfoContext(ctx, "API: Get user request received", "param_id", c.Param("id"))

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.ErrorContext(ctx, "API: Invalid user ID format", "error", err, "param_id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.Service.GetUser(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "API: Failed to get user", "error", err, "id", id)

		c.AbortWithStatusJSON(http.StatusNotFound, err)
		// c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	slog.InfoContext(ctx, "API: User retrieved successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	slog.InfoContext(ctx, "API: Update user request received", "param_id", c.Param("id"))

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.ErrorContext(ctx, "API: Invalid user ID format", "error", err, "param_id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		slog.ErrorContext(ctx, "API: Invalid JSON in update user request", "error", err, "id", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = id
	if err := h.Service.UpdateUser(ctx, &user); err != nil {
		slog.ErrorContext(ctx, "API: Failed to update user", "error", err, "id", id, "name", user.Name, "email", user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.InfoContext(ctx, "API: User updated successfully", "id", user.ID, "name", user.Name, "email", user.Email)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	slog.InfoContext(ctx, "API: Delete user request received", "param_id", c.Param("id"))

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		slog.ErrorContext(ctx, "API: Invalid user ID format", "error", err, "param_id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.Service.DeleteUser(ctx, id); err != nil {
		slog.ErrorContext(ctx, "API: Failed to delete user", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.InfoContext(ctx, "API: User deleted successfully", "id", id)
	c.Status(http.StatusNoContent)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()
	slog.InfoContext(ctx, "API: List users request received")

	users, err := h.Service.ListUsers(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "API: Failed to list users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.InfoContext(ctx, "API: Users listed successfully", "count", len(users))
	c.JSON(http.StatusOK, users)
}
