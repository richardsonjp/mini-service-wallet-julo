package repos

import (
	"go-skeleton/internal/repositories/credential"
	"go-skeleton/internal/repositories/customer"
	"go-skeleton/internal/repositories/transaction"
	"go-skeleton/internal/repositories/tx"
	"go-skeleton/internal/repositories/wallet"
)

// put repos alias
type (
	CustomerRepo    = customer.CustomerRepo
	CredentialRepo  = credential.CredentialRepo
	WalletRepo      = wallet.WalletRepo
	TransactionRepo = transaction.TransactionRepo

	TxRepo = tx.TxRepo
)

var (
	NewCustomerRepo    = customer.NewCustomerRepo
	NewCredentialRepo  = credential.NewCredentialRepo
	NewWalletRepo      = wallet.NewWalletRepo
	NewTransactionRepo = transaction.NewTransactionRepo

	NewTxRepo = tx.NewTxRepo
)
