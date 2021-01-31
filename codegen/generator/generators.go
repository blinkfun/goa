package generator

import (
	"context"
	"fmt"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

// Genfunc is the type of the functions invoked to generate code.
type Genfunc func(ctx context.Context, genpkg string, roots []eval.Root) ([]*codegen.File, error)

// Generators returns the qualified paths (including the package name) to the
// code generator functions for the given command, an error if the command is
// not supported. Generators is a public variable so that external code (e.g.
// plugins) may override the default generators.
var Generators = generators

// generators returns the generator functions exposed by the generator package
// for the given command.
func generators(cmd string) ([]Genfunc, error) {
	switch cmd {
	case "gen":
		return []Genfunc{Service, Transport, OpenAPI}, nil
	case "gen-client":
		return []Genfunc{Service, ClientTransport}, nil
	case "gen-server":
		return []Genfunc{Service, ServerTransport, OpenAPI}, nil
	case "example":
		return []Genfunc{Example}, nil
	default:
		return nil, fmt.Errorf("unknown command %q", cmd)
	}
}
