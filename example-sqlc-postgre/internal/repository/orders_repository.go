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

func (o *OrderRepo) Save(ctx context.Context) {
}
