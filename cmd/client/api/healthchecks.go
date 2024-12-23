package api

import (
	"net/http"

	"github.com/eduardobcolombo/learning-grpc/foundation/web"
)

// Liveness provide health check
func Liveness(w http.ResponseWriter, r *http.Request) {
	web.ResponseSuccess(w, http.StatusOK, "liveness")
}

// Readiness provide health check
func Readiness(w http.ResponseWriter, r *http.Request) {
	web.ResponseSuccess(w, http.StatusOK, "readiness")
}
