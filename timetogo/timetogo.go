package timetogo

import "time"

// Second A second, the tiniest unit of timetogo
type Second int

// Minute  A minute, between 0-59
type Minute int

// Hour An hour, given by its position in the day, between 0-23
type Hour int

// Day A day, given by its position in the month, between 1-31
type Day int

// Monthday A day in a month, do not check if it exists (30 february)
type Monthday struct {
	Day   Day
	Month time.Month
}

// NewMonthday Generate a new Monthday
func NewMonthday(d Day, m time.Month) (md *Monthday) {
	md = &Monthday{d, m}
	return
}
