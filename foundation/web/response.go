package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response will handle the http responses
func ResponseSuccess(w http.ResponseWriter, status int, payload interface{}) {
	var response []byte
	if payload != nil {
		byteResponse, err := json.Marshal(payload)
		if err != nil {
			ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response = byteResponse
		if _, err := w.Write(response); err != nil {
			return
		}
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}

// Error will handle the http responses
func ResponseError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", err.Error())
}
