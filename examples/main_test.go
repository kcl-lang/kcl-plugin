// Copyright The KCL Authors. All rights reserved.

//go:build cgo
// +build cgo

package main

import (
	"strings"
	"testing"

	"kcl-lang.io/kcl-go/pkg/kcl"
	"kcl-lang.io/kcl-go/pkg/plugin"
)

func TestBlankImportRegistersPlugins(t *testing.T) {
	for _, name := range []string{"hello", "http"} {
		if _, ok := plugin.GetPlugin(name); !ok {
			t.Fatalf("plugin %q is not registered; blank-import its package before kcl.Run", name)
		}
	}
	for _, method := range []string{"kcl_plugin.hello.add", "kcl_plugin.http.get"} {
		if _, ok := plugin.GetMethodSpec(method); !ok {
			t.Fatalf("method %q is not registered", method)
		}
	}
}

func TestExampleModule(t *testing.T) {
	yaml, err := run("module")
	if err != nil {
		t.Fatal(err)
	}
	got := strings.TrimSpace(yaml)
	want := "name: kcl\nthree: 3"
	if got != want {
		t.Fatalf("module result\n%s\nwant\n%s", got, want)
	}
}

func TestHTTPSubpackage(t *testing.T) {
	yaml, err := run("module/http")
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(yaml) != "status: 200" {
		t.Fatalf("http subpackage result\n%s", yaml)
	}
}

func TestUnregisteredPluginIsNotCallable(t *testing.T) {
	_, err := kcl.Run("main.k", kcl.WithCode(`
import kcl_plugin.not_loaded

value = not_loaded.add(1, 2)
`))
	if err == nil {
		t.Fatal("expected a call to an unregistered plugin to fail")
	}
	// Plugin mode is on (the import was accepted). Discovery still misses
	// the method because nothing called RegisterPlugin for "not_loaded".
	if !strings.Contains(err.Error(), "kcl_plugin.not_loaded.add") {
		t.Fatal(err)
	}
}
