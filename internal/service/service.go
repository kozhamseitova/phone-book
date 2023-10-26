package service

import (
	"context"

	"github.com/kozhamseitova/phone-book/internal/models"
	"github.com/kozhamseitova/phone-book/internal/repository"
	"github.com/kozhamseitova/phone-book/pkg/logger"
)

type Service interface {
	Search(ctx context.Context, search models.Search) ([]*models.Search, error)
	Create(ctx context.Context, search models.Search) error
}

type service struct {
	repository repository.Repository
	logger logger.Logger
}

func New(repository repository.Repository, logger logger.Logger) Service {
	return &service{
		repository: repository,
		logger: logger,
	}
}