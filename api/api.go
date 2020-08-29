package api

import "net/http"

type Services struct {
}

func (s *Services) healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("OK from Container"))
}

func Start(port string) {
	services := &Services{}

	r := routes(services)
	server := newServer(port, r)
	server.Start()

}
