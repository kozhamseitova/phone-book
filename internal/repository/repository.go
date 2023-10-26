package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kozhamseitova/phone-book/internal/models"
	"github.com/kozhamseitova/phone-book/pkg/logger"
)

type Repository interface {
	GetByPhoneAndName(ctx context.Context, search models.Search) ([]*models.Search, error)
	Create(ctx context.Context, search models.Search) error
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