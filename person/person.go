package person

import (
	"sort"
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/habit"
)

var count = 0

// Person A person
type Person struct {
	ID     int             // unique identifier of the person
	Name   string          // Person name, for understanding purposes
	Habits []habit.Habiter // habits of an user
}

// New Generate a new person
func New(name string) (p *Person) {
	p = &Person{count, name, nil}
	count++
	return
}

// AddHabit Add an habit to a person
func (p *Person) AddHabit(h habit.Habiter) {
	p.Habits = append(p.Habits, h)
}

// Act Generate event according to the person habit, during a time frame
func (p *Person) Act(begin, end time.Time) (evts event.Events) {
	for _, h := range p.Habits {
		habitEvts := h.Generate(begin, end)
		// add events
		evts = append(evts, habitEvts...)
	}
	// sort Events by trigger time
	sort.Sort(evts)
	return
}
