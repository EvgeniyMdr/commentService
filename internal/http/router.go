package http

import (
	"github.com/gorilla/mux"
)

func SetupRouter(service services.Service) *mux.Router {
	h := handlers.Handlers{}

	r := mux.NewRouter()
	return r
}
