package models

import (
	"log"
	"time"
)

// Defines the shape of an Event
type Event struct {
	ID          int
	Name        string    `binding:"required` //`` = struct tags
	Description string    `binding:"required`
	Location    string    `binding:"required`
	DateTime    time.Time `binding:"required`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	log.Println("Save function executed.")
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
