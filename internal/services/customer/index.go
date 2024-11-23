package customer

import (
	"context"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/services/credential"
	"go-skeleton/internal/services/wallet"
)

type CustomerService interface {
	InitCustomer(ctx context.Context, idx InitCustomerPayload) (*InitCustomerResponse, error)
}

type customerService struct {
	customerRepo      repos.CustomerRepo
	credentialService credential.CredentialService
	WalletService     wallet.WalletService
	tx                repos.TxRepo
}

func NewCustomerService(customerRepo repos.CustomerRepo, credentialService credential.CredentialService, WalletService wallet.WalletService, tx repos.TxRepo) CustomerService {
	return &customerService{
		customerRepo:      customerRepo,
		credentialService: credentialService,
		WalletService:     WalletService,
		tx:                tx,
	}
}
