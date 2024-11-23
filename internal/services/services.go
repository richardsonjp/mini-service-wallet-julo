package service

import (
	"go-skeleton/internal/services/credential"
	"go-skeleton/internal/services/customer"
	"go-skeleton/internal/services/transaction"
	"go-skeleton/internal/services/wallet"
)

// put handlers alias
type (
	CustomerService    = customer.CustomerService
	CredentialService  = credential.CredentialService
	WalletService      = wallet.WalletService
	TransactionService = transaction.TransactionService
)

var (
	NewCustomerService    = customer.NewCustomerService
	NewCredentialService  = credential.NewCredentialService
	NewWalletService      = wallet.NewWalletService
	NewTransactionService = transaction.NewTransactionService
)
