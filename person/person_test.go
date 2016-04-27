package person

import (
	"testing"
	"time"

	"github.com/eric-burel/pog/stubs"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := New("Bernard")
	assert.Equal(t, "Bernard", p.Name)
	assert.Equal(t, 0, p.ID)
}

func TestUniquePersonID(t *testing.T) {
	henri := New("Henri")
	didier := New("Didier")
	assert.NotEqual(t, henri.ID, didier.ID)
}

func TestAddPeriodHabit(t *testing.T) {
	p := New("Bernard")
	h := stubs.NewExamPeriodic()
	p.AddHabit(h)
	assert.Equal(t, h, p.Habits[0])

}
func TestAddRythmicHabit(t *testing.T) {
	p := New("Bernard")
	h := stubs.NewFacebookRythmic()
	p.AddHabit(h)
	assert.Equal(t, h, p.Habits[0])
}

func TestAct(t *testing.T) {
	p := New("Bernard")
	facebook := stubs.NewFacebookRythmic()
	exam := stubs.NewExamPeriodic()
	p.AddHabit(facebook)
	p.AddHabit(exam)
	begin := time.Date(2016, 1, 3, 7, 30, 0, 0, time.UTC)
	end := time.Date(2016, 1, 3, 22, 30, 0, 0, time.UTC)
	evts := p.Act(begin, end)
	assert.True(t, evts.Len() > 0)
	for _, evt := range evts {
		t.Logf("******************************\n")
		t.Logf("Event: %v\n", evt.Family)
		t.Logf("Triggered at: %v\n", evt.Date)
		t.Logf("Data: %v\n", evt.Data)
		t.Logf("******************************\n\n")
	}
}
