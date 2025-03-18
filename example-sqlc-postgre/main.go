package main

import (
	"context"
	"example-sqlc-postgre/app"
	"example-sqlc-postgre/handler"
	"example-sqlc-postgre/internal/db/sqlcdb"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

func main() {
	// 配置 zerolog
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// 初始化数据库连接池
	connString := "postgresql://root:root@localhost:5432/dev?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer pool.Close()

	// 检查数据库连接
	if err := pool.Ping(context.Background()); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping database")
	}

	// 创建 sqlc queries
	queries := sqlcdb.New(pool)

	// 初始化服务容器
	services := app.NewServiceContainer(pool, &logger, queries)

	// 创建 order handler
	orderHandler := handler.NewOrderHandler(&logger, services.OrderService)

	// 初始化路由
	r := chi.NewRouter()

	// 注册路由
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/order/create", orderHandler.CreateOrderHandler)
		r.Post("/order/list-by-status", orderHandler.ListOrdersByStatusHandler)
		r.Post("/order/list-by-status-use-index", orderHandler.ListOrdersByStatusUseIndexHandler)
		r.Post("/order/list-by-status-use-index2", orderHandler.ListOrdersByStatusUseIndex2Handler)
	})

	// 启动服务器
	serverAddr := ":8080"

	logger.Info().Msgf("Server starting on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, r); err != nil {
		logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
