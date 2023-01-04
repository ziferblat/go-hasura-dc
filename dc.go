package dc

type Agent struct {
	Health HealthService
	Schema SchemaService
}

// Config represents the configuration properties of the DataConnector Agent.
type Config struct {
	// DB is the database name,
	// e.g. "/ru-central1/HwYLaqGaBt8hmM85wQq2/eCGNADsfxsbFyNrWPZCz".
	DB string `json:"db"`

	// Tables is a list of tables to be exposed.
	Tables []string `json:"tables,omitempty"`

	// InludeMetaTable is whether to expose meta tables.
	InludeMetaTable bool `json:"include_sqlite_meta_tables,omitempty"`

	// InludeMainSchema is whether to opt into exposing tables
	// by their fully qualified names,
	// e.g. ["main", "table42"] instead of just ["table42"].
	ExplicitMainSchema bool `json:"explicit_main_schema,omitempty"`
}

// TableName is the fully qualified name of a table,
// where the last item in the array is the table name
// and any earlier items represent the namespacing of the table name.
type TableName []string
