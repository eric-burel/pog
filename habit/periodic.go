package habit

import (
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/timetogo"
	"github.com/eric-burel/rgo/rand"
)

// PeriodicDay A periodic habit is either defined as a weekly habit, or as a daily habit,
// hence this type
// NOTE : this behaviour might change, using only Day and using methods like generateDays(Weekday)
type PeriodicDay struct {
	IsWeekly bool
	Weekdays []time.Weekday
	Days     []timetogo.Day
}

// Periodic An event that happens at a certain times
// TODO : So far, it is not randomized, we should be able to use timetogo._er to generate Monthes, days and all
type Periodic struct {
	Months    []timetogo.Month //
	Days      PeriodicDay
	Hours     []timetogo.Hour
	Minutes   []timetogo.Minute
	Seconds   []timetogo.Second
	Happens   float64                              // Average probability that then event happens
	Condition func(p Periodic, e event.Event) bool // Cancel the event
	// useful to implement stuffs such ass "First tuesday of the month",
	// if event.Date does not match this condition, we cancel its firing
	Randomizer timetogo.Durationer // Add randomness
	Eventer    event.Eventer       // Generate a randomized event
}

// Generate Generate a set of periodic events
func (p Periodic) Generate(begin time.Time, end time.Time) (evts event.Events) {
	if !begin.Before(end) {
		panic("error : specified begin time was after end time")
	}
	t := begin
	evtTime := timetogo.NewZeroTime()
	y := begin.Year()
	// TODO I am a huge mess, refactor me
	for y <= end.Year() {
		evtTime.Year = timetogo.Year(y)
		for _, m := range p.Months {
			evtTime.Month = m
			var days []timetogo.Day
			if p.Days.IsWeekly {
				// TODO : generate all the days in the Month
				// for each weekday
				// generate them for this month and year
				days = make([]timetogo.Day, 0) // TODO generateDays(p.Days.Weekdays)
			} else {
				days = p.Days.Days
			}
			for _, d := range days {
				evtTime.Day = d
				for _, h := range p.Hours {
					evtTime.Hour = h
					for _, min := range p.Minutes {
						evtTime.Minute = min
						for _, s := range p.Seconds {
							evtTime.Second = s
							// skip if the generated time is before begin
							if evtTime.Before(begin) {
								continue
							} else if end.Before(evtTime.ToTime()) {
								// we break if event time is after the loop
								break
							}
							// Check if event should happen
							var happen = rand.NewBern(p.Happens).R()
							if happen == 0 {
								continue
							} else {
								// convert timetogo.Time to time.Time
								var stdTime = evtTime.ToTime()
								// add randomized duration to it
								stdTime = stdTime.Add(p.Randomizer.Duration())
								// generate an event
								var evt = p.Eventer.Generate()
								// set its time to the given m d h min sec
								evt.SetDate(stdTime)
								// add it to the slice
								evts = append(evts, evt)
								// increase time
								t = stdTime
							}
						}
					}
				}
			}
		}
		// if time haven't been increase, it means that p contained no Date
		// => we quit
		if t == begin {
			break
		}
		// when we iterated on every month, we increase year
		y++
	}
	return
}
