package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/api"
)

// TestLiveness checks if the liveness endpoint returns the correct response
func TestLiveness(t *testing.T) {
	req, err := http.NewRequest("GET", "/liveness", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.Liveness)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	expected := `"liveness"`
	if rr.Body.String() != expected {
		t.Errorf("expected body %v, got %v", expected, rr.Body.String())
	}
}

// TestReadiness checks if the readiness endpoint returns the correct response
func TestReadiness(t *testing.T) {
	req, err := http.NewRequest("GET", "/readiness", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.Readiness)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	expected := `"readiness"`
	if rr.Body.String() != expected {
		t.Errorf("expected body %v, got %v", expected, rr.Body.String())
	}
}
