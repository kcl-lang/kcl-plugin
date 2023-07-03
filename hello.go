// Copyright 2021 The KCL Authors. All rights reserved.

//go:build ignore
// +build ignore

package main

import (
	"kcl-lang.io/kcl-plugin"
)

func main() {
	kcl_plugin.InstallPlugins("_kclvm_plugins_")
}
