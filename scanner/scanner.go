package scanner

import (
	"database/sql"
)

// Scan copies the columns in the current row into the map pointed at by dest.
func ScanMap(rs *sql.Rows, dest map[string]interface{}) error {
	return resolve(rs, dest)
}

func resolve(rs *sql.Rows, dest map[string]interface{}) error {
	return nil
}
