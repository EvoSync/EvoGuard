package database

import "errors"

// CreateUser will insert the user into the database
func (adapter *Adapter) CreateUser(user *User, parent *User) error {
	exists, err := adapter.GetUser(user.Username)
	if err == nil && exists != nil{
		return errors.New("user already exists")
	}

	user.Salt = Salt(32)
	user.Password = Hash(user.Password, user.Salt)
	return adapter.NoExec("INSERT INTO `users` (`username`, `password`, `email`, `salt`, `accountLevel`, `parent`) VALUES (?, ?, ?, ?, ?, ?)", user.Username, user.Password, user.Email, user.Salt, user.AccountLevel, parent.ID)
}

// DefaultUser will create the default user into the database
func (adapter *Adapter) DefaultUser() error {
	var user = &User{
		Email: "system@evoguard.cc",
		Username: "root",
		Password: []byte("root"),
		AccountLevel: 0,
	}

	// The parent of the orginal user is always it's self
	return adapter.CreateUser(user, &User{ID: 1})
}

// RemoveUser will attempt to remove a user from the database
// To execute the RemoveUser operation you are required to
// include a yourself parameter, which will be used and authenticated
// to check if you can remove that user from inside the database.
func (adapter *Adapter) RemoveUser(self *User, user string) error {
	account, err := adapter.GetUser(user)
	if err != nil || account == nil {
		return errors.New("unable to find user")
	}

	if adapter.IsChild(account, self) {
		return errors.New("you are this users child")
	}

	return adapter.NoExec("DELETE FROM `users` WHERE `username` = ?", account.Username)
}

// GetParent will try find the user via there assigned user parent account
// which means it will return there entire information structure and helps
// with management of permissions and modification rights.
func (adapter *Adapter) GetParent(self *User) (*User, error) {
	return adapter.GetByID(self.Parent)
}

// IsChild is an error safe function, which will only return false when it absolutly 
// can be false, meaning any errors which occur during function runtime, only if the
// user is not there child, it won't cause any other issues only when there can be one
func (adapter *Adapter) IsChild(parent *User, self *User) bool {
	childs, err := adapter.GetChildren(parent)
	if err != nil || childs == nil {
		return true
	}

	// Checks if the self user is the child of the parent user
	for _, child := range childs {
		if child.ID != self.ID {
			continue
		}

		return true
	}
	
	return false
}

// FamilyTree will take the lowest scoring ancestor and return find
// all the parents, granparents etc of that ancestor, this means at
// the end of this function, we will have the entire family tree of
// a user, including there parents, parents of parents etc...
func (adapter *Adapter) FamilyTree(self *User) ([]*User, error) {
	destination := append(make([]*User, 0), self)

	for {
		ancestor, err := adapter.GetByID(destination[len(destination) - 1].Parent)
		if err != nil || ancestor == nil {
			return destination, nil
		}

		destination = append(destination, ancestor)
	}
}