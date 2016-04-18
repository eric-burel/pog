package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerson(t *testing.T) {
	p := NewPerson("Bernard")
	assert.Equal(t, "Bernard", p.Name)
	assert.Equal(t, 0, p.ID)
}

func TestUniquePersonID(t *testing.T) {
	henri := NewPerson("Henri")
	didier := NewPerson("Didier")
	assert.NotEqual(t, henri.ID, didier.ID)
}
