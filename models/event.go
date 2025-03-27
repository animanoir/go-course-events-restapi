package models

import (
	"log"
	"time"
)

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	log.Println("Save function executed.")
	events = append(events, e)
}
