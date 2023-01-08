package dc

import "context"

type QueryService interface {
	Exec(ctx context.Context, conf any, q Query) (*ExecQueryResponse, error)
}

type QueryRequest struct {
	Table              TableName            `json:"table"`
	Query              Query                `json:"query"`
	TableRelationships []TableRelationships `json:"table_relationships"`
}

type ExecQueryResponse struct{}

type Query struct{}

type TableRelationships struct{}
