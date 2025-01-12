package main

import (
	"context"
	"log"
	"net/http"
	"os"

	httpHandler "movie-catalog-go/api/http"
	"movie-catalog-go/internal/adapter/repository"
	"movie-catalog-go/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	_ = godotenv.Load()

	// Configurar conexão com o banco
	ctx := context.Background()
	db, err := repository.NewDatabaseConnection(
		ctx,
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Pool.Close()

	// Inicializar o serviço de saúde
	healthService := app.NewHealthService(db)

	// Configurar manipuladores
	healthHandler := httpHandler.NewHealthHandler(healthService)

	// Registrar rotas
	http.HandleFunc("/health", healthHandler.HandleHealth)
	http.HandleFunc("/health-db", healthHandler.HandleHealthDB)

	// Iniciar servidor
	port := ":8080"
	log.Printf("Servidor rodando na porta %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
