package hasuradc

import "context"

type HealthService interface {
	// Check checks if the agent is running, or if particular data source
	// is healthy.
	Check(ctx context.Context, conf *Config) error
}
