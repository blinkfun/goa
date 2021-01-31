package goa

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

type (
	// ServiceError is the default error type used by the goa package to
	// encode and decode error responses.
	ServiceError struct {
		// Name is a name for that class of errors.
		Name string
		// ID is a unique value for each occurrence of the error.
		ID string
		// Message contains the specific error details.
		Message string
		// format is only as key for i18n
		format string
		// arguments for message format
		arguments []interface{}
	}

	// ErrorName is implemented by error to provide the error name.
	ErrorNamer interface {
		ErrorName() string
	}
)

// Fault creates an error given a format and values a la fmt.Printf. The error
// has the Fault field set to true.
func Fault(format string, v ...interface{}) *ServiceError {
	return newError(ErrNameFault, format, v...)
}

// PermanentError creates an error given a name and a format and values a la
// fmt.Printf.
func PermanentError(name, format string, v ...interface{}) *ServiceError {
	return newError(name, format, v...)
}

// TemporaryError is an error class that indicates that the error is temporary
// and that retrying the request may be successful. TemporaryError creates an
// error given a name and a format and values a la fmt.Printf. The error has the
// Temporary field set to true.
func TemporaryError(name, format string, v ...interface{}) *ServiceError {
	return newError(name, format, v...)
}

// PermanentTimeoutError creates an error given a name and a format and values a
// la fmt.Printf. The error has the Timeout field set to true.
func PermanentTimeoutError(name, format string, v ...interface{}) *ServiceError {
	return newError(name, format, v...)
}

// TemporaryTimeoutError creates an error given a name and a format and values a
// la fmt.Printf. The error has both the Timeout and Temporary fields set to
// true.
func TemporaryTimeoutError(name, format string, v ...interface{}) *ServiceError {
	return newError(name, format, v...)
}

// MissingPayloadError is the error produced by the generated code when a
// request is missing a required payload.
func MissingPayloadError() error {
	return PermanentError(ErrNameMissingPayload, "missing required payload")
}

// DecodePayloadError is the error produced by the generated code when a request
// body cannot be decoded successfully.
func DecodePayloadError(msg string) error {
	return PermanentError(ErrNameDecodePayload, msg)
}

// InvalidFieldTypeError is the error produced by the generated code when the
// type of a payload field does not match the type defined in the design.
func InvalidFieldTypeError(name string, val interface{}, expected string) error {
	return PermanentError(ErrNameInvalidFieldType, "invalid value %#v for %q, must be a %s", val, name, expected)
}

// MissingFieldError is the error produced by the generated code when a payload
// is missing a required field.
func MissingFieldError(name, context string) error {
	return PermanentError(ErrNameMissingField, "%q is missing from %s", name, context)
}

// InvalidEnumValueError is the error produced by the generated code when the
// value of a payload field does not match one the values defined in the design
// Enum validation.
func InvalidEnumValueError(name string, val interface{}, allowed []interface{}) error {
	elems := make([]string, len(allowed))
	for i, a := range allowed {
		elems[i] = fmt.Sprintf("%#v", a)
	}
	return PermanentError(ErrNameInvalidEnumValue, "value of %s must be one of %s but got value %#v", name, strings.Join(elems, ", "), val)
}

// InvalidFormatError is the error produced by the generated code when the value
// of a payload field does not match the format validation defined in the
// design.
func InvalidFormatError(name, target string, format Format, formatError error) error {
	return PermanentError(ErrNameInvalidFormat, "%s must be formatted as a %s but got value %q, %s", name, format, target, formatError.Error())
}

// InvalidPatternError is the error produced by the generated code when the
// value of a payload field does not match the pattern validation defined in the
// design.
func InvalidPatternError(name, target string, pattern string) error {
	return PermanentError(ErrNameInvalidPattern, "%s must match the regexp %q but got value %q", name, pattern, target)
}

// InvalidRangeError is the error produced by the generated code when the value
// of a payload field does not match the range validation defined in the design.
// value may be an int or a float64.
func InvalidRangeError(name string, target interface{}, value interface{}, min bool) error {
	format := "%s must be greater or equal than %d but got value %#v"
	if !min {
		format = "%s must be lesser or equal than %d but got value %#v"
	}
	return PermanentError(ErrNameInvalidRange, format, name, value, target)
}

// InvalidLengthError is the error produced by the generated code when the value
// of a payload field does not match the length validation defined in the
// design.
func InvalidLengthError(name string, target interface{}, ln, value int, min bool) error {
	format := "length of %s must be greater or equal than %d but got value %#v (len=%d)"
	if !min {
		format = "length of %s must be lesser or equal than %d but got value %#v (len=%d)"
	}
	return PermanentError(ErrNameInvalidLength, format, name, value, target, ln)
}

// NewErrorID creates a unique 8 character ID that is well suited to use as an
// error identifier.
func NewErrorID() string {
	// for the curious - simplifying a bit - the probability of 2 values
	// being equal for n 6-bytes values is n^2 / 2^49. For n = 1 million
	// this gives around 1 chance in 500. 6 bytes seems to be a good
	// trade-off between probability of clashes and length of ID (6 * 4/3 =
	// 8 chars) since clashes are not catastrophic.
	b := make([]byte, 6)
	io.ReadFull(rand.Reader, b)
	return base64.RawURLEncoding.EncodeToString(b)
}

// MergeErrors updates an error by merging another into it. It first converts
// other into a ServiceError if not already one. The merge algorithm then:
//
// * uses the name of err if a ServiceError, the name of other otherwise.
//
// * appends both error messages.
//
// * computes Timeout and Temporary by "and"ing the fields of both errors.
//
// Merge returns the updated error. This makes it possible to return other when
// err is nil.
func MergeErrors(err, other error) error {
	if err == nil {
		if other == nil {
			return nil
		}
		return other
	}
	if other == nil {
		return err
	}
	e := asError(err)
	o := asError(other)
	if e.Name == ErrNameError {
		e.Name = o.Name
	}
	e.Message = e.Message + "; " + o.Message

	return e
}

// Error returns the error message.
func (s *ServiceError) Error() string { return s.Message }

// ErrorName returns the error name.
func (s *ServiceError) ErrorName() string { return s.Name }

func (s *ServiceError) Format() string {
	return s.format
}

func (s *ServiceError) Arguments() []interface{} {
	return s.arguments
}

func newError(name string, format string, v ...interface{}) *ServiceError {
	return &ServiceError{
		Name:      name,
		ID:        NewErrorID(),
		Message:   fmt.Sprintf(format, v...),
		format:    format,
		arguments: v,
	}
}

func asError(err error) *ServiceError {
	e, ok := err.(*ServiceError)
	if !ok {
		return &ServiceError{
			Name:    ErrNameError,
			ID:      NewErrorID(),
			Message: err.Error(),
		}
	}
	return e
}
