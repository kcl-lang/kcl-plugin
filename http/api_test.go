package http

import (
	"encoding/json"
	"testing"

	"kcl-lang.io/kcl-go/pkg/plugin"
)

func TestHttpGet(t *testing.T) {
	resultJson := plugin.Invoke("kcl_plugin.http.get", []interface{}{"https://www.kcl-lang.io/"}, nil)
	var data map[string]any
	err := json.Unmarshal([]byte(resultJson), &data)
	if err != nil {
		t.Fatal(err)
	}
	if int(data["status"].(float64)) != 200 {
		t.Fatal(data["status"])
	}
}
