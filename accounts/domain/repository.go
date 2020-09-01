package domain

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Account, error)
	Save(account Account) (Account, error)
}
