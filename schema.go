package hasuradc

// TableType models a table type.
type TableType string

// This is the set of table types.
const (
	TableTypeTable TableType = "table"
	TableTypeView  TableType = "view"
)
