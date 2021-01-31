package goa

// Error names are for internal use and can not be overridden by RegisterCode
const (
	// Client side codes
	ErrNameDecodePayload    = "decode_payload"
	ErrNameDecodingError    = "decoding_error"
	ErrNameEncodingError    = "encoding_error"
	ErrNameInvalidEnumValue = "invalid_enum_value"
	ErrNameInvalidFieldType = "invalid_field_type"
	ErrNameInvalidFormat    = "invalid_format"
	ErrNameInvalidLength    = "invalid_length"
	ErrNameInvalidPattern   = "invalid_pattern"
	ErrNameInvalidRange     = "invalid_range"
	ErrNameInvalidResponse  = "invalid_response"
	ErrNameInvalidType      = "invalid_type"
	ErrNameInvalidURL       = "invalid_url"
	ErrNameMissingField     = "missing_field"
	ErrNameMissingPayload   = "missing_payload"
	ErrNameRequestError     = "request_error"
	ErrNameValidationError  = "validation_error"
	// Server side codes
	ErrNameError = "error"
	ErrNameFault = "fault"
)

var ReservedErrNames = []string{
	// Client side codes
	ErrNameDecodePayload,
	ErrNameInvalidEnumValue,
	ErrNameInvalidFieldType,
	ErrNameInvalidFormat,
	ErrNameInvalidLength,
	ErrNameInvalidPattern,
	ErrNameInvalidRange,
	ErrNameInvalidType,
	ErrNameMissingField,
	ErrNameMissingPayload,
	// Server side codes
	ErrNameFault,
}
