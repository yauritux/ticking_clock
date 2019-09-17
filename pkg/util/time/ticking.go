package time

import (
	"fmt"
	"time"
)

type TickingClock struct {
	duration    int
	end         time.Time
	Initialized bool
}

func (tc *TickingClock) GetElapsedTime(hour int) (timeRemained string) {
	if tc.Initialized == false || tc.duration != hour {
		tc.duration = hour
		tc.end = time.Now().Local().Add(time.Hour * time.Duration(hour))
		tc.Initialized = true
	}

	remaining := tc.end.Sub(time.Now().Local())
	time.Sleep(1000 * time.Millisecond)
	timeRemained = fmt.Sprintf(
		"%02d:%02d:%02d",
		int(remaining.Hours()),
		int(remaining.Minutes())%60,
		int(remaining.Seconds())%60,
	)

	return
}
