package handlers

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	mock_core_port "github.com/eduardobcolombo/learning-grpc/cmd/client/core/port/mock_port"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/portpb"
	"github.com/golang/mock/gomock"
)

func TestRetrievePorts(t *testing.T) {
	tests := []struct {
		name           string
		mockRetrieve   func() ([]*portpb.Port, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful retrieval",
			mockRetrieve: func() ([]*portpb.Port, error) {
				return []*portpb.Port{
					{Name: "Port1"},
					{Name: "Port2"},
				}, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"name":"Port1"},{"name":"Port2"}]`,
		},
		{
			name: "internal server error",
			mockRetrieve: func() ([]*portpb.Port, error) {
				return nil, fmt.Errorf("internal error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   http.StatusText(http.StatusInternalServerError),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCoreService := mock_core_port.NewMockCoreService(ctrl)
			testLog, _ := foundation.NewLogger(&foundation.LoggerConfig{})
			handler := NewHandler(mockCoreService, testLog)

			mockCoreService.EXPECT().Retrieve().Return(tt.mockRetrieve())

			req, err := http.NewRequest("GET", "/ports", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.Retrieve(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
func TestUpload(t *testing.T) {
	tests := []struct {
		name           string
		fileContent    string
		expectedError  bool
		expectedPrefix string
	}{
		{
			name:           "successful upload",
			fileContent:    `{"name":"Port1"}`,
			expectedError:  false,
			expectedPrefix: "upload-",
		},
		{
			name:          "no file in request",
			fileContent:   "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			if tt.fileContent != "" {
				part, err := writer.CreateFormFile("file", "test.json")
				if err != nil {
					t.Fatal(err)
				}
				_, err = part.Write([]byte(tt.fileContent))
				if err != nil {
					t.Fatal(err)
				}
			}
			writer.Close()

			req, err := http.NewRequest("POST", "/upload", body)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())

			fileName, err := upload(req)
			if (err != nil) != tt.expectedError {
				t.Errorf("upload() error = %v, expectedError %v", err, tt.expectedError)
				return
			}

			if !tt.expectedError {
				if !strings.HasPrefix(filepath.Base(fileName), tt.expectedPrefix) {
					t.Errorf("uploaded file name = %v, expected prefix %v", filepath.Base(fileName), tt.expectedPrefix)
				}
				err = os.Remove(fileName)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
func TestUpdatePorts(t *testing.T) {
	tests := []struct {
		name           string
		fileContent    string
		setup          func(*mock_core_port.MockCoreService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "successful update",
			fileContent: `{"name":"Port1"}`,
			setup: func(mockCoreService *mock_core_port.MockCoreService) {
				mockCoreService.EXPECT().Update(gomock.Any()).Return("update successful", nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"update successful"`,
		},
		{
			name:        "return error calling update",
			fileContent: `{"xpto""INVALID"}`,
			setup: func(mockCoreService *mock_core_port.MockCoreService) {
				mockCoreService.EXPECT().Update(gomock.Any()).Return("", fmt.Errorf("error..."))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   http.StatusText(http.StatusInternalServerError),
		},
		{
			name:           "internal server error on upload",
			fileContent:    "",
			setup:          func(mockCoreService *mock_core_port.MockCoreService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrBadRequestFileMustBeSent.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCoreService := mock_core_port.NewMockCoreService(ctrl)
			testLog, _ := foundation.NewLogger(&foundation.LoggerConfig{})
			handler := NewHandler(mockCoreService, testLog)

			tt.setup(mockCoreService)

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			if tt.fileContent != "" {
				part, err := writer.CreateFormFile("file", "test.json")
				if err != nil {
					t.Fatal(err)
				}
				_, err = part.Write([]byte(tt.fileContent))
				if err != nil {
					t.Fatal(err)
				}
			}
			writer.Close()

			req, err := http.NewRequest("POST", "/update", body)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())

			rr := httptest.NewRecorder()
			handler.Update(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if strings.TrimSpace(rr.Body.String()) != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
