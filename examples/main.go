// Copyright The KCL Authors. All rights reserved.

//go:build cgo
// +build cgo

// Command examples loads the hello and http plugins and evaluates the KCL
// module in ./module.
//
// Blank-importing a plugin package is what registers it. kcl.Run is what
// turns plugin mode on: the native client passes plugin.GetInvokeJsonProxyPtr()
// to the runtime, and a non-zero pointer makes "import kcl_plugin.<name>"
// resolve. The kcl CLI uses a zero pointer, so the same import fails there.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"kcl-lang.io/kcl-go/pkg/kcl"

	// Registration. Each init calls plugin.RegisterPlugin. Without these
	// imports the module still parses (plugin mode is on), but hello.add
	// fails because kcl_plugin.hello.add was never published.
	_ "kcl-lang.io/kcl-plugin/hello"
	_ "kcl-lang.io/kcl-plugin/http"
)

func main() {
	target := "module"
	if len(os.Args) > 1 {
		target = filepath.Join("module", os.Args[1])
	}
	yaml, err := run(target)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(yaml)
	if len(yaml) == 0 || yaml[len(yaml)-1] != '\n' {
		fmt.Println()
	}
}

// run evaluates a path relative to this file. go run and go test can use any
// working directory and still find the module.
func run(rel string) (string, error) {
	result, err := kcl.Run(filepath.Join(sourceDir(), rel))
	if err != nil {
		return "", err
	}
	return result.GetRawYamlResult(), nil
}

func sourceDir() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "."
	}
	return filepath.Dir(file)
}
