package doomsday

import (
	"log"
	"testing"
)

// This is testing generated code, perhaps bizarre but I want to make sure it
// all works
func StringerTest(t *testing.T) {

	if res := Monday.String(); res != "Monday" {
		log.Println("Expected string to be Monday, got %s", res)
	}

	if res := Sunday.String(); res != "Sunday" {
		log.Println("Expected string to be Sunday, got %s", res)
	}
}

// Tests year anchor and century anchor (although it can be tested
// separately, it is just a rule based function
func TestYearAnchor(t *testing.T) {

	if res := yearAnchor(1966, centuryAnchor(1966)); res != Monday {
		t.Errorf("Expected day to be %d, got %d", Monday, res)
	}
}

// Performs some basic leapyear tests
func TestIsLeapYear(t *testing.T) {

	if !isLeapYear(2016) {
		t.Errorf("Expected 2016 to return true as leap year")
	}

	if !isLeapYear(2012) {
		t.Errorf("Expected 2012 to return true as leap year")
	}

	if isLeapYear(2015) {
		t.Errorf("Expected 2015 to return false as leap year")
	}

	if isLeapYear(2013) {
		t.Errorf("Expected 2013 to return false as leap year")
	}
}

// Checks if the last two digits are what we would expectto be removed
func TestLastTwoDigits(t *testing.T) {

	if res := lastTwoDigits(2015); res != 15 {
		t.Errorf("Expected last two digits of 2015 to be 15, got %d", res)
	}

	if res := lastTwoDigits(1994); res != 94 {
		t.Errorf("Expected last two digits of 1994 to be 94, got %d", res)
	}
}

// The built in modulus allows for negative
// range, that's not useful in this situation
// and so we create a small wrap around
func TestMod(t *testing.T) {

	if res := mod(-11, 7); res != 3 {
		t.Errorf("Expected mod to be 3, got %d", res)
	}

	if res := mod(11, 7); res != 4 {
		t.Errorf("Expected mod to be 4, got %d", res)
	}
}
