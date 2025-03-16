package orderservice

import (
	"context"
	"example-sqlc-postgre/internal/db/sqlcdb"
	"github.com/shopspring/decimal"
)

type OrderService interface {
	CreateOrder(ctx context.Context, param *OrderCreateParam) error
}

type OrderCreateParam struct {
	AccountId string
	Symbol    string
	Quantity  decimal.Decimal
	Type      sqlcdb.OrderType
}

type Fee struct {
	Type   string          `json:"type"`
	Amount decimal.Decimal `json:"amount"`
}
