package models

import (
	"errors"

	"test.com/event-api/db"
)

type Registration struct {
	ID      int64 `json:"id"`
	EventId int64 `json:"event_id"`
	UserId  int64 `json:"user_id"`
}

func (event Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations (event_id, user_id) VALUES ($1, $2)
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = db.DB.Exec(query, event.ID, userId)
	return err
}

func (event Event) GetRegistrationByEvent(userId int64) (*Registration, error) {
	query := `
	SELECT * FROM registrations WHERE user_id = $1 AND event_id = $2
	`
	row := db.DB.QueryRow(query, userId, event.ID)

	userRegistration := &Registration{}
	err := row.Scan(&userRegistration.ID, &userRegistration.EventId, &userRegistration.UserId)
	if err != nil {
		return nil, err
	}
	return userRegistration, nil

}

func (event Event) Cancel(userId int64) error {

	query := `
		DELETE FROM registrations WHERE event_id = $1 AND user_id = $2
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("error preparing statement")
	}
	defer statement.Close()

	_, err = db.DB.Exec(query, event.ID, userId)
	if err != nil {
		return errors.New("error deleting registration")
	}
	return nil

}
