package api

import (
	"net/http"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/mid"
	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
}

func newAPI() *API {
	return &API{
		router: mux.NewRouter(),
	}
}

func (api *API) Handler() http.Handler {
	return api.router
}

func (api *API) routes(core CoreConfig, cfg *Config) {
	ver := "/v1"

	portHandler := Handler{
		core: core.Port,
	}

	api.router.HandleFunc(ver+"/ports", portHandler.RetrievePorts).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc(ver+"/ports", portHandler.UpdatePorts).Methods(http.MethodPost, http.MethodOptions)

	api.router.HandleFunc("/debug/liveness", liveness)
	api.router.HandleFunc("/debug/readiness", readiness)
}

func (api *API) mid(cfg *Config) {
	api.router.Use(mux.CORSMethodMiddleware(api.router))
	api.router.Use(mid.CORS())
	api.router.Use(mid.Authenticate)
}
