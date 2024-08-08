package ref

import (
	"fmt"
	"reflect"
	"unicode"

	"github.com/10antz-inc/pf-go-ref/option"
)

type ScanHandle func(field reflect.StructField, value reflect.Value) error

// ScanField は v のフィールドをスキャンして handle に渡す
func ScanField(v any, handle ScanHandle, options ...option.ScanOption) error {
	o := option.Scan{}
	for _, opt := range options {
		opt(&o)
	}

	e := reflect.Indirect(reflect.ValueOf(v))
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		field := t.Field(i)
		value := e.Field(i)
		if !isPublic(field.Name) && !bool(o.PrivateField) {
			continue
		}
		if err := handle(field, value); err != nil {
			return fmt.Errorf("failed to handle: %w", err)
		}
	}

	return nil
}

// ToMap は v のフィールドをマップにして返す
func ToMap(v any, options ...option.ScanOption) (map[string]any, error) {
	dst := map[string]any{}

	if err := ScanField(v, func(field reflect.StructField, value reflect.Value) error {
		if value.Kind() == reflect.Ptr {
			if !value.IsNil() {
				dst[field.Name] = value.Interface()
			} else {
				dst[field.Name] = nil
			}
		} else {
			dst[field.Name] = value
		}
		return nil
	}, options...); err != nil {
		return nil, fmt.Errorf("failed to scan field: %w", err)
	}

	return dst, nil
}

func isPublic(fieldName string) bool {
	if fieldName == "" {
		return false
	}
	r := rune(fieldName[0])
	return unicode.IsUpper(r)
}
