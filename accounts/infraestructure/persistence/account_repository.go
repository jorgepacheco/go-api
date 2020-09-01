package persistence

import (
	"context"
	"github.com/go-first-steps/accounts/domain"
	"github.com/go-first-steps/internal/data"
	"github.com/go-first-steps/internal/logs"
	"github.com/google/uuid"
)

type AccountRepository struct {
	Data *data.Data
}

func (ur *AccountRepository) GetAll(ctx context.Context) ([]domain.Account, error) {
	q := `
	SELECT id, balance, iban
		FROM accounts;
	`

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []domain.Account
	for rows.Next() {
		var u domain.Account
		rows.Scan(&u.Id, &u.Balance, &u.Iban)
		users = append(users, u)
	}

	return users, nil
}

func (ur *AccountRepository) Save(account domain.Account) (domain.Account, error) {

	uuid := uuid.New()

	newAccount := domain.Account{
		Id:      uuid,
		Iban:    account.Iban,
		Balance: 0,
	}

	tx, err := ur.Data.DB.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return account, err
	}

	logs.Log().Info("::: INSERT ACCOUNT {0}")

	q := `
	insert into accounts (id, iban, balance) values ($1, $2, $3);
	`

	_, err = tx.Exec(q, uuid, account.Iban, 0.0)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return account, err
	}

	_ = tx.Commit()

	return newAccount, nil
}
