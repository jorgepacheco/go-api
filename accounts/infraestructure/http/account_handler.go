package http

import (
	"github.com/go-chi/chi"
	"github.com/go-first-steps/accounts/app"
	"net/http"
)

type AccountHandler struct {
	Service app.AccountService
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

// Routes returns post router with each endpoint.
func (srv *AccountHandler) Routes() http.Handler {

	r := chi.NewRouter()

	r.Get("/", srv.GetAll)

	return r
}
