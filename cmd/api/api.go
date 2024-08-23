package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"repair-queue/service/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (a *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	fmt.Println("\033[32m✅ Listening on port", a.addr, "\033[0m")

	return http.ListenAndServe(a.addr, router)
}
