// Package types defines the various types and data structures used throughout the application.
package types

import "time"

// UserStore defines the methods for interacting with user data.
type UserStore interface {
	GetUserByUserName(userName string) (*User, error)
	CreateUser(User) error
}

// AppointmentStore defines the methods for interacting with appointments data.
type AppointmentStore interface {
	CreateAppointment(Appointment) error
	GetMinimizedAppointments() ([]*MinimizedAppointment, error)
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

// Appointment represents a appointment in the system.
type Appointment struct {
	ID               int       `json:"id"`
	Reason           string    `json:"reason"`
	Model            string    `json:"model"`
	Make             string    `json:"make"`
	LicencePlate     string    `json:"licencePlate"`
	ManufactureYear  int       `json:"manufactureYear"`
	Status           string    `json:"status"`
	OwnerPhoneNumber string    `json:"ownerPhoneNumber"`
	CreatedAt        time.Time `json:"createdAt"`
}

// CreateAppointmentPayload represents the payload structure for a create appointment request.
type CreateAppointmentPayload struct {
	Reason           string `json:"reason" validate:"required"`
	Model            string `json:"model" validate:"required"`
	Make             string `json:"make" validate:"required"`
	LicencePlate     string `json:"licencePlate" validate:"required"`
	ManufactureYear  int    `json:"manufactureYear" validate:"required"`
	OwnerPhoneNumber string `json:"ownerPhoneNumber" validate:"required"`
}

// MinimizedAppointment represents the structure for minimized appointments.
type MinimizedAppointment struct {
	ID        int       `json:"id"`
	Model     string    `json:"model"`
	Make      string    `json:"make"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// StatusReceived Types of statuses for appointments.
const (
	StatusReceived = "RECEIVED"
	// STATUS_DIAGNOSIS         = "DIAGNOSIS"
	// STATUS_REPAIR_IN_PROCESS = "REPAIR_IN_PROCESS"
	// STATUS_READY_FOR_PICKUP  = "READY_FOR_PICKUP"
)
