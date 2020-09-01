package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-first-steps/accounts/app"
	http2 "github.com/go-first-steps/accounts/infraestructure/http"
	"github.com/go-first-steps/accounts/infraestructure/persistence"
	"github.com/go-first-steps/internal/data"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("OK from Container"))
}

func Start(port string) {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)

	r.Get("/heartbeat", healthCheck)

	accountController := accountController()

	r.Mount("/accounts", accountController.Routes())

	server := newServer(port, r)

	server.Start()

}

func accountController() http2.AccountHandler {

	repo := persistence.AccountRepository{Data: data.New()}

	accountService := app.AccountService{Repo: &repo}

	return http2.AccountHandler{Service: accountService}

}
