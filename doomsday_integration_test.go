//build +integration

package doomsday

import "testing"

// Runs some integration tests against the day of the week
func TestDayOfWeek(t *testing.T) {

	// Common day tests
	dow := DayOfWeek(2015, 11, 24)
	if dow != Tuesday {
		t.Errorf("Day of the week does not match expected: %d", dow)
	}

	dow = DayOfWeek(2015, 11, 25)
	if dow != Wednesday {
		t.Errorf("Day of the week does not match expected: %d", dow)
	}

	dow = DayOfWeek(1990, 3, 16)
	if dow != Friday {
		t.Errorf("Day of the week does not match expected: %d", dow)
	}

	dow = DayOfWeek(1837, 12, 25)
	if dow != Monday {
		t.Errorf("Day of the week does not match expected: %d", dow)
	}

	// Leap year tests
	dow = DayOfWeek(2000, 11, 10)
	if dow != Friday {
		t.Errorf("Day of the week does not match expected: %d", dow)
	}

	dow = DayOfWeek(2004, 01, 10)
	if dow != Saturday {
		t.Errorf("Day of the week does not match expected: %d", dow)
	}

	// On doomsday test
	dow = DayOfWeek(2005, 01, 3)
	if dow != Monday {
		t.Errorf("Day of the week does not match expected: %d", dow)
	}

}
