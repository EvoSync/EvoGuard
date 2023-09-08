package pages

import (
	"EvoGuard/source/website/functions"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ActivateHandler is what we trigger whenever someone reaches the /activate
// endpoint, within this endpoint it requires a open cookie, if not we will
// redirect them directly to the / endpoint to login with a url query
// being our own activate point again
func ActivateHandler(response http.ResponseWriter, request *http.Request) {
	token, ok := functions.ExtractJWTCrums(request)
	if !ok || token == nil { // if no cookie is found, we redirect them to a login page
		http.Redirect(response, request, "/?url="+request.URL.String(), http.StatusFound)
		return
	}

	owner, err := token.GetOwnershipUser()
	if err != nil || owner == nil {
		// TODO: write some html pages
		return
	}

	activationID := mux.Vars(request)["id"]
	ID, err := strconv.Atoi(activationID)
	if err != nil || ID == 0 || len(activationID) == 0 {
		// TODO: write some html pages
		return
	}

	fmt.Printf("Activation endpoint reached by %s via %d\r\n", owner.Username, ID)
}
