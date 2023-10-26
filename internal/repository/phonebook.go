package repository

import (
	"context"

	"github.com/kozhamseitova/phone-book/internal/models"
)

func(r *repository) GetByPhoneAndName(ctx context.Context, search models.Search) ([]*models.Search, error){
	return nil, nil 
}
 
func(r *repository) Create(ctx context.Context, search models.Search) error {
	return nil
}