package eventsdata

import "time"

type Event struct {
	input        string
	description  string
	CreationTime time.Time
}

func CreateEvent() Event {
	event := Event{}
	return event
}

func (e *Event) SetEventInput(input string) {
	e.input = input
}

func (e *Event) SetEventDescription(description string) {
	e.description = description
}

func (e *Event) SetEventCreationTime(creationTime time.Time) {
	e.CreationTime = creationTime
}
