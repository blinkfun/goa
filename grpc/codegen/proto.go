package codegen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/expr"
	goa "goa.design/goa/v3/pkg"
)

const (
	// ProtoVersion is the protocol buffer version used to generate .proto files
	ProtoVersion = "proto3"
)

// ProtoFiles returns a *.proto file for each gRPC service.
func ProtoFiles(genpkg string, root *expr.RootExpr) []*codegen.File {
	fw := make([]*codegen.File, len(root.API.GRPC.Services))
	for i, svc := range root.API.GRPC.Services {
		fw[i] = protoFile(genpkg, svc)
	}
	return fw
}

func protoFile(genpkg string, svc *expr.GRPCServiceExpr) *codegen.File {
	data := GRPCServices.Get(svc.Name())
	svcName := codegen.SnakeCase(data.Service.VarName)
	path := filepath.Join(codegen.Gendir, "grpc", svcName, pbPkgName, svcName+".proto")
	meta := svc.Meta
	pkg := codegen.SnakeCase(codegen.Goify(svcName, false))

	sections := []*codegen.SectionTemplate{
		// header comments
		&codegen.SectionTemplate{
			Name:   "proto-header",
			Source: protoHeaderT,
			Data: map[string]interface{}{
				"Title":       fmt.Sprintf("%s protocol buffer definition", svc.Name()),
				"ToolVersion": goa.Version(),
			},
		},
		// proto syntax and package
		&codegen.SectionTemplate{
			Name:   "proto-start",
			Source: protoStartT,
			Data: map[string]interface{}{
				"ProtoVersion": ProtoVersion,
				"Pkg":          pkg,
				"Options":      OptionsFromExpr(pkg, meta),
			},
		},
		// service definition
		&codegen.SectionTemplate{Name: "grpc-service", Source: serviceT, Data: data},
	}

	// message definition
	for _, m := range data.Messages {
		sections = append(sections, &codegen.SectionTemplate{Name: "grpc-message", Source: messageT, Data: m})
	}

	return &codegen.File{
		Path:             path,
		SectionTemplates: sections,
		FinalizeFunc:     protoc,
	}
}

func protoc(path string) error {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, 0777)

	args := []string{
		"--go_out=.",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=.",
		"--go-grpc_opt=paths=source_relative",
		"--proto_path",
		dir,
		path}
	cmd := exec.Command("protoc", args...)
	cmd.Dir = filepath.Dir(path)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to run protoc: %s: %s", err, output)
	}

	return nil
}

func OptionsFromExpr(pkg string, mdata expr.MetaExpr) map[string]string {
	return optionsFromExprWithPrefix(pkg, mdata, "rpc:option:")
}

func optionsFromExprWithPrefix(pkg string, mdata expr.MetaExpr, prefix string) map[string]string {
	if !strings.HasSuffix(prefix, ":") {
		prefix += ":"
	}
	options := make(map[string]string)
	options["go_package"] = pkg + "pb"
	for key, value := range mdata {
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		name := key[len(prefix):]
		if strings.Contains(name, ":") {
			continue
		}
		val := value[0]
		if !validOption(name, val) {
			continue
		}

		options[name] = val
	}
	return options
}

const (
	protoHeaderT = `{{ if .Title -}}
// Code generated with goa {{ .ToolVersion }}, DO NOT EDIT.
//
// {{ .Title }}
//
// Command:
{{ comment commandLine }}
{{- end }}
`

	protoStartT = `
syntax = {{ printf "%q" .ProtoVersion }};

package {{ .Pkg }};
{{ if .Options }}
	{{- range $key, $value := .Options }}
option {{ $key }} = "{{ $value }}";
	{{- end }}
{{- end }}
`

	// input: ServiceData
	serviceT = `
{{ .Description | comment }}
service {{ .Name }} {
	{{- range .Endpoints }}
	{{ if .Method.Description }}{{ .Method.Description | comment }}{{ end }}
	{{- $serverStream := or (eq .Method.StreamKind 3) (eq .Method.StreamKind 4) }}
	{{- $clientStream := or (eq .Method.StreamKind 2) (eq .Method.StreamKind 4) }}
	rpc {{ .Method.VarName }} ({{ if $clientStream }}stream {{ end }}{{ .Request.Message.VarName }}) returns ({{ if $serverStream }}stream {{ end }}{{ .Response.Message.VarName }});
	{{- end }}
}
`

	// input: service.UserTypeData
	messageT = `{{ comment .Description }}
message {{ .VarName }}{{ .Def }}
`
)
