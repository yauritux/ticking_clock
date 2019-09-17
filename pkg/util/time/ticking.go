package time

import (
	"fmt"
	"time"
)

type TickingClock struct {
	duration int
	end time.Time
	Initialized bool\
}