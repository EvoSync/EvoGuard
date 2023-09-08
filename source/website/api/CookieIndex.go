package api

import (
	"EvoGuard/source/website/functions"
	"encoding/json"
	"net/http"
)

/*
	Myself will return information about the cookies owner like
	username, accountLevel etc. this will help us to determine
	certain information
*/

func Myself(writer http.ResponseWriter, request *http.Request) {
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

	contents, err := json.Marshal(owner)
	if err != nil || contents == nil {
		WriteJson(map[string]any{"status": false, "error": "event occurred which haulted the marshalling of response"}, http.StatusInternalServerError, writer)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(contents)
}
