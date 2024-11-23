package store

import (
	"go-skeleton/cmd/apiserver/app/handlers"
	"go-skeleton/config"
	"go-skeleton/internal/middlewares"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/services"
	"go-skeleton/pkg/clients/db"
)

var (
	// handlers
	CustomerHandler    *handler.CustomerHandler
	WalletHandler      *handler.WalletHandler
	TransactionHandler *handler.TransactionHandler

	// services
	CustomerService    service.CustomerService
	CredentialService  service.CredentialService
	WalletService      service.WalletService
	TransactionService service.TransactionService

	// repos
	CustomerRepo    repos.CustomerRepo
	CredentialRepo  repos.CredentialRepo
	WalletRepo      repos.WalletRepo
	TransactionRepo repos.TransactionRepo

	TxRepo repos.TxRepo

	//middleware
	MiddlewareAccess middlewares.MiddlewareAccess
)

// Init application global variable with single instance
func InitDI() {
	// setup resources
	dbdget := db.NewDBdelegate(config.Config.DB.Debug)
	dbdget.Init()

	// repos
	CustomerRepo = repos.NewCustomerRepo(dbdget)
	CredentialRepo = repos.NewCredentialRepo(dbdget)
	WalletRepo = repos.NewWalletRepo(dbdget)
	TransactionRepo = repos.NewTransactionRepo(dbdget)
	TxRepo = repos.NewTxRepo(dbdget)

	// servicesWalletRepo
	CredentialService = service.NewCredentialService(CredentialRepo)
	WalletService = service.NewWalletService(WalletRepo)
	TransactionService = service.NewTransactionService(TransactionRepo, WalletService, TxRepo)
	CustomerService = service.NewCustomerService(CustomerRepo, CredentialService, WalletService, TxRepo)

	// handlers
	CustomerHandler = handler.NewCustomerHandler(CustomerService)
	WalletHandler = handler.NewWalletHandler(WalletService)
	TransactionHandler = handler.NewTransactionHandler(TransactionService)

	// middlewares
	MiddlewareAccess = middlewares.NewMiddlewareAccess(CredentialService)
}
