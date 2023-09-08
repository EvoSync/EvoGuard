package config

import "EvoGuard/packages/goconfig"

/*
	Config.go will wrap the goconfig package functions
	which allows us to use within our current workspace.
*/

// Options will be used for accessing the memory being stored
var Options *goconfig.Options = new(goconfig.Options)

// OpenConfigDriver will attempt to execute GoConfig within the directories
func OpenConfigDriver() error {
	config := goconfig.NewConfig()
	config.NewInclusion(".sql", func(b []byte, s string, m map[string]any) error {
		return nil
	})
	
	err := config.Parse(directories...)
	if err != nil {
		return err
	}

	Options, err = config.Options()
	if err != nil {
		return err
	}

	return nil
}