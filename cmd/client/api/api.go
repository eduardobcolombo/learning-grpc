package api

import (
	"net/http"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/handlers"
	"github.com/eduardobcolombo/learning-grpc/cmd/client/mid"
	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
}

// New creates a new API struct
func New() *API {
	return &API{
		router: mux.NewRouter(),
	}
}

// Handler will return the router
func (a *API) Handler() http.Handler {
	return a.router
}

// Routes write the routes in the router
func (a *API) Routes(coreCfg CoreConfig, cfg *Config) {
	ver := "/v1"

	portHandler := handlers.Handler{
		Core: coreCfg.Port,
	}

	a.router.HandleFunc(ver+"/ports", portHandler.RetrievePorts).Methods(http.MethodGet, http.MethodOptions)
	a.router.HandleFunc(ver+"/ports", portHandler.UpdatePorts).Methods(http.MethodPost, http.MethodOptions)

	a.router.HandleFunc("/debug/liveness", liveness)
	a.router.HandleFunc("/debug/readiness", readiness)
}

// Mid set all middlewares
func (a *API) Mid(cfg *Config) {
	a.router.Use(mux.CORSMethodMiddleware(a.router))
	a.router.Use(mid.CORS())
	a.router.Use(mid.Authenticate)
}
