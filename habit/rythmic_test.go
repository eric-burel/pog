package habit

import (
	"testing"
	"time"

	"github.com/eric-burel/pog/event"
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

func TestGenerate(t *testing.T) {
	// TODO So far, it simply checks that functions work well, but does
	// validate results
	l := lookAtFacebook{rand.NewBinom(100, 0.2)}
	// event that happens every half an hour, with a 5 minutes variance
	rate := facebookRate{rand.NewNorm(30, 5)}
	p := Rythmic{Habit{}, rate, l}
	begin := time.Now()
	end := begin.Add(timetogo.Hours(5))
	evts := p.Generate(begin, end)
	for _, evt := range evts {
		t.Logf("Event : %v pages liked at time %v \n", evt.Data["liked pages"], evt.Date)
	}
}
