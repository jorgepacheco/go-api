package app

import (
	"context"
	"github.com/go-first-steps/accounts/domain"
)

type AccountService struct {
	Repo domain.Repository
}

func (srv *AccountService) GetAll(ctx context.Context) ([]domain.Account, error) {

	return srv.Repo.GetAll(ctx)
}

func (srv *AccountService) Create(account domain.Account) (domain.Account, error) {

	return srv.Repo.Save(account)
}
