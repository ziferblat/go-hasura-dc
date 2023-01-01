package dc

import "context"

type HealthService interface {
	// Check executes a query to check if a connection to the data source
	// is availabile.
	Check(ctx context.Context, conf any) error
}
