package dc

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Test whether the marshaling of v produces JSON that corresponds
// to the want string.
func testJSONMarshal(t *testing.T, v interface{}, s string) {
	t.Helper()

	// Unmarshal the wanted JSON, to verify its correctness,
	// and marshal it back to sort the keys
	u := reflect.New(reflect.TypeOf(v)).Interface()
	if err := json.Unmarshal([]byte(s), &u); err != nil {
		t.Fatalf("unable to unmarshal JSON for \n%s\nerror: %v", s, err)
	}
	want, err := json.Marshal(u)
	if err != nil {
		t.Fatalf("unable to marshal JSON for \n%#v\nerror: %v", u, err)
	}

	// Marshal the target value
	got, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("unable to marshal JSON for \n%#v\nerror: %v", v, err)
	}

	if string(got) != string(want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
