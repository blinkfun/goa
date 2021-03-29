package grpc

import (
	"fmt"

	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

type (
	// ClientError is an error returned by a gRPC service client.
	ClientError struct {
		// Name is a name for this class of errors.
		Name string
		// Message contains the specific error details.
		Message string
		// Service is the name of the service.
		Service string
		// Method is the name of the service method.
		Method string
	}
)

var (
	DefaultCode = codes.Unknown
)

// NewErrorResponse creates a new ErrorResponse protocol buffer message from
// the given error. If the given error is a goa ServiceError, the ErrorResponse
// message will be set with the corresponding Timeout, Temporary, and Fault
// characteristics. If the error is not a goa ServiceError, it creates an
// ErrorResponse message with the Fault field set to true.
func NewErrorResponse(err error) *goapb.ErrorResponse {
	if gerr, ok := err.(*goa.ServiceError); ok {
		return &goapb.ErrorResponse{
			Name:    gerr.Name,
			ID:      gerr.ID,
			Message: gerr.Message,
		}
	}
	return NewErrorResponse(goa.Fault(err.Error()))
}

// NewServiceError returns a goa ServiceError type for the given ErrorResponse
// message.
func NewServiceError(resp *goapb.ErrorResponse) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    resp.Name,
		ID:      resp.ID,
		Message: resp.Message,
	}
}

// NewStatusError creates a gRPC status error with the error response
// messages added to its details.
func NewStatusError(code codes.Code, err error, details ...protoiface.MessageV1) error {
	st := status.New(code, err.Error())
	if s, err := st.WithDetails(details...); err == nil {
		return s.Err()
	}
	return st.Err()
}

// EncodeError returns a gRPC status error from the given error with the error
// response encoded in the status details. If error is a goa ServiceError type
// it implements a heuristic to compute the status code from the Timeout,
// Fault, and Temporary characteristics of the ServiceError. If error is not a
// ServiceError or a gRPC status error it returns a gRPC status error with
// Unknown code and Fault characteristic set.
func EncodeError(err error) error {
	if st, ok := status.FromError(err); ok {
		if s, err := st.WithDetails(NewErrorResponse(err)); err == nil {
			return s.Err()
		}
		return st.Err()
	}
	if gerr, ok := err.(*goa.ServiceError); ok {
		// goa service error type. Compute the status code from the service error
		// characteristics and create a new detailed gRPC status error.
		c, ok := GetCode(gerr.Name)

		var code = codes.Unknown
		if ok {
			code = c
		}

		return NewStatusError(code, err, NewErrorResponse(err))
	}
	// Return an unknown gRPC status error with fault characteristic set.
	return NewStatusError(DefaultCode, err, NewErrorResponse(err))
}

// DecodeError returns the error message encoded in the status details if error
// is a gRPC status error. It assumes that the error message is encoded as the
// first item in the details. It returns nil if the error is not a gRPC status
// error or if no detail is found.
func DecodeError(err error) proto.Message {
	st, ok := status.FromError(err)
	if !ok {
		return nil
	}
	details := st.Details()
	if len(details) == 0 {
		return nil
	}
	return details[0].(proto.Message)
}

// ErrInvalidType is the error returned when the wrong type is given to a
// encoder or decoder.
func ErrInvalidType(svc, m, expected string, actual interface{}) error {
	msg := fmt.Sprintf("invalid value expected %s, got %v", expected, actual)
	return &ClientError{Name: goa.ErrNameInvalidType, Message: msg, Service: svc, Method: m}
}

// Error builds an error message.
func (c *ClientError) Error() string {
	return fmt.Sprintf("[%s %s]: %s", c.Service, c.Method, c.Message)
}
