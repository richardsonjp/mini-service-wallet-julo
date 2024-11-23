package customer

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type CustomerRepo interface {
	Create(ctx context.Context, m *model.Customer) (*model.Customer, error)
	GetOne(ctx context.Context, m *model.Customer) (*model.Customer, error)
	Update(ctx context.Context, param *model.Customer, updatedFields ...string) (int64, error)
}

type customerRepo struct {
	dbdget db.DBGormDelegate
}

func NewCustomerRepo(dbdget db.DBGormDelegate) CustomerRepo {
	return &customerRepo{
		dbdget: dbdget,
	}
}
