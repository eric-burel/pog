package person

import "github.com/eric-burel/pog/habit"

var count = 0

// Person A person
type Person struct {
	ID     int           // unique identifier of the person
	Name   string        // Person name, for understanding purposes
	Habits []habit.Habit // habits of an user
}

// NewPerson Generate a new person
func NewPerson(name string) (p *Person) {
	p = &Person{count, name, make([]habit.Habit, 0)}
	count++
	return
}
