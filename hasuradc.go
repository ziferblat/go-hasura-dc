package hasuradc

type Agent struct {
	Health HealthService
}

type Config struct {
	DB string
}
