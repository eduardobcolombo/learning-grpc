package api

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func CORS() func(http.Handler) http.Handler {

	allowedHeaders := []string{"X-Requested-With", "Authorization", "Content-Type", "X-Total-Count", "Location"}
	allowedOrigins := []string{"*"}
	allowedMethods := []string{"POST", "GET", "OPTIONS"}
	exposedHeaders := []string{"X-Requested-With", "Authorization", "Content-Type", "X-Total-Count", "Location"}

	return handlers.CORS(handlers.AllowedOrigins(allowedOrigins), handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowedMethods(allowedMethods), handlers.ExposedHeaders(exposedHeaders))
}

// TODO: implement some auth
func (e *Environment) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
