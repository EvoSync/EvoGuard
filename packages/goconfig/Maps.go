

package goconfig

import (
	"encoding/json"
	"reflect"
)

// Marshal will convert the entire options map into a map, then attempt to convert into the corresponding type given
func (GC *Options) MarshalEntire(v any) error {
	return GC.toValue(GC.config.configs, v)
}

// MarshalFromPath allows you to marshal from a specific point inside the map
func (GC *Options) MarshalFromPath(v any, path ...string) error {
	item, err := GC.Get(reflect.Map, path...)
	if err != nil {
		return err
	}

	return GC.toValue(item.(map[string]any), v)
}

// Allows for the people using this package to use both the builtin get function and fill structures fields with this package.
func (GC *Options) toValue(scope map[string]any, v any) error {
	entire, err := json.Marshal(scope)
	if err != nil {
		return err
	}

	return json.Unmarshal(entire, &v)
}
