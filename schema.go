package dc

import "context"

type SchemaService interface {
	Get(ctx context.Context, conf any) (*SchemaResponse, error)
}

type SchemaResponse struct {
	// Tables is a list of available tables.
	Tables []TableInfo `json:"tables"`
}

type TableInfo struct {
	// Name is the fully qualified name of the table.
	Name TableName `json:"name"`

	// Columns ...
	Columns []ColumnInfo `json:"columns"`

	Type EntryType `json:"type,omitempty"`

	// Insertable is whether or not new rows can be inserted into the table.
	Insertable *bool `json:"insertable,omitempty"`

	// Updatable is whether or not existing rows can be updated in the table.
	Updatable *bool `json:"updatable,omitempty"`

	// Deletable is whether or not existing rows can be deleted in the table.
	Deletable *bool `json:"deletable,omitempty"`

	// ForeignKeys ...
	ForeignKeys map[string]Constraint `json:"foreign_keys,omitempty"`

	// PrimaryKey ...
	PrimaryKey []string `json:"primary_key,omitempty"`

	// Description ...
	Description *string `json:"description,omitempty"`
}

// ColumnInfo contains a column information.
type ColumnInfo struct {
	// Name is the column name.
	Name string `json:"name"`

	// Nullable if true, the column is allowed to contain null values.
	// if false, it is not allowed.
	Nullable bool `json:"nullable"`

	// Type is the data type of the column.
	Type string `json:"type"`

	// Description is the column description.
	Description *string `json:"description,omitempty"`

	// Insertable is whether or not the column can be inserted into.
	Insertable bool `json:"insertable,omitempty"`

	// Updatable is whether or not the column can be updated.
	Updatable bool `json:"updatable,omitempty"`
}

// EntryType models a entry type.
type EntryType string

// This is the set of table types.
const (
	EntryTypeTable EntryType = "table"
	EntryTypeView  EntryType = "view"
)

type Constraint struct {
	// ForeignTable is the fully qualified name of a foreign table.
	ForeignTable TableName `json:"foreign_table"`

	// ColumnMapping ...
	ColumnMapping map[string]string `json:"column_mapping"`
}
