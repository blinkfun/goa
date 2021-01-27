package grpc

import (
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

var registeredCodes = make(map[string]codes.Code)

func init() {
	registeredCodes[goa.ErrNameDecodePayload] = codes.InvalidArgument
	registeredCodes[goa.ErrNameInvalidFieldType] = codes.InvalidArgument
	registeredCodes[goa.ErrNameInvalidEnumValue] = codes.InvalidArgument
	registeredCodes[goa.ErrNameInvalidFormat] = codes.InvalidArgument
	registeredCodes[goa.ErrNameInvalidPattern] = codes.InvalidArgument
	registeredCodes[goa.ErrNameInvalidRange] = codes.InvalidArgument
	registeredCodes[goa.ErrNameInvalidLength] = codes.InvalidArgument
	registeredCodes[goa.ErrNameInvalidType] = codes.InvalidArgument
	registeredCodes[goa.ErrNameMissingField] = codes.InvalidArgument
	registeredCodes[goa.ErrNameMissingPayload] = codes.InvalidArgument
	registeredCodes[goa.ErrNameError] = codes.Internal
	registeredCodes[goa.ErrNameFault] = codes.Unavailable
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
		registeredCodes[name] = code
	}
}

func GetCode(name string) (codes.Code, bool) {
	value, ok := registeredCodes[name]
	return value, ok
}

func AllCodes() map[string]codes.Code {
	return registeredCodes
}
