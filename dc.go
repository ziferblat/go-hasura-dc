package dc

type Agent struct {
	Health HealthService
	Schema SchemaService
}

// Config represents a model of the X-Hasura-DataConnector-Config header.
type Config struct {
	// DB is the database name,
	// e.g. "/ru-central1/HwYLaqGaBt8hmM85wQq2/eCGNADsfxsbFyNrWPZCz".
	DB string
}

// TableName is the fully qualified name of a table,
// where the last item in the array is the table name
// and any earlier items represent the namespacing of the table name.
type TableName []string
