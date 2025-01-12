package in

type HealthPort interface {
	CheckHealth() (string, error)
}
