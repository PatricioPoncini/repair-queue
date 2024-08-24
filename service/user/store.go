package user

import (
	"database/sql"
	"fmt"
	"repair-queue/types"
)

// Store provides an interface for interacting with the database and managing operations related to the "users" table.
type Store struct {
	db *sql.DB
}

// NewStore creates and returns a new instance of Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetUserByUserName retrieves a user from the database based on the provided username.
func (s *Store) GetUserByUserName(userName string) (*types.User, error) {
	rows, err := s.db.Query("SELECT id, firstName, lastName, userName, password, createdAt FROM users WHERE userName = ?", userName)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// CreateUser inserts a new user record into the database with the provided user details.
func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, userName, password) VALUES (?,?,?,?)", user.FirstName, user.LastName, user.UserName, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
