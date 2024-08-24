// Package types defines the various types and data structures used throughout the application.
package types

import "time"

// UserStore defines the methods for interacting with user data.
type UserStore interface {
	GetUserByUserName(userName string) (*User, error)
	CreateUser(User) error
}

// LoginUserPayload represents the payload structure for a user login request.
type LoginUserPayload struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterUserPayload represents the payload structure for a user register request.
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	UserName  string `json:"userName" validate:"required"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}

// User represents a user in the system.
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"userName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
