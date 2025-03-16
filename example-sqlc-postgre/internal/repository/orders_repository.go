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

func NewOrderRepo(pool *pgxpool.Pool, logger *zerolog.Logger, queries *sqlcdb.Queries) *OrderRepo {
	return &OrderRepo{
		pool:    pool,
		logger:  logger,
		queries: queries,
	}
}
