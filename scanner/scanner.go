package scanner

import (
	"database/sql"
	"strconv"
)

// Scan copies the columns in the current row into the map pointed at by dest.
// If possible, it converts val to scalar type.
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
		dest[col] = decode(*val)
	}

	return nil
}

func decode(val interface{}) interface{} {
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
