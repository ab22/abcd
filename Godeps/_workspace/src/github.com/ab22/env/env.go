package env

import (
	"os"
	"reflect"
	"strconv"
)

// Parse is the main function of this package. It takes any valid interface{}
// value as parameter and expects it to be a pointer to any structure.
// It attempts to read the 'env' and 'envDefault'
//
// An error is returned if the value passed as parameter is not a pointer to a
// structure and it returns an InvalidInterfaceError.
func Parse(i interface{}) error {
	if isInvalidInterface(i) {
		return InvalidInterfaceError
	}

	elem := reflect.ValueOf(i).Elem()
	return setStructValues(&elem)
}

// isInvalidInterface determines whether i is a pointer to a structure.
// Returns true if it's not a pointer to a structure.
func isInvalidInterface(i interface{}) bool {
	if i == nil {
		return true
	}

	interfaceValue := reflect.ValueOf(i)
	interfaceKind := interfaceValue.Kind()

	if interfaceKind != reflect.Ptr {
		return true
	}

	interfaceKind = interfaceValue.Elem().Kind()

	return interfaceKind != reflect.Struct
}

// setStructValues iterates over each of the structures elements to determine
// what value must be set to the field. If no env value of default env value
// was found, then we skip the current field.
func setStructValues(structElem *reflect.Value) error {
	numFields := structElem.NumField()
	structType := structElem.Type()

	for i := 0; i < numFields; i++ {
		structField := structType.Field(i)
		fieldValue := structElem.Field(i)

		if fieldValue.Kind() == reflect.Struct {
			if err := setValue(&fieldValue, structField.Name, ""); err != nil {
				return err
			}

			continue
		}

		envValue := getEnvOrDefaultValue(&structField)

		if envValue == "" {
			continue
		}

		if err := setValue(&fieldValue, structField.Name, envValue); err != nil {
			return err
		}
	}

	return nil
}

// getEnvOrDefaultValue attempts to read the env tags from the structure field,
// and returns the environment value found or the default value.
//
// Two env tags can be set, and those are the 'env' tag and the 'envDefault'
// tags. The 'env' tag contains the name of the environment key, and the
// 'envDefault' tag contains the default value for the structure field.
func getEnvOrDefaultValue(field *reflect.StructField) string {
	envKey := field.Tag.Get("env")
	defaultValue := field.Tag.Get("envDefault")

	if envKey == "" {
		return defaultValue
	}

	envValue := os.Getenv(envKey)

	if envValue == "" {
		return defaultValue
	}

	return envValue
}

// setValue function is in charge of determining if the field is an
// exported field (accessable). We return a FieldMustBeAssignableError if
// the field is unaccessable.
//
// For each of the supported types, we just parse the string value into the
// corresponding type and assigned it to the field. If there's an error when
// parsing a type, then we return that error.
//
// For structures, we recursively call the Parse function in this package
// to attempt to set all values.
//
// If an unsupported field is found, then we return an
// UnsupportedFieldKindError.
func setValue(field *reflect.Value, fieldName string, envValue string) error {
	if !field.CanSet() {
		return FieldMustBeAssignableError(fieldName)
	}

	fieldKind := field.Kind()

	switch fieldKind {
	case reflect.String:
		field.SetString(envValue)
	case reflect.Int:
		intValue, err := strconv.ParseInt(envValue, 10, 32)

		if err != nil {
			return err
		}

		field.SetInt(intValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(envValue)

		if err != nil {
			return err
		}

		field.SetBool(boolValue)
	case reflect.Float32:

		floatValue, err := strconv.ParseFloat(envValue, 32)

		if err != nil {
			return err
		}

		field.SetFloat(floatValue)
	case reflect.Struct:
		return Parse(field.Addr().Interface())
	default:
		return &UnsupportedFieldKindError{
			FieldName: fieldName,
			FieldKind: fieldKind.String(),
		}
	}

	return nil
}
