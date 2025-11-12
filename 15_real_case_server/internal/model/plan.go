package model

import "github.com/shopspring/decimal"

type Plan struct {
	ID      int64           `json:"id" db:"id"`
	Code    string          `json:"code" db:"code" binding:"required,min=1,max=50"`
	Name    string          `json:"name" db:"name" binding:"required,min=1,max=200"`
	Premium decimal.Decimal `json:"premium" db:"premium" binding:"required,gt=0"`
}
