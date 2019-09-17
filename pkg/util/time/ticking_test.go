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
}
