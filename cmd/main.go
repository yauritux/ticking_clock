package main

import (
	"fmt"
	"os"
	"strconv"

	"janio.asia/ticking_clock/pkg/util/time"
)

func main() {
	args := os.Args[1:][0]

	hour, err := strconv.Atoi(args)

	if err != nil {
		panic(err.Error())
	}

	tc := &time.TickingClock{
		Initialized: true,
	}

	for {
		timeRemained := tc.GetElapsedTime(hour)
		if timeRemained == "00:00:00" {
			break
		}
		fmt.Printf("\r%s", timeRemained)
	}

	fmt.Printf("\r%s\n", "00:00:00")
}
