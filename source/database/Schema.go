package database

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/*
	Schema.go will implement the required functions for handling
	and inserting each unquie table on startup, this will mean
	that when the program is spawned. we will automatically
	try to insert said bounds
*/

// SchemaEngine will attempt to scan the schema dirctory for all schemas within it
func (adapter *Adapter) SchemaEngine() error {
	return filepath.Walk(adapter.Config.Schema, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil || filepath.Ext(path) != ".sql" {
			return err
		}

		contents, err := os.ReadFile(path)
		if err != nil || contents == nil {
			return err
		}

		name := strings.Split(string(contents), "`")[1:2][0]
		query, err := adapter.DB.Query(fmt.Sprintf("SELECT * FROM %s", name))
		if err == nil && query != nil {
			return query.Close()
		}

		if err := adapter.NoExec(string(contents)); err != nil {
			return err
		}

		switch strings.ToLower(name) {

		// When users table created introduce some new values
		case "users":
			return adapter.DefaultUser()

		default:
			return nil
		}
	})
}