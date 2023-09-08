package api

import (
	"EvoGuard/source/database"
	"EvoGuard/source/website/functions"
	"encoding/json"
	"net/http"
)

/*
	GetUsers.go is the api routing which implements the required functionality
	for fetching all the children of said user. this will mean they can only
	view users which they have created.
*/

// GetUsers is the api route for returning all of the users children
func GetUsers(writer http.ResponseWriter, request *http.Request) {
	biscuit, ok := functions.ExtractJWTCrums(request)
	if !ok || biscuit == nil {
		WriteJson(map[string]any{"status": false, "error": "missing authorization"}, http.StatusUnauthorized, writer)
		return
	}

	owner, err := biscuit.GetOwnershipUser()
	if err != nil || owner == nil {
		WriteJson(map[string]any{"status": false, "error": "unable to find signed user inside your token"}, http.StatusUnauthorized, writer)
		return
	}

	children, err := database.DB.GetChildren(owner)
	if err != nil || children == nil {
		WriteJson(map[string]any{"status": false, "error": "error occurred which stopped us from fetching your users"}, http.StatusInternalServerError, writer)
		return
	}

	contents, err := json.Marshal(children)
	if err != nil || contents == nil {
		WriteJson(map[string]any{"status": true, "error": "event occurred which haulted the marshalling of response"}, http.StatusInternalServerError, writer)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(contents)
}
