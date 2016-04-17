package env

import (
	"errors"
	"fmt"
)

// ErrInvalidInterface is returned when the value passed to the Parse function
// is not a pointer to a struct.
var ErrInvalidInterface = errors.New("env: struct parsing: expected pointer to struct")

// ErrUnsupportedFieldKind is returned when the structure contains an
// unsupported variable type. Currently, only string, int32, bool, float32 are
// supported. Also, recursiveness is supported now.
type ErrUnsupportedFieldKind struct {
	FieldName string
	FieldKind string
}

func (e ErrUnsupportedFieldKind) Error() string {
	return fmt.Sprintf("env: set value '%s': unsupported field kind '%s'", e.FieldName, e.FieldKind)
}

// ErrFieldMustBeAssignable is returned when a field in the structure is
// unexported (field name starts with lowercase) or there was an error while
// trying to write to it.
type ErrFieldMustBeAssignable string

func (field ErrFieldMustBeAssignable) Error() string {
	return fmt.Sprintf("env: set value '%s': cannot set value to unexported field", string(field))
}
