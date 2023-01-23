package dc

import (
	"testing"
)

func TestQueryRequest_Marshal(t *testing.T) {
	var tests = map[string]struct {
		qr   QueryRequest
		want string
	}{
		"no relationships": {
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
		"with relationships": {
			qr: QueryRequest{
				Table: []string{"Track"},
				TableRelationships: []TableRelationships{
					{
						SourceTable: []string{"Track"},
						Relationships: map[string]Relationship{
							"Album": {
								TargetTable:      []string{"Album"},
								RelationshipType: "object",
								ColumnMapping: map[string]string{
									"AlbumId": "AlbumId",
								},
							},
						},
					},
					{
						SourceTable: []string{"Album"},
						Relationships: map[string]Relationship{
							"Artist": {
								TargetTable:      []string{"Artist"},
								RelationshipType: "object",
								ColumnMapping: map[string]string{
									"ArtistId": "ArtistId",
								},
							},
						},
					},
				},
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
						"Album": {
							Type:         FieldTypeRelationship,
							Relationship: "Album",
							Query: Query{
								Fields: map[string]Field{
									"Artist": {
										Type:         FieldTypeRelationship,
										Relationship: "Artist",
										Query: Query{
											Fields: map[string]Field{
												"Name": {
													Type:       FieldTypeColumn,
													Column:     "Name",
													ColumnType: "string",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: `{
				"table": [
					"Track"
				],
				"table_relationships": [
					{
						"source_table": [
							"Track"
						],
						"relationships": {
							"Album": {
								"target_table": [
									"Album"
								],
								"relationship_type": "object",
								"column_mapping": {
									"AlbumId": "AlbumId"
								}
							}
						}
					},
					{
						"source_table": [
							"Album"
						],
						"relationships": {
							"Artist": {
								"target_table": [
									"Artist"
								],
								"relationship_type": "object",
								"column_mapping": {
									"ArtistId": "ArtistId"
								}
							}
						}
					}
				],
				"query": {
					"fields": {
						"Name": {
							"type": "column",
							"column": "Name",
							"column_type": "string"
						},
						"Album": {
							"type": "relationship",
							"relationship": "Album",
							"query": {
								"fields": {
									"Artist": {
										"type": "relationship",
										"relationship": "Artist",
										"query": {
											"fields": {
												"Name": {
													"type": "column",
													"column": "Name",
													"column_type": "string"
												}
											}
										}
									}
								}
							}
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
