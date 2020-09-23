package grpc

import (
	"google.golang.org/grpc/codes"

	goa "goa.design/goa/v3/pkg"
)

var registeredCode = make(map[string]codes.Code)

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
