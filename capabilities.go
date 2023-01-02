package dc

import "context"

type CapabilitiesService interface {
	Get(ctx context.Context, conf any) (*CapabilitiesResponse, error)
}

type CapabilitiesResponse struct {
	DisplayName string `json:"display_name,omitempty"`

	ReleaseName string `json:"release_name,omitempty"`

	// ConfigSchemas represents an agent configuration.
	ConfigSchemas ConfigSchemasResponse `json:"config_schemas"`

	Capabilities Capabilities `json:"capabilities"`
}

type Capabilities struct {
	DataSchema DataSchemaCapabilities `json:"data_schema,omitempty"`
}

type DataSchemaCapabilities struct {
	ColumnNullability ColumnNullability `json:"column_nullability,omitempty"`

	// SupportsForeignKeys is whether tables can have foreign keys.
	SupportsForeignKeys bool `json:"supports_foreign_keys,omitempty"`

	// SupportsPrimaryKeys is whether tables can have primary keys.
	SupportsPrimaryKeys bool `json:"supports_primary_keys,omitempty"`
}

type ColumnNullability string

const (
	ColumnNullabilityNull           ColumnNullability = "only_nullable"
	ColumnNullabilityNullAndNotNull ColumnNullability = "nullable_and_non_nullable"
)

// ConfigSchemasResponse ...
//
// Note. OpenAPI specification name: ConfigSchemaResponse.
type ConfigSchemasResponse struct {
	ConfigSchema OpenApiSchema `json:"config_schema"`

	OtherSchemas map[string]OpenApiSchema `json:"other_schemas"`
}

type OpenApiSchema struct {
	// Nullable is whether the value may be null.
	Nullable bool `json:"nullable,omitempty"`

	// Type defines the data type of an OpenAPI schema.
	Type OASDataType `json:"type,omitempty"`
}

type OASDataType string

const (
	OASDataTypeString  OASDataType = "string"
	OASDataTypeNumber  OASDataType = "number"
	OASDataTypeInteger OASDataType = "integer"
	OASDataTypeBoolean OASDataType = "boolean"
	OASDataTypeArray   OASDataType = "array"
	OASDataTypeObject  OASDataType = "object"
)
