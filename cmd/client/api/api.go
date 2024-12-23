package api

import (
	"net/http"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/handlers"
	"github.com/eduardobcolombo/learning-grpc/cmd/client/mid"
	"github.com/gorilla/mux"
)

type API struct {
	Router *mux.Router
}

// New creates a new API struct.
func New() *API {
	return &API{
		Router: mux.NewRouter(),
	}
}

// Routes write the routes in the Router
func (a *API) Routes(handlerPort handlers.Handler) {
	ver := "/v1"

	a.Router.HandleFunc(ver+"/ports", handlerPort.Retrieve).Methods(http.MethodGet, http.MethodOptions)
	a.Router.HandleFunc(ver+"/ports", handlerPort.Update).Methods(http.MethodPost, http.MethodOptions)
	// a.Router.HandleFunc(ver+"/ports", handlerPort.Create).Methods(http.MethodPost, http.MethodOptions)
	// a.Router.HandleFunc(ver+"/ports/{id}", handlerPort.RetrievePort).Methods(http.MethodGet, http.MethodOptions)
	// a.Router.HandleFunc(ver+"/ports/{id}", handlerPort.UpdatePort).Methods(http.MethodPut, http.MethodOptions)
	// a.Router.HandleFunc(ver+"/ports/{id}", handlerPort.DeletePort).Methods(http.MethodDelete, http.MethodOptions)

	a.Router.HandleFunc("/liveness", Liveness)
	a.Router.HandleFunc("/readiness", Readiness)
}

// Mid set all middlewares
func (a *API) Mid(cfg *Config) {
	a.Router.Use(mux.CORSMethodMiddleware(a.Router))
	a.Router.Use(mid.Authenticate)
}
