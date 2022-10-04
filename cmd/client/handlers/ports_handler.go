package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/core/port"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/foundation"
)

type Handler struct {
	Core port.Core
}

// RetrievePorts will retrieve Ports using Core
func (h Handler) RetrievePorts(w http.ResponseWriter, r *http.Request) {
	ports, err := h.Core.RetrievePortsFromServer()
	if err != nil {
		foundation.Response(w, http.StatusInternalServerError, err)
	}

	foundation.Response(w, http.StatusOK, ports)
}

// UpdatePorts will update Ports using Core
func (h Handler) UpdatePorts(w http.ResponseWriter, r *http.Request) {
	// fileName := "data/ports.json"
	fileName, err := upload(w, r)
	if err != nil {
		foundation.Response(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("Importing the file: %s\n\n", fileName)
	msg, err := h.Core.UpdatePortsOnServer(fileName)
	if err != nil {
		foundation.Response(w, http.StatusInternalServerError, err)
		return
	}

	os.Remove(fileName)

	foundation.Response(w, http.StatusOK, msg)
}

// upload will do the file upload when received from request
func upload(w http.ResponseWriter, r *http.Request) (string, error) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 30 specifies a maximum
	// upload of 10 GB files.
	r.ParseMultipartForm(10 << 30)
	file, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("./", "upload-*.json")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	tempFile.Write(fileBytes)

	return tempFile.Name(), nil
}
