package transaction

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type TransactionRepo interface {
	Create(ctx context.Context, m *model.Transaction) (*model.Transaction, error)
	GetOne(ctx context.Context, m *model.Transaction) (*model.Transaction, error)
	ListTransaction(ctx context.Context, walletID string) ([]model.Transaction, error)
}

type transactionRepo struct {
	dbdget db.DBGormDelegate
}

func NewTransactionRepo(dbdget db.DBGormDelegate) TransactionRepo {
	return &transactionRepo{
		dbdget: dbdget,
	}
}
