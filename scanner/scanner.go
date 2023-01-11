package scanner

import (
	"database/sql"
)

// Scan copies the columns in the current row into the map pointed at by dest.
func ScanMap(rs *sql.Rows, dest map[string]interface{}) error {
	return copy(rs, dest)
}

func copy(rs *sql.Rows, dest map[string]interface{}) error {
	cols, err := rs.Columns()
	if err != nil {
		return err
	}
	len := len(cols)
	vals := make([]interface{}, len)
	ptrs := make([]interface{}, len)
	for i := 0; i < len; i++ {
		ptrs[i] = &vals[i]
	}
	if err := rs.Scan(ptrs...); err != nil {
		return err
	}
	for i, col := range cols {
		val := ptrs[i].(*interface{})
		dest[col] = *val
	}

	return nil
}
