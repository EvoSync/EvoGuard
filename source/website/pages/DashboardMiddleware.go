package pages

import (
	"EvoGuard/source/website/functions"
	"net/http"

	"github.com/gorilla/mux"
)

// DashboardMiddleware will implement the middleware required for the dashboard
func DashboardMiddleware(header http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contents, ok := functions.ExtractJWTCrums(request)
		if !ok || contents == nil || mux.Vars(request)["token"] != contents.Token || !contents.Pair(request) {
			http.Redirect(writer, request, "/", http.StatusFound)
			return
		}

		// Once authentication has been performed we redirect.
		header.ServeHTTP(writer, request)
	})
}
