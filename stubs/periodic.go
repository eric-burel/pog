/**
 * A periodic event example
 */
package stubs

import (
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/habit"
	"github.com/eric-burel/pog/timetogo"
	"github.com/eric-burel/rgo/rand"
)

type Exam struct{}

func (Exam) Generate() event.Event {
	students := rand.NewBinom(100, 0.98).R()
	absents := 100 - students
	duration := rand.NewGeom(0.5).R() + 1
	data := event.Data{"presents": students, "absents": absents, "duration": duration}
	e := event.New(nil, "exam", data)
	return *e
}

type ExamTime struct{}

// Duration The begining of the exam, either at 0,15,30,45
func (e ExamTime) Duration() (d time.Duration) {
	// 0 is the more likely, then 30, 15 and 45 seldom happens
	possibilities := map[int]float64{0: 0.5, 15: 0.1, 30: 0.3, 45: 0.1}
	startAt := rand.NewFiniteInt(possibilities).R()
	d = timetogo.Minutes(startAt)
	return
}

// NewExamPeriodic Generate a new habit that consists in generating exams
func NewExamPeriodic() habit.Periodic {
	m := []timetogo.Month{1, 5}
	d := []timetogo.Day{3, 4, 5, 6, 7}
	h := []timetogo.Hour{8, 13}
	min := []timetogo.Minute{0}
	s := []timetogo.Second{0}
	happens := 0.99
	randomizer := new(ExamTime)
	eventer := new(Exam)
	p := habit.Periodic{Months: m,
		Days:       habit.PeriodicDay{IsWeekly: false, Weekdays: nil, Days: d},
		Hours:      h,
		Minutes:    min,
		Seconds:    s,
		Happens:    happens,
		Condition:  nil,
		Randomizer: randomizer,
		Eventer:    eventer}
	return p
}
