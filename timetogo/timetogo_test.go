package timetogo

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSecond(t *testing.T) {
	assert.Equal(t, time.Duration(2)*time.Second, Seconds(2))
	// in nanoseconds
	assert.Equal(t, time.Duration(25*int(math.Pow10(8))), Seconds(2.5))
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
