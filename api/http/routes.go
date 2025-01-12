package http

import (
	"net/http"

	"movie-catalog-go/internal/core/port/in"
)

func RegisterRoutes(mux *http.ServeMux, healthService in.HealthPort) {
	healthHandler := NewHealthHandler(healthService)

	// Registrar rotas
	mux.HandleFunc("/health", healthHandler.HandleHealth)
	mux.HandleFunc("/health-db", healthHandler.HandleHealthDB)
}
