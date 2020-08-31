package accounts

type Repository interface {
	GetAll() ([]Account, error)
}
