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

	Fields map[string]ColumnField `json:"fields,omitempty"`

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

type ColumnField struct {
	Column string `json:"column"`

	ColumnType ScalarType `json:"column_type"`

	Type string `json:"type"`
}

type OrderBy struct{}

type Where struct{}

type TableRelationships struct{}

type QueryResponse struct {
	Aggregates map[string]any `json:"aggregates,omitempty"`

	// Rows is a set of rows returned by the query, corresponding
	// to the query's fields.
	//
	// FIXME: Mark this prop as required in the Agent OpenAPI specification.
	Rows []map[string]any `json:"rows"`
}
