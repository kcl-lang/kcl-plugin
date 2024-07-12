package http

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"kcl-lang.io/kcl-go/pkg/plugin"
)

func getCallArgs(p *plugin.MethodArgs, key string, index int) any {
	if val, ok := p.KwArgs[key]; ok {
		return val
	}
	if index < len(p.Args) {
		return p.Args[index]
	}
	return nil
}

func init() {
	plugin.RegisterPlugin(plugin.Plugin{
		Name: "http",
		MethodMap: map[string]plugin.MethodSpec{
			"get": {
				// http.get(url)
				Body: func(args *plugin.MethodArgs) (*plugin.MethodResult, error) {
					url := fmt.Sprintf("%s", getCallArgs(args, "url", 0))
					status, body, err := fasthttp.Get(nil, url)
					if err != nil {
						return nil, err
					}
					return &plugin.MethodResult{V: map[string]any{
						"status": status,
						"body":   string(body),
					}}, nil
				},
			},
		},
	})
}
