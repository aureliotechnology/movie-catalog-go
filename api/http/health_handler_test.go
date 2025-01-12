package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"movie-catalog-go/internal/app"
)

func TestHealthHandler(t *testing.T) {
	// Mock do serviço de saúde
	healthService := app.NewHealthService()
	handler := NewHealthHandler(healthService)

	// Criar uma requisição HTTP simulada
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatalf("Não foi possível criar a requisição: %v", err)
	}

	// Criar um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()

	// Chamar o manipulador diretamente
	handler.HandleHealth(rr, req)

	// Verificar o status da resposta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de status incorreto: obtido %v, esperado %v", status, http.StatusOK)
	}

	// Verificar o corpo da resposta
	expected := map[string]string{"message": "ok"}
	var actual map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
	}

	if actual["message"] != expected["message"] {
		t.Errorf("Resposta incorreta: obtido %v, esperado %v", actual, expected)
	}
}
