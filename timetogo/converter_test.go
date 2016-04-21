package timetogo

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToTime(t *testing.T) {
	t1 := Time{2016, 4, 21, 17, 30, 1}
	t2 := t1.ToTime()
	assert.Equal(t, 2016, t2.Year())
	assert.Equal(t, time.Month(4), t2.Month())
	assert.Equal(t, 21, t2.Day())
	assert.Equal(t, 17, t2.Hour())
	assert.Equal(t, 30, t2.Minute())
	assert.Equal(t, 1, t2.Second())
	assert.Equal(t, 0, t2.Nanosecond())
}

func TestTime(t *testing.T) {
	assert.Equal(t, time.Duration(25*int(math.Pow10(8))), times(float64(2.5), time.Second))
	assert.Equal(t, time.Duration(25*int(math.Pow10(8))), times(float32(2.5), time.Second))
	assert.Equal(t, time.Duration(2*int(math.Pow10(9))), times(2, time.Second))
	// case not handled yet
	assert.Panics(t, func() { times(uint(2), time.Second) })
}
func TestSecond(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Second, Seconds(2))
	// in nanoseconds
	assert.Equal(t, time.Duration(25*int(math.Pow10(8))), Seconds(2.5))
	assert.Equal(t, time.Duration(25*int(math.Pow10(8))), Seconds(float32(2.5)))
}
func TestMinute(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Minute, Minutes(2))
	assert.Equal(t, time.Duration(2500000000*60), Minutes(2.5))
}
func TestHour(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Hour, Hours(2))
	assert.Equal(t, time.Duration(2500000000*60*60), Hours(2.5))
}
func TestDay(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Hour*24, Days(2))
}
