package goa

// Error names are for internal use and can not be overridden by RegisterCode
const (
	// Client side codes
	ErrNameDecodePayload    = "decode_payload"
	ErrNameInvalidEnumValue = "invalid_enum_value"
	ErrNameInvalidFieldType = "invalid_field_type"
	ErrNameInvalidFormat    = "invalid_format"
	ErrNameInvalidLength    = "invalid_length"
	ErrNameInvalidPattern   = "invalid_pattern"
	ErrNameInvalidRange     = "invalid_range"
	ErrNameInvalidType      = "invalid_type"
	ErrNameMissingField     = "missing_field"
	ErrNameMissingPayload   = "missing_payload"
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
