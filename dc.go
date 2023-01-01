package dc

type Agent struct {
	Capabilities CapabilitiesService
	Health       HealthService
	Schema       SchemaService
}

// TableName is the fully qualified name of a table,
// where the last item in the array is the table name
// and any earlier items represent the namespacing of the table name.
type TableName []string
