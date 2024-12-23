package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	core_port "github.com/eduardobcolombo/learning-grpc/cmd/client/core/port"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/foundation/web"
)

var (
	ErrInternalServerError      = fmt.Errorf("internal server error")
	ErrBadRequestFileMustBeSent = fmt.Errorf("file must be sent")
)

type Handler struct {
	PortCore core_port.CoreService
	log      *foundation.Logger
}

// NewHandler will create a new Handler.
func NewHandler(pc core_port.CoreService, log *foundation.Logger) Handler {
	return Handler{PortCore: pc, log: log}
}

// Create will create a new Port using Core.
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {

	var port PortRequest
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&port); err != nil {
	// 	web.ResponseError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// port, err := h.PortCore.Create(port.toPortCore())
	// if err != nil {
	// 	web.ResponseError(w, http.StatusInternalServerError, err)
	// }

	web.ResponseSuccess(w, http.StatusOK, port)
}

// Retrieve will retrieve Ports using Core.
func (h Handler) Retrieve(w http.ResponseWriter, r *http.Request) {
	ports, err := h.PortCore.Retrieve()
	if err != nil {
		log.Printf("error while calling Retrieve: %v", err)
		web.ResponseError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("%s", http.StatusText(http.StatusInternalServerError)),
		)
		return
	}

	web.ResponseSuccess(w, http.StatusOK, ports)
}

// Update will update Ports using Core.
func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	// upload the file
	fileName, err := upload(r)
	if err != nil {
		log.Printf("error while calling Update in the upload: %v", err)
		web.ResponseError(w, http.StatusBadRequest, ErrBadRequestFileMustBeSent)
		return
	}
	// update the ports
	msg, err := h.PortCore.Update(fileName)
	if err != nil {
		log.Printf("error while calling Update Core: %v", err)
		web.ResponseError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("%s", http.StatusText(http.StatusInternalServerError)),
		)
		os.Remove(fileName)

		return
	}

	os.Remove(fileName)

	web.ResponseSuccess(w, http.StatusOK, msg)
}

// upload will do the file upload when received from request
func upload(r *http.Request) (string, error) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 30 specifies a maximum
	// upload of 10 GB files.
	err := r.ParseMultipartForm(10 << 30)
	if err != nil {
		return "", err

	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("./", "upload-*.json")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if _, err := tempFile.Write(fileBytes); err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

// portRequest is the struct to receive the request from the client.
type PortRequest struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
