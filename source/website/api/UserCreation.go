package api

import (
	"EvoGuard/source/database"
	"EvoGuard/source/website/functions"
	"encoding/json"
	"net/http"
	"strconv"
)

/*
	CreateUser.go implements the api routing for
	making an user via an API endpoint. this will
	implement the creation of users
*/

// CreateUser is the routine used to actually make a user
func CreateUser(writer http.ResponseWriter, request *http.Request) {
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

	// checks if all the fields exist, if not we return an error
	if !request.URL.Query().Has("username") || !request.URL.Query().Has("password") || !request.URL.Query().Has("email") || !request.URL.Query().Has("accountlevel") {
		WriteJson(map[string]any{"status": false, "error": "missing fields for request"}, http.StatusBadRequest, writer)
		return
	}

	converted, err := strconv.Atoi(request.URL.Query().Get("accountlevel"))
	if err != nil {
		WriteJson(map[string]any{"status": false, "error": "accountLevel field must be an integer"}, http.StatusBadRequest, writer)
		return
	}

	user := &database.User{
		Username:     request.URL.Query().Get("username"),
		Password:     []byte(request.URL.Query().Get("password")),
		Email:        request.URL.Query().Get("email"),
		AccountLevel: converted,
		Parent:       owner.ID,
	}

	if err := database.DB.CreateUser(user, owner); err != nil {
		WriteJson(map[string]any{"status": false, "error": "error occurred while trying to create user"}, http.StatusInternalServerError, writer)
		return
	}

	module, err := json.Marshal(map[string]any{"status": true})
	if err != nil || module == nil {
		WriteJson(map[string]any{"status": true, "error": "event occurred which haulted the marshalling of response"}, http.StatusInternalServerError, writer)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(module)
}
