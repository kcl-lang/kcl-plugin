# Copyright The KCL Authors. All rights reserved.
#
# Subpackage of plugin_examples. The host loads it with:
#   go run ./examples http
# which blank-imports kcl-lang.io/kcl-plugin/http and evaluates this file.
# The module root does not import this package, so `go run ./examples`
# does not perform the request.

import kcl_plugin.http

status = http.get("https://www.kcl-lang.io/").status
