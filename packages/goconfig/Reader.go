

package goconfig

import (
	"io/fs"
	"os"
	"path/filepath"
)

// execDir will attempt to scan an entire directory for any extensions inside the inclusion map
func (GC *GoConfig) execDir(dir string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		/* Checks for options inside that config */
		config, ok := GC.inclusions[filepath.Ext(info.Name())]
		if !ok || config == nil {
			return nil 
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		GC.renders[path] = contents
		return config(contents, path, GC.configs)
	})
}