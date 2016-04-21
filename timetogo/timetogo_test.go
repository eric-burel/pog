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
