package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"repair-queue/types"
	"testing"

	"github.com/gorilla/mux"
)

type mockUserStore struct{}

func (m *mockUserStore) GetUserByUserName(_ string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the body payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "test_firstName",
			LastName:  "test_lastName",
			UserName:  "",
			Password:  "ab",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register the yuseuser", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "test_firstName",
			LastName:  "test_lastName",
			UserName:  "test_username",
			Password:  "test_password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}
