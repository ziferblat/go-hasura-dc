package sqlext

import (
	"database/sql"
	"strconv"
)

// Scan copies the columns in the current row into the map pointed at by dest.
// If possible, it converts val to scalar type.
func ScanMap(rs *sql.Rows, dest map[string]any) error {
	return copy(rs, dest)
}

func copy(rs *sql.Rows, dest map[string]any) error {
	cols, err := rs.Columns()
	if err != nil {
		return err
	}
	len := len(cols)
	vals := make([]any, len)
	ptrs := make([]any, len)
	for i := 0; i < len; i++ {
		ptrs[i] = &vals[i]
	}
	if err := rs.Scan(ptrs...); err != nil {
		return err
	}
	for i, col := range cols {
		val := ptrs[i].(*any)
		dest[col] = decode(*val)
	}

	return nil
}

func decode(val any) any {
	bb, ok := val.([]byte)
	if !ok {
		return val
	}
	s := string(bb)

	if rval, err := strconv.Atoi(s); err == nil {
		return rval
	}
	if rval, err := strconv.ParseFloat(s, 64); err == nil {
		return rval
	}
	if rval, err := strconv.ParseBool(s); err == nil {
		return rval
	}

	return s
}
