package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eduardobcolombo/learning-grpc/portpb"
)

type Environment struct {
	psc portpb.PortServiceClient
}

func (e *Environment) Response(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			log.Fatalf("Error trying to write the error response body: %s ", err)
		}
		return
	}
	// TODO: check the response format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write([]byte(response)); err != nil {
		log.Fatalf("Error trying to write the response body: %s ", err)
	}
	return

}
