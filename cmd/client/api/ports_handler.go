package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	// fileName := "data/ports.json"
	fileName, err := upload(w, r)
	if err != nil {
		foundation.Response(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("Importing the file: %s\n\n", fileName)
	msg, err := h.core.UpdatePortsOnServer(fileName)
	if err != nil {
		foundation.Response(w, http.StatusInternalServerError, err)
		return
	}

	os.Remove(fileName)

	foundation.Response(w, http.StatusOK, msg)
}

func upload(w http.ResponseWriter, r *http.Request) (string, error) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
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
