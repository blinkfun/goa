package goa

// Error names are for internal use and can not be overridden by RegisterCode
const (
	ErrNameMissingPayload   = "missing_payload"
	ErrNameDecodePayload    = "decode_payload"
	ErrNameInvalidFieldType = "invalid_field_type"
	ErrNameMissingField     = "missing_field"
	ErrNameInvalidEnumValue = "invalid_enum_value"
	ErrNameInvalidFormat    = "invalid_format"
	ErrNameInvalidPattern   = "invalid_pattern"
	ErrNameInvalidRange     = "invalid_range"
	ErrNameInvalidLength    = "invalid_length"
	ErrNameFault            = "fault"
)

var ReservedErrNames = []string{
	ErrNameMissingPayload,
	ErrNameDecodePayload,
	ErrNameInvalidFieldType,
	ErrNameMissingField,
	ErrNameInvalidEnumValue,
	ErrNameInvalidFormat,
	ErrNameInvalidPattern,
	ErrNameInvalidLength,
	ErrNameFault,
}
