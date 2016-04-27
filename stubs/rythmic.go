package stubs

import (
	"time"

	"github.com/eric-burel/pog/event"
	"github.com/eric-burel/pog/habit"
	"github.com/eric-burel/pog/timetogo"
	"github.com/eric-burel/rgo/rand"
)

// user behaviour when on facebook
type lookAtFacebook struct {
	PagesLiked rand.Discreter // Number of pages liked
}

func (l lookAtFacebook) Generate() (e event.Event) {
	e = *event.New(nil, "visited facebook", event.Data{"liked pages": l.PagesLiked.R()})
	return
}

// time between two connexion on facebook
type facebookRate struct {
	Randomizer rand.Continuouser
}

func (r facebookRate) Duration() time.Duration {
	return timetogo.Minutes(r.Randomizer.R())
}

// NewFacebookRythmic Generate a new rythmic event that consist in looking at facebool
func NewFacebookRythmic() habit.Rythmic {
	l := lookAtFacebook{rand.NewBinom(100, 0.2)}
	// event that happens every half an hour, with a 5 minutes variance
	rate := facebookRate{rand.NewNorm(30, 5)}
	r := habit.Rythmic{habit.Habit{}, rate, l}
	return r
}
