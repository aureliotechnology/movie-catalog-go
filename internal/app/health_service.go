package app

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (h *HealthService) GetHealthStatus() string {
	return "ok"
}
