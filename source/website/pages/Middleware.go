package pages

import "net/http"

// MiddlewareFunc is what continues the middleware chain within the source
func MiddlewareFunc(header http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !AddRequest(request) {
			writer.Write([]byte("Ratelimited"))
			return
		}

		/*
			Behond this point we will do some basic checks to confirm
			if we want to handle this request with authenticating it
			or not, we will decide whether we will if it's requesting
			a certain path or not, for example anything containing
			login or register will not need authenticating.
		*/

		header.ServeHTTP(writer, request)
	})
}
