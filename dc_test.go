package dc

import "testing"

func TestTableName_StringWithSep(t *testing.T) {
	tests := []struct {
		tbl  TableName
		sep  string
		want string
	}{
		{
			TableName{"film"},
			"",
			"film",
		},
		{
			TableName{"public", "film"},
			".",
			"public.film",
		},
		{
			TableName{"foo", "bar", "table559"},
			"/",
			"foo/bar/table559",
		},
	}
	for _, tt := range tests {
		if got := tt.tbl.StringWithSep(tt.sep); got != tt.want {
			t.Errorf("got %s, want %s", got, tt.want)
		}
	}
}
