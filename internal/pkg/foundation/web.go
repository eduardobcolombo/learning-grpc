package foundation

import (
	"encoding/json"
	"net/http"
)

// Response will handle the http responses
func Response(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			return
		}

	}
	// TODO: check the response format
	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write([]byte(response)); err != nil {
		return
	}

}
