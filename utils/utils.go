// Package utils provides utility functions and helpers used across the application.
package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Validate is a global instance of the validator.Validation struct.
var Validate = validator.New()

// ParseJSON reads and parses the JSON payload from an HTTP request body
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON writes a JSON response to the HTTP response writer.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// WriteError writes an error response to the HTTP response writer.
func WriteError(w http.ResponseWriter, status int, err error) {
	if e := WriteJSON(w, status, map[string]string{"error": err.Error()}); e != nil {
		log.Printf("error writing JSON response: %v", e)
	}
}
