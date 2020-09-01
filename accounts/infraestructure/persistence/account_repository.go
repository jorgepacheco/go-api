package persistence

import (
	"context"
	"github.com/go-first-steps/accounts/domain"
	"github.com/go-first-steps/internal/data"
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
