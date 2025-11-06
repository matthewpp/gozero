package api

import (
	"gozero/server/internal/service"

	"github.com/gin-gonic/gin"
)

type PlanHandler struct {
	Service service.PlanService
}

func NewPlanHandler(s service.PlanService) *PlanHandler {
	return &PlanHandler{
		Service: s,
	}
}

func (h *PlanHandler) RegisterRoutes(r *gin.Engine) {
	plans := r.Group("/plans")
	{
		plans.POST("", h.CreatePlan)
		plans.GET(":id", h.GetPlan)
		plans.PUT(":id", h.UpdatePlan)
		plans.GET("", h.ListPlans)
	}
}

func (h *PlanHandler) CreatePlan(c *gin.Context) {
	// TODO: implement CreatePlan handler
}

func (h *PlanHandler) GetPlan(c *gin.Context) {
	// TODO: implement GetPlan handler
}

func (h *PlanHandler) UpdatePlan(c *gin.Context) {
	// TODO: implement UpdatePlan handler
}

func (h *PlanHandler) ListPlans(c *gin.Context) {
	// TODO: implement ListPlans handler
}
