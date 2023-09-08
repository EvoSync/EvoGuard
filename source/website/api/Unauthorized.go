package api

import (
	"encoding/json"
	"net/http"
)

// Error is a universally understood response for invalid authentication attempts.
type Error struct {
	Status bool   `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

// UnauthorizedRequest is the function which gets the hand off from when requests aren't complete
func UnauthorizedRequest(writer http.ResponseWriter, request *http.Request, message string) {
	errors := Error{
		Status: false,
		Error:  message,
	}

	contents, err := json.Marshal(errors)
	if err != nil || contents == nil {
		UnauthorizedRequest(writer, request, message)
		return
	}

	writer.WriteHeader(http.StatusUnauthorized)
	writer.Write([]byte(contents))
}
