package event

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	event := NewEvent(nil, Data{"email": "foo@bar.com"})
	assert.NotNil(t, event)
	assert.Equal(t, "foo@bar.com", event.Data["email"])
}
func TestEventUniqueID(t *testing.T) {
	event1 := NewEvent(nil, Data{"email": "foo@bar.com"})
	event2 := NewEvent(nil, Data{"email": "foo@bar.com"})
	assert.NotEqual(t, event1.ID, event2.ID)
}

func TestTriggerNow(t *testing.T) {
	event := NewEvent(nil, Data{"email": "foo@bar.com"})
	before := time.Now()
	event.TriggerNow()
	after := time.Now()
	assert.True(t, event.Date.After(before), "Event trigger date %v is too early", event.Date)
	assert.True(t, event.Date.Before(after), "Event trigger date %v is too late", event.Date)
}

func TestTrigger(t *testing.T) {
	event := NewEvent(nil, Data{"email": "foo@bar.com"})
	now := time.Now()
	event.Trigger(now)
	assert.Equal(t, now, event.Date)
}
