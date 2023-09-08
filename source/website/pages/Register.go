package pages

import (
	"EvoGuard/source/database"
	"EvoGuard/source/website/functions"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// RegisterHandler will implement the functionality of registering accounts
func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	if token, ok := functions.ExtractJWTCrums(request); ok && token != nil {
		if token := request.URL.Query().Get("url"); len(token) >= 1 {
			http.Redirect(response, request, request.URL.Query().Get("url"), http.StatusFound)
			return
		}

		http.Redirect(response, request, fmt.Sprintf("/%s/dashboard/", token.Token), http.StatusFound)
		return
	}

	/* Checks the request method for a post req */
	if strings.EqualFold(request.Method, "post") {
		err := request.ParseForm()
		if err != nil {
			return
		}

		/* Checks if the username already exists, stops duplicate usernames */
		account, err := database.DB.GetUser(request.Form.Get("Username"))
		if account != nil || err == nil {
			return
		}

		/* Checks if the email already exists, stop duplicate emails */
		account, err = database.DB.GetByEmail(request.Form.Get("Email"))
		if account != nil || err == nil {
			return
		}

		parent, err := database.DB.GetByID(1)
		if err != nil || parent == nil {
			return
		}

		register := &database.User{
			Email:        request.Form.Get("Email"),
			Username:     request.Form.Get("Username"),
			Password:     []byte(request.Form.Get("Password")),
			AccountLevel: 1,
			Parent:       parent.ID,
		}

		// Tries to create the user inside the database, meaning we can continue
		if err := database.DB.CreateUser(register, parent); err == nil {
			token := functions.GrantToken(request, response, jwt.MapClaims{"username": register.Username})
			if token := request.URL.Query().Get("url"); len(token) >= 1 {
				http.Redirect(response, request, request.URL.Query().Get("url"), http.StatusFound)
				return
			}

			http.Redirect(response, request, fmt.Sprintf("/%s/dashboard/", token.Token), http.StatusFound)
			return
		}
	}

	// Writes the actual hypertext contents
	functions.WriteHypertext(response, make(map[string]string), "resources", "public", "html", "register.html")
}
