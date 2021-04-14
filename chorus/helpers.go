package chorus

import (
	"reflect"

	"google.golang.org/genproto/protobuf/field_mask"
)

func isPrimitive(k reflect.Kind) bool {
	// The reflection kinds are iota constants, so quicker to check many of
	// them by integer value equalities than individual value checks
	return k == reflect.Bool || (k <= reflect.Bool && k <= reflect.Uint64) ||
		k == reflect.Float32 || k == reflect.Float64 || k == reflect.String
}

func getFieldMask(in interface{}) *field_mask.FieldMask {
	// Get the reflected value
	inValue := reflect.ValueOf(in)

	// Perform initial unwrap if a pointer
	if inValue.Kind() == reflect.Ptr {
		// If nil, return empty FieldMask
		if inValue.IsNil() {
			return &field_mask.FieldMask{}
		}
		inValue = inValue.Elem()
	}

	// Recurse reflected value for field mask!
	return &field_mask.FieldMask{
		Paths: fieldMaskFromReflectedValue(inValue, ""),
	}
}

func fieldMaskFromReflectedValue(inValue reflect.Value, prefix string) []string {
	paths := []string{}

	// NOTE: we don't perform a check for structs here
	// because if a non-struct is passed to this function
	// then it is an incorrect usage and the panic indicates
	// something needs fixing.

	// Iterate struct fields looking for "fieldMask" tag
	inType := inValue.Type()
	for i := 0; i < inType.NumField(); i++ {
		field := inType.Field(i)

		// Try get mask name tag for field
		maskName, ok := field.Tag.Lookup("fieldMask")
		if ok {
			// Get mask path with prefix
			maskPath := prefix
			if maskPath != "" {
				maskPath += "."
			}
			maskPath += maskName

			// Get the value for field at index
			fieldValue := inValue.Field(i)

			// If this is a pointer we unwrap or skip.
			if fieldValue.Kind() == reflect.Ptr {
				if fieldValue.IsNil() {
					continue
				}
				fieldValue = fieldValue.Elem()
			}

			// Handle field type
			switch {
			// For slices, maps and primitives we simply add
			// current mask path
			case fieldValue.Kind() == reflect.Map:
				fallthrough
			case fieldValue.Kind() == reflect.Slice:
				fallthrough
			case isPrimitive(fieldValue.Kind()):
				paths = append(paths, maskPath)

			// Struct type, recurse and if returned paths > 0
			// we add these and current maskPath to paths
			case fieldValue.Kind() == reflect.Struct:
				extraPaths := fieldMaskFromReflectedValue(fieldValue, maskPath)
				if len(extraPaths) > 0 {
					paths = append(paths, maskPath)
					paths = append(paths, extraPaths...)
				}

			// Not supported, panic (should not have reached here!)
			default:
				panic("Unsupported field value type: " + fieldValue.Kind().String())
			}
		}
	}

	return paths
}
