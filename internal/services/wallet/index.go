package wallet

import (
	"context"
	repos "go-skeleton/internal/repositories"
)

type WalletService interface {
	CreateWallet(ctx context.Context, customerID string) error
	ToggleWallet(ctx context.Context, payload ToggleWallet) (*WalletResponse, error)
	GetWalletBalance(ctx context.Context, customerID string) (*WalletResponse, error)
	Deduct(ctx context.Context, walletID string, amount int64) (string, error)
	Add(ctx context.Context, walletID string, amount int64) (string, error)
}

type walletService struct {
	walletRepo repos.WalletRepo
}

func NewWalletService(walletRepo repos.WalletRepo) WalletService {
	return &walletService{
		walletRepo: walletRepo,
	}
}
