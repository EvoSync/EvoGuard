

package goconfig

import (
	"encoding/json"
	"errors"

	"github.com/BurntSushi/toml"
)


// Defaults represents all the default inclusions registered
var Defaults map[string]Inclusion = map[string]Inclusion{

	/* Json configuration files support*/
	".json": func(b []byte, p string, m map[string]any) error {
		return json.Unmarshal(b, &m)
	},

	/* Toml configuration files support */
	".toml": func(b []byte, p string, m map[string]any) error {
		return toml.Unmarshal(b, &m)
	},
}

// Inclusion allows for you to register your own functions
type Inclusion func([]byte, string, map[string]any) error

// NewInclusion attempts to register the inclusion into the support map
func (GC *GoConfig) NewInclusion(ext string, exec Inclusion) error {
	if e, ok := GC.inclusions[ext]; e != nil || ok {
		return errors.New("file type already included")
	}

	GC.inclusions[ext] = exec
	return nil
}