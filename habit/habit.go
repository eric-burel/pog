package habit

import (
	"time"

	"github.com/eric-burel/pog/event"
)

// Habit An habit is a set of one or more event that happens in a pseudo-deterministic
// way
// For example "leaving work" always occur around 17pm, (Timed), every day (Periodical)
// and is followed by "taking car" (chaining)
// An habit can cumulate many types
type Habit struct {
}

// Habiter An object able to generate a set of event
type Habiter interface {
	Trigger(begin time.Time) []event.Event // function that generate a set of events
}
