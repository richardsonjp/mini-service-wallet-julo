package credential

import (
	"context"
	"go-skeleton/internal/model"
)

func (r *credentialRepo) Create(ctx context.Context, m *model.Credential) (*model.Credential, error) {
	if err := r.dbdget.Get(ctx).
		Create(m).
		Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *credentialRepo) GetOne(ctx context.Context, m *model.Credential) (*model.Credential, error) {
	query := r.dbdget.Get(ctx).Where(m)

	if err := query.Last(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}
