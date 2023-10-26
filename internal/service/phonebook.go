package service

import (
	"context"

	"github.com/kozhamseitova/phone-book/internal/models"
)

func(s *service) Search(ctx context.Context, search models.Search) ([]*models.Search, error) {
	return s.repository.GetByPhoneAndName(ctx, search)
}

func(s *service) Create(ctx context.Context, search models.Search) error {
	return s.repository.Create(ctx, search)
}