// The doomsday package implements a somewhat human approach to calculating
// Conway's Doomsday Algorithm. It exposes a core function:
//
// DayOfWeek(year, month, day) int
//
// It returns a Day (base-type integer)
//
// It exposes two types, Day and Month, which represent the underlying dates as
// integers from 0-7 and 1-12 respectively. These have a stringer implemented
// and so can be converted to their string representative
//
// Finally it can say if a given year is a leap year: IsLeapYear(year) bool
package doomsday

import "math"

// The day of the week is an integer than can be operated on, ranging 0 - 6. It
// may be converted into it's string representative
type Day int

// The days of the week (0 - 6)
const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Month int

// The months of the year (1 - 12)
const (
	_ Month = iota
	January
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// Day runs the doomsday algorithm to find the day of the week.
// It takes a 'human' approach, so tries to work things out
// descriptively as a human might, rather than being more elegant
func DayOfWeek(year, month, day int) Day {

	// Find the anchor for the year
	yearAnchor := yearAnchor(year, centuryAnchor(year))

	// Calculate the closest day
	_, closestDay := closestDoomsday(month, IsLeapYear(year))

	// Distance from the closest day anchor
	dist := (day - closestDay)

	// Return the distance from anchor mod week
	return Day(mod((int(yearAnchor) + dist), 7))
}

// Returns true if it is a leap year, otherwise false
func IsLeapYear(year int) bool {
	return (year % 4) == 0
}

// Based on a given month, will return the naively closest month and
// day, again this is based on how a human might do it
func closestDoomsday(month int, isLeap bool) (int, int) {

	switch Month(month) {
	case January:
		if isLeap {
			return int(January), 4
		} else {
			return int(January), 3
		}
	case February:
		if isLeap {
			return int(February), 29
		} else {
			return int(February), 28
		}
	case March:
		return int(February), 28
	case April:
		return int(April), 4
	case May:
		return int(May), 9
	case June:
		return int(June), 6
	case July:
		return int(July), 11
	case August:
		return int(August), 8
	case September:
		return int(September), 5
	case October:
		return int(October), 10
	case November:
		return int(November), 7
	case December:
		return int(December), 12
	}

	return 0, 0
}

// Grabs the anchor for a given year
func yearAnchor(year int, centuryAnchor Day) Day {

	lastDigits := lastTwoDigits(year)

	// Divide the last two digits of the year by 12
	y := float64(lastDigits) / 12.0

	// Let a be the floor of the quotient
	a := int(math.Floor(float64(y)))

	// Let b be the remainder of the quotient
	b := lastDigits % 12

	// Divide the remainder by 4 and let c be the floor of the quotient
	c := int(math.Floor(float64(b) / 4.0))

	// Let d be the sum of the three numbers
	d := a + b + c

	// Count forward the specified number of days from the anchor day
	doomsday := (d%7 + int(centuryAnchor)) % 7

	return Day(doomsday)
}

// Gets the anchor day for a given year, calculated how a human might
// calculate it
func centuryAnchor(year int) Day {

	if year >= 1800 && year <= 1899 {
		return Friday
	}

	if year >= 1900 && year <= 1999 {
		return Wednesday
	}

	if year >= 2000 && year <= 2099 {
		return Tuesday
	}

	if year >= 2100 && year <= 2199 {
		return Sunday
	}

	return -1
}

// Grabs the last two digits of a year
func lastTwoDigits(year int) int {

	// Last two digits of the year are a maximum of 99
	hundreds := year

	// Could optimize by removing 1000s first
	for ; hundreds >= 99; hundreds /= 100 {
	}

	// Cut off the hundreds to leave the last digits
	lastDigits := year - (hundreds * 100)

	return lastDigits
}

// Positive truncating modulus
func mod(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}

//go:generate stringer -type=Day
//go:generate stringer -type=Month
