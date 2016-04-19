package habit

import (
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/timetogo"
	"github.com/eric-burel/rgo"
)

// Chaining An habit that consists in chaining a list of event, with a certain
// amount of time between each
type Chaining []chainingElement

// chainingElement Element of a chaining, ie an event
type chainingElement struct {
	Eventer    event.Eventer // The event generator
	TimeToNext rgo.Discreter // Time between this event and the next one
}

// AppendNew Create a new chaining element and adds it to the chaining
func (c Chaining) AppendNew(e event.Eventer, t rgo.Discreter) {
	elmt := chainingElement{e, t}
	c = append(c, elmt)
}

// Trigger Trigger the events of the chaining
func (c Chaining) Trigger(begin time.Time) {
	t := begin
	for _, elmt := range c {
		// generate the event random attributes
		event := elmt.Eventer.Generate()
		// trigger the event
		event.SetDate(t)
		nextDuration := timetogo.Seconds(elmt.TimeToNext.R()) // discrete random value is converted to a duration, in second
		// increase time between each event
		t.Add(nextDuration)
	}
}
