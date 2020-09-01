package domain

type Account struct {
	Id      int     `json:"id,omitempty"`
	Iban    string  `json:"iban,omitempty"`
	Balance float32 `json:"balance,omitempty"`
}
