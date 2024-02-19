// Copyright 2021 The KCL Authors. All rights reserved.

// kclvm plugins.
package kcl_plugin

import (
	"embed"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed README.md
//go:embed hello
//go:embed project_context
//go:embed http
var PluginFS embed.FS

func InstallPlugins(root string) error {
	embedFS := PluginFS
	err := os.MkdirAll(root, 0777)
	// If permission denied, ignore it.
	if errors.Is(err, fs.ErrPermission) {
		return nil
	}
	if err != nil {
		return err
	}
	return fs.WalkDir(embedFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		abspath := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(abspath), 0777); err != nil {
			_ = err
		}
		data, err := fs.ReadFile(embedFS, path)
		if err != nil {
			return err
		}
		err = os.WriteFile(abspath, data, 0777)
		if err != nil {
			return err
		}
		return nil
	})
}
