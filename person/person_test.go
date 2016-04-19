package person

import (
	"testing"

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
