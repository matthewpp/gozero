package model

import "github.com/shopspring/decimal"

type Plan struct {
	ID      int64           `json:"id" db:"id"`
	Code    string          `json:"code" db:"code"`
	Name    string          `json:"name" db:"name"`
	Premium decimal.Decimal `json:"premium" db:"premium"`
}
