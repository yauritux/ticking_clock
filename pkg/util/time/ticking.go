package time

import (
	"time"
)

type TickingClock struct {
	duration    int
	end         time.Time
	Initialized bool
}

func (tc *TickingClock) GetElapsedTime(hour int) (elapsedTime string) {
	if tc.Initialized == false || tc.duration != hour {
		tc.duration = hour
		tc.end = time.Now().Local().Add(time.Hour * time.Duration(hour))
		tc.Initialized = true
	}

	return
}
