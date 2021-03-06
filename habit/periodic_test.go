package habit

import (
	"testing"
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/timetogo"
	"github.com/eric-burel/rgo/rand"
	"github.com/stretchr/testify/assert"
)

type exam struct{}

func (exam) Generate() event.Event {
	students := rand.NewBinom(100, 0.98).R()
	absents := 100 - students
	duration := rand.NewGeom(0.5).R() + 1
	data := event.Data{"presents": students, "absents": absents, "duration": duration}
	e := event.New(nil, "exam", data)
	return *e
}

type examTime struct{}

// Duration The begining of the exam, either at 0,15,30,45
func (examTime) Duration() (d time.Duration) {
	// 0 is the more likely, then 30, 15 and 45 seldom happens
	possibilities := map[int]float64{0: 0.5, 15: 0.1, 30: 0.3, 45: 0.1}
	startAt := rand.NewFiniteInt(possibilities).R()
	d = timetogo.Minutes(startAt)
	return
}

// Test when no time is specified
func TestPeriodicGenerateNeverHappens(t *testing.T) {
	m := []timetogo.Month{1, 5}
	d := []timetogo.Day{3, 4, 5, 6, 7}
	//h := []timetogo.Hour{8, 13}
	min := []timetogo.Minute{0}
	s := []timetogo.Second{0}
	happens := 0.99
	randomizer := new(examTime)
	eventer := new(exam)

	p := Periodic{m, PeriodicDay{false, nil, d}, nil, min, s, happens, nil, randomizer, eventer}

	begin := time.Now()
	end := begin.Add(timetogo.Days(365 * 2))
	evts := p.Generate(begin, end)
	for _, evt := range evts {
		t.Logf("Event : exam started at %v and lasted %v hours, with %v students there, and %v absents.\n",
			evt.Date, evt.Data["duration"], evt.Data["presents"], evt.Data["absents"])
	}

}
func TestPeriodicGenerate(t *testing.T) {
	m := []timetogo.Month{1, 5}
	d := []timetogo.Day{3, 4, 5, 6, 7}
	h := []timetogo.Hour{8, 13}
	min := []timetogo.Minute{0}
	s := []timetogo.Second{0}
	happens := 0.99
	randomizer := new(examTime)
	eventer := new(exam)

	p := Periodic{m, PeriodicDay{false, nil, d}, h, min, s, happens, nil, randomizer, eventer}

	begin := time.Now()
	end := begin.Add(timetogo.Days(365 * 2))
	evts := p.Generate(begin, end)
	for _, evt := range evts {
		t.Logf("Event : exam started at %v and lasted %v hours, with %v students there, and %v absents.\n",
			evt.Date, evt.Data["duration"], evt.Data["presents"], evt.Data["absents"])
	}
}

func TestPeriodicLength(t *testing.T) {
	m := []timetogo.Month{1, 5}
	d := []timetogo.Day{3, 4, 5, 6, 7}
	h := []timetogo.Hour{8, 13}
	min := []timetogo.Minute{0}
	s := []timetogo.Second{0}
	happens := 0.99
	randomizer := new(examTime)
	eventer := new(exam)

	p := Periodic{m, PeriodicDay{false, nil, d}, h, min, s, happens, nil, randomizer, eventer}
	begin := time.Date(2016, 1, 3, 5, 30, 0, 0, time.UTC)
	end := time.Date(2016, 1, 3, 23, 30, 0, 0, time.UTC)
	evts := p.Generate(begin, end)
	assert.Equal(t, 2, evts.Len())

}

func TestPanicBeforeEnd(t *testing.T) {
	m := []timetogo.Month{1, 5}
	d := []timetogo.Day{3, 4, 5, 6, 7}
	h := []timetogo.Hour{8, 13}
	min := []timetogo.Minute{0}
	s := []timetogo.Second{0}
	happens := 0.99
	randomizer := new(examTime)
	eventer := new(exam)

	p := Periodic{m, PeriodicDay{false, nil, d}, h, min, s, happens, nil, randomizer, eventer}
	begin := time.Date(2016, 1, 3, 7, 30, 0, 0, time.UTC)
	end := time.Date(2016, 1, 2, 7, 30, 0, 0, time.UTC)
	assert.Panics(t, func() { p.Generate(begin, end) })

}
