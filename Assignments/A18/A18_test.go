package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	gotOutput1 := isLeapYear(2012)
	// expectedOutput := true
	if !gotOutput1 {
		t.Error("Invalid Output. Expected: true, Got:", gotOutput1)	
	}

	gotOutput2 := isLeapYear(2013)
	if gotOutput2 {
		t.Error("Invalid Output. Expected: False, Got:", gotOutput2)	
	}
}