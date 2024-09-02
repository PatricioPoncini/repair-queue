// Package appointment defines the routes and handlers for appointment-related endpoints.
package appointment

import (
	"fmt"
	"net/http"
	"repair-queue/types"
	"repair-queue/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Handler manages HTTP requests related to user operations.
type Handler struct {
	store types.AppointmentStore
}

// NewHandler creates and returns a new instance of appointment Handler.
func NewHandler(store types.AppointmentStore) *Handler {
	return &Handler{
		store: store,
	}
}

// RegisterRoutes sets up the HTTP routes for the Handler
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/appointment", h.createAppointment).Methods("POST")
	router.HandleFunc("/appointment", h.getMinimizedAppointments).Methods("GET")
}

func (h *Handler) createAppointment(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateAppointmentPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid body payload %v", errors))
		return
	}

	err := h.store.CreateAppointment(types.Appointment{
		Reason:           payload.Reason,
		Model:            payload.Model,
		Make:             payload.Make,
		LicencePlate:     payload.LicencePlate,
		ManufactureYear:  payload.ManufactureYear,
		OwnerPhoneNumber: payload.OwnerPhoneNumber,
		Status:           types.StatusReceived,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, nil); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error generating response"))
	}
}

func (h *Handler) getMinimizedAppointments(w http.ResponseWriter, _ *http.Request) {
	appointments, err := h.store.GetMinimizedAppointments()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, appointments); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error generating response"))
	}
}
