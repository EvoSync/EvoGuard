package pages

import (
	"EvoGuard/source/website/functions"
	"net/http"
)

// Dashboard home page is represented within this function
func Dashboard(writer http.ResponseWriter, request *http.Request) {
	biscuit, ok := functions.ExtractJWTCrums(request)
	if !ok || biscuit == nil {
		http.Redirect(writer, request, "/", http.StatusFound)
		return
	}

	/* Tries to get the username from the local package */
	account, err := biscuit.GetOwnershipUser()
	if err != nil || account == nil {
		http.Redirect(writer, request, "/", http.StatusFound)
		return
	}

	switch account.AccountLevel {

	case 0: // When the account is admin
		AdminDashboard(writer, request, account)
		return

	default: // When the account level is invalid
		writer.Write([]byte("Invalid account level"))
		return
	}

}
