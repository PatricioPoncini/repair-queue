package types

import "time"

type UserStore interface {
	GetUserByUserName(userName string) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	UserName  string `json:"userName" validate:"required"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"userName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
