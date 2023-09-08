package database

import "net/http"

/*
	Logins.go will implement the functions required for
	consuming logins and queries them into the database
	which allows us to index for them in the future.
*/

// Login is the formula we use when trying to insert a user into the database
type Login struct {
	Email			string `json:"email"`
	Token			string `json:"token"`
	Username		string `json:"username"`
	RemoteAddress	string `json:"remote_address"`
}

// GetUser will attempt to return the user from the database
func (adapter *Adapter) LoginUser(login *Login) (*User, error) {
	return adapter.GetUser(login.Username)
}

// CommitLogin will attempt to commit said login into the database
func (adapter *Adapter) CommitLogin(user *User, token string, request *http.Request) error {
	return adapter.NoExec("INSERT INTO `logins` (`email`, `token`, `username`, `remoteAddress`) VALUES (?, ?, ?, ?)", user.Email, token, user.Username, request.RemoteAddr)
} 