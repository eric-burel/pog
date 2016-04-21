package habit

import (
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/timetogo"
	"github.com/eric-burel/rgo"
)

// Periodic An event that happens at a certain times
// TODO : So far, it is not randomized, we should be able to use timetogo._er to generate Monthes, days and all
type Periodic struct {
	Months []timetogo.Month //
	Days   struct {
		isWeekly bool
		Weekdays []timetogo.Weekday
		Days     []timetogo.Day
	}
	Hours     []timetogo.Hour
	Minutes   []timetogo.Minute
	Seconds   []timetogo.Second
	Happens   float64                        // Average probability that then event happens
	Condition func(p Periodic, e Event) bool // Cancel the event
	// useful to implement stuffs such ass "First tuesday of the month",
	// if event.Date does not match this condition, we cancel its firing
	Randomizer timetogo.Durationer // Add randomness
	Eventer    event.Eventer       // Generate a randomized event
}

// Generate Generate a set of periodic events
func (p Periodic) Generate(begin time.Time, end time.Time) (evts []event.Event) {
	t := begin
	y := t.Year()
	// TODO I am a huge mess, refactor me
	for t.Before(end) {
		for _, m := range p.Months {
			// skip if current time is before begin
			if y == begin.Year() && m < begin.Month() {
				continue
			}
			var days []Day
			if p.Days.isWeekly {
				// TODO : generate all the days in the Month
				// for each weekday
				// generate them for this month and year
				days = make([]timetogo.Weekday, 0) // TODO generateDays(p.Days.Weekdays)
			} else {
				days = p.Days.Days
			}
			for _, d := range days {
				if y == begin.Year() && m == begin.Month() && d < begin.Day() {
					continue
				}
				// TODO : skip if year y, month m, day d are  < to t
				for _, h := range p.Hours {
					if y == begin.Year() && m == begin.Month() && d < begin.Hour() {
						continue
					}
					// TODO : skip if year y, month m, day d, hour h are  < to t
					for _, min := range p.Minutes {
						if y == begin.Year() && m == begin.Month() && h == begin.Hour() && d < begin.Minute() {
							continue
						}
						// TODO : skip if year y, month m, day d, hour h,minute min are  < to t
						for _, s := range p.Seconds {
							if y == begin.Year() && m == begin.Month() && h == begin.Hour() && d == begin.Minute() && s < begin.Second() {
								continue
							}
							// Check if event should happen
							var happen = rgo.NewBern(p.Happens).R()
							if happen == 0 {
								// TODO break for loop
							} else {
								// TODO : skip if year y, month m, day d, hour h,minute min, second s are  < to t
								var evtTime = time.Date(y, m, d, h, min, s, 0, time.UTC)
								// TODO check if this time respects p.condition
								// generate an event
								var evt = p.Eventer.Generate()
								// set its time to the given m d h min sec
								evt.SetDate(evtTime)
								// add it to the slice
								append(evts, evt)
								// increase time
								t = evtTime
							}
						}
					}
				}
			}
		}
		// when we iterated on every month, we increase year
		y++
	}
	return
}
