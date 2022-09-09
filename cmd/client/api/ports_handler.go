package api

import (
	"fmt"
	"net/http"

	"github.com/eduardobcolombo/learning-grpc/internal/pkg/foundation"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/port"
)

type Handler struct {
	core port.Core
}

func (h Handler) RetrievePorts(w http.ResponseWriter, r *http.Request) {
	ports, err := h.core.RetrievePortsFromServer()
	if err != nil {
		foundation.Response(w, http.StatusInternalServerError, err)
	}

	foundation.Response(w, http.StatusOK, ports)
}

func (h Handler) UpdatePorts(w http.ResponseWriter, r *http.Request) {
	// TODO: Read the file from upload or URL if it is the case
	fileName := "data/ports.json"
	fmt.Printf("Importing the file: %s\n\n", fileName)
	msg, err := h.core.UpdatePortsOnServer(fileName)
	if err != nil {
		foundation.Response(w, http.StatusInternalServerError, err)
		return
	}

	foundation.Response(w, http.StatusOK, msg)
}
