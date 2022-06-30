package ref

import (
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {
	foo := "test"
	bar := Ptr(foo)

	if foo != *bar {
		t.Errorf("not match value. got: %s, want: %s", *bar, foo)
	}

	if reflect.ValueOf(bar).Kind() != reflect.Ptr {
		t.Errorf("not pointer")
	}

	if reflect.Indirect(reflect.ValueOf(bar)).Kind() != reflect.String {
		t.Errorf("not string")
	}
}
