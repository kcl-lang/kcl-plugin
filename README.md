# KCL Plugin

KCL calls host functions through plugins. This repository contains the `hello` and `http` plugins. A KCL program imports one with `import kcl_plugin.<name>` and calls its methods by that name.

Language reference: [Plugin overview](https://kcl-lang.io/docs/reference/plugin/overview/).

## Load a plugin

Registering a plugin in Go is half of the setup. KCL loads it only when the process that evaluates the program installs a plugin agent and that process has registered the plugin's methods.

The `kcl` CLI leaves the agent pointer at zero. Running a module that imports a plugin then fails with:

```text
the plugin package `kcl_plugin.hello` is not found, please confirm if plugin mode is enabled
```

A Go host fixes that in two steps.

1. **Register the methods.** Blank-import the plugin package. The import runs `init`, and `plugin.RegisterPlugin` publishes every method as `kcl_plugin.<name>.<method>`. The `Name` field is the name used in the KCL import (`hello` becomes `import kcl_plugin.hello`).

2. **Evaluate through the Go SDK.** `kcl.Run` and `kcl.MustRun` (`kcl-lang.io/kcl-go`) build the native service client. That client passes `plugin.GetInvokeJsonProxyPtr()` into the runtime. The runner sets `load_plugins` when this pointer is non-zero, which is what makes `import kcl_plugin.<name>` resolve. CGO is required. A build with `CGO_ENABLED=0` reports a zero pointer, and the import fails with the message above.

```go
import (
    "fmt"

    "kcl-lang.io/kcl-go/pkg/kcl"

    _ "kcl-lang.io/kcl-plugin/hello"
    _ "kcl-lang.io/kcl-plugin/http"
)

func main() {
    // examples/module is a KCL module whose main.k imports kcl_plugin.hello.
    yaml := kcl.MustRun("examples/module").GetRawYamlResult()
    fmt.Println(yaml)
}
```

KCL does not scan a plugin directory and does not read a `kcl.mod` dependency to find these plugins. `import kcl_plugin.hello` resolves because plugin mode is on. The call `hello.add(1, 2)` is dispatched to the registered method `kcl_plugin.hello.add`. A call whose package was never blank-imported fails at runtime, because the agent has no method under that name.

`examples/` is a complete module that uses both plugins in this repository, plus the host that loads them:

```shell
go run ./examples
```

```yaml
name: kcl
three: 3
```

`go run ./examples http` runs the `http` subpackage (`http.get`).

To use the same module from a command line tool, build that tool the same way: blank-import the plugin packages and call the Go SDK. The stock `kcl` binary has no agent, so it cannot see plugins registered in another process.

Python hosts follow the same split. Call `kcl_lib.plugin.register_plugin` before `exec_program`, then import `kcl_plugin.<name>` from KCL. `hello/plugin.py` and `http/plugin.py` are those Python implementations. The Go host above does not load them.

## Plugins

| Directory | KCL import | Blank import |
| --- | --- | --- |
| [`hello`](hello/api.go) | `kcl_plugin.hello` | `kcl-lang.io/kcl-plugin/hello` |
| [`http`](http/api.go) | `kcl_plugin.http` | `kcl-lang.io/kcl-plugin/http` |
