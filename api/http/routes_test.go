package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock para a interface HealthPort
type MockHealthService struct{}

func (m *MockHealthService) CheckHealth() (string, error) {
	return "ok", nil
}

func TestRegisterRoutes(t *testing.T) {
	mux := http.NewServeMux()
	mockHealthService := &MockHealthService{}

	// Registrar rotas
	RegisterRoutes(mux, mockHealthService)

	tests := []struct {
		route   string
		wantErr bool
	}{
		{route: "/health", wantErr: false},
		{route: "/health-db", wantErr: false},
		{route: "/invalid", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.route, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.route, nil)
			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			if tt.wantErr && rr.Code != http.StatusNotFound {
				t.Errorf("esperado 404 para %s, mas obteve %d", tt.route, rr.Code)
			}

			if !tt.wantErr && rr.Code != http.StatusOK {
				t.Errorf("esperado 200 para %s, mas obteve %d", tt.route, rr.Code)
			}
		})
	}
}
