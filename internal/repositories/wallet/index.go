package wallet

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type WalletRepo interface {
	Create(ctx context.Context, m *model.Wallet) error
	GetOne(ctx context.Context, m *model.Wallet) (*model.Wallet, error)
	Update(ctx context.Context, param *model.Wallet, updatedFields ...string) (int64, error)
}

type walletRepo struct {
	dbdget db.DBGormDelegate
}

func NewWalletRepo(dbdget db.DBGormDelegate) WalletRepo {
	return &walletRepo{
		dbdget: dbdget,
	}
}
