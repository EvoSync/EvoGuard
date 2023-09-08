package api

import (
	"EvoGuard/source/website/functions"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// MiddlewareFunc is the interaction and handshake which happens to confirm the request can happen
// within the auth guidelines
func MiddlewareFunc(header http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contents, ok := functions.ExtractJWTCrums(request)
		if !ok || contents == nil || mux.Vars(request)["token"] != contents.Token || !contents.Pair(request) {
			UnauthorizedRequest(writer, request, "missing correct authorization, receive this before continuing again")
			return
		}

		// checks if the cookie contains a valid user
		biscuit, err := contents.GetOwnershipUser()
		if err != nil || biscuit == nil {
			UnauthorizedRequest(writer, request, "missing correct authorization, receive this before continuing again")
			return
		}

		header.ServeHTTP(writer, request)
	})
}

// WriteJson writes the JSON representation to the given writer
func WriteJson(elements any, code int, writer http.ResponseWriter) error {
	components, err := json.Marshal(elements)
	if err != nil {
		return err
	}

	writer.WriteHeader(code)
	return functions.ExportLast(writer.Write(components))
}
