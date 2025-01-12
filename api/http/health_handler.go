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
		"message": h.healthService.GetHealthStatus(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
