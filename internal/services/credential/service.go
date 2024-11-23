package credential

import (
	"context"
	"go-skeleton/internal/model"
)

func (s *credentialService) CreateCredential(ctx context.Context, customerID string) (*CredentialResponse, error) {
	currentCredential, _ := s.GetCredential(ctx, customerID)
	if currentCredential != nil {
		return currentCredential, nil
	}

	credential, err := s.credentialRepo.Create(ctx, &model.Credential{
		CustomerID: customerID,
	})
	if err != nil {
		return nil, err
	}

	return &CredentialResponse{
		Token: credential.ID,
	}, nil
}

func (s *credentialService) GetCredential(ctx context.Context, customerID string) (*CredentialResponse, error) {
	credential, err := s.credentialRepo.GetOne(ctx, &model.Credential{CustomerID: customerID})
	if err != nil {
		return nil, err
	}

	return &CredentialResponse{
		Token: credential.ID,
	}, nil
}

func (s *credentialService) AuthenticateToken(ctx context.Context, token string) (string, error) {
	data, err := s.credentialRepo.GetOne(ctx, &model.Credential{ID: token})
	if err != nil {
		return "", err
	}

	return data.CustomerID, nil
}
