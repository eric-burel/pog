package timetogo

import "time"

// Durationer An object that can be converted to a duration
type Durationer interface {
	Duration() time.Duration
}

// Weekdayer Generate a week day, between Monday-Sunday
type Weekdayer interface {
	Weekday() time.Weekday
}

// Dayer Generate a day, between 0-31
type Dayer interface {
	Day() Day
}

// Hourer Generate an hour, between 0-23
type Hourer interface {
	Hour() Hour
}

// Minuter Generate minutes, between 0-59
type Minuter interface {
	Minuter() Minute
}

// Seconder Generate seconds, berween 0-59
type Seconder interface {
	Seconder() Second
}
