// Package db contains functions and types related to database operations.
package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// NewMySQLStorage creates a new instance of a MySQL database connection using the provided configuration.
func NewMySQLStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
