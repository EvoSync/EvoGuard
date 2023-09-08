package functions

import (
	"EvoGuard/source/config"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// ExtractJWTCrums will extract all the JWT tokens from cookies being included
func ExtractJWTCrums(request *http.Request) (*Token, bool) {
	biscuit, err := request.Cookie("token")
	if err != nil {
		return nil, false
	}

	/* Tries to parse the JWT value */
	token, err := jwt.Parse(biscuit.Value, func(recv *jwt.Token) (interface{}, error) {
		if _, ok := recv.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.Options.String("website", "jwt")), nil
	})

	/* Tries to find the current token */
	if err != nil || token == nil {
		return nil, false
	}

	return FindTokenViaJWT(token)
}
