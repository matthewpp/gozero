package service

import (
	"context"
	"gozero/server/internal/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=./plan.go -destination=./mock_services/plan.go
type PlanService interface {
	CreatePlan(ctx context.Context, plan *model.Plan) error
	GetPlan(ctx context.Context, id int64) (*model.Plan, error)
	UpdatePlan(ctx context.Context, plan *model.Plan) error
	ListPlans(ctx context.Context) ([]*model.Plan, error)
}

// TODO : 1. Create a concrete implementation of PlanService for PostgreSQL or SQLite
// Your code here ...
