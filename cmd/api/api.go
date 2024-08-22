package api

import (
	"database/sql"
	"log"
	"net/http"

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
	//subrouter := router.PathPrefix("/api")

	log.Println("LIstening on: ", a.addr)

	return http.ListenAndServe(a.addr, router)
}
