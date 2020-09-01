package domain

import "github.com/google/uuid"

type Account struct {
	Id      uuid.UUID `json:"id,omitempty"`
	Iban    string    `json:"iban,omitempty"`
	Balance float32   `json:"balance,omitempty"`
}
