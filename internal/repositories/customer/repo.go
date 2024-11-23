package customer

import (
	"context"
	"go-skeleton/internal/model"
)

func (r *customerRepo) Create(ctx context.Context, m *model.Customer) (*model.Customer, error) {
	if err := r.dbdget.Get(ctx).
		Create(m).
		Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *customerRepo) GetOne(ctx context.Context, m *model.Customer) (*model.Customer, error) {
	query := r.dbdget.Get(ctx).Where(m)

	if err := query.Last(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *customerRepo) Update(ctx context.Context, param *model.Customer, updatedFields ...string) (int64, error) {
	query := r.dbdget.Get(ctx)
	if len(updatedFields) > 0 {
		updatedFields = append(updatedFields, "updated_at")
		query = query.Select(updatedFields)
	}

	query.Updates(param)

	// execute query
	if err := query.Find(&param).Error; err != nil {
		return 0, err
	}

	return query.RowsAffected, nil
}
