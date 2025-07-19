package db

import (
	"context"
	"effective_mobile/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Init() (*pgxpool.Pool, error) {

	dsn := GetConfig()

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Error("Database connection error", zap.Error(err))
		return nil, err
	}

	logger.Info("Successful database connection")
	return pool, nil
}

func Close(pool *pgxpool.Pool) {
	pool.Close()
}
