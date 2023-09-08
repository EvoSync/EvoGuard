package pages

import (
	"EvoGuard/source/database"
	"EvoGuard/source/website/functions"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

/*
	Login.go is the listened and watched internal package for the
	actual server runtime, this will maintain all the required
	information which handles the login process
*/

// Login implements the actual login interactions with the remote node
func Login(writer http.ResponseWriter, request *http.Request) {
	if token, ok := functions.ExtractJWTCrums(request); ok && token != nil {
		if token := request.URL.Query().Get("url"); len(token) >= 1 {
			http.Redirect(writer, request, request.URL.Query().Get("url"), http.StatusFound)
			return
		}

		http.Redirect(writer, request, fmt.Sprintf("/%s/dashboard/", token.Token), http.StatusFound)
		return
	}

	if strings.ToLower(request.Method) == "post" {
		err := request.ParseForm()
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		/* Gets both the username & password */
		username := request.Form.Get("Username")
		password := request.Form.Get("Password")

		user, err := database.DB.GetUser(username)
		if err != nil {
			functions.WriteHypertext(writer, make(map[string]string), "resources", "public", "html", "login.html")
			return
		}

		/* Attempts to perform the password comparison */
		if database.MatchHash(user.Password, []byte(password), user.Salt) {
			token := functions.GrantToken(request, writer, jwt.MapClaims{"username": user.Username})
			if token := request.URL.Query().Get("url"); len(token) >= 1 {
				http.Redirect(writer, request, request.URL.Query().Get("url"), http.StatusFound)
				return
			}

			http.Redirect(writer, request, fmt.Sprintf("/%s/dashboard/", token.Token), http.StatusFound)
			return
		}

		functions.WriteHypertext(writer, make(map[string]string), "resources", "public", "html", "login.html")
		return
	}

	// WriteHypertext will execute said hyper interface onto the terminal
	functions.WriteHypertext(writer, make(map[string]string), "resources", "public", "html", "login.html")
}
