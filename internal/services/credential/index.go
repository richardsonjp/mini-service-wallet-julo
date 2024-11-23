package credential

import (
	"context"
	repos "go-skeleton/internal/repositories"
)

type CredentialService interface {
	CreateCredential(ctx context.Context, customerID string) (*CredentialResponse, error)
	GetCredential(ctx context.Context, customerID string) (*CredentialResponse, error)
	AuthenticateToken(ctx context.Context, token string) (string, error)
}

type credentialService struct {
	credentialRepo repos.CredentialRepo
}

func NewCredentialService(credentialRepo repos.CredentialRepo) CredentialService {
	return &credentialService{
		credentialRepo: credentialRepo,
	}
}
