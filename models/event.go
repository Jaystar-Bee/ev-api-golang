package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date" binding:"required"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() {
	// Save an event to the database
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}
