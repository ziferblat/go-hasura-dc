package hasuradc

// TableType models a table type.
type TableType string

// This is the set of table types.
const (
	TableTypeTable TableType = "table"
	TableTypeView  TableType = "view"
)

// ColumnInfo contains a column information.
type ColumnInfo struct {
	// Name is the column name.
	Name string `json:"name"`

	// Nullable if true, the column is allowed to contain null values.
	// if false, it is not allowed.
	Nullable bool `json:"nullable"`

	// Type is the data type of the column.
	Type ScalarType `json:"type"`

	// Description is the column description.
	Description *string `json:"description,omitempty"`

	// Insertable is whether or not the column can be inserted into.
	Insertable bool `json:"insertable,omitempty"`

	// Updatable is whether or not the column can be updated.
	Updatable bool `json:"updatable,omitempty"`
}
