package orderservice

import (
	"context"
	"example-sqlc-postgre/internal/db/sqlcdb"
	"github.com/shopspring/decimal"
)

type OrderService interface {
	CreateOrder(ctx context.Context, param *OrderCreateParam) error
	ListOrdersByStatus(ctx context.Context, param *FilterStatusParam) ([]*sqlcdb.Order, error)
	ListOrdersByStatusUserIndex(ctx context.Context, param *FilterStatusParam) ([]*sqlcdb.Order, error)
	ListOrdersByStatusUserIndex2(ctx context.Context, param *FilterStatusParam) ([]*sqlcdb.Order, error)
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

type FilterStatusParam struct {
	Status []sqlcdb.OrderStatus
}
