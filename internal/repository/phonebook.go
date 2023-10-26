package repository

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/kozhamseitova/phone-book/internal/models"
	"github.com/kozhamseitova/phone-book/utils"

	sq "github.com/Masterminds/squirrel"
)

const phonebookTable = "phone_book"

func(r *repository) GetByPhoneAndName(ctx context.Context, search models.Search) ([]*models.Search, error){
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("phone, name").From(phonebookTable)

	if search.Phone != "" {
		builder = builder.Where("phone LIKE ?", fmt.Sprint("%", search.Phone, "%"))
	}
	if search.Name != "" {
		builder = builder.Where("name LIKE ?", fmt.Sprint("%", search.Name, "%"))
	}

	var phoneBook []*models.Search
	sql, args, err := builder.ToSql()
	if err != nil {
		r.logger.Errorf(ctx, "[GetByPhoneAndName] err: %v", err)
		return nil, utils.ErrInternalError
	}
	
	err = pgxscan.Select(ctx, r.pool, &phoneBook, sql, args...)
	if err != nil {
		r.logger.Errorf(ctx, "[GetByPhoneAndName] err: %v", err)
		return nil, utils.ErrInternalError
	}

	return phoneBook, nil
}
 
func(r *repository) Create(ctx context.Context, search models.Search) error {
	query := fmt.Sprintf(`INSERT INTO %s (phone, name) VALUES ($1, $2)`, phonebookTable)

	_, err := r.pool.Exec(ctx, query, search.Phone, search.Name)
	if err != nil {
		r.logger.Errorf(ctx, "[Create] err: %v", err)
		return utils.ErrInternalError
	}

	return nil
}