package wallet

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/internal/model/enum"
	"go-skeleton/pkg/utils/errors"
	timeutil "go-skeleton/pkg/utils/time"
)

func (s *walletService) CreateWallet(ctx context.Context, customerID string) error {
	err := s.walletRepo.Create(ctx, &model.Wallet{CustomerID: customerID, Status: enum.ENABLED})
	if err != nil {
		return err
	}

	return nil
}

func (s *walletService) ToggleWallet(ctx context.Context, payload ToggleWallet) (*WalletResponse, error) {
	walletData, err := s.walletRepo.GetOne(ctx, &model.Wallet{
		CustomerID: payload.CustomerID,
	})
	if err != nil {
		return nil, err
	}

	if walletData.Status == enum.DISABLED && !payload.Enabled {
		return nil, errors.NewGenericError(errors.WALLET_DISABLED)
	}

	walletData.Status = enum.DISABLED
	if payload.Enabled {
		walletData.Status = enum.ENABLED
	}

	_, err = s.walletRepo.Update(ctx, walletData, "status")
	if err != nil {
		return nil, err
	}

	result, err := s.GetByCustomerID(ctx, payload.CustomerID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *walletService) GetByCustomerID(ctx context.Context, customerID string) (*WalletResponse, error) {
	walletData, err := s.walletRepo.GetOne(ctx, &model.Wallet{
		CustomerID: customerID,
	})
	if err != nil {
		return nil, err
	}

	result := &WalletResponse{
		ID:      walletData.ID,
		OwnedBy: walletData.CustomerID,
		Balance: walletData.Balance,
		Status:  walletData.Status.String(),
	}

	if walletData.Status == enum.ENABLED {
		result.EnabledAt = timeutil.StrFormat(walletData.UpdatedAt)
	} else {
		result.DisabledAt = timeutil.StrFormat(walletData.UpdatedAt)
	}

	return result, err
}

func (s *walletService) GetWalletBalance(ctx context.Context, customerID string) (*WalletResponse, error) {
	walletData, err := s.GetByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if walletData.Status == enum.DISABLED.String() {
		return nil, errors.NewGenericError(errors.WALLET_DISABLED)
	}

	return walletData, err
}

func (s *walletService) Deduct(ctx context.Context, walletID string, amount int64) (string, error) {
	currentWallet, err := s.walletRepo.GetOne(ctx, &model.Wallet{ID: walletID})
	if err != nil {
		return "", err
	}

	if currentWallet.Balance < amount {
		return enum.FAIL.String(), nil
	}

	currentWallet.Balance = currentWallet.Balance - amount
	_, err = s.walletRepo.Update(ctx, currentWallet, "balance")
	if err != nil {
		return "", err
	}
	return enum.SUCCESS.String(), nil
}

func (s *walletService) Add(ctx context.Context, walletID string, amount int64) (string, error) {
	currentWallet, err := s.walletRepo.GetOne(ctx, &model.Wallet{ID: walletID})
	if err != nil {
		return "", err
	}

	currentWallet.Balance = currentWallet.Balance + amount
	_, err = s.walletRepo.Update(ctx, currentWallet, "balance")
	if err != nil {
		return "", err
	}
	return enum.SUCCESS.String(), nil
}
