package hasuradc

type Agent struct {
	Health HealthService
	Schema SchemaService
}

type Config struct {
	DB string
}
