package http

import (
	"net/http"

	goa "goa.design/goa/v3/pkg"
)

type (
	// ErrorResponse is the default data structure encoded in HTTP responses
	// that correspond to errors created by the generated code. This struct is
	// mainly intended for clients to decode error responses.
	ErrorResponse struct {
		// Name is a name for that class of errors.
		Name string `json:"name" xml:"name" form:"name"`
		// ID is the unique error instance identifier.
		ID string `json:"id" xml:"id" form:"id"`
		// Message describes the specific error occurrence.
		Message string `json:"message" xml:"message" form:"message"`
	}

	// Statuser is implemented by error response object to provide the response
	// HTTP status code.
	Statuser interface {
		// StatusCode return the HTTP status code used to encode the response
		// when not defined in the design.
		StatusCode() int
	}
)

// NewErrorResponse creates a HTTP response from the given error.
func NewErrorResponse(err error) Statuser {
	if gerr, ok := err.(*goa.ServiceError); ok {
		return &ErrorResponse{
			Name:    gerr.Name,
			ID:      gerr.ID,
			Message: gerr.Message,
		}
	}
	return NewErrorResponse(goa.Fault(err.Error()))
}

// StatusCode implements a heuristic that computes a HTTP response status code
// appropriate for the timeout, temporary and fault characteristics of the
// error. This method is used by the generated server code when the error is not
// described explicitly in the design.
func (resp *ErrorResponse) StatusCode() int {
	name := resp.Name
	value, ok := GetCode(name)
	if ok {
		return value
	}
	return http.StatusInternalServerError
}
