package dc

import (
	"testing"
)

func TestQueryRequest_Marshal(t *testing.T) {
	var tests = map[string]struct {
		qr   QueryRequest
		want string
	}{
		"no nested fields": {
			qr: QueryRequest{
				Table:              []string{"Track"},
				TableRelationships: []TableRelationships{},
				Query: Query{
					Fields: map[string]Field{
						"TrackId": {
							Type:       FieldTypeColumn,
							Column:     "TrackId",
							ColumnType: "number",
						},
						"Name": {
							Type:       FieldTypeColumn,
							Column:     "Name",
							ColumnType: "string",
						},
					},
				},
			},
			want: `{
				"table": [
					"Track"
				],
				"table_relationships": [],
				"query": {
					"fields": {
						"Name": {
							"type": "column",
							"column": "Name",
							"column_type": "string"
						},
						"TrackId": {
							"type": "column",
							"column": "TrackId",
							"column_type": "number"
						}
					}
				}
			}`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			testJSONMarshal(t, tt.qr, tt.want)
		})
	}
}
