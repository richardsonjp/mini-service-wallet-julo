package transaction

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/internal/model/enum"
	timeutil "go-skeleton/pkg/utils/time"
)

func (s *transactionService) CreateTransaction(ctx context.Context, payload CreateTransactionPayload) (interface{}, error) {
	var tData *model.Transaction
	errTx := s.tx.Run(ctx, func(ctx context.Context) error {
		walletData, err := s.walletService.GetWalletBalance(ctx, payload.CustomerID)
		if err != nil {
			return err
		}
		var tType model.TransactionType
		var tStatus string
		if payload.TransactionType == enum.WITHDRAWAL.String() {
			tType = enum.WITHDRAWAL
			tStatus, err = s.walletService.Deduct(ctx, walletData.ID, payload.Amount)
			if err != nil {
				return err
			}
		} else if payload.TransactionType == enum.DEPOSIT.String() {
			tType = enum.DEPOSIT
			tStatus, err = s.walletService.Add(ctx, walletData.ID, payload.Amount)
			if err != nil {
				return err
			}
		}

		tData, err = s.transactionRepo.Create(ctx, &model.Transaction{
			WalletID:    walletData.ID,
			Amount:      payload.Amount,
			Status:      enum.NewGenericStatus(tStatus),
			Type:        tType,
			ReferenceID: payload.ReferenceID,
		})
		return err
	})
	if errTx != nil {
		return nil, errTx
	}

	var result interface{}
	if payload.TransactionType == enum.WITHDRAWAL.String() {
		result = WithdrawalResponse{
			ID:          tData.ID,
			WithdrawBy:  tData.WalletID,
			Status:      tData.Status.String(),
			WithdrawAt:  timeutil.StrFormat(tData.CreatedAt),
			Amount:      tData.Amount,
			ReferenceID: tData.ReferenceID,
		}
	} else if payload.TransactionType == enum.DEPOSIT.String() {
		result = DepositResponse{
			ID:          tData.ID,
			DepositedBy: tData.WalletID,
			Status:      tData.Status.String(),
			DepositedAt: timeutil.StrFormat(tData.CreatedAt),
			Amount:      tData.Amount,
			ReferenceID: tData.ReferenceID,
		}
	}

	return result, nil
}

func (s *transactionService) ListTransactions(ctx context.Context, customerID string) ([]ListTransactionsResponse, error) {
	walletData, err := s.walletService.GetWalletBalance(ctx, customerID)
	if err != nil {
		return nil, err
	}

	list, err := s.transactionRepo.ListTransaction(ctx, walletData.ID)
	if err != nil {
		return nil, err
	}

	result := []ListTransactionsResponse{}
	for _, item := range list {
		result = append(result, ListTransactionsResponse{
			ID:           item.ID,
			Status:       item.Status.String(),
			TransactedAt: timeutil.StrFormat(item.CreatedAt),
			Type:         item.Type.String(),
			Amount:       item.Amount,
			ReferenceID:  item.ReferenceID,
		})
	}

	return result, nil
}
