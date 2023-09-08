package database

import "EvoGuard/source/config"

/*
	Config.go will maintain the information for accessing the
	current state of the database. this allows us to configure
	the database for different envs.
*/

// DBConfig represents the current information of the configuration structure
type DBConfig struct {
	Address string `json:"address"`
	Schema  string `json:"schema"`
}

// Marshals into the destination structure
func NewDBConfig(dest *DBConfig) error {
	return config.Options.MarshalFromPath(dest, "database")
} 

// NewConfig will attempt to create a new database configuration structure
func NewConfig() (*DBConfig, error) {
	destination := new(DBConfig)
	err := NewDBConfig(destination)
	if err != nil {
		return nil, err
	}

	return destination, nil
}