package functions

import (
	"EvoGuard/source/config"
	"EvoGuard/source/database"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// tokens will store all the tokens which have been granted to sessions
var tokens map[string]*Token = make(map[string]*Token)

type Token struct {
	// JWT is the storage container which is locally held
	// and allows us to retrieve information about any
	// session without depending a request
	JWT *jwt.Token

	// JWTToken will be the raw string which we assign to
	// the sessions cookie, this will help authenticate
	// and confirm the identity of the said request.
	JWTToken string

	// Request is a cached iteration of the first request
	// that was received from the remote host, this can help
	// with indentifying via the IP and cookies
	Request *http.Request

	// Token is the string which we present and assign to
	// the remote session, this allows us to distinguish
	// between multiple different incoming resource requests
	Token string

	// Claimed is what we will generate and include into the JWT token headers which
	// this can also help with the authentication process of the incoming requests.
	Claimed string

	// Claims is a locally stored iteration of what we assign to the cookie, this will
	// allow us to try distinguish when cookies have been modified and at this point
	// we will try to block said request from performing any more actions
	Claims jwt.MapClaims

	// Expiry is the time we force the cookie to maintain its validity until
	// it can no longer be used and becomes obsolete
	Expiry time.Time
}

// GrantToken will attempt to grant the token to the request
func GrantToken(request *http.Request, writer http.ResponseWriter, claims jwt.MapClaims) *Token {
	EncodeSecret := hex.EncodeToString(database.Salt(32))
	claims["secret"] = hex.EncodeToString(sha512.New().Sum([]byte(EncodeSecret)))

	session := &Token{
		JWT:     jwt.NewWithClaims(jwt.SigningMethodHS256, claims),
		Token:   hex.EncodeToString(database.Salt(32)),
		Expiry:  time.Now().Add(1 * time.Hour),
		Claims:  claims,
		Request: request,
		Claimed: EncodeSecret,
	}

	var err error = nil
	session.JWTToken, err = session.JWT.SignedString([]byte(config.Options.String("website", "jwt")))
	if err != nil {
		return GrantToken(request, writer, claims)
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   session.JWTToken,
		Secure:  true,
		Expires: session.Expiry,
	}

	user, err := session.GetOwnershipUser()
	if err != nil {
		return GrantToken(request, writer, claims)
	}

	if err := database.DB.CommitLogin(user, session.Token, request); err != nil {
		return GrantToken(request, writer, claims)
	}

	http.SetCookie(writer, cookie)
	tokens[session.Token] = session
	return session
}

// FindToken attempts to index via the token inside the array
func FindToken(token string) (*Token, bool) {
	content, ok := tokens[token]
	if !ok || content == nil {
		return nil, false
	}

	if content.Expiry.Unix() <= time.Now().Unix() {
		delete(tokens, token)
		return nil, false
	}

	return content, true
}

// FindTokenViaJWT will try to find the token via the JWT token
func FindTokenViaJWT(token *jwt.Token) (*Token, bool) {
	signed, err := token.SignedString([]byte(config.Options.String("website", "jwt")))
	if err != nil {
		return nil, false
	}

	signatures, ok := token.Claims.(jwt.MapClaims)
	if !ok || signatures == nil {
		return nil, false
	}

	// secret is the cryptographic signature we assign to each session's cookie.
	secret, ok := signatures["secret"]
	if !ok || secret == nil {
		return nil, false
	}

	for _, token := range tokens {
		if token.JWTToken != signed || secret != hex.EncodeToString(sha512.New().Sum([]byte(token.Claimed))) || token.Expiry.Unix() <= time.Now().Unix() {
			if token.Expiry.Unix() <= time.Now().Unix() {
				token.RemoveToken()
				continue
			}

			continue
		}

		return token, true
	}

	return nil, false
}

// GetOwnershipUser will return the user who is associated with said cookie/token
func (token *Token) GetOwnershipUser() (*database.User, error) {
	return database.DB.GetUser(fmt.Sprint(token.Claims["username"]))
}

// RemoveToken will remove the token from the map
func (token *Token) RemoveToken() {
	delete(tokens, token.Token)
}

// Pair compares both of the stored request's IP & the current request's IP
func (token *Token) Pair(request *http.Request) bool {
	return RemovePortFromIP(token.Request.RemoteAddr) == RemovePortFromIP(request.RemoteAddr)
}
