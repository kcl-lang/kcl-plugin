# Plugin example module

`module/` is a KCL module. `main.go` is the host that loads the plugins in this repository and evaluates the module.

KCL discovers a plugin from that host, not from the module path:

1. The blank imports in `main.go` run `init` in [`hello`](../hello/api.go) and [`http`](../http/api.go). Each `plugin.RegisterPlugin` call stores methods under `kcl_plugin.<name>.<method>`.
2. `kcl.Run` passes `plugin.GetInvokeJsonProxyPtr()` to the runtime. The pointer is non-zero when CGO is enabled, and the runner then accepts `import kcl_plugin.<name>`.
3. `module/main.k` calls `hello.add`. The runtime sends the method name `kcl_plugin.hello.add` to the agent, which looks it up in the registry from step 1.

`module/http/main.k` is the same path for `kcl_plugin.http.get`. It is a subpackage, so evaluating the module root does not call it.

## Run

From the repository root, with CGO enabled:

```shell
go run ./examples
go run ./examples http
```

The first command prints:

```yaml
name: kcl
three: 3
```

The second command fetches `https://www.kcl-lang.io/` and prints the HTTP status.

`go test ./examples` checks that both plugins are registered, that the module evaluates to `three: 3`, that the http subpackage evaluates to `status: 200`, and that a plugin the host did not import cannot be called.
