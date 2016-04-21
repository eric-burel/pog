package habit

import (
	"time"

	"github.com/eric-burel/pog/timetogo"
)

// Periodic An event that happens at a certain times
type Periodic struct {
	Months   []time.Month
	Weekdays []time.Weekday
	Days     []timetogo.Day
	Hours    []timetogo.Hour
	Minutes  []timetogo.Minute
	Seconds  []timetogo.Second
}
