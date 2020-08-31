package data

import (
	"context"
	"github.com/go-first-steps/accounts"
)

type AccountRepository struct {
	Data *Data
}

func (ur *AccountRepository) GetAll() ([]accounts.Account, error) {
	q := `
	SELECT id, balance, iban
		FROM accounts;
	`

	rows, err := ur.Data.DB.QueryContext(context.Background(), q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []accounts.Account
	for rows.Next() {
		var u accounts.Account
		rows.Scan(&u.Id, &u.Balance, &u.Iban)
		users = append(users, u)
	}

	return users, nil
}
