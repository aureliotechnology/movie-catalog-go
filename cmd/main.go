package main

import (
	"log"
	"net/http"

	httpHandler "movie-catalog-go/api/http"
	"movie-catalog-go/internal/app"
)

func main() {
	// Inicializar o serviço de saúde
	healthService := app.NewHealthService()

	// Configurar manipulador HTTP
	healthHandler := httpHandler.NewHealthHandler(healthService)

	// Registrar rotas
	http.HandleFunc("/health", healthHandler.HandleHealth)

	// Iniciar servidor
	port := ":8080"
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
