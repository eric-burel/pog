package timetogo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMonthday(t *testing.T) {
	md := NewMonthday(21, time.January)
	assert.Equal(t, Day(21), md.Day)
	assert.Equal(t, time.Month(1), md.Month)
}

func TestNewZeroTime(t *testing.T) {
	time := NewZeroTime()
	assert.Equal(t, Year(0), time.Year)
	assert.Equal(t, Month(1), time.Month)
	assert.Equal(t, Day(1), time.Day)
	assert.Equal(t, Hour(0), time.Hour)
	assert.Equal(t, Minute(0), time.Minute)
	assert.Equal(t, Second(0), time.Second)
}

func TestBefore(t *testing.T) {
	t1 := Time{2016, 4, 21, 17, 30, 2}
	before := time.Date(2016, 4, 21, 17, 30, 1, 0, time.UTC)
	after := time.Date(2016, 4, 21, 17, 30, 3, 0, time.UTC)
	assert.False(t, t1.Before(before))
	assert.True(t, t1.Before(after))
}
