package http

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-first-steps/accounts/app"
	"github.com/go-first-steps/accounts/domain"
	"github.com/go-first-steps/internal/logs"
	"net/http"
)

type AccountHandler struct {
	Service app.AccountService
}

type AccountRequest struct {
	Iban string `json:"iban"`
}

// GetAllHandler response all the users.
func (srv *AccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := srv.Service.GetAll(ctx)
	if err != nil {
		HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	JSON(w, r, http.StatusOK, Map{"accounts": users})
}

// GetAllHandler response all the users.
func (srv *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {

	logs.Log().Info("Create Account...")

	accountRequest := parseRequest(r)

	logs.Log().Info("Create Account... " + accountRequest.Iban)

	account := domain.Account{
		Iban: accountRequest.Iban,
	}

	accountCreated, err := srv.Service.Create(account)
	if err != nil {
		HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	JSON(w, r, http.StatusCreated, accountCreated)
}

func parseRequest(r *http.Request) AccountRequest {
	body := r.Body

	defer body.Close()

	var accountRequest AccountRequest

	_ = json.NewDecoder(body).Decode(&accountRequest)

	return accountRequest
}

// Routes returns post router with each endpoint.
func (srv *AccountHandler) Routes() http.Handler {

	r := chi.NewRouter()

	r.Get("/", srv.GetAll)
	r.Post("/account", srv.Create)

	return r
}
