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

// ConfigSchemasResponse ...
//
// Note. OpenAPI specification name: ConfigSchemaResponse.
type ConfigSchemasResponse struct {
	ConfigSchema OpenApiSchema `json:"config_schema"`

	// OtherSchemas ...
	OtherSchemas map[string]OpenApiSchema `json:"other_schemas"`
}

type OpenApiSchema struct {
	// FIXME: define all supported properties of OpenApiSchema.

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

type Capabilities struct {
	// FIXME: define all supported properties of Capabilities.
	//
	// + comparisons: ComparisonCapabilities;
	// + data_schema: DataSchemaCapabilities;
	// - explain: ExplainCapabilities;
	// - metrics: MetricsCapabilities;
	// + mutations: MutationCapabilities;
	// - queries: QueryCapabilities;
	// - raw: RawCapabilities;
	// + relationships: RelationshipCapabilities;
	// + scalar_types: ScalarTypesCapabilities;
	// - subscriptions: SubscriptionCapabilities;

	// Comparisons ...
	Comparisons *ComparisonCapabilities `json:"comparisons,omitempty"`

	// DataSchema ...
	DataSchema *DataSchemaCapabilities `json:"data_schema,omitempty"`

	Mutations *MutationCapabilities `json:"mutations,omitempty"`

	// Relationships is whether DC supports executing of relationships queries.
	//
	// If the corresponding node is not in the capabilities response,
	// relationships queries to Hasura GraphQL Engine will return the error
	//
	//	"error": {
	//		"code": "not-supported",
	//		"error": "Agents must provide their own dataloader.",
	//		"path": "$"
	//	}
	Relationships *RelationshipCapabilities `json:"relationships,omitempty"`

	ScalarTypes map[ScalarType]ScalarTypeCapabilities `json:"scalar_types,omitempty"`
}

type DataSchemaCapabilities struct {
	// SupportsPrimaryKeys is whether tables can have primary keys.
	SupportsPrimaryKeys bool `json:"supports_primary_keys,omitempty"`

	// SupportsForeignKeys is whether tables can have foreign keys.
	SupportsForeignKeys bool `json:"supports_foreign_keys,omitempty"`

	ColumnNullability ColumnNullability `json:"column_nullability,omitempty"`
}

type ColumnNullability string

const (
	ColumnNullabilityNull           ColumnNullability = "only_nullable"
	ColumnNullabilityNullAndNotNull ColumnNullability = "nullable_and_non_nullable"
)

type ScalarTypeCapabilities struct {
	AggregateFunctions map[string]ScalarType `json:"aggregate_functions,omitempty"`

	ComparisonOperators map[string]ScalarType `json:"comparison_operators,omitempty"`

	GraphQLType GraphQLType `json:"graphql_type,omitempty"`
}

type ComparisonCapabilities struct {
	// Subquery ...
	Subquery *SubqueryComparisonCapabilities `json:"subquery,omitempty"`
}

type SubqueryComparisonCapabilities struct {
	// SupportsRelations is whether support comparisons
	// that involve related tables, i.e. joins.
	SupportsRelations bool `json:"supports_relations,omitempty"`
}

type MutationCapabilities struct {
	AtomicitySupportLevel AtomicitySupportLevel `json:"atomicity_support_level,omitempty"`

	Delete DeleteCapabilities `json:"delete,omitempty"`

	Insert *InsertCapabilities `json:"insert,omitempty"`

	Returning ReturningCapabilities `json:"returning,omitempty"`

	Update UpdateCapabilities `json:"update,omitempty"`
}

// AtomicitySupportLevel describes the level of transactional
// atomicity the agent supports for mutation operations.
type AtomicitySupportLevel string

// List of AtomicitySupportLevel values.
const (
	// FIXME: rename constants

	// If multiple rows are affected in a single operation but one fails,
	// only the failed row's changes will be reverted.
	AtomicitySupportLevelRow AtomicitySupportLevel = "row"

	// If multiple rows are affected in a single operation but one fails,
	// all affected rows in the operation will be reverted.
	AtomicitySupportLevelSingleOperation AtomicitySupportLevel = "single_operation"

	// If multiple operations of only the same type exist in the one mutation request,
	// a failure in one will result in all changes being reverted
	AtomicitySupportLevelHomogeneousOperations AtomicitySupportLevel = "homogeneous_operations"

	// If multiple operations of any type exist in the one mutation request,
	// a failure in one will result in all changes being reverted.
	AtomicitySupportLevelHeterogeneousOperations AtomicitySupportLevel = "heterogeneous_operations"
)

type DeleteCapabilities any

type InsertCapabilities struct {
	// SupportsNestedInserts is whether nested inserts
	// to related tables are supported.
	SupportsNestedInserts bool `json:"supports_nested_inserts,omitempty"`
}

type ReturningCapabilities any

type UpdateCapabilities any

// RelationshipCapabilities is an empty struct.
type RelationshipCapabilities struct{}
