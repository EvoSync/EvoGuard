package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
	Adapter.go will implement the database adapters required
	for the current interface, this will allow us to constantly
	check the active state of said database.
*/

// Adapter will store all the current information about the database
type Adapter struct {
	DB	 	*sql.DB
	Config	*DBConfig
	Init	    	time.Time
}

// DB is the structure which is required
var DB *Adapter = new(Adapter)

// NewDatabaseAdapter will attempt to spawn a brand new database
func NewDatabaseAdapter() error {
	config, err := NewConfig()
	if err != nil || config == nil {
		return err
	}

	DB.Config, DB.Init = config, time.Now()
	DB.DB, err = sql.Open("sqlite3", config.Address)
	if err != nil || DB.DB == nil {
		return err
	}
	
	return DB.SchemaEngine()
}

// NoExec will return an error and ignore the argument within the exec function
func (adapter *Adapter) NoExec(str string, args ...interface{}) error {
	if _, err := adapter.DB.Exec(str, args...); err != nil {
		return err
	}

	return nil
}