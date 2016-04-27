package event

import "time"

var count = 0

// Data Data that comes with an event
// Works exactly like a JSON object
// []interface{} would be a JSON array
type Data map[string]interface{}

// Event General type for an event
type Event struct {
	ID      int          // event unique Id
	Date    time.Time    // time of declenchment, allow to sort events
	Family  string       // event family
	Creator *interface{} // Object that triggered the event
	Data    Data         // Data that comes with the event, can be of any type
}

// Events An array of events
type Events []Event

// Eventer An object able to generate an event
type Eventer interface {
	Generate() Event // Genereate an event
}

// New Constructor that sets up event date to Now
func New(creator *interface{}, family string, data Data) (event *Event) {
	event = &Event{count, time.Now(), family, creator, data}
	count++
	return
}

// SetDateNow Trigger an event now
func (event *Event) SetDateNow() {
	event.Date = time.Now()
}

// SetDate Trigger an event at the specified date
func (event *Event) SetDate(date time.Time) {
	event.Date = date
}

// Len Number of events
func (events Events) Len() int {
	return len(events)
}

// Less return true if event i appends before event j
func (events Events) Less(i, j int) bool {
	return events[i].Date.Before(events[j].Date)
}

// Swap Swap two elements of the array
func (events Events) Swap(i, j int) {
	events[i], events[j] = events[j], events[i] //idiomatic
	return
}
