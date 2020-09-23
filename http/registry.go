package http

import (
	"net/http"

	"goa.design/goa/v3/pkg"
)

var registeredCode = make(map[string]int)

func init() {
	RegisterCode(goa.ErrNameMissingPayload, http.StatusBadRequest)
	RegisterCode(goa.ErrNameDecodePayload, http.StatusBadRequest)
	RegisterCode(goa.ErrNameInvalidFieldType, http.StatusBadRequest)
	RegisterCode(goa.ErrNameMissingField, http.StatusBadRequest)
	RegisterCode(goa.ErrNameInvalidEnumValue, http.StatusBadRequest)
	RegisterCode(goa.ErrNameInvalidFormat, http.StatusBadRequest)
	RegisterCode(goa.ErrNameInvalidPattern, http.StatusBadRequest)
	RegisterCode(goa.ErrNameInvalidRange, http.StatusBadRequest)
	RegisterCode(goa.ErrNameInvalidLength, http.StatusBadRequest)
	RegisterCode(goa.ErrNameFault, http.StatusServiceUnavailable)
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
