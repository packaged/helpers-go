package helpers

import (
	"fmt"
	"reflect"
)

// PPull takes a slice of structs and a field name, returning a map keyed by that field's value.
// The field must be exported and its value must be a valid map key type.
func PPull[T any](list []T, keyProperty string) (map[string]T, error) {
	result := make(map[string]T, len(list))
	for _, item := range list {
		v := reflect.ValueOf(item)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("ppull: expected struct, got %s", v.Kind())
		}

		field := v.FieldByName(keyProperty)
		if !field.IsValid() {
			return nil, fmt.Errorf("ppull: field %q not found", keyProperty)
		}

		key := fmt.Sprintf("%v", field.Interface())
		result[key] = item
	}
	return result, nil
}
