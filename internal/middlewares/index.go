package middlewares

import "go-skeleton/internal/services/credential"

type MiddlewareAccess struct {
	credentialService credential.CredentialService
}

func NewMiddlewareAccess(credentialService credential.CredentialService) MiddlewareAccess {
	return MiddlewareAccess{
		credentialService: credentialService,
	}
}
