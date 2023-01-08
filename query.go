package dc

import "context"

type QueryService interface {
	Exec(ctx context.Context, conf any, q any) (*ExecQueryResponse, error)
}

type ExecQueryResponse struct{}
