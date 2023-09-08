package api

import (
	"EvoGuard/source/website/functions"
	"encoding/json"
	"net/http"
)

/*
	Signout.go will implement the disabling of said token, this will still keep
	the entire token within the clients browser, but removes it from our side so
	we no longer recognize there token.
*/

// Signout implements the signing out and disabling of token functionality
func Signout(response http.ResponseWriter, request *http.Request) {
	biscuit, ok := functions.ExtractJWTCrums(request)
	if !ok || biscuit == nil {
		WriteJson(map[string]any{"status": false, "error": "missing authorization"}, http.StatusUnauthorized, response)
		return
	}

	owner, err := biscuit.GetOwnershipUser()
	if err != nil || owner == nil {
		WriteJson(map[string]any{"status": false, "error": "unable to find signed user inside your token"}, http.StatusUnauthorized, response)
		return
	}

	biscuit.RemoveToken()
	contents, err := json.Marshal(map[string]any{"status": true})
	if err != nil {
		WriteJson(map[string]any{"status": true, "error": "event occurred which haulted the marshalling of response"}, http.StatusInternalServerError, response)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(contents)
}
