package pages

import (
	"EvoGuard/source/website/functions"
	"fmt"
	"net/http"
)

// DashboardClients is the management and rendering item for the users
// management dashboard, this will allow users to modify and management
// other users
func DashboardClients(response http.ResponseWriter, request *http.Request) {
	biscuit, ok := functions.ExtractJWTCrums(request)
	if !ok || biscuit == nil {
		return
	}

	admin, err := biscuit.GetOwnershipUser()
	if err != nil || admin == nil {
		return
	}

	// Has() will check if the query has a user parameter, this
	// will mean we handle the request with a different method.
	if request.URL.Query().Has("user") {
		fmt.Println(request.URL.Query().Get("user"))
	}

	functions.WriteHypertext(response, make(map[string]string), "resources", "public", "html", "admin", "clients.html")
}
