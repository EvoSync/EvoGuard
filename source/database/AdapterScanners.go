package database

import "database/sql"

// scanUser will scan directly from the database through the rows interface
func (adapter *Adapter) scanUser(rows *sql.Rows) (*User, error) {
	destination := new(User)

	if err := rows.Scan(&destination.ID, &destination.Username, &destination.Password, &destination.Email, &destination.Salt, &destination.AccountLevel, &destination.Parent); err != nil {
		return nil, err
	}
	
	return destination, nil
}