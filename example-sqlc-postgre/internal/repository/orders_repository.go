package repository

import (
	"context"
	"example-sqlc-postgre/internal/db/sqlcdb"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type OrderRepo struct {
	pool    *pgxpool.Pool
	logger  *zerolog.Logger
	queries *sqlcdb.Queries
}

func (o *OrderRepo) Save(ctx context.Context, u *sqlcdb.UpsertOrdersParams) error {
	err := o.queries.UpsertOrders(ctx, u)
	if err != nil {
		o.logger.Error().Err(err).Msgf("save order failed, params: %v", u)
		return err
	}
	return nil
}

func (o *OrderRepo) GetOrdersByStatuses(ctx context.Context, params []string) ([]*sqlcdb.Order, error) {
	statuses, err := o.queries.GetOrdersByStatuses(ctx, params)
	if err != nil {
		o.logger.Error().Err(err).Msgf("get orders by status failed, params: %v", params)
	}

	return statuses, err

}

func (o *OrderRepo) GetOrdersByStatusesUseIndex(ctx context.Context, params []string) ([]*sqlcdb.Order, error) {
	statuses, err := o.queries.GetOrdersByStatusesAsStrings(ctx, params)
	if err != nil {
		o.logger.Error().Err(err).Msgf("get orders by status failed, params: %v", params)
	}

	return statuses, err

}

func (o *OrderRepo) GetOrdersByStatusesUseIndex2(ctx context.Context, params []string) ([]*sqlcdb.Order, error) {
	statuses, err := o.queries.GetOrdersByStatusesAsStrings2(ctx, params)
	if err != nil {
		o.logger.Error().Err(err).Msgf("get orders by status failed, params: %v", params)
	}

	return statuses, err

}

func NewOrderRepo(pool *pgxpool.Pool, logger *zerolog.Logger, queries *sqlcdb.Queries) *OrderRepo {
	return &OrderRepo{
		pool:    pool,
		logger:  logger,
		queries: queries,
	}
}
