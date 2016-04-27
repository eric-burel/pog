package event

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	event := New(nil, "e", Data{"email": "foo@bar.com"})
	assert.NotNil(t, event)
	assert.Equal(t, "foo@bar.com", event.Data["email"])
}
func TestEventUniqueID(t *testing.T) {
	event1 := New(nil, "e", Data{"email": "foo@bar.com"})
	event2 := New(nil, "e", Data{"email": "foo@bar.com"})
	assert.NotEqual(t, event1.ID, event2.ID)
}

func TestTriggerNow(t *testing.T) {
	event := New(nil, "e", Data{"email": "foo@bar.com"})
	before := time.Now()
	event.SetDateNow()
	after := time.Now()
	assert.True(t, event.Date.After(before), "Event trigger date %v is too early", event.Date)
	assert.True(t, event.Date.Before(after), "Event trigger date %v is too late", event.Date)
}

func TestTrigger(t *testing.T) {
	event := New(nil, "e", Data{"email": "foo@bar.com"})
	now := time.Now()
	event.SetDate(now)
	assert.Equal(t, now, event.Date)
}

func TestLess(t *testing.T) {
	e1 := New(nil, "e", Data{"email": "foo@bar.com"})
	e1.SetDateNow()
	e2 := New(nil, "e", Data{"email": "foo@bar.com"})
	e2.SetDateNow()
	evts := Events{*e2, *e1}
	assert.True(t, evts.Less(1, 0))
	assert.False(t, evts.Less(0, 1))
}
func TestSwap(t *testing.T) {
	e1 := New(nil, "e", Data{"email": "foo@bar.com"})
	e1.SetDateNow()
	e2 := New(nil, "e", Data{"email": "foo@bar.com"})
	e2.SetDateNow()
	evts := Events{*e1, *e2}
	evts.Swap(0, 1)
	// evtnes tshould be swapped
	assert.Equal(t, e1.ID, evts[1].ID)
	assert.Equal(t, e2.ID, evts[0].ID)
}

// TestSort An event array should be sortable by trigger time
func TestSort(t *testing.T) {
	e1 := New(nil, "e", Data{"email": "foo@bar.com"})
	e1.SetDateNow()
	e2 := New(nil, "e", Data{"email": "foo@bar.com"})
	e2.SetDateNow()
	evts := Events{*e2, *e1}
	sort.Sort(evts)
	// evtnes tshould be swapped
	assert.Equal(t, e1.ID, evts[0].ID)
	assert.Equal(t, e2.ID, evts[1].ID)
}
