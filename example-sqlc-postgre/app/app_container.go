package app

import (
	"example-sqlc-postgre/internal/db/sqlcdb"
	"example-sqlc-postgre/internal/repository"
	"example-sqlc-postgre/orderservice"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// ServiceContainer holds all application services
type ServiceContainer struct {
	OrderService orderservice.OrderService
}

// NewServiceContainer initializes all services
func NewServiceContainer(pool *pgxpool.Pool, logger *zerolog.Logger, queries *sqlcdb.Queries) *ServiceContainer {
	// Initialize repositories
	orderRepo := repository.NewOrderRepo(pool, logger, queries)

	// Initialize services
	orderService := orderservice.NewOrderService(logger, orderRepo)

	return &ServiceContainer{
		OrderService: orderService,
	}
}
