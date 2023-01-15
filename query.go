package dc

import "context"

type QueryService interface {
	Exec(ctx context.Context, conf any, rq QueryRequest) (*QueryResponse, error)
}

type QueryRequest struct {
	Table              TableName            `json:"table"`
	Query              Query                `json:"query"`
	TableRelationships []TableRelationships `json:"table_relationships"`
}

type Query struct {
	Aggregates map[string]Aggregate `json:"aggregates,omitempty"`

	Fields map[string]Field `json:"fields,omitempty"`

	// Limit is the maximum number of rows to return.
	// Null value is treated as no limit.
	Limit *uint `json:"limit,omitempty"`

	// Offset is the number of rows to skip before returning.
	// Null value is treated as 0.
	Offset *uint `json:"offset,omitempty"`

	OrderBy *OrderBy `json:"order_by,omitempty"`

	Where *Where `json:"where,omitempty"`
}

type Aggregate struct{}

type Field struct {
	Column string `json:"column,omitempty"`

	ColumnType ScalarType `json:"column_type,omitempty"`

	Query Query `json:"query,omitempty"`

	Relationship string `json:"relationship,omitempty"`

	Type FieldType `json:"type"`
}

type FieldType string

const (
	FieldTypeColumn       FieldType = "column"
	FieldTypeRelationship FieldType = "relationship"
)

type OrderBy struct{}

type Where struct{}

type TableRelationships struct {
	Relationships map[string]Relationship `json:"relationships"`

	SourceTable TableName `json:"source_table"`
}

type Relationship struct {
	ColumnMapping map[string]string `json:"column_mapping"`

	RelationshipType RelationshipType `json:"relationship_type"`

	TargetTable TableName `json:"target_table"`
}

type RelationshipType string

const (
	RelationshipTypeObject RelationshipType = "object"
	RelationshipTypeArray  RelationshipType = "array"
)

type QueryResponse struct {
	Aggregates map[string]any `json:"aggregates,omitempty"`

	// Rows is a set of rows returned by the query, corresponding
	// to the query's fields.
	//
	// FIXME: Mark this prop as required in the Agent OpenAPI specification.
	// Also look at https://github.com/hasura/graphql-engine/issues/9380
	Rows []map[string]any `json:"rows"`
}
