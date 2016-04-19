package timetogo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSecond(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Second, Seconds(2))
}
func TestMinute(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Minute, Minutes(2))
}
func TestHour(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Hour, Hours(2))
}
func TestDay(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Hour*24, Days(2))
}
