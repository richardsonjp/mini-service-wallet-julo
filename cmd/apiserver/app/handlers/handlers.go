package handler

import (
	"go-skeleton/cmd/apiserver/app/handlers/customer"
	"go-skeleton/cmd/apiserver/app/handlers/transaction"
	"go-skeleton/cmd/apiserver/app/handlers/wallet"
)

// put handlers alias
type (
	CustomerHandler    = customer.CustomerHandler
	WalletHandler      = wallet.WalletHandler
	TransactionHandler = transaction.TransactionHandler
)

var (
	NewCustomerHandler    = customer.NewCustomerHandler
	NewWalletHandler      = wallet.NewWalletHandler
	NewTransactionHandler = transaction.NewTransactionHandler
)
