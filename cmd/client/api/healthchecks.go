package api

import "net/http"

// liveness provide health check
func liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("liveness"))
}

// readiness provide health check
func readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("readiness"))
}
