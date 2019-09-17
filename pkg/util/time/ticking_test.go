package time

import "testing"

func TestGetElapsedTime(t *testing.T) {

	var tc = &TickingClock{
		Initialized: false,
	}

	t.Run("Initialized changed to true after first call", func(t *testing.T) {
		tc.GetElapsedTime(1)
		if tc.Initialized != true {
			t.Errorf("ERROR: tc struct should be initialized to true after first call")
		}
	})

	t.Run("provided hour should be recorded in struct duration", func(t *testing.T) {
		tc.GetElapsedTime(1)
		if tc.duration != 1 {
			t.Errorf("ERROR: duration is expected to be equal with hour arg")
		}
	})

	t.Run("should return remaining time in the format of hh:mm:ss", func(t *testing.T) {
		got := tc.GetElapsedTime(1)
		want := "00:59:57"
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})

	t.Run("4 seconds has passed after 4 times call", func(t *testing.T) {
		got := tc.GetElapsedTime(1)
		want := "00:59:56"
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
}
