package timetogo

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
