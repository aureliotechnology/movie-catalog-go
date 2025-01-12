package app

import (
	"context"
	"time"
)

type HealthService struct {
	db DatabaseChecker
}

type DatabaseChecker interface {
	Ping(ctx context.Context) error
}

// Ajustar a assinatura para aceitar um DatabaseChecker
func NewHealthService(db DatabaseChecker) *HealthService {
	return &HealthService{db: db}
}

func (h *HealthService) CheckHealth() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := h.db.Ping(ctx)
	if err != nil {
		return "database connection error", err
	}

	return "ok", nil
}
