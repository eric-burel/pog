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

// Month A month
type Month int

// Year A year
type Year int

// Monthday A day in a month, do not check if it exists (30 february)
type Monthday struct {
	Day   Day
	Month time.Month
}

// Time A strut representation of an instant, previse to the second
type Time struct {
	Year   Year
	Month  Month
	Day    Day
	Hour   Hour
	Minute Minute
	Second Second
}

// NewZeroTime Generate a new time, initialized to year 0, month 1, day 1, hour 0,
// minute 0, second 0
func NewZeroTime() (t *Time) {
	// first month and first day are 1, not 0
	t = &Time{0, 1, 1, 0, 0, 0}
	return
}

// Before Check whether a time is before another
func (t Time) Before(tt time.Time) bool {
	return t.ToTime().Before(tt)
}

// NewMonthday Generate a new Monthday
func NewMonthday(d Day, m time.Month) (md *Monthday) {
	md = &Monthday{d, m}
	return
}
