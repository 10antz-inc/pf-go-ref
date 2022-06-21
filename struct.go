package ref

import (
	"fmt"
	"reflect"
	"unicode"

	"github.com/tys-muta/go-opt"
	"github.com/tys-muta/go-ref/option"
)

type ScanHandle func(field reflect.StructField, value reflect.Value) error

func ScanField(v interface{}, handle ScanHandle, options ...opt.Option) error {
	o := &option.StructOptions{}
	if err := opt.Reflect(o, options...); err != nil {
		return fmt.Errorf("failed to reflect: %w", err)
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

func ToMap(v interface{}, options ...opt.Option) (map[string]interface{}, error) {
	dst := map[string]interface{}{}

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
