package http

import (
	"net/http"

	goa "goa.design/goa/v3/pkg"
)

var (
	registeredCodes = make(map[string]int)
)

func init() {
	registeredCodes[goa.ErrNameDecodePayload] = http.StatusBadRequest
	registeredCodes[goa.ErrNameInvalidFieldType] = http.StatusBadRequest
	registeredCodes[goa.ErrNameInvalidEnumValue] = http.StatusBadRequest
	registeredCodes[goa.ErrNameInvalidFormat] = http.StatusBadRequest
	registeredCodes[goa.ErrNameInvalidPattern] = http.StatusBadRequest
	registeredCodes[goa.ErrNameInvalidRange] = http.StatusBadRequest
	registeredCodes[goa.ErrNameInvalidLength] = http.StatusBadRequest
	registeredCodes[goa.ErrNameInvalidType] = http.StatusBadRequest
	registeredCodes[goa.ErrNameMissingField] = http.StatusBadRequest
	registeredCodes[goa.ErrNameMissingPayload] = http.StatusBadRequest
	registeredCodes[goa.ErrNameError] = http.StatusInternalServerError
	registeredCodes[goa.ErrNameFault] = http.StatusServiceUnavailable
}

func RegisterCode(name string, code int) {
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

func GetCode(name string) (int, bool) {
	value, ok := registeredCodes[name]
	return value, ok
}

func AllCodes() map[string]int {
	return registeredCodes
}
