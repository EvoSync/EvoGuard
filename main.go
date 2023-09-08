package main

import (
	"EvoGuard/source/config"
	"EvoGuard/source/database"
	"EvoGuard/source/website"
)

// main function for the application
func main() {
	config.Log("\x1b[38;5;10;1mInitializing\x1b[0m EvoGuard \x1b[3m%s\x1b[0m...", config.Version)

	// OpenConfigDriver will execute goconfigs required functions
	if err := config.OpenConfigDriver(); err != nil {
		config.Log("\x1b[38;5;9;1mShutting down\x1b[0m because of error: %v", err)
		return
	}

	config.Log("\x1b[38;5;10;1mSuccessfully\x1b[0m opened and scanned \x1b[38;5;13;3;1m%d\x1b[0m scopes", len(config.Options.Back().Files()))
	
	// NewDatabaseAdapter will try to connect and configure the database
	if err := database.NewDatabaseAdapter(); err != nil {
		config.Log("\x1b[38;5;9;1mShutting down\x1b[0m because of error: %v", err)
		return
	}

	config.Log("\x1b[38;5;10;1mSuccessfully\x1b[0m opened the database and profiled it: \"\x1b[0;38;5;13;3;1m%s\x1b[0m\"", database.DB.Config.Address)

	// NewHandler will try to spawn a new HTTPs server
	if err := website.NewHandler(); err != nil {
		config.Log("\x1b[38;5;9;1mShutting down\x1b[0m because of error: %v", err)
		return
	}
}