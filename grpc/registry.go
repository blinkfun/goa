package grpc

import (
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

var registeredCode = make(map[string]codes.Code)

func init() {
	registeredCode[goa.ErrNameDecodePayload] = codes.InvalidArgument
	registeredCode[goa.ErrNameInvalidFieldType] = codes.InvalidArgument
	registeredCode[goa.ErrNameInvalidEnumValue] = codes.InvalidArgument
	registeredCode[goa.ErrNameInvalidFormat] = codes.InvalidArgument
	registeredCode[goa.ErrNameInvalidPattern] = codes.InvalidArgument
	registeredCode[goa.ErrNameInvalidRange] = codes.InvalidArgument
	registeredCode[goa.ErrNameInvalidLength] = codes.InvalidArgument
	registeredCode[goa.ErrNameMissingField] = codes.InvalidArgument
	registeredCode[goa.ErrNameMissingPayload] = codes.InvalidArgument
	registeredCode[goa.ErrNameFault] = codes.Unavailable
}

func RegisterCode(name string, code codes.Code) {
	var exists = false
	for _, n := range goa.ReservedErrNames {
		if n == name {
			exists = true
			break
		}
	}
	if !exists {
		registeredCode[name] = code
	}
}

func GetCode(name string) (codes.Code, bool) {
	value, ok := registeredCode[name]
	return value, ok
}
