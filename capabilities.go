package dc

import "context"

type CapabilitiesService interface {
	Get(ctx context.Context, conf any) (*CapabilitiesResponse, error)
}

type CapabilitiesResponse struct {
	Capabilities Capabilities `json:"capabilities"`

	// ConfigSchemas represents an agent configuration.
	ConfigSchemas ConfigSchemasResponse `json:"config_schemas"`

	DisplayName string `json:"display_name,omitempty"`

	ReleaseName string `json:"release_name,omitempty"`
}

type Capabilities struct {
	ConfigSchema OpenApiSchema `json:"config_schema"`

	OtherSchemas map[string]OpenApiSchema `json:"other_schemas"`
}

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

// OASDataType is a set of OpenAPI schema data types.
type OASDataType string

const (
	OASDataTypeString  OASDataType = "string"
	OASDataTypeNumber  OASDataType = "number"
	OASDataTypeInteger OASDataType = "integer"
	OASDataTypeBoolean OASDataType = "boolean"
	OASDataTypeArray   OASDataType = "array"
	OASDataTypeObject  OASDataType = "object"
)
