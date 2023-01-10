package dc

import "strings"

type Agent struct {
	Capabilities CapabilitiesService
	Health       HealthService
	Schema       SchemaService
	Query        QueryService
}

type GraphQLType string

const (
	GraphQLTypeInt     GraphQLType = "Int"
	GraphQLTypeFloat   GraphQLType = "Float"
	GraphQLTypeString  GraphQLType = "String"
	GraphQLTypeBoolean GraphQLType = "Boolean"
	GraphQLTypeID      GraphQLType = "ID"
)

type ScalarType string

// TableName is the fully qualified name of a table,
// where the last item in the array is the table name
// and any earlier items represent the namespacing of the table name.
type TableName []string

// StringWithSep returns a string representaion of the table name.
func (t *TableName) StringWithSep(sep string) string {
	return strings.Join(*t, sep)
}
