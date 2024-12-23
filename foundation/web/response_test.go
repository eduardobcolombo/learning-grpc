package web

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSuccess(t *testing.T) {
	tests := []struct {
		name       string
		status     int
		payload    interface{}
		wantStatus int
		wantBody   string
	}{
		{
			name:       "should return success with a json payload",
			status:     http.StatusOK,
			payload:    map[string]string{"message": "ok"},
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"ok"}`,
		},
		{
			name:       "should return no content with a nil payload",
			status:     http.StatusNoContent,
			payload:    nil,
			wantStatus: http.StatusNoContent,
			wantBody:   "",
		},
		{
			name:       "should return 500 with invalid payload",
			status:     http.StatusInternalServerError,
			payload:    make(chan int),
			wantStatus: http.StatusInternalServerError,
			wantBody:   "json: unsupported type: chan int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ResponseSuccess(w, tt.status, tt.payload)
			})

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("expected status code %v, got %v", tt.wantStatus, rr.Code)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("expected body %v, got %v", tt.wantBody, rr.Body.String())
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name       string
		status     int
		err        error
		wantStatus int
		wantBody   string
	}{
		{
			name:       "should return bad request with a valid error",
			status:     http.StatusBadRequest,
			err:        errors.New("bad request"),
			wantStatus: http.StatusBadRequest,
			wantBody:   "bad request",
		},
		{
			name:       "should return forbidden with a valid error",
			status:     http.StatusForbidden,
			err:        errors.New("forbidden"),
			wantStatus: http.StatusForbidden,
			wantBody:   "forbidden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ResponseError(w, tt.status, tt.err)
			})

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("expected status code %v, got %v", tt.wantStatus, rr.Code)
			}

			if strings.TrimSpace(rr.Body.String()) != tt.wantBody {
				t.Errorf("expected body %v, got %v", tt.wantBody, rr.Body.String())
			}
		})
	}
}
