package orderservice

import (
	"context"
	"encoding/json"
	"example-sqlc-postgre/common"
	"example-sqlc-postgre/internal/db/sqlcdb"
	"example-sqlc-postgre/internal/repository"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

const InitVersion = 1

type OrderServiceImpl struct {
	OrderRepo *repository.OrderRepo
	logger    *zerolog.Logger
}

func (o OrderServiceImpl) ListOrdersByStatus(ctx context.Context, param *FilterStatusParam) ([]*sqlcdb.Order, error) {
	var q *FilterStatusParam

	if param != nil {
		q = param
	} else {
		statuses := []sqlcdb.OrderStatus{
			sqlcdb.OrderStatusFailed, sqlcdb.OrderStatusDone,
		}
		q = &FilterStatusParam{Status: statuses}
	}

	statuses := make([]string, 0, len(q.Status))
	for _, status := range q.Status {
		statusStr := string(status)
		statuses = append(statuses, statusStr)
	}

	r, err := o.OrderRepo.GetOrdersByStatuses(ctx, statuses)
	if err != nil {
		return nil, fmt.Errorf("listOrdersByStatus error: %w", err)
	}

	return r, nil
}

func (o OrderServiceImpl) CreateOrder(ctx context.Context, param *OrderCreateParam) error {
	id := common.GenerateOrderULID()

	accountId, err := uuid.FromString(param.AccountId)
	if err != nil {
		return fmt.Errorf("parse error for accountId: %s, error: %w", param.AccountId, err)
	}

	// fee
	fees := []Fee{
		{Type: "fee1", Amount: decimal.NewFromFloat(0.99)},
		{Type: "fee2", Amount: decimal.NewFromFloat(0.99)},
	}
	feesJSON, err := json.Marshal(fees)
	if err != nil {
		return fmt.Errorf("marshal fees error: %w", err)
	}

	saveParam := &sqlcdb.UpsertOrdersParams{
		ID:        id,
		AccountID: accountId,
		Symbol:    param.Symbol,
		Quantity:  param.Quantity,
		Fees:      feesJSON,
		Status:    sqlcdb.OrderStatusInit,
		Type:      param.Type,
		Version:   InitVersion,
	}

	o.logger.Debug().Str("id", saveParam.ID).
		Str("accountId", saveParam.AccountID.String()).
		Interface("fees", fees).
		Msg("Creating new order")

	saveErr := o.OrderRepo.Save(ctx, saveParam)
	if saveErr != nil {
		return fmt.Errorf("CreateOrder save error, param: %v", param)
	}

	return nil
}

var (
	_ OrderService = (*OrderServiceImpl)(nil)
)

func NewOrderService(logger *zerolog.Logger, repo *repository.OrderRepo) *OrderServiceImpl {
	return &OrderServiceImpl{
		repo,
		logger,
	}
}
