package habit

import (
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/timetogo"
)

// Rythmic An habit that consist in triggering the same event at a given rate :
// connecting every day, cleaning mail every month and so one
type Rythmic struct {
	Habit
	Rate    timetogo.Durationer // Generate time between two occurence of the event
	Eventer event.Eventer       // Generate the event
}

// Generate Generate a set of periodic events
func (p Rythmic) Generate(begin time.Time, end time.Time) (evts []event.Event) {
	t := begin
	for t.Before(end) {
		evt := p.Eventer.Generate()
		evt.SetDate(t)
		evts = append(evts, evt)
		nextDuration := p.Rate.Duration()
		t = t.Add(nextDuration)
	}
	return
}
