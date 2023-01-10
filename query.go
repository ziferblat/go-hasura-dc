package dc

import "context"

type QueryService interface {
	Exec(ctx context.Context, conf any, q QueryRequest) (*QueryResponse, error)
}

type QueryRequest struct {
	Table              TableName            `json:"table"`
	Query              Query                `json:"query"`
	TableRelationships []TableRelationships `json:"table_relationships"`
}

type Query struct {
	Aggregates map[string]Aggregate `json:"aggregates,omitempty"`

	Fields map[string]ColumnField `json:"fields,omitempty"`

	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

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

	Rows []map[string]any `json:"rows,omitempty"`
}
