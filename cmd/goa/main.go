package main

import (
	"fmt"
	"go/build"
	"os"
	"strings"

	"flag"

	"goa.design/goa/v3/codegen"
	goa "goa.design/goa/v3/pkg"
)

func main() {
	var (
		cmd    string
		path   string
		offset int
	)
	{
		if len(os.Args) == 1 {
			usage()
		}

		switch os.Args[1] {
		case "version":
			fmt.Println("Goa version " + goa.Version())
			os.Exit(0)
		case "gen", "example", "gen-server":
			if len(os.Args) == 2 {
				usage()
			}
			cmd = os.Args[1]
			path = os.Args[2]
			offset = 2
		default:
			usage()
		}
	}

	var (
		output = "."
		gendir = ""
		debug  bool
	)
	if len(os.Args) > offset+1 {
		var (
			fset = flag.NewFlagSet("default", flag.ExitOnError)
			o    = fset.String("o", "", "output `directory`")
			out  = fset.String("output", output, "output `directory`")
			//
			g   = fset.String("g", "", "gen directory")
			gen = fset.String("gen", gendir, "gen directory")
		)
		fset.BoolVar(&debug, "debug", false, "Print debug information")

		fset.Usage = usage
		fset.Parse(os.Args[offset+1:])

		output = *o
		if output == "" {
			output = *out
		}

		gendir = *g
		if gendir == "" {
			gendir = *gen
		}
	}

	gen(cmd, path, output, gendir, debug)
}

// help with tests
var (
	usage = help
	gen   = generate
)

func generate(cmd, path, output, gendir string, debug bool) {
	var (
		files []string
		err   error
		tmp   *Generator
	)

	if len(gendir) > 0 {
		codegen.Gendir = gendir
		fmt.Printf("The gen directory is: %s\n", gendir)
	}

	if _, err = build.Import(path, ".", 0); err != nil {
		goto fail
	}

	tmp = NewGenerator(cmd, path, output)
	if !debug {
		defer tmp.Remove()
	}

	if err = tmp.Write(debug); err != nil {
		goto fail
	}

	if err = tmp.Compile(); err != nil {
		goto fail
	}

	if files, err = tmp.Run(); err != nil {
		goto fail
	}

	fmt.Println(strings.Join(files, "\n"))
	return
fail:
	fmt.Fprintln(os.Stderr, err.Error())
	if !debug && tmp != nil {
		tmp.Remove()
	}
	os.Exit(1)
}

func help() {
	fmt.Fprint(os.Stderr, `goa is the code generation tool for the Goa framework.
Learn more at https://goa.design.

Usage:
  goa gen PACKAGE [--output DIRECTORY] [--debug]
  goa example PACKAGE [--output DIRECTORY] [--debug]
  goa version

Commands:
  gen
        Generate service interfaces, endpoints, transport code and OpenAPI spec.
  gen-server
        Generate service interfaces, server's endpoints, server's transport code only and OpenAPI spec.
  example
        Generate example server and client tool.
  version
        Print version information.

Args:
  PACKAGE
        Go import path to design package

Flags:
  -o, -output DIRECTORY
        output directory, defaults to the current working directory
  -g, -gen DIRECTORY
        gen directory, defaults is gen
  -debug
        Print debug information (mainly intended for Goa developers)

Example:

  goa gen goa.design/examples/cellar/design -o gendir -g generated

`)
	os.Exit(1)
}
