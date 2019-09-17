package main

import (
	"fmt"
	"janio.asia/ticking_clock/pkg/util/time"
)

func main() {
	tc := &time.TickingClock{
		Initialized: true,
	}

	for {
		timeRemained := tc.GetElapsedTime(1)
		if timeRemained == "00:00:00" {
			break
		}
		fmt.Printf("\r%s", timeRemained)
	}

	fmt.Printf("\r%s\n", "00:00:00")
}

