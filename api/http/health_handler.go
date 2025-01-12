package http

import (
	"encoding/json"
	"net/http"

	"movie-catalog-go/internal/core/port/in"
)

type HealthHandler struct {
	healthService in.HealthPort
}

func NewHealthHandler(healthService in.HealthPort) *HealthHandler {
	return &HealthHandler{healthService: healthService}
}

func (h *HealthHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *HealthHandler) HandleHealthDB(w http.ResponseWriter, r *http.Request) {
	status, err := h.healthService.CheckHealth()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response := map[string]string{
		"message": status,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
