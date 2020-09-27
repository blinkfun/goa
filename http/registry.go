package http

import (
	"net/http"

	goa "goa.design/goa/v3/pkg"
)

var registeredCode = make(map[string]int)

func init() {
	registeredCode[goa.ErrNameDecodePayload] = http.StatusBadRequest
	registeredCode[goa.ErrNameInvalidFieldType] = http.StatusBadRequest
	registeredCode[goa.ErrNameInvalidEnumValue] = http.StatusBadRequest
	registeredCode[goa.ErrNameInvalidFormat] = http.StatusBadRequest
	registeredCode[goa.ErrNameInvalidPattern] = http.StatusBadRequest
	registeredCode[goa.ErrNameInvalidRange] = http.StatusBadRequest
	registeredCode[goa.ErrNameInvalidLength] = http.StatusBadRequest
	registeredCode[goa.ErrNameInvalidType] = http.StatusBadRequest
	registeredCode[goa.ErrNameMissingField] = http.StatusBadRequest
	registeredCode[goa.ErrNameMissingPayload] = http.StatusBadRequest
	registeredCode[goa.ErrNameError] = http.StatusInternalServerError
	registeredCode[goa.ErrNameFault] = http.StatusServiceUnavailable
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
		registeredCode[name] = code
	}
}

func GetCode(name string) (int, bool) {
	value, ok := registeredCode[name]
	return value, ok
}
