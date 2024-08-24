// Package user defines the routes and handlers for user-related endpoints.
package user

import (
	"fmt"
	"net/http"
	"repair-queue/config"
	"repair-queue/service/auth"
	"repair-queue/types"
	"repair-queue/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Handler manages HTTP requests related to user operations.
type Handler struct {
	store types.UserStore
}

// NewHandler creates and returns a new instance of user Handler.
func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

// RegisterRoutes sets up the HTTP routes for the Handler
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid body payload %v", errors))
		return
	}

	u, err := h.store.GetUserByUserName(payload.UserName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid username or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid username or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error generating response"))
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid body payload %v", errors))
		return
	}

	_, err := h.store.GetUserByUserName(payload.UserName)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with this username already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		UserName:  payload.UserName,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, nil); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error generating response"))
	}
}
