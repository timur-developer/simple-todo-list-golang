package eventsdata

type EventData struct {
	events []Event
}

func CreateEventData() EventData {
	ed := EventData{}
	return ed
}

func (ed *EventData) AddEvent(event Event) {
	ed.events = append(ed.events, event)
}

func (ed *EventData) GetAllEvents() []Event {
	return ed.events
}
