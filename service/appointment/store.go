package appointment

import (
	"database/sql"
	"fmt"
	"repair-queue/types"
)

// Store provides an interface for interacting with the database and managing operations related to the "appointment" table.
type Store struct {
	db *sql.DB
}

// NewStore creates and returns a new instance of Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// CreateAppointment inserts a new user record into the database with the provided appointment details.
func (s *Store) CreateAppointment(appointment types.Appointment) error {
	_, err := s.db.Exec("INSERT INTO appointment (reason, model, make, licencePlate, manufactureYear, status, ownerPhoneNumber) VALUES (?,?,?,?,?,?,?)",
		appointment.Reason, appointment.Model, appointment.Make, appointment.LicencePlate, appointment.ManufactureYear, appointment.Status, appointment.OwnerPhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

// GetMinimizedAppointments retrieves appointments from the database in order by createdAt date.
func (s *Store) GetMinimizedAppointments() ([]*types.MinimizedAppointment, error) {
	rows, err := s.db.Query(`
        SELECT id, model, make, status, createdAt
        FROM appointment 
        ORDER BY createdAt ASC
    `)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var appointments []*types.MinimizedAppointment

	for rows.Next() {
		var appointment types.MinimizedAppointment
		if err := rows.Scan(
			&appointment.ID,
			&appointment.Model,
			&appointment.Make,
			&appointment.Status,
			&appointment.CreatedAt); err != nil {
			return nil, err
		}
		appointments = append(appointments, &appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(appointments) == 0 {
		return nil, fmt.Errorf("appointments not found")
	}

	return appointments, nil
}
