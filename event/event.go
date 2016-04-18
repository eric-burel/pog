package event

import (
	"time"
)

var count = 0

// Data Data that comes with an event
// Works exactly like a JSON object
// []interface{} would be a JSON array
type Data map[string]interface{}

// Event General type for an event
type Event struct {
	ID      int          // event unique Id
	Date    time.Time    // time of declenchment, allow to sort events
	Creator *interface{} // Object that triggered the event
	Data    Data         // Data that comes with the event, can be of any type
}

// Eventer An object able to generate an event
type Eventer interface {
	Generate() Event // Genereate an event
}

// NewEvent Constructor that sets up event date to Now
func NewEvent(creator *interface{}, data Data) (event *Event) {
	event = &Event{count, time.Now(), creator, data}
	count++
	return
}

// TriggerNow Trigger an event now
func (event *Event) TriggerNow() {
	event.Date = time.Now()
}

// Trigger Trigger an event at the specified date
func (event *Event) Trigger(date time.Time) {
	event.Date = date
}
