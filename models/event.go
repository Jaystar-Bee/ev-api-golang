package models

import (
	"time"

	"test.com/event-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UserId      int64     `json:"user_id"`
}

// var events = []Event{}

func (e *Event) Save() error {
	create_query := `INSERT INTO events (name, description, location, date_time, created_at, user_id) VALUES (?,?,?,?,?,?)`

	statement, err := db.DB.Prepare(create_query)
	if err != nil {
		return err
	}
	defer statement.Close()
	time := time.Now()
	data, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, time, e.UserId)
	if err != nil {
		return err
	}
	id, err := data.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	e.CreatedAt = time
	return nil
}

func GetEvent(id int64) (*Event, error) {
	const query = `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	event := &Event{}
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.CreatedAt, &event.UserId)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func DeleteEvent(id int64) error {
	const query = `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents() ([]Event, error) {

	const query = `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.CreatedAt, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (event *Event) UpdateEvent() error {

	query := `
	Update events
	SET name = ?, description = ?, location = ?, date_time = ?, user_id = ?, created_at = ?
	WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.CreatedAt, event.ID)
	if err != nil {
		return err
	}
	return nil

}
