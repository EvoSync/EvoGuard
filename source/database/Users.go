package database

// User is the structure which allows us to interact with said user interface easily.
type User struct {
	ID 			int     `json:"id"`
	Username		string	`json:"username"`
	Password		[]byte	`json:"password"`
	Email			string  `json:"email"`
	Salt			[]byte	`json:"salt"`
	AccountLevel		int	`json:"account_level"`
	Parent        		int     `json:"parent"`
}

// GetUser will scan from the database into this structure
func (adapter *Adapter) GetUser(username string) (*User, error) {
 	rows, err := adapter.DB.Query("SELECT `id`, `username`, `password`, `email`, `salt`, `accountLevel`, `parent` FROM `users` WHERE `username` = ?", username)
 	if err != nil || rows == nil {
		return nil, err
	}

	rows.Next()
	defer rows.Close()
	return adapter.scanUser(rows)
}

// GetUsers will return all the users found inside the database
func (adapter *Adapter) GetUsers() ([]*User, error) {
	rows, err := adapter.DB.Query("SELECT `id`, `username`, `password`, `email`, `salt`, `accountLevel`, `parent` FROM `users`")
	if err != nil || rows == nil {
		return nil, err
	}

	destination := make([]*User, 0)

	defer rows.Close()
	for rows.Next() {
		user, err := adapter.scanUser(rows)
		if err != nil || user == nil {
			return nil, err
		}
		
		destination = append(destination, user)
	}
	
	return destination, nil
}


// GetChildren will return all the users created and all the users they have created, it's a recursive function
func (adapter *Adapter) GetChildren(self *User) ([]*User, error) {
	rows, err := adapter.DB.Query("SELECT `id`, `username`, `password`, `email`, `salt`, `accountLevel`, `parent` FROM `users` WHERE `parent` = ?", self.ID)
	if err != nil || rows == nil {
		return nil, err
	}

	defer rows.Close()
	destination := make([]*User, 0)

	for rows.Next() {
		index, err := adapter.scanUser(rows)
		if err != nil || index == nil {
			return nil, err
		}

		destination = append(destination, index)
		if index.Username == self.Username {
			continue
		}

		thereChildren, err := adapter.GetChildren(index)
		if err != nil || thereChildren == nil {
			return nil, err
		}

		destination = append(destination, thereChildren...)
	}

	return destination, nil 
}

// GetByID will find the user who has the specified ID from the database.
func (adapter *Adapter) GetByID(id int) (*User, error) {
	rows, err := adapter.DB.Query("SELECT `id`, `username`, `password`, `email`, `salt`, `accountLevel`, `parent` FROM `users` WHERE `id` = ?", id)
	if err != nil || rows == nil {
		return nil, err
       	}

       	rows.Next()
       	defer rows.Close()
       	return adapter.scanUser(rows)
}

// GetByEmail will try to find the user via there email address
func (adapter *Adapter) GetByEmail(email string) (*User, error) {
	rows, err := adapter.DB.Query("SELECT `id`, `username`, `password`, `email`, `salt`, `accountLevel`, `parent` FROM `users` WHERE `email` = ?", email)
	if err != nil || rows == nil {
		return nil, err
       	}

       	rows.Next()
       	defer rows.Close()
       	return adapter.scanUser(rows)
}