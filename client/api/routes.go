package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (e *Environment) GetRoutes(router *mux.Router) {
	ver := "/v1"
	router.HandleFunc(ver+"/ports", e.RetrievePortsHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc(ver+"/ports", e.UpdatePortsHandler).Methods(http.MethodPost, http.MethodOptions)
}

func (e *Environment) GetMiddlewares(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(CORS())
	router.Use(e.Auth)
}

func (e *Environment) RetrievePortsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		e.RetrievePorts(w, r)
	}
}
func (e *Environment) UpdatePortsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		e.UpdatePorts(w, r)
	}
}
