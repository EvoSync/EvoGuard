package config

/*
	Constants.go will include all the required constants of
	information which aren't allowed to be changed, these
	are required to be static within runtime.
*/

// Version is the current application state
const Version string = "v1.0"

// directories is a list of directorys which we will index within with goconfig
var directories []string = []string{"resources"}