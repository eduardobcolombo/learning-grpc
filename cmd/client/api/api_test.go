package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock_core_port "github.com/eduardobcolombo/learning-grpc/cmd/client/core/port/mock_port"
	"github.com/eduardobcolombo/learning-grpc/cmd/client/handlers"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	cfg := &Config{}
	log, _ := foundation.NewLogger(&foundation.LoggerConfig{})

	ctrl := gomock.NewController(t)
	coreService := mock_core_port.NewMockCoreService(ctrl)
	handlerPort := handlers.NewHandler(coreService, log)
	api := New()
	api.Routes(handlerPort)
	api.Mid(cfg)

	tests := []struct {
		method        string
		url           string
		expectedCalls func()
	}{
		{method: http.MethodGet, url: "/v1/ports", expectedCalls: func() {
			coreService.EXPECT().Retrieve().AnyTimes()
		}},
		{method: http.MethodPost, url: "/v1/ports", expectedCalls: func() {}},
		{method: http.MethodOptions, url: "/v1/ports", expectedCalls: func() {}},
		{method: http.MethodGet, url: "/liveness", expectedCalls: func() {}},
		{method: http.MethodGet, url: "/readiness", expectedCalls: func() {}},
	}

	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, tt.url, nil)
		assert.NoError(t, err)

		tt.expectedCalls()

		rr := httptest.NewRecorder()
		api.Router.ServeHTTP(rr, req)

		assert.NotEqual(t, http.StatusNotFound, rr.Code, "Expected route to be registered")
	}
}
