package hasuradc

type Agent struct {
	Health HealthService
}

type Config struct {
	DB string
}

// ScalarType models a scalar type.
type ScalarType string

// This is the set of scalar types.
const (
	ScalarTypeString ScalarType = "string"
	ScalarTypeNumber ScalarType = "number"
	ScalarTypeBool   ScalarType = "bool"
)
