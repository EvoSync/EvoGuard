
package goconfig

import (
	"errors"
	"reflect"
)

type Options struct {
	config *GoConfig
}

func (O *Options) Back() *GoConfig {
	return O.config
}

// Options creates a new options window from the config structure
func (GC *GoConfig) Options() (*Options, error) {
	if GC.configs == nil || len(GC.configs) <= 0 {
		return nil, errors.New("must execute Parse() once with targets before calling this function")
	}

	return &Options{
		config: GC,
	}, nil
}

// Get will attempt to retrive the value from the configuration options
func (O *Options) Get(t reflect.Kind, p ...string) (any, error) {
	scope := O.config.configs

	for i, element := range p {
		switch reflect.ValueOf(scope[element]).Kind() {

		case t: // object located
			if i != len(p) - 1 {
				if t == reflect.Map {
					scope = scope[element].(map[string]any)
					continue
				}

				return nil, errors.New("item appears before suggested")
			}

			return scope[element], nil

		case reflect.Map: // new scope with same name as path
			scope = scope[element].(map[string]any)

		default: // lost trace
			return nil, errors.New("bad item")
		}
	}

	return nil, errors.New("item not found")
}

// GetFromVectors will attempt to find all the possible matches from the multiply choices
func (O *Options) GetFromVectors(vectors []reflect.Kind, p ...string) (any, error) {
	for _, vector := range vectors {
		if item, err := O.Get(vector, p...); err == nil {
			return item, nil
		}
	}

	return nil, errors.New("item not found")
}