package ref

import "reflect"

// IsNil は、v が ( typed ) nil かどうかを返す
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return value.IsNil()
	}

	return false
}
