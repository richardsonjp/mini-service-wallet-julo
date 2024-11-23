package customer

import (
	"context"
	"go-skeleton/internal/model"
)

func (s *customerService) InitCustomer(ctx context.Context, payload InitCustomerPayload) (*InitCustomerResponse, error) {
	var token string
	errTx := s.tx.Run(ctx, func(ctx context.Context) error {
		customerResp, err := s.customerRepo.GetOne(ctx, &model.Customer{XID: payload.CustomerXID})
		if customerResp == nil {
			customerResp, err = s.customerRepo.Create(ctx, &model.Customer{
				XID: payload.CustomerXID,
			})
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}

		credential, err := s.credentialService.CreateCredential(ctx, customerResp.ID)
		if err != nil {
			return err
		}
		token = credential.Token

		err = s.WalletService.CreateWallet(ctx, customerResp.ID)
		return err
	})

	if errTx != nil {
		return nil, errTx
	}

	return &InitCustomerResponse{Token: token}, nil
}
