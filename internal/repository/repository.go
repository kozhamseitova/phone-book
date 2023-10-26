package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kozhamseitova/phone-book/pkg/logger"
)

type Repository interface {

}

type repository struct {
	pool *pgxpool.Pool
	logger logger.Logger
}

func New(pool *pgxpool.Pool, logger logger.Logger) Repository {
	return &repository{
		pool: pool,
		logger: logger,
	}
}