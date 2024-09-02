// Package api contains the implementation of the API server.
package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"repair-queue/service/appointment"
	"repair-queue/service/user"

	"github.com/gorilla/mux"
)

// Server represents the API server configuration and provides methods
// for running the server and handling HTTP requests.
type Server struct {
	addr string
	db   *sql.DB
}

// NewAPIServer creates and returns a new instance of Server, initialized with the provided address and database connection.
func NewAPIServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server and listens for incoming HTTP requests.
func (a *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	// User
	userStore := user.NewStore(a.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Appointment
	appointmentStore := appointment.NewStore(a.db)
	appointmentHandler := appointment.NewHandler(appointmentStore)
	appointmentHandler.RegisterRoutes(subrouter)

	fmt.Println("\033[32m✅ Listening on port", a.addr, "\033[0m")

	return http.ListenAndServe(a.addr, router)
}
