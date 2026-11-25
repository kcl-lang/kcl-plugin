# Copyright The KCL Authors. All rights reserved.
#
# Module entry. hello.add runs only when the host in examples/main.go has
# blank-imported kcl-lang.io/kcl-plugin/hello and evaluates this module with
# the Go SDK. See examples/README.md.

import kcl_plugin.hello

name = "kcl"
three = hello.add(1, 2)
