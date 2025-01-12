package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"movie-catalog-go/internal/app"
)

// Mock do DatabaseChecker
type MockDatabaseChecker struct {
	ShouldFail bool
}

func (m *MockDatabaseChecker) Ping(ctx context.Context) error {
	if m.ShouldFail {
		return errors.New("database connection error")
	}
	return nil
}

func TestHealthHandler_HandleHealthDB(t *testing.T) {
	tests := []struct {
		name           string
		mockShouldFail bool
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Database connection is healthy",
			mockShouldFail: false,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"ok"}`,
		},
		{
			name:           "Database connection fails",
			mockShouldFail: true,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"database connection error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock do serviço de saúde
			mockDB := &MockDatabaseChecker{ShouldFail: tt.mockShouldFail}
			healthService := app.NewHealthService(mockDB)
			handler := NewHealthHandler(healthService)

			// Criar uma requisição simulada
			req, err := http.NewRequest(http.MethodGet, "/health-db", nil)
			if err != nil {
				t.Fatalf("Não foi possível criar a requisição: %v", err)
			}

			// Criar um ResponseRecorder para capturar a resposta
			rr := httptest.NewRecorder()

			// Chamar o manipulador diretamente
			handler.HandleHealthDB(rr, req)

			// Verificar o status da resposta
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Código de status incorreto: obtido %v, esperado %v", status, tt.expectedStatus)
			}

			// Verificar o corpo da resposta
			actualBody := rr.Body.String()
			if actualBody != tt.expectedBody+"\n" {
				t.Errorf("Corpo da resposta incorreto: obtido %v, esperado %v", actualBody, tt.expectedBody)
			}
		})
	}
}
