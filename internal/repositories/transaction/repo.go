package transaction

import (
	"context"
	"go-skeleton/internal/model"
)

func (r *transactionRepo) Create(ctx context.Context, m *model.Transaction) (*model.Transaction, error) {
	if err := r.dbdget.Get(ctx).
		Create(m).
		Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *transactionRepo) GetOne(ctx context.Context, m *model.Transaction) (*model.Transaction, error) {
	query := r.dbdget.Get(ctx).Where(m)

	if err := query.Last(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *transactionRepo) ListTransaction(ctx context.Context, walletID string) ([]model.Transaction, error) {
	var data []model.Transaction
	err := r.dbdget.Get(ctx).
		Where("transaction.wallet_id = ?", walletID).
		Order("transaction.created_at DESC").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
