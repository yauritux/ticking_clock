package time

import (
	"time"
)

type TickingClock struct {
	duration    int
	end         time.Time
	Initialized bool
}

func (tc TickingClock) GetElapsedTime(hour int) (elapsedTime string) {
	return
}
