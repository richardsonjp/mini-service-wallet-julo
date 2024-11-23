package transaction

import (
	"context"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/services/wallet"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, payload CreateTransactionPayload) (interface{}, error)
	ListTransactions(ctx context.Context, customerID string) ([]ListTransactionsResponse, error)
}

type transactionService struct {
	transactionRepo repos.TransactionRepo
	walletService   wallet.WalletService
	tx              repos.TxRepo
}

func NewTransactionService(transactionRepo repos.TransactionRepo, walletService wallet.WalletService, tx repos.TxRepo) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		walletService:   walletService,
		tx:              tx,
	}
}
